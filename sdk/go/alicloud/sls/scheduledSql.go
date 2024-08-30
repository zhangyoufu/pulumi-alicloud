// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SLS Scheduled SQL resource. Scheduled SQL task.
//
// For information about SLS Scheduled SQL and how to use it, see [What is Scheduled SQL](https://www.alibabacloud.com/help/zh/sls/developer-reference/api-sls-2020-12-30-createscheduledsql).
//
// > **NOTE:** Available since v1.224.0.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sls"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			defaultKIe4KV, err := log.NewProject(ctx, "defaultKIe4KV", &log.ProjectArgs{
//				Description: pulumi.Sprintf("%v-%v", name, _default.Result),
//				ProjectName: pulumi.Sprintf("%v-%v", name, _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			default1LI9we, err := log.NewStore(ctx, "default1LI9we", &log.StoreArgs{
//				HotTtl:          pulumi.Int(8),
//				RetentionPeriod: pulumi.Int(30),
//				ShardCount:      pulumi.Int(2),
//				ProjectName:     defaultKIe4KV.ProjectName,
//				LogstoreName:    pulumi.Sprintf("%v-%v", name, _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sls.NewScheduledSql(ctx, "default", &sls.ScheduledSqlArgs{
//				Description: pulumi.String("example-tf-scheduled-sql-0006"),
//				Schedule: &sls.ScheduledSqlScheduleArgs{
//					Type:           pulumi.String("Cron"),
//					TimeZone:       pulumi.String("+0700"),
//					Delay:          pulumi.Int(20),
//					CronExpression: pulumi.String("0 0/1 * * *"),
//				},
//				DisplayName: pulumi.String("example-tf-scheduled-sql-0006"),
//				ScheduledSqlConfiguration: &sls.ScheduledSqlScheduledSqlConfigurationArgs{
//					Script:              pulumi.String("* | select * from log"),
//					SqlType:             pulumi.String("searchQuery"),
//					DestEndpoint:        pulumi.String("ap-northeast-1.log.aliyuncs.com"),
//					DestProject:         pulumi.String("job-e2e-project-jj78kur-ap-southeast-1"),
//					SourceLogstore:      default1LI9we.LogstoreName,
//					DestLogstore:        pulumi.String("example-open-api02"),
//					RoleArn:             pulumi.String("acs:ram::1395894005868720:role/aliyunlogetlrole"),
//					DestRoleArn:         pulumi.String("acs:ram::1395894005868720:role/aliyunlogetlrole"),
//					FromTimeExpr:        pulumi.String("@m-1m"),
//					ToTimeExpr:          pulumi.String("@m"),
//					MaxRunTimeInSeconds: pulumi.Int(1800),
//					ResourcePool:        pulumi.String("enhanced"),
//					MaxRetries:          pulumi.Int(5),
//					FromTime:            pulumi.Int(1713196800),
//					ToTime:              pulumi.Int(0),
//					DataFormat:          pulumi.String("log2log"),
//				},
//				ScheduledSqlName: pulumi.String(name),
//				Project:          defaultKIe4KV.ProjectName,
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
// SLS Scheduled SQL can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:sls/scheduledSql:ScheduledSql example <project>:<scheduled_sql_name>
// ```
type ScheduledSql struct {
	pulumi.CustomResourceState

	// Task Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Task Display Name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Log project.
	Project pulumi.StringOutput `pulumi:"project"`
	// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
	Schedule ScheduledSqlScheduleOutput `pulumi:"schedule"`
	// Task Configuration. See `scheduledSqlConfiguration` below.
	ScheduledSqlConfiguration ScheduledSqlScheduledSqlConfigurationOutput `pulumi:"scheduledSqlConfiguration"`
	// Timed SQL name.
	ScheduledSqlName pulumi.StringOutput `pulumi:"scheduledSqlName"`
}

// NewScheduledSql registers a new resource with the given unique name, arguments, and options.
func NewScheduledSql(ctx *pulumi.Context,
	name string, args *ScheduledSqlArgs, opts ...pulumi.ResourceOption) (*ScheduledSql, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.ScheduledSqlConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledSqlConfiguration'")
	}
	if args.ScheduledSqlName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledSqlName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScheduledSql
	err := ctx.RegisterResource("alicloud:sls/scheduledSql:ScheduledSql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledSql gets an existing ScheduledSql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledSql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledSqlState, opts ...pulumi.ResourceOption) (*ScheduledSql, error) {
	var resource ScheduledSql
	err := ctx.ReadResource("alicloud:sls/scheduledSql:ScheduledSql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledSql resources.
type scheduledSqlState struct {
	// Task Description.
	Description *string `pulumi:"description"`
	// Task Display Name.
	DisplayName *string `pulumi:"displayName"`
	// Log project.
	Project *string `pulumi:"project"`
	// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
	Schedule *ScheduledSqlSchedule `pulumi:"schedule"`
	// Task Configuration. See `scheduledSqlConfiguration` below.
	ScheduledSqlConfiguration *ScheduledSqlScheduledSqlConfiguration `pulumi:"scheduledSqlConfiguration"`
	// Timed SQL name.
	ScheduledSqlName *string `pulumi:"scheduledSqlName"`
}

type ScheduledSqlState struct {
	// Task Description.
	Description pulumi.StringPtrInput
	// Task Display Name.
	DisplayName pulumi.StringPtrInput
	// Log project.
	Project pulumi.StringPtrInput
	// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
	Schedule ScheduledSqlSchedulePtrInput
	// Task Configuration. See `scheduledSqlConfiguration` below.
	ScheduledSqlConfiguration ScheduledSqlScheduledSqlConfigurationPtrInput
	// Timed SQL name.
	ScheduledSqlName pulumi.StringPtrInput
}

func (ScheduledSqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSqlState)(nil)).Elem()
}

type scheduledSqlArgs struct {
	// Task Description.
	Description *string `pulumi:"description"`
	// Task Display Name.
	DisplayName string `pulumi:"displayName"`
	// Log project.
	Project string `pulumi:"project"`
	// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
	Schedule ScheduledSqlSchedule `pulumi:"schedule"`
	// Task Configuration. See `scheduledSqlConfiguration` below.
	ScheduledSqlConfiguration ScheduledSqlScheduledSqlConfiguration `pulumi:"scheduledSqlConfiguration"`
	// Timed SQL name.
	ScheduledSqlName string `pulumi:"scheduledSqlName"`
}

// The set of arguments for constructing a ScheduledSql resource.
type ScheduledSqlArgs struct {
	// Task Description.
	Description pulumi.StringPtrInput
	// Task Display Name.
	DisplayName pulumi.StringInput
	// Log project.
	Project pulumi.StringInput
	// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
	Schedule ScheduledSqlScheduleInput
	// Task Configuration. See `scheduledSqlConfiguration` below.
	ScheduledSqlConfiguration ScheduledSqlScheduledSqlConfigurationInput
	// Timed SQL name.
	ScheduledSqlName pulumi.StringInput
}

func (ScheduledSqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledSqlArgs)(nil)).Elem()
}

type ScheduledSqlInput interface {
	pulumi.Input

	ToScheduledSqlOutput() ScheduledSqlOutput
	ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput
}

func (*ScheduledSql) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledSql)(nil)).Elem()
}

func (i *ScheduledSql) ToScheduledSqlOutput() ScheduledSqlOutput {
	return i.ToScheduledSqlOutputWithContext(context.Background())
}

func (i *ScheduledSql) ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlOutput)
}

// ScheduledSqlArrayInput is an input type that accepts ScheduledSqlArray and ScheduledSqlArrayOutput values.
// You can construct a concrete instance of `ScheduledSqlArrayInput` via:
//
//	ScheduledSqlArray{ ScheduledSqlArgs{...} }
type ScheduledSqlArrayInput interface {
	pulumi.Input

	ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput
	ToScheduledSqlArrayOutputWithContext(context.Context) ScheduledSqlArrayOutput
}

type ScheduledSqlArray []ScheduledSqlInput

func (ScheduledSqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledSql)(nil)).Elem()
}

func (i ScheduledSqlArray) ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput {
	return i.ToScheduledSqlArrayOutputWithContext(context.Background())
}

func (i ScheduledSqlArray) ToScheduledSqlArrayOutputWithContext(ctx context.Context) ScheduledSqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlArrayOutput)
}

// ScheduledSqlMapInput is an input type that accepts ScheduledSqlMap and ScheduledSqlMapOutput values.
// You can construct a concrete instance of `ScheduledSqlMapInput` via:
//
//	ScheduledSqlMap{ "key": ScheduledSqlArgs{...} }
type ScheduledSqlMapInput interface {
	pulumi.Input

	ToScheduledSqlMapOutput() ScheduledSqlMapOutput
	ToScheduledSqlMapOutputWithContext(context.Context) ScheduledSqlMapOutput
}

type ScheduledSqlMap map[string]ScheduledSqlInput

func (ScheduledSqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledSql)(nil)).Elem()
}

func (i ScheduledSqlMap) ToScheduledSqlMapOutput() ScheduledSqlMapOutput {
	return i.ToScheduledSqlMapOutputWithContext(context.Background())
}

func (i ScheduledSqlMap) ToScheduledSqlMapOutputWithContext(ctx context.Context) ScheduledSqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSqlMapOutput)
}

type ScheduledSqlOutput struct{ *pulumi.OutputState }

func (ScheduledSqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlOutput) ToScheduledSqlOutput() ScheduledSqlOutput {
	return o
}

func (o ScheduledSqlOutput) ToScheduledSqlOutputWithContext(ctx context.Context) ScheduledSqlOutput {
	return o
}

// Task Description.
func (o ScheduledSqlOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Task Display Name.
func (o ScheduledSqlOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Log project.
func (o ScheduledSqlOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The scheduling type is generally not required by default. If there is a strong timing requirement, if it must be imported every Monday at 8 o'clock, cron can be used. See `schedule` below.
func (o ScheduledSqlOutput) Schedule() ScheduledSqlScheduleOutput {
	return o.ApplyT(func(v *ScheduledSql) ScheduledSqlScheduleOutput { return v.Schedule }).(ScheduledSqlScheduleOutput)
}

// Task Configuration. See `scheduledSqlConfiguration` below.
func (o ScheduledSqlOutput) ScheduledSqlConfiguration() ScheduledSqlScheduledSqlConfigurationOutput {
	return o.ApplyT(func(v *ScheduledSql) ScheduledSqlScheduledSqlConfigurationOutput { return v.ScheduledSqlConfiguration }).(ScheduledSqlScheduledSqlConfigurationOutput)
}

// Timed SQL name.
func (o ScheduledSqlOutput) ScheduledSqlName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledSql) pulumi.StringOutput { return v.ScheduledSqlName }).(pulumi.StringOutput)
}

type ScheduledSqlArrayOutput struct{ *pulumi.OutputState }

func (ScheduledSqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlArrayOutput) ToScheduledSqlArrayOutput() ScheduledSqlArrayOutput {
	return o
}

func (o ScheduledSqlArrayOutput) ToScheduledSqlArrayOutputWithContext(ctx context.Context) ScheduledSqlArrayOutput {
	return o
}

func (o ScheduledSqlArrayOutput) Index(i pulumi.IntInput) ScheduledSqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScheduledSql {
		return vs[0].([]*ScheduledSql)[vs[1].(int)]
	}).(ScheduledSqlOutput)
}

type ScheduledSqlMapOutput struct{ *pulumi.OutputState }

func (ScheduledSqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledSql)(nil)).Elem()
}

func (o ScheduledSqlMapOutput) ToScheduledSqlMapOutput() ScheduledSqlMapOutput {
	return o
}

func (o ScheduledSqlMapOutput) ToScheduledSqlMapOutputWithContext(ctx context.Context) ScheduledSqlMapOutput {
	return o
}

func (o ScheduledSqlMapOutput) MapIndex(k pulumi.StringInput) ScheduledSqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScheduledSql {
		return vs[0].(map[string]*ScheduledSql)[vs[1].(string)]
	}).(ScheduledSqlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlInput)(nil)).Elem(), &ScheduledSql{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlArrayInput)(nil)).Elem(), ScheduledSqlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSqlMapInput)(nil)).Elem(), ScheduledSqlMap{})
	pulumi.RegisterOutputType(ScheduledSqlOutput{})
	pulumi.RegisterOutputType(ScheduledSqlArrayOutput{})
	pulumi.RegisterOutputType(ScheduledSqlMapOutput{})
}
