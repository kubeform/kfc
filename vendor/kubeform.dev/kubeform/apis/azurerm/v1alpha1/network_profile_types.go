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

type NetworkProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkProfileSpec   `json:"spec,omitempty"`
	Status            NetworkProfileStatus `json:"status,omitempty"`
}

type NetworkProfileSpecContainerNetworkInterfaceIpConfiguration struct {
	Name     string `json:"name" tf:"name"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

type NetworkProfileSpecContainerNetworkInterface struct {
	IpConfiguration []NetworkProfileSpecContainerNetworkInterfaceIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Name            string                                                       `json:"name" tf:"name"`
}

type NetworkProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +kubebuilder:validation:MaxItems=1
	ContainerNetworkInterface []NetworkProfileSpecContainerNetworkInterface `json:"containerNetworkInterface" tf:"container_network_interface"`
	Location                  string                                        `json:"location" tf:"location"`
	Name                      string                                        `json:"name" tf:"name"`
	ResourceGroupName         string                                        `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkProfileList is a list of NetworkProfiles
type NetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkProfile CRD objects
	Items []NetworkProfile `json:"items,omitempty"`
}