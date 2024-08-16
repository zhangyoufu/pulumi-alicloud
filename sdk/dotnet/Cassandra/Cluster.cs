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
    /// Provides a Cassandra cluster resource supports replica set clusters only. The Cassandra provides stable, reliable, and automatic scalable database services.
    /// It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
    /// You can see detail product introduction [here](https://www.alibabacloud.com/help/product/49055.htm).
    /// 
    /// &gt; **NOTE:**  Available in 1.88.0+.
    /// 
    /// &gt; **NOTE:**  The following regions support create Vpc network Cassandra cluster.
    /// The official website mark more regions. Or you can call [DescribeRegions](https://help.aliyun.com/document_detail/157540.html).
    /// 
    /// &gt; **NOTE:**  Create Cassandra cluster or change cluster type and storage would cost 30 minutes. Please make full preparation.
    /// 
    /// &gt; **DEPRECATED:**  This resource has been [deprecated](https://www.alibabacloud.com/help/en/apsaradb-for-cassandra/latest/cassandra-delisting-notice) from version `1.220.0`.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a cassandra cluster
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new AliCloud.Cassandra.Cluster("default", new()
    ///     {
    ///         ClusterName = "cassandra-cluster-name-tf",
    ///         DataCenterName = "dc-1",
    ///         AutoRenew = false,
    ///         InstanceType = "cassandra.c.large",
    ///         MajorVersion = "3.11",
    ///         NodeCount = 2,
    ///         PayType = "PayAsYouGo",
    ///         VswitchId = "vsw-xxxx",
    ///         DiskSize = 160,
    ///         DiskType = "cloud_ssd",
    ///         MaintainStartTime = "18:00Z",
    ///         MaintainEndTime = "20:00Z",
    ///         IpWhite = "127.0.0.1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// This is a example for class netType cluster. You can find more detail with the examples/cassandra_cluster dir.
    /// 
    /// ## Import
    /// 
    /// Cassandra cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cassandra/cluster:Cluster example cds-wz9sr400dd7xxxxx
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cassandra/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Output("clusterName")]
        public Output<string?> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Output("dataCenterName")]
        public Output<string?> DataCenterName { get; private set; } = null!;

        /// <summary>
        /// User-defined Cassandra dataCenter-1 one node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Output("diskSize")]
        public Output<int?> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
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
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        [Output("ipWhite")]
        public Output<string> IpWhite { get; private set; } = null!;

        /// <summary>
        /// The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Output("maintainEndTime")]
        public Output<string?> MaintainEndTime { get; private set; } = null!;

        /// <summary>
        /// The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Output("maintainStartTime")]
        public Output<string?> MaintainStartTime { get; private set; } = null!;

        /// <summary>
        /// Cassandra major version. Now only support version `3.11`.
        /// </summary>
        [Output("majorVersion")]
        public Output<string> MajorVersion { get; private set; } = null!;

        /// <summary>
        /// The node count of Cassandra dataCenter-1 default to 2.
        /// </summary>
        [Output("nodeCount")]
        public Output<int> NodeCount { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
        /// </summary>
        [Output("payType")]
        public Output<string> PayType { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        [Output("publicPoints")]
        public Output<ImmutableArray<string>> PublicPoints { get; private set; } = null!;

        /// <summary>
        /// A list of security group ids to associate with.
        /// 
        /// &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The vswitch_id of dataCenter-1, can not empty.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cassandra/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cassandra/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("dataCenterName")]
        public Input<string>? DataCenterName { get; set; }

        /// <summary>
        /// User-defined Cassandra dataCenter-1 one node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
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
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        [Input("ipWhite")]
        public Input<string>? IpWhite { get; set; }

        /// <summary>
        /// The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// Cassandra major version. Now only support version `3.11`.
        /// </summary>
        [Input("majorVersion", required: true)]
        public Input<string> MajorVersion { get; set; } = null!;

        /// <summary>
        /// The node count of Cassandra dataCenter-1 default to 2.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;
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
        /// The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
        /// </summary>
        [Input("payType", required: true)]
        public Input<string> PayType { get; set; } = null!;

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group ids to associate with.
        /// 
        /// &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The vswitch_id of dataCenter-1, can not empty.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        /// <summary>
        /// The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto renew of dataCenter-1,`true` or `false`. System default to `false`, valid when pay_type = PrePaid.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Period of dataCenter-1 auto renew, if auto renew is `true`, one of `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60`, valid when pay_type = Subscription. Unit: month.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period `.`, underline `_`, or dash `-` are permitted.
        /// </summary>
        [Input("dataCenterName")]
        public Input<string>? DataCenterName { get; set; }

        /// <summary>
        /// User-defined Cassandra dataCenter-1 one node's storage space.Unit: GB. Value range:
        /// - Custom storage space; value range: [160, 2000].
        /// - 80-GB increments.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The disk type of Cassandra dataCenter-1. Valid values are `cloud_ssd`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`, local_disk size is fixed.
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
        /// Set the instance's IP whitelist in VPC network.
        /// </summary>
        [Input("ipWhite")]
        public Input<string>? IpWhite { get; set; }

        /// <summary>
        /// The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainEndTime")]
        public Input<string>? MaintainEndTime { get; set; }

        /// <summary>
        /// The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).
        /// </summary>
        [Input("maintainStartTime")]
        public Input<string>? MaintainStartTime { get; set; }

        /// <summary>
        /// Cassandra major version. Now only support version `3.11`.
        /// </summary>
        [Input("majorVersion")]
        public Input<string>? MajorVersion { get; set; }

        /// <summary>
        /// The node count of Cassandra dataCenter-1 default to 2.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        [Input("password")]
        private Input<string>? _password;
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
        /// The pay type of Cassandra dataCenter-1. Valid values are `Subscription`, `PayAsYouGo`,System default to `PayAsYouGo`.
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

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group ids to associate with.
        /// 
        /// &gt; **NOTE:** Now cluster_name,data_center_name,instance_type,node_count,disk_type,disk_size,maintain_start_time,maintain_end_time,tags,ip_white,security_groups can be change. The others(auto_renew, auto_renew_period and so on) will be supported in the furture.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The vswitch_id of dataCenter-1, can not empty.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be "" or consistent.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
