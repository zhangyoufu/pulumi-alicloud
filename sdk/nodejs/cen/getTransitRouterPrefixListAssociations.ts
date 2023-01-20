// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cen Transit Router Prefix List Associations of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.188.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cen.getTransitRouterPrefixListAssociations({
 *     transitRouterId: "tr-6ehx7q2jze8ch5ji0****",
 *     transitRouterTableId: "vtb-6ehgc262hr170qgyc****",
 * });
 * export const cenTransitRouterPrefixListAssociationId = _default.then(_default => _default.associations?.[0]?.id);
 * ```
 */
export function getTransitRouterPrefixListAssociations(args: GetTransitRouterPrefixListAssociationsArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRouterPrefixListAssociationsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouterPrefixListAssociations:getTransitRouterPrefixListAssociations", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "ownerUid": args.ownerUid,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "prefixListId": args.prefixListId,
        "status": args.status,
        "transitRouterId": args.transitRouterId,
        "transitRouterTableId": args.transitRouterTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouterPrefixListAssociations.
 */
export interface GetTransitRouterPrefixListAssociationsArgs {
    /**
     * A list of Cen Transit Router Prefix List Association IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The ID of the Alibaba Cloud account to which the prefix list belongs.
     */
    ownerUid?: number;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The ID of the prefix list.
     */
    prefixListId?: string;
    /**
     * The status of the prefix list.
     */
    status?: string;
    /**
     * The ID of the transit router.
     */
    transitRouterId: string;
    /**
     * The ID of the route table of the transit router.
     */
    transitRouterTableId: string;
}

/**
 * A collection of values returned by getTransitRouterPrefixListAssociations.
 */
export interface GetTransitRouterPrefixListAssociationsResult {
    readonly associations: outputs.cen.GetTransitRouterPrefixListAssociationsAssociation[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly ownerUid?: number;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    readonly prefixListId?: string;
    readonly status?: string;
    readonly transitRouterId: string;
    readonly transitRouterTableId: string;
}
/**
 * This data source provides the Cen Transit Router Prefix List Associations of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.188.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const default = alicloud.cen.getTransitRouterPrefixListAssociations({
 *     transitRouterId: "tr-6ehx7q2jze8ch5ji0****",
 *     transitRouterTableId: "vtb-6ehgc262hr170qgyc****",
 * });
 * export const cenTransitRouterPrefixListAssociationId = _default.then(_default => _default.associations?.[0]?.id);
 * ```
 */
export function getTransitRouterPrefixListAssociationsOutput(args: GetTransitRouterPrefixListAssociationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRouterPrefixListAssociationsResult> {
    return pulumi.output(args).apply((a: any) => getTransitRouterPrefixListAssociations(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouterPrefixListAssociations.
 */
export interface GetTransitRouterPrefixListAssociationsOutputArgs {
    /**
     * A list of Cen Transit Router Prefix List Association IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of the Alibaba Cloud account to which the prefix list belongs.
     */
    ownerUid?: pulumi.Input<number>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The ID of the prefix list.
     */
    prefixListId?: pulumi.Input<string>;
    /**
     * The status of the prefix list.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The ID of the route table of the transit router.
     */
    transitRouterTableId: pulumi.Input<string>;
}
