// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Alb Ascript resource.
//
// For information about Alb Ascript and how to use it, see [What is AScript](https://www.alibabacloud.com/help/en/server-load-balancer/latest/api-doc-alb-2020-06-16-api-doc-createascripts).
//
// > **NOTE:** Available since v1.195.0.
//
// ## Import
//
// Alb AScript can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:alb/aScript:AScript example <id>
//
// ```
type AScript struct {
	pulumi.CustomResourceState

	// The name of AScript.
	AscriptName pulumi.StringOutput `pulumi:"ascriptName"`
	// Whether scripts are enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Whether extension parameters are enabled.
	ExtAttributeEnabled pulumi.BoolOutput `pulumi:"extAttributeEnabled"`
	// Extended attribute list. See `extAttributes` below for details.
	ExtAttributes AScriptExtAttributeArrayOutput `pulumi:"extAttributes"`
	// Listener ID of script attribution
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The ID of load balancer instance.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Execution location of AScript.
	Position pulumi.StringOutput `pulumi:"position"`
	// The content of AScript.
	ScriptContent pulumi.StringOutput `pulumi:"scriptContent"`
	// The status of AScript.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAScript registers a new resource with the given unique name, arguments, and options.
func NewAScript(ctx *pulumi.Context,
	name string, args *AScriptArgs, opts ...pulumi.ResourceOption) (*AScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscriptName == nil {
		return nil, errors.New("invalid value for required argument 'AscriptName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.Position == nil {
		return nil, errors.New("invalid value for required argument 'Position'")
	}
	if args.ScriptContent == nil {
		return nil, errors.New("invalid value for required argument 'ScriptContent'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AScript
	err := ctx.RegisterResource("alicloud:alb/aScript:AScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAScript gets an existing AScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AScriptState, opts ...pulumi.ResourceOption) (*AScript, error) {
	var resource AScript
	err := ctx.ReadResource("alicloud:alb/aScript:AScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AScript resources.
type ascriptState struct {
	// The name of AScript.
	AscriptName *string `pulumi:"ascriptName"`
	// Whether scripts are enabled.
	Enabled *bool `pulumi:"enabled"`
	// Whether extension parameters are enabled.
	ExtAttributeEnabled *bool `pulumi:"extAttributeEnabled"`
	// Extended attribute list. See `extAttributes` below for details.
	ExtAttributes []AScriptExtAttribute `pulumi:"extAttributes"`
	// Listener ID of script attribution
	ListenerId *string `pulumi:"listenerId"`
	// The ID of load balancer instance.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Execution location of AScript.
	Position *string `pulumi:"position"`
	// The content of AScript.
	ScriptContent *string `pulumi:"scriptContent"`
	// The status of AScript.
	Status *string `pulumi:"status"`
}

type AScriptState struct {
	// The name of AScript.
	AscriptName pulumi.StringPtrInput
	// Whether scripts are enabled.
	Enabled pulumi.BoolPtrInput
	// Whether extension parameters are enabled.
	ExtAttributeEnabled pulumi.BoolPtrInput
	// Extended attribute list. See `extAttributes` below for details.
	ExtAttributes AScriptExtAttributeArrayInput
	// Listener ID of script attribution
	ListenerId pulumi.StringPtrInput
	// The ID of load balancer instance.
	LoadBalancerId pulumi.StringPtrInput
	// Execution location of AScript.
	Position pulumi.StringPtrInput
	// The content of AScript.
	ScriptContent pulumi.StringPtrInput
	// The status of AScript.
	Status pulumi.StringPtrInput
}

func (AScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*ascriptState)(nil)).Elem()
}

type ascriptArgs struct {
	// The name of AScript.
	AscriptName string `pulumi:"ascriptName"`
	// Whether scripts are enabled.
	Enabled bool `pulumi:"enabled"`
	// Whether extension parameters are enabled.
	ExtAttributeEnabled *bool `pulumi:"extAttributeEnabled"`
	// Extended attribute list. See `extAttributes` below for details.
	ExtAttributes []AScriptExtAttribute `pulumi:"extAttributes"`
	// Listener ID of script attribution
	ListenerId string `pulumi:"listenerId"`
	// Execution location of AScript.
	Position string `pulumi:"position"`
	// The content of AScript.
	ScriptContent string `pulumi:"scriptContent"`
}

// The set of arguments for constructing a AScript resource.
type AScriptArgs struct {
	// The name of AScript.
	AscriptName pulumi.StringInput
	// Whether scripts are enabled.
	Enabled pulumi.BoolInput
	// Whether extension parameters are enabled.
	ExtAttributeEnabled pulumi.BoolPtrInput
	// Extended attribute list. See `extAttributes` below for details.
	ExtAttributes AScriptExtAttributeArrayInput
	// Listener ID of script attribution
	ListenerId pulumi.StringInput
	// Execution location of AScript.
	Position pulumi.StringInput
	// The content of AScript.
	ScriptContent pulumi.StringInput
}

func (AScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ascriptArgs)(nil)).Elem()
}

type AScriptInput interface {
	pulumi.Input

	ToAScriptOutput() AScriptOutput
	ToAScriptOutputWithContext(ctx context.Context) AScriptOutput
}

func (*AScript) ElementType() reflect.Type {
	return reflect.TypeOf((**AScript)(nil)).Elem()
}

func (i *AScript) ToAScriptOutput() AScriptOutput {
	return i.ToAScriptOutputWithContext(context.Background())
}

func (i *AScript) ToAScriptOutputWithContext(ctx context.Context) AScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AScriptOutput)
}

func (i *AScript) ToOutput(ctx context.Context) pulumix.Output[*AScript] {
	return pulumix.Output[*AScript]{
		OutputState: i.ToAScriptOutputWithContext(ctx).OutputState,
	}
}

// AScriptArrayInput is an input type that accepts AScriptArray and AScriptArrayOutput values.
// You can construct a concrete instance of `AScriptArrayInput` via:
//
//	AScriptArray{ AScriptArgs{...} }
type AScriptArrayInput interface {
	pulumi.Input

	ToAScriptArrayOutput() AScriptArrayOutput
	ToAScriptArrayOutputWithContext(context.Context) AScriptArrayOutput
}

type AScriptArray []AScriptInput

func (AScriptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AScript)(nil)).Elem()
}

func (i AScriptArray) ToAScriptArrayOutput() AScriptArrayOutput {
	return i.ToAScriptArrayOutputWithContext(context.Background())
}

func (i AScriptArray) ToAScriptArrayOutputWithContext(ctx context.Context) AScriptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AScriptArrayOutput)
}

func (i AScriptArray) ToOutput(ctx context.Context) pulumix.Output[[]*AScript] {
	return pulumix.Output[[]*AScript]{
		OutputState: i.ToAScriptArrayOutputWithContext(ctx).OutputState,
	}
}

// AScriptMapInput is an input type that accepts AScriptMap and AScriptMapOutput values.
// You can construct a concrete instance of `AScriptMapInput` via:
//
//	AScriptMap{ "key": AScriptArgs{...} }
type AScriptMapInput interface {
	pulumi.Input

	ToAScriptMapOutput() AScriptMapOutput
	ToAScriptMapOutputWithContext(context.Context) AScriptMapOutput
}

type AScriptMap map[string]AScriptInput

func (AScriptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AScript)(nil)).Elem()
}

func (i AScriptMap) ToAScriptMapOutput() AScriptMapOutput {
	return i.ToAScriptMapOutputWithContext(context.Background())
}

func (i AScriptMap) ToAScriptMapOutputWithContext(ctx context.Context) AScriptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AScriptMapOutput)
}

func (i AScriptMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AScript] {
	return pulumix.Output[map[string]*AScript]{
		OutputState: i.ToAScriptMapOutputWithContext(ctx).OutputState,
	}
}

type AScriptOutput struct{ *pulumi.OutputState }

func (AScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AScript)(nil)).Elem()
}

func (o AScriptOutput) ToAScriptOutput() AScriptOutput {
	return o
}

func (o AScriptOutput) ToAScriptOutputWithContext(ctx context.Context) AScriptOutput {
	return o
}

func (o AScriptOutput) ToOutput(ctx context.Context) pulumix.Output[*AScript] {
	return pulumix.Output[*AScript]{
		OutputState: o.OutputState,
	}
}

// The name of AScript.
func (o AScriptOutput) AscriptName() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.AscriptName }).(pulumi.StringOutput)
}

// Whether scripts are enabled.
func (o AScriptOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AScript) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Whether extension parameters are enabled.
func (o AScriptOutput) ExtAttributeEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AScript) pulumi.BoolOutput { return v.ExtAttributeEnabled }).(pulumi.BoolOutput)
}

// Extended attribute list. See `extAttributes` below for details.
func (o AScriptOutput) ExtAttributes() AScriptExtAttributeArrayOutput {
	return o.ApplyT(func(v *AScript) AScriptExtAttributeArrayOutput { return v.ExtAttributes }).(AScriptExtAttributeArrayOutput)
}

// Listener ID of script attribution
func (o AScriptOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The ID of load balancer instance.
func (o AScriptOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

// Execution location of AScript.
func (o AScriptOutput) Position() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.Position }).(pulumi.StringOutput)
}

// The content of AScript.
func (o AScriptOutput) ScriptContent() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.ScriptContent }).(pulumi.StringOutput)
}

// The status of AScript.
func (o AScriptOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AScript) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AScriptArrayOutput struct{ *pulumi.OutputState }

func (AScriptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AScript)(nil)).Elem()
}

func (o AScriptArrayOutput) ToAScriptArrayOutput() AScriptArrayOutput {
	return o
}

func (o AScriptArrayOutput) ToAScriptArrayOutputWithContext(ctx context.Context) AScriptArrayOutput {
	return o
}

func (o AScriptArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AScript] {
	return pulumix.Output[[]*AScript]{
		OutputState: o.OutputState,
	}
}

func (o AScriptArrayOutput) Index(i pulumi.IntInput) AScriptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AScript {
		return vs[0].([]*AScript)[vs[1].(int)]
	}).(AScriptOutput)
}

type AScriptMapOutput struct{ *pulumi.OutputState }

func (AScriptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AScript)(nil)).Elem()
}

func (o AScriptMapOutput) ToAScriptMapOutput() AScriptMapOutput {
	return o
}

func (o AScriptMapOutput) ToAScriptMapOutputWithContext(ctx context.Context) AScriptMapOutput {
	return o
}

func (o AScriptMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AScript] {
	return pulumix.Output[map[string]*AScript]{
		OutputState: o.OutputState,
	}
}

func (o AScriptMapOutput) MapIndex(k pulumi.StringInput) AScriptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AScript {
		return vs[0].(map[string]*AScript)[vs[1].(string)]
	}).(AScriptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AScriptInput)(nil)).Elem(), &AScript{})
	pulumi.RegisterInputType(reflect.TypeOf((*AScriptArrayInput)(nil)).Elem(), AScriptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AScriptMapInput)(nil)).Elem(), AScriptMap{})
	pulumi.RegisterOutputType(AScriptOutput{})
	pulumi.RegisterOutputType(AScriptArrayOutput{})
	pulumi.RegisterOutputType(AScriptMapOutput{})
}
