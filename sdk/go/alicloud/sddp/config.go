// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sddp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Data Security Center Config resource.
//
// For information about Data Security Center Config and how to use it, see [What is Config](https://help.aliyun.com/product/88674.html).
//
// > **NOTE:** Available in v1.133.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sddp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sddp.NewConfig(ctx, "default", &sddp.ConfigArgs{
//				Code:  pulumi.String("access_failed_cnt"),
//				Value: pulumi.String("10"),
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
// Data Security Center Config can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sddp/config:Config example <code>
//
// ```
type Config struct {
	pulumi.CustomResourceState

	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringOutput `pulumi:"description"`
	// The language of the request and response. Valid values: `zh`,`en`.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		args = &ConfigArgs{}
	}

	var resource Config
	err := ctx.RegisterResource("alicloud:sddp/config:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("alicloud:sddp/config:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code *string `pulumi:"code"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description *string `pulumi:"description"`
	// The language of the request and response. Valid values: `zh`,`en`.
	Lang *string `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	Value *string `pulumi:"value"`
}

type ConfigState struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrInput
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringPtrInput
	// The language of the request and response. Valid values: `zh`,`en`.
	Lang pulumi.StringPtrInput
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	Value pulumi.StringPtrInput
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code *string `pulumi:"code"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description *string `pulumi:"description"`
	// The language of the request and response. Valid values: `zh`,`en`.
	Lang *string `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrInput
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringPtrInput
	// The language of the request and response. Valid values: `zh`,`en`.
	Lang pulumi.StringPtrInput
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	Value pulumi.StringPtrInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (*Config) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

// ConfigArrayInput is an input type that accepts ConfigArray and ConfigArrayOutput values.
// You can construct a concrete instance of `ConfigArrayInput` via:
//
//	ConfigArray{ ConfigArgs{...} }
type ConfigArrayInput interface {
	pulumi.Input

	ToConfigArrayOutput() ConfigArrayOutput
	ToConfigArrayOutputWithContext(context.Context) ConfigArrayOutput
}

type ConfigArray []ConfigInput

func (ConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Config)(nil)).Elem()
}

func (i ConfigArray) ToConfigArrayOutput() ConfigArrayOutput {
	return i.ToConfigArrayOutputWithContext(context.Background())
}

func (i ConfigArray) ToConfigArrayOutputWithContext(ctx context.Context) ConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigArrayOutput)
}

// ConfigMapInput is an input type that accepts ConfigMap and ConfigMapOutput values.
// You can construct a concrete instance of `ConfigMapInput` via:
//
//	ConfigMap{ "key": ConfigArgs{...} }
type ConfigMapInput interface {
	pulumi.Input

	ToConfigMapOutput() ConfigMapOutput
	ToConfigMapOutputWithContext(context.Context) ConfigMapOutput
}

type ConfigMap map[string]ConfigInput

func (ConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Config)(nil)).Elem()
}

func (i ConfigMap) ToConfigMapOutput() ConfigMapOutput {
	return i.ToConfigMapOutputWithContext(context.Background())
}

func (i ConfigMap) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapOutput)
}

type ConfigOutput struct{ *pulumi.OutputState }

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
func (o ConfigOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Config) pulumi.StringPtrOutput { return v.Code }).(pulumi.StringPtrOutput)
}

// Abnormal Alarm General Description of the Configuration Item.
func (o ConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Config) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The language of the request and response. Valid values: `zh`,`en`.
func (o ConfigOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Config) pulumi.StringPtrOutput { return v.Lang }).(pulumi.StringPtrOutput)
}

// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
func (o ConfigOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Config) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigArrayOutput struct{ *pulumi.OutputState }

func (ConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Config)(nil)).Elem()
}

func (o ConfigArrayOutput) ToConfigArrayOutput() ConfigArrayOutput {
	return o
}

func (o ConfigArrayOutput) ToConfigArrayOutputWithContext(ctx context.Context) ConfigArrayOutput {
	return o
}

func (o ConfigArrayOutput) Index(i pulumi.IntInput) ConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Config {
		return vs[0].([]*Config)[vs[1].(int)]
	}).(ConfigOutput)
}

type ConfigMapOutput struct{ *pulumi.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Config)(nil)).Elem()
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) MapIndex(k pulumi.StringInput) ConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Config {
		return vs[0].(map[string]*Config)[vs[1].(string)]
	}).(ConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInput)(nil)).Elem(), &Config{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigArrayInput)(nil)).Elem(), ConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), ConfigMap{})
	pulumi.RegisterOutputType(ConfigOutput{})
	pulumi.RegisterOutputType(ConfigArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapOutput{})
}
