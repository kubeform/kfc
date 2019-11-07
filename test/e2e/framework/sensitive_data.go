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
