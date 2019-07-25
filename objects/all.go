/*
Copyright 2019 The Kubernetes Authors.

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

package objects

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/kustomize/v3/k8sdeps/kunstruct"
	"sigs.k8s.io/kustomize/v3/k8sdeps/transformer"
	"sigs.k8s.io/kustomize/v3/k8sdeps/validator"
	"sigs.k8s.io/kustomize/v3/pkg/fs"
	"sigs.k8s.io/kustomize/v3/pkg/ifc"
	"sigs.k8s.io/kustomize/v3/pkg/loader"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/resource"
	"sigs.k8s.io/kustomize/v3/pkg/target"
	"sigs.k8s.io/kustomize/v3/pkg/transformers"
)

type Provider struct {
	// Name is the name of the provider, the Github project name
	Name string
	// Manager image the name of the manager image for this provider
	ManagerImage string

	// Version is a git ref that you want to use
	Version string

	// Organization is the github organization
	Organization string

	// The local or remote path to a CRD
	CRDPath string

	// CustomImage will override whatever image this provier uses for the manager
	CustomImage string
	// The kubernetes Kind of the manager (Deployment, StatefulSet, etc)
	ManagerKind string
}

func (p *Provider) GetCRDPath() string {
	if p.CRDPath != "" {
		return p.CRDPath
	}
	return fmt.Sprintf("https://github.com/%s/%s/config/default?ref=%s", p.Organization, p.Name, p.Version)
}

// OverrideImage is a MutateFunction found in Kustomize's transform.
func (p *Provider) OverrideImage(item interface{}) (interface{}, error) {
	// Oh yes, we are sure this is a string
	s := item.(string)
	// IMAGE_URL is the default provided by kubebuilder without ever generating a real patch
	if strings.Contains(s, "manager") || s == "IMAGE_URL" || strings.Contains(s, "cluster-api-controller") {
		return p.CustomImage, nil
	}
	// We did not find what we were looking for
	return item, nil
}

func (p *Provider) GetCRDs() ([]runtime.Object, error) {
	v := validator.NewKustValidator()
	fSys := fs.MakeRealFS()

	uf := kunstruct.NewKunstructuredFactoryImpl()

	pf := transformer.NewFactoryImpl()
	rf := resmap.NewFactory(resource.NewFactory(uf), pf)

	var l ifc.Loader
	var err error
	crdPath := p.GetCRDPath()
	l, err = loader.NewFileLoaderAtCwd(v, fSys).New(p.CRDPath)
	if strings.HasPrefix(crdPath, "https://") {
		l, err = loader.NewLoader(loader.RestrictionNone, v, crdPath, fSys)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	kt, err := target.NewKustTarget(l, rf, nil, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m, err := kt.MakePruneConfigMap()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, res := range m.Resources() {
		// Don't do anything if there isn't anything to do
		if p.CustomImage == "" {
			continue
		}

		// Only run the transform on something that is probably the manager
		if res.GetKind() != p.ManagerKind {
			continue
		}

		m := res.Map()
		if err := transformers.MutateField(m, []string{"spec", "template", "spec", "containers", "image[]"}, false, p.OverrideImage); err != nil {
			return nil, errors.WithStack(err)
		}
		// Update the resource
		res.SetMap(m)
	}
	y, err := m.AsYaml()
	m.Resources()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return DecodeCAPIObjects(bytes.NewReader(y))
}

// GetManagementCluster returns all the objects needed to create a working management cluster
func GetManagementCluster(capiImage, capiRef, ipImage, ipRef, bpImage, bpRef string) ([]runtime.Object, error) {
	infrastructureDocker := &Provider{
		Organization: "kubernetes-sigs",
		Name:         "cluster-api-provider-docker",
		CRDPath:      "config/default",
		Version:      ipRef,
		ManagerKind:  "Deployment",
		CustomImage:  ipImage,
	}
	bootstrap := &Provider{
		Organization: "kubernetes-sigs",
		Name:         "cluster-api-bootstrap-provider-kubeadm",
		Version:      bpRef,
		ManagerKind:  "Deployment",
		CustomImage:  bpImage,
	}
	capi := &Provider{
		Organization: "kubernetes-sigs",
		Name:         "cluster-api",
		Version:      capiRef,
		ManagerKind:  "StatefulSet",
		CustomImage:  capiImage,
	}

	providers := []*Provider{
		infrastructureDocker,
		bootstrap,
		capi,
	}

	out := []runtime.Object{}

	for _, provider := range providers {
		objs, err := provider.GetCRDs()
		if err != nil {
			return nil, err
		}
		out = append(out, objs...)
	}
	return out, nil
}
