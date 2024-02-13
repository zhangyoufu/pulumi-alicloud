// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Ipv6 Egress Rule resource. IPv6 address addition only active exit rule.
 *
 * For information about VPC Ipv6 Egress Rule and how to use it, see [What is Ipv6 Egress Rule](https://www.alibabacloud.com/help/doc-detail/102200.htm).
 *
 * > **NOTE:** Available since v1.142.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     systemDiskCategory: "cloud_efficiency",
 *     cpuCoreCount: 4,
 *     minimumEniIpv6AddressQuantity: 1,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     enableIpv6: true,
 *     cidrBlock: "172.16.0.0/12",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/21",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     vswitchName: name,
 *     ipv6CidrBlockMask: 64,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {
 *     description: name,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstance = new alicloud.ecs.Instance("defaultInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     ipv6AddressCount: 1,
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     systemDiskCategory: "cloud_efficiency",
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceName: name,
 *     vswitchId: defaultSwitch.id,
 *     internetMaxBandwidthOut: 10,
 *     securityGroups: [defaultSecurityGroup.id],
 * });
 * const defaultIpv6Gateway = new alicloud.vpc.Ipv6Gateway("defaultIpv6Gateway", {
 *     ipv6GatewayName: name,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultIpv6Addresses = alicloud.vpc.getIpv6AddressesOutput({
 *     associatedInstanceId: defaultInstance.id,
 *     status: "Available",
 * });
 * const defaultIpv6InternetBandwidth = new alicloud.vpc.Ipv6InternetBandwidth("defaultIpv6InternetBandwidth", {
 *     ipv6AddressId: defaultIpv6Addresses.apply(defaultIpv6Addresses => defaultIpv6Addresses.addresses?.[0]?.id),
 *     ipv6GatewayId: defaultIpv6Gateway.ipv6GatewayId,
 *     internetChargeType: "PayByBandwidth",
 *     bandwidth: 20,
 * });
 * const defaultIpv6EgressRule = new alicloud.vpc.Ipv6EgressRule("defaultIpv6EgressRule", {
 *     instanceId: defaultIpv6InternetBandwidth.ipv6AddressId,
 *     ipv6EgressRuleName: name,
 *     description: name,
 *     ipv6GatewayId: defaultIpv6InternetBandwidth.ipv6GatewayId,
 *     instanceType: "Ipv6Address",
 * });
 * ```
 *
 * ## Import
 *
 * VPC Ipv6 Egress Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/ipv6EgressRule:Ipv6EgressRule example <ipv6_gateway_id>:<ipv6_egress_rule_id>
 * ```
 */
export class Ipv6EgressRule extends pulumi.CustomResource {
    /**
     * Get an existing Ipv6EgressRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv6EgressRuleState, opts?: pulumi.CustomResourceOptions): Ipv6EgressRule {
        return new Ipv6EgressRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/ipv6EgressRule:Ipv6EgressRule';

    /**
     * Returns true if the given object is an instance of Ipv6EgressRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv6EgressRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv6EgressRule.__pulumiType;
    }

    /**
     * The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the IPv6 address to which you want to apply the egress-only rule.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     */
    public readonly ipv6EgressRuleName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the IPv6 gateway.
     */
    public readonly ipv6GatewayId!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Ipv6EgressRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ipv6EgressRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv6EgressRuleArgs | Ipv6EgressRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv6EgressRuleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["ipv6EgressRuleName"] = state ? state.ipv6EgressRuleName : undefined;
            resourceInputs["ipv6GatewayId"] = state ? state.ipv6GatewayId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as Ipv6EgressRuleArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.ipv6GatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipv6GatewayId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["ipv6EgressRuleName"] = args ? args.ipv6EgressRuleName : undefined;
            resourceInputs["ipv6GatewayId"] = args ? args.ipv6GatewayId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv6EgressRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv6EgressRule resources.
 */
export interface Ipv6EgressRuleState {
    /**
     * The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the IPv6 address to which you want to apply the egress-only rule.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     */
    ipv6EgressRuleName?: pulumi.Input<string>;
    /**
     * The ID of the IPv6 gateway.
     */
    ipv6GatewayId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv6EgressRule resource.
 */
export interface Ipv6EgressRuleArgs {
    /**
     * The description of the egress-only rule. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the IPv6 address to which you want to apply the egress-only rule.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The type of instance to which you want to apply the egress-only rule. Valid values: `Ipv6Address`. `Ipv6Address` (default): an IPv6 address.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the egress-only rule. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     */
    ipv6EgressRuleName?: pulumi.Input<string>;
    /**
     * The ID of the IPv6 gateway.
     */
    ipv6GatewayId: pulumi.Input<string>;
}
