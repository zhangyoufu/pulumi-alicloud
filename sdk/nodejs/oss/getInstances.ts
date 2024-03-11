// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ots instances of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instancesDs = alicloud.ots.getInstances({
 *     nameRegex: "sample-instance",
 *     outputFile: "instances.txt",
 * });
 * export const firstInstanceId = instancesDs.then(instancesDs => instancesDs.instances?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
/** @deprecated alicloud.oss.getInstances has been deprecated in favor of alicloud.ots.getInstances */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    pulumi.log.warn("getInstances is deprecated: alicloud.oss.getInstances has been deprecated in favor of alicloud.ots.getInstances")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:oss/getInstances:getInstances", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * A list of instance IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by instance name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * A map of tags assigned to the instance. It must be in the format:
     * <!--Start PulumiCodeChooser -->
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as alicloud from "@pulumi/alicloud";
     *
     * const instancesDs = alicloud.ots.getInstances({
     *     tags: {
     *         tagKey1: "tagValue1",
     *         tagKey2: "tagValue2",
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
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of instance IDs.
     */
    readonly ids: string[];
    /**
     * A list of instances. Each element contains the following attributes:
     */
    readonly instances: outputs.oss.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * A list of instance names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The tags of the instance.
     */
    readonly tags?: {[key: string]: any};
}
/**
 * This data source provides the ots instances of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instancesDs = alicloud.ots.getInstances({
 *     nameRegex: "sample-instance",
 *     outputFile: "instances.txt",
 * });
 * export const firstInstanceId = instancesDs.then(instancesDs => instancesDs.instances?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
/** @deprecated alicloud.oss.getInstances has been deprecated in favor of alicloud.ots.getInstances */
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * A list of instance IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by instance name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A map of tags assigned to the instance. It must be in the format:
     * <!--Start PulumiCodeChooser -->
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as alicloud from "@pulumi/alicloud";
     *
     * const instancesDs = alicloud.ots.getInstances({
     *     tags: {
     *         tagKey1: "tagValue1",
     *         tagKey2: "tagValue2",
     *     },
     * });
     * ```
     * <!--End PulumiCodeChooser -->
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
