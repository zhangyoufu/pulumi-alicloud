// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Enterprise Network (CEN) Transit Router Vpn Attachment resource.
 *
 * For information about Cloud Enterprise Network (CEN) Transit Router Vpn Attachment and how to use it, see [What is Transit Router Vpn Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutervpnattachment).
 *
 * > **NOTE:** Available since v1.183.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf_example";
 * const default = alicloud.cen.getTransitRouterAvailableResources({});
 * const exampleInstance = new alicloud.cen.Instance("exampleInstance", {cenInstanceName: name});
 * const exampleTransitRouter = new alicloud.cen.TransitRouter("exampleTransitRouter", {
 *     cenId: exampleInstance.id,
 *     transitRouterDescription: name,
 *     transitRouterName: name,
 * });
 * const exampleCustomerGateway = new alicloud.vpn.CustomerGateway("exampleCustomerGateway", {
 *     ipAddress: "42.104.22.210",
 *     asn: "45014",
 *     description: name,
 * });
 * const exampleGatewayVpnAttachment = new alicloud.vpn.GatewayVpnAttachment("exampleGatewayVpnAttachment", {
 *     customerGatewayId: exampleCustomerGateway.id,
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
 * const exampleTransitRouterCidr = new alicloud.cen.TransitRouterCidr("exampleTransitRouterCidr", {
 *     transitRouterId: exampleTransitRouter.transitRouterId,
 *     cidr: "192.168.0.0/16",
 *     transitRouterCidrName: name,
 *     description: name,
 *     publishCidrRoute: true,
 * });
 * const exampleTransitRouterVpnAttachment = new alicloud.cen.TransitRouterVpnAttachment("exampleTransitRouterVpnAttachment", {
 *     autoPublishRouteEnabled: false,
 *     transitRouterAttachmentDescription: name,
 *     transitRouterAttachmentName: name,
 *     cenId: exampleTransitRouter.cenId,
 *     transitRouterId: exampleTransitRouterCidr.transitRouterId,
 *     vpnId: exampleGatewayVpnAttachment.id,
 *     zones: [{
 *         zoneId: _default.then(_default => _default.resources?.[0]?.masterZones?.[0]),
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Enterprise Network (CEN) Transit Router Vpn Attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cen/transitRouterVpnAttachment:TransitRouterVpnAttachment example <id>
 * ```
 */
export class TransitRouterVpnAttachment extends pulumi.CustomResource {
    /**
     * Get an existing TransitRouterVpnAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitRouterVpnAttachmentState, opts?: pulumi.CustomResourceOptions): TransitRouterVpnAttachment {
        return new TransitRouterVpnAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/transitRouterVpnAttachment:TransitRouterVpnAttachment';

    /**
     * Returns true if the given object is an instance of TransitRouterVpnAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitRouterVpnAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitRouterVpnAttachment.__pulumiType;
    }

    /**
     * Whether to allow the forwarding router instance to automatically publish routing entries to IPsec connections.
     */
    public readonly autoPublishRouteEnabled!: pulumi.Output<boolean>;
    /**
     * The id of the cen.
     */
    public readonly cenId!: pulumi.Output<string | undefined>;
    /**
     * The associating status of the network.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The description of the VPN connection. The description can contain `2` to `256` characters. The description must start with English letters, but cannot start with `http://` or `https://`.
     */
    public readonly transitRouterAttachmentDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the VPN connection. The name must be `2` to `128` characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    public readonly transitRouterAttachmentName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the forwarding router instance.
     */
    public readonly transitRouterId!: pulumi.Output<string>;
    /**
     * The id of the vpn.
     */
    public readonly vpnId!: pulumi.Output<string>;
    /**
     * The owner id of vpn. **NOTE:** You must set `vpnOwnerId`, if you want to connect the transit router to an IPsec-VPN connection that belongs to another Alibaba Cloud account.
     */
    public readonly vpnOwnerId!: pulumi.Output<string>;
    /**
     * The list of zone mapping. See `zone` below.
     */
    public readonly zones!: pulumi.Output<outputs.cen.TransitRouterVpnAttachmentZone[]>;

    /**
     * Create a TransitRouterVpnAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitRouterVpnAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitRouterVpnAttachmentArgs | TransitRouterVpnAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitRouterVpnAttachmentState | undefined;
            resourceInputs["autoPublishRouteEnabled"] = state ? state.autoPublishRouteEnabled : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["transitRouterAttachmentDescription"] = state ? state.transitRouterAttachmentDescription : undefined;
            resourceInputs["transitRouterAttachmentName"] = state ? state.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["vpnId"] = state ? state.vpnId : undefined;
            resourceInputs["vpnOwnerId"] = state ? state.vpnOwnerId : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as TransitRouterVpnAttachmentArgs | undefined;
            if ((!args || args.transitRouterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitRouterId'");
            }
            if ((!args || args.vpnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnId'");
            }
            if ((!args || args.zones === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zones'");
            }
            resourceInputs["autoPublishRouteEnabled"] = args ? args.autoPublishRouteEnabled : undefined;
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitRouterAttachmentDescription"] = args ? args.transitRouterAttachmentDescription : undefined;
            resourceInputs["transitRouterAttachmentName"] = args ? args.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = args ? args.transitRouterId : undefined;
            resourceInputs["vpnId"] = args ? args.vpnId : undefined;
            resourceInputs["vpnOwnerId"] = args ? args.vpnOwnerId : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitRouterVpnAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitRouterVpnAttachment resources.
 */
export interface TransitRouterVpnAttachmentState {
    /**
     * Whether to allow the forwarding router instance to automatically publish routing entries to IPsec connections.
     */
    autoPublishRouteEnabled?: pulumi.Input<boolean>;
    /**
     * The id of the cen.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The associating status of the network.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The description of the VPN connection. The description can contain `2` to `256` characters. The description must start with English letters, but cannot start with `http://` or `https://`.
     */
    transitRouterAttachmentDescription?: pulumi.Input<string>;
    /**
     * The name of the VPN connection. The name must be `2` to `128` characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The ID of the forwarding router instance.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The id of the vpn.
     */
    vpnId?: pulumi.Input<string>;
    /**
     * The owner id of vpn. **NOTE:** You must set `vpnOwnerId`, if you want to connect the transit router to an IPsec-VPN connection that belongs to another Alibaba Cloud account.
     */
    vpnOwnerId?: pulumi.Input<string>;
    /**
     * The list of zone mapping. See `zone` below.
     */
    zones?: pulumi.Input<pulumi.Input<inputs.cen.TransitRouterVpnAttachmentZone>[]>;
}

/**
 * The set of arguments for constructing a TransitRouterVpnAttachment resource.
 */
export interface TransitRouterVpnAttachmentArgs {
    /**
     * Whether to allow the forwarding router instance to automatically publish routing entries to IPsec connections.
     */
    autoPublishRouteEnabled?: pulumi.Input<boolean>;
    /**
     * The id of the cen.
     */
    cenId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The description of the VPN connection. The description can contain `2` to `256` characters. The description must start with English letters, but cannot start with `http://` or `https://`.
     */
    transitRouterAttachmentDescription?: pulumi.Input<string>;
    /**
     * The name of the VPN connection. The name must be `2` to `128` characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The ID of the forwarding router instance.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The id of the vpn.
     */
    vpnId: pulumi.Input<string>;
    /**
     * The owner id of vpn. **NOTE:** You must set `vpnOwnerId`, if you want to connect the transit router to an IPsec-VPN connection that belongs to another Alibaba Cloud account.
     */
    vpnOwnerId?: pulumi.Input<string>;
    /**
     * The list of zone mapping. See `zone` below.
     */
    zones: pulumi.Input<pulumi.Input<inputs.cen.TransitRouterVpnAttachmentZone>[]>;
}
