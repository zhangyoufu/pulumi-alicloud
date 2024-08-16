// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a SLS Project resource.
 *
 * For information about SLS Project and how to use it, see [What is Project](https://www.alibabacloud.com/help/en/sls/developer-reference/api-createproject).
 *
 * > **NOTE:** Available since v1.9.5.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const _default = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const example = new alicloud.log.Project("example", {
 *     name: `terraform-example-${_default.result}`,
 *     description: "terraform-example",
 *     tags: {
 *         Created: "TF",
 *         For: "example",
 *     },
 * });
 * ```
 *
 * Project With Policy Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const _default = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const examplePolicy = new alicloud.log.Project("example_policy", {
 *     name: `terraform-example-${_default.result}`,
 *     description: "terraform-example",
 *     policy: `{
 *   "Statement": [
 *     {
 *       "Action": [
 *         "log:PostLogStoreLogs"
 *       ],
 *       "Condition": {
 *         "StringNotLike": {
 *           "acs:SourceVpc": [
 *             "vpc-*"
 *           ]
 *         }
 *       },
 *       "Effect": "Deny",
 *       "Resource": "acs:log:*:*:project/tf-log/*"
 *     }
 *   ],
 *   "Version": "1"
 * }
 * `,
 * });
 * ```
 *
 * ## Module Support
 *
 * You can use the existing sls module
 * to create SLS project, store and store index one-click, like ECS instances.
 *
 * ## Import
 *
 * SLS Project can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:log/project:Project example <id>
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:log/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * CreateTime.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
     * - The project name must be globally unique.
     * - The name can contain only lowercase letters, digits, and hyphens (-).
     * - It must start and end with a lowercase letter or number.
     * - The value contains 3 to 63 characters.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The status of the resource.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * CreateTime.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    policy?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
     * - The project name must be globally unique.
     * - The name can contain only lowercase letters, digits, and hyphens (-).
     * - It must start and end with a lowercase letter or number.
     * - The value contains 3 to 63 characters.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The status of the resource.
     */
    status?: pulumi.Input<string>;
    /**
     * Tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
     *
     * @deprecated Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
     */
    name?: pulumi.Input<string>;
    /**
     * Log project policy, used to set a policy for a project.
     */
    policy?: pulumi.Input<string>;
    /**
     * The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
     * - The project name must be globally unique.
     * - The name can contain only lowercase letters, digits, and hyphens (-).
     * - It must start and end with a lowercase letter or number.
     * - The value contains 3 to 63 characters.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Tag.
     *
     * The following arguments will be discarded. Please use new fields as soon as possible:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
