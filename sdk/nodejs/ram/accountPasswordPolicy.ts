// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Empty resource sets defaults values for every property.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.ram.AccountPasswordPolicy("default", {});
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const corporate = new alicloud.ram.AccountPasswordPolicy("corporate", {
 *     hardExpiry: true,
 *     maxLoginAttempts: 3,
 *     maxPasswordAge: 12,
 *     minimumPasswordLength: 9,
 *     passwordReusePrevention: 5,
 *     requireLowercaseCharacters: false,
 *     requireNumbers: false,
 *     requireSymbols: false,
 *     requireUppercaseCharacters: false,
 * });
 * ```
 * For not specified values sets defaults.
 *
 * ## Import
 *
 * RAM account password policy can be imported using the `id`, e.g.
 *
 *  bash
 *
 * ```sh
 * $ pulumi import alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy example ram-account-password-policy
 * ```
 */
export class AccountPasswordPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccountPasswordPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountPasswordPolicyState, opts?: pulumi.CustomResourceOptions): AccountPasswordPolicy {
        return new AccountPasswordPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/accountPasswordPolicy:AccountPasswordPolicy';

    /**
     * Returns true if the given object is an instance of AccountPasswordPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountPasswordPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountPasswordPolicy.__pulumiType;
    }

    /**
     * Specifies if a password can expire in a hard way. Default to false.
     */
    public readonly hardExpiry!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
     */
    public readonly maxLoginAttempts!: pulumi.Output<number | undefined>;
    /**
     * The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
     */
    public readonly maxPasswordAge!: pulumi.Output<number | undefined>;
    /**
     * Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
     */
    public readonly minimumPasswordLength!: pulumi.Output<number | undefined>;
    /**
     * User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
     */
    public readonly passwordReusePrevention!: pulumi.Output<number | undefined>;
    /**
     * Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
     */
    public readonly requireLowercaseCharacters!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the occurrence of a number in the password is mandatory. Default to true.
     */
    public readonly requireNumbers!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the occurrence of a special character in the password is mandatory. Default to true.
     */
    public readonly requireSymbols!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
     */
    public readonly requireUppercaseCharacters!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AccountPasswordPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountPasswordPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountPasswordPolicyArgs | AccountPasswordPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountPasswordPolicyState | undefined;
            resourceInputs["hardExpiry"] = state ? state.hardExpiry : undefined;
            resourceInputs["maxLoginAttempts"] = state ? state.maxLoginAttempts : undefined;
            resourceInputs["maxPasswordAge"] = state ? state.maxPasswordAge : undefined;
            resourceInputs["minimumPasswordLength"] = state ? state.minimumPasswordLength : undefined;
            resourceInputs["passwordReusePrevention"] = state ? state.passwordReusePrevention : undefined;
            resourceInputs["requireLowercaseCharacters"] = state ? state.requireLowercaseCharacters : undefined;
            resourceInputs["requireNumbers"] = state ? state.requireNumbers : undefined;
            resourceInputs["requireSymbols"] = state ? state.requireSymbols : undefined;
            resourceInputs["requireUppercaseCharacters"] = state ? state.requireUppercaseCharacters : undefined;
        } else {
            const args = argsOrState as AccountPasswordPolicyArgs | undefined;
            resourceInputs["hardExpiry"] = args ? args.hardExpiry : undefined;
            resourceInputs["maxLoginAttempts"] = args ? args.maxLoginAttempts : undefined;
            resourceInputs["maxPasswordAge"] = args ? args.maxPasswordAge : undefined;
            resourceInputs["minimumPasswordLength"] = args ? args.minimumPasswordLength : undefined;
            resourceInputs["passwordReusePrevention"] = args ? args.passwordReusePrevention : undefined;
            resourceInputs["requireLowercaseCharacters"] = args ? args.requireLowercaseCharacters : undefined;
            resourceInputs["requireNumbers"] = args ? args.requireNumbers : undefined;
            resourceInputs["requireSymbols"] = args ? args.requireSymbols : undefined;
            resourceInputs["requireUppercaseCharacters"] = args ? args.requireUppercaseCharacters : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountPasswordPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountPasswordPolicy resources.
 */
export interface AccountPasswordPolicyState {
    /**
     * Specifies if a password can expire in a hard way. Default to false.
     */
    hardExpiry?: pulumi.Input<boolean>;
    /**
     * Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
     */
    maxLoginAttempts?: pulumi.Input<number>;
    /**
     * The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
     */
    maxPasswordAge?: pulumi.Input<number>;
    /**
     * Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
     */
    minimumPasswordLength?: pulumi.Input<number>;
    /**
     * User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
     */
    passwordReusePrevention?: pulumi.Input<number>;
    /**
     * Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
     */
    requireLowercaseCharacters?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of a number in the password is mandatory. Default to true.
     */
    requireNumbers?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of a special character in the password is mandatory. Default to true.
     */
    requireSymbols?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
     */
    requireUppercaseCharacters?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AccountPasswordPolicy resource.
 */
export interface AccountPasswordPolicyArgs {
    /**
     * Specifies if a password can expire in a hard way. Default to false.
     */
    hardExpiry?: pulumi.Input<boolean>;
    /**
     * Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5.
     */
    maxLoginAttempts?: pulumi.Input<number>;
    /**
     * The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.
     */
    maxPasswordAge?: pulumi.Input<number>;
    /**
     * Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.
     */
    minimumPasswordLength?: pulumi.Input<number>;
    /**
     * User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.
     */
    passwordReusePrevention?: pulumi.Input<number>;
    /**
     * Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.
     */
    requireLowercaseCharacters?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of a number in the password is mandatory. Default to true.
     */
    requireNumbers?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of a special character in the password is mandatory. Default to true.
     */
    requireSymbols?: pulumi.Input<boolean>;
    /**
     * Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.
     */
    requireUppercaseCharacters?: pulumi.Input<boolean>;
}
