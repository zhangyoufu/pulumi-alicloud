// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Express Connect Vbr Pconn Association resource.
//
// For information about Express Connect Vbr Pconn Association and how to use it, see [What is Vbr Pconn Association](https://www.alibabacloud.com/help/en/express-connect/latest/associatephysicalconnectiontovirtualborderrouter#doc-api-Vpc-AssociatePhysicalConnectionToVirtualBorderRouter).
//
// > **NOTE:** Available since v1.196.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			examplePhysicalConnections, err := expressconnect.GetPhysicalConnections(ctx, &expressconnect.GetPhysicalConnectionsArgs{
//				NameRegex: pulumi.StringRef("^preserved-NODELETING"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = expressconnect.NewVirtualBorderRouter(ctx, "default", &expressconnect.VirtualBorderRouterArgs{
//				LocalGatewayIp:          pulumi.String("10.0.0.1"),
//				PeerGatewayIp:           pulumi.String("10.0.0.2"),
//				PeeringSubnetMask:       pulumi.String("255.255.255.252"),
//				PhysicalConnectionId:    *pulumi.String(examplePhysicalConnections.Connections[0].Id),
//				VirtualBorderRouterName: pulumi.String(name),
//				VlanId:                  pulumi.Int(110),
//				MinRxInterval:           pulumi.Int(1000),
//				MinTxInterval:           pulumi.Int(1000),
//				DetectMultiplier:        pulumi.Int(10),
//				EnableIpv6:              pulumi.Bool(true),
//				LocalIpv6GatewayIp:      pulumi.String("2408:4004:cc:400::1"),
//				PeerIpv6GatewayIp:       pulumi.String("2408:4004:cc:400::2"),
//				PeeringIpv6SubnetMask:   pulumi.String("2408:4004:cc:400::/56"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = expressconnect.NewVbrPconnAssociation(ctx, "exampleVbrPconnAssociation", &expressconnect.VbrPconnAssociationArgs{
//				PeerGatewayIp:         pulumi.String("10.0.0.6"),
//				LocalGatewayIp:        pulumi.String("10.0.0.5"),
//				PhysicalConnectionId:  *pulumi.String(examplePhysicalConnections.Connections[2].Id),
//				VbrId:                 _default.ID(),
//				PeeringSubnetMask:     pulumi.String("255.255.255.252"),
//				VlanId:                pulumi.Int(1122),
//				EnableIpv6:            pulumi.Bool(true),
//				LocalIpv6GatewayIp:    pulumi.String("2408:4004:cc::3"),
//				PeerIpv6GatewayIp:     pulumi.String("2408:4004:cc::4"),
//				PeeringIpv6SubnetMask: pulumi.String("2408:4004:cc::/56"),
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
// Express Connect Vbr Pconn Association can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation example <VbrId>:<PhysicalConnectionId>
// ```
type VbrPconnAssociation struct {
	pulumi.CustomResourceState

	// The circuit code provided by the operator for the physical connection.
	CircuitCode pulumi.StringOutput `pulumi:"circuitCode"`
	// Whether IPv6 is enabled. Value:
	// - **true**: on.
	// - **false** (default): Off.
	EnableIpv6 pulumi.BoolOutput `pulumi:"enableIpv6"`
	// The Alibaba cloud IP address of the VBR instance.
	LocalGatewayIp pulumi.StringPtrOutput `pulumi:"localGatewayIp"`
	// The IPv6 address on the Alibaba Cloud side of the VBR instance.
	LocalIpv6GatewayIp pulumi.StringPtrOutput `pulumi:"localIpv6GatewayIp"`
	// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerGatewayIp pulumi.StringPtrOutput `pulumi:"peerGatewayIp"`
	// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerIpv6GatewayIp pulumi.StringPtrOutput `pulumi:"peerIpv6GatewayIp"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
	PeeringIpv6SubnetMask pulumi.StringPtrOutput `pulumi:"peeringIpv6SubnetMask"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
	PeeringSubnetMask pulumi.StringPtrOutput `pulumi:"peeringSubnetMask"`
	// The ID of the leased line instance.
	PhysicalConnectionId pulumi.StringOutput `pulumi:"physicalConnectionId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the VBR instance.
	VbrId pulumi.StringOutput `pulumi:"vbrId"`
	// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewVbrPconnAssociation registers a new resource with the given unique name, arguments, and options.
func NewVbrPconnAssociation(ctx *pulumi.Context,
	name string, args *VbrPconnAssociationArgs, opts ...pulumi.ResourceOption) (*VbrPconnAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PhysicalConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'PhysicalConnectionId'")
	}
	if args.VbrId == nil {
		return nil, errors.New("invalid value for required argument 'VbrId'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VbrPconnAssociation
	err := ctx.RegisterResource("alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVbrPconnAssociation gets an existing VbrPconnAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVbrPconnAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VbrPconnAssociationState, opts ...pulumi.ResourceOption) (*VbrPconnAssociation, error) {
	var resource VbrPconnAssociation
	err := ctx.ReadResource("alicloud:expressconnect/vbrPconnAssociation:VbrPconnAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VbrPconnAssociation resources.
type vbrPconnAssociationState struct {
	// The circuit code provided by the operator for the physical connection.
	CircuitCode *string `pulumi:"circuitCode"`
	// Whether IPv6 is enabled. Value:
	// - **true**: on.
	// - **false** (default): Off.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The Alibaba cloud IP address of the VBR instance.
	LocalGatewayIp *string `pulumi:"localGatewayIp"`
	// The IPv6 address on the Alibaba Cloud side of the VBR instance.
	LocalIpv6GatewayIp *string `pulumi:"localIpv6GatewayIp"`
	// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerGatewayIp *string `pulumi:"peerGatewayIp"`
	// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerIpv6GatewayIp *string `pulumi:"peerIpv6GatewayIp"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
	PeeringIpv6SubnetMask *string `pulumi:"peeringIpv6SubnetMask"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
	PeeringSubnetMask *string `pulumi:"peeringSubnetMask"`
	// The ID of the leased line instance.
	PhysicalConnectionId *string `pulumi:"physicalConnectionId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the VBR instance.
	VbrId *string `pulumi:"vbrId"`
	// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
	VlanId *int `pulumi:"vlanId"`
}

type VbrPconnAssociationState struct {
	// The circuit code provided by the operator for the physical connection.
	CircuitCode pulumi.StringPtrInput
	// Whether IPv6 is enabled. Value:
	// - **true**: on.
	// - **false** (default): Off.
	EnableIpv6 pulumi.BoolPtrInput
	// The Alibaba cloud IP address of the VBR instance.
	LocalGatewayIp pulumi.StringPtrInput
	// The IPv6 address on the Alibaba Cloud side of the VBR instance.
	LocalIpv6GatewayIp pulumi.StringPtrInput
	// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerGatewayIp pulumi.StringPtrInput
	// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerIpv6GatewayIp pulumi.StringPtrInput
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
	PeeringIpv6SubnetMask pulumi.StringPtrInput
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
	PeeringSubnetMask pulumi.StringPtrInput
	// The ID of the leased line instance.
	PhysicalConnectionId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The ID of the VBR instance.
	VbrId pulumi.StringPtrInput
	// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
	VlanId pulumi.IntPtrInput
}

func (VbrPconnAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vbrPconnAssociationState)(nil)).Elem()
}

type vbrPconnAssociationArgs struct {
	// Whether IPv6 is enabled. Value:
	// - **true**: on.
	// - **false** (default): Off.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The Alibaba cloud IP address of the VBR instance.
	LocalGatewayIp *string `pulumi:"localGatewayIp"`
	// The IPv6 address on the Alibaba Cloud side of the VBR instance.
	LocalIpv6GatewayIp *string `pulumi:"localIpv6GatewayIp"`
	// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerGatewayIp *string `pulumi:"peerGatewayIp"`
	// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerIpv6GatewayIp *string `pulumi:"peerIpv6GatewayIp"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
	PeeringIpv6SubnetMask *string `pulumi:"peeringIpv6SubnetMask"`
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
	PeeringSubnetMask *string `pulumi:"peeringSubnetMask"`
	// The ID of the leased line instance.
	PhysicalConnectionId string `pulumi:"physicalConnectionId"`
	// The ID of the VBR instance.
	VbrId string `pulumi:"vbrId"`
	// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a VbrPconnAssociation resource.
type VbrPconnAssociationArgs struct {
	// Whether IPv6 is enabled. Value:
	// - **true**: on.
	// - **false** (default): Off.
	EnableIpv6 pulumi.BoolPtrInput
	// The Alibaba cloud IP address of the VBR instance.
	LocalGatewayIp pulumi.StringPtrInput
	// The IPv6 address on the Alibaba Cloud side of the VBR instance.
	LocalIpv6GatewayIp pulumi.StringPtrInput
	// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerGatewayIp pulumi.StringPtrInput
	// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
	PeerIpv6GatewayIp pulumi.StringPtrInput
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
	PeeringIpv6SubnetMask pulumi.StringPtrInput
	// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
	PeeringSubnetMask pulumi.StringPtrInput
	// The ID of the leased line instance.
	PhysicalConnectionId pulumi.StringInput
	// The ID of the VBR instance.
	VbrId pulumi.StringInput
	// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
	VlanId pulumi.IntInput
}

func (VbrPconnAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vbrPconnAssociationArgs)(nil)).Elem()
}

type VbrPconnAssociationInput interface {
	pulumi.Input

	ToVbrPconnAssociationOutput() VbrPconnAssociationOutput
	ToVbrPconnAssociationOutputWithContext(ctx context.Context) VbrPconnAssociationOutput
}

func (*VbrPconnAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**VbrPconnAssociation)(nil)).Elem()
}

func (i *VbrPconnAssociation) ToVbrPconnAssociationOutput() VbrPconnAssociationOutput {
	return i.ToVbrPconnAssociationOutputWithContext(context.Background())
}

func (i *VbrPconnAssociation) ToVbrPconnAssociationOutputWithContext(ctx context.Context) VbrPconnAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VbrPconnAssociationOutput)
}

// VbrPconnAssociationArrayInput is an input type that accepts VbrPconnAssociationArray and VbrPconnAssociationArrayOutput values.
// You can construct a concrete instance of `VbrPconnAssociationArrayInput` via:
//
//	VbrPconnAssociationArray{ VbrPconnAssociationArgs{...} }
type VbrPconnAssociationArrayInput interface {
	pulumi.Input

	ToVbrPconnAssociationArrayOutput() VbrPconnAssociationArrayOutput
	ToVbrPconnAssociationArrayOutputWithContext(context.Context) VbrPconnAssociationArrayOutput
}

type VbrPconnAssociationArray []VbrPconnAssociationInput

func (VbrPconnAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VbrPconnAssociation)(nil)).Elem()
}

func (i VbrPconnAssociationArray) ToVbrPconnAssociationArrayOutput() VbrPconnAssociationArrayOutput {
	return i.ToVbrPconnAssociationArrayOutputWithContext(context.Background())
}

func (i VbrPconnAssociationArray) ToVbrPconnAssociationArrayOutputWithContext(ctx context.Context) VbrPconnAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VbrPconnAssociationArrayOutput)
}

// VbrPconnAssociationMapInput is an input type that accepts VbrPconnAssociationMap and VbrPconnAssociationMapOutput values.
// You can construct a concrete instance of `VbrPconnAssociationMapInput` via:
//
//	VbrPconnAssociationMap{ "key": VbrPconnAssociationArgs{...} }
type VbrPconnAssociationMapInput interface {
	pulumi.Input

	ToVbrPconnAssociationMapOutput() VbrPconnAssociationMapOutput
	ToVbrPconnAssociationMapOutputWithContext(context.Context) VbrPconnAssociationMapOutput
}

type VbrPconnAssociationMap map[string]VbrPconnAssociationInput

func (VbrPconnAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VbrPconnAssociation)(nil)).Elem()
}

func (i VbrPconnAssociationMap) ToVbrPconnAssociationMapOutput() VbrPconnAssociationMapOutput {
	return i.ToVbrPconnAssociationMapOutputWithContext(context.Background())
}

func (i VbrPconnAssociationMap) ToVbrPconnAssociationMapOutputWithContext(ctx context.Context) VbrPconnAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VbrPconnAssociationMapOutput)
}

type VbrPconnAssociationOutput struct{ *pulumi.OutputState }

func (VbrPconnAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VbrPconnAssociation)(nil)).Elem()
}

func (o VbrPconnAssociationOutput) ToVbrPconnAssociationOutput() VbrPconnAssociationOutput {
	return o
}

func (o VbrPconnAssociationOutput) ToVbrPconnAssociationOutputWithContext(ctx context.Context) VbrPconnAssociationOutput {
	return o
}

// The circuit code provided by the operator for the physical connection.
func (o VbrPconnAssociationOutput) CircuitCode() pulumi.StringOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringOutput { return v.CircuitCode }).(pulumi.StringOutput)
}

// Whether IPv6 is enabled. Value:
// - **true**: on.
// - **false** (default): Off.
func (o VbrPconnAssociationOutput) EnableIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.BoolOutput { return v.EnableIpv6 }).(pulumi.BoolOutput)
}

// The Alibaba cloud IP address of the VBR instance.
func (o VbrPconnAssociationOutput) LocalGatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.LocalGatewayIp }).(pulumi.StringPtrOutput)
}

// The IPv6 address on the Alibaba Cloud side of the VBR instance.
func (o VbrPconnAssociationOutput) LocalIpv6GatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.LocalIpv6GatewayIp }).(pulumi.StringPtrOutput)
}

// The client IP address of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
func (o VbrPconnAssociationOutput) PeerGatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.PeerGatewayIp }).(pulumi.StringPtrOutput)
}

// The IPv6 address of the client side of the VBR instance. This attribute only allows the VBR owner to specify or modify. **NOTE:** Required when creating a VBR instance for the physical connection owner.
func (o VbrPconnAssociationOutput) PeerIpv6GatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.PeerIpv6GatewayIp }).(pulumi.StringPtrOutput)
}

// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.Two IPv6 addresses must be in the same subnet.
func (o VbrPconnAssociationOutput) PeeringIpv6SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.PeeringIpv6SubnetMask }).(pulumi.StringPtrOutput)
}

// The subnet mask of the Alibaba Cloud side and the client side of the VBR instance.The two IP addresses must be in the same subnet.
func (o VbrPconnAssociationOutput) PeeringSubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringPtrOutput { return v.PeeringSubnetMask }).(pulumi.StringPtrOutput)
}

// The ID of the leased line instance.
func (o VbrPconnAssociationOutput) PhysicalConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringOutput { return v.PhysicalConnectionId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o VbrPconnAssociationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the VBR instance.
func (o VbrPconnAssociationOutput) VbrId() pulumi.StringOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.StringOutput { return v.VbrId }).(pulumi.StringOutput)
}

// VLAN ID of the VBR. Valid values: **0 to 2999**. **NOTE:** only the owner of the physical connection can specify this parameter. The VLAN ID of two VBRs under the same physical connection cannot be the same.
func (o VbrPconnAssociationOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *VbrPconnAssociation) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

type VbrPconnAssociationArrayOutput struct{ *pulumi.OutputState }

func (VbrPconnAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VbrPconnAssociation)(nil)).Elem()
}

func (o VbrPconnAssociationArrayOutput) ToVbrPconnAssociationArrayOutput() VbrPconnAssociationArrayOutput {
	return o
}

func (o VbrPconnAssociationArrayOutput) ToVbrPconnAssociationArrayOutputWithContext(ctx context.Context) VbrPconnAssociationArrayOutput {
	return o
}

func (o VbrPconnAssociationArrayOutput) Index(i pulumi.IntInput) VbrPconnAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VbrPconnAssociation {
		return vs[0].([]*VbrPconnAssociation)[vs[1].(int)]
	}).(VbrPconnAssociationOutput)
}

type VbrPconnAssociationMapOutput struct{ *pulumi.OutputState }

func (VbrPconnAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VbrPconnAssociation)(nil)).Elem()
}

func (o VbrPconnAssociationMapOutput) ToVbrPconnAssociationMapOutput() VbrPconnAssociationMapOutput {
	return o
}

func (o VbrPconnAssociationMapOutput) ToVbrPconnAssociationMapOutputWithContext(ctx context.Context) VbrPconnAssociationMapOutput {
	return o
}

func (o VbrPconnAssociationMapOutput) MapIndex(k pulumi.StringInput) VbrPconnAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VbrPconnAssociation {
		return vs[0].(map[string]*VbrPconnAssociation)[vs[1].(string)]
	}).(VbrPconnAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VbrPconnAssociationInput)(nil)).Elem(), &VbrPconnAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VbrPconnAssociationArrayInput)(nil)).Elem(), VbrPconnAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VbrPconnAssociationMapInput)(nil)).Elem(), VbrPconnAssociationMap{})
	pulumi.RegisterOutputType(VbrPconnAssociationOutput{})
	pulumi.RegisterOutputType(VbrPconnAssociationArrayOutput{})
	pulumi.RegisterOutputType(VbrPconnAssociationMapOutput{})
}
