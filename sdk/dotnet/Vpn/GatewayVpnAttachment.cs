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
    /// Provides a VPN Gateway Vpn Attachment resource.
    /// 
    /// For information about VPN Gateway Vpn Attachment and how to use it, see [What is Vpn Attachment](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/createvpnattachment).
    /// 
    /// &gt; **NOTE:** Available since v1.181.0.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultCustomerGateway = new AliCloud.Vpn.CustomerGateway("defaultCustomerGateway", new()
    ///     {
    ///         IpAddress = "42.104.22.210",
    ///         Asn = "45014",
    ///         Description = name,
    ///     });
    /// 
    ///     var defaultGatewayVpnAttachment = new AliCloud.Vpn.GatewayVpnAttachment("defaultGatewayVpnAttachment", new()
    ///     {
    ///         CustomerGatewayId = defaultCustomerGateway.Id,
    ///         NetworkType = "public",
    ///         LocalSubnet = "0.0.0.0/0",
    ///         RemoteSubnet = "0.0.0.0/0",
    ///         EffectImmediately = false,
    ///         IkeConfig = new AliCloud.Vpn.Inputs.GatewayVpnAttachmentIkeConfigArgs
    ///         {
    ///             IkeAuthAlg = "md5",
    ///             IkeEncAlg = "des",
    ///             IkeVersion = "ikev2",
    ///             IkeMode = "main",
    ///             IkeLifetime = 86400,
    ///             Psk = "tf-testvpn2",
    ///             IkePfs = "group1",
    ///             RemoteId = "testbob2",
    ///             LocalId = "testalice2",
    ///         },
    ///         IpsecConfig = new AliCloud.Vpn.Inputs.GatewayVpnAttachmentIpsecConfigArgs
    ///         {
    ///             IpsecPfs = "group5",
    ///             IpsecEncAlg = "des",
    ///             IpsecAuthAlg = "md5",
    ///             IpsecLifetime = 86400,
    ///         },
    ///         BgpConfig = new AliCloud.Vpn.Inputs.GatewayVpnAttachmentBgpConfigArgs
    ///         {
    ///             Enable = true,
    ///             LocalAsn = 45014,
    ///             TunnelCidr = "169.254.11.0/30",
    ///             LocalBgpIp = "169.254.11.1",
    ///         },
    ///         HealthCheckConfig = new AliCloud.Vpn.Inputs.GatewayVpnAttachmentHealthCheckConfigArgs
    ///         {
    ///             Enable = true,
    ///             Sip = "192.168.1.1",
    ///             Dip = "10.0.0.1",
    ///             Interval = 10,
    ///             Retry = 10,
    ///             Policy = "revoke_route",
    ///         },
    ///         EnableDpd = true,
    ///         EnableNatTraversal = true,
    ///         VpnAttachmentName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPN Gateway Vpn Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment")]
    public partial class GatewayVpnAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bgp configuration information. See `bgp_config` below.
        /// </summary>
        [Output("bgpConfig")]
        public Output<Outputs.GatewayVpnAttachmentBgpConfig> BgpConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the customer gateway. From version 1.196.0, `customer_gateway_id` can be modified.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        /// </summary>
        [Output("effectImmediately")]
        public Output<bool?> EffectImmediately { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the DPD (peer survival detection) function.
        /// </summary>
        [Output("enableDpd")]
        public Output<bool> EnableDpd { get; private set; } = null!;

        /// <summary>
        /// Allow NAT penetration.
        /// </summary>
        [Output("enableNatTraversal")]
        public Output<bool> EnableNatTraversal { get; private set; } = null!;

        /// <summary>
        /// Health check configuration information. See `health_check_config` below.
        /// </summary>
        [Output("healthCheckConfig")]
        public Output<Outputs.GatewayVpnAttachmentHealthCheckConfig> HealthCheckConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration negotiated in the second stage. See `ike_config` below.
        /// </summary>
        [Output("ikeConfig")]
        public Output<Outputs.GatewayVpnAttachmentIkeConfig> IkeConfig { get; private set; } = null!;

        /// <summary>
        /// The VPN gateway IP.
        /// </summary>
        [Output("internetIp")]
        public Output<string> InternetIp { get; private set; } = null!;

        /// <summary>
        /// Configuration negotiated in the second stage. See `ipsec_config` below.
        /// </summary>
        [Output("ipsecConfig")]
        public Output<Outputs.GatewayVpnAttachmentIpsecConfig> IpsecConfig { get; private set; } = null!;

        /// <summary>
        /// The CIDR block of the virtual private cloud (VPC).
        /// </summary>
        [Output("localSubnet")]
        public Output<string> LocalSubnet { get; private set; } = null!;

        /// <summary>
        /// The network type of the IPsec connection. Valid values: `public`, `private`.
        /// </summary>
        [Output("networkType")]
        public Output<string> NetworkType { get; private set; } = null!;

        /// <summary>
        /// The CIDR block of the on-premises data center.
        /// </summary>
        [Output("remoteSubnet")]
        public Output<string> RemoteSubnet { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of the vpn attachment.
        /// </summary>
        [Output("vpnAttachmentName")]
        public Output<string?> VpnAttachmentName { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayVpnAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayVpnAttachment(string name, GatewayVpnAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment", name, args ?? new GatewayVpnAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayVpnAttachment(string name, Input<string> id, GatewayVpnAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gatewayVpnAttachment:GatewayVpnAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayVpnAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayVpnAttachment Get(string name, Input<string> id, GatewayVpnAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayVpnAttachment(name, id, state, options);
        }
    }

    public sealed class GatewayVpnAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bgp configuration information. See `bgp_config` below.
        /// </summary>
        [Input("bgpConfig")]
        public Input<Inputs.GatewayVpnAttachmentBgpConfigArgs>? BgpConfig { get; set; }

        /// <summary>
        /// The ID of the customer gateway. From version 1.196.0, `customer_gateway_id` can be modified.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        /// <summary>
        /// Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        /// </summary>
        [Input("effectImmediately")]
        public Input<bool>? EffectImmediately { get; set; }

        /// <summary>
        /// Whether to enable the DPD (peer survival detection) function.
        /// </summary>
        [Input("enableDpd")]
        public Input<bool>? EnableDpd { get; set; }

        /// <summary>
        /// Allow NAT penetration.
        /// </summary>
        [Input("enableNatTraversal")]
        public Input<bool>? EnableNatTraversal { get; set; }

        /// <summary>
        /// Health check configuration information. See `health_check_config` below.
        /// </summary>
        [Input("healthCheckConfig")]
        public Input<Inputs.GatewayVpnAttachmentHealthCheckConfigArgs>? HealthCheckConfig { get; set; }

        /// <summary>
        /// Configuration negotiated in the second stage. See `ike_config` below.
        /// </summary>
        [Input("ikeConfig")]
        public Input<Inputs.GatewayVpnAttachmentIkeConfigArgs>? IkeConfig { get; set; }

        /// <summary>
        /// Configuration negotiated in the second stage. See `ipsec_config` below.
        /// </summary>
        [Input("ipsecConfig")]
        public Input<Inputs.GatewayVpnAttachmentIpsecConfigArgs>? IpsecConfig { get; set; }

        /// <summary>
        /// The CIDR block of the virtual private cloud (VPC).
        /// </summary>
        [Input("localSubnet", required: true)]
        public Input<string> LocalSubnet { get; set; } = null!;

        /// <summary>
        /// The network type of the IPsec connection. Valid values: `public`, `private`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The CIDR block of the on-premises data center.
        /// </summary>
        [Input("remoteSubnet", required: true)]
        public Input<string> RemoteSubnet { get; set; } = null!;

        /// <summary>
        /// The name of the vpn attachment.
        /// </summary>
        [Input("vpnAttachmentName")]
        public Input<string>? VpnAttachmentName { get; set; }

        public GatewayVpnAttachmentArgs()
        {
        }
        public static new GatewayVpnAttachmentArgs Empty => new GatewayVpnAttachmentArgs();
    }

    public sealed class GatewayVpnAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bgp configuration information. See `bgp_config` below.
        /// </summary>
        [Input("bgpConfig")]
        public Input<Inputs.GatewayVpnAttachmentBgpConfigGetArgs>? BgpConfig { get; set; }

        /// <summary>
        /// The ID of the customer gateway. From version 1.196.0, `customer_gateway_id` can be modified.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        /// <summary>
        /// Indicates whether IPsec-VPN negotiations are initiated immediately. Valid values.
        /// </summary>
        [Input("effectImmediately")]
        public Input<bool>? EffectImmediately { get; set; }

        /// <summary>
        /// Whether to enable the DPD (peer survival detection) function.
        /// </summary>
        [Input("enableDpd")]
        public Input<bool>? EnableDpd { get; set; }

        /// <summary>
        /// Allow NAT penetration.
        /// </summary>
        [Input("enableNatTraversal")]
        public Input<bool>? EnableNatTraversal { get; set; }

        /// <summary>
        /// Health check configuration information. See `health_check_config` below.
        /// </summary>
        [Input("healthCheckConfig")]
        public Input<Inputs.GatewayVpnAttachmentHealthCheckConfigGetArgs>? HealthCheckConfig { get; set; }

        /// <summary>
        /// Configuration negotiated in the second stage. See `ike_config` below.
        /// </summary>
        [Input("ikeConfig")]
        public Input<Inputs.GatewayVpnAttachmentIkeConfigGetArgs>? IkeConfig { get; set; }

        /// <summary>
        /// The VPN gateway IP.
        /// </summary>
        [Input("internetIp")]
        public Input<string>? InternetIp { get; set; }

        /// <summary>
        /// Configuration negotiated in the second stage. See `ipsec_config` below.
        /// </summary>
        [Input("ipsecConfig")]
        public Input<Inputs.GatewayVpnAttachmentIpsecConfigGetArgs>? IpsecConfig { get; set; }

        /// <summary>
        /// The CIDR block of the virtual private cloud (VPC).
        /// </summary>
        [Input("localSubnet")]
        public Input<string>? LocalSubnet { get; set; }

        /// <summary>
        /// The network type of the IPsec connection. Valid values: `public`, `private`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The CIDR block of the on-premises data center.
        /// </summary>
        [Input("remoteSubnet")]
        public Input<string>? RemoteSubnet { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the vpn attachment.
        /// </summary>
        [Input("vpnAttachmentName")]
        public Input<string>? VpnAttachmentName { get; set; }

        public GatewayVpnAttachmentState()
        {
        }
        public static new GatewayVpnAttachmentState Empty => new GatewayVpnAttachmentState();
    }
}
