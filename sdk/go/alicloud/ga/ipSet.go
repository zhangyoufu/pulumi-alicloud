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

// Provides a Global Accelerator (GA) Ip Set resource.
//
// For information about Global Accelerator (GA) Ip Set and how to use it, see [What is Ip Set](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createipsets).
//
// > **NOTE:** Available since v1.113.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			region := "cn-hangzhou"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			defaultAccelerator, err := ga.NewAccelerator(ctx, "defaultAccelerator", &ga.AcceleratorArgs{
//				Duration:      pulumi.Int(1),
//				AutoUseCoupon: pulumi.Bool(true),
//				Spec:          pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "defaultBandwidthPackage", &ga.BandwidthPackageArgs{
//				Bandwidth:     pulumi.Int(100),
//				Type:          pulumi.String("Basic"),
//				BandwidthType: pulumi.String("Basic"),
//				PaymentType:   pulumi.String("PayAsYouGo"),
//				BillingType:   pulumi.String("PayBy95"),
//				Ratio:         pulumi.Int(30),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "defaultBandwidthPackageAttachment", &ga.BandwidthPackageAttachmentArgs{
//				AcceleratorId:      defaultAccelerator.ID(),
//				BandwidthPackageId: defaultBandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewIpSet(ctx, "example", &ga.IpSetArgs{
//				AccelerateRegionId: pulumi.String(region),
//				Bandwidth:          pulumi.Int(5),
//				AcceleratorId:      defaultBandwidthPackageAttachment.AcceleratorId,
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
// Ga Ip Set can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/ipSet:IpSet example <id>
// ```
type IpSet struct {
	pulumi.CustomResourceState

	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringOutput `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists pulumi.StringArrayOutput `pulumi:"ipAddressLists"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
	IspType pulumi.StringPtrOutput `pulumi:"ispType"`
	// The status of the acceleration region.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccelerateRegionId == nil {
		return nil, errors.New("invalid value for required argument 'AccelerateRegionId'")
	}
	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSet
	err := ctx.RegisterResource("alicloud:ga/ipSet:IpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSetState, opts ...pulumi.ResourceOption) (*IpSet, error) {
	var resource IpSet
	err := ctx.ReadResource("alicloud:ga/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// The ID of an acceleration region.
	AccelerateRegionId *string `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
	Bandwidth *int `pulumi:"bandwidth"`
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists []string `pulumi:"ipAddressLists"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
	IspType *string `pulumi:"ispType"`
	// The status of the acceleration region.
	Status *string `pulumi:"status"`
}

type IpSetState struct {
	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringPtrInput
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringPtrInput
	// The bandwidth allocated to the acceleration region.
	// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
	Bandwidth pulumi.IntPtrInput
	// The list of accelerated IP addresses in the acceleration region.
	IpAddressLists pulumi.StringArrayInput
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
	IpVersion pulumi.StringPtrInput
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
	IspType pulumi.StringPtrInput
	// The status of the acceleration region.
	Status pulumi.StringPtrInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// The ID of an acceleration region.
	AccelerateRegionId string `pulumi:"accelerateRegionId"`
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The bandwidth allocated to the acceleration region.
	// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
	Bandwidth *int `pulumi:"bandwidth"`
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
	IpVersion *string `pulumi:"ipVersion"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
	IspType *string `pulumi:"ispType"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// The ID of an acceleration region.
	AccelerateRegionId pulumi.StringInput
	// The ID of the Global Accelerator (GA) instance.
	AcceleratorId pulumi.StringInput
	// The bandwidth allocated to the acceleration region.
	// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
	Bandwidth pulumi.IntPtrInput
	// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
	IpVersion pulumi.StringPtrInput
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
	IspType pulumi.StringPtrInput
}

func (IpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetArgs)(nil)).Elem()
}

type IpSetInput interface {
	pulumi.Input

	ToIpSetOutput() IpSetOutput
	ToIpSetOutputWithContext(ctx context.Context) IpSetOutput
}

func (*IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (i *IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i *IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

// IpSetArrayInput is an input type that accepts IpSetArray and IpSetArrayOutput values.
// You can construct a concrete instance of `IpSetArrayInput` via:
//
//	IpSetArray{ IpSetArgs{...} }
type IpSetArrayInput interface {
	pulumi.Input

	ToIpSetArrayOutput() IpSetArrayOutput
	ToIpSetArrayOutputWithContext(context.Context) IpSetArrayOutput
}

type IpSetArray []IpSetInput

func (IpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSet)(nil)).Elem()
}

func (i IpSetArray) ToIpSetArrayOutput() IpSetArrayOutput {
	return i.ToIpSetArrayOutputWithContext(context.Background())
}

func (i IpSetArray) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetArrayOutput)
}

// IpSetMapInput is an input type that accepts IpSetMap and IpSetMapOutput values.
// You can construct a concrete instance of `IpSetMapInput` via:
//
//	IpSetMap{ "key": IpSetArgs{...} }
type IpSetMapInput interface {
	pulumi.Input

	ToIpSetMapOutput() IpSetMapOutput
	ToIpSetMapOutputWithContext(context.Context) IpSetMapOutput
}

type IpSetMap map[string]IpSetInput

func (IpSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSet)(nil)).Elem()
}

func (i IpSetMap) ToIpSetMapOutput() IpSetMapOutput {
	return i.ToIpSetMapOutputWithContext(context.Background())
}

func (i IpSetMap) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetMapOutput)
}

type IpSetOutput struct{ *pulumi.OutputState }

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

// The ID of an acceleration region.
func (o IpSetOutput) AccelerateRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.AccelerateRegionId }).(pulumi.StringOutput)
}

// The ID of the Global Accelerator (GA) instance.
func (o IpSetOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The bandwidth allocated to the acceleration region.
// > **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
func (o IpSetOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpSet) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// The list of accelerated IP addresses in the acceleration region.
func (o IpSetOutput) IpAddressLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringArrayOutput { return v.IpAddressLists }).(pulumi.StringArrayOutput)
}

// The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value: `IPv4`.
func (o IpSetOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.IpVersion }).(pulumi.StringOutput)
}

// The line type of the elastic IP address (EIP) in the acceleration region. Valid values: `BGP`, `BGP_PRO`.
func (o IpSetOutput) IspType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringPtrOutput { return v.IspType }).(pulumi.StringPtrOutput)
}

// The status of the acceleration region.
func (o IpSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type IpSetArrayOutput struct{ *pulumi.OutputState }

func (IpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSet)(nil)).Elem()
}

func (o IpSetArrayOutput) ToIpSetArrayOutput() IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) Index(i pulumi.IntInput) IpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpSet {
		return vs[0].([]*IpSet)[vs[1].(int)]
	}).(IpSetOutput)
}

type IpSetMapOutput struct{ *pulumi.OutputState }

func (IpSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSet)(nil)).Elem()
}

func (o IpSetMapOutput) ToIpSetMapOutput() IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) MapIndex(k pulumi.StringInput) IpSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpSet {
		return vs[0].(map[string]*IpSet)[vs[1].(string)]
	}).(IpSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetInput)(nil)).Elem(), &IpSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetArrayInput)(nil)).Elem(), IpSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetMapInput)(nil)).Elem(), IpSetMap{})
	pulumi.RegisterOutputType(IpSetOutput{})
	pulumi.RegisterOutputType(IpSetArrayOutput{})
	pulumi.RegisterOutputType(IpSetMapOutput{})
}
