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
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Alicloud Config Delivery Channel can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cfg/deliveryChannel:DeliveryChannel example cdc-49a2ad756057********
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cfg/deliveryChannel:DeliveryChannel")]
    public partial class DeliveryChannel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
        /// </summary>
        [Output("deliveryChannelAssumeRoleArn")]
        public Output<string> DeliveryChannelAssumeRoleArn { get; private set; } = null!;

        /// <summary>
        /// The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
        /// </summary>
        [Output("deliveryChannelCondition")]
        public Output<string> DeliveryChannelCondition { get; private set; } = null!;

        /// <summary>
        /// The name of the delivery channel.
        /// </summary>
        [Output("deliveryChannelName")]
        public Output<string> DeliveryChannelName { get; private set; } = null!;

        /// <summary>
        /// The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
        /// - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Output("deliveryChannelTargetArn")]
        public Output<string> DeliveryChannelTargetArn { get; private set; } = null!;

        /// <summary>
        /// The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Output("deliveryChannelType")]
        public Output<string> DeliveryChannelType { get; private set; } = null!;

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DeliveryChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeliveryChannel(string name, DeliveryChannelArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cfg/deliveryChannel:DeliveryChannel", name, args ?? new DeliveryChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeliveryChannel(string name, Input<string> id, DeliveryChannelState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cfg/deliveryChannel:DeliveryChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeliveryChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeliveryChannel Get(string name, Input<string> id, DeliveryChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new DeliveryChannel(name, id, state, options);
        }
    }

    public sealed class DeliveryChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
        /// </summary>
        [Input("deliveryChannelAssumeRoleArn", required: true)]
        public Input<string> DeliveryChannelAssumeRoleArn { get; set; } = null!;

        /// <summary>
        /// The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
        /// </summary>
        [Input("deliveryChannelCondition")]
        public Input<string>? DeliveryChannelCondition { get; set; }

        /// <summary>
        /// The name of the delivery channel.
        /// </summary>
        [Input("deliveryChannelName")]
        public Input<string>? DeliveryChannelName { get; set; }

        /// <summary>
        /// The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
        /// - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Input("deliveryChannelTargetArn", required: true)]
        public Input<string> DeliveryChannelTargetArn { get; set; } = null!;

        /// <summary>
        /// The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Input("deliveryChannelType", required: true)]
        public Input<string> DeliveryChannelType { get; set; } = null!;

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public DeliveryChannelArgs()
        {
        }
        public static new DeliveryChannelArgs Empty => new DeliveryChannelArgs();
    }

    public sealed class DeliveryChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method.
        /// </summary>
        [Input("deliveryChannelAssumeRoleArn")]
        public Input<string>? DeliveryChannelAssumeRoleArn { get; set; }

        /// <summary>
        /// The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type. Please refer to api [PutDeliveryChannel](https://www.alibabacloud.com/help/en/doc-detail/174253.htm) for example format.
        /// </summary>
        [Input("deliveryChannelCondition")]
        public Input<string>? DeliveryChannelCondition { get; set; }

        /// <summary>
        /// The name of the delivery channel.
        /// </summary>
        [Input("deliveryChannelName")]
        public Input<string>? DeliveryChannelName { get; set; }

        /// <summary>
        /// The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
        /// - `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
        /// - `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
        /// - `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
        /// </summary>
        [Input("deliveryChannelTargetArn")]
        public Input<string>? DeliveryChannelTargetArn { get; set; }

        /// <summary>
        /// The type of the delivery method. This parameter is required when you create a delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
        /// </summary>
        [Input("deliveryChannelType")]
        public Input<string>? DeliveryChannelType { get; set; }

        /// <summary>
        /// The description of the delivery method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The status of the delivery method. Valid values: `0`: The delivery method is disabled., `1`: The delivery destination is enabled. This is the default value.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public DeliveryChannelState()
        {
        }
        public static new DeliveryChannelState Empty => new DeliveryChannelState();
    }
}
