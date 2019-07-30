module github.com/appscode-cloud/kfc

go 1.12

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.0 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.1.0 // indirect
	github.com/Azure/go-autorest/autorest/mocks v0.2.0 // indirect
	github.com/appscode/go v0.0.0-20190621064509-6b292c9166e3
	github.com/chzyer/logex v1.1.11-0.20160617073814-96a4d311aa9b // indirect
	github.com/fatih/structs v1.1.0
	github.com/go-toolsmith/astcast v1.0.0 // indirect
	github.com/go-toolsmith/astcopy v1.0.0 // indirect
	github.com/go-toolsmith/astfmt v1.0.0 // indirect
	github.com/go-toolsmith/astp v1.0.0 // indirect
	github.com/go-toolsmith/pkgload v1.0.0 // indirect
	github.com/go-toolsmith/typep v1.0.0 // indirect
	github.com/gobuffalo/flect v0.1.5
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.2-0.20171022020050-c20a0b1b1ea9 // indirect
	github.com/hashicorp/terraform v0.12.4
	github.com/json-iterator/go v1.1.6
	github.com/miekg/dns v1.0.14 // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/terraform-providers/terraform-provider-openstack v1.20.0 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	k8s.io/api v0.0.0-20190711103429-37c3b8b1ca65
	k8s.io/apiextensions-apiserver v0.0.0-20190516231611-bf6753f2aa24
	k8s.io/apimachinery v0.0.0-20190711222657-391ed67afa7b
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v0.3.1
	k8s.io/kube-openapi v0.0.0-20190502190224-411b2483e503 // indirect
	kmodules.xyz/client-go v0.0.0-20190715080709-7162a6c90b04
	kubeform.dev/kubeform v0.0.2-0.20190723075747-9001a322126a
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
