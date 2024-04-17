// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databasefilesystem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Dbfs Auto Snap Shot Policy resource.
//
// For information about Dbfs Auto Snap Shot Policy and how to use it.
//
// > **NOTE:** Available since v1.202.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/databasefilesystem"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := databasefilesystem.NewAutoSnapShotPolicy(ctx, "default", &databasefilesystem.AutoSnapShotPolicyArgs{
//				TimePoints: pulumi.StringArray{
//					pulumi.String("01"),
//				},
//				PolicyName:    pulumi.String("tf-example"),
//				RetentionDays: pulumi.Int(1),
//				RepeatWeekdays: pulumi.StringArray{
//					pulumi.String("2"),
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
// Dbfs Auto Snap Shot Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy example <id>
// ```
type AutoSnapShotPolicy struct {
	pulumi.CustomResourceState

	// The number of database file systems set by the automatic snapshot policy.
	AppliedDbfsNumber pulumi.IntOutput `pulumi:"appliedDbfsNumber"`
	// The creation time of the resource
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Last modification time of automatic snapshot policy
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// Automatic snapshot policy ID
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Automatic snapshot policy name
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
	RepeatWeekdays pulumi.StringArrayOutput `pulumi:"repeatWeekdays"`
	// Automatic snapshot retention days.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// Automatic snapshot policy status
	Status pulumi.StringOutput `pulumi:"status"`
	// Automatic snapshot policy status details
	StatusDetail pulumi.StringOutput `pulumi:"statusDetail"`
	// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
	TimePoints pulumi.StringArrayOutput `pulumi:"timePoints"`
}

// NewAutoSnapShotPolicy registers a new resource with the given unique name, arguments, and options.
func NewAutoSnapShotPolicy(ctx *pulumi.Context,
	name string, args *AutoSnapShotPolicyArgs, opts ...pulumi.ResourceOption) (*AutoSnapShotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
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
	var resource AutoSnapShotPolicy
	err := ctx.RegisterResource("alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoSnapShotPolicy gets an existing AutoSnapShotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoSnapShotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoSnapShotPolicyState, opts ...pulumi.ResourceOption) (*AutoSnapShotPolicy, error) {
	var resource AutoSnapShotPolicy
	err := ctx.ReadResource("alicloud:databasefilesystem/autoSnapShotPolicy:AutoSnapShotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoSnapShotPolicy resources.
type autoSnapShotPolicyState struct {
	// The number of database file systems set by the automatic snapshot policy.
	AppliedDbfsNumber *int `pulumi:"appliedDbfsNumber"`
	// The creation time of the resource
	CreateTime *string `pulumi:"createTime"`
	// Last modification time of automatic snapshot policy
	LastModified *string `pulumi:"lastModified"`
	// Automatic snapshot policy ID
	PolicyId *string `pulumi:"policyId"`
	// Automatic snapshot policy name
	PolicyName *string `pulumi:"policyName"`
	// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// Automatic snapshot retention days.
	RetentionDays *int `pulumi:"retentionDays"`
	// Automatic snapshot policy status
	Status *string `pulumi:"status"`
	// Automatic snapshot policy status details
	StatusDetail *string `pulumi:"statusDetail"`
	// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
	TimePoints []string `pulumi:"timePoints"`
}

type AutoSnapShotPolicyState struct {
	// The number of database file systems set by the automatic snapshot policy.
	AppliedDbfsNumber pulumi.IntPtrInput
	// The creation time of the resource
	CreateTime pulumi.StringPtrInput
	// Last modification time of automatic snapshot policy
	LastModified pulumi.StringPtrInput
	// Automatic snapshot policy ID
	PolicyId pulumi.StringPtrInput
	// Automatic snapshot policy name
	PolicyName pulumi.StringPtrInput
	// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
	RepeatWeekdays pulumi.StringArrayInput
	// Automatic snapshot retention days.
	RetentionDays pulumi.IntPtrInput
	// Automatic snapshot policy status
	Status pulumi.StringPtrInput
	// Automatic snapshot policy status details
	StatusDetail pulumi.StringPtrInput
	// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
	TimePoints pulumi.StringArrayInput
}

func (AutoSnapShotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapShotPolicyState)(nil)).Elem()
}

type autoSnapShotPolicyArgs struct {
	// Automatic snapshot policy name
	PolicyName string `pulumi:"policyName"`
	// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// Automatic snapshot retention days.
	RetentionDays int `pulumi:"retentionDays"`
	// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
	TimePoints []string `pulumi:"timePoints"`
}

// The set of arguments for constructing a AutoSnapShotPolicy resource.
type AutoSnapShotPolicyArgs struct {
	// Automatic snapshot policy name
	PolicyName pulumi.StringInput
	// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
	RepeatWeekdays pulumi.StringArrayInput
	// Automatic snapshot retention days.
	RetentionDays pulumi.IntInput
	// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
	TimePoints pulumi.StringArrayInput
}

func (AutoSnapShotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoSnapShotPolicyArgs)(nil)).Elem()
}

type AutoSnapShotPolicyInput interface {
	pulumi.Input

	ToAutoSnapShotPolicyOutput() AutoSnapShotPolicyOutput
	ToAutoSnapShotPolicyOutputWithContext(ctx context.Context) AutoSnapShotPolicyOutput
}

func (*AutoSnapShotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapShotPolicy)(nil)).Elem()
}

func (i *AutoSnapShotPolicy) ToAutoSnapShotPolicyOutput() AutoSnapShotPolicyOutput {
	return i.ToAutoSnapShotPolicyOutputWithContext(context.Background())
}

func (i *AutoSnapShotPolicy) ToAutoSnapShotPolicyOutputWithContext(ctx context.Context) AutoSnapShotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapShotPolicyOutput)
}

// AutoSnapShotPolicyArrayInput is an input type that accepts AutoSnapShotPolicyArray and AutoSnapShotPolicyArrayOutput values.
// You can construct a concrete instance of `AutoSnapShotPolicyArrayInput` via:
//
//	AutoSnapShotPolicyArray{ AutoSnapShotPolicyArgs{...} }
type AutoSnapShotPolicyArrayInput interface {
	pulumi.Input

	ToAutoSnapShotPolicyArrayOutput() AutoSnapShotPolicyArrayOutput
	ToAutoSnapShotPolicyArrayOutputWithContext(context.Context) AutoSnapShotPolicyArrayOutput
}

type AutoSnapShotPolicyArray []AutoSnapShotPolicyInput

func (AutoSnapShotPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapShotPolicy)(nil)).Elem()
}

func (i AutoSnapShotPolicyArray) ToAutoSnapShotPolicyArrayOutput() AutoSnapShotPolicyArrayOutput {
	return i.ToAutoSnapShotPolicyArrayOutputWithContext(context.Background())
}

func (i AutoSnapShotPolicyArray) ToAutoSnapShotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapShotPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapShotPolicyArrayOutput)
}

// AutoSnapShotPolicyMapInput is an input type that accepts AutoSnapShotPolicyMap and AutoSnapShotPolicyMapOutput values.
// You can construct a concrete instance of `AutoSnapShotPolicyMapInput` via:
//
//	AutoSnapShotPolicyMap{ "key": AutoSnapShotPolicyArgs{...} }
type AutoSnapShotPolicyMapInput interface {
	pulumi.Input

	ToAutoSnapShotPolicyMapOutput() AutoSnapShotPolicyMapOutput
	ToAutoSnapShotPolicyMapOutputWithContext(context.Context) AutoSnapShotPolicyMapOutput
}

type AutoSnapShotPolicyMap map[string]AutoSnapShotPolicyInput

func (AutoSnapShotPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapShotPolicy)(nil)).Elem()
}

func (i AutoSnapShotPolicyMap) ToAutoSnapShotPolicyMapOutput() AutoSnapShotPolicyMapOutput {
	return i.ToAutoSnapShotPolicyMapOutputWithContext(context.Background())
}

func (i AutoSnapShotPolicyMap) ToAutoSnapShotPolicyMapOutputWithContext(ctx context.Context) AutoSnapShotPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoSnapShotPolicyMapOutput)
}

type AutoSnapShotPolicyOutput struct{ *pulumi.OutputState }

func (AutoSnapShotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoSnapShotPolicy)(nil)).Elem()
}

func (o AutoSnapShotPolicyOutput) ToAutoSnapShotPolicyOutput() AutoSnapShotPolicyOutput {
	return o
}

func (o AutoSnapShotPolicyOutput) ToAutoSnapShotPolicyOutputWithContext(ctx context.Context) AutoSnapShotPolicyOutput {
	return o
}

// The number of database file systems set by the automatic snapshot policy.
func (o AutoSnapShotPolicyOutput) AppliedDbfsNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.IntOutput { return v.AppliedDbfsNumber }).(pulumi.IntOutput)
}

// The creation time of the resource
func (o AutoSnapShotPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Last modification time of automatic snapshot policy
func (o AutoSnapShotPolicyOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// Automatic snapshot policy ID
func (o AutoSnapShotPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// Automatic snapshot policy name
func (o AutoSnapShotPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// A collection of automatic snapshots performed on several days of the week. Value range: 1~7, for example, `1` means Monday.
func (o AutoSnapShotPolicyOutput) RepeatWeekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringArrayOutput { return v.RepeatWeekdays }).(pulumi.StringArrayOutput)
}

// Automatic snapshot retention days.
func (o AutoSnapShotPolicyOutput) RetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.IntOutput { return v.RetentionDays }).(pulumi.IntOutput)
}

// Automatic snapshot policy status
func (o AutoSnapShotPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Automatic snapshot policy status details
func (o AutoSnapShotPolicyOutput) StatusDetail() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringOutput { return v.StatusDetail }).(pulumi.StringOutput)
}

// The set of times at which the snapshot is taken on the day the automatic snapshot is executed. Value range: `00` to `23`, representing 24 time points from 00:00 to 23:00, for example, `01` indicates 01:00.
func (o AutoSnapShotPolicyOutput) TimePoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoSnapShotPolicy) pulumi.StringArrayOutput { return v.TimePoints }).(pulumi.StringArrayOutput)
}

type AutoSnapShotPolicyArrayOutput struct{ *pulumi.OutputState }

func (AutoSnapShotPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoSnapShotPolicy)(nil)).Elem()
}

func (o AutoSnapShotPolicyArrayOutput) ToAutoSnapShotPolicyArrayOutput() AutoSnapShotPolicyArrayOutput {
	return o
}

func (o AutoSnapShotPolicyArrayOutput) ToAutoSnapShotPolicyArrayOutputWithContext(ctx context.Context) AutoSnapShotPolicyArrayOutput {
	return o
}

func (o AutoSnapShotPolicyArrayOutput) Index(i pulumi.IntInput) AutoSnapShotPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoSnapShotPolicy {
		return vs[0].([]*AutoSnapShotPolicy)[vs[1].(int)]
	}).(AutoSnapShotPolicyOutput)
}

type AutoSnapShotPolicyMapOutput struct{ *pulumi.OutputState }

func (AutoSnapShotPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoSnapShotPolicy)(nil)).Elem()
}

func (o AutoSnapShotPolicyMapOutput) ToAutoSnapShotPolicyMapOutput() AutoSnapShotPolicyMapOutput {
	return o
}

func (o AutoSnapShotPolicyMapOutput) ToAutoSnapShotPolicyMapOutputWithContext(ctx context.Context) AutoSnapShotPolicyMapOutput {
	return o
}

func (o AutoSnapShotPolicyMapOutput) MapIndex(k pulumi.StringInput) AutoSnapShotPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoSnapShotPolicy {
		return vs[0].(map[string]*AutoSnapShotPolicy)[vs[1].(string)]
	}).(AutoSnapShotPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapShotPolicyInput)(nil)).Elem(), &AutoSnapShotPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapShotPolicyArrayInput)(nil)).Elem(), AutoSnapShotPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoSnapShotPolicyMapInput)(nil)).Elem(), AutoSnapShotPolicyMap{})
	pulumi.RegisterOutputType(AutoSnapShotPolicyOutput{})
	pulumi.RegisterOutputType(AutoSnapShotPolicyArrayOutput{})
	pulumi.RegisterOutputType(AutoSnapShotPolicyMapOutput{})
}
