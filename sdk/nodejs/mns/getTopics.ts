// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const topics = alicloud.mns.getTopics({
 *     namePrefix: "tf-",
 * });
 * export const firstTopicId = topics.then(topics => topics.topics?.[0]?.id);
 * ```
 */
export function getTopics(args?: GetTopicsArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:mns/getTopics:getTopics", {
        "namePrefix": args.namePrefix,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsArgs {
    /**
     * A string to filter resulting topics by their name prefixs.
     */
    namePrefix?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getTopics.
 */
export interface GetTopicsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namePrefix?: string;
    /**
     * A list of topic names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of topics. Each element contains the following attributes:
     */
    readonly topics: outputs.mns.GetTopicsTopic[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const topics = alicloud.mns.getTopics({
 *     namePrefix: "tf-",
 * });
 * export const firstTopicId = topics.then(topics => topics.topics?.[0]?.id);
 * ```
 */
export function getTopicsOutput(args?: GetTopicsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicsResult> {
    return pulumi.output(args).apply((a: any) => getTopics(a, opts))
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsOutputArgs {
    /**
     * A string to filter resulting topics by their name prefixs.
     */
    namePrefix?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
