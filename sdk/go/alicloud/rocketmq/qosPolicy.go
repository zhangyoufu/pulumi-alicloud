// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sag qos policy resource.
// You need to create a QoS policy to set priorities, rate limits, and quintuple rules for different messages.
//
// For information about Sag Qos Policy and how to use it, see [What is Qos Policy](https://www.alibabacloud.com/help/doc-detail/140065.htm).
//
// > **NOTE:** Available in 1.60.0+
//
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/rocketmq"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultQos, err := rocketmq.NewQos(ctx, "defaultQos", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rocketmq.NewQosPolicy(ctx, "defaultQosPolicy", &rocketmq.QosPolicyArgs{
// 			QosId:           defaultQos.ID(),
// 			Description:     pulumi.String("tf-testSagQosPolicyDescription"),
// 			Priority:        pulumi.Int(1),
// 			IpProtocol:      pulumi.String("ALL"),
// 			SourceCidr:      pulumi.String("192.168.0.0/24"),
// 			SourcePortRange: pulumi.String("-1/-1"),
// 			DestCidr:        pulumi.String("10.10.0.0/24"),
// 			DestPortRange:   pulumi.String("-1/-1"),
// 			StartTime:       pulumi.String("2019-10-25T16:41:33+0800"),
// 			EndTime:         pulumi.String("2019-10-26T16:41:33+0800"),
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
// The Sag Qos Policy can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:rocketmq/qosPolicy:QosPolicy example qos-abc123456:qospy-abc123456
// ```
type QosPolicy struct {
	pulumi.CustomResourceState

	// The description of the QoS policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination CIDR block.
	DestCidr pulumi.StringOutput `pulumi:"destCidr"`
	// The destination port range.
	DestPortRange pulumi.StringOutput `pulumi:"destPortRange"`
	// The expiration time of the quintuple rule.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The transport layer protocol.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The name of the QoS policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The instance ID of the QoS policy to which the quintuple rule is created.
	QosId pulumi.StringOutput `pulumi:"qosId"`
	// The source CIDR block.
	SourceCidr pulumi.StringOutput `pulumi:"sourceCidr"`
	// The source port range of the transport layer.
	SourcePortRange pulumi.StringOutput `pulumi:"sourcePortRange"`
	// The time when the quintuple rule takes effect.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
}

// NewQosPolicy registers a new resource with the given unique name, arguments, and options.
func NewQosPolicy(ctx *pulumi.Context,
	name string, args *QosPolicyArgs, opts ...pulumi.ResourceOption) (*QosPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestCidr == nil {
		return nil, errors.New("invalid value for required argument 'DestCidr'")
	}
	if args.DestPortRange == nil {
		return nil, errors.New("invalid value for required argument 'DestPortRange'")
	}
	if args.IpProtocol == nil {
		return nil, errors.New("invalid value for required argument 'IpProtocol'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.QosId == nil {
		return nil, errors.New("invalid value for required argument 'QosId'")
	}
	if args.SourceCidr == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidr'")
	}
	if args.SourcePortRange == nil {
		return nil, errors.New("invalid value for required argument 'SourcePortRange'")
	}
	var resource QosPolicy
	err := ctx.RegisterResource("alicloud:rocketmq/qosPolicy:QosPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosPolicy gets an existing QosPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosPolicyState, opts ...pulumi.ResourceOption) (*QosPolicy, error) {
	var resource QosPolicy
	err := ctx.ReadResource("alicloud:rocketmq/qosPolicy:QosPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosPolicy resources.
type qosPolicyState struct {
	// The description of the QoS policy.
	Description *string `pulumi:"description"`
	// The destination CIDR block.
	DestCidr *string `pulumi:"destCidr"`
	// The destination port range.
	DestPortRange *string `pulumi:"destPortRange"`
	// The expiration time of the quintuple rule.
	EndTime *string `pulumi:"endTime"`
	// The transport layer protocol.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The name of the QoS policy.
	Name *string `pulumi:"name"`
	// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
	Priority *int `pulumi:"priority"`
	// The instance ID of the QoS policy to which the quintuple rule is created.
	QosId *string `pulumi:"qosId"`
	// The source CIDR block.
	SourceCidr *string `pulumi:"sourceCidr"`
	// The source port range of the transport layer.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The time when the quintuple rule takes effect.
	StartTime *string `pulumi:"startTime"`
}

type QosPolicyState struct {
	// The description of the QoS policy.
	Description pulumi.StringPtrInput
	// The destination CIDR block.
	DestCidr pulumi.StringPtrInput
	// The destination port range.
	DestPortRange pulumi.StringPtrInput
	// The expiration time of the quintuple rule.
	EndTime pulumi.StringPtrInput
	// The transport layer protocol.
	IpProtocol pulumi.StringPtrInput
	// The name of the QoS policy.
	Name pulumi.StringPtrInput
	// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
	Priority pulumi.IntPtrInput
	// The instance ID of the QoS policy to which the quintuple rule is created.
	QosId pulumi.StringPtrInput
	// The source CIDR block.
	SourceCidr pulumi.StringPtrInput
	// The source port range of the transport layer.
	SourcePortRange pulumi.StringPtrInput
	// The time when the quintuple rule takes effect.
	StartTime pulumi.StringPtrInput
}

func (QosPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosPolicyState)(nil)).Elem()
}

type qosPolicyArgs struct {
	// The description of the QoS policy.
	Description *string `pulumi:"description"`
	// The destination CIDR block.
	DestCidr string `pulumi:"destCidr"`
	// The destination port range.
	DestPortRange string `pulumi:"destPortRange"`
	// The expiration time of the quintuple rule.
	EndTime *string `pulumi:"endTime"`
	// The transport layer protocol.
	IpProtocol string `pulumi:"ipProtocol"`
	// The name of the QoS policy.
	Name *string `pulumi:"name"`
	// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
	Priority int `pulumi:"priority"`
	// The instance ID of the QoS policy to which the quintuple rule is created.
	QosId string `pulumi:"qosId"`
	// The source CIDR block.
	SourceCidr string `pulumi:"sourceCidr"`
	// The source port range of the transport layer.
	SourcePortRange string `pulumi:"sourcePortRange"`
	// The time when the quintuple rule takes effect.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a QosPolicy resource.
type QosPolicyArgs struct {
	// The description of the QoS policy.
	Description pulumi.StringPtrInput
	// The destination CIDR block.
	DestCidr pulumi.StringInput
	// The destination port range.
	DestPortRange pulumi.StringInput
	// The expiration time of the quintuple rule.
	EndTime pulumi.StringPtrInput
	// The transport layer protocol.
	IpProtocol pulumi.StringInput
	// The name of the QoS policy.
	Name pulumi.StringPtrInput
	// The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
	Priority pulumi.IntInput
	// The instance ID of the QoS policy to which the quintuple rule is created.
	QosId pulumi.StringInput
	// The source CIDR block.
	SourceCidr pulumi.StringInput
	// The source port range of the transport layer.
	SourcePortRange pulumi.StringInput
	// The time when the quintuple rule takes effect.
	StartTime pulumi.StringPtrInput
}

func (QosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosPolicyArgs)(nil)).Elem()
}

type QosPolicyInput interface {
	pulumi.Input

	ToQosPolicyOutput() QosPolicyOutput
	ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput
}

func (*QosPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*QosPolicy)(nil))
}

func (i *QosPolicy) ToQosPolicyOutput() QosPolicyOutput {
	return i.ToQosPolicyOutputWithContext(context.Background())
}

func (i *QosPolicy) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyOutput)
}

func (i *QosPolicy) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return i.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (i *QosPolicy) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyPtrOutput)
}

type QosPolicyPtrInput interface {
	pulumi.Input

	ToQosPolicyPtrOutput() QosPolicyPtrOutput
	ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput
}

type qosPolicyPtrType QosPolicyArgs

func (*qosPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil))
}

func (i *qosPolicyPtrType) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return i.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (i *qosPolicyPtrType) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyPtrOutput)
}

// QosPolicyArrayInput is an input type that accepts QosPolicyArray and QosPolicyArrayOutput values.
// You can construct a concrete instance of `QosPolicyArrayInput` via:
//
//          QosPolicyArray{ QosPolicyArgs{...} }
type QosPolicyArrayInput interface {
	pulumi.Input

	ToQosPolicyArrayOutput() QosPolicyArrayOutput
	ToQosPolicyArrayOutputWithContext(context.Context) QosPolicyArrayOutput
}

type QosPolicyArray []QosPolicyInput

func (QosPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*QosPolicy)(nil))
}

func (i QosPolicyArray) ToQosPolicyArrayOutput() QosPolicyArrayOutput {
	return i.ToQosPolicyArrayOutputWithContext(context.Background())
}

func (i QosPolicyArray) ToQosPolicyArrayOutputWithContext(ctx context.Context) QosPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyArrayOutput)
}

// QosPolicyMapInput is an input type that accepts QosPolicyMap and QosPolicyMapOutput values.
// You can construct a concrete instance of `QosPolicyMapInput` via:
//
//          QosPolicyMap{ "key": QosPolicyArgs{...} }
type QosPolicyMapInput interface {
	pulumi.Input

	ToQosPolicyMapOutput() QosPolicyMapOutput
	ToQosPolicyMapOutputWithContext(context.Context) QosPolicyMapOutput
}

type QosPolicyMap map[string]QosPolicyInput

func (QosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*QosPolicy)(nil))
}

func (i QosPolicyMap) ToQosPolicyMapOutput() QosPolicyMapOutput {
	return i.ToQosPolicyMapOutputWithContext(context.Background())
}

func (i QosPolicyMap) ToQosPolicyMapOutputWithContext(ctx context.Context) QosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyMapOutput)
}

type QosPolicyOutput struct {
	*pulumi.OutputState
}

func (QosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QosPolicy)(nil))
}

func (o QosPolicyOutput) ToQosPolicyOutput() QosPolicyOutput {
	return o
}

func (o QosPolicyOutput) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return o
}

func (o QosPolicyOutput) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return o.ToQosPolicyPtrOutputWithContext(context.Background())
}

func (o QosPolicyOutput) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return o.ApplyT(func(v QosPolicy) *QosPolicy {
		return &v
	}).(QosPolicyPtrOutput)
}

type QosPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (QosPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil))
}

func (o QosPolicyPtrOutput) ToQosPolicyPtrOutput() QosPolicyPtrOutput {
	return o
}

func (o QosPolicyPtrOutput) ToQosPolicyPtrOutputWithContext(ctx context.Context) QosPolicyPtrOutput {
	return o
}

type QosPolicyArrayOutput struct{ *pulumi.OutputState }

func (QosPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QosPolicy)(nil))
}

func (o QosPolicyArrayOutput) ToQosPolicyArrayOutput() QosPolicyArrayOutput {
	return o
}

func (o QosPolicyArrayOutput) ToQosPolicyArrayOutputWithContext(ctx context.Context) QosPolicyArrayOutput {
	return o
}

func (o QosPolicyArrayOutput) Index(i pulumi.IntInput) QosPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QosPolicy {
		return vs[0].([]QosPolicy)[vs[1].(int)]
	}).(QosPolicyOutput)
}

type QosPolicyMapOutput struct{ *pulumi.OutputState }

func (QosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QosPolicy)(nil))
}

func (o QosPolicyMapOutput) ToQosPolicyMapOutput() QosPolicyMapOutput {
	return o
}

func (o QosPolicyMapOutput) ToQosPolicyMapOutputWithContext(ctx context.Context) QosPolicyMapOutput {
	return o
}

func (o QosPolicyMapOutput) MapIndex(k pulumi.StringInput) QosPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QosPolicy {
		return vs[0].(map[string]QosPolicy)[vs[1].(string)]
	}).(QosPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(QosPolicyOutput{})
	pulumi.RegisterOutputType(QosPolicyPtrOutput{})
	pulumi.RegisterOutputType(QosPolicyArrayOutput{})
	pulumi.RegisterOutputType(QosPolicyMapOutput{})
}
