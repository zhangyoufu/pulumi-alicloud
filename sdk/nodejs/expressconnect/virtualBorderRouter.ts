// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Express Connect Virtual Border Router resource.
 *
 * For information about Express Connect Virtual Border Router and how to use it, see [What is Virtual Border Router](https://www.alibabacloud.com/help/en/doc-detail/44854.htm).
 *
 * > **NOTE:** Available in v1.134.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const nameRegex = alicloud.expressconnect.getPhysicalConnections({
 *     nameRegex: "^my-PhysicalConnection",
 * });
 * const example = new alicloud.expressconnect.VirtualBorderRouter("example", {
 *     localGatewayIp: "10.0.0.1",
 *     peerGatewayIp: "10.0.0.2",
 *     peeringSubnetMask: "255.255.255.252",
 *     physicalConnectionId: nameRegex.then(nameRegex => nameRegex.connections?.[0]?.id),
 *     virtualBorderRouterName: "example_value",
 *     vlanId: 1,
 *     minRxInterval: 1000,
 *     minTxInterval: 1000,
 *     detectMultiplier: 10,
 * });
 * ```
 *
 * ## Import
 *
 * Express Connect Virtual Border Router can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:expressconnect/virtualBorderRouter:VirtualBorderRouter example <id>
 * ```
 */
export class VirtualBorderRouter extends pulumi.CustomResource {
    /**
     * Get an existing VirtualBorderRouter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualBorderRouterState, opts?: pulumi.CustomResourceOptions): VirtualBorderRouter {
        return new VirtualBorderRouter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:expressconnect/virtualBorderRouter:VirtualBorderRouter';

    /**
     * Returns true if the given object is an instance of VirtualBorderRouter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualBorderRouter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualBorderRouter.__pulumiType;
    }

    /**
     * The associated physical connections.
     */
    public readonly associatedPhysicalConnections!: pulumi.Output<string | undefined>;
    /**
     * The bandwidth.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Operators for physical connection circuit provided coding.
     */
    public readonly circuitCode!: pulumi.Output<string | undefined>;
    /**
     * The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
     */
    public readonly detectMultiplier!: pulumi.Output<number>;
    /**
     * Whether to Enable IPv6. Valid values: `false`, `true`.
     */
    public readonly enableIpv6!: pulumi.Output<boolean>;
    /**
     * Whether cross account border routers are included. Valid values: `false`, `true`. Default: `true`.
     */
    public readonly includeCrossAccountVbr!: pulumi.Output<boolean>;
    /**
     * Alibaba Cloud-Connected IPv4 address.
     */
    public readonly localGatewayIp!: pulumi.Output<string>;
    /**
     * Alibaba Cloud-Connected IPv6 Address.
     */
    public readonly localIpv6GatewayIp!: pulumi.Output<string | undefined>;
    /**
     * Configure BFD packet reception interval of values include: 200~1000, unit: ms.
     */
    public readonly minRxInterval!: pulumi.Output<number>;
    /**
     * Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
     */
    public readonly minTxInterval!: pulumi.Output<number>;
    /**
     * The Client-Side Interconnection IPv4 Address.
     */
    public readonly peerGatewayIp!: pulumi.Output<string>;
    /**
     * The Client-Side Interconnection IPv6 Address.
     */
    public readonly peerIpv6GatewayIp!: pulumi.Output<string | undefined>;
    /**
     * Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
     */
    public readonly peeringIpv6SubnetMask!: pulumi.Output<string | undefined>;
    /**
     * Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
     */
    public readonly peeringSubnetMask!: pulumi.Output<string>;
    /**
     * The ID of the Physical Connection to Which the ID.
     */
    public readonly physicalConnectionId!: pulumi.Output<string>;
    /**
     * (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
     */
    public /*out*/ readonly routeTableId!: pulumi.Output<string>;
    /**
     * The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The vbr owner id.
     */
    public readonly vbrOwnerId!: pulumi.Output<string | undefined>;
    /**
     * The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
     */
    public readonly virtualBorderRouterName!: pulumi.Output<string | undefined>;
    /**
     * The VLAN ID of the VBR. Value range: 0~2999.
     */
    public readonly vlanId!: pulumi.Output<number>;

    /**
     * Create a VirtualBorderRouter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualBorderRouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualBorderRouterArgs | VirtualBorderRouterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualBorderRouterState | undefined;
            resourceInputs["associatedPhysicalConnections"] = state ? state.associatedPhysicalConnections : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["circuitCode"] = state ? state.circuitCode : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["detectMultiplier"] = state ? state.detectMultiplier : undefined;
            resourceInputs["enableIpv6"] = state ? state.enableIpv6 : undefined;
            resourceInputs["includeCrossAccountVbr"] = state ? state.includeCrossAccountVbr : undefined;
            resourceInputs["localGatewayIp"] = state ? state.localGatewayIp : undefined;
            resourceInputs["localIpv6GatewayIp"] = state ? state.localIpv6GatewayIp : undefined;
            resourceInputs["minRxInterval"] = state ? state.minRxInterval : undefined;
            resourceInputs["minTxInterval"] = state ? state.minTxInterval : undefined;
            resourceInputs["peerGatewayIp"] = state ? state.peerGatewayIp : undefined;
            resourceInputs["peerIpv6GatewayIp"] = state ? state.peerIpv6GatewayIp : undefined;
            resourceInputs["peeringIpv6SubnetMask"] = state ? state.peeringIpv6SubnetMask : undefined;
            resourceInputs["peeringSubnetMask"] = state ? state.peeringSubnetMask : undefined;
            resourceInputs["physicalConnectionId"] = state ? state.physicalConnectionId : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vbrOwnerId"] = state ? state.vbrOwnerId : undefined;
            resourceInputs["virtualBorderRouterName"] = state ? state.virtualBorderRouterName : undefined;
            resourceInputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as VirtualBorderRouterArgs | undefined;
            if ((!args || args.localGatewayIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localGatewayIp'");
            }
            if ((!args || args.peerGatewayIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerGatewayIp'");
            }
            if ((!args || args.peeringSubnetMask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peeringSubnetMask'");
            }
            if ((!args || args.physicalConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'physicalConnectionId'");
            }
            if ((!args || args.vlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlanId'");
            }
            resourceInputs["associatedPhysicalConnections"] = args ? args.associatedPhysicalConnections : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["circuitCode"] = args ? args.circuitCode : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["detectMultiplier"] = args ? args.detectMultiplier : undefined;
            resourceInputs["enableIpv6"] = args ? args.enableIpv6 : undefined;
            resourceInputs["includeCrossAccountVbr"] = args ? args.includeCrossAccountVbr : undefined;
            resourceInputs["localGatewayIp"] = args ? args.localGatewayIp : undefined;
            resourceInputs["localIpv6GatewayIp"] = args ? args.localIpv6GatewayIp : undefined;
            resourceInputs["minRxInterval"] = args ? args.minRxInterval : undefined;
            resourceInputs["minTxInterval"] = args ? args.minTxInterval : undefined;
            resourceInputs["peerGatewayIp"] = args ? args.peerGatewayIp : undefined;
            resourceInputs["peerIpv6GatewayIp"] = args ? args.peerIpv6GatewayIp : undefined;
            resourceInputs["peeringIpv6SubnetMask"] = args ? args.peeringIpv6SubnetMask : undefined;
            resourceInputs["peeringSubnetMask"] = args ? args.peeringSubnetMask : undefined;
            resourceInputs["physicalConnectionId"] = args ? args.physicalConnectionId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vbrOwnerId"] = args ? args.vbrOwnerId : undefined;
            resourceInputs["virtualBorderRouterName"] = args ? args.virtualBorderRouterName : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["routeTableId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualBorderRouter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualBorderRouter resources.
 */
export interface VirtualBorderRouterState {
    /**
     * The associated physical connections.
     */
    associatedPhysicalConnections?: pulumi.Input<string>;
    /**
     * The bandwidth.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Operators for physical connection circuit provided coding.
     */
    circuitCode?: pulumi.Input<string>;
    /**
     * The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
     */
    description?: pulumi.Input<string>;
    /**
     * Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
     */
    detectMultiplier?: pulumi.Input<number>;
    /**
     * Whether to Enable IPv6. Valid values: `false`, `true`.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * Whether cross account border routers are included. Valid values: `false`, `true`. Default: `true`.
     */
    includeCrossAccountVbr?: pulumi.Input<boolean>;
    /**
     * Alibaba Cloud-Connected IPv4 address.
     */
    localGatewayIp?: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv6 Address.
     */
    localIpv6GatewayIp?: pulumi.Input<string>;
    /**
     * Configure BFD packet reception interval of values include: 200~1000, unit: ms.
     */
    minRxInterval?: pulumi.Input<number>;
    /**
     * Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
     */
    minTxInterval?: pulumi.Input<number>;
    /**
     * The Client-Side Interconnection IPv4 Address.
     */
    peerGatewayIp?: pulumi.Input<string>;
    /**
     * The Client-Side Interconnection IPv6 Address.
     */
    peerIpv6GatewayIp?: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
     */
    peeringIpv6SubnetMask?: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
     */
    peeringSubnetMask?: pulumi.Input<string>;
    /**
     * The ID of the Physical Connection to Which the ID.
     */
    physicalConnectionId?: pulumi.Input<string>;
    /**
     * (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
     */
    status?: pulumi.Input<string>;
    /**
     * The vbr owner id.
     */
    vbrOwnerId?: pulumi.Input<string>;
    /**
     * The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
     */
    virtualBorderRouterName?: pulumi.Input<string>;
    /**
     * The VLAN ID of the VBR. Value range: 0~2999.
     */
    vlanId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VirtualBorderRouter resource.
 */
export interface VirtualBorderRouterArgs {
    /**
     * The associated physical connections.
     */
    associatedPhysicalConnections?: pulumi.Input<string>;
    /**
     * The bandwidth.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Operators for physical connection circuit provided coding.
     */
    circuitCode?: pulumi.Input<string>;
    /**
     * The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
     */
    description?: pulumi.Input<string>;
    /**
     * Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
     */
    detectMultiplier?: pulumi.Input<number>;
    /**
     * Whether to Enable IPv6. Valid values: `false`, `true`.
     */
    enableIpv6?: pulumi.Input<boolean>;
    /**
     * Whether cross account border routers are included. Valid values: `false`, `true`. Default: `true`.
     */
    includeCrossAccountVbr?: pulumi.Input<boolean>;
    /**
     * Alibaba Cloud-Connected IPv4 address.
     */
    localGatewayIp: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv6 Address.
     */
    localIpv6GatewayIp?: pulumi.Input<string>;
    /**
     * Configure BFD packet reception interval of values include: 200~1000, unit: ms.
     */
    minRxInterval?: pulumi.Input<number>;
    /**
     * Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
     */
    minTxInterval?: pulumi.Input<number>;
    /**
     * The Client-Side Interconnection IPv4 Address.
     */
    peerGatewayIp: pulumi.Input<string>;
    /**
     * The Client-Side Interconnection IPv6 Address.
     */
    peerIpv6GatewayIp?: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
     */
    peeringIpv6SubnetMask?: pulumi.Input<string>;
    /**
     * Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
     */
    peeringSubnetMask: pulumi.Input<string>;
    /**
     * The ID of the Physical Connection to Which the ID.
     */
    physicalConnectionId: pulumi.Input<string>;
    /**
     * The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
     */
    status?: pulumi.Input<string>;
    /**
     * The vbr owner id.
     */
    vbrOwnerId?: pulumi.Input<string>;
    /**
     * The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
     */
    virtualBorderRouterName?: pulumi.Input<string>;
    /**
     * The VLAN ID of the VBR. Value range: 0~2999.
     */
    vlanId: pulumi.Input<number>;
}
