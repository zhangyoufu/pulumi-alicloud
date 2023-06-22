// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an elastic IP resource.
//
// > **DEPRECATED:**  This resource  has been deprecated from version `1.126.0`. Please use new resource alicloud_eip_address.
//
// > **NOTE:** The resource only supports to create `PostPaid PayByTraffic`  or `PrePaid PayByBandwidth` elastic IP for international account. Otherwise, you will happened error `COMMODITY.INVALID_COMPONENT`.
// Your account is international if you can use it to login in [International Web Console](https://account.alibabacloud.com/login/login.htm).
//
// > **NOTE:** From version 1.10.1, this resource supports creating "PrePaid" EIP. In addition, it supports setting EIP name and description.
//
// ## Example Usage
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
//			_, err := ecs.NewEip(ctx, "example", &ecs.EipArgs{
//				Bandwidth:          pulumi.String("10"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Module Support
//
// You can use the existing eip module
// to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.
//
// ## Import
//
// Elastic IP address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/eip:Eip example eip-abc12345678
//
// ```
//
// Deprecated: This resource has been deprecated in favour of the EipAddress resource
type Eip struct {
	pulumi.CustomResourceState

	ActivityId pulumi.StringPtrOutput `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringOutput  `pulumi:"addressName"`
	AutoPay     pulumi.BoolPtrOutput `pulumi:"autoPay"`
	// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
	Bandwidth  pulumi.StringOutput `pulumi:"bandwidth"`
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description                    pulumi.StringPtrOutput `pulumi:"description"`
	HighDefinitionMonitorLogStatus pulumi.StringOutput    `pulumi:"highDefinitionMonitorLogStatus"`
	// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// The elastic ip address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
	Isp        pulumi.StringOutput    `pulumi:"isp"`
	LogProject pulumi.StringPtrOutput `pulumi:"logProject"`
	LogStore   pulumi.StringPtrOutput `pulumi:"logStore"`
	// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
	Name    pulumi.StringOutput `pulumi:"name"`
	Netmode pulumi.StringOutput `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period                pulumi.IntPtrOutput    `pulumi:"period"`
	PricingCycle          pulumi.StringPtrOutput `pulumi:"pricingCycle"`
	PublicIpAddressPoolId pulumi.StringPtrOutput `pulumi:"publicIpAddressPoolId"`
	// The Id of resource group which the eip belongs.
	ResourceGroupId         pulumi.StringOutput      `pulumi:"resourceGroupId"`
	SecurityProtectionTypes pulumi.StringArrayOutput `pulumi:"securityProtectionTypes"`
	// The EIP current status.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput    `pulumi:"tags"`
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewEip registers a new resource with the given unique name, arguments, and options.
func NewEip(ctx *pulumi.Context,
	name string, args *EipArgs, opts ...pulumi.ResourceOption) (*Eip, error) {
	if args == nil {
		args = &EipArgs{}
	}

	var resource Eip
	err := ctx.RegisterResource("alicloud:ecs/eip:Eip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEip gets an existing Eip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipState, opts ...pulumi.ResourceOption) (*Eip, error) {
	var resource Eip
	err := ctx.ReadResource("alicloud:ecs/eip:Eip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eip resources.
type eipState struct {
	ActivityId *string `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName *string `pulumi:"addressName"`
	AutoPay     *bool   `pulumi:"autoPay"`
	// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
	Bandwidth  *string `pulumi:"bandwidth"`
	CreateTime *string `pulumi:"createTime"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description                    *string `pulumi:"description"`
	HighDefinitionMonitorLogStatus *string `pulumi:"highDefinitionMonitorLogStatus"`
	// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The elastic ip address
	IpAddress *string `pulumi:"ipAddress"`
	// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
	Isp        *string `pulumi:"isp"`
	LogProject *string `pulumi:"logProject"`
	LogStore   *string `pulumi:"logStore"`
	// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
	Name    *string `pulumi:"name"`
	Netmode *string `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period                *int    `pulumi:"period"`
	PricingCycle          *string `pulumi:"pricingCycle"`
	PublicIpAddressPoolId *string `pulumi:"publicIpAddressPoolId"`
	// The Id of resource group which the eip belongs.
	ResourceGroupId         *string  `pulumi:"resourceGroupId"`
	SecurityProtectionTypes []string `pulumi:"securityProtectionTypes"`
	// The EIP current status.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	Zone *string                `pulumi:"zone"`
}

type EipState struct {
	ActivityId pulumi.StringPtrInput
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringPtrInput
	AutoPay     pulumi.BoolPtrInput
	// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
	Bandwidth  pulumi.StringPtrInput
	CreateTime pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description                    pulumi.StringPtrInput
	HighDefinitionMonitorLogStatus pulumi.StringPtrInput
	// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
	InternetChargeType pulumi.StringPtrInput
	// The elastic ip address
	IpAddress pulumi.StringPtrInput
	// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
	Isp        pulumi.StringPtrInput
	LogProject pulumi.StringPtrInput
	LogStore   pulumi.StringPtrInput
	// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
	Name    pulumi.StringPtrInput
	Netmode pulumi.StringPtrInput
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period                pulumi.IntPtrInput
	PricingCycle          pulumi.StringPtrInput
	PublicIpAddressPoolId pulumi.StringPtrInput
	// The Id of resource group which the eip belongs.
	ResourceGroupId         pulumi.StringPtrInput
	SecurityProtectionTypes pulumi.StringArrayInput
	// The EIP current status.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	Zone pulumi.StringPtrInput
}

func (EipState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	ActivityId *string `pulumi:"activityId"`
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName *string `pulumi:"addressName"`
	AutoPay     *bool   `pulumi:"autoPay"`
	// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
	Bandwidth *string `pulumi:"bandwidth"`
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description                    *string `pulumi:"description"`
	HighDefinitionMonitorLogStatus *string `pulumi:"highDefinitionMonitorLogStatus"`
	// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
	Isp        *string `pulumi:"isp"`
	LogProject *string `pulumi:"logProject"`
	LogStore   *string `pulumi:"logStore"`
	// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
	Name    *string `pulumi:"name"`
	Netmode *string `pulumi:"netmode"`
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period                *int    `pulumi:"period"`
	PricingCycle          *string `pulumi:"pricingCycle"`
	PublicIpAddressPoolId *string `pulumi:"publicIpAddressPoolId"`
	// The Id of resource group which the eip belongs.
	ResourceGroupId         *string  `pulumi:"resourceGroupId"`
	SecurityProtectionTypes []string `pulumi:"securityProtectionTypes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	Zone *string                `pulumi:"zone"`
}

// The set of arguments for constructing a Eip resource.
type EipArgs struct {
	ActivityId pulumi.StringPtrInput
	// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
	AddressName pulumi.StringPtrInput
	AutoPay     pulumi.BoolPtrInput
	// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
	Bandwidth pulumi.StringPtrInput
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description                    pulumi.StringPtrInput
	HighDefinitionMonitorLogStatus pulumi.StringPtrInput
	// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
	//
	// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
	InstanceChargeType pulumi.StringPtrInput
	// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
	InternetChargeType pulumi.StringPtrInput
	// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
	Isp        pulumi.StringPtrInput
	LogProject pulumi.StringPtrInput
	LogStore   pulumi.StringPtrInput
	// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
	Name    pulumi.StringPtrInput
	Netmode pulumi.StringPtrInput
	// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
	// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period                pulumi.IntPtrInput
	PricingCycle          pulumi.StringPtrInput
	PublicIpAddressPoolId pulumi.StringPtrInput
	// The Id of resource group which the eip belongs.
	ResourceGroupId         pulumi.StringPtrInput
	SecurityProtectionTypes pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	Zone pulumi.StringPtrInput
}

func (EipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EipInput interface {
	pulumi.Input

	ToEipOutput() EipOutput
	ToEipOutputWithContext(ctx context.Context) EipOutput
}

func (*Eip) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (i *Eip) ToEipOutput() EipOutput {
	return i.ToEipOutputWithContext(context.Background())
}

func (i *Eip) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipOutput)
}

// EipArrayInput is an input type that accepts EipArray and EipArrayOutput values.
// You can construct a concrete instance of `EipArrayInput` via:
//
//	EipArray{ EipArgs{...} }
type EipArrayInput interface {
	pulumi.Input

	ToEipArrayOutput() EipArrayOutput
	ToEipArrayOutputWithContext(context.Context) EipArrayOutput
}

type EipArray []EipInput

func (EipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (i EipArray) ToEipArrayOutput() EipArrayOutput {
	return i.ToEipArrayOutputWithContext(context.Background())
}

func (i EipArray) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipArrayOutput)
}

// EipMapInput is an input type that accepts EipMap and EipMapOutput values.
// You can construct a concrete instance of `EipMapInput` via:
//
//	EipMap{ "key": EipArgs{...} }
type EipMapInput interface {
	pulumi.Input

	ToEipMapOutput() EipMapOutput
	ToEipMapOutputWithContext(context.Context) EipMapOutput
}

type EipMap map[string]EipInput

func (EipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (i EipMap) ToEipMapOutput() EipMapOutput {
	return i.ToEipMapOutputWithContext(context.Background())
}

func (i EipMap) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipMapOutput)
}

type EipOutput struct{ *pulumi.OutputState }

func (EipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (o EipOutput) ToEipOutput() EipOutput {
	return o
}

func (o EipOutput) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return o
}

func (o EipOutput) ActivityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.ActivityId }).(pulumi.StringPtrOutput)
}

// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
func (o EipOutput) AddressName() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.AddressName }).(pulumi.StringOutput)
}

func (o EipOutput) AutoPay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.BoolPtrOutput { return v.AutoPay }).(pulumi.BoolPtrOutput)
}

// Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
func (o EipOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Bandwidth }).(pulumi.StringOutput)
}

func (o EipOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether enable the deletion protection or not. Default value: `false`.
// - true: Enable deletion protection.
// - false: Disable deletion protection.
func (o EipOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Eip) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
func (o EipOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EipOutput) HighDefinitionMonitorLogStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.HighDefinitionMonitorLogStatus }).(pulumi.StringOutput)
}

// (It has been deprecated from version 1.126.0 and using new attribute `paymentType` instead) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".
//
// Deprecated: Field 'instance_charge_type' has been deprecated since provider version 1.126.0. New field 'payment_type' instead.
func (o EipOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.InstanceChargeType }).(pulumi.StringOutput)
}

// Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only "PayByBandwidth" when `instanceChargeType` is PrePaid.
func (o EipOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.InternetChargeType }).(pulumi.StringOutput)
}

// The elastic ip address
func (o EipOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
func (o EipOutput) Isp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Isp }).(pulumi.StringOutput)
}

func (o EipOutput) LogProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.LogProject }).(pulumi.StringPtrOutput)
}

func (o EipOutput) LogStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.LogStore }).(pulumi.StringPtrOutput)
}

// It has been deprecated from version 1.126.0 and using new attribute `addressName` instead.
//
// Deprecated: Field 'name' has been deprecated since provider version 1.126.0. New field 'address_name' instead.
func (o EipOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EipOutput) Netmode() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Netmode }).(pulumi.StringOutput)
}

// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
func (o EipOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
// **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
func (o EipOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o EipOutput) PricingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.PricingCycle }).(pulumi.StringPtrOutput)
}

func (o EipOutput) PublicIpAddressPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.PublicIpAddressPoolId }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the eip belongs.
func (o EipOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o EipOutput) SecurityProtectionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringArrayOutput { return v.SecurityProtectionTypes }).(pulumi.StringArrayOutput)
}

// The EIP current status.
func (o EipOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o EipOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Eip) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o EipOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type EipArrayOutput struct{ *pulumi.OutputState }

func (EipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (o EipArrayOutput) ToEipArrayOutput() EipArrayOutput {
	return o
}

func (o EipArrayOutput) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return o
}

func (o EipArrayOutput) Index(i pulumi.IntInput) EipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].([]*Eip)[vs[1].(int)]
	}).(EipOutput)
}

type EipMapOutput struct{ *pulumi.OutputState }

func (EipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (o EipMapOutput) ToEipMapOutput() EipMapOutput {
	return o
}

func (o EipMapOutput) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return o
}

func (o EipMapOutput) MapIndex(k pulumi.StringInput) EipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].(map[string]*Eip)[vs[1].(string)]
	}).(EipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipInput)(nil)).Elem(), &Eip{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipArrayInput)(nil)).Elem(), EipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipMapInput)(nil)).Elem(), EipMap{})
	pulumi.RegisterOutputType(EipOutput{})
	pulumi.RegisterOutputType(EipArrayOutput{})
	pulumi.RegisterOutputType(EipMapOutput{})
}
