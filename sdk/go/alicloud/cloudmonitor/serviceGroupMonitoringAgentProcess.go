// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudmonitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Group Monitoring Agent Process resource.
//
// For information about Cloud Monitor Service Group Monitoring Agent Process and how to use it, see [What is Group Monitoring Agent Process](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-creategroupmonitoringagentprocess).
//
// > **NOTE:** Available since v1.212.0.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudmonitor"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultAlarmContactGroup, err := cms.NewAlarmContactGroup(ctx, "defaultAlarmContactGroup", &cms.AlarmContactGroupArgs{
//				AlarmContactGroupName: pulumi.String(name),
//				Contacts: pulumi.StringArray{
//					pulumi.String("user"),
//					pulumi.String("user1"),
//					pulumi.String("user2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultMonitorGroup, err := cms.NewMonitorGroup(ctx, "defaultMonitorGroup", &cms.MonitorGroupArgs{
//				MonitorGroupName: pulumi.String(name),
//				ContactGroups: pulumi.StringArray{
//					defaultAlarmContactGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudmonitor.NewServiceGroupMonitoringAgentProcess(ctx, "defaultServiceGroupMonitoringAgentProcess", &cloudmonitor.ServiceGroupMonitoringAgentProcessArgs{
//				GroupId:                    defaultMonitorGroup.ID(),
//				ProcessName:                pulumi.String(name),
//				MatchExpressFilterRelation: pulumi.String("or"),
//				MatchExpresses: cloudmonitor.ServiceGroupMonitoringAgentProcessMatchExpressArray{
//					&cloudmonitor.ServiceGroupMonitoringAgentProcessMatchExpressArgs{
//						Name:     pulumi.String(name),
//						Value:    pulumi.String("*"),
//						Function: pulumi.String("all"),
//					},
//				},
//				AlertConfigs: cloudmonitor.ServiceGroupMonitoringAgentProcessAlertConfigArray{
//					&cloudmonitor.ServiceGroupMonitoringAgentProcessAlertConfigArgs{
//						EscalationsLevel:   pulumi.String("critical"),
//						ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
//						Statistics:         pulumi.String("Average"),
//						Threshold:          pulumi.String("20"),
//						Times:              pulumi.String("100"),
//						EffectiveInterval:  pulumi.String("00:00-22:59"),
//						SilenceTime:        pulumi.Int(85800),
//						Webhook:            pulumi.String("https://www.aliyun.com"),
//						TargetLists: cloudmonitor.ServiceGroupMonitoringAgentProcessAlertConfigTargetListArray{
//							&cloudmonitor.ServiceGroupMonitoringAgentProcessAlertConfigTargetListArgs{
//								TargetListId: pulumi.String("1"),
//								JsonParams:   pulumi.String("{}"),
//								Level:        pulumi.String("WARN"),
//								Arn:          pulumi.String("acs:mns:cn-hangzhou:120886317861****:/queues/test123/message"),
//							},
//						},
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
//
// ## Import
//
// Cloud Monitor Service Group Monitoring Agent Process can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess example <group_id>:<group_monitoring_agent_process_id>
// ```
type ServiceGroupMonitoringAgentProcess struct {
	pulumi.CustomResourceState

	// The alert rule configurations. See `alertConfig` below.
	AlertConfigs ServiceGroupMonitoringAgentProcessAlertConfigArrayOutput `pulumi:"alertConfigs"`
	// The ID of the application group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of the Group Monitoring Agent Process.
	GroupMonitoringAgentProcessId pulumi.StringOutput `pulumi:"groupMonitoringAgentProcessId"`
	// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
	MatchExpressFilterRelation pulumi.StringOutput `pulumi:"matchExpressFilterRelation"`
	// The expressions used to match instances. See `matchExpress` below.
	MatchExpresses ServiceGroupMonitoringAgentProcessMatchExpressArrayOutput `pulumi:"matchExpresses"`
	// The name of the process.
	ProcessName pulumi.StringOutput `pulumi:"processName"`
}

// NewServiceGroupMonitoringAgentProcess registers a new resource with the given unique name, arguments, and options.
func NewServiceGroupMonitoringAgentProcess(ctx *pulumi.Context,
	name string, args *ServiceGroupMonitoringAgentProcessArgs, opts ...pulumi.ResourceOption) (*ServiceGroupMonitoringAgentProcess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AlertConfigs'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.ProcessName == nil {
		return nil, errors.New("invalid value for required argument 'ProcessName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceGroupMonitoringAgentProcess
	err := ctx.RegisterResource("alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceGroupMonitoringAgentProcess gets an existing ServiceGroupMonitoringAgentProcess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceGroupMonitoringAgentProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceGroupMonitoringAgentProcessState, opts ...pulumi.ResourceOption) (*ServiceGroupMonitoringAgentProcess, error) {
	var resource ServiceGroupMonitoringAgentProcess
	err := ctx.ReadResource("alicloud:cloudmonitor/serviceGroupMonitoringAgentProcess:ServiceGroupMonitoringAgentProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceGroupMonitoringAgentProcess resources.
type serviceGroupMonitoringAgentProcessState struct {
	// The alert rule configurations. See `alertConfig` below.
	AlertConfigs []ServiceGroupMonitoringAgentProcessAlertConfig `pulumi:"alertConfigs"`
	// The ID of the application group.
	GroupId *string `pulumi:"groupId"`
	// The ID of the Group Monitoring Agent Process.
	GroupMonitoringAgentProcessId *string `pulumi:"groupMonitoringAgentProcessId"`
	// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
	MatchExpressFilterRelation *string `pulumi:"matchExpressFilterRelation"`
	// The expressions used to match instances. See `matchExpress` below.
	MatchExpresses []ServiceGroupMonitoringAgentProcessMatchExpress `pulumi:"matchExpresses"`
	// The name of the process.
	ProcessName *string `pulumi:"processName"`
}

type ServiceGroupMonitoringAgentProcessState struct {
	// The alert rule configurations. See `alertConfig` below.
	AlertConfigs ServiceGroupMonitoringAgentProcessAlertConfigArrayInput
	// The ID of the application group.
	GroupId pulumi.StringPtrInput
	// The ID of the Group Monitoring Agent Process.
	GroupMonitoringAgentProcessId pulumi.StringPtrInput
	// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
	MatchExpressFilterRelation pulumi.StringPtrInput
	// The expressions used to match instances. See `matchExpress` below.
	MatchExpresses ServiceGroupMonitoringAgentProcessMatchExpressArrayInput
	// The name of the process.
	ProcessName pulumi.StringPtrInput
}

func (ServiceGroupMonitoringAgentProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceGroupMonitoringAgentProcessState)(nil)).Elem()
}

type serviceGroupMonitoringAgentProcessArgs struct {
	// The alert rule configurations. See `alertConfig` below.
	AlertConfigs []ServiceGroupMonitoringAgentProcessAlertConfig `pulumi:"alertConfigs"`
	// The ID of the application group.
	GroupId string `pulumi:"groupId"`
	// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
	MatchExpressFilterRelation *string `pulumi:"matchExpressFilterRelation"`
	// The expressions used to match instances. See `matchExpress` below.
	MatchExpresses []ServiceGroupMonitoringAgentProcessMatchExpress `pulumi:"matchExpresses"`
	// The name of the process.
	ProcessName string `pulumi:"processName"`
}

// The set of arguments for constructing a ServiceGroupMonitoringAgentProcess resource.
type ServiceGroupMonitoringAgentProcessArgs struct {
	// The alert rule configurations. See `alertConfig` below.
	AlertConfigs ServiceGroupMonitoringAgentProcessAlertConfigArrayInput
	// The ID of the application group.
	GroupId pulumi.StringInput
	// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
	MatchExpressFilterRelation pulumi.StringPtrInput
	// The expressions used to match instances. See `matchExpress` below.
	MatchExpresses ServiceGroupMonitoringAgentProcessMatchExpressArrayInput
	// The name of the process.
	ProcessName pulumi.StringInput
}

func (ServiceGroupMonitoringAgentProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceGroupMonitoringAgentProcessArgs)(nil)).Elem()
}

type ServiceGroupMonitoringAgentProcessInput interface {
	pulumi.Input

	ToServiceGroupMonitoringAgentProcessOutput() ServiceGroupMonitoringAgentProcessOutput
	ToServiceGroupMonitoringAgentProcessOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessOutput
}

func (*ServiceGroupMonitoringAgentProcess) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (i *ServiceGroupMonitoringAgentProcess) ToServiceGroupMonitoringAgentProcessOutput() ServiceGroupMonitoringAgentProcessOutput {
	return i.ToServiceGroupMonitoringAgentProcessOutputWithContext(context.Background())
}

func (i *ServiceGroupMonitoringAgentProcess) ToServiceGroupMonitoringAgentProcessOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupMonitoringAgentProcessOutput)
}

// ServiceGroupMonitoringAgentProcessArrayInput is an input type that accepts ServiceGroupMonitoringAgentProcessArray and ServiceGroupMonitoringAgentProcessArrayOutput values.
// You can construct a concrete instance of `ServiceGroupMonitoringAgentProcessArrayInput` via:
//
//	ServiceGroupMonitoringAgentProcessArray{ ServiceGroupMonitoringAgentProcessArgs{...} }
type ServiceGroupMonitoringAgentProcessArrayInput interface {
	pulumi.Input

	ToServiceGroupMonitoringAgentProcessArrayOutput() ServiceGroupMonitoringAgentProcessArrayOutput
	ToServiceGroupMonitoringAgentProcessArrayOutputWithContext(context.Context) ServiceGroupMonitoringAgentProcessArrayOutput
}

type ServiceGroupMonitoringAgentProcessArray []ServiceGroupMonitoringAgentProcessInput

func (ServiceGroupMonitoringAgentProcessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (i ServiceGroupMonitoringAgentProcessArray) ToServiceGroupMonitoringAgentProcessArrayOutput() ServiceGroupMonitoringAgentProcessArrayOutput {
	return i.ToServiceGroupMonitoringAgentProcessArrayOutputWithContext(context.Background())
}

func (i ServiceGroupMonitoringAgentProcessArray) ToServiceGroupMonitoringAgentProcessArrayOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupMonitoringAgentProcessArrayOutput)
}

// ServiceGroupMonitoringAgentProcessMapInput is an input type that accepts ServiceGroupMonitoringAgentProcessMap and ServiceGroupMonitoringAgentProcessMapOutput values.
// You can construct a concrete instance of `ServiceGroupMonitoringAgentProcessMapInput` via:
//
//	ServiceGroupMonitoringAgentProcessMap{ "key": ServiceGroupMonitoringAgentProcessArgs{...} }
type ServiceGroupMonitoringAgentProcessMapInput interface {
	pulumi.Input

	ToServiceGroupMonitoringAgentProcessMapOutput() ServiceGroupMonitoringAgentProcessMapOutput
	ToServiceGroupMonitoringAgentProcessMapOutputWithContext(context.Context) ServiceGroupMonitoringAgentProcessMapOutput
}

type ServiceGroupMonitoringAgentProcessMap map[string]ServiceGroupMonitoringAgentProcessInput

func (ServiceGroupMonitoringAgentProcessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (i ServiceGroupMonitoringAgentProcessMap) ToServiceGroupMonitoringAgentProcessMapOutput() ServiceGroupMonitoringAgentProcessMapOutput {
	return i.ToServiceGroupMonitoringAgentProcessMapOutputWithContext(context.Background())
}

func (i ServiceGroupMonitoringAgentProcessMap) ToServiceGroupMonitoringAgentProcessMapOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceGroupMonitoringAgentProcessMapOutput)
}

type ServiceGroupMonitoringAgentProcessOutput struct{ *pulumi.OutputState }

func (ServiceGroupMonitoringAgentProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (o ServiceGroupMonitoringAgentProcessOutput) ToServiceGroupMonitoringAgentProcessOutput() ServiceGroupMonitoringAgentProcessOutput {
	return o
}

func (o ServiceGroupMonitoringAgentProcessOutput) ToServiceGroupMonitoringAgentProcessOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessOutput {
	return o
}

// The alert rule configurations. See `alertConfig` below.
func (o ServiceGroupMonitoringAgentProcessOutput) AlertConfigs() ServiceGroupMonitoringAgentProcessAlertConfigArrayOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) ServiceGroupMonitoringAgentProcessAlertConfigArrayOutput {
		return v.AlertConfigs
	}).(ServiceGroupMonitoringAgentProcessAlertConfigArrayOutput)
}

// The ID of the application group.
func (o ServiceGroupMonitoringAgentProcessOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The ID of the Group Monitoring Agent Process.
func (o ServiceGroupMonitoringAgentProcessOutput) GroupMonitoringAgentProcessId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) pulumi.StringOutput {
		return v.GroupMonitoringAgentProcessId
	}).(pulumi.StringOutput)
}

// The logical operator used between conditional expressions that are used to match instances. Valid values: `all`, `and`, `or`.
func (o ServiceGroupMonitoringAgentProcessOutput) MatchExpressFilterRelation() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) pulumi.StringOutput { return v.MatchExpressFilterRelation }).(pulumi.StringOutput)
}

// The expressions used to match instances. See `matchExpress` below.
func (o ServiceGroupMonitoringAgentProcessOutput) MatchExpresses() ServiceGroupMonitoringAgentProcessMatchExpressArrayOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) ServiceGroupMonitoringAgentProcessMatchExpressArrayOutput {
		return v.MatchExpresses
	}).(ServiceGroupMonitoringAgentProcessMatchExpressArrayOutput)
}

// The name of the process.
func (o ServiceGroupMonitoringAgentProcessOutput) ProcessName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceGroupMonitoringAgentProcess) pulumi.StringOutput { return v.ProcessName }).(pulumi.StringOutput)
}

type ServiceGroupMonitoringAgentProcessArrayOutput struct{ *pulumi.OutputState }

func (ServiceGroupMonitoringAgentProcessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (o ServiceGroupMonitoringAgentProcessArrayOutput) ToServiceGroupMonitoringAgentProcessArrayOutput() ServiceGroupMonitoringAgentProcessArrayOutput {
	return o
}

func (o ServiceGroupMonitoringAgentProcessArrayOutput) ToServiceGroupMonitoringAgentProcessArrayOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessArrayOutput {
	return o
}

func (o ServiceGroupMonitoringAgentProcessArrayOutput) Index(i pulumi.IntInput) ServiceGroupMonitoringAgentProcessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceGroupMonitoringAgentProcess {
		return vs[0].([]*ServiceGroupMonitoringAgentProcess)[vs[1].(int)]
	}).(ServiceGroupMonitoringAgentProcessOutput)
}

type ServiceGroupMonitoringAgentProcessMapOutput struct{ *pulumi.OutputState }

func (ServiceGroupMonitoringAgentProcessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceGroupMonitoringAgentProcess)(nil)).Elem()
}

func (o ServiceGroupMonitoringAgentProcessMapOutput) ToServiceGroupMonitoringAgentProcessMapOutput() ServiceGroupMonitoringAgentProcessMapOutput {
	return o
}

func (o ServiceGroupMonitoringAgentProcessMapOutput) ToServiceGroupMonitoringAgentProcessMapOutputWithContext(ctx context.Context) ServiceGroupMonitoringAgentProcessMapOutput {
	return o
}

func (o ServiceGroupMonitoringAgentProcessMapOutput) MapIndex(k pulumi.StringInput) ServiceGroupMonitoringAgentProcessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceGroupMonitoringAgentProcess {
		return vs[0].(map[string]*ServiceGroupMonitoringAgentProcess)[vs[1].(string)]
	}).(ServiceGroupMonitoringAgentProcessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupMonitoringAgentProcessInput)(nil)).Elem(), &ServiceGroupMonitoringAgentProcess{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupMonitoringAgentProcessArrayInput)(nil)).Elem(), ServiceGroupMonitoringAgentProcessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceGroupMonitoringAgentProcessMapInput)(nil)).Elem(), ServiceGroupMonitoringAgentProcessMap{})
	pulumi.RegisterOutputType(ServiceGroupMonitoringAgentProcessOutput{})
	pulumi.RegisterOutputType(ServiceGroupMonitoringAgentProcessArrayOutput{})
	pulumi.RegisterOutputType(ServiceGroupMonitoringAgentProcessMapOutput{})
}
