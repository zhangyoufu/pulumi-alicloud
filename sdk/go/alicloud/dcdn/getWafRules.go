// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Dcdn Waf Rule available to the user.[What is Waf Rule](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchcreatedcdnwafrules)
//
// > **NOTE:** Available since v1.201.0.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _default, err := dcdn.GetWafRules(ctx, &dcdn.GetWafRulesArgs{
// Ids: interface{}{
// defaultAlicloudDcdnWafRule.Id,
// },
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("alicloudDcdnWafRuleExampleId", _default.WafRules[0].Id)
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
func GetWafRules(ctx *pulumi.Context, args *GetWafRulesArgs, opts ...pulumi.InvokeOption) (*GetWafRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWafRulesResult
	err := ctx.Invoke("alicloud:dcdn/getWafRules:getWafRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafRules.
type GetWafRulesArgs struct {
	// A list of Waf Rule IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The query conditions. The value is a string in the JSON format.
	QueryArgs *string `pulumi:"queryArgs"`
}

// A collection of values returned by getWafRules.
type GetWafRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	QueryArgs  *string  `pulumi:"queryArgs"`
	// A list of Waf Rule Entries. Each element contains the following attributes:
	WafRules []GetWafRulesWafRule `pulumi:"wafRules"`
}

func GetWafRulesOutput(ctx *pulumi.Context, args GetWafRulesOutputArgs, opts ...pulumi.InvokeOption) GetWafRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWafRulesResult, error) {
			args := v.(GetWafRulesArgs)
			r, err := GetWafRules(ctx, &args, opts...)
			var s GetWafRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWafRulesResultOutput)
}

// A collection of arguments for invoking getWafRules.
type GetWafRulesOutputArgs struct {
	// A list of Waf Rule IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The query conditions. The value is a string in the JSON format.
	QueryArgs pulumi.StringPtrInput `pulumi:"queryArgs"`
}

func (GetWafRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafRulesArgs)(nil)).Elem()
}

// A collection of values returned by getWafRules.
type GetWafRulesResultOutput struct{ *pulumi.OutputState }

func (GetWafRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafRulesResult)(nil)).Elem()
}

func (o GetWafRulesResultOutput) ToGetWafRulesResultOutput() GetWafRulesResultOutput {
	return o
}

func (o GetWafRulesResultOutput) ToGetWafRulesResultOutputWithContext(ctx context.Context) GetWafRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWafRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetWafRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetWafRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetWafRulesResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWafRulesResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetWafRulesResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetWafRulesResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetWafRulesResultOutput) QueryArgs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafRulesResult) *string { return v.QueryArgs }).(pulumi.StringPtrOutput)
}

// A list of Waf Rule Entries. Each element contains the following attributes:
func (o GetWafRulesResultOutput) WafRules() GetWafRulesWafRuleArrayOutput {
	return o.ApplyT(func(v GetWafRulesResult) []GetWafRulesWafRule { return v.WafRules }).(GetWafRulesWafRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafRulesResultOutput{})
}
