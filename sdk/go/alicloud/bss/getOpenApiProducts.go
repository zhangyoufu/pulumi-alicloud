// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Bss Open Api Product available to the user.[What is Product](https://www.alibabacloud.com/help/en/bss-openapi/latest/api-bssopenapi-2017-12-14-queryproductlist)
//
// > **NOTE:** Available in 1.195.0+
func GetOpenApiProducts(ctx *pulumi.Context, args *GetOpenApiProductsArgs, opts ...pulumi.InvokeOption) (*GetOpenApiProductsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOpenApiProductsResult
	err := ctx.Invoke("alicloud:bss/getOpenApiProducts:getOpenApiProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOpenApiProducts.
type GetOpenApiProductsArgs struct {
	// A list of product IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Product name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getOpenApiProducts.
type GetOpenApiProductsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of name of Products.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of Product Entries. Each element contains the following attributes:
	Products []GetOpenApiProductsProduct `pulumi:"products"`
}

func GetOpenApiProductsOutput(ctx *pulumi.Context, args GetOpenApiProductsOutputArgs, opts ...pulumi.InvokeOption) GetOpenApiProductsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOpenApiProductsResult, error) {
			args := v.(GetOpenApiProductsArgs)
			r, err := GetOpenApiProducts(ctx, &args, opts...)
			var s GetOpenApiProductsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOpenApiProductsResultOutput)
}

// A collection of arguments for invoking getOpenApiProducts.
type GetOpenApiProductsOutputArgs struct {
	// A list of product IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Product name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetOpenApiProductsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpenApiProductsArgs)(nil)).Elem()
}

// A collection of values returned by getOpenApiProducts.
type GetOpenApiProductsResultOutput struct{ *pulumi.OutputState }

func (GetOpenApiProductsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOpenApiProductsResult)(nil)).Elem()
}

func (o GetOpenApiProductsResultOutput) ToGetOpenApiProductsResultOutput() GetOpenApiProductsResultOutput {
	return o
}

func (o GetOpenApiProductsResultOutput) ToGetOpenApiProductsResultOutputWithContext(ctx context.Context) GetOpenApiProductsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOpenApiProductsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOpenApiProductsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetOpenApiProductsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of name of Products.
func (o GetOpenApiProductsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetOpenApiProductsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of Product Entries. Each element contains the following attributes:
func (o GetOpenApiProductsResultOutput) Products() GetOpenApiProductsProductArrayOutput {
	return o.ApplyT(func(v GetOpenApiProductsResult) []GetOpenApiProductsProduct { return v.Products }).(GetOpenApiProductsProductArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOpenApiProductsResultOutput{})
}
