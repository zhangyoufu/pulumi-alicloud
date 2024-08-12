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

// Provides a OSS Bucket Data Redundancy Transition resource. Create a storage redundancy transition task to convert local redundant storage(LRS) to zone redundant storage(ZRS).
//
// For information about OSS Bucket Data Redundancy Transition and how to use it, see [What is Bucket Data Redundancy Transition](https://www.alibabacloud.com/help/en/oss/developer-reference/createbucketdataredundancytransition).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
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
//			createBucket, err := oss.NewBucket(ctx, "CreateBucket", &oss.BucketArgs{
//				StorageClass: pulumi.String("Standard"),
//				Bucket:       pulumi.Sprintf("%v-%v", name, _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oss.NewBucketDataRedundancyTransition(ctx, "default", &oss.BucketDataRedundancyTransitionArgs{
//				Bucket: createBucket.Bucket,
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
// OSS Bucket Data Redundancy Transition can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oss/bucketDataRedundancyTransition:BucketDataRedundancyTransition example <bucket>:<task_id>
// ```
type BucketDataRedundancyTransition struct {
	pulumi.CustomResourceState

	// Storage space name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Stores the creation time of the redundant transformation task.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
	Status pulumi.StringOutput `pulumi:"status"`
	// Unique identification of the storage redundancy conversion task.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
}

// NewBucketDataRedundancyTransition registers a new resource with the given unique name, arguments, and options.
func NewBucketDataRedundancyTransition(ctx *pulumi.Context,
	name string, args *BucketDataRedundancyTransitionArgs, opts ...pulumi.ResourceOption) (*BucketDataRedundancyTransition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketDataRedundancyTransition
	err := ctx.RegisterResource("alicloud:oss/bucketDataRedundancyTransition:BucketDataRedundancyTransition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketDataRedundancyTransition gets an existing BucketDataRedundancyTransition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketDataRedundancyTransition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketDataRedundancyTransitionState, opts ...pulumi.ResourceOption) (*BucketDataRedundancyTransition, error) {
	var resource BucketDataRedundancyTransition
	err := ctx.ReadResource("alicloud:oss/bucketDataRedundancyTransition:BucketDataRedundancyTransition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketDataRedundancyTransition resources.
type bucketDataRedundancyTransitionState struct {
	// Storage space name.
	Bucket *string `pulumi:"bucket"`
	// Stores the creation time of the redundant transformation task.
	CreateTime *string `pulumi:"createTime"`
	// Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
	Status *string `pulumi:"status"`
	// Unique identification of the storage redundancy conversion task.
	TaskId *string `pulumi:"taskId"`
}

type BucketDataRedundancyTransitionState struct {
	// Storage space name.
	Bucket pulumi.StringPtrInput
	// Stores the creation time of the redundant transformation task.
	CreateTime pulumi.StringPtrInput
	// Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
	Status pulumi.StringPtrInput
	// Unique identification of the storage redundancy conversion task.
	TaskId pulumi.StringPtrInput
}

func (BucketDataRedundancyTransitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketDataRedundancyTransitionState)(nil)).Elem()
}

type bucketDataRedundancyTransitionArgs struct {
	// Storage space name.
	Bucket string `pulumi:"bucket"`
}

// The set of arguments for constructing a BucketDataRedundancyTransition resource.
type BucketDataRedundancyTransitionArgs struct {
	// Storage space name.
	Bucket pulumi.StringInput
}

func (BucketDataRedundancyTransitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketDataRedundancyTransitionArgs)(nil)).Elem()
}

type BucketDataRedundancyTransitionInput interface {
	pulumi.Input

	ToBucketDataRedundancyTransitionOutput() BucketDataRedundancyTransitionOutput
	ToBucketDataRedundancyTransitionOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionOutput
}

func (*BucketDataRedundancyTransition) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketDataRedundancyTransition)(nil)).Elem()
}

func (i *BucketDataRedundancyTransition) ToBucketDataRedundancyTransitionOutput() BucketDataRedundancyTransitionOutput {
	return i.ToBucketDataRedundancyTransitionOutputWithContext(context.Background())
}

func (i *BucketDataRedundancyTransition) ToBucketDataRedundancyTransitionOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDataRedundancyTransitionOutput)
}

// BucketDataRedundancyTransitionArrayInput is an input type that accepts BucketDataRedundancyTransitionArray and BucketDataRedundancyTransitionArrayOutput values.
// You can construct a concrete instance of `BucketDataRedundancyTransitionArrayInput` via:
//
//	BucketDataRedundancyTransitionArray{ BucketDataRedundancyTransitionArgs{...} }
type BucketDataRedundancyTransitionArrayInput interface {
	pulumi.Input

	ToBucketDataRedundancyTransitionArrayOutput() BucketDataRedundancyTransitionArrayOutput
	ToBucketDataRedundancyTransitionArrayOutputWithContext(context.Context) BucketDataRedundancyTransitionArrayOutput
}

type BucketDataRedundancyTransitionArray []BucketDataRedundancyTransitionInput

func (BucketDataRedundancyTransitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketDataRedundancyTransition)(nil)).Elem()
}

func (i BucketDataRedundancyTransitionArray) ToBucketDataRedundancyTransitionArrayOutput() BucketDataRedundancyTransitionArrayOutput {
	return i.ToBucketDataRedundancyTransitionArrayOutputWithContext(context.Background())
}

func (i BucketDataRedundancyTransitionArray) ToBucketDataRedundancyTransitionArrayOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDataRedundancyTransitionArrayOutput)
}

// BucketDataRedundancyTransitionMapInput is an input type that accepts BucketDataRedundancyTransitionMap and BucketDataRedundancyTransitionMapOutput values.
// You can construct a concrete instance of `BucketDataRedundancyTransitionMapInput` via:
//
//	BucketDataRedundancyTransitionMap{ "key": BucketDataRedundancyTransitionArgs{...} }
type BucketDataRedundancyTransitionMapInput interface {
	pulumi.Input

	ToBucketDataRedundancyTransitionMapOutput() BucketDataRedundancyTransitionMapOutput
	ToBucketDataRedundancyTransitionMapOutputWithContext(context.Context) BucketDataRedundancyTransitionMapOutput
}

type BucketDataRedundancyTransitionMap map[string]BucketDataRedundancyTransitionInput

func (BucketDataRedundancyTransitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketDataRedundancyTransition)(nil)).Elem()
}

func (i BucketDataRedundancyTransitionMap) ToBucketDataRedundancyTransitionMapOutput() BucketDataRedundancyTransitionMapOutput {
	return i.ToBucketDataRedundancyTransitionMapOutputWithContext(context.Background())
}

func (i BucketDataRedundancyTransitionMap) ToBucketDataRedundancyTransitionMapOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDataRedundancyTransitionMapOutput)
}

type BucketDataRedundancyTransitionOutput struct{ *pulumi.OutputState }

func (BucketDataRedundancyTransitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketDataRedundancyTransition)(nil)).Elem()
}

func (o BucketDataRedundancyTransitionOutput) ToBucketDataRedundancyTransitionOutput() BucketDataRedundancyTransitionOutput {
	return o
}

func (o BucketDataRedundancyTransitionOutput) ToBucketDataRedundancyTransitionOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionOutput {
	return o
}

// Storage space name.
func (o BucketDataRedundancyTransitionOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketDataRedundancyTransition) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Stores the creation time of the redundant transformation task.
func (o BucketDataRedundancyTransitionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketDataRedundancyTransition) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
func (o BucketDataRedundancyTransitionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketDataRedundancyTransition) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Unique identification of the storage redundancy conversion task.
func (o BucketDataRedundancyTransitionOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketDataRedundancyTransition) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

type BucketDataRedundancyTransitionArrayOutput struct{ *pulumi.OutputState }

func (BucketDataRedundancyTransitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketDataRedundancyTransition)(nil)).Elem()
}

func (o BucketDataRedundancyTransitionArrayOutput) ToBucketDataRedundancyTransitionArrayOutput() BucketDataRedundancyTransitionArrayOutput {
	return o
}

func (o BucketDataRedundancyTransitionArrayOutput) ToBucketDataRedundancyTransitionArrayOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionArrayOutput {
	return o
}

func (o BucketDataRedundancyTransitionArrayOutput) Index(i pulumi.IntInput) BucketDataRedundancyTransitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketDataRedundancyTransition {
		return vs[0].([]*BucketDataRedundancyTransition)[vs[1].(int)]
	}).(BucketDataRedundancyTransitionOutput)
}

type BucketDataRedundancyTransitionMapOutput struct{ *pulumi.OutputState }

func (BucketDataRedundancyTransitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketDataRedundancyTransition)(nil)).Elem()
}

func (o BucketDataRedundancyTransitionMapOutput) ToBucketDataRedundancyTransitionMapOutput() BucketDataRedundancyTransitionMapOutput {
	return o
}

func (o BucketDataRedundancyTransitionMapOutput) ToBucketDataRedundancyTransitionMapOutputWithContext(ctx context.Context) BucketDataRedundancyTransitionMapOutput {
	return o
}

func (o BucketDataRedundancyTransitionMapOutput) MapIndex(k pulumi.StringInput) BucketDataRedundancyTransitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketDataRedundancyTransition {
		return vs[0].(map[string]*BucketDataRedundancyTransition)[vs[1].(string)]
	}).(BucketDataRedundancyTransitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketDataRedundancyTransitionInput)(nil)).Elem(), &BucketDataRedundancyTransition{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketDataRedundancyTransitionArrayInput)(nil)).Elem(), BucketDataRedundancyTransitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketDataRedundancyTransitionMapInput)(nil)).Elem(), BucketDataRedundancyTransitionMap{})
	pulumi.RegisterOutputType(BucketDataRedundancyTransitionOutput{})
	pulumi.RegisterOutputType(BucketDataRedundancyTransitionArrayOutput{})
	pulumi.RegisterOutputType(BucketDataRedundancyTransitionMapOutput{})
}
