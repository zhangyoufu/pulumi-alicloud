// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a ESS scaling rule resource.
 *
 * For information about ess scaling rule, see [CreateScalingRule](https://www.alibabacloud.com/help/en/auto-scaling/latest/createscalingrule).
 *
 * > **NOTE:** Available since v1.39.0.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultInteger = new random.index.Integer("default", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const myName = `${name}-${defaultInteger.result}`;
 * const default = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultGetInstanceTypes = _default.then(_default => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: _default.zones?.[0]?.id,
 *     cpuCoreCount: 2,
 *     memorySize: 4,
 * }));
 * const defaultGetImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_18.*64",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     vpcName: myName,
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: _default.then(_default => _default.zones?.[0]?.id),
 *     vswitchName: myName,
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     name: myName,
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultSecurityGroupRule = new alicloud.ecs.SecurityGroupRule("default", {
 *     type: "ingress",
 *     ipProtocol: "tcp",
 *     nicType: "intranet",
 *     policy: "accept",
 *     portRange: "22/22",
 *     priority: 1,
 *     securityGroupId: defaultSecurityGroup.id,
 *     cidrIp: "172.16.0.0/24",
 * });
 * const defaultScalingGroup = new alicloud.ess.ScalingGroup("default", {
 *     minSize: 1,
 *     maxSize: 1,
 *     scalingGroupName: myName,
 *     vswitchIds: [defaultSwitch.id],
 *     removalPolicies: [
 *         "OldestInstance",
 *         "NewestInstance",
 *     ],
 * });
 * const defaultScalingConfiguration = new alicloud.ess.ScalingConfiguration("default", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     imageId: defaultGetImages.then(defaultGetImages => defaultGetImages.images?.[0]?.id),
 *     instanceType: defaultGetInstanceTypes.then(defaultGetInstanceTypes => defaultGetInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroupId: defaultSecurityGroup.id,
 *     forceDelete: true,
 * });
 * const defaultScalingRule = new alicloud.ess.ScalingRule("default", {
 *     scalingGroupId: defaultScalingGroup.id,
 *     adjustmentType: "TotalCapacity",
 *     adjustmentValue: 1,
 * });
 * ```
 *
 * ## Module Support
 *
 * You can use to the existing autoscaling-rule module
 * to create different type rules, alarm task and scheduled task one-click.
 *
 * ## Import
 *
 * ESS scaling rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ess/scalingRule:ScalingRule example abc123456
 * ```
 */
export class ScalingRule extends pulumi.CustomResource {
    /**
     * Get an existing ScalingRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingRuleState, opts?: pulumi.CustomResourceOptions): ScalingRule {
        return new ScalingRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ess/scalingRule:ScalingRule';

    /**
     * Returns true if the given object is an instance of ScalingRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingRule.__pulumiType;
    }

    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    public readonly adjustmentType!: pulumi.Output<string | undefined>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    public readonly adjustmentValue!: pulumi.Output<number | undefined>;
    /**
     * AlarmDimension for StepScalingRule. See `alarmDimension` below.
     */
    public readonly alarmDimension!: pulumi.Output<outputs.ess.ScalingRuleAlarmDimension | undefined>;
    /**
     * The unique identifier of the scaling rule.
     */
    public /*out*/ readonly ari!: pulumi.Output<string>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    public readonly cooldown!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    public readonly disableScaleIn!: pulumi.Output<boolean | undefined>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    public readonly estimatedInstanceWarmup!: pulumi.Output<number>;
    /**
     * The maximum number of ECS instances that can be added to the scaling group. If you specify InitialMaxSize, you must also specify PredictiveValueBehavior.
     */
    public readonly initialMaxSize!: pulumi.Output<number>;
    /**
     * A CloudMonitor metric name.
     */
    public readonly metricName!: pulumi.Output<string | undefined>;
    /**
     * The minimum number of instances that must be scaled. This parameter takes effect if you set ScalingRuleType to SimpleScalingRule or StepScalingRule, and AdjustmentType to PercentChangeInCapacity.
     */
    public readonly minAdjustmentMagnitude!: pulumi.Output<number | undefined>;
    /**
     * The mode of the predictive scaling rule. Valid values: PredictAndScale, PredictOnly.
     */
    public readonly predictiveScalingMode!: pulumi.Output<string>;
    /**
     * The amount of buffer time before the prediction task runs. By default, all prediction tasks that are automatically created by a predictive scaling rule run on the hour. You can specify a buffer time to run prediction tasks and prepare resources in advance. Valid values: 0 to 60. Unit: minutes.
     */
    public readonly predictiveTaskBufferTime!: pulumi.Output<number>;
    /**
     * The action on the predicted maximum value. Valid values: MaxOverridePredictiveValue, PredictiveValueOverrideMax, PredictiveValueOverrideMaxWithBuffer.
     */
    public readonly predictiveValueBehavior!: pulumi.Output<string>;
    /**
     * The ratio based on which the predicted value is increased if you set PredictiveValueBehavior to PredictiveValueOverrideMaxWithBuffer. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
     */
    public readonly predictiveValueBuffer!: pulumi.Output<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-ins must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    public readonly scaleInEvaluationCount!: pulumi.Output<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-outs must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    public readonly scaleOutEvaluationCount!: pulumi.Output<number>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    public readonly scalingRuleName!: pulumi.Output<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule", "PredictiveScalingRule". Default to "SimpleScalingRule".
     */
    public readonly scalingRuleType!: pulumi.Output<string | undefined>;
    /**
     * Steps for StepScalingRule. See `stepAdjustment` below.
     */
    public readonly stepAdjustments!: pulumi.Output<outputs.ess.ScalingRuleStepAdjustment[] | undefined>;
    /**
     * The target value for the metric.
     */
    public readonly targetValue!: pulumi.Output<number | undefined>;

    /**
     * Create a ScalingRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingRuleArgs | ScalingRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingRuleState | undefined;
            resourceInputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            resourceInputs["adjustmentValue"] = state ? state.adjustmentValue : undefined;
            resourceInputs["alarmDimension"] = state ? state.alarmDimension : undefined;
            resourceInputs["ari"] = state ? state.ari : undefined;
            resourceInputs["cooldown"] = state ? state.cooldown : undefined;
            resourceInputs["disableScaleIn"] = state ? state.disableScaleIn : undefined;
            resourceInputs["estimatedInstanceWarmup"] = state ? state.estimatedInstanceWarmup : undefined;
            resourceInputs["initialMaxSize"] = state ? state.initialMaxSize : undefined;
            resourceInputs["metricName"] = state ? state.metricName : undefined;
            resourceInputs["minAdjustmentMagnitude"] = state ? state.minAdjustmentMagnitude : undefined;
            resourceInputs["predictiveScalingMode"] = state ? state.predictiveScalingMode : undefined;
            resourceInputs["predictiveTaskBufferTime"] = state ? state.predictiveTaskBufferTime : undefined;
            resourceInputs["predictiveValueBehavior"] = state ? state.predictiveValueBehavior : undefined;
            resourceInputs["predictiveValueBuffer"] = state ? state.predictiveValueBuffer : undefined;
            resourceInputs["scaleInEvaluationCount"] = state ? state.scaleInEvaluationCount : undefined;
            resourceInputs["scaleOutEvaluationCount"] = state ? state.scaleOutEvaluationCount : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            resourceInputs["scalingRuleName"] = state ? state.scalingRuleName : undefined;
            resourceInputs["scalingRuleType"] = state ? state.scalingRuleType : undefined;
            resourceInputs["stepAdjustments"] = state ? state.stepAdjustments : undefined;
            resourceInputs["targetValue"] = state ? state.targetValue : undefined;
        } else {
            const args = argsOrState as ScalingRuleArgs | undefined;
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            resourceInputs["adjustmentValue"] = args ? args.adjustmentValue : undefined;
            resourceInputs["alarmDimension"] = args ? args.alarmDimension : undefined;
            resourceInputs["cooldown"] = args ? args.cooldown : undefined;
            resourceInputs["disableScaleIn"] = args ? args.disableScaleIn : undefined;
            resourceInputs["estimatedInstanceWarmup"] = args ? args.estimatedInstanceWarmup : undefined;
            resourceInputs["initialMaxSize"] = args ? args.initialMaxSize : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["minAdjustmentMagnitude"] = args ? args.minAdjustmentMagnitude : undefined;
            resourceInputs["predictiveScalingMode"] = args ? args.predictiveScalingMode : undefined;
            resourceInputs["predictiveTaskBufferTime"] = args ? args.predictiveTaskBufferTime : undefined;
            resourceInputs["predictiveValueBehavior"] = args ? args.predictiveValueBehavior : undefined;
            resourceInputs["predictiveValueBuffer"] = args ? args.predictiveValueBuffer : undefined;
            resourceInputs["scaleInEvaluationCount"] = args ? args.scaleInEvaluationCount : undefined;
            resourceInputs["scaleOutEvaluationCount"] = args ? args.scaleOutEvaluationCount : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["scalingRuleName"] = args ? args.scalingRuleName : undefined;
            resourceInputs["scalingRuleType"] = args ? args.scalingRuleType : undefined;
            resourceInputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            resourceInputs["targetValue"] = args ? args.targetValue : undefined;
            resourceInputs["ari"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingRule resources.
 */
export interface ScalingRuleState {
    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    adjustmentValue?: pulumi.Input<number>;
    /**
     * AlarmDimension for StepScalingRule. See `alarmDimension` below.
     */
    alarmDimension?: pulumi.Input<inputs.ess.ScalingRuleAlarmDimension>;
    /**
     * The unique identifier of the scaling rule.
     */
    ari?: pulumi.Input<string>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    disableScaleIn?: pulumi.Input<boolean>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * The maximum number of ECS instances that can be added to the scaling group. If you specify InitialMaxSize, you must also specify PredictiveValueBehavior.
     */
    initialMaxSize?: pulumi.Input<number>;
    /**
     * A CloudMonitor metric name.
     */
    metricName?: pulumi.Input<string>;
    /**
     * The minimum number of instances that must be scaled. This parameter takes effect if you set ScalingRuleType to SimpleScalingRule or StepScalingRule, and AdjustmentType to PercentChangeInCapacity.
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The mode of the predictive scaling rule. Valid values: PredictAndScale, PredictOnly.
     */
    predictiveScalingMode?: pulumi.Input<string>;
    /**
     * The amount of buffer time before the prediction task runs. By default, all prediction tasks that are automatically created by a predictive scaling rule run on the hour. You can specify a buffer time to run prediction tasks and prepare resources in advance. Valid values: 0 to 60. Unit: minutes.
     */
    predictiveTaskBufferTime?: pulumi.Input<number>;
    /**
     * The action on the predicted maximum value. Valid values: MaxOverridePredictiveValue, PredictiveValueOverrideMax, PredictiveValueOverrideMaxWithBuffer.
     */
    predictiveValueBehavior?: pulumi.Input<string>;
    /**
     * The ratio based on which the predicted value is increased if you set PredictiveValueBehavior to PredictiveValueOverrideMaxWithBuffer. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
     */
    predictiveValueBuffer?: pulumi.Input<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-ins must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    scaleInEvaluationCount?: pulumi.Input<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-outs must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    scaleOutEvaluationCount?: pulumi.Input<number>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    scalingRuleName?: pulumi.Input<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule", "PredictiveScalingRule". Default to "SimpleScalingRule".
     */
    scalingRuleType?: pulumi.Input<string>;
    /**
     * Steps for StepScalingRule. See `stepAdjustment` below.
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.ess.ScalingRuleStepAdjustment>[]>;
    /**
     * The target value for the metric.
     */
    targetValue?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ScalingRule resource.
 */
export interface ScalingRuleArgs {
    /**
     * Adjustment mode of a scaling rule. Optional values:
     * - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
     * - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
     * - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range:
     * - QuantityChangeInCapacity：(0, 500] U (-500, 0]
     * - PercentChangeInCapacity：[0, 10000] U [-100, 0]
     * - TotalCapacity：[0, 1000]
     */
    adjustmentValue?: pulumi.Input<number>;
    /**
     * AlarmDimension for StepScalingRule. See `alarmDimension` below.
     */
    alarmDimension?: pulumi.Input<inputs.ess.ScalingRuleAlarmDimension>;
    /**
     * The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Indicates whether scale in by the target tracking policy is disabled. Default to false.
     */
    disableScaleIn?: pulumi.Input<boolean>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * The maximum number of ECS instances that can be added to the scaling group. If you specify InitialMaxSize, you must also specify PredictiveValueBehavior.
     */
    initialMaxSize?: pulumi.Input<number>;
    /**
     * A CloudMonitor metric name.
     */
    metricName?: pulumi.Input<string>;
    /**
     * The minimum number of instances that must be scaled. This parameter takes effect if you set ScalingRuleType to SimpleScalingRule or StepScalingRule, and AdjustmentType to PercentChangeInCapacity.
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The mode of the predictive scaling rule. Valid values: PredictAndScale, PredictOnly.
     */
    predictiveScalingMode?: pulumi.Input<string>;
    /**
     * The amount of buffer time before the prediction task runs. By default, all prediction tasks that are automatically created by a predictive scaling rule run on the hour. You can specify a buffer time to run prediction tasks and prepare resources in advance. Valid values: 0 to 60. Unit: minutes.
     */
    predictiveTaskBufferTime?: pulumi.Input<number>;
    /**
     * The action on the predicted maximum value. Valid values: MaxOverridePredictiveValue, PredictiveValueOverrideMax, PredictiveValueOverrideMaxWithBuffer.
     */
    predictiveValueBehavior?: pulumi.Input<string>;
    /**
     * The ratio based on which the predicted value is increased if you set PredictiveValueBehavior to PredictiveValueOverrideMaxWithBuffer. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
     */
    predictiveValueBuffer?: pulumi.Input<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-ins must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    scaleInEvaluationCount?: pulumi.Input<number>;
    /**
     * The number of consecutive times that the event-triggered task created for scale-outs must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
     */
    scaleOutEvaluationCount?: pulumi.Input<number>;
    /**
     * ID of the scaling group of a scaling rule.
     */
    scalingGroupId: pulumi.Input<string>;
    /**
     * Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
     */
    scalingRuleName?: pulumi.Input<string>;
    /**
     * The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule", "PredictiveScalingRule". Default to "SimpleScalingRule".
     */
    scalingRuleType?: pulumi.Input<string>;
    /**
     * Steps for StepScalingRule. See `stepAdjustment` below.
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.ess.ScalingRuleStepAdjustment>[]>;
    /**
     * The target value for the metric.
     */
    targetValue?: pulumi.Input<number>;
}
