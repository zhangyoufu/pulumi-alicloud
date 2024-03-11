// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RabbitMQ (AMQP) Binding resource to bind tha exchange with another exchange or queue.
//
// For information about RabbitMQ (AMQP) Binding and how to use it, see [What is Binding](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createbinding).
//
// > **NOTE:** Available since v1.135.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInstance, err := amqp.NewInstance(ctx, "defaultInstance", &amqp.InstanceArgs{
//				InstanceType:  pulumi.String("enterprise"),
//				MaxTps:        pulumi.String("3000"),
//				QueueCapacity: pulumi.String("200"),
//				StorageSize:   pulumi.String("700"),
//				SupportEip:    pulumi.Bool(false),
//				MaxEipTps:     pulumi.String("128"),
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVirtualHost, err := amqp.NewVirtualHost(ctx, "defaultVirtualHost", &amqp.VirtualHostArgs{
//				InstanceId:      defaultInstance.ID(),
//				VirtualHostName: pulumi.String("tf-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultExchange, err := amqp.NewExchange(ctx, "defaultExchange", &amqp.ExchangeArgs{
//				AutoDeleteState: pulumi.Bool(false),
//				ExchangeName:    pulumi.String("tf-example"),
//				ExchangeType:    pulumi.String("HEADERS"),
//				InstanceId:      defaultInstance.ID(),
//				Internal:        pulumi.Bool(false),
//				VirtualHostName: defaultVirtualHost.VirtualHostName,
//			})
//			if err != nil {
//				return err
//			}
//			defaultQueue, err := amqp.NewQueue(ctx, "defaultQueue", &amqp.QueueArgs{
//				InstanceId:      defaultInstance.ID(),
//				QueueName:       pulumi.String("tf-example"),
//				VirtualHostName: defaultVirtualHost.VirtualHostName,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = amqp.NewBinding(ctx, "defaultBinding", &amqp.BindingArgs{
//				Argument:        pulumi.String("x-match:all"),
//				BindingKey:      defaultQueue.QueueName,
//				BindingType:     pulumi.String("QUEUE"),
//				DestinationName: pulumi.String("tf-example"),
//				InstanceId:      defaultInstance.ID(),
//				SourceExchange:  defaultExchange.ExchangeName,
//				VirtualHostName: defaultVirtualHost.VirtualHostName,
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
// RabbitMQ (AMQP) Binding can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:amqp/binding:Binding example <instance_id>:<virtual_host_name>:<source_exchange>:<destination_name>
// ```
type Binding struct {
	pulumi.CustomResourceState

	// X-match Attributes. Valid Values:
	// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
	// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
	//
	// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
	Argument pulumi.StringOutput `pulumi:"argument"`
	// The Binding Key.
	// * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   The binding key must be 1 to 255 characters in length.
	// * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
	//   The binding key must be 1 to 255 characters in length.
	BindingKey pulumi.StringOutput `pulumi:"bindingKey"`
	// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
	BindingType pulumi.StringOutput `pulumi:"bindingType"`
	// The Target Queue Or Exchange of the Name.
	DestinationName pulumi.StringOutput `pulumi:"destinationName"`
	// Instance Id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The Source Exchange Name.
	SourceExchange pulumi.StringOutput `pulumi:"sourceExchange"`
	// Virtualhost Name.
	VirtualHostName pulumi.StringOutput `pulumi:"virtualHostName"`
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOption) (*Binding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BindingKey == nil {
		return nil, errors.New("invalid value for required argument 'BindingKey'")
	}
	if args.BindingType == nil {
		return nil, errors.New("invalid value for required argument 'BindingType'")
	}
	if args.DestinationName == nil {
		return nil, errors.New("invalid value for required argument 'DestinationName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.SourceExchange == nil {
		return nil, errors.New("invalid value for required argument 'SourceExchange'")
	}
	if args.VirtualHostName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHostName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Binding
	err := ctx.RegisterResource("alicloud:amqp/binding:Binding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindingState, opts ...pulumi.ResourceOption) (*Binding, error) {
	var resource Binding
	err := ctx.ReadResource("alicloud:amqp/binding:Binding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Binding resources.
type bindingState struct {
	// X-match Attributes. Valid Values:
	// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
	// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
	//
	// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
	Argument *string `pulumi:"argument"`
	// The Binding Key.
	// * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   The binding key must be 1 to 255 characters in length.
	// * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
	//   The binding key must be 1 to 255 characters in length.
	BindingKey *string `pulumi:"bindingKey"`
	// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
	BindingType *string `pulumi:"bindingType"`
	// The Target Queue Or Exchange of the Name.
	DestinationName *string `pulumi:"destinationName"`
	// Instance Id.
	InstanceId *string `pulumi:"instanceId"`
	// The Source Exchange Name.
	SourceExchange *string `pulumi:"sourceExchange"`
	// Virtualhost Name.
	VirtualHostName *string `pulumi:"virtualHostName"`
}

type BindingState struct {
	// X-match Attributes. Valid Values:
	// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
	// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
	//
	// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
	Argument pulumi.StringPtrInput
	// The Binding Key.
	// * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   The binding key must be 1 to 255 characters in length.
	// * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
	//   The binding key must be 1 to 255 characters in length.
	BindingKey pulumi.StringPtrInput
	// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
	BindingType pulumi.StringPtrInput
	// The Target Queue Or Exchange of the Name.
	DestinationName pulumi.StringPtrInput
	// Instance Id.
	InstanceId pulumi.StringPtrInput
	// The Source Exchange Name.
	SourceExchange pulumi.StringPtrInput
	// Virtualhost Name.
	VirtualHostName pulumi.StringPtrInput
}

func (BindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingState)(nil)).Elem()
}

type bindingArgs struct {
	// X-match Attributes. Valid Values:
	// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
	// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
	//
	// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
	Argument *string `pulumi:"argument"`
	// The Binding Key.
	// * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   The binding key must be 1 to 255 characters in length.
	// * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
	//   The binding key must be 1 to 255 characters in length.
	BindingKey string `pulumi:"bindingKey"`
	// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
	BindingType string `pulumi:"bindingType"`
	// The Target Queue Or Exchange of the Name.
	DestinationName string `pulumi:"destinationName"`
	// Instance Id.
	InstanceId string `pulumi:"instanceId"`
	// The Source Exchange Name.
	SourceExchange string `pulumi:"sourceExchange"`
	// Virtualhost Name.
	VirtualHostName string `pulumi:"virtualHostName"`
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// X-match Attributes. Valid Values:
	// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
	// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
	//
	// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
	Argument pulumi.StringPtrInput
	// The Binding Key.
	// * For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   The binding key must be 1 to 255 characters in length.
	// * For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	//   If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
	//   The binding key must be 1 to 255 characters in length.
	BindingKey pulumi.StringInput
	// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
	BindingType pulumi.StringInput
	// The Target Queue Or Exchange of the Name.
	DestinationName pulumi.StringInput
	// Instance Id.
	InstanceId pulumi.StringInput
	// The Source Exchange Name.
	SourceExchange pulumi.StringInput
	// Virtualhost Name.
	VirtualHostName pulumi.StringInput
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindingArgs)(nil)).Elem()
}

type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(ctx context.Context) BindingOutput
}

func (*Binding) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (i *Binding) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i *Binding) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

// BindingArrayInput is an input type that accepts BindingArray and BindingArrayOutput values.
// You can construct a concrete instance of `BindingArrayInput` via:
//
//	BindingArray{ BindingArgs{...} }
type BindingArrayInput interface {
	pulumi.Input

	ToBindingArrayOutput() BindingArrayOutput
	ToBindingArrayOutputWithContext(context.Context) BindingArrayOutput
}

type BindingArray []BindingInput

func (BindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (i BindingArray) ToBindingArrayOutput() BindingArrayOutput {
	return i.ToBindingArrayOutputWithContext(context.Background())
}

func (i BindingArray) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingArrayOutput)
}

// BindingMapInput is an input type that accepts BindingMap and BindingMapOutput values.
// You can construct a concrete instance of `BindingMapInput` via:
//
//	BindingMap{ "key": BindingArgs{...} }
type BindingMapInput interface {
	pulumi.Input

	ToBindingMapOutput() BindingMapOutput
	ToBindingMapOutputWithContext(context.Context) BindingMapOutput
}

type BindingMap map[string]BindingInput

func (BindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (i BindingMap) ToBindingMapOutput() BindingMapOutput {
	return i.ToBindingMapOutputWithContext(context.Background())
}

func (i BindingMap) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingMapOutput)
}

type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

// X-match Attributes. Valid Values:
// * "x-match:all": Default Value, All the Message Header of Key-Value Pairs Stored in the Must Match.
// * "x-match:any": at Least One Pair of the Message Header of Key-Value Pairs Stored in the Must Match.
//
// **NOTE:** This Parameter Applies Only to Headers Exchange Other Types of Exchange Is Invalid. Other Types of Exchange Here Can Either Be an Arbitrary Value.
func (o BindingOutput) Argument() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.Argument }).(pulumi.StringOutput)
}

// The Binding Key.
//   - For a non-topic source exchange: The binding key can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
//     The binding key must be 1 to 255 characters in length.
//   - For a topic source exchange: The binding key can contain letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
//     If the binding key contains a number sign (#), the binding key must start with a number sign (#) followed by a period (.) or end with a number sign (#) that follows a period (.).
//     The binding key must be 1 to 255 characters in length.
func (o BindingOutput) BindingKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.BindingKey }).(pulumi.StringOutput)
}

// The Target Binding Types. Valid values: `EXCHANGE`, `QUEUE`.
func (o BindingOutput) BindingType() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.BindingType }).(pulumi.StringOutput)
}

// The Target Queue Or Exchange of the Name.
func (o BindingOutput) DestinationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.DestinationName }).(pulumi.StringOutput)
}

// Instance Id.
func (o BindingOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The Source Exchange Name.
func (o BindingOutput) SourceExchange() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.SourceExchange }).(pulumi.StringOutput)
}

// Virtualhost Name.
func (o BindingOutput) VirtualHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Binding) pulumi.StringOutput { return v.VirtualHostName }).(pulumi.StringOutput)
}

type BindingArrayOutput struct{ *pulumi.OutputState }

func (BindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Binding)(nil)).Elem()
}

func (o BindingArrayOutput) ToBindingArrayOutput() BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) Index(i pulumi.IntInput) BindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].([]*Binding)[vs[1].(int)]
	}).(BindingOutput)
}

type BindingMapOutput struct{ *pulumi.OutputState }

func (BindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Binding)(nil)).Elem()
}

func (o BindingMapOutput) ToBindingMapOutput() BindingMapOutput {
	return o
}

func (o BindingMapOutput) ToBindingMapOutputWithContext(ctx context.Context) BindingMapOutput {
	return o
}

func (o BindingMapOutput) MapIndex(k pulumi.StringInput) BindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Binding {
		return vs[0].(map[string]*Binding)[vs[1].(string)]
	}).(BindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingInput)(nil)).Elem(), &Binding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingArrayInput)(nil)).Elem(), BindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingMapInput)(nil)).Elem(), BindingMap{})
	pulumi.RegisterOutputType(BindingOutput{})
	pulumi.RegisterOutputType(BindingArrayOutput{})
	pulumi.RegisterOutputType(BindingMapOutput{})
}
