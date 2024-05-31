// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// OSS Bucket Access Monitor can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oss/bucketAccessMonitor:BucketAccessMonitor example <id>
// ```
type BucketAccessMonitor struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBucketAccessMonitor registers a new resource with the given unique name, arguments, and options.
func NewBucketAccessMonitor(ctx *pulumi.Context,
	name string, args *BucketAccessMonitorArgs, opts ...pulumi.ResourceOption) (*BucketAccessMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketAccessMonitor
	err := ctx.RegisterResource("alicloud:oss/bucketAccessMonitor:BucketAccessMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketAccessMonitor gets an existing BucketAccessMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketAccessMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketAccessMonitorState, opts ...pulumi.ResourceOption) (*BucketAccessMonitor, error) {
	var resource BucketAccessMonitor
	err := ctx.ReadResource("alicloud:oss/bucketAccessMonitor:BucketAccessMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketAccessMonitor resources.
type bucketAccessMonitorState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
	Status *string `pulumi:"status"`
}

type BucketAccessMonitorState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
	Status pulumi.StringPtrInput
}

func (BucketAccessMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessMonitorState)(nil)).Elem()
}

type bucketAccessMonitorArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a BucketAccessMonitor resource.
type BucketAccessMonitorArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
	Status pulumi.StringInput
}

func (BucketAccessMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessMonitorArgs)(nil)).Elem()
}

type BucketAccessMonitorInput interface {
	pulumi.Input

	ToBucketAccessMonitorOutput() BucketAccessMonitorOutput
	ToBucketAccessMonitorOutputWithContext(ctx context.Context) BucketAccessMonitorOutput
}

func (*BucketAccessMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessMonitor)(nil)).Elem()
}

func (i *BucketAccessMonitor) ToBucketAccessMonitorOutput() BucketAccessMonitorOutput {
	return i.ToBucketAccessMonitorOutputWithContext(context.Background())
}

func (i *BucketAccessMonitor) ToBucketAccessMonitorOutputWithContext(ctx context.Context) BucketAccessMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessMonitorOutput)
}

// BucketAccessMonitorArrayInput is an input type that accepts BucketAccessMonitorArray and BucketAccessMonitorArrayOutput values.
// You can construct a concrete instance of `BucketAccessMonitorArrayInput` via:
//
//	BucketAccessMonitorArray{ BucketAccessMonitorArgs{...} }
type BucketAccessMonitorArrayInput interface {
	pulumi.Input

	ToBucketAccessMonitorArrayOutput() BucketAccessMonitorArrayOutput
	ToBucketAccessMonitorArrayOutputWithContext(context.Context) BucketAccessMonitorArrayOutput
}

type BucketAccessMonitorArray []BucketAccessMonitorInput

func (BucketAccessMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccessMonitor)(nil)).Elem()
}

func (i BucketAccessMonitorArray) ToBucketAccessMonitorArrayOutput() BucketAccessMonitorArrayOutput {
	return i.ToBucketAccessMonitorArrayOutputWithContext(context.Background())
}

func (i BucketAccessMonitorArray) ToBucketAccessMonitorArrayOutputWithContext(ctx context.Context) BucketAccessMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessMonitorArrayOutput)
}

// BucketAccessMonitorMapInput is an input type that accepts BucketAccessMonitorMap and BucketAccessMonitorMapOutput values.
// You can construct a concrete instance of `BucketAccessMonitorMapInput` via:
//
//	BucketAccessMonitorMap{ "key": BucketAccessMonitorArgs{...} }
type BucketAccessMonitorMapInput interface {
	pulumi.Input

	ToBucketAccessMonitorMapOutput() BucketAccessMonitorMapOutput
	ToBucketAccessMonitorMapOutputWithContext(context.Context) BucketAccessMonitorMapOutput
}

type BucketAccessMonitorMap map[string]BucketAccessMonitorInput

func (BucketAccessMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccessMonitor)(nil)).Elem()
}

func (i BucketAccessMonitorMap) ToBucketAccessMonitorMapOutput() BucketAccessMonitorMapOutput {
	return i.ToBucketAccessMonitorMapOutputWithContext(context.Background())
}

func (i BucketAccessMonitorMap) ToBucketAccessMonitorMapOutputWithContext(ctx context.Context) BucketAccessMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessMonitorMapOutput)
}

type BucketAccessMonitorOutput struct{ *pulumi.OutputState }

func (BucketAccessMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessMonitor)(nil)).Elem()
}

func (o BucketAccessMonitorOutput) ToBucketAccessMonitorOutput() BucketAccessMonitorOutput {
	return o
}

func (o BucketAccessMonitorOutput) ToBucketAccessMonitorOutputWithContext(ctx context.Context) BucketAccessMonitorOutput {
	return o
}

// The name of the bucket.
func (o BucketAccessMonitorOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessMonitor) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
func (o BucketAccessMonitorOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessMonitor) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BucketAccessMonitorArrayOutput struct{ *pulumi.OutputState }

func (BucketAccessMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketAccessMonitor)(nil)).Elem()
}

func (o BucketAccessMonitorArrayOutput) ToBucketAccessMonitorArrayOutput() BucketAccessMonitorArrayOutput {
	return o
}

func (o BucketAccessMonitorArrayOutput) ToBucketAccessMonitorArrayOutputWithContext(ctx context.Context) BucketAccessMonitorArrayOutput {
	return o
}

func (o BucketAccessMonitorArrayOutput) Index(i pulumi.IntInput) BucketAccessMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketAccessMonitor {
		return vs[0].([]*BucketAccessMonitor)[vs[1].(int)]
	}).(BucketAccessMonitorOutput)
}

type BucketAccessMonitorMapOutput struct{ *pulumi.OutputState }

func (BucketAccessMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketAccessMonitor)(nil)).Elem()
}

func (o BucketAccessMonitorMapOutput) ToBucketAccessMonitorMapOutput() BucketAccessMonitorMapOutput {
	return o
}

func (o BucketAccessMonitorMapOutput) ToBucketAccessMonitorMapOutputWithContext(ctx context.Context) BucketAccessMonitorMapOutput {
	return o
}

func (o BucketAccessMonitorMapOutput) MapIndex(k pulumi.StringInput) BucketAccessMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketAccessMonitor {
		return vs[0].(map[string]*BucketAccessMonitor)[vs[1].(string)]
	}).(BucketAccessMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessMonitorInput)(nil)).Elem(), &BucketAccessMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessMonitorArrayInput)(nil)).Elem(), BucketAccessMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessMonitorMapInput)(nil)).Elem(), BucketAccessMonitorMap{})
	pulumi.RegisterOutputType(BucketAccessMonitorOutput{})
	pulumi.RegisterOutputType(BucketAccessMonitorArrayOutput{})
	pulumi.RegisterOutputType(BucketAccessMonitorMapOutput{})
}
