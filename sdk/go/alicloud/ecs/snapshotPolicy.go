// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:** This resource has been renamed to ecs.AutoSnapshotPolicy from version 1.117.0.
//
// Provides an ECS snapshot policy resource.
//
// For information about snapshot policy and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
//
// > **NOTE:** Available in 1.42.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewSnapshotPolicy(ctx, "sp", &ecs.SnapshotPolicyArgs{
//				Name: pulumi.String("tf-testAcc-sp"),
//				RepeatWeekdays: pulumi.StringArray{
//					pulumi.String("1"),
//					pulumi.String("2"),
//					pulumi.String("3"),
//				},
//				RetentionDays: int(-1),
//				TimePoints: pulumi.StringArray{
//					pulumi.String("1"),
//					pulumi.String("22"),
//					pulumi.String("23"),
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
// Snapshot can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/snapshotPolicy:SnapshotPolicy snapshot sp-abc1234567890000
// ```
type SnapshotPolicy struct {
	pulumi.CustomResourceState

	CopiedSnapshotsRetentionDays pulumi.IntPtrOutput  `pulumi:"copiedSnapshotsRetentionDays"`
	EnableCrossRegionCopy        pulumi.BoolPtrOutput `pulumi:"enableCrossRegionCopy"`
	// The snapshot policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayOutput `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	//
	// Default value: -1.
	RetentionDays     pulumi.IntOutput         `pulumi:"retentionDays"`
	Status            pulumi.StringOutput      `pulumi:"status"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	TargetCopyRegions pulumi.StringArrayOutput `pulumi:"targetCopyRegions"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayOutput `pulumi:"timePoints"`
}

// NewSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotPolicy(ctx *pulumi.Context,
	name string, args *SnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepeatWeekdays == nil {
		return nil, errors.New("invalid value for required argument 'RepeatWeekdays'")
	}
	if args.RetentionDays == nil {
		return nil, errors.New("invalid value for required argument 'RetentionDays'")
	}
	if args.TimePoints == nil {
		return nil, errors.New("invalid value for required argument 'TimePoints'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SnapshotPolicy
	err := ctx.RegisterResource("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotPolicy gets an existing SnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyState, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	var resource SnapshotPolicy
	err := ctx.ReadResource("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotPolicy resources.
type snapshotPolicyState struct {
	CopiedSnapshotsRetentionDays *int  `pulumi:"copiedSnapshotsRetentionDays"`
	EnableCrossRegionCopy        *bool `pulumi:"enableCrossRegionCopy"`
	// The snapshot policy name.
	Name *string `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	//
	// Default value: -1.
	RetentionDays     *int              `pulumi:"retentionDays"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	TargetCopyRegions []string          `pulumi:"targetCopyRegions"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

type SnapshotPolicyState struct {
	CopiedSnapshotsRetentionDays pulumi.IntPtrInput
	EnableCrossRegionCopy        pulumi.BoolPtrInput
	// The snapshot policy name.
	Name pulumi.StringPtrInput
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	//
	// Default value: -1.
	RetentionDays     pulumi.IntPtrInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TargetCopyRegions pulumi.StringArrayInput
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (SnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyState)(nil)).Elem()
}

type snapshotPolicyArgs struct {
	CopiedSnapshotsRetentionDays *int  `pulumi:"copiedSnapshotsRetentionDays"`
	EnableCrossRegionCopy        *bool `pulumi:"enableCrossRegionCopy"`
	// The snapshot policy name.
	Name *string `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	//
	// Default value: -1.
	RetentionDays     int               `pulumi:"retentionDays"`
	Tags              map[string]string `pulumi:"tags"`
	TargetCopyRegions []string          `pulumi:"targetCopyRegions"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

// The set of arguments for constructing a SnapshotPolicy resource.
type SnapshotPolicyArgs struct {
	CopiedSnapshotsRetentionDays pulumi.IntPtrInput
	EnableCrossRegionCopy        pulumi.BoolPtrInput
	// The snapshot policy name.
	Name pulumi.StringPtrInput
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	//
	// Default value: -1.
	RetentionDays     pulumi.IntInput
	Tags              pulumi.StringMapInput
	TargetCopyRegions pulumi.StringArrayInput
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (SnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyArgs)(nil)).Elem()
}

type SnapshotPolicyInput interface {
	pulumi.Input

	ToSnapshotPolicyOutput() SnapshotPolicyOutput
	ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput
}

func (*SnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return i.ToSnapshotPolicyOutputWithContext(context.Background())
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyOutput)
}

// SnapshotPolicyArrayInput is an input type that accepts SnapshotPolicyArray and SnapshotPolicyArrayOutput values.
// You can construct a concrete instance of `SnapshotPolicyArrayInput` via:
//
//	SnapshotPolicyArray{ SnapshotPolicyArgs{...} }
type SnapshotPolicyArrayInput interface {
	pulumi.Input

	ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput
	ToSnapshotPolicyArrayOutputWithContext(context.Context) SnapshotPolicyArrayOutput
}

type SnapshotPolicyArray []SnapshotPolicyInput

func (SnapshotPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicy)(nil)).Elem()
}

func (i SnapshotPolicyArray) ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput {
	return i.ToSnapshotPolicyArrayOutputWithContext(context.Background())
}

func (i SnapshotPolicyArray) ToSnapshotPolicyArrayOutputWithContext(ctx context.Context) SnapshotPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyArrayOutput)
}

// SnapshotPolicyMapInput is an input type that accepts SnapshotPolicyMap and SnapshotPolicyMapOutput values.
// You can construct a concrete instance of `SnapshotPolicyMapInput` via:
//
//	SnapshotPolicyMap{ "key": SnapshotPolicyArgs{...} }
type SnapshotPolicyMapInput interface {
	pulumi.Input

	ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput
	ToSnapshotPolicyMapOutputWithContext(context.Context) SnapshotPolicyMapOutput
}

type SnapshotPolicyMap map[string]SnapshotPolicyInput

func (SnapshotPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicy)(nil)).Elem()
}

func (i SnapshotPolicyMap) ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput {
	return i.ToSnapshotPolicyMapOutputWithContext(context.Background())
}

func (i SnapshotPolicyMap) ToSnapshotPolicyMapOutputWithContext(ctx context.Context) SnapshotPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyMapOutput)
}

type SnapshotPolicyOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) CopiedSnapshotsRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.IntPtrOutput { return v.CopiedSnapshotsRetentionDays }).(pulumi.IntPtrOutput)
}

func (o SnapshotPolicyOutput) EnableCrossRegionCopy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.BoolPtrOutput { return v.EnableCrossRegionCopy }).(pulumi.BoolPtrOutput)
}

// The snapshot policy name.
func (o SnapshotPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
// - A maximum of seven time points can be selected.
// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
func (o SnapshotPolicyOutput) RepeatWeekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringArrayOutput { return v.RepeatWeekdays }).(pulumi.StringArrayOutput)
}

// The snapshot retention time, and the unit of measurement is day. Optional values:
// - -1: The automatic snapshots are retained permanently.
// - [1, 65536]: The number of days retained.
//
// Default value: -1.
func (o SnapshotPolicyOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

func (o SnapshotPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SnapshotPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SnapshotPolicyOutput) TargetCopyRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringArrayOutput { return v.TargetCopyRegions }).(pulumi.StringArrayOutput)
}

// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
// - A maximum of 24 time points can be selected.
// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
func (o SnapshotPolicyOutput) TimePoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringArrayOutput { return v.TimePoints }).(pulumi.StringArrayOutput)
}

type SnapshotPolicyArrayOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyArrayOutput) ToSnapshotPolicyArrayOutput() SnapshotPolicyArrayOutput {
	return o
}

func (o SnapshotPolicyArrayOutput) ToSnapshotPolicyArrayOutputWithContext(ctx context.Context) SnapshotPolicyArrayOutput {
	return o
}

func (o SnapshotPolicyArrayOutput) Index(i pulumi.IntInput) SnapshotPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotPolicy {
		return vs[0].([]*SnapshotPolicy)[vs[1].(int)]
	}).(SnapshotPolicyOutput)
}

type SnapshotPolicyMapOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyMapOutput) ToSnapshotPolicyMapOutput() SnapshotPolicyMapOutput {
	return o
}

func (o SnapshotPolicyMapOutput) ToSnapshotPolicyMapOutputWithContext(ctx context.Context) SnapshotPolicyMapOutput {
	return o
}

func (o SnapshotPolicyMapOutput) MapIndex(k pulumi.StringInput) SnapshotPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotPolicy {
		return vs[0].(map[string]*SnapshotPolicy)[vs[1].(string)]
	}).(SnapshotPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyInput)(nil)).Elem(), &SnapshotPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyArrayInput)(nil)).Elem(), SnapshotPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyMapInput)(nil)).Elem(), SnapshotPolicyMap{})
	pulumi.RegisterOutputType(SnapshotPolicyOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyArrayOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyMapOutput{})
}
