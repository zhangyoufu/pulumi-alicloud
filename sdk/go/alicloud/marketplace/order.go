// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/marketplace"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := marketplace.NewOrder(ctx, "order", &marketplace.OrderArgs{
//				CouponId:       pulumi.String(""),
//				Duration:       pulumi.Int(1),
//				PackageVersion: pulumi.String("yuncode2713600001"),
//				PayType:        pulumi.String("prepay"),
//				PricingCycle:   pulumi.String("Month"),
//				ProductCode:    pulumi.String("cmapi033136"),
//				Quantity:       pulumi.Int(1),
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
// Market order can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:marketplace/order:Order order your-order-id
//
// ```
type Order struct {
	pulumi.CustomResourceState

	// Service providers customize additional components.
	Components pulumi.MapOutput `pulumi:"components"`
	// The coupon id of the market product.
	CouponId pulumi.StringPtrOutput `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion pulumi.StringOutput `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrOutput `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringOutput `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrOutput `pulumi:"quantity"`
}

// NewOrder registers a new resource with the given unique name, arguments, and options.
func NewOrder(ctx *pulumi.Context,
	name string, args *OrderArgs, opts ...pulumi.ResourceOption) (*Order, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PackageVersion == nil {
		return nil, errors.New("invalid value for required argument 'PackageVersion'")
	}
	if args.PricingCycle == nil {
		return nil, errors.New("invalid value for required argument 'PricingCycle'")
	}
	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	var resource Order
	err := ctx.RegisterResource("alicloud:marketplace/order:Order", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrder gets an existing Order resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderState, opts ...pulumi.ResourceOption) (*Order, error) {
	var resource Order
	err := ctx.ReadResource("alicloud:marketplace/order:Order", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Order resources.
type orderState struct {
	// Service providers customize additional components.
	Components map[string]interface{} `pulumi:"components"`
	// The coupon id of the market product.
	CouponId *string `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration *int `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion *string `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType *string `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode *string `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity *int `pulumi:"quantity"`
}

type OrderState struct {
	// Service providers customize additional components.
	Components pulumi.MapInput
	// The coupon id of the market product.
	CouponId pulumi.StringPtrInput
	// The number of purchase cycles.
	Duration pulumi.IntPtrInput
	// The package version of the market product.
	PackageVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrInput
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringPtrInput
	// The productCode of market place product.
	ProductCode pulumi.StringPtrInput
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrInput
}

func (OrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderState)(nil)).Elem()
}

type orderArgs struct {
	// Service providers customize additional components.
	Components map[string]interface{} `pulumi:"components"`
	// The coupon id of the market product.
	CouponId *string `pulumi:"couponId"`
	// The number of purchase cycles.
	Duration *int `pulumi:"duration"`
	// The package version of the market product.
	PackageVersion string `pulumi:"packageVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType *string `pulumi:"payType"`
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle string `pulumi:"pricingCycle"`
	// The productCode of market place product.
	ProductCode string `pulumi:"productCode"`
	// The quantity of the market product will be purchased.
	Quantity *int `pulumi:"quantity"`
}

// The set of arguments for constructing a Order resource.
type OrderArgs struct {
	// Service providers customize additional components.
	Components pulumi.MapInput
	// The coupon id of the market product.
	CouponId pulumi.StringPtrInput
	// The number of purchase cycles.
	Duration pulumi.IntPtrInput
	// The package version of the market product.
	PackageVersion pulumi.StringInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	PayType pulumi.StringPtrInput
	// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
	PricingCycle pulumi.StringInput
	// The productCode of market place product.
	ProductCode pulumi.StringInput
	// The quantity of the market product will be purchased.
	Quantity pulumi.IntPtrInput
}

func (OrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderArgs)(nil)).Elem()
}

type OrderInput interface {
	pulumi.Input

	ToOrderOutput() OrderOutput
	ToOrderOutputWithContext(ctx context.Context) OrderOutput
}

func (*Order) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (i *Order) ToOrderOutput() OrderOutput {
	return i.ToOrderOutputWithContext(context.Background())
}

func (i *Order) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderOutput)
}

// OrderArrayInput is an input type that accepts OrderArray and OrderArrayOutput values.
// You can construct a concrete instance of `OrderArrayInput` via:
//
//	OrderArray{ OrderArgs{...} }
type OrderArrayInput interface {
	pulumi.Input

	ToOrderArrayOutput() OrderArrayOutput
	ToOrderArrayOutputWithContext(context.Context) OrderArrayOutput
}

type OrderArray []OrderInput

func (OrderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Order)(nil)).Elem()
}

func (i OrderArray) ToOrderArrayOutput() OrderArrayOutput {
	return i.ToOrderArrayOutputWithContext(context.Background())
}

func (i OrderArray) ToOrderArrayOutputWithContext(ctx context.Context) OrderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderArrayOutput)
}

// OrderMapInput is an input type that accepts OrderMap and OrderMapOutput values.
// You can construct a concrete instance of `OrderMapInput` via:
//
//	OrderMap{ "key": OrderArgs{...} }
type OrderMapInput interface {
	pulumi.Input

	ToOrderMapOutput() OrderMapOutput
	ToOrderMapOutputWithContext(context.Context) OrderMapOutput
}

type OrderMap map[string]OrderInput

func (OrderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Order)(nil)).Elem()
}

func (i OrderMap) ToOrderMapOutput() OrderMapOutput {
	return i.ToOrderMapOutputWithContext(context.Background())
}

func (i OrderMap) ToOrderMapOutputWithContext(ctx context.Context) OrderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderMapOutput)
}

type OrderOutput struct{ *pulumi.OutputState }

func (OrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (o OrderOutput) ToOrderOutput() OrderOutput {
	return o
}

func (o OrderOutput) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return o
}

// Service providers customize additional components.
func (o OrderOutput) Components() pulumi.MapOutput {
	return o.ApplyT(func(v *Order) pulumi.MapOutput { return v.Components }).(pulumi.MapOutput)
}

// The coupon id of the market product.
func (o OrderOutput) CouponId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.StringPtrOutput { return v.CouponId }).(pulumi.StringPtrOutput)
}

// The number of purchase cycles.
func (o OrderOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// The package version of the market product.
func (o OrderOutput) PackageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.PackageVersion }).(pulumi.StringOutput)
}

// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
func (o OrderOutput) PayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.StringPtrOutput { return v.PayType }).(pulumi.StringPtrOutput)
}

// The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
func (o OrderOutput) PricingCycle() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.PricingCycle }).(pulumi.StringOutput)
}

// The productCode of market place product.
func (o OrderOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.ProductCode }).(pulumi.StringOutput)
}

// The quantity of the market product will be purchased.
func (o OrderOutput) Quantity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.IntPtrOutput { return v.Quantity }).(pulumi.IntPtrOutput)
}

type OrderArrayOutput struct{ *pulumi.OutputState }

func (OrderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Order)(nil)).Elem()
}

func (o OrderArrayOutput) ToOrderArrayOutput() OrderArrayOutput {
	return o
}

func (o OrderArrayOutput) ToOrderArrayOutputWithContext(ctx context.Context) OrderArrayOutput {
	return o
}

func (o OrderArrayOutput) Index(i pulumi.IntInput) OrderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Order {
		return vs[0].([]*Order)[vs[1].(int)]
	}).(OrderOutput)
}

type OrderMapOutput struct{ *pulumi.OutputState }

func (OrderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Order)(nil)).Elem()
}

func (o OrderMapOutput) ToOrderMapOutput() OrderMapOutput {
	return o
}

func (o OrderMapOutput) ToOrderMapOutputWithContext(ctx context.Context) OrderMapOutput {
	return o
}

func (o OrderMapOutput) MapIndex(k pulumi.StringInput) OrderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Order {
		return vs[0].(map[string]*Order)[vs[1].(string)]
	}).(OrderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrderInput)(nil)).Elem(), &Order{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderArrayInput)(nil)).Elem(), OrderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderMapInput)(nil)).Elem(), OrderMap{})
	pulumi.RegisterOutputType(OrderOutput{})
	pulumi.RegisterOutputType(OrderArrayOutput{})
	pulumi.RegisterOutputType(OrderMapOutput{})
}
