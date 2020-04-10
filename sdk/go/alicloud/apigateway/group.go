// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	// The description of the api gateway group. Defaults to null.
	Description pulumi.StringOutput `pulumi:"description"`
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
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
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
	// The name of the api gateway group. Defaults to null.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description of the api gateway group. Defaults to null.
	Description pulumi.StringInput
	// The name of the api gateway group. Defaults to null.
	Name pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
