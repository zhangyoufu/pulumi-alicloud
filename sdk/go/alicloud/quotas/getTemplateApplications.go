// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Quotas Template Applications available to the user.[What is Template Applications](https://www.alibabacloud.com/help/en/quota-center/developer-reference/api-quotas-2020-05-10-createquotaapplicationsfortemplate)
//
// > **NOTE:** Available since v1.214.0.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/quotas"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := resourcemanager.GetAccounts(ctx, &resourcemanager.GetAccountsArgs{
//				Status: pulumi.StringRef("CreateSuccess"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultTemplateApplications, err := quotas.NewTemplateApplications(ctx, "default", &quotas.TemplateApplicationsArgs{
//				QuotaActionCode: pulumi.String("vpc_whitelist/ha_vip_whitelist"),
//				ProductCode:     pulumi.String("vpc"),
//				QuotaCategory:   pulumi.String("FlowControl"),
//				AliyunUids: pulumi.StringArray{
//					pulumi.String(_default.Ids[0]),
//				},
//				DesireValue: pulumi.Float64(6),
//				NoticeType:  pulumi.Int(0),
//				EnvLanguage: pulumi.String("zh"),
//				Reason:      pulumi.String("example"),
//				Dimensions: quotas.TemplateApplicationsDimensionArray{
//					&quotas.TemplateApplicationsDimensionArgs{
//						Key:   pulumi.String("apiName"),
//						Value: pulumi.String("GetProductQuotaDimension"),
//					},
//					&quotas.TemplateApplicationsDimensionArgs{
//						Key:   pulumi.String("apiVersion"),
//						Value: pulumi.String("2020-05-10"),
//					},
//					&quotas.TemplateApplicationsDimensionArgs{
//						Key:   pulumi.String("regionId"),
//						Value: pulumi.String("cn-hangzhou"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultGetTemplateApplications := quotas.LookupTemplateApplicationsOutput(ctx, quotas.GetTemplateApplicationsOutputArgs{
//				Ids: pulumi.StringArray{
//					defaultTemplateApplications.ID(),
//				},
//				ProductCode:     pulumi.String("vpc"),
//				QuotaActionCode: pulumi.String("vpc_whitelist/ha_vip_whitelist"),
//				QuotaCategory:   pulumi.String("FlowControl"),
//			}, nil)
//			ctx.Export("alicloudQuotasTemplateApplicationsExampleId", defaultGetTemplateApplications.ApplyT(func(defaultGetTemplateApplications quotas.GetTemplateApplicationsResult) (*string, error) {
//				return &defaultGetTemplateApplications.Applications[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupTemplateApplications(ctx *pulumi.Context, args *LookupTemplateApplicationsArgs, opts ...pulumi.InvokeOption) (*LookupTemplateApplicationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateApplicationsResult
	err := ctx.Invoke("alicloud:quotas/getTemplateApplications:getTemplateApplications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplateApplications.
type LookupTemplateApplicationsArgs struct {
	// The ID of the quota application batch.
	BatchQuotaApplicationId *string `pulumi:"batchQuotaApplicationId"`
	// A list of Template Applications IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode *string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
	QuotaCategory *string `pulumi:"quotaCategory"`
}

// A collection of values returned by getTemplateApplications.
type LookupTemplateApplicationsResult struct {
	// A list of Template Applications Entries. Each element contains the following attributes:
	Applications []GetTemplateApplicationsApplication `pulumi:"applications"`
	// The ID of the quota application batch.
	BatchQuotaApplicationId *string `pulumi:"batchQuotaApplicationId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Template Applications IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode *string `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode *string `pulumi:"quotaActionCode"`
	// The quota type. Value:-CommonQuota (default): Generic quota.-FlowControl:API rate quota.-WhiteListLabel: Equity quota.
	QuotaCategory *string `pulumi:"quotaCategory"`
}

func LookupTemplateApplicationsOutput(ctx *pulumi.Context, args LookupTemplateApplicationsOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateApplicationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateApplicationsResult, error) {
			args := v.(LookupTemplateApplicationsArgs)
			r, err := LookupTemplateApplications(ctx, &args, opts...)
			var s LookupTemplateApplicationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateApplicationsResultOutput)
}

// A collection of arguments for invoking getTemplateApplications.
type LookupTemplateApplicationsOutputArgs struct {
	// The ID of the quota application batch.
	BatchQuotaApplicationId pulumi.StringPtrInput `pulumi:"batchQuotaApplicationId"`
	// A list of Template Applications IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
	ProductCode pulumi.StringPtrInput `pulumi:"productCode"`
	// The quota ID.
	QuotaActionCode pulumi.StringPtrInput `pulumi:"quotaActionCode"`
	// The quota type. Value: `CommonQuota`, `FlowControl` and `WhiteListLabel`.
	QuotaCategory pulumi.StringPtrInput `pulumi:"quotaCategory"`
}

func (LookupTemplateApplicationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateApplicationsArgs)(nil)).Elem()
}

// A collection of values returned by getTemplateApplications.
type LookupTemplateApplicationsResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateApplicationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateApplicationsResult)(nil)).Elem()
}

func (o LookupTemplateApplicationsResultOutput) ToLookupTemplateApplicationsResultOutput() LookupTemplateApplicationsResultOutput {
	return o
}

func (o LookupTemplateApplicationsResultOutput) ToLookupTemplateApplicationsResultOutputWithContext(ctx context.Context) LookupTemplateApplicationsResultOutput {
	return o
}

// A list of Template Applications Entries. Each element contains the following attributes:
func (o LookupTemplateApplicationsResultOutput) Applications() GetTemplateApplicationsApplicationArrayOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) []GetTemplateApplicationsApplication { return v.Applications }).(GetTemplateApplicationsApplicationArrayOutput)
}

// The ID of the quota application batch.
func (o LookupTemplateApplicationsResultOutput) BatchQuotaApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) *string { return v.BatchQuotaApplicationId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTemplateApplicationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Template Applications IDs.
func (o LookupTemplateApplicationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupTemplateApplicationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Cloud service name abbreviation.> For more information about cloud services that support quota centers, see Cloud services that support quota centers.
func (o LookupTemplateApplicationsResultOutput) ProductCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) *string { return v.ProductCode }).(pulumi.StringPtrOutput)
}

// The quota ID.
func (o LookupTemplateApplicationsResultOutput) QuotaActionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) *string { return v.QuotaActionCode }).(pulumi.StringPtrOutput)
}

// The quota type. Value:-CommonQuota (default): Generic quota.-FlowControl:API rate quota.-WhiteListLabel: Equity quota.
func (o LookupTemplateApplicationsResultOutput) QuotaCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateApplicationsResult) *string { return v.QuotaCategory }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateApplicationsResultOutput{})
}
