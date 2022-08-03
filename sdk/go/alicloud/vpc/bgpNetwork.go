// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Bgp Network resource.
//
// For information about VPC Bgp Network and how to use it, see [What is Bgp Network](https://www.alibabacloud.com/help/en/doc-detail/91267.html).
//
// > **NOTE:** Available in v1.153.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultPhysicalConnections, err := expressconnect.GetPhysicalConnections(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultVirtualBorderRouter, err := expressconnect.NewVirtualBorderRouter(ctx, "defaultVirtualBorderRouter", &expressconnect.VirtualBorderRouterArgs{
// 			LocalGatewayIp:          pulumi.String("10.0.0.1"),
// 			PeerGatewayIp:           pulumi.String("10.0.0.2"),
// 			PeeringSubnetMask:       pulumi.String("255.255.255.252"),
// 			PhysicalConnectionId:    pulumi.String(defaultPhysicalConnections.Connections[0].Id),
// 			VirtualBorderRouterName: pulumi.Any(_var.Name),
// 			VlanId:                  pulumi.Int(120),
// 			MinRxInterval:           pulumi.Int(1000),
// 			MinTxInterval:           pulumi.Int(1000),
// 			DetectMultiplier:        pulumi.Int(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vpc.NewBgpNetwork(ctx, "example", &vpc.BgpNetworkArgs{
// 			DstCidrBlock: pulumi.String("example_value"),
// 			RouterId:     defaultVirtualBorderRouter.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// VPC Bgp Network can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpc/bgpNetwork:BgpNetwork example <router_id>:<dst_cidr_block>
// ```
type BgpNetwork struct {
	pulumi.CustomResourceState

	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	DstCidrBlock pulumi.StringOutput `pulumi:"dstCidrBlock"`
	// The ID of the vRouter associated with the router interface.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// The state of the advertised BGP network.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBgpNetwork registers a new resource with the given unique name, arguments, and options.
func NewBgpNetwork(ctx *pulumi.Context,
	name string, args *BgpNetworkArgs, opts ...pulumi.ResourceOption) (*BgpNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DstCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DstCidrBlock'")
	}
	if args.RouterId == nil {
		return nil, errors.New("invalid value for required argument 'RouterId'")
	}
	var resource BgpNetwork
	err := ctx.RegisterResource("alicloud:vpc/bgpNetwork:BgpNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpNetwork gets an existing BgpNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpNetworkState, opts ...pulumi.ResourceOption) (*BgpNetwork, error) {
	var resource BgpNetwork
	err := ctx.ReadResource("alicloud:vpc/bgpNetwork:BgpNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpNetwork resources.
type bgpNetworkState struct {
	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	DstCidrBlock *string `pulumi:"dstCidrBlock"`
	// The ID of the vRouter associated with the router interface.
	RouterId *string `pulumi:"routerId"`
	// The state of the advertised BGP network.
	Status *string `pulumi:"status"`
}

type BgpNetworkState struct {
	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	DstCidrBlock pulumi.StringPtrInput
	// The ID of the vRouter associated with the router interface.
	RouterId pulumi.StringPtrInput
	// The state of the advertised BGP network.
	Status pulumi.StringPtrInput
}

func (BgpNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpNetworkState)(nil)).Elem()
}

type bgpNetworkArgs struct {
	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	DstCidrBlock string `pulumi:"dstCidrBlock"`
	// The ID of the vRouter associated with the router interface.
	RouterId string `pulumi:"routerId"`
}

// The set of arguments for constructing a BgpNetwork resource.
type BgpNetworkArgs struct {
	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	DstCidrBlock pulumi.StringInput
	// The ID of the vRouter associated with the router interface.
	RouterId pulumi.StringInput
}

func (BgpNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpNetworkArgs)(nil)).Elem()
}

type BgpNetworkInput interface {
	pulumi.Input

	ToBgpNetworkOutput() BgpNetworkOutput
	ToBgpNetworkOutputWithContext(ctx context.Context) BgpNetworkOutput
}

func (*BgpNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpNetwork)(nil)).Elem()
}

func (i *BgpNetwork) ToBgpNetworkOutput() BgpNetworkOutput {
	return i.ToBgpNetworkOutputWithContext(context.Background())
}

func (i *BgpNetwork) ToBgpNetworkOutputWithContext(ctx context.Context) BgpNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpNetworkOutput)
}

// BgpNetworkArrayInput is an input type that accepts BgpNetworkArray and BgpNetworkArrayOutput values.
// You can construct a concrete instance of `BgpNetworkArrayInput` via:
//
//          BgpNetworkArray{ BgpNetworkArgs{...} }
type BgpNetworkArrayInput interface {
	pulumi.Input

	ToBgpNetworkArrayOutput() BgpNetworkArrayOutput
	ToBgpNetworkArrayOutputWithContext(context.Context) BgpNetworkArrayOutput
}

type BgpNetworkArray []BgpNetworkInput

func (BgpNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpNetwork)(nil)).Elem()
}

func (i BgpNetworkArray) ToBgpNetworkArrayOutput() BgpNetworkArrayOutput {
	return i.ToBgpNetworkArrayOutputWithContext(context.Background())
}

func (i BgpNetworkArray) ToBgpNetworkArrayOutputWithContext(ctx context.Context) BgpNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpNetworkArrayOutput)
}

// BgpNetworkMapInput is an input type that accepts BgpNetworkMap and BgpNetworkMapOutput values.
// You can construct a concrete instance of `BgpNetworkMapInput` via:
//
//          BgpNetworkMap{ "key": BgpNetworkArgs{...} }
type BgpNetworkMapInput interface {
	pulumi.Input

	ToBgpNetworkMapOutput() BgpNetworkMapOutput
	ToBgpNetworkMapOutputWithContext(context.Context) BgpNetworkMapOutput
}

type BgpNetworkMap map[string]BgpNetworkInput

func (BgpNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpNetwork)(nil)).Elem()
}

func (i BgpNetworkMap) ToBgpNetworkMapOutput() BgpNetworkMapOutput {
	return i.ToBgpNetworkMapOutputWithContext(context.Background())
}

func (i BgpNetworkMap) ToBgpNetworkMapOutputWithContext(ctx context.Context) BgpNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpNetworkMapOutput)
}

type BgpNetworkOutput struct{ *pulumi.OutputState }

func (BgpNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpNetwork)(nil)).Elem()
}

func (o BgpNetworkOutput) ToBgpNetworkOutput() BgpNetworkOutput {
	return o
}

func (o BgpNetworkOutput) ToBgpNetworkOutputWithContext(ctx context.Context) BgpNetworkOutput {
	return o
}

// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
func (o BgpNetworkOutput) DstCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpNetwork) pulumi.StringOutput { return v.DstCidrBlock }).(pulumi.StringOutput)
}

// The ID of the vRouter associated with the router interface.
func (o BgpNetworkOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpNetwork) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

// The state of the advertised BGP network.
func (o BgpNetworkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpNetwork) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BgpNetworkArrayOutput struct{ *pulumi.OutputState }

func (BgpNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpNetwork)(nil)).Elem()
}

func (o BgpNetworkArrayOutput) ToBgpNetworkArrayOutput() BgpNetworkArrayOutput {
	return o
}

func (o BgpNetworkArrayOutput) ToBgpNetworkArrayOutputWithContext(ctx context.Context) BgpNetworkArrayOutput {
	return o
}

func (o BgpNetworkArrayOutput) Index(i pulumi.IntInput) BgpNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpNetwork {
		return vs[0].([]*BgpNetwork)[vs[1].(int)]
	}).(BgpNetworkOutput)
}

type BgpNetworkMapOutput struct{ *pulumi.OutputState }

func (BgpNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpNetwork)(nil)).Elem()
}

func (o BgpNetworkMapOutput) ToBgpNetworkMapOutput() BgpNetworkMapOutput {
	return o
}

func (o BgpNetworkMapOutput) ToBgpNetworkMapOutputWithContext(ctx context.Context) BgpNetworkMapOutput {
	return o
}

func (o BgpNetworkMapOutput) MapIndex(k pulumi.StringInput) BgpNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpNetwork {
		return vs[0].(map[string]*BgpNetwork)[vs[1].(string)]
	}).(BgpNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpNetworkInput)(nil)).Elem(), &BgpNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpNetworkArrayInput)(nil)).Elem(), BgpNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpNetworkMapInput)(nil)).Elem(), BgpNetworkMap{})
	pulumi.RegisterOutputType(BgpNetworkOutput{})
	pulumi.RegisterOutputType(BgpNetworkArrayOutput{})
	pulumi.RegisterOutputType(BgpNetworkMapOutput{})
}
