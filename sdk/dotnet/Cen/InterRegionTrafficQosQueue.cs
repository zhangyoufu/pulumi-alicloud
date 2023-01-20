// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a Cen Inter Region Traffic Qos Queue resource.
    /// 
    /// For information about Cen Inter Region Traffic Qos Queue and how to use it, see [What is Inter Region Traffic Qos Queue](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-createceninterregiontrafficqosqueue).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new AliCloud.Cen.InterRegionTrafficQosQueue("default", new()
    ///     {
    ///         Dscps = new[]
    ///         {
    ///             "1",
    ///             "2",
    ///         },
    ///         InterRegionTrafficQosQueueDescription = "test",
    ///         RemainBandwidthPercent = 20,
    ///         TrafficQosPolicyId = "qos-xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cen Inter Region Traffic Qos Queue can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue")]
    public partial class InterRegionTrafficQosQueue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
        /// </summary>
        [Output("dscps")]
        public Output<ImmutableArray<string>> Dscps { get; private set; } = null!;

        /// <summary>
        /// The description information of the traffic scheduling policy.
        /// </summary>
        [Output("interRegionTrafficQosQueueDescription")]
        public Output<string?> InterRegionTrafficQosQueueDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the traffic scheduling policy.
        /// </summary>
        [Output("interRegionTrafficQosQueueName")]
        public Output<string?> InterRegionTrafficQosQueueName { get; private set; } = null!;

        /// <summary>
        /// The percentage of cross-region bandwidth that the current queue can use.
        /// </summary>
        [Output("remainBandwidthPercent")]
        public Output<int> RemainBandwidthPercent { get; private set; } = null!;

        /// <summary>
        /// The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the traffic scheduling policy.
        /// </summary>
        [Output("trafficQosPolicyId")]
        public Output<string> TrafficQosPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a InterRegionTrafficQosQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterRegionTrafficQosQueue(string name, InterRegionTrafficQosQueueArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue", name, args ?? new InterRegionTrafficQosQueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterRegionTrafficQosQueue(string name, Input<string> id, InterRegionTrafficQosQueueState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterRegionTrafficQosQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterRegionTrafficQosQueue Get(string name, Input<string> id, InterRegionTrafficQosQueueState? state = null, CustomResourceOptions? options = null)
        {
            return new InterRegionTrafficQosQueue(name, id, state, options);
        }
    }

    public sealed class InterRegionTrafficQosQueueArgs : global::Pulumi.ResourceArgs
    {
        [Input("dscps", required: true)]
        private InputList<string>? _dscps;

        /// <summary>
        /// The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
        /// </summary>
        public InputList<string> Dscps
        {
            get => _dscps ?? (_dscps = new InputList<string>());
            set => _dscps = value;
        }

        /// <summary>
        /// The description information of the traffic scheduling policy.
        /// </summary>
        [Input("interRegionTrafficQosQueueDescription")]
        public Input<string>? InterRegionTrafficQosQueueDescription { get; set; }

        /// <summary>
        /// The name of the traffic scheduling policy.
        /// </summary>
        [Input("interRegionTrafficQosQueueName")]
        public Input<string>? InterRegionTrafficQosQueueName { get; set; }

        /// <summary>
        /// The percentage of cross-region bandwidth that the current queue can use.
        /// </summary>
        [Input("remainBandwidthPercent", required: true)]
        public Input<int> RemainBandwidthPercent { get; set; } = null!;

        /// <summary>
        /// The ID of the traffic scheduling policy.
        /// </summary>
        [Input("trafficQosPolicyId", required: true)]
        public Input<string> TrafficQosPolicyId { get; set; } = null!;

        public InterRegionTrafficQosQueueArgs()
        {
        }
        public static new InterRegionTrafficQosQueueArgs Empty => new InterRegionTrafficQosQueueArgs();
    }

    public sealed class InterRegionTrafficQosQueueState : global::Pulumi.ResourceArgs
    {
        [Input("dscps")]
        private InputList<string>? _dscps;

        /// <summary>
        /// The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
        /// </summary>
        public InputList<string> Dscps
        {
            get => _dscps ?? (_dscps = new InputList<string>());
            set => _dscps = value;
        }

        /// <summary>
        /// The description information of the traffic scheduling policy.
        /// </summary>
        [Input("interRegionTrafficQosQueueDescription")]
        public Input<string>? InterRegionTrafficQosQueueDescription { get; set; }

        /// <summary>
        /// The name of the traffic scheduling policy.
        /// </summary>
        [Input("interRegionTrafficQosQueueName")]
        public Input<string>? InterRegionTrafficQosQueueName { get; set; }

        /// <summary>
        /// The percentage of cross-region bandwidth that the current queue can use.
        /// </summary>
        [Input("remainBandwidthPercent")]
        public Input<int>? RemainBandwidthPercent { get; set; }

        /// <summary>
        /// The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the traffic scheduling policy.
        /// </summary>
        [Input("trafficQosPolicyId")]
        public Input<string>? TrafficQosPolicyId { get; set; }

        public InterRegionTrafficQosQueueState()
        {
        }
        public static new InterRegionTrafficQosQueueState Empty => new InterRegionTrafficQosQueueState();
    }
}
