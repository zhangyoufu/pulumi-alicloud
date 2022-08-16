// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Function Compute trigger can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
//
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrOutput `pulumi:"configMns"`
	// The Function Compute function name.
	Function pulumi.StringOutput `pulumi:"function"`
	// The date this resource was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringOutput `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// The Function Compute service name.
	Service pulumi.StringOutput `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrOutput `pulumi:"sourceArn"`
	// The Function Compute trigger ID.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Function == nil {
		return nil, errors.New("invalid value for required argument 'Function'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Trigger
	err := ctx.RegisterResource("alicloud:fc/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("alicloud:fc/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config *string `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns *string `pulumi:"configMns"`
	// The Function Compute function name.
	Function *string `pulumi:"function"`
	// The date this resource was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role *string `pulumi:"role"`
	// The Function Compute service name.
	Service *string `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn *string `pulumi:"sourceArn"`
	// The Function Compute trigger ID.
	TriggerId *string `pulumi:"triggerId"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
	Type *string `pulumi:"type"`
}

type TriggerState struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrInput
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrInput
	// The Function Compute function name.
	Function pulumi.StringPtrInput
	// The date this resource was last modified.
	LastModified pulumi.StringPtrInput
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrInput
	// The Function Compute service name.
	Service pulumi.StringPtrInput
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrInput
	// The Function Compute trigger ID.
	TriggerId pulumi.StringPtrInput
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
	Type pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config *string `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns *string `pulumi:"configMns"`
	// The Function Compute function name.
	Function string `pulumi:"function"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role *string `pulumi:"role"`
	// The Function Compute service name.
	Service string `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn *string `pulumi:"sourceArn"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrInput
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrInput
	// The Function Compute function name.
	Function pulumi.StringInput
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrInput
	// The Function Compute service name.
	Service pulumi.StringInput
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrInput
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
	Type pulumi.StringInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//	TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//	TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
func (o TriggerOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Config }).(pulumi.StringPtrOutput)
}

// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
func (o TriggerOutput) ConfigMns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.ConfigMns }).(pulumi.StringPtrOutput)
}

// The Function Compute function name.
func (o TriggerOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// The date this resource was last modified.
func (o TriggerOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Setting a prefix to get a only trigger name. It is conflict with "name".
func (o TriggerOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
func (o TriggerOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// The Function Compute service name.
func (o TriggerOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
func (o TriggerOutput) SourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.SourceArn }).(pulumi.StringPtrOutput)
}

// The Function Compute trigger ID.
func (o TriggerOutput) TriggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.TriggerId }).(pulumi.StringOutput)
}

// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents", "eventbridge"].
func (o TriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
