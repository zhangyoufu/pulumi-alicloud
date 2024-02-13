// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Message Notification Service Topic resource.
 *
 * For information about Message Notification Service Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/message-service/latest/createtopic).
 *
 * > **NOTE:** Available since v1.188.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-example";
 * const _default = new alicloud.message.ServiceTopic("default", {
 *     topicName: name,
 *     maxMessageSize: 12357,
 *     loggingEnabled: true,
 * });
 * ```
 *
 * ## Import
 *
 * Message Notification Service Topic can be imported using the id or topic_name, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:message/serviceTopic:ServiceTopic example <topic_name>
 * ```
 */
export class ServiceTopic extends pulumi.CustomResource {
    /**
     * Get an existing ServiceTopic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceTopicState, opts?: pulumi.CustomResourceOptions): ServiceTopic {
        return new ServiceTopic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:message/serviceTopic:ServiceTopic';

    /**
     * Returns true if the given object is an instance of ServiceTopic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceTopic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceTopic.__pulumiType;
    }

    /**
     * Specifies whether to enable the log management feature. Default value: false. Valid values:
     */
    public readonly loggingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
     */
    public readonly maxMessageSize!: pulumi.Output<number>;
    /**
     * Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a ServiceTopic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceTopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceTopicArgs | ServiceTopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceTopicState | undefined;
            resourceInputs["loggingEnabled"] = state ? state.loggingEnabled : undefined;
            resourceInputs["maxMessageSize"] = state ? state.maxMessageSize : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as ServiceTopicArgs | undefined;
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["loggingEnabled"] = args ? args.loggingEnabled : undefined;
            resourceInputs["maxMessageSize"] = args ? args.maxMessageSize : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceTopic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceTopic resources.
 */
export interface ServiceTopicState {
    /**
     * Specifies whether to enable the log management feature. Default value: false. Valid values:
     */
    loggingEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
     */
    maxMessageSize?: pulumi.Input<number>;
    /**
     * Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceTopic resource.
 */
export interface ServiceTopicArgs {
    /**
     * Specifies whether to enable the log management feature. Default value: false. Valid values:
     */
    loggingEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
     */
    maxMessageSize?: pulumi.Input<number>;
    /**
     * Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
     */
    topicName: pulumi.Input<string>;
}
