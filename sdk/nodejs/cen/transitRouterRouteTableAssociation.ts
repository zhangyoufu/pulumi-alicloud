// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN transit router route table association resource.[What is Cen Transit Router Route Table Association](https://help.aliyun.com/document_detail/261242.html)
 *
 * > **NOTE:** Available in 1.126.0+
 *
 * ## Import
 *
 * CEN transit router route table association can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation default tr-********:tr-attach-********
 * ```
 */
export class TransitRouterRouteTableAssociation extends pulumi.CustomResource {
    /**
     * Get an existing TransitRouterRouteTableAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitRouterRouteTableAssociationState, opts?: pulumi.CustomResourceOptions): TransitRouterRouteTableAssociation {
        return new TransitRouterRouteTableAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/transitRouterRouteTableAssociation:TransitRouterRouteTableAssociation';

    /**
     * Returns true if the given object is an instance of TransitRouterRouteTableAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitRouterRouteTableAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitRouterRouteTableAssociation.__pulumiType;
    }

    /**
     * The dry run.
     *
     * > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * The associating status of the network.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID the transit router attachment.
     */
    public readonly transitRouterAttachmentId!: pulumi.Output<string>;
    /**
     * The ID of the transit router route table.
     */
    public readonly transitRouterRouteTableId!: pulumi.Output<string>;

    /**
     * Create a TransitRouterRouteTableAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitRouterRouteTableAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitRouterRouteTableAssociationArgs | TransitRouterRouteTableAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitRouterRouteTableAssociationState | undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transitRouterAttachmentId"] = state ? state.transitRouterAttachmentId : undefined;
            resourceInputs["transitRouterRouteTableId"] = state ? state.transitRouterRouteTableId : undefined;
        } else {
            const args = argsOrState as TransitRouterRouteTableAssociationArgs | undefined;
            if ((!args || args.transitRouterAttachmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitRouterAttachmentId'");
            }
            if ((!args || args.transitRouterRouteTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitRouterRouteTableId'");
            }
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["transitRouterAttachmentId"] = args ? args.transitRouterAttachmentId : undefined;
            resourceInputs["transitRouterRouteTableId"] = args ? args.transitRouterRouteTableId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitRouterRouteTableAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitRouterRouteTableAssociation resources.
 */
export interface TransitRouterRouteTableAssociationState {
    /**
     * The dry run.
     *
     * > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The associating status of the network.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID the transit router attachment.
     */
    transitRouterAttachmentId?: pulumi.Input<string>;
    /**
     * The ID of the transit router route table.
     */
    transitRouterRouteTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransitRouterRouteTableAssociation resource.
 */
export interface TransitRouterRouteTableAssociationArgs {
    /**
     * The dry run.
     *
     * > **NOTE:** The Zone of CEN has MasterZone and SlaveZone, first zoneId of zoneMapping need be MasterZone. We have a API to describeZones[API](https://help.aliyun.com/document_detail/261356.html)
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * The ID the transit router attachment.
     */
    transitRouterAttachmentId: pulumi.Input<string>;
    /**
     * The ID of the transit router route table.
     */
    transitRouterRouteTableId: pulumi.Input<string>;
}
