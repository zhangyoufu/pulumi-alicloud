// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn
{
    public static class GetGatewayVcoRoutes
    {
        /// <summary>
        /// This data source provides the Vpn Gateway Vco Routes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.183.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new()
        ///     {
        ///         CenInstanceName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouter = new AliCloud.Cen.TransitRouter("defaultTransitRouter", new()
        ///     {
        ///         CenId = defaultInstance.Id,
        ///         TransitRouterDescription = "desd",
        ///         TransitRouterName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouterAvailableResources = AliCloud.Cen.GetTransitRouterAvailableResources.Invoke();
        /// 
        ///     var defaultCustomerGateway = new AliCloud.Vpn.CustomerGateway("defaultCustomerGateway", new()
        ///     {
        ///         IpAddress = "42.104.22.210",
        ///         Asn = "45014",
        ///         Description = "testAccVpnConnectionDesc",
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
        ///         VpnAttachmentName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouterVpnAttachment = new AliCloud.Cen.TransitRouterVpnAttachment("defaultTransitRouterVpnAttachment", new()
        ///     {
        ///         AutoPublishRouteEnabled = false,
        ///         TransitRouterAttachmentDescription = @var.Name,
        ///         TransitRouterAttachmentName = @var.Name,
        ///         CenId = defaultTransitRouter.CenId,
        ///         TransitRouterId = defaultTransitRouter.TransitRouterId,
        ///         VpnId = defaultGatewayVpnAttachment.Id,
        ///         Zones = new[]
        ///         {
        ///             new AliCloud.Cen.Inputs.TransitRouterVpnAttachmentZoneArgs
        ///             {
        ///                 ZoneId = defaultTransitRouterAvailableResources.Apply(getTransitRouterAvailableResourcesResult =&gt; getTransitRouterAvailableResourcesResult.Resources[0]?.MasterZones[0]),
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultGatewayVcoRoute = new AliCloud.Vpn.GatewayVcoRoute("defaultGatewayVcoRoute", new()
        ///     {
        ///         RouteDest = "192.168.12.0/24",
        ///         NextHop = defaultTransitRouterVpnAttachment.VpnId,
        ///         VpnConnectionId = defaultTransitRouterVpnAttachment.VpnId,
        ///         Weight = 100,
        ///     });
        /// 
        ///     var defaultGatewayVcoRoutes = AliCloud.Vpn.GetGatewayVcoRoutes.Invoke(new()
        ///     {
        ///         VpnConnectionId = defaultTransitRouterVpnAttachment.VpnId,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpnGatewayVcoRouteId1"] = data.Alicloud_vpn_gateway_vco_routes.Ids.Routes[0].Id,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewayVcoRoutesResult> InvokeAsync(GetGatewayVcoRoutesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayVcoRoutesResult>("alicloud:vpn/getGatewayVcoRoutes:getGatewayVcoRoutes", args ?? new GetGatewayVcoRoutesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Vpn Gateway Vco Routes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.183.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new()
        ///     {
        ///         CenInstanceName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouter = new AliCloud.Cen.TransitRouter("defaultTransitRouter", new()
        ///     {
        ///         CenId = defaultInstance.Id,
        ///         TransitRouterDescription = "desd",
        ///         TransitRouterName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouterAvailableResources = AliCloud.Cen.GetTransitRouterAvailableResources.Invoke();
        /// 
        ///     var defaultCustomerGateway = new AliCloud.Vpn.CustomerGateway("defaultCustomerGateway", new()
        ///     {
        ///         IpAddress = "42.104.22.210",
        ///         Asn = "45014",
        ///         Description = "testAccVpnConnectionDesc",
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
        ///         VpnAttachmentName = @var.Name,
        ///     });
        /// 
        ///     var defaultTransitRouterVpnAttachment = new AliCloud.Cen.TransitRouterVpnAttachment("defaultTransitRouterVpnAttachment", new()
        ///     {
        ///         AutoPublishRouteEnabled = false,
        ///         TransitRouterAttachmentDescription = @var.Name,
        ///         TransitRouterAttachmentName = @var.Name,
        ///         CenId = defaultTransitRouter.CenId,
        ///         TransitRouterId = defaultTransitRouter.TransitRouterId,
        ///         VpnId = defaultGatewayVpnAttachment.Id,
        ///         Zones = new[]
        ///         {
        ///             new AliCloud.Cen.Inputs.TransitRouterVpnAttachmentZoneArgs
        ///             {
        ///                 ZoneId = defaultTransitRouterAvailableResources.Apply(getTransitRouterAvailableResourcesResult =&gt; getTransitRouterAvailableResourcesResult.Resources[0]?.MasterZones[0]),
        ///             },
        ///         },
        ///     });
        /// 
        ///     var defaultGatewayVcoRoute = new AliCloud.Vpn.GatewayVcoRoute("defaultGatewayVcoRoute", new()
        ///     {
        ///         RouteDest = "192.168.12.0/24",
        ///         NextHop = defaultTransitRouterVpnAttachment.VpnId,
        ///         VpnConnectionId = defaultTransitRouterVpnAttachment.VpnId,
        ///         Weight = 100,
        ///     });
        /// 
        ///     var defaultGatewayVcoRoutes = AliCloud.Vpn.GetGatewayVcoRoutes.Invoke(new()
        ///     {
        ///         VpnConnectionId = defaultTransitRouterVpnAttachment.VpnId,
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["vpnGatewayVcoRouteId1"] = data.Alicloud_vpn_gateway_vco_routes.Ids.Routes[0].Id,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGatewayVcoRoutesResult> Invoke(GetGatewayVcoRoutesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayVcoRoutesResult>("alicloud:vpn/getGatewayVcoRoutes:getGatewayVcoRoutes", args ?? new GetGatewayVcoRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayVcoRoutesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Vco Route IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The Routing input type. Valid values: `custom`, `bgp`.
        /// </summary>
        [Input("routeEntryType")]
        public string? RouteEntryType { get; set; }

        /// <summary>
        /// The status of the vpn route entry.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The id of the vpn connection.
        /// </summary>
        [Input("vpnConnectionId", required: true)]
        public string VpnConnectionId { get; set; } = null!;

        public GetGatewayVcoRoutesArgs()
        {
        }
        public static new GetGatewayVcoRoutesArgs Empty => new GetGatewayVcoRoutesArgs();
    }

    public sealed class GetGatewayVcoRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Vco Route IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The Routing input type. Valid values: `custom`, `bgp`.
        /// </summary>
        [Input("routeEntryType")]
        public Input<string>? RouteEntryType { get; set; }

        /// <summary>
        /// The status of the vpn route entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the vpn connection.
        /// </summary>
        [Input("vpnConnectionId", required: true)]
        public Input<string> VpnConnectionId { get; set; } = null!;

        public GetGatewayVcoRoutesInvokeArgs()
        {
        }
        public static new GetGatewayVcoRoutesInvokeArgs Empty => new GetGatewayVcoRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayVcoRoutesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? RouteEntryType;
        public readonly ImmutableArray<Outputs.GetGatewayVcoRoutesRouteResult> Routes;
        public readonly string? Status;
        public readonly string VpnConnectionId;

        [OutputConstructor]
        private GetGatewayVcoRoutesResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            int? pageNumber,

            int? pageSize,

            string? routeEntryType,

            ImmutableArray<Outputs.GetGatewayVcoRoutesRouteResult> routes,

            string? status,

            string vpnConnectionId)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            RouteEntryType = routeEntryType;
            Routes = routes;
            Status = status;
            VpnConnectionId = vpnConnectionId;
        }
    }
}
