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
	"strings"
	"time"

	"github.com/appscode-cloud/kfc/controllers"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/sample-controller/pkg/signals"

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

	exampleInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicClient, time.Second*30)

	extClient, err := apiextension.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes api extension clientset: %s", err.Error())
	}

	crds, err := extClient.ApiextensionsV1beta1().CustomResourceDefinitions().List(v1.ListOptions{})
	if err != nil {
		klog.Fatalf("Error listing CRDs: %s", err.Error())
	}

	var gvrs []schema.GroupVersionResource

	for _, crd := range crds.Items {
		if strings.Contains(crd.Spec.Group, "kubeform.com") {
			gvr := schema.GroupVersionResource{
				Group:    crd.Spec.Group,
				Version:  crd.Spec.Version,
				Resource: crd.Spec.Names.Plural,
			}

			gvrs = append(gvrs, gvr)
		}
	}

	controller := controllers.NewController(kubeClient, dynamicClient, exampleInformerFactory, gvrs)

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	exampleInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "/home/fahim/.kube/config", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
