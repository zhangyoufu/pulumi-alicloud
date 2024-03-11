// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    /// <summary>
    /// Information about RDS database exclusive agent and its usage, see [What is RDS DB Proxy](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-modifydbproxy).
    /// 
    /// &gt; **NOTE:** Available since v1.193.0.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultZones = AliCloud.Rds.GetZones.Invoke(new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.6",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Rds.Instance("defaultInstance", new()
    ///     {
    ///         Engine = "MySQL",
    ///         EngineVersion = "5.7",
    ///         InstanceType = "rds.mysql.c1.large",
    ///         InstanceStorage = 20,
    ///         InstanceChargeType = "Postpaid",
    ///         InstanceName = name,
    ///         VswitchId = defaultSwitch.Id,
    ///         DbInstanceStorageType = "local_ssd",
    ///     });
    /// 
    ///     var defaultReadOnlyInstance = new AliCloud.Rds.ReadOnlyInstance("defaultReadOnlyInstance", new()
    ///     {
    ///         ZoneId = defaultInstance.ZoneId,
    ///         MasterDbInstanceId = defaultInstance.Id,
    ///         EngineVersion = defaultInstance.EngineVersion,
    ///         InstanceStorage = defaultInstance.InstanceStorage,
    ///         InstanceType = defaultInstance.InstanceType,
    ///         InstanceName = $"{name}readonly",
    ///         VswitchId = defaultSwitch.Id,
    ///     });
    /// 
    ///     var defaultRdsDbProxy = new AliCloud.Rds.RdsDbProxy("defaultRdsDbProxy", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         InstanceNetworkType = "VPC",
    ///         VpcId = defaultInstance.VpcId,
    ///         VswitchId = defaultInstance.VswitchId,
    ///         DbProxyInstanceNum = 2,
    ///         DbProxyConnectionPrefix = "example",
    ///         DbProxyConnectStringPort = 3306,
    ///         DbProxyEndpointReadWriteMode = "ReadWrite",
    ///         ReadOnlyInstanceMaxDelayTime = 90,
    ///         DbProxyFeatures = "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:1",
    ///         ReadOnlyInstanceDistributionType = "Custom",
    ///         ReadOnlyInstanceWeights = new[]
    ///         {
    ///             new AliCloud.Rds.Inputs.RdsDbProxyReadOnlyInstanceWeightArgs
    ///             {
    ///                 InstanceId = defaultInstance.Id,
    ///                 Weight = "100",
    ///             },
    ///             new AliCloud.Rds.Inputs.RdsDbProxyReadOnlyInstanceWeightArgs
    ///             {
    ///                 InstanceId = defaultReadOnlyInstance.Id,
    ///                 Weight = "500",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// &gt; **NOTE:** Resource `alicloud.rds.RdsDbProxy` should be created after `alicloud.rds.ReadOnlyInstance`, so the `depends_on` statement is necessary.
    /// 
    /// ## Import
    /// 
    /// RDS database proxy feature can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rds/rdsDbProxy:RdsDbProxy example abc12345678
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rds/rdsDbProxy:RdsDbProxy")]
    public partial class RdsDbProxy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The port number that is associated with the proxy endpoint.
        /// </summary>
        [Output("dbProxyConnectStringPort")]
        public Output<int> DbProxyConnectStringPort { get; private set; } = null!;

        /// <summary>
        /// The new prefix of the proxy endpoint. Enter a prefix.
        /// </summary>
        [Output("dbProxyConnectionPrefix")]
        public Output<string> DbProxyConnectionPrefix { get; private set; } = null!;

        /// <summary>
        /// Connection instance string.
        /// </summary>
        [Output("dbProxyConnectionString")]
        public Output<string> DbProxyConnectionString { get; private set; } = null!;

        /// <summary>
        /// Remarks of agent terminal.
        /// </summary>
        [Output("dbProxyEndpointAliases")]
        public Output<string> DbProxyEndpointAliases { get; private set; } = null!;

        /// <summary>
        /// Proxy connection address ID.
        /// </summary>
        [Output("dbProxyEndpointId")]
        public Output<string> DbProxyEndpointId { get; private set; } = null!;

        /// <summary>
        /// The read and write attributes of the proxy terminal. Valid values:
        /// - ReadWrite: The proxy terminal connects to the primary instance and can receive both read and write requests.
        /// - ReadOnly: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
        /// 
        /// &gt; **NOTE:** Note This setting causes your instance to restart. Proceed with caution.
        /// </summary>
        [Output("dbProxyEndpointReadWriteMode")]
        public Output<string> DbProxyEndpointReadWriteMode { get; private set; } = null!;

        /// <summary>
        /// The features that you want to enable for the proxy endpoint. If you specify more than one feature, separate the features with semicolons (;). Format: Feature 1:Status;Feature 2:Status;.... Do not add a semicolon (;) at the end of the last value. Valid feature values:
        /// - ReadWriteSpliting: read/write splitting.
        /// - ConnectionPersist: connection pooling.
        /// - TransactionReadSqlRouteOptimizeStatus: transaction splitting.
        /// Valid status values:
        /// - 1: enabled.
        /// - 0: disabled.
        /// 
        /// &gt; **NOTE:** Note You must specify this parameter only when the read/write splitting feature is enabled.
        /// </summary>
        [Output("dbProxyFeatures")]
        public Output<string> DbProxyFeatures { get; private set; } = null!;

        /// <summary>
        /// The number of proxy instances that are enabled. Valid values: 1 to 60.
        /// </summary>
        [Output("dbProxyInstanceNum")]
        public Output<int> DbProxyInstanceNum { get; private set; } = null!;

        /// <summary>
        /// The SSL configuration setting that you want to apply on the instance. Valid values:
        /// - Close: disables SSL encryption.
        /// - Open: enables SSL encryption or modifies the endpoint that requires SSL encryption.
        /// - Update: updates the validity period of the SSL certificate.
        /// </summary>
        [Output("dbProxySslEnabled")]
        public Output<string> DbProxySslEnabled { get; private set; } = null!;

        /// <summary>
        /// The point in time at which you want to apply the new database proxy settings. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Output("effectiveSpecificTime")]
        public Output<string> EffectiveSpecificTime { get; private set; } = null!;

        /// <summary>
        /// When modifying the number of proxy instances,The time when you want to apply the new database proxy settings.Valid values:
        /// - Immediate: ApsaraDB RDS immediately applies the new settings.
        /// - MaintainTime: ApsaraDB RDS applies the new settings during the maintenance window that you specified. For more information, see Modify the maintenance window.
        /// - SpecificTime: ApsaraDB RDS applies the new settings at a specified point in time.
        /// 
        /// &gt; **NOTE:** Note If you set the EffectiveTime parameter to SpecificTime, you must specify the EffectiveSpecificTime parameter.
        /// </summary>
        [Output("effectiveTime")]
        public Output<string> EffectiveTime { get; private set; } = null!;

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The network type of the instance. Set the value to VPC.
        /// </summary>
        [Output("instanceNetworkType")]
        public Output<string> InstanceNetworkType { get; private set; } = null!;

        /// <summary>
        /// Network type of proxy connection address.
        /// </summary>
        [Output("netType")]
        public Output<string> NetType { get; private set; } = null!;

        /// <summary>
        /// The policy that is used to allocate read weights. Valid values:
        /// - Standard: ApsaraDB RDS automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
        /// - Custom: You must manually allocate read weights to the instance and its read-only instances.
        /// 
        /// &gt; **NOTE:** Note If you set the ReadOnlyInstanceDistributionType parameter to Custom, you must specify the ReadOnlyInstanceWeight parameter.
        /// </summary>
        [Output("readOnlyInstanceDistributionType")]
        public Output<string> ReadOnlyInstanceDistributionType { get; private set; } = null!;

        /// <summary>
        /// The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, ApsaraDB RDS no longer forwards read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained. Unit: seconds. Valid values: 0 to 3600.
        /// 
        /// &gt; **NOTE:** Note If the instance runs PostgreSQL, you can enable only the read/write splitting feature, which is specified by ReadWriteSpliting.
        /// </summary>
        [Output("readOnlyInstanceMaxDelayTime")]
        public Output<int> ReadOnlyInstanceMaxDelayTime { get; private set; } = null!;

        /// <summary>
        /// A list of the read weights of the instance and its read-only instances.  It contains two sub-fields(instance_id and weight). Read weights increase in increments of 100, and the maximum read weight is 10000. See `read_only_instance_weight` below.
        /// </summary>
        [Output("readOnlyInstanceWeights")]
        public Output<ImmutableArray<Outputs.RdsDbProxyReadOnlyInstanceWeight>> ReadOnlyInstanceWeights { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The time when the certificate expires.
        /// </summary>
        [Output("sslExpiredTime")]
        public Output<string> SslExpiredTime { get; private set; } = null!;

        /// <summary>
        /// The point in time at which you want to upgrade the database proxy version of the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Output("switchTime")]
        public Output<string?> SwitchTime { get; private set; } = null!;

        /// <summary>
        /// The time when you want to upgrade the database proxy version of the instance. Valid values:
        /// - MaintainTime: ApsaraDB RDS performs the upgrade during the maintenance window that you specified. This is the default value. For more information, see Modify the maintenance window.
        /// - Immediate: ApsaraDB RDS immediately performs the upgrade.
        /// - SpecificTime: ApsaraDB RDS performs the upgrade at a specified point in time.
        /// </summary>
        [Output("upgradeTime")]
        public Output<string?> UpgradeTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual private cloud (VPC) to which the instance belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The ID of the vSwitch that is associated with the specified VPC.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a RdsDbProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdsDbProxy(string name, RdsDbProxyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsDbProxy:RdsDbProxy", name, args ?? new RdsDbProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdsDbProxy(string name, Input<string> id, RdsDbProxyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rds/rdsDbProxy:RdsDbProxy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RdsDbProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdsDbProxy Get(string name, Input<string> id, RdsDbProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new RdsDbProxy(name, id, state, options);
        }
    }

    public sealed class RdsDbProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port number that is associated with the proxy endpoint.
        /// </summary>
        [Input("dbProxyConnectStringPort")]
        public Input<int>? DbProxyConnectStringPort { get; set; }

        /// <summary>
        /// The new prefix of the proxy endpoint. Enter a prefix.
        /// </summary>
        [Input("dbProxyConnectionPrefix")]
        public Input<string>? DbProxyConnectionPrefix { get; set; }

        /// <summary>
        /// The read and write attributes of the proxy terminal. Valid values:
        /// - ReadWrite: The proxy terminal connects to the primary instance and can receive both read and write requests.
        /// - ReadOnly: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
        /// 
        /// &gt; **NOTE:** Note This setting causes your instance to restart. Proceed with caution.
        /// </summary>
        [Input("dbProxyEndpointReadWriteMode")]
        public Input<string>? DbProxyEndpointReadWriteMode { get; set; }

        /// <summary>
        /// The features that you want to enable for the proxy endpoint. If you specify more than one feature, separate the features with semicolons (;). Format: Feature 1:Status;Feature 2:Status;.... Do not add a semicolon (;) at the end of the last value. Valid feature values:
        /// - ReadWriteSpliting: read/write splitting.
        /// - ConnectionPersist: connection pooling.
        /// - TransactionReadSqlRouteOptimizeStatus: transaction splitting.
        /// Valid status values:
        /// - 1: enabled.
        /// - 0: disabled.
        /// 
        /// &gt; **NOTE:** Note You must specify this parameter only when the read/write splitting feature is enabled.
        /// </summary>
        [Input("dbProxyFeatures")]
        public Input<string>? DbProxyFeatures { get; set; }

        /// <summary>
        /// The number of proxy instances that are enabled. Valid values: 1 to 60.
        /// </summary>
        [Input("dbProxyInstanceNum", required: true)]
        public Input<int> DbProxyInstanceNum { get; set; } = null!;

        /// <summary>
        /// The SSL configuration setting that you want to apply on the instance. Valid values:
        /// - Close: disables SSL encryption.
        /// - Open: enables SSL encryption or modifies the endpoint that requires SSL encryption.
        /// - Update: updates the validity period of the SSL certificate.
        /// </summary>
        [Input("dbProxySslEnabled")]
        public Input<string>? DbProxySslEnabled { get; set; }

        /// <summary>
        /// The point in time at which you want to apply the new database proxy settings. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("effectiveSpecificTime")]
        public Input<string>? EffectiveSpecificTime { get; set; }

        /// <summary>
        /// When modifying the number of proxy instances,The time when you want to apply the new database proxy settings.Valid values:
        /// - Immediate: ApsaraDB RDS immediately applies the new settings.
        /// - MaintainTime: ApsaraDB RDS applies the new settings during the maintenance window that you specified. For more information, see Modify the maintenance window.
        /// - SpecificTime: ApsaraDB RDS applies the new settings at a specified point in time.
        /// 
        /// &gt; **NOTE:** Note If you set the EffectiveTime parameter to SpecificTime, you must specify the EffectiveSpecificTime parameter.
        /// </summary>
        [Input("effectiveTime")]
        public Input<string>? EffectiveTime { get; set; }

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The network type of the instance. Set the value to VPC.
        /// </summary>
        [Input("instanceNetworkType", required: true)]
        public Input<string> InstanceNetworkType { get; set; } = null!;

        /// <summary>
        /// The policy that is used to allocate read weights. Valid values:
        /// - Standard: ApsaraDB RDS automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
        /// - Custom: You must manually allocate read weights to the instance and its read-only instances.
        /// 
        /// &gt; **NOTE:** Note If you set the ReadOnlyInstanceDistributionType parameter to Custom, you must specify the ReadOnlyInstanceWeight parameter.
        /// </summary>
        [Input("readOnlyInstanceDistributionType")]
        public Input<string>? ReadOnlyInstanceDistributionType { get; set; }

        /// <summary>
        /// The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, ApsaraDB RDS no longer forwards read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained. Unit: seconds. Valid values: 0 to 3600.
        /// 
        /// &gt; **NOTE:** Note If the instance runs PostgreSQL, you can enable only the read/write splitting feature, which is specified by ReadWriteSpliting.
        /// </summary>
        [Input("readOnlyInstanceMaxDelayTime")]
        public Input<int>? ReadOnlyInstanceMaxDelayTime { get; set; }

        [Input("readOnlyInstanceWeights")]
        private InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightArgs>? _readOnlyInstanceWeights;

        /// <summary>
        /// A list of the read weights of the instance and its read-only instances.  It contains two sub-fields(instance_id and weight). Read weights increase in increments of 100, and the maximum read weight is 10000. See `read_only_instance_weight` below.
        /// </summary>
        public InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightArgs> ReadOnlyInstanceWeights
        {
            get => _readOnlyInstanceWeights ?? (_readOnlyInstanceWeights = new InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightArgs>());
            set => _readOnlyInstanceWeights = value;
        }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The point in time at which you want to upgrade the database proxy version of the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("switchTime")]
        public Input<string>? SwitchTime { get; set; }

        /// <summary>
        /// The time when you want to upgrade the database proxy version of the instance. Valid values:
        /// - MaintainTime: ApsaraDB RDS performs the upgrade during the maintenance window that you specified. This is the default value. For more information, see Modify the maintenance window.
        /// - Immediate: ApsaraDB RDS immediately performs the upgrade.
        /// - SpecificTime: ApsaraDB RDS performs the upgrade at a specified point in time.
        /// </summary>
        [Input("upgradeTime")]
        public Input<string>? UpgradeTime { get; set; }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) to which the instance belongs.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The ID of the vSwitch that is associated with the specified VPC.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public RdsDbProxyArgs()
        {
        }
        public static new RdsDbProxyArgs Empty => new RdsDbProxyArgs();
    }

    public sealed class RdsDbProxyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port number that is associated with the proxy endpoint.
        /// </summary>
        [Input("dbProxyConnectStringPort")]
        public Input<int>? DbProxyConnectStringPort { get; set; }

        /// <summary>
        /// The new prefix of the proxy endpoint. Enter a prefix.
        /// </summary>
        [Input("dbProxyConnectionPrefix")]
        public Input<string>? DbProxyConnectionPrefix { get; set; }

        /// <summary>
        /// Connection instance string.
        /// </summary>
        [Input("dbProxyConnectionString")]
        public Input<string>? DbProxyConnectionString { get; set; }

        /// <summary>
        /// Remarks of agent terminal.
        /// </summary>
        [Input("dbProxyEndpointAliases")]
        public Input<string>? DbProxyEndpointAliases { get; set; }

        /// <summary>
        /// Proxy connection address ID.
        /// </summary>
        [Input("dbProxyEndpointId")]
        public Input<string>? DbProxyEndpointId { get; set; }

        /// <summary>
        /// The read and write attributes of the proxy terminal. Valid values:
        /// - ReadWrite: The proxy terminal connects to the primary instance and can receive both read and write requests.
        /// - ReadOnly: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
        /// 
        /// &gt; **NOTE:** Note This setting causes your instance to restart. Proceed with caution.
        /// </summary>
        [Input("dbProxyEndpointReadWriteMode")]
        public Input<string>? DbProxyEndpointReadWriteMode { get; set; }

        /// <summary>
        /// The features that you want to enable for the proxy endpoint. If you specify more than one feature, separate the features with semicolons (;). Format: Feature 1:Status;Feature 2:Status;.... Do not add a semicolon (;) at the end of the last value. Valid feature values:
        /// - ReadWriteSpliting: read/write splitting.
        /// - ConnectionPersist: connection pooling.
        /// - TransactionReadSqlRouteOptimizeStatus: transaction splitting.
        /// Valid status values:
        /// - 1: enabled.
        /// - 0: disabled.
        /// 
        /// &gt; **NOTE:** Note You must specify this parameter only when the read/write splitting feature is enabled.
        /// </summary>
        [Input("dbProxyFeatures")]
        public Input<string>? DbProxyFeatures { get; set; }

        /// <summary>
        /// The number of proxy instances that are enabled. Valid values: 1 to 60.
        /// </summary>
        [Input("dbProxyInstanceNum")]
        public Input<int>? DbProxyInstanceNum { get; set; }

        /// <summary>
        /// The SSL configuration setting that you want to apply on the instance. Valid values:
        /// - Close: disables SSL encryption.
        /// - Open: enables SSL encryption or modifies the endpoint that requires SSL encryption.
        /// - Update: updates the validity period of the SSL certificate.
        /// </summary>
        [Input("dbProxySslEnabled")]
        public Input<string>? DbProxySslEnabled { get; set; }

        /// <summary>
        /// The point in time at which you want to apply the new database proxy settings. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("effectiveSpecificTime")]
        public Input<string>? EffectiveSpecificTime { get; set; }

        /// <summary>
        /// When modifying the number of proxy instances,The time when you want to apply the new database proxy settings.Valid values:
        /// - Immediate: ApsaraDB RDS immediately applies the new settings.
        /// - MaintainTime: ApsaraDB RDS applies the new settings during the maintenance window that you specified. For more information, see Modify the maintenance window.
        /// - SpecificTime: ApsaraDB RDS applies the new settings at a specified point in time.
        /// 
        /// &gt; **NOTE:** Note If you set the EffectiveTime parameter to SpecificTime, you must specify the EffectiveSpecificTime parameter.
        /// </summary>
        [Input("effectiveTime")]
        public Input<string>? EffectiveTime { get; set; }

        /// <summary>
        /// The Id of instance that can run database.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The network type of the instance. Set the value to VPC.
        /// </summary>
        [Input("instanceNetworkType")]
        public Input<string>? InstanceNetworkType { get; set; }

        /// <summary>
        /// Network type of proxy connection address.
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        /// <summary>
        /// The policy that is used to allocate read weights. Valid values:
        /// - Standard: ApsaraDB RDS automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
        /// - Custom: You must manually allocate read weights to the instance and its read-only instances.
        /// 
        /// &gt; **NOTE:** Note If you set the ReadOnlyInstanceDistributionType parameter to Custom, you must specify the ReadOnlyInstanceWeight parameter.
        /// </summary>
        [Input("readOnlyInstanceDistributionType")]
        public Input<string>? ReadOnlyInstanceDistributionType { get; set; }

        /// <summary>
        /// The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, ApsaraDB RDS no longer forwards read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained. Unit: seconds. Valid values: 0 to 3600.
        /// 
        /// &gt; **NOTE:** Note If the instance runs PostgreSQL, you can enable only the read/write splitting feature, which is specified by ReadWriteSpliting.
        /// </summary>
        [Input("readOnlyInstanceMaxDelayTime")]
        public Input<int>? ReadOnlyInstanceMaxDelayTime { get; set; }

        [Input("readOnlyInstanceWeights")]
        private InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightGetArgs>? _readOnlyInstanceWeights;

        /// <summary>
        /// A list of the read weights of the instance and its read-only instances.  It contains two sub-fields(instance_id and weight). Read weights increase in increments of 100, and the maximum read weight is 10000. See `read_only_instance_weight` below.
        /// </summary>
        public InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightGetArgs> ReadOnlyInstanceWeights
        {
            get => _readOnlyInstanceWeights ?? (_readOnlyInstanceWeights = new InputList<Inputs.RdsDbProxyReadOnlyInstanceWeightGetArgs>());
            set => _readOnlyInstanceWeights = value;
        }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The time when the certificate expires.
        /// </summary>
        [Input("sslExpiredTime")]
        public Input<string>? SslExpiredTime { get; set; }

        /// <summary>
        /// The point in time at which you want to upgrade the database proxy version of the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
        /// </summary>
        [Input("switchTime")]
        public Input<string>? SwitchTime { get; set; }

        /// <summary>
        /// The time when you want to upgrade the database proxy version of the instance. Valid values:
        /// - MaintainTime: ApsaraDB RDS performs the upgrade during the maintenance window that you specified. This is the default value. For more information, see Modify the maintenance window.
        /// - Immediate: ApsaraDB RDS immediately performs the upgrade.
        /// - SpecificTime: ApsaraDB RDS performs the upgrade at a specified point in time.
        /// </summary>
        [Input("upgradeTime")]
        public Input<string>? UpgradeTime { get; set; }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) to which the instance belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The ID of the vSwitch that is associated with the specified VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public RdsDbProxyState()
        {
        }
        public static new RdsDbProxyState Empty => new RdsDbProxyState();
    }
}
