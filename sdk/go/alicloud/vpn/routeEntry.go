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

// ## Example Usage
//
// # Basic Usage
//
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
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := vpn.GetGateways(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultCustomerGateway, err := vpn.NewCustomerGateway(ctx, "defaultCustomerGateway", &vpn.CustomerGatewayArgs{
//				Description:         pulumi.String("defaultCustomerGateway"),
//				IpAddress:           pulumi.String("2.2.2.5"),
//				Asn:                 pulumi.String("2224"),
//				CustomerGatewayName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpn.NewCustomerGateway(ctx, "changeCustomerGateway", &vpn.CustomerGatewayArgs{
//				Description:         pulumi.String("changeCustomerGateway"),
//				IpAddress:           pulumi.String("2.2.2.6"),
//				Asn:                 pulumi.String("2225"),
//				CustomerGatewayName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultConnection, err := vpn.NewConnection(ctx, "default", &vpn.ConnectionArgs{
//				VpnGatewayId:      pulumi.String(_default.Ids[0]),
//				VpnConnectionName: pulumi.String(name),
//				LocalSubnets: pulumi.StringArray{
//					pulumi.String("3.0.0.0/24"),
//				},
//				RemoteSubnets: pulumi.StringArray{
//					pulumi.String("10.0.0.0/24"),
//					pulumi.String("10.0.1.0/24"),
//				},
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//				EnableTunnelsBgp: pulumi.Bool(true),
//				TunnelOptionsSpecifications: vpn.ConnectionTunnelOptionsSpecificationArray{
//					&vpn.ConnectionTunnelOptionsSpecificationArgs{
//						TunnelIpsecConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelIpsecConfigArgs{
//							IpsecAuthAlg:  pulumi.String("md5"),
//							IpsecEncAlg:   pulumi.String("aes256"),
//							IpsecLifetime: pulumi.Int(16400),
//							IpsecPfs:      pulumi.String("group5"),
//						},
//						CustomerGatewayId: defaultCustomerGateway.ID(),
//						Role:              pulumi.String("master"),
//						TunnelBgpConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelBgpConfigArgs{
//							LocalAsn:   pulumi.String("1219002"),
//							TunnelCidr: pulumi.String("169.254.30.0/30"),
//							LocalBgpIp: pulumi.String("169.254.30.1"),
//						},
//						TunnelIkeConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelIkeConfigArgs{
//							IkeMode:     pulumi.String("aggressive"),
//							IkeVersion:  pulumi.String("ikev2"),
//							LocalId:     pulumi.String("localid_tunnel2"),
//							Psk:         pulumi.String("12345678"),
//							RemoteId:    pulumi.String("remote2"),
//							IkeAuthAlg:  pulumi.String("md5"),
//							IkeEncAlg:   pulumi.String("aes256"),
//							IkeLifetime: pulumi.Int(3600),
//							IkePfs:      pulumi.String("group14"),
//						},
//					},
//					&vpn.ConnectionTunnelOptionsSpecificationArgs{
//						TunnelIkeConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelIkeConfigArgs{
//							RemoteId:    pulumi.String("remote24"),
//							IkeEncAlg:   pulumi.String("aes256"),
//							IkeLifetime: pulumi.Int(27000),
//							IkeMode:     pulumi.String("aggressive"),
//							IkePfs:      pulumi.String("group5"),
//							IkeAuthAlg:  pulumi.String("md5"),
//							IkeVersion:  pulumi.String("ikev2"),
//							LocalId:     pulumi.String("localid_tunnel2"),
//							Psk:         pulumi.String("12345678"),
//						},
//						TunnelIpsecConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelIpsecConfigArgs{
//							IpsecLifetime: pulumi.Int(2700),
//							IpsecPfs:      pulumi.String("group14"),
//							IpsecAuthAlg:  pulumi.String("md5"),
//							IpsecEncAlg:   pulumi.String("aes256"),
//						},
//						CustomerGatewayId: defaultCustomerGateway.ID(),
//						Role:              pulumi.String("slave"),
//						TunnelBgpConfig: &vpn.ConnectionTunnelOptionsSpecificationTunnelBgpConfigArgs{
//							LocalAsn:   pulumi.String("1219002"),
//							LocalBgpIp: pulumi.String("169.254.40.1"),
//							TunnelCidr: pulumi.String("169.254.40.0/30"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpn.NewRouteEntry(ctx, "default", &vpn.RouteEntryArgs{
//				VpnGatewayId: pulumi.String(_default.Ids[0]),
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
//
// ## Import
//
// VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g.
//
// ```sh
// $ pulumi import alicloud:vpn/routeEntry:RouteEntry example vpn-abc123456:vco-abc123456:10.0.0.10/24
// ```
type RouteEntry struct {
	pulumi.CustomResourceState

	// The next hop of the destination route.
	NextHop pulumi.StringOutput `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolOutput `pulumi:"publishVpc"`
	// The destination network segment of the destination route.
	RouteDest pulumi.StringOutput `pulumi:"routeDest"`
	// (Available in 1.161.0+) The type of the vpn route entry.
	RouteEntryType pulumi.StringOutput `pulumi:"routeEntryType"`
	// (Available in 1.161.0+) The status of the vpn route entry.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the vpn gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
	// The value should be 0 or 100.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
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
	if args.VpnGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VpnGatewayId'")
	}
	if args.Weight == nil {
		return nil, errors.New("invalid value for required argument 'Weight'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteEntry
	err := ctx.RegisterResource("alicloud:vpn/routeEntry:RouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteEntryState, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	var resource RouteEntry
	err := ctx.ReadResource("alicloud:vpn/routeEntry:RouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteEntry resources.
type routeEntryState struct {
	// The next hop of the destination route.
	NextHop *string `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc *bool `pulumi:"publishVpc"`
	// The destination network segment of the destination route.
	RouteDest *string `pulumi:"routeDest"`
	// (Available in 1.161.0+) The type of the vpn route entry.
	RouteEntryType *string `pulumi:"routeEntryType"`
	// (Available in 1.161.0+) The status of the vpn route entry.
	Status *string `pulumi:"status"`
	// The id of the vpn gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
	// The value should be 0 or 100.
	Weight *int `pulumi:"weight"`
}

type RouteEntryState struct {
	// The next hop of the destination route.
	NextHop pulumi.StringPtrInput
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolPtrInput
	// The destination network segment of the destination route.
	RouteDest pulumi.StringPtrInput
	// (Available in 1.161.0+) The type of the vpn route entry.
	RouteEntryType pulumi.StringPtrInput
	// (Available in 1.161.0+) The status of the vpn route entry.
	Status pulumi.StringPtrInput
	// The id of the vpn gateway.
	VpnGatewayId pulumi.StringPtrInput
	// The value should be 0 or 100.
	Weight pulumi.IntPtrInput
}

func (RouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryState)(nil)).Elem()
}

type routeEntryArgs struct {
	// The next hop of the destination route.
	NextHop string `pulumi:"nextHop"`
	// Whether to issue the destination route to the VPC.
	PublishVpc bool `pulumi:"publishVpc"`
	// The destination network segment of the destination route.
	RouteDest string `pulumi:"routeDest"`
	// The id of the vpn gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
	// The value should be 0 or 100.
	Weight int `pulumi:"weight"`
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// The next hop of the destination route.
	NextHop pulumi.StringInput
	// Whether to issue the destination route to the VPC.
	PublishVpc pulumi.BoolInput
	// The destination network segment of the destination route.
	RouteDest pulumi.StringInput
	// The id of the vpn gateway.
	VpnGatewayId pulumi.StringInput
	// The value should be 0 or 100.
	Weight pulumi.IntInput
}

func (RouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryArgs)(nil)).Elem()
}

type RouteEntryInput interface {
	pulumi.Input

	ToRouteEntryOutput() RouteEntryOutput
	ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput
}

func (*RouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (i *RouteEntry) ToRouteEntryOutput() RouteEntryOutput {
	return i.ToRouteEntryOutputWithContext(context.Background())
}

func (i *RouteEntry) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryOutput)
}

// RouteEntryArrayInput is an input type that accepts RouteEntryArray and RouteEntryArrayOutput values.
// You can construct a concrete instance of `RouteEntryArrayInput` via:
//
//	RouteEntryArray{ RouteEntryArgs{...} }
type RouteEntryArrayInput interface {
	pulumi.Input

	ToRouteEntryArrayOutput() RouteEntryArrayOutput
	ToRouteEntryArrayOutputWithContext(context.Context) RouteEntryArrayOutput
}

type RouteEntryArray []RouteEntryInput

func (RouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryArray) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return i.ToRouteEntryArrayOutputWithContext(context.Background())
}

func (i RouteEntryArray) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryArrayOutput)
}

// RouteEntryMapInput is an input type that accepts RouteEntryMap and RouteEntryMapOutput values.
// You can construct a concrete instance of `RouteEntryMapInput` via:
//
//	RouteEntryMap{ "key": RouteEntryArgs{...} }
type RouteEntryMapInput interface {
	pulumi.Input

	ToRouteEntryMapOutput() RouteEntryMapOutput
	ToRouteEntryMapOutputWithContext(context.Context) RouteEntryMapOutput
}

type RouteEntryMap map[string]RouteEntryInput

func (RouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryMap) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return i.ToRouteEntryMapOutputWithContext(context.Background())
}

func (i RouteEntryMap) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryMapOutput)
}

type RouteEntryOutput struct{ *pulumi.OutputState }

func (RouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (o RouteEntryOutput) ToRouteEntryOutput() RouteEntryOutput {
	return o
}

func (o RouteEntryOutput) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return o
}

// The next hop of the destination route.
func (o RouteEntryOutput) NextHop() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.NextHop }).(pulumi.StringOutput)
}

// Whether to issue the destination route to the VPC.
func (o RouteEntryOutput) PublishVpc() pulumi.BoolOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.BoolOutput { return v.PublishVpc }).(pulumi.BoolOutput)
}

// The destination network segment of the destination route.
func (o RouteEntryOutput) RouteDest() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.RouteDest }).(pulumi.StringOutput)
}

// (Available in 1.161.0+) The type of the vpn route entry.
func (o RouteEntryOutput) RouteEntryType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.RouteEntryType }).(pulumi.StringOutput)
}

// (Available in 1.161.0+) The status of the vpn route entry.
func (o RouteEntryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the vpn gateway.
func (o RouteEntryOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.VpnGatewayId }).(pulumi.StringOutput)
}

// The value should be 0 or 100.
func (o RouteEntryOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type RouteEntryArrayOutput struct{ *pulumi.OutputState }

func (RouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) Index(i pulumi.IntInput) RouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].([]*RouteEntry)[vs[1].(int)]
	}).(RouteEntryOutput)
}

type RouteEntryMapOutput struct{ *pulumi.OutputState }

func (RouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) MapIndex(k pulumi.StringInput) RouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].(map[string]*RouteEntry)[vs[1].(string)]
	}).(RouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryInput)(nil)).Elem(), &RouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryArrayInput)(nil)).Elem(), RouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryMapInput)(nil)).Elem(), RouteEntryMap{})
	pulumi.RegisterOutputType(RouteEntryOutput{})
	pulumi.RegisterOutputType(RouteEntryArrayOutput{})
	pulumi.RegisterOutputType(RouteEntryMapOutput{})
}
