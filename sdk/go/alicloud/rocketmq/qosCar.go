// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Sag qos car resource.
// You need to create a QoS car to set priorities, rate limits, and quintuple rules for different messages.
//
// For information about Sag Qos Car and how to use it, see [What is Qos Car](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/createqoscar).
//
// > **NOTE:** Available since v1.60.0.
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultQos, err := rocketmq.NewQos(ctx, "defaultQos", nil)
//			if err != nil {
//				return err
//			}
//			_, err = rocketmq.NewQosCar(ctx, "defaultQosCar", &rocketmq.QosCarArgs{
//				QosId:               defaultQos.ID(),
//				Description:         pulumi.String(name),
//				Priority:            pulumi.Int(1),
//				LimitType:           pulumi.String("Absolute"),
//				MinBandwidthAbs:     pulumi.Int(10),
//				MaxBandwidthAbs:     pulumi.Int(20),
//				MinBandwidthPercent: pulumi.Int(10),
//				MaxBandwidthPercent: pulumi.Int(20),
//				PercentSourceType:   pulumi.String("InternetUpBandwidth"),
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
// The Sag Qos Car can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rocketmq/qosCar:QosCar example qos-abc123456:qoscar-abc123456
//
// ```
type QosCar struct {
	pulumi.CustomResourceState

	// The description of the QoS speed limiting rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType pulumi.StringOutput `pulumi:"limitType"`
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs pulumi.IntPtrOutput `pulumi:"maxBandwidthAbs"`
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent pulumi.IntPtrOutput `pulumi:"maxBandwidthPercent"`
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs pulumi.IntPtrOutput `pulumi:"minBandwidthAbs"`
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent pulumi.IntPtrOutput `pulumi:"minBandwidthPercent"`
	// The name of the QoS speed limiting rule..
	Name pulumi.StringOutput `pulumi:"name"`
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType pulumi.StringPtrOutput `pulumi:"percentSourceType"`
	// The priority of the specified stream.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The instance ID of the QoS.
	QosId pulumi.StringOutput `pulumi:"qosId"`
}

// NewQosCar registers a new resource with the given unique name, arguments, and options.
func NewQosCar(ctx *pulumi.Context,
	name string, args *QosCarArgs, opts ...pulumi.ResourceOption) (*QosCar, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LimitType == nil {
		return nil, errors.New("invalid value for required argument 'LimitType'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.QosId == nil {
		return nil, errors.New("invalid value for required argument 'QosId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QosCar
	err := ctx.RegisterResource("alicloud:rocketmq/qosCar:QosCar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosCar gets an existing QosCar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosCar(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosCarState, opts ...pulumi.ResourceOption) (*QosCar, error) {
	var resource QosCar
	err := ctx.ReadResource("alicloud:rocketmq/qosCar:QosCar", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosCar resources.
type qosCarState struct {
	// The description of the QoS speed limiting rule.
	Description *string `pulumi:"description"`
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType *string `pulumi:"limitType"`
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs *int `pulumi:"maxBandwidthAbs"`
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent *int `pulumi:"maxBandwidthPercent"`
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs *int `pulumi:"minBandwidthAbs"`
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent *int `pulumi:"minBandwidthPercent"`
	// The name of the QoS speed limiting rule..
	Name *string `pulumi:"name"`
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType *string `pulumi:"percentSourceType"`
	// The priority of the specified stream.
	Priority *int `pulumi:"priority"`
	// The instance ID of the QoS.
	QosId *string `pulumi:"qosId"`
}

type QosCarState struct {
	// The description of the QoS speed limiting rule.
	Description pulumi.StringPtrInput
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType pulumi.StringPtrInput
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs pulumi.IntPtrInput
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent pulumi.IntPtrInput
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs pulumi.IntPtrInput
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent pulumi.IntPtrInput
	// The name of the QoS speed limiting rule..
	Name pulumi.StringPtrInput
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType pulumi.StringPtrInput
	// The priority of the specified stream.
	Priority pulumi.IntPtrInput
	// The instance ID of the QoS.
	QosId pulumi.StringPtrInput
}

func (QosCarState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosCarState)(nil)).Elem()
}

type qosCarArgs struct {
	// The description of the QoS speed limiting rule.
	Description *string `pulumi:"description"`
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType string `pulumi:"limitType"`
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs *int `pulumi:"maxBandwidthAbs"`
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent *int `pulumi:"maxBandwidthPercent"`
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs *int `pulumi:"minBandwidthAbs"`
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent *int `pulumi:"minBandwidthPercent"`
	// The name of the QoS speed limiting rule..
	Name *string `pulumi:"name"`
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType *string `pulumi:"percentSourceType"`
	// The priority of the specified stream.
	Priority int `pulumi:"priority"`
	// The instance ID of the QoS.
	QosId string `pulumi:"qosId"`
}

// The set of arguments for constructing a QosCar resource.
type QosCarArgs struct {
	// The description of the QoS speed limiting rule.
	Description pulumi.StringPtrInput
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType pulumi.StringInput
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs pulumi.IntPtrInput
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent pulumi.IntPtrInput
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs pulumi.IntPtrInput
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent pulumi.IntPtrInput
	// The name of the QoS speed limiting rule..
	Name pulumi.StringPtrInput
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType pulumi.StringPtrInput
	// The priority of the specified stream.
	Priority pulumi.IntInput
	// The instance ID of the QoS.
	QosId pulumi.StringInput
}

func (QosCarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosCarArgs)(nil)).Elem()
}

type QosCarInput interface {
	pulumi.Input

	ToQosCarOutput() QosCarOutput
	ToQosCarOutputWithContext(ctx context.Context) QosCarOutput
}

func (*QosCar) ElementType() reflect.Type {
	return reflect.TypeOf((**QosCar)(nil)).Elem()
}

func (i *QosCar) ToQosCarOutput() QosCarOutput {
	return i.ToQosCarOutputWithContext(context.Background())
}

func (i *QosCar) ToQosCarOutputWithContext(ctx context.Context) QosCarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosCarOutput)
}

func (i *QosCar) ToOutput(ctx context.Context) pulumix.Output[*QosCar] {
	return pulumix.Output[*QosCar]{
		OutputState: i.ToQosCarOutputWithContext(ctx).OutputState,
	}
}

// QosCarArrayInput is an input type that accepts QosCarArray and QosCarArrayOutput values.
// You can construct a concrete instance of `QosCarArrayInput` via:
//
//	QosCarArray{ QosCarArgs{...} }
type QosCarArrayInput interface {
	pulumi.Input

	ToQosCarArrayOutput() QosCarArrayOutput
	ToQosCarArrayOutputWithContext(context.Context) QosCarArrayOutput
}

type QosCarArray []QosCarInput

func (QosCarArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosCar)(nil)).Elem()
}

func (i QosCarArray) ToQosCarArrayOutput() QosCarArrayOutput {
	return i.ToQosCarArrayOutputWithContext(context.Background())
}

func (i QosCarArray) ToQosCarArrayOutputWithContext(ctx context.Context) QosCarArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosCarArrayOutput)
}

func (i QosCarArray) ToOutput(ctx context.Context) pulumix.Output[[]*QosCar] {
	return pulumix.Output[[]*QosCar]{
		OutputState: i.ToQosCarArrayOutputWithContext(ctx).OutputState,
	}
}

// QosCarMapInput is an input type that accepts QosCarMap and QosCarMapOutput values.
// You can construct a concrete instance of `QosCarMapInput` via:
//
//	QosCarMap{ "key": QosCarArgs{...} }
type QosCarMapInput interface {
	pulumi.Input

	ToQosCarMapOutput() QosCarMapOutput
	ToQosCarMapOutputWithContext(context.Context) QosCarMapOutput
}

type QosCarMap map[string]QosCarInput

func (QosCarMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosCar)(nil)).Elem()
}

func (i QosCarMap) ToQosCarMapOutput() QosCarMapOutput {
	return i.ToQosCarMapOutputWithContext(context.Background())
}

func (i QosCarMap) ToQosCarMapOutputWithContext(ctx context.Context) QosCarMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosCarMapOutput)
}

func (i QosCarMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QosCar] {
	return pulumix.Output[map[string]*QosCar]{
		OutputState: i.ToQosCarMapOutputWithContext(ctx).OutputState,
	}
}

type QosCarOutput struct{ *pulumi.OutputState }

func (QosCarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosCar)(nil)).Elem()
}

func (o QosCarOutput) ToQosCarOutput() QosCarOutput {
	return o
}

func (o QosCarOutput) ToQosCarOutputWithContext(ctx context.Context) QosCarOutput {
	return o
}

func (o QosCarOutput) ToOutput(ctx context.Context) pulumix.Output[*QosCar] {
	return pulumix.Output[*QosCar]{
		OutputState: o.OutputState,
	}
}

// The description of the QoS speed limiting rule.
func (o QosCarOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The speed limiting method. Valid values: Absolute, Percent.
func (o QosCarOutput) LimitType() pulumi.StringOutput {
	return o.ApplyT(func(v *QosCar) pulumi.StringOutput { return v.LimitType }).(pulumi.StringOutput)
}

// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
func (o QosCarOutput) MaxBandwidthAbs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.IntPtrOutput { return v.MaxBandwidthAbs }).(pulumi.IntPtrOutput)
}

// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
func (o QosCarOutput) MaxBandwidthPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.IntPtrOutput { return v.MaxBandwidthPercent }).(pulumi.IntPtrOutput)
}

// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
func (o QosCarOutput) MinBandwidthAbs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.IntPtrOutput { return v.MinBandwidthAbs }).(pulumi.IntPtrOutput)
}

// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
func (o QosCarOutput) MinBandwidthPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.IntPtrOutput { return v.MinBandwidthPercent }).(pulumi.IntPtrOutput)
}

// The name of the QoS speed limiting rule..
func (o QosCarOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QosCar) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
func (o QosCarOutput) PercentSourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosCar) pulumi.StringPtrOutput { return v.PercentSourceType }).(pulumi.StringPtrOutput)
}

// The priority of the specified stream.
func (o QosCarOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *QosCar) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The instance ID of the QoS.
func (o QosCarOutput) QosId() pulumi.StringOutput {
	return o.ApplyT(func(v *QosCar) pulumi.StringOutput { return v.QosId }).(pulumi.StringOutput)
}

type QosCarArrayOutput struct{ *pulumi.OutputState }

func (QosCarArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosCar)(nil)).Elem()
}

func (o QosCarArrayOutput) ToQosCarArrayOutput() QosCarArrayOutput {
	return o
}

func (o QosCarArrayOutput) ToQosCarArrayOutputWithContext(ctx context.Context) QosCarArrayOutput {
	return o
}

func (o QosCarArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QosCar] {
	return pulumix.Output[[]*QosCar]{
		OutputState: o.OutputState,
	}
}

func (o QosCarArrayOutput) Index(i pulumi.IntInput) QosCarOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QosCar {
		return vs[0].([]*QosCar)[vs[1].(int)]
	}).(QosCarOutput)
}

type QosCarMapOutput struct{ *pulumi.OutputState }

func (QosCarMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosCar)(nil)).Elem()
}

func (o QosCarMapOutput) ToQosCarMapOutput() QosCarMapOutput {
	return o
}

func (o QosCarMapOutput) ToQosCarMapOutputWithContext(ctx context.Context) QosCarMapOutput {
	return o
}

func (o QosCarMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QosCar] {
	return pulumix.Output[map[string]*QosCar]{
		OutputState: o.OutputState,
	}
}

func (o QosCarMapOutput) MapIndex(k pulumi.StringInput) QosCarOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QosCar {
		return vs[0].(map[string]*QosCar)[vs[1].(string)]
	}).(QosCarOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QosCarInput)(nil)).Elem(), &QosCar{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosCarArrayInput)(nil)).Elem(), QosCarArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosCarMapInput)(nil)).Elem(), QosCarMap{})
	pulumi.RegisterOutputType(QosCarOutput{})
	pulumi.RegisterOutputType(QosCarArrayOutput{})
	pulumi.RegisterOutputType(QosCarMapOutput{})
}
