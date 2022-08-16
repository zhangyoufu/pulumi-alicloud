// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Express Connect Virtual Border Router resource.
//
// For information about Express Connect Virtual Border Router and how to use it, see [What is Virtual Border Router](https://www.alibabacloud.com/help/en/doc-detail/44854.htm).
//
// > **NOTE:** Available in v1.134.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nameRegex, err := expressconnect.GetPhysicalConnections(ctx, &expressconnect.GetPhysicalConnectionsArgs{
//				NameRegex: pulumi.StringRef("^my-PhysicalConnection"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = expressconnect.NewVirtualBorderRouter(ctx, "example", &expressconnect.VirtualBorderRouterArgs{
//				LocalGatewayIp:          pulumi.String("10.0.0.1"),
//				PeerGatewayIp:           pulumi.String("10.0.0.2"),
//				PeeringSubnetMask:       pulumi.String("255.255.255.252"),
//				PhysicalConnectionId:    pulumi.String(nameRegex.Connections[0].Id),
//				VirtualBorderRouterName: pulumi.String("example_value"),
//				VlanId:                  pulumi.Int(1),
//				MinRxInterval:           pulumi.Int(1000),
//				MinTxInterval:           pulumi.Int(1000),
//				DetectMultiplier:        pulumi.Int(10),
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
// Express Connect Virtual Border Router can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:expressconnect/virtualBorderRouter:VirtualBorderRouter example <id>
//
// ```
type VirtualBorderRouter struct {
	pulumi.CustomResourceState

	// The associated physical connections.
	AssociatedPhysicalConnections pulumi.StringPtrOutput `pulumi:"associatedPhysicalConnections"`
	// The bandwidth.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// Operators for physical connection circuit provided coding.
	CircuitCode pulumi.StringPtrOutput `pulumi:"circuitCode"`
	// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
	DetectMultiplier pulumi.IntOutput `pulumi:"detectMultiplier"`
	// Whether to Enable IPv6. Valid values: `false`, `true`.
	EnableIpv6 pulumi.BoolOutput `pulumi:"enableIpv6"`
	// Alibaba Cloud-Connected IPv4 address.
	LocalGatewayIp pulumi.StringOutput `pulumi:"localGatewayIp"`
	// Alibaba Cloud-Connected IPv6 Address.
	LocalIpv6GatewayIp pulumi.StringPtrOutput `pulumi:"localIpv6GatewayIp"`
	// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
	MinRxInterval pulumi.IntOutput `pulumi:"minRxInterval"`
	// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
	MinTxInterval pulumi.IntOutput `pulumi:"minTxInterval"`
	// The Client-Side Interconnection IPv4 Address.
	PeerGatewayIp pulumi.StringOutput `pulumi:"peerGatewayIp"`
	// The Client-Side Interconnection IPv6 Address.
	PeerIpv6GatewayIp pulumi.StringPtrOutput `pulumi:"peerIpv6GatewayIp"`
	// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
	PeeringIpv6SubnetMask pulumi.StringPtrOutput `pulumi:"peeringIpv6SubnetMask"`
	// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
	PeeringSubnetMask pulumi.StringOutput `pulumi:"peeringSubnetMask"`
	// The ID of the Physical Connection to Which the ID.
	PhysicalConnectionId pulumi.StringOutput `pulumi:"physicalConnectionId"`
	// (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The vbr owner id.
	VbrOwnerId pulumi.StringPtrOutput `pulumi:"vbrOwnerId"`
	// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
	VirtualBorderRouterName pulumi.StringPtrOutput `pulumi:"virtualBorderRouterName"`
	// The VLAN ID of the VBR. Value range: 0~2999.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewVirtualBorderRouter registers a new resource with the given unique name, arguments, and options.
func NewVirtualBorderRouter(ctx *pulumi.Context,
	name string, args *VirtualBorderRouterArgs, opts ...pulumi.ResourceOption) (*VirtualBorderRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalGatewayIp == nil {
		return nil, errors.New("invalid value for required argument 'LocalGatewayIp'")
	}
	if args.PeerGatewayIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerGatewayIp'")
	}
	if args.PeeringSubnetMask == nil {
		return nil, errors.New("invalid value for required argument 'PeeringSubnetMask'")
	}
	if args.PhysicalConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'PhysicalConnectionId'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	var resource VirtualBorderRouter
	err := ctx.RegisterResource("alicloud:expressconnect/virtualBorderRouter:VirtualBorderRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualBorderRouter gets an existing VirtualBorderRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualBorderRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualBorderRouterState, opts ...pulumi.ResourceOption) (*VirtualBorderRouter, error) {
	var resource VirtualBorderRouter
	err := ctx.ReadResource("alicloud:expressconnect/virtualBorderRouter:VirtualBorderRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualBorderRouter resources.
type virtualBorderRouterState struct {
	// The associated physical connections.
	AssociatedPhysicalConnections *string `pulumi:"associatedPhysicalConnections"`
	// The bandwidth.
	Bandwidth *int `pulumi:"bandwidth"`
	// Operators for physical connection circuit provided coding.
	CircuitCode *string `pulumi:"circuitCode"`
	// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
	Description *string `pulumi:"description"`
	// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
	DetectMultiplier *int `pulumi:"detectMultiplier"`
	// Whether to Enable IPv6. Valid values: `false`, `true`.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// Alibaba Cloud-Connected IPv4 address.
	LocalGatewayIp *string `pulumi:"localGatewayIp"`
	// Alibaba Cloud-Connected IPv6 Address.
	LocalIpv6GatewayIp *string `pulumi:"localIpv6GatewayIp"`
	// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
	MinRxInterval *int `pulumi:"minRxInterval"`
	// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
	MinTxInterval *int `pulumi:"minTxInterval"`
	// The Client-Side Interconnection IPv4 Address.
	PeerGatewayIp *string `pulumi:"peerGatewayIp"`
	// The Client-Side Interconnection IPv6 Address.
	PeerIpv6GatewayIp *string `pulumi:"peerIpv6GatewayIp"`
	// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
	PeeringIpv6SubnetMask *string `pulumi:"peeringIpv6SubnetMask"`
	// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
	PeeringSubnetMask *string `pulumi:"peeringSubnetMask"`
	// The ID of the Physical Connection to Which the ID.
	PhysicalConnectionId *string `pulumi:"physicalConnectionId"`
	// (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
	RouteTableId *string `pulumi:"routeTableId"`
	// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
	Status *string `pulumi:"status"`
	// The vbr owner id.
	VbrOwnerId *string `pulumi:"vbrOwnerId"`
	// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
	VirtualBorderRouterName *string `pulumi:"virtualBorderRouterName"`
	// The VLAN ID of the VBR. Value range: 0~2999.
	VlanId *int `pulumi:"vlanId"`
}

type VirtualBorderRouterState struct {
	// The associated physical connections.
	AssociatedPhysicalConnections pulumi.StringPtrInput
	// The bandwidth.
	Bandwidth pulumi.IntPtrInput
	// Operators for physical connection circuit provided coding.
	CircuitCode pulumi.StringPtrInput
	// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
	Description pulumi.StringPtrInput
	// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
	DetectMultiplier pulumi.IntPtrInput
	// Whether to Enable IPv6. Valid values: `false`, `true`.
	EnableIpv6 pulumi.BoolPtrInput
	// Alibaba Cloud-Connected IPv4 address.
	LocalGatewayIp pulumi.StringPtrInput
	// Alibaba Cloud-Connected IPv6 Address.
	LocalIpv6GatewayIp pulumi.StringPtrInput
	// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
	MinRxInterval pulumi.IntPtrInput
	// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
	MinTxInterval pulumi.IntPtrInput
	// The Client-Side Interconnection IPv4 Address.
	PeerGatewayIp pulumi.StringPtrInput
	// The Client-Side Interconnection IPv6 Address.
	PeerIpv6GatewayIp pulumi.StringPtrInput
	// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
	PeeringIpv6SubnetMask pulumi.StringPtrInput
	// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
	PeeringSubnetMask pulumi.StringPtrInput
	// The ID of the Physical Connection to Which the ID.
	PhysicalConnectionId pulumi.StringPtrInput
	// (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
	RouteTableId pulumi.StringPtrInput
	// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
	Status pulumi.StringPtrInput
	// The vbr owner id.
	VbrOwnerId pulumi.StringPtrInput
	// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
	VirtualBorderRouterName pulumi.StringPtrInput
	// The VLAN ID of the VBR. Value range: 0~2999.
	VlanId pulumi.IntPtrInput
}

func (VirtualBorderRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualBorderRouterState)(nil)).Elem()
}

type virtualBorderRouterArgs struct {
	// The associated physical connections.
	AssociatedPhysicalConnections *string `pulumi:"associatedPhysicalConnections"`
	// The bandwidth.
	Bandwidth *int `pulumi:"bandwidth"`
	// Operators for physical connection circuit provided coding.
	CircuitCode *string `pulumi:"circuitCode"`
	// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
	Description *string `pulumi:"description"`
	// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
	DetectMultiplier *int `pulumi:"detectMultiplier"`
	// Whether to Enable IPv6. Valid values: `false`, `true`.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// Alibaba Cloud-Connected IPv4 address.
	LocalGatewayIp string `pulumi:"localGatewayIp"`
	// Alibaba Cloud-Connected IPv6 Address.
	LocalIpv6GatewayIp *string `pulumi:"localIpv6GatewayIp"`
	// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
	MinRxInterval *int `pulumi:"minRxInterval"`
	// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
	MinTxInterval *int `pulumi:"minTxInterval"`
	// The Client-Side Interconnection IPv4 Address.
	PeerGatewayIp string `pulumi:"peerGatewayIp"`
	// The Client-Side Interconnection IPv6 Address.
	PeerIpv6GatewayIp *string `pulumi:"peerIpv6GatewayIp"`
	// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
	PeeringIpv6SubnetMask *string `pulumi:"peeringIpv6SubnetMask"`
	// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
	PeeringSubnetMask string `pulumi:"peeringSubnetMask"`
	// The ID of the Physical Connection to Which the ID.
	PhysicalConnectionId string `pulumi:"physicalConnectionId"`
	// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
	Status *string `pulumi:"status"`
	// The vbr owner id.
	VbrOwnerId *string `pulumi:"vbrOwnerId"`
	// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
	VirtualBorderRouterName *string `pulumi:"virtualBorderRouterName"`
	// The VLAN ID of the VBR. Value range: 0~2999.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a VirtualBorderRouter resource.
type VirtualBorderRouterArgs struct {
	// The associated physical connections.
	AssociatedPhysicalConnections pulumi.StringPtrInput
	// The bandwidth.
	Bandwidth pulumi.IntPtrInput
	// Operators for physical connection circuit provided coding.
	CircuitCode pulumi.StringPtrInput
	// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
	Description pulumi.StringPtrInput
	// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
	DetectMultiplier pulumi.IntPtrInput
	// Whether to Enable IPv6. Valid values: `false`, `true`.
	EnableIpv6 pulumi.BoolPtrInput
	// Alibaba Cloud-Connected IPv4 address.
	LocalGatewayIp pulumi.StringInput
	// Alibaba Cloud-Connected IPv6 Address.
	LocalIpv6GatewayIp pulumi.StringPtrInput
	// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
	MinRxInterval pulumi.IntPtrInput
	// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
	MinTxInterval pulumi.IntPtrInput
	// The Client-Side Interconnection IPv4 Address.
	PeerGatewayIp pulumi.StringInput
	// The Client-Side Interconnection IPv6 Address.
	PeerIpv6GatewayIp pulumi.StringPtrInput
	// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
	PeeringIpv6SubnetMask pulumi.StringPtrInput
	// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
	PeeringSubnetMask pulumi.StringInput
	// The ID of the Physical Connection to Which the ID.
	PhysicalConnectionId pulumi.StringInput
	// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
	Status pulumi.StringPtrInput
	// The vbr owner id.
	VbrOwnerId pulumi.StringPtrInput
	// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
	VirtualBorderRouterName pulumi.StringPtrInput
	// The VLAN ID of the VBR. Value range: 0~2999.
	VlanId pulumi.IntInput
}

func (VirtualBorderRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualBorderRouterArgs)(nil)).Elem()
}

type VirtualBorderRouterInput interface {
	pulumi.Input

	ToVirtualBorderRouterOutput() VirtualBorderRouterOutput
	ToVirtualBorderRouterOutputWithContext(ctx context.Context) VirtualBorderRouterOutput
}

func (*VirtualBorderRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualBorderRouter)(nil)).Elem()
}

func (i *VirtualBorderRouter) ToVirtualBorderRouterOutput() VirtualBorderRouterOutput {
	return i.ToVirtualBorderRouterOutputWithContext(context.Background())
}

func (i *VirtualBorderRouter) ToVirtualBorderRouterOutputWithContext(ctx context.Context) VirtualBorderRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBorderRouterOutput)
}

// VirtualBorderRouterArrayInput is an input type that accepts VirtualBorderRouterArray and VirtualBorderRouterArrayOutput values.
// You can construct a concrete instance of `VirtualBorderRouterArrayInput` via:
//
//	VirtualBorderRouterArray{ VirtualBorderRouterArgs{...} }
type VirtualBorderRouterArrayInput interface {
	pulumi.Input

	ToVirtualBorderRouterArrayOutput() VirtualBorderRouterArrayOutput
	ToVirtualBorderRouterArrayOutputWithContext(context.Context) VirtualBorderRouterArrayOutput
}

type VirtualBorderRouterArray []VirtualBorderRouterInput

func (VirtualBorderRouterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualBorderRouter)(nil)).Elem()
}

func (i VirtualBorderRouterArray) ToVirtualBorderRouterArrayOutput() VirtualBorderRouterArrayOutput {
	return i.ToVirtualBorderRouterArrayOutputWithContext(context.Background())
}

func (i VirtualBorderRouterArray) ToVirtualBorderRouterArrayOutputWithContext(ctx context.Context) VirtualBorderRouterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBorderRouterArrayOutput)
}

// VirtualBorderRouterMapInput is an input type that accepts VirtualBorderRouterMap and VirtualBorderRouterMapOutput values.
// You can construct a concrete instance of `VirtualBorderRouterMapInput` via:
//
//	VirtualBorderRouterMap{ "key": VirtualBorderRouterArgs{...} }
type VirtualBorderRouterMapInput interface {
	pulumi.Input

	ToVirtualBorderRouterMapOutput() VirtualBorderRouterMapOutput
	ToVirtualBorderRouterMapOutputWithContext(context.Context) VirtualBorderRouterMapOutput
}

type VirtualBorderRouterMap map[string]VirtualBorderRouterInput

func (VirtualBorderRouterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualBorderRouter)(nil)).Elem()
}

func (i VirtualBorderRouterMap) ToVirtualBorderRouterMapOutput() VirtualBorderRouterMapOutput {
	return i.ToVirtualBorderRouterMapOutputWithContext(context.Background())
}

func (i VirtualBorderRouterMap) ToVirtualBorderRouterMapOutputWithContext(ctx context.Context) VirtualBorderRouterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualBorderRouterMapOutput)
}

type VirtualBorderRouterOutput struct{ *pulumi.OutputState }

func (VirtualBorderRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualBorderRouter)(nil)).Elem()
}

func (o VirtualBorderRouterOutput) ToVirtualBorderRouterOutput() VirtualBorderRouterOutput {
	return o
}

func (o VirtualBorderRouterOutput) ToVirtualBorderRouterOutputWithContext(ctx context.Context) VirtualBorderRouterOutput {
	return o
}

// The associated physical connections.
func (o VirtualBorderRouterOutput) AssociatedPhysicalConnections() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.AssociatedPhysicalConnections }).(pulumi.StringPtrOutput)
}

// The bandwidth.
func (o VirtualBorderRouterOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// Operators for physical connection circuit provided coding.
func (o VirtualBorderRouterOutput) CircuitCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.CircuitCode }).(pulumi.StringPtrOutput)
}

// The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.
func (o VirtualBorderRouterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.
func (o VirtualBorderRouterOutput) DetectMultiplier() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.IntOutput { return v.DetectMultiplier }).(pulumi.IntOutput)
}

// Whether to Enable IPv6. Valid values: `false`, `true`.
func (o VirtualBorderRouterOutput) EnableIpv6() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.BoolOutput { return v.EnableIpv6 }).(pulumi.BoolOutput)
}

// Alibaba Cloud-Connected IPv4 address.
func (o VirtualBorderRouterOutput) LocalGatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.LocalGatewayIp }).(pulumi.StringOutput)
}

// Alibaba Cloud-Connected IPv6 Address.
func (o VirtualBorderRouterOutput) LocalIpv6GatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.LocalIpv6GatewayIp }).(pulumi.StringPtrOutput)
}

// Configure BFD packet reception interval of values include: 200~1000, unit: ms.
func (o VirtualBorderRouterOutput) MinRxInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.IntOutput { return v.MinRxInterval }).(pulumi.IntOutput)
}

// Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.
func (o VirtualBorderRouterOutput) MinTxInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.IntOutput { return v.MinTxInterval }).(pulumi.IntOutput)
}

// The Client-Side Interconnection IPv4 Address.
func (o VirtualBorderRouterOutput) PeerGatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.PeerGatewayIp }).(pulumi.StringOutput)
}

// The Client-Side Interconnection IPv6 Address.
func (o VirtualBorderRouterOutput) PeerIpv6GatewayIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.PeerIpv6GatewayIp }).(pulumi.StringPtrOutput)
}

// Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.
func (o VirtualBorderRouterOutput) PeeringIpv6SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.PeeringIpv6SubnetMask }).(pulumi.StringPtrOutput)
}

// Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.
func (o VirtualBorderRouterOutput) PeeringSubnetMask() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.PeeringSubnetMask }).(pulumi.StringOutput)
}

// The ID of the Physical Connection to Which the ID.
func (o VirtualBorderRouterOutput) PhysicalConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.PhysicalConnectionId }).(pulumi.StringOutput)
}

// (Available in v1.166.0+) The Route Table ID Of the Virtual Border Router.
func (o VirtualBorderRouterOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// The instance state. Valid values: `active`, `deleting`, `recovering`, `terminated`, `terminating`, `unconfirmed`.
func (o VirtualBorderRouterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The vbr owner id.
func (o VirtualBorderRouterOutput) VbrOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.VbrOwnerId }).(pulumi.StringPtrOutput)
}

// The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.
func (o VirtualBorderRouterOutput) VirtualBorderRouterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.StringPtrOutput { return v.VirtualBorderRouterName }).(pulumi.StringPtrOutput)
}

// The VLAN ID of the VBR. Value range: 0~2999.
func (o VirtualBorderRouterOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualBorderRouter) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

type VirtualBorderRouterArrayOutput struct{ *pulumi.OutputState }

func (VirtualBorderRouterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualBorderRouter)(nil)).Elem()
}

func (o VirtualBorderRouterArrayOutput) ToVirtualBorderRouterArrayOutput() VirtualBorderRouterArrayOutput {
	return o
}

func (o VirtualBorderRouterArrayOutput) ToVirtualBorderRouterArrayOutputWithContext(ctx context.Context) VirtualBorderRouterArrayOutput {
	return o
}

func (o VirtualBorderRouterArrayOutput) Index(i pulumi.IntInput) VirtualBorderRouterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualBorderRouter {
		return vs[0].([]*VirtualBorderRouter)[vs[1].(int)]
	}).(VirtualBorderRouterOutput)
}

type VirtualBorderRouterMapOutput struct{ *pulumi.OutputState }

func (VirtualBorderRouterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualBorderRouter)(nil)).Elem()
}

func (o VirtualBorderRouterMapOutput) ToVirtualBorderRouterMapOutput() VirtualBorderRouterMapOutput {
	return o
}

func (o VirtualBorderRouterMapOutput) ToVirtualBorderRouterMapOutputWithContext(ctx context.Context) VirtualBorderRouterMapOutput {
	return o
}

func (o VirtualBorderRouterMapOutput) MapIndex(k pulumi.StringInput) VirtualBorderRouterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualBorderRouter {
		return vs[0].(map[string]*VirtualBorderRouter)[vs[1].(string)]
	}).(VirtualBorderRouterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBorderRouterInput)(nil)).Elem(), &VirtualBorderRouter{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBorderRouterArrayInput)(nil)).Elem(), VirtualBorderRouterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualBorderRouterMapInput)(nil)).Elem(), VirtualBorderRouterMap{})
	pulumi.RegisterOutputType(VirtualBorderRouterOutput{})
	pulumi.RegisterOutputType(VirtualBorderRouterArrayOutput{})
	pulumi.RegisterOutputType(VirtualBorderRouterMapOutput{})
}
