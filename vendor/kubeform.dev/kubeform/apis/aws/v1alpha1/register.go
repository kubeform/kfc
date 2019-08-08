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

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&WafIpset{},
		&WafIpsetList{},

		&DbInstance{},
		&DbInstanceList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&KmsKey{},
		&KmsKeyList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&PinpointApp{},
		&PinpointAppList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&TransferServer{},
		&TransferServerList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&Instance{},
		&InstanceList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&Elb{},
		&ElbList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&Vpc{},
		&VpcList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&IamGroup{},
		&IamGroupList{},

		&IamRole{},
		&IamRoleList{},

		&IotCertificate{},
		&IotCertificateList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&AthenaWorkgroup{},
		&AthenaWorkgroupList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&IotThing{},
		&IotThingList{},

		&IotThingType{},
		&IotThingTypeList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&FlowLog{},
		&FlowLogList{},

		&IamUser{},
		&IamUserList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&Ec2TransitGatewayVpcAttachmentAccepter{},
		&Ec2TransitGatewayVpcAttachmentAccepterList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&TransferUser{},
		&TransferUserList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&IotPolicy{},
		&IotPolicyList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&GlacierVault{},
		&GlacierVaultList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&AppsyncFunction{},
		&AppsyncFunctionList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&WafRule{},
		&WafRuleList{},

		&Alb{},
		&AlbList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&KeyPair{},
		&KeyPairList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IamPolicy{},
		&IamPolicyList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&EbsDefaultKmsKey{},
		&EbsDefaultKmsKeyList{},

		&EbsEncryptionByDefault{},
		&EbsEncryptionByDefaultList{},

		&S3Bucket{},
		&S3BucketList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&EipAssociation{},
		&EipAssociationList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&MskCluster{},
		&MskClusterList{},

		&Route53Record{},
		&Route53RecordList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&GlobalacceleratorEndpointGroup{},
		&GlobalacceleratorEndpointGroupList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&Ami{},
		&AmiList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&KmsAlias{},
		&KmsAliasList{},

		&MqBroker{},
		&MqBrokerList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&SesEmailIdentity{},
		&SesEmailIdentityList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&EcsService{},
		&EcsServiceList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&RouteTable{},
		&RouteTableList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&AmiCopy{},
		&AmiCopyList{},

		&BackupVault{},
		&BackupVaultList{},

		&DatapipelinePipeline{},
		&DatapipelinePipelineList{},

		&DirectoryServiceLogSubscription{},
		&DirectoryServiceLogSubscriptionList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&AlbListener{},
		&AlbListenerList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&BackupPlan{},
		&BackupPlanList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&Route{},
		&RouteList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&LbSslNegotiationPolicy{},
		&LbSslNegotiationPolicyList{},

		&MskConfiguration{},
		&MskConfigurationList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&DxGateway{},
		&DxGatewayList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&LbListener{},
		&LbListenerList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&SesIdentityPolicy{},
		&SesIdentityPolicyList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&SfnActivity{},
		&SfnActivityList{},

		&Lb{},
		&LbList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&EcsCluster{},
		&EcsClusterList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&NatGateway{},
		&NatGatewayList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&SwfDomain{},
		&SwfDomainList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&EksCluster{},
		&EksClusterList{},

		&GlueJob{},
		&GlueJobList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&SsmParameter{},
		&SsmParameterList{},

		&DxConnection{},
		&DxConnectionList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&DaxCluster{},
		&DaxClusterList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&WafWebACL{},
		&WafWebACLList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&Codepipeline{},
		&CodepipelineList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&EmrCluster{},
		&EmrClusterList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&KmsGrant{},
		&KmsGrantList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&SqsQueue{},
		&SqsQueueList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SsmActivation{},
		&SsmActivationList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&SnsTopic{},
		&SnsTopicList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&SesTemplate{},
		&SesTemplateList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&NetworkACL{},
		&NetworkACLList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&Subnet{},
		&SubnetList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&InternetGateway{},
		&InternetGatewayList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&Eip{},
		&EipList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DxLag{},
		&DxLagList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&ShieldProtection{},
		&ShieldProtectionList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&RdsCluster{},
		&RdsClusterList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
