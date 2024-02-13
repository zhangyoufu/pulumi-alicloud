// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OOS Parameter resource.
//
// For information about OOS Parameter and how to use it, see [What is Parameter](https://www.alibabacloud.com/help/en/doc-detail/183408.html).
//
// > **NOTE:** Available in v1.147.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oos"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = oos.NewParameter(ctx, "example", &oos.ParameterArgs{
//				ParameterName: pulumi.String("my-Parameter"),
//				Type:          pulumi.String("String"),
//				Value:         pulumi.String("example_value"),
//				Description:   pulumi.String("example_value"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("OosParameter"),
//				},
//				ResourceGroupId: *pulumi.String(_default.Groups[0].Id),
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
// OOS Parameter can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oos/parameter:Parameter example <parameter_name>
// ```
type Parameter struct {
	pulumi.CustomResourceState

	// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrOutput `pulumi:"constraints"`
	// The description of the common parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	ParameterName pulumi.StringOutput `pulumi:"parameterName"`
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The data type of the common parameter. Valid values: `String` and `StringList`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the common parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewParameter registers a new resource with the given unique name, arguments, and options.
func NewParameter(ctx *pulumi.Context,
	name string, args *ParameterArgs, opts ...pulumi.ResourceOption) (*Parameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParameterName == nil {
		return nil, errors.New("invalid value for required argument 'ParameterName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Parameter
	err := ctx.RegisterResource("alicloud:oos/parameter:Parameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameter gets an existing Parameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterState, opts ...pulumi.ResourceOption) (*Parameter, error) {
	var resource Parameter
	err := ctx.ReadResource("alicloud:oos/parameter:Parameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameter resources.
type parameterState struct {
	// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
	Constraints *string `pulumi:"constraints"`
	// The description of the common parameter. The description must be `1` to `200` characters in length.
	Description *string `pulumi:"description"`
	// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	ParameterName *string `pulumi:"parameterName"`
	// The ID of the Resource Group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data type of the common parameter. Valid values: `String` and `StringList`.
	Type *string `pulumi:"type"`
	// The value of the common parameter. The value must be `1` to `4096` characters in length.
	Value *string `pulumi:"value"`
}

type ParameterState struct {
	// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrInput
	// The description of the common parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringPtrInput
	// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	ParameterName pulumi.StringPtrInput
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The data type of the common parameter. Valid values: `String` and `StringList`.
	Type pulumi.StringPtrInput
	// The value of the common parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringPtrInput
}

func (ParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterState)(nil)).Elem()
}

type parameterArgs struct {
	// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
	Constraints *string `pulumi:"constraints"`
	// The description of the common parameter. The description must be `1` to `200` characters in length.
	Description *string `pulumi:"description"`
	// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	ParameterName string `pulumi:"parameterName"`
	// The ID of the Resource Group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The data type of the common parameter. Valid values: `String` and `StringList`.
	Type string `pulumi:"type"`
	// The value of the common parameter. The value must be `1` to `4096` characters in length.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Parameter resource.
type ParameterArgs struct {
	// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
	Constraints pulumi.StringPtrInput
	// The description of the common parameter. The description must be `1` to `200` characters in length.
	Description pulumi.StringPtrInput
	// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
	ParameterName pulumi.StringInput
	// The ID of the Resource Group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The data type of the common parameter. Valid values: `String` and `StringList`.
	Type pulumi.StringInput
	// The value of the common parameter. The value must be `1` to `4096` characters in length.
	Value pulumi.StringInput
}

func (ParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterArgs)(nil)).Elem()
}

type ParameterInput interface {
	pulumi.Input

	ToParameterOutput() ParameterOutput
	ToParameterOutputWithContext(ctx context.Context) ParameterOutput
}

func (*Parameter) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (i *Parameter) ToParameterOutput() ParameterOutput {
	return i.ToParameterOutputWithContext(context.Background())
}

func (i *Parameter) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterOutput)
}

// ParameterArrayInput is an input type that accepts ParameterArray and ParameterArrayOutput values.
// You can construct a concrete instance of `ParameterArrayInput` via:
//
//	ParameterArray{ ParameterArgs{...} }
type ParameterArrayInput interface {
	pulumi.Input

	ToParameterArrayOutput() ParameterArrayOutput
	ToParameterArrayOutputWithContext(context.Context) ParameterArrayOutput
}

type ParameterArray []ParameterInput

func (ParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (i ParameterArray) ToParameterArrayOutput() ParameterArrayOutput {
	return i.ToParameterArrayOutputWithContext(context.Background())
}

func (i ParameterArray) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterArrayOutput)
}

// ParameterMapInput is an input type that accepts ParameterMap and ParameterMapOutput values.
// You can construct a concrete instance of `ParameterMapInput` via:
//
//	ParameterMap{ "key": ParameterArgs{...} }
type ParameterMapInput interface {
	pulumi.Input

	ToParameterMapOutput() ParameterMapOutput
	ToParameterMapOutputWithContext(context.Context) ParameterMapOutput
}

type ParameterMap map[string]ParameterInput

func (ParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (i ParameterMap) ToParameterMapOutput() ParameterMapOutput {
	return i.ToParameterMapOutputWithContext(context.Background())
}

func (i ParameterMap) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterMapOutput)
}

type ParameterOutput struct{ *pulumi.OutputState }

func (ParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (o ParameterOutput) ToParameterOutput() ParameterOutput {
	return o
}

func (o ParameterOutput) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return o
}

// The constraints of the common parameter. This value follows the json format. By default, this parameter is null. Valid values:
func (o ParameterOutput) Constraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringPtrOutput { return v.Constraints }).(pulumi.StringPtrOutput)
}

// The description of the common parameter. The description must be `1` to `200` characters in length.
func (o ParameterOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the common parameter. The name must be `2` to `180` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/) and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, `ALICLOUD`, or `OOS`.
func (o ParameterOutput) ParameterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.ParameterName }).(pulumi.StringOutput)
}

// The ID of the Resource Group.
func (o ParameterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o ParameterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Parameter) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The data type of the common parameter. Valid values: `String` and `StringList`.
func (o ParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The value of the common parameter. The value must be `1` to `4096` characters in length.
func (o ParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type ParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (o ParameterArrayOutput) ToParameterArrayOutput() ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) Index(i pulumi.IntInput) ParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].([]*Parameter)[vs[1].(int)]
	}).(ParameterOutput)
}

type ParameterMapOutput struct{ *pulumi.OutputState }

func (ParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (o ParameterMapOutput) ToParameterMapOutput() ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) MapIndex(k pulumi.StringInput) ParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].(map[string]*Parameter)[vs[1].(string)]
	}).(ParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterInput)(nil)).Elem(), &Parameter{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterArrayInput)(nil)).Elem(), ParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterMapInput)(nil)).Elem(), ParameterMap{})
	pulumi.RegisterOutputType(ParameterOutput{})
	pulumi.RegisterOutputType(ParameterArrayOutput{})
	pulumi.RegisterOutputType(ParameterMapOutput{})
}
