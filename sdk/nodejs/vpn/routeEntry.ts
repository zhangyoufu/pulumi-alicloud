// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
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
 * const name = config.get("name") || "tf-example";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "^default-NODELETING$",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?.[0],
 *     zoneId: defaultZones.ids?.[0],
 * }));
 * const defaultGateway = new alicloud.vpn.Gateway("defaultGateway", {
 *     vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?.[0]),
 *     bandwidth: 10,
 *     instanceChargeType: "PrePaid",
 *     enableSsl: false,
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?.[0]),
 * });
 * const defaultCustomerGateway = new alicloud.vpn.CustomerGateway("defaultCustomerGateway", {ipAddress: "192.168.1.1"});
 * const defaultConnection = new alicloud.vpn.Connection("defaultConnection", {
 *     customerGatewayId: defaultCustomerGateway.id,
 *     vpnGatewayId: defaultGateway.id,
 *     localSubnets: ["192.168.2.0/24"],
 *     remoteSubnets: ["192.168.3.0/24"],
 * });
 * const defaultRouteEntry = new alicloud.vpn.RouteEntry("defaultRouteEntry", {
 *     vpnGatewayId: defaultGateway.id,
 *     routeDest: "10.0.0.0/24",
 *     nextHop: defaultConnection.id,
 *     weight: 0,
 *     publishVpc: false,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpn/routeEntry:RouteEntry example vpn-abc123456:vco-abc123456:10.0.0.10/24
 * ```
 */
export class RouteEntry extends pulumi.CustomResource {
    /**
     * Get an existing RouteEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteEntryState, opts?: pulumi.CustomResourceOptions): RouteEntry {
        return new RouteEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/routeEntry:RouteEntry';

    /**
     * Returns true if the given object is an instance of RouteEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteEntry.__pulumiType;
    }

    /**
     * The next hop of the destination route.
     */
    public readonly nextHop!: pulumi.Output<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    public readonly publishVpc!: pulumi.Output<boolean>;
    /**
     * The destination network segment of the destination route.
     */
    public readonly routeDest!: pulumi.Output<string>;
    /**
     * (Available in 1.161.0+) The type of the vpn route entry.
     */
    public /*out*/ readonly routeEntryType!: pulumi.Output<string>;
    /**
     * (Available in 1.161.0+) The status of the vpn route entry.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of the vpn gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string>;
    /**
     * The value should be 0 or 100.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a RouteEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteEntryArgs | RouteEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteEntryState | undefined;
            resourceInputs["nextHop"] = state ? state.nextHop : undefined;
            resourceInputs["publishVpc"] = state ? state.publishVpc : undefined;
            resourceInputs["routeDest"] = state ? state.routeDest : undefined;
            resourceInputs["routeEntryType"] = state ? state.routeEntryType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as RouteEntryArgs | undefined;
            if ((!args || args.nextHop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHop'");
            }
            if ((!args || args.publishVpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publishVpc'");
            }
            if ((!args || args.routeDest === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeDest'");
            }
            if ((!args || args.vpnGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnGatewayId'");
            }
            if ((!args || args.weight === undefined) && !opts.urn) {
                throw new Error("Missing required property 'weight'");
            }
            resourceInputs["nextHop"] = args ? args.nextHop : undefined;
            resourceInputs["publishVpc"] = args ? args.publishVpc : undefined;
            resourceInputs["routeDest"] = args ? args.routeDest : undefined;
            resourceInputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
            resourceInputs["routeEntryType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteEntry resources.
 */
export interface RouteEntryState {
    /**
     * The next hop of the destination route.
     */
    nextHop?: pulumi.Input<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    publishVpc?: pulumi.Input<boolean>;
    /**
     * The destination network segment of the destination route.
     */
    routeDest?: pulumi.Input<string>;
    /**
     * (Available in 1.161.0+) The type of the vpn route entry.
     */
    routeEntryType?: pulumi.Input<string>;
    /**
     * (Available in 1.161.0+) The status of the vpn route entry.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the vpn gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
    /**
     * The value should be 0 or 100.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RouteEntry resource.
 */
export interface RouteEntryArgs {
    /**
     * The next hop of the destination route.
     */
    nextHop: pulumi.Input<string>;
    /**
     * Whether to issue the destination route to the VPC.
     */
    publishVpc: pulumi.Input<boolean>;
    /**
     * The destination network segment of the destination route.
     */
    routeDest: pulumi.Input<string>;
    /**
     * The id of the vpn gateway.
     */
    vpnGatewayId: pulumi.Input<string>;
    /**
     * The value should be 0 or 100.
     */
    weight: pulumi.Input<number>;
}
