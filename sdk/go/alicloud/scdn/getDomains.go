// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Scdn Domains of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.131.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/scdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nameRegex, err := scdn.GetDomains(ctx, &scdn.GetDomainsArgs{
//				NameRegex: pulumi.StringRef("^my-Domain"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("scdnDomainId", nameRegex.Domains[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainsResult
	err := ctx.Invoke("alicloud:scdn/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Domain IDs. Its element value is same as Domain Name.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Domain name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The Resource Group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	Domains       []GetDomainsDomain `pulumi:"domains"`
	EnableDetails *bool              `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	Ids             []string `pulumi:"ids"`
	NameRegex       *string  `pulumi:"nameRegex"`
	Names           []string `pulumi:"names"`
	OutputFile      *string  `pulumi:"outputFile"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	Status          *string  `pulumi:"status"`
}

func GetDomainsOutput(ctx *pulumi.Context, args GetDomainsOutputArgs, opts ...pulumi.InvokeOption) GetDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainsResult, error) {
			args := v.(GetDomainsArgs)
			r, err := GetDomains(ctx, &args, opts...)
			var s GetDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainsResultOutput)
}

// A collection of arguments for invoking getDomains.
type GetDomainsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Domain IDs. Its element value is same as Domain Name.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Domain name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Resource Group ID.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getDomains.
type GetDomainsResultOutput struct{ *pulumi.OutputState }

func (GetDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsResult)(nil)).Elem()
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutput() GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutputWithContext(ctx context.Context) GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) Domains() GetDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []GetDomainsDomain { return v.Domains }).(GetDomainsDomainArrayOutput)
}

func (o GetDomainsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDomainsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetDomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainsResultOutput{})
}
