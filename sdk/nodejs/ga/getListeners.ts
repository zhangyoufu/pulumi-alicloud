// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Global Accelerator (GA) Listeners of the current Alibaba Cloud user.
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
 * const example = alicloud.ga.getListeners({
 *     acceleratorId: "example_value",
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstGaListenerId = example.then(example => example.listeners?.[0]?.id);
 * ```
 */
export function getListeners(args: GetListenersArgs, opts?: pulumi.InvokeOptions): Promise<GetListenersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ga/getListeners:getListeners", {
        "acceleratorId": args.acceleratorId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getListeners.
 */
export interface GetListenersArgs {
    /**
     * The accelerator id.
     */
    acceleratorId: string;
    /**
     * A list of Listener IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Listener name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The status of the listener.
     */
    status?: string;
}

/**
 * A collection of values returned by getListeners.
 */
export interface GetListenersResult {
    readonly acceleratorId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly listeners: outputs.ga.GetListenersListener[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly status?: string;
}
/**
 * This data source provides the Global Accelerator (GA) Listeners of the current Alibaba Cloud user.
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
 * const example = alicloud.ga.getListeners({
 *     acceleratorId: "example_value",
 *     ids: ["example_value"],
 *     nameRegex: "the_resource_name",
 * });
 * export const firstGaListenerId = example.then(example => example.listeners?.[0]?.id);
 * ```
 */
export function getListenersOutput(args: GetListenersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetListenersResult> {
    return pulumi.output(args).apply((a: any) => getListeners(a, opts))
}

/**
 * A collection of arguments for invoking getListeners.
 */
export interface GetListenersOutputArgs {
    /**
     * The accelerator id.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * A list of Listener IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Listener name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the listener.
     */
    status?: pulumi.Input<string>;
}
