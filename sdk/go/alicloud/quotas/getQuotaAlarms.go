// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Quotas Quota Alarms of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.116.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := quotas.GetQuotaAlarms(ctx, &quotas.GetQuotaAlarmsArgs{
//				Ids: []string{
//					"5VR90-421F886-81E9-xxx",
//				},
//				NameRegex: pulumi.StringRef("tf-testAcc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstQuotasQuotaAlarmId", example.Alarms[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetQuotaAlarms(ctx *pulumi.Context, args *GetQuotaAlarmsArgs, opts ...pulumi.InvokeOption) (*GetQuotaAlarmsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetQuotaAlarmsResult
	err := ctx.Invoke("alicloud:quotas/getQuotaAlarms:getQuotaAlarms", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQuotaAlarms.
type GetQuotaAlarmsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Quota Alarm IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Quota Alarm name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The Product Code.
	ProductCode *string `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName *string `pulumi:"quotaAlarmName"`
	// The Quota Dimensions.
	QuotaDimensions []GetQuotaAlarmsQuotaDimension `pulumi:"quotaDimensions"`
}

// A collection of values returned by getQuotaAlarms.
type GetQuotaAlarmsResult struct {
	Alarms        []GetQuotaAlarmsAlarm `pulumi:"alarms"`
	EnableDetails *bool                 `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                         `pulumi:"id"`
	Ids             []string                       `pulumi:"ids"`
	NameRegex       *string                        `pulumi:"nameRegex"`
	Names           []string                       `pulumi:"names"`
	OutputFile      *string                        `pulumi:"outputFile"`
	ProductCode     *string                        `pulumi:"productCode"`
	QuotaActionCode *string                        `pulumi:"quotaActionCode"`
	QuotaAlarmName  *string                        `pulumi:"quotaAlarmName"`
	QuotaDimensions []GetQuotaAlarmsQuotaDimension `pulumi:"quotaDimensions"`
}

func GetQuotaAlarmsOutput(ctx *pulumi.Context, args GetQuotaAlarmsOutputArgs, opts ...pulumi.InvokeOption) GetQuotaAlarmsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetQuotaAlarmsResult, error) {
			args := v.(GetQuotaAlarmsArgs)
			r, err := GetQuotaAlarms(ctx, &args, opts...)
			var s GetQuotaAlarmsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetQuotaAlarmsResultOutput)
}

// A collection of arguments for invoking getQuotaAlarms.
type GetQuotaAlarmsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Quota Alarm IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Quota Alarm name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Product Code.
	ProductCode pulumi.StringPtrInput `pulumi:"productCode"`
	// The Quota Action Code.
	QuotaActionCode pulumi.StringPtrInput `pulumi:"quotaActionCode"`
	// The name of Quota Alarm.
	QuotaAlarmName pulumi.StringPtrInput `pulumi:"quotaAlarmName"`
	// The Quota Dimensions.
	QuotaDimensions GetQuotaAlarmsQuotaDimensionArrayInput `pulumi:"quotaDimensions"`
}

func (GetQuotaAlarmsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQuotaAlarmsArgs)(nil)).Elem()
}

// A collection of values returned by getQuotaAlarms.
type GetQuotaAlarmsResultOutput struct{ *pulumi.OutputState }

func (GetQuotaAlarmsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQuotaAlarmsResult)(nil)).Elem()
}

func (o GetQuotaAlarmsResultOutput) ToGetQuotaAlarmsResultOutput() GetQuotaAlarmsResultOutput {
	return o
}

func (o GetQuotaAlarmsResultOutput) ToGetQuotaAlarmsResultOutputWithContext(ctx context.Context) GetQuotaAlarmsResultOutput {
	return o
}

func (o GetQuotaAlarmsResultOutput) Alarms() GetQuotaAlarmsAlarmArrayOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) []GetQuotaAlarmsAlarm { return v.Alarms }).(GetQuotaAlarmsAlarmArrayOutput)
}

func (o GetQuotaAlarmsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetQuotaAlarmsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetQuotaAlarmsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetQuotaAlarmsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetQuotaAlarmsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetQuotaAlarmsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetQuotaAlarmsResultOutput) ProductCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *string { return v.ProductCode }).(pulumi.StringPtrOutput)
}

func (o GetQuotaAlarmsResultOutput) QuotaActionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *string { return v.QuotaActionCode }).(pulumi.StringPtrOutput)
}

func (o GetQuotaAlarmsResultOutput) QuotaAlarmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) *string { return v.QuotaAlarmName }).(pulumi.StringPtrOutput)
}

func (o GetQuotaAlarmsResultOutput) QuotaDimensions() GetQuotaAlarmsQuotaDimensionArrayOutput {
	return o.ApplyT(func(v GetQuotaAlarmsResult) []GetQuotaAlarmsQuotaDimension { return v.QuotaDimensions }).(GetQuotaAlarmsQuotaDimensionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetQuotaAlarmsResultOutput{})
}
