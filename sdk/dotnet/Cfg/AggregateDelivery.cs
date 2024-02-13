// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cfg
{
    /// <summary>
    /// Provides a Cloud Config Aggregate Delivery resource.
    /// 
    /// For information about Cloud Config Aggregate Delivery and how to use it, see [What is Aggregate Delivery](https://www.alibabacloud.com/help/en/cloud-config/latest/api-config-2020-09-07-createaggregateconfigdeliverychannel).
    /// 
    /// &gt; **NOTE:** Available since v1.172.0.
    /// 
    /// ## Import
    /// 
    /// Cloud Config Aggregate Delivery can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cfg/aggregateDelivery:AggregateDelivery example &lt;aggregator_id&gt;:&lt;delivery_channel_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cfg/aggregateDelivery:AggregateDelivery")]
    public partial class AggregateDelivery : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Aggregator.
        /// </summary>
        [Output("aggregatorId")]
        public Output<string> AggregatorId { get; private set; } = null!;

        /// <summary>
        /// Open or close delivery configuration change history.
        /// </summary>
        [Output("configurationItemChangeNotification")]
        public Output<bool> ConfigurationItemChangeNotification { get; private set; } = null!;

        /// <summary>
        /// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        /// </summary>
        [Output("configurationSnapshot")]
        public Output<bool> ConfigurationSnapshot { get; private set; } = null!;

        /// <summary>
        /// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        /// </summary>
        [Output("deliveryChannelCondition")]
        public Output<string?> DeliveryChannelCondition { get; private set; } = null!;

        /// <summary>
        /// The ID of the delivery method.
        /// </summary>
        [Output("deliveryChannelId")]
        public Output<string> DeliveryChannelId { get; private set; } = null!;

        /// <summary>
        /// The name of the delivery method.
        /// </summary>
        [Output("deliveryChannelName")]
        public Output<string?> DeliveryChannelName { get; private set; } = null!;

        /// <summary>
        /// The ARN of the delivery destination. The value must be in one of the following formats:
        /// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Output("deliveryChannelTargetArn")]
        public Output<string> DeliveryChannelTargetArn { get; private set; } = null!;

        /// <summary>
        /// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Output("deliveryChannelType")]
        public Output<string> DeliveryChannelType { get; private set; } = null!;

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        /// </summary>
        [Output("nonCompliantNotification")]
        public Output<bool> NonCompliantNotification { get; private set; } = null!;

        /// <summary>
        /// The oss ARN of the delivery channel when the value data oversized limit.
        /// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        /// </summary>
        [Output("oversizedDataOssTargetArn")]
        public Output<string?> OversizedDataOssTargetArn { get; private set; } = null!;

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AggregateDelivery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AggregateDelivery(string name, AggregateDeliveryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cfg/aggregateDelivery:AggregateDelivery", name, args ?? new AggregateDeliveryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AggregateDelivery(string name, Input<string> id, AggregateDeliveryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cfg/aggregateDelivery:AggregateDelivery", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AggregateDelivery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AggregateDelivery Get(string name, Input<string> id, AggregateDeliveryState? state = null, CustomResourceOptions? options = null)
        {
            return new AggregateDelivery(name, id, state, options);
        }
    }

    public sealed class AggregateDeliveryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Aggregator.
        /// </summary>
        [Input("aggregatorId", required: true)]
        public Input<string> AggregatorId { get; set; } = null!;

        /// <summary>
        /// Open or close delivery configuration change history.
        /// </summary>
        [Input("configurationItemChangeNotification")]
        public Input<bool>? ConfigurationItemChangeNotification { get; set; }

        /// <summary>
        /// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        /// </summary>
        [Input("configurationSnapshot")]
        public Input<bool>? ConfigurationSnapshot { get; set; }

        /// <summary>
        /// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        /// </summary>
        [Input("deliveryChannelCondition")]
        public Input<string>? DeliveryChannelCondition { get; set; }

        /// <summary>
        /// The name of the delivery method.
        /// </summary>
        [Input("deliveryChannelName")]
        public Input<string>? DeliveryChannelName { get; set; }

        /// <summary>
        /// The ARN of the delivery destination. The value must be in one of the following formats:
        /// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Input("deliveryChannelTargetArn", required: true)]
        public Input<string> DeliveryChannelTargetArn { get; set; } = null!;

        /// <summary>
        /// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Input("deliveryChannelType", required: true)]
        public Input<string> DeliveryChannelType { get; set; } = null!;

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        /// </summary>
        [Input("nonCompliantNotification")]
        public Input<bool>? NonCompliantNotification { get; set; }

        /// <summary>
        /// The oss ARN of the delivery channel when the value data oversized limit.
        /// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        /// </summary>
        [Input("oversizedDataOssTargetArn")]
        public Input<string>? OversizedDataOssTargetArn { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public AggregateDeliveryArgs()
        {
        }
        public static new AggregateDeliveryArgs Empty => new AggregateDeliveryArgs();
    }

    public sealed class AggregateDeliveryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Aggregator.
        /// </summary>
        [Input("aggregatorId")]
        public Input<string>? AggregatorId { get; set; }

        /// <summary>
        /// Open or close delivery configuration change history.
        /// </summary>
        [Input("configurationItemChangeNotification")]
        public Input<bool>? ConfigurationItemChangeNotification { get; set; }

        /// <summary>
        /// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `OSS`.
        /// </summary>
        [Input("configurationSnapshot")]
        public Input<bool>? ConfigurationSnapshot { get; set; }

        /// <summary>
        /// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `MNS`.
        /// </summary>
        [Input("deliveryChannelCondition")]
        public Input<string>? DeliveryChannelCondition { get; set; }

        /// <summary>
        /// The ID of the delivery method.
        /// </summary>
        [Input("deliveryChannelId")]
        public Input<string>? DeliveryChannelId { get; set; }

        /// <summary>
        /// The name of the delivery method.
        /// </summary>
        [Input("deliveryChannelName")]
        public Input<string>? DeliveryChannelName { get; set; }

        /// <summary>
        /// The ARN of the delivery destination. The value must be in one of the following formats:
        /// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Input("deliveryChannelTargetArn")]
        public Input<string>? DeliveryChannelTargetArn { get; set; }

        /// <summary>
        /// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Input("deliveryChannelType")]
        public Input<string>? DeliveryChannelType { get; set; }

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `delivery_channel_type` is `SLS` or `MNS`.
        /// </summary>
        [Input("nonCompliantNotification")]
        public Input<bool>? NonCompliantNotification { get; set; }

        /// <summary>
        /// The oss ARN of the delivery channel when the value data oversized limit.
        /// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
        /// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
        /// </summary>
        [Input("oversizedDataOssTargetArn")]
        public Input<string>? OversizedDataOssTargetArn { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public AggregateDeliveryState()
        {
        }
        public static new AggregateDeliveryState Empty => new AggregateDeliveryState();
    }
}
