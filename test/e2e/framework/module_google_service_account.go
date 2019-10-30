package framework

import (
	"os"
	"time"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"
	"kubeform.dev/kubeform/apis/modules/v1alpha1"

	. "github.com/onsi/gomega"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/constants/google"
)

func (i *Invocation) ModuleServiceAccount(name string) *v1alpha1.GoogleServiceAccount {
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
				Name: GoogleProviderRef,
			},
			Names:        []string{"single-account"},
			Prefix:       "kfc-e2e-testing",
			ProjectID:    os.Getenv(google.GOOGLE_PROJECT_ID),
			ProjectRoles: []string{os.Getenv(google.GOOGLE_PROJECT_ID) + "=>roles/viewer"},
		},
	}
}

func (f *Framework) CreateModuleServiceAccount(obj *v1alpha1.GoogleServiceAccount) error {
	_, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteModuleServiceAccount(meta metav1.ObjectMeta) error {
	return f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyModuleServiceAccountRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			serviceAccount, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
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
			_, err := f.kubeformClient.ModulesV1alpha1().GoogleServiceAccounts(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
