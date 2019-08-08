package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubeform.dev/kubeform/apis/azurerm"
)

var SchemeGroupVersion = schema.GroupVersion{Group: azurerm.GroupName, Version: "v1alpha1"}

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

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&CdnProfile{},
		&CdnProfileList{},

		&Image{},
		&ImageList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&NotificationHub{},
		&NotificationHubList{},

		&StorageShare{},
		&StorageShareList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&Route{},
		&RouteList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&StorageContainer{},
		&StorageContainerList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&DataFactory{},
		&DataFactoryList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&StorageTable{},
		&StorageTableList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&StorageBlob{},
		&StorageBlobList{},

		&AppService{},
		&AppServiceList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&Firewall{},
		&FirewallList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&BatchAccount{},
		&BatchAccountList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&LbProbe{},
		&LbProbeList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&DnsARecord{},
		&DnsARecordList{},

		&EventhubNamespace{},
		&EventhubNamespaceList{},

		&Eventhub{},
		&EventhubList{},

		&IotDps{},
		&IotDpsList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&BatchPool{},
		&BatchPoolList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&Snapshot{},
		&SnapshotList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&SearchService{},
		&SearchServiceList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&LbRule{},
		&LbRuleList{},

		&Lb{},
		&LbList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&MariadbServer{},
		&MariadbServerList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&ContainerService{},
		&ContainerServiceList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&Iothub{},
		&IothubList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&RedisCache{},
		&RedisCacheList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&SqlServer{},
		&SqlServerList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MysqlServer{},
		&MysqlServerList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&SharedImage{},
		&SharedImageList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&PublicIP{},
		&PublicIPList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DnsZone{},
		&DnsZoneList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&StorageAccount{},
		&StorageAccountList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&NotificationHubNamespace{},
		&NotificationHubNamespaceList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&ManagementLock{},
		&ManagementLockList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&KeyVault{},
		&KeyVaultList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&DevTestLab{},
		&DevTestLabList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&FunctionApp{},
		&FunctionAppList{},

		&ApiManagement{},
		&ApiManagementList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&SignalrService{},
		&SignalrServiceList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
