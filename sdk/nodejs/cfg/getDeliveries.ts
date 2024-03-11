// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Config Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.171.0+.
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
 * const ids = alicloud.cfg.getDeliveries({
 *     ids: ["example_id"],
 * });
 * export const configDeliveryId1 = ids.then(ids => ids.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDeliveries(args?: GetDeliveriesArgs, opts?: pulumi.InvokeOptions): Promise<GetDeliveriesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cfg/getDeliveries:getDeliveries", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeliveries.
 */
export interface GetDeliveriesArgs {
    /**
     * A list of Delivery IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by delivery channel name.
     */
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
 * A collection of values returned by getDeliveries.
 */
export interface GetDeliveriesResult {
    readonly deliveries: outputs.cfg.GetDeliveriesDelivery[];
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
 * This data source provides the Config Deliveries of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.171.0+.
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
 * const ids = alicloud.cfg.getDeliveries({
 *     ids: ["example_id"],
 * });
 * export const configDeliveryId1 = ids.then(ids => ids.deliveries?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDeliveriesOutput(args?: GetDeliveriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeliveriesResult> {
    return pulumi.output(args).apply((a: any) => getDeliveries(a, opts))
}

/**
 * A collection of arguments for invoking getDeliveries.
 */
export interface GetDeliveriesOutputArgs {
    /**
     * A list of Delivery IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by delivery channel name.
     */
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
