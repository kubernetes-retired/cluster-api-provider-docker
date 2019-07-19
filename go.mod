module sigs.k8s.io/cluster-api-provider-docker

go 1.12

require (
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/go-openapi/strfmt v0.19.2 // indirect
	github.com/go-openapi/validate v0.19.2 // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.0
	github.com/prometheus/client_golang v0.9.4 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/grpc v1.22.0 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
	k8s.io/api v0.0.0-20190718062839-c8a0b81cb10e
	k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476 // indirect
	k8s.io/apimachinery v0.0.0-20190717022731-0bb8574e0887
	k8s.io/apiserver v0.0.0-20190718202433-cfb21c5903d4 // indirect
	k8s.io/cli-runtime v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/cloud-provider v0.0.0-20190717025205-585d8110a88f // indirect
	k8s.io/cluster-bootstrap v0.0.0-20181213155137-5f9271efc2e7 // indirect
	k8s.io/csi-api v0.0.0-20190313123203-94ac839bf26c // indirect
	k8s.io/klog v0.3.1
	k8s.io/kube-aggregator v0.0.0-20190718063459-502bbc9a656c // indirect
	k8s.io/kubernetes v1.13.1
	k8s.io/utils v0.0.0-20190712204705-3dccf664f023 // indirect
	sigs.k8s.io/cluster-api v0.0.0-20190607141803-aacb0c613ffb
	sigs.k8s.io/controller-runtime v0.1.10
	sigs.k8s.io/kind v0.4.0
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
	sigs.k8s.io/yaml v1.1.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20181213150558-05914d821849
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190704050804-bd8686edbd81
	k8s.io/client-go => k8s.io/client-go v10.0.0+incompatible
	k8s.io/kubernetes => k8s.io/kubernetes v1.13.1
)
