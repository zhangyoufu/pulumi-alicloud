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
    /// Provides a VPC Nat Ip Cidr resource.
    /// 
    /// For information about VPC Nat Ip Cidr and how to use it, see [What is Nat Ip Cidr](https://www.alibabacloud.com/help/doc-detail/281972.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.136.0+.
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
    ///     var example = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var exampleNetwork = new AliCloud.Vpc.Network("example", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var exampleSwitch = new AliCloud.Vpc.Switch("example", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = example.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = "terraform-example",
    ///     });
    /// 
    ///     var exampleNatGateway = new AliCloud.Vpc.NatGateway("example", new()
    ///     {
    ///         VpcId = exampleNetwork.Id,
    ///         InternetChargeType = "PayByLcu",
    ///         NatGatewayName = "terraform-example",
    ///         Description = "terraform-example",
    ///         NatType = "Enhanced",
    ///         VswitchId = exampleSwitch.Id,
    ///         NetworkType = "intranet",
    ///     });
    /// 
    ///     var exampleNatIpCidr = new AliCloud.Vpc.NatIpCidr("example", new()
    ///     {
    ///         NatGatewayId = exampleNatGateway.Id,
    ///         NatIpCidrName = "terraform-example",
    ///         NatIpCidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPC Nat Ip Cidr can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/natIpCidr:NatIpCidr example &lt;nat_gateway_id&gt;:&lt;nat_ip_cidr&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/natIpCidr:NatIpCidr")]
    public partial class NatIpCidr : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` and `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        /// </summary>
        [Output("natGatewayId")]
        public Output<string> NatGatewayId { get; private set; } = null!;

        /// <summary>
        /// The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        /// </summary>
        [Output("natIpCidr")]
        public Output<string?> NatIpCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Output("natIpCidrDescription")]
        public Output<string?> NatIpCidrDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Output("natIpCidrName")]
        public Output<string?> NatIpCidrName { get; private set; } = null!;

        /// <summary>
        /// The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a NatIpCidr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NatIpCidr(string name, NatIpCidrArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/natIpCidr:NatIpCidr", name, args ?? new NatIpCidrArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NatIpCidr(string name, Input<string> id, NatIpCidrState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/natIpCidr:NatIpCidr", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NatIpCidr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NatIpCidr Get(string name, Input<string> id, NatIpCidrState? state = null, CustomResourceOptions? options = null)
        {
            return new NatIpCidr(name, id, state, options);
        }
    }

    public sealed class NatIpCidrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        /// </summary>
        [Input("natGatewayId", required: true)]
        public Input<string> NatGatewayId { get; set; } = null!;

        /// <summary>
        /// The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        /// </summary>
        [Input("natIpCidr")]
        public Input<string>? NatIpCidrBlock { get; set; }

        /// <summary>
        /// The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("natIpCidrDescription")]
        public Input<string>? NatIpCidrDescription { get; set; }

        /// <summary>
        /// The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("natIpCidrName")]
        public Input<string>? NatIpCidrName { get; set; }

        public NatIpCidrArgs()
        {
        }
        public static new NatIpCidrArgs Empty => new NatIpCidrArgs();
    }

    public sealed class NatIpCidrState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` and `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        /// </summary>
        [Input("natIpCidr")]
        public Input<string>? NatIpCidrBlock { get; set; }

        /// <summary>
        /// The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("natIpCidrDescription")]
        public Input<string>? NatIpCidrDescription { get; set; }

        /// <summary>
        /// The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("natIpCidrName")]
        public Input<string>? NatIpCidrName { get; set; }

        /// <summary>
        /// The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public NatIpCidrState()
        {
        }
        public static new NatIpCidrState Empty => new NatIpCidrState();
    }
}
