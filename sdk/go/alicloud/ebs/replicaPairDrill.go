// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a EBS Replica Pair Drill resource.
//
// For information about EBS Replica Pair Drill and how to use it, see [What is Replica Pair Drill](https://www.alibabacloud.com/help/en/).
//
// > **NOTE:** Available since v1.215.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ebs"
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
//			_, err := ebs.NewReplicaPairDrill(ctx, "default", &ebs.ReplicaPairDrillArgs{
//				PairId: pulumi.String("pair-cn-wwo3kjfq5001"),
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
// EBS Replica Pair Drill can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ebs/replicaPairDrill:ReplicaPairDrill example <pair_id>:<replica_pair_drill_id>
// ```
type ReplicaPairDrill struct {
	pulumi.CustomResourceState

	// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
	PairId pulumi.StringOutput `pulumi:"pairId"`
	// The first ID of the resource.
	ReplicaPairDrillId pulumi.StringOutput `pulumi:"replicaPairDrillId"`
	// Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewReplicaPairDrill registers a new resource with the given unique name, arguments, and options.
func NewReplicaPairDrill(ctx *pulumi.Context,
	name string, args *ReplicaPairDrillArgs, opts ...pulumi.ResourceOption) (*ReplicaPairDrill, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PairId == nil {
		return nil, errors.New("invalid value for required argument 'PairId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReplicaPairDrill
	err := ctx.RegisterResource("alicloud:ebs/replicaPairDrill:ReplicaPairDrill", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaPairDrill gets an existing ReplicaPairDrill resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaPairDrill(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaPairDrillState, opts ...pulumi.ResourceOption) (*ReplicaPairDrill, error) {
	var resource ReplicaPairDrill
	err := ctx.ReadResource("alicloud:ebs/replicaPairDrill:ReplicaPairDrill", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaPairDrill resources.
type replicaPairDrillState struct {
	// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
	PairId *string `pulumi:"pairId"`
	// The first ID of the resource.
	ReplicaPairDrillId *string `pulumi:"replicaPairDrillId"`
	// Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
	Status *string `pulumi:"status"`
}

type ReplicaPairDrillState struct {
	// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
	PairId pulumi.StringPtrInput
	// The first ID of the resource.
	ReplicaPairDrillId pulumi.StringPtrInput
	// Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
	Status pulumi.StringPtrInput
}

func (ReplicaPairDrillState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaPairDrillState)(nil)).Elem()
}

type replicaPairDrillArgs struct {
	// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
	PairId string `pulumi:"pairId"`
}

// The set of arguments for constructing a ReplicaPairDrill resource.
type ReplicaPairDrillArgs struct {
	// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
	PairId pulumi.StringInput
}

func (ReplicaPairDrillArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaPairDrillArgs)(nil)).Elem()
}

type ReplicaPairDrillInput interface {
	pulumi.Input

	ToReplicaPairDrillOutput() ReplicaPairDrillOutput
	ToReplicaPairDrillOutputWithContext(ctx context.Context) ReplicaPairDrillOutput
}

func (*ReplicaPairDrill) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaPairDrill)(nil)).Elem()
}

func (i *ReplicaPairDrill) ToReplicaPairDrillOutput() ReplicaPairDrillOutput {
	return i.ToReplicaPairDrillOutputWithContext(context.Background())
}

func (i *ReplicaPairDrill) ToReplicaPairDrillOutputWithContext(ctx context.Context) ReplicaPairDrillOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaPairDrillOutput)
}

// ReplicaPairDrillArrayInput is an input type that accepts ReplicaPairDrillArray and ReplicaPairDrillArrayOutput values.
// You can construct a concrete instance of `ReplicaPairDrillArrayInput` via:
//
//	ReplicaPairDrillArray{ ReplicaPairDrillArgs{...} }
type ReplicaPairDrillArrayInput interface {
	pulumi.Input

	ToReplicaPairDrillArrayOutput() ReplicaPairDrillArrayOutput
	ToReplicaPairDrillArrayOutputWithContext(context.Context) ReplicaPairDrillArrayOutput
}

type ReplicaPairDrillArray []ReplicaPairDrillInput

func (ReplicaPairDrillArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaPairDrill)(nil)).Elem()
}

func (i ReplicaPairDrillArray) ToReplicaPairDrillArrayOutput() ReplicaPairDrillArrayOutput {
	return i.ToReplicaPairDrillArrayOutputWithContext(context.Background())
}

func (i ReplicaPairDrillArray) ToReplicaPairDrillArrayOutputWithContext(ctx context.Context) ReplicaPairDrillArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaPairDrillArrayOutput)
}

// ReplicaPairDrillMapInput is an input type that accepts ReplicaPairDrillMap and ReplicaPairDrillMapOutput values.
// You can construct a concrete instance of `ReplicaPairDrillMapInput` via:
//
//	ReplicaPairDrillMap{ "key": ReplicaPairDrillArgs{...} }
type ReplicaPairDrillMapInput interface {
	pulumi.Input

	ToReplicaPairDrillMapOutput() ReplicaPairDrillMapOutput
	ToReplicaPairDrillMapOutputWithContext(context.Context) ReplicaPairDrillMapOutput
}

type ReplicaPairDrillMap map[string]ReplicaPairDrillInput

func (ReplicaPairDrillMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaPairDrill)(nil)).Elem()
}

func (i ReplicaPairDrillMap) ToReplicaPairDrillMapOutput() ReplicaPairDrillMapOutput {
	return i.ToReplicaPairDrillMapOutputWithContext(context.Background())
}

func (i ReplicaPairDrillMap) ToReplicaPairDrillMapOutputWithContext(ctx context.Context) ReplicaPairDrillMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaPairDrillMapOutput)
}

type ReplicaPairDrillOutput struct{ *pulumi.OutputState }

func (ReplicaPairDrillOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaPairDrill)(nil)).Elem()
}

func (o ReplicaPairDrillOutput) ToReplicaPairDrillOutput() ReplicaPairDrillOutput {
	return o
}

func (o ReplicaPairDrillOutput) ToReplicaPairDrillOutputWithContext(ctx context.Context) ReplicaPairDrillOutput {
	return o
}

// Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
func (o ReplicaPairDrillOutput) PairId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaPairDrill) pulumi.StringOutput { return v.PairId }).(pulumi.StringOutput)
}

// The first ID of the resource.
func (o ReplicaPairDrillOutput) ReplicaPairDrillId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaPairDrill) pulumi.StringOutput { return v.ReplicaPairDrillId }).(pulumi.StringOutput)
}

// Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
func (o ReplicaPairDrillOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicaPairDrill) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ReplicaPairDrillArrayOutput struct{ *pulumi.OutputState }

func (ReplicaPairDrillArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicaPairDrill)(nil)).Elem()
}

func (o ReplicaPairDrillArrayOutput) ToReplicaPairDrillArrayOutput() ReplicaPairDrillArrayOutput {
	return o
}

func (o ReplicaPairDrillArrayOutput) ToReplicaPairDrillArrayOutputWithContext(ctx context.Context) ReplicaPairDrillArrayOutput {
	return o
}

func (o ReplicaPairDrillArrayOutput) Index(i pulumi.IntInput) ReplicaPairDrillOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReplicaPairDrill {
		return vs[0].([]*ReplicaPairDrill)[vs[1].(int)]
	}).(ReplicaPairDrillOutput)
}

type ReplicaPairDrillMapOutput struct{ *pulumi.OutputState }

func (ReplicaPairDrillMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicaPairDrill)(nil)).Elem()
}

func (o ReplicaPairDrillMapOutput) ToReplicaPairDrillMapOutput() ReplicaPairDrillMapOutput {
	return o
}

func (o ReplicaPairDrillMapOutput) ToReplicaPairDrillMapOutputWithContext(ctx context.Context) ReplicaPairDrillMapOutput {
	return o
}

func (o ReplicaPairDrillMapOutput) MapIndex(k pulumi.StringInput) ReplicaPairDrillOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReplicaPairDrill {
		return vs[0].(map[string]*ReplicaPairDrill)[vs[1].(string)]
	}).(ReplicaPairDrillOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaPairDrillInput)(nil)).Elem(), &ReplicaPairDrill{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaPairDrillArrayInput)(nil)).Elem(), ReplicaPairDrillArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicaPairDrillMapInput)(nil)).Elem(), ReplicaPairDrillMap{})
	pulumi.RegisterOutputType(ReplicaPairDrillOutput{})
	pulumi.RegisterOutputType(ReplicaPairDrillArrayOutput{})
	pulumi.RegisterOutputType(ReplicaPairDrillMapOutput{})
}
