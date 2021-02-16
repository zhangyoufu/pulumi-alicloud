// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Ess alarm can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ess/alarm:Alarm example asg-2ze500_045efffe-4d05
 * ```
 */
export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmState, opts?: pulumi.CustomResourceOptions): Alarm {
        return new Alarm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/alarm:Alarm';

    /**
     * Returns true if the given object is an instance of Alarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alarm.__pulumiType;
    }

    /**
     * The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
     */
    public readonly alarmActions!: pulumi.Output<string[]>;
    /**
     * Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
     */
    public readonly cloudMonitorGroupId!: pulumi.Output<number | undefined>;
    /**
     * The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
     */
    public readonly comparisonOperator!: pulumi.Output<string | undefined>;
    /**
     * The description for the alarm.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
     */
    public readonly dimensions!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether to enable specific ess alarm. Default to true.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
     */
    public readonly evaluationCount!: pulumi.Output<number | undefined>;
    /**
     * The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
     */
    public readonly metricType!: pulumi.Output<string | undefined>;
    /**
     * The name for ess alarm.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * The state of specified alarm.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
     */
    public readonly statistics!: pulumi.Output<string | undefined>;
    /**
     * The value against which the specified statistics is compared.
     */
    public readonly threshold!: pulumi.Output<string>;

    /**
     * Create a Alarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmArgs | AlarmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmState | undefined;
            inputs["alarmActions"] = state ? state.alarmActions : undefined;
            inputs["cloudMonitorGroupId"] = state ? state.cloudMonitorGroupId : undefined;
            inputs["comparisonOperator"] = state ? state.comparisonOperator : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dimensions"] = state ? state.dimensions : undefined;
            inputs["enable"] = state ? state.enable : undefined;
            inputs["evaluationCount"] = state ? state.evaluationCount : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["metricType"] = state ? state.metricType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["statistics"] = state ? state.statistics : undefined;
            inputs["threshold"] = state ? state.threshold : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if ((!args || args.alarmActions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alarmActions'");
            }
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            if ((!args || args.threshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'threshold'");
            }
            inputs["alarmActions"] = args ? args.alarmActions : undefined;
            inputs["cloudMonitorGroupId"] = args ? args.cloudMonitorGroupId : undefined;
            inputs["comparisonOperator"] = args ? args.comparisonOperator : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dimensions"] = args ? args.dimensions : undefined;
            inputs["enable"] = args ? args.enable : undefined;
            inputs["evaluationCount"] = args ? args.evaluationCount : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["metricType"] = args ? args.metricType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            inputs["statistics"] = args ? args.statistics : undefined;
            inputs["threshold"] = args ? args.threshold : undefined;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Alarm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alarm resources.
 */
export interface AlarmState {
    /**
     * The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
     */
    readonly alarmActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
     */
    readonly cloudMonitorGroupId?: pulumi.Input<number>;
    /**
     * The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
     */
    readonly comparisonOperator?: pulumi.Input<string>;
    /**
     * The description for the alarm.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
     */
    readonly dimensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether to enable specific ess alarm. Default to true.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
     */
    readonly evaluationCount?: pulumi.Input<number>;
    /**
     * The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
     */
    readonly metricType?: pulumi.Input<string>;
    /**
     * The name for ess alarm.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
     */
    readonly scalingGroupId?: pulumi.Input<string>;
    /**
     * The state of specified alarm.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
     */
    readonly statistics?: pulumi.Input<string>;
    /**
     * The value against which the specified statistics is compared.
     */
    readonly threshold?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.
     */
    readonly alarmActions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.
     */
    readonly cloudMonitorGroupId?: pulumi.Input<number>;
    /**
     * The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.
     */
    readonly comparisonOperator?: pulumi.Input<string>;
    /**
     * The description for the alarm.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scalingGroup" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users.
     */
    readonly dimensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether to enable specific ess alarm. Default to true.
     */
    readonly enable?: pulumi.Input<boolean>;
    /**
     * The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.
     */
    readonly evaluationCount?: pulumi.Input<number>;
    /**
     * The name for the alarm's associated metric. See Block_metricNames_and_dimensions below for details.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.
     */
    readonly metricType?: pulumi.Input<string>;
    /**
     * The name for ess alarm.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.
     */
    readonly scalingGroupId: pulumi.Input<string>;
    /**
     * The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.
     */
    readonly statistics?: pulumi.Input<string>;
    /**
     * The value against which the specified statistics is compared.
     */
    readonly threshold: pulumi.Input<string>;
}
