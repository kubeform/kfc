/*
Copyright The Kubeform Authors.

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

package cmds

import (
	"strings"
	"time"

	"kubeform.dev/kfc/pkg/controllers"
	aws "kubeform.dev/kubeform/apis/aws/install"
	azurerm "kubeform.dev/kubeform/apis/azurerm/install"
	digitalocean "kubeform.dev/kubeform/apis/digitalocean/install"
	google "kubeform.dev/kubeform/apis/google/install"
	linode "kubeform.dev/kubeform/apis/linode/install"
	modules "kubeform.dev/kubeform/apis/modules/install"

	"github.com/appscode/go/log"
	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	license "go.bytebuilders.dev/license-verifier/kubernetes"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"kmodules.xyz/client-go/tools/cli"
)

var (
	masterURL   string
	kubeconfig  string
	licenseFile string
)

func NewCmdRun(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Vault operator",
		DisableAutoGenTag: true,
		PreRun: func(c *cobra.Command, args []string) {
			cli.SendPeriodicAnalytics(c, v.Version.Version)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("Starting operator version %s+%s ...", v.Version.Version, v.Version.CommitHash)

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
			modules.Install(clientsetscheme.Scheme)

			// Start periodic license verification
			//nolint:errcheck
			go license.VerifyLicensePeriodically(cfg, licenseFile, stopCh)

			controller := controllers.NewController(kubeClient, dynamicClient)
			watchCRD(cfg, stopCh, controller, dynamicClient)

			return controller.Run(stopCh)
		},
	}

	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	cmd.Flags().StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	cmd.Flags().StringVar(&licenseFile, "license-file", licenseFile, "Path to license file")
	cmd.Flags().StringVar(&controllers.SecretKey, "secret-key", "", "A base64-encoded key, of length 32 bytes when decoded.")

	return cmd
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
