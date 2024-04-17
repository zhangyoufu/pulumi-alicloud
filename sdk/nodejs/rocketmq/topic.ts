// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ONS topic resource.
 *
 * For more information about how to use it, see [RocketMQ Topic Management API](https://www.alibabacloud.com/help/doc-detail/29591.html).
 *
 * > **NOTE:** Available in 1.53.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "onsInstanceName";
 * const topic = config.get("topic") || "onsTopicName";
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultInstance = new alicloud.rocketmq.Instance("default", {
 *     name: `${name}-${_default.result}`,
 *     remark: "default_ons_instance_remark",
 * });
 * const defaultTopic = new alicloud.rocketmq.Topic("default", {
 *     topicName: topic,
 *     instanceId: defaultInstance.id,
 *     messageType: 0,
 *     remark: "dafault_ons_topic_remark",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ONS TOPIC can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:rocketmq/topic:Topic topic MQ_INST_1234567890_Baso1234567:onsTopicDemo
 * ```
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rocketmq/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * ID of the ONS Instance that owns the topics.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
     */
    public readonly messageType!: pulumi.Output<number>;
    /**
     * This attribute has been deprecated.
     *
     * @deprecated Attribute perm has been deprecated and suggest removing it from your template.
     */
    public readonly perm!: pulumi.Output<number>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 128.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     *
     * > **NOTE:** At least one of `topicName` and `topic` should be set.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Replaced by `topicName` after version 1.97.0.
     *
     * @deprecated Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
     */
    public readonly topic!: pulumi.Output<string>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["messageType"] = state ? state.messageType : undefined;
            resourceInputs["perm"] = state ? state.perm : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.messageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'messageType'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["messageType"] = args ? args.messageType : undefined;
            resourceInputs["perm"] = args ? args.perm : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * ID of the ONS Instance that owns the topics.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
     */
    messageType?: pulumi.Input<number>;
    /**
     * This attribute has been deprecated.
     *
     * @deprecated Attribute perm has been deprecated and suggest removing it from your template.
     */
    perm?: pulumi.Input<number>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 128.
     */
    remark?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     *
     * > **NOTE:** At least one of `topicName` and `topic` should be set.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Replaced by `topicName` after version 1.97.0.
     *
     * @deprecated Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
     */
    topic?: pulumi.Input<string>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * ID of the ONS Instance that owns the topics.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
     */
    messageType: pulumi.Input<number>;
    /**
     * This attribute has been deprecated.
     *
     * @deprecated Attribute perm has been deprecated and suggest removing it from your template.
     */
    perm?: pulumi.Input<number>;
    /**
     * This attribute is a concise description of topic. The length cannot exceed 128.
     */
    remark?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     *
     * > **NOTE:** At least one of `topicName` and `topic` should be set.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Replaced by `topicName` after version 1.97.0.
     *
     * @deprecated Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
     */
    topic?: pulumi.Input<string>;
    /**
     * Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
     */
    topicName?: pulumi.Input<string>;
}
