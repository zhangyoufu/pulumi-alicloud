// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mns
{
    public partial class TopicSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint has three format. Available values format:
        /// - HTTP Format: http://xxx.com/xxx
        /// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        /// - Email Format: mail:directmail:{MailAddress}
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The length should be shorter than 16.
        /// </summary>
        [Output("filterTag")]
        public Output<string?> FilterTag { get; private set; } = null!;

        /// <summary>
        /// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
        /// </summary>
        [Output("notifyContentFormat")]
        public Output<string?> NotifyContentFormat { get; private set; } = null!;

        /// <summary>
        /// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
        /// </summary>
        [Output("notifyStrategy")]
        public Output<string?> NotifyStrategy { get; private set; } = null!;

        /// <summary>
        /// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a TopicSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TopicSubscription(string name, TopicSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mns/topicSubscription:TopicSubscription", name, args ?? new TopicSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TopicSubscription(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mns/topicSubscription:TopicSubscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TopicSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TopicSubscription Get(string name, Input<string> id, TopicSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new TopicSubscription(name, id, state, options);
        }
    }

    public sealed class TopicSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint has three format. Available values format:
        /// - HTTP Format: http://xxx.com/xxx
        /// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        /// - Email Format: mail:directmail:{MailAddress}
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// The length should be shorter than 16.
        /// </summary>
        [Input("filterTag")]
        public Input<string>? FilterTag { get; set; }

        /// <summary>
        /// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
        /// </summary>
        [Input("notifyContentFormat")]
        public Input<string>? NotifyContentFormat { get; set; }

        /// <summary>
        /// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
        /// </summary>
        [Input("notifyStrategy")]
        public Input<string>? NotifyStrategy { get; set; }

        /// <summary>
        /// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public TopicSubscriptionArgs()
        {
        }
    }

    public sealed class TopicSubscriptionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint has three format. Available values format:
        /// - HTTP Format: http://xxx.com/xxx
        /// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
        /// - Email Format: mail:directmail:{MailAddress}
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The length should be shorter than 16.
        /// </summary>
        [Input("filterTag")]
        public Input<string>? FilterTag { get; set; }

        /// <summary>
        /// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
        /// </summary>
        [Input("notifyContentFormat")]
        public Input<string>? NotifyContentFormat { get; set; }

        /// <summary>
        /// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
        /// </summary>
        [Input("notifyStrategy")]
        public Input<string>? NotifyStrategy { get; set; }

        /// <summary>
        /// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public TopicSubscriptionState()
        {
        }
    }
}
