module github.com/appscode-cloud/kfc

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/hashicorp/terraform v0.11.14
	github.com/mitchellh/cli v0.0.0-20171129193617-33edc47170b5
	github.com/mitchellh/colorstring v0.0.0-20150917214807-8631ce90f286
	github.com/onsi/ginkgo v1.6.0
	github.com/onsi/gomega v1.4.2
	github.com/pkg/errors v0.8.1
	golang.org/x/net v0.0.0-20190501004415-9ce7a6920f09
	golang.org/x/sys v0.0.0-20190502145724-3ef323f4f1fd // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190425155659-357c62f0e4bb // indirect
	google.golang.org/grpc v1.20.1 // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0-beta.2
)

replace cloud.google.com/go => cloud.google.com/go v0.15.0
