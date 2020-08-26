// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of MNS topic subscriptions in an Alibaba Cloud account according to the specified parameters.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const subscriptions = pulumi.output(alicloud.mns.getTopicSubscriptions({
 *     namePrefix: "tf-",
 *     topicName: "topic_name",
 * }, { async: true }));
 *
 * export const firstTopicSubscriptionId = subscriptions.subscriptions[0].id;
 * ```
 */
export function getTopicSubscriptions(args: GetTopicSubscriptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicSubscriptionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:mns/getTopicSubscriptions:getTopicSubscriptions", {
        "namePrefix": args.namePrefix,
        "outputFile": args.outputFile,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopicSubscriptions.
 */
export interface GetTopicSubscriptionsArgs {
    /**
     * A string to filter resulting subscriptions of the topic by their name prefixs.
     */
    readonly namePrefix?: string;
    readonly outputFile?: string;
    /**
     * Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
     */
    readonly topicName: string;
}

/**
 * A collection of values returned by getTopicSubscriptions.
 */
export interface GetTopicSubscriptionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namePrefix?: string;
    /**
     * A list of subscription names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of subscriptions. Each element contains the following attributes:
     */
    readonly subscriptions: outputs.mns.GetTopicSubscriptionsSubscription[];
    readonly topicName: string;
}
