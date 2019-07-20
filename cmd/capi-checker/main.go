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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"sigs.k8s.io/cluster-api-provider-docker/cmd/capi-checker/pkg/check"
	"sigs.k8s.io/cluster-api-provider-docker/cmd/versioninfo"
)

// Constants
const (
	KUBECONFIG = "KUBECONFIG"
)

func main() {
	fmt.Print("Running ", versioninfo.VersionInfo(`"capi-checker"`))
	kubeConfigEnv := os.Getenv(KUBECONFIG)

	checkConfig := flag.String("check-yaml", "", "Path to the YAML file with the check configuration")
	kubeConfig := flag.String("kube-config", "~/.kube/config", "Path to the kube config file for the cluster for which checks need to be run")
	flag.Parse()

	if *checkConfig == "" {
		fmt.Printf("check-yaml is required but not supplied\n")
		flag.PrintDefaults()
		os.Exit(101)
	}

	if kubeConfigEnv != "" {
		fmt.Printf("Using kubeconfig %q from env var %q\n", kubeConfigEnv, KUBECONFIG)
		*kubeConfig = kubeConfigEnv
	}

	fmt.Printf("Reading check config from %q\n", *checkConfig)

	data, err := ioutil.ReadFile(*checkConfig)

	if err != nil {
		fmt.Printf("Failed to read file %q; %v\n", *checkConfig, err)
		os.Exit(201)
	}

	conf, err := check.ConfigFromYaml(data)
	if err != nil {
		fmt.Printf("Invalid yaml file %q\n", *checkConfig)
		os.Exit(301)
	}

	fmt.Printf("Running checks %q\n", conf.Name)
	for _, check := range conf.Checks {
		fmt.Printf("🧪  Checking for %v...\n", check)
		result, err := check.Run(*kubeConfig)
		if err != nil {
			fmt.Printf("\t 🛑    Failed, Reason=%q\n", err)
			continue
		}
		fmt.Printf("\t ✅    Result: %s\n", result)
	}
}
