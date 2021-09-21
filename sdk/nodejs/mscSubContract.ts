// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Msc Sub Contact resource.
 *
 * > **NOTE:** Available in v1.132.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.MscSubContract("default", {
 *     contactName: example_value,
 *     position: "CEO",
 *     email: "123@163.com",
 *     mobile: "153xxxxx906",
 * });
 * ```
 *
 * ## Import
 *
 * Msc Sub Contact can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:index/mscSubContract:MscSubContract example <id>
 * ```
 */
export class MscSubContract extends pulumi.CustomResource {
    /**
     * Get an existing MscSubContract resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MscSubContractState, opts?: pulumi.CustomResourceOptions): MscSubContract {
        return new MscSubContract(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:index/mscSubContract:MscSubContract';

    /**
     * Returns true if the given object is an instance of MscSubContract.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MscSubContract {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MscSubContract.__pulumiType;
    }

    /**
     * The User's Contact Name. **Note:** The name must be 2 to 12 characters in length, and can contain uppercase and lowercase letters.
     */
    public readonly contactName!: pulumi.Output<string>;
    /**
     * The User's Contact Email Address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The User's Telephone.
     */
    public readonly mobile!: pulumi.Output<string>;
    /**
     * The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Others`.
     */
    public readonly position!: pulumi.Output<string>;

    /**
     * Create a MscSubContract resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MscSubContractArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MscSubContractArgs | MscSubContractState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MscSubContractState | undefined;
            inputs["contactName"] = state ? state.contactName : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["mobile"] = state ? state.mobile : undefined;
            inputs["position"] = state ? state.position : undefined;
        } else {
            const args = argsOrState as MscSubContractArgs | undefined;
            if ((!args || args.contactName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactName'");
            }
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.mobile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mobile'");
            }
            if ((!args || args.position === undefined) && !opts.urn) {
                throw new Error("Missing required property 'position'");
            }
            inputs["contactName"] = args ? args.contactName : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["mobile"] = args ? args.mobile : undefined;
            inputs["position"] = args ? args.position : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MscSubContract.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MscSubContract resources.
 */
export interface MscSubContractState {
    /**
     * The User's Contact Name. **Note:** The name must be 2 to 12 characters in length, and can contain uppercase and lowercase letters.
     */
    readonly contactName?: pulumi.Input<string>;
    /**
     * The User's Contact Email Address.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The User's Telephone.
     */
    readonly mobile?: pulumi.Input<string>;
    /**
     * The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Others`.
     */
    readonly position?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MscSubContract resource.
 */
export interface MscSubContractArgs {
    /**
     * The User's Contact Name. **Note:** The name must be 2 to 12 characters in length, and can contain uppercase and lowercase letters.
     */
    readonly contactName: pulumi.Input<string>;
    /**
     * The User's Contact Email Address.
     */
    readonly email: pulumi.Input<string>;
    /**
     * The User's Telephone.
     */
    readonly mobile: pulumi.Input<string>;
    /**
     * The User's Position. Valid values: `CEO`, `Technical Director`, `Maintenance Director`, `Project Director`,`Finance Director` and `Others`.
     */
    readonly position: pulumi.Input<string>;
}
