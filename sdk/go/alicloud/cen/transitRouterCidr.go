// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Cloud Enterprise Network (CEN) Transit Router Cidr resource.
//
// For information about Cloud Enterprise Network (CEN) Transit Router Cidr and how to use it, see [What is Transit Router Cidr](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/createtransitroutercidr).
//
// > **NOTE:** Available since v1.193.0.
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
//			exampleInstance, err := cen.NewInstance(ctx, "exampleInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String("tf_example"),
//				Description:     pulumi.String("an example for cen"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouter, err := cen.NewTransitRouter(ctx, "exampleTransitRouter", &cen.TransitRouterArgs{
//				TransitRouterName: pulumi.String("tf_example"),
//				CenId:             exampleInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterCidr(ctx, "exampleTransitRouterCidr", &cen.TransitRouterCidrArgs{
//				TransitRouterId:       exampleTransitRouter.TransitRouterId,
//				Cidr:                  pulumi.String("192.168.0.0/16"),
//				TransitRouterCidrName: pulumi.String("tf_example"),
//				Description:           pulumi.String("tf_example"),
//				PublishCidrRoute:      pulumi.Bool(true),
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
// Cloud Enterprise Network (CEN) Transit Router Cidr can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/transitRouterCidr:TransitRouterCidr default <transit_router_id>:<transit_router_cidr_id>.
//
// ```
type TransitRouterCidr struct {
	pulumi.CustomResourceState

	// The cidr of the transit router.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
	PublishCidrRoute pulumi.BoolOutput `pulumi:"publishCidrRoute"`
	// The ID of the transit router cidr.
	TransitRouterCidrId pulumi.StringOutput `pulumi:"transitRouterCidrId"`
	// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	TransitRouterCidrName pulumi.StringPtrOutput `pulumi:"transitRouterCidrName"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
}

// NewTransitRouterCidr registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterCidr(ctx *pulumi.Context,
	name string, args *TransitRouterCidrArgs, opts ...pulumi.ResourceOption) (*TransitRouterCidr, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cidr == nil {
		return nil, errors.New("invalid value for required argument 'Cidr'")
	}
	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterCidr
	err := ctx.RegisterResource("alicloud:cen/transitRouterCidr:TransitRouterCidr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterCidr gets an existing TransitRouterCidr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterCidr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterCidrState, opts ...pulumi.ResourceOption) (*TransitRouterCidr, error) {
	var resource TransitRouterCidr
	err := ctx.ReadResource("alicloud:cen/transitRouterCidr:TransitRouterCidr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterCidr resources.
type transitRouterCidrState struct {
	// The cidr of the transit router.
	Cidr *string `pulumi:"cidr"`
	// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
	PublishCidrRoute *bool `pulumi:"publishCidrRoute"`
	// The ID of the transit router cidr.
	TransitRouterCidrId *string `pulumi:"transitRouterCidrId"`
	// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	TransitRouterCidrName *string `pulumi:"transitRouterCidrName"`
	// The ID of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

type TransitRouterCidrState struct {
	// The cidr of the transit router.
	Cidr pulumi.StringPtrInput
	// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
	PublishCidrRoute pulumi.BoolPtrInput
	// The ID of the transit router cidr.
	TransitRouterCidrId pulumi.StringPtrInput
	// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	TransitRouterCidrName pulumi.StringPtrInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringPtrInput
}

func (TransitRouterCidrState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterCidrState)(nil)).Elem()
}

type transitRouterCidrArgs struct {
	// The cidr of the transit router.
	Cidr string `pulumi:"cidr"`
	// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
	PublishCidrRoute *bool `pulumi:"publishCidrRoute"`
	// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	TransitRouterCidrName *string `pulumi:"transitRouterCidrName"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// The set of arguments for constructing a TransitRouterCidr resource.
type TransitRouterCidrArgs struct {
	// The cidr of the transit router.
	Cidr pulumi.StringInput
	// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
	PublishCidrRoute pulumi.BoolPtrInput
	// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	TransitRouterCidrName pulumi.StringPtrInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput
}

func (TransitRouterCidrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterCidrArgs)(nil)).Elem()
}

type TransitRouterCidrInput interface {
	pulumi.Input

	ToTransitRouterCidrOutput() TransitRouterCidrOutput
	ToTransitRouterCidrOutputWithContext(ctx context.Context) TransitRouterCidrOutput
}

func (*TransitRouterCidr) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterCidr)(nil)).Elem()
}

func (i *TransitRouterCidr) ToTransitRouterCidrOutput() TransitRouterCidrOutput {
	return i.ToTransitRouterCidrOutputWithContext(context.Background())
}

func (i *TransitRouterCidr) ToTransitRouterCidrOutputWithContext(ctx context.Context) TransitRouterCidrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterCidrOutput)
}

func (i *TransitRouterCidr) ToOutput(ctx context.Context) pulumix.Output[*TransitRouterCidr] {
	return pulumix.Output[*TransitRouterCidr]{
		OutputState: i.ToTransitRouterCidrOutputWithContext(ctx).OutputState,
	}
}

// TransitRouterCidrArrayInput is an input type that accepts TransitRouterCidrArray and TransitRouterCidrArrayOutput values.
// You can construct a concrete instance of `TransitRouterCidrArrayInput` via:
//
//	TransitRouterCidrArray{ TransitRouterCidrArgs{...} }
type TransitRouterCidrArrayInput interface {
	pulumi.Input

	ToTransitRouterCidrArrayOutput() TransitRouterCidrArrayOutput
	ToTransitRouterCidrArrayOutputWithContext(context.Context) TransitRouterCidrArrayOutput
}

type TransitRouterCidrArray []TransitRouterCidrInput

func (TransitRouterCidrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterCidr)(nil)).Elem()
}

func (i TransitRouterCidrArray) ToTransitRouterCidrArrayOutput() TransitRouterCidrArrayOutput {
	return i.ToTransitRouterCidrArrayOutputWithContext(context.Background())
}

func (i TransitRouterCidrArray) ToTransitRouterCidrArrayOutputWithContext(ctx context.Context) TransitRouterCidrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterCidrArrayOutput)
}

func (i TransitRouterCidrArray) ToOutput(ctx context.Context) pulumix.Output[[]*TransitRouterCidr] {
	return pulumix.Output[[]*TransitRouterCidr]{
		OutputState: i.ToTransitRouterCidrArrayOutputWithContext(ctx).OutputState,
	}
}

// TransitRouterCidrMapInput is an input type that accepts TransitRouterCidrMap and TransitRouterCidrMapOutput values.
// You can construct a concrete instance of `TransitRouterCidrMapInput` via:
//
//	TransitRouterCidrMap{ "key": TransitRouterCidrArgs{...} }
type TransitRouterCidrMapInput interface {
	pulumi.Input

	ToTransitRouterCidrMapOutput() TransitRouterCidrMapOutput
	ToTransitRouterCidrMapOutputWithContext(context.Context) TransitRouterCidrMapOutput
}

type TransitRouterCidrMap map[string]TransitRouterCidrInput

func (TransitRouterCidrMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterCidr)(nil)).Elem()
}

func (i TransitRouterCidrMap) ToTransitRouterCidrMapOutput() TransitRouterCidrMapOutput {
	return i.ToTransitRouterCidrMapOutputWithContext(context.Background())
}

func (i TransitRouterCidrMap) ToTransitRouterCidrMapOutputWithContext(ctx context.Context) TransitRouterCidrMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterCidrMapOutput)
}

func (i TransitRouterCidrMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitRouterCidr] {
	return pulumix.Output[map[string]*TransitRouterCidr]{
		OutputState: i.ToTransitRouterCidrMapOutputWithContext(ctx).OutputState,
	}
}

type TransitRouterCidrOutput struct{ *pulumi.OutputState }

func (TransitRouterCidrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterCidr)(nil)).Elem()
}

func (o TransitRouterCidrOutput) ToTransitRouterCidrOutput() TransitRouterCidrOutput {
	return o
}

func (o TransitRouterCidrOutput) ToTransitRouterCidrOutputWithContext(ctx context.Context) TransitRouterCidrOutput {
	return o
}

func (o TransitRouterCidrOutput) ToOutput(ctx context.Context) pulumix.Output[*TransitRouterCidr] {
	return pulumix.Output[*TransitRouterCidr]{
		OutputState: o.OutputState,
	}
}

// The cidr of the transit router.
func (o TransitRouterCidrOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
func (o TransitRouterCidrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
func (o TransitRouterCidrOutput) PublishCidrRoute() pulumi.BoolOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.BoolOutput { return v.PublishCidrRoute }).(pulumi.BoolOutput)
}

// The ID of the transit router cidr.
func (o TransitRouterCidrOutput) TransitRouterCidrId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.StringOutput { return v.TransitRouterCidrId }).(pulumi.StringOutput)
}

// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
func (o TransitRouterCidrOutput) TransitRouterCidrName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.StringPtrOutput { return v.TransitRouterCidrName }).(pulumi.StringPtrOutput)
}

// The ID of the transit router.
func (o TransitRouterCidrOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterCidr) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

type TransitRouterCidrArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterCidrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterCidr)(nil)).Elem()
}

func (o TransitRouterCidrArrayOutput) ToTransitRouterCidrArrayOutput() TransitRouterCidrArrayOutput {
	return o
}

func (o TransitRouterCidrArrayOutput) ToTransitRouterCidrArrayOutputWithContext(ctx context.Context) TransitRouterCidrArrayOutput {
	return o
}

func (o TransitRouterCidrArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TransitRouterCidr] {
	return pulumix.Output[[]*TransitRouterCidr]{
		OutputState: o.OutputState,
	}
}

func (o TransitRouterCidrArrayOutput) Index(i pulumi.IntInput) TransitRouterCidrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterCidr {
		return vs[0].([]*TransitRouterCidr)[vs[1].(int)]
	}).(TransitRouterCidrOutput)
}

type TransitRouterCidrMapOutput struct{ *pulumi.OutputState }

func (TransitRouterCidrMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterCidr)(nil)).Elem()
}

func (o TransitRouterCidrMapOutput) ToTransitRouterCidrMapOutput() TransitRouterCidrMapOutput {
	return o
}

func (o TransitRouterCidrMapOutput) ToTransitRouterCidrMapOutputWithContext(ctx context.Context) TransitRouterCidrMapOutput {
	return o
}

func (o TransitRouterCidrMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TransitRouterCidr] {
	return pulumix.Output[map[string]*TransitRouterCidr]{
		OutputState: o.OutputState,
	}
}

func (o TransitRouterCidrMapOutput) MapIndex(k pulumi.StringInput) TransitRouterCidrOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterCidr {
		return vs[0].(map[string]*TransitRouterCidr)[vs[1].(string)]
	}).(TransitRouterCidrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterCidrInput)(nil)).Elem(), &TransitRouterCidr{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterCidrArrayInput)(nil)).Elem(), TransitRouterCidrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterCidrMapInput)(nil)).Elem(), TransitRouterCidrMap{})
	pulumi.RegisterOutputType(TransitRouterCidrOutput{})
	pulumi.RegisterOutputType(TransitRouterCidrArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterCidrMapOutput{})
}
