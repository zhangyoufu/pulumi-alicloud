// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cassandra
{
    /// <summary>
    /// Provides a Cassandra dataCenter resource supports replica set dataCenters only. The Cassandra provides stable, reliable, and automatic scalable database services.
    /// It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
    /// You can see detail product introduction [here](https://www.alibabacloud.com/help/product/49055.htm).
    /// 
    /// &gt; **NOTE:**  Available in 1.88.0+.
    /// 
    /// &gt; **NOTE:**  Create a cassandra dataCenter need a clusterId,so need create a cassandra cluster first.
    /// 
    /// &gt; **NOTE:**  The following regions support create Vpc network Cassandra cluster.
    /// The official website mark  more regions. Or you can call [DescribeRegions](https://help.aliyun.com/document_detail/157540.html).
    /// 
    /// &gt; **NOTE:**  Create Cassandra dataCenter or change dataCenter type and storage would cost 30 minutes. Please make full preparation.
    /// 
    /// ## Example Usage
    /// ### Create a cassandra dataCenter
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultCluster = new AliCloud.Cassandra.Cluster("defaultCluster", new()
    ///     {
    ///         ClusterName = "cassandra-cluster-name-tf",
    ///         DataCenterName = "dc-1",
    ///         AutoRenew = false,
    ///         InstanceType = "cassandra.c.large",
    ///         MajorVersion = "3.11",
    ///         NodeCount = 2,
    ///         PayType = "PayAsYouGo",
    ///         VswitchId = "vsw-xxxx1",
    ///         DiskSize = 160,
    ///         DiskType = "cloud_ssd",
    ///         MaintainStartTime = "18:00Z",
    ///         MaintainEndTime = "20:00Z",
    ///         IpWhite = "127.0.0.1",
    ///     });
    /// 
    ///     var defaultDataCenter = new AliCloud.Cassandra.DataCenter("defaultDataCenter", new()
    ///     {
    ///         ClusterId = defaultCluster.Id,
    ///         DataCenterName = "dc-2",
    ///         AutoRenew = false,
    ///         InstanceType = "cassandra.c.large",
    ///         NodeCount = 2,
    ///         PayType = "PayAsYouGo",
    ///         VswitchId = "vsw-xxxx2",
    ///         DiskSize = 160,
    ///         DiskType = "cloud_ssd",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// This is a example for class netType dataCenter. You can find more detail with the examples/cassandra_data_center dir.
    /// 
    /// ## Import
    /// 
    /// If you need full function, please import Cassandra cluster first.
    /// 
    ///  Cassandra dataCenter can be imported using the dcId:clusterId, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cassandra/dataCenter:DataCenter dc_2 cn-shenxxxx-x:cds-wz933ryoaurxxxxx
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cassandra/dataCenter:DataCenter")]
    public partial class DataCenter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// Cassandra cluster id of dataCenter-2 belongs to.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        [Output("dataCenterId")]
        public Output<string> DataCenterId { get; private set; } = null!;

        /// <summary>
        /// Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Output("dataCenterName")]
        public Output<string?> DataCenterName { get; private set; } = null!;

        /// <summary>
        /// User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Output("diskSize")]
        public Output<int?> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        /// </summary>
        [Output("diskType")]
        public Output<string?> DiskType { get; private set; } = null!;

        [Output("enablePublic")]
        public Output<bool?> EnablePublic { get; private set; } = null!;

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The node count of Cassandra dataCenter-2, default to 2.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        /// <summary>
        /// The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        /// </summary>
        [Output("payType")]
        public Output<string> PayType { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        [Output("publicPoints")]
        public Output<ImmutableArray<string>> PublicPoints { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        /// 
        /// &gt; **NOTE:** Now data_center_name,instance_type,node_count,disk_type,disk_size can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DataCenter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataCenter(string name, DataCenterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cassandra/dataCenter:DataCenter", name, args ?? new DataCenterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataCenter(string name, Input<string> id, DataCenterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cassandra/dataCenter:DataCenter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataCenter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataCenter Get(string name, Input<string> id, DataCenterState? state = null, CustomResourceOptions? options = null)
        {
            return new DataCenter(name, id, state, options);
        }
    }

    public sealed class DataCenterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cassandra cluster id of dataCenter-2 belongs to.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("dataCenterName")]
        public Input<string>? DataCenterName { get; set; }

        /// <summary>
        /// User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("enablePublic")]
        public Input<bool>? EnablePublic { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The node count of Cassandra dataCenter-2, default to 2.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        /// </summary>
        [Input("payType", required: true)]
        public Input<string> PayType { get; set; } = null!;

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        /// 
        /// &gt; **NOTE:** Now data_center_name,instance_type,node_count,disk_type,disk_size can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DataCenterArgs()
        {
        }
        public static new DataCenterArgs Empty => new DataCenterArgs();
    }

    public sealed class DataCenterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew of dataCenter-2,`true` or `false`. System default to `false`, valid when pay_type = Subscription.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Period of dataCenter-2 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cassandra cluster id of dataCenter-2 belongs to.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("dataCenterId")]
        public Input<string>? DataCenterId { get; set; }

        /// <summary>
        /// Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("dataCenterName")]
        public Input<string>? DataCenterName { get; set; }

        /// <summary>
        /// User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The disk type of Cassandra dataCenter-2. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("enablePublic")]
        public Input<bool>? EnablePublic { get; set; }

        /// <summary>
        /// Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The node count of Cassandra dataCenter-2, default to 2.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The pay type of Cassandra dataCenter-2. Valid values are `Subscription`, `PayAsYouGo`. System default to `PayAsYouGo`.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("publicPoints")]
        private InputList<string>? _publicPoints;
        public InputList<string> PublicPoints
        {
            get => _publicPoints ?? (_publicPoints = new InputList<string>());
            set => _publicPoints = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty.
        /// 
        /// &gt; **NOTE:** Now data_center_name,instance_type,node_count,disk_type,disk_size can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DataCenterState()
        {
        }
        public static new DataCenterState Empty => new DataCenterState();
    }
}
