// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource provides a alarm rule resource and it can be used to monitor several cloud services according different metrics.
// Details for [alarm rule](https://www.alibabacloud.com/help/doc-detail/28608.htm).
type Alarm struct {
	pulumi.CustomResourceState

	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayOutput `pulumi:"contactGroups"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Dimensions pulumi.MapOutput `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrOutput `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringOutput `pulumi:"metric"`
	// The alarm rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".
	Operator pulumi.StringPtrOutput `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringOutput `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrOutput `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".
	Statistics pulumi.StringPtrOutput `pulumi:"statistics"`
	// The current alarm rule status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Alarm threshold value, which must be a numeric value currently.
	Threshold pulumi.StringOutput `pulumi:"threshold"`
	// Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.
	TriggeredCount pulumi.IntPtrOutput `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrOutput `pulumi:"webhook"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil || args.ContactGroups == nil {
		return nil, errors.New("missing required argument 'ContactGroups'")
	}
	if args == nil || args.Dimensions == nil {
		return nil, errors.New("missing required argument 'Dimensions'")
	}
	if args == nil || args.Metric == nil {
		return nil, errors.New("missing required argument 'Metric'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Threshold == nil {
		return nil, errors.New("missing required argument 'Threshold'")
	}
	if args == nil {
		args = &AlarmArgs{}
	}
	var resource Alarm
	err := ctx.RegisterResource("alicloud:cms/alarm:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("alicloud:cms/alarm:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups []string `pulumi:"contactGroups"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled *bool `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	EndTime *int `pulumi:"endTime"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric *string `pulumi:"metric"`
	// The alarm rule name.
	Name *string `pulumi:"name"`
	// Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".
	Operator *string `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period *int `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project *string `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime *int `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	StartTime *int `pulumi:"startTime"`
	// Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".
	Statistics *string `pulumi:"statistics"`
	// The current alarm rule status.
	Status *string `pulumi:"status"`
	// Alarm threshold value, which must be a numeric value currently.
	Threshold *string `pulumi:"threshold"`
	// Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.
	TriggeredCount *int `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook *string `pulumi:"webhook"`
}

type AlarmState struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayInput
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Dimensions pulumi.MapInput
	// The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrInput
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	EndTime pulumi.IntPtrInput
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringPtrInput
	// The alarm rule name.
	Name pulumi.StringPtrInput
	// Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".
	Operator pulumi.StringPtrInput
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrInput
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringPtrInput
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	StartTime pulumi.IntPtrInput
	// Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".
	Statistics pulumi.StringPtrInput
	// The current alarm rule status.
	Status pulumi.StringPtrInput
	// Alarm threshold value, which must be a numeric value currently.
	Threshold pulumi.StringPtrInput
	// Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.
	TriggeredCount pulumi.IntPtrInput
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrInput
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups []string `pulumi:"contactGroups"`
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Dimensions map[string]interface{} `pulumi:"dimensions"`
	// The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval *string `pulumi:"effectiveInterval"`
	// Whether to enable alarm rule. Default to true.
	Enabled *bool `pulumi:"enabled"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	EndTime *int `pulumi:"endTime"`
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric string `pulumi:"metric"`
	// The alarm rule name.
	Name *string `pulumi:"name"`
	// Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".
	Operator *string `pulumi:"operator"`
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period *int `pulumi:"period"`
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project string `pulumi:"project"`
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime *int `pulumi:"silenceTime"`
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	StartTime *int `pulumi:"startTime"`
	// Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".
	Statistics *string `pulumi:"statistics"`
	// Alarm threshold value, which must be a numeric value currently.
	Threshold string `pulumi:"threshold"`
	// Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.
	TriggeredCount *int `pulumi:"triggeredCount"`
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook *string `pulumi:"webhook"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// List contact groups of the alarm rule, which must have been created on the console.
	ContactGroups pulumi.StringArrayInput
	// Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Dimensions pulumi.MapInput
	// The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
	EffectiveInterval pulumi.StringPtrInput
	// Whether to enable alarm rule. Default to true.
	Enabled pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	EndTime pulumi.IntPtrInput
	// Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkinRate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Metric pulumi.StringInput
	// The alarm rule name.
	Name pulumi.StringPtrInput
	// Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".
	Operator pulumi.StringPtrInput
	// Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
	Period pulumi.IntPtrInput
	// Monitor project name, such as "acsEcsDashboard" and "acsRdsDashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
	Project pulumi.StringInput
	// Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400
	SilenceTime pulumi.IntPtrInput
	// It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.
	StartTime pulumi.IntPtrInput
	// Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".
	Statistics pulumi.StringPtrInput
	// Alarm threshold value, which must be a numeric value currently.
	Threshold pulumi.StringInput
	// Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.
	TriggeredCount pulumi.IntPtrInput
	// The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
	Webhook pulumi.StringPtrInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}
