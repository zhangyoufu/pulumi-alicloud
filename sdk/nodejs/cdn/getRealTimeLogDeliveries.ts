// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.134.0+.
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
 * const example = alicloud.cdn.getRealTimeLogDeliveries({
 *     domain: "example_value",
 * });
 * export const cdnRealTimeLogDelivery1 = example.then(example => example.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRealTimeLogDeliveries(args: GetRealTimeLogDeliveriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRealTimeLogDeliveriesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", {
        "domain": args.domain,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getRealTimeLogDeliveries.
 */
export interface GetRealTimeLogDeliveriesArgs {
    /**
     * Real-Time Log Service Domain.
     */
    domain: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
     */
    status?: string;
}

/**
 * A collection of values returned by getRealTimeLogDeliveries.
 */
export interface GetRealTimeLogDeliveriesResult {
    readonly deliveries: outputs.cdn.GetRealTimeLogDeliveriesDelivery[];
    readonly domain: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.134.0+.
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
 * const example = alicloud.cdn.getRealTimeLogDeliveries({
 *     domain: "example_value",
 * });
 * export const cdnRealTimeLogDelivery1 = example.then(example => example.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRealTimeLogDeliveriesOutput(args: GetRealTimeLogDeliveriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRealTimeLogDeliveriesResult> {
    return pulumi.output(args).apply((a: any) => getRealTimeLogDeliveries(a, opts))
}

/**
 * A collection of arguments for invoking getRealTimeLogDeliveries.
 */
export interface GetRealTimeLogDeliveriesOutputArgs {
    /**
     * Real-Time Log Service Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
     */
    status?: pulumi.Input<string>;
}
