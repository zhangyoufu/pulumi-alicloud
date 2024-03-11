// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
 *
 * > **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const queues = alicloud.mns.getQueues({
 *     namePrefix: "tf-",
 * });
 * export const firstQueueId = queues.then(queues => queues.queues?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQueues(args?: GetQueuesArgs, opts?: pulumi.InvokeOptions): Promise<GetQueuesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
    namePrefix?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getQueues.
 */
export interface GetQueuesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
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
}
/**
 * This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
 *
 * > **DEPRECATED:**  This datasource has been deprecated from version `1.188.0`. Please use new datasource message_service_queues.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const queues = alicloud.mns.getQueues({
 *     namePrefix: "tf-",
 * });
 * export const firstQueueId = queues.then(queues => queues.queues?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQueuesOutput(args?: GetQueuesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueuesResult> {
    return pulumi.output(args).apply((a: any) => getQueues(a, opts))
}

/**
 * A collection of arguments for invoking getQueues.
 */
export interface GetQueuesOutputArgs {
    /**
     * A string to filter resulting queues by their name prefixs.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
