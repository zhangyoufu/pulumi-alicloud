// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Elastic Desktop Service (ECD) User resource.
 *
 * For information about Elastic Desktop Service (ECD) User and how to use it, see [What is User](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-eds-user-2021-03-08-createusers-desktop).
 *
 * > **NOTE:** Available since v1.142.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.eds.User("default", {
 *     email: "tf.example@abc.com",
 *     endUserId: "terraform_example123",
 *     password: "Example_123",
 *     phone: "18888888888",
 * });
 * ```
 *
 * ## Import
 *
 * ECD User can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/user:User example <end_user_id>
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The email of the user email.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
     */
    public readonly endUserId!: pulumi.Output<string>;
    /**
     * The password of the user password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The phone of the mobile phone number.
     */
    public readonly phone!: pulumi.Output<string | undefined>;
    /**
     * The status of the resource. Valid values: `Unlocked`, `Locked`.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["endUserId"] = state ? state.endUserId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["phone"] = state ? state.phone : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.endUserId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endUserId'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["endUserId"] = args ? args.endUserId : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The email of the user email.
     */
    email?: pulumi.Input<string>;
    /**
     * The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
     */
    endUserId?: pulumi.Input<string>;
    /**
     * The password of the user password.
     */
    password?: pulumi.Input<string>;
    /**
     * The phone of the mobile phone number.
     */
    phone?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Unlocked`, `Locked`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The email of the user email.
     */
    email: pulumi.Input<string>;
    /**
     * The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
     */
    endUserId: pulumi.Input<string>;
    /**
     * The password of the user password.
     */
    password?: pulumi.Input<string>;
    /**
     * The phone of the mobile phone number.
     */
    phone?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `Unlocked`, `Locked`.
     */
    status?: pulumi.Input<string>;
}
