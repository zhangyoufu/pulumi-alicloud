// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Group Metric Rule resource.
 *
 * For information about Cloud Monitor Service Group Metric Rule and how to use it, see [What is Group Metric Rule](https://www.alibabacloud.com/help/en/doc-detail/114943.htm).
 *
 * > **NOTE:** Available in v1.104.0+.
 *
 * ## Import
 *
 * Cloud Monitor Service Group Metric Rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cms/groupMetricRule:GroupMetricRule example <rule_id>
 * ```
 */
export class GroupMetricRule extends pulumi.CustomResource {
    /**
     * Get an existing GroupMetricRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMetricRuleState, opts?: pulumi.CustomResourceOptions): GroupMetricRule {
        return new GroupMetricRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/groupMetricRule:GroupMetricRule';

    /**
     * Returns true if the given object is an instance of GroupMetricRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMetricRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMetricRule.__pulumiType;
    }

    /**
     * The abbreviation of the service name.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * Alarm contact group.
     */
    public readonly contactGroups!: pulumi.Output<string>;
    /**
     * The dimensions that specify the resources to be associated with the alert rule.
     */
    public readonly dimensions!: pulumi.Output<string>;
    /**
     * The time period during which the alert rule is effective.
     */
    public readonly effectiveInterval!: pulumi.Output<string | undefined>;
    /**
     * The subject of the alert notification email.                                         .
     */
    public readonly emailSubject!: pulumi.Output<string>;
    /**
     * Alarm level. See the block for escalations.
     */
    public readonly escalations!: pulumi.Output<outputs.cms.GroupMetricRuleEscalations>;
    /**
     * The ID of the application group.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The name of the alert rule.
     */
    public readonly groupMetricRuleName!: pulumi.Output<string>;
    /**
     * The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
     */
    public readonly interval!: pulumi.Output<string | undefined>;
    /**
     * The name of the metric.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The namespace of the service.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * The time period during which the alert rule is ineffective.
     */
    public readonly noEffectiveInterval!: pulumi.Output<string | undefined>;
    /**
     * The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The ID of the alert rule.
     */
    public readonly ruleId!: pulumi.Output<string>;
    /**
     * The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
     */
    public readonly silenceTime!: pulumi.Output<number | undefined>;
    /**
     * The status of Group Metric Rule.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The callback URL.
     */
    public readonly webhook!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMetricRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMetricRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMetricRuleArgs | GroupMetricRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMetricRuleState | undefined;
            inputs["category"] = state ? state.category : undefined;
            inputs["contactGroups"] = state ? state.contactGroups : undefined;
            inputs["dimensions"] = state ? state.dimensions : undefined;
            inputs["effectiveInterval"] = state ? state.effectiveInterval : undefined;
            inputs["emailSubject"] = state ? state.emailSubject : undefined;
            inputs["escalations"] = state ? state.escalations : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["groupMetricRuleName"] = state ? state.groupMetricRuleName : undefined;
            inputs["interval"] = state ? state.interval : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
            inputs["noEffectiveInterval"] = state ? state.noEffectiveInterval : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["silenceTime"] = state ? state.silenceTime : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["webhook"] = state ? state.webhook : undefined;
        } else {
            const args = argsOrState as GroupMetricRuleArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            if ((!args || args.escalations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'escalations'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.groupMetricRuleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupMetricRuleName'");
            }
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            inputs["category"] = args ? args.category : undefined;
            inputs["contactGroups"] = args ? args.contactGroups : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["effectiveInterval"] = args ? args.effectiveInterval : undefined;
            inputs["emailSubject"] = args ? args.emailSubject : undefined;
            inputs["escalations"] = args ? args.escalations : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["groupMetricRuleName"] = args ? args.groupMetricRuleName : undefined;
            inputs["interval"] = args ? args.interval : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["noEffectiveInterval"] = args ? args.noEffectiveInterval : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["ruleId"] = args ? args.ruleId : undefined;
            inputs["silenceTime"] = args ? args.silenceTime : undefined;
            inputs["webhook"] = args ? args.webhook : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupMetricRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMetricRule resources.
 */
export interface GroupMetricRuleState {
    /**
     * The abbreviation of the service name.
     */
    readonly category?: pulumi.Input<string>;
    /**
     * Alarm contact group.
     */
    readonly contactGroups?: pulumi.Input<string>;
    /**
     * The dimensions that specify the resources to be associated with the alert rule.
     */
    readonly dimensions?: pulumi.Input<string>;
    /**
     * The time period during which the alert rule is effective.
     */
    readonly effectiveInterval?: pulumi.Input<string>;
    /**
     * The subject of the alert notification email.                                         .
     */
    readonly emailSubject?: pulumi.Input<string>;
    /**
     * Alarm level. See the block for escalations.
     */
    readonly escalations?: pulumi.Input<inputs.cms.GroupMetricRuleEscalations>;
    /**
     * The ID of the application group.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The name of the alert rule.
     */
    readonly groupMetricRuleName?: pulumi.Input<string>;
    /**
     * The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
     */
    readonly interval?: pulumi.Input<string>;
    /**
     * The name of the metric.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The namespace of the service.
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * The time period during which the alert rule is ineffective.
     */
    readonly noEffectiveInterval?: pulumi.Input<string>;
    /**
     * The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The ID of the alert rule.
     */
    readonly ruleId?: pulumi.Input<string>;
    /**
     * The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
     */
    readonly silenceTime?: pulumi.Input<number>;
    /**
     * The status of Group Metric Rule.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The callback URL.
     */
    readonly webhook?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMetricRule resource.
 */
export interface GroupMetricRuleArgs {
    /**
     * The abbreviation of the service name.
     */
    readonly category: pulumi.Input<string>;
    /**
     * Alarm contact group.
     */
    readonly contactGroups?: pulumi.Input<string>;
    /**
     * The dimensions that specify the resources to be associated with the alert rule.
     */
    readonly dimensions?: pulumi.Input<string>;
    /**
     * The time period during which the alert rule is effective.
     */
    readonly effectiveInterval?: pulumi.Input<string>;
    /**
     * The subject of the alert notification email.                                         .
     */
    readonly emailSubject?: pulumi.Input<string>;
    /**
     * Alarm level. See the block for escalations.
     */
    readonly escalations: pulumi.Input<inputs.cms.GroupMetricRuleEscalations>;
    /**
     * The ID of the application group.
     */
    readonly groupId: pulumi.Input<string>;
    /**
     * The name of the alert rule.
     */
    readonly groupMetricRuleName: pulumi.Input<string>;
    /**
     * The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
     */
    readonly interval?: pulumi.Input<string>;
    /**
     * The name of the metric.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The namespace of the service.
     */
    readonly namespace: pulumi.Input<string>;
    /**
     * The time period during which the alert rule is ineffective.
     */
    readonly noEffectiveInterval?: pulumi.Input<string>;
    /**
     * The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The ID of the alert rule.
     */
    readonly ruleId: pulumi.Input<string>;
    /**
     * The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
     */
    readonly silenceTime?: pulumi.Input<number>;
    /**
     * The callback URL.
     */
    readonly webhook?: pulumi.Input<string>;
}
