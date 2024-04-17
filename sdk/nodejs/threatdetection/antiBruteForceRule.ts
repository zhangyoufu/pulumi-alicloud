// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Threat Detection Anti Brute Force Rule resource.
 *
 * For information about Threat Detection Anti Brute Force Rule and how to use it, see [What is Anti Brute Force Rule](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createantibruteforcerule).
 *
 * > **NOTE:** Available since v1.195.0.
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
 * const _default = new alicloud.threatdetection.AntiBruteForceRule("default", {
 *     antiBruteForceRuleName: "apispec_example",
 *     forbiddenTime: 360,
 *     uuidLists: ["032b618f-b220-4a0d-bd37-fbdc6ef58b6a"],
 *     failCount: 80,
 *     span: 10,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Threat Detection Anti Brute Force Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule example <id>
 * ```
 */
export class AntiBruteForceRule extends pulumi.CustomResource {
    /**
     * Get an existing AntiBruteForceRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AntiBruteForceRuleState, opts?: pulumi.CustomResourceOptions): AntiBruteForceRule {
        return new AntiBruteForceRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:threatdetection/antiBruteForceRule:AntiBruteForceRule';

    /**
     * Returns true if the given object is an instance of AntiBruteForceRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AntiBruteForceRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AntiBruteForceRule.__pulumiType;
    }

    /**
     * The ID of the defense rule.
     */
    public /*out*/ readonly antiBruteForceRuleId!: pulumi.Output<string>;
    /**
     * The name of the defense rule.
     */
    public readonly antiBruteForceRuleName!: pulumi.Output<string>;
    /**
     * Specifies whether to set the defense rule as the default rule.
     */
    public readonly defaultRule!: pulumi.Output<boolean>;
    /**
     * The threshold for the number of failed user logins when the brute-force defense rule takes effect.
     */
    public readonly failCount!: pulumi.Output<number>;
    /**
     * The period of time during which logons from an account are not allowed. Unit: minutes.
     */
    public readonly forbiddenTime!: pulumi.Output<number>;
    /**
     * The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
     */
    public readonly span!: pulumi.Output<number>;
    /**
     * An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
     */
    public readonly uuidLists!: pulumi.Output<string[]>;

    /**
     * Create a AntiBruteForceRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AntiBruteForceRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AntiBruteForceRuleArgs | AntiBruteForceRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AntiBruteForceRuleState | undefined;
            resourceInputs["antiBruteForceRuleId"] = state ? state.antiBruteForceRuleId : undefined;
            resourceInputs["antiBruteForceRuleName"] = state ? state.antiBruteForceRuleName : undefined;
            resourceInputs["defaultRule"] = state ? state.defaultRule : undefined;
            resourceInputs["failCount"] = state ? state.failCount : undefined;
            resourceInputs["forbiddenTime"] = state ? state.forbiddenTime : undefined;
            resourceInputs["span"] = state ? state.span : undefined;
            resourceInputs["uuidLists"] = state ? state.uuidLists : undefined;
        } else {
            const args = argsOrState as AntiBruteForceRuleArgs | undefined;
            if ((!args || args.antiBruteForceRuleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'antiBruteForceRuleName'");
            }
            if ((!args || args.failCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'failCount'");
            }
            if ((!args || args.forbiddenTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'forbiddenTime'");
            }
            if ((!args || args.span === undefined) && !opts.urn) {
                throw new Error("Missing required property 'span'");
            }
            if ((!args || args.uuidLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uuidLists'");
            }
            resourceInputs["antiBruteForceRuleName"] = args ? args.antiBruteForceRuleName : undefined;
            resourceInputs["defaultRule"] = args ? args.defaultRule : undefined;
            resourceInputs["failCount"] = args ? args.failCount : undefined;
            resourceInputs["forbiddenTime"] = args ? args.forbiddenTime : undefined;
            resourceInputs["span"] = args ? args.span : undefined;
            resourceInputs["uuidLists"] = args ? args.uuidLists : undefined;
            resourceInputs["antiBruteForceRuleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AntiBruteForceRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AntiBruteForceRule resources.
 */
export interface AntiBruteForceRuleState {
    /**
     * The ID of the defense rule.
     */
    antiBruteForceRuleId?: pulumi.Input<string>;
    /**
     * The name of the defense rule.
     */
    antiBruteForceRuleName?: pulumi.Input<string>;
    /**
     * Specifies whether to set the defense rule as the default rule.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * The threshold for the number of failed user logins when the brute-force defense rule takes effect.
     */
    failCount?: pulumi.Input<number>;
    /**
     * The period of time during which logons from an account are not allowed. Unit: minutes.
     */
    forbiddenTime?: pulumi.Input<number>;
    /**
     * The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
     */
    span?: pulumi.Input<number>;
    /**
     * An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
     */
    uuidLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AntiBruteForceRule resource.
 */
export interface AntiBruteForceRuleArgs {
    /**
     * The name of the defense rule.
     */
    antiBruteForceRuleName: pulumi.Input<string>;
    /**
     * Specifies whether to set the defense rule as the default rule.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * The threshold for the number of failed user logins when the brute-force defense rule takes effect.
     */
    failCount: pulumi.Input<number>;
    /**
     * The period of time during which logons from an account are not allowed. Unit: minutes.
     */
    forbiddenTime: pulumi.Input<number>;
    /**
     * The period of time during which logon failures from an account are measured. Unit: minutes. If Span is set to 10, the defense rule takes effect when the logon failures measured within 10 minutes reaches the specified threshold. The IP address of attackers cannot be used to log on to the server in the specified period of time.
     */
    span: pulumi.Input<number>;
    /**
     * An array consisting of the UUIDs of servers to which the defense rule is applied.**The binding status must be Enterprise Edition.**
     */
    uuidLists: pulumi.Input<pulumi.Input<string>[]>;
}
