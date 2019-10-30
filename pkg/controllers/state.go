/*
Copyright The Kubeform Authors.

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
package controllers

import "encoding/json"

type state struct {
	Version          json.Number       `json:"version"`
	TerraformVersion string            `json:"terraform_version"`
	Serial           uint64            `json:"serial"`
	Lineage          string            `json:"lineage"`
	RootOutputs      map[string]output `json:"outputs"`
	Resources        []resource        `json:"resources"`
}

type output struct {
	ValueRaw     json.RawMessage `json:"value"`
	ValueTypeRaw json.RawMessage `json:"type"`
	Sensitive    bool            `json:"sensitive,omitempty"`
}

type resource struct {
	Module         string     `json:"module,omitempty"`
	Mode           string     `json:"mode"`
	Type           string     `json:"type"`
	Name           string     `json:"name"`
	EachMode       string     `json:"each,omitempty"`
	ProviderConfig string     `json:"provider"`
	Instances      []instance `json:"instances"`
}

type instance struct {
	IndexKey interface{} `json:"index_key,omitempty"`
	Status   string      `json:"status,omitempty"`
	Deposed  string      `json:"deposed,omitempty"`

	SchemaVersion  uint64            `json:"schema_version"`
	AttributesRaw  json.RawMessage   `json:"attributes,omitempty"`
	AttributesFlat map[string]string `json:"attributes_flat,omitempty"`

	PrivateRaw []byte `json:"private,omitempty"`

	Dependencies []string `json:"depends_on,omitempty"`
}
