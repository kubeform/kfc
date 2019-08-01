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
