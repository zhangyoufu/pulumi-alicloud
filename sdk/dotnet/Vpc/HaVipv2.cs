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
    /// Provides a Vpc Ha Vip resource. Highly available virtual IP
    /// 
    /// For information about Vpc Ha Vip and how to use it, see [What is Ha Vip](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createhavip).
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
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
    ///     var name = config.Get("name") ?? "tf-testacc-example";
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultVpc = new AliCloud.Vpc.Network("defaultVpc", new()
    ///     {
    ///         Description = "tf-test-acc-vpc",
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    ///     var defaultVswitch = new AliCloud.Vpc.Switch("defaultVswitch", new()
    ///     {
    ///         VpcId = defaultVpc.Id,
    ///         CidrBlock = "192.168.0.0/21",
    ///         VswitchName = $"{name}1",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///         Description = "tf-testacc-vswitch",
    ///     });
    /// 
    ///     var defaultRg = new AliCloud.ResourceManager.ResourceGroup("defaultRg", new()
    ///     {
    ///         DisplayName = "tf-testacc-rg819",
    ///         ResourceGroupName = $"{name}2",
    ///     });
    /// 
    ///     var changeRg = new AliCloud.ResourceManager.ResourceGroup("changeRg", new()
    ///     {
    ///         DisplayName = "tf-testacc-changerg670",
    ///         ResourceGroupName = $"{name}3",
    ///     });
    /// 
    ///     var defaultHaVipv2 = new AliCloud.Vpc.HaVipv2("default", new()
    ///     {
    ///         Description = "test",
    ///         VswitchId = defaultVswitch.Id,
    ///         HaVipName = name,
    ///         IpAddress = "192.168.1.101",
    ///         ResourceGroupId = defaultRg.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Vpc Ha Vip can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/haVipv2:HaVipv2 example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/haVipv2:HaVipv2")]
    public partial class HaVipv2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// EIP bound to HaVip.
        /// </summary>
        [Output("associatedEipAddresses")]
        public Output<ImmutableArray<string>> AssociatedEipAddresses { get; private set; } = null!;

        /// <summary>
        /// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
        /// </summary>
        [Output("associatedInstanceType")]
        public Output<string> AssociatedInstanceType { get; private set; } = null!;

        /// <summary>
        /// An ECS instance that is bound to HaVip.
        /// </summary>
        [Output("associatedInstances")]
        public Output<ImmutableArray<string>> AssociatedInstances { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the HaVip instance. The length is 2 to 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource.
        /// </summary>
        [Output("haVipId")]
        public Output<string> HaVipId { get; private set; } = null!;

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Output("haVipName")]
        public Output<string> HaVipName { get; private set; } = null!;

        /// <summary>
        /// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
        /// </summary>
        [Output("havipName")]
        public Output<string> HavipName { get; private set; } = null!;

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The primary instance ID bound to HaVip.
        /// </summary>
        [Output("masterInstanceId")]
        public Output<string> MasterInstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of this resource instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of HaVip.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC ID to which the HaVip instance belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The switch ID to which the HaVip instance belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a HaVipv2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HaVipv2(string name, HaVipv2Args args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/haVipv2:HaVipv2", name, args ?? new HaVipv2Args(), MakeResourceOptions(options, ""))
        {
        }

        private HaVipv2(string name, Input<string> id, HaVipv2State? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/haVipv2:HaVipv2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HaVipv2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HaVipv2 Get(string name, Input<string> id, HaVipv2State? state = null, CustomResourceOptions? options = null)
        {
            return new HaVipv2(name, id, state, options);
        }
    }

    public sealed class HaVipv2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the HaVip instance. The length is 2 to 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Input("haVipName")]
        public Input<string>? HaVipName { get; set; }

        /// <summary>
        /// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
        /// </summary>
        [Input("havipName")]
        public Input<string>? HavipName { get; set; }

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of HaVip.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The switch ID to which the HaVip instance belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public HaVipv2Args()
        {
        }
        public static new HaVipv2Args Empty => new HaVipv2Args();
    }

    public sealed class HaVipv2State : global::Pulumi.ResourceArgs
    {
        [Input("associatedEipAddresses")]
        private InputList<string>? _associatedEipAddresses;

        /// <summary>
        /// EIP bound to HaVip.
        /// </summary>
        public InputList<string> AssociatedEipAddresses
        {
            get => _associatedEipAddresses ?? (_associatedEipAddresses = new InputList<string>());
            set => _associatedEipAddresses = value;
        }

        /// <summary>
        /// The type of the instance that is bound to the HaVip. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
        /// </summary>
        [Input("associatedInstanceType")]
        public Input<string>? AssociatedInstanceType { get; set; }

        [Input("associatedInstances")]
        private InputList<string>? _associatedInstances;

        /// <summary>
        /// An ECS instance that is bound to HaVip.
        /// </summary>
        public InputList<string> AssociatedInstances
        {
            get => _associatedInstances ?? (_associatedInstances = new InputList<string>());
            set => _associatedInstances = value;
        }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the HaVip instance. The length is 2 to 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the resource.
        /// </summary>
        [Input("haVipId")]
        public Input<string>? HaVipId { get; set; }

        /// <summary>
        /// The name of the HaVip instance.
        /// </summary>
        [Input("haVipName")]
        public Input<string>? HaVipName { get; set; }

        /// <summary>
        /// Field 'havip_name' has been deprecated from provider version 1.205.0. New field 'ha_vip_name' instead.
        /// </summary>
        [Input("havipName")]
        public Input<string>? HavipName { get; set; }

        /// <summary>
        /// The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The primary instance ID bound to HaVip.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of this resource instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of HaVip.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID to which the HaVip instance belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The switch ID to which the HaVip instance belongs.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public HaVipv2State()
        {
        }
        public static new HaVipv2State Empty => new HaVipv2State();
    }
}
