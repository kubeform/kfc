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
	"time"

	"kubeform.dev/kubeform/apis/aws/v1alpha1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Invocation) DBInstance(name string) *v1alpha1.DbInstance {
	return &v1alpha1.DbInstance{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.DbInstanceSpec{
			AllocatedStorage: 5,
			ProviderRef: corev1.LocalObjectReference{
				Name: AwsProviderRef,
			},
			StorageType:        "gp2",
			Engine:             "mysql",
			EngineVersion:      "5.7",
			SkipFinalSnapshot:  true,
			InstanceClass:      "db.t2.micro",
			Name:               name,
			Username:           "foo",
			ParameterGroupName: "default.mysql5.7",
			SecretRef: &corev1.LocalObjectReference{
				Name: DBInstanceSecretName,
			},
		},
	}
}

func (f *Framework) CreateDBInstance(obj *v1alpha1.DbInstance) error {
	_, err := f.kubeformClient.AwsV1alpha1().DbInstances(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteDBInstance(meta metav1.ObjectMeta) error {
	return f.kubeformClient.AwsV1alpha1().DbInstances(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyDbInstanceRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			dbInstance, err := f.kubeformClient.AwsV1alpha1().DbInstances(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return dbInstance.Status.Phase == base.PhaseRunning
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyDbInstanceDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.AwsV1alpha1().DbInstances(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
