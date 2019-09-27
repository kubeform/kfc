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

		&BatchCertificate{},
		&BatchCertificateList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&IotDps{},
		&IotDpsList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&DataFactory{},
		&DataFactoryList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&PacketCapture{},
		&PacketCaptureList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&ApiManagement{},
		&ApiManagementList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&Iothub{},
		&IothubList{},

		&Route{},
		&RouteList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&Snapshot{},
		&SnapshotList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&LbRule{},
		&LbRuleList{},

		&NotificationHub{},
		&NotificationHubList{},

		&BatchAccount{},
		&BatchAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&RedisCache{},
		&RedisCacheList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&StorageAccount{},
		&StorageAccountList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&FunctionApp{},
		&FunctionAppList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&BatchPool{},
		&BatchPoolList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&ManagementLock{},
		&ManagementLockList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&ContainerService{},
		&ContainerServiceList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&AppService{},
		&AppServiceList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&MapsAccount{},
		&MapsAccountList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&SearchService{},
		&SearchServiceList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&RouteTable{},
		&RouteTableList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&Eventhub{},
		&EventhubList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&Image{},
		&ImageList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&DevTestLab{},
		&DevTestLabList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&StorageContainer{},
		&StorageContainerList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&KeyVault{},
		&KeyVaultList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&LbProbe{},
		&LbProbeList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&StorageQueue{},
		&StorageQueueList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&Subnet{},
		&SubnetList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&MysqlServer{},
		&MysqlServerList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&MariadbServer{},
		&MariadbServerList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&StorageTable{},
		&StorageTableList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&Firewall{},
		&FirewallList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&PublicIP{},
		&PublicIPList{},

		&StorageShare{},
		&StorageShareList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DnsZone{},
		&DnsZoneList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&SharedImage{},
		&SharedImageList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&CdnProfile{},
		&CdnProfileList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&Lb{},
		&LbList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
