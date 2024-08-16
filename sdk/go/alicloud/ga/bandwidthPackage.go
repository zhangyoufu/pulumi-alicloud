// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Bandwidth Package resource.
//
// For information about Global Accelerator (GA) Bandwidth Package and how to use it, see [What is Bandwidth Package](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbandwidthpackage).
//
// > **NOTE:** At present, The `ga.BandwidthPackage` created with `Subscription` cannot be deleted. you need to wait until the resource is outdated and released automatically.
//
// > **NOTE:** Available since v1.112.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ga.NewBandwidthPackage(ctx, "example", &ga.BandwidthPackageArgs{
//				Bandwidth:     pulumi.Int(20),
//				Type:          pulumi.String("Basic"),
//				BandwidthType: pulumi.String("Basic"),
//				Duration:      pulumi.String("1"),
//				AutoPay:       pulumi.Bool(true),
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
// Ga Bandwidth Package can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/bandwidthPackage:BandwidthPackage example <id>
// ```
type BandwidthPackage struct {
	pulumi.CustomResourceState

	// Whether to pay automatically. Valid values:
	AutoPay pulumi.BoolPtrOutput `pulumi:"autoPay"`
	// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
	AutoRenewDuration pulumi.IntOutput `pulumi:"autoRenewDuration"`
	// Whether use vouchers. Default value: `false`. Valid values:
	AutoUseCoupon pulumi.BoolPtrOutput `pulumi:"autoUseCoupon"`
	// The bandwidth value of bandwidth packet.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The name of the bandwidth packet.
	BandwidthPackageName pulumi.StringPtrOutput `pulumi:"bandwidthPackageName"`
	// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
	// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
	BandwidthType pulumi.StringPtrOutput `pulumi:"bandwidthType"`
	// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
	BillingType pulumi.StringPtrOutput `pulumi:"billingType"`
	// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
	CbnGeographicRegionIda pulumi.StringOutput `pulumi:"cbnGeographicRegionIda"`
	// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
	CbnGeographicRegionIdb pulumi.StringOutput `pulumi:"cbnGeographicRegionIdb"`
	// The description of bandwidth package.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrOutput `pulumi:"paymentType"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrOutput `pulumi:"promotionOptionNo"`
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
	Ratio pulumi.IntPtrOutput `pulumi:"ratio"`
	// Whether to renew a bandwidth packet. automatically or not. Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringOutput `pulumi:"renewalStatus"`
	// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the Bandwidth Package.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBandwidthPackage registers a new resource with the given unique name, arguments, and options.
func NewBandwidthPackage(ctx *pulumi.Context,
	name string, args *BandwidthPackageArgs, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BandwidthPackage
	err := ctx.RegisterResource("alicloud:ga/bandwidthPackage:BandwidthPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthPackage gets an existing BandwidthPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthPackageState, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	var resource BandwidthPackage
	err := ctx.ReadResource("alicloud:ga/bandwidthPackage:BandwidthPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthPackage resources.
type bandwidthPackageState struct {
	// Whether to pay automatically. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Whether use vouchers. Default value: `false`. Valid values:
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// The bandwidth value of bandwidth packet.
	Bandwidth *int `pulumi:"bandwidth"`
	// The name of the bandwidth packet.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
	// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
	BandwidthType *string `pulumi:"bandwidthType"`
	// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
	BillingType *string `pulumi:"billingType"`
	// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
	CbnGeographicRegionIda *string `pulumi:"cbnGeographicRegionIda"`
	// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
	CbnGeographicRegionIdb *string `pulumi:"cbnGeographicRegionIdb"`
	// The description of bandwidth package.
	Description *string `pulumi:"description"`
	// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
	Duration *string `pulumi:"duration"`
	// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo *string `pulumi:"promotionOptionNo"`
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
	Ratio *int `pulumi:"ratio"`
	// Whether to renew a bandwidth packet. automatically or not. Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the Bandwidth Package.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
	Type *string `pulumi:"type"`
}

type BandwidthPackageState struct {
	// Whether to pay automatically. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
	AutoRenewDuration pulumi.IntPtrInput
	// Whether use vouchers. Default value: `false`. Valid values:
	AutoUseCoupon pulumi.BoolPtrInput
	// The bandwidth value of bandwidth packet.
	Bandwidth pulumi.IntPtrInput
	// The name of the bandwidth packet.
	BandwidthPackageName pulumi.StringPtrInput
	// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
	// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
	BandwidthType pulumi.StringPtrInput
	// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
	BillingType pulumi.StringPtrInput
	// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
	CbnGeographicRegionIda pulumi.StringPtrInput
	// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
	CbnGeographicRegionIdb pulumi.StringPtrInput
	// The description of bandwidth package.
	Description pulumi.StringPtrInput
	// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
	Duration pulumi.StringPtrInput
	// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrInput
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
	Ratio pulumi.IntPtrInput
	// Whether to renew a bandwidth packet. automatically or not. Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringPtrInput
	// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the Bandwidth Package.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
	Type pulumi.StringPtrInput
}

func (BandwidthPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageState)(nil)).Elem()
}

type bandwidthPackageArgs struct {
	// Whether to pay automatically. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Whether use vouchers. Default value: `false`. Valid values:
	AutoUseCoupon *bool `pulumi:"autoUseCoupon"`
	// The bandwidth value of bandwidth packet.
	Bandwidth int `pulumi:"bandwidth"`
	// The name of the bandwidth packet.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
	// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
	BandwidthType *string `pulumi:"bandwidthType"`
	// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
	BillingType *string `pulumi:"billingType"`
	// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
	CbnGeographicRegionIda *string `pulumi:"cbnGeographicRegionIda"`
	// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
	CbnGeographicRegionIdb *string `pulumi:"cbnGeographicRegionIdb"`
	// The description of bandwidth package.
	Description *string `pulumi:"description"`
	// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
	Duration *string `pulumi:"duration"`
	// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo *string `pulumi:"promotionOptionNo"`
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
	Ratio *int `pulumi:"ratio"`
	// Whether to renew a bandwidth packet. automatically or not. Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a BandwidthPackage resource.
type BandwidthPackageArgs struct {
	// Whether to pay automatically. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
	AutoRenewDuration pulumi.IntPtrInput
	// Whether use vouchers. Default value: `false`. Valid values:
	AutoUseCoupon pulumi.BoolPtrInput
	// The bandwidth value of bandwidth packet.
	Bandwidth pulumi.IntInput
	// The name of the bandwidth packet.
	BandwidthPackageName pulumi.StringPtrInput
	// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
	// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
	BandwidthType pulumi.StringPtrInput
	// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
	BillingType pulumi.StringPtrInput
	// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
	CbnGeographicRegionIda pulumi.StringPtrInput
	// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
	CbnGeographicRegionIdb pulumi.StringPtrInput
	// The description of bandwidth package.
	Description pulumi.StringPtrInput
	// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
	Duration pulumi.StringPtrInput
	// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrInput
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
	Ratio pulumi.IntPtrInput
	// Whether to renew a bandwidth packet. automatically or not. Valid values:
	// - `AutoRenewal`: Enable auto renewal.
	// - `Normal`: Disable auto renewal.
	// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
	RenewalStatus pulumi.StringPtrInput
	// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
	Type pulumi.StringInput
}

func (BandwidthPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageArgs)(nil)).Elem()
}

type BandwidthPackageInput interface {
	pulumi.Input

	ToBandwidthPackageOutput() BandwidthPackageOutput
	ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput
}

func (*BandwidthPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackage)(nil)).Elem()
}

func (i *BandwidthPackage) ToBandwidthPackageOutput() BandwidthPackageOutput {
	return i.ToBandwidthPackageOutputWithContext(context.Background())
}

func (i *BandwidthPackage) ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageOutput)
}

// BandwidthPackageArrayInput is an input type that accepts BandwidthPackageArray and BandwidthPackageArrayOutput values.
// You can construct a concrete instance of `BandwidthPackageArrayInput` via:
//
//	BandwidthPackageArray{ BandwidthPackageArgs{...} }
type BandwidthPackageArrayInput interface {
	pulumi.Input

	ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput
	ToBandwidthPackageArrayOutputWithContext(context.Context) BandwidthPackageArrayOutput
}

type BandwidthPackageArray []BandwidthPackageInput

func (BandwidthPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackage)(nil)).Elem()
}

func (i BandwidthPackageArray) ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput {
	return i.ToBandwidthPackageArrayOutputWithContext(context.Background())
}

func (i BandwidthPackageArray) ToBandwidthPackageArrayOutputWithContext(ctx context.Context) BandwidthPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageArrayOutput)
}

// BandwidthPackageMapInput is an input type that accepts BandwidthPackageMap and BandwidthPackageMapOutput values.
// You can construct a concrete instance of `BandwidthPackageMapInput` via:
//
//	BandwidthPackageMap{ "key": BandwidthPackageArgs{...} }
type BandwidthPackageMapInput interface {
	pulumi.Input

	ToBandwidthPackageMapOutput() BandwidthPackageMapOutput
	ToBandwidthPackageMapOutputWithContext(context.Context) BandwidthPackageMapOutput
}

type BandwidthPackageMap map[string]BandwidthPackageInput

func (BandwidthPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackage)(nil)).Elem()
}

func (i BandwidthPackageMap) ToBandwidthPackageMapOutput() BandwidthPackageMapOutput {
	return i.ToBandwidthPackageMapOutputWithContext(context.Background())
}

func (i BandwidthPackageMap) ToBandwidthPackageMapOutputWithContext(ctx context.Context) BandwidthPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageMapOutput)
}

type BandwidthPackageOutput struct{ *pulumi.OutputState }

func (BandwidthPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageOutput) ToBandwidthPackageOutput() BandwidthPackageOutput {
	return o
}

func (o BandwidthPackageOutput) ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput {
	return o
}

// Whether to pay automatically. Valid values:
func (o BandwidthPackageOutput) AutoPay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.BoolPtrOutput { return v.AutoPay }).(pulumi.BoolPtrOutput)
}

// Auto renewal period of a bandwidth packet, in the unit of month. Valid values: `1` to `12`.
func (o BandwidthPackageOutput) AutoRenewDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntOutput { return v.AutoRenewDuration }).(pulumi.IntOutput)
}

// Whether use vouchers. Default value: `false`. Valid values:
func (o BandwidthPackageOutput) AutoUseCoupon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.BoolPtrOutput { return v.AutoUseCoupon }).(pulumi.BoolPtrOutput)
}

// The bandwidth value of bandwidth packet.
func (o BandwidthPackageOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The name of the bandwidth packet.
func (o BandwidthPackageOutput) BandwidthPackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.BandwidthPackageName }).(pulumi.StringPtrOutput)
}

// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
// > **NOTE:** At present, only basic can be configured to enhanced, but not enhanced and advanced to other types of accelerated bandwidth.
func (o BandwidthPackageOutput) BandwidthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.BandwidthType }).(pulumi.StringPtrOutput)
}

// The billing type. Valid values: `PayBy95`, `PayByTraffic`. **NOTE:** `billingType` is valid only when `paymentType` is set to `PayAsYouGo`.
func (o BandwidthPackageOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.BillingType }).(pulumi.StringPtrOutput)
}

// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value: `China-mainland`.
func (o BandwidthPackageOutput) CbnGeographicRegionIda() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.CbnGeographicRegionIda }).(pulumi.StringOutput)
}

// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value: `Global`.
func (o BandwidthPackageOutput) CbnGeographicRegionIdb() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.CbnGeographicRegionIdb }).(pulumi.StringOutput)
}

// The description of bandwidth package.
func (o BandwidthPackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `paymentType` is set to `Subscription`, this parameter is required.
func (o BandwidthPackageOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// The payment type of the bandwidth. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
func (o BandwidthPackageOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.PaymentType }).(pulumi.StringPtrOutput)
}

// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
func (o BandwidthPackageOutput) PromotionOptionNo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.PromotionOptionNo }).(pulumi.StringPtrOutput)
}

// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: `30` to `100`. **NOTE:** `ratio` is valid only when `billingType` is set to `PayBy95`.
func (o BandwidthPackageOutput) Ratio() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntPtrOutput { return v.Ratio }).(pulumi.IntPtrOutput)
}

// Whether to renew a bandwidth packet. automatically or not. Valid values:
// - `AutoRenewal`: Enable auto renewal.
// - `Normal`: Disable auto renewal.
// - `NotRenewal`: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.
func (o BandwidthPackageOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

// The ID of the resource group. **Note:** Once you set a value of this property, you cannot set it to an empty string anymore.
func (o BandwidthPackageOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the Bandwidth Package.
func (o BandwidthPackageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o BandwidthPackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
func (o BandwidthPackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type BandwidthPackageArrayOutput struct{ *pulumi.OutputState }

func (BandwidthPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageArrayOutput) ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput {
	return o
}

func (o BandwidthPackageArrayOutput) ToBandwidthPackageArrayOutputWithContext(ctx context.Context) BandwidthPackageArrayOutput {
	return o
}

func (o BandwidthPackageArrayOutput) Index(i pulumi.IntInput) BandwidthPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BandwidthPackage {
		return vs[0].([]*BandwidthPackage)[vs[1].(int)]
	}).(BandwidthPackageOutput)
}

type BandwidthPackageMapOutput struct{ *pulumi.OutputState }

func (BandwidthPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageMapOutput) ToBandwidthPackageMapOutput() BandwidthPackageMapOutput {
	return o
}

func (o BandwidthPackageMapOutput) ToBandwidthPackageMapOutputWithContext(ctx context.Context) BandwidthPackageMapOutput {
	return o
}

func (o BandwidthPackageMapOutput) MapIndex(k pulumi.StringInput) BandwidthPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BandwidthPackage {
		return vs[0].(map[string]*BandwidthPackage)[vs[1].(string)]
	}).(BandwidthPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageInput)(nil)).Elem(), &BandwidthPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageArrayInput)(nil)).Elem(), BandwidthPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageMapInput)(nil)).Elem(), BandwidthPackageMap{})
	pulumi.RegisterOutputType(BandwidthPackageOutput{})
	pulumi.RegisterOutputType(BandwidthPackageArrayOutput{})
	pulumi.RegisterOutputType(BandwidthPackageMapOutput{})
}
