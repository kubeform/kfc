package v1alpha1

// xref: https://github.com/hashicorp/terraform/blob/d7bce857cf7e2ab8cd6378196ee95b9a935acb5c/states/statefile/version4.go#L497

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces
type State struct {
	Version          int64  `json:"version"`
	TerraformVersion string `json:"terraform_version"`
	Serial           uint64 `json:"serial"`
	Lineage          string `json:"lineage"`
}
