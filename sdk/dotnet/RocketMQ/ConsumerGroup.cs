// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    /// <summary>
    /// Provides a RocketMQ Consumer Group resource.
    /// 
    /// For information about RocketMQ Consumer Group and how to use it, see [What is Consumer Group](https://www.alibabacloud.com/help/en/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/api-rocketmq-2022-08-01-createconsumergroup).
    /// 
    /// &gt; **NOTE:** Available since v1.212.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var createVpc = new AliCloud.Vpc.Network("createVpc", new()
    ///     {
    ///         Description = "example",
    ///         CidrBlock = "172.16.0.0/12",
    ///         VpcName = name,
    ///     });
    /// 
    ///     var createVswitch = new AliCloud.Vpc.Switch("createVswitch", new()
    ///     {
    ///         Description = "example",
    ///         VpcId = createVpc.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CidrBlock = "172.16.0.0/24",
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var createInstance = new AliCloud.RocketMQ.RocketMQInstance("createInstance", new()
    ///     {
    ///         AutoRenewPeriod = 1,
    ///         ProductInfo = new AliCloud.RocketMQ.Inputs.RocketMQInstanceProductInfoArgs
    ///         {
    ///             MsgProcessSpec = "rmq.p2.4xlarge",
    ///             SendReceiveRatio = 0.3,
    ///             MessageRetentionTime = 70,
    ///         },
    ///         NetworkInfo = new AliCloud.RocketMQ.Inputs.RocketMQInstanceNetworkInfoArgs
    ///         {
    ///             VpcInfo = new AliCloud.RocketMQ.Inputs.RocketMQInstanceNetworkInfoVpcInfoArgs
    ///             {
    ///                 VpcId = createVpc.Id,
    ///                 VswitchId = createVswitch.Id,
    ///             },
    ///             InternetInfo = new AliCloud.RocketMQ.Inputs.RocketMQInstanceNetworkInfoInternetInfoArgs
    ///             {
    ///                 InternetSpec = "enable",
    ///                 FlowOutType = "payByBandwidth",
    ///                 FlowOutBandwidth = 30,
    ///             },
    ///         },
    ///         Period = 1,
    ///         SubSeriesCode = "cluster_ha",
    ///         Remark = "example",
    ///         InstanceName = name,
    ///         ServiceCode = "rmq",
    ///         SeriesCode = "professional",
    ///         PaymentType = "PayAsYouGo",
    ///         PeriodUnit = "Month",
    ///     });
    /// 
    ///     var defaultConsumerGroup = new AliCloud.RocketMQ.ConsumerGroup("defaultConsumerGroup", new()
    ///     {
    ///         ConsumerGroupId = name,
    ///         InstanceId = createInstance.Id,
    ///         ConsumeRetryPolicy = new AliCloud.RocketMQ.Inputs.ConsumerGroupConsumeRetryPolicyArgs
    ///         {
    ///             RetryPolicy = "DefaultRetryPolicy",
    ///             MaxRetryTimes = 10,
    ///         },
    ///         DeliveryOrderType = "Concurrently",
    ///         Remark = "example",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// RocketMQ Consumer Group can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:rocketmq/consumerGroup:ConsumerGroup example &lt;instance_id&gt;:&lt;consumer_group_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rocketmq/consumerGroup:ConsumerGroup")]
    public partial class ConsumerGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Consumption retry strategy. See `consume_retry_policy` below.
        /// </summary>
        [Output("consumeRetryPolicy")]
        public Output<Outputs.ConsumerGroupConsumeRetryPolicy> ConsumeRetryPolicy { get; private set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Output("consumerGroupId")]
        public Output<string> ConsumerGroupId { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Delivery order.
        /// </summary>
        [Output("deliveryOrderType")]
        public Output<string?> DeliveryOrderType { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Custom remarks.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerGroup(string name, ConsumerGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/consumerGroup:ConsumerGroup", name, args ?? new ConsumerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerGroup(string name, Input<string> id, ConsumerGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/consumerGroup:ConsumerGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerGroup Get(string name, Input<string> id, ConsumerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerGroup(name, id, state, options);
        }
    }

    public sealed class ConsumerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Consumption retry strategy. See `consume_retry_policy` below.
        /// </summary>
        [Input("consumeRetryPolicy", required: true)]
        public Input<Inputs.ConsumerGroupConsumeRetryPolicyArgs> ConsumeRetryPolicy { get; set; } = null!;

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("consumerGroupId", required: true)]
        public Input<string> ConsumerGroupId { get; set; } = null!;

        /// <summary>
        /// Delivery order.
        /// </summary>
        [Input("deliveryOrderType")]
        public Input<string>? DeliveryOrderType { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Custom remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        public ConsumerGroupArgs()
        {
        }
        public static new ConsumerGroupArgs Empty => new ConsumerGroupArgs();
    }

    public sealed class ConsumerGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Consumption retry strategy. See `consume_retry_policy` below.
        /// </summary>
        [Input("consumeRetryPolicy")]
        public Input<Inputs.ConsumerGroupConsumeRetryPolicyGetArgs>? ConsumeRetryPolicy { get; set; }

        /// <summary>
        /// The first ID of the resource.
        /// </summary>
        [Input("consumerGroupId")]
        public Input<string>? ConsumerGroupId { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Delivery order.
        /// </summary>
        [Input("deliveryOrderType")]
        public Input<string>? DeliveryOrderType { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Custom remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ConsumerGroupState()
        {
        }
        public static new ConsumerGroupState Empty => new ConsumerGroupState();
    }
}
