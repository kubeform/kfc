package framework

import (
	"time"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"
	"kubeform.dev/kubeform/apis/google/v1alpha1"

	. "github.com/onsi/gomega"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Invocation) ServiceAccount(name string) *v1alpha1.ServiceAccount {
	return &v1alpha1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.ServiceAccountSpec{
			AccountID: "test-account",
			ProviderRef: v12.LocalObjectReference{
				Name: GoogleProviderRef,
			},
			DisplayName: "kfc-e2e-test",
			Project:     "appscode-testing",
		},
	}
}

func (f *Framework) CreateServiceAccount(obj *v1alpha1.ServiceAccount) error {
	_, err := f.kubeformClient.GoogleV1alpha1().ServiceAccounts(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteServiceAccount(meta metav1.ObjectMeta) error {
	return f.kubeformClient.GoogleV1alpha1().ServiceAccounts(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyServiceAccountRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			serviceAccount, err := f.kubeformClient.GoogleV1alpha1().ServiceAccounts(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return serviceAccount.Status.Phase == base.PhaseRunning
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyServiceAccountDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.GoogleV1alpha1().ServiceAccounts(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
