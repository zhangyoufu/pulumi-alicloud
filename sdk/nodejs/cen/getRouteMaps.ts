// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Route Maps available to the user.
 *
 * > **NOTE:** Available in v1.87.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const this = alicloud.cen.getRouteMaps({
 *     cenId: "cen-ihdlgo87ai********",
 *     ids: ["cen-ihdlgo87ai:cenrmap-bnh97kb3mn********"],
 *     descriptionRegex: "datasource_test",
 *     cenRegionId: "cn-hangzhou",
 *     transmitDirection: "RegionIn",
 *     status: "Active",
 * });
 * export const firstCenRouteMapId = _this.then(_this => _this.maps?.[0]?.routeMapId);
 * ```
 */
export function getRouteMaps(args: GetRouteMapsArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteMapsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cen/getRouteMaps:getRouteMaps", {
        "cenId": args.cenId,
        "cenRegionId": args.cenRegionId,
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "transmitDirection": args.transmitDirection,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteMaps.
 */
export interface GetRouteMapsArgs {
    /**
     * The ID of the CEN instance.
     */
    cenId: string;
    /**
     * The ID of the region to which the CEN instance belongs.
     */
    cenRegionId?: string;
    /**
     * A regex string to filter CEN route map by description.
     */
    descriptionRegex?: string;
    /**
     * A list of CEN route map IDs. Each item formats as `<cen_id>:<route_map_id>`.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of the route map, including `Creating`, `Active` and `Deleting`.
     */
    status?: string;
    /**
     * The direction in which the route map is applied, including `RegionIn` and `RegionOut`.
     */
    transmitDirection?: string;
}

/**
 * A collection of values returned by getRouteMaps.
 */
export interface GetRouteMapsResult {
    /**
     * The ID of the CEN instance.
     */
    readonly cenId: string;
    /**
     * The ID of the region to which the CEN instance belongs.
     */
    readonly cenRegionId?: string;
    readonly descriptionRegex?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of CEN route map IDs. Each item formats as `<cen_id>:<route_map_id>`. Before 1.161.0, its element is `routeMapId`.
     */
    readonly ids: string[];
    /**
     * A list of CEN instances. Each element contains the following attributes:
     */
    readonly maps: outputs.cen.GetRouteMapsMap[];
    readonly outputFile?: string;
    /**
     * The status of the route map.
     */
    readonly status?: string;
    /**
     * The direction in which the route map is applied.
     */
    readonly transmitDirection?: string;
}
/**
 * This data source provides CEN Route Maps available to the user.
 *
 * > **NOTE:** Available in v1.87.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const this = alicloud.cen.getRouteMaps({
 *     cenId: "cen-ihdlgo87ai********",
 *     ids: ["cen-ihdlgo87ai:cenrmap-bnh97kb3mn********"],
 *     descriptionRegex: "datasource_test",
 *     cenRegionId: "cn-hangzhou",
 *     transmitDirection: "RegionIn",
 *     status: "Active",
 * });
 * export const firstCenRouteMapId = _this.then(_this => _this.maps?.[0]?.routeMapId);
 * ```
 */
export function getRouteMapsOutput(args: GetRouteMapsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteMapsResult> {
    return pulumi.output(args).apply((a: any) => getRouteMaps(a, opts))
}

/**
 * A collection of arguments for invoking getRouteMaps.
 */
export interface GetRouteMapsOutputArgs {
    /**
     * The ID of the CEN instance.
     */
    cenId: pulumi.Input<string>;
    /**
     * The ID of the region to which the CEN instance belongs.
     */
    cenRegionId?: pulumi.Input<string>;
    /**
     * A regex string to filter CEN route map by description.
     */
    descriptionRegex?: pulumi.Input<string>;
    /**
     * A list of CEN route map IDs. Each item formats as `<cen_id>:<route_map_id>`.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the route map, including `Creating`, `Active` and `Deleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * The direction in which the route map is applied, including `RegionIn` and `RegionOut`.
     */
    transmitDirection?: pulumi.Input<string>;
}
