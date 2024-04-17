// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sag Qos resource. Smart Access Gateway (SAG) supports quintuple-based QoS functions to differentiate traffic of different services and ensure high-priority traffic bandwidth.
//
// For information about Sag Qos and how to use it, see [What is Qos](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createqos).
//
// > **NOTE:** Available since v1.60.0.
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rocketmq.NewQos(ctx, "default", &rocketmq.QosArgs{
//				Name: pulumi.String("terraform-example"),
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
// The Sag Qos can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:rocketmq/qos:Qos example qos-abc123456
// ```
type Qos struct {
	pulumi.CustomResourceState

	// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewQos registers a new resource with the given unique name, arguments, and options.
func NewQos(ctx *pulumi.Context,
	name string, args *QosArgs, opts ...pulumi.ResourceOption) (*Qos, error) {
	if args == nil {
		args = &QosArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Qos
	err := ctx.RegisterResource("alicloud:rocketmq/qos:Qos", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQos gets an existing Qos resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQos(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosState, opts ...pulumi.ResourceOption) (*Qos, error) {
	var resource Qos
	err := ctx.ReadResource("alicloud:rocketmq/qos:Qos", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Qos resources.
type qosState struct {
	// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
	Name *string `pulumi:"name"`
}

type QosState struct {
	// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
	Name pulumi.StringPtrInput
}

func (QosState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosState)(nil)).Elem()
}

type qosArgs struct {
	// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Qos resource.
type QosArgs struct {
	// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
	Name pulumi.StringPtrInput
}

func (QosArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosArgs)(nil)).Elem()
}

type QosInput interface {
	pulumi.Input

	ToQosOutput() QosOutput
	ToQosOutputWithContext(ctx context.Context) QosOutput
}

func (*Qos) ElementType() reflect.Type {
	return reflect.TypeOf((**Qos)(nil)).Elem()
}

func (i *Qos) ToQosOutput() QosOutput {
	return i.ToQosOutputWithContext(context.Background())
}

func (i *Qos) ToQosOutputWithContext(ctx context.Context) QosOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosOutput)
}

// QosArrayInput is an input type that accepts QosArray and QosArrayOutput values.
// You can construct a concrete instance of `QosArrayInput` via:
//
//	QosArray{ QosArgs{...} }
type QosArrayInput interface {
	pulumi.Input

	ToQosArrayOutput() QosArrayOutput
	ToQosArrayOutputWithContext(context.Context) QosArrayOutput
}

type QosArray []QosInput

func (QosArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qos)(nil)).Elem()
}

func (i QosArray) ToQosArrayOutput() QosArrayOutput {
	return i.ToQosArrayOutputWithContext(context.Background())
}

func (i QosArray) ToQosArrayOutputWithContext(ctx context.Context) QosArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosArrayOutput)
}

// QosMapInput is an input type that accepts QosMap and QosMapOutput values.
// You can construct a concrete instance of `QosMapInput` via:
//
//	QosMap{ "key": QosArgs{...} }
type QosMapInput interface {
	pulumi.Input

	ToQosMapOutput() QosMapOutput
	ToQosMapOutputWithContext(context.Context) QosMapOutput
}

type QosMap map[string]QosInput

func (QosMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qos)(nil)).Elem()
}

func (i QosMap) ToQosMapOutput() QosMapOutput {
	return i.ToQosMapOutputWithContext(context.Background())
}

func (i QosMap) ToQosMapOutputWithContext(ctx context.Context) QosMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosMapOutput)
}

type QosOutput struct{ *pulumi.OutputState }

func (QosOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Qos)(nil)).Elem()
}

func (o QosOutput) ToQosOutput() QosOutput {
	return o
}

func (o QosOutput) ToQosOutputWithContext(ctx context.Context) QosOutput {
	return o
}

// The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.
func (o QosOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Qos) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type QosArrayOutput struct{ *pulumi.OutputState }

func (QosArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Qos)(nil)).Elem()
}

func (o QosArrayOutput) ToQosArrayOutput() QosArrayOutput {
	return o
}

func (o QosArrayOutput) ToQosArrayOutputWithContext(ctx context.Context) QosArrayOutput {
	return o
}

func (o QosArrayOutput) Index(i pulumi.IntInput) QosOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Qos {
		return vs[0].([]*Qos)[vs[1].(int)]
	}).(QosOutput)
}

type QosMapOutput struct{ *pulumi.OutputState }

func (QosMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Qos)(nil)).Elem()
}

func (o QosMapOutput) ToQosMapOutput() QosMapOutput {
	return o
}

func (o QosMapOutput) ToQosMapOutputWithContext(ctx context.Context) QosMapOutput {
	return o
}

func (o QosMapOutput) MapIndex(k pulumi.StringInput) QosOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Qos {
		return vs[0].(map[string]*Qos)[vs[1].(string)]
	}).(QosOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QosInput)(nil)).Elem(), &Qos{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosArrayInput)(nil)).Elem(), QosArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosMapInput)(nil)).Elem(), QosMap{})
	pulumi.RegisterOutputType(QosOutput{})
	pulumi.RegisterOutputType(QosArrayOutput{})
	pulumi.RegisterOutputType(QosMapOutput{})
}
