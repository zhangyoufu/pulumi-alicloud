// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a OOS Template resource. For information about Alicloud OOS Template and how to use it, see [What is Resource Alicloud OOS Template](https://www.alibabacloud.com/help/doc-detail/120761.htm).
 *
 * > **NOTE:** Available in 1.92.0+.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.oos.Template("example", {
 *     content: `  {
 *     "FormatVersion": "OOS-2019-06-01",
 *     "Description": "Update Describe instances of given status",
 *     "Parameters":{
 *       "Status":{
 *         "Type": "String",
 *         "Description": "(Required) The status of the Ecs instance."
 *       }
 *     },
 *     "Tasks": [
 *       {
 *         "Properties" :{
 *           "Parameters":{
 *             "Status": "{{ Status }}"
 *           },
 *           "API": "DescribeInstances",
 *           "Service": "Ecs"
 *         },
 *         "Name": "foo",
 *         "Action": "ACS::ExecuteApi"
 *       }]
 *   }
 * `,
 *     templateName: "test-name",
 *     versionName: "test",
 *     tags: {
 *         Created: "TF",
 *         For: "acceptance Test",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * OOS Template can be imported using the id or template_name, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:oos/template:Template example template_name
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:oos/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * When deleting a template, whether to delete its related executions. Default to `false`.
     */
    public readonly autoDeleteExecutions!: pulumi.Output<boolean | undefined>;
    /**
     * The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The creator of the template.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * The time when the template is created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The description of the template.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Is it triggered successfully.
     */
    public /*out*/ readonly hasTrigger!: pulumi.Output<boolean>;
    /**
     * The ID of resource group which the template belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
     */
    public /*out*/ readonly shareType!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
     */
    public /*out*/ readonly templateFormat!: pulumi.Output<string>;
    /**
     * The id of OOS Template.
     */
    public /*out*/ readonly templateId!: pulumi.Output<string>;
    /**
     * The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
     */
    public /*out*/ readonly templateType!: pulumi.Output<string>;
    /**
     * The version of OOS Template.
     */
    public /*out*/ readonly templateVersion!: pulumi.Output<string>;
    /**
     * The user who updated the template.
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<string>;
    /**
     * The time when the template was updated.
     */
    public /*out*/ readonly updatedDate!: pulumi.Output<string>;
    /**
     * The name of template version.
     */
    public readonly versionName!: pulumi.Output<string | undefined>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateState | undefined;
            resourceInputs["autoDeleteExecutions"] = state ? state.autoDeleteExecutions : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["createdBy"] = state ? state.createdBy : undefined;
            resourceInputs["createdDate"] = state ? state.createdDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hasTrigger"] = state ? state.hasTrigger : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["shareType"] = state ? state.shareType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["templateFormat"] = state ? state.templateFormat : undefined;
            resourceInputs["templateId"] = state ? state.templateId : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["templateType"] = state ? state.templateType : undefined;
            resourceInputs["templateVersion"] = state ? state.templateVersion : undefined;
            resourceInputs["updatedBy"] = state ? state.updatedBy : undefined;
            resourceInputs["updatedDate"] = state ? state.updatedDate : undefined;
            resourceInputs["versionName"] = state ? state.versionName : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["autoDeleteExecutions"] = args ? args.autoDeleteExecutions : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["versionName"] = args ? args.versionName : undefined;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["createdDate"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hasTrigger"] = undefined /*out*/;
            resourceInputs["shareType"] = undefined /*out*/;
            resourceInputs["templateFormat"] = undefined /*out*/;
            resourceInputs["templateId"] = undefined /*out*/;
            resourceInputs["templateType"] = undefined /*out*/;
            resourceInputs["templateVersion"] = undefined /*out*/;
            resourceInputs["updatedBy"] = undefined /*out*/;
            resourceInputs["updatedDate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * When deleting a template, whether to delete its related executions. Default to `false`.
     */
    autoDeleteExecutions?: pulumi.Input<boolean>;
    /**
     * The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
     */
    content?: pulumi.Input<string>;
    /**
     * The creator of the template.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * The time when the template is created.
     */
    createdDate?: pulumi.Input<string>;
    /**
     * The description of the template.
     */
    description?: pulumi.Input<string>;
    /**
     * Is it triggered successfully.
     */
    hasTrigger?: pulumi.Input<boolean>;
    /**
     * The ID of resource group which the template belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
     */
    shareType?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
     */
    templateFormat?: pulumi.Input<string>;
    /**
     * The id of OOS Template.
     */
    templateId?: pulumi.Input<string>;
    /**
     * The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
     */
    templateName?: pulumi.Input<string>;
    /**
     * The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
     */
    templateType?: pulumi.Input<string>;
    /**
     * The version of OOS Template.
     */
    templateVersion?: pulumi.Input<string>;
    /**
     * The user who updated the template.
     */
    updatedBy?: pulumi.Input<string>;
    /**
     * The time when the template was updated.
     */
    updatedDate?: pulumi.Input<string>;
    /**
     * The name of template version.
     */
    versionName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * When deleting a template, whether to delete its related executions. Default to `false`.
     */
    autoDeleteExecutions?: pulumi.Input<boolean>;
    /**
     * The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
     */
    content: pulumi.Input<string>;
    /**
     * The ID of resource group which the template belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
     */
    templateName: pulumi.Input<string>;
    /**
     * The name of template version.
     */
    versionName?: pulumi.Input<string>;
}
