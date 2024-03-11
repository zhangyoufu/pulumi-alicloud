// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package message

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Message Notification Service Topic resource.
//
// For information about Message Notification Service Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/message-service/latest/createtopic).
//
// > **NOTE:** Available since v1.188.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/message"
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
//			_, err := message.NewServiceTopic(ctx, "default", &message.ServiceTopicArgs{
//				TopicName:      pulumi.String(name),
//				MaxMessageSize: pulumi.Int(12357),
//				LoggingEnabled: pulumi.Bool(true),
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
// Message Notification Service Topic can be imported using the id or topic_name, e.g.
//
// ```sh
// $ pulumi import alicloud:message/serviceTopic:ServiceTopic example <topic_name>
// ```
type ServiceTopic struct {
	pulumi.CustomResourceState

	// Specifies whether to enable the log management feature. Default value: false. Valid values:
	LoggingEnabled pulumi.BoolPtrOutput `pulumi:"loggingEnabled"`
	// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
	MaxMessageSize pulumi.IntOutput `pulumi:"maxMessageSize"`
	// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewServiceTopic registers a new resource with the given unique name, arguments, and options.
func NewServiceTopic(ctx *pulumi.Context,
	name string, args *ServiceTopicArgs, opts ...pulumi.ResourceOption) (*ServiceTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceTopic
	err := ctx.RegisterResource("alicloud:message/serviceTopic:ServiceTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTopic gets an existing ServiceTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTopicState, opts ...pulumi.ResourceOption) (*ServiceTopic, error) {
	var resource ServiceTopic
	err := ctx.ReadResource("alicloud:message/serviceTopic:ServiceTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTopic resources.
type serviceTopicState struct {
	// Specifies whether to enable the log management feature. Default value: false. Valid values:
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
	TopicName *string `pulumi:"topicName"`
}

type ServiceTopicState struct {
	// Specifies whether to enable the log management feature. Default value: false. Valid values:
	LoggingEnabled pulumi.BoolPtrInput
	// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
	MaxMessageSize pulumi.IntPtrInput
	// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
	TopicName pulumi.StringPtrInput
}

func (ServiceTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopicState)(nil)).Elem()
}

type serviceTopicArgs struct {
	// Specifies whether to enable the log management feature. Default value: false. Valid values:
	LoggingEnabled *bool `pulumi:"loggingEnabled"`
	// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
	MaxMessageSize *int `pulumi:"maxMessageSize"`
	// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a ServiceTopic resource.
type ServiceTopicArgs struct {
	// Specifies whether to enable the log management feature. Default value: false. Valid values:
	LoggingEnabled pulumi.BoolPtrInput
	// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
	MaxMessageSize pulumi.IntPtrInput
	// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
	TopicName pulumi.StringInput
}

func (ServiceTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopicArgs)(nil)).Elem()
}

type ServiceTopicInput interface {
	pulumi.Input

	ToServiceTopicOutput() ServiceTopicOutput
	ToServiceTopicOutputWithContext(ctx context.Context) ServiceTopicOutput
}

func (*ServiceTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTopic)(nil)).Elem()
}

func (i *ServiceTopic) ToServiceTopicOutput() ServiceTopicOutput {
	return i.ToServiceTopicOutputWithContext(context.Background())
}

func (i *ServiceTopic) ToServiceTopicOutputWithContext(ctx context.Context) ServiceTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopicOutput)
}

// ServiceTopicArrayInput is an input type that accepts ServiceTopicArray and ServiceTopicArrayOutput values.
// You can construct a concrete instance of `ServiceTopicArrayInput` via:
//
//	ServiceTopicArray{ ServiceTopicArgs{...} }
type ServiceTopicArrayInput interface {
	pulumi.Input

	ToServiceTopicArrayOutput() ServiceTopicArrayOutput
	ToServiceTopicArrayOutputWithContext(context.Context) ServiceTopicArrayOutput
}

type ServiceTopicArray []ServiceTopicInput

func (ServiceTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceTopic)(nil)).Elem()
}

func (i ServiceTopicArray) ToServiceTopicArrayOutput() ServiceTopicArrayOutput {
	return i.ToServiceTopicArrayOutputWithContext(context.Background())
}

func (i ServiceTopicArray) ToServiceTopicArrayOutputWithContext(ctx context.Context) ServiceTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopicArrayOutput)
}

// ServiceTopicMapInput is an input type that accepts ServiceTopicMap and ServiceTopicMapOutput values.
// You can construct a concrete instance of `ServiceTopicMapInput` via:
//
//	ServiceTopicMap{ "key": ServiceTopicArgs{...} }
type ServiceTopicMapInput interface {
	pulumi.Input

	ToServiceTopicMapOutput() ServiceTopicMapOutput
	ToServiceTopicMapOutputWithContext(context.Context) ServiceTopicMapOutput
}

type ServiceTopicMap map[string]ServiceTopicInput

func (ServiceTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceTopic)(nil)).Elem()
}

func (i ServiceTopicMap) ToServiceTopicMapOutput() ServiceTopicMapOutput {
	return i.ToServiceTopicMapOutputWithContext(context.Background())
}

func (i ServiceTopicMap) ToServiceTopicMapOutputWithContext(ctx context.Context) ServiceTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTopicMapOutput)
}

type ServiceTopicOutput struct{ *pulumi.OutputState }

func (ServiceTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTopic)(nil)).Elem()
}

func (o ServiceTopicOutput) ToServiceTopicOutput() ServiceTopicOutput {
	return o
}

func (o ServiceTopicOutput) ToServiceTopicOutputWithContext(ctx context.Context) ServiceTopicOutput {
	return o
}

// Specifies whether to enable the log management feature. Default value: false. Valid values:
func (o ServiceTopicOutput) LoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceTopic) pulumi.BoolPtrOutput { return v.LoggingEnabled }).(pulumi.BoolPtrOutput)
}

// The maximum size of a message body that can be sent to the topic. Unit: bytes. Valid values: 1024-65536. Default value: 65536.
func (o ServiceTopicOutput) MaxMessageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ServiceTopic) pulumi.IntOutput { return v.MaxMessageSize }).(pulumi.IntOutput)
}

// Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 255 characters.
func (o ServiceTopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTopic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type ServiceTopicArrayOutput struct{ *pulumi.OutputState }

func (ServiceTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceTopic)(nil)).Elem()
}

func (o ServiceTopicArrayOutput) ToServiceTopicArrayOutput() ServiceTopicArrayOutput {
	return o
}

func (o ServiceTopicArrayOutput) ToServiceTopicArrayOutputWithContext(ctx context.Context) ServiceTopicArrayOutput {
	return o
}

func (o ServiceTopicArrayOutput) Index(i pulumi.IntInput) ServiceTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceTopic {
		return vs[0].([]*ServiceTopic)[vs[1].(int)]
	}).(ServiceTopicOutput)
}

type ServiceTopicMapOutput struct{ *pulumi.OutputState }

func (ServiceTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceTopic)(nil)).Elem()
}

func (o ServiceTopicMapOutput) ToServiceTopicMapOutput() ServiceTopicMapOutput {
	return o
}

func (o ServiceTopicMapOutput) ToServiceTopicMapOutputWithContext(ctx context.Context) ServiceTopicMapOutput {
	return o
}

func (o ServiceTopicMapOutput) MapIndex(k pulumi.StringInput) ServiceTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceTopic {
		return vs[0].(map[string]*ServiceTopic)[vs[1].(string)]
	}).(ServiceTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTopicInput)(nil)).Elem(), &ServiceTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTopicArrayInput)(nil)).Elem(), ServiceTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTopicMapInput)(nil)).Elem(), ServiceTopicMap{})
	pulumi.RegisterOutputType(ServiceTopicOutput{})
	pulumi.RegisterOutputType(ServiceTopicArrayOutput{})
	pulumi.RegisterOutputType(ServiceTopicMapOutput{})
}
