// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Dcdn Waf Rule resource.
 *
 * For information about Dcdn Waf Rule and how to use it, see [What is Waf Rule](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchcreatedcdnwafrules).
 *
 * > **NOTE:** Available since v1.201.0.
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
 * const name = config.get("name") || "tf_example";
 * const _default = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const example = new alicloud.dcdn.WafPolicy("example", {
 *     defenseScene: "waf_group",
 *     policyName: `${name}_${_default.result}`,
 *     policyType: "custom",
 *     status: "on",
 * });
 * const exampleWafRule = new alicloud.dcdn.WafRule("example", {
 *     policyId: example.id,
 *     ruleName: name,
 *     conditions: [
 *         {
 *             key: "URI",
 *             opValue: "ne",
 *             values: "/login.php",
 *         },
 *         {
 *             key: "Header",
 *             subKey: "a",
 *             opValue: "eq",
 *             values: "b",
 *         },
 *     ],
 *     status: "on",
 *     action: "monitor",
 *     rateLimit: {
 *         target: "IP",
 *         interval: 5,
 *         threshold: 5,
 *         ttl: 1800,
 *         status: {
 *             code: "200",
 *             ratio: 60,
 *         },
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Dcdn Waf Rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dcdn/wafRule:WafRule example <id>
 * ```
 */
export class WafRule extends pulumi.CustomResource {
    /**
     * Get an existing WafRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafRuleState, opts?: pulumi.CustomResourceOptions): WafRule {
        return new WafRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dcdn/wafRule:WafRule';

    /**
     * Returns true if the given object is an instance of WafRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafRule.__pulumiType;
    }

    /**
     * Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
     */
    public readonly ccStatus!: pulumi.Output<string>;
    /**
     * The blocked regions in the Chinese mainland, separated by commas (,).
     */
    public readonly cnRegionList!: pulumi.Output<string | undefined>;
    /**
     * Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
     */
    public readonly conditions!: pulumi.Output<outputs.dcdn.WafRuleCondition[] | undefined>;
    /**
     * The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
     */
    public /*out*/ readonly defenseScene!: pulumi.Output<string>;
    /**
     * The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
     */
    public readonly effect!: pulumi.Output<string | undefined>;
    /**
     * Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
     */
    public /*out*/ readonly gmtModified!: pulumi.Output<string>;
    /**
     * Blocked regions outside the Chinese mainland, separated by commas (,).
     */
    public readonly otherRegionList!: pulumi.Output<string | undefined>;
    /**
     * The protection policy ID.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See `rateLimit` below.
     */
    public readonly rateLimit!: pulumi.Output<outputs.dcdn.WafRuleRateLimit | undefined>;
    /**
     * The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
     */
    public readonly regularRules!: pulumi.Output<string[] | undefined>;
    /**
     * Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
     */
    public readonly regularTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Filter by IP address.
     */
    public readonly remoteAddrs!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * The types of the protection policies.
     */
    public readonly scenes!: pulumi.Output<string[] | undefined>;
    /**
     * The status of the waf rule. Valid values: `on` and `off`. Default value: on.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
     */
    public readonly wafGroupIds!: pulumi.Output<string | undefined>;

    /**
     * Create a WafRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WafRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafRuleArgs | WafRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["ccStatus"] = state ? state.ccStatus : undefined;
            resourceInputs["cnRegionList"] = state ? state.cnRegionList : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["defenseScene"] = state ? state.defenseScene : undefined;
            resourceInputs["effect"] = state ? state.effect : undefined;
            resourceInputs["gmtModified"] = state ? state.gmtModified : undefined;
            resourceInputs["otherRegionList"] = state ? state.otherRegionList : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["rateLimit"] = state ? state.rateLimit : undefined;
            resourceInputs["regularRules"] = state ? state.regularRules : undefined;
            resourceInputs["regularTypes"] = state ? state.regularTypes : undefined;
            resourceInputs["remoteAddrs"] = state ? state.remoteAddrs : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["scenes"] = state ? state.scenes : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["wafGroupIds"] = state ? state.wafGroupIds : undefined;
        } else {
            const args = argsOrState as WafRuleArgs | undefined;
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["ccStatus"] = args ? args.ccStatus : undefined;
            resourceInputs["cnRegionList"] = args ? args.cnRegionList : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["effect"] = args ? args.effect : undefined;
            resourceInputs["otherRegionList"] = args ? args.otherRegionList : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["rateLimit"] = args ? args.rateLimit : undefined;
            resourceInputs["regularRules"] = args ? args.regularRules : undefined;
            resourceInputs["regularTypes"] = args ? args.regularTypes : undefined;
            resourceInputs["remoteAddrs"] = args ? args.remoteAddrs : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["scenes"] = args ? args.scenes : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["wafGroupIds"] = args ? args.wafGroupIds : undefined;
            resourceInputs["defenseScene"] = undefined /*out*/;
            resourceInputs["gmtModified"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafRule resources.
 */
export interface WafRuleState {
    /**
     * Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
     */
    action?: pulumi.Input<string>;
    /**
     * Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
     */
    ccStatus?: pulumi.Input<string>;
    /**
     * The blocked regions in the Chinese mainland, separated by commas (,).
     */
    cnRegionList?: pulumi.Input<string>;
    /**
     * Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.dcdn.WafRuleCondition>[]>;
    /**
     * The type of protection policy. The following scenarios are supported:-waf_group:Web basic protection-custom_acl: Custom protection policy-whitelist: whitelist
     */
    defenseScene?: pulumi.Input<string>;
    /**
     * The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
     */
    effect?: pulumi.Input<string>;
    /**
     * Revised the time. The date format is based on ISO8601 notation and uses UTC +0 time in the format of yyyy-MM-ddTHH:mm:ssZ.
     */
    gmtModified?: pulumi.Input<string>;
    /**
     * Blocked regions outside the Chinese mainland, separated by commas (,).
     */
    otherRegionList?: pulumi.Input<string>;
    /**
     * The protection policy ID.
     */
    policyId?: pulumi.Input<string>;
    /**
     * The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See `rateLimit` below.
     */
    rateLimit?: pulumi.Input<inputs.dcdn.WafRuleRateLimit>;
    /**
     * The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
     */
    regularRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
     */
    regularTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filter by IP address.
     */
    remoteAddrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * The types of the protection policies.
     */
    scenes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the waf rule. Valid values: `on` and `off`. Default value: on.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
     */
    wafGroupIds?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafRule resource.
 */
export interface WafRuleArgs {
    /**
     * Specifies the action of the rule. Valid values: `block`, `monitor`, `js`.
     */
    action?: pulumi.Input<string>;
    /**
     * Specifies whether to enable rate limiting. Valid values: `on` and `off`. **NOTE:** This parameter is required when policy is of type `customAcl`.
     */
    ccStatus?: pulumi.Input<string>;
    /**
     * The blocked regions in the Chinese mainland, separated by commas (,).
     */
    cnRegionList?: pulumi.Input<string>;
    /**
     * Conditions that trigger the rule. See `conditions` below. **NOTE:** This parameter is required when policy is of type `customAcl` or `whitelist`.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.dcdn.WafRuleCondition>[]>;
    /**
     * The effective scope of the rate limiting blacklist. If you set ccStatus to on, you must configure this parameter. Valid values: `rule` (takes effect for the current rule) and `service` (takes effect globally).
     */
    effect?: pulumi.Input<string>;
    /**
     * Blocked regions outside the Chinese mainland, separated by commas (,).
     */
    otherRegionList?: pulumi.Input<string>;
    /**
     * The protection policy ID.
     */
    policyId: pulumi.Input<string>;
    /**
     * The rules of rate limiting. If you set `ccStatus` to on, you must configure this parameter. See `rateLimit` below.
     */
    rateLimit?: pulumi.Input<inputs.dcdn.WafRuleRateLimit>;
    /**
     * The regular expression.e, when wafGroup appears in tags, this value can be filled in, and only one list of six digits in string format can appear with regultypes.
     */
    regularRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Regular rule type, when wafGroup appears in tags, this value can be filled in, optional values:["sqli", "xss", "codeExec", "crlf", "lfileii", "rfileii", "webshell", "vvip", "other"]
     */
    regularTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filter by IP address.
     */
    remoteAddrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the protection rule. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_). **NOTE:** This parameter cannot be modified when policy is of type `regionBlock`.
     */
    ruleName: pulumi.Input<string>;
    /**
     * The types of the protection policies.
     */
    scenes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the waf rule. Valid values: `on` and `off`. Default value: on.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the waf rule group. The default value is "1012". Multiple rules are separated by commas.
     */
    wafGroupIds?: pulumi.Input<string>;
}
