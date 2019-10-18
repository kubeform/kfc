package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type WafRegexMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRegexMatchSetSpec   `json:"spec,omitempty"`
	Status            WafRegexMatchSetStatus `json:"status,omitempty"`
}

type WafRegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafRegexMatchSetSpecRegexMatchTuple struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch       []WafRegexMatchSetSpecRegexMatchTupleFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	RegexPatternSetID  string                                            `json:"regexPatternSetID" tf:"regex_pattern_set_id"`
	TextTransformation string                                            `json:"textTransformation" tf:"text_transformation"`
}

type WafRegexMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	RegexMatchTuple []WafRegexMatchSetSpecRegexMatchTuple `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple,omitempty"`
}

type WafRegexMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafRegexMatchSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafRegexMatchSetList is a list of WafRegexMatchSets
type WafRegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafRegexMatchSet CRD objects
	Items []WafRegexMatchSet `json:"items,omitempty"`
}
