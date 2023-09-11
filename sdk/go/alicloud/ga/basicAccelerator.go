// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Global Accelerator (GA) Basic Accelerator resource.
//
// For information about Global Accelerator (GA) Basic Accelerator and how to use it, see [What is Basic Accelerator](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicaccelerator).
//
// > **NOTE:** Available since v1.194.0.
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
//			_, err := ga.NewBasicAccelerator(ctx, "default", &ga.BasicAcceleratorArgs{
//				AutoPay:              pulumi.Bool(true),
//				AutoUseCoupon:        pulumi.String("true"),
//				BandwidthBillingType: pulumi.String("BandwidthPackage"),
//				BasicAcceleratorName: pulumi.String("tf-example-value"),
//				Description:          pulumi.String("tf-example-value"),
//				Duration:             pulumi.Int(1),
//				PricingCycle:         pulumi.String("Month"),
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
// Global Accelerator (GA) Basic Accelerator can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/basicAccelerator:BasicAccelerator example <id>
//
// ```
type BasicAccelerator struct {
	pulumi.CustomResourceState

	// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
	AutoPay pulumi.BoolPtrOutput `pulumi:"autoPay"`
	// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
	AutoRenewDuration pulumi.IntPtrOutput `pulumi:"autoRenewDuration"`
	// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
	AutoUseCoupon pulumi.StringPtrOutput `pulumi:"autoUseCoupon"`
	// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
	BandwidthBillingType pulumi.StringPtrOutput `pulumi:"bandwidthBillingType"`
	// The name of the Global Accelerator Basic Accelerator instance.
	BasicAcceleratorName pulumi.StringPtrOutput `pulumi:"basicAcceleratorName"`
	// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
	CrossBorderStatus pulumi.BoolPtrOutput `pulumi:"crossBorderStatus"`
	// The description of the Global Accelerator Basic Accelerator instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The subscription duration. Default value: `1`.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
	PricingCycle pulumi.StringPtrOutput `pulumi:"pricingCycle"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrOutput `pulumi:"promotionOptionNo"`
	// The status of the Basic Accelerator instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewBasicAccelerator registers a new resource with the given unique name, arguments, and options.
func NewBasicAccelerator(ctx *pulumi.Context,
	name string, args *BasicAcceleratorArgs, opts ...pulumi.ResourceOption) (*BasicAccelerator, error) {
	if args == nil {
		args = &BasicAcceleratorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BasicAccelerator
	err := ctx.RegisterResource("alicloud:ga/basicAccelerator:BasicAccelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicAccelerator gets an existing BasicAccelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicAcceleratorState, opts ...pulumi.ResourceOption) (*BasicAccelerator, error) {
	var resource BasicAccelerator
	err := ctx.ReadResource("alicloud:ga/basicAccelerator:BasicAccelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicAccelerator resources.
type basicAcceleratorState struct {
	// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
	AutoRenew *bool `pulumi:"autoRenew"`
	// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
	AutoUseCoupon *string `pulumi:"autoUseCoupon"`
	// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
	BandwidthBillingType *string `pulumi:"bandwidthBillingType"`
	// The name of the Global Accelerator Basic Accelerator instance.
	BasicAcceleratorName *string `pulumi:"basicAcceleratorName"`
	// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
	CrossBorderStatus *bool `pulumi:"crossBorderStatus"`
	// The description of the Global Accelerator Basic Accelerator instance.
	Description *string `pulumi:"description"`
	// The subscription duration. Default value: `1`.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
	Duration *int `pulumi:"duration"`
	// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo *string `pulumi:"promotionOptionNo"`
	// The status of the Basic Accelerator instance.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type BasicAcceleratorState struct {
	// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
	AutoRenew pulumi.BoolPtrInput
	// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
	AutoRenewDuration pulumi.IntPtrInput
	// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
	AutoUseCoupon pulumi.StringPtrInput
	// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
	BandwidthBillingType pulumi.StringPtrInput
	// The name of the Global Accelerator Basic Accelerator instance.
	BasicAcceleratorName pulumi.StringPtrInput
	// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
	CrossBorderStatus pulumi.BoolPtrInput
	// The description of the Global Accelerator Basic Accelerator instance.
	Description pulumi.StringPtrInput
	// The subscription duration. Default value: `1`.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
	Duration pulumi.IntPtrInput
	// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
	PricingCycle pulumi.StringPtrInput
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrInput
	// The status of the Basic Accelerator instance.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (BasicAcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicAcceleratorState)(nil)).Elem()
}

type basicAcceleratorArgs struct {
	// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
	AutoPay *bool `pulumi:"autoPay"`
	// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
	AutoRenew *bool `pulumi:"autoRenew"`
	// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
	AutoRenewDuration *int `pulumi:"autoRenewDuration"`
	// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
	AutoUseCoupon *string `pulumi:"autoUseCoupon"`
	// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
	BandwidthBillingType *string `pulumi:"bandwidthBillingType"`
	// The name of the Global Accelerator Basic Accelerator instance.
	BasicAcceleratorName *string `pulumi:"basicAcceleratorName"`
	// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
	CrossBorderStatus *bool `pulumi:"crossBorderStatus"`
	// The description of the Global Accelerator Basic Accelerator instance.
	Description *string `pulumi:"description"`
	// The subscription duration. Default value: `1`.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
	Duration *int `pulumi:"duration"`
	// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo *string `pulumi:"promotionOptionNo"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a BasicAccelerator resource.
type BasicAcceleratorArgs struct {
	// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
	AutoPay pulumi.BoolPtrInput
	// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
	AutoRenew pulumi.BoolPtrInput
	// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
	AutoRenewDuration pulumi.IntPtrInput
	// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
	AutoUseCoupon pulumi.StringPtrInput
	// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
	BandwidthBillingType pulumi.StringPtrInput
	// The name of the Global Accelerator Basic Accelerator instance.
	BasicAcceleratorName pulumi.StringPtrInput
	// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
	CrossBorderStatus pulumi.BoolPtrInput
	// The description of the Global Accelerator Basic Accelerator instance.
	Description pulumi.StringPtrInput
	// The subscription duration. Default value: `1`.
	// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
	// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
	Duration pulumi.IntPtrInput
	// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
	PricingCycle pulumi.StringPtrInput
	// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
	PromotionOptionNo pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (BasicAcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicAcceleratorArgs)(nil)).Elem()
}

type BasicAcceleratorInput interface {
	pulumi.Input

	ToBasicAcceleratorOutput() BasicAcceleratorOutput
	ToBasicAcceleratorOutputWithContext(ctx context.Context) BasicAcceleratorOutput
}

func (*BasicAccelerator) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicAccelerator)(nil)).Elem()
}

func (i *BasicAccelerator) ToBasicAcceleratorOutput() BasicAcceleratorOutput {
	return i.ToBasicAcceleratorOutputWithContext(context.Background())
}

func (i *BasicAccelerator) ToBasicAcceleratorOutputWithContext(ctx context.Context) BasicAcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAcceleratorOutput)
}

func (i *BasicAccelerator) ToOutput(ctx context.Context) pulumix.Output[*BasicAccelerator] {
	return pulumix.Output[*BasicAccelerator]{
		OutputState: i.ToBasicAcceleratorOutputWithContext(ctx).OutputState,
	}
}

// BasicAcceleratorArrayInput is an input type that accepts BasicAcceleratorArray and BasicAcceleratorArrayOutput values.
// You can construct a concrete instance of `BasicAcceleratorArrayInput` via:
//
//	BasicAcceleratorArray{ BasicAcceleratorArgs{...} }
type BasicAcceleratorArrayInput interface {
	pulumi.Input

	ToBasicAcceleratorArrayOutput() BasicAcceleratorArrayOutput
	ToBasicAcceleratorArrayOutputWithContext(context.Context) BasicAcceleratorArrayOutput
}

type BasicAcceleratorArray []BasicAcceleratorInput

func (BasicAcceleratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicAccelerator)(nil)).Elem()
}

func (i BasicAcceleratorArray) ToBasicAcceleratorArrayOutput() BasicAcceleratorArrayOutput {
	return i.ToBasicAcceleratorArrayOutputWithContext(context.Background())
}

func (i BasicAcceleratorArray) ToBasicAcceleratorArrayOutputWithContext(ctx context.Context) BasicAcceleratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAcceleratorArrayOutput)
}

func (i BasicAcceleratorArray) ToOutput(ctx context.Context) pulumix.Output[[]*BasicAccelerator] {
	return pulumix.Output[[]*BasicAccelerator]{
		OutputState: i.ToBasicAcceleratorArrayOutputWithContext(ctx).OutputState,
	}
}

// BasicAcceleratorMapInput is an input type that accepts BasicAcceleratorMap and BasicAcceleratorMapOutput values.
// You can construct a concrete instance of `BasicAcceleratorMapInput` via:
//
//	BasicAcceleratorMap{ "key": BasicAcceleratorArgs{...} }
type BasicAcceleratorMapInput interface {
	pulumi.Input

	ToBasicAcceleratorMapOutput() BasicAcceleratorMapOutput
	ToBasicAcceleratorMapOutputWithContext(context.Context) BasicAcceleratorMapOutput
}

type BasicAcceleratorMap map[string]BasicAcceleratorInput

func (BasicAcceleratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicAccelerator)(nil)).Elem()
}

func (i BasicAcceleratorMap) ToBasicAcceleratorMapOutput() BasicAcceleratorMapOutput {
	return i.ToBasicAcceleratorMapOutputWithContext(context.Background())
}

func (i BasicAcceleratorMap) ToBasicAcceleratorMapOutputWithContext(ctx context.Context) BasicAcceleratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicAcceleratorMapOutput)
}

func (i BasicAcceleratorMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BasicAccelerator] {
	return pulumix.Output[map[string]*BasicAccelerator]{
		OutputState: i.ToBasicAcceleratorMapOutputWithContext(ctx).OutputState,
	}
}

type BasicAcceleratorOutput struct{ *pulumi.OutputState }

func (BasicAcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicAccelerator)(nil)).Elem()
}

func (o BasicAcceleratorOutput) ToBasicAcceleratorOutput() BasicAcceleratorOutput {
	return o
}

func (o BasicAcceleratorOutput) ToBasicAcceleratorOutputWithContext(ctx context.Context) BasicAcceleratorOutput {
	return o
}

func (o BasicAcceleratorOutput) ToOutput(ctx context.Context) pulumix.Output[*BasicAccelerator] {
	return pulumix.Output[*BasicAccelerator]{
		OutputState: o.OutputState,
	}
}

// Specifies whether to enable automatic payment. Default value: `false`. Valid values:
func (o BasicAcceleratorOutput) AutoPay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.BoolPtrOutput { return v.AutoPay }).(pulumi.BoolPtrOutput)
}

// Specifies whether to enable auto-renewal for the GA Basic Accelerator instance. Default value: `false`. Valid values:
func (o BasicAcceleratorOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

// The auto-renewal period. Unit: months. Default value: `1`. Valid values: `1` to `12`. **NOTE:** This parameter is required only if `autoRenew` is set to `true`.
func (o BasicAcceleratorOutput) AutoRenewDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.IntPtrOutput { return v.AutoRenewDuration }).(pulumi.IntPtrOutput)
}

// Specifies whether to automatically pay bills by using coupons. Default value: `false`. **NOTE:** This parameter is required only if `autoPay` is set to `true`.
func (o BasicAcceleratorOutput) AutoUseCoupon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.AutoUseCoupon }).(pulumi.StringPtrOutput)
}

// The bandwidth billing method. Valid values: `BandwidthPackage`, `CDT`, `CDT95`.
func (o BasicAcceleratorOutput) BandwidthBillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.BandwidthBillingType }).(pulumi.StringPtrOutput)
}

// The name of the Global Accelerator Basic Accelerator instance.
func (o BasicAcceleratorOutput) BasicAcceleratorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.BasicAcceleratorName }).(pulumi.StringPtrOutput)
}

// Indicates whether cross-border acceleration is enabled. Default value: `false`. Valid values:
func (o BasicAcceleratorOutput) CrossBorderStatus() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.BoolPtrOutput { return v.CrossBorderStatus }).(pulumi.BoolPtrOutput)
}

// The description of the Global Accelerator Basic Accelerator instance.
func (o BasicAcceleratorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The subscription duration. Default value: `1`.
// * If the `pricingCycle` parameter is set to `Month`, the valid values for the `duration` parameter are `1` to `9`.
// * If the `pricingCycle` parameter is set to `Year`, the valid values for the `duration` parameter are `1` to `3`.
func (o BasicAcceleratorOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// The payment type. Default value: `Subscription`. Valid values: `PayAsYouGo`, `Subscription`.
func (o BasicAcceleratorOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The billing cycle. Default value: `Month`. Valid values: `Month`, `Year`.
func (o BasicAcceleratorOutput) PricingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.PricingCycle }).(pulumi.StringPtrOutput)
}

// The code of the coupon. **NOTE:** The `promotionOptionNo` takes effect only for accounts registered on the international site (alibabacloud.com).
func (o BasicAcceleratorOutput) PromotionOptionNo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringPtrOutput { return v.PromotionOptionNo }).(pulumi.StringPtrOutput)
}

// The status of the Basic Accelerator instance.
func (o BasicAcceleratorOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o BasicAcceleratorOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *BasicAccelerator) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type BasicAcceleratorArrayOutput struct{ *pulumi.OutputState }

func (BasicAcceleratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicAccelerator)(nil)).Elem()
}

func (o BasicAcceleratorArrayOutput) ToBasicAcceleratorArrayOutput() BasicAcceleratorArrayOutput {
	return o
}

func (o BasicAcceleratorArrayOutput) ToBasicAcceleratorArrayOutputWithContext(ctx context.Context) BasicAcceleratorArrayOutput {
	return o
}

func (o BasicAcceleratorArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BasicAccelerator] {
	return pulumix.Output[[]*BasicAccelerator]{
		OutputState: o.OutputState,
	}
}

func (o BasicAcceleratorArrayOutput) Index(i pulumi.IntInput) BasicAcceleratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicAccelerator {
		return vs[0].([]*BasicAccelerator)[vs[1].(int)]
	}).(BasicAcceleratorOutput)
}

type BasicAcceleratorMapOutput struct{ *pulumi.OutputState }

func (BasicAcceleratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicAccelerator)(nil)).Elem()
}

func (o BasicAcceleratorMapOutput) ToBasicAcceleratorMapOutput() BasicAcceleratorMapOutput {
	return o
}

func (o BasicAcceleratorMapOutput) ToBasicAcceleratorMapOutputWithContext(ctx context.Context) BasicAcceleratorMapOutput {
	return o
}

func (o BasicAcceleratorMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BasicAccelerator] {
	return pulumix.Output[map[string]*BasicAccelerator]{
		OutputState: o.OutputState,
	}
}

func (o BasicAcceleratorMapOutput) MapIndex(k pulumi.StringInput) BasicAcceleratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicAccelerator {
		return vs[0].(map[string]*BasicAccelerator)[vs[1].(string)]
	}).(BasicAcceleratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAcceleratorInput)(nil)).Elem(), &BasicAccelerator{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAcceleratorArrayInput)(nil)).Elem(), BasicAcceleratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicAcceleratorMapInput)(nil)).Elem(), BasicAcceleratorMap{})
	pulumi.RegisterOutputType(BasicAcceleratorOutput{})
	pulumi.RegisterOutputType(BasicAcceleratorArrayOutput{})
	pulumi.RegisterOutputType(BasicAcceleratorMapOutput{})
}
