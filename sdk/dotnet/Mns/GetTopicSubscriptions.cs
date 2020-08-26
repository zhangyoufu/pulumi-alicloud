// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mns
{
    public static class GetTopicSubscriptions
    {
        /// <summary>
        /// This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var subscriptions = Output.Create(AliCloud.Mns.GetTopicSubscriptions.InvokeAsync(new AliCloud.Mns.GetTopicSubscriptionsArgs
        ///         {
        ///             NamePrefix = "tf-",
        ///             TopicName = "topic_name",
        ///         }));
        ///         this.FirstTopicSubscriptionId = subscriptions.Apply(subscriptions =&gt; subscriptions.Subscriptions[0].Id);
        ///     }
        /// 
        ///     [Output("firstTopicSubscriptionId")]
        ///     public Output&lt;string&gt; FirstTopicSubscriptionId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTopicSubscriptionsResult> InvokeAsync(GetTopicSubscriptionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTopicSubscriptionsResult>("alicloud:mns/getTopicSubscriptions:getTopicSubscriptions", args ?? new GetTopicSubscriptionsArgs(), options.WithVersion());
    }


    public sealed class GetTopicSubscriptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A string to filter resulting subscriptions of the topic by their name prefixs.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
        /// </summary>
        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetTopicSubscriptionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTopicSubscriptionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NamePrefix;
        /// <summary>
        /// A list of subscription names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of subscriptions. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicSubscriptionsSubscriptionResult> Subscriptions;
        public readonly string TopicName;

        [OutputConstructor]
        private GetTopicSubscriptionsResult(
            string id,

            string? namePrefix,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetTopicSubscriptionsSubscriptionResult> subscriptions,

            string topicName)
        {
            Id = id;
            NamePrefix = namePrefix;
            Names = names;
            OutputFile = outputFile;
            Subscriptions = subscriptions;
            TopicName = topicName;
        }
    }
}
