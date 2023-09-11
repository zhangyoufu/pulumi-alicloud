// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a Vpc Flow Log resource. While it uses alicloud.vpc.FlowLog to build a vpc flow log resource, it will be active by default.
    /// 
    /// For information about Vpc Flow Log and how to use it, see [What is Flow Log](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/flow-logs-overview).
    /// 
    /// &gt; **NOTE:** Available since v1.117.0.
    /// 
    /// ## Import
    /// 
    /// Vpc Flow Log can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/flowLog:FlowLog example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/flowLog:FlowLog")]
    public partial class FlowLog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Data aggregation interval.
        /// </summary>
        [Output("aggregationInterval")]
        public Output<string> AggregationInterval { get; private set; } = null!;

        /// <summary>
        /// Business status.
        /// </summary>
        [Output("businessStatus")]
        public Output<string> BusinessStatus { get; private set; } = null!;

        /// <summary>
        /// Creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The Description of the VPC Flow Log.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The flow log ID.
        /// </summary>
        [Output("flowLogId")]
        public Output<string> FlowLogId { get; private set; } = null!;

        /// <summary>
        /// The Name of the VPC Flow Log.
        /// </summary>
        [Output("flowLogName")]
        public Output<string?> FlowLogName { get; private set; } = null!;

        /// <summary>
        /// The name of the logstore.
        /// </summary>
        [Output("logStoreName")]
        public Output<string> LogStoreName { get; private set; } = null!;

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the current instance resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        /// </summary>
        [Output("trafficPaths")]
        public Output<ImmutableArray<string>> TrafficPaths { get; private set; } = null!;

        /// <summary>
        /// The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        /// </summary>
        [Output("trafficType")]
        public Output<string> TrafficType { get; private set; } = null!;


        /// <summary>
        /// Create a FlowLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowLog(string name, FlowLogArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/flowLog:FlowLog", name, args ?? new FlowLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlowLog(string name, Input<string> id, FlowLogState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/flowLog:FlowLog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FlowLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowLog Get(string name, Input<string> id, FlowLogState? state = null, CustomResourceOptions? options = null)
        {
            return new FlowLog(name, id, state, options);
        }
    }

    public sealed class FlowLogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data aggregation interval.
        /// </summary>
        [Input("aggregationInterval")]
        public Input<string>? AggregationInterval { get; set; }

        /// <summary>
        /// The Description of the VPC Flow Log.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Name of the VPC Flow Log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// The name of the logstore.
        /// </summary>
        [Input("logStoreName", required: true)]
        public Input<string> LogStoreName { get; set; } = null!;

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the current instance resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("trafficPaths")]
        private InputList<string>? _trafficPaths;

        /// <summary>
        /// The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        /// </summary>
        public InputList<string> TrafficPaths
        {
            get => _trafficPaths ?? (_trafficPaths = new InputList<string>());
            set => _trafficPaths = value;
        }

        /// <summary>
        /// The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        /// </summary>
        [Input("trafficType", required: true)]
        public Input<string> TrafficType { get; set; } = null!;

        public FlowLogArgs()
        {
        }
        public static new FlowLogArgs Empty => new FlowLogArgs();
    }

    public sealed class FlowLogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data aggregation interval.
        /// </summary>
        [Input("aggregationInterval")]
        public Input<string>? AggregationInterval { get; set; }

        /// <summary>
        /// Business status.
        /// </summary>
        [Input("businessStatus")]
        public Input<string>? BusinessStatus { get; set; }

        /// <summary>
        /// Creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The Description of the VPC Flow Log.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The flow log ID.
        /// </summary>
        [Input("flowLogId")]
        public Input<string>? FlowLogId { get; set; }

        /// <summary>
        /// The Name of the VPC Flow Log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// The name of the logstore.
        /// </summary>
        [Input("logStoreName")]
        public Input<string>? LogStoreName { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The ID of the resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The resource type of the traffic captured by the flow log:-**NetworkInterface**: ENI.-**VSwitch**: All ENIs in the VSwitch.-**VPC**: All ENIs in the VPC.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The status of the VPC Flow Log. Valid values: **Active** and **Inactive**.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the current instance resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("trafficPaths")]
        private InputList<string>? _trafficPaths;

        /// <summary>
        /// The collected flow path. Value:**all**: indicates full acquisition.**internetGateway**: indicates public network traffic collection.
        /// </summary>
        public InputList<string> TrafficPaths
        {
            get => _trafficPaths ?? (_trafficPaths = new InputList<string>());
            set => _trafficPaths = value;
        }

        /// <summary>
        /// The type of traffic collected. Valid values:**All**: All traffic.**Allow**: Access control allowedtraffic.**Drop**: Access control denied traffic.
        /// </summary>
        [Input("trafficType")]
        public Input<string>? TrafficType { get; set; }

        public FlowLogState()
        {
        }
        public static new FlowLogState Empty => new FlowLogState();
    }
}
