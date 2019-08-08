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

		&StorageBucketIamMember{},
		&StorageBucketIamMemberList{},

		&StorageBucketIamBinding{},
		&StorageBucketIamBindingList{},

		&ServiceAccountIamMember{},
		&ServiceAccountIamMemberList{},

		&PubsubTopicIamPolicy{},
		&PubsubTopicIamPolicyList{},

		&ProjectOrganizationPolicy{},
		&ProjectOrganizationPolicyList{},

		&SourcerepoRepository{},
		&SourcerepoRepositoryList{},

		&SpannerDatabaseIamBinding{},
		&SpannerDatabaseIamBindingList{},

		&ResourceManagerLien{},
		&ResourceManagerLienList{},

		&ContainerNodePool{},
		&ContainerNodePoolList{},

		&FolderOrganizationPolicy{},
		&FolderOrganizationPolicyList{},

		&SpannerInstanceIamMember{},
		&SpannerInstanceIamMemberList{},

		&ProjectIamMember{},
		&ProjectIamMemberList{},

		&Project{},
		&ProjectList{},

		&KmsKeyRingIamBinding{},
		&KmsKeyRingIamBindingList{},

		&ComputeTargetSslProxy{},
		&ComputeTargetSslProxyList{},

		&ComputeHTTPHealthCheck{},
		&ComputeHTTPHealthCheckList{},

		&MonitoringUptimeCheckConfig{},
		&MonitoringUptimeCheckConfigList{},

		&BillingAccountIamBinding{},
		&BillingAccountIamBindingList{},

		&OrganizationIamBinding{},
		&OrganizationIamBindingList{},

		&LoggingOrganizationExclusion{},
		&LoggingOrganizationExclusionList{},

		&KmsKeyRingIamMember{},
		&KmsKeyRingIamMemberList{},

		&LoggingFolderSink{},
		&LoggingFolderSinkList{},

		&ComputeFirewall{},
		&ComputeFirewallList{},

		&StorageObjectACL{},
		&StorageObjectACLList{},

		&ComputeRouterPeer{},
		&ComputeRouterPeerList{},

		&ComputeHealthCheck{},
		&ComputeHealthCheckList{},

		&ComputeTargetTcpProxy{},
		&ComputeTargetTcpProxyList{},

		&ComputeDisk{},
		&ComputeDiskList{},

		&MonitoringAlertPolicy{},
		&MonitoringAlertPolicyList{},

		&ComputeInstance{},
		&ComputeInstanceList{},

		&ComputeNetworkPeering{},
		&ComputeNetworkPeeringList{},

		&DnsManagedZone{},
		&DnsManagedZoneList{},

		&ComputeRegionDisk{},
		&ComputeRegionDiskList{},

		&ComputeImage{},
		&ComputeImageList{},

		&StorageBucketACL{},
		&StorageBucketACLList{},

		&ServiceAccountIamPolicy{},
		&ServiceAccountIamPolicyList{},

		&ContainerCluster{},
		&ContainerClusterList{},

		&ComputeRegionInstanceGroupManager{},
		&ComputeRegionInstanceGroupManagerList{},

		&OrganizationIamCustomRole{},
		&OrganizationIamCustomRoleList{},

		&StorageDefaultObjectAccessControl{},
		&StorageDefaultObjectAccessControlList{},

		&LoggingProjectExclusion{},
		&LoggingProjectExclusionList{},

		&OrganizationIamMember{},
		&OrganizationIamMemberList{},

		&CloudfunctionsFunction{},
		&CloudfunctionsFunctionList{},

		&BigtableInstance{},
		&BigtableInstanceList{},

		&ComputeAttachedDisk{},
		&ComputeAttachedDiskList{},

		&ComposerEnvironment{},
		&ComposerEnvironmentList{},

		&PubsubSubscriptionIamBinding{},
		&PubsubSubscriptionIamBindingList{},

		&PubsubSubscriptionIamPolicy{},
		&PubsubSubscriptionIamPolicyList{},

		&DnsRecordSet{},
		&DnsRecordSetList{},

		&ComputeURLMap{},
		&ComputeURLMapList{},

		&BigtableTable{},
		&BigtableTableList{},

		&FolderIamBinding{},
		&FolderIamBindingList{},

		&ProjectUsageExportBucket{},
		&ProjectUsageExportBucketList{},

		&SqlDatabaseInstance{},
		&SqlDatabaseInstanceList{},

		&DataprocJob{},
		&DataprocJobList{},

		&LoggingOrganizationSink{},
		&LoggingOrganizationSinkList{},

		&BillingAccountIamMember{},
		&BillingAccountIamMemberList{},

		&BillingAccountIamPolicy{},
		&BillingAccountIamPolicyList{},

		&ComputeInterconnectAttachment{},
		&ComputeInterconnectAttachmentList{},

		&ComputeProjectMetadataItem{},
		&ComputeProjectMetadataItemList{},

		&ProjectIamPolicy{},
		&ProjectIamPolicyList{},

		&RedisInstance{},
		&RedisInstanceList{},

		&ContainerAnalysisNote{},
		&ContainerAnalysisNoteList{},

		&SqlUser{},
		&SqlUserList{},

		&ComputeSecurityPolicy{},
		&ComputeSecurityPolicyList{},

		&BinaryAuthorizationPolicy{},
		&BinaryAuthorizationPolicyList{},

		&ComputeForwardingRule{},
		&ComputeForwardingRuleList{},

		&ServiceAccount{},
		&ServiceAccountList{},

		&ComputeTargetHTTPProxy{},
		&ComputeTargetHTTPProxyList{},

		&KmsCryptoKey{},
		&KmsCryptoKeyList{},

		&ComputeInstanceGroupManager{},
		&ComputeInstanceGroupManagerList{},

		&EndpointsService{},
		&EndpointsServiceList{},

		&ComputeSubnetworkIamMember{},
		&ComputeSubnetworkIamMemberList{},

		&ComputeAutoscaler{},
		&ComputeAutoscalerList{},

		&ComputeTargetPool{},
		&ComputeTargetPoolList{},

		&DataflowJob{},
		&DataflowJobList{},

		&AppEngineApplication{},
		&AppEngineApplicationList{},

		&LoggingBillingAccountExclusion{},
		&LoggingBillingAccountExclusionList{},

		&ComputeBackendService{},
		&ComputeBackendServiceList{},

		&PubsubTopicIamMember{},
		&PubsubTopicIamMemberList{},

		&FolderIamPolicy{},
		&FolderIamPolicyList{},

		&ProjectIamBinding{},
		&ProjectIamBindingList{},

		&ComputeInstanceGroup{},
		&ComputeInstanceGroupList{},

		&OrganizationPolicy{},
		&OrganizationPolicyList{},

		&ComputeGlobalAddress{},
		&ComputeGlobalAddressList{},

		&FolderIamMember{},
		&FolderIamMemberList{},

		&PubsubSubscription{},
		&PubsubSubscriptionList{},

		&BigqueryTable{},
		&BigqueryTableList{},

		&KmsKeyRing{},
		&KmsKeyRingList{},

		&ComputeRegionBackendService{},
		&ComputeRegionBackendServiceList{},

		&PubsubSubscriptionIamMember{},
		&PubsubSubscriptionIamMemberList{},

		&ComputeRouterInterface{},
		&ComputeRouterInterfaceList{},

		&LoggingProjectSink{},
		&LoggingProjectSinkList{},

		&ProjectIamCustomRole{},
		&ProjectIamCustomRoleList{},

		&SpannerInstance{},
		&SpannerInstanceList{},

		&ComputeSubnetwork{},
		&ComputeSubnetworkList{},

		&ComputeRouterNAT{},
		&ComputeRouterNATList{},

		&KmsCryptoKeyIamMember{},
		&KmsCryptoKeyIamMemberList{},

		&ComputeSubnetworkIamBinding{},
		&ComputeSubnetworkIamBindingList{},

		&PubsubTopic{},
		&PubsubTopicList{},

		&StorageDefaultObjectACL{},
		&StorageDefaultObjectACLList{},

		&MonitoringNotificationChannel{},
		&MonitoringNotificationChannelList{},

		&FilestoreInstance{},
		&FilestoreInstanceList{},

		&CloudbuildTrigger{},
		&CloudbuildTriggerList{},

		&SpannerDatabaseIamPolicy{},
		&SpannerDatabaseIamPolicyList{},

		&ComputeSharedVpcServiceProject{},
		&ComputeSharedVpcServiceProjectList{},

		&ComputeInstanceFromTemplate{},
		&ComputeInstanceFromTemplateList{},

		&ServiceAccountKey{},
		&ServiceAccountKeyList{},

		&ComputeRoute{},
		&ComputeRouteList{},

		&ServiceAccountIamBinding{},
		&ServiceAccountIamBindingList{},

		&RuntimeconfigConfig{},
		&RuntimeconfigConfigList{},

		&StorageNotification{},
		&StorageNotificationList{},

		&BigqueryDataset{},
		&BigqueryDatasetList{},

		&StorageBucket{},
		&StorageBucketList{},

		&DataprocCluster{},
		&DataprocClusterList{},

		&KmsCryptoKeyIamBinding{},
		&KmsCryptoKeyIamBindingList{},

		&ComputeAddress{},
		&ComputeAddressList{},

		&ProjectService{},
		&ProjectServiceList{},

		&SpannerInstanceIamBinding{},
		&SpannerInstanceIamBindingList{},

		&SpannerInstanceIamPolicy{},
		&SpannerInstanceIamPolicyList{},

		&ComputeSslPolicy{},
		&ComputeSslPolicyList{},

		&ComputeHTTPSHealthCheck{},
		&ComputeHTTPSHealthCheckList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlSslCert{},
		&SqlSslCertList{},

		&PubsubTopicIamBinding{},
		&PubsubTopicIamBindingList{},

		&ComputeVPNGateway{},
		&ComputeVPNGatewayList{},

		&ComputeSharedVpcHostProject{},
		&ComputeSharedVpcHostProjectList{},

		&ComputeBackendBucket{},
		&ComputeBackendBucketList{},

		&CloudiotRegistry{},
		&CloudiotRegistryList{},

		&ComputeSubnetworkIamPolicy{},
		&ComputeSubnetworkIamPolicyList{},

		&StorageBucketIamPolicy{},
		&StorageBucketIamPolicyList{},

		&ComputeRegionAutoscaler{},
		&ComputeRegionAutoscalerList{},

		&ComputeRouter{},
		&ComputeRouterList{},

		&ComputeSslCertificate{},
		&ComputeSslCertificateList{},

		&StorageObjectAccessControl{},
		&StorageObjectAccessControlList{},

		&ComputeNetwork{},
		&ComputeNetworkList{},

		&SpannerDatabase{},
		&SpannerDatabaseList{},

		&StorageBucketObject{},
		&StorageBucketObjectList{},

		&Folder{},
		&FolderList{},

		&ComputeVPNTunnel{},
		&ComputeVPNTunnelList{},

		&OrganizationIamPolicy{},
		&OrganizationIamPolicyList{},

		&ProjectServices{},
		&ProjectServicesList{},

		&ComputeProjectMetadata{},
		&ComputeProjectMetadataList{},

		&BinaryAuthorizationAttestor{},
		&BinaryAuthorizationAttestorList{},

		&ComputeGlobalForwardingRule{},
		&ComputeGlobalForwardingRuleList{},

		&LoggingFolderExclusion{},
		&LoggingFolderExclusionList{},

		&SpannerDatabaseIamMember{},
		&SpannerDatabaseIamMemberList{},

		&RuntimeconfigVariable{},
		&RuntimeconfigVariableList{},

		&ComputeSnapshot{},
		&ComputeSnapshotList{},

		&KmsKeyRingIamPolicy{},
		&KmsKeyRingIamPolicyList{},

		&MonitoringGroup{},
		&MonitoringGroupList{},

		&ComputeInstanceTemplate{},
		&ComputeInstanceTemplateList{},

		&LoggingBillingAccountSink{},
		&LoggingBillingAccountSinkList{},

		&ComputeTargetHTTPSProxy{},
		&ComputeTargetHTTPSProxyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
