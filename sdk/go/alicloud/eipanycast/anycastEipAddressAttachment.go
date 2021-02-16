// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eipanycast

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Eipanycast Anycast Eip Address Attachment resource.
//
// For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://help.aliyun.com/document_detail/171857.html).
//
// > **NOTE:** Available in v1.113.0+.
//
// > **NOTE:** The following regions support currently while Slb instance support bound.
// [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/eipanycast"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAnycastEipAddress, err := eipanycast.NewAnycastEipAddress(ctx, "exampleAnycastEipAddress", &eipanycast.AnycastEipAddressArgs{
// 			ServiceLocation: pulumi.String("international"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eipanycast.NewAnycastEipAddressAttachment(ctx, "exampleAnycastEipAddressAttachment", &eipanycast.AnycastEipAddressAttachmentArgs{
// 			AnycastId:            exampleAnycastEipAddress.ID(),
// 			BindInstanceId:       pulumi.String("lb-j6chlcr8lffy7********"),
// 			BindInstanceRegionId: pulumi.String("cn-hongkong"),
// 			BindInstanceType:     pulumi.String("SlbInstance"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example `anycast_id`:`bind_instance_id`:`bind_instance_region_id`:`bind_instance_type`
// ```
type AnycastEipAddressAttachment struct {
	pulumi.CustomResourceState

	// The ID of Anycast EIP.
	AnycastId pulumi.StringOutput `pulumi:"anycastId"`
	// The ID of bound instance.
	BindInstanceId pulumi.StringOutput `pulumi:"bindInstanceId"`
	// The region ID of bound instance.
	BindInstanceRegionId pulumi.StringOutput `pulumi:"bindInstanceRegionId"`
	// The type of bound instance. Valid value: `SlbInstance`.
	BindInstanceType pulumi.StringOutput `pulumi:"bindInstanceType"`
	// The time of bound instance.
	BindTime pulumi.StringOutput `pulumi:"bindTime"`
}

// NewAnycastEipAddressAttachment registers a new resource with the given unique name, arguments, and options.
func NewAnycastEipAddressAttachment(ctx *pulumi.Context,
	name string, args *AnycastEipAddressAttachmentArgs, opts ...pulumi.ResourceOption) (*AnycastEipAddressAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnycastId == nil {
		return nil, errors.New("invalid value for required argument 'AnycastId'")
	}
	if args.BindInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceId'")
	}
	if args.BindInstanceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceRegionId'")
	}
	if args.BindInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceType'")
	}
	var resource AnycastEipAddressAttachment
	err := ctx.RegisterResource("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnycastEipAddressAttachment gets an existing AnycastEipAddressAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnycastEipAddressAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnycastEipAddressAttachmentState, opts ...pulumi.ResourceOption) (*AnycastEipAddressAttachment, error) {
	var resource AnycastEipAddressAttachment
	err := ctx.ReadResource("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnycastEipAddressAttachment resources.
type anycastEipAddressAttachmentState struct {
	// The ID of Anycast EIP.
	AnycastId *string `pulumi:"anycastId"`
	// The ID of bound instance.
	BindInstanceId *string `pulumi:"bindInstanceId"`
	// The region ID of bound instance.
	BindInstanceRegionId *string `pulumi:"bindInstanceRegionId"`
	// The type of bound instance. Valid value: `SlbInstance`.
	BindInstanceType *string `pulumi:"bindInstanceType"`
	// The time of bound instance.
	BindTime *string `pulumi:"bindTime"`
}

type AnycastEipAddressAttachmentState struct {
	// The ID of Anycast EIP.
	AnycastId pulumi.StringPtrInput
	// The ID of bound instance.
	BindInstanceId pulumi.StringPtrInput
	// The region ID of bound instance.
	BindInstanceRegionId pulumi.StringPtrInput
	// The type of bound instance. Valid value: `SlbInstance`.
	BindInstanceType pulumi.StringPtrInput
	// The time of bound instance.
	BindTime pulumi.StringPtrInput
}

func (AnycastEipAddressAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressAttachmentState)(nil)).Elem()
}

type anycastEipAddressAttachmentArgs struct {
	// The ID of Anycast EIP.
	AnycastId string `pulumi:"anycastId"`
	// The ID of bound instance.
	BindInstanceId string `pulumi:"bindInstanceId"`
	// The region ID of bound instance.
	BindInstanceRegionId string `pulumi:"bindInstanceRegionId"`
	// The type of bound instance. Valid value: `SlbInstance`.
	BindInstanceType string `pulumi:"bindInstanceType"`
}

// The set of arguments for constructing a AnycastEipAddressAttachment resource.
type AnycastEipAddressAttachmentArgs struct {
	// The ID of Anycast EIP.
	AnycastId pulumi.StringInput
	// The ID of bound instance.
	BindInstanceId pulumi.StringInput
	// The region ID of bound instance.
	BindInstanceRegionId pulumi.StringInput
	// The type of bound instance. Valid value: `SlbInstance`.
	BindInstanceType pulumi.StringInput
}

func (AnycastEipAddressAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressAttachmentArgs)(nil)).Elem()
}

type AnycastEipAddressAttachmentInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput
	ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput
}

func (*AnycastEipAddressAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*AnycastEipAddressAttachment)(nil))
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput {
	return i.ToAnycastEipAddressAttachmentOutputWithContext(context.Background())
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentOutput)
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentPtrOutput() AnycastEipAddressAttachmentPtrOutput {
	return i.ToAnycastEipAddressAttachmentPtrOutputWithContext(context.Background())
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentPtrOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentPtrOutput)
}

type AnycastEipAddressAttachmentPtrInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentPtrOutput() AnycastEipAddressAttachmentPtrOutput
	ToAnycastEipAddressAttachmentPtrOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentPtrOutput
}

type anycastEipAddressAttachmentPtrType AnycastEipAddressAttachmentArgs

func (*anycastEipAddressAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddressAttachment)(nil))
}

func (i *anycastEipAddressAttachmentPtrType) ToAnycastEipAddressAttachmentPtrOutput() AnycastEipAddressAttachmentPtrOutput {
	return i.ToAnycastEipAddressAttachmentPtrOutputWithContext(context.Background())
}

func (i *anycastEipAddressAttachmentPtrType) ToAnycastEipAddressAttachmentPtrOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentPtrOutput)
}

// AnycastEipAddressAttachmentArrayInput is an input type that accepts AnycastEipAddressAttachmentArray and AnycastEipAddressAttachmentArrayOutput values.
// You can construct a concrete instance of `AnycastEipAddressAttachmentArrayInput` via:
//
//          AnycastEipAddressAttachmentArray{ AnycastEipAddressAttachmentArgs{...} }
type AnycastEipAddressAttachmentArrayInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput
	ToAnycastEipAddressAttachmentArrayOutputWithContext(context.Context) AnycastEipAddressAttachmentArrayOutput
}

type AnycastEipAddressAttachmentArray []AnycastEipAddressAttachmentInput

func (AnycastEipAddressAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AnycastEipAddressAttachment)(nil))
}

func (i AnycastEipAddressAttachmentArray) ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput {
	return i.ToAnycastEipAddressAttachmentArrayOutputWithContext(context.Background())
}

func (i AnycastEipAddressAttachmentArray) ToAnycastEipAddressAttachmentArrayOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentArrayOutput)
}

// AnycastEipAddressAttachmentMapInput is an input type that accepts AnycastEipAddressAttachmentMap and AnycastEipAddressAttachmentMapOutput values.
// You can construct a concrete instance of `AnycastEipAddressAttachmentMapInput` via:
//
//          AnycastEipAddressAttachmentMap{ "key": AnycastEipAddressAttachmentArgs{...} }
type AnycastEipAddressAttachmentMapInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput
	ToAnycastEipAddressAttachmentMapOutputWithContext(context.Context) AnycastEipAddressAttachmentMapOutput
}

type AnycastEipAddressAttachmentMap map[string]AnycastEipAddressAttachmentInput

func (AnycastEipAddressAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AnycastEipAddressAttachment)(nil))
}

func (i AnycastEipAddressAttachmentMap) ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput {
	return i.ToAnycastEipAddressAttachmentMapOutputWithContext(context.Background())
}

func (i AnycastEipAddressAttachmentMap) ToAnycastEipAddressAttachmentMapOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentMapOutput)
}

type AnycastEipAddressAttachmentOutput struct {
	*pulumi.OutputState
}

func (AnycastEipAddressAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnycastEipAddressAttachment)(nil))
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput {
	return o
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput {
	return o
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentPtrOutput() AnycastEipAddressAttachmentPtrOutput {
	return o.ToAnycastEipAddressAttachmentPtrOutputWithContext(context.Background())
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentPtrOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentPtrOutput {
	return o.ApplyT(func(v AnycastEipAddressAttachment) *AnycastEipAddressAttachment {
		return &v
	}).(AnycastEipAddressAttachmentPtrOutput)
}

type AnycastEipAddressAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (AnycastEipAddressAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddressAttachment)(nil))
}

func (o AnycastEipAddressAttachmentPtrOutput) ToAnycastEipAddressAttachmentPtrOutput() AnycastEipAddressAttachmentPtrOutput {
	return o
}

func (o AnycastEipAddressAttachmentPtrOutput) ToAnycastEipAddressAttachmentPtrOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentPtrOutput {
	return o
}

type AnycastEipAddressAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnycastEipAddressAttachment)(nil))
}

func (o AnycastEipAddressAttachmentArrayOutput) ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput {
	return o
}

func (o AnycastEipAddressAttachmentArrayOutput) ToAnycastEipAddressAttachmentArrayOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentArrayOutput {
	return o
}

func (o AnycastEipAddressAttachmentArrayOutput) Index(i pulumi.IntInput) AnycastEipAddressAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AnycastEipAddressAttachment {
		return vs[0].([]AnycastEipAddressAttachment)[vs[1].(int)]
	}).(AnycastEipAddressAttachmentOutput)
}

type AnycastEipAddressAttachmentMapOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnycastEipAddressAttachment)(nil))
}

func (o AnycastEipAddressAttachmentMapOutput) ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput {
	return o
}

func (o AnycastEipAddressAttachmentMapOutput) ToAnycastEipAddressAttachmentMapOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentMapOutput {
	return o
}

func (o AnycastEipAddressAttachmentMapOutput) MapIndex(k pulumi.StringInput) AnycastEipAddressAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AnycastEipAddressAttachment {
		return vs[0].(map[string]AnycastEipAddressAttachment)[vs[1].(string)]
	}).(AnycastEipAddressAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentPtrOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentMapOutput{})
}
