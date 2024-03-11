// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **NOTE:** From the version 1.132.0, the data source has been renamed to `alicloud.bastionhost.getInstances`.
 *
 * This data source provides a list of cloud Bastionhost instances in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.63.0+ .
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instanceInstances = alicloud.bastionhost.getInstances({
 *     descriptionRegex: "^bastionhost",
 * });
 * export const instance = alicloud_bastionhost_instances.instance.map(__item => __item.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:bastionhost/getInstances:getInstances", {
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * A regex string to filter results by the instance description.
     */
    descriptionRegex?: string;
    /**
     * Matched instance IDs to filter data source result.
     */
    ids?: string[];
    /**
     * File name to persist data source output.
     */
    outputFile?: string;
    /**
     * A map of tags assigned to the bastionhost instance. It must be in the format:
     * <!--Start PulumiCodeChooser -->
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as alicloud from "@pulumi/alicloud";
     *
     * const instance = alicloud.bastionhost.getInstances({
     *     tags: {
     *         tagKey1: "tagValue1",
     *     },
     * });
     * ```
     * <!--End PulumiCodeChooser -->
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly descriptionRegex?: string;
    readonly descriptions: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * A list of apis. Each element contains the following attributes:
     */
    readonly instances: outputs.bastionhost.GetInstancesInstance[];
    readonly outputFile?: string;
    /**
     * A map of tags assigned to the bastionhost instance.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * > **NOTE:** From the version 1.132.0, the data source has been renamed to `alicloud.bastionhost.getInstances`.
 *
 * This data source provides a list of cloud Bastionhost instances in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.63.0+ .
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instanceInstances = alicloud.bastionhost.getInstances({
 *     descriptionRegex: "^bastionhost",
 * });
 * export const instance = alicloud_bastionhost_instances.instance.map(__item => __item.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * A regex string to filter results by the instance description.
     */
    descriptionRegex?: pulumi.Input<string>;
    /**
     * Matched instance IDs to filter data source result.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name to persist data source output.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A map of tags assigned to the bastionhost instance. It must be in the format:
     * <!--Start PulumiCodeChooser -->
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as alicloud from "@pulumi/alicloud";
     *
     * const instance = alicloud.bastionhost.getInstances({
     *     tags: {
     *         tagKey1: "tagValue1",
     *     },
     * });
     * ```
     * <!--End PulumiCodeChooser -->
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
