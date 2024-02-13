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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			defaultCommonBandwithPackage, err := vpc.NewCommonBandwithPackage(ctx, "defaultCommonBandwithPackage", &vpc.CommonBandwithPackageArgs{
//				Bandwidth:          pulumi.String("3"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEipAddress, err := ecs.NewEipAddress(ctx, "defaultEipAddress", &ecs.EipAddressArgs{
//				Bandwidth:          pulumi.String("3"),
//				InternetChargeType: pulumi.String("PayByTraffic"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewCommonBandwithPackageAttachment(ctx, "defaultCommonBandwithPackageAttachment", &vpc.CommonBandwithPackageAttachmentArgs{
//				BandwidthPackageId:        defaultCommonBandwithPackage.ID(),
//				InstanceId:                defaultEipAddress.ID(),
//				BandwidthPackageBandwidth: pulumi.String("2"),
//				IpType:                    pulumi.String("EIP"),
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
// cbwp Common Bandwidth Package Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment example <bandwidth_package_id>:<instance_id>
// ```
type CommonBandwithPackageAttachment struct {
	pulumi.CustomResourceState

	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringOutput `pulumi:"bandwidthPackageBandwidth"`
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
	// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
	CancelCommonBandwidthPackageIpBandwidth pulumi.BoolPtrOutput `pulumi:"cancelCommonBandwidthPackageIpBandwidth"`
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
	IpType pulumi.StringPtrOutput `pulumi:"ipType"`
	// The status of the Internet Shared Bandwidth instance.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCommonBandwithPackageAttachment registers a new resource with the given unique name, arguments, and options.
func NewCommonBandwithPackageAttachment(ctx *pulumi.Context,
	name string, args *CommonBandwithPackageAttachmentArgs, opts ...pulumi.ResourceOption) (*CommonBandwithPackageAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BandwidthPackageId == nil {
		return nil, errors.New("invalid value for required argument 'BandwidthPackageId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CommonBandwithPackageAttachment
	err := ctx.RegisterResource("alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommonBandwithPackageAttachment gets an existing CommonBandwithPackageAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommonBandwithPackageAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommonBandwithPackageAttachmentState, opts ...pulumi.ResourceOption) (*CommonBandwithPackageAttachment, error) {
	var resource CommonBandwithPackageAttachment
	err := ctx.ReadResource("alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommonBandwithPackageAttachment resources.
type commonBandwithPackageAttachmentState struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth *string `pulumi:"bandwidthPackageBandwidth"`
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
	CancelCommonBandwidthPackageIpBandwidth *bool `pulumi:"cancelCommonBandwidthPackageIpBandwidth"`
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId *string `pulumi:"instanceId"`
	// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
	IpType *string `pulumi:"ipType"`
	// The status of the Internet Shared Bandwidth instance.
	Status *string `pulumi:"status"`
}

type CommonBandwithPackageAttachmentState struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringPtrInput
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringPtrInput
	// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
	CancelCommonBandwidthPackageIpBandwidth pulumi.BoolPtrInput
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringPtrInput
	// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
	IpType pulumi.StringPtrInput
	// The status of the Internet Shared Bandwidth instance.
	Status pulumi.StringPtrInput
}

func (CommonBandwithPackageAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageAttachmentState)(nil)).Elem()
}

type commonBandwithPackageAttachmentArgs struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth *string `pulumi:"bandwidthPackageBandwidth"`
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId string `pulumi:"bandwidthPackageId"`
	// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
	CancelCommonBandwidthPackageIpBandwidth *bool `pulumi:"cancelCommonBandwidthPackageIpBandwidth"`
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId string `pulumi:"instanceId"`
	// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
	IpType *string `pulumi:"ipType"`
}

// The set of arguments for constructing a CommonBandwithPackageAttachment resource.
type CommonBandwithPackageAttachmentArgs struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringPtrInput
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringInput
	// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
	CancelCommonBandwidthPackageIpBandwidth pulumi.BoolPtrInput
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringInput
	// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
	IpType pulumi.StringPtrInput
}

func (CommonBandwithPackageAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageAttachmentArgs)(nil)).Elem()
}

type CommonBandwithPackageAttachmentInput interface {
	pulumi.Input

	ToCommonBandwithPackageAttachmentOutput() CommonBandwithPackageAttachmentOutput
	ToCommonBandwithPackageAttachmentOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentOutput
}

func (*CommonBandwithPackageAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackageAttachment)(nil)).Elem()
}

func (i *CommonBandwithPackageAttachment) ToCommonBandwithPackageAttachmentOutput() CommonBandwithPackageAttachmentOutput {
	return i.ToCommonBandwithPackageAttachmentOutputWithContext(context.Background())
}

func (i *CommonBandwithPackageAttachment) ToCommonBandwithPackageAttachmentOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageAttachmentOutput)
}

// CommonBandwithPackageAttachmentArrayInput is an input type that accepts CommonBandwithPackageAttachmentArray and CommonBandwithPackageAttachmentArrayOutput values.
// You can construct a concrete instance of `CommonBandwithPackageAttachmentArrayInput` via:
//
//	CommonBandwithPackageAttachmentArray{ CommonBandwithPackageAttachmentArgs{...} }
type CommonBandwithPackageAttachmentArrayInput interface {
	pulumi.Input

	ToCommonBandwithPackageAttachmentArrayOutput() CommonBandwithPackageAttachmentArrayOutput
	ToCommonBandwithPackageAttachmentArrayOutputWithContext(context.Context) CommonBandwithPackageAttachmentArrayOutput
}

type CommonBandwithPackageAttachmentArray []CommonBandwithPackageAttachmentInput

func (CommonBandwithPackageAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackageAttachment)(nil)).Elem()
}

func (i CommonBandwithPackageAttachmentArray) ToCommonBandwithPackageAttachmentArrayOutput() CommonBandwithPackageAttachmentArrayOutput {
	return i.ToCommonBandwithPackageAttachmentArrayOutputWithContext(context.Background())
}

func (i CommonBandwithPackageAttachmentArray) ToCommonBandwithPackageAttachmentArrayOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageAttachmentArrayOutput)
}

// CommonBandwithPackageAttachmentMapInput is an input type that accepts CommonBandwithPackageAttachmentMap and CommonBandwithPackageAttachmentMapOutput values.
// You can construct a concrete instance of `CommonBandwithPackageAttachmentMapInput` via:
//
//	CommonBandwithPackageAttachmentMap{ "key": CommonBandwithPackageAttachmentArgs{...} }
type CommonBandwithPackageAttachmentMapInput interface {
	pulumi.Input

	ToCommonBandwithPackageAttachmentMapOutput() CommonBandwithPackageAttachmentMapOutput
	ToCommonBandwithPackageAttachmentMapOutputWithContext(context.Context) CommonBandwithPackageAttachmentMapOutput
}

type CommonBandwithPackageAttachmentMap map[string]CommonBandwithPackageAttachmentInput

func (CommonBandwithPackageAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackageAttachment)(nil)).Elem()
}

func (i CommonBandwithPackageAttachmentMap) ToCommonBandwithPackageAttachmentMapOutput() CommonBandwithPackageAttachmentMapOutput {
	return i.ToCommonBandwithPackageAttachmentMapOutputWithContext(context.Background())
}

func (i CommonBandwithPackageAttachmentMap) ToCommonBandwithPackageAttachmentMapOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageAttachmentMapOutput)
}

type CommonBandwithPackageAttachmentOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackageAttachment)(nil)).Elem()
}

func (o CommonBandwithPackageAttachmentOutput) ToCommonBandwithPackageAttachmentOutput() CommonBandwithPackageAttachmentOutput {
	return o
}

func (o CommonBandwithPackageAttachmentOutput) ToCommonBandwithPackageAttachmentOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentOutput {
	return o
}

// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
func (o CommonBandwithPackageAttachmentOutput) BandwidthPackageBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringOutput { return v.BandwidthPackageBandwidth }).(pulumi.StringOutput)
}

// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
func (o CommonBandwithPackageAttachmentOutput) BandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringOutput { return v.BandwidthPackageId }).(pulumi.StringOutput)
}

// Whether to cancel the maximum bandwidth configuration for the EIP. Default: false.
func (o CommonBandwithPackageAttachmentOutput) CancelCommonBandwidthPackageIpBandwidth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.BoolPtrOutput {
		return v.CancelCommonBandwidthPackageIpBandwidth
	}).(pulumi.BoolPtrOutput)
}

// The instanceId of the common bandwidth package attachment, the field can't be changed.
func (o CommonBandwithPackageAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// IP type. Set the value to **EIP**, which indicates that the EIP is added to the Internet shared bandwidth.
func (o CommonBandwithPackageAttachmentOutput) IpType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringPtrOutput { return v.IpType }).(pulumi.StringPtrOutput)
}

// The status of the Internet Shared Bandwidth instance.
func (o CommonBandwithPackageAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CommonBandwithPackageAttachmentArrayOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackageAttachment)(nil)).Elem()
}

func (o CommonBandwithPackageAttachmentArrayOutput) ToCommonBandwithPackageAttachmentArrayOutput() CommonBandwithPackageAttachmentArrayOutput {
	return o
}

func (o CommonBandwithPackageAttachmentArrayOutput) ToCommonBandwithPackageAttachmentArrayOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentArrayOutput {
	return o
}

func (o CommonBandwithPackageAttachmentArrayOutput) Index(i pulumi.IntInput) CommonBandwithPackageAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CommonBandwithPackageAttachment {
		return vs[0].([]*CommonBandwithPackageAttachment)[vs[1].(int)]
	}).(CommonBandwithPackageAttachmentOutput)
}

type CommonBandwithPackageAttachmentMapOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackageAttachment)(nil)).Elem()
}

func (o CommonBandwithPackageAttachmentMapOutput) ToCommonBandwithPackageAttachmentMapOutput() CommonBandwithPackageAttachmentMapOutput {
	return o
}

func (o CommonBandwithPackageAttachmentMapOutput) ToCommonBandwithPackageAttachmentMapOutputWithContext(ctx context.Context) CommonBandwithPackageAttachmentMapOutput {
	return o
}

func (o CommonBandwithPackageAttachmentMapOutput) MapIndex(k pulumi.StringInput) CommonBandwithPackageAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CommonBandwithPackageAttachment {
		return vs[0].(map[string]*CommonBandwithPackageAttachment)[vs[1].(string)]
	}).(CommonBandwithPackageAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageAttachmentInput)(nil)).Elem(), &CommonBandwithPackageAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageAttachmentArrayInput)(nil)).Elem(), CommonBandwithPackageAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageAttachmentMapInput)(nil)).Elem(), CommonBandwithPackageAttachmentMap{})
	pulumi.RegisterOutputType(CommonBandwithPackageAttachmentOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageAttachmentArrayOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageAttachmentMapOutput{})
}
