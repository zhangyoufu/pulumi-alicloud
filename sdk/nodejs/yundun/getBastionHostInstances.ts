// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of cloud Bastionhost instances in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.63.0+ .
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const instanceBastionHostInstances = pulumi.output(alicloud.yundun.getBastionHostInstances({
 *     nameRegex: "^bastionhost",
 * }, { async: true }));
 *
 * export const instance = alicloud_yundun_bastionhost_instances_instance.map(v => v.id);
 * ```
 */
export function getBastionHostInstances(args?: GetBastionHostInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetBastionHostInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:yundun/getBastionHostInstances:getBastionHostInstances", {
        "descriptionRegex": args.descriptionRegex,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getBastionHostInstances.
 */
export interface GetBastionHostInstancesArgs {
    /**
     * A regex string to filter results by the instance description.
     */
    readonly descriptionRegex?: string;
    /**
     * Matched instance IDs to filter data source result.
     */
    readonly ids?: string[];
    /**
     * File name to persist data source output.
     */
    readonly outputFile?: string;
    /**
     * A map of tags assigned to the bastionhost instance. It must be in the format:
     * ```
     * data "alicloud.yundun.getBastionHostInstances" "instance" {
     * tags = {
     * tagKey1 = "tagValue1"
     * }
     * }
     * ```
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getBastionHostInstances.
 */
export interface GetBastionHostInstancesResult {
    readonly descriptionRegex?: string;
    readonly descriptions: string[];
    readonly ids: string[];
    /**
     * A list of apis. Each element contains the following attributes:
     */
    readonly instances: outputs.yundun.GetBastionHostInstancesInstance[];
    readonly outputFile?: string;
    /**
     * A map of tags assigned to the bastionhost instance.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
