// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Nat gateway can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/natGateway:NatGateway example <id>
// ```
type NatGateway struct {
	pulumi.CustomResourceState

	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to only precheck this request. Default value: `false`.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
	EipBindMode pulumi.StringOutput `pulumi:"eipBindMode"`
	// Specifies whether to forcefully delete the NAT gateway.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The nat gateway will auto create a forward item.
	ForwardTableIds pulumi.StringOutput `pulumi:"forwardTableIds"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringOutput `pulumi:"natGatewayName"`
	// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
	NatType pulumi.StringOutput `pulumi:"natType"`
	// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
	// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The nat gateway will auto create a snat item.
	SnatTableIds pulumi.StringOutput `pulumi:"snatTableIds"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
	Specification pulumi.StringOutput `pulumi:"specification"`
	// (Available since v1.121.0) The status of NAT gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of NAT gateway.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NatGateway
	err := ctx.RegisterResource("alicloud:vpc/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("alicloud:vpc/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description *string `pulumi:"description"`
	// Specifies whether to only precheck this request. Default value: `false`.
	DryRun *bool `pulumi:"dryRun"`
	// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
	EipBindMode *string `pulumi:"eipBindMode"`
	// Specifies whether to forcefully delete the NAT gateway.
	Force *bool `pulumi:"force"`
	// The nat gateway will auto create a forward item.
	ForwardTableIds *string `pulumi:"forwardTableIds"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name *string `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
	NatType *string `pulumi:"natType"`
	// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
	NetworkType *string `pulumi:"networkType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
	// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period *int `pulumi:"period"`
	// The nat gateway will auto create a snat item.
	SnatTableIds *string `pulumi:"snatTableIds"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
	Specification *string `pulumi:"specification"`
	// (Available since v1.121.0) The status of NAT gateway.
	Status *string `pulumi:"status"`
	// The tags of NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type NatGatewayState struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrInput
	// Specifies whether to only precheck this request. Default value: `false`.
	DryRun pulumi.BoolPtrInput
	// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
	EipBindMode pulumi.StringPtrInput
	// Specifies whether to forcefully delete the NAT gateway.
	Force pulumi.BoolPtrInput
	// The nat gateway will auto create a forward item.
	ForwardTableIds pulumi.StringPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringPtrInput
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
	InternetChargeType pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringPtrInput
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringPtrInput
	// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
	NatType pulumi.StringPtrInput
	// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
	NetworkType pulumi.StringPtrInput
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
	// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period pulumi.IntPtrInput
	// The nat gateway will auto create a snat item.
	SnatTableIds pulumi.StringPtrInput
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
	Specification pulumi.StringPtrInput
	// (Available since v1.121.0) The status of NAT gateway.
	Status pulumi.StringPtrInput
	// The tags of NAT gateway.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description *string `pulumi:"description"`
	// Specifies whether to only precheck this request. Default value: `false`.
	DryRun *bool `pulumi:"dryRun"`
	// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
	EipBindMode *string `pulumi:"eipBindMode"`
	// Specifies whether to forcefully delete the NAT gateway.
	Force *bool `pulumi:"force"`
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name *string `pulumi:"name"`
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
	NatType *string `pulumi:"natType"`
	// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
	NetworkType *string `pulumi:"networkType"`
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
	// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period *int `pulumi:"period"`
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
	Specification *string `pulumi:"specification"`
	// The tags of NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// Whether enable the deletion protection or not. Default value: `false`.
	// - true: Enable deletion protection.
	// - false: Disable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
	Description pulumi.StringPtrInput
	// Specifies whether to only precheck this request. Default value: `false`.
	DryRun pulumi.BoolPtrInput
	// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
	EipBindMode pulumi.StringPtrInput
	// Specifies whether to forcefully delete the NAT gateway.
	Force pulumi.BoolPtrInput
	// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
	InstanceChargeType pulumi.StringPtrInput
	// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
	InternetChargeType pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
	Name pulumi.StringPtrInput
	// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
	NatGatewayName pulumi.StringPtrInput
	// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
	NatType pulumi.StringPtrInput
	// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
	NetworkType pulumi.StringPtrInput
	// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
	// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
	Period pulumi.IntPtrInput
	// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
	Specification pulumi.StringPtrInput
	// The tags of NAT gateway.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

// NatGatewayArrayInput is an input type that accepts NatGatewayArray and NatGatewayArrayOutput values.
// You can construct a concrete instance of `NatGatewayArrayInput` via:
//
//	NatGatewayArray{ NatGatewayArgs{...} }
type NatGatewayArrayInput interface {
	pulumi.Input

	ToNatGatewayArrayOutput() NatGatewayArrayOutput
	ToNatGatewayArrayOutputWithContext(context.Context) NatGatewayArrayOutput
}

type NatGatewayArray []NatGatewayInput

func (NatGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGateway)(nil)).Elem()
}

func (i NatGatewayArray) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return i.ToNatGatewayArrayOutputWithContext(context.Background())
}

func (i NatGatewayArray) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayArrayOutput)
}

// NatGatewayMapInput is an input type that accepts NatGatewayMap and NatGatewayMapOutput values.
// You can construct a concrete instance of `NatGatewayMapInput` via:
//
//	NatGatewayMap{ "key": NatGatewayArgs{...} }
type NatGatewayMapInput interface {
	pulumi.Input

	ToNatGatewayMapOutput() NatGatewayMapOutput
	ToNatGatewayMapOutputWithContext(context.Context) NatGatewayMapOutput
}

type NatGatewayMap map[string]NatGatewayInput

func (NatGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGateway)(nil)).Elem()
}

func (i NatGatewayMap) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return i.ToNatGatewayMapOutputWithContext(context.Background())
}

func (i NatGatewayMap) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayMapOutput)
}

type NatGatewayOutput struct{ *pulumi.OutputState }

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

// Whether enable the deletion protection or not. Default value: `false`.
// - true: Enable deletion protection.
// - false: Disable deletion protection.
func (o NatGatewayOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.
func (o NatGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to only precheck this request. Default value: `false`.
func (o NatGatewayOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The EIP binding mode of the NAT gateway. Default value: `MULTI_BINDED`. Valid values:
func (o NatGatewayOutput) EipBindMode() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.EipBindMode }).(pulumi.StringOutput)
}

// Specifies whether to forcefully delete the NAT gateway.
func (o NatGatewayOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// The nat gateway will auto create a forward item.
func (o NatGatewayOutput) ForwardTableIds() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.ForwardTableIds }).(pulumi.StringOutput)
}

// Field `instanceChargeType` has been deprecated from provider version 1.121.0. New field `paymentType` instead.
func (o NatGatewayOutput) InstanceChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.InstanceChargeType }).(pulumi.StringOutput)
}

// The internet charge type. Valid values `PayByLcu` and `PayBySpec`. The `PayByLcu` is only support enhanced NAT. **NOTE:** From 1.137.0+, The `PayBySpec` has been deprecated.
func (o NatGatewayOutput) InternetChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.InternetChargeType }).(pulumi.StringOutput)
}

// Field `name` has been deprecated from provider version 1.121.0. New field `natGatewayName` instead.
func (o NatGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.
func (o NatGatewayOutput) NatGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.NatGatewayName }).(pulumi.StringOutput)
}

// The type of NAT gateway. Valid values: `Normal` and `Enhanced`. **NOTE:** From 1.137.0+,  The `Normal` has been deprecated.
func (o NatGatewayOutput) NatType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.NatType }).(pulumi.StringOutput)
}

// Indicates the type of the created NAT gateway. Valid values `internet` and `intranet`. `internet`: Internet NAT Gateway. `intranet`: VPC NAT Gateway.
func (o NatGatewayOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// The billing method of the NAT gateway. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`.
func (o NatGatewayOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The duration that you will buy the resource, in month. It is valid when `paymentType` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. **NOTE:** International station only supports `Subscription`.
// > **NOTE:** The attribute `period` is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription. Once effect, it will not be modified that means running `pulumi up` will not effect the resource.
func (o NatGatewayOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The nat gateway will auto create a snat item.
func (o NatGatewayOutput) SnatTableIds() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.SnatTableIds }).(pulumi.StringOutput)
}

// The specification of the nat gateway. Valid values are `Small`, `Middle` and `Large`. Effective when `internetChargeType` is `PayBySpec` and `networkType` is `internet`. Details refer to [Nat Gateway Specification](https://help.aliyun.com/document_detail/203500.html).
func (o NatGatewayOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Specification }).(pulumi.StringOutput)
}

// (Available since v1.121.0) The status of NAT gateway.
func (o NatGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of NAT gateway.
func (o NatGatewayOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The VPC ID.
func (o NatGatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The id of VSwitch.
func (o NatGatewayOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.VswitchId }).(pulumi.StringPtrOutput)
}

type NatGatewayArrayOutput struct{ *pulumi.OutputState }

func (NatGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGateway)(nil)).Elem()
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) Index(i pulumi.IntInput) NatGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NatGateway {
		return vs[0].([]*NatGateway)[vs[1].(int)]
	}).(NatGatewayOutput)
}

type NatGatewayMapOutput struct{ *pulumi.OutputState }

func (NatGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGateway)(nil)).Elem()
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) MapIndex(k pulumi.StringInput) NatGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NatGateway {
		return vs[0].(map[string]*NatGateway)[vs[1].(string)]
	}).(NatGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayInput)(nil)).Elem(), &NatGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayArrayInput)(nil)).Elem(), NatGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayMapInput)(nil)).Elem(), NatGatewayMap{})
	pulumi.RegisterOutputType(NatGatewayOutput{})
	pulumi.RegisterOutputType(NatGatewayArrayOutput{})
	pulumi.RegisterOutputType(NatGatewayMapOutput{})
}
