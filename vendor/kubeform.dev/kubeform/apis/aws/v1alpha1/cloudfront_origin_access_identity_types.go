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
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type CloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontOriginAccessIdentitySpec   `json:"spec,omitempty"`
	Status            CloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

type CloudfrontOriginAccessIdentitySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CallerReference string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`
	// +optional
	CloudfrontAccessIdentityPath string `json:"cloudfrontAccessIdentityPath,omitempty" tf:"cloudfront_access_identity_path,omitempty"`
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +optional
	IamArn string `json:"iamArn,omitempty" tf:"iam_arn,omitempty"`
	// +optional
	S3CanonicalUserID string `json:"s3CanonicalUserID,omitempty" tf:"s3_canonical_user_id,omitempty"`
}

type CloudfrontOriginAccessIdentityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudfrontOriginAccessIdentitySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfrontOriginAccessIdentityList is a list of CloudfrontOriginAccessIdentitys
type CloudfrontOriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfrontOriginAccessIdentity CRD objects
	Items []CloudfrontOriginAccessIdentity `json:"items,omitempty"`
}
