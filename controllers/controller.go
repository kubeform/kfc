/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform/command"

	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/dynamic"

	"k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	terraformv1alpha1 "github.com/appscode-cloud/kfc/api/v1alpha1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const KFCFinalizer = "kfc.io"

var (
	homePath = os.Getenv("HOME")
	basePath = filepath.Join(homePath, ".kfc")
)

// ResourceReconciler reconciles a Resource object
type ResourceReconciler struct {
	client.Client
	Log        logr.Logger
	DynClient  dynamic.Interface
	Kubeclient kubernetes.Interface
	CrdClient  *crdclientset.Clientset
}

// +kubebuilder:rbac:groups=terraform.kfc.io,resources=resources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=terraform.kfc.io,resources=resources/status,verbs=get;update;patch

func (r *ResourceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("kfc-controller", req.NamespacedName)
	log.Info("Reconciling Resource")

	var resource terraformv1alpha1.Resource
	if err := r.Get(ctx, req.NamespacedName, &resource); err != nil {
		log.Error(err, "unable to fetch resource")
	}

	obj, err := r.DynClient.Resource(resource.Spec.GVR).Namespace(resource.Spec.Namespace).Get(resource.Spec.Name, metav1.GetOptions{})
	if err != nil {
		log.Error(err, "unable to get resource", "ns", resource.Spec.Namespace, "name", resource.Spec.Name, "gvr", resource.Spec.GVR)
		return ctrl.Result{}, nil
	}

	// TODO: make a namer
	resPath := filepath.Join(basePath, resource.Spec.GVR.Resource+"."+resource.Spec.Namespace+"."+resource.Spec.Name)
	providerFile := filepath.Join(resPath, "provider.tf.json")
	mainFile := filepath.Join(resPath, "main.tf.json")
	stateFile := filepath.Join(resPath, "terraform.tfstate")

	if hasFinalizer(obj.GetFinalizers(), KFCFinalizer) {
		if obj.GetDeletionTimestamp() != nil {
			err := terraformDestroy(resPath, stateFile)
			if err != nil {
				log.Error(err, "failed to terraform destroy")
			}

			err = deleteFiles(resPath)
			if err != nil {
				log.Error(err, "failed to delete files")
			}

			err = r.Client.Delete(ctx, &resource)
			if err != nil {
				log.Error(err, "failed to delete resource")
			}

			err = removeFinalizer(obj)
			if err != nil {
				log.Error(err, "failed to remove finalizer")
			}

			r.updateResource(resource.Spec.GVR, obj)
			return ctrl.Result{}, nil
		}
	} else {
		err := addFinalizer(obj, KFCFinalizer)
		if err != nil {
			log.Error(err, "failed to add finalizer")
		}
	}

	err = createFiles(resPath, stateFile, providerFile, mainFile)
	if err != nil {
		log.Error(err, "failed to create files")
		return ctrl.Result{}, nil
	}

	// TODO: how to handle provider name?
	configMap, err := r.Kubeclient.CoreV1().ConfigMaps("default").Get(strings.Split(kindToResouceMap[resource.Spec.GVR.Resource], "_")[0], metav1.GetOptions{})
	if err != nil {
		log.Error(err, "unable to fetch configmap")
	}

	err = configmapToTFProvider(configMap, providerFile)
	if err != nil {
		log.Error(err, "unable to get configmap")
	}

	err = crdToTFResource(resource.Spec.GVR.Resource, obj, mainFile)
	if err != nil {
		log.Error(err, "unable to get crd resource")
	}

	init, _, err := unstructured.NestedBool(obj.Object, "status", "initialized")
	if !init {
		err = terraformInit(resPath)
		if err != nil {
			log.Error(err, "unable to initialize terraform")
		} else {
			err := updateInitializedField(obj)
			if err != nil {
				log.Error(err, "failed to initialize field")
			}
		}
	}

	err = terraformApply(resPath, stateFile)
	if err != nil {
		log.Error(err, "unable to apply terraform")
	}

	err = updateStatusOut(obj, stateFile)
	if err != nil {
		log.Error(err, "unable to update status out field")
	}

	r.updateResource(resource.Spec.GVR, obj)
	return ctrl.Result{}, nil
}

func (r *ResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&terraformv1alpha1.Resource{}).
		Watches(&source.Kind{Type: &terraformv1alpha1.LinodeInstance{}}, &handler.Funcs{
			CreateFunc: func(event event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.createResource("linodeinstances", event.Meta.GetName(), event.Meta.GetNamespace(), event.Meta.GetResourceVersion())
			},
			DeleteFunc: func(deleteEvent event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.updateResourceVersion("linodeinstances", deleteEvent.Meta.GetName(), deleteEvent.Meta.GetNamespace(), deleteEvent.Meta.GetResourceVersion())
			},
			UpdateFunc: func(updateEvent event.UpdateEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.updateResourceVersion("linodeinstances", updateEvent.MetaNew.GetName(), updateEvent.MetaNew.GetNamespace(), updateEvent.MetaNew.GetResourceVersion())
			},
		}).
		Watches(&source.Kind{Type: &terraformv1alpha1.DigitaloceanDroplet{}}, &handler.Funcs{
			CreateFunc: func(event event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.createResource("digitaloceandroplets", event.Meta.GetName(), event.Meta.GetNamespace(), event.Meta.GetResourceVersion())
			},
			DeleteFunc: func(deleteEvent event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.updateResourceVersion("digitaloceandroplets", deleteEvent.Meta.GetName(), deleteEvent.Meta.GetNamespace(), deleteEvent.Meta.GetResourceVersion())
			},
			UpdateFunc: func(updateEvent event.UpdateEvent, limitingInterface workqueue.RateLimitingInterface) {
				r.updateResourceVersion("digitaloceandroplets", updateEvent.MetaNew.GetName(), updateEvent.MetaNew.GetNamespace(), updateEvent.MetaNew.GetResourceVersion())
			},
		}).
		Complete(r)
}

func updateStatusOut(u *unstructured.Unstructured, resPath string) error {
	//TODO: Handle output
	//codeUi := &CodeUi{
	//	OutputBuffer: new(bytes.Buffer),
	//}
	//showCmd := command.ShowCommand{
	//	Meta: command.Meta{
	//		Ui: codeUi,
	//	},
	//}
	//
	//args := []string{
	//	"-json",
	//	resPath,
	//}
	//
	//x := showCmd.Run(args)
	//if x != 0 {
	//	return errors.New("failed to run terraform show command")
	//}
	//
	//out := codeUi.OutputBuffer.String()

	out := "updated"

	return unstructured.SetNestedField(u.Object, out, "status", "out")
}

func configmapToTFProvider(config *corev1.ConfigMap, providerFile string) error {
	d1 := []byte(`{ "provider": { "` + config.Name + `":`)
	providerJson, err := json.Marshal(config.Data)
	if err != nil {
		return err
	}
	d1 = append(d1, providerJson...)
	d1 = append(d1, []byte("} }")...)

	prettyData, err := prettyJSON(d1)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(providerFile, prettyData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func crdToTFResource(kind string, obj *unstructured.Unstructured, mainFile string) error {
	objMap := obj.Object

	d1 := []byte(`{"resource":{ "` + kindToResouceMap[kind] + `":{"` + obj.GetName() + `":`)

	instanceSpecJson, err := json.Marshal(objMap["spec"])
	if err != nil {
		return err
	}
	d1 = append(d1, instanceSpecJson...)
	d1 = append(d1, []byte("} } }")...)
	prettyData, err := prettyJSON(d1)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(mainFile, prettyData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func terraformInit(resPath string) error {
	// TODO: This initializes terraform in the current directory. Shouldn't it be moved to the resource directory?
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}
	initCommand := command.InitCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	args := []string{
		resPath,
	}

	x := initCommand.Run(args)

	if x != 0 {
		return errors.New("failed to run terraform init command")
	}

	return nil
}

func terraformApply(resPath, stateFile string) error {
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}

	cmd := command.ApplyCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
	}

	args := []string{
		"-auto-approve",
		"-state",
		stateFile,
		resPath,
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform apply command")
	}

	return nil
}

func terraformDestroy(resPath, stateFile string) error {
	codeUi := &CodeUi{
		OutputBuffer: new(bytes.Buffer),
	}

	cmd := command.ApplyCommand{
		Meta: command.Meta{
			Ui: codeUi,
		},
		Destroy: true,
	}

	args := []string{
		"-auto-approve",
		"-state",
		stateFile,
		resPath,
	}
	x := cmd.Run(args)
	if x != 0 {
		return errors.New("failed to run terraform destroy command")
	}

	return nil
}

func prettyJSON(byteJson []byte) ([]byte, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, byteJson, "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), err
}

func hasFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}

	return false
}

func addFinalizer(u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	finalizers = append(finalizers, finalizer)
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	return nil
}

func removeFinalizer(u *unstructured.Unstructured) error {
	err := unstructured.SetNestedStringSlice(u.Object, []string{}, "metadata", "finalizers")
	if err != nil {
		return err
	}

	return nil
}

func createFiles(resPath, stateFile, providerFile, mainFile string) error {
	_, err := os.Stat(resPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(resPath, 0777)
		if err != nil {
			return err
		}
		_, err = os.Create(providerFile)
		if err != nil {
			return err
		}

		_, err = os.Create(mainFile)
		if err != nil {
			return err
		}

		_, err = os.Create(stateFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteFiles(resPath string) error {
	err := os.RemoveAll(resPath)
	if err != nil {
		return err
	}

	return nil
}

func updateInitializedField(obj *unstructured.Unstructured) error {
	err := unstructured.SetNestedField(obj.Object, true, "status", "initialized")
	if err != nil {
		return errors.Wrap(err, "failed to update field")
	}

	return nil
}

func (r *ResourceReconciler) createResource(kind, name, namespace, rv string) {
	// TODO: make a namer
	resName := kind + "-" + name + "-" + namespace

	err := r.Client.Get(context.Background(), types.NamespacedName{
		Namespace: corev1.NamespaceDefault,
		Name:      resName,
	}, &terraformv1alpha1.Resource{})

	if err != nil {
		err = r.Client.Create(context.Background(), &terraformv1alpha1.Resource{
			ObjectMeta: metav1.ObjectMeta{
				Name:      resName,
				Namespace: corev1.NamespaceDefault,
			},
			Spec: terraformv1alpha1.ResourceSpec{
				GVR: schema.GroupVersionResource{
					Group:    terraformv1alpha1.GroupVersion.Group,
					Version:  terraformv1alpha1.GroupVersion.Version,
					Resource: kind,
				},
				Name:            name,
				Namespace:       namespace,
				ResourceVersion: rv,
			},
		})

		if err != nil {
			r.Log.Error(err, "failed to create resource")
		}
	}
}

func (r *ResourceReconciler) updateResourceVersion(kind, name, namespace, rv string) {
	r.Log.Info("Updating resource version")

	var resource terraformv1alpha1.Resource
	err := r.Client.Get(context.Background(), types.NamespacedName{
		Namespace: corev1.NamespaceDefault,
		Name:      kind + "-" + name + "-" + namespace,
	}, &resource)
	if err != nil {
		r.Log.Error(err, "failed to get resource")
	}

	resource.Spec.ResourceVersion = rv

	err = r.Client.Update(context.Background(), &resource)
	if err != nil {
		r.Log.Error(err, "failed to update resource")
	}
}

func (r *ResourceReconciler) updateResource(gvr schema.GroupVersionResource, u *unstructured.Unstructured) {
	_, err := r.DynClient.Resource(gvr).Namespace(u.GetNamespace()).Update(u, metav1.UpdateOptions{})
	if err != nil {
		r.Log.Error(err, "failed to update resource")
	}
}

// TODO: how to handle resource name?
var kindToResouceMap = map[string]string{
	"linodeinstances":      "linode_instance",
	"digitaloceandroplets": "digitalocean_droplet",
}
