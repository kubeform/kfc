package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorAutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorAutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            MonitorAutoscaleSettingStatus `json:"status,omitempty"`
}

type MonitorAutoscaleSettingSpecNotificationEmail struct {
	// +optional
	CustomEmails []string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`
	// +optional
	SendToSubscriptionAdministrator bool `json:"sendToSubscriptionAdministrator,omitempty" tf:"send_to_subscription_administrator,omitempty"`
	// +optional
	SendToSubscriptionCoAdministrator bool `json:"sendToSubscriptionCoAdministrator,omitempty" tf:"send_to_subscription_co_administrator,omitempty"`
}

type MonitorAutoscaleSettingSpecNotificationWebhook struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	ServiceURI string            `json:"serviceURI" tf:"service_uri"`
}

type MonitorAutoscaleSettingSpecNotification struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Email []MonitorAutoscaleSettingSpecNotificationEmail `json:"email,omitempty" tf:"email,omitempty"`
	// +optional
	Webhook []MonitorAutoscaleSettingSpecNotificationWebhook `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileCapacity struct {
	Default int `json:"default" tf:"default"`
	Maximum int `json:"maximum" tf:"maximum"`
	Minimum int `json:"minimum" tf:"minimum"`
}

type MonitorAutoscaleSettingSpecProfileFixedDate struct {
	End   string `json:"end" tf:"end"`
	Start string `json:"start" tf:"start"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileRecurrence struct {
	Days []string `json:"days" tf:"days"`
	// +kubebuilder:validation:MaxItems=1
	Hours []int64 `json:"hours" tf:"hours"`
	// +kubebuilder:validation:MaxItems=1
	Minutes []int64 `json:"minutes" tf:"minutes"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type MonitorAutoscaleSettingSpecProfileRuleMetricTrigger struct {
	MetricName       string      `json:"metricName" tf:"metric_name"`
	MetricResourceID string      `json:"metricResourceID" tf:"metric_resource_id"`
	Operator         string      `json:"operator" tf:"operator"`
	Statistic        string      `json:"statistic" tf:"statistic"`
	Threshold        json.Number `json:"threshold" tf:"threshold"`
	TimeAggregation  string      `json:"timeAggregation" tf:"time_aggregation"`
	TimeGrain        string      `json:"timeGrain" tf:"time_grain"`
	TimeWindow       string      `json:"timeWindow" tf:"time_window"`
}

type MonitorAutoscaleSettingSpecProfileRuleScaleAction struct {
	Cooldown  string `json:"cooldown" tf:"cooldown"`
	Direction string `json:"direction" tf:"direction"`
	Type      string `json:"type" tf:"type"`
	Value     int    `json:"value" tf:"value"`
}

type MonitorAutoscaleSettingSpecProfileRule struct {
	// +kubebuilder:validation:MaxItems=1
	MetricTrigger []MonitorAutoscaleSettingSpecProfileRuleMetricTrigger `json:"metricTrigger" tf:"metric_trigger"`
	// +kubebuilder:validation:MaxItems=1
	ScaleAction []MonitorAutoscaleSettingSpecProfileRuleScaleAction `json:"scaleAction" tf:"scale_action"`
}

type MonitorAutoscaleSettingSpecProfile struct {
	// +kubebuilder:validation:MaxItems=1
	Capacity []MonitorAutoscaleSettingSpecProfileCapacity `json:"capacity" tf:"capacity"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedDate []MonitorAutoscaleSettingSpecProfileFixedDate `json:"fixedDate,omitempty" tf:"fixed_date,omitempty"`
	Name      string                                        `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Recurrence []MonitorAutoscaleSettingSpecProfileRecurrence `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Rule []MonitorAutoscaleSettingSpecProfileRule `json:"rule,omitempty" tf:"rule,omitempty"`
}

type MonitorAutoscaleSettingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Enabled  bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Notification []MonitorAutoscaleSettingSpecNotification `json:"notification,omitempty" tf:"notification,omitempty"`
	// +kubebuilder:validation:MaxItems=20
	Profile           []MonitorAutoscaleSettingSpecProfile `json:"profile" tf:"profile"`
	ResourceGroupName string                               `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags             map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TargetResourceID string            `json:"targetResourceID" tf:"target_resource_id"`
}

type MonitorAutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorAutoscaleSettingList is a list of MonitorAutoscaleSettings
type MonitorAutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorAutoscaleSetting CRD objects
	Items []MonitorAutoscaleSetting `json:"items,omitempty"`
}
