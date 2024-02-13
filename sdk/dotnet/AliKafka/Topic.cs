// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.AliKafka
{
    /// <summary>
    /// Provides an ALIKAFKA topic resource, see [What is Alikafka topic ](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/api-alikafka-2019-09-16-createtopic).
    /// 
    /// &gt; **NOTE:** Available since v1.56.0.
    /// 
    /// &gt; **NOTE:**  Only the following regions support create alikafka topic.
    /// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var instanceName = config.Get("instanceName") ?? "tf-example";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.AliKafka.Instance("defaultInstance", new()
    ///     {
    ///         PartitionNum = 50,
    ///         DiskType = 1,
    ///         DiskSize = 500,
    ///         DeployType = 5,
    ///         IoMax = 20,
    ///         VswitchId = defaultSwitch.Id,
    ///         SecurityGroup = defaultSecurityGroup.Id,
    ///     });
    /// 
    ///     var defaultTopic = new AliCloud.AliKafka.Topic("defaultTopic", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         TopicName = "example-topic",
    ///         LocalTopic = false,
    ///         CompactTopic = false,
    ///         PartitionNum = 12,
    ///         Remark = "dafault_kafka_topic_remark",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ALIKAFKA TOPIC can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:alikafka/topic:Topic topic alikafka_post-cn-123455abc:topicName
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alikafka/topic:Topic")]
    public partial class Topic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the topic is compactTopic or not. Compact topic must be a localTopic.
        /// </summary>
        [Output("compactTopic")]
        public Output<bool?> CompactTopic { get; private set; } = null!;

        /// <summary>
        /// InstanceId of your Kafka resource, the topic will create in this instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Whether the topic is localTopic or not.
        /// </summary>
        [Output("localTopic")]
        public Output<bool?> LocalTopic { get; private set; } = null!;

        /// <summary>
        /// The number of partitions of the topic. The number should between 1 and 48.
        /// </summary>
        [Output("partitionNum")]
        public Output<int?> PartitionNum { get; private set; } = null!;

        /// <summary>
        /// This attribute is a concise description of topic. The length cannot exceed 64.
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 249 characters.
        /// </summary>
        [Output("topic")]
        public Output<string> TopicName { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the topic is compactTopic or not. Compact topic must be a localTopic.
        /// </summary>
        [Input("compactTopic")]
        public Input<bool>? CompactTopic { get; set; }

        /// <summary>
        /// InstanceId of your Kafka resource, the topic will create in this instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Whether the topic is localTopic or not.
        /// </summary>
        [Input("localTopic")]
        public Input<bool>? LocalTopic { get; set; }

        /// <summary>
        /// The number of partitions of the topic. The number should between 1 and 48.
        /// </summary>
        [Input("partitionNum")]
        public Input<int>? PartitionNum { get; set; }

        /// <summary>
        /// This attribute is a concise description of topic. The length cannot exceed 64.
        /// </summary>
        [Input("remark", required: true)]
        public Input<string> Remark { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 249 characters.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public TopicArgs()
        {
        }
        public static new TopicArgs Empty => new TopicArgs();
    }

    public sealed class TopicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the topic is compactTopic or not. Compact topic must be a localTopic.
        /// </summary>
        [Input("compactTopic")]
        public Input<bool>? CompactTopic { get; set; }

        /// <summary>
        /// InstanceId of your Kafka resource, the topic will create in this instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Whether the topic is localTopic or not.
        /// </summary>
        [Input("localTopic")]
        public Input<bool>? LocalTopic { get; set; }

        /// <summary>
        /// The number of partitions of the topic. The number should between 1 and 48.
        /// </summary>
        [Input("partitionNum")]
        public Input<int>? PartitionNum { get; set; }

        /// <summary>
        /// This attribute is a concise description of topic. The length cannot exceed 64.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 249 characters.
        /// </summary>
        [Input("topic")]
        public Input<string>? TopicName { get; set; }

        public TopicState()
        {
        }
        public static new TopicState Empty => new TopicState();
    }
}
