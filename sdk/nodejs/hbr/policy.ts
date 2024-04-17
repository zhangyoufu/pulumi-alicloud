// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a HBR Policy resource.
 *
 * For information about HBR Policy and how to use it, see [What is Policy](https://www.alibabacloud.com/help/en/cloud-backup/developer-reference/api-hbr-2017-09-08-createpolicyv2).
 *
 * > **NOTE:** Available since v1.221.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new random.index.Integer("default", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultyk84Hc = new alicloud.hbr.Vault("defaultyk84Hc", {
 *     vaultType: "STANDARD",
 *     vaultName: `example-value-${_default.result}`,
 * });
 * const defaultoqWvHQ = new alicloud.hbr.Policy("defaultoqWvHQ", {
 *     policyName: `example-value-${_default.result}`,
 *     rules: [{
 *         ruleType: "BACKUP",
 *         backupType: "COMPLETE",
 *         schedule: "I|1631685600|P1D",
 *         retention: 7,
 *         archiveDays: 0,
 *         vaultId: defaultyk84Hc.id,
 *     }],
 *     policyDescription: "policy example",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * HBR Policy can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:hbr/policy:Policy example <id>
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:hbr/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Policy creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The policy description.
     */
    public readonly policyDescription!: pulumi.Output<string | undefined>;
    /**
     * Policy Name.
     */
    public readonly policyName!: pulumi.Output<string | undefined>;
    /**
     * A list of policy rules. See `rules` below.
     */
    public readonly rules!: pulumi.Output<outputs.hbr.PolicyRule[] | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["policyDescription"] = state ? state.policyDescription : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            resourceInputs["policyDescription"] = args ? args.policyDescription : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Policy creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The policy description.
     */
    policyDescription?: pulumi.Input<string>;
    /**
     * Policy Name.
     */
    policyName?: pulumi.Input<string>;
    /**
     * A list of policy rules. See `rules` below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.hbr.PolicyRule>[]>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The policy description.
     */
    policyDescription?: pulumi.Input<string>;
    /**
     * Policy Name.
     */
    policyName?: pulumi.Input<string>;
    /**
     * A list of policy rules. See `rules` below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.hbr.PolicyRule>[]>;
}
