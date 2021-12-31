// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sddp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sddp.NewConfig(ctx, "_default", &sddp.ConfigArgs{
// 			Code:  pulumi.String("access_failed_cnt"),
// 			Value: pulumi.String("10"),
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
// Data Security Center Config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:sddp/config:Config example <code>
// ```
type Config struct {
	pulumi.CustomResourceState

	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringOutput `pulumi:"description"`
	// The language of the request and response. Valid values: `zh`,`en`.
	// * `zh`: Chinese.
	// * `en`: English.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	// * `accessFailedCnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
	// * `accessPermissionExprieMaxDays`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
	// * `logDatasizeAvgDays`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
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
	// * `zh`: Chinese.
	// * `en`: English.
	Lang *string `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	// * `accessFailedCnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
	// * `accessPermissionExprieMaxDays`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
	// * `logDatasizeAvgDays`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
	Value *string `pulumi:"value"`
}

type ConfigState struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrInput
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringPtrInput
	// The language of the request and response. Valid values: `zh`,`en`.
	// * `zh`: Chinese.
	// * `en`: English.
	Lang pulumi.StringPtrInput
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	// * `accessFailedCnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
	// * `accessPermissionExprieMaxDays`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
	// * `logDatasizeAvgDays`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
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
	// * `zh`: Chinese.
	// * `en`: English.
	Lang *string `pulumi:"lang"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	// * `accessFailedCnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
	// * `accessPermissionExprieMaxDays`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
	// * `logDatasizeAvgDays`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringPtrInput
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringPtrInput
	// The language of the request and response. Valid values: `zh`,`en`.
	// * `zh`: Chinese.
	// * `en`: English.
	Lang pulumi.StringPtrInput
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
	// * `accessFailedCnt`: Value Represents the Non-Authorized Resource Repeatedly Attempts to Access the Threshold.
	// * `accessPermissionExprieMaxDays`: Value Represents the Permissions during Periods of Inactivity Exceeding a Threshold.
	// * `logDatasizeAvgDays`: Value Represents the Date Certain Log Output Is Less than 10 Days before the Average Value of the Threshold.
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
	return reflect.TypeOf((*Config)(nil))
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

func (i *Config) ToConfigPtrOutput() ConfigPtrOutput {
	return i.ToConfigPtrOutputWithContext(context.Background())
}

func (i *Config) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigPtrOutput)
}

type ConfigPtrInput interface {
	pulumi.Input

	ToConfigPtrOutput() ConfigPtrOutput
	ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput
}

type configPtrType ConfigArgs

func (*configPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil))
}

func (i *configPtrType) ToConfigPtrOutput() ConfigPtrOutput {
	return i.ToConfigPtrOutputWithContext(context.Background())
}

func (i *configPtrType) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigPtrOutput)
}

// ConfigArrayInput is an input type that accepts ConfigArray and ConfigArrayOutput values.
// You can construct a concrete instance of `ConfigArrayInput` via:
//
//          ConfigArray{ ConfigArgs{...} }
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
//          ConfigMap{ "key": ConfigArgs{...} }
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
	return reflect.TypeOf((*Config)(nil))
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigPtrOutput() ConfigPtrOutput {
	return o.ToConfigPtrOutputWithContext(context.Background())
}

func (o ConfigOutput) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Config) *Config {
		return &v
	}).(ConfigPtrOutput)
}

type ConfigPtrOutput struct{ *pulumi.OutputState }

func (ConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Config)(nil))
}

func (o ConfigPtrOutput) ToConfigPtrOutput() ConfigPtrOutput {
	return o
}

func (o ConfigPtrOutput) ToConfigPtrOutputWithContext(ctx context.Context) ConfigPtrOutput {
	return o
}

func (o ConfigPtrOutput) Elem() ConfigOutput {
	return o.ApplyT(func(v *Config) Config {
		if v != nil {
			return *v
		}
		var ret Config
		return ret
	}).(ConfigOutput)
}

type ConfigArrayOutput struct{ *pulumi.OutputState }

func (ConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Config)(nil))
}

func (o ConfigArrayOutput) ToConfigArrayOutput() ConfigArrayOutput {
	return o
}

func (o ConfigArrayOutput) ToConfigArrayOutputWithContext(ctx context.Context) ConfigArrayOutput {
	return o
}

func (o ConfigArrayOutput) Index(i pulumi.IntInput) ConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Config {
		return vs[0].([]Config)[vs[1].(int)]
	}).(ConfigOutput)
}

type ConfigMapOutput struct{ *pulumi.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Config)(nil))
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) MapIndex(k pulumi.StringInput) ConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Config {
		return vs[0].(map[string]Config)[vs[1].(string)]
	}).(ConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigInput)(nil)).Elem(), &Config{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigPtrInput)(nil)).Elem(), &Config{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigArrayInput)(nil)).Elem(), ConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), ConfigMap{})
	pulumi.RegisterOutputType(ConfigOutput{})
	pulumi.RegisterOutputType(ConfigPtrOutput{})
	pulumi.RegisterOutputType(ConfigArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapOutput{})
}
