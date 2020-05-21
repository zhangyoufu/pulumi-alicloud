// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const queues = pulumi.output(alicloud.mns.getQueues({
 *     namePrefix: "tf-",
 * }, { async: true }));
 *
 * export const firstQueueId = queues.queues[0].id;
 * ```
 */
export function getQueues(args?: GetQueuesArgs, opts?: pulumi.InvokeOptions): Promise<GetQueuesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:mns/getQueues:getQueues", {
        "namePrefix": args.namePrefix,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getQueues.
 */
export interface GetQueuesArgs {
    /**
     * A string to filter resulting queues by their name prefixs.
     */
    readonly namePrefix?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getQueues.
 */
export interface GetQueuesResult {
    readonly namePrefix?: string;
    /**
     * A list of queue names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of queues. Each element contains the following attributes:
     */
    readonly queues: outputs.mns.GetQueuesQueue[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
