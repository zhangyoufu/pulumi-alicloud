// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * The router interface can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/routerInterface:RouterInterface interface ri-abc123456
 * ```
 */
export class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceState, opts?: pulumi.CustomResourceOptions): RouterInterface {
        return new RouterInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/routerInterface:RouterInterface';

    /**
     * Returns true if the given object is an instance of RouterInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterInterface.__pulumiType;
    }

    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    public /*out*/ readonly accessPointId!: pulumi.Output<string>;
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    public readonly healthCheckSourceIp!: pulumi.Output<string | undefined>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    public readonly healthCheckTargetIp!: pulumi.Output<string | undefined>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    public readonly oppositeAccessPointId!: pulumi.Output<string | undefined>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     */
    public /*out*/ readonly oppositeInterfaceId!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_owner_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_owner_id' instead.
     */
    public /*out*/ readonly oppositeInterfaceOwnerId!: pulumi.Output<string>;
    /**
     * The Region of peer side.
     */
    public readonly oppositeRegion!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_router_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     */
    public /*out*/ readonly oppositeRouterId!: pulumi.Output<string>;
    /**
     * It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     *
     * @deprecated Attribute 'opposite_router_type' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     */
    public /*out*/ readonly oppositeRouterType!: pulumi.Output<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The Router ID.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    public readonly routerType!: pulumi.Output<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    public readonly specification!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterInterfaceArgs | RouterInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterInterfaceState | undefined;
            inputs["accessPointId"] = state ? state.accessPointId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["healthCheckSourceIp"] = state ? state.healthCheckSourceIp : undefined;
            inputs["healthCheckTargetIp"] = state ? state.healthCheckTargetIp : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["oppositeAccessPointId"] = state ? state.oppositeAccessPointId : undefined;
            inputs["oppositeInterfaceId"] = state ? state.oppositeInterfaceId : undefined;
            inputs["oppositeInterfaceOwnerId"] = state ? state.oppositeInterfaceOwnerId : undefined;
            inputs["oppositeRegion"] = state ? state.oppositeRegion : undefined;
            inputs["oppositeRouterId"] = state ? state.oppositeRouterId : undefined;
            inputs["oppositeRouterType"] = state ? state.oppositeRouterType : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
            inputs["routerType"] = state ? state.routerType : undefined;
            inputs["specification"] = state ? state.specification : undefined;
        } else {
            const args = argsOrState as RouterInterfaceArgs | undefined;
            if ((!args || args.oppositeRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oppositeRegion'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.routerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerId'");
            }
            if ((!args || args.routerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerType'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["healthCheckSourceIp"] = args ? args.healthCheckSourceIp : undefined;
            inputs["healthCheckTargetIp"] = args ? args.healthCheckTargetIp : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oppositeAccessPointId"] = args ? args.oppositeAccessPointId : undefined;
            inputs["oppositeRegion"] = args ? args.oppositeRegion : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["routerId"] = args ? args.routerId : undefined;
            inputs["routerType"] = args ? args.routerType : undefined;
            inputs["specification"] = args ? args.specification : undefined;
            inputs["accessPointId"] = undefined /*out*/;
            inputs["oppositeInterfaceId"] = undefined /*out*/;
            inputs["oppositeInterfaceOwnerId"] = undefined /*out*/;
            inputs["oppositeRouterId"] = undefined /*out*/;
            inputs["oppositeRouterType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouterInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterInterface resources.
 */
export interface RouterInterfaceState {
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    readonly accessPointId?: pulumi.Input<string>;
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    readonly healthCheckSourceIp?: pulumi.Input<string>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    readonly healthCheckTargetIp?: pulumi.Input<string>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    readonly oppositeAccessPointId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     */
    readonly oppositeInterfaceId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.
     *
     * @deprecated Attribute 'opposite_interface_owner_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_owner_id' instead.
     */
    readonly oppositeInterfaceOwnerId?: pulumi.Input<string>;
    /**
     * The Region of peer side.
     */
    readonly oppositeRegion?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     *
     * @deprecated Attribute 'opposite_router_id' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.
     */
    readonly oppositeRouterId?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     *
     * @deprecated Attribute 'opposite_router_type' has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_type' instead.
     */
    readonly oppositeRouterType?: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The Router ID.
     */
    readonly routerId?: pulumi.Input<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    readonly routerType?: pulumi.Input<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    readonly specification?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterInterface resource.
 */
export interface RouterInterfaceArgs {
    /**
     * Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckTargetIp` must be specified at the same time.
     */
    readonly healthCheckSourceIp?: pulumi.Input<string>;
    /**
     * Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when `routerType` is `VBR`. The IP must be an unused IP in the local VPC. It and `healthCheckSourceIp` must be specified at the same time.
     */
    readonly healthCheckTargetIp?: pulumi.Input<string>;
    /**
     * The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and oppositeRegion are the same.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * It has been deprecated from version 1.11.0.
     *
     * @deprecated Attribute 'opposite_access_point_id' has been deprecated from version 1.11.0.
     */
    readonly oppositeAccessPointId?: pulumi.Input<string>;
    /**
     * The Region of peer side.
     */
    readonly oppositeRegion: pulumi.Input<string>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The role the router interface plays. Optional value: `InitiatingSide`, `AcceptingSide`.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The Router ID.
     */
    readonly routerId: pulumi.Input<string>;
    /**
     * Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.
     */
    readonly routerType: pulumi.Input<string>;
    /**
     * Specification of router interfaces. It is valid when `role` is `InitiatingSide`. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).
     */
    readonly specification?: pulumi.Input<string>;
}
