// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Actiontrail History Delivery Jobs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.139.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.actiontrail.getHistoryDeliveryJobs({
 *     ids: ["example_id"],
 * });
 * export const actiontrailHistoryDeliveryJobId1 = ids.then(ids => ids.jobs?.[0]?.id);
 * const status = alicloud.actiontrail.getHistoryDeliveryJobs({
 *     ids: ["example_id"],
 *     status: 2,
 * });
 * export const actiontrailHistoryDeliveryJobId2 = status.then(status => status.jobs?.[0]?.id);
 * ```
 */
export function getHistoryDeliveryJobs(args?: GetHistoryDeliveryJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetHistoryDeliveryJobsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:actiontrail/getHistoryDeliveryJobs:getHistoryDeliveryJobs", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getHistoryDeliveryJobs.
 */
export interface GetHistoryDeliveryJobsArgs {
    enableDetails?: boolean;
    /**
     * A list of History Delivery Job IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
     */
    status?: number;
}

/**
 * A collection of values returned by getHistoryDeliveryJobs.
 */
export interface GetHistoryDeliveryJobsResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly jobs: outputs.actiontrail.GetHistoryDeliveryJobsJob[];
    readonly outputFile?: string;
    readonly status?: number;
}
/**
 * This data source provides the Actiontrail History Delivery Jobs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.139.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.actiontrail.getHistoryDeliveryJobs({
 *     ids: ["example_id"],
 * });
 * export const actiontrailHistoryDeliveryJobId1 = ids.then(ids => ids.jobs?.[0]?.id);
 * const status = alicloud.actiontrail.getHistoryDeliveryJobs({
 *     ids: ["example_id"],
 *     status: 2,
 * });
 * export const actiontrailHistoryDeliveryJobId2 = status.then(status => status.jobs?.[0]?.id);
 * ```
 */
export function getHistoryDeliveryJobsOutput(args?: GetHistoryDeliveryJobsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHistoryDeliveryJobsResult> {
    return pulumi.output(args).apply((a: any) => getHistoryDeliveryJobs(a, opts))
}

/**
 * A collection of arguments for invoking getHistoryDeliveryJobs.
 */
export interface GetHistoryDeliveryJobsOutputArgs {
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of History Delivery Job IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
     */
    status?: pulumi.Input<number>;
}
