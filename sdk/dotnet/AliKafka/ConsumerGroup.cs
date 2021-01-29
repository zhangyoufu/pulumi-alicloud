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
    /// Provides an ALIKAFKA consumer group resource.
    /// 
    /// &gt; **NOTE:** Available in 1.56.0+
    /// 
    /// &gt; **NOTE:**  Only the following regions support create alikafka consumer group.
    /// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`ap-southeast-1`,`ap-south-1`,`ap-southeast-5`]
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var consumerId = config.Get("consumerId") ?? "CID-alikafkaGroupDatasourceName";
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableResourceCreation = "VSwitch",
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///             CidrBlock = "172.16.0.0/24",
    ///             AvailabilityZone = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///         });
    ///         var defaultInstance = new AliCloud.AliKafka.Instance("defaultInstance", new AliCloud.AliKafka.InstanceArgs
    ///         {
    ///             TopicQuota = 50,
    ///             DiskType = 1,
    ///             DiskSize = 500,
    ///             DeployType = 5,
    ///             IoMax = 20,
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///         var defaultConsumerGroup = new AliCloud.AliKafka.ConsumerGroup("defaultConsumerGroup", new AliCloud.AliKafka.ConsumerGroupArgs
    ///         {
    ///             ConsumerId = consumerId,
    ///             InstanceId = defaultInstance.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ALIKAFKA GROUP can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alikafka/consumerGroup:ConsumerGroup group alikafka_post-cn-123455abc:consumerId
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alikafka/consumerGroup:ConsumerGroup")]
    public partial class ConsumerGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the consumer group. The length cannot exceed 64 characters.
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerGroup(string name, ConsumerGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/consumerGroup:ConsumerGroup", name, args ?? new ConsumerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerGroup(string name, Input<string> id, ConsumerGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/consumerGroup:ConsumerGroup", name, state, MakeResourceOptions(options, id))
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

    public sealed class ConsumerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the consumer group. The length cannot exceed 64 characters.
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

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

        public ConsumerGroupArgs()
        {
        }
    }

    public sealed class ConsumerGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the consumer group. The length cannot exceed 64 characters.
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

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

        public ConsumerGroupState()
        {
        }
    }
}
