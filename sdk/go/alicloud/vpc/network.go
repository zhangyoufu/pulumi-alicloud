// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// VPC can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/network:Network example vpc-abc123456
// ```
type Network struct {
	pulumi.CustomResourceState

	// The CIDR block for the VPC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the VPC. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The route table ID of the router created by default on VPC creation.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The ID of the router created by default on VPC creation.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId pulumi.StringOutput `pulumi:"routerTableId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayOutput `pulumi:"secondaryCidrBlocks"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	var resource Network
	err := ctx.RegisterResource("alicloud:vpc/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("alicloud:vpc/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// The CIDR block for the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description *string `pulumi:"description"`
	// The name of the VPC. Defaults to null.
	Name *string `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The route table ID of the router created by default on VPC creation.
	RouteTableId *string `pulumi:"routeTableId"`
	// The ID of the router created by default on VPC creation.
	RouterId *string `pulumi:"routerId"`
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId *string `pulumi:"routerTableId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks []string `pulumi:"secondaryCidrBlocks"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type NetworkState struct {
	// The CIDR block for the VPC.
	CidrBlock pulumi.StringPtrInput
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrInput
	// The name of the VPC. Defaults to null.
	Name pulumi.StringPtrInput
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The route table ID of the router created by default on VPC creation.
	RouteTableId pulumi.StringPtrInput
	// The ID of the router created by default on VPC creation.
	RouterId pulumi.StringPtrInput
	// Deprecated: Attribute router_table_id has been deprecated and replaced with route_table_id.
	RouterTableId pulumi.StringPtrInput
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// The CIDR block for the VPC.
	CidrBlock string `pulumi:"cidrBlock"`
	// The VPC description. Defaults to null.
	Description *string `pulumi:"description"`
	// The name of the VPC. Defaults to null.
	Name *string `pulumi:"name"`
	// The Id of resource group which the VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks []string `pulumi:"secondaryCidrBlocks"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The CIDR block for the VPC.
	CidrBlock pulumi.StringInput
	// The VPC description. Defaults to null.
	Description pulumi.StringPtrInput
	// The name of the VPC. Defaults to null.
	Name pulumi.StringPtrInput
	// The Id of resource group which the VPC belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The secondary CIDR blocks for the VPC.
	SecondaryCidrBlocks pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

type NetworkOutput struct {
	*pulumi.OutputState
}

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
}
