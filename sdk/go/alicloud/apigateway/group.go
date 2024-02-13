// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewGroup(ctx, "default", &apigateway.GroupArgs{
//				Description: pulumi.String("tf_example"),
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
// Api gateway group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:apigateway/group:Group example "ab2351f2ce904edaa8d92a0510832b91"
// ```
type Group struct {
	pulumi.CustomResourceState

	// The description of the api gateway group. Defaults to null.
	Description pulumi.StringOutput `pulumi:"description"`
	// The id of the api gateway.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the api gateway group. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Available in 1.69.0+)	Second-level domain name automatically assigned to the API group.
	SubDomain pulumi.StringOutput `pulumi:"subDomain"`
	// (Available in 1.69.0+)	Second-level VPC domain name automatically assigned to the API group.
	VpcDomain pulumi.StringOutput `pulumi:"vpcDomain"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("alicloud:apigateway/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("alicloud:apigateway/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The description of the api gateway group. Defaults to null.
	Description *string `pulumi:"description"`
	// The id of the api gateway.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the api gateway group. Defaults to null.
	Name *string `pulumi:"name"`
	// (Available in 1.69.0+)	Second-level domain name automatically assigned to the API group.
	SubDomain *string `pulumi:"subDomain"`
	// (Available in 1.69.0+)	Second-level VPC domain name automatically assigned to the API group.
	VpcDomain *string `pulumi:"vpcDomain"`
}

type GroupState struct {
	// The description of the api gateway group. Defaults to null.
	Description pulumi.StringPtrInput
	// The id of the api gateway.
	InstanceId pulumi.StringPtrInput
	// The name of the api gateway group. Defaults to null.
	Name pulumi.StringPtrInput
	// (Available in 1.69.0+)	Second-level domain name automatically assigned to the API group.
	SubDomain pulumi.StringPtrInput
	// (Available in 1.69.0+)	Second-level VPC domain name automatically assigned to the API group.
	VpcDomain pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The description of the api gateway group. Defaults to null.
	Description string `pulumi:"description"`
	// The id of the api gateway.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the api gateway group. Defaults to null.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description of the api gateway group. Defaults to null.
	Description pulumi.StringInput
	// The id of the api gateway.
	InstanceId pulumi.StringPtrInput
	// The name of the api gateway group. Defaults to null.
	Name pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The description of the api gateway group. Defaults to null.
func (o GroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The id of the api gateway.
func (o GroupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of the api gateway group. Defaults to null.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Available in 1.69.0+)	Second-level domain name automatically assigned to the API group.
func (o GroupOutput) SubDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.SubDomain }).(pulumi.StringOutput)
}

// (Available in 1.69.0+)	Second-level VPC domain name automatically assigned to the API group.
func (o GroupOutput) VpcDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.VpcDomain }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
