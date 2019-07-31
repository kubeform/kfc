package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OpsworksApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksApplicationSpec   `json:"spec,omitempty"`
	Status            OpsworksApplicationStatus `json:"status,omitempty"`
}

type OpsworksApplicationSpecAppSource struct {
	// +optional
	// +optional
	Revision string `json:"revision,omitempty" tf:"revision,omitempty"`
	// +optional
	SshKey string `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`
	Type   string `json:"type" tf:"type"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type OpsworksApplicationSpecEnvironment struct {
	Key string `json:"key" tf:"key"`
	// +optional
	Secure bool   `json:"secure,omitempty" tf:"secure,omitempty"`
	Value  string `json:"value" tf:"value"`
}

type OpsworksApplicationSpecSslConfiguration struct {
	Certificate string `json:"certificate" tf:"certificate"`
	// +optional
	Chain string `json:"chain,omitempty" tf:"chain,omitempty"`
}

type OpsworksApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	AppSource []OpsworksApplicationSpecAppSource `json:"appSource,omitempty" tf:"app_source,omitempty"`
	// +optional
	AutoBundleOnDeploy string `json:"autoBundleOnDeploy,omitempty" tf:"auto_bundle_on_deploy,omitempty"`
	// +optional
	AwsFlowRubySettings string `json:"awsFlowRubySettings,omitempty" tf:"aws_flow_ruby_settings,omitempty"`
	// +optional
	DataSourceArn string `json:"dataSourceArn,omitempty" tf:"data_source_arn,omitempty"`
	// +optional
	DataSourceDatabaseName string `json:"dataSourceDatabaseName,omitempty" tf:"data_source_database_name,omitempty"`
	// +optional
	DataSourceType string `json:"dataSourceType,omitempty" tf:"data_source_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DocumentRoot string `json:"documentRoot,omitempty" tf:"document_root,omitempty"`
	// +optional
	Domains []string `json:"domains,omitempty" tf:"domains,omitempty"`
	// +optional
	EnableSsl bool `json:"enableSsl,omitempty" tf:"enable_ssl,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Environment []OpsworksApplicationSpecEnvironment `json:"environment,omitempty" tf:"environment,omitempty"`
	Name        string                               `json:"name" tf:"name"`
	// +optional
	RailsEnv string `json:"railsEnv,omitempty" tf:"rails_env,omitempty"`
	// +optional
	ShortName string `json:"shortName,omitempty" tf:"short_name,omitempty"`
	// +optional
	SslConfiguration []OpsworksApplicationSpecSslConfiguration `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`
	StackID          string                                    `json:"stackID" tf:"stack_id"`
	Type             string                                    `json:"type" tf:"type"`
}

type OpsworksApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksApplicationList is a list of OpsworksApplications
type OpsworksApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksApplication CRD objects
	Items []OpsworksApplication `json:"items,omitempty"`
}