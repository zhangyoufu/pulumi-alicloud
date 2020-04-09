// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rocketmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sag qos car resource.
// You need to create a QoS car to set priorities, rate limits, and quintuple rules for different messages.
//
// For information about Sag Qos Car and how to use it, see [What is Qos Car](https://www.alibabacloud.com/help/doc-detail/140065.htm).
//
// > **NOTE:** Available in 1.60.0+
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_qos_car.html.markdown.
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
	if args == nil || args.LimitType == nil {
		return nil, errors.New("missing required argument 'LimitType'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.QosId == nil {
		return nil, errors.New("missing required argument 'QosId'")
	}
	if args == nil {
		args = &QosCarArgs{}
	}
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
