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

// Provides a RabbitMQ (AMQP) Virtual Host resource.
//
// For information about RabbitMQ (AMQP) Virtual Host and how to use it, see [What is Virtual Host](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/createvirtualhost).
//
// > **NOTE:** Available since v1.126.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInstance, err := amqp.NewInstance(ctx, "defaultInstance", &amqp.InstanceArgs{
//				InstanceType:  pulumi.String("professional"),
//				MaxTps:        pulumi.String("1000"),
//				QueueCapacity: pulumi.String("50"),
//				SupportEip:    pulumi.Bool(true),
//				MaxEipTps:     pulumi.String("128"),
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = amqp.NewVirtualHost(ctx, "defaultVirtualHost", &amqp.VirtualHostArgs{
//				InstanceId:      defaultInstance.ID(),
//				VirtualHostName: pulumi.String("tf-example"),
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
// RabbitMQ (AMQP) Virtual Host can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:amqp/virtualHost:VirtualHost example <instance_id>:<virtual_host_name>
// ```
type VirtualHost struct {
	pulumi.CustomResourceState

	// InstanceId.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// VirtualHostName.
	VirtualHostName pulumi.StringOutput `pulumi:"virtualHostName"`
}

// NewVirtualHost registers a new resource with the given unique name, arguments, and options.
func NewVirtualHost(ctx *pulumi.Context,
	name string, args *VirtualHostArgs, opts ...pulumi.ResourceOption) (*VirtualHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VirtualHostName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHostName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualHost
	err := ctx.RegisterResource("alicloud:amqp/virtualHost:VirtualHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHost gets an existing VirtualHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHostState, opts ...pulumi.ResourceOption) (*VirtualHost, error) {
	var resource VirtualHost
	err := ctx.ReadResource("alicloud:amqp/virtualHost:VirtualHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHost resources.
type virtualHostState struct {
	// InstanceId.
	InstanceId *string `pulumi:"instanceId"`
	// VirtualHostName.
	VirtualHostName *string `pulumi:"virtualHostName"`
}

type VirtualHostState struct {
	// InstanceId.
	InstanceId pulumi.StringPtrInput
	// VirtualHostName.
	VirtualHostName pulumi.StringPtrInput
}

func (VirtualHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHostState)(nil)).Elem()
}

type virtualHostArgs struct {
	// InstanceId.
	InstanceId string `pulumi:"instanceId"`
	// VirtualHostName.
	VirtualHostName string `pulumi:"virtualHostName"`
}

// The set of arguments for constructing a VirtualHost resource.
type VirtualHostArgs struct {
	// InstanceId.
	InstanceId pulumi.StringInput
	// VirtualHostName.
	VirtualHostName pulumi.StringInput
}

func (VirtualHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHostArgs)(nil)).Elem()
}

type VirtualHostInput interface {
	pulumi.Input

	ToVirtualHostOutput() VirtualHostOutput
	ToVirtualHostOutputWithContext(ctx context.Context) VirtualHostOutput
}

func (*VirtualHost) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHost)(nil)).Elem()
}

func (i *VirtualHost) ToVirtualHostOutput() VirtualHostOutput {
	return i.ToVirtualHostOutputWithContext(context.Background())
}

func (i *VirtualHost) ToVirtualHostOutputWithContext(ctx context.Context) VirtualHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHostOutput)
}

// VirtualHostArrayInput is an input type that accepts VirtualHostArray and VirtualHostArrayOutput values.
// You can construct a concrete instance of `VirtualHostArrayInput` via:
//
//	VirtualHostArray{ VirtualHostArgs{...} }
type VirtualHostArrayInput interface {
	pulumi.Input

	ToVirtualHostArrayOutput() VirtualHostArrayOutput
	ToVirtualHostArrayOutputWithContext(context.Context) VirtualHostArrayOutput
}

type VirtualHostArray []VirtualHostInput

func (VirtualHostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHost)(nil)).Elem()
}

func (i VirtualHostArray) ToVirtualHostArrayOutput() VirtualHostArrayOutput {
	return i.ToVirtualHostArrayOutputWithContext(context.Background())
}

func (i VirtualHostArray) ToVirtualHostArrayOutputWithContext(ctx context.Context) VirtualHostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHostArrayOutput)
}

// VirtualHostMapInput is an input type that accepts VirtualHostMap and VirtualHostMapOutput values.
// You can construct a concrete instance of `VirtualHostMapInput` via:
//
//	VirtualHostMap{ "key": VirtualHostArgs{...} }
type VirtualHostMapInput interface {
	pulumi.Input

	ToVirtualHostMapOutput() VirtualHostMapOutput
	ToVirtualHostMapOutputWithContext(context.Context) VirtualHostMapOutput
}

type VirtualHostMap map[string]VirtualHostInput

func (VirtualHostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHost)(nil)).Elem()
}

func (i VirtualHostMap) ToVirtualHostMapOutput() VirtualHostMapOutput {
	return i.ToVirtualHostMapOutputWithContext(context.Background())
}

func (i VirtualHostMap) ToVirtualHostMapOutputWithContext(ctx context.Context) VirtualHostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHostMapOutput)
}

type VirtualHostOutput struct{ *pulumi.OutputState }

func (VirtualHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHost)(nil)).Elem()
}

func (o VirtualHostOutput) ToVirtualHostOutput() VirtualHostOutput {
	return o
}

func (o VirtualHostOutput) ToVirtualHostOutputWithContext(ctx context.Context) VirtualHostOutput {
	return o
}

// InstanceId.
func (o VirtualHostOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHost) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// VirtualHostName.
func (o VirtualHostOutput) VirtualHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHost) pulumi.StringOutput { return v.VirtualHostName }).(pulumi.StringOutput)
}

type VirtualHostArrayOutput struct{ *pulumi.OutputState }

func (VirtualHostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualHost)(nil)).Elem()
}

func (o VirtualHostArrayOutput) ToVirtualHostArrayOutput() VirtualHostArrayOutput {
	return o
}

func (o VirtualHostArrayOutput) ToVirtualHostArrayOutputWithContext(ctx context.Context) VirtualHostArrayOutput {
	return o
}

func (o VirtualHostArrayOutput) Index(i pulumi.IntInput) VirtualHostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualHost {
		return vs[0].([]*VirtualHost)[vs[1].(int)]
	}).(VirtualHostOutput)
}

type VirtualHostMapOutput struct{ *pulumi.OutputState }

func (VirtualHostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualHost)(nil)).Elem()
}

func (o VirtualHostMapOutput) ToVirtualHostMapOutput() VirtualHostMapOutput {
	return o
}

func (o VirtualHostMapOutput) ToVirtualHostMapOutputWithContext(ctx context.Context) VirtualHostMapOutput {
	return o
}

func (o VirtualHostMapOutput) MapIndex(k pulumi.StringInput) VirtualHostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualHost {
		return vs[0].(map[string]*VirtualHost)[vs[1].(string)]
	}).(VirtualHostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHostInput)(nil)).Elem(), &VirtualHost{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHostArrayInput)(nil)).Elem(), VirtualHostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHostMapInput)(nil)).Elem(), VirtualHostMap{})
	pulumi.RegisterOutputType(VirtualHostOutput{})
	pulumi.RegisterOutputType(VirtualHostArrayOutput{})
	pulumi.RegisterOutputType(VirtualHostMapOutput{})
}
