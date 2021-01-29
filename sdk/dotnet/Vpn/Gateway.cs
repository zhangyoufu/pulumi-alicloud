// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn
{
    /// <summary>
    /// ## Import
    /// 
    /// VPN gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpn/gateway:Gateway example vpn-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpn/gateway:Gateway")]
    public partial class Gateway : Pulumi.CustomResource
    {
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The business status of the VPN gateway.
        /// </summary>
        [Output("businessStatus")]
        public Output<string> BusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The description of the VPN instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
        /// </summary>
        [Output("enableIpsec")]
        public Output<bool?> EnableIpsec { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
        /// </summary>
        [Output("enableSsl")]
        public Output<bool?> EnableSsl { get; private set; } = null!;

        /// <summary>
        /// The charge type for instance. If it is an international site account, the valid value is PostPaid, otherwise PrePaid. 
        /// Default to PostPaid.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The internet ip of the VPN.
        /// </summary>
        [Output("internetIp")]
        public Output<string> InternetIp { get; private set; } = null!;

        /// <summary>
        /// The name of the VPN. Defaults to null.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The max connections of SSL VPN. Default to 5. The number of connections supported by each account is different. 
        /// This field is ignored when enable_ssl is false.
        /// </summary>
        [Output("sslConnections")]
        public Output<int?> SslConnections { get; private set; } = null!;

        /// <summary>
        /// The status of the VPN gateway.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The VPN belongs the vpc_id, the field can't be changed.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The VPN belongs the vswitch_id, the field can't be changed.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : Pulumi.ResourceArgs
    {
        [Input("bandwidth", required: true)]
        public Input<int> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The description of the VPN instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
        /// </summary>
        [Input("enableIpsec")]
        public Input<bool>? EnableIpsec { get; set; }

        /// <summary>
        /// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
        /// </summary>
        [Input("enableSsl")]
        public Input<bool>? EnableSsl { get; set; }

        /// <summary>
        /// The charge type for instance. If it is an international site account, the valid value is PostPaid, otherwise PrePaid. 
        /// Default to PostPaid.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The name of the VPN. Defaults to null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The max connections of SSL VPN. Default to 5. The number of connections supported by each account is different. 
        /// This field is ignored when enable_ssl is false.
        /// </summary>
        [Input("sslConnections")]
        public Input<int>? SslConnections { get; set; }

        /// <summary>
        /// The VPN belongs the vpc_id, the field can't be changed.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The VPN belongs the vswitch_id, the field can't be changed.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GatewayArgs()
        {
        }
    }

    public sealed class GatewayState : Pulumi.ResourceArgs
    {
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The business status of the VPN gateway.
        /// </summary>
        [Input("businessStatus")]
        public Input<string>? BusinessStatus { get; set; }

        /// <summary>
        /// The description of the VPN instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
        /// </summary>
        [Input("enableIpsec")]
        public Input<bool>? EnableIpsec { get; set; }

        /// <summary>
        /// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
        /// </summary>
        [Input("enableSsl")]
        public Input<bool>? EnableSsl { get; set; }

        /// <summary>
        /// The charge type for instance. If it is an international site account, the valid value is PostPaid, otherwise PrePaid. 
        /// Default to PostPaid.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The internet ip of the VPN.
        /// </summary>
        [Input("internetIp")]
        public Input<string>? InternetIp { get; set; }

        /// <summary>
        /// The name of the VPN. Defaults to null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The max connections of SSL VPN. Default to 5. The number of connections supported by each account is different. 
        /// This field is ignored when enable_ssl is false.
        /// </summary>
        [Input("sslConnections")]
        public Input<int>? SslConnections { get; set; }

        /// <summary>
        /// The status of the VPN gateway.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The VPN belongs the vpc_id, the field can't be changed.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The VPN belongs the vswitch_id, the field can't be changed.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GatewayState()
        {
        }
    }
}
