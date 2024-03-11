// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN transit router VBR attachment resource that associate the VBR with the CEN instance.[What is Cen Transit Router VBR Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutervbrattachment)
 *
 * > **NOTE:** Available since v1.126.0.
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
 * const name = config.get("name") || "terraform-example";
 * const defaultInstance = new alicloud.cen.Instance("defaultInstance", {
 *     cenInstanceName: name,
 *     protectionLevel: "REDUCED",
 * });
 * const defaultTransitRouter = new alicloud.cen.TransitRouter("defaultTransitRouter", {cenId: defaultInstance.id});
 * const nameRegex = alicloud.expressconnect.getPhysicalConnections({
 *     nameRegex: "^preserved-NODELETING",
 * });
 * const defaultVirtualBorderRouter = new alicloud.expressconnect.VirtualBorderRouter("defaultVirtualBorderRouter", {
 *     localGatewayIp: "10.0.0.1",
 *     peerGatewayIp: "10.0.0.2",
 *     peeringSubnetMask: "255.255.255.252",
 *     physicalConnectionId: nameRegex.then(nameRegex => nameRegex.connections?.[0]?.id),
 *     virtualBorderRouterName: name,
 *     vlanId: 2420,
 *     minRxInterval: 1000,
 *     minTxInterval: 1000,
 *     detectMultiplier: 10,
 * });
 * const defaultTransitRouterVbrAttachment = new alicloud.cen.TransitRouterVbrAttachment("defaultTransitRouterVbrAttachment", {
 *     transitRouterId: defaultTransitRouter.transitRouterId,
 *     transitRouterAttachmentName: "example",
 *     transitRouterAttachmentDescription: "example",
 *     vbrId: defaultVirtualBorderRouter.id,
 *     cenId: defaultInstance.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * CEN transit router VBR attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cen/transitRouterVbrAttachment:TransitRouterVbrAttachment example tr-********:tr-attach-********
 * ```
 */
export class TransitRouterVbrAttachment extends pulumi.CustomResource {
    /**
     * Get an existing TransitRouterVbrAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitRouterVbrAttachmentState, opts?: pulumi.CustomResourceOptions): TransitRouterVbrAttachment {
        return new TransitRouterVbrAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/transitRouterVbrAttachment:TransitRouterVbrAttachment';

    /**
     * Returns true if the given object is an instance of TransitRouterVbrAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitRouterVbrAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitRouterVbrAttachment.__pulumiType;
    }

    /**
     * Auto publish route enabled.Default value is `false`.
     */
    public readonly autoPublishRouteEnabled!: pulumi.Output<boolean>;
    /**
     * The ID of the CEN.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The resource type of the transit router vbr attachment.  Valid values: `VPC`, `CCN`, `VBR`, `TR`.
     *
     * ->**NOTE:** Ensure that the vbr is not used in Express Connect.
     */
    public readonly resourceType!: pulumi.Output<string | undefined>;
    /**
     * Whether to enabled route table association. The system default value is `true`.
     */
    public readonly routeTableAssociationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enabled route table propagation. The system default value is `true`.
     */
    public readonly routeTablePropagationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The associating status of the network.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The description of the transit router vbr attachment.
     */
    public readonly transitRouterAttachmentDescription!: pulumi.Output<string | undefined>;
    /**
     * The id of the transit router vbr attachment.
     */
    public /*out*/ readonly transitRouterAttachmentId!: pulumi.Output<string>;
    /**
     * The name of the transit router vbr attachment.
     */
    public readonly transitRouterAttachmentName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the transit router.
     */
    public readonly transitRouterId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VBR.
     */
    public readonly vbrId!: pulumi.Output<string>;
    /**
     * The owner id of the transit router vbr attachment.
     */
    public readonly vbrOwnerId!: pulumi.Output<string>;

    /**
     * Create a TransitRouterVbrAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitRouterVbrAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitRouterVbrAttachmentArgs | TransitRouterVbrAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitRouterVbrAttachmentState | undefined;
            resourceInputs["autoPublishRouteEnabled"] = state ? state.autoPublishRouteEnabled : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["routeTableAssociationEnabled"] = state ? state.routeTableAssociationEnabled : undefined;
            resourceInputs["routeTablePropagationEnabled"] = state ? state.routeTablePropagationEnabled : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["transitRouterAttachmentDescription"] = state ? state.transitRouterAttachmentDescription : undefined;
            resourceInputs["transitRouterAttachmentId"] = state ? state.transitRouterAttachmentId : undefined;
            resourceInputs["transitRouterAttachmentName"] = state ? state.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["vbrId"] = state ? state.vbrId : undefined;
            resourceInputs["vbrOwnerId"] = state ? state.vbrOwnerId : undefined;
        } else {
            const args = argsOrState as TransitRouterVbrAttachmentArgs | undefined;
            if ((!args || args.cenId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.vbrId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vbrId'");
            }
            resourceInputs["autoPublishRouteEnabled"] = args ? args.autoPublishRouteEnabled : undefined;
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["routeTableAssociationEnabled"] = args ? args.routeTableAssociationEnabled : undefined;
            resourceInputs["routeTablePropagationEnabled"] = args ? args.routeTablePropagationEnabled : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitRouterAttachmentDescription"] = args ? args.transitRouterAttachmentDescription : undefined;
            resourceInputs["transitRouterAttachmentName"] = args ? args.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = args ? args.transitRouterId : undefined;
            resourceInputs["vbrId"] = args ? args.vbrId : undefined;
            resourceInputs["vbrOwnerId"] = args ? args.vbrOwnerId : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["transitRouterAttachmentId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitRouterVbrAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitRouterVbrAttachment resources.
 */
export interface TransitRouterVbrAttachmentState {
    /**
     * Auto publish route enabled.Default value is `false`.
     */
    autoPublishRouteEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the CEN.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The resource type of the transit router vbr attachment.  Valid values: `VPC`, `CCN`, `VBR`, `TR`.
     *
     * ->**NOTE:** Ensure that the vbr is not used in Express Connect.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Whether to enabled route table association. The system default value is `true`.
     */
    routeTableAssociationEnabled?: pulumi.Input<boolean>;
    /**
     * Whether to enabled route table propagation. The system default value is `true`.
     */
    routeTablePropagationEnabled?: pulumi.Input<boolean>;
    /**
     * The associating status of the network.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The description of the transit router vbr attachment.
     */
    transitRouterAttachmentDescription?: pulumi.Input<string>;
    /**
     * The id of the transit router vbr attachment.
     */
    transitRouterAttachmentId?: pulumi.Input<string>;
    /**
     * The name of the transit router vbr attachment.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The ID of the transit router.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The ID of the VBR.
     */
    vbrId?: pulumi.Input<string>;
    /**
     * The owner id of the transit router vbr attachment.
     */
    vbrOwnerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransitRouterVbrAttachment resource.
 */
export interface TransitRouterVbrAttachmentArgs {
    /**
     * Auto publish route enabled.Default value is `false`.
     */
    autoPublishRouteEnabled?: pulumi.Input<boolean>;
    /**
     * The ID of the CEN.
     */
    cenId: pulumi.Input<string>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The resource type of the transit router vbr attachment.  Valid values: `VPC`, `CCN`, `VBR`, `TR`.
     *
     * ->**NOTE:** Ensure that the vbr is not used in Express Connect.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Whether to enabled route table association. The system default value is `true`.
     */
    routeTableAssociationEnabled?: pulumi.Input<boolean>;
    /**
     * Whether to enabled route table propagation. The system default value is `true`.
     */
    routeTablePropagationEnabled?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The description of the transit router vbr attachment.
     */
    transitRouterAttachmentDescription?: pulumi.Input<string>;
    /**
     * The name of the transit router vbr attachment.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The ID of the transit router.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The ID of the VBR.
     */
    vbrId: pulumi.Input<string>;
    /**
     * The owner id of the transit router vbr attachment.
     */
    vbrOwnerId?: pulumi.Input<string>;
}
