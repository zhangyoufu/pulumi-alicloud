// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RAM Group Policy attachment resource. 
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a RAM Group Policy attachment.
 * const group = new alicloud.ram.Group("group", {
 *     comments: "this is a group comments.",
 *     force: true,
 * });
 * const policy = new alicloud.ram.Policy("policy", {
 *     description: "this is a policy test",
 *     document: `    {
 *       "Statement": [
 *         {
 *           "Action": [
 *             "oss:ListObjects",
 *             "oss:GetObject"
 *           ],
 *           "Effect": "Allow",
 *           "Resource": [
 *             "acs:oss:*:*:mybucket",
 *             "acs:oss:*:*:mybucket/*"
 *           ]
 *         }
 *       ],
 *         "Version": "1"
 *     }
 *   `,
 *     force: true,
 * });
 * const attach = new alicloud.ram.GroupPolicyAttachment("attach", {
 *     groupName: group.name,
 *     policyName: policy.name,
 *     policyType: policy.type,
 * });
 * ```
 */
export class GroupPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing GroupPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): GroupPolicyAttachment {
        return new GroupPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/groupPolicyAttachment:GroupPolicyAttachment';

    /**
     * Returns true if the given object is an instance of GroupPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupPolicyAttachment.__pulumiType;
    }

    /**
     * Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    public readonly policyType!: pulumi.Output<string>;

    /**
     * Create a GroupPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupPolicyAttachmentArgs | GroupPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupPolicyAttachmentState | undefined;
            inputs["groupName"] = state ? state.groupName : undefined;
            inputs["policyName"] = state ? state.policyName : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
        } else {
            const args = argsOrState as GroupPolicyAttachmentArgs | undefined;
            if (!args || args.groupName === undefined) {
                throw new Error("Missing required property 'groupName'");
            }
            if (!args || args.policyName === undefined) {
                throw new Error("Missing required property 'policyName'");
            }
            if (!args || args.policyType === undefined) {
                throw new Error("Missing required property 'policyType'");
            }
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupPolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupPolicyAttachment resources.
 */
export interface GroupPolicyAttachmentState {
    /**
     * Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly groupName?: pulumi.Input<string>;
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly policyName?: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    readonly policyType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupPolicyAttachment resource.
 */
export interface GroupPolicyAttachmentArgs {
    /**
     * Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly groupName: pulumi.Input<string>;
    /**
     * Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
     */
    readonly policyName: pulumi.Input<string>;
    /**
     * Type of the RAM policy. It must be `Custom` or `System`.
     */
    readonly policyType: pulumi.Input<string>;
}
