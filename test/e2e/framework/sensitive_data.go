package framework

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DBInstanceSecretName = "dbinstance-secret"
	InstanceSecretName   = "instance-secret"
)

func (i *Invocation) DBInstanceSensitiveData() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DBInstanceSecretName,
			Namespace: i.Namespace(),
		},
		Type: "kubeform.com/aws",
		Data: map[string][]byte{
			"password": []byte("thisIsAPassword123!"),
		},
	}
}

func (i *Invocation) InstanceSensitiveData() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      InstanceSecretName,
			Namespace: i.Namespace(),
		},
		Type: "kubeform.com/linode",
		Data: map[string][]byte{
			"root_pass": []byte("thisIsAPassword123!"),
			"stackscript_data": []byte(`{
"user" : "Fahim Abrar"
}`),
		},
	}
}
