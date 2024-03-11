// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Bandwidth Limits available to the user.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const bwl = alicloud.cen.getBandwidthLimits({
 *     instanceIds: ["cen-id1"],
 * });
 * export const firstCenBandwidthLimitsLocalRegionId = bwl.then(bwl => bwl.limits?.[0]?.localRegionId);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBandwidthLimits(args?: GetBandwidthLimitsArgs, opts?: pulumi.InvokeOptions): Promise<GetBandwidthLimitsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cen/getBandwidthLimits:getBandwidthLimits", {
        "instanceIds": args.instanceIds,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBandwidthLimits.
 */
export interface GetBandwidthLimitsArgs {
    /**
     * A list of CEN instances IDs.
     */
    instanceIds?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getBandwidthLimits.
 */
export interface GetBandwidthLimitsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    /**
     * A list of CEN Bandwidth Limits. Each element contains the following attributes:
     */
    readonly limits: outputs.cen.GetBandwidthLimitsLimit[];
    readonly outputFile?: string;
}
/**
 * This data source provides CEN Bandwidth Limits available to the user.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const bwl = alicloud.cen.getBandwidthLimits({
 *     instanceIds: ["cen-id1"],
 * });
 * export const firstCenBandwidthLimitsLocalRegionId = bwl.then(bwl => bwl.limits?.[0]?.localRegionId);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBandwidthLimitsOutput(args?: GetBandwidthLimitsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBandwidthLimitsResult> {
    return pulumi.output(args).apply((a: any) => getBandwidthLimits(a, opts))
}

/**
 * A collection of arguments for invoking getBandwidthLimits.
 */
export interface GetBandwidthLimitsOutputArgs {
    /**
     * A list of CEN instances IDs.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
