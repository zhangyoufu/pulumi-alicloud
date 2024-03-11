// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Api Gateway Instance resource.
//
// For information about Api Gateway Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/api-gateway/product-overview/dedicated-instances).
//
// > **NOTE:** Available since v1.218.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
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
//			_, err := apigateway.NewInstance(ctx, "default", &apigateway.InstanceArgs{
//				InstanceName: pulumi.String(name),
//				InstanceSpec: pulumi.String("api.s1.small"),
//				HttpsPolicy:  pulumi.String("HTTPS2_TLS1_0"),
//				ZoneId:       pulumi.String("cn-hangzhou-MAZ6"),
//				PaymentType:  pulumi.String("PayAsYouGo"),
//				UserVpcId:    pulumi.String("1709116870"),
//				InstanceType: pulumi.String("normal"),
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
// Api Gateway Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:apigateway/instance:Instance example <id>
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time of the instance package. Valid values:
	// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
	// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
	//
	// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// Does IPV6 Capability Support.
	EgressIpv6Enable pulumi.BoolPtrOutput `pulumi:"egressIpv6Enable"`
	// Https policy.
	HttpsPolicy pulumi.StringOutput `pulumi:"httpsPolicy"`
	// Instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Instance type.
	InstanceSpec pulumi.StringOutput `pulumi:"instanceSpec"`
	// Instance type-normal: traditional exclusive instance.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The payment type of the resource.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The subscription instance is of the subscription year or month type. The value range is as follows:
	// - **year**: year
	// - **month**: month
	// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
	PricingCycle pulumi.StringPtrOutput `pulumi:"pricingCycle"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Does ipv6 support.
	SupportIpv6 pulumi.BoolPtrOutput `pulumi:"supportIpv6"`
	// User's VpcID.
	UserVpcId pulumi.StringPtrOutput `pulumi:"userVpcId"`
	// Whether the slb of the Vpc supports.
	VpcSlbIntranetEnable pulumi.BoolPtrOutput `pulumi:"vpcSlbIntranetEnable"`
	// The zone where the instance is deployed.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpsPolicy == nil {
		return nil, errors.New("invalid value for required argument 'HttpsPolicy'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.InstanceSpec == nil {
		return nil, errors.New("invalid value for required argument 'InstanceSpec'")
	}
	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:apigateway/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:apigateway/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// The time of the instance package. Valid values:
	// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
	// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
	//
	// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
	Duration *int `pulumi:"duration"`
	// Does IPV6 Capability Support.
	EgressIpv6Enable *bool `pulumi:"egressIpv6Enable"`
	// Https policy.
	HttpsPolicy *string `pulumi:"httpsPolicy"`
	// Instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Instance type.
	InstanceSpec *string `pulumi:"instanceSpec"`
	// Instance type-normal: traditional exclusive instance.
	InstanceType *string `pulumi:"instanceType"`
	// The payment type of the resource.
	PaymentType *string `pulumi:"paymentType"`
	// The subscription instance is of the subscription year or month type. The value range is as follows:
	// - **year**: year
	// - **month**: month
	// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Does ipv6 support.
	SupportIpv6 *bool `pulumi:"supportIpv6"`
	// User's VpcID.
	UserVpcId *string `pulumi:"userVpcId"`
	// Whether the slb of the Vpc supports.
	VpcSlbIntranetEnable *bool `pulumi:"vpcSlbIntranetEnable"`
	// The zone where the instance is deployed.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// The time of the instance package. Valid values:
	// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
	// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
	//
	// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
	Duration pulumi.IntPtrInput
	// Does IPV6 Capability Support.
	EgressIpv6Enable pulumi.BoolPtrInput
	// Https policy.
	HttpsPolicy pulumi.StringPtrInput
	// Instance name.
	InstanceName pulumi.StringPtrInput
	// Instance type.
	InstanceSpec pulumi.StringPtrInput
	// Instance type-normal: traditional exclusive instance.
	InstanceType pulumi.StringPtrInput
	// The payment type of the resource.
	PaymentType pulumi.StringPtrInput
	// The subscription instance is of the subscription year or month type. The value range is as follows:
	// - **year**: year
	// - **month**: month
	// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
	PricingCycle pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Does ipv6 support.
	SupportIpv6 pulumi.BoolPtrInput
	// User's VpcID.
	UserVpcId pulumi.StringPtrInput
	// Whether the slb of the Vpc supports.
	VpcSlbIntranetEnable pulumi.BoolPtrInput
	// The zone where the instance is deployed.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The time of the instance package. Valid values:
	// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
	// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
	//
	// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
	Duration *int `pulumi:"duration"`
	// Does IPV6 Capability Support.
	EgressIpv6Enable *bool `pulumi:"egressIpv6Enable"`
	// Https policy.
	HttpsPolicy string `pulumi:"httpsPolicy"`
	// Instance name.
	InstanceName string `pulumi:"instanceName"`
	// Instance type.
	InstanceSpec string `pulumi:"instanceSpec"`
	// Instance type-normal: traditional exclusive instance.
	InstanceType *string `pulumi:"instanceType"`
	// The payment type of the resource.
	PaymentType string `pulumi:"paymentType"`
	// The subscription instance is of the subscription year or month type. The value range is as follows:
	// - **year**: year
	// - **month**: month
	// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
	PricingCycle *string `pulumi:"pricingCycle"`
	// Does ipv6 support.
	SupportIpv6 *bool `pulumi:"supportIpv6"`
	// User's VpcID.
	UserVpcId *string `pulumi:"userVpcId"`
	// Whether the slb of the Vpc supports.
	VpcSlbIntranetEnable *bool `pulumi:"vpcSlbIntranetEnable"`
	// The zone where the instance is deployed.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The time of the instance package. Valid values:
	// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
	// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
	//
	// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
	Duration pulumi.IntPtrInput
	// Does IPV6 Capability Support.
	EgressIpv6Enable pulumi.BoolPtrInput
	// Https policy.
	HttpsPolicy pulumi.StringInput
	// Instance name.
	InstanceName pulumi.StringInput
	// Instance type.
	InstanceSpec pulumi.StringInput
	// Instance type-normal: traditional exclusive instance.
	InstanceType pulumi.StringPtrInput
	// The payment type of the resource.
	PaymentType pulumi.StringInput
	// The subscription instance is of the subscription year or month type. The value range is as follows:
	// - **year**: year
	// - **month**: month
	// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
	PricingCycle pulumi.StringPtrInput
	// Does ipv6 support.
	SupportIpv6 pulumi.BoolPtrInput
	// User's VpcID.
	UserVpcId pulumi.StringPtrInput
	// Whether the slb of the Vpc supports.
	VpcSlbIntranetEnable pulumi.BoolPtrInput
	// The zone where the instance is deployed.
	ZoneId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Creation time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The time of the instance package. Valid values:
// - PricingCycle is **Month**, indicating monthly payment. The value range is **1** to **9**.
// - PricingCycle is **Year**, indicating annual payment. The value range is **1** to **3**.
//
// When the value of> ChargeType is **PrePaid**, this parameter is available and must be passed in.
func (o InstanceOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// Does IPV6 Capability Support.
func (o InstanceOutput) EgressIpv6Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.EgressIpv6Enable }).(pulumi.BoolPtrOutput)
}

// Https policy.
func (o InstanceOutput) HttpsPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.HttpsPolicy }).(pulumi.StringOutput)
}

// Instance name.
func (o InstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Instance type.
func (o InstanceOutput) InstanceSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceSpec }).(pulumi.StringOutput)
}

// Instance type-normal: traditional exclusive instance.
func (o InstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The payment type of the resource.
func (o InstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The subscription instance is of the subscription year or month type. The value range is as follows:
// - **year**: year
// - **month**: month
// > **NOTE:**  If the Payment type is PrePaid, this parameter is required.
func (o InstanceOutput) PricingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PricingCycle }).(pulumi.StringPtrOutput)
}

// The status of the resource.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Does ipv6 support.
func (o InstanceOutput) SupportIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.SupportIpv6 }).(pulumi.BoolPtrOutput)
}

// User's VpcID.
func (o InstanceOutput) UserVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.UserVpcId }).(pulumi.StringPtrOutput)
}

// Whether the slb of the Vpc supports.
func (o InstanceOutput) VpcSlbIntranetEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.VpcSlbIntranetEnable }).(pulumi.BoolPtrOutput)
}

// The zone where the instance is deployed.
func (o InstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
