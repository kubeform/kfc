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

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeTargetSSLProxy{},
		&ComputeTargetSSLProxyList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SqlSSLCert{},
		&SqlSSLCertList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&DataflowJob{},
		&DataflowJobList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&DataprocJob{},
		&DataprocJobList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&Project{},
		&ProjectList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&StorageBucket{},
		&StorageBucketList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&SqlUser{},
		&SqlUserList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&Folder{},
		&FolderList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&BigtableTable{},
		&BigtableTableList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ProjectService{},
		&ProjectServiceList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ComputeImage{},
		&ComputeImageList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&ComputeSSLCertificate{},
		&ComputeSSLCertificateList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeSSLPolicy{},
		&ComputeSSLPolicyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
