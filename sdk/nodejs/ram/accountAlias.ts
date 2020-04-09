// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RAM cloud account alias.
 * 
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * // Create a alias for cloud account.
 * const alias = new alicloud.ram.AccountAlias("alias", {
 *     accountAlias: "hallo",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ram_account_alias.html.markdown.
 */
export class AccountAlias extends pulumi.CustomResource {
    /**
     * Get an existing AccountAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountAliasState, opts?: pulumi.CustomResourceOptions): AccountAlias {
        return new AccountAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/accountAlias:AccountAlias';

    /**
     * Returns true if the given object is an instance of AccountAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountAlias.__pulumiType;
    }

    /**
     * Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
     */
    public readonly accountAlias!: pulumi.Output<string>;

    /**
     * Create a AccountAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountAliasArgs | AccountAliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccountAliasState | undefined;
            inputs["accountAlias"] = state ? state.accountAlias : undefined;
        } else {
            const args = argsOrState as AccountAliasArgs | undefined;
            if (!args || args.accountAlias === undefined) {
                throw new Error("Missing required property 'accountAlias'");
            }
            inputs["accountAlias"] = args ? args.accountAlias : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccountAlias.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountAlias resources.
 */
export interface AccountAliasState {
    /**
     * Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
     */
    readonly accountAlias?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountAlias resource.
 */
export interface AccountAliasArgs {
    /**
     * Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen.
     */
    readonly accountAlias: pulumi.Input<string>;
}
