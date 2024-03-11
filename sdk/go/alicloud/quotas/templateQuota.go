// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Quotas Template Quota resource.
//
// For information about Quotas Template Quota and how to use it, see [What is Template Quota](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createtemplatequotaitem).
//
// > **NOTE:** Available since v1.206.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := quotas.NewTemplateQuota(ctx, "default", &quotas.TemplateQuotaArgs{
//				DesireValue: pulumi.Float64(1001),
//				Dimensions: quotas.TemplateQuotaDimensionArray{
//					&quotas.TemplateQuotaDimensionArgs{
//						Key:   pulumi.String("regionId"),
//						Value: pulumi.String("cn-hangzhou"),
//					},
//				},
//				EnvLanguage:     pulumi.String("zh"),
//				NoticeType:      pulumi.Int(3),
//				ProductCode:     pulumi.String("gws"),
//				QuotaActionCode: pulumi.String("q_desktop-count"),
//				QuotaCategory:   pulumi.String("CommonQuota"),
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
// Quotas Template Quota can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:quotas/templateQuota:TemplateQuota example <id>
// ```
type TemplateQuota struct {
	pulumi.CustomResourceState

	// Quota application value.
	DesireValue pulumi.Float64Output `pulumi:"desireValue"`
	// The Quota Dimensions. See `dimensions` below.
	Dimensions TemplateQuotaDimensionArrayOutput `pulumi:"dimensions"`
	// The UTC time when the quota takes effect.
	EffectiveTime pulumi.StringPtrOutput `pulumi:"effectiveTime"`
	// The language of the quota alert notification. Value:
	// - zh: Chinese.
	// - en: English.
	EnvLanguage pulumi.StringOutput `pulumi:"envLanguage"`
	// The UTC time when the quota expires.
	ExpireTime pulumi.StringPtrOutput `pulumi:"expireTime"`
	// Whether to notify the result of quota promotion application. Value:
	// - 0: No.
	// - 3: Yes.
	NoticeType pulumi.IntOutput `pulumi:"noticeType"`
	// The abbreviation of the cloud service name.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode pulumi.StringOutput `pulumi:"quotaActionCode"`
	// Type of quota. Value:
	// - CommonQuota : Generic quota.
	// - WhiteListLabel: Equity quota.
	// - FlowControl:API rate quota.
	QuotaCategory pulumi.StringPtrOutput `pulumi:"quotaCategory"`
}

// NewTemplateQuota registers a new resource with the given unique name, arguments, and options.
func NewTemplateQuota(ctx *pulumi.Context,
	name string, args *TemplateQuotaArgs, opts ...pulumi.ResourceOption) (*TemplateQuota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DesireValue == nil {
		return nil, errors.New("invalid value for required argument 'DesireValue'")
	}
	if args.ProductCode == nil {
		return nil, errors.New("invalid value for required argument 'ProductCode'")
	}
	if args.QuotaActionCode == nil {
		return nil, errors.New("invalid value for required argument 'QuotaActionCode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TemplateQuota
	err := ctx.RegisterResource("alicloud:quotas/templateQuota:TemplateQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateQuota gets an existing TemplateQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateQuotaState, opts ...pulumi.ResourceOption) (*TemplateQuota, error) {
	var resource TemplateQuota
	err := ctx.ReadResource("alicloud:quotas/templateQuota:TemplateQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateQuota resources.
type templateQuotaState struct {
	// Quota application value.
	DesireValue *float64 `pulumi:"desireValue"`
	// The Quota Dimensions. See `dimensions` below.
	Dimensions []TemplateQuotaDimension `pulumi:"dimensions"`
	// The UTC time when the quota takes effect.
	EffectiveTime *string `pulumi:"effectiveTime"`
	// The language of the quota alert notification. Value:
	// - zh: Chinese.
	// - en: English.
	EnvLanguage *string `pulumi:"envLanguage"`
	// The UTC time when the quota expires.
	ExpireTime *string `pulumi:"expireTime"`
	// Whether to notify the result of quota promotion application. Value:
	// - 0: No.
	// - 3: Yes.
	NoticeType *int `pulumi:"noticeType"`
	// The abbreviation of the cloud service name.
	ProductCode *string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// Type of quota. Value:
	// - CommonQuota : Generic quota.
	// - WhiteListLabel: Equity quota.
	// - FlowControl:API rate quota.
	QuotaCategory *string `pulumi:"quotaCategory"`
}

type TemplateQuotaState struct {
	// Quota application value.
	DesireValue pulumi.Float64PtrInput
	// The Quota Dimensions. See `dimensions` below.
	Dimensions TemplateQuotaDimensionArrayInput
	// The UTC time when the quota takes effect.
	EffectiveTime pulumi.StringPtrInput
	// The language of the quota alert notification. Value:
	// - zh: Chinese.
	// - en: English.
	EnvLanguage pulumi.StringPtrInput
	// The UTC time when the quota expires.
	ExpireTime pulumi.StringPtrInput
	// Whether to notify the result of quota promotion application. Value:
	// - 0: No.
	// - 3: Yes.
	NoticeType pulumi.IntPtrInput
	// The abbreviation of the cloud service name.
	ProductCode pulumi.StringPtrInput
	// The quota ID.
	QuotaActionCode pulumi.StringPtrInput
	// Type of quota. Value:
	// - CommonQuota : Generic quota.
	// - WhiteListLabel: Equity quota.
	// - FlowControl:API rate quota.
	QuotaCategory pulumi.StringPtrInput
}

func (TemplateQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateQuotaState)(nil)).Elem()
}

type templateQuotaArgs struct {
	// Quota application value.
	DesireValue float64 `pulumi:"desireValue"`
	// The Quota Dimensions. See `dimensions` below.
	Dimensions []TemplateQuotaDimension `pulumi:"dimensions"`
	// The UTC time when the quota takes effect.
	EffectiveTime *string `pulumi:"effectiveTime"`
	// The language of the quota alert notification. Value:
	// - zh: Chinese.
	// - en: English.
	EnvLanguage *string `pulumi:"envLanguage"`
	// The UTC time when the quota expires.
	ExpireTime *string `pulumi:"expireTime"`
	// Whether to notify the result of quota promotion application. Value:
	// - 0: No.
	// - 3: Yes.
	NoticeType *int `pulumi:"noticeType"`
	// The abbreviation of the cloud service name.
	ProductCode string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode string `pulumi:"quotaActionCode"`
	// Type of quota. Value:
	// - CommonQuota : Generic quota.
	// - WhiteListLabel: Equity quota.
	// - FlowControl:API rate quota.
	QuotaCategory *string `pulumi:"quotaCategory"`
}

// The set of arguments for constructing a TemplateQuota resource.
type TemplateQuotaArgs struct {
	// Quota application value.
	DesireValue pulumi.Float64Input
	// The Quota Dimensions. See `dimensions` below.
	Dimensions TemplateQuotaDimensionArrayInput
	// The UTC time when the quota takes effect.
	EffectiveTime pulumi.StringPtrInput
	// The language of the quota alert notification. Value:
	// - zh: Chinese.
	// - en: English.
	EnvLanguage pulumi.StringPtrInput
	// The UTC time when the quota expires.
	ExpireTime pulumi.StringPtrInput
	// Whether to notify the result of quota promotion application. Value:
	// - 0: No.
	// - 3: Yes.
	NoticeType pulumi.IntPtrInput
	// The abbreviation of the cloud service name.
	ProductCode pulumi.StringInput
	// The quota ID.
	QuotaActionCode pulumi.StringInput
	// Type of quota. Value:
	// - CommonQuota : Generic quota.
	// - WhiteListLabel: Equity quota.
	// - FlowControl:API rate quota.
	QuotaCategory pulumi.StringPtrInput
}

func (TemplateQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateQuotaArgs)(nil)).Elem()
}

type TemplateQuotaInput interface {
	pulumi.Input

	ToTemplateQuotaOutput() TemplateQuotaOutput
	ToTemplateQuotaOutputWithContext(ctx context.Context) TemplateQuotaOutput
}

func (*TemplateQuota) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateQuota)(nil)).Elem()
}

func (i *TemplateQuota) ToTemplateQuotaOutput() TemplateQuotaOutput {
	return i.ToTemplateQuotaOutputWithContext(context.Background())
}

func (i *TemplateQuota) ToTemplateQuotaOutputWithContext(ctx context.Context) TemplateQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateQuotaOutput)
}

// TemplateQuotaArrayInput is an input type that accepts TemplateQuotaArray and TemplateQuotaArrayOutput values.
// You can construct a concrete instance of `TemplateQuotaArrayInput` via:
//
//	TemplateQuotaArray{ TemplateQuotaArgs{...} }
type TemplateQuotaArrayInput interface {
	pulumi.Input

	ToTemplateQuotaArrayOutput() TemplateQuotaArrayOutput
	ToTemplateQuotaArrayOutputWithContext(context.Context) TemplateQuotaArrayOutput
}

type TemplateQuotaArray []TemplateQuotaInput

func (TemplateQuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateQuota)(nil)).Elem()
}

func (i TemplateQuotaArray) ToTemplateQuotaArrayOutput() TemplateQuotaArrayOutput {
	return i.ToTemplateQuotaArrayOutputWithContext(context.Background())
}

func (i TemplateQuotaArray) ToTemplateQuotaArrayOutputWithContext(ctx context.Context) TemplateQuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateQuotaArrayOutput)
}

// TemplateQuotaMapInput is an input type that accepts TemplateQuotaMap and TemplateQuotaMapOutput values.
// You can construct a concrete instance of `TemplateQuotaMapInput` via:
//
//	TemplateQuotaMap{ "key": TemplateQuotaArgs{...} }
type TemplateQuotaMapInput interface {
	pulumi.Input

	ToTemplateQuotaMapOutput() TemplateQuotaMapOutput
	ToTemplateQuotaMapOutputWithContext(context.Context) TemplateQuotaMapOutput
}

type TemplateQuotaMap map[string]TemplateQuotaInput

func (TemplateQuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateQuota)(nil)).Elem()
}

func (i TemplateQuotaMap) ToTemplateQuotaMapOutput() TemplateQuotaMapOutput {
	return i.ToTemplateQuotaMapOutputWithContext(context.Background())
}

func (i TemplateQuotaMap) ToTemplateQuotaMapOutputWithContext(ctx context.Context) TemplateQuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateQuotaMapOutput)
}

type TemplateQuotaOutput struct{ *pulumi.OutputState }

func (TemplateQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateQuota)(nil)).Elem()
}

func (o TemplateQuotaOutput) ToTemplateQuotaOutput() TemplateQuotaOutput {
	return o
}

func (o TemplateQuotaOutput) ToTemplateQuotaOutputWithContext(ctx context.Context) TemplateQuotaOutput {
	return o
}

// Quota application value.
func (o TemplateQuotaOutput) DesireValue() pulumi.Float64Output {
	return o.ApplyT(func(v *TemplateQuota) pulumi.Float64Output { return v.DesireValue }).(pulumi.Float64Output)
}

// The Quota Dimensions. See `dimensions` below.
func (o TemplateQuotaOutput) Dimensions() TemplateQuotaDimensionArrayOutput {
	return o.ApplyT(func(v *TemplateQuota) TemplateQuotaDimensionArrayOutput { return v.Dimensions }).(TemplateQuotaDimensionArrayOutput)
}

// The UTC time when the quota takes effect.
func (o TemplateQuotaOutput) EffectiveTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringPtrOutput { return v.EffectiveTime }).(pulumi.StringPtrOutput)
}

// The language of the quota alert notification. Value:
// - zh: Chinese.
// - en: English.
func (o TemplateQuotaOutput) EnvLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringOutput { return v.EnvLanguage }).(pulumi.StringOutput)
}

// The UTC time when the quota expires.
func (o TemplateQuotaOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringPtrOutput { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

// Whether to notify the result of quota promotion application. Value:
// - 0: No.
// - 3: Yes.
func (o TemplateQuotaOutput) NoticeType() pulumi.IntOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.IntOutput { return v.NoticeType }).(pulumi.IntOutput)
}

// The abbreviation of the cloud service name.
func (o TemplateQuotaOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringOutput { return v.ProductCode }).(pulumi.StringOutput)
}

// The quota ID.
func (o TemplateQuotaOutput) QuotaActionCode() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringOutput { return v.QuotaActionCode }).(pulumi.StringOutput)
}

// Type of quota. Value:
// - CommonQuota : Generic quota.
// - WhiteListLabel: Equity quota.
// - FlowControl:API rate quota.
func (o TemplateQuotaOutput) QuotaCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateQuota) pulumi.StringPtrOutput { return v.QuotaCategory }).(pulumi.StringPtrOutput)
}

type TemplateQuotaArrayOutput struct{ *pulumi.OutputState }

func (TemplateQuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateQuota)(nil)).Elem()
}

func (o TemplateQuotaArrayOutput) ToTemplateQuotaArrayOutput() TemplateQuotaArrayOutput {
	return o
}

func (o TemplateQuotaArrayOutput) ToTemplateQuotaArrayOutputWithContext(ctx context.Context) TemplateQuotaArrayOutput {
	return o
}

func (o TemplateQuotaArrayOutput) Index(i pulumi.IntInput) TemplateQuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemplateQuota {
		return vs[0].([]*TemplateQuota)[vs[1].(int)]
	}).(TemplateQuotaOutput)
}

type TemplateQuotaMapOutput struct{ *pulumi.OutputState }

func (TemplateQuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateQuota)(nil)).Elem()
}

func (o TemplateQuotaMapOutput) ToTemplateQuotaMapOutput() TemplateQuotaMapOutput {
	return o
}

func (o TemplateQuotaMapOutput) ToTemplateQuotaMapOutputWithContext(ctx context.Context) TemplateQuotaMapOutput {
	return o
}

func (o TemplateQuotaMapOutput) MapIndex(k pulumi.StringInput) TemplateQuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemplateQuota {
		return vs[0].(map[string]*TemplateQuota)[vs[1].(string)]
	}).(TemplateQuotaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateQuotaInput)(nil)).Elem(), &TemplateQuota{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateQuotaArrayInput)(nil)).Elem(), TemplateQuotaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateQuotaMapInput)(nil)).Elem(), TemplateQuotaMap{})
	pulumi.RegisterOutputType(TemplateQuotaOutput{})
	pulumi.RegisterOutputType(TemplateQuotaArrayOutput{})
	pulumi.RegisterOutputType(TemplateQuotaMapOutput{})
}
