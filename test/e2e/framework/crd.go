package framework

import (
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
			if _, err := f.kubeformClient.GoogleV1alpha1().ServiceAccounts(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check RedisCache CRD
			if _, err := f.kubeformClient.AzurermV1alpha1().RedisCaches(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check DbInstance CRD
			if _, err := f.kubeformClient.AwsV1alpha1().DbInstances(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check Instances CRD
			if _, err := f.kubeformClient.LinodeV1alpha1().Instances(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check DatabaseCluster CRD
			if _, err := f.kubeformClient.DigitaloceanV1alpha1().DatabaseClusters(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}

			// Check GoogleServiceAccount CRD
			if _, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(core.NamespaceAll).List(metav1.ListOptions{}); err != nil {
				return errors.New("CRD ServiceAccount is not ready")
			}
			return nil
		},
		time.Minute*2,
		time.Second*10,
	)
}
