// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Nat Ip resource.
 *
 * For information about VPC Nat Ip and how to use it, see [What is Nat Ip](https://www.alibabacloud.com/help/doc-detail/281976.htm).
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     vpcName: "example_value",
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const exampleSwitch = new alicloud.vpc.Switch("exampleSwitch", {
 *     vpcId: alicloud_vpc["default"].id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: exampleZones.then(exampleZones => exampleZones.zones?[0]?.id),
 *     vswitchName: "example_value",
 * });
 * const exampleNatGateway = new alicloud.vpc.NatGateway("exampleNatGateway", {
 *     vpcId: exampleNetwork.id,
 *     internetChargeType: "PayByLcu",
 *     natGatewayName: "example_value",
 *     description: "example_value",
 *     natType: "Enhanced",
 *     vswitchId: exampleSwitch.id,
 *     networkType: "intranet",
 * });
 * const exampleNatIpCidr = new alicloud.vpc.NatIpCidr("exampleNatIpCidr", {
 *     natIpCidr: "192.168.0.0/16",
 *     natGatewayId: exampleNatGateway.id,
 *     natIpCidrDescription: "example_value",
 *     natIpCidrName: "example_value",
 * });
 * const exampleNatIp = new alicloud.vpc.NatIp("exampleNatIp", {
 *     natIp: "192.168.0.37",
 *     natGatewayId: exampleNatGateway.id,
 *     natIpDescription: "example_value",
 *     natIpName: "example_value",
 *     natIpCidr: exampleNatIpCidr.natIpCidr,
 * });
 * ```
 *
 * ## Import
 *
 * VPC Nat Ip can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/natIp:NatIp example <nat_gateway_id>:<nat_ip_id>
 * ```
 */
export class NatIp extends pulumi.CustomResource {
    /**
     * Get an existing NatIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatIpState, opts?: pulumi.CustomResourceOptions): NatIp {
        return new NatIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/natIp:NatIp';

    /**
     * Returns true if the given object is an instance of NatIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatIp.__pulumiType;
    }

    /**
     * Specifies whether to check the validity of the request without actually making the request.
     */
    public readonly dryRun!: pulumi.Output<boolean>;
    /**
     * The ID of the Virtual Private Cloud (VPC) NAT gateway for which you want to create the NAT IP address.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * The NAT IP address that you want to create. If you do not specify an IP address, the system selects a random IP address from the specified CIDR block.
     */
    public readonly natIp!: pulumi.Output<string>;
    /**
     * NAT IP ADDRESS of the address segment.
     */
    public readonly natIpCidr!: pulumi.Output<string | undefined>;
    /**
     * The ID of the CIDR block to which the NAT IP address belongs.
     */
    public readonly natIpCidrId!: pulumi.Output<string | undefined>;
    /**
     * NAT IP ADDRESS description of information. Length is from `2` to `256` characters, must start with a letter or the Chinese at the beginning, but not at the` http://` Or `https://` at the beginning.
     */
    public readonly natIpDescription!: pulumi.Output<string | undefined>;
    /**
     * Ihe ID of the Nat Ip.
     */
    public /*out*/ readonly natIpId!: pulumi.Output<string>;
    /**
     * NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
     */
    public readonly natIpName!: pulumi.Output<string | undefined>;
    /**
     * The status of the NAT IP address. Valid values: `Available`, `Deleting`, `Creating` and `Deleted`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a NatIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatIpArgs | NatIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NatIpState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            resourceInputs["natIp"] = state ? state.natIp : undefined;
            resourceInputs["natIpCidr"] = state ? state.natIpCidr : undefined;
            resourceInputs["natIpCidrId"] = state ? state.natIpCidrId : undefined;
            resourceInputs["natIpDescription"] = state ? state.natIpDescription : undefined;
            resourceInputs["natIpId"] = state ? state.natIpId : undefined;
            resourceInputs["natIpName"] = state ? state.natIpName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as NatIpArgs | undefined;
            if ((!args || args.natGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            resourceInputs["natIp"] = args ? args.natIp : undefined;
            resourceInputs["natIpCidr"] = args ? args.natIpCidr : undefined;
            resourceInputs["natIpCidrId"] = args ? args.natIpCidrId : undefined;
            resourceInputs["natIpDescription"] = args ? args.natIpDescription : undefined;
            resourceInputs["natIpName"] = args ? args.natIpName : undefined;
            resourceInputs["natIpId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NatIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatIp resources.
 */
export interface NatIpState {
    /**
     * Specifies whether to check the validity of the request without actually making the request.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Private Cloud (VPC) NAT gateway for which you want to create the NAT IP address.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * The NAT IP address that you want to create. If you do not specify an IP address, the system selects a random IP address from the specified CIDR block.
     */
    natIp?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS of the address segment.
     */
    natIpCidr?: pulumi.Input<string>;
    /**
     * The ID of the CIDR block to which the NAT IP address belongs.
     */
    natIpCidrId?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS description of information. Length is from `2` to `256` characters, must start with a letter or the Chinese at the beginning, but not at the` http://` Or `https://` at the beginning.
     */
    natIpDescription?: pulumi.Input<string>;
    /**
     * Ihe ID of the Nat Ip.
     */
    natIpId?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
     */
    natIpName?: pulumi.Input<string>;
    /**
     * The status of the NAT IP address. Valid values: `Available`, `Deleting`, `Creating` and `Deleted`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NatIp resource.
 */
export interface NatIpArgs {
    /**
     * Specifies whether to check the validity of the request without actually making the request.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the Virtual Private Cloud (VPC) NAT gateway for which you want to create the NAT IP address.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * The NAT IP address that you want to create. If you do not specify an IP address, the system selects a random IP address from the specified CIDR block.
     */
    natIp?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS of the address segment.
     */
    natIpCidr?: pulumi.Input<string>;
    /**
     * The ID of the CIDR block to which the NAT IP address belongs.
     */
    natIpCidrId?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS description of information. Length is from `2` to `256` characters, must start with a letter or the Chinese at the beginning, but not at the` http://` Or `https://` at the beginning.
     */
    natIpDescription?: pulumi.Input<string>;
    /**
     * NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
     */
    natIpName?: pulumi.Input<string>;
}
