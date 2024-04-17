// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPN Pbr Route Entry resource.
//
// > **NOTE:** Available since v1.162.0+.
//
// For information about VPN Pbr Route Entry and how to use it, see [What is VPN Pbr Route Entry](https://www.alibabacloud.com/help/en/doc-detail/127248.html).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tfacc"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := vpn.GetGateways(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultCustomerGateway, err := vpn.NewCustomerGateway(ctx, "default", &vpn.CustomerGatewayArgs{
//				Name:      pulumi.String(name),
//				IpAddress: pulumi.String("192.168.1.1"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultConnection, err := vpn.NewConnection(ctx, "default", &vpn.ConnectionArgs{
//				Name:              pulumi.String(name),
//				CustomerGatewayId: defaultCustomerGateway.ID(),
//				VpnGatewayId:      pulumi.String(_default.Ids[0]),
//				LocalSubnets: pulumi.StringArray{
//					pulumi.String("192.168.2.0/24"),
//				},
//				RemoteSubnets: pulumi.StringArray{
//					pulumi.String("192.168.3.0/24"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpn.NewPbrRouteEntry(ctx, "default", &vpn.PbrRouteEntryArgs{
//				VpnGatewayId: pulumi.String(_default.Ids[0]),
//				RouteSource:  pulumi.String("192.168.1.0/24"),
//				RouteDest:    pulumi.String("10.0.0.0/24"),
//				NextHop:      defaultConnection.ID(),
//				Weight:       pulumi.Int(0),
//				PublishVpc:   pulumi.Bool(false),
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
// VPN Pbr route entry can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpn/pbrRouteEntry:PbrRouteEntry example <vpn_gateway_id>:<next_hop>:<route_source>:<route_dest>
// ```
type PbrRouteEntry struct {
	pulumi.CustomResourceState

	// The next hop of the policy-based route.
	NextHop pulumi.StringOutput `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolOutput `pulumi:"publishVpc"`
	// The destination CIDR block of the policy-based route.
	RouteDest pulumi.StringOutput `pulumi:"routeDest"`
	// The source CIDR block of the policy-based route.
	RouteSource pulumi.StringOutput `pulumi:"routeSource"`
	// The status of the vpn pbr route entry.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the vpn gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
	// The weight of the policy-based route. Valid values: 0 and 100.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewPbrRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewPbrRouteEntry(ctx *pulumi.Context,
	name string, args *PbrRouteEntryArgs, opts ...pulumi.ResourceOption) (*PbrRouteEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NextHop == nil {
		return nil, errors.New("invalid value for required argument 'NextHop'")
	}
	if args.PublishVpc == nil {
		return nil, errors.New("invalid value for required argument 'PublishVpc'")
	}
	if args.RouteDest == nil {
		return nil, errors.New("invalid value for required argument 'RouteDest'")
	}
	if args.RouteSource == nil {
		return nil, errors.New("invalid value for required argument 'RouteSource'")
	}
	if args.VpnGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VpnGatewayId'")
	}
	if args.Weight == nil {
		return nil, errors.New("invalid value for required argument 'Weight'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PbrRouteEntry
	err := ctx.RegisterResource("alicloud:vpn/pbrRouteEntry:PbrRouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPbrRouteEntry gets an existing PbrRouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPbrRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PbrRouteEntryState, opts ...pulumi.ResourceOption) (*PbrRouteEntry, error) {
	var resource PbrRouteEntry
	err := ctx.ReadResource("alicloud:vpn/pbrRouteEntry:PbrRouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PbrRouteEntry resources.
type pbrRouteEntryState struct {
	// The next hop of the policy-based route.
	NextHop *string `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc *bool `pulumi:"publishVpc"`
	// The destination CIDR block of the policy-based route.
	RouteDest *string `pulumi:"routeDest"`
	// The source CIDR block of the policy-based route.
	RouteSource *string `pulumi:"routeSource"`
	// The status of the vpn pbr route entry.
	Status *string `pulumi:"status"`
	// The ID of the vpn gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
	// The weight of the policy-based route. Valid values: 0 and 100.
	Weight *int `pulumi:"weight"`
}

type PbrRouteEntryState struct {
	// The next hop of the policy-based route.
	NextHop pulumi.StringPtrInput
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolPtrInput
	// The destination CIDR block of the policy-based route.
	RouteDest pulumi.StringPtrInput
	// The source CIDR block of the policy-based route.
	RouteSource pulumi.StringPtrInput
	// The status of the vpn pbr route entry.
	Status pulumi.StringPtrInput
	// The ID of the vpn gateway.
	VpnGatewayId pulumi.StringPtrInput
	// The weight of the policy-based route. Valid values: 0 and 100.
	Weight pulumi.IntPtrInput
}

func (PbrRouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*pbrRouteEntryState)(nil)).Elem()
}

type pbrRouteEntryArgs struct {
	// The next hop of the policy-based route.
	NextHop string `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc bool `pulumi:"publishVpc"`
	// The destination CIDR block of the policy-based route.
	RouteDest string `pulumi:"routeDest"`
	// The source CIDR block of the policy-based route.
	RouteSource string `pulumi:"routeSource"`
	// The ID of the vpn gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
	// The weight of the policy-based route. Valid values: 0 and 100.
	Weight int `pulumi:"weight"`
}

// The set of arguments for constructing a PbrRouteEntry resource.
type PbrRouteEntryArgs struct {
	// The next hop of the policy-based route.
	NextHop pulumi.StringInput
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolInput
	// The destination CIDR block of the policy-based route.
	RouteDest pulumi.StringInput
	// The source CIDR block of the policy-based route.
	RouteSource pulumi.StringInput
	// The ID of the vpn gateway.
	VpnGatewayId pulumi.StringInput
	// The weight of the policy-based route. Valid values: 0 and 100.
	Weight pulumi.IntInput
}

func (PbrRouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pbrRouteEntryArgs)(nil)).Elem()
}

type PbrRouteEntryInput interface {
	pulumi.Input

	ToPbrRouteEntryOutput() PbrRouteEntryOutput
	ToPbrRouteEntryOutputWithContext(ctx context.Context) PbrRouteEntryOutput
}

func (*PbrRouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**PbrRouteEntry)(nil)).Elem()
}

func (i *PbrRouteEntry) ToPbrRouteEntryOutput() PbrRouteEntryOutput {
	return i.ToPbrRouteEntryOutputWithContext(context.Background())
}

func (i *PbrRouteEntry) ToPbrRouteEntryOutputWithContext(ctx context.Context) PbrRouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PbrRouteEntryOutput)
}

// PbrRouteEntryArrayInput is an input type that accepts PbrRouteEntryArray and PbrRouteEntryArrayOutput values.
// You can construct a concrete instance of `PbrRouteEntryArrayInput` via:
//
//	PbrRouteEntryArray{ PbrRouteEntryArgs{...} }
type PbrRouteEntryArrayInput interface {
	pulumi.Input

	ToPbrRouteEntryArrayOutput() PbrRouteEntryArrayOutput
	ToPbrRouteEntryArrayOutputWithContext(context.Context) PbrRouteEntryArrayOutput
}

type PbrRouteEntryArray []PbrRouteEntryInput

func (PbrRouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PbrRouteEntry)(nil)).Elem()
}

func (i PbrRouteEntryArray) ToPbrRouteEntryArrayOutput() PbrRouteEntryArrayOutput {
	return i.ToPbrRouteEntryArrayOutputWithContext(context.Background())
}

func (i PbrRouteEntryArray) ToPbrRouteEntryArrayOutputWithContext(ctx context.Context) PbrRouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PbrRouteEntryArrayOutput)
}

// PbrRouteEntryMapInput is an input type that accepts PbrRouteEntryMap and PbrRouteEntryMapOutput values.
// You can construct a concrete instance of `PbrRouteEntryMapInput` via:
//
//	PbrRouteEntryMap{ "key": PbrRouteEntryArgs{...} }
type PbrRouteEntryMapInput interface {
	pulumi.Input

	ToPbrRouteEntryMapOutput() PbrRouteEntryMapOutput
	ToPbrRouteEntryMapOutputWithContext(context.Context) PbrRouteEntryMapOutput
}

type PbrRouteEntryMap map[string]PbrRouteEntryInput

func (PbrRouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PbrRouteEntry)(nil)).Elem()
}

func (i PbrRouteEntryMap) ToPbrRouteEntryMapOutput() PbrRouteEntryMapOutput {
	return i.ToPbrRouteEntryMapOutputWithContext(context.Background())
}

func (i PbrRouteEntryMap) ToPbrRouteEntryMapOutputWithContext(ctx context.Context) PbrRouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PbrRouteEntryMapOutput)
}

type PbrRouteEntryOutput struct{ *pulumi.OutputState }

func (PbrRouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PbrRouteEntry)(nil)).Elem()
}

func (o PbrRouteEntryOutput) ToPbrRouteEntryOutput() PbrRouteEntryOutput {
	return o
}

func (o PbrRouteEntryOutput) ToPbrRouteEntryOutputWithContext(ctx context.Context) PbrRouteEntryOutput {
	return o
}

// The next hop of the policy-based route.
func (o PbrRouteEntryOutput) NextHop() pulumi.StringOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.StringOutput { return v.NextHop }).(pulumi.StringOutput)
}

// Whether to issue the destination route to the VPC.
func (o PbrRouteEntryOutput) PublishVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.BoolOutput { return v.PublishVpc }).(pulumi.BoolOutput)
}

// The destination CIDR block of the policy-based route.
func (o PbrRouteEntryOutput) RouteDest() pulumi.StringOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.StringOutput { return v.RouteDest }).(pulumi.StringOutput)
}

// The source CIDR block of the policy-based route.
func (o PbrRouteEntryOutput) RouteSource() pulumi.StringOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.StringOutput { return v.RouteSource }).(pulumi.StringOutput)
}

// The status of the vpn pbr route entry.
func (o PbrRouteEntryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the vpn gateway.
func (o PbrRouteEntryOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.StringOutput { return v.VpnGatewayId }).(pulumi.StringOutput)
}

// The weight of the policy-based route. Valid values: 0 and 100.
func (o PbrRouteEntryOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *PbrRouteEntry) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type PbrRouteEntryArrayOutput struct{ *pulumi.OutputState }

func (PbrRouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PbrRouteEntry)(nil)).Elem()
}

func (o PbrRouteEntryArrayOutput) ToPbrRouteEntryArrayOutput() PbrRouteEntryArrayOutput {
	return o
}

func (o PbrRouteEntryArrayOutput) ToPbrRouteEntryArrayOutputWithContext(ctx context.Context) PbrRouteEntryArrayOutput {
	return o
}

func (o PbrRouteEntryArrayOutput) Index(i pulumi.IntInput) PbrRouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PbrRouteEntry {
		return vs[0].([]*PbrRouteEntry)[vs[1].(int)]
	}).(PbrRouteEntryOutput)
}

type PbrRouteEntryMapOutput struct{ *pulumi.OutputState }

func (PbrRouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PbrRouteEntry)(nil)).Elem()
}

func (o PbrRouteEntryMapOutput) ToPbrRouteEntryMapOutput() PbrRouteEntryMapOutput {
	return o
}

func (o PbrRouteEntryMapOutput) ToPbrRouteEntryMapOutputWithContext(ctx context.Context) PbrRouteEntryMapOutput {
	return o
}

func (o PbrRouteEntryMapOutput) MapIndex(k pulumi.StringInput) PbrRouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PbrRouteEntry {
		return vs[0].(map[string]*PbrRouteEntry)[vs[1].(string)]
	}).(PbrRouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PbrRouteEntryInput)(nil)).Elem(), &PbrRouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*PbrRouteEntryArrayInput)(nil)).Elem(), PbrRouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PbrRouteEntryMapInput)(nil)).Elem(), PbrRouteEntryMap{})
	pulumi.RegisterOutputType(PbrRouteEntryOutput{})
	pulumi.RegisterOutputType(PbrRouteEntryArrayOutput{})
	pulumi.RegisterOutputType(PbrRouteEntryMapOutput{})
}
