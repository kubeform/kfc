package controllers

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/appscode/go/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/dynamic/dynamiclister"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"kmodules.xyz/client-go/tools/queue"
)

const controllerAgentName = "kfc"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Resource is synced
	SuccessSynced = "Synced"

	// MessageResourceSynced is the message used for an Event fired when a Resource
	// is synced successfully
	MessageResourceSynced = "Resource synced successfully"
)

// Controller is the controller implementation for KubeForm resources
type Controller struct {
	sync.Mutex

	kubeclientset kubernetes.Interface
	dynamicclient dynamic.Interface

	crdListers map[schema.GroupVersionResource]dynamiclister.Lister
	crdWorkers map[schema.GroupVersionResource]*queue.Worker
	syncedFns  []cache.InformerSynced

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	dynamicclient dynamic.Interface,
	factory dynamicinformer.DynamicSharedInformerFactory,
	gvrs []schema.GroupVersionResource) *Controller {
	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	//utilruntime.Must(kubeformscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset: kubeclientset,
		dynamicclient: dynamicclient,
		crdListers:    make(map[schema.GroupVersionResource]dynamiclister.Lister),
		crdWorkers:    make(map[schema.GroupVersionResource]*queue.Worker),
		recorder:      recorder,
	}
	klog.Info("Setting up event handlers")
	maxNumRequeues := 5
	numThreads := 5
	for i := range gvrs {
		gvr := gvrs[i]
		i := factory.ForResource(gvr)
		q := queue.New(gvr.String(), maxNumRequeues, numThreads, func(key string) error {
			return controller.reconcile(gvr, key)
		})

		i.Informer().AddEventHandler(queue.DefaultEventHandler(q.GetQueue()))

		controller.crdListers[gvr] = dynamiclister.New(i.Informer().GetIndexer(), gvr)
		controller.crdWorkers[gvr] = q
		controller.syncedFns = append(controller.syncedFns, i.Informer().HasSynced)
	}

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting KubeForm controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.syncedFns...); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch workers to process KubeForm resources
	for _, w := range c.crdWorkers {
		w.Run(stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *Controller) Lister(gvr schema.GroupVersionResource) dynamiclister.Lister {
	c.Lock()
	defer c.Unlock()
	return c.crdListers[gvr]
}

// reconcile compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the KubeForm resource
// with the current status of the resource.
func (c *Controller) reconcile(gvr schema.GroupVersionResource, key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the resource with this namespace/name
	obj, err := c.Lister(gvr).Namespace(namespace).Get(name)
	if err != nil {
		// The resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("resource '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	obj.GetKind()

	// TODO: make a namer
	resPath := filepath.Join(basePath, gvr.Resource+"."+namespace+"."+name)
	providerFile := filepath.Join(resPath, "provider.tf.json")
	mainFile := filepath.Join(resPath, "main.tf.json")
	stateFile := filepath.Join(resPath, "terraform.tfstate")

	if hasFinalizer(obj.GetFinalizers(), KFCFinalizer) {
		if obj.GetDeletionTimestamp() != nil {
			err := terraformDestroy(resPath)
			if err != nil {
				log.Error(err, "failed to terraform destroy")
			}

			err = deleteFiles(resPath)
			if err != nil {
				log.Error(err, "failed to delete files")
			}

			err = removeFinalizer(obj, KFCFinalizer)
			if err != nil {
				log.Error(err, "failed to remove finalizer")
			}

			c.updateResource(gvr, obj)
			return nil
		}
	} else {
		err := addFinalizer(obj, KFCFinalizer)
		if err != nil {
			log.Error(err, "failed to add finalizer")
		}
	}

	err = createFiles(resPath, providerFile, mainFile)
	if err != nil {
		log.Error(err, "failed to create files")
		return nil
	}

	providerRef, _, err := unstructured.NestedString(obj.Object, "spec", "providerRef", "name")
	if err != nil {
		log.Error(err, "failed to get providerRef")
		return nil
	}

	secret, err := c.kubeclientset.CoreV1().Secrets(namespace).Get(providerRef, metav1.GetOptions{})
	if err != nil {
		log.Error(err, "unable to fetch configmap")
	}

	providerName := strings.Split(gvr.Group, ".")[0]
	err = secretToTFProvider(secret, providerName, providerFile)
	if err != nil {
		log.Error(err, "unable to get configmap")
	}

	err = crdToTFResource(gvr.GroupVersion(), obj.GetKind(), namespace, providerName, c.kubeclientset, obj, mainFile)
	if err != nil {
		log.Error(err, "unable to get crd resource")
	}

	err = terraformInit(resPath)
	if err != nil {
		log.Error(err, "unable to initialize terraform")
	}

	createTFState(stateFile, gvr.GroupVersion(), obj)

	err = terraformApply(resPath)
	if err != nil {
		log.Error(err, "unable to apply terraform")
	}
	updateTFState(stateFile, gvr.GroupVersion(), obj)

	err = updateStatusOut(obj, resPath)
	if err != nil {
		log.Error(err, "unable to update status out field")
	}

	c.updateResource(gvr, obj)

	c.recorder.Event(obj, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateResource(gvr schema.GroupVersionResource, u *unstructured.Unstructured) {
	_, err := c.dynamicclient.Resource(gvr).Namespace(u.GetNamespace()).Update(u, metav1.UpdateOptions{})
	if err != nil {
		klog.Errorf("failed to update resource, reason : %s", err.Error())
	}
}
