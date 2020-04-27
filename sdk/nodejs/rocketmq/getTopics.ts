// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ONS Topics in an Alibaba Cloud account according to the specified filters.
 * 
 * > **NOTE:** Available in 1.53.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "onsInstanceName";
 * const topic = config.get("topic") || "onsTopicDatasourceName";
 * 
 * const defaultInstance = new alicloud.rocketmq.Instance("default", {
 *     remark: "defaultOnsInstanceRemark",
 * });
 * const defaultTopic = new alicloud.rocketmq.Topic("default", {
 *     instanceId: defaultInstance.id,
 *     messageType: 0,
 *     remark: "dafaultOnsTopicRemark",
 *     topic: topic,
 * });
 * const topicsDs = defaultTopic.instanceId.apply(instanceId => alicloud.rocketmq.getTopics({
 *     instanceId: instanceId,
 *     nameRegex: topic,
 *     outputFile: "topics.txt",
 * }, { async: true }));
 * 
 * export const firstTopicName = topicsDs.topics[0].topic;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ons_topics.html.markdown.
 */
export function getTopics(args: GetTopicsArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:rocketmq/getTopics:getTopics", {
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsArgs {
    /**
     * ID of the ONS Instance that owns the topics.
     */
    readonly instanceId: string;
    /**
     * A regex string to filter results by the topic name. 
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getTopics.
 */
export interface GetTopicsResult {
    readonly instanceId: string;
    readonly nameRegex?: string;
    /**
     * A list of topic names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of topics. Each element contains the following attributes:
     */
    readonly topics: outputs.rocketmq.GetTopicsTopic[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
