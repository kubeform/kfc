package framework

import (
	kfclient "kubeform.dev/kubeform/client/clientset/versioned"

	"github.com/appscode/go/crypto/rand"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Framework struct {
	restConfig     *rest.Config
	kubeClient     kubernetes.Interface
	kubeformClient kfclient.Interface
	namespace      string
	name           string
}

func New(
	restConfig *rest.Config,
	kubeClient kubernetes.Interface,
	kubeformClient kfclient.Interface,
) *Framework {
	return &Framework{
		restConfig:     restConfig,
		kubeClient:     kubeClient,
		kubeformClient: kubeformClient,
		name:           "kfc",
		namespace:      rand.WithUniqSuffix("kubeform"),
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("kfc-e2e"),
	}
}

func (fi *Invocation) KubeClient() kubernetes.Interface {
	return fi.kubeClient
}

func (fi *Invocation) RestConfig() *rest.Config {
	return fi.restConfig
}

func (fi *Invocation) KubeformClient() kfclient.Interface {
	return fi.kubeformClient
}

type Invocation struct {
	*Framework
	app string
}
