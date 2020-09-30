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
    /// Provides an ALIKAFKA instance resource.
    /// 
    /// &gt; **NOTE:** Available in 1.59.0+
    /// 
    /// &gt; **NOTE:** Creation or modification may took about 10-40 minutes.
    /// 
    /// &gt; **NOTE:** Only the following regions support create alikafka pre paid instance.
    /// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
    /// 
    /// &gt; **NOTE:** Only the following regions support create alikafka post paid instance(International account is not support to buy post paid instance currently).
    /// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
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
    ///         var instanceName = config.Get("instanceName") ?? "alikafkaInstanceName";
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
    ///             DeployType = 4,
    ///             IoMax = 20,
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
        /// </summary>
        [Output("deployType")]
        public Output<int> DeployType { get; private set; } = null!;

        /// <summary>
        /// The disk size of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
        /// </summary>
        [Output("diskType")]
        public Output<int> DiskType { get; private set; } = null!;

        /// <summary>
        /// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Output("eipMax")]
        public Output<int?> EipMax { get; private set; } = null!;

        /// <summary>
        /// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Output("ioMax")]
        public Output<int> IoMax { get; private set; } = null!;

        /// <summary>
        /// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
        /// </summary>
        [Output("paidType")]
        public Output<string?> PaidType { get; private set; } = null!;

        /// <summary>
        /// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
        /// </summary>
        [Output("securityGroup")]
        public Output<string?> SecurityGroup { get; private set; } = null!;

        /// <summary>
        /// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
        /// </summary>
        [Output("specType")]
        public Output<string?> SpecType { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
        /// </summary>
        [Output("topicQuota")]
        public Output<int> TopicQuota { get; private set; } = null!;

        /// <summary>
        /// The ID of attaching VPC to instance.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The ID of attaching vswitch to instance.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;

        /// <summary>
        /// The Zone to launch the kafka instance.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
        /// </summary>
        [Input("deployType", required: true)]
        public Input<int> DeployType { get; set; } = null!;

        /// <summary>
        /// The disk size of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
        /// </summary>
        [Input("diskType", required: true)]
        public Input<int> DiskType { get; set; } = null!;

        /// <summary>
        /// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("eipMax")]
        public Input<int>? EipMax { get; set; }

        /// <summary>
        /// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("ioMax", required: true)]
        public Input<int> IoMax { get; set; } = null!;

        /// <summary>
        /// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
        /// </summary>
        [Input("paidType")]
        public Input<string>? PaidType { get; set; }

        /// <summary>
        /// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
        /// </summary>
        [Input("specType")]
        public Input<string>? SpecType { get; set; }

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
        /// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
        /// </summary>
        [Input("topicQuota", required: true)]
        public Input<int> TopicQuota { get; set; } = null!;

        /// <summary>
        /// The ID of attaching vswitch to instance.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
        /// </summary>
        [Input("deployType")]
        public Input<int>? DeployType { get; set; }

        /// <summary>
        /// The disk size of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
        /// </summary>
        [Input("diskType")]
        public Input<int>? DiskType { get; set; }

        /// <summary>
        /// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("eipMax")]
        public Input<int>? EipMax { get; set; }

        /// <summary>
        /// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
        /// </summary>
        [Input("ioMax")]
        public Input<int>? IoMax { get; set; }

        /// <summary>
        /// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
        /// </summary>
        [Input("paidType")]
        public Input<string>? PaidType { get; set; }

        /// <summary>
        /// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
        /// </summary>
        [Input("specType")]
        public Input<string>? SpecType { get; set; }

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
        /// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
        /// </summary>
        [Input("topicQuota")]
        public Input<int>? TopicQuota { get; set; }

        /// <summary>
        /// The ID of attaching VPC to instance.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The ID of attaching vswitch to instance.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        /// <summary>
        /// The Zone to launch the kafka instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
    }
}
