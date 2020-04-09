// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Adb
{
    /// <summary>
    /// Provides a ADB cluster resource. A ADB cluster is an isolated database
    /// environment in the cloud. A ADB cluster can contain multiple user-created
    /// databases.
    /// 
    /// &gt; **NOTE:** Available in v1.71.0+.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/adb_cluster.html.markdown.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Output("autoRenewPeriod")]
        public Output<int?> AutoRenewPeriod { get; private set; } = null!;

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Output("dbClusterCategory")]
        public Output<string> DbClusterCategory { get; private set; } = null!;

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Output("dbClusterVersion")]
        public Output<string?> DbClusterVersion { get; private set; } = null!;

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Output("dbNodeClass")]
        public Output<string> DbNodeClass { get; private set; } = null!;

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Output("dbNodeCount")]
        public Output<int> DbNodeCount { get; private set; } = null!;

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Output("dbNodeStorage")]
        public Output<int> DbNodeStorage { get; private set; } = null!;

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Output("maintainTime")]
        public Output<string> MaintainTime { get; private set; } = null!;

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Output("payType")]
        public Output<string?> PayType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Output("renewalStatus")]
        public Output<string?> RenewalStatus { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        [Output("securityIps")]
        public Output<ImmutableArray<string>> SecurityIps { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the DB cluster.
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
            : base("alicloud:adb/cluster:Cluster", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:adb/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Input("dbClusterCategory", required: true)]
        public Input<string> DbClusterCategory { get; set; } = null!;

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Input("dbClusterVersion")]
        public Input<string>? DbClusterVersion { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass", required: true)]
        public Input<string> DbNodeClass { get; set; } = null!;

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Input("dbNodeCount", required: true)]
        public Input<int> DbNodeCount { get; set; } = null!;

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Input("dbNodeStorage", required: true)]
        public Input<int> DbNodeStorage { get; set; } = null!;

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
        /// </summary>
        [Input("autoRenewPeriod")]
        public Input<int>? AutoRenewPeriod { get; set; }

        /// <summary>
        /// Cluster category. Value options: `Basic`, `Cluster`.
        /// </summary>
        [Input("dbClusterCategory")]
        public Input<string>? DbClusterCategory { get; set; }

        /// <summary>
        /// Cluster version. Value options: `3.0`, Default to `3.0`.
        /// </summary>
        [Input("dbClusterVersion")]
        public Input<string>? DbClusterVersion { get; set; }

        /// <summary>
        /// The db_node_class of cluster node.
        /// </summary>
        [Input("dbNodeClass")]
        public Input<string>? DbNodeClass { get; set; }

        /// <summary>
        /// The db_node_count of cluster node.
        /// </summary>
        [Input("dbNodeCount")]
        public Input<int>? DbNodeCount { get; set; }

        /// <summary>
        /// The db_node_storage of cluster node.
        /// </summary>
        [Input("dbNodeStorage")]
        public Input<int>? DbNodeStorage { get; set; }

        /// <summary>
        /// The description of cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
        /// </summary>
        [Input("maintainTime")]
        public Input<string>? MaintainTime { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
        /// </summary>
        [Input("payType")]
        public Input<string>? PayType { get; set; }

        /// <summary>
        /// The duration that you will buy DB cluster (in month). It is valid when pay_type is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("securityIps")]
        private InputList<string>? _securityIps;

        /// <summary>
        /// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        /// </summary>
        public InputList<string> SecurityIps
        {
            get => _securityIps ?? (_securityIps = new InputList<string>());
            set => _securityIps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
        /// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The virtual switch ID to launch DB instances in one VPC.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the DB cluster.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ClusterState()
        {
        }
    }
}
