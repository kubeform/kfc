module kubeform.dev/kfc

go 1.12

require (
	ekyu.moe/base91 v0.2.3
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/fatih/structs v1.1.0
	github.com/gobuffalo/flect v0.2.1
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gophercloud/gophercloud v0.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/json-iterator/go v1.1.8
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/spf13/cobra v0.0.5
	gocloud.dev v0.20.0
	k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog v1.0.0
	kmodules.xyz/client-go v0.0.0-20200630053911-20d035822d35
	kmodules.xyz/constants v0.0.0-20200506032633-a21e58ceec72
	kubeform.dev/kubeform v0.1.1-0.20200630094658-eeabd42236da
)

replace github.com/linode/linodego => github.com/linode/linodego v0.19.0

replace github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.25.4

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.2.0

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/keybase/go-crypto v0.0.0-20190523171820-b785b22cc757 => github.com/keybase/go-crypto v0.0.0-20190416182011-b785b22cc757

replace github.com/terraform-providers/terraform-provider-google v2.17.0+incompatible => github.com/terraform-providers/terraform-provider-google v1.20.1-0.20191008212436-363f2d283518

replace github.com/terraform-providers/terraform-provider-aws v2.32.0+incompatible => github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20191010190908-1261a98537f2

replace github.com/terraform-providers/terraform-provider-random v2.2.1+incompatible => github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f

replace (
	bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5
	// github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.0.0
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200513171258-e048e166ab9c
	// google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	// google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => github.com/kmodules/api v0.18.4-0.20200524125823-c8bc107809b9
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423
	k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200521033432-49d3646051ad
)
