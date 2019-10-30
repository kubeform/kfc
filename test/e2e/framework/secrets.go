package framework

import (
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) CreateSecret(obj *core.Secret) error {
	_, err := f.kubeClient.CoreV1().Secrets(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteSecret(obj v1.ObjectMeta) error {
	return f.kubeClient.CoreV1().Secrets(obj.Namespace).Delete(obj.Name, nil)
}
