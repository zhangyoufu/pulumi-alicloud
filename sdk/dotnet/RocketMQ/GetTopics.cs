// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    public static class GetTopics
    {
        /// <summary>
        /// This data source provides a list of ONS Topics in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.53.0+
        /// 
        /// ## Example Usage
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
        ///     var name = config.Get("name") ?? "onsInstanceName";
        ///     var topic = config.Get("topic") ?? "onsTopicDatasourceName";
        ///     var defaultInstance = new AliCloud.RocketMQ.Instance("defaultInstance", new()
        ///     {
        ///         InstanceName = name,
        ///         Remark = "default_ons_instance_remark",
        ///     });
        /// 
        ///     var defaultTopic = new AliCloud.RocketMQ.Topic("defaultTopic", new()
        ///     {
        ///         TopicName = topic,
        ///         InstanceId = defaultInstance.Id,
        ///         MessageType = 0,
        ///         Remark = "dafault_ons_topic_remark",
        ///     });
        /// 
        ///     var topicsDs = AliCloud.RocketMQ.GetTopics.Invoke(new()
        ///     {
        ///         InstanceId = defaultTopic.InstanceId,
        ///         NameRegex = topic,
        ///         OutputFile = "topics.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTopicName"] = topicsDs.Apply(getTopicsResult =&gt; getTopicsResult.Topics[0]?.TopicName),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTopicsResult> InvokeAsync(GetTopicsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTopicsResult>("alicloud:rocketmq/getTopics:getTopics", args ?? new GetTopicsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of ONS Topics in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.53.0+
        /// 
        /// ## Example Usage
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
        ///     var name = config.Get("name") ?? "onsInstanceName";
        ///     var topic = config.Get("topic") ?? "onsTopicDatasourceName";
        ///     var defaultInstance = new AliCloud.RocketMQ.Instance("defaultInstance", new()
        ///     {
        ///         InstanceName = name,
        ///         Remark = "default_ons_instance_remark",
        ///     });
        /// 
        ///     var defaultTopic = new AliCloud.RocketMQ.Topic("defaultTopic", new()
        ///     {
        ///         TopicName = topic,
        ///         InstanceId = defaultInstance.Id,
        ///         MessageType = 0,
        ///         Remark = "dafault_ons_topic_remark",
        ///     });
        /// 
        ///     var topicsDs = AliCloud.RocketMQ.GetTopics.Invoke(new()
        ///     {
        ///         InstanceId = defaultTopic.InstanceId,
        ///         NameRegex = topic,
        ///         OutputFile = "topics.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTopicName"] = topicsDs.Apply(getTopicsResult =&gt; getTopicsResult.Topics[0]?.TopicName),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTopicsResult> Invoke(GetTopicsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicsResult>("alicloud:rocketmq/getTopics:getTopics", args ?? new GetTopicsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicsArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of topic IDs to filter results.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the ONS Instance that owns the topics.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the topic name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetTopicsArgs()
        {
        }
        public static new GetTopicsArgs Empty => new GetTopicsArgs();
    }

    public sealed class GetTopicsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of topic IDs to filter results.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// ID of the ONS Instance that owns the topics.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the topic name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetTopicsInvokeArgs()
        {
        }
        public static new GetTopicsInvokeArgs Empty => new GetTopicsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTopicsResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of topic names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;
        /// <summary>
        /// A list of topics. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicsTopicResult> Topics;

        [OutputConstructor]
        private GetTopicsResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableDictionary<string, object>? tags,

            ImmutableArray<Outputs.GetTopicsTopicResult> topics)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Tags = tags;
            Topics = topics;
        }
    }
}
