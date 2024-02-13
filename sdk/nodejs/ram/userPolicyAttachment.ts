// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RAM User Policy attachment resource.
 *
 * > **NOTE:** Available since v1.0.0.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a RAM User Policy attachment.
 * const user = new alicloud.ram.User("user", {
 *     displayName: "user_display_name",
 *     mobile: "86-18688888888",
 *     email: "hello.uuu@aaa.com",
 *     comments: "yoyoyo",
 * });
 * const policy = new alicloud.ram.Policy("policy", {
 *     document: `  {
 *     "Statement": [
 *       {
 *         "Action": [
 *           "oss:ListObjects",
 *           "oss:GetObject"
 *         ],
 *         "Effect": "Allow",
 *         "Resource": [
 *           "acs:oss:*:*:mybucket",
 *           "acs:oss:*:*:mybucket/*"
 *         ]
 *       }
 *     ],
 *       "Version": "1"
 *   }
 * `,
 *     description: "this is a policy test",
 * });
 * const attach = new alicloud.ram.UserPolicyAttachment("attach", {
 *     policyName: policy.name,
 *     policyType: policy.type,
 *     userName: user.name,
 * });
 * ```
 *
 * ## Import
 *
 * RAM User Policy attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ram/userPolicyAttachment:UserPolicyAttachment example user:my-policy:Custom:my-user
 * ```
 */
export class UserPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing UserPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): UserPolicyAttachment {
        return new UserPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/userPolicyAttachment:UserPolicyAttachment';

    /**
     * Returns true if the given object is an instance of UserPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPolicyAttachment.__pulumiType;
    }

    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    public readonly policyType!: pulumi.Output<string>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a UserPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPolicyAttachmentArgs | UserPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserPolicyAttachmentState | undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserPolicyAttachmentArgs | undefined;
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            if ((!args || args.policyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyType'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPolicyAttachment resources.
 */
export interface UserPolicyAttachmentState {
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    policyName?: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    policyType?: pulumi.Input<string>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPolicyAttachment resource.
 */
export interface UserPolicyAttachmentArgs {
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    policyName: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    policyType: pulumi.Input<string>;
    /**
     * Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
     */
    userName: pulumi.Input<string>;
}
