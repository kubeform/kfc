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

package framework

import (
	"context"
	"time"

	"kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
)

func (i *Invocation) ResourceGroup(name string, secretName string) *v1alpha1.ResourceGroup {
	return &v1alpha1.ResourceGroup{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.ResourceGroupSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: secretName,
			},
			Name:     name,
			Location: "East US",
			Tags: map[string]string{
				"env": "testing",
			},
		},
	}
}

func (f *Framework) CreateResourceGroup(obj *v1alpha1.ResourceGroup) error {
	_, err := f.kubeformClient.AzurermV1alpha1().ResourceGroups(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) DeleteResourceGroup(meta metav1.ObjectMeta) error {
	return f.kubeformClient.AzurermV1alpha1().ResourceGroups(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInForeground())
}

func (f *Framework) EventuallyResourceGroupRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			resourceGroup, err := f.kubeformClient.AzurermV1alpha1().ResourceGroups(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return resourceGroup.Status.Phase == base.PhaseRunning
		},
		time.Minute*30,
		time.Second*10,
	)
}

func (f *Framework) EventuallyResourceGroupDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.AzurermV1alpha1().ResourceGroups(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*30,
		time.Second*10,
	)
}
