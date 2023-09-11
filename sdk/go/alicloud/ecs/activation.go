// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a ECS Activation resource.
//
// For information about ECS Activation and how to use it, see [What is Activation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/createactivation#doc-api-Ecs-CreateActivation).
//
// > **NOTE:** Available in v1.177.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewActivation(ctx, "example", &ecs.ActivationArgs{
//				Description:       pulumi.String("terraform-example"),
//				InstanceCount:     pulumi.Int(10),
//				InstanceName:      pulumi.String("terraform-example"),
//				IpAddressRange:    pulumi.String("0.0.0.0/0"),
//				TimeToLiveInHours: pulumi.Int(4),
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
// ECS Activation can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/activation:Activation example <id>
//
// ```
type Activation struct {
	pulumi.CustomResourceState

	// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
	// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	IpAddressRange pulumi.StringOutput `pulumi:"ipAddressRange"`
	// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
	TimeToLiveInHours pulumi.IntOutput `pulumi:"timeToLiveInHours"`
}

// NewActivation registers a new resource with the given unique name, arguments, and options.
func NewActivation(ctx *pulumi.Context,
	name string, args *ActivationArgs, opts ...pulumi.ResourceOption) (*Activation, error) {
	if args == nil {
		args = &ActivationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Activation
	err := ctx.RegisterResource("alicloud:ecs/activation:Activation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivation gets an existing Activation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivationState, opts ...pulumi.ResourceOption) (*Activation, error) {
	var resource Activation
	err := ctx.ReadResource("alicloud:ecs/activation:Activation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Activation resources.
type activationState struct {
	// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
	InstanceCount *int `pulumi:"instanceCount"`
	// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
	// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	InstanceName *string `pulumi:"instanceName"`
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	IpAddressRange *string `pulumi:"ipAddressRange"`
	// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
	TimeToLiveInHours *int `pulumi:"timeToLiveInHours"`
}

type ActivationState struct {
	// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
	InstanceCount pulumi.IntPtrInput
	// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
	// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	InstanceName pulumi.StringPtrInput
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	IpAddressRange pulumi.StringPtrInput
	// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
	TimeToLiveInHours pulumi.IntPtrInput
}

func (ActivationState) ElementType() reflect.Type {
	return reflect.TypeOf((*activationState)(nil)).Elem()
}

type activationArgs struct {
	// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
	InstanceCount *int `pulumi:"instanceCount"`
	// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
	// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	InstanceName *string `pulumi:"instanceName"`
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	IpAddressRange *string `pulumi:"ipAddressRange"`
	// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
	TimeToLiveInHours *int `pulumi:"timeToLiveInHours"`
}

// The set of arguments for constructing a Activation resource.
type ActivationArgs struct {
	// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
	InstanceCount pulumi.IntPtrInput
	// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
	// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	InstanceName pulumi.StringPtrInput
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	IpAddressRange pulumi.StringPtrInput
	// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
	TimeToLiveInHours pulumi.IntPtrInput
}

func (ActivationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activationArgs)(nil)).Elem()
}

type ActivationInput interface {
	pulumi.Input

	ToActivationOutput() ActivationOutput
	ToActivationOutputWithContext(ctx context.Context) ActivationOutput
}

func (*Activation) ElementType() reflect.Type {
	return reflect.TypeOf((**Activation)(nil)).Elem()
}

func (i *Activation) ToActivationOutput() ActivationOutput {
	return i.ToActivationOutputWithContext(context.Background())
}

func (i *Activation) ToActivationOutputWithContext(ctx context.Context) ActivationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationOutput)
}

func (i *Activation) ToOutput(ctx context.Context) pulumix.Output[*Activation] {
	return pulumix.Output[*Activation]{
		OutputState: i.ToActivationOutputWithContext(ctx).OutputState,
	}
}

// ActivationArrayInput is an input type that accepts ActivationArray and ActivationArrayOutput values.
// You can construct a concrete instance of `ActivationArrayInput` via:
//
//	ActivationArray{ ActivationArgs{...} }
type ActivationArrayInput interface {
	pulumi.Input

	ToActivationArrayOutput() ActivationArrayOutput
	ToActivationArrayOutputWithContext(context.Context) ActivationArrayOutput
}

type ActivationArray []ActivationInput

func (ActivationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Activation)(nil)).Elem()
}

func (i ActivationArray) ToActivationArrayOutput() ActivationArrayOutput {
	return i.ToActivationArrayOutputWithContext(context.Background())
}

func (i ActivationArray) ToActivationArrayOutputWithContext(ctx context.Context) ActivationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationArrayOutput)
}

func (i ActivationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Activation] {
	return pulumix.Output[[]*Activation]{
		OutputState: i.ToActivationArrayOutputWithContext(ctx).OutputState,
	}
}

// ActivationMapInput is an input type that accepts ActivationMap and ActivationMapOutput values.
// You can construct a concrete instance of `ActivationMapInput` via:
//
//	ActivationMap{ "key": ActivationArgs{...} }
type ActivationMapInput interface {
	pulumi.Input

	ToActivationMapOutput() ActivationMapOutput
	ToActivationMapOutputWithContext(context.Context) ActivationMapOutput
}

type ActivationMap map[string]ActivationInput

func (ActivationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Activation)(nil)).Elem()
}

func (i ActivationMap) ToActivationMapOutput() ActivationMapOutput {
	return i.ToActivationMapOutputWithContext(context.Background())
}

func (i ActivationMap) ToActivationMapOutputWithContext(ctx context.Context) ActivationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationMapOutput)
}

func (i ActivationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Activation] {
	return pulumix.Output[map[string]*Activation]{
		OutputState: i.ToActivationMapOutputWithContext(ctx).OutputState,
	}
}

type ActivationOutput struct{ *pulumi.OutputState }

func (ActivationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Activation)(nil)).Elem()
}

func (o ActivationOutput) ToActivationOutput() ActivationOutput {
	return o
}

func (o ActivationOutput) ToActivationOutputWithContext(ctx context.Context) ActivationOutput {
	return o
}

func (o ActivationOutput) ToOutput(ctx context.Context) pulumix.Output[*Activation] {
	return pulumix.Output[*Activation]{
		OutputState: o.OutputState,
	}
}

// The description of the activation code. The description can be 1 to 100 characters in length and cannot start with `http://` or `https://`.
func (o ActivationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Activation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The maximum number of times that the activation code can be used to register managed instances. Valid values: `1` to `1000`. Default value: `10`.
func (o ActivationOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Activation) pulumi.IntOutput { return v.InstanceCount }).(pulumi.IntOutput)
}

// The default instance name prefix. The instance name prefix must be 1 to 50 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The instance name prefix can contain only letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
// - If you use the activation code created by the CreateActivation operation to register managed instances, the instances are assigned sequential names that are prefixed by the value of this parameter. You can also specify a new instance name to override the assigned sequential name when you register a managed instance.
// - If you specify InstanceName when you register a managed instance, an instance name in the format of `<InstanceName>-<Number>` is generated. The number of digits in the <Number> value is determined by that in the InstanceCount value. Example: 001. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
func (o ActivationOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Activation) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
func (o ActivationOutput) IpAddressRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Activation) pulumi.StringOutput { return v.IpAddressRange }).(pulumi.StringOutput)
}

// The validity period of the activation code. The activation code cannot be used to register new instances after the validity period expires. Unit: hours. Valid values: `1` to `24`. Default value: `4`.
func (o ActivationOutput) TimeToLiveInHours() pulumi.IntOutput {
	return o.ApplyT(func(v *Activation) pulumi.IntOutput { return v.TimeToLiveInHours }).(pulumi.IntOutput)
}

type ActivationArrayOutput struct{ *pulumi.OutputState }

func (ActivationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Activation)(nil)).Elem()
}

func (o ActivationArrayOutput) ToActivationArrayOutput() ActivationArrayOutput {
	return o
}

func (o ActivationArrayOutput) ToActivationArrayOutputWithContext(ctx context.Context) ActivationArrayOutput {
	return o
}

func (o ActivationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Activation] {
	return pulumix.Output[[]*Activation]{
		OutputState: o.OutputState,
	}
}

func (o ActivationArrayOutput) Index(i pulumi.IntInput) ActivationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Activation {
		return vs[0].([]*Activation)[vs[1].(int)]
	}).(ActivationOutput)
}

type ActivationMapOutput struct{ *pulumi.OutputState }

func (ActivationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Activation)(nil)).Elem()
}

func (o ActivationMapOutput) ToActivationMapOutput() ActivationMapOutput {
	return o
}

func (o ActivationMapOutput) ToActivationMapOutputWithContext(ctx context.Context) ActivationMapOutput {
	return o
}

func (o ActivationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Activation] {
	return pulumix.Output[map[string]*Activation]{
		OutputState: o.OutputState,
	}
}

func (o ActivationMapOutput) MapIndex(k pulumi.StringInput) ActivationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Activation {
		return vs[0].(map[string]*Activation)[vs[1].(string)]
	}).(ActivationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActivationInput)(nil)).Elem(), &Activation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivationArrayInput)(nil)).Elem(), ActivationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivationMapInput)(nil)).Elem(), ActivationMap{})
	pulumi.RegisterOutputType(ActivationOutput{})
	pulumi.RegisterOutputType(ActivationArrayOutput{})
	pulumi.RegisterOutputType(ActivationMapOutput{})
}
