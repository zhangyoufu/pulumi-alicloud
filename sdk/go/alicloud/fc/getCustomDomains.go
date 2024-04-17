// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Function Compute custom domains of the current Alibaba Cloud user.
//
// > **NOTE:** Available in 1.98.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fc.GetCustomDomains(ctx, &fc.GetCustomDomainsArgs{
//				NameRegex: pulumi.StringRef("sample_fc_custom_domain"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstFcCustomDomainName", fcDomainsDs.Domains[0].DomainName)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetCustomDomains(ctx *pulumi.Context, args *GetCustomDomainsArgs, opts ...pulumi.InvokeOption) (*GetCustomDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomDomainsResult
	err := ctx.Invoke("alicloud:fc/getCustomDomains:getCustomDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomDomains.
type GetCustomDomainsArgs struct {
	// A list of functions ids.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Function Compute custom domain name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getCustomDomains.
type GetCustomDomainsResult struct {
	// A list of custom domains, including the following attributes:
	Domains []GetCustomDomainsDomain `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of custom domain ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of custom domain names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetCustomDomainsOutput(ctx *pulumi.Context, args GetCustomDomainsOutputArgs, opts ...pulumi.InvokeOption) GetCustomDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomDomainsResult, error) {
			args := v.(GetCustomDomainsArgs)
			r, err := GetCustomDomains(ctx, &args, opts...)
			var s GetCustomDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomDomainsResultOutput)
}

// A collection of arguments for invoking getCustomDomains.
type GetCustomDomainsOutputArgs struct {
	// A list of functions ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Function Compute custom domain name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetCustomDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getCustomDomains.
type GetCustomDomainsResultOutput struct{ *pulumi.OutputState }

func (GetCustomDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomDomainsResult)(nil)).Elem()
}

func (o GetCustomDomainsResultOutput) ToGetCustomDomainsResultOutput() GetCustomDomainsResultOutput {
	return o
}

func (o GetCustomDomainsResultOutput) ToGetCustomDomainsResultOutputWithContext(ctx context.Context) GetCustomDomainsResultOutput {
	return o
}

// A list of custom domains, including the following attributes:
func (o GetCustomDomainsResultOutput) Domains() GetCustomDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) []GetCustomDomainsDomain { return v.Domains }).(GetCustomDomainsDomainArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of custom domain ids.
func (o GetCustomDomainsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCustomDomainsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of custom domain names.
func (o GetCustomDomainsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetCustomDomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomDomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomDomainsResultOutput{})
}
