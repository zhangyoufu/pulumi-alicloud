// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Resource Manager Resource Share resource.
//
// For information about Resource Manager Resource Share and how to use it, see [What is Resource Share](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
//
// > **NOTE:** Available since v1.111.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tfexample"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := resourcemanager.NewResourceShare(ctx, "example", &resourcemanager.ResourceShareArgs{
//				ResourceShareName: pulumi.String(name),
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
// Resource Manager Resource Share can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:resourcemanager/resourceShare:ResourceShare example <id>
//
// ```
type ResourceShare struct {
	pulumi.CustomResourceState

	// The name of resource share.
	ResourceShareName pulumi.StringOutput `pulumi:"resourceShareName"`
	// The owner of resource share.
	ResourceShareOwner pulumi.StringOutput `pulumi:"resourceShareOwner"`
	// The status of resource share.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewResourceShare registers a new resource with the given unique name, arguments, and options.
func NewResourceShare(ctx *pulumi.Context,
	name string, args *ResourceShareArgs, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceShareName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceShareName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceShare
	err := ctx.RegisterResource("alicloud:resourcemanager/resourceShare:ResourceShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceShare gets an existing ResourceShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceShareState, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	var resource ResourceShare
	err := ctx.ReadResource("alicloud:resourcemanager/resourceShare:ResourceShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceShare resources.
type resourceShareState struct {
	// The name of resource share.
	ResourceShareName *string `pulumi:"resourceShareName"`
	// The owner of resource share.
	ResourceShareOwner *string `pulumi:"resourceShareOwner"`
	// The status of resource share.
	Status *string `pulumi:"status"`
}

type ResourceShareState struct {
	// The name of resource share.
	ResourceShareName pulumi.StringPtrInput
	// The owner of resource share.
	ResourceShareOwner pulumi.StringPtrInput
	// The status of resource share.
	Status pulumi.StringPtrInput
}

func (ResourceShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareState)(nil)).Elem()
}

type resourceShareArgs struct {
	// The name of resource share.
	ResourceShareName string `pulumi:"resourceShareName"`
}

// The set of arguments for constructing a ResourceShare resource.
type ResourceShareArgs struct {
	// The name of resource share.
	ResourceShareName pulumi.StringInput
}

func (ResourceShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareArgs)(nil)).Elem()
}

type ResourceShareInput interface {
	pulumi.Input

	ToResourceShareOutput() ResourceShareOutput
	ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput
}

func (*ResourceShare) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceShare)(nil)).Elem()
}

func (i *ResourceShare) ToResourceShareOutput() ResourceShareOutput {
	return i.ToResourceShareOutputWithContext(context.Background())
}

func (i *ResourceShare) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareOutput)
}

func (i *ResourceShare) ToOutput(ctx context.Context) pulumix.Output[*ResourceShare] {
	return pulumix.Output[*ResourceShare]{
		OutputState: i.ToResourceShareOutputWithContext(ctx).OutputState,
	}
}

// ResourceShareArrayInput is an input type that accepts ResourceShareArray and ResourceShareArrayOutput values.
// You can construct a concrete instance of `ResourceShareArrayInput` via:
//
//	ResourceShareArray{ ResourceShareArgs{...} }
type ResourceShareArrayInput interface {
	pulumi.Input

	ToResourceShareArrayOutput() ResourceShareArrayOutput
	ToResourceShareArrayOutputWithContext(context.Context) ResourceShareArrayOutput
}

type ResourceShareArray []ResourceShareInput

func (ResourceShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceShare)(nil)).Elem()
}

func (i ResourceShareArray) ToResourceShareArrayOutput() ResourceShareArrayOutput {
	return i.ToResourceShareArrayOutputWithContext(context.Background())
}

func (i ResourceShareArray) ToResourceShareArrayOutputWithContext(ctx context.Context) ResourceShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareArrayOutput)
}

func (i ResourceShareArray) ToOutput(ctx context.Context) pulumix.Output[[]*ResourceShare] {
	return pulumix.Output[[]*ResourceShare]{
		OutputState: i.ToResourceShareArrayOutputWithContext(ctx).OutputState,
	}
}

// ResourceShareMapInput is an input type that accepts ResourceShareMap and ResourceShareMapOutput values.
// You can construct a concrete instance of `ResourceShareMapInput` via:
//
//	ResourceShareMap{ "key": ResourceShareArgs{...} }
type ResourceShareMapInput interface {
	pulumi.Input

	ToResourceShareMapOutput() ResourceShareMapOutput
	ToResourceShareMapOutputWithContext(context.Context) ResourceShareMapOutput
}

type ResourceShareMap map[string]ResourceShareInput

func (ResourceShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceShare)(nil)).Elem()
}

func (i ResourceShareMap) ToResourceShareMapOutput() ResourceShareMapOutput {
	return i.ToResourceShareMapOutputWithContext(context.Background())
}

func (i ResourceShareMap) ToResourceShareMapOutputWithContext(ctx context.Context) ResourceShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareMapOutput)
}

func (i ResourceShareMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResourceShare] {
	return pulumix.Output[map[string]*ResourceShare]{
		OutputState: i.ToResourceShareMapOutputWithContext(ctx).OutputState,
	}
}

type ResourceShareOutput struct{ *pulumi.OutputState }

func (ResourceShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceShare)(nil)).Elem()
}

func (o ResourceShareOutput) ToResourceShareOutput() ResourceShareOutput {
	return o
}

func (o ResourceShareOutput) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return o
}

func (o ResourceShareOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceShare] {
	return pulumix.Output[*ResourceShare]{
		OutputState: o.OutputState,
	}
}

// The name of resource share.
func (o ResourceShareOutput) ResourceShareName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceShare) pulumi.StringOutput { return v.ResourceShareName }).(pulumi.StringOutput)
}

// The owner of resource share.
func (o ResourceShareOutput) ResourceShareOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceShare) pulumi.StringOutput { return v.ResourceShareOwner }).(pulumi.StringOutput)
}

// The status of resource share.
func (o ResourceShareOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceShare) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ResourceShareArrayOutput struct{ *pulumi.OutputState }

func (ResourceShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceShare)(nil)).Elem()
}

func (o ResourceShareArrayOutput) ToResourceShareArrayOutput() ResourceShareArrayOutput {
	return o
}

func (o ResourceShareArrayOutput) ToResourceShareArrayOutputWithContext(ctx context.Context) ResourceShareArrayOutput {
	return o
}

func (o ResourceShareArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ResourceShare] {
	return pulumix.Output[[]*ResourceShare]{
		OutputState: o.OutputState,
	}
}

func (o ResourceShareArrayOutput) Index(i pulumi.IntInput) ResourceShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceShare {
		return vs[0].([]*ResourceShare)[vs[1].(int)]
	}).(ResourceShareOutput)
}

type ResourceShareMapOutput struct{ *pulumi.OutputState }

func (ResourceShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceShare)(nil)).Elem()
}

func (o ResourceShareMapOutput) ToResourceShareMapOutput() ResourceShareMapOutput {
	return o
}

func (o ResourceShareMapOutput) ToResourceShareMapOutputWithContext(ctx context.Context) ResourceShareMapOutput {
	return o
}

func (o ResourceShareMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ResourceShare] {
	return pulumix.Output[map[string]*ResourceShare]{
		OutputState: o.OutputState,
	}
}

func (o ResourceShareMapOutput) MapIndex(k pulumi.StringInput) ResourceShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceShare {
		return vs[0].(map[string]*ResourceShare)[vs[1].(string)]
	}).(ResourceShareOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareInput)(nil)).Elem(), &ResourceShare{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareArrayInput)(nil)).Elem(), ResourceShareArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareMapInput)(nil)).Elem(), ResourceShareMap{})
	pulumi.RegisterOutputType(ResourceShareOutput{})
	pulumi.RegisterOutputType(ResourceShareArrayOutput{})
	pulumi.RegisterOutputType(ResourceShareMapOutput{})
}
