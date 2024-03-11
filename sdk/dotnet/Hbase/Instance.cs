// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbase
{
    /// <summary>
    /// Provides a HBase instance resource supports replica set instances only. The HBase provides stable, reliable, and automatic scalable database services.
    /// It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
    /// You can see detail product introduction [here](https://www.alibabacloud.com/help/en/apsaradb-for-hbase/latest/createcluster)
    /// 
    /// &gt; **NOTE:** Available since v1.67.0.
    /// 
    /// &gt; **NOTE:**  The following regions don't support create Classic network HBase instance.
    /// [`cn-hangzhou`,`cn-shanghai`,`cn-qingdao`,`cn-beijing`,`cn-shenzhen`,`ap-southeast-1a`,.....]
    /// The official website mark  more regions. or you can call [DescribeRegions](https://www.alibabacloud.com/help/en/apsaradb-for-hbase/latest/describeregions)
    /// 
    /// &gt; **NOTE:**  Create HBase instance or change instance type and storage would cost 15 minutes. Please make full preparation
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a hbase instance
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
    ///     var defaultZones = AliCloud.Hbase.GetZones.Invoke();
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "^default-NODELETING$",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Hbase.Instance("defaultInstance", new()
    ///     {
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         Engine = "hbaseue",
    ///         EngineVersion = "2.0",
    ///         MasterInstanceType = "hbase.sn2.2xlarge",
    ///         CoreInstanceType = "hbase.sn2.2xlarge",
    ///         CoreInstanceQuantity = 2,
    ///         CoreDiskType = "cloud_efficiency",
    ///         CoreDiskSize = 400,
    ///         PayType = "PostPaid",
    ///         ColdStorageSize = 0,
    ///         DeletionProtection = false,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// this is a example for class netType instance. you can find more detail with the examples/hbase dir.
    /// 
    /// ## Import
    /// 
    /// HBase can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:hbase/instance:Instance example hb-wz96815u13k659fvd
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:hbase/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The account of the cluster web ui. Size [0-128].
        /// </summary>
        [Output("account")]
        public Output<string?> Account { get; private set; } = null!;

        /// <summary>
        /// Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
        /// </summary>
        [Output("coldStorageSize")]
        public Output<int?> ColdStorageSize { get; private set; } = null!;

        /// <summary>
        /// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
        /// - Custom storage space, value range: [20, 64000].
        /// - Cluster [400, 64000], step:40-GB increments.
        /// - Single [20-500GB], step:1-GB increments.
        /// </summary>
        [Output("coreDiskSize")]
        public Output<int?> CoreDiskSize { get; private set; } = null!;

        /// <summary>
        /// Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
        /// </summary>
        [Output("coreDiskType")]
        public Output<string?> CoreDiskType { get; private set; } = null!;

        /// <summary>
        /// Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster's instance. If core_instance_quantity = 1, this is a single instance.
        /// </summary>
        [Output("coreInstanceQuantity")]
        public Output<int?> CoreInstanceQuantity { get; private set; } = null!;

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Output("coreInstanceType")]
        public Output<string> CoreInstanceType { get; private set; } = null!;

        /// <summary>
        /// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
        /// </summary>
        [Output("duration")]
        public Output<int> Duration { get; private set; } = null!;

        /// <summary>
        /// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
        /// </summary>
        [Output("engine")]
        public Output<string?> Engine { get; private set; } = null!;

        /// <summary>
        /// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://www.alibabacloud.com/help/en/data-lake-analytics/latest/createinstance).
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
        /// </summary>
        [Output("immediateDeleteFlag")]
        public Output<bool?> ImmediateDeleteFlag { get; private set; } = null!;

        /// <summary>
        /// The white ip list of the cluster.
        /// </summary>
        [Output("ipWhite")]
        public Output<string> IpWhite { get; private set; } = null!;

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
        /// </summary>
        [Output("maintainEndTime")]
        public Output<string> MaintainEndTime { get; private set; } = null!;

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Output("maintainStartTime")]
        public Output<string> MaintainStartTime { get; private set; } = null!;

        /// <summary>
        /// Count nodes of the master node.
        /// </summary>
        [Output("masterInstanceQuantity")]
        public Output<int> MasterInstanceQuantity { get; private set; } = null!;

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Output("masterInstanceType")]
        public Output<string> MasterInstanceType { get; private set; } = null!;

        /// <summary>
        /// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password of the cluster web ui account. Size [0-128].
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
        /// </summary>
        [Output("payType")]
        public Output<string?> PayType { get; private set; } = null!;

        /// <summary>
        /// The security group resource of the cluster.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The slb service addresses of the cluster. See `slb_conn_addrs` below.
        /// 
        /// &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
        /// </summary>
        [Output("slbConnAddrs")]
        public Output<ImmutableArray<Outputs.InstanceSlbConnAddr>> SlbConnAddrs { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The Web UI proxy addresses of the cluster. See `ui_proxy_conn_addrs` below.
        /// </summary>
        [Output("uiProxyConnAddrs")]
        public Output<ImmutableArray<Outputs.InstanceUiProxyConnAddr>> UiProxyConnAddrs { get; private set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The zookeeper addresses of the cluster. See `zk_conn_addrs` below.
        /// </summary>
        [Output("zkConnAddrs")]
        public Output<ImmutableArray<Outputs.InstanceZkConnAddr>> ZkConnAddrs { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:hbase/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:hbase/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account of the cluster web ui. Size [0-128].
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
        /// </summary>
        [Input("coldStorageSize")]
        public Input<int>? ColdStorageSize { get; set; }

        /// <summary>
        /// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
        /// - Custom storage space, value range: [20, 64000].
        /// - Cluster [400, 64000], step:40-GB increments.
        /// - Single [20-500GB], step:1-GB increments.
        /// </summary>
        [Input("coreDiskSize")]
        public Input<int>? CoreDiskSize { get; set; }

        /// <summary>
        /// Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
        /// </summary>
        [Input("coreDiskType")]
        public Input<string>? CoreDiskType { get; set; }

        /// <summary>
        /// Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster's instance. If core_instance_quantity = 1, this is a single instance.
        /// </summary>
        [Input("coreInstanceQuantity")]
        public Input<int>? CoreInstanceQuantity { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Input("coreInstanceType", required: true)]
        public Input<string> CoreInstanceType { get; set; } = null!;

        /// <summary>
        /// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://www.alibabacloud.com/help/en/data-lake-analytics/latest/createinstance).
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
        /// </summary>
        [Input("immediateDeleteFlag")]
        public Input<bool>? ImmediateDeleteFlag { get; set; }

        /// <summary>
        /// The white ip list of the cluster.
        /// </summary>
        [Input("ipWhite")]
        public Input<string>? IpWhite { get; set; }

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Input("masterInstanceType", required: true)]
        public Input<string> MasterInstanceType { get; set; } = null!;

        /// <summary>
        /// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the cluster web ui account. Size [0-128].
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The security group resource of the cluster.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account of the cluster web ui. Size [0-128].
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
        /// </summary>
        [Input("coldStorageSize")]
        public Input<int>? ColdStorageSize { get; set; }

        /// <summary>
        /// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
        /// - Custom storage space, value range: [20, 64000].
        /// - Cluster [400, 64000], step:40-GB increments.
        /// - Single [20-500GB], step:1-GB increments.
        /// </summary>
        [Input("coreDiskSize")]
        public Input<int>? CoreDiskSize { get; set; }

        /// <summary>
        /// Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
        /// </summary>
        [Input("coreDiskType")]
        public Input<string>? CoreDiskType { get; set; }

        /// <summary>
        /// Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster's instance. If core_instance_quantity = 1, this is a single instance.
        /// </summary>
        [Input("coreInstanceQuantity")]
        public Input<int>? CoreInstanceQuantity { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Input("coreInstanceType")]
        public Input<string>? CoreInstanceType { get; set; }

        /// <summary>
        /// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://www.alibabacloud.com/help/en/data-lake-analytics/latest/createinstance).
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
        /// </summary>
        [Input("immediateDeleteFlag")]
        public Input<bool>? ImmediateDeleteFlag { get; set; }

        /// <summary>
        /// The white ip list of the cluster.
        /// </summary>
        [Input("ipWhite")]
        public Input<string>? IpWhite { get; set; }

        /// <summary>
        /// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// Count nodes of the master node.
        /// </summary>
        [Input("masterInstanceQuantity")]
        public Input<int>? MasterInstanceQuantity { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
        /// </summary>
        [Input("masterInstanceType")]
        public Input<string>? MasterInstanceType { get; set; }

        /// <summary>
        /// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the cluster web ui account. Size [0-128].
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// The security group resource of the cluster.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("slbConnAddrs")]
        private InputList<Inputs.InstanceSlbConnAddrGetArgs>? _slbConnAddrs;

        /// <summary>
        /// The slb service addresses of the cluster. See `slb_conn_addrs` below.
        /// 
        /// &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
        /// </summary>
        public InputList<Inputs.InstanceSlbConnAddrGetArgs> SlbConnAddrs
        {
            get => _slbConnAddrs ?? (_slbConnAddrs = new InputList<Inputs.InstanceSlbConnAddrGetArgs>());
            set => _slbConnAddrs = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("uiProxyConnAddrs")]
        private InputList<Inputs.InstanceUiProxyConnAddrGetArgs>? _uiProxyConnAddrs;

        /// <summary>
        /// The Web UI proxy addresses of the cluster. See `ui_proxy_conn_addrs` below.
        /// </summary>
        public InputList<Inputs.InstanceUiProxyConnAddrGetArgs> UiProxyConnAddrs
        {
            get => _uiProxyConnAddrs ?? (_uiProxyConnAddrs = new InputList<Inputs.InstanceUiProxyConnAddrGetArgs>());
            set => _uiProxyConnAddrs = value;
        }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        [Input("zkConnAddrs")]
        private InputList<Inputs.InstanceZkConnAddrGetArgs>? _zkConnAddrs;

        /// <summary>
        /// The zookeeper addresses of the cluster. See `zk_conn_addrs` below.
        /// </summary>
        public InputList<Inputs.InstanceZkConnAddrGetArgs> ZkConnAddrs
        {
            get => _zkConnAddrs ?? (_zkConnAddrs = new InputList<Inputs.InstanceZkConnAddrGetArgs>());
            set => _zkConnAddrs = value;
        }

        /// <summary>
        /// The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
