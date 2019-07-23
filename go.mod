module sigs.k8s.io/cluster-api-provider-docker

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/pkg/errors v0.8.1
	k8s.io/api v0.0.0-20190516230258-a675ac48af67
	k8s.io/apiextensions-apiserver v0.0.0-20190516231611-bf6753f2aa24 // indirect
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190516232516-d7d78ab2cfe7 // indirect
	k8s.io/component-base v0.0.0-20190516230545-9b07ebd4193b // indirect
	k8s.io/kubernetes v1.14.2
	sigs.k8s.io/cluster-api v0.0.0-20190723184414-6afa7aee7f42
	sigs.k8s.io/controller-runtime v0.2.0-beta.4
	sigs.k8s.io/kind v0.4.0
)
