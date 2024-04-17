// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Intelligent Media Management Project resource.
 *
 * For information about Intelligent Media Management Project and how to use it, see [What is Project](https://www.alibabacloud.com/help/en/network-intelligence-service/latest/user-overview).
 *
 * > **NOTE:** Available since v1.134.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tfexample";
 * const role = new alicloud.ram.Role("role", {
 *     name: name,
 *     document: `  {
 *     "Statement": [
 *       {
 *         "Action": "sts:AssumeRole",
 *         "Effect": "Allow",
 *         "Principal": {
 *           "Service": [
 *             "imm.aliyuncs.com"
 *           ]
 *         }
 *       }
 *     ],
 *     "Version": "1"
 *   }
 * `,
 *     description: "this is a role test.",
 *     force: true,
 * });
 * const example = new alicloud.imm.Project("example", {
 *     project: name,
 *     serviceRole: role.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Intelligent Media Management Project can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:imm/project:Project example <project>
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
    public static readonly __pulumiType = 'alicloud:imm/project:Project';

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
     * The name of Project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The service role authorized to the Intelligent Media Management service to access other cloud resources. Default value: `AliyunIMMDefaultRole`. You can also create authorization  roles through the `alicloud.ram.Role`.
     */
    public readonly serviceRole!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["serviceRole"] = state ? state.serviceRole : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceRole"] = args ? args.serviceRole : undefined;
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
     * The name of Project.
     */
    project?: pulumi.Input<string>;
    /**
     * The service role authorized to the Intelligent Media Management service to access other cloud resources. Default value: `AliyunIMMDefaultRole`. You can also create authorization  roles through the `alicloud.ram.Role`.
     */
    serviceRole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The name of Project.
     */
    project: pulumi.Input<string>;
    /**
     * The service role authorized to the Intelligent Media Management service to access other cloud resources. Default value: `AliyunIMMDefaultRole`. You can also create authorization  roles through the `alicloud.ram.Role`.
     */
    serviceRole?: pulumi.Input<string>;
}
