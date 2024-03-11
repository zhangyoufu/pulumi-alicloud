// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Chatbot Publish Task resource.
 *
 * For information about Chatbot Publish Task and how to use it, see [What is Publish Task](https://help.aliyun.com/document_detail/433996.html).
 *
 * > **NOTE:** Available since v1.203.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAgents = alicloud.chatbot.getAgents({});
 * const defaultPublishTask = new alicloud.chatbot.PublishTask("defaultPublishTask", {
 *     bizType: "faq",
 *     agentKey: defaultAgents.then(defaultAgents => defaultAgents.agents?.[0]?.agentKey),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Chatbot Publish Task can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:chatbot/publishTask:PublishTask example <id>
 * ```
 */
export class PublishTask extends pulumi.CustomResource {
    /**
     * Get an existing PublishTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublishTaskState, opts?: pulumi.CustomResourceOptions): PublishTask {
        return new PublishTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:chatbot/publishTask:PublishTask';

    /**
     * Returns true if the given object is an instance of PublishTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublishTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublishTask.__pulumiType;
    }

    /**
     * The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
     */
    public readonly agentKey!: pulumi.Output<string>;
    /**
     * The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
     */
    public readonly bizType!: pulumi.Output<string>;
    /**
     * UTC time of task creation
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
     */
    public readonly dataIdLists!: pulumi.Output<string[] | undefined>;
    /**
     * UTC time for task modification
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * The status of the task.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a PublishTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublishTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublishTaskArgs | PublishTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublishTaskState | undefined;
            resourceInputs["agentKey"] = state ? state.agentKey : undefined;
            resourceInputs["bizType"] = state ? state.bizType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dataIdLists"] = state ? state.dataIdLists : undefined;
            resourceInputs["modifyTime"] = state ? state.modifyTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as PublishTaskArgs | undefined;
            if ((!args || args.bizType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bizType'");
            }
            resourceInputs["agentKey"] = args ? args.agentKey : undefined;
            resourceInputs["bizType"] = args ? args.bizType : undefined;
            resourceInputs["dataIdLists"] = args ? args.dataIdLists : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["modifyTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublishTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublishTask resources.
 */
export interface PublishTaskState {
    /**
     * The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
     */
    agentKey?: pulumi.Input<string>;
    /**
     * The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
     */
    bizType?: pulumi.Input<string>;
    /**
     * UTC time of task creation
     */
    createTime?: pulumi.Input<string>;
    /**
     * Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
     */
    dataIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * UTC time for task modification
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * The status of the task.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublishTask resource.
 */
export interface PublishTaskArgs {
    /**
     * The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
     */
    agentKey?: pulumi.Input<string>;
    /**
     * The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
     */
    bizType: pulumi.Input<string>;
    /**
     * Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
     */
    dataIdLists?: pulumi.Input<pulumi.Input<string>[]>;
}
