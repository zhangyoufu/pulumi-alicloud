// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Dcdn Ipa Domains of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.158.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := dcdn.GetIpaDomains(ctx, &dcdn.GetIpaDomainsArgs{
//				DomainName: pulumi.StringRef("example_value"),
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("dcdnIpaDomainId1", ids.Domains[0].Id)
//			status, err := dcdn.GetIpaDomains(ctx, &dcdn.GetIpaDomainsArgs{
//				Status: pulumi.StringRef("online"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("dcdnIpaDomainId2", status.Domains[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetIpaDomains(ctx *pulumi.Context, args *GetIpaDomainsArgs, opts ...pulumi.InvokeOption) (*GetIpaDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpaDomainsResult
	err := ctx.Invoke("alicloud:dcdn/getIpaDomains:getIpaDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpaDomains.
type GetIpaDomainsArgs struct {
	// The accelerated domain names.
	DomainName *string `pulumi:"domainName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Ipa Domain IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the accelerated domain name.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getIpaDomains.
type GetIpaDomainsResult struct {
	DomainName    *string               `pulumi:"domainName"`
	Domains       []GetIpaDomainsDomain `pulumi:"domains"`
	EnableDetails *bool                 `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetIpaDomainsOutput(ctx *pulumi.Context, args GetIpaDomainsOutputArgs, opts ...pulumi.InvokeOption) GetIpaDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpaDomainsResult, error) {
			args := v.(GetIpaDomainsArgs)
			r, err := GetIpaDomains(ctx, &args, opts...)
			var s GetIpaDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpaDomainsResultOutput)
}

// A collection of arguments for invoking getIpaDomains.
type GetIpaDomainsOutputArgs struct {
	// The accelerated domain names.
	DomainName pulumi.StringPtrInput `pulumi:"domainName"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Ipa Domain IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the accelerated domain name.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetIpaDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpaDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getIpaDomains.
type GetIpaDomainsResultOutput struct{ *pulumi.OutputState }

func (GetIpaDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpaDomainsResult)(nil)).Elem()
}

func (o GetIpaDomainsResultOutput) ToGetIpaDomainsResultOutput() GetIpaDomainsResultOutput {
	return o
}

func (o GetIpaDomainsResultOutput) ToGetIpaDomainsResultOutputWithContext(ctx context.Context) GetIpaDomainsResultOutput {
	return o
}

func (o GetIpaDomainsResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o GetIpaDomainsResultOutput) Domains() GetIpaDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) []GetIpaDomainsDomain { return v.Domains }).(GetIpaDomainsDomainArrayOutput)
}

func (o GetIpaDomainsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpaDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpaDomainsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetIpaDomainsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetIpaDomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetIpaDomainsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpaDomainsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpaDomainsResultOutput{})
}
