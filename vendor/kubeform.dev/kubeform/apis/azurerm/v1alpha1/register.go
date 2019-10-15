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

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&Route{},
		&RouteList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&MysqlServer{},
		&MysqlServerList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&Snapshot{},
		&SnapshotList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&StorageBlob{},
		&StorageBlobList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&StorageTable{},
		&StorageTableList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&DnsARecord{},
		&DnsARecordList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&Iothub{},
		&IothubList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&StorageShare{},
		&StorageShareList{},

		&AppService{},
		&AppServiceList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&Lb{},
		&LbList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&RouteTable{},
		&RouteTableList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&BatchAccount{},
		&BatchAccountList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&SqlServer{},
		&SqlServerList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&StorageContainer{},
		&StorageContainerList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&MariadbServer{},
		&MariadbServerList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&MapsAccount{},
		&MapsAccountList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&RedisCache{},
		&RedisCacheList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&Firewall{},
		&FirewallList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&Subnet{},
		&SubnetList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&DataFactory{},
		&DataFactoryList{},

		&DevTestLab{},
		&DevTestLabList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&SharedImage{},
		&SharedImageList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&BatchPool{},
		&BatchPoolList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&IotDps{},
		&IotDpsList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&NotificationHub{},
		&NotificationHubList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&SearchService{},
		&SearchServiceList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&DnsZone{},
		&DnsZoneList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&LbRule{},
		&LbRuleList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&Eventhub{},
		&EventhubList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&Image{},
		&ImageList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&FunctionApp{},
		&FunctionAppList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApiManagement{},
		&ApiManagementList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&LbProbe{},
		&LbProbeList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&KeyVault{},
		&KeyVaultList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&StorageAccount{},
		&StorageAccountList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&PublicIP{},
		&PublicIPList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
