// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ess

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ScalingRule struct {
	pulumi.CustomResourceState

	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrOutput `pulumi:"adjustmentType"`
	// Adjusted value of a scaling rule. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrOutput `pulumi:"adjustmentValue"`
	Ari             pulumi.StringOutput `pulumi:"ari"`
	// Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrOutput `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrOutput `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntOutput `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrOutput `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringOutput `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrOutput `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayOutput `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue pulumi.Float64PtrOutput `pulumi:"targetValue"`
}

// NewScalingRule registers a new resource with the given unique name, arguments, and options.
func NewScalingRule(ctx *pulumi.Context,
	name string, args *ScalingRuleArgs, opts ...pulumi.ResourceOption) (*ScalingRule, error) {
	if args == nil || args.ScalingGroupId == nil {
		return nil, errors.New("missing required argument 'ScalingGroupId'")
	}
	if args == nil {
		args = &ScalingRuleArgs{}
	}
	var resource ScalingRule
	err := ctx.RegisterResource("alicloud:ess/scalingRule:ScalingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingRule gets an existing ScalingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingRuleState, opts ...pulumi.ResourceOption) (*ScalingRule, error) {
	var resource ScalingRule
	err := ctx.ReadResource("alicloud:ess/scalingRule:ScalingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingRule resources.
type scalingRuleState struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// Adjusted value of a scaling rule. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue *int    `pulumi:"adjustmentValue"`
	Ari             *string `pulumi:"ari"`
	// Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown *int `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn *bool `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName *string `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName *string `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType *string `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments []ScalingRuleStepAdjustment `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue *float64 `pulumi:"targetValue"`
}

type ScalingRuleState struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrInput
	// Adjusted value of a scaling rule. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrInput
	Ari             pulumi.StringPtrInput
	// Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrInput
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrInput
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringPtrInput
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringPtrInput
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrInput
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayInput
	// The target value for the metric.
	TargetValue pulumi.Float64PtrInput
}

func (ScalingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingRuleState)(nil)).Elem()
}

type scalingRuleArgs struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// Adjusted value of a scaling rule. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue *int `pulumi:"adjustmentValue"`
	// Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown *int `pulumi:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn *bool `pulumi:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// A CloudMonitor metric name.
	MetricName *string `pulumi:"metricName"`
	// ID of the scaling group of a scaling rule.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName *string `pulumi:"scalingRuleName"`
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType *string `pulumi:"scalingRuleType"`
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments []ScalingRuleStepAdjustment `pulumi:"stepAdjustments"`
	// The target value for the metric.
	TargetValue *float64 `pulumi:"targetValue"`
}

// The set of arguments for constructing a ScalingRule resource.
type ScalingRuleArgs struct {
	// Adjustment mode of a scaling rule. Optional values:
	// - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances.
	// - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances.
	// - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.
	AdjustmentType pulumi.StringPtrInput
	// Adjusted value of a scaling rule. Value range:
	// - QuantityChangeInCapacity：(0, 500] U (-500, 0]
	// - PercentChangeInCapacity：[0, 10000] U [-100, 0]
	// - TotalCapacity：[0, 1000]
	AdjustmentValue pulumi.IntPtrInput
	// Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.
	Cooldown pulumi.IntPtrInput
	// Indicates whether scale in by the target tracking policy is disabled. Default to false.
	DisableScaleIn pulumi.BoolPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// A CloudMonitor metric name.
	MetricName pulumi.StringPtrInput
	// ID of the scaling group of a scaling rule.
	ScalingGroupId pulumi.StringInput
	// Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is scaling rule id.
	ScalingRuleName pulumi.StringPtrInput
	// The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".
	ScalingRuleType pulumi.StringPtrInput
	// Steps for StepScalingRule. See Block stepAdjustment below for details.
	StepAdjustments ScalingRuleStepAdjustmentArrayInput
	// The target value for the metric.
	TargetValue pulumi.Float64PtrInput
}

func (ScalingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingRuleArgs)(nil)).Elem()
}
