// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Global Accelerator (GA) Accelerators of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ga.getAccelerators({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstGaAcceleratorId = example.then(example => example.accelerators?.[0]?.id);
 * ```
 */
export function getAccelerators(args?: GetAcceleratorsArgs, opts?: pulumi.InvokeOptions): Promise<GetAcceleratorsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ga/getAccelerators:getAccelerators", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccelerators.
 */
export interface GetAcceleratorsArgs {
    /**
     * A list of Accelerator IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Accelerator name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the GA instance.
     */
    status?: string;
}

/**
 * A collection of values returned by getAccelerators.
 */
export interface GetAcceleratorsResult {
    readonly accelerators: outputs.ga.GetAcceleratorsAccelerator[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Global Accelerator (GA) Accelerators of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ga.getAccelerators({
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstGaAcceleratorId = example.then(example => example.accelerators?.[0]?.id);
 * ```
 */
export function getAcceleratorsOutput(args?: GetAcceleratorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAcceleratorsResult> {
    return pulumi.output(args).apply((a: any) => getAccelerators(a, opts))
}

/**
 * A collection of arguments for invoking getAccelerators.
 */
export interface GetAcceleratorsOutputArgs {
    /**
     * A list of Accelerator IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Accelerator name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the GA instance.
     */
    status?: pulumi.Input<string>;
}
