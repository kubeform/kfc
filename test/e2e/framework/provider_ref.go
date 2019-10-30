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

const (
	GoogleProviderRef       = "google"
	AwsProviderRef          = "aws"
	DigitalOceanProviderRef = "digitalocean"
	LinodeProviderRef       = "linode"
	AzureProviderRef        = "azure"
)

func (i *Invocation) GoogleProviderRef() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      GoogleProviderRef,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"credentials": []byte(google.ServiceAccountFromEnv()),
			"region":      []byte("us-central1"),
			"project":     []byte(os.Getenv(google.GOOGLE_PROJECT_ID)),
		},
	}
}

func (i *Invocation) AwsProviderRef() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      AwsProviderRef,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"access_key": []byte(os.Getenv(aws.AWS_ACCESS_KEY_ID)),
			"secret_key": []byte(os.Getenv(aws.AWS_SECRET_ACCESS_KEY)),
			"region":     []byte("us-east-1"),
		},
	}
}

func (i *Invocation) DigitalOceanProviderRef() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DigitalOceanProviderRef,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"token": []byte(os.Getenv(digitalocean.DIGITALOCEAN_TOKEN)),
		},
	}
}

func (i *Invocation) LinodeProviderRef() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      LinodeProviderRef,
			Namespace: i.Namespace(),
		},
		Data: map[string][]byte{
			"token": []byte(os.Getenv(linode.LINODE_API_TOKEN)),
		},
	}
}

func (i *Invocation) AzureProviderRef() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      AzureProviderRef,
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
