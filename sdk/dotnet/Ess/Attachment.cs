// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess
{
    /// <summary>
    /// Attaches several ECS instances to a specified scaling group or remove them from it.
    /// 
    /// &gt; **NOTE:** ECS instances can be attached or remove only when the scaling group is active, and it has no scaling activity in progress.
    /// 
    /// &gt; **NOTE:** There are two types ECS instances in a scaling group: "AutoCreated" and "Attached". The total number of them can not larger than the scaling group "MaxSize".
    /// 
    /// &gt; **NOTE:** Available since v1.6.0.
    /// 
    /// ## Example Usage
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
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableDiskCategory = "cloud_efficiency",
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 2,
    ///         MemorySize = 4,
    ///     });
    /// 
    ///     var defaultImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = name,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("defaultSecurityGroup", new()
    ///     {
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultSecurityGroupRule = new AliCloud.Ecs.SecurityGroupRule("defaultSecurityGroupRule", new()
    ///     {
    ///         Type = "ingress",
    ///         IpProtocol = "tcp",
    ///         NicType = "intranet",
    ///         Policy = "accept",
    ///         PortRange = "22/22",
    ///         Priority = 1,
    ///         SecurityGroupId = defaultSecurityGroup.Id,
    ///         CidrIp = "172.16.0.0/24",
    ///     });
    /// 
    ///     var defaultScalingGroup = new AliCloud.Ess.ScalingGroup("defaultScalingGroup", new()
    ///     {
    ///         MinSize = 0,
    ///         MaxSize = 2,
    ///         ScalingGroupName = name,
    ///         RemovalPolicies = new[]
    ///         {
    ///             "OldestInstance",
    ///             "NewestInstance",
    ///         },
    ///         VswitchIds = new[]
    ///         {
    ///             defaultSwitch.Id,
    ///         },
    ///     });
    /// 
    ///     var defaultScalingConfiguration = new AliCloud.Ess.ScalingConfiguration("defaultScalingConfiguration", new()
    ///     {
    ///         ScalingGroupId = defaultScalingGroup.Id,
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = defaultInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         SecurityGroupId = defaultSecurityGroup.Id,
    ///         ForceDelete = true,
    ///         Active = true,
    ///         Enable = true,
    ///     });
    /// 
    ///     var defaultInstance = new List&lt;AliCloud.Ecs.Instance&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         defaultInstance.Add(new AliCloud.Ecs.Instance($"defaultInstance-{range.Value}", new()
    ///         {
    ///             ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///             InstanceType = defaultInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///             SecurityGroups = new[]
    ///             {
    ///                 defaultSecurityGroup.Id,
    ///             },
    ///             InternetChargeType = "PayByTraffic",
    ///             InternetMaxBandwidthOut = 10,
    ///             InstanceChargeType = "PostPaid",
    ///             SystemDiskCategory = "cloud_efficiency",
    ///             VswitchId = defaultSwitch.Id,
    ///             InstanceName = name,
    ///         }));
    ///     }
    ///     var defaultAttachment = new AliCloud.Ess.Attachment("defaultAttachment", new()
    ///     {
    ///         ScalingGroupId = defaultScalingGroup.Id,
    ///         InstanceIds = new[]
    ///         {
    ///             defaultInstance[0].Id,
    ///             defaultInstance[1].Id,
    ///         },
    ///         Force = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ESS attachment can be imported using the id or scaling group id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ess/attachment:Attachment example asg-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ess/attachment:Attachment")]
    public partial class Attachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
        /// 
        /// &gt; **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
        /// 
        /// &gt; **NOTE:** Restrictions on attaching ECS instances:
        /// 
        /// - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
        /// - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
        /// - The attached ECS instances must in the running state.
        /// - The attached ECS instances has not been attached to other scaling groups.
        /// - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// ID of the scaling group of a scaling configuration.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attachment(string name, AttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ess/attachment:Attachment", name, args ?? new AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attachment(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ess/attachment:Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attachment Get(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Attachment(name, id, state, options);
        }
    }

    public sealed class AttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
        /// 
        /// &gt; **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
        /// 
        /// &gt; **NOTE:** Restrictions on attaching ECS instances:
        /// 
        /// - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
        /// - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
        /// - The attached ECS instances must in the running state.
        /// - The attached ECS instances has not been attached to other scaling groups.
        /// - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the scaling group of a scaling configuration.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        public AttachmentArgs()
        {
        }
        public static new AttachmentArgs Empty => new AttachmentArgs();
    }

    public sealed class AttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false.
        /// 
        /// &gt; **NOTE:** "AutoCreated" ECS instance will be deleted after it is removed from scaling group, but "Attached" will be not.
        /// 
        /// &gt; **NOTE:** Restrictions on attaching ECS instances:
        /// 
        /// - The attached ECS instances and the scaling group must have the same region and network type(`Classic` or `VPC`).
        /// - The attached ECS instances and the instance with active scaling configurations must have the same instance type.
        /// - The attached ECS instances must in the running state.
        /// - The attached ECS instances has not been attached to other scaling groups.
        /// - The attached ECS instances supports Subscription and Pay-As-You-Go payment methods.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// ID of the scaling group of a scaling configuration.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public AttachmentState()
        {
        }
        public static new AttachmentState Empty => new AttachmentState();
    }
}
