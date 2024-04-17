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
    /// Provides a VPN Gateway Vco Route resource.
    /// 
    /// For information about VPN Gateway Vco Route and how to use it, see [What is Vco Route](https://www.alibabacloud.com/help/zh/virtual-private-cloud/latest/createvcorouteentry).
    /// 
    /// &gt; **NOTE:** Available in v1.183.0+.
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
    ///     var defaultInstance = new AliCloud.Cen.Instance("default", new()
    ///     {
    ///         CenInstanceName = name,
    ///     });
    /// 
    ///     var defaultTransitRouter = new AliCloud.Cen.TransitRouter("default", new()
    ///     {
    ///         CenId = defaultInstance.Id,
    ///         TransitRouterDescription = "desd",
    ///         TransitRouterName = name,
    ///     });
    /// 
    ///     var @default = AliCloud.Cen.GetTransitRouterAvailableResources.Invoke();
    /// 
    ///     var defaultCustomerGateway = new AliCloud.Vpn.CustomerGateway("default", new()
    ///     {
    ///         Name = name,
    ///         IpAddress = "42.104.22.210",
    ///         Asn = "45014",
    ///         Description = "testAccVpnConnectionDesc",
    ///     });
    /// 
    ///     var defaultGatewayVpnAttachment = new AliCloud.Vpn.GatewayVpnAttachment("default", new()
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
    ///     var defaultTransitRouterVpnAttachment = new AliCloud.Cen.TransitRouterVpnAttachment("default", new()
    ///     {
    ///         AutoPublishRouteEnabled = false,
    ///         TransitRouterAttachmentDescription = name,
    ///         TransitRouterAttachmentName = name,
    ///         CenId = defaultTransitRouter.CenId,
    ///         TransitRouterId = defaultTransitRouter.TransitRouterId,
    ///         VpnId = defaultGatewayVpnAttachment.Id,
    ///         Zones = new[]
    ///         {
    ///             new AliCloud.Cen.Inputs.TransitRouterVpnAttachmentZoneArgs
    ///             {
    ///                 ZoneId = @default.Apply(@default =&gt; @default.Apply(getTransitRouterAvailableResourcesResult =&gt; getTransitRouterAvailableResourcesResult.Resources[0]?.MasterZones[0])),
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultGatewayVcoRoute = new AliCloud.Vpn.GatewayVcoRoute("default", new()
    ///     {
    ///         RouteDest = "192.168.12.0/24",
    ///         NextHop = defaultTransitRouterVpnAttachment.VpnId,
    ///         VpnConnectionId = defaultTransitRouterVpnAttachment.VpnId,
    ///         Weight = 100,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPN Gateway Vco Route can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute example &lt;vpn_connection_id&gt;:&lt;route_dest&gt;:&lt;next_hop&gt;:&lt;weight&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute")]
    public partial class GatewayVcoRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The next hop of the destination route.
        /// </summary>
        [Output("nextHop")]
        public Output<string> NextHop { get; private set; } = null!;

        /// <summary>
        /// The destination network segment of the destination route.
        /// </summary>
        [Output("routeDest")]
        public Output<string> RouteDest { get; private set; } = null!;

        /// <summary>
        /// The status of the vpn route entry.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of the vpn attachment.
        /// </summary>
        [Output("vpnConnectionId")]
        public Output<string> VpnConnectionId { get; private set; } = null!;

        /// <summary>
        /// The weight value of the destination route. Valid values: `0`, `100`.
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayVcoRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayVcoRoute(string name, GatewayVcoRouteArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute", name, args ?? new GatewayVcoRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayVcoRoute(string name, Input<string> id, GatewayVcoRouteState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpn/gatewayVcoRoute:GatewayVcoRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayVcoRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayVcoRoute Get(string name, Input<string> id, GatewayVcoRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayVcoRoute(name, id, state, options);
        }
    }

    public sealed class GatewayVcoRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The next hop of the destination route.
        /// </summary>
        [Input("nextHop", required: true)]
        public Input<string> NextHop { get; set; } = null!;

        /// <summary>
        /// The destination network segment of the destination route.
        /// </summary>
        [Input("routeDest", required: true)]
        public Input<string> RouteDest { get; set; } = null!;

        /// <summary>
        /// The id of the vpn attachment.
        /// </summary>
        [Input("vpnConnectionId", required: true)]
        public Input<string> VpnConnectionId { get; set; } = null!;

        /// <summary>
        /// The weight value of the destination route. Valid values: `0`, `100`.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public GatewayVcoRouteArgs()
        {
        }
        public static new GatewayVcoRouteArgs Empty => new GatewayVcoRouteArgs();
    }

    public sealed class GatewayVcoRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The next hop of the destination route.
        /// </summary>
        [Input("nextHop")]
        public Input<string>? NextHop { get; set; }

        /// <summary>
        /// The destination network segment of the destination route.
        /// </summary>
        [Input("routeDest")]
        public Input<string>? RouteDest { get; set; }

        /// <summary>
        /// The status of the vpn route entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the vpn attachment.
        /// </summary>
        [Input("vpnConnectionId")]
        public Input<string>? VpnConnectionId { get; set; }

        /// <summary>
        /// The weight value of the destination route. Valid values: `0`, `100`.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public GatewayVcoRouteState()
        {
        }
        public static new GatewayVcoRouteState Empty => new GatewayVcoRouteState();
    }
}
