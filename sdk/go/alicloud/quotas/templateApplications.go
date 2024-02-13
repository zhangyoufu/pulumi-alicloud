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

// ## Import
//
// Quotas Template Applications can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:quotas/templateApplications:TemplateApplications example <id>
// ```
type TemplateApplications struct {
	pulumi.CustomResourceState

	// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
	// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
	AliyunUids pulumi.StringArrayOutput `pulumi:"aliyunUids"`
	// The value of the quota request.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	DesireValue pulumi.Float64Output `pulumi:"desireValue"`
	// Quota dimension. See `dimensions` below.
	Dimensions TemplateApplicationsDimensionArrayOutput `pulumi:"dimensions"`
	// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
	EffectiveTime pulumi.StringPtrOutput `pulumi:"effectiveTime"`
	// The language of the quota application result notification. Value:
	// - zh (default): Chinese.
	// - en: English.
	EnvLanguage pulumi.StringPtrOutput `pulumi:"envLanguage"`
	// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
	ExpireTime pulumi.StringPtrOutput `pulumi:"expireTime"`
	// Whether to send notification of quota application result. Value:
	// - 0 (default): No.
	// - 3: Yes.
	NoticeType pulumi.IntPtrOutput `pulumi:"noticeType"`
	// Cloud service name abbreviation.
	// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode pulumi.StringOutput `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode pulumi.StringOutput `pulumi:"quotaActionCode"`
	// List of quota application details.
	QuotaApplicationDetails TemplateApplicationsQuotaApplicationDetailArrayOutput `pulumi:"quotaApplicationDetails"`
	// The quota type. Value:
	// - CommonQuota (default): Generic quota.
	// - FlowControl:API rate quota.
	// - WhiteListLabel: Equity quota.
	QuotaCategory pulumi.StringOutput `pulumi:"quotaCategory"`
	// Reason for quota application.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	Reason pulumi.StringOutput `pulumi:"reason"`
}

// NewTemplateApplications registers a new resource with the given unique name, arguments, and options.
func NewTemplateApplications(ctx *pulumi.Context,
	name string, args *TemplateApplicationsArgs, opts ...pulumi.ResourceOption) (*TemplateApplications, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AliyunUids == nil {
		return nil, errors.New("invalid value for required argument 'AliyunUids'")
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
	if args.QuotaCategory == nil {
		return nil, errors.New("invalid value for required argument 'QuotaCategory'")
	}
	if args.Reason == nil {
		return nil, errors.New("invalid value for required argument 'Reason'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TemplateApplications
	err := ctx.RegisterResource("alicloud:quotas/templateApplications:TemplateApplications", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplateApplications gets an existing TemplateApplications resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplateApplications(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateApplicationsState, opts ...pulumi.ResourceOption) (*TemplateApplications, error) {
	var resource TemplateApplications
	err := ctx.ReadResource("alicloud:quotas/templateApplications:TemplateApplications", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TemplateApplications resources.
type templateApplicationsState struct {
	// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
	// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
	AliyunUids []string `pulumi:"aliyunUids"`
	// The value of the quota request.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	DesireValue *float64 `pulumi:"desireValue"`
	// Quota dimension. See `dimensions` below.
	Dimensions []TemplateApplicationsDimension `pulumi:"dimensions"`
	// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
	EffectiveTime *string `pulumi:"effectiveTime"`
	// The language of the quota application result notification. Value:
	// - zh (default): Chinese.
	// - en: English.
	EnvLanguage *string `pulumi:"envLanguage"`
	// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
	ExpireTime *string `pulumi:"expireTime"`
	// Whether to send notification of quota application result. Value:
	// - 0 (default): No.
	// - 3: Yes.
	NoticeType *int `pulumi:"noticeType"`
	// Cloud service name abbreviation.
	// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode *string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// List of quota application details.
	QuotaApplicationDetails []TemplateApplicationsQuotaApplicationDetail `pulumi:"quotaApplicationDetails"`
	// The quota type. Value:
	// - CommonQuota (default): Generic quota.
	// - FlowControl:API rate quota.
	// - WhiteListLabel: Equity quota.
	QuotaCategory *string `pulumi:"quotaCategory"`
	// Reason for quota application.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	Reason *string `pulumi:"reason"`
}

type TemplateApplicationsState struct {
	// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
	// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
	AliyunUids pulumi.StringArrayInput
	// The value of the quota request.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	DesireValue pulumi.Float64PtrInput
	// Quota dimension. See `dimensions` below.
	Dimensions TemplateApplicationsDimensionArrayInput
	// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
	EffectiveTime pulumi.StringPtrInput
	// The language of the quota application result notification. Value:
	// - zh (default): Chinese.
	// - en: English.
	EnvLanguage pulumi.StringPtrInput
	// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
	ExpireTime pulumi.StringPtrInput
	// Whether to send notification of quota application result. Value:
	// - 0 (default): No.
	// - 3: Yes.
	NoticeType pulumi.IntPtrInput
	// Cloud service name abbreviation.
	// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode pulumi.StringPtrInput
	// The quota ID.
	QuotaActionCode pulumi.StringPtrInput
	// List of quota application details.
	QuotaApplicationDetails TemplateApplicationsQuotaApplicationDetailArrayInput
	// The quota type. Value:
	// - CommonQuota (default): Generic quota.
	// - FlowControl:API rate quota.
	// - WhiteListLabel: Equity quota.
	QuotaCategory pulumi.StringPtrInput
	// Reason for quota application.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	Reason pulumi.StringPtrInput
}

func (TemplateApplicationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateApplicationsState)(nil)).Elem()
}

type templateApplicationsArgs struct {
	// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
	// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
	AliyunUids []string `pulumi:"aliyunUids"`
	// The value of the quota request.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	DesireValue float64 `pulumi:"desireValue"`
	// Quota dimension. See `dimensions` below.
	Dimensions []TemplateApplicationsDimension `pulumi:"dimensions"`
	// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
	EffectiveTime *string `pulumi:"effectiveTime"`
	// The language of the quota application result notification. Value:
	// - zh (default): Chinese.
	// - en: English.
	EnvLanguage *string `pulumi:"envLanguage"`
	// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
	ExpireTime *string `pulumi:"expireTime"`
	// Whether to send notification of quota application result. Value:
	// - 0 (default): No.
	// - 3: Yes.
	NoticeType *int `pulumi:"noticeType"`
	// Cloud service name abbreviation.
	// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode string `pulumi:"quotaActionCode"`
	// The quota type. Value:
	// - CommonQuota (default): Generic quota.
	// - FlowControl:API rate quota.
	// - WhiteListLabel: Equity quota.
	QuotaCategory string `pulumi:"quotaCategory"`
	// Reason for quota application.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	Reason string `pulumi:"reason"`
}

// The set of arguments for constructing a TemplateApplications resource.
type TemplateApplicationsArgs struct {
	// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
	// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
	AliyunUids pulumi.StringArrayInput
	// The value of the quota request.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	DesireValue pulumi.Float64Input
	// Quota dimension. See `dimensions` below.
	Dimensions TemplateApplicationsDimensionArrayInput
	// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
	EffectiveTime pulumi.StringPtrInput
	// The language of the quota application result notification. Value:
	// - zh (default): Chinese.
	// - en: English.
	EnvLanguage pulumi.StringPtrInput
	// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
	// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
	ExpireTime pulumi.StringPtrInput
	// Whether to send notification of quota application result. Value:
	// - 0 (default): No.
	// - 3: Yes.
	NoticeType pulumi.IntPtrInput
	// Cloud service name abbreviation.
	// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode pulumi.StringInput
	// The quota ID.
	QuotaActionCode pulumi.StringInput
	// The quota type. Value:
	// - CommonQuota (default): Generic quota.
	// - FlowControl:API rate quota.
	// - WhiteListLabel: Equity quota.
	QuotaCategory pulumi.StringInput
	// Reason for quota application.
	// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
	Reason pulumi.StringInput
}

func (TemplateApplicationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateApplicationsArgs)(nil)).Elem()
}

type TemplateApplicationsInput interface {
	pulumi.Input

	ToTemplateApplicationsOutput() TemplateApplicationsOutput
	ToTemplateApplicationsOutputWithContext(ctx context.Context) TemplateApplicationsOutput
}

func (*TemplateApplications) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateApplications)(nil)).Elem()
}

func (i *TemplateApplications) ToTemplateApplicationsOutput() TemplateApplicationsOutput {
	return i.ToTemplateApplicationsOutputWithContext(context.Background())
}

func (i *TemplateApplications) ToTemplateApplicationsOutputWithContext(ctx context.Context) TemplateApplicationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateApplicationsOutput)
}

// TemplateApplicationsArrayInput is an input type that accepts TemplateApplicationsArray and TemplateApplicationsArrayOutput values.
// You can construct a concrete instance of `TemplateApplicationsArrayInput` via:
//
//	TemplateApplicationsArray{ TemplateApplicationsArgs{...} }
type TemplateApplicationsArrayInput interface {
	pulumi.Input

	ToTemplateApplicationsArrayOutput() TemplateApplicationsArrayOutput
	ToTemplateApplicationsArrayOutputWithContext(context.Context) TemplateApplicationsArrayOutput
}

type TemplateApplicationsArray []TemplateApplicationsInput

func (TemplateApplicationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateApplications)(nil)).Elem()
}

func (i TemplateApplicationsArray) ToTemplateApplicationsArrayOutput() TemplateApplicationsArrayOutput {
	return i.ToTemplateApplicationsArrayOutputWithContext(context.Background())
}

func (i TemplateApplicationsArray) ToTemplateApplicationsArrayOutputWithContext(ctx context.Context) TemplateApplicationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateApplicationsArrayOutput)
}

// TemplateApplicationsMapInput is an input type that accepts TemplateApplicationsMap and TemplateApplicationsMapOutput values.
// You can construct a concrete instance of `TemplateApplicationsMapInput` via:
//
//	TemplateApplicationsMap{ "key": TemplateApplicationsArgs{...} }
type TemplateApplicationsMapInput interface {
	pulumi.Input

	ToTemplateApplicationsMapOutput() TemplateApplicationsMapOutput
	ToTemplateApplicationsMapOutputWithContext(context.Context) TemplateApplicationsMapOutput
}

type TemplateApplicationsMap map[string]TemplateApplicationsInput

func (TemplateApplicationsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateApplications)(nil)).Elem()
}

func (i TemplateApplicationsMap) ToTemplateApplicationsMapOutput() TemplateApplicationsMapOutput {
	return i.ToTemplateApplicationsMapOutputWithContext(context.Background())
}

func (i TemplateApplicationsMap) ToTemplateApplicationsMapOutputWithContext(ctx context.Context) TemplateApplicationsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateApplicationsMapOutput)
}

type TemplateApplicationsOutput struct{ *pulumi.OutputState }

func (TemplateApplicationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateApplications)(nil)).Elem()
}

func (o TemplateApplicationsOutput) ToTemplateApplicationsOutput() TemplateApplicationsOutput {
	return o
}

func (o TemplateApplicationsOutput) ToTemplateApplicationsOutputWithContext(ctx context.Context) TemplateApplicationsOutput {
	return o
}

// The list of Alibaba Cloud accounts (primary accounts) of the resource directory members to which the quota is applied.
// > **NOTE:**  Only 50 members can apply for quota increase in batch at a time. For more information about the members of the resource directory, see Query the list of all members in the resource directory.
func (o TemplateApplicationsOutput) AliyunUids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringArrayOutput { return v.AliyunUids }).(pulumi.StringArrayOutput)
}

// The value of the quota request.
// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
func (o TemplateApplicationsOutput) DesireValue() pulumi.Float64Output {
	return o.ApplyT(func(v *TemplateApplications) pulumi.Float64Output { return v.DesireValue }).(pulumi.Float64Output)
}

// Quota dimension. See `dimensions` below.
func (o TemplateApplicationsOutput) Dimensions() TemplateApplicationsDimensionArrayOutput {
	return o.ApplyT(func(v *TemplateApplications) TemplateApplicationsDimensionArrayOutput { return v.Dimensions }).(TemplateApplicationsDimensionArrayOutput)
}

// The UTC time when the quota takes effect. This parameter applies only to the equity quota (WhiteListLabel).
// > **NOTE:**  If the current account does not select the effective time, the default is the submission time.
func (o TemplateApplicationsOutput) EffectiveTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringPtrOutput { return v.EffectiveTime }).(pulumi.StringPtrOutput)
}

// The language of the quota application result notification. Value:
// - zh (default): Chinese.
// - en: English.
func (o TemplateApplicationsOutput) EnvLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringPtrOutput { return v.EnvLanguage }).(pulumi.StringPtrOutput)
}

// The UTC time when the quota expires. This parameter applies only to the equity quota (WhiteListLabel).
// > **NOTE:**  If No Expiration Time is selected for the current account, the expiration time is 99 years from the effective time of the current quota.
func (o TemplateApplicationsOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringPtrOutput { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

// Whether to send notification of quota application result. Value:
// - 0 (default): No.
// - 3: Yes.
func (o TemplateApplicationsOutput) NoticeType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.IntPtrOutput { return v.NoticeType }).(pulumi.IntPtrOutput)
}

// Cloud service name abbreviation.
// > **NOTE:**  For more information about cloud services that support quota centers, see Cloud services that support quota centers.
func (o TemplateApplicationsOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringOutput { return v.ProductCode }).(pulumi.StringOutput)
}

// The quota ID.
func (o TemplateApplicationsOutput) QuotaActionCode() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringOutput { return v.QuotaActionCode }).(pulumi.StringOutput)
}

// List of quota application details.
func (o TemplateApplicationsOutput) QuotaApplicationDetails() TemplateApplicationsQuotaApplicationDetailArrayOutput {
	return o.ApplyT(func(v *TemplateApplications) TemplateApplicationsQuotaApplicationDetailArrayOutput {
		return v.QuotaApplicationDetails
	}).(TemplateApplicationsQuotaApplicationDetailArrayOutput)
}

// The quota type. Value:
// - CommonQuota (default): Generic quota.
// - FlowControl:API rate quota.
// - WhiteListLabel: Equity quota.
func (o TemplateApplicationsOutput) QuotaCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringOutput { return v.QuotaCategory }).(pulumi.StringOutput)
}

// Reason for quota application.
// > **NOTE:**  The quota request is approved by the technical support of each cloud service. If you want to increase the chance of passing, please fill in a reasonable application value and detailed application reasons when applying for quota.
func (o TemplateApplicationsOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateApplications) pulumi.StringOutput { return v.Reason }).(pulumi.StringOutput)
}

type TemplateApplicationsArrayOutput struct{ *pulumi.OutputState }

func (TemplateApplicationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TemplateApplications)(nil)).Elem()
}

func (o TemplateApplicationsArrayOutput) ToTemplateApplicationsArrayOutput() TemplateApplicationsArrayOutput {
	return o
}

func (o TemplateApplicationsArrayOutput) ToTemplateApplicationsArrayOutputWithContext(ctx context.Context) TemplateApplicationsArrayOutput {
	return o
}

func (o TemplateApplicationsArrayOutput) Index(i pulumi.IntInput) TemplateApplicationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TemplateApplications {
		return vs[0].([]*TemplateApplications)[vs[1].(int)]
	}).(TemplateApplicationsOutput)
}

type TemplateApplicationsMapOutput struct{ *pulumi.OutputState }

func (TemplateApplicationsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TemplateApplications)(nil)).Elem()
}

func (o TemplateApplicationsMapOutput) ToTemplateApplicationsMapOutput() TemplateApplicationsMapOutput {
	return o
}

func (o TemplateApplicationsMapOutput) ToTemplateApplicationsMapOutputWithContext(ctx context.Context) TemplateApplicationsMapOutput {
	return o
}

func (o TemplateApplicationsMapOutput) MapIndex(k pulumi.StringInput) TemplateApplicationsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TemplateApplications {
		return vs[0].(map[string]*TemplateApplications)[vs[1].(string)]
	}).(TemplateApplicationsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateApplicationsInput)(nil)).Elem(), &TemplateApplications{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateApplicationsArrayInput)(nil)).Elem(), TemplateApplicationsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateApplicationsMapInput)(nil)).Elem(), TemplateApplicationsMap{})
	pulumi.RegisterOutputType(TemplateApplicationsOutput{})
	pulumi.RegisterOutputType(TemplateApplicationsArrayOutput{})
	pulumi.RegisterOutputType(TemplateApplicationsMapOutput{})
}
