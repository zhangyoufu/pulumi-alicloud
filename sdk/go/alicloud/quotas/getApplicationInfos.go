// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApplicationInfos(ctx *pulumi.Context, args *GetApplicationInfosArgs, opts ...pulumi.InvokeOption) (*GetApplicationInfosResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationInfosResult
	err := ctx.Invoke("alicloud:quotas/getApplicationInfos:getApplicationInfos", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationInfos.
type GetApplicationInfosArgs struct {
	Dimensions      []GetApplicationInfosDimension `pulumi:"dimensions"`
	EnableDetails   *bool                          `pulumi:"enableDetails"`
	Ids             []string                       `pulumi:"ids"`
	KeyWord         *string                        `pulumi:"keyWord"`
	OutputFile      *string                        `pulumi:"outputFile"`
	ProductCode     string                         `pulumi:"productCode"`
	QuotaActionCode *string                        `pulumi:"quotaActionCode"`
	QuotaCategory   *string                        `pulumi:"quotaCategory"`
	Status          *string                        `pulumi:"status"`
}

// A collection of values returned by getApplicationInfos.
type GetApplicationInfosResult struct {
	Applications  []GetApplicationInfosApplication `pulumi:"applications"`
	Dimensions    []GetApplicationInfosDimension   `pulumi:"dimensions"`
	EnableDetails *bool                            `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	Ids             []string `pulumi:"ids"`
	KeyWord         *string  `pulumi:"keyWord"`
	OutputFile      *string  `pulumi:"outputFile"`
	ProductCode     string   `pulumi:"productCode"`
	QuotaActionCode *string  `pulumi:"quotaActionCode"`
	QuotaCategory   *string  `pulumi:"quotaCategory"`
	Status          *string  `pulumi:"status"`
}

func GetApplicationInfosOutput(ctx *pulumi.Context, args GetApplicationInfosOutputArgs, opts ...pulumi.InvokeOption) GetApplicationInfosResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationInfosResult, error) {
			args := v.(GetApplicationInfosArgs)
			r, err := GetApplicationInfos(ctx, &args, opts...)
			var s GetApplicationInfosResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationInfosResultOutput)
}

// A collection of arguments for invoking getApplicationInfos.
type GetApplicationInfosOutputArgs struct {
	Dimensions      GetApplicationInfosDimensionArrayInput `pulumi:"dimensions"`
	EnableDetails   pulumi.BoolPtrInput                    `pulumi:"enableDetails"`
	Ids             pulumi.StringArrayInput                `pulumi:"ids"`
	KeyWord         pulumi.StringPtrInput                  `pulumi:"keyWord"`
	OutputFile      pulumi.StringPtrInput                  `pulumi:"outputFile"`
	ProductCode     pulumi.StringInput                     `pulumi:"productCode"`
	QuotaActionCode pulumi.StringPtrInput                  `pulumi:"quotaActionCode"`
	QuotaCategory   pulumi.StringPtrInput                  `pulumi:"quotaCategory"`
	Status          pulumi.StringPtrInput                  `pulumi:"status"`
}

func (GetApplicationInfosOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationInfosArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationInfos.
type GetApplicationInfosResultOutput struct{ *pulumi.OutputState }

func (GetApplicationInfosResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationInfosResult)(nil)).Elem()
}

func (o GetApplicationInfosResultOutput) ToGetApplicationInfosResultOutput() GetApplicationInfosResultOutput {
	return o
}

func (o GetApplicationInfosResultOutput) ToGetApplicationInfosResultOutputWithContext(ctx context.Context) GetApplicationInfosResultOutput {
	return o
}

func (o GetApplicationInfosResultOutput) Applications() GetApplicationInfosApplicationArrayOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) []GetApplicationInfosApplication { return v.Applications }).(GetApplicationInfosApplicationArrayOutput)
}

func (o GetApplicationInfosResultOutput) Dimensions() GetApplicationInfosDimensionArrayOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) []GetApplicationInfosDimension { return v.Dimensions }).(GetApplicationInfosDimensionArrayOutput)
}

func (o GetApplicationInfosResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationInfosResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApplicationInfosResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetApplicationInfosResultOutput) KeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *string { return v.KeyWord }).(pulumi.StringPtrOutput)
}

func (o GetApplicationInfosResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetApplicationInfosResultOutput) ProductCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) string { return v.ProductCode }).(pulumi.StringOutput)
}

func (o GetApplicationInfosResultOutput) QuotaActionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *string { return v.QuotaActionCode }).(pulumi.StringPtrOutput)
}

func (o GetApplicationInfosResultOutput) QuotaCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *string { return v.QuotaCategory }).(pulumi.StringPtrOutput)
}

func (o GetApplicationInfosResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationInfosResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationInfosResultOutput{})
}
