module kubeform.dev/kfc

go 1.12

require (
	ekyu.moe/base91 v0.2.3
	github.com/appscode/go v0.0.0-20190808133642-1d4ef1f1c1e0
	github.com/fatih/structs v1.1.0
	github.com/gobuffalo/flect v0.1.5
	github.com/json-iterator/go v1.1.6
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	gocloud.dev v0.15.0
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
	k8s.io/api v0.0.0-20190711103429-37c3b8b1ca65
	k8s.io/apiextensions-apiserver v0.0.0-20190516231611-bf6753f2aa24
	k8s.io/apimachinery v0.0.0-20190711222657-391ed67afa7b
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v0.3.1
	kmodules.xyz/client-go v0.0.0-20190715080709-7162a6c90b04
	kubeform.dev/kubeform v0.0.3-0.20190926104403-a9109847c227i
	sigs.k8s.io/controller-runtime v0.2.0-beta.4
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest/autorest v0.5.0
	github.com/hashicorp/go-azure-helpers => github.com/hashicorp/go-azure-helpers v0.3.2
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190311093542-50b561225d70
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190314000054-4a91899592f4
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/kubernetes => k8s.io/kubernetes v1.14.0
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
)
