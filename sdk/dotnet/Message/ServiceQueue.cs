// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Message
{
    /// <summary>
    /// Provides a Message Service Queue resource.
    /// 
    /// For information about Message Service Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/en/message-service/latest/createqueue).
    /// 
    /// &gt; **NOTE:** Available since v1.188.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
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
    ///     var @default = new AliCloud.Message.ServiceQueue("default", new()
    ///     {
    ///         DelaySeconds = 2,
    ///         PollingWaitSeconds = 2,
    ///         MessageRetentionPeriod = 566,
    ///         MaximumMessageSize = 1123,
    ///         VisibilityTimeout = 30,
    ///         QueueName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Message Service Queue can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:message/serviceQueue:ServiceQueue example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:message/serviceQueue:ServiceQueue")]
    public partial class ServiceQueue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Represents the time when the Queue was created.
        /// </summary>
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        /// <summary>
        /// This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
        /// </summary>
        [Output("delaySeconds")]
        public Output<int> DelaySeconds { get; private set; } = null!;

        /// <summary>
        /// Represents whether the log management function is enabled.
        /// </summary>
        [Output("loggingEnabled")]
        public Output<bool?> LoggingEnabled { get; private set; } = null!;

        /// <summary>
        /// Represents the maximum length of the message body sent to the Queue, in Byte.
        /// </summary>
        [Output("maximumMessageSize")]
        public Output<int> MaximumMessageSize { get; private set; } = null!;

        /// <summary>
        /// Represents the longest life time of the message in the Queue.
        /// </summary>
        [Output("messageRetentionPeriod")]
        public Output<int> MessageRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The longest waiting time for a Queue request when the number of messages is empty, in seconds.
        /// </summary>
        [Output("pollingWaitSeconds")]
        public Output<int> PollingWaitSeconds { get; private set; } = null!;

        /// <summary>
        /// Representative resources.
        /// </summary>
        [Output("queueName")]
        public Output<string> QueueName { get; private set; } = null!;

        /// <summary>
        /// Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
        /// </summary>
        [Output("visibilityTimeout")]
        public Output<int> VisibilityTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceQueue(string name, ServiceQueueArgs args, CustomResourceOptions? options = null)
            : base("alicloud:message/serviceQueue:ServiceQueue", name, args ?? new ServiceQueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceQueue(string name, Input<string> id, ServiceQueueState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:message/serviceQueue:ServiceQueue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceQueue Get(string name, Input<string> id, ServiceQueueState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceQueue(name, id, state, options);
        }
    }

    public sealed class ServiceQueueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
        /// </summary>
        [Input("delaySeconds")]
        public Input<int>? DelaySeconds { get; set; }

        /// <summary>
        /// Represents whether the log management function is enabled.
        /// </summary>
        [Input("loggingEnabled")]
        public Input<bool>? LoggingEnabled { get; set; }

        /// <summary>
        /// Represents the maximum length of the message body sent to the Queue, in Byte.
        /// </summary>
        [Input("maximumMessageSize")]
        public Input<int>? MaximumMessageSize { get; set; }

        /// <summary>
        /// Represents the longest life time of the message in the Queue.
        /// </summary>
        [Input("messageRetentionPeriod")]
        public Input<int>? MessageRetentionPeriod { get; set; }

        /// <summary>
        /// The longest waiting time for a Queue request when the number of messages is empty, in seconds.
        /// </summary>
        [Input("pollingWaitSeconds")]
        public Input<int>? PollingWaitSeconds { get; set; }

        /// <summary>
        /// Representative resources.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
        /// </summary>
        [Input("visibilityTimeout")]
        public Input<int>? VisibilityTimeout { get; set; }

        public ServiceQueueArgs()
        {
        }
        public static new ServiceQueueArgs Empty => new ServiceQueueArgs();
    }

    public sealed class ServiceQueueState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents the time when the Queue was created.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// This means that messages sent to the queue can only be consumed after the delay time set by this parameter, in seconds.
        /// </summary>
        [Input("delaySeconds")]
        public Input<int>? DelaySeconds { get; set; }

        /// <summary>
        /// Represents whether the log management function is enabled.
        /// </summary>
        [Input("loggingEnabled")]
        public Input<bool>? LoggingEnabled { get; set; }

        /// <summary>
        /// Represents the maximum length of the message body sent to the Queue, in Byte.
        /// </summary>
        [Input("maximumMessageSize")]
        public Input<int>? MaximumMessageSize { get; set; }

        /// <summary>
        /// Represents the longest life time of the message in the Queue.
        /// </summary>
        [Input("messageRetentionPeriod")]
        public Input<int>? MessageRetentionPeriod { get; set; }

        /// <summary>
        /// The longest waiting time for a Queue request when the number of messages is empty, in seconds.
        /// </summary>
        [Input("pollingWaitSeconds")]
        public Input<int>? PollingWaitSeconds { get; set; }

        /// <summary>
        /// Representative resources.
        /// </summary>
        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        /// <summary>
        /// Represents the duration after the message is removed from the Queue and changed from the Active state to the Inactive state.
        /// </summary>
        [Input("visibilityTimeout")]
        public Input<int>? VisibilityTimeout { get; set; }

        public ServiceQueueState()
        {
        }
        public static new ServiceQueueState Empty => new ServiceQueueState();
    }
}
