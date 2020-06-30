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
	"errors"
	"time"

	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCRD() GomegaAsyncAssertion {
	return Eventually(
		func() error {
			// Check ServiceAccount CRD
			if _, err := f.kubeformClient.GoogleV1alpha1().ServiceAccounts(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check ResourceGroup CRD
			if _, err := f.kubeformClient.AzurermV1alpha1().ResourceGroups(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD ResourceGroup is not ready")
			}

			// Check DbInstance CRD
			if _, err := f.kubeformClient.AwsV1alpha1().S3Buckets(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD S3Buckets is not ready")
			}

			// Check Instances CRD
			if _, err := f.kubeformClient.LinodeV1alpha1().Instances(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD Instances is not ready")
			}

			// Check Droplets CRD
			if _, err := f.kubeformClient.DigitaloceanV1alpha1().Droplets(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD Droplets is not ready")
			}

			// Check GoogleServiceAccount CRD
			if _, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD GoogleServiceAccount is not ready")
			}
			return nil
		},
		time.Minute*2,
		time.Second*10,
	)
}
