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
    /// Provides a Sag qos policy resource.
    /// You need to create a QoS policy to set priorities, rate limits, and quintuple rules for different messages.
    /// 
    /// For information about Sag Qos Policy and how to use it, see [What is Qos Policy](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createqospolicy).
    /// 
    /// &gt; **NOTE:** Available since v1.60.0.
    /// 
    /// &gt; **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
    /// 
    /// ## Import
    /// 
    /// The Sag Qos Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:rocketmq/qosPolicy:QosPolicy example qos-abc123456:qospy-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:rocketmq/qosPolicy:QosPolicy")]
    public partial class QosPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the QoS policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination CIDR block.
        /// </summary>
        [Output("destCidr")]
        public Output<string> DestCidr { get; private set; } = null!;

        /// <summary>
        /// The destination port range.
        /// </summary>
        [Output("destPortRange")]
        public Output<string> DestPortRange { get; private set; } = null!;

        /// <summary>
        /// The expiration time of the quintuple rule.
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// The transport layer protocol.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The name of the QoS policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The instance ID of the QoS policy to which the quintuple rule is created.
        /// </summary>
        [Output("qosId")]
        public Output<string> QosId { get; private set; } = null!;

        /// <summary>
        /// The source CIDR block.
        /// </summary>
        [Output("sourceCidr")]
        public Output<string> SourceCidr { get; private set; } = null!;

        /// <summary>
        /// The source port range of the transport layer.
        /// </summary>
        [Output("sourcePortRange")]
        public Output<string> SourcePortRange { get; private set; } = null!;

        /// <summary>
        /// The time when the quintuple rule takes effect.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a QosPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosPolicy(string name, QosPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/qosPolicy:QosPolicy", name, args ?? new QosPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QosPolicy(string name, Input<string> id, QosPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:rocketmq/qosPolicy:QosPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosPolicy Get(string name, Input<string> id, QosPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new QosPolicy(name, id, state, options);
        }
    }

    public sealed class QosPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the QoS policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination CIDR block.
        /// </summary>
        [Input("destCidr", required: true)]
        public Input<string> DestCidr { get; set; } = null!;

        /// <summary>
        /// The destination port range.
        /// </summary>
        [Input("destPortRange", required: true)]
        public Input<string> DestPortRange { get; set; } = null!;

        /// <summary>
        /// The expiration time of the quintuple rule.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The transport layer protocol.
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// The name of the QoS policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// The instance ID of the QoS policy to which the quintuple rule is created.
        /// </summary>
        [Input("qosId", required: true)]
        public Input<string> QosId { get; set; } = null!;

        /// <summary>
        /// The source CIDR block.
        /// </summary>
        [Input("sourceCidr", required: true)]
        public Input<string> SourceCidr { get; set; } = null!;

        /// <summary>
        /// The source port range of the transport layer.
        /// </summary>
        [Input("sourcePortRange", required: true)]
        public Input<string> SourcePortRange { get; set; } = null!;

        /// <summary>
        /// The time when the quintuple rule takes effect.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public QosPolicyArgs()
        {
        }
        public static new QosPolicyArgs Empty => new QosPolicyArgs();
    }

    public sealed class QosPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the QoS policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination CIDR block.
        /// </summary>
        [Input("destCidr")]
        public Input<string>? DestCidr { get; set; }

        /// <summary>
        /// The destination port range.
        /// </summary>
        [Input("destPortRange")]
        public Input<string>? DestPortRange { get; set; }

        /// <summary>
        /// The expiration time of the quintuple rule.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The transport layer protocol.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The name of the QoS policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The instance ID of the QoS policy to which the quintuple rule is created.
        /// </summary>
        [Input("qosId")]
        public Input<string>? QosId { get; set; }

        /// <summary>
        /// The source CIDR block.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// The source port range of the transport layer.
        /// </summary>
        [Input("sourcePortRange")]
        public Input<string>? SourcePortRange { get; set; }

        /// <summary>
        /// The time when the quintuple rule takes effect.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public QosPolicyState()
        {
        }
        public static new QosPolicyState Empty => new QosPolicyState();
    }
}
