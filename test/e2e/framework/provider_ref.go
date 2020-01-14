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
	"os"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/constants/aws"
	"kmodules.xyz/constants/azure"
	"kmodules.xyz/constants/digitalocean"
	"kmodules.xyz/constants/google"
	"kmodules.xyz/constants/linode"
)

func (i *Invocation) GoogleProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"credentials": []byte(google.ServiceAccountFromEnv()),
			"region":      []byte("us-central1"),
			"project":     []byte(os.Getenv(google.GOOGLE_PROJECT_ID)),
		},
	}
}

func (i *Invocation) AwsProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"access_key": []byte(os.Getenv(aws.AWS_ACCESS_KEY_ID)),
			"secret_key": []byte(os.Getenv(aws.AWS_SECRET_ACCESS_KEY)),
			"region":     []byte("us-east-1"),
		},
	}
}

func (i *Invocation) DigitalOceanProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"token": []byte(os.Getenv(digitalocean.DIGITALOCEAN_TOKEN)),
		},
	}
}

func (i *Invocation) LinodeProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"token": []byte(os.Getenv(linode.LINODE_API_TOKEN)),
		},
	}
}

func (i *Invocation) AzureProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"subscription_id": []byte(os.Getenv(azure.AZURE_SUBSCRIPTION_ID)),
			"client_id":       []byte(os.Getenv(azure.AZURE_CLIENT_ID)),
			"client_secret":   []byte(os.Getenv(azure.AZURE_CLIENT_SECRET)),
			"tenant_id":       []byte(os.Getenv(azure.AZURE_TENANT_ID)),
		},
	}
}
