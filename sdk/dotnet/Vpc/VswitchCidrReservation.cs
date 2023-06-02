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
    /// Provides a Vpc Vswitch Cidr Reservation resource. The reserved network segment of the vswitch. This resource type can be used only in ap-southeast region.
    /// 
    /// For information about Vpc Vswitch Cidr Reservation and how to use it, see [What is Vswitch Cidr Reservation](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/610154).
    /// 
    /// &gt; **NOTE:** Available in v1.205.0+.
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
    ///     var name = config.Get("name") ?? "tf-testacc-example";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultVpc = new AliCloud.Vpc.Network("defaultVpc", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.0.0.0/8",
    ///     });
    /// 
    ///     var defaultVSwitch = new AliCloud.Vpc.Switch("defaultVSwitch", new()
    ///     {
    ///         VpcId = defaultVpc.Id,
    ///         CidrBlock = "10.0.0.0/20",
    ///         VswitchName = $"{name}1",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultVswitchCidrReservation = new AliCloud.Vpc.VswitchCidrReservation("defaultVswitchCidrReservation", new()
    ///     {
    ///         IpVersion = "IPv4",
    ///         VswitchId = defaultVSwitch.Id,
    ///         CidrReservationDescription = "test",
    ///         CidrReservationCidr = "10.0.10.0/24",
    ///         VswitchCidrReservationName = name,
    ///         CidrReservationType = "Prefix",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Vpc Vswitch Cidr Reservation can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/vswitchCidrReservation:VswitchCidrReservation example &lt;vswitch_id&gt;:&lt;vswitch_cidr_reservation_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/vswitchCidrReservation:VswitchCidrReservation")]
    public partial class VswitchCidrReservation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Reserved network segment CIdrBlock.
        /// </summary>
        [Output("cidrReservationCidr")]
        public Output<string> CidrReservationCidr { get; private set; } = null!;

        /// <summary>
        /// The description of the reserved CIDR block.
        /// </summary>
        [Output("cidrReservationDescription")]
        public Output<string?> CidrReservationDescription { get; private set; } = null!;

        /// <summary>
        /// Reserved segment mask.
        /// </summary>
        [Output("cidrReservationMask")]
        public Output<string?> CidrReservationMask { get; private set; } = null!;

        /// <summary>
        /// Reserved CIDR Block Type.Valid values: `Prefix`. Default value: Prefix.
        /// </summary>
        [Output("cidrReservationType")]
        public Output<string> CidrReservationType { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Reserved ip version of network segment, valid values: `IPv4`, `IPv6`, default IPv4.
        /// </summary>
        [Output("ipVersion")]
        public Output<string> IpVersion { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of the vpc instance to which the reserved CIDR block belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The resource attribute field of the resource ID.
        /// </summary>
        [Output("vswitchCidrReservationId")]
        public Output<string> VswitchCidrReservationId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("vswitchCidrReservationName")]
        public Output<string?> VswitchCidrReservationName { get; private set; } = null!;

        /// <summary>
        /// The Id of the switch instance.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a VswitchCidrReservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VswitchCidrReservation(string name, VswitchCidrReservationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/vswitchCidrReservation:VswitchCidrReservation", name, args ?? new VswitchCidrReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VswitchCidrReservation(string name, Input<string> id, VswitchCidrReservationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/vswitchCidrReservation:VswitchCidrReservation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VswitchCidrReservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VswitchCidrReservation Get(string name, Input<string> id, VswitchCidrReservationState? state = null, CustomResourceOptions? options = null)
        {
            return new VswitchCidrReservation(name, id, state, options);
        }
    }

    public sealed class VswitchCidrReservationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reserved network segment CIdrBlock.
        /// </summary>
        [Input("cidrReservationCidr")]
        public Input<string>? CidrReservationCidr { get; set; }

        /// <summary>
        /// The description of the reserved CIDR block.
        /// </summary>
        [Input("cidrReservationDescription")]
        public Input<string>? CidrReservationDescription { get; set; }

        /// <summary>
        /// Reserved segment mask.
        /// </summary>
        [Input("cidrReservationMask")]
        public Input<string>? CidrReservationMask { get; set; }

        /// <summary>
        /// Reserved CIDR Block Type.Valid values: `Prefix`. Default value: Prefix.
        /// </summary>
        [Input("cidrReservationType")]
        public Input<string>? CidrReservationType { get; set; }

        /// <summary>
        /// Reserved ip version of network segment, valid values: `IPv4`, `IPv6`, default IPv4.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("vswitchCidrReservationName")]
        public Input<string>? VswitchCidrReservationName { get; set; }

        /// <summary>
        /// The Id of the switch instance.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public VswitchCidrReservationArgs()
        {
        }
        public static new VswitchCidrReservationArgs Empty => new VswitchCidrReservationArgs();
    }

    public sealed class VswitchCidrReservationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reserved network segment CIdrBlock.
        /// </summary>
        [Input("cidrReservationCidr")]
        public Input<string>? CidrReservationCidr { get; set; }

        /// <summary>
        /// The description of the reserved CIDR block.
        /// </summary>
        [Input("cidrReservationDescription")]
        public Input<string>? CidrReservationDescription { get; set; }

        /// <summary>
        /// Reserved segment mask.
        /// </summary>
        [Input("cidrReservationMask")]
        public Input<string>? CidrReservationMask { get; set; }

        /// <summary>
        /// Reserved CIDR Block Type.Valid values: `Prefix`. Default value: Prefix.
        /// </summary>
        [Input("cidrReservationType")]
        public Input<string>? CidrReservationType { get; set; }

        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Reserved ip version of network segment, valid values: `IPv4`, `IPv6`, default IPv4.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the vpc instance to which the reserved CIDR block belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The resource attribute field of the resource ID.
        /// </summary>
        [Input("vswitchCidrReservationId")]
        public Input<string>? VswitchCidrReservationId { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("vswitchCidrReservationName")]
        public Input<string>? VswitchCidrReservationName { get; set; }

        /// <summary>
        /// The Id of the switch instance.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public VswitchCidrReservationState()
        {
        }
        public static new VswitchCidrReservationState Empty => new VswitchCidrReservationState();
    }
}
