// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cen Transit Router Multicast Domains of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.195.0+.
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
 * const ids = alicloud.cen.getTransitRouterMulticastDomains({
 *     ids: ["example_id"],
 *     transitRouterId: "your_transit_router_id",
 * });
 * export const cenTransitRouterMulticastDomainId0 = ids.then(ids => ids.domains?.[0]?.id);
 * const nameRegex = alicloud.cen.getTransitRouterMulticastDomains({
 *     nameRegex: "^my-name",
 *     transitRouterId: "your_transit_router_id",
 * });
 * export const cenTransitRouterMulticastDomainId1 = nameRegex.then(nameRegex => nameRegex.domains?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTransitRouterMulticastDomains(args: GetTransitRouterMulticastDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitRouterMulticastDomainsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cen/getTransitRouterMulticastDomains:getTransitRouterMulticastDomains", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "transitRouterId": args.transitRouterId,
        "transitRouterMulticastDomainId": args.transitRouterMulticastDomainId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitRouterMulticastDomains.
 */
export interface GetTransitRouterMulticastDomainsArgs {
    /**
     * A list of Transit Router Multicast Domain IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Transit Router Multicast Domain name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the multicast domain. Valid Value: `Active`.
     */
    status?: string;
    /**
     * The ID of the transit router.
     */
    transitRouterId: string;
    /**
     * The ID of the multicast domain.
     */
    transitRouterMulticastDomainId?: string;
}

/**
 * A collection of values returned by getTransitRouterMulticastDomains.
 */
export interface GetTransitRouterMulticastDomainsResult {
    /**
     * A list of Cen Transit Router Multicast Domains. Each element contains the following attributes:
     */
    readonly domains: outputs.cen.GetTransitRouterMulticastDomainsDomain[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Transit Router Multicast Domain names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The status of the Transit Router Multicast Domain.
     */
    readonly status?: string;
    /**
     * The ID of the transit router.
     */
    readonly transitRouterId: string;
    /**
     * The ID of the Transit Router Multicast Domain.
     */
    readonly transitRouterMulticastDomainId?: string;
}
/**
 * This data source provides the Cen Transit Router Multicast Domains of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.195.0+.
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
 * const ids = alicloud.cen.getTransitRouterMulticastDomains({
 *     ids: ["example_id"],
 *     transitRouterId: "your_transit_router_id",
 * });
 * export const cenTransitRouterMulticastDomainId0 = ids.then(ids => ids.domains?.[0]?.id);
 * const nameRegex = alicloud.cen.getTransitRouterMulticastDomains({
 *     nameRegex: "^my-name",
 *     transitRouterId: "your_transit_router_id",
 * });
 * export const cenTransitRouterMulticastDomainId1 = nameRegex.then(nameRegex => nameRegex.domains?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTransitRouterMulticastDomainsOutput(args: GetTransitRouterMulticastDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitRouterMulticastDomainsResult> {
    return pulumi.output(args).apply((a: any) => getTransitRouterMulticastDomains(a, opts))
}

/**
 * A collection of arguments for invoking getTransitRouterMulticastDomains.
 */
export interface GetTransitRouterMulticastDomainsOutputArgs {
    /**
     * A list of Transit Router Multicast Domain IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Transit Router Multicast Domain name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the multicast domain. Valid Value: `Active`.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The ID of the multicast domain.
     */
    transitRouterMulticastDomainId?: pulumi.Input<string>;
}
