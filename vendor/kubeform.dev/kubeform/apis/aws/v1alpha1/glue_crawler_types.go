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

type GlueCrawler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueCrawlerSpec   `json:"spec,omitempty"`
	Status            GlueCrawlerStatus `json:"status,omitempty"`
}

type GlueCrawlerSpecDynamodbTarget struct {
	Path string `json:"path" tf:"path"`
}

type GlueCrawlerSpecJdbcTarget struct {
	ConnectionName string `json:"connectionName" tf:"connection_name"`
	// +optional
	Exclusions []string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`
	Path       string   `json:"path" tf:"path"`
}

type GlueCrawlerSpecS3Target struct {
	// +optional
	Exclusions []string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`
	Path       string   `json:"path" tf:"path"`
}

type GlueCrawlerSpecSchemaChangePolicy struct {
	// +optional
	DeleteBehavior string `json:"deleteBehavior,omitempty" tf:"delete_behavior,omitempty"`
	// +optional
	UpdateBehavior string `json:"updateBehavior,omitempty" tf:"update_behavior,omitempty"`
}

type GlueCrawlerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Classifiers []string `json:"classifiers,omitempty" tf:"classifiers,omitempty"`
	// +optional
	Configuration string `json:"configuration,omitempty" tf:"configuration,omitempty"`
	DatabaseName  string `json:"databaseName" tf:"database_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	DynamodbTarget []GlueCrawlerSpecDynamodbTarget `json:"dynamodbTarget,omitempty" tf:"dynamodb_target,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	JdbcTarget []GlueCrawlerSpecJdbcTarget `json:"jdbcTarget,omitempty" tf:"jdbc_target,omitempty"`
	Name       string                      `json:"name" tf:"name"`
	Role       string                      `json:"role" tf:"role"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	S3Target []GlueCrawlerSpecS3Target `json:"s3Target,omitempty" tf:"s3_target,omitempty"`
	// +optional
	Schedule string `json:"schedule,omitempty" tf:"schedule,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SchemaChangePolicy []GlueCrawlerSpecSchemaChangePolicy `json:"schemaChangePolicy,omitempty" tf:"schema_change_policy,omitempty"`
	// +optional
	SecurityConfiguration string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`
	// +optional
	TablePrefix string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`
}

type GlueCrawlerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueCrawlerList is a list of GlueCrawlers
type GlueCrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueCrawler CRD objects
	Items []GlueCrawler `json:"items,omitempty"`
}