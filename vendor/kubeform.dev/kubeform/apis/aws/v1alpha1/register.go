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

		&DxHostedPrivateVirtualInterfaceAccepter{},
		&DxHostedPrivateVirtualInterfaceAccepterList{},

		&Ec2TransitGateway{},
		&Ec2TransitGatewayList{},

		&ApiGatewayRestAPI{},
		&ApiGatewayRestAPIList{},

		&GuarddutyThreatintelset{},
		&GuarddutyThreatintelsetList{},

		&IotRoleAlias{},
		&IotRoleAliasList{},

		&WafXssMatchSet{},
		&WafXssMatchSetList{},

		&DxConnectionAssociation{},
		&DxConnectionAssociationList{},

		&GlobalacceleratorListener{},
		&GlobalacceleratorListenerList{},

		&IamUserLoginProfile{},
		&IamUserLoginProfileList{},

		&IotPolicy{},
		&IotPolicyList{},

		&AppsyncAPIKey{},
		&AppsyncAPIKeyList{},

		&AppsyncGraphqlAPI{},
		&AppsyncGraphqlAPIList{},

		&DbInstanceRoleAssociation{},
		&DbInstanceRoleAssociationList{},

		&Elb{},
		&ElbList{},

		&NeptuneCluster{},
		&NeptuneClusterList{},

		&ApiGatewayRequestValidator{},
		&ApiGatewayRequestValidatorList{},

		&RedshiftSecurityGroup{},
		&RedshiftSecurityGroupList{},

		&SesDomainMailFrom{},
		&SesDomainMailFromList{},

		&VpcIpv4CIDRBlockAssociation{},
		&VpcIpv4CIDRBlockAssociationList{},

		&PinpointGcmChannel{},
		&PinpointGcmChannelList{},

		&DaxCluster{},
		&DaxClusterList{},

		&GlueSecurityConfiguration{},
		&GlueSecurityConfigurationList{},

		&IotThingPrincipalAttachment{},
		&IotThingPrincipalAttachmentList{},

		&MediaStoreContainerPolicy{},
		&MediaStoreContainerPolicyList{},

		&ProxyProtocolPolicy{},
		&ProxyProtocolPolicyList{},

		&StoragegatewayWorkingStorage{},
		&StoragegatewayWorkingStorageList{},

		&VpnGatewayAttachment{},
		&VpnGatewayAttachmentList{},

		&RedshiftSnapshotCopyGrant{},
		&RedshiftSnapshotCopyGrantList{},

		&CloudwatchLogGroup{},
		&CloudwatchLogGroupList{},

		&CustomerGateway{},
		&CustomerGatewayList{},

		&DaxParameterGroup{},
		&DaxParameterGroupList{},

		&DbOptionGroup{},
		&DbOptionGroupList{},

		&DocdbClusterSnapshot{},
		&DocdbClusterSnapshotList{},

		&Ec2TransitGatewayVpcAttachment{},
		&Ec2TransitGatewayVpcAttachmentList{},

		&EipAssociation{},
		&EipAssociationList{},

		&AlbListener{},
		&AlbListenerList{},

		&ApiGatewayVpcLink{},
		&ApiGatewayVpcLinkList{},

		&AppCookieStickinessPolicy{},
		&AppCookieStickinessPolicyList{},

		&ConfigConfigurationAggregator{},
		&ConfigConfigurationAggregatorList{},

		&DxGateway{},
		&DxGatewayList{},

		&ElasticBeanstalkConfigurationTemplate{},
		&ElasticBeanstalkConfigurationTemplateList{},

		&GlueCatalogDatabase{},
		&GlueCatalogDatabaseList{},

		&SecurityhubProductSubscription{},
		&SecurityhubProductSubscriptionList{},

		&ElasticsearchDomain{},
		&ElasticsearchDomainList{},

		&Instance{},
		&InstanceList{},

		&NatGateway{},
		&NatGatewayList{},

		&OpsworksHaproxyLayer{},
		&OpsworksHaproxyLayerList{},

		&WafByteMatchSet{},
		&WafByteMatchSetList{},

		&PinpointSmsChannel{},
		&PinpointSmsChannelList{},

		&AppsyncDatasource{},
		&AppsyncDatasourceList{},

		&DocdbClusterParameterGroup{},
		&DocdbClusterParameterGroupList{},

		&IamGroupMembership{},
		&IamGroupMembershipList{},

		&WafregionalRateBasedRule{},
		&WafregionalRateBasedRuleList{},

		&Cloud9EnvironmentEc2{},
		&Cloud9EnvironmentEc2List{},

		&InternetGateway{},
		&InternetGatewayList{},

		&NeptuneClusterSnapshot{},
		&NeptuneClusterSnapshotList{},

		&NeptuneSubnetGroup{},
		&NeptuneSubnetGroupList{},

		&RamResourceShare{},
		&RamResourceShareList{},

		&ServiceDiscoveryPublicDNSNamespace{},
		&ServiceDiscoveryPublicDNSNamespaceList{},

		&AthenaNamedQuery{},
		&AthenaNamedQueryList{},

		&AutoscalingSchedule{},
		&AutoscalingScheduleList{},

		&DynamodbTable{},
		&DynamodbTableList{},

		&IamGroup{},
		&IamGroupList{},

		&IamUserGroupMembership{},
		&IamUserGroupMembershipList{},

		&InspectorResourceGroup{},
		&InspectorResourceGroupList{},

		&Route53ZoneAssociation{},
		&Route53ZoneAssociationList{},

		&CloudfrontPublicKey{},
		&CloudfrontPublicKeyList{},

		&CloudwatchMetricAlarm{},
		&CloudwatchMetricAlarmList{},

		&IotTopicRule{},
		&IotTopicRuleList{},

		&RdsClusterParameterGroup{},
		&RdsClusterParameterGroupList{},

		&CloudfrontDistribution{},
		&CloudfrontDistributionList{},

		&SagemakerNotebookInstance{},
		&SagemakerNotebookInstanceList{},

		&SesDomainDkim{},
		&SesDomainDkimList{},

		&ElasticacheReplicationGroup{},
		&ElasticacheReplicationGroupList{},

		&SpotDatafeedSubscription{},
		&SpotDatafeedSubscriptionList{},

		&DefaultVpcDHCPOptions{},
		&DefaultVpcDHCPOptionsList{},

		&OpsworksNodejsAppLayer{},
		&OpsworksNodejsAppLayerList{},

		&ConfigConfigurationRecorderStatus_{},
		&ConfigConfigurationRecorderStatus_List{},

		&GlacierVault{},
		&GlacierVaultList{},

		&IamRolePolicyAttachment{},
		&IamRolePolicyAttachmentList{},

		&BatchComputeEnvironment{},
		&BatchComputeEnvironmentList{},

		&ApiGatewayDocumentationPart{},
		&ApiGatewayDocumentationPartList{},

		&EmrSecurityConfiguration{},
		&EmrSecurityConfigurationList{},

		&SecretsmanagerSecret{},
		&SecretsmanagerSecretList{},

		&LbTargetGroupAttachment{},
		&LbTargetGroupAttachmentList{},

		&ApiGatewayModel{},
		&ApiGatewayModelList{},

		&CloudwatchLogMetricFilter{},
		&CloudwatchLogMetricFilterList{},

		&DocdbCluster{},
		&DocdbClusterList{},

		&IamAccountPasswordPolicy{},
		&IamAccountPasswordPolicyList{},

		&LoadBalancerListenerPolicy{},
		&LoadBalancerListenerPolicyList{},

		&WafGeoMatchSet{},
		&WafGeoMatchSetList{},

		&PinpointAdmChannel{},
		&PinpointAdmChannelList{},

		&ElbAttachment{},
		&ElbAttachmentList{},

		&GlueTrigger{},
		&GlueTriggerList{},

		&RedshiftEventSubscription{},
		&RedshiftEventSubscriptionList{},

		&Route53Record{},
		&Route53RecordList{},

		&SesReceiptFilter{},
		&SesReceiptFilterList{},

		&SwfDomain{},
		&SwfDomainList{},

		&AlbListenerCertificate{},
		&AlbListenerCertificateList{},

		&CodebuildProject{},
		&CodebuildProjectList{},

		&LoadBalancerPolicy{},
		&LoadBalancerPolicyList{},

		&RdsCluster{},
		&RdsClusterList{},

		&TransferServer{},
		&TransferServerList{},

		&ApiGatewayAPIKey{},
		&ApiGatewayAPIKeyList{},

		&Ec2Fleet{},
		&Ec2FleetList{},

		&CodebuildWebhook{},
		&CodebuildWebhookList{},

		&GlacierVaultLock{},
		&GlacierVaultLockList{},

		&SnsTopicPolicy{},
		&SnsTopicPolicyList{},

		&OpsworksStaticWebLayer{},
		&OpsworksStaticWebLayerList{},

		&OpsworksMysqlLayer{},
		&OpsworksMysqlLayerList{},

		&RdsClusterInstance{},
		&RdsClusterInstanceList{},

		&LbCookieStickinessPolicy{},
		&LbCookieStickinessPolicyList{},

		&S3BucketNotification{},
		&S3BucketNotificationList{},

		&TransferUser{},
		&TransferUserList{},

		&ApiGatewayAuthorizer{},
		&ApiGatewayAuthorizerList{},

		&CloudwatchLogResourcePolicy{},
		&CloudwatchLogResourcePolicyList{},

		&ElasticBeanstalkApplication{},
		&ElasticBeanstalkApplicationList{},

		&GuarddutyIpset{},
		&GuarddutyIpsetList{},

		&NeptuneClusterParameterGroup{},
		&NeptuneClusterParameterGroupList{},

		&NetworkInterfaceAttachment{},
		&NetworkInterfaceAttachmentList{},

		&RamResourceAssociation{},
		&RamResourceAssociationList{},

		&CognitoUserPoolDomain{},
		&CognitoUserPoolDomainList{},

		&Ec2ClientVPNNetworkAssociation{},
		&Ec2ClientVPNNetworkAssociationList{},

		&GuarddutyDetector{},
		&GuarddutyDetectorList{},

		&CloudwatchEventRule{},
		&CloudwatchEventRuleList{},

		&SsmAssociation{},
		&SsmAssociationList{},

		&LambdaFunction{},
		&LambdaFunctionList{},

		&OpsworksCustomLayer{},
		&OpsworksCustomLayerList{},

		&SsmActivation{},
		&SsmActivationList{},

		&DatasyncLocationEfs{},
		&DatasyncLocationEfsList{},

		&DevicefarmProject{},
		&DevicefarmProjectList{},

		&IamGroupPolicy{},
		&IamGroupPolicyList{},

		&KmsCiphertext{},
		&KmsCiphertextList{},

		&SesReceiptRule{},
		&SesReceiptRuleList{},

		&ApiGatewayDocumentationVersion{},
		&ApiGatewayDocumentationVersionList{},

		&MacieMemberAccountAssociation{},
		&MacieMemberAccountAssociationList{},

		&OpsworksMemcachedLayer{},
		&OpsworksMemcachedLayerList{},

		&SnsTopic{},
		&SnsTopicList{},

		&DefaultSubnet{},
		&DefaultSubnetList{},

		&ApiGatewayMethodSettings{},
		&ApiGatewayMethodSettingsList{},

		&DbSnapshot{},
		&DbSnapshotList{},

		&IotPolicyAttachment{},
		&IotPolicyAttachmentList{},

		&AppmeshVirtualService{},
		&AppmeshVirtualServiceList{},

		&IotThing{},
		&IotThingList{},

		&SecurityGroupRule{},
		&SecurityGroupRuleList{},

		&CognitoIdentityProvider{},
		&CognitoIdentityProviderList{},

		&ElastictranscoderPipeline{},
		&ElastictranscoderPipelineList{},

		&IamAccessKey{},
		&IamAccessKeyList{},

		&CodecommitTrigger{},
		&CodecommitTriggerList{},

		&DmsReplicationInstance{},
		&DmsReplicationInstanceList{},

		&EmrInstanceGroup{},
		&EmrInstanceGroupList{},

		&RedshiftParameterGroup{},
		&RedshiftParameterGroupList{},

		&SsmParameter{},
		&SsmParameterList{},

		&WafRateBasedRule{},
		&WafRateBasedRuleList{},

		&BudgetsBudget{},
		&BudgetsBudgetList{},

		&RouteTableAssociation{},
		&RouteTableAssociationList{},

		&LbListenerCertificate{},
		&LbListenerCertificateList{},

		&DxGatewayAssociation{},
		&DxGatewayAssociationList{},

		&DefaultSecurityGroup{},
		&DefaultSecurityGroupList{},

		&EbsSnapshot{},
		&EbsSnapshotList{},

		&EcrRepositoryPolicy{},
		&EcrRepositoryPolicyList{},

		&IamUserPolicy{},
		&IamUserPolicyList{},

		&AppautoscalingTarget{},
		&AppautoscalingTargetList{},

		&RouteTable{},
		&RouteTableList{},

		&WafregionalXssMatchSet{},
		&WafregionalXssMatchSetList{},

		&WorklinkFleet{},
		&WorklinkFleetList{},

		&CognitoResourceServer{},
		&CognitoResourceServerList{},

		&DmsCertificate{},
		&DmsCertificateList{},

		&DxConnection{},
		&DxConnectionList{},

		&NetworkInterfaceSgAttachment{},
		&NetworkInterfaceSgAttachmentList{},

		&TransferSSHKey{},
		&TransferSSHKeyList{},

		&DbSubnetGroup{},
		&DbSubnetGroupList{},

		&LoadBalancerBackendServerPolicy{},
		&LoadBalancerBackendServerPolicyList{},

		&OpsworksRailsAppLayer{},
		&OpsworksRailsAppLayerList{},

		&VpcPeeringConnectionAccepter{},
		&VpcPeeringConnectionAccepterList{},

		&SfnActivity{},
		&SfnActivityList{},

		&WafregionalRuleGroup{},
		&WafregionalRuleGroupList{},

		&SecurityhubAccount{},
		&SecurityhubAccountList{},

		&GameliftFleet{},
		&GameliftFleetList{},

		&S3AccountPublicAccessBlock{},
		&S3AccountPublicAccessBlockList{},

		&StoragegatewaySmbFileShare{},
		&StoragegatewaySmbFileShareList{},

		&Lb{},
		&LbList{},

		&ApiGatewayIntegration{},
		&ApiGatewayIntegrationList{},

		&ApiGatewayMethodResponse{},
		&ApiGatewayMethodResponseList{},

		&AppmeshVirtualNode{},
		&AppmeshVirtualNodeList{},

		&WafSizeConstraintSet{},
		&WafSizeConstraintSetList{},

		&CloudformationStackSetInstance{},
		&CloudformationStackSetInstanceList{},

		&EmrCluster{},
		&EmrClusterList{},

		&Route53ResolverRule{},
		&Route53ResolverRuleList{},

		&ApiGatewayClientCertificate{},
		&ApiGatewayClientCertificateList{},

		&CloudhsmV2Cluster{},
		&CloudhsmV2ClusterList{},

		&GlueCrawler{},
		&GlueCrawlerList{},

		&Route{},
		&RouteList{},

		&AmiCopy{},
		&AmiCopyList{},

		&DocdbClusterInstance{},
		&DocdbClusterInstanceList{},

		&DefaultNetworkACL{},
		&DefaultNetworkACLList{},

		&SecurityGroup{},
		&SecurityGroupList{},

		&PinpointApnsSandboxChannel{},
		&PinpointApnsSandboxChannelList{},

		&AppmeshMesh{},
		&AppmeshMeshList{},

		&ConfigConfigurationRecorder{},
		&ConfigConfigurationRecorderList{},

		&DbSecurityGroup{},
		&DbSecurityGroupList{},

		&Ec2CapacityReservation{},
		&Ec2CapacityReservationList{},

		&OpsworksJavaAppLayer{},
		&OpsworksJavaAppLayerList{},

		&OrganizationsOrganization{},
		&OrganizationsOrganizationList{},

		&StoragegatewayCache{},
		&StoragegatewayCacheList{},

		&CognitoIdentityPool{},
		&CognitoIdentityPoolList{},

		&CodedeployDeploymentGroup{},
		&CodedeployDeploymentGroupList{},

		&ElasticacheCluster{},
		&ElasticacheClusterList{},

		&KmsKey{},
		&KmsKeyList{},

		&LaunchConfiguration{},
		&LaunchConfigurationList{},

		&SqsQueue{},
		&SqsQueueList{},

		&BackupPlan{},
		&BackupPlanList{},

		&GameliftAlias{},
		&GameliftAliasList{},

		&OpsworksUserProfile{},
		&OpsworksUserProfileList{},

		&WafregionalRegexPatternSet{},
		&WafregionalRegexPatternSetList{},

		&AppmeshRoute{},
		&AppmeshRouteList{},

		&ElasticBeanstalkEnvironment{},
		&ElasticBeanstalkEnvironmentList{},

		&AppautoscalingScheduledAction{},
		&AppautoscalingScheduledActionList{},

		&DbEventSubscription{},
		&DbEventSubscriptionList{},

		&GlobalacceleratorAccelerator{},
		&GlobalacceleratorAcceleratorList{},

		&IamRole{},
		&IamRoleList{},

		&Route53ResolverEndpoint{},
		&Route53ResolverEndpointList{},

		&PinpointApnsVoipSandboxChannel{},
		&PinpointApnsVoipSandboxChannelList{},

		&KmsExternalKey{},
		&KmsExternalKeyList{},

		&SsmMaintenanceWindowTask{},
		&SsmMaintenanceWindowTaskList{},

		&SsmResourceDataSync{},
		&SsmResourceDataSyncList{},

		&AutoscalingNotification{},
		&AutoscalingNotificationList{},

		&CodepipelineWebhook{},
		&CodepipelineWebhookList{},

		&NeptuneClusterInstance{},
		&NeptuneClusterInstanceList{},

		&SecurityhubStandardsSubscription{},
		&SecurityhubStandardsSubscriptionList{},

		&SqsQueuePolicy{},
		&SqsQueuePolicyList{},

		&VpcPeeringConnection{},
		&VpcPeeringConnectionList{},

		&AutoscalingPolicy{},
		&AutoscalingPolicyList{},

		&ElasticacheParameterGroup{},
		&ElasticacheParameterGroupList{},

		&SpotFleetRequest{},
		&SpotFleetRequestList{},

		&SnsTopicSubscription{},
		&SnsTopicSubscriptionList{},

		&IamPolicy{},
		&IamPolicyList{},

		&IamSamlProvider{},
		&IamSamlProviderList{},

		&MediaStoreContainer{},
		&MediaStoreContainerList{},

		&NeptuneParameterGroup{},
		&NeptuneParameterGroupList{},

		&OpsworksInstance{},
		&OpsworksInstanceList{},

		&WafregionalByteMatchSet{},
		&WafregionalByteMatchSetList{},

		&PinpointEventStream{},
		&PinpointEventStreamList{},

		&AppmeshVirtualRouter{},
		&AppmeshVirtualRouterList{},

		&DbParameterGroup{},
		&DbParameterGroupList{},

		&DmsReplicationSubnetGroup{},
		&DmsReplicationSubnetGroupList{},

		&GuarddutyMember{},
		&GuarddutyMemberList{},

		&Route53DelegationSet{},
		&Route53DelegationSetList{},

		&StoragegatewayUploadBuffer{},
		&StoragegatewayUploadBufferList{},

		&VpcEndpointServiceAllowedPrincipal{},
		&VpcEndpointServiceAllowedPrincipalList{},

		&ApiGatewayAccount{},
		&ApiGatewayAccountList{},

		&CloudwatchEventPermission{},
		&CloudwatchEventPermissionList{},

		&DynamodbTableItem{},
		&DynamodbTableItemList{},

		&KinesisStream{},
		&KinesisStreamList{},

		&NetworkACL{},
		&NetworkACLList{},

		&WafIpset{},
		&WafIpsetList{},

		&ApiGatewayUsagePlan{},
		&ApiGatewayUsagePlanList{},

		&PinpointApnsVoipChannel{},
		&PinpointApnsVoipChannelList{},

		&DirectoryServiceConditionalForwarder{},
		&DirectoryServiceConditionalForwarderList{},

		&OpsworksPermission{},
		&OpsworksPermissionList{},

		&DefaultRouteTable{},
		&DefaultRouteTableList{},

		&SfnStateMachine{},
		&SfnStateMachineList{},

		&Ami{},
		&AmiList{},

		&CloudwatchLogStream{},
		&CloudwatchLogStreamList{},

		&IamUserSSHKey{},
		&IamUserSSHKeyList{},

		&IotThingType{},
		&IotThingTypeList{},

		&Route53QueryLog{},
		&Route53QueryLogList{},

		&VpcEndpointConnectionNotification{},
		&VpcEndpointConnectionNotificationList{},

		&WafRegexPatternSet{},
		&WafRegexPatternSetList{},

		&Vpc{},
		&VpcList{},

		&PinpointApnsChannel{},
		&PinpointApnsChannelList{},

		&LambdaLayerVersion{},
		&LambdaLayerVersionList{},

		&OrganizationsPolicyAttachment{},
		&OrganizationsPolicyAttachmentList{},

		&ConfigDeliveryChannel{},
		&ConfigDeliveryChannelList{},

		&CodedeployDeploymentConfig{},
		&CodedeployDeploymentConfigList{},

		&PinpointEmailChannel{},
		&PinpointEmailChannelList{},

		&AlbListenerRule{},
		&AlbListenerRuleList{},

		&DbInstance{},
		&DbInstanceList{},

		&IamGroupPolicyAttachment{},
		&IamGroupPolicyAttachmentList{},

		&IamRolePolicy{},
		&IamRolePolicyList{},

		&VpcEndpointService{},
		&VpcEndpointServiceList{},

		&Cloudtrail{},
		&CloudtrailList{},

		&SagemakerNotebookInstanceLifecycleConfiguration{},
		&SagemakerNotebookInstanceLifecycleConfigurationList{},

		&SesActiveReceiptRuleSet{},
		&SesActiveReceiptRuleSetList{},

		&SpotInstanceRequest{},
		&SpotInstanceRequestList{},

		&VpnConnectionRoute{},
		&VpnConnectionRouteList{},

		&Ec2TransitGatewayRouteTablePropagation{},
		&Ec2TransitGatewayRouteTablePropagationList{},

		&FlowLog{},
		&FlowLogList{},

		&GuarddutyInviteAccepter{},
		&GuarddutyInviteAccepterList{},

		&MainRouteTableAssociation{},
		&MainRouteTableAssociationList{},

		&SecretsmanagerSecretVersion{},
		&SecretsmanagerSecretVersionList{},

		&WafWebACL{},
		&WafWebACLList{},

		&Ec2ClientVPNEndpoint{},
		&Ec2ClientVPNEndpointList{},

		&Ec2TransitGatewayRoute{},
		&Ec2TransitGatewayRouteList{},

		&RedshiftSubnetGroup{},
		&RedshiftSubnetGroupList{},

		&VpnGatewayRoutePropagation{},
		&VpnGatewayRoutePropagationList{},

		&WafRule{},
		&WafRuleList{},

		&WafregionalGeoMatchSet{},
		&WafregionalGeoMatchSetList{},

		&WorklinkWebsiteCertificateAuthorityAssociation{},
		&WorklinkWebsiteCertificateAuthorityAssociationList{},

		&CloudhsmV2Hsm{},
		&CloudhsmV2HsmList{},

		&Ec2TransitGatewayRouteTable{},
		&Ec2TransitGatewayRouteTableList{},

		&ElastictranscoderPreset{},
		&ElastictranscoderPresetList{},

		&IamServiceLinkedRole{},
		&IamServiceLinkedRoleList{},

		&LicensemanagerAssociation{},
		&LicensemanagerAssociationList{},

		&SagemakerModel{},
		&SagemakerModelList{},

		&VpcEndpoint{},
		&VpcEndpointList{},

		&EgressOnlyInternetGateway{},
		&EgressOnlyInternetGatewayList{},

		&Eip{},
		&EipList{},

		&SimpledbDomain{},
		&SimpledbDomainList{},

		&PinpointBaiduChannel{},
		&PinpointBaiduChannelList{},

		&AppautoscalingPolicy{},
		&AppautoscalingPolicyList{},

		&GlueJob{},
		&GlueJobList{},

		&OpsworksApplication{},
		&OpsworksApplicationList{},

		&ApiGatewayIntegrationResponse{},
		&ApiGatewayIntegrationResponseList{},

		&OpsworksPhpAppLayer{},
		&OpsworksPhpAppLayerList{},

		&SesConfigurationSet{},
		&SesConfigurationSetList{},

		&VpcDHCPOptions{},
		&VpcDHCPOptionsList{},

		&AmiFromInstance{},
		&AmiFromInstanceList{},

		&AutoscalingAttachment{},
		&AutoscalingAttachmentList{},

		&BackupVault{},
		&BackupVaultList{},

		&IamUserPolicyAttachment{},
		&IamUserPolicyAttachmentList{},

		&WafregionalWebACL{},
		&WafregionalWebACLList{},

		&EcrRepository{},
		&EcrRepositoryList{},

		&IotCertificate{},
		&IotCertificateList{},

		&KeyPair{},
		&KeyPairList{},

		&DmsReplicationTask{},
		&DmsReplicationTaskList{},

		&RdsClusterEndpoint{},
		&RdsClusterEndpointList{},

		&ResourcegroupsGroup{},
		&ResourcegroupsGroupList{},

		&CodedeployApp{},
		&CodedeployAppList{},

		&LambdaPermission{},
		&LambdaPermissionList{},

		&S3BucketInventory{},
		&S3BucketInventoryList{},

		&Subnet{},
		&SubnetList{},

		&WafregionalRegexMatchSet{},
		&WafregionalRegexMatchSetList{},

		&AmiLaunchPermission{},
		&AmiLaunchPermissionList{},

		&DatasyncAgent{},
		&DatasyncAgentList{},

		&DbClusterSnapshot{},
		&DbClusterSnapshotList{},

		&ElasticsearchDomainPolicy{},
		&ElasticsearchDomainPolicyList{},

		&OpsworksGangliaLayer{},
		&OpsworksGangliaLayerList{},

		&WafRuleGroup{},
		&WafRuleGroupList{},

		&GameliftBuild{},
		&GameliftBuildList{},

		&LaunchTemplate{},
		&LaunchTemplateList{},

		&LbSSLNegotiationPolicy{},
		&LbSSLNegotiationPolicyList{},

		&RedshiftCluster{},
		&RedshiftClusterList{},

		&SagemakerEndpoint{},
		&SagemakerEndpointList{},

		&WafregionalRule{},
		&WafregionalRuleList{},

		&ApiGatewayDomainName{},
		&ApiGatewayDomainNameList{},

		&AutoscalingLifecycleHook{},
		&AutoscalingLifecycleHookList{},

		&CloudwatchLogSubscriptionFilter{},
		&CloudwatchLogSubscriptionFilterList{},

		&Codepipeline{},
		&CodepipelineList{},

		&EbsSnapshotCopy{},
		&EbsSnapshotCopyList{},

		&EcsTaskDefinition{},
		&EcsTaskDefinitionList{},

		&VpcEndpointRouteTableAssociation{},
		&VpcEndpointRouteTableAssociationList{},

		&WafregionalSizeConstraintSet{},
		&WafregionalSizeConstraintSetList{},

		&CloudfrontOriginAccessIdentity{},
		&CloudfrontOriginAccessIdentityList{},

		&DynamodbGlobalTable{},
		&DynamodbGlobalTableList{},

		&Ec2TransitGatewayRouteTableAssociation{},
		&Ec2TransitGatewayRouteTableAssociationList{},

		&EfsMountTarget{},
		&EfsMountTargetList{},

		&LightsailDomain{},
		&LightsailDomainList{},

		&SsmMaintenanceWindow{},
		&SsmMaintenanceWindowList{},

		&AcmCertificateValidation{},
		&AcmCertificateValidationList{},

		&ApiGatewayResource{},
		&ApiGatewayResourceList{},

		&IamUser{},
		&IamUserList{},

		&ServiceDiscoveryService{},
		&ServiceDiscoveryServiceList{},

		&ApiGatewayStage{},
		&ApiGatewayStageList{},

		&CloudformationStack{},
		&CloudformationStackList{},

		&DirectoryServiceDirectory{},
		&DirectoryServiceDirectoryList{},

		&NetworkACLRule{},
		&NetworkACLRuleList{},

		&OrganizationsPolicy{},
		&OrganizationsPolicyList{},

		&CurReportDefinition{},
		&CurReportDefinitionList{},

		&LightsailKeyPair{},
		&LightsailKeyPairList{},

		&S3BucketObject{},
		&S3BucketObjectList{},

		&StoragegatewayCachedIscsiVolume{},
		&StoragegatewayCachedIscsiVolumeList{},

		&ApiGatewayBasePathMapping{},
		&ApiGatewayBasePathMappingList{},

		&CognitoUserPoolClient{},
		&CognitoUserPoolClientList{},

		&DxHostedPrivateVirtualInterface{},
		&DxHostedPrivateVirtualInterfaceList{},

		&DxLag{},
		&DxLagList{},

		&ElasticBeanstalkApplicationVersion{},
		&ElasticBeanstalkApplicationVersionList{},

		&LbListenerRule{},
		&LbListenerRuleList{},

		&LbTargetGroup{},
		&LbTargetGroupList{},

		&ApiGatewayUsagePlanKey{},
		&ApiGatewayUsagePlanKeyList{},

		&AthenaDatabase{},
		&AthenaDatabaseList{},

		&LicensemanagerLicenseConfiguration{},
		&LicensemanagerLicenseConfigurationList{},

		&WafregionalWebACLAssociation{},
		&WafregionalWebACLAssociationList{},

		&ConfigConfigRule{},
		&ConfigConfigRuleList{},

		&IamServerCertificate{},
		&IamServerCertificateList{},

		&KinesisAnalyticsApplication{},
		&KinesisAnalyticsApplicationList{},

		&RamPrincipalAssociation{},
		&RamPrincipalAssociationList{},

		&DxHostedPublicVirtualInterface{},
		&DxHostedPublicVirtualInterfaceList{},

		&KinesisFirehoseDeliveryStream{},
		&KinesisFirehoseDeliveryStreamList{},

		&KmsAlias{},
		&KmsAliasList{},

		&S3BucketPolicy{},
		&S3BucketPolicyList{},

		&Route53Zone{},
		&Route53ZoneList{},

		&StoragegatewayNfsFileShare{},
		&StoragegatewayNfsFileShareList{},

		&VolumeAttachment{},
		&VolumeAttachmentList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&BatchJobDefinition{},
		&BatchJobDefinitionList{},

		&EksCluster{},
		&EksClusterList{},

		&KmsGrant{},
		&KmsGrantList{},

		&LightsailInstance{},
		&LightsailInstanceList{},

		&LightsailStaticIP{},
		&LightsailStaticIPList{},

		&NeptuneEventSubscription{},
		&NeptuneEventSubscriptionList{},

		&OrganizationsAccount{},
		&OrganizationsAccountList{},

		&DatasyncLocationNfs{},
		&DatasyncLocationNfsList{},

		&GlueClassifier{},
		&GlueClassifierList{},

		&PlacementGroup{},
		&PlacementGroupList{},

		&SesTemplate{},
		&SesTemplateList{},

		&ServiceDiscoveryPrivateDNSNamespace{},
		&ServiceDiscoveryPrivateDNSNamespaceList{},

		&CloudwatchDashboard{},
		&CloudwatchDashboardList{},

		&DatasyncLocationS3{},
		&DatasyncLocationS3List{},

		&InspectorAssessmentTemplate{},
		&InspectorAssessmentTemplateList{},

		&AcmpcaCertificateAuthority{},
		&AcmpcaCertificateAuthorityList{},

		&ApiGatewayDeployment{},
		&ApiGatewayDeploymentList{},

		&DxPrivateVirtualInterface{},
		&DxPrivateVirtualInterfaceList{},

		&SsmDocument{},
		&SsmDocumentList{},

		&Alb{},
		&AlbList{},

		&CloudwatchLogDestination{},
		&CloudwatchLogDestinationList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&S3BucketPublicAccessBlock{},
		&S3BucketPublicAccessBlockList{},

		&GlueConnection{},
		&GlueConnectionList{},

		&Route53HealthCheck{},
		&Route53HealthCheckList{},

		&AppsyncResolver{},
		&AppsyncResolverList{},

		&OpsworksStack{},
		&OpsworksStackList{},

		&SnapshotCreateVolumePermission{},
		&SnapshotCreateVolumePermissionList{},

		&SnsSmsPreferences{},
		&SnsSmsPreferencesList{},

		&WafSQLInjectionMatchSet{},
		&WafSQLInjectionMatchSetList{},

		&DaxSubnetGroup{},
		&DaxSubnetGroupList{},

		&InspectorAssessmentTarget{},
		&InspectorAssessmentTargetList{},

		&ServicecatalogPortfolio{},
		&ServicecatalogPortfolioList{},

		&AcmCertificate{},
		&AcmCertificateList{},

		&EcrLifecyclePolicy{},
		&EcrLifecyclePolicyList{},

		&IamPolicyAttachment{},
		&IamPolicyAttachmentList{},

		&SesDomainIdentity{},
		&SesDomainIdentityList{},

		&ApiGatewayGatewayResponse{},
		&ApiGatewayGatewayResponseList{},

		&DxBGPPeer{},
		&DxBGPPeerList{},

		&DxPublicVirtualInterface{},
		&DxPublicVirtualInterfaceList{},

		&MediaPackageChannel{},
		&MediaPackageChannelList{},

		&S3Bucket{},
		&S3BucketList{},

		&SsmPatchBaseline{},
		&SsmPatchBaselineList{},

		&WafregionalSQLInjectionMatchSet{},
		&WafregionalSQLInjectionMatchSetList{},

		&CodecommitRepository{},
		&CodecommitRepositoryList{},

		&DatasyncTask{},
		&DatasyncTaskList{},

		&DxHostedPublicVirtualInterfaceAccepter{},
		&DxHostedPublicVirtualInterfaceAccepterList{},

		&SesReceiptRuleSet{},
		&SesReceiptRuleSetList{},

		&GameliftGameSessionQueue{},
		&GameliftGameSessionQueueList{},

		&LambdaEventSourceMapping{},
		&LambdaEventSourceMappingList{},

		&CognitoUserPool{},
		&CognitoUserPoolList{},

		&SsmMaintenanceWindowTarget{},
		&SsmMaintenanceWindowTargetList{},

		&StoragegatewayGateway{},
		&StoragegatewayGatewayList{},

		&ConfigAggregateAuthorization{},
		&ConfigAggregateAuthorizationList{},

		&CognitoIdentityPoolRolesAttachment{},
		&CognitoIdentityPoolRolesAttachmentList{},

		&SnsPlatformApplication{},
		&SnsPlatformApplicationList{},

		&XraySamplingRule{},
		&XraySamplingRuleList{},

		&AutoscalingGroup{},
		&AutoscalingGroupList{},

		&CloudwatchLogDestinationPolicy{},
		&CloudwatchLogDestinationPolicyList{},

		&VpcDHCPOptionsAssociation{},
		&VpcDHCPOptionsAssociationList{},

		&VpnConnection{},
		&VpnConnectionList{},

		&BackupSelection{},
		&BackupSelectionList{},

		&DxGatewayAssociationProposal{},
		&DxGatewayAssociationProposalList{},

		&EfsFileSystem{},
		&EfsFileSystemList{},

		&RdsGlobalCluster{},
		&RdsGlobalClusterList{},

		&DefaultVpc{},
		&DefaultVpcList{},

		&SesEventDestination{},
		&SesEventDestinationList{},

		&AlbTargetGroup{},
		&AlbTargetGroupList{},

		&DmsEndpoint{},
		&DmsEndpointList{},

		&EcsService{},
		&EcsServiceList{},

		&SesDomainIdentityVerification{},
		&SesDomainIdentityVerificationList{},

		&AlbTargetGroupAttachment{},
		&AlbTargetGroupAttachmentList{},

		&DocdbSubnetGroup{},
		&DocdbSubnetGroupList{},

		&ElasticacheSubnetGroup{},
		&ElasticacheSubnetGroupList{},

		&GlueCatalogTable{},
		&GlueCatalogTableList{},

		&OpsworksRdsDbInstance{},
		&OpsworksRdsDbInstanceList{},

		&ServiceDiscoveryHTTPNamespace{},
		&ServiceDiscoveryHTTPNamespaceList{},

		&VpcEndpointSubnetAssociation{},
		&VpcEndpointSubnetAssociationList{},

		&LbListener{},
		&LbListenerList{},

		&EbsVolume{},
		&EbsVolumeList{},

		&PinpointApp{},
		&PinpointAppList{},

		&ElasticacheSecurityGroup{},
		&ElasticacheSecurityGroupList{},

		&MacieS3BucketAssociation{},
		&MacieS3BucketAssociationList{},

		&CloudwatchEventTarget{},
		&CloudwatchEventTargetList{},

		&MqConfiguration{},
		&MqConfigurationList{},

		&SagemakerEndpointConfiguration{},
		&SagemakerEndpointConfigurationList{},

		&CognitoUserGroup{},
		&CognitoUserGroupList{},

		&EcsCluster{},
		&EcsClusterList{},

		&IamOpenidConnectProvider{},
		&IamOpenidConnectProviderList{},

		&S3BucketMetric{},
		&S3BucketMetricList{},

		&DlmLifecyclePolicy{},
		&DlmLifecyclePolicyList{},

		&MqBroker{},
		&MqBrokerList{},

		&Route53ResolverRuleAssociation{},
		&Route53ResolverRuleAssociationList{},

		&VpcPeeringConnectionOptions{},
		&VpcPeeringConnectionOptionsList{},

		&IamAccountAlias{},
		&IamAccountAliasList{},

		&LightsailStaticIPAttachment{},
		&LightsailStaticIPAttachmentList{},

		&SesIdentityNotificationTopic{},
		&SesIdentityNotificationTopicList{},

		&WafregionalIpset{},
		&WafregionalIpsetList{},

		&ApiGatewayMethod{},
		&ApiGatewayMethodList{},

		&CloudformationStackSet{},
		&CloudformationStackSetList{},

		&IamInstanceProfile{},
		&IamInstanceProfileList{},

		&LambdaAlias{},
		&LambdaAliasList{},

		&OrganizationsOrganizationalUnit{},
		&OrganizationsOrganizationalUnitList{},

		&SsmPatchGroup{},
		&SsmPatchGroupList{},

		&WafRegexMatchSet{},
		&WafRegexMatchSetList{},

		&BatchJobQueue{},
		&BatchJobQueueList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
