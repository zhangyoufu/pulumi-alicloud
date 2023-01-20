// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.188.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.message.getServiceSubscriptions({
 *     ids: ["example_id"],
 *     topicName: "tf-example",
 * });
 * export const subscriptionId1 = ids.then(ids => ids.subscriptions?.[0]?.id);
 * const name = alicloud.message.getServiceSubscriptions({
 *     topicName: "tf-example",
 * });
 * export const subscriptionId2 = name.then(name => name.subscriptions?.[0]?.id);
 * ```
 */
export function getServiceSubscriptions(args: GetServiceSubscriptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceSubscriptionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:message/getServiceSubscriptions:getServiceSubscriptions", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "subscriptionName": args.subscriptionName,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceSubscriptions.
 */
export interface GetServiceSubscriptionsArgs {
    /**
     * A list of Subscription IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Subscription name.
     */
    nameRegex?: string;
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * The name of the subscription.
     */
    subscriptionName?: string;
    /**
     * The name of the topic.
     */
    topicName: string;
}

/**
 * A collection of values returned by getServiceSubscriptions.
 */
export interface GetServiceSubscriptionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Subscription names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    /**
     * The name of the subscription.
     */
    readonly subscriptionName?: string;
    /**
     * A list of Subscriptions. Each element contains the following attributes:
     */
    readonly subscriptions: outputs.message.GetServiceSubscriptionsSubscription[];
    /**
     * The name of the topic.
     */
    readonly topicName: string;
}
/**
 * This data source provides the Message Notification Service Subscriptions of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.188.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.message.getServiceSubscriptions({
 *     ids: ["example_id"],
 *     topicName: "tf-example",
 * });
 * export const subscriptionId1 = ids.then(ids => ids.subscriptions?.[0]?.id);
 * const name = alicloud.message.getServiceSubscriptions({
 *     topicName: "tf-example",
 * });
 * export const subscriptionId2 = name.then(name => name.subscriptions?.[0]?.id);
 * ```
 */
export function getServiceSubscriptionsOutput(args: GetServiceSubscriptionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceSubscriptionsResult> {
    return pulumi.output(args).apply((a: any) => getServiceSubscriptions(a, opts))
}

/**
 * A collection of arguments for invoking getServiceSubscriptions.
 */
export interface GetServiceSubscriptionsOutputArgs {
    /**
     * A list of Subscription IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Subscription name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * The name of the subscription.
     */
    subscriptionName?: pulumi.Input<string>;
    /**
     * The name of the topic.
     */
    topicName: pulumi.Input<string>;
}
