// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DTS Consumer Channel resource.
//
// For information about DTS Consumer Channel and how to use it, see [What is Consumer Channel](https://www.alibabacloud.com/help/en/doc-detail/264593.htm).
//
// > **NOTE:** Available in v1.146.0+.
//
// ## Import
//
// DTS Consumer Channel can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dts/consumerChannel:ConsumerChannel example <dts_instance_id>:<consumer_group_id>
// ```
type ConsumerChannel struct {
	pulumi.CustomResourceState

	// The ID of the consumer group.
	ConsumerGroupId pulumi.StringOutput `pulumi:"consumerGroupId"`
	// The name of the consumer group.
	ConsumerGroupName pulumi.StringOutput `pulumi:"consumerGroupName"`
	// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	ConsumerGroupPassword pulumi.StringOutput `pulumi:"consumerGroupPassword"`
	// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	ConsumerGroupUserName pulumi.StringOutput `pulumi:"consumerGroupUserName"`
	// The ID of the subscription instance.
	DtsInstanceId pulumi.StringOutput `pulumi:"dtsInstanceId"`
}

// NewConsumerChannel registers a new resource with the given unique name, arguments, and options.
func NewConsumerChannel(ctx *pulumi.Context,
	name string, args *ConsumerChannelArgs, opts ...pulumi.ResourceOption) (*ConsumerChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.ConsumerGroupPassword == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupPassword'")
	}
	if args.ConsumerGroupUserName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupUserName'")
	}
	if args.DtsInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DtsInstanceId'")
	}
	var resource ConsumerChannel
	err := ctx.RegisterResource("alicloud:dts/consumerChannel:ConsumerChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerChannel gets an existing ConsumerChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerChannelState, opts ...pulumi.ResourceOption) (*ConsumerChannel, error) {
	var resource ConsumerChannel
	err := ctx.ReadResource("alicloud:dts/consumerChannel:ConsumerChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerChannel resources.
type consumerChannelState struct {
	// The ID of the consumer group.
	ConsumerGroupId *string `pulumi:"consumerGroupId"`
	// The name of the consumer group.
	ConsumerGroupName *string `pulumi:"consumerGroupName"`
	// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	ConsumerGroupPassword *string `pulumi:"consumerGroupPassword"`
	// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	ConsumerGroupUserName *string `pulumi:"consumerGroupUserName"`
	// The ID of the subscription instance.
	DtsInstanceId *string `pulumi:"dtsInstanceId"`
}

type ConsumerChannelState struct {
	// The ID of the consumer group.
	ConsumerGroupId pulumi.StringPtrInput
	// The name of the consumer group.
	ConsumerGroupName pulumi.StringPtrInput
	// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	ConsumerGroupPassword pulumi.StringPtrInput
	// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	ConsumerGroupUserName pulumi.StringPtrInput
	// The ID of the subscription instance.
	DtsInstanceId pulumi.StringPtrInput
}

func (ConsumerChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerChannelState)(nil)).Elem()
}

type consumerChannelArgs struct {
	// The name of the consumer group.
	ConsumerGroupName string `pulumi:"consumerGroupName"`
	// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	ConsumerGroupPassword string `pulumi:"consumerGroupPassword"`
	// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	ConsumerGroupUserName string `pulumi:"consumerGroupUserName"`
	// The ID of the subscription instance.
	DtsInstanceId string `pulumi:"dtsInstanceId"`
}

// The set of arguments for constructing a ConsumerChannel resource.
type ConsumerChannelArgs struct {
	// The name of the consumer group.
	ConsumerGroupName pulumi.StringInput
	// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	ConsumerGroupPassword pulumi.StringInput
	// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
	ConsumerGroupUserName pulumi.StringInput
	// The ID of the subscription instance.
	DtsInstanceId pulumi.StringInput
}

func (ConsumerChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerChannelArgs)(nil)).Elem()
}

type ConsumerChannelInput interface {
	pulumi.Input

	ToConsumerChannelOutput() ConsumerChannelOutput
	ToConsumerChannelOutputWithContext(ctx context.Context) ConsumerChannelOutput
}

func (*ConsumerChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerChannel)(nil)).Elem()
}

func (i *ConsumerChannel) ToConsumerChannelOutput() ConsumerChannelOutput {
	return i.ToConsumerChannelOutputWithContext(context.Background())
}

func (i *ConsumerChannel) ToConsumerChannelOutputWithContext(ctx context.Context) ConsumerChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerChannelOutput)
}

// ConsumerChannelArrayInput is an input type that accepts ConsumerChannelArray and ConsumerChannelArrayOutput values.
// You can construct a concrete instance of `ConsumerChannelArrayInput` via:
//
//          ConsumerChannelArray{ ConsumerChannelArgs{...} }
type ConsumerChannelArrayInput interface {
	pulumi.Input

	ToConsumerChannelArrayOutput() ConsumerChannelArrayOutput
	ToConsumerChannelArrayOutputWithContext(context.Context) ConsumerChannelArrayOutput
}

type ConsumerChannelArray []ConsumerChannelInput

func (ConsumerChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerChannel)(nil)).Elem()
}

func (i ConsumerChannelArray) ToConsumerChannelArrayOutput() ConsumerChannelArrayOutput {
	return i.ToConsumerChannelArrayOutputWithContext(context.Background())
}

func (i ConsumerChannelArray) ToConsumerChannelArrayOutputWithContext(ctx context.Context) ConsumerChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerChannelArrayOutput)
}

// ConsumerChannelMapInput is an input type that accepts ConsumerChannelMap and ConsumerChannelMapOutput values.
// You can construct a concrete instance of `ConsumerChannelMapInput` via:
//
//          ConsumerChannelMap{ "key": ConsumerChannelArgs{...} }
type ConsumerChannelMapInput interface {
	pulumi.Input

	ToConsumerChannelMapOutput() ConsumerChannelMapOutput
	ToConsumerChannelMapOutputWithContext(context.Context) ConsumerChannelMapOutput
}

type ConsumerChannelMap map[string]ConsumerChannelInput

func (ConsumerChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerChannel)(nil)).Elem()
}

func (i ConsumerChannelMap) ToConsumerChannelMapOutput() ConsumerChannelMapOutput {
	return i.ToConsumerChannelMapOutputWithContext(context.Background())
}

func (i ConsumerChannelMap) ToConsumerChannelMapOutputWithContext(ctx context.Context) ConsumerChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerChannelMapOutput)
}

type ConsumerChannelOutput struct{ *pulumi.OutputState }

func (ConsumerChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerChannel)(nil)).Elem()
}

func (o ConsumerChannelOutput) ToConsumerChannelOutput() ConsumerChannelOutput {
	return o
}

func (o ConsumerChannelOutput) ToConsumerChannelOutputWithContext(ctx context.Context) ConsumerChannelOutput {
	return o
}

// The ID of the consumer group.
func (o ConsumerChannelOutput) ConsumerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerChannel) pulumi.StringOutput { return v.ConsumerGroupId }).(pulumi.StringOutput)
}

// The name of the consumer group.
func (o ConsumerChannelOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerChannel) pulumi.StringOutput { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

// The password of the consumer group account. The length of the `consumerGroupPassword` is limited to `8` to `32` characters. It can contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
func (o ConsumerChannelOutput) ConsumerGroupPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerChannel) pulumi.StringOutput { return v.ConsumerGroupPassword }).(pulumi.StringOutput)
}

// The username of the consumer group. The length of the `consumerGroupUserName` is limited to `1` to `16` characters. It can contain one or more of the following characters: uppercase letters, lowercase letters, digits, and underscores (_).
func (o ConsumerChannelOutput) ConsumerGroupUserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerChannel) pulumi.StringOutput { return v.ConsumerGroupUserName }).(pulumi.StringOutput)
}

// The ID of the subscription instance.
func (o ConsumerChannelOutput) DtsInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerChannel) pulumi.StringOutput { return v.DtsInstanceId }).(pulumi.StringOutput)
}

type ConsumerChannelArrayOutput struct{ *pulumi.OutputState }

func (ConsumerChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerChannel)(nil)).Elem()
}

func (o ConsumerChannelArrayOutput) ToConsumerChannelArrayOutput() ConsumerChannelArrayOutput {
	return o
}

func (o ConsumerChannelArrayOutput) ToConsumerChannelArrayOutputWithContext(ctx context.Context) ConsumerChannelArrayOutput {
	return o
}

func (o ConsumerChannelArrayOutput) Index(i pulumi.IntInput) ConsumerChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumerChannel {
		return vs[0].([]*ConsumerChannel)[vs[1].(int)]
	}).(ConsumerChannelOutput)
}

type ConsumerChannelMapOutput struct{ *pulumi.OutputState }

func (ConsumerChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerChannel)(nil)).Elem()
}

func (o ConsumerChannelMapOutput) ToConsumerChannelMapOutput() ConsumerChannelMapOutput {
	return o
}

func (o ConsumerChannelMapOutput) ToConsumerChannelMapOutputWithContext(ctx context.Context) ConsumerChannelMapOutput {
	return o
}

func (o ConsumerChannelMapOutput) MapIndex(k pulumi.StringInput) ConsumerChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumerChannel {
		return vs[0].(map[string]*ConsumerChannel)[vs[1].(string)]
	}).(ConsumerChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerChannelInput)(nil)).Elem(), &ConsumerChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerChannelArrayInput)(nil)).Elem(), ConsumerChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerChannelMapInput)(nil)).Elem(), ConsumerChannelMap{})
	pulumi.RegisterOutputType(ConsumerChannelOutput{})
	pulumi.RegisterOutputType(ConsumerChannelArrayOutput{})
	pulumi.RegisterOutputType(ConsumerChannelMapOutput{})
}
