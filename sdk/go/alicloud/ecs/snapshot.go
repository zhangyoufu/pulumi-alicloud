// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := ecs.NewSnapshot(ctx, "snapshot", &ecs.SnapshotArgs{
//				DiskId:      pulumi.Any(alicloud_disk_attachment.InstanceAttachment.Disk_id),
//				Description: pulumi.String("this snapshot is created for testing"),
//				Tags: pulumi.AnyMap{
//					"version": pulumi.Any("1.2"),
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
//
//	$ pulumi import alicloud:ecs/snapshot:Snapshot snapshot s-abc1234567890000
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The source disk ID.
	DiskId                     pulumi.StringOutput  `pulumi:"diskId"`
	Force                      pulumi.BoolPtrOutput `pulumi:"force"`
	InstantAccess              pulumi.BoolPtrOutput `pulumi:"instantAccess"`
	InstantAccessRetentionDays pulumi.IntPtrOutput  `pulumi:"instantAccessRetentionDays"`
	// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	RetentionDays   pulumi.IntPtrOutput    `pulumi:"retentionDays"`
	SnapshotName    pulumi.StringOutput    `pulumi:"snapshotName"`
	Status          pulumi.StringOutput    `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	var resource Snapshot
	err := ctx.RegisterResource("alicloud:ecs/snapshot:Snapshot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:ecs/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	Category *string `pulumi:"category"`
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// The source disk ID.
	DiskId                     *string `pulumi:"diskId"`
	Force                      *bool   `pulumi:"force"`
	InstantAccess              *bool   `pulumi:"instantAccess"`
	InstantAccessRetentionDays *int    `pulumi:"instantAccessRetentionDays"`
	// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name *string `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	RetentionDays   *int    `pulumi:"retentionDays"`
	SnapshotName    *string `pulumi:"snapshotName"`
	Status          *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type SnapshotState struct {
	Category pulumi.StringPtrInput
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// The source disk ID.
	DiskId                     pulumi.StringPtrInput
	Force                      pulumi.BoolPtrInput
	InstantAccess              pulumi.BoolPtrInput
	InstantAccessRetentionDays pulumi.IntPtrInput
	// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	RetentionDays   pulumi.IntPtrInput
	SnapshotName    pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	Category *string `pulumi:"category"`
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	// The source disk ID.
	DiskId                     string `pulumi:"diskId"`
	Force                      *bool  `pulumi:"force"`
	InstantAccess              *bool  `pulumi:"instantAccess"`
	InstantAccessRetentionDays *int   `pulumi:"instantAccessRetentionDays"`
	// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name *string `pulumi:"name"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	RetentionDays   *int    `pulumi:"retentionDays"`
	SnapshotName    *string `pulumi:"snapshotName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	Category pulumi.StringPtrInput
	// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	// The source disk ID.
	DiskId                     pulumi.StringInput
	Force                      pulumi.BoolPtrInput
	InstantAccess              pulumi.BoolPtrInput
	InstantAccessRetentionDays pulumi.IntPtrInput
	// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
	Name pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	RetentionDays   pulumi.IntPtrInput
	SnapshotName    pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
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

func (o SnapshotOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The source disk ID.
func (o SnapshotOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o SnapshotOutput) InstantAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolPtrOutput { return v.InstantAccess }).(pulumi.BoolPtrOutput)
}

func (o SnapshotOutput) InstantAccessRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.InstantAccessRetentionDays }).(pulumi.IntPtrOutput)
}

// The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
// It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o SnapshotOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o SnapshotOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

func (o SnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o SnapshotOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
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
