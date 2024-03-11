// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EipAnycast
{
    /// <summary>
    /// Provides a Eipanycast Anycast Eip Address Attachment resource.
    /// 
    /// For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://www.alibabacloud.com/help/en/anycast-eip/latest/api-eipanycast-2020-03-09-associateanycasteipaddress).
    /// 
    /// &gt; **NOTE:** Available since v1.113.0.
    /// 
    /// &gt; **NOTE:** The following regions support currently while Slb instance support bound.
    /// [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.Slb.GetZones.Invoke(new()
    ///     {
    ///         AvailableSlbAddressType = "vpc",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.0.0.0/8",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "10.1.0.0/16",
    ///         VpcId = defaultNetwork.Id,
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("defaultApplicationLoadBalancer", new()
    ///     {
    ///         AddressType = "intranet",
    ///         VswitchId = defaultSwitch.Id,
    ///         LoadBalancerName = name,
    ///         LoadBalancerSpec = "slb.s1.small",
    ///         MasterZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultAnycastEipAddress = new AliCloud.EipAnycast.AnycastEipAddress("defaultAnycastEipAddress", new()
    ///     {
    ///         AnycastEipAddressName = name,
    ///         ServiceLocation = "ChineseMainland",
    ///     });
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultAnycastEipAddressAttachment = new AliCloud.EipAnycast.AnycastEipAddressAttachment("defaultAnycastEipAddressAttachment", new()
    ///     {
    ///         BindInstanceId = defaultApplicationLoadBalancer.Id,
    ///         BindInstanceType = "SlbInstance",
    ///         BindInstanceRegionId = defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id),
    ///         AnycastId = defaultAnycastEipAddress.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Multiple Usage
    /// 
    /// &gt; **NOTE:**  Anycast EIP supports binding cloud resource instances in multiple regions. Only one cloud resource instance is supported as the default origin station, and the rest are normal origin stations. When no access point is specified or an access point is added, the access request is forwarded to the default origin by default.  If you are bound for the first time, the Default value of the binding mode is **Default * *. /li&gt; li&gt; If you are not binding for the first time, you can set the binding mode to **Default**, and the new Default origin will take effect. The original Default origin will be changed to a common origin.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var beijing = new AliCloud.Provider("beijing", new()
    ///     {
    ///         Region = "cn-beijing",
    ///     });
    /// 
    ///     var hangzhou = new AliCloud.Provider("hangzhou", new()
    ///     {
    ///         Region = "cn-hangzhou",
    ///     });
    /// 
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableDiskCategory = "cloud_efficiency",
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var defaultInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 1,
    ///         MemorySize = 2,
    ///     });
    /// 
    ///     var defaultVpc = new AliCloud.Vpc.Network("defaultVpc", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/16",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.beijing",
    ///     });
    /// 
    ///     var defaultVsw = new AliCloud.Vpc.Switch("defaultVsw", new()
    ///     {
    ///         VpcId = defaultVpc.Id,
    ///         CidrBlock = "192.168.0.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.beijing",
    ///     });
    /// 
    ///     var defaultuBsECI = new AliCloud.Ecs.SecurityGroup("defaultuBsECI", new()
    ///     {
    ///         VpcId = defaultVpc.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.beijing",
    ///     });
    /// 
    ///     var default9KDlN7 = new AliCloud.Ecs.Instance("default9KDlN7", new()
    ///     {
    ///         ImageId = defaultImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = defaultInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         InstanceName = name,
    ///         SecurityGroups = new[]
    ///         {
    ///             defaultuBsECI.Id,
    ///         },
    ///         AvailabilityZone = defaultVsw.ZoneId,
    ///         InstanceChargeType = "PostPaid",
    ///         SystemDiskCategory = "cloud_efficiency",
    ///         VswitchId = defaultVsw.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.beijing",
    ///     });
    /// 
    ///     var defaultXkpFRs = new AliCloud.EipAnycast.AnycastEipAddress("defaultXkpFRs", new()
    ///     {
    ///         ServiceLocation = "ChineseMainland",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    ///     var defaultVpc2 = new AliCloud.Vpc.Network("defaultVpc2", new()
    ///     {
    ///         VpcName = $"{name}6",
    ///         CidrBlock = "192.168.0.0/16",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    ///     var default2Zones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableDiskCategory = "cloud_efficiency",
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var default2Images = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var default2InstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = default2Zones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 1,
    ///         MemorySize = 2,
    ///     });
    /// 
    ///     var defaultdsVsw2 = new AliCloud.Vpc.Switch("defaultdsVsw2", new()
    ///     {
    ///         VpcId = defaultVpc2.Id,
    ///         CidrBlock = "192.168.0.0/24",
    ///         ZoneId = default2Zones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    ///     var defaultuBsECI2 = new AliCloud.Ecs.SecurityGroup("defaultuBsECI2", new()
    ///     {
    ///         VpcId = defaultVpc2.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    ///     var defaultEcs2 = new AliCloud.Ecs.Instance("defaultEcs2", new()
    ///     {
    ///         ImageId = default2Images.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = default2InstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         InstanceName = name,
    ///         SecurityGroups = new[]
    ///         {
    ///             defaultuBsECI2.Id,
    ///         },
    ///         AvailabilityZone = defaultdsVsw2.ZoneId,
    ///         InstanceChargeType = "PostPaid",
    ///         SystemDiskCategory = "cloud_efficiency",
    ///         VswitchId = defaultdsVsw2.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    ///     var defaultEfYBJY = new AliCloud.EipAnycast.AnycastEipAddressAttachment("defaultEfYBJY", new()
    ///     {
    ///         BindInstanceId = default9KDlN7.NetworkInterfaceId,
    ///         BindInstanceType = "NetworkInterface",
    ///         BindInstanceRegionId = "cn-beijing",
    ///         AnycastId = defaultXkpFRs.Id,
    ///         AssociationMode = "Default",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.beijing",
    ///     });
    /// 
    ///     var normal = new AliCloud.EipAnycast.AnycastEipAddressAttachment("normal", new()
    ///     {
    ///         BindInstanceId = defaultEcs2.NetworkInterfaceId,
    ///         BindInstanceType = "NetworkInterface",
    ///         BindInstanceRegionId = "cn-hangzhou",
    ///         AnycastId = defaultEfYBJY.AnycastId,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = "alicloud.hangzhou",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example &lt;anycast_id&gt;:&lt;bind_instance_id&gt;:&lt;bind_instance_region_id&gt;:&lt;bind_instance_type&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment")]
    public partial class AnycastEipAddressAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Anycast EIP instance.
        /// </summary>
        [Output("anycastId")]
        public Output<string> AnycastId { get; private set; } = null!;

        /// <summary>
        /// Binding mode, value:
        /// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
        /// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
        /// </summary>
        [Output("associationMode")]
        public Output<string> AssociationMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the cloud resource instance to be bound.
        /// </summary>
        [Output("bindInstanceId")]
        public Output<string> BindInstanceId { get; private set; } = null!;

        /// <summary>
        /// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
        /// </summary>
        [Output("bindInstanceRegionId")]
        public Output<string> BindInstanceRegionId { get; private set; } = null!;

        /// <summary>
        /// The type of the cloud resource instance to be bound. Value:
        /// - **SlbInstance**: a private network SLB instance.
        /// - **NetworkInterface**: ENI.
        /// </summary>
        [Output("bindInstanceType")]
        public Output<string> BindInstanceType { get; private set; } = null!;

        /// <summary>
        /// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
        /// </summary>
        [Output("bindTime")]
        public Output<string> BindTime { get; private set; } = null!;

        /// <summary>
        /// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `pop_locations` below.
        /// </summary>
        [Output("popLocations")]
        public Output<ImmutableArray<Outputs.AnycastEipAddressAttachmentPopLocation>> PopLocations { get; private set; } = null!;

        /// <summary>
        /// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string?> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AnycastEipAddressAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnycastEipAddressAttachment(string name, AnycastEipAddressAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, args ?? new AnycastEipAddressAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnycastEipAddressAttachment(string name, Input<string> id, AnycastEipAddressAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AnycastEipAddressAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnycastEipAddressAttachment Get(string name, Input<string> id, AnycastEipAddressAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AnycastEipAddressAttachment(name, id, state, options);
        }
    }

    public sealed class AnycastEipAddressAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Anycast EIP instance.
        /// </summary>
        [Input("anycastId", required: true)]
        public Input<string> AnycastId { get; set; } = null!;

        /// <summary>
        /// Binding mode, value:
        /// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
        /// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
        /// </summary>
        [Input("associationMode")]
        public Input<string>? AssociationMode { get; set; }

        /// <summary>
        /// The ID of the cloud resource instance to be bound.
        /// </summary>
        [Input("bindInstanceId", required: true)]
        public Input<string> BindInstanceId { get; set; } = null!;

        /// <summary>
        /// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
        /// </summary>
        [Input("bindInstanceRegionId", required: true)]
        public Input<string> BindInstanceRegionId { get; set; } = null!;

        /// <summary>
        /// The type of the cloud resource instance to be bound. Value:
        /// - **SlbInstance**: a private network SLB instance.
        /// - **NetworkInterface**: ENI.
        /// </summary>
        [Input("bindInstanceType", required: true)]
        public Input<string> BindInstanceType { get; set; } = null!;

        [Input("popLocations")]
        private InputList<Inputs.AnycastEipAddressAttachmentPopLocationArgs>? _popLocations;

        /// <summary>
        /// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `pop_locations` below.
        /// </summary>
        public InputList<Inputs.AnycastEipAddressAttachmentPopLocationArgs> PopLocations
        {
            get => _popLocations ?? (_popLocations = new InputList<Inputs.AnycastEipAddressAttachmentPopLocationArgs>());
            set => _popLocations = value;
        }

        /// <summary>
        /// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public AnycastEipAddressAttachmentArgs()
        {
        }
        public static new AnycastEipAddressAttachmentArgs Empty => new AnycastEipAddressAttachmentArgs();
    }

    public sealed class AnycastEipAddressAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Anycast EIP instance.
        /// </summary>
        [Input("anycastId")]
        public Input<string>? AnycastId { get; set; }

        /// <summary>
        /// Binding mode, value:
        /// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
        /// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
        /// </summary>
        [Input("associationMode")]
        public Input<string>? AssociationMode { get; set; }

        /// <summary>
        /// The ID of the cloud resource instance to be bound.
        /// </summary>
        [Input("bindInstanceId")]
        public Input<string>? BindInstanceId { get; set; }

        /// <summary>
        /// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
        /// </summary>
        [Input("bindInstanceRegionId")]
        public Input<string>? BindInstanceRegionId { get; set; }

        /// <summary>
        /// The type of the cloud resource instance to be bound. Value:
        /// - **SlbInstance**: a private network SLB instance.
        /// - **NetworkInterface**: ENI.
        /// </summary>
        [Input("bindInstanceType")]
        public Input<string>? BindInstanceType { get; set; }

        /// <summary>
        /// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
        /// </summary>
        [Input("bindTime")]
        public Input<string>? BindTime { get; set; }

        [Input("popLocations")]
        private InputList<Inputs.AnycastEipAddressAttachmentPopLocationGetArgs>? _popLocations;

        /// <summary>
        /// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `pop_locations` below.
        /// </summary>
        public InputList<Inputs.AnycastEipAddressAttachmentPopLocationGetArgs> PopLocations
        {
            get => _popLocations ?? (_popLocations = new InputList<Inputs.AnycastEipAddressAttachmentPopLocationGetArgs>());
            set => _popLocations = value;
        }

        /// <summary>
        /// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AnycastEipAddressAttachmentState()
        {
        }
        public static new AnycastEipAddressAttachmentState Empty => new AnycastEipAddressAttachmentState();
    }
}
