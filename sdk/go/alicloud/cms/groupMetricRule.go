// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Group Metric Rule resource.
//
// For information about Cloud Monitor Service Group Metric Rule and how to use it, see [What is Group Metric Rule](https://www.alibabacloud.com/help/en/cloudmonitor/latest/putgroupmetricrule).
//
// > **NOTE:** Available since v1.104.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := cms.NewAlarmContactGroup(ctx, "default", &cms.AlarmContactGroupArgs{
//				AlarmContactGroupName: pulumi.String(name),
//				Describe:              pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultMonitorGroup, err := cms.NewMonitorGroup(ctx, "default", &cms.MonitorGroupArgs{
//				MonitorGroupName: pulumi.String(name),
//				ContactGroups: pulumi.StringArray{
//					_default.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cms.NewGroupMetricRule(ctx, "this", &cms.GroupMetricRuleArgs{
//				GroupId:             defaultMonitorGroup.ID(),
//				GroupMetricRuleName: pulumi.String(name),
//				Category:            pulumi.String("ecs"),
//				MetricName:          pulumi.String("cpu_total"),
//				Namespace:           pulumi.String("acs_ecs_dashboard"),
//				RuleId:              pulumi.String(name),
//				Period:              pulumi.Int(60),
//				Interval:            pulumi.String("3600"),
//				SilenceTime:         pulumi.Int(85800),
//				NoEffectiveInterval: pulumi.String("00:00-05:30"),
//				Webhook:             pulumi.String("http://www.aliyun.com"),
//				Escalations: &cms.GroupMetricRuleEscalationsArgs{
//					Warn: &cms.GroupMetricRuleEscalationsWarnArgs{
//						ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
//						Statistics:         pulumi.String("Average"),
//						Threshold:          pulumi.String("90"),
//						Times:              pulumi.Int(3),
//					},
//					Info: &cms.GroupMetricRuleEscalationsInfoArgs{
//						ComparisonOperator: pulumi.String("LessThanLastWeek"),
//						Statistics:         pulumi.String("Average"),
//						Threshold:          pulumi.String("90"),
//						Times:              pulumi.Int(5),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Cloud Monitor Service Group Metric Rule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cms/groupMetricRule:GroupMetricRule example <rule_id>
// ```
type GroupMetricRule struct {
	pulumi.CustomResourceState

	// The abbreviation of the service name.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Alarm contact group.
	ContactGroups pulumi.StringOutput `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringOutput `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrOutput `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.
	EmailSubject pulumi.StringOutput `pulumi:"emailSubject"`
	// Alarm level. See `escalations` below.
	Escalations GroupMetricRuleEscalationsOutput `pulumi:"escalations"`
	// The ID of the application group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringOutput `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrOutput `pulumi:"interval"`
	// The name of the metric.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The namespace of the service.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrOutput `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntOutput `pulumi:"period"`
	// The ID of the alert rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntOutput `pulumi:"silenceTime"`
	// The status of Group Metric Rule.
	Status pulumi.StringOutput `pulumi:"status"`
	// The information about the resource for which alerts are triggered. See `targets` below.
	Targets GroupMetricRuleTargetArrayOutput `pulumi:"targets"`
	// The callback URL.
	Webhook pulumi.StringPtrOutput `pulumi:"webhook"`
}

// NewGroupMetricRule registers a new resource with the given unique name, arguments, and options.
func NewGroupMetricRule(ctx *pulumi.Context,
	name string, args *GroupMetricRuleArgs, opts ...pulumi.ResourceOption) (*GroupMetricRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Escalations == nil {
		return nil, errors.New("invalid value for required argument 'Escalations'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.GroupMetricRuleName == nil {
		return nil, errors.New("invalid value for required argument 'GroupMetricRuleName'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMetricRule
	err := ctx.RegisterResource("alicloud:cms/groupMetricRule:GroupMetricRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMetricRule gets an existing GroupMetricRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMetricRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMetricRuleState, opts ...pulumi.ResourceOption) (*GroupMetricRule, error) {
	var resource GroupMetricRule
	err := ctx.ReadResource("alicloud:cms/groupMetricRule:GroupMetricRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMetricRule resources.
type groupMetricRuleState struct {
	// The abbreviation of the service name.
	Category *string `pulumi:"category"`
	// Alarm contact group.
	ContactGroups *string `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions *string `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.
	EmailSubject *string `pulumi:"emailSubject"`
	// Alarm level. See `escalations` below.
	Escalations *GroupMetricRuleEscalations `pulumi:"escalations"`
	// The ID of the application group.
	GroupId *string `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName *string `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval *string `pulumi:"interval"`
	// The name of the metric.
	MetricName *string `pulumi:"metricName"`
	// The namespace of the service.
	Namespace *string `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval *string `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period *int `pulumi:"period"`
	// The ID of the alert rule.
	RuleId *string `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime *int `pulumi:"silenceTime"`
	// The status of Group Metric Rule.
	Status *string `pulumi:"status"`
	// The information about the resource for which alerts are triggered. See `targets` below.
	Targets []GroupMetricRuleTarget `pulumi:"targets"`
	// The callback URL.
	Webhook *string `pulumi:"webhook"`
}

type GroupMetricRuleState struct {
	// The abbreviation of the service name.
	Category pulumi.StringPtrInput
	// Alarm contact group.
	ContactGroups pulumi.StringPtrInput
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringPtrInput
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrInput
	// The subject of the alert notification email.
	EmailSubject pulumi.StringPtrInput
	// Alarm level. See `escalations` below.
	Escalations GroupMetricRuleEscalationsPtrInput
	// The ID of the application group.
	GroupId pulumi.StringPtrInput
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringPtrInput
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrInput
	// The name of the metric.
	MetricName pulumi.StringPtrInput
	// The namespace of the service.
	Namespace pulumi.StringPtrInput
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrInput
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntPtrInput
	// The ID of the alert rule.
	RuleId pulumi.StringPtrInput
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntPtrInput
	// The status of Group Metric Rule.
	Status pulumi.StringPtrInput
	// The information about the resource for which alerts are triggered. See `targets` below.
	Targets GroupMetricRuleTargetArrayInput
	// The callback URL.
	Webhook pulumi.StringPtrInput
}

func (GroupMetricRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMetricRuleState)(nil)).Elem()
}

type groupMetricRuleArgs struct {
	// The abbreviation of the service name.
	Category *string `pulumi:"category"`
	// Alarm contact group.
	ContactGroups *string `pulumi:"contactGroups"`
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions *string `pulumi:"dimensions"`
	// The time period during which the alert rule is effective.
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// The subject of the alert notification email.
	EmailSubject *string `pulumi:"emailSubject"`
	// Alarm level. See `escalations` below.
	Escalations GroupMetricRuleEscalations `pulumi:"escalations"`
	// The ID of the application group.
	GroupId string `pulumi:"groupId"`
	// The name of the alert rule.
	GroupMetricRuleName string `pulumi:"groupMetricRuleName"`
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval *string `pulumi:"interval"`
	// The name of the metric.
	MetricName string `pulumi:"metricName"`
	// The namespace of the service.
	Namespace string `pulumi:"namespace"`
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval *string `pulumi:"noEffectiveInterval"`
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period *int `pulumi:"period"`
	// The ID of the alert rule.
	RuleId string `pulumi:"ruleId"`
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime *int `pulumi:"silenceTime"`
	// The information about the resource for which alerts are triggered. See `targets` below.
	Targets []GroupMetricRuleTarget `pulumi:"targets"`
	// The callback URL.
	Webhook *string `pulumi:"webhook"`
}

// The set of arguments for constructing a GroupMetricRule resource.
type GroupMetricRuleArgs struct {
	// The abbreviation of the service name.
	Category pulumi.StringPtrInput
	// Alarm contact group.
	ContactGroups pulumi.StringPtrInput
	// The dimensions that specify the resources to be associated with the alert rule.
	Dimensions pulumi.StringPtrInput
	// The time period during which the alert rule is effective.
	EffectiveInterval pulumi.StringPtrInput
	// The subject of the alert notification email.
	EmailSubject pulumi.StringPtrInput
	// Alarm level. See `escalations` below.
	Escalations GroupMetricRuleEscalationsInput
	// The ID of the application group.
	GroupId pulumi.StringInput
	// The name of the alert rule.
	GroupMetricRuleName pulumi.StringInput
	// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
	Interval pulumi.StringPtrInput
	// The name of the metric.
	MetricName pulumi.StringInput
	// The namespace of the service.
	Namespace pulumi.StringInput
	// The time period during which the alert rule is ineffective.
	NoEffectiveInterval pulumi.StringPtrInput
	// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
	Period pulumi.IntPtrInput
	// The ID of the alert rule.
	RuleId pulumi.StringInput
	// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
	SilenceTime pulumi.IntPtrInput
	// The information about the resource for which alerts are triggered. See `targets` below.
	Targets GroupMetricRuleTargetArrayInput
	// The callback URL.
	Webhook pulumi.StringPtrInput
}

func (GroupMetricRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMetricRuleArgs)(nil)).Elem()
}

type GroupMetricRuleInput interface {
	pulumi.Input

	ToGroupMetricRuleOutput() GroupMetricRuleOutput
	ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput
}

func (*GroupMetricRule) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMetricRule)(nil)).Elem()
}

func (i *GroupMetricRule) ToGroupMetricRuleOutput() GroupMetricRuleOutput {
	return i.ToGroupMetricRuleOutputWithContext(context.Background())
}

func (i *GroupMetricRule) ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMetricRuleOutput)
}

// GroupMetricRuleArrayInput is an input type that accepts GroupMetricRuleArray and GroupMetricRuleArrayOutput values.
// You can construct a concrete instance of `GroupMetricRuleArrayInput` via:
//
//	GroupMetricRuleArray{ GroupMetricRuleArgs{...} }
type GroupMetricRuleArrayInput interface {
	pulumi.Input

	ToGroupMetricRuleArrayOutput() GroupMetricRuleArrayOutput
	ToGroupMetricRuleArrayOutputWithContext(context.Context) GroupMetricRuleArrayOutput
}

type GroupMetricRuleArray []GroupMetricRuleInput

func (GroupMetricRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMetricRule)(nil)).Elem()
}

func (i GroupMetricRuleArray) ToGroupMetricRuleArrayOutput() GroupMetricRuleArrayOutput {
	return i.ToGroupMetricRuleArrayOutputWithContext(context.Background())
}

func (i GroupMetricRuleArray) ToGroupMetricRuleArrayOutputWithContext(ctx context.Context) GroupMetricRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMetricRuleArrayOutput)
}

// GroupMetricRuleMapInput is an input type that accepts GroupMetricRuleMap and GroupMetricRuleMapOutput values.
// You can construct a concrete instance of `GroupMetricRuleMapInput` via:
//
//	GroupMetricRuleMap{ "key": GroupMetricRuleArgs{...} }
type GroupMetricRuleMapInput interface {
	pulumi.Input

	ToGroupMetricRuleMapOutput() GroupMetricRuleMapOutput
	ToGroupMetricRuleMapOutputWithContext(context.Context) GroupMetricRuleMapOutput
}

type GroupMetricRuleMap map[string]GroupMetricRuleInput

func (GroupMetricRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMetricRule)(nil)).Elem()
}

func (i GroupMetricRuleMap) ToGroupMetricRuleMapOutput() GroupMetricRuleMapOutput {
	return i.ToGroupMetricRuleMapOutputWithContext(context.Background())
}

func (i GroupMetricRuleMap) ToGroupMetricRuleMapOutputWithContext(ctx context.Context) GroupMetricRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMetricRuleMapOutput)
}

type GroupMetricRuleOutput struct{ *pulumi.OutputState }

func (GroupMetricRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMetricRule)(nil)).Elem()
}

func (o GroupMetricRuleOutput) ToGroupMetricRuleOutput() GroupMetricRuleOutput {
	return o
}

func (o GroupMetricRuleOutput) ToGroupMetricRuleOutputWithContext(ctx context.Context) GroupMetricRuleOutput {
	return o
}

// The abbreviation of the service name.
func (o GroupMetricRuleOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// Alarm contact group.
func (o GroupMetricRuleOutput) ContactGroups() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.ContactGroups }).(pulumi.StringOutput)
}

// The dimensions that specify the resources to be associated with the alert rule.
func (o GroupMetricRuleOutput) Dimensions() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.Dimensions }).(pulumi.StringOutput)
}

// The time period during which the alert rule is effective.
func (o GroupMetricRuleOutput) EffectiveInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringPtrOutput { return v.EffectiveInterval }).(pulumi.StringPtrOutput)
}

// The subject of the alert notification email.
func (o GroupMetricRuleOutput) EmailSubject() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.EmailSubject }).(pulumi.StringOutput)
}

// Alarm level. See `escalations` below.
func (o GroupMetricRuleOutput) Escalations() GroupMetricRuleEscalationsOutput {
	return o.ApplyT(func(v *GroupMetricRule) GroupMetricRuleEscalationsOutput { return v.Escalations }).(GroupMetricRuleEscalationsOutput)
}

// The ID of the application group.
func (o GroupMetricRuleOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the alert rule.
func (o GroupMetricRuleOutput) GroupMetricRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.GroupMetricRuleName }).(pulumi.StringOutput)
}

// The interval at which Cloud Monitor checks whether the alert rule is triggered. Unit: seconds.
func (o GroupMetricRuleOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringPtrOutput { return v.Interval }).(pulumi.StringPtrOutput)
}

// The name of the metric.
func (o GroupMetricRuleOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

// The namespace of the service.
func (o GroupMetricRuleOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The time period during which the alert rule is ineffective.
func (o GroupMetricRuleOutput) NoEffectiveInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringPtrOutput { return v.NoEffectiveInterval }).(pulumi.StringPtrOutput)
}

// The aggregation period of the monitoring data. Unit: seconds. The value is an integral multiple of 60. Default value: `300`.
func (o GroupMetricRuleOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// The ID of the alert rule.
func (o GroupMetricRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// The mute period during which new alerts are not reported even if the alert trigger conditions are met. Unit: seconds. Default value: `86400`, which is equivalent to one day.
func (o GroupMetricRuleOutput) SilenceTime() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.IntOutput { return v.SilenceTime }).(pulumi.IntOutput)
}

// The status of Group Metric Rule.
func (o GroupMetricRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The information about the resource for which alerts are triggered. See `targets` below.
func (o GroupMetricRuleOutput) Targets() GroupMetricRuleTargetArrayOutput {
	return o.ApplyT(func(v *GroupMetricRule) GroupMetricRuleTargetArrayOutput { return v.Targets }).(GroupMetricRuleTargetArrayOutput)
}

// The callback URL.
func (o GroupMetricRuleOutput) Webhook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMetricRule) pulumi.StringPtrOutput { return v.Webhook }).(pulumi.StringPtrOutput)
}

type GroupMetricRuleArrayOutput struct{ *pulumi.OutputState }

func (GroupMetricRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMetricRule)(nil)).Elem()
}

func (o GroupMetricRuleArrayOutput) ToGroupMetricRuleArrayOutput() GroupMetricRuleArrayOutput {
	return o
}

func (o GroupMetricRuleArrayOutput) ToGroupMetricRuleArrayOutputWithContext(ctx context.Context) GroupMetricRuleArrayOutput {
	return o
}

func (o GroupMetricRuleArrayOutput) Index(i pulumi.IntInput) GroupMetricRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMetricRule {
		return vs[0].([]*GroupMetricRule)[vs[1].(int)]
	}).(GroupMetricRuleOutput)
}

type GroupMetricRuleMapOutput struct{ *pulumi.OutputState }

func (GroupMetricRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMetricRule)(nil)).Elem()
}

func (o GroupMetricRuleMapOutput) ToGroupMetricRuleMapOutput() GroupMetricRuleMapOutput {
	return o
}

func (o GroupMetricRuleMapOutput) ToGroupMetricRuleMapOutputWithContext(ctx context.Context) GroupMetricRuleMapOutput {
	return o
}

func (o GroupMetricRuleMapOutput) MapIndex(k pulumi.StringInput) GroupMetricRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMetricRule {
		return vs[0].(map[string]*GroupMetricRule)[vs[1].(string)]
	}).(GroupMetricRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMetricRuleInput)(nil)).Elem(), &GroupMetricRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMetricRuleArrayInput)(nil)).Elem(), GroupMetricRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMetricRuleMapInput)(nil)).Elem(), GroupMetricRuleMap{})
	pulumi.RegisterOutputType(GroupMetricRuleOutput{})
	pulumi.RegisterOutputType(GroupMetricRuleArrayOutput{})
	pulumi.RegisterOutputType(GroupMetricRuleMapOutput{})
}
