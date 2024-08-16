// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var example = new AliCloud.Vpc.Network("example", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.4.0.0/16",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("example", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.4.0.0/24",
    ///         VpcId = example.Id,
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///     });
    /// 
    ///     var exampleHAVip = new AliCloud.Vpc.HAVip("example", new()
    ///     {
    ///         VswitchId = exampleSwitch.Id,
    ///         Description = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The havip can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/hAVip:HAVip foo havip-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/hAVip:HAVip")]
    public partial class HAVip : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The elastic IP address (EIP) associated with the HAVIP.
        /// </summary>
        [Output("associatedEipAddresses")]
        public Output<ImmutableArray<string>> AssociatedEipAddresses { get; private set; } = null!;

        /// <summary>
        /// The type of the instance with which the HAVIP is associated. Valid values:
        /// - `EcsInstance`: an ECS instance.
        /// - `NetworkInterface`: an ENI.
        /// </summary>
        [Output("associatedInstanceType")]
        public Output<string> AssociatedInstanceType { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance with which the HAVIP is associated.
        /// </summary>
        [Output("associatedInstances")]
        public Output<ImmutableArray<string>> AssociatedInstances { get; private set; } = null!;

        /// <summary>
        /// The time when the HAVIP was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the HaVip instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the HAVIP.
        /// </summary>
        [Output("haVipId")]
        public Output<string> HaVipId { get; private set; } = null!;

        /// <summary>
        /// The name of the HAVIP.
        /// </summary>
        [Output("haVipName")]
        public Output<string> HaVipName { get; private set; } = null!;

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Output("havipName")]
        public Output<string> HavipName { get; private set; } = null!;

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the active instance that is associated with the HAVIP.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the HAVIP belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// (Available in v1.120.0+) The status of the HaVip instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The list of tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC to which the HAVIP belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The vswitch_id of the HaVip, the field can't be changed.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a HAVip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HAVip(string name, HAVipArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/hAVip:HAVip", name, args ?? new HAVipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HAVip(string name, Input<string> id, HAVipState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/hAVip:HAVip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HAVip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HAVip Get(string name, Input<string> id, HAVipState? state = null, CustomResourceOptions? options = null)
        {
            return new HAVip(name, id, state, options);
        }
    }

    public sealed class HAVipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the HaVip instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the HAVIP.
        /// </summary>
        [Input("haVipName")]
        public Input<string>? HaVipName { get; set; }

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Input("havipName")]
        public Input<string>? HavipName { get; set; }

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The ID of the resource group to which the HAVIP belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The vswitch_id of the HaVip, the field can't be changed.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public HAVipArgs()
        {
        }
        public static new HAVipArgs Empty => new HAVipArgs();
    }

    public sealed class HAVipState : global::Pulumi.ResourceArgs
    {
        [Input("associatedEipAddresses")]
        private InputList<string>? _associatedEipAddresses;

        /// <summary>
        /// The elastic IP address (EIP) associated with the HAVIP.
        /// </summary>
        public InputList<string> AssociatedEipAddresses
        {
            get => _associatedEipAddresses ?? (_associatedEipAddresses = new InputList<string>());
            set => _associatedEipAddresses = value;
        }

        /// <summary>
        /// The type of the instance with which the HAVIP is associated. Valid values:
        /// - `EcsInstance`: an ECS instance.
        /// - `NetworkInterface`: an ENI.
        /// </summary>
        [Input("associatedInstanceType")]
        public Input<string>? AssociatedInstanceType { get; set; }

        [Input("associatedInstances")]
        private InputList<string>? _associatedInstances;

        /// <summary>
        /// The ID of the instance with which the HAVIP is associated.
        /// </summary>
        public InputList<string> AssociatedInstances
        {
            get => _associatedInstances ?? (_associatedInstances = new InputList<string>());
            set => _associatedInstances = value;
        }

        /// <summary>
        /// The time when the HAVIP was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the HaVip instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the HAVIP.
        /// </summary>
        [Input("haVipId")]
        public Input<string>? HaVipId { get; set; }

        /// <summary>
        /// The name of the HAVIP.
        /// </summary>
        [Input("haVipName")]
        public Input<string>? HaVipName { get; set; }

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Input("havipName")]
        public Input<string>? HavipName { get; set; }

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The ID of the active instance that is associated with the HAVIP.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        /// <summary>
        /// The ID of the resource group to which the HAVIP belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// (Available in v1.120.0+) The status of the HaVip instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC to which the HAVIP belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The vswitch_id of the HaVip, the field can't be changed.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public HAVipState()
        {
        }
        public static new HAVipState Empty => new HAVipState();
    }
}
