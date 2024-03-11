// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Config Aggregate Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.172.0+.
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
 * const ids = alicloud.cfg.getAggregateDeliveries({
 *     aggregatorId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const configAggregateDeliveryId1 = ids.then(ids => ids.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAggregateDeliveries(args: GetAggregateDeliveriesArgs, opts?: pulumi.InvokeOptions): Promise<GetAggregateDeliveriesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cfg/getAggregateDeliveries:getAggregateDeliveries", {
        "aggregatorId": args.aggregatorId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAggregateDeliveries.
 */
export interface GetAggregateDeliveriesArgs {
    /**
     * The ID of the Aggregator.
     */
    aggregatorId: string;
    /**
     * A list of Aggregate Delivery IDs.
     */
    ids?: string[];
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
     */
    status?: number;
}

/**
 * A collection of values returned by getAggregateDeliveries.
 */
export interface GetAggregateDeliveriesResult {
    readonly aggregatorId: string;
    readonly deliveries: outputs.cfg.GetAggregateDeliveriesDelivery[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: number;
}
/**
 * This data source provides the Config Aggregate Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.172.0+.
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
 * const ids = alicloud.cfg.getAggregateDeliveries({
 *     aggregatorId: "example_value",
 *     ids: [
 *         "example_value-1",
 *         "example_value-2",
 *     ],
 * });
 * export const configAggregateDeliveryId1 = ids.then(ids => ids.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAggregateDeliveriesOutput(args: GetAggregateDeliveriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAggregateDeliveriesResult> {
    return pulumi.output(args).apply((a: any) => getAggregateDeliveries(a, opts))
}

/**
 * A collection of arguments for invoking getAggregateDeliveries.
 */
export interface GetAggregateDeliveriesOutputArgs {
    /**
     * The ID of the Aggregator.
     */
    aggregatorId: pulumi.Input<string>;
    /**
     * A list of Aggregate Delivery IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled.
     */
    status?: pulumi.Input<number>;
}
