// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Vpn Gateway Vco Routes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.183.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.cen.Instance("default", {cenInstanceName: name});
 * const defaultTransitRouter = new alicloud.cen.TransitRouter("default", {
 *     cenId: defaultInstance.id,
 *     transitRouterDescription: "desd",
 *     transitRouterName: name,
 * });
 * const default = alicloud.cen.getTransitRouterAvailableResources({});
 * const defaultCustomerGateway = new alicloud.vpn.CustomerGateway("default", {
 *     name: name,
 *     ipAddress: "42.104.22.210",
 *     asn: "45014",
 *     description: "testAccVpnConnectionDesc",
 * });
 * const defaultGatewayVpnAttachment = new alicloud.vpn.GatewayVpnAttachment("default", {
 *     customerGatewayId: defaultCustomerGateway.id,
 *     networkType: "public",
 *     localSubnet: "0.0.0.0/0",
 *     remoteSubnet: "0.0.0.0/0",
 *     effectImmediately: false,
 *     ikeConfig: {
 *         ikeAuthAlg: "md5",
 *         ikeEncAlg: "des",
 *         ikeVersion: "ikev2",
 *         ikeMode: "main",
 *         ikeLifetime: 86400,
 *         psk: "tf-testvpn2",
 *         ikePfs: "group1",
 *         remoteId: "testbob2",
 *         localId: "testalice2",
 *     },
 *     ipsecConfig: {
 *         ipsecPfs: "group5",
 *         ipsecEncAlg: "des",
 *         ipsecAuthAlg: "md5",
 *         ipsecLifetime: 86400,
 *     },
 *     bgpConfig: {
 *         enable: true,
 *         localAsn: 45014,
 *         tunnelCidr: "169.254.11.0/30",
 *         localBgpIp: "169.254.11.1",
 *     },
 *     healthCheckConfig: {
 *         enable: true,
 *         sip: "192.168.1.1",
 *         dip: "10.0.0.1",
 *         interval: 10,
 *         retry: 10,
 *         policy: "revoke_route",
 *     },
 *     enableDpd: true,
 *     enableNatTraversal: true,
 *     vpnAttachmentName: name,
 * });
 * const defaultTransitRouterVpnAttachment = new alicloud.cen.TransitRouterVpnAttachment("default", {
 *     autoPublishRouteEnabled: false,
 *     transitRouterAttachmentDescription: name,
 *     transitRouterAttachmentName: name,
 *     cenId: defaultTransitRouter.cenId,
 *     transitRouterId: defaultTransitRouter.transitRouterId,
 *     vpnId: defaultGatewayVpnAttachment.id,
 *     zones: [{
 *         zoneId: _default.then(_default => _default.resources?.[0]?.masterZones?.[0]),
 *     }],
 * });
 * const defaultGatewayVcoRoute = new alicloud.vpn.GatewayVcoRoute("default", {
 *     routeDest: "192.168.12.0/24",
 *     nextHop: defaultTransitRouterVpnAttachment.vpnId,
 *     vpnConnectionId: defaultTransitRouterVpnAttachment.vpnId,
 *     weight: 100,
 * });
 * const defaultGetGatewayVcoRoutes = alicloud.vpn.getGatewayVcoRoutesOutput({
 *     vpnConnectionId: defaultTransitRouterVpnAttachment.vpnId,
 * });
 * export const vpnGatewayVcoRouteId1 = ids.routes[0].id;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGatewayVcoRoutes(args: GetGatewayVcoRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayVcoRoutesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:vpn/getGatewayVcoRoutes:getGatewayVcoRoutes", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "routeEntryType": args.routeEntryType,
        "status": args.status,
        "vpnConnectionId": args.vpnConnectionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGatewayVcoRoutes.
 */
export interface GetGatewayVcoRoutesArgs {
    /**
     * A list of Vco Route IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The Routing input type. Valid values: `custom`, `bgp`.
     */
    routeEntryType?: string;
    /**
     * The status of the vpn route entry.
     */
    status?: string;
    /**
     * The id of the vpn connection.
     */
    vpnConnectionId: string;
}

/**
 * A collection of values returned by getGatewayVcoRoutes.
 */
export interface GetGatewayVcoRoutesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly routeEntryType?: string;
    readonly routes: outputs.vpn.GetGatewayVcoRoutesRoute[];
    readonly status?: string;
    readonly vpnConnectionId: string;
}
/**
 * This data source provides the Vpn Gateway Vco Routes of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.183.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.cen.Instance("default", {cenInstanceName: name});
 * const defaultTransitRouter = new alicloud.cen.TransitRouter("default", {
 *     cenId: defaultInstance.id,
 *     transitRouterDescription: "desd",
 *     transitRouterName: name,
 * });
 * const default = alicloud.cen.getTransitRouterAvailableResources({});
 * const defaultCustomerGateway = new alicloud.vpn.CustomerGateway("default", {
 *     name: name,
 *     ipAddress: "42.104.22.210",
 *     asn: "45014",
 *     description: "testAccVpnConnectionDesc",
 * });
 * const defaultGatewayVpnAttachment = new alicloud.vpn.GatewayVpnAttachment("default", {
 *     customerGatewayId: defaultCustomerGateway.id,
 *     networkType: "public",
 *     localSubnet: "0.0.0.0/0",
 *     remoteSubnet: "0.0.0.0/0",
 *     effectImmediately: false,
 *     ikeConfig: {
 *         ikeAuthAlg: "md5",
 *         ikeEncAlg: "des",
 *         ikeVersion: "ikev2",
 *         ikeMode: "main",
 *         ikeLifetime: 86400,
 *         psk: "tf-testvpn2",
 *         ikePfs: "group1",
 *         remoteId: "testbob2",
 *         localId: "testalice2",
 *     },
 *     ipsecConfig: {
 *         ipsecPfs: "group5",
 *         ipsecEncAlg: "des",
 *         ipsecAuthAlg: "md5",
 *         ipsecLifetime: 86400,
 *     },
 *     bgpConfig: {
 *         enable: true,
 *         localAsn: 45014,
 *         tunnelCidr: "169.254.11.0/30",
 *         localBgpIp: "169.254.11.1",
 *     },
 *     healthCheckConfig: {
 *         enable: true,
 *         sip: "192.168.1.1",
 *         dip: "10.0.0.1",
 *         interval: 10,
 *         retry: 10,
 *         policy: "revoke_route",
 *     },
 *     enableDpd: true,
 *     enableNatTraversal: true,
 *     vpnAttachmentName: name,
 * });
 * const defaultTransitRouterVpnAttachment = new alicloud.cen.TransitRouterVpnAttachment("default", {
 *     autoPublishRouteEnabled: false,
 *     transitRouterAttachmentDescription: name,
 *     transitRouterAttachmentName: name,
 *     cenId: defaultTransitRouter.cenId,
 *     transitRouterId: defaultTransitRouter.transitRouterId,
 *     vpnId: defaultGatewayVpnAttachment.id,
 *     zones: [{
 *         zoneId: _default.then(_default => _default.resources?.[0]?.masterZones?.[0]),
 *     }],
 * });
 * const defaultGatewayVcoRoute = new alicloud.vpn.GatewayVcoRoute("default", {
 *     routeDest: "192.168.12.0/24",
 *     nextHop: defaultTransitRouterVpnAttachment.vpnId,
 *     vpnConnectionId: defaultTransitRouterVpnAttachment.vpnId,
 *     weight: 100,
 * });
 * const defaultGetGatewayVcoRoutes = alicloud.vpn.getGatewayVcoRoutesOutput({
 *     vpnConnectionId: defaultTransitRouterVpnAttachment.vpnId,
 * });
 * export const vpnGatewayVcoRouteId1 = ids.routes[0].id;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGatewayVcoRoutesOutput(args: GetGatewayVcoRoutesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayVcoRoutesResult> {
    return pulumi.output(args).apply((a: any) => getGatewayVcoRoutes(a, opts))
}

/**
 * A collection of arguments for invoking getGatewayVcoRoutes.
 */
export interface GetGatewayVcoRoutesOutputArgs {
    /**
     * A list of Vco Route IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The Routing input type. Valid values: `custom`, `bgp`.
     */
    routeEntryType?: pulumi.Input<string>;
    /**
     * The status of the vpn route entry.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the vpn connection.
     */
    vpnConnectionId: pulumi.Input<string>;
}
