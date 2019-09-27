package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/aws"
)

var SchemeGroupVersion = schema.GroupVersion{Group: aws.GroupName, Version: "v1alpha1"}

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

		&SnsTopic{},
		&SnsTopicList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&EipAssociation{},
		&EipAssociationList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&Alb{},
		&AlbList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&Subnet{},
		&SubnetList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&TransferUser{},
		&TransferUserList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&DbInstance{},
		&DbInstanceList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&IotCertificate{},
		&IotCertificateList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&AmiCopy{},
		&AmiCopyList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&KmsAlias{},
		&KmsAliasList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&GlueJob{},
		&GlueJobList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&FlowLog{},
		&FlowLogList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&RouteTable{},
		&RouteTableList{},

		&PinpointApp{},
		&PinpointAppList{},

		&AlbListener{},
		&AlbListenerList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&EmrCluster{},
		&EmrClusterList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&BackupPlan{},
		&BackupPlanList{},

		&SwfDomain{},
		&SwfDomainList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&Codepipeline{},
		&CodepipelineList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&RdsCluster{},
		&RdsClusterList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesTemplate{},
		&SesTemplateList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&Route{},
		&RouteList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&Vpc{},
		&VpcList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&WafRule{},
		&WafRuleList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&DaxCluster{},
		&DaxClusterList{},

		&EksCluster{},
		&EksClusterList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&IamGroup{},
		&IamGroupList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&NatGateway{},
		&NatGatewayList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&KeyPair{},
		&KeyPairList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&IotThing{},
		&IotThingList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&Ami{},
		&AmiList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&SsmActivation{},
		&SsmActivationList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&WafWebACL{},
		&WafWebACLList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&DxConnection{},
		&DxConnectionList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&Route53Record{},
		&Route53RecordList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&Eip{},
		&EipList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&Instance{},
		&InstanceList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&EcsCluster{},
		&EcsClusterList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&IamRole{},
		&IamRoleList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&MqBroker{},
		&MqBrokerList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&LbListener{},
		&LbListenerList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&TransferServer{},
		&TransferServerList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SfnActivity{},
		&SfnActivityList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IamUser{},
		&IamUserList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DxLag{},
		&DxLagList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&KmsKey{},
		&KmsKeyList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&DxGateway{},
		&DxGatewayList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&NetworkACL{},
		&NetworkACLList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&EcsService{},
		&EcsServiceList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&SsmParameter{},
		&SsmParameterList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&IotPolicy{},
		&IotPolicyList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&BackupVault{},
		&BackupVaultList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&Lb{},
		&LbList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&IotThingType{},
		&IotThingTypeList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&Elb{},
		&ElbList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&S3Bucket{},
		&S3BucketList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SqsQueue{},
		&SqsQueueList{},

		&WafIpset{},
		&WafIpsetList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
