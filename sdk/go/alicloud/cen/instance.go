// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a CEN instance resource. Cloud Enterprise Network (CEN) is a service that allows you to create a global network for rapidly building a distributed business system with a hybrid cloud computing solution. CEN enables you to build a secure, private, and enterprise-class interconnected network between VPCs in different regions and your local data centers. CEN provides enterprise-class scalability that automatically responds to your dynamic computing requirements.
//
// For information about CEN and how to use it, see [What is Cloud Enterprise Network](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-createcen).
//
// > **NOTE:** Available since v1.15.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewInstance(ctx, "example", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String("tf_example"),
//				Description:     pulumi.String("an example for cen"),
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
// CEN instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/instance:Instance example cen-abc123456
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
	CenInstanceName pulumi.StringOutput `pulumi:"cenInstanceName"`
	// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
	//
	// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
	ProtectionLevel pulumi.StringOutput `pulumi:"protectionLevel"`
	// The Cen Instance current status.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:cen/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:cen/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
	CenInstanceName *string `pulumi:"cenInstanceName"`
	// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
	Description *string `pulumi:"description"`
	// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
	//
	// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
	Name *string `pulumi:"name"`
	// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The Cen Instance current status.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type InstanceState struct {
	// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
	CenInstanceName pulumi.StringPtrInput
	// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
	Description pulumi.StringPtrInput
	// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
	//
	// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
	Name pulumi.StringPtrInput
	// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
	ProtectionLevel pulumi.StringPtrInput
	// The Cen Instance current status.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
	CenInstanceName *string `pulumi:"cenInstanceName"`
	// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
	Description *string `pulumi:"description"`
	// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
	//
	// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
	Name *string `pulumi:"name"`
	// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
	CenInstanceName pulumi.StringPtrInput
	// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
	Description pulumi.StringPtrInput
	// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
	//
	// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
	Name pulumi.StringPtrInput
	// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
	ProtectionLevel pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: i.ToInstanceOutputWithContext(ctx).OutputState,
	}
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

func (i InstanceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Instance] {
	return pulumix.Output[[]*Instance]{
		OutputState: i.ToInstanceArrayOutputWithContext(ctx).OutputState,
	}
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

func (i InstanceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Instance] {
	return pulumix.Output[map[string]*Instance]{
		OutputState: i.ToInstanceMapOutputWithContext(ctx).OutputState,
	}
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*Instance] {
	return pulumix.Output[*Instance]{
		OutputState: o.OutputState,
	}
}

// The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.
func (o InstanceOutput) CenInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CenInstanceName }).(pulumi.StringOutput)
}

// The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.
func (o InstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Field `name` has been deprecated from version 1.98.0. Use `cenInstanceName` instead.
//
// Deprecated: attribute 'name' has been deprecated from version 1.98.0. Use 'cen_instance_name' instead.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates the allowed level of CIDR block overlapping. Default value: `REDUCE`: Overlapping CIDR blocks are allowed. However, the overlapping CIDR blocks cannot be identical.
func (o InstanceOutput) ProtectionLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProtectionLevel }).(pulumi.StringOutput)
}

// The Cen Instance current status.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Instance] {
	return pulumix.Output[[]*Instance]{
		OutputState: o.OutputState,
	}
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Instance] {
	return pulumix.Output[map[string]*Instance]{
		OutputState: o.OutputState,
	}
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
