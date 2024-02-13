// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Vpc Ipv6 Gateway resource. Gateway Based on Internet Protocol Version 6.
 *
 * For information about Vpc Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv6gateway).
 *
 * > **NOTE:** Available in v1.142.0+.
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
 * const name = config.get("name") || "tf-testacc-example";
 * const defaultVpc = new alicloud.vpc.Network("defaultVpc", {
 *     description: "tf-testacc",
 *     enableIpv6: true,
 * });
 * const defaultRg = new alicloud.resourcemanager.ResourceGroup("defaultRg", {
 *     displayName: "tf-testacc-ipv6gateway503",
 *     resourceGroupName: `${name}1`,
 * });
 * const changeRg = new alicloud.resourcemanager.ResourceGroup("changeRg", {
 *     displayName: "tf-testacc-ipv6gateway311",
 *     resourceGroupName: `${name}2`,
 * });
 * const _default = new alicloud.vpc.Ipv6Gateway("default", {
 *     description: "test",
 *     ipv6GatewayName: name,
 *     vpcId: defaultVpc.id,
 *     resourceGroupId: defaultRg.id,
 * });
 * ```
 *
 * ## Import
 *
 * Vpc Ipv6 Gateway can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:vpc/ipv6Gateway:Ipv6Gateway example <id>
 * ```
 */
export class Ipv6Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Ipv6Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv6GatewayState, opts?: pulumi.CustomResourceOptions): Ipv6Gateway {
        return new Ipv6Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/ipv6Gateway:Ipv6Gateway';

    /**
     * Returns true if the given object is an instance of Ipv6Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv6Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv6Gateway.__pulumiType;
    }

    /**
     * The status of the IPv6 gateway.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the IPv6 gateway. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The expiration time of IPv6 gateway.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * The charge type of IPv6 gateway.
     */
    public /*out*/ readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * Resource primary key attribute field.
     */
    public /*out*/ readonly ipv6GatewayId!: pulumi.Output<string>;
    /**
     * The name of the IPv6 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
     */
    public readonly ipv6GatewayName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     *
     * @deprecated Field 'Spec' has been deprecated from provider version 1.205.0. IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     */
    public readonly spec!: pulumi.Output<string>;
    /**
     * The status of the resource. Valid values: Available, Pending and Deleting.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags for the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Ipv6Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ipv6GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv6GatewayArgs | Ipv6GatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv6GatewayState | undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["ipv6GatewayId"] = state ? state.ipv6GatewayId : undefined;
            resourceInputs["ipv6GatewayName"] = state ? state.ipv6GatewayName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as Ipv6GatewayArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipv6GatewayName"] = args ? args.ipv6GatewayName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["instanceChargeType"] = undefined /*out*/;
            resourceInputs["ipv6GatewayId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv6Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv6Gateway resources.
 */
export interface Ipv6GatewayState {
    /**
     * The status of the IPv6 gateway.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the IPv6 gateway. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * The expiration time of IPv6 gateway.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * The charge type of IPv6 gateway.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * Resource primary key attribute field.
     */
    ipv6GatewayId?: pulumi.Input<string>;
    /**
     * The name of the IPv6 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
     */
    ipv6GatewayName?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     *
     * @deprecated Field 'Spec' has been deprecated from provider version 1.205.0. IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     */
    spec?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: Available, Pending and Deleting.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags for the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv6Gateway resource.
 */
export interface Ipv6GatewayArgs {
    /**
     * The description of the IPv6 gateway. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the IPv6 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
     */
    ipv6GatewayName?: pulumi.Input<string>;
    /**
     * The ID of the resource group to which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     *
     * @deprecated Field 'Spec' has been deprecated from provider version 1.205.0. IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
     */
    spec?: pulumi.Input<string>;
    /**
     * The tags for the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
     */
    vpcId: pulumi.Input<string>;
}
