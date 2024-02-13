// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.
//
// > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
//
// > **NOTE:** Available since v1.183.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ddos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := ddos.NewDdosBgpInstance(ctx, "instance", &ddos.DdosBgpInstanceArgs{
//				BaseBandwidth:   pulumi.Int(20),
//				Bandwidth:       -1,
//				IpCount:         pulumi.Int(100),
//				IpType:          pulumi.String("IPv4"),
//				NormalBandwidth: pulumi.Int(100),
//				Type:            pulumi.String("Enterprise"),
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
// Ddosbgp instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dns/ddosBgpInstance:DdosBgpInstance example ddosbgp-abc123456
// ```
//
// Deprecated: alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance
type DdosBgpInstance struct {
	pulumi.CustomResourceState

	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrOutput `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntOutput `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringOutput `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Normal defend bandwidth of the instance. The unit is Gbps.
	NormalBandwidth pulumi.IntOutput `pulumi:"normalBandwidth"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDdosBgpInstance registers a new resource with the given unique name, arguments, and options.
func NewDdosBgpInstance(ctx *pulumi.Context,
	name string, args *DdosBgpInstanceArgs, opts ...pulumi.ResourceOption) (*DdosBgpInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.IpCount == nil {
		return nil, errors.New("invalid value for required argument 'IpCount'")
	}
	if args.IpType == nil {
		return nil, errors.New("invalid value for required argument 'IpType'")
	}
	if args.NormalBandwidth == nil {
		return nil, errors.New("invalid value for required argument 'NormalBandwidth'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DdosBgpInstance
	err := ctx.RegisterResource("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosBgpInstance gets an existing DdosBgpInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosBgpInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosBgpInstanceState, opts ...pulumi.ResourceOption) (*DdosBgpInstance, error) {
	var resource DdosBgpInstance
	err := ctx.ReadResource("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosBgpInstance resources.
type ddosBgpInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth *int `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount *int `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType *string `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// Normal defend bandwidth of the instance. The unit is Gbps.
	NormalBandwidth *int `pulumi:"normalBandwidth"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
	Type *string `pulumi:"type"`
}

type DdosBgpInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntPtrInput
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrInput
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntPtrInput
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringPtrInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// Normal defend bandwidth of the instance. The unit is Gbps.
	NormalBandwidth pulumi.IntPtrInput
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
	Type pulumi.StringPtrInput
}

func (DdosBgpInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosBgpInstanceState)(nil)).Elem()
}

type ddosBgpInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth int `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth *int `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount int `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType string `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// Normal defend bandwidth of the instance. The unit is Gbps.
	NormalBandwidth int `pulumi:"normalBandwidth"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DdosBgpInstance resource.
type DdosBgpInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntInput
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrInput
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntInput
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// Normal defend bandwidth of the instance. The unit is Gbps.
	NormalBandwidth pulumi.IntInput
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
	Type pulumi.StringPtrInput
}

func (DdosBgpInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosBgpInstanceArgs)(nil)).Elem()
}

type DdosBgpInstanceInput interface {
	pulumi.Input

	ToDdosBgpInstanceOutput() DdosBgpInstanceOutput
	ToDdosBgpInstanceOutputWithContext(ctx context.Context) DdosBgpInstanceOutput
}

func (*DdosBgpInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosBgpInstance)(nil)).Elem()
}

func (i *DdosBgpInstance) ToDdosBgpInstanceOutput() DdosBgpInstanceOutput {
	return i.ToDdosBgpInstanceOutputWithContext(context.Background())
}

func (i *DdosBgpInstance) ToDdosBgpInstanceOutputWithContext(ctx context.Context) DdosBgpInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosBgpInstanceOutput)
}

// DdosBgpInstanceArrayInput is an input type that accepts DdosBgpInstanceArray and DdosBgpInstanceArrayOutput values.
// You can construct a concrete instance of `DdosBgpInstanceArrayInput` via:
//
//	DdosBgpInstanceArray{ DdosBgpInstanceArgs{...} }
type DdosBgpInstanceArrayInput interface {
	pulumi.Input

	ToDdosBgpInstanceArrayOutput() DdosBgpInstanceArrayOutput
	ToDdosBgpInstanceArrayOutputWithContext(context.Context) DdosBgpInstanceArrayOutput
}

type DdosBgpInstanceArray []DdosBgpInstanceInput

func (DdosBgpInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosBgpInstance)(nil)).Elem()
}

func (i DdosBgpInstanceArray) ToDdosBgpInstanceArrayOutput() DdosBgpInstanceArrayOutput {
	return i.ToDdosBgpInstanceArrayOutputWithContext(context.Background())
}

func (i DdosBgpInstanceArray) ToDdosBgpInstanceArrayOutputWithContext(ctx context.Context) DdosBgpInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosBgpInstanceArrayOutput)
}

// DdosBgpInstanceMapInput is an input type that accepts DdosBgpInstanceMap and DdosBgpInstanceMapOutput values.
// You can construct a concrete instance of `DdosBgpInstanceMapInput` via:
//
//	DdosBgpInstanceMap{ "key": DdosBgpInstanceArgs{...} }
type DdosBgpInstanceMapInput interface {
	pulumi.Input

	ToDdosBgpInstanceMapOutput() DdosBgpInstanceMapOutput
	ToDdosBgpInstanceMapOutputWithContext(context.Context) DdosBgpInstanceMapOutput
}

type DdosBgpInstanceMap map[string]DdosBgpInstanceInput

func (DdosBgpInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosBgpInstance)(nil)).Elem()
}

func (i DdosBgpInstanceMap) ToDdosBgpInstanceMapOutput() DdosBgpInstanceMapOutput {
	return i.ToDdosBgpInstanceMapOutputWithContext(context.Background())
}

func (i DdosBgpInstanceMap) ToDdosBgpInstanceMapOutputWithContext(ctx context.Context) DdosBgpInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosBgpInstanceMapOutput)
}

type DdosBgpInstanceOutput struct{ *pulumi.OutputState }

func (DdosBgpInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosBgpInstance)(nil)).Elem()
}

func (o DdosBgpInstanceOutput) ToDdosBgpInstanceOutput() DdosBgpInstanceOutput {
	return o
}

func (o DdosBgpInstanceOutput) ToDdosBgpInstanceOutputWithContext(ctx context.Context) DdosBgpInstanceOutput {
	return o
}

// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
func (o DdosBgpInstanceOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
func (o DdosBgpInstanceOutput) BaseBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.IntPtrOutput { return v.BaseBandwidth }).(pulumi.IntPtrOutput)
}

// IP count of the instance. Valid values: 100.
func (o DdosBgpInstanceOutput) IpCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.IntOutput { return v.IpCount }).(pulumi.IntOutput)
}

// IP version of the instance. Valid values: IPv4,IPv6.
func (o DdosBgpInstanceOutput) IpType() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.StringOutput { return v.IpType }).(pulumi.StringOutput)
}

// Name of the instance. This name can have a string of 1 to 63 characters.
func (o DdosBgpInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Normal defend bandwidth of the instance. The unit is Gbps.
func (o DdosBgpInstanceOutput) NormalBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.IntOutput { return v.NormalBandwidth }).(pulumi.IntOutput)
}

// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
func (o DdosBgpInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Type of the instance. Valid values: `Enterprise`, `Professional`. Default to `Enterprise`
func (o DdosBgpInstanceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosBgpInstance) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type DdosBgpInstanceArrayOutput struct{ *pulumi.OutputState }

func (DdosBgpInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosBgpInstance)(nil)).Elem()
}

func (o DdosBgpInstanceArrayOutput) ToDdosBgpInstanceArrayOutput() DdosBgpInstanceArrayOutput {
	return o
}

func (o DdosBgpInstanceArrayOutput) ToDdosBgpInstanceArrayOutputWithContext(ctx context.Context) DdosBgpInstanceArrayOutput {
	return o
}

func (o DdosBgpInstanceArrayOutput) Index(i pulumi.IntInput) DdosBgpInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdosBgpInstance {
		return vs[0].([]*DdosBgpInstance)[vs[1].(int)]
	}).(DdosBgpInstanceOutput)
}

type DdosBgpInstanceMapOutput struct{ *pulumi.OutputState }

func (DdosBgpInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosBgpInstance)(nil)).Elem()
}

func (o DdosBgpInstanceMapOutput) ToDdosBgpInstanceMapOutput() DdosBgpInstanceMapOutput {
	return o
}

func (o DdosBgpInstanceMapOutput) ToDdosBgpInstanceMapOutputWithContext(ctx context.Context) DdosBgpInstanceMapOutput {
	return o
}

func (o DdosBgpInstanceMapOutput) MapIndex(k pulumi.StringInput) DdosBgpInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdosBgpInstance {
		return vs[0].(map[string]*DdosBgpInstance)[vs[1].(string)]
	}).(DdosBgpInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdosBgpInstanceInput)(nil)).Elem(), &DdosBgpInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosBgpInstanceArrayInput)(nil)).Elem(), DdosBgpInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosBgpInstanceMapInput)(nil)).Elem(), DdosBgpInstanceMap{})
	pulumi.RegisterOutputType(DdosBgpInstanceOutput{})
	pulumi.RegisterOutputType(DdosBgpInstanceArrayOutput{})
	pulumi.RegisterOutputType(DdosBgpInstanceMapOutput{})
}
