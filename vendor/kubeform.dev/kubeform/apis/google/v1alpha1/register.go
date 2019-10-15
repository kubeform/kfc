package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/google"
)

var SchemeGroupVersion = schema.GroupVersion{Group: google.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&BigtableTable{},
		&BigtableTableList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&StorageBucket{},
		&StorageBucketList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&Project{},
		&ProjectList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&Folder{},
		&FolderList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&DataflowJob{},
		&DataflowJobList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ProjectService{},
		&ProjectServiceList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
