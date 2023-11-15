// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sddp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetConfigsConfig struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code string `pulumi:"code"`
	// Configure the Number.
	ConfigId string `pulumi:"configId"`
	// Default Value.
	DefaultValue string `pulumi:"defaultValue"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description string `pulumi:"description"`
	// The ID of the Config.
	Id string `pulumi:"id"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
	Value string `pulumi:"value"`
}

// GetConfigsConfigInput is an input type that accepts GetConfigsConfigArgs and GetConfigsConfigOutput values.
// You can construct a concrete instance of `GetConfigsConfigInput` via:
//
//	GetConfigsConfigArgs{...}
type GetConfigsConfigInput interface {
	pulumi.Input

	ToGetConfigsConfigOutput() GetConfigsConfigOutput
	ToGetConfigsConfigOutputWithContext(context.Context) GetConfigsConfigOutput
}

type GetConfigsConfigArgs struct {
	// Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
	Code pulumi.StringInput `pulumi:"code"`
	// Configure the Number.
	ConfigId pulumi.StringInput `pulumi:"configId"`
	// Default Value.
	DefaultValue pulumi.StringInput `pulumi:"defaultValue"`
	// Abnormal Alarm General Description of the Configuration Item.
	Description pulumi.StringInput `pulumi:"description"`
	// The ID of the Config.
	Id pulumi.StringInput `pulumi:"id"`
	// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetConfigsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigsConfig)(nil)).Elem()
}

func (i GetConfigsConfigArgs) ToGetConfigsConfigOutput() GetConfigsConfigOutput {
	return i.ToGetConfigsConfigOutputWithContext(context.Background())
}

func (i GetConfigsConfigArgs) ToGetConfigsConfigOutputWithContext(ctx context.Context) GetConfigsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetConfigsConfigOutput)
}

// GetConfigsConfigArrayInput is an input type that accepts GetConfigsConfigArray and GetConfigsConfigArrayOutput values.
// You can construct a concrete instance of `GetConfigsConfigArrayInput` via:
//
//	GetConfigsConfigArray{ GetConfigsConfigArgs{...} }
type GetConfigsConfigArrayInput interface {
	pulumi.Input

	ToGetConfigsConfigArrayOutput() GetConfigsConfigArrayOutput
	ToGetConfigsConfigArrayOutputWithContext(context.Context) GetConfigsConfigArrayOutput
}

type GetConfigsConfigArray []GetConfigsConfigInput

func (GetConfigsConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetConfigsConfig)(nil)).Elem()
}

func (i GetConfigsConfigArray) ToGetConfigsConfigArrayOutput() GetConfigsConfigArrayOutput {
	return i.ToGetConfigsConfigArrayOutputWithContext(context.Background())
}

func (i GetConfigsConfigArray) ToGetConfigsConfigArrayOutputWithContext(ctx context.Context) GetConfigsConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetConfigsConfigArrayOutput)
}

type GetConfigsConfigOutput struct{ *pulumi.OutputState }

func (GetConfigsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigsConfig)(nil)).Elem()
}

func (o GetConfigsConfigOutput) ToGetConfigsConfigOutput() GetConfigsConfigOutput {
	return o
}

func (o GetConfigsConfigOutput) ToGetConfigsConfigOutputWithContext(ctx context.Context) GetConfigsConfigOutput {
	return o
}

// Abnormal Alarm General Configuration Module by Using the Encoding.Valid values: `accessFailedCnt`, `accessPermissionExprieMaxDays`, `logDatasizeAvgDays`.
func (o GetConfigsConfigOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.Code }).(pulumi.StringOutput)
}

// Configure the Number.
func (o GetConfigsConfigOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.ConfigId }).(pulumi.StringOutput)
}

// Default Value.
func (o GetConfigsConfigOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.DefaultValue }).(pulumi.StringOutput)
}

// Abnormal Alarm General Description of the Configuration Item.
func (o GetConfigsConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the Config.
func (o GetConfigsConfigOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.Id }).(pulumi.StringOutput)
}

// The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different.
func (o GetConfigsConfigOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigsConfig) string { return v.Value }).(pulumi.StringOutput)
}

type GetConfigsConfigArrayOutput struct{ *pulumi.OutputState }

func (GetConfigsConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetConfigsConfig)(nil)).Elem()
}

func (o GetConfigsConfigArrayOutput) ToGetConfigsConfigArrayOutput() GetConfigsConfigArrayOutput {
	return o
}

func (o GetConfigsConfigArrayOutput) ToGetConfigsConfigArrayOutputWithContext(ctx context.Context) GetConfigsConfigArrayOutput {
	return o
}

func (o GetConfigsConfigArrayOutput) Index(i pulumi.IntInput) GetConfigsConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetConfigsConfig {
		return vs[0].([]GetConfigsConfig)[vs[1].(int)]
	}).(GetConfigsConfigOutput)
}

type GetDataLimitsLimit struct {
	// Whether to enable the log auditing feature.
	AuditStatus int `pulumi:"auditStatus"`
	// The status of the connectivity test between the data asset and SDDP.
	CheckStatus int `pulumi:"checkStatus"`
	// The first ID of the resource.
	DataLimitId string `pulumi:"dataLimitId"`
	// The type of the database.
	EngineType string `pulumi:"engineType"`
	// The ID of the Data Limit.
	Id string `pulumi:"id"`
	// The name of the service to which the data asset belongs.
	LocalName string `pulumi:"localName"`
	// The retention period of raw logs after you enable the log auditing feature.
	LogStoreDay int `pulumi:"logStoreDay"`
	// The ID of the data asset.
	ParentId string `pulumi:"parentId"`
	// The port that is used to connect to the database.
	Port int `pulumi:"port"`
	// The type of the service to which the data asset belongs.
	ResourceType string `pulumi:"resourceType"`
	// The name of the user who owns the data asset.
	UserName string `pulumi:"userName"`
}

// GetDataLimitsLimitInput is an input type that accepts GetDataLimitsLimitArgs and GetDataLimitsLimitOutput values.
// You can construct a concrete instance of `GetDataLimitsLimitInput` via:
//
//	GetDataLimitsLimitArgs{...}
type GetDataLimitsLimitInput interface {
	pulumi.Input

	ToGetDataLimitsLimitOutput() GetDataLimitsLimitOutput
	ToGetDataLimitsLimitOutputWithContext(context.Context) GetDataLimitsLimitOutput
}

type GetDataLimitsLimitArgs struct {
	// Whether to enable the log auditing feature.
	AuditStatus pulumi.IntInput `pulumi:"auditStatus"`
	// The status of the connectivity test between the data asset and SDDP.
	CheckStatus pulumi.IntInput `pulumi:"checkStatus"`
	// The first ID of the resource.
	DataLimitId pulumi.StringInput `pulumi:"dataLimitId"`
	// The type of the database.
	EngineType pulumi.StringInput `pulumi:"engineType"`
	// The ID of the Data Limit.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the service to which the data asset belongs.
	LocalName pulumi.StringInput `pulumi:"localName"`
	// The retention period of raw logs after you enable the log auditing feature.
	LogStoreDay pulumi.IntInput `pulumi:"logStoreDay"`
	// The ID of the data asset.
	ParentId pulumi.StringInput `pulumi:"parentId"`
	// The port that is used to connect to the database.
	Port pulumi.IntInput `pulumi:"port"`
	// The type of the service to which the data asset belongs.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// The name of the user who owns the data asset.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (GetDataLimitsLimitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataLimitsLimit)(nil)).Elem()
}

func (i GetDataLimitsLimitArgs) ToGetDataLimitsLimitOutput() GetDataLimitsLimitOutput {
	return i.ToGetDataLimitsLimitOutputWithContext(context.Background())
}

func (i GetDataLimitsLimitArgs) ToGetDataLimitsLimitOutputWithContext(ctx context.Context) GetDataLimitsLimitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDataLimitsLimitOutput)
}

// GetDataLimitsLimitArrayInput is an input type that accepts GetDataLimitsLimitArray and GetDataLimitsLimitArrayOutput values.
// You can construct a concrete instance of `GetDataLimitsLimitArrayInput` via:
//
//	GetDataLimitsLimitArray{ GetDataLimitsLimitArgs{...} }
type GetDataLimitsLimitArrayInput interface {
	pulumi.Input

	ToGetDataLimitsLimitArrayOutput() GetDataLimitsLimitArrayOutput
	ToGetDataLimitsLimitArrayOutputWithContext(context.Context) GetDataLimitsLimitArrayOutput
}

type GetDataLimitsLimitArray []GetDataLimitsLimitInput

func (GetDataLimitsLimitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDataLimitsLimit)(nil)).Elem()
}

func (i GetDataLimitsLimitArray) ToGetDataLimitsLimitArrayOutput() GetDataLimitsLimitArrayOutput {
	return i.ToGetDataLimitsLimitArrayOutputWithContext(context.Background())
}

func (i GetDataLimitsLimitArray) ToGetDataLimitsLimitArrayOutputWithContext(ctx context.Context) GetDataLimitsLimitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDataLimitsLimitArrayOutput)
}

type GetDataLimitsLimitOutput struct{ *pulumi.OutputState }

func (GetDataLimitsLimitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataLimitsLimit)(nil)).Elem()
}

func (o GetDataLimitsLimitOutput) ToGetDataLimitsLimitOutput() GetDataLimitsLimitOutput {
	return o
}

func (o GetDataLimitsLimitOutput) ToGetDataLimitsLimitOutputWithContext(ctx context.Context) GetDataLimitsLimitOutput {
	return o
}

// Whether to enable the log auditing feature.
func (o GetDataLimitsLimitOutput) AuditStatus() pulumi.IntOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) int { return v.AuditStatus }).(pulumi.IntOutput)
}

// The status of the connectivity test between the data asset and SDDP.
func (o GetDataLimitsLimitOutput) CheckStatus() pulumi.IntOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) int { return v.CheckStatus }).(pulumi.IntOutput)
}

// The first ID of the resource.
func (o GetDataLimitsLimitOutput) DataLimitId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.DataLimitId }).(pulumi.StringOutput)
}

// The type of the database.
func (o GetDataLimitsLimitOutput) EngineType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.EngineType }).(pulumi.StringOutput)
}

// The ID of the Data Limit.
func (o GetDataLimitsLimitOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the service to which the data asset belongs.
func (o GetDataLimitsLimitOutput) LocalName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.LocalName }).(pulumi.StringOutput)
}

// The retention period of raw logs after you enable the log auditing feature.
func (o GetDataLimitsLimitOutput) LogStoreDay() pulumi.IntOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) int { return v.LogStoreDay }).(pulumi.IntOutput)
}

// The ID of the data asset.
func (o GetDataLimitsLimitOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.ParentId }).(pulumi.StringOutput)
}

// The port that is used to connect to the database.
func (o GetDataLimitsLimitOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) int { return v.Port }).(pulumi.IntOutput)
}

// The type of the service to which the data asset belongs.
func (o GetDataLimitsLimitOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.ResourceType }).(pulumi.StringOutput)
}

// The name of the user who owns the data asset.
func (o GetDataLimitsLimitOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataLimitsLimit) string { return v.UserName }).(pulumi.StringOutput)
}

type GetDataLimitsLimitArrayOutput struct{ *pulumi.OutputState }

func (GetDataLimitsLimitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDataLimitsLimit)(nil)).Elem()
}

func (o GetDataLimitsLimitArrayOutput) ToGetDataLimitsLimitArrayOutput() GetDataLimitsLimitArrayOutput {
	return o
}

func (o GetDataLimitsLimitArrayOutput) ToGetDataLimitsLimitArrayOutputWithContext(ctx context.Context) GetDataLimitsLimitArrayOutput {
	return o
}

func (o GetDataLimitsLimitArrayOutput) Index(i pulumi.IntInput) GetDataLimitsLimitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDataLimitsLimit {
		return vs[0].([]GetDataLimitsLimit)[vs[1].(int)]
	}).(GetDataLimitsLimitOutput)
}

type GetInstancesInstance struct {
	// Whether the required RAM authorization is configured.
	Authed bool   `pulumi:"authed"`
	Id     string `pulumi:"id"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The number of instances.
	InstanceNum string `pulumi:"instanceNum"`
	// Whether the authorized MaxCompute (ODPS) assets.
	OdpsSet bool `pulumi:"odpsSet"`
	// Whether the authorized oss assets.
	OssBucketSet bool `pulumi:"ossBucketSet"`
	// The OSS size of the instance.
	OssSize string `pulumi:"ossSize"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType string `pulumi:"paymentType"`
	// Whether the authorized rds assets.
	RdsSet bool `pulumi:"rdsSet"`
	// The status of the resource.
	Status string `pulumi:"status"`
}

// GetInstancesInstanceInput is an input type that accepts GetInstancesInstanceArgs and GetInstancesInstanceOutput values.
// You can construct a concrete instance of `GetInstancesInstanceInput` via:
//
//	GetInstancesInstanceArgs{...}
type GetInstancesInstanceInput interface {
	pulumi.Input

	ToGetInstancesInstanceOutput() GetInstancesInstanceOutput
	ToGetInstancesInstanceOutputWithContext(context.Context) GetInstancesInstanceOutput
}

type GetInstancesInstanceArgs struct {
	// Whether the required RAM authorization is configured.
	Authed pulumi.BoolInput   `pulumi:"authed"`
	Id     pulumi.StringInput `pulumi:"id"`
	// The ID of the instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The number of instances.
	InstanceNum pulumi.StringInput `pulumi:"instanceNum"`
	// Whether the authorized MaxCompute (ODPS) assets.
	OdpsSet pulumi.BoolInput `pulumi:"odpsSet"`
	// Whether the authorized oss assets.
	OssBucketSet pulumi.BoolInput `pulumi:"ossBucketSet"`
	// The OSS size of the instance.
	OssSize pulumi.StringInput `pulumi:"ossSize"`
	// The payment type of the resource. Valid values: `Subscription`.
	PaymentType pulumi.StringInput `pulumi:"paymentType"`
	// Whether the authorized rds assets.
	RdsSet pulumi.BoolInput `pulumi:"rdsSet"`
	// The status of the resource.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetInstancesInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesInstance)(nil)).Elem()
}

func (i GetInstancesInstanceArgs) ToGetInstancesInstanceOutput() GetInstancesInstanceOutput {
	return i.ToGetInstancesInstanceOutputWithContext(context.Background())
}

func (i GetInstancesInstanceArgs) ToGetInstancesInstanceOutputWithContext(ctx context.Context) GetInstancesInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesInstanceOutput)
}

// GetInstancesInstanceArrayInput is an input type that accepts GetInstancesInstanceArray and GetInstancesInstanceArrayOutput values.
// You can construct a concrete instance of `GetInstancesInstanceArrayInput` via:
//
//	GetInstancesInstanceArray{ GetInstancesInstanceArgs{...} }
type GetInstancesInstanceArrayInput interface {
	pulumi.Input

	ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput
	ToGetInstancesInstanceArrayOutputWithContext(context.Context) GetInstancesInstanceArrayOutput
}

type GetInstancesInstanceArray []GetInstancesInstanceInput

func (GetInstancesInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesInstance)(nil)).Elem()
}

func (i GetInstancesInstanceArray) ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput {
	return i.ToGetInstancesInstanceArrayOutputWithContext(context.Background())
}

func (i GetInstancesInstanceArray) ToGetInstancesInstanceArrayOutputWithContext(ctx context.Context) GetInstancesInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInstancesInstanceArrayOutput)
}

type GetInstancesInstanceOutput struct{ *pulumi.OutputState }

func (GetInstancesInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesInstance)(nil)).Elem()
}

func (o GetInstancesInstanceOutput) ToGetInstancesInstanceOutput() GetInstancesInstanceOutput {
	return o
}

func (o GetInstancesInstanceOutput) ToGetInstancesInstanceOutputWithContext(ctx context.Context) GetInstancesInstanceOutput {
	return o
}

// Whether the required RAM authorization is configured.
func (o GetInstancesInstanceOutput) Authed() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstancesInstance) bool { return v.Authed }).(pulumi.BoolOutput)
}

func (o GetInstancesInstanceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o GetInstancesInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The number of instances.
func (o GetInstancesInstanceOutput) InstanceNum() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.InstanceNum }).(pulumi.StringOutput)
}

// Whether the authorized MaxCompute (ODPS) assets.
func (o GetInstancesInstanceOutput) OdpsSet() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstancesInstance) bool { return v.OdpsSet }).(pulumi.BoolOutput)
}

// Whether the authorized oss assets.
func (o GetInstancesInstanceOutput) OssBucketSet() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstancesInstance) bool { return v.OssBucketSet }).(pulumi.BoolOutput)
}

// The OSS size of the instance.
func (o GetInstancesInstanceOutput) OssSize() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.OssSize }).(pulumi.StringOutput)
}

// The payment type of the resource. Valid values: `Subscription`.
func (o GetInstancesInstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.PaymentType }).(pulumi.StringOutput)
}

// Whether the authorized rds assets.
func (o GetInstancesInstanceOutput) RdsSet() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstancesInstance) bool { return v.RdsSet }).(pulumi.BoolOutput)
}

// The status of the resource.
func (o GetInstancesInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesInstance) string { return v.Status }).(pulumi.StringOutput)
}

type GetInstancesInstanceArrayOutput struct{ *pulumi.OutputState }

func (GetInstancesInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInstancesInstance)(nil)).Elem()
}

func (o GetInstancesInstanceArrayOutput) ToGetInstancesInstanceArrayOutput() GetInstancesInstanceArrayOutput {
	return o
}

func (o GetInstancesInstanceArrayOutput) ToGetInstancesInstanceArrayOutputWithContext(ctx context.Context) GetInstancesInstanceArrayOutput {
	return o
}

func (o GetInstancesInstanceArrayOutput) Index(i pulumi.IntInput) GetInstancesInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInstancesInstance {
		return vs[0].([]GetInstancesInstance)[vs[1].(int)]
	}).(GetInstancesInstanceOutput)
}

type GetRulesRule struct {
	// Sensitive Data Identification Rules for the Type of.
	Category int `pulumi:"category"`
	// Sensitive Data Identification Rules Belongs Type Name.
	CategoryName string `pulumi:"categoryName"`
	// Sensitive Data Identification Rules the Content.
	Content string `pulumi:"content"`
	// The Content Classification.
	ContentCategory string `pulumi:"contentCategory"`
	// Sensitive Data Identification Rules the Creation Time of the Number of Milliseconds.
	CreateTime string `pulumi:"createTime"`
	// Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
	CustomType int `pulumi:"customType"`
	// Sensitive Data Identification a Description of the Rule Information.
	Description string `pulumi:"description"`
	// Sensitive Data Identification Rules, Founder of Account Display Name.
	DisplayName string `pulumi:"displayName"`
	// Sensitive Data Identification Rules to the Modified Time of the Number of Milliseconds.
	GmtModified string `pulumi:"gmtModified"`
	// The ID of the Rule.
	Id string `pulumi:"id"`
	// Sensitive Data Identification Rules, Founder Of Account Login.
	LoginName string `pulumi:"loginName"`
	// The Primary Key.
	MajorKey string `pulumi:"majorKey"`
	// The name of rule.
	Name string `pulumi:"name"`
	// Product Code.
	ProductCode string `pulumi:"productCode"`
	// Product ID.
	ProductId string `pulumi:"productId"`
	// Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
	RiskLevelId string `pulumi:"riskLevelId"`
	// Sensitive Data Identification Rules the Risk Level of. S1: Weak Risk Level S2: Moderate Risk Level S3: High Risk Level S4: the Highest Risk Level.
	RiskLevelName string `pulumi:"riskLevelName"`
	// The first ID of the resource.
	RuleId string `pulumi:"ruleId"`
	// Triggered the Alarm Conditions.
	StatExpress string `pulumi:"statExpress"`
	// Sensitive Data Identification Rules Detection State of.
	Status int `pulumi:"status"`
	// The Target.
	Target string `pulumi:"target"`
	// The User ID.
	UserId string `pulumi:"userId"`
	// The Level of Risk.
	WarnLevel int `pulumi:"warnLevel"`
}

// GetRulesRuleInput is an input type that accepts GetRulesRuleArgs and GetRulesRuleOutput values.
// You can construct a concrete instance of `GetRulesRuleInput` via:
//
//	GetRulesRuleArgs{...}
type GetRulesRuleInput interface {
	pulumi.Input

	ToGetRulesRuleOutput() GetRulesRuleOutput
	ToGetRulesRuleOutputWithContext(context.Context) GetRulesRuleOutput
}

type GetRulesRuleArgs struct {
	// Sensitive Data Identification Rules for the Type of.
	Category pulumi.IntInput `pulumi:"category"`
	// Sensitive Data Identification Rules Belongs Type Name.
	CategoryName pulumi.StringInput `pulumi:"categoryName"`
	// Sensitive Data Identification Rules the Content.
	Content pulumi.StringInput `pulumi:"content"`
	// The Content Classification.
	ContentCategory pulumi.StringInput `pulumi:"contentCategory"`
	// Sensitive Data Identification Rules the Creation Time of the Number of Milliseconds.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
	CustomType pulumi.IntInput `pulumi:"customType"`
	// Sensitive Data Identification a Description of the Rule Information.
	Description pulumi.StringInput `pulumi:"description"`
	// Sensitive Data Identification Rules, Founder of Account Display Name.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Sensitive Data Identification Rules to the Modified Time of the Number of Milliseconds.
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The ID of the Rule.
	Id pulumi.StringInput `pulumi:"id"`
	// Sensitive Data Identification Rules, Founder Of Account Login.
	LoginName pulumi.StringInput `pulumi:"loginName"`
	// The Primary Key.
	MajorKey pulumi.StringInput `pulumi:"majorKey"`
	// The name of rule.
	Name pulumi.StringInput `pulumi:"name"`
	// Product Code.
	ProductCode pulumi.StringInput `pulumi:"productCode"`
	// Product ID.
	ProductId pulumi.StringInput `pulumi:"productId"`
	// Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
	RiskLevelId pulumi.StringInput `pulumi:"riskLevelId"`
	// Sensitive Data Identification Rules the Risk Level of. S1: Weak Risk Level S2: Moderate Risk Level S3: High Risk Level S4: the Highest Risk Level.
	RiskLevelName pulumi.StringInput `pulumi:"riskLevelName"`
	// The first ID of the resource.
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// Triggered the Alarm Conditions.
	StatExpress pulumi.StringInput `pulumi:"statExpress"`
	// Sensitive Data Identification Rules Detection State of.
	Status pulumi.IntInput `pulumi:"status"`
	// The Target.
	Target pulumi.StringInput `pulumi:"target"`
	// The User ID.
	UserId pulumi.StringInput `pulumi:"userId"`
	// The Level of Risk.
	WarnLevel pulumi.IntInput `pulumi:"warnLevel"`
}

func (GetRulesRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesRule)(nil)).Elem()
}

func (i GetRulesRuleArgs) ToGetRulesRuleOutput() GetRulesRuleOutput {
	return i.ToGetRulesRuleOutputWithContext(context.Background())
}

func (i GetRulesRuleArgs) ToGetRulesRuleOutputWithContext(ctx context.Context) GetRulesRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRulesRuleOutput)
}

// GetRulesRuleArrayInput is an input type that accepts GetRulesRuleArray and GetRulesRuleArrayOutput values.
// You can construct a concrete instance of `GetRulesRuleArrayInput` via:
//
//	GetRulesRuleArray{ GetRulesRuleArgs{...} }
type GetRulesRuleArrayInput interface {
	pulumi.Input

	ToGetRulesRuleArrayOutput() GetRulesRuleArrayOutput
	ToGetRulesRuleArrayOutputWithContext(context.Context) GetRulesRuleArrayOutput
}

type GetRulesRuleArray []GetRulesRuleInput

func (GetRulesRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRulesRule)(nil)).Elem()
}

func (i GetRulesRuleArray) ToGetRulesRuleArrayOutput() GetRulesRuleArrayOutput {
	return i.ToGetRulesRuleArrayOutputWithContext(context.Background())
}

func (i GetRulesRuleArray) ToGetRulesRuleArrayOutputWithContext(ctx context.Context) GetRulesRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRulesRuleArrayOutput)
}

type GetRulesRuleOutput struct{ *pulumi.OutputState }

func (GetRulesRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesRule)(nil)).Elem()
}

func (o GetRulesRuleOutput) ToGetRulesRuleOutput() GetRulesRuleOutput {
	return o
}

func (o GetRulesRuleOutput) ToGetRulesRuleOutputWithContext(ctx context.Context) GetRulesRuleOutput {
	return o
}

// Sensitive Data Identification Rules for the Type of.
func (o GetRulesRuleOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesRule) int { return v.Category }).(pulumi.IntOutput)
}

// Sensitive Data Identification Rules Belongs Type Name.
func (o GetRulesRuleOutput) CategoryName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.CategoryName }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules the Content.
func (o GetRulesRuleOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.Content }).(pulumi.StringOutput)
}

// The Content Classification.
func (o GetRulesRuleOutput) ContentCategory() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.ContentCategory }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules the Creation Time of the Number of Milliseconds.
func (o GetRulesRuleOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules of Type. 0: the Built-in 1: The User-Defined.
func (o GetRulesRuleOutput) CustomType() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesRule) int { return v.CustomType }).(pulumi.IntOutput)
}

// Sensitive Data Identification a Description of the Rule Information.
func (o GetRulesRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.Description }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules, Founder of Account Display Name.
func (o GetRulesRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules to the Modified Time of the Number of Milliseconds.
func (o GetRulesRuleOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The ID of the Rule.
func (o GetRulesRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.Id }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules, Founder Of Account Login.
func (o GetRulesRuleOutput) LoginName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.LoginName }).(pulumi.StringOutput)
}

// The Primary Key.
func (o GetRulesRuleOutput) MajorKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.MajorKey }).(pulumi.StringOutput)
}

// The name of rule.
func (o GetRulesRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.Name }).(pulumi.StringOutput)
}

// Product Code.
func (o GetRulesRuleOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.ProductCode }).(pulumi.StringOutput)
}

// Product ID.
func (o GetRulesRuleOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.ProductId }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules of Risk Level ID. Valid values:1:S1, Weak Risk Level. 2:S2, Medium Risk Level. 3:S3 High Risk Level. 4:S4, the Highest Risk Level.
func (o GetRulesRuleOutput) RiskLevelId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.RiskLevelId }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules the Risk Level of. S1: Weak Risk Level S2: Moderate Risk Level S3: High Risk Level S4: the Highest Risk Level.
func (o GetRulesRuleOutput) RiskLevelName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.RiskLevelName }).(pulumi.StringOutput)
}

// The first ID of the resource.
func (o GetRulesRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.RuleId }).(pulumi.StringOutput)
}

// Triggered the Alarm Conditions.
func (o GetRulesRuleOutput) StatExpress() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.StatExpress }).(pulumi.StringOutput)
}

// Sensitive Data Identification Rules Detection State of.
func (o GetRulesRuleOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesRule) int { return v.Status }).(pulumi.IntOutput)
}

// The Target.
func (o GetRulesRuleOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.Target }).(pulumi.StringOutput)
}

// The User ID.
func (o GetRulesRuleOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesRule) string { return v.UserId }).(pulumi.StringOutput)
}

// The Level of Risk.
func (o GetRulesRuleOutput) WarnLevel() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesRule) int { return v.WarnLevel }).(pulumi.IntOutput)
}

type GetRulesRuleArrayOutput struct{ *pulumi.OutputState }

func (GetRulesRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRulesRule)(nil)).Elem()
}

func (o GetRulesRuleArrayOutput) ToGetRulesRuleArrayOutput() GetRulesRuleArrayOutput {
	return o
}

func (o GetRulesRuleArrayOutput) ToGetRulesRuleArrayOutputWithContext(ctx context.Context) GetRulesRuleArrayOutput {
	return o
}

func (o GetRulesRuleArrayOutput) Index(i pulumi.IntInput) GetRulesRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetRulesRule {
		return vs[0].([]GetRulesRule)[vs[1].(int)]
	}).(GetRulesRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetConfigsConfigInput)(nil)).Elem(), GetConfigsConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetConfigsConfigArrayInput)(nil)).Elem(), GetConfigsConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDataLimitsLimitInput)(nil)).Elem(), GetDataLimitsLimitArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDataLimitsLimitArrayInput)(nil)).Elem(), GetDataLimitsLimitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstancesInstanceInput)(nil)).Elem(), GetInstancesInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInstancesInstanceArrayInput)(nil)).Elem(), GetInstancesInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRulesRuleInput)(nil)).Elem(), GetRulesRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRulesRuleArrayInput)(nil)).Elem(), GetRulesRuleArray{})
	pulumi.RegisterOutputType(GetConfigsConfigOutput{})
	pulumi.RegisterOutputType(GetConfigsConfigArrayOutput{})
	pulumi.RegisterOutputType(GetDataLimitsLimitOutput{})
	pulumi.RegisterOutputType(GetDataLimitsLimitArrayOutput{})
	pulumi.RegisterOutputType(GetInstancesInstanceOutput{})
	pulumi.RegisterOutputType(GetInstancesInstanceArrayOutput{})
	pulumi.RegisterOutputType(GetRulesRuleOutput{})
	pulumi.RegisterOutputType(GetRulesRuleArrayOutput{})
}
