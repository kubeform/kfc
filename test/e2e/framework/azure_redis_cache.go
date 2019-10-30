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

	"kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Invocation) RedisCache(name string) *v1alpha1.RedisCache {
	return &v1alpha1.RedisCache{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.RedisCacheSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: AzureProviderRef,
			},
			Name:              "example-cache",
			ResourceGroupName: "dev",
			Location:          "East US",
			Capacity:          2,
			Family:            "C",
			SkuName:           "Standard",
			EnableNonSSLPort:  false,
			MinimumTLSVersion: "1.2",
		},
	}
}

func (f *Framework) CreateRedisCache(obj *v1alpha1.RedisCache) error {
	_, err := f.kubeformClient.AzurermV1alpha1().RedisCaches(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteRedisCache(meta metav1.ObjectMeta) error {
	return f.kubeformClient.AzurermV1alpha1().RedisCaches(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyRedisCacheRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			redisCache, err := f.kubeformClient.AzurermV1alpha1().RedisCaches(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return redisCache.Status.Phase == base.PhaseRunning
		},
		time.Minute*30,
		time.Second*10,
	)
}

func (f *Framework) EventuallyRedisCacheDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.AzurermV1alpha1().RedisCaches(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*30,
		time.Second*10,
	)
}
