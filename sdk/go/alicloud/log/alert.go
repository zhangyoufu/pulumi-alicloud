// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package log

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Log alert is a unit of log service, which is used to monitor and alert the user's logstore status information.
// Log Service enables you to configure alerts based on the charts in a dashboard to monitor the service status in real time.
//
// > **NOTE:** Available in 1.78.0
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/log_alert.html.markdown.
type Alert struct {
	pulumi.CustomResourceState

	// Alert description.
	AlertDescription pulumi.StringPtrOutput `pulumi:"alertDescription"`
	// Alert displayname.
	AlertDisplayname pulumi.StringOutput `pulumi:"alertDisplayname"`
	// Name of logstore for configuring alarm service.
	AlertName pulumi.StringOutput `pulumi:"alertName"`
	// Conditional expression, such as: count> 100.
	Condition pulumi.StringOutput `pulumi:"condition"`
	Dashboard pulumi.StringOutput `pulumi:"dashboard"`
	// Timestamp, notifications before closing again.
	MuteUntil pulumi.IntPtrOutput `pulumi:"muteUntil"`
	// Alarm information notification list.
	NotificationLists AlertNotificationListArrayOutput `pulumi:"notificationLists"`
	// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
	NotifyThreshold pulumi.IntPtrOutput `pulumi:"notifyThreshold"`
	// The project name.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Multiple conditions for configured alarm query.
	QueryLists AlertQueryListArrayOutput `pulumi:"queryLists"`
	// Execution interval. 60 seconds minimum, such as 60s, 1h.
	ScheduleInterval pulumi.StringPtrOutput `pulumi:"scheduleInterval"`
	// Default FixedRate. No need to configure this parameter.
	ScheduleType pulumi.StringPtrOutput `pulumi:"scheduleType"`
	// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
	Throttling pulumi.StringPtrOutput `pulumi:"throttling"`
}

// NewAlert registers a new resource with the given unique name, arguments, and options.
func NewAlert(ctx *pulumi.Context,
	name string, args *AlertArgs, opts ...pulumi.ResourceOption) (*Alert, error) {
	if args == nil || args.AlertDisplayname == nil {
		return nil, errors.New("missing required argument 'AlertDisplayname'")
	}
	if args == nil || args.AlertName == nil {
		return nil, errors.New("missing required argument 'AlertName'")
	}
	if args == nil || args.Condition == nil {
		return nil, errors.New("missing required argument 'Condition'")
	}
	if args == nil || args.Dashboard == nil {
		return nil, errors.New("missing required argument 'Dashboard'")
	}
	if args == nil || args.NotificationLists == nil {
		return nil, errors.New("missing required argument 'NotificationLists'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.QueryLists == nil {
		return nil, errors.New("missing required argument 'QueryLists'")
	}
	if args == nil {
		args = &AlertArgs{}
	}
	var resource Alert
	err := ctx.RegisterResource("alicloud:log/alert:Alert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlert gets an existing Alert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertState, opts ...pulumi.ResourceOption) (*Alert, error) {
	var resource Alert
	err := ctx.ReadResource("alicloud:log/alert:Alert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alert resources.
type alertState struct {
	// Alert description.
	AlertDescription *string `pulumi:"alertDescription"`
	// Alert displayname.
	AlertDisplayname *string `pulumi:"alertDisplayname"`
	// Name of logstore for configuring alarm service.
	AlertName *string `pulumi:"alertName"`
	// Conditional expression, such as: count> 100.
	Condition *string `pulumi:"condition"`
	Dashboard *string `pulumi:"dashboard"`
	// Timestamp, notifications before closing again.
	MuteUntil *int `pulumi:"muteUntil"`
	// Alarm information notification list.
	NotificationLists []AlertNotificationList `pulumi:"notificationLists"`
	// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
	NotifyThreshold *int `pulumi:"notifyThreshold"`
	// The project name.
	ProjectName *string `pulumi:"projectName"`
	// Multiple conditions for configured alarm query.
	QueryLists []AlertQueryList `pulumi:"queryLists"`
	// Execution interval. 60 seconds minimum, such as 60s, 1h.
	ScheduleInterval *string `pulumi:"scheduleInterval"`
	// Default FixedRate. No need to configure this parameter.
	ScheduleType *string `pulumi:"scheduleType"`
	// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
	Throttling *string `pulumi:"throttling"`
}

type AlertState struct {
	// Alert description.
	AlertDescription pulumi.StringPtrInput
	// Alert displayname.
	AlertDisplayname pulumi.StringPtrInput
	// Name of logstore for configuring alarm service.
	AlertName pulumi.StringPtrInput
	// Conditional expression, such as: count> 100.
	Condition pulumi.StringPtrInput
	Dashboard pulumi.StringPtrInput
	// Timestamp, notifications before closing again.
	MuteUntil pulumi.IntPtrInput
	// Alarm information notification list.
	NotificationLists AlertNotificationListArrayInput
	// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
	NotifyThreshold pulumi.IntPtrInput
	// The project name.
	ProjectName pulumi.StringPtrInput
	// Multiple conditions for configured alarm query.
	QueryLists AlertQueryListArrayInput
	// Execution interval. 60 seconds minimum, such as 60s, 1h.
	ScheduleInterval pulumi.StringPtrInput
	// Default FixedRate. No need to configure this parameter.
	ScheduleType pulumi.StringPtrInput
	// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
	Throttling pulumi.StringPtrInput
}

func (AlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertState)(nil)).Elem()
}

type alertArgs struct {
	// Alert description.
	AlertDescription *string `pulumi:"alertDescription"`
	// Alert displayname.
	AlertDisplayname string `pulumi:"alertDisplayname"`
	// Name of logstore for configuring alarm service.
	AlertName string `pulumi:"alertName"`
	// Conditional expression, such as: count> 100.
	Condition string `pulumi:"condition"`
	Dashboard string `pulumi:"dashboard"`
	// Timestamp, notifications before closing again.
	MuteUntil *int `pulumi:"muteUntil"`
	// Alarm information notification list.
	NotificationLists []AlertNotificationList `pulumi:"notificationLists"`
	// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
	NotifyThreshold *int `pulumi:"notifyThreshold"`
	// The project name.
	ProjectName string `pulumi:"projectName"`
	// Multiple conditions for configured alarm query.
	QueryLists []AlertQueryList `pulumi:"queryLists"`
	// Execution interval. 60 seconds minimum, such as 60s, 1h.
	ScheduleInterval *string `pulumi:"scheduleInterval"`
	// Default FixedRate. No need to configure this parameter.
	ScheduleType *string `pulumi:"scheduleType"`
	// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
	Throttling *string `pulumi:"throttling"`
}

// The set of arguments for constructing a Alert resource.
type AlertArgs struct {
	// Alert description.
	AlertDescription pulumi.StringPtrInput
	// Alert displayname.
	AlertDisplayname pulumi.StringInput
	// Name of logstore for configuring alarm service.
	AlertName pulumi.StringInput
	// Conditional expression, such as: count> 100.
	Condition pulumi.StringInput
	Dashboard pulumi.StringInput
	// Timestamp, notifications before closing again.
	MuteUntil pulumi.IntPtrInput
	// Alarm information notification list.
	NotificationLists AlertNotificationListArrayInput
	// Notification threshold, which is not notified until the number of triggers is reached. The default is 1.
	NotifyThreshold pulumi.IntPtrInput
	// The project name.
	ProjectName pulumi.StringInput
	// Multiple conditions for configured alarm query.
	QueryLists AlertQueryListArrayInput
	// Execution interval. 60 seconds minimum, such as 60s, 1h.
	ScheduleInterval pulumi.StringPtrInput
	// Default FixedRate. No need to configure this parameter.
	ScheduleType pulumi.StringPtrInput
	// Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.
	Throttling pulumi.StringPtrInput
}

func (AlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertArgs)(nil)).Elem()
}
