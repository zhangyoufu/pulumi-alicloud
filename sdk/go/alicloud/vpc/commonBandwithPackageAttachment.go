// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooCommonBandwithPackage, err := vpc.NewCommonBandwithPackage(ctx, "fooCommonBandwithPackage", &vpc.CommonBandwithPackageArgs{
//				Bandwidth:   pulumi.String("2"),
//				Description: pulumi.String("test_common_bandwidth_package"),
//			})
//			if err != nil {
//				return err
//			}
//			fooEipAddress, err := ecs.NewEipAddress(ctx, "fooEipAddress", &ecs.EipAddressArgs{
//				Bandwidth:          pulumi.String("2"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewCommonBandwithPackageAttachment(ctx, "fooCommonBandwithPackageAttachment", &vpc.CommonBandwithPackageAttachmentArgs{
//				BandwidthPackageId: fooCommonBandwithPackage.ID(),
//				InstanceId:         fooEipAddress.ID(),
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
// The common bandwidth package attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment foo cbwp-abc123456:eip-abc123456
//
// ```
type CommonBandwithPackageAttachment struct {
	pulumi.CustomResourceState

	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringOutput `pulumi:"bandwidthPackageBandwidth"`
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
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
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId *string `pulumi:"instanceId"`
}

type CommonBandwithPackageAttachmentState struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringPtrInput
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringPtrInput
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringPtrInput
}

func (CommonBandwithPackageAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageAttachmentState)(nil)).Elem()
}

type commonBandwithPackageAttachmentArgs struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth *string `pulumi:"bandwidthPackageBandwidth"`
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId string `pulumi:"bandwidthPackageId"`
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a CommonBandwithPackageAttachment resource.
type CommonBandwithPackageAttachmentArgs struct {
	// The maximum bandwidth for the EIP. This value cannot be larger than the maximum bandwidth of the EIP bandwidth plan. Unit: Mbit/s.
	BandwidthPackageBandwidth pulumi.StringPtrInput
	// The bandwidthPackageId of the common bandwidth package attachment, the field can't be changed.
	BandwidthPackageId pulumi.StringInput
	// The instanceId of the common bandwidth package attachment, the field can't be changed.
	InstanceId pulumi.StringInput
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

// The instanceId of the common bandwidth package attachment, the field can't be changed.
func (o CommonBandwithPackageAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackageAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
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
