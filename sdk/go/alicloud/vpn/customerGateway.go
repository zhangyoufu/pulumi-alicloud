// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vpn.NewCustomerGateway(ctx, "foo", &vpn.CustomerGatewayArgs{
// 			Description: pulumi.String("vpnCgwDescriptionExample"),
// 			IpAddress:   pulumi.String("43.104.22.228"),
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
// VPN customer gateway can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example cgw-abc123456
// ```
type CustomerGateway struct {
	pulumi.CustomResourceState

	// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
	Asn pulumi.StringPtrOutput `pulumi:"asn"`
	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	var resource CustomerGateway
	err := ctx.RegisterResource("alicloud:vpn/customerGateway:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("alicloud:vpn/customerGateway:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
	// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
	Asn *string `pulumi:"asn"`
	// The description of the VPN customer gateway instance.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress *string `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name *string `pulumi:"name"`
}

type CustomerGatewayState struct {
	// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
	Asn pulumi.StringPtrInput
	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringPtrInput
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringPtrInput
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
	Asn *string `pulumi:"asn"`
	// The description of the VPN customer gateway instance.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress string `pulumi:"ipAddress"`
	// The name of the VPN customer gateway. Defaults to null.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
	Asn pulumi.StringPtrInput
	// The description of the VPN customer gateway instance.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringInput
	// The name of the VPN customer gateway. Defaults to null.
	Name pulumi.StringPtrInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

// CustomerGatewayArrayInput is an input type that accepts CustomerGatewayArray and CustomerGatewayArrayOutput values.
// You can construct a concrete instance of `CustomerGatewayArrayInput` via:
//
//          CustomerGatewayArray{ CustomerGatewayArgs{...} }
type CustomerGatewayArrayInput interface {
	pulumi.Input

	ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput
	ToCustomerGatewayArrayOutputWithContext(context.Context) CustomerGatewayArrayOutput
}

type CustomerGatewayArray []CustomerGatewayInput

func (CustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return i.ToCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayArrayOutput)
}

// CustomerGatewayMapInput is an input type that accepts CustomerGatewayMap and CustomerGatewayMapOutput values.
// You can construct a concrete instance of `CustomerGatewayMapInput` via:
//
//          CustomerGatewayMap{ "key": CustomerGatewayArgs{...} }
type CustomerGatewayMapInput interface {
	pulumi.Input

	ToCustomerGatewayMapOutput() CustomerGatewayMapOutput
	ToCustomerGatewayMapOutputWithContext(context.Context) CustomerGatewayMapOutput
}

type CustomerGatewayMap map[string]CustomerGatewayInput

func (CustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return i.ToCustomerGatewayMapOutputWithContext(context.Background())
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayMapOutput)
}

type CustomerGatewayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

// The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
func (o CustomerGatewayOutput) Asn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.Asn }).(pulumi.StringPtrOutput)
}

// The description of the VPN customer gateway instance.
func (o CustomerGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IP address of the customer gateway.
func (o CustomerGatewayOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The name of the VPN customer gateway. Defaults to null.
func (o CustomerGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type CustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) Index(i pulumi.IntInput) CustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].([]*CustomerGateway)[vs[1].(int)]
	}).(CustomerGatewayOutput)
}

type CustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (CustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) CustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].(map[string]*CustomerGateway)[vs[1].(string)]
	}).(CustomerGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayInput)(nil)).Elem(), &CustomerGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayArrayInput)(nil)).Elem(), CustomerGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayMapInput)(nil)).Elem(), CustomerGatewayMap{})
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayMapOutput{})
}
