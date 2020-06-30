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
	"os"
	"time"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"
	"kubeform.dev/kubeform/apis/modules/v1alpha1"

	. "github.com/onsi/gomega"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/constants/google"
)

func (i *Invocation) ModuleServiceAccount(name string, secretName string) *v1alpha1.GoogleServiceAccount {
	return &v1alpha1.GoogleServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.GoogleServiceAccountSpec{
			ProviderRef: v12.LocalObjectReference{
				Name: secretName,
			},
			Names:        []string{"single-account"},
			Prefix:       name,
			ProjectID:    os.Getenv(google.GOOGLE_PROJECT_ID),
			ProjectRoles: []string{os.Getenv(google.GOOGLE_PROJECT_ID) + "=>roles/viewer"},
		},
	}
}

func (f *Framework) CreateModuleServiceAccount(obj *v1alpha1.GoogleServiceAccount) error {
	_, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) DeleteModuleServiceAccount(meta metav1.ObjectMeta) error {
	return f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInForeground())
}

func (f *Framework) EventuallyModuleServiceAccountRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			serviceAccount, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return serviceAccount.Status.Phase == base.PhaseRunning
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyModuleServiceAccountDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
