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

// ## Import
//
// VPN gateway can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpn/gateway:Gateway example <id>
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// Whether to pay automatically. Default value: `true`. Valid values:
	AutoPay pulumi.BoolPtrOutput `pulumi:"autoPay"`
	// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
	AutoPropagate pulumi.BoolPtrOutput `pulumi:"autoPropagate"`
	// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The business status of the VPN gateway.
	BusinessStatus pulumi.StringOutput `pulumi:"businessStatus"`
	// The time when the VPN gateway was created.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// The description of the VPN gateway.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The backup public IP address of the VPN gateway. The second IP address assigned by the system to create an IPsec-VPN connection. This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
	DisasterRecoveryInternetIp pulumi.StringOutput `pulumi:"disasterRecoveryInternetIp"`
	// The ID of the backup VSwitch to which the VPN gateway is attached.
	DisasterRecoveryVswitchId pulumi.StringOutput `pulumi:"disasterRecoveryVswitchId"`
	// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
	EnableIpsec pulumi.BoolPtrOutput `pulumi:"enableIpsec"`
	// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	EnableSsl pulumi.BoolPtrOutput `pulumi:"enableSsl"`
	// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// The internet ip of the VPN.
	InternetIp pulumi.StringOutput `pulumi:"internetIp"`
	// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Maximum number of clients.
	SslConnections pulumi.IntOutput `pulumi:"sslConnections"`
	// The IP address of the SSL-VPN connection. This parameter is returned only when the VPN gateway is a public VPN gateway and supports only the single-tunnel mode. In addition, the VPN gateway must have the SSL-VPN feature enabled.
	SslVpnInternetIp pulumi.StringOutput `pulumi:"sslVpnInternetIp"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The Tag of.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the VPC to which the VPN gateway belongs.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringOutput `pulumi:"vpnGatewayName"`
	// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
	VpnType pulumi.StringOutput `pulumi:"vpnType"`
	// The ID of the VSwitch to which the VPN gateway is attached.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("alicloud:vpn/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("alicloud:vpn/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// Whether to pay automatically. Default value: `true`. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
	AutoPropagate *bool `pulumi:"autoPropagate"`
	// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
	Bandwidth *int `pulumi:"bandwidth"`
	// The business status of the VPN gateway.
	BusinessStatus *string `pulumi:"businessStatus"`
	// The time when the VPN gateway was created.
	CreateTime *int `pulumi:"createTime"`
	// The description of the VPN gateway.
	Description *string `pulumi:"description"`
	// The backup public IP address of the VPN gateway. The second IP address assigned by the system to create an IPsec-VPN connection. This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
	DisasterRecoveryInternetIp *string `pulumi:"disasterRecoveryInternetIp"`
	// The ID of the backup VSwitch to which the VPN gateway is attached.
	DisasterRecoveryVswitchId *string `pulumi:"disasterRecoveryVswitchId"`
	// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
	EnableIpsec *bool `pulumi:"enableIpsec"`
	// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	EnableSsl *bool `pulumi:"enableSsl"`
	// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The internet ip of the VPN.
	InternetIp *string `pulumi:"internetIp"`
	// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
	Name *string `pulumi:"name"`
	// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
	NetworkType *string `pulumi:"networkType"`
	// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
	PaymentType *string `pulumi:"paymentType"`
	// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
	Period *int `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Maximum number of clients.
	SslConnections *int `pulumi:"sslConnections"`
	// The IP address of the SSL-VPN connection. This parameter is returned only when the VPN gateway is a public VPN gateway and supports only the single-tunnel mode. In addition, the VPN gateway must have the SSL-VPN feature enabled.
	SslVpnInternetIp *string `pulumi:"sslVpnInternetIp"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The Tag of.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the VPC to which the VPN gateway belongs.
	VpcId *string `pulumi:"vpcId"`
	// The name of the VPN gateway.
	VpnGatewayName *string `pulumi:"vpnGatewayName"`
	// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
	VpnType *string `pulumi:"vpnType"`
	// The ID of the VSwitch to which the VPN gateway is attached.
	VswitchId *string `pulumi:"vswitchId"`
}

type GatewayState struct {
	// Whether to pay automatically. Default value: `true`. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
	AutoPropagate pulumi.BoolPtrInput
	// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
	Bandwidth pulumi.IntPtrInput
	// The business status of the VPN gateway.
	BusinessStatus pulumi.StringPtrInput
	// The time when the VPN gateway was created.
	CreateTime pulumi.IntPtrInput
	// The description of the VPN gateway.
	Description pulumi.StringPtrInput
	// The backup public IP address of the VPN gateway. The second IP address assigned by the system to create an IPsec-VPN connection. This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
	DisasterRecoveryInternetIp pulumi.StringPtrInput
	// The ID of the backup VSwitch to which the VPN gateway is attached.
	DisasterRecoveryVswitchId pulumi.StringPtrInput
	// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
	EnableIpsec pulumi.BoolPtrInput
	// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	EnableSsl pulumi.BoolPtrInput
	// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// The internet ip of the VPN.
	InternetIp pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
	Name pulumi.StringPtrInput
	// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
	NetworkType pulumi.StringPtrInput
	// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
	PaymentType pulumi.StringPtrInput
	// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
	Period pulumi.IntPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Maximum number of clients.
	SslConnections pulumi.IntPtrInput
	// The IP address of the SSL-VPN connection. This parameter is returned only when the VPN gateway is a public VPN gateway and supports only the single-tunnel mode. In addition, the VPN gateway must have the SSL-VPN feature enabled.
	SslVpnInternetIp pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The Tag of.
	Tags pulumi.StringMapInput
	// The ID of the VPC to which the VPN gateway belongs.
	VpcId pulumi.StringPtrInput
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringPtrInput
	// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
	VpnType pulumi.StringPtrInput
	// The ID of the VSwitch to which the VPN gateway is attached.
	VswitchId pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// Whether to pay automatically. Default value: `true`. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
	AutoPropagate *bool `pulumi:"autoPropagate"`
	// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
	Bandwidth int `pulumi:"bandwidth"`
	// The description of the VPN gateway.
	Description *string `pulumi:"description"`
	// The ID of the backup VSwitch to which the VPN gateway is attached.
	DisasterRecoveryVswitchId *string `pulumi:"disasterRecoveryVswitchId"`
	// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
	EnableIpsec *bool `pulumi:"enableIpsec"`
	// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	EnableSsl *bool `pulumi:"enableSsl"`
	// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
	Name *string `pulumi:"name"`
	// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
	NetworkType *string `pulumi:"networkType"`
	// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
	PaymentType *string `pulumi:"paymentType"`
	// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
	Period *int `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Maximum number of clients.
	SslConnections *int `pulumi:"sslConnections"`
	// The Tag of.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the VPC to which the VPN gateway belongs.
	VpcId string `pulumi:"vpcId"`
	// The name of the VPN gateway.
	VpnGatewayName *string `pulumi:"vpnGatewayName"`
	// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
	VpnType *string `pulumi:"vpnType"`
	// The ID of the VSwitch to which the VPN gateway is attached.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// Whether to pay automatically. Default value: `true`. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
	AutoPropagate pulumi.BoolPtrInput
	// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
	Bandwidth pulumi.IntInput
	// The description of the VPN gateway.
	Description pulumi.StringPtrInput
	// The ID of the backup VSwitch to which the VPN gateway is attached.
	DisasterRecoveryVswitchId pulumi.StringPtrInput
	// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
	EnableIpsec pulumi.BoolPtrInput
	// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	EnableSsl pulumi.BoolPtrInput
	// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
	Name pulumi.StringPtrInput
	// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
	NetworkType pulumi.StringPtrInput
	// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
	PaymentType pulumi.StringPtrInput
	// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
	Period pulumi.IntPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Maximum number of clients.
	SslConnections pulumi.IntPtrInput
	// The Tag of.
	Tags pulumi.StringMapInput
	// The ID of the VPC to which the VPN gateway belongs.
	VpcId pulumi.StringInput
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringPtrInput
	// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
	VpnType pulumi.StringPtrInput
	// The ID of the VSwitch to which the VPN gateway is attached.
	VswitchId pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//	GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//	GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// Whether to pay automatically. Default value: `true`. Valid values:
func (o GatewayOutput) AutoPay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.AutoPay }).(pulumi.BoolPtrOutput)
}

// Whether to automatically propagate the BGP route to the VPC. Value:  true: Propagate automatically.  false: does not propagate automatically.
func (o GatewayOutput) AutoPropagate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.AutoPropagate }).(pulumi.BoolPtrOutput)
}

// The Bandwidth specification of the VPN gateway. Unit: Mbps.  If you want to create a public VPN gateway, the value is 5, 10, 20, 50, 100, 200, 500, or 1000. If you want to create a private VPN gateway, the value is 200 or 1000.
func (o GatewayOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The business status of the VPN gateway.
func (o GatewayOutput) BusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.BusinessStatus }).(pulumi.StringOutput)
}

// The time when the VPN gateway was created.
func (o GatewayOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// The description of the VPN gateway.
func (o GatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The backup public IP address of the VPN gateway. The second IP address assigned by the system to create an IPsec-VPN connection. This parameter is returned only when the VPN gateway supports the dual-tunnel mode.
func (o GatewayOutput) DisasterRecoveryInternetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.DisasterRecoveryInternetIp }).(pulumi.StringOutput)
}

// The ID of the backup VSwitch to which the VPN gateway is attached.
func (o GatewayOutput) DisasterRecoveryVswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.DisasterRecoveryVswitchId }).(pulumi.StringOutput)
}

// Enable or Disable IPSec VPN. At least one type of VPN should be enabled.
func (o GatewayOutput) EnableIpsec() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.EnableIpsec }).(pulumi.BoolPtrOutput)
}

// Enable or Disable SSL VPN.  At least one type of VPN should be enabled.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o GatewayOutput) EnableSsl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.BoolPtrOutput { return v.EnableSsl }).(pulumi.BoolPtrOutput)
}

// . Field 'instance_charge_type' has been deprecated from provider version 1.216.0. New field 'payment_type' instead.
//
// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.215.0. New field 'payment_type' instead.
func (o GatewayOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.InstanceChargeType }).(pulumi.StringOutput)
}

// The internet ip of the VPN.
func (o GatewayOutput) InternetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.InternetIp }).(pulumi.StringOutput)
}

// . Field 'name' has been deprecated from provider version 1.216.0. New field 'vpn_gateway_name' instead.
//
// Deprecated: Field 'name' has been deprecated since provider version 1.215.0. New field 'vpn_gateway_name' instead.
func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network type of the VPN gateway. Value:  public (default): public VPN gateway. private: private network VPN gateway.
func (o GatewayOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// Type of payment. Value: Subscription: prepaid PayAsYouGo: Post-paid.
func (o GatewayOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.
func (o GatewayOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The ID of the resource group.
func (o GatewayOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Maximum number of clients.
func (o GatewayOutput) SslConnections() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.SslConnections }).(pulumi.IntOutput)
}

// The IP address of the SSL-VPN connection. This parameter is returned only when the VPN gateway is a public VPN gateway and supports only the single-tunnel mode. In addition, the VPN gateway must have the SSL-VPN feature enabled.
func (o GatewayOutput) SslVpnInternetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.SslVpnInternetIp }).(pulumi.StringOutput)
}

// The status of the resource.
func (o GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The Tag of.
func (o GatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the VPC to which the VPN gateway belongs.
func (o GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The name of the VPN gateway.
func (o GatewayOutput) VpnGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpnGatewayName }).(pulumi.StringOutput)
}

// The VPN gateway type. Value:  Normal (default): Normal type. NationalStandard: National Secret type.
func (o GatewayOutput) VpnType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpnType }).(pulumi.StringOutput)
}

// The ID of the VSwitch to which the VPN gateway is attached.
func (o GatewayOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
