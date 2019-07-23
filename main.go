/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"

	aws "kubeform.dev/kubeform/apis/aws/install"
	azurerm "kubeform.dev/kubeform/apis/azurerm/install"
	digitalocean "kubeform.dev/kubeform/apis/digitalocean/install"
	google "kubeform.dev/kubeform/apis/google/install"

	"github.com/appscode/go/log"

	"github.com/appscode-cloud/kfc/signals"

	"k8s.io/client-go/rest"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/client-go/tools/cache"

	"github.com/appscode-cloud/kfc/controllers"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	linode "kubeform.dev/kubeform/apis/linode/install"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	//klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes dynamic clientset: %s", err.Error())
	}

	linode.Install(clientsetscheme.Scheme)
	aws.Install(clientsetscheme.Scheme)
	azurerm.Install(clientsetscheme.Scheme)
	digitalocean.Install(clientsetscheme.Scheme)
	google.Install(clientsetscheme.Scheme)

	controller := controllers.NewController(kubeClient, dynamicClient)

	watchCRD(cfg, stopCh, controller, dynamicClient)

	if err = controller.Run(stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube/config"), "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func watchCRD(cfg *rest.Config, stopCh <-chan struct{}, controller *controllers.Controller, dynamicClient dynamic.Interface) {
	crdClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1beta1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1beta1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string

			if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
				log.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				log.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				log.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "kubeform.com") {
				gvr := schema.GroupVersionResource{
					Group:    crd.Spec.Group,
					Version:  crd.Spec.Version,
					Resource: crd.Spec.Names.Plural,
				}

				err = controller.AddNewCRD(gvr, dynamicClient, stopCh)
				if err != nil {
					log.Error(err)
					return
				}
			}
		},
	})

	informerFactory.Start(stopCh)
}
