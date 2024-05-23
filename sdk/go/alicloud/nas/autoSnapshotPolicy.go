// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NAS Auto Snapshot Policy resource. Automatic snapshot policy.
//
// For information about NAS Auto Snapshot Policy and how to use it, see [What is Auto Snapshot Policy](https://www.alibabacloud.com/help/en/doc-detail/135662.html)).
//
// > **NOTE:** Available since v1.153.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
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
//			_, err := nas.NewAutoSnapshotPolicy(ctx, "default", &nas.AutoSnapshotPolicyArgs{
//				TimePoints: pulumi.StringArray{
//					pulumi.String("0"),
//					pulumi.String("1"),
//					pulumi.String("2"),
//				},
//				RetentionDays: pulumi.Int(1),
//				RepeatWeekdays: pulumi.StringArray{
//					pulumi.String("2"),
//					pulumi.String("3"),
//					pulumi.String("4"),
//				},
//				AutoSnapshotPolicyName: pulumi.String(name),
//				FileSystemType:         pulumi.String("extreme"),
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
// NAS Auto Snapshot Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy example <id>
// ```
type AutoSnapshotPolicy struct {
	pulumi.CustomResourceState

	// The name of the automatic snapshot policy. Limits:
	// - The name must be `2` to `128` characters in length,
	// - The name must start with a letter.
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	// - The value of this parameter is empty by default.
	AutoSnapshotPolicyName pulumi.StringPtrOutput `pulumi:"autoSnapshotPolicyName"`
	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The file system type.
	FileSystemType pulumi.StringOutput `pulumi:"fileSystemType"`
	// The day on which an auto snapshot is created.
	// - A maximum of 7 time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayOutput `pulumi:"repeatWeekdays"`
	// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
	// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The status of the automatic snapshot policy.
	Status pulumi.StringOutput `pulumi:"status"`
	// The point in time at which an auto snapshot is created.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayOutput `pulumi:"timePoints"`
}

// NewAutoSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoSnapshotPolicy(ctx *pulumi.Context,
	name string, args *AutoSnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepeatWeekdays == nil {
		return nil, errors.New("invalid value for required argument 'RepeatWeekdays'")
	}
	if args.TimePoints == nil {
		return nil, errors.New("invalid value for required argument 'TimePoints'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoSnapshotPolicy
	err := ctx.RegisterResource("alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoSnapshotPolicy gets an existing AutoSnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoSnapshotPolicyState, opts ...pulumi.ResourceOption) (*AutoSnapshotPolicy, error) {
	var resource AutoSnapshotPolicy
	err := ctx.ReadResource("alicloud:nas/autoSnapshotPolicy:AutoSnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoSnapshotPolicy resources.
type autoSnapshotPolicyState struct {
	// The name of the automatic snapshot policy. Limits:
	// - The name must be `2` to `128` characters in length,
	// - The name must start with a letter.
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	// - The value of this parameter is empty by default.
	AutoSnapshotPolicyName *string `pulumi:"autoSnapshotPolicyName"`
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// The file system type.
	FileSystemType *string `pulumi:"fileSystemType"`
	// The day on which an auto snapshot is created.
	// - A maximum of 7 time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
	// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	RetentionDays *int `pulumi:"retentionDays"`
	// The status of the automatic snapshot policy.
	Status *string `pulumi:"status"`
	// The point in time at which an auto snapshot is created.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

type AutoSnapshotPolicyState struct {
	// The name of the automatic snapshot policy. Limits:
	// - The name must be `2` to `128` characters in length,
	// - The name must start with a letter.
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	// - The value of this parameter is empty by default.
	AutoSnapshotPolicyName pulumi.StringPtrInput
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// The file system type.
	FileSystemType pulumi.StringPtrInput
	// The day on which an auto snapshot is created.
	// - A maximum of 7 time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
	// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	RetentionDays pulumi.IntPtrInput
	// The status of the automatic snapshot policy.
	Status pulumi.StringPtrInput
	// The point in time at which an auto snapshot is created.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (AutoSnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyState)(nil)).Elem()
}

type autoSnapshotPolicyArgs struct {
	// The name of the automatic snapshot policy. Limits:
	// - The name must be `2` to `128` characters in length,
	// - The name must start with a letter.
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	// - The value of this parameter is empty by default.
	AutoSnapshotPolicyName *string `pulumi:"autoSnapshotPolicyName"`
	// The file system type.
	FileSystemType *string `pulumi:"fileSystemType"`
	// The day on which an auto snapshot is created.
	// - A maximum of 7 time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
	// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	RetentionDays *int `pulumi:"retentionDays"`
	// The point in time at which an auto snapshot is created.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

// The set of arguments for constructing a AutoSnapshotPolicy resource.
type AutoSnapshotPolicyArgs struct {
	// The name of the automatic snapshot policy. Limits:
	// - The name must be `2` to `128` characters in length,
	// - The name must start with a letter.
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	// - The value of this parameter is empty by default.
	AutoSnapshotPolicyName pulumi.StringPtrInput
	// The file system type.
	FileSystemType pulumi.StringPtrInput
	// The day on which an auto snapshot is created.
	// - A maximum of 7 time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
	// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	RetentionDays pulumi.IntPtrInput
	// The point in time at which an auto snapshot is created.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (AutoSnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapshotPolicyArgs)(nil)).Elem()
}

type AutoSnapshotPolicyInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput
	ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput
}

func (*AutoSnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicy)(nil)).Elem()
}

func (i *AutoSnapshotPolicy) ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput {
	return i.ToAutoSnapshotPolicyOutputWithContext(context.Background())
}

func (i *AutoSnapshotPolicy) ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyOutput)
}

// AutoSnapshotPolicyArrayInput is an input type that accepts AutoSnapshotPolicyArray and AutoSnapshotPolicyArrayOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyArrayInput` via:
//
//	AutoSnapshotPolicyArray{ AutoSnapshotPolicyArgs{...} }
type AutoSnapshotPolicyArrayInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput
	ToAutoSnapshotPolicyArrayOutputWithContext(context.Context) AutoSnapshotPolicyArrayOutput
}

type AutoSnapshotPolicyArray []AutoSnapshotPolicyInput

func (AutoSnapshotPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicy)(nil)).Elem()
}

func (i AutoSnapshotPolicyArray) ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput {
	return i.ToAutoSnapshotPolicyArrayOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyArray) ToAutoSnapshotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyArrayOutput)
}

// AutoSnapshotPolicyMapInput is an input type that accepts AutoSnapshotPolicyMap and AutoSnapshotPolicyMapOutput values.
// You can construct a concrete instance of `AutoSnapshotPolicyMapInput` via:
//
//	AutoSnapshotPolicyMap{ "key": AutoSnapshotPolicyArgs{...} }
type AutoSnapshotPolicyMapInput interface {
	pulumi.Input

	ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput
	ToAutoSnapshotPolicyMapOutputWithContext(context.Context) AutoSnapshotPolicyMapOutput
}

type AutoSnapshotPolicyMap map[string]AutoSnapshotPolicyInput

func (AutoSnapshotPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicy)(nil)).Elem()
}

func (i AutoSnapshotPolicyMap) ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput {
	return i.ToAutoSnapshotPolicyMapOutputWithContext(context.Background())
}

func (i AutoSnapshotPolicyMap) ToAutoSnapshotPolicyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapshotPolicyMapOutput)
}

type AutoSnapshotPolicyOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyOutput) ToAutoSnapshotPolicyOutput() AutoSnapshotPolicyOutput {
	return o
}

func (o AutoSnapshotPolicyOutput) ToAutoSnapshotPolicyOutputWithContext(ctx context.Context) AutoSnapshotPolicyOutput {
	return o
}

// The name of the automatic snapshot policy. Limits:
// - The name must be `2` to `128` characters in length,
// - The name must start with a letter.
// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
// - The value of this parameter is empty by default.
func (o AutoSnapshotPolicyOutput) AutoSnapshotPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringPtrOutput { return v.AutoSnapshotPolicyName }).(pulumi.StringPtrOutput)
}

// Creation time.
func (o AutoSnapshotPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The file system type.
func (o AutoSnapshotPolicyOutput) FileSystemType() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.FileSystemType }).(pulumi.StringOutput)
}

// The day on which an auto snapshot is created.
// - A maximum of 7 time points can be selected.
// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
func (o AutoSnapshotPolicyOutput) RepeatWeekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringArrayOutput { return v.RepeatWeekdays }).(pulumi.StringArrayOutput)
}

// The number of days for which you want to retain auto snapshots. Unit: days. Valid values:
// - `-1`: the default value. Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
func (o AutoSnapshotPolicyOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

// The status of the automatic snapshot policy.
func (o AutoSnapshotPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The point in time at which an auto snapshot is created.
// - A maximum of 24 time points can be selected.
// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
func (o AutoSnapshotPolicyOutput) TimePoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapshotPolicy) pulumi.StringArrayOutput { return v.TimePoints }).(pulumi.StringArrayOutput)
}

type AutoSnapshotPolicyArrayOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyArrayOutput) ToAutoSnapshotPolicyArrayOutput() AutoSnapshotPolicyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyArrayOutput) ToAutoSnapshotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapshotPolicyArrayOutput {
	return o
}

func (o AutoSnapshotPolicyArrayOutput) Index(i pulumi.IntInput) AutoSnapshotPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoSnapshotPolicy {
		return vs[0].([]*AutoSnapshotPolicy)[vs[1].(int)]
	}).(AutoSnapshotPolicyOutput)
}

type AutoSnapshotPolicyMapOutput struct{ *pulumi.OutputState }

func (AutoSnapshotPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapshotPolicy)(nil)).Elem()
}

func (o AutoSnapshotPolicyMapOutput) ToAutoSnapshotPolicyMapOutput() AutoSnapshotPolicyMapOutput {
	return o
}

func (o AutoSnapshotPolicyMapOutput) ToAutoSnapshotPolicyMapOutputWithContext(ctx context.Context) AutoSnapshotPolicyMapOutput {
	return o
}

func (o AutoSnapshotPolicyMapOutput) MapIndex(k pulumi.StringInput) AutoSnapshotPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoSnapshotPolicy {
		return vs[0].(map[string]*AutoSnapshotPolicy)[vs[1].(string)]
	}).(AutoSnapshotPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyInput)(nil)).Elem(), &AutoSnapshotPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyArrayInput)(nil)).Elem(), AutoSnapshotPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapshotPolicyMapInput)(nil)).Elem(), AutoSnapshotPolicyMap{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyArrayOutput{})
	pulumi.RegisterOutputType(AutoSnapshotPolicyMapOutput{})
}
