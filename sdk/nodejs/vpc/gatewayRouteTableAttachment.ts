// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Gateway Route Table Attachment resource.
 *
 * For information about VPC Gateway Route Table Attachment and how to use it, see [What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 *
 * > **NOTE:** Available in v1.194.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const exampleNetwork = new alicloud.vpc.Network("exampleNetwork", {
 *     cidrBlock: "172.16.0.0/12",
 *     vpcName: "terraform-example",
 * });
 * const exampleRouteTable = new alicloud.vpc.RouteTable("exampleRouteTable", {
 *     vpcId: exampleNetwork.id,
 *     routeTableName: "terraform-example",
 *     description: "terraform-example",
 *     associateType: "Gateway",
 * });
 * const exampleIpv4Gateway = new alicloud.vpc.Ipv4Gateway("exampleIpv4Gateway", {
 *     ipv4GatewayName: "terraform-example",
 *     vpcId: exampleNetwork.id,
 *     enabled: true,
 * });
 * const exampleGatewayRouteTableAttachment = new alicloud.vpc.GatewayRouteTableAttachment("exampleGatewayRouteTableAttachment", {
 *     ipv4GatewayId: exampleIpv4Gateway.id,
 *     routeTableId: exampleRouteTable.id,
 * });
 * ```
 *
 * ## Import
 *
 * VPC Gateway Route Table Attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment example <route_table_id>:<ipv4_gateway_id>
 * ```
 */
export class GatewayRouteTableAttachment extends pulumi.CustomResource {
    /**
     * Get an existing GatewayRouteTableAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayRouteTableAttachmentState, opts?: pulumi.CustomResourceOptions): GatewayRouteTableAttachment {
        return new GatewayRouteTableAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment';

    /**
     * Returns true if the given object is an instance of GatewayRouteTableAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayRouteTableAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayRouteTableAttachment.__pulumiType;
    }

    /**
     * The creation time of the resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the IPv4 Gateway instance.
     */
    public readonly ipv4GatewayId!: pulumi.Output<string>;
    /**
     * The ID of the Gateway route table to be bound.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The status of the IPv4 Gateway instance. Value:
     * - **Creating**: The function is being created.
     * - **Created**: Created and available.
     * - **Modifying**: is being modified.
     * - **Deleting**: Deleting.
     * - **Deleted**: Deleted.
     * - **Activating**: enabled.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a GatewayRouteTableAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayRouteTableAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayRouteTableAttachmentArgs | GatewayRouteTableAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayRouteTableAttachmentState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["ipv4GatewayId"] = state ? state.ipv4GatewayId : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as GatewayRouteTableAttachmentArgs | undefined;
            if ((!args || args.ipv4GatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipv4GatewayId'");
            }
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["ipv4GatewayId"] = args ? args.ipv4GatewayId : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewayRouteTableAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayRouteTableAttachment resources.
 */
export interface GatewayRouteTableAttachmentState {
    /**
     * The creation time of the resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the IPv4 Gateway instance.
     */
    ipv4GatewayId?: pulumi.Input<string>;
    /**
     * The ID of the Gateway route table to be bound.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The status of the IPv4 Gateway instance. Value:
     * - **Creating**: The function is being created.
     * - **Created**: Created and available.
     * - **Modifying**: is being modified.
     * - **Deleting**: Deleting.
     * - **Deleted**: Deleted.
     * - **Activating**: enabled.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewayRouteTableAttachment resource.
 */
export interface GatewayRouteTableAttachmentArgs {
    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID of the IPv4 Gateway instance.
     */
    ipv4GatewayId: pulumi.Input<string>;
    /**
     * The ID of the Gateway route table to be bound.
     */
    routeTableId: pulumi.Input<string>;
}
