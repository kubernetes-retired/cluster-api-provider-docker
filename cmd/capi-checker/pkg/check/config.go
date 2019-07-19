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

package check

import (
	"sigs.k8s.io/yaml"
)

// Config defines configuration for a check
type Config struct {
	Name   string  `yaml:"name"`
	Checks []Check `yaml:"checks"`
}

// ConfigFromYaml marshalls suppied yaml to Checks
func ConfigFromYaml(rawYaml []byte) (Config, error) {
	var c Config
	if err := yaml.Unmarshal(rawYaml, &c); err != nil {
		return c, err
	}

	return c, nil
}
