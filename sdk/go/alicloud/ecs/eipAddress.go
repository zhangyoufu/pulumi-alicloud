// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a EIP Address resource.
//
// For information about EIP Address and how to use it, see [What is EIP Address](https://www.alibabacloud.com/help/en/doc-detail/36016.htm).
//
// > **NOTE:** Available in v1.126.0+.
//
// > **NOTE:** BGP (Multi-ISP) lines are supported in all regions. BGP (Multi-ISP) Pro lines are supported only in the China (Hong Kong) region.
//
// > **NOTE:** The resource only supports to create `PayAsYouGo PayByTraffic`  or `Subscription PayByBandwidth` elastic IP for international account. Otherwise, you will happened error `COMMODITY.INVALID_COMPONENT`.
// Your account is international if you can use it to login in [International Web Console](https://account.alibabacloud.com/login/login.htm).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewEipAddress(ctx, "example", &ecs.EipAddressArgs{
//				AddressName:        pulumi.String("tf-testAcc1234"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//				Isp:                pulumi.String("BGP"),
//				PaymentType:        pulumi.String("PayAsYouGo"),
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
// EIP Address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/eipAddress:EipAddress example <id>
//
// ```
type EipAddress struct {
	pulumi.CustomResourceState

	// The activity id.
	ActivityId pulumi.StringPtrOutput `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringOutput `pulumi:"addressName"`
	// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
	AutoPay pulumi.BoolPtrOutput `pulumi:"autoPay"`
	// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// Whether enable the deletion protection or not. Default value: `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// The description of the EIP.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// The metering method of the EIP.
	// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// The address of the EIP.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
	Isp pulumi.StringOutput `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the network. Valid value is `public` (Internet).
	Netmode pulumi.StringPtrOutput `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEipAddress registers a new resource with the given unique name, arguments, and options.
func NewEipAddress(ctx *pulumi.Context,
	name string, args *EipAddressArgs, opts ...pulumi.ResourceOption) (*EipAddress, error) {
	if args == nil {
		args = &EipAddressArgs{}
	}

	var resource EipAddress
	err := ctx.RegisterResource("alicloud:ecs/eipAddress:EipAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEipAddress gets an existing EipAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEipAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipAddressState, opts ...pulumi.ResourceOption) (*EipAddress, error) {
	var resource EipAddress
	err := ctx.ReadResource("alicloud:ecs/eipAddress:EipAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EipAddress resources.
type eipAddressState struct {
	// The activity id.
	ActivityId *string `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName *string `pulumi:"addressName"`
	// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
	AutoPay *bool `pulumi:"autoPay"`
	// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
	Bandwidth *string `pulumi:"bandwidth"`
	// Whether enable the deletion protection or not. Default value: `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the EIP.
	Description *string `pulumi:"description"`
	// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The metering method of the EIP.
	// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The address of the EIP.
	IpAddress *string `pulumi:"ipAddress"`
	// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
	Isp *string `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
	Name *string `pulumi:"name"`
	// The type of the network. Valid value is `public` (Internet).
	Netmode *string `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period *int `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EipAddressState struct {
	// The activity id.
	ActivityId pulumi.StringPtrInput
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringPtrInput
	// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
	AutoPay pulumi.BoolPtrInput
	// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
	Bandwidth pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	DeletionProtection pulumi.BoolPtrInput
	// The description of the EIP.
	Description pulumi.StringPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// The metering method of the EIP.
	// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
	InternetChargeType pulumi.StringPtrInput
	// The address of the EIP.
	IpAddress pulumi.StringPtrInput
	// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
	Isp pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
	Name pulumi.StringPtrInput
	// The type of the network. Valid value is `public` (Internet).
	Netmode pulumi.StringPtrInput
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period pulumi.IntPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EipAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAddressState)(nil)).Elem()
}

type eipAddressArgs struct {
	// The activity id.
	ActivityId *string `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName *string `pulumi:"addressName"`
	// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
	AutoPay *bool `pulumi:"autoPay"`
	// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
	Bandwidth *string `pulumi:"bandwidth"`
	// Whether enable the deletion protection or not. Default value: `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the EIP.
	Description *string `pulumi:"description"`
	// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The metering method of the EIP.
	// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
	Isp *string `pulumi:"isp"`
	// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
	Name *string `pulumi:"name"`
	// The type of the network. Valid value is `public` (Internet).
	Netmode *string `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period *int `pulumi:"period"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EipAddress resource.
type EipAddressArgs struct {
	// The activity id.
	ActivityId pulumi.StringPtrInput
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringPtrInput
	// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
	AutoPay pulumi.BoolPtrInput
	// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
	Bandwidth pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	DeletionProtection pulumi.BoolPtrInput
	// The description of the EIP.
	Description pulumi.StringPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// The metering method of the EIP.
	// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
	InternetChargeType pulumi.StringPtrInput
	// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
	Isp pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
	Name pulumi.StringPtrInput
	// The type of the network. Valid value is `public` (Internet).
	Netmode pulumi.StringPtrInput
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	Period pulumi.IntPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EipAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAddressArgs)(nil)).Elem()
}

type EipAddressInput interface {
	pulumi.Input

	ToEipAddressOutput() EipAddressOutput
	ToEipAddressOutputWithContext(ctx context.Context) EipAddressOutput
}

func (*EipAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAddress)(nil)).Elem()
}

func (i *EipAddress) ToEipAddressOutput() EipAddressOutput {
	return i.ToEipAddressOutputWithContext(context.Background())
}

func (i *EipAddress) ToEipAddressOutputWithContext(ctx context.Context) EipAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAddressOutput)
}

// EipAddressArrayInput is an input type that accepts EipAddressArray and EipAddressArrayOutput values.
// You can construct a concrete instance of `EipAddressArrayInput` via:
//
//	EipAddressArray{ EipAddressArgs{...} }
type EipAddressArrayInput interface {
	pulumi.Input

	ToEipAddressArrayOutput() EipAddressArrayOutput
	ToEipAddressArrayOutputWithContext(context.Context) EipAddressArrayOutput
}

type EipAddressArray []EipAddressInput

func (EipAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EipAddress)(nil)).Elem()
}

func (i EipAddressArray) ToEipAddressArrayOutput() EipAddressArrayOutput {
	return i.ToEipAddressArrayOutputWithContext(context.Background())
}

func (i EipAddressArray) ToEipAddressArrayOutputWithContext(ctx context.Context) EipAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAddressArrayOutput)
}

// EipAddressMapInput is an input type that accepts EipAddressMap and EipAddressMapOutput values.
// You can construct a concrete instance of `EipAddressMapInput` via:
//
//	EipAddressMap{ "key": EipAddressArgs{...} }
type EipAddressMapInput interface {
	pulumi.Input

	ToEipAddressMapOutput() EipAddressMapOutput
	ToEipAddressMapOutputWithContext(context.Context) EipAddressMapOutput
}

type EipAddressMap map[string]EipAddressInput

func (EipAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EipAddress)(nil)).Elem()
}

func (i EipAddressMap) ToEipAddressMapOutput() EipAddressMapOutput {
	return i.ToEipAddressMapOutputWithContext(context.Background())
}

func (i EipAddressMap) ToEipAddressMapOutputWithContext(ctx context.Context) EipAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAddressMapOutput)
}

type EipAddressOutput struct{ *pulumi.OutputState }

func (EipAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAddress)(nil)).Elem()
}

func (o EipAddressOutput) ToEipAddressOutput() EipAddressOutput {
	return o
}

func (o EipAddressOutput) ToEipAddressOutputWithContext(ctx context.Context) EipAddressOutput {
	return o
}

// The activity id.
func (o EipAddressOutput) ActivityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringPtrOutput { return v.ActivityId }).(pulumi.StringPtrOutput)
}

// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
func (o EipAddressOutput) AddressName() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.AddressName }).(pulumi.StringOutput)
}

// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `autoPay` is `true`, The order will be automatically paid. When `autoPay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `paymentType` is `Subscription`, this parameter is valid.
func (o EipAddressOutput) AutoPay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.BoolPtrOutput { return v.AutoPay }).(pulumi.BoolPtrOutput)
}

// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
func (o EipAddressOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.Bandwidth }).(pulumi.StringOutput)
}

// Whether enable the deletion protection or not. Default value: `false`.
func (o EipAddressOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// The description of the EIP.
func (o EipAddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Field `instanceChargeType` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `paymentType` instead.
//
// Deprecated: Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
func (o EipAddressOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.InstanceChargeType }).(pulumi.StringOutput)
}

// The metering method of the EIP.
// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `paymentType` is "Subscription".
func (o EipAddressOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.InternetChargeType }).(pulumi.StringOutput)
}

// The address of the EIP.
func (o EipAddressOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
func (o EipAddressOutput) Isp() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.Isp }).(pulumi.StringOutput)
}

// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `addressName` instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead.
func (o EipAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the network. Valid value is `public` (Internet).
func (o EipAddressOutput) Netmode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringPtrOutput { return v.Netmode }).(pulumi.StringPtrOutput)
}

// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
func (o EipAddressOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
func (o EipAddressOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The ID of the resource group.
func (o EipAddressOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
func (o EipAddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o EipAddressOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EipAddress) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type EipAddressArrayOutput struct{ *pulumi.OutputState }

func (EipAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EipAddress)(nil)).Elem()
}

func (o EipAddressArrayOutput) ToEipAddressArrayOutput() EipAddressArrayOutput {
	return o
}

func (o EipAddressArrayOutput) ToEipAddressArrayOutputWithContext(ctx context.Context) EipAddressArrayOutput {
	return o
}

func (o EipAddressArrayOutput) Index(i pulumi.IntInput) EipAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EipAddress {
		return vs[0].([]*EipAddress)[vs[1].(int)]
	}).(EipAddressOutput)
}

type EipAddressMapOutput struct{ *pulumi.OutputState }

func (EipAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EipAddress)(nil)).Elem()
}

func (o EipAddressMapOutput) ToEipAddressMapOutput() EipAddressMapOutput {
	return o
}

func (o EipAddressMapOutput) ToEipAddressMapOutputWithContext(ctx context.Context) EipAddressMapOutput {
	return o
}

func (o EipAddressMapOutput) MapIndex(k pulumi.StringInput) EipAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EipAddress {
		return vs[0].(map[string]*EipAddress)[vs[1].(string)]
	}).(EipAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipAddressInput)(nil)).Elem(), &EipAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipAddressArrayInput)(nil)).Elem(), EipAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipAddressMapInput)(nil)).Elem(), EipAddressMap{})
	pulumi.RegisterOutputType(EipAddressOutput{})
	pulumi.RegisterOutputType(EipAddressArrayOutput{})
	pulumi.RegisterOutputType(EipAddressMapOutput{})
}
