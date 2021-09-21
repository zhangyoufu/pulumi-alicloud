// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a HBR Ecs Backup Plan resource.
//
// For information about HBR Ecs Backup Plan and how to use it, see [What is Ecs Backup Plan](https://www.alibabacloud.com/help/doc-detail/186574.htm).
//
// > **NOTE:** Available in v1.132.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hbr.NewEcsBackupPlan(ctx, "example", &hbr.EcsBackupPlanArgs{
// 			BackupType:        pulumi.String("COMPLETE"),
// 			EcsBackupPlanName: pulumi.String("example_value"),
// 			Exclude:           pulumi.String(fmt.Sprintf("%v%v", "  [\"/home/exclude\"]\n", "  \n")),
// 			Include:           pulumi.String(fmt.Sprintf("%v%v", "  [\"/home/include\"]\n", "  \n")),
// 			InstanceId:        pulumi.String("i-bp1567rc0oxxxxxxxxxx"),
// 			Paths: pulumi.StringArray{
// 				pulumi.String("/home"),
// 				pulumi.String("/var"),
// 			},
// 			Retention:  pulumi.String("1"),
// 			Schedule:   pulumi.String("I|1602673264|PT2H"),
// 			SpeedLimit: pulumi.String("0:24:5120"),
// 			VaultId:    pulumi.String("v-0003gxoksflhxxxxxxxx"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// HBR Ecs Backup Plan can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:hbr/ecsBackupPlan:EcsBackupPlan example <id>
// ```
type EcsBackupPlan struct {
	pulumi.CustomResourceState

	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringOutput    `pulumi:"backupType"`
	Detail     pulumi.StringPtrOutput `pulumi:"detail"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	EcsBackupPlanName pulumi.StringOutput `pulumi:"ecsBackupPlanName"`
	// Exclude path. String of Json list, up to 255 characters. e.g. `"[\"/home/work\"]"`
	Exclude pulumi.StringPtrOutput `pulumi:"exclude"`
	// Include path. String of Json list, up to 255 characters. e.g. `"[\"/var\"]"`
	Include pulumi.StringPtrOutput `pulumi:"include"`
	// The ID of ECS instance. The ecs backup client must have been installed on the host.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Windows operating system with application consistency using VSS. eg: {`UseVSS`:false}.
	Options pulumi.StringPtrOutput `pulumi:"options"`
	// Backup path. e.g. `["/home", "/var"]`
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringOutput `pulumi:"retention"`
	// Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Flow control. The format is: {start}|{end}|{bandwidth}. Use `|` to separate multiple flow control configurations, multiple flow control configurations not allowed to have overlapping times.
	SpeedLimit  pulumi.StringPtrOutput `pulumi:"speedLimit"`
	UpdatePaths pulumi.BoolPtrOutput   `pulumi:"updatePaths"`
	// The ID of Backup vault.
	VaultId pulumi.StringPtrOutput `pulumi:"vaultId"`
}

// NewEcsBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewEcsBackupPlan(ctx *pulumi.Context,
	name string, args *EcsBackupPlanArgs, opts ...pulumi.ResourceOption) (*EcsBackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EcsBackupPlanName == nil {
		return nil, errors.New("invalid value for required argument 'EcsBackupPlanName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource EcsBackupPlan
	err := ctx.RegisterResource("alicloud:hbr/ecsBackupPlan:EcsBackupPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsBackupPlan gets an existing EcsBackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsBackupPlanState, opts ...pulumi.ResourceOption) (*EcsBackupPlan, error) {
	var resource EcsBackupPlan
	err := ctx.ReadResource("alicloud:hbr/ecsBackupPlan:EcsBackupPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsBackupPlan resources.
type ecsBackupPlanState struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType *string `pulumi:"backupType"`
	Detail     *string `pulumi:"detail"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled *bool `pulumi:"disabled"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	EcsBackupPlanName *string `pulumi:"ecsBackupPlanName"`
	// Exclude path. String of Json list, up to 255 characters. e.g. `"[\"/home/work\"]"`
	Exclude *string `pulumi:"exclude"`
	// Include path. String of Json list, up to 255 characters. e.g. `"[\"/var\"]"`
	Include *string `pulumi:"include"`
	// The ID of ECS instance. The ecs backup client must have been installed on the host.
	InstanceId *string `pulumi:"instanceId"`
	// Windows operating system with application consistency using VSS. eg: {`UseVSS`:false}.
	Options *string `pulumi:"options"`
	// Backup path. e.g. `["/home", "/var"]`
	Paths []string `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention *string `pulumi:"retention"`
	// Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
	Schedule *string `pulumi:"schedule"`
	// Flow control. The format is: {start}|{end}|{bandwidth}. Use `|` to separate multiple flow control configurations, multiple flow control configurations not allowed to have overlapping times.
	SpeedLimit  *string `pulumi:"speedLimit"`
	UpdatePaths *bool   `pulumi:"updatePaths"`
	// The ID of Backup vault.
	VaultId *string `pulumi:"vaultId"`
}

type EcsBackupPlanState struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringPtrInput
	Detail     pulumi.StringPtrInput
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolPtrInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	EcsBackupPlanName pulumi.StringPtrInput
	// Exclude path. String of Json list, up to 255 characters. e.g. `"[\"/home/work\"]"`
	Exclude pulumi.StringPtrInput
	// Include path. String of Json list, up to 255 characters. e.g. `"[\"/var\"]"`
	Include pulumi.StringPtrInput
	// The ID of ECS instance. The ecs backup client must have been installed on the host.
	InstanceId pulumi.StringPtrInput
	// Windows operating system with application consistency using VSS. eg: {`UseVSS`:false}.
	Options pulumi.StringPtrInput
	// Backup path. e.g. `["/home", "/var"]`
	Paths pulumi.StringArrayInput
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringPtrInput
	// Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringPtrInput
	// Flow control. The format is: {start}|{end}|{bandwidth}. Use `|` to separate multiple flow control configurations, multiple flow control configurations not allowed to have overlapping times.
	SpeedLimit  pulumi.StringPtrInput
	UpdatePaths pulumi.BoolPtrInput
	// The ID of Backup vault.
	VaultId pulumi.StringPtrInput
}

func (EcsBackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsBackupPlanState)(nil)).Elem()
}

type ecsBackupPlanArgs struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType *string `pulumi:"backupType"`
	Detail     *string `pulumi:"detail"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled *bool `pulumi:"disabled"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	EcsBackupPlanName string `pulumi:"ecsBackupPlanName"`
	// Exclude path. String of Json list, up to 255 characters. e.g. `"[\"/home/work\"]"`
	Exclude *string `pulumi:"exclude"`
	// Include path. String of Json list, up to 255 characters. e.g. `"[\"/var\"]"`
	Include *string `pulumi:"include"`
	// The ID of ECS instance. The ecs backup client must have been installed on the host.
	InstanceId string `pulumi:"instanceId"`
	// Windows operating system with application consistency using VSS. eg: {`UseVSS`:false}.
	Options *string `pulumi:"options"`
	// Backup path. e.g. `["/home", "/var"]`
	Paths []string `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention string `pulumi:"retention"`
	// Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
	Schedule string `pulumi:"schedule"`
	// Flow control. The format is: {start}|{end}|{bandwidth}. Use `|` to separate multiple flow control configurations, multiple flow control configurations not allowed to have overlapping times.
	SpeedLimit  *string `pulumi:"speedLimit"`
	UpdatePaths *bool   `pulumi:"updatePaths"`
	// The ID of Backup vault.
	VaultId *string `pulumi:"vaultId"`
}

// The set of arguments for constructing a EcsBackupPlan resource.
type EcsBackupPlanArgs struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringPtrInput
	Detail     pulumi.StringPtrInput
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolPtrInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	EcsBackupPlanName pulumi.StringInput
	// Exclude path. String of Json list, up to 255 characters. e.g. `"[\"/home/work\"]"`
	Exclude pulumi.StringPtrInput
	// Include path. String of Json list, up to 255 characters. e.g. `"[\"/var\"]"`
	Include pulumi.StringPtrInput
	// The ID of ECS instance. The ecs backup client must have been installed on the host.
	InstanceId pulumi.StringInput
	// Windows operating system with application consistency using VSS. eg: {`UseVSS`:false}.
	Options pulumi.StringPtrInput
	// Backup path. e.g. `["/home", "/var"]`
	Paths pulumi.StringArrayInput
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringInput
	// Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringInput
	// Flow control. The format is: {start}|{end}|{bandwidth}. Use `|` to separate multiple flow control configurations, multiple flow control configurations not allowed to have overlapping times.
	SpeedLimit  pulumi.StringPtrInput
	UpdatePaths pulumi.BoolPtrInput
	// The ID of Backup vault.
	VaultId pulumi.StringPtrInput
}

func (EcsBackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsBackupPlanArgs)(nil)).Elem()
}

type EcsBackupPlanInput interface {
	pulumi.Input

	ToEcsBackupPlanOutput() EcsBackupPlanOutput
	ToEcsBackupPlanOutputWithContext(ctx context.Context) EcsBackupPlanOutput
}

func (*EcsBackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*EcsBackupPlan)(nil))
}

func (i *EcsBackupPlan) ToEcsBackupPlanOutput() EcsBackupPlanOutput {
	return i.ToEcsBackupPlanOutputWithContext(context.Background())
}

func (i *EcsBackupPlan) ToEcsBackupPlanOutputWithContext(ctx context.Context) EcsBackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsBackupPlanOutput)
}

func (i *EcsBackupPlan) ToEcsBackupPlanPtrOutput() EcsBackupPlanPtrOutput {
	return i.ToEcsBackupPlanPtrOutputWithContext(context.Background())
}

func (i *EcsBackupPlan) ToEcsBackupPlanPtrOutputWithContext(ctx context.Context) EcsBackupPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsBackupPlanPtrOutput)
}

type EcsBackupPlanPtrInput interface {
	pulumi.Input

	ToEcsBackupPlanPtrOutput() EcsBackupPlanPtrOutput
	ToEcsBackupPlanPtrOutputWithContext(ctx context.Context) EcsBackupPlanPtrOutput
}

type ecsBackupPlanPtrType EcsBackupPlanArgs

func (*ecsBackupPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsBackupPlan)(nil))
}

func (i *ecsBackupPlanPtrType) ToEcsBackupPlanPtrOutput() EcsBackupPlanPtrOutput {
	return i.ToEcsBackupPlanPtrOutputWithContext(context.Background())
}

func (i *ecsBackupPlanPtrType) ToEcsBackupPlanPtrOutputWithContext(ctx context.Context) EcsBackupPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsBackupPlanPtrOutput)
}

// EcsBackupPlanArrayInput is an input type that accepts EcsBackupPlanArray and EcsBackupPlanArrayOutput values.
// You can construct a concrete instance of `EcsBackupPlanArrayInput` via:
//
//          EcsBackupPlanArray{ EcsBackupPlanArgs{...} }
type EcsBackupPlanArrayInput interface {
	pulumi.Input

	ToEcsBackupPlanArrayOutput() EcsBackupPlanArrayOutput
	ToEcsBackupPlanArrayOutputWithContext(context.Context) EcsBackupPlanArrayOutput
}

type EcsBackupPlanArray []EcsBackupPlanInput

func (EcsBackupPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EcsBackupPlan)(nil))
}

func (i EcsBackupPlanArray) ToEcsBackupPlanArrayOutput() EcsBackupPlanArrayOutput {
	return i.ToEcsBackupPlanArrayOutputWithContext(context.Background())
}

func (i EcsBackupPlanArray) ToEcsBackupPlanArrayOutputWithContext(ctx context.Context) EcsBackupPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsBackupPlanArrayOutput)
}

// EcsBackupPlanMapInput is an input type that accepts EcsBackupPlanMap and EcsBackupPlanMapOutput values.
// You can construct a concrete instance of `EcsBackupPlanMapInput` via:
//
//          EcsBackupPlanMap{ "key": EcsBackupPlanArgs{...} }
type EcsBackupPlanMapInput interface {
	pulumi.Input

	ToEcsBackupPlanMapOutput() EcsBackupPlanMapOutput
	ToEcsBackupPlanMapOutputWithContext(context.Context) EcsBackupPlanMapOutput
}

type EcsBackupPlanMap map[string]EcsBackupPlanInput

func (EcsBackupPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EcsBackupPlan)(nil))
}

func (i EcsBackupPlanMap) ToEcsBackupPlanMapOutput() EcsBackupPlanMapOutput {
	return i.ToEcsBackupPlanMapOutputWithContext(context.Background())
}

func (i EcsBackupPlanMap) ToEcsBackupPlanMapOutputWithContext(ctx context.Context) EcsBackupPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsBackupPlanMapOutput)
}

type EcsBackupPlanOutput struct {
	*pulumi.OutputState
}

func (EcsBackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EcsBackupPlan)(nil))
}

func (o EcsBackupPlanOutput) ToEcsBackupPlanOutput() EcsBackupPlanOutput {
	return o
}

func (o EcsBackupPlanOutput) ToEcsBackupPlanOutputWithContext(ctx context.Context) EcsBackupPlanOutput {
	return o
}

func (o EcsBackupPlanOutput) ToEcsBackupPlanPtrOutput() EcsBackupPlanPtrOutput {
	return o.ToEcsBackupPlanPtrOutputWithContext(context.Background())
}

func (o EcsBackupPlanOutput) ToEcsBackupPlanPtrOutputWithContext(ctx context.Context) EcsBackupPlanPtrOutput {
	return o.ApplyT(func(v EcsBackupPlan) *EcsBackupPlan {
		return &v
	}).(EcsBackupPlanPtrOutput)
}

type EcsBackupPlanPtrOutput struct {
	*pulumi.OutputState
}

func (EcsBackupPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsBackupPlan)(nil))
}

func (o EcsBackupPlanPtrOutput) ToEcsBackupPlanPtrOutput() EcsBackupPlanPtrOutput {
	return o
}

func (o EcsBackupPlanPtrOutput) ToEcsBackupPlanPtrOutputWithContext(ctx context.Context) EcsBackupPlanPtrOutput {
	return o
}

type EcsBackupPlanArrayOutput struct{ *pulumi.OutputState }

func (EcsBackupPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EcsBackupPlan)(nil))
}

func (o EcsBackupPlanArrayOutput) ToEcsBackupPlanArrayOutput() EcsBackupPlanArrayOutput {
	return o
}

func (o EcsBackupPlanArrayOutput) ToEcsBackupPlanArrayOutputWithContext(ctx context.Context) EcsBackupPlanArrayOutput {
	return o
}

func (o EcsBackupPlanArrayOutput) Index(i pulumi.IntInput) EcsBackupPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EcsBackupPlan {
		return vs[0].([]EcsBackupPlan)[vs[1].(int)]
	}).(EcsBackupPlanOutput)
}

type EcsBackupPlanMapOutput struct{ *pulumi.OutputState }

func (EcsBackupPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EcsBackupPlan)(nil))
}

func (o EcsBackupPlanMapOutput) ToEcsBackupPlanMapOutput() EcsBackupPlanMapOutput {
	return o
}

func (o EcsBackupPlanMapOutput) ToEcsBackupPlanMapOutputWithContext(ctx context.Context) EcsBackupPlanMapOutput {
	return o
}

func (o EcsBackupPlanMapOutput) MapIndex(k pulumi.StringInput) EcsBackupPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EcsBackupPlan {
		return vs[0].(map[string]EcsBackupPlan)[vs[1].(string)]
	}).(EcsBackupPlanOutput)
}

func init() {
	pulumi.RegisterOutputType(EcsBackupPlanOutput{})
	pulumi.RegisterOutputType(EcsBackupPlanPtrOutput{})
	pulumi.RegisterOutputType(EcsBackupPlanArrayOutput{})
	pulumi.RegisterOutputType(EcsBackupPlanMapOutput{})
}
