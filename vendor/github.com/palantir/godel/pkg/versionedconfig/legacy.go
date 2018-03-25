// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package versionedconfig

import (
	"gopkg.in/yaml.v2"
)

// ConfigWithLegacy is a struct with a "legacy" YAML field that stores a boolean that indicates whether or not the
// configuration is "legacy" configuration.
type ConfigWithLegacy struct {
	Legacy bool `yaml:"legacy-config"`
}

func IsLegacyConfig(cfgBytes []byte) bool {
	var cfg ConfigWithLegacy
	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		return false
	}
	return cfg.Legacy
}
