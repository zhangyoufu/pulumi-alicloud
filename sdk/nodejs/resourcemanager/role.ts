// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Resource Manager role resource. Members are resource containers in the resource directory, which can physically isolate resources to form an independent resource grouping unit. You can create members in the resource folder to manage them in a unified manner.
 * For information about Resource Manager role and how to use it, see [What is Resource Manager role](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).
 *
 * > **NOTE:** Available in v1.82.0+.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Add a Resource Manager role.
 * const example = new alicloud.resourcemanager.Role("example", {
 *     assumeRolePolicyDocument: `     {
 *           "Statement": [
 *                {
 *                     "Action": "sts:AssumeRole",
 *                     "Effect": "Allow",
 *                     "Principal": {
 *                         "RAM":[
 *                                 "acs:ram::103755469187****:root"，
 *                                 "acs:ram::104408977069****:root"
 *                         ]
 *                     }
 *                 }
 *           ],
 *           "Version": "1"
 *      }
 * 	 `,
 *     roleName: "testrd",
 * });
 * ```
 *
 * ## Import
 *
 * Resource Manager can be imported using the id or role_name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:resourcemanager/role:Role example testrd
 * ```
 */
export class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState, opts?: pulumi.CustomResourceOptions): Role {
        return new Role(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:resourcemanager/role:Role';

    /**
     * Returns true if the given object is an instance of Role.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Role {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Role.__pulumiType;
    }

    /**
     * The resource descriptor of the role.
     * * `createDate` (Removed form v1.114.0) - Role creation time.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The content of the permissions strategy that plays a role.
     */
    public readonly assumeRolePolicyDocument!: pulumi.Output<string>;
    /**
     * The description of the Resource Manager role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     */
    public readonly maxSessionDuration!: pulumi.Output<number | undefined>;
    public /*out*/ readonly roleId!: pulumi.Output<string>;
    /**
     * Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * Role update time.
     */
    public /*out*/ readonly updateDate!: pulumi.Output<string>;

    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleArgs | RoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["assumeRolePolicyDocument"] = state ? state.assumeRolePolicyDocument : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxSessionDuration"] = state ? state.maxSessionDuration : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["updateDate"] = state ? state.updateDate : undefined;
        } else {
            const args = argsOrState as RoleArgs | undefined;
            if ((!args || args.assumeRolePolicyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assumeRolePolicyDocument'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["assumeRolePolicyDocument"] = args ? args.assumeRolePolicyDocument : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxSessionDuration"] = args ? args.maxSessionDuration : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["roleId"] = undefined /*out*/;
            inputs["updateDate"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Role.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * The resource descriptor of the role.
     * * `createDate` (Removed form v1.114.0) - Role creation time.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The content of the permissions strategy that plays a role.
     */
    readonly assumeRolePolicyDocument?: pulumi.Input<string>;
    /**
     * The description of the Resource Manager role.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     */
    readonly maxSessionDuration?: pulumi.Input<number>;
    readonly roleId?: pulumi.Input<string>;
    /**
     * Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * Role update time.
     */
    readonly updateDate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * The content of the permissions strategy that plays a role.
     */
    readonly assumeRolePolicyDocument: pulumi.Input<string>;
    /**
     * The description of the Resource Manager role.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Role maximum session time. Valid values: [3600-43200]. Default to `3600`.
     */
    readonly maxSessionDuration?: pulumi.Input<number>;
    /**
     * Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-".
     */
    readonly roleName: pulumi.Input<string>;
}
