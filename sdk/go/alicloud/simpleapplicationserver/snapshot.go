// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleapplicationserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Simple Application Server Snapshot resource.
//
// For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).
//
// > **NOTE:** Available since v1.143.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/simpleapplicationserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultImages, err := simpleapplicationserver.GetImages(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultServerPlans, err := simpleapplicationserver.GetServerPlans(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := simpleapplicationserver.NewInstance(ctx, "defaultInstance", &simpleapplicationserver.InstanceArgs{
//				PaymentType:  pulumi.String("Subscription"),
//				PlanId:       *pulumi.String(defaultServerPlans.Plans[0].Id),
//				InstanceName: pulumi.String(name),
//				ImageId:      *pulumi.String(defaultImages.Images[0].Id),
//				Period:       pulumi.Int(1),
//				DataDiskSize: pulumi.Int(100),
//			})
//			if err != nil {
//				return err
//			}
//			defaultServerDisks := simpleapplicationserver.GetServerDisksOutput(ctx, simpleapplicationserver.GetServerDisksOutputArgs{
//				InstanceId: defaultInstance.ID(),
//			}, nil)
//			_, err = simpleapplicationserver.NewSnapshot(ctx, "defaultSnapshot", &simpleapplicationserver.SnapshotArgs{
//				DiskId: defaultServerDisks.ApplyT(func(defaultServerDisks simpleapplicationserver.GetServerDisksResult) (*string, error) {
//					return &defaultServerDisks.Ids[0], nil
//				}).(pulumi.StringPtrOutput),
//				SnapshotName: pulumi.String(name),
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
// Simple Application Server Snapshot can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example <id>
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The ID of the disk.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.SnapshotName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("alicloud:simpleapplicationserver/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("alicloud:simpleapplicationserver/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The ID of the disk.
	DiskId *string `pulumi:"diskId"`
	// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
	SnapshotName *string `pulumi:"snapshotName"`
	// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
	Status *string `pulumi:"status"`
}

type SnapshotState struct {
	// The ID of the disk.
	DiskId pulumi.StringPtrInput
	// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
	SnapshotName pulumi.StringPtrInput
	// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
	Status pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The ID of the disk.
	DiskId string `pulumi:"diskId"`
	// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
	SnapshotName string `pulumi:"snapshotName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The ID of the disk.
	DiskId pulumi.StringInput
	// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
	SnapshotName pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The ID of the disk.
func (o SnapshotOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

// The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
func (o SnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
