// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides the generation of txt records to realize the retrieval and verification of domain names.
//
// > **NOTE:** Available in v1.80.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			this, err := dns.GetDomainTxtGuid(ctx, &dns.GetDomainTxtGuidArgs{
//				DomainName: "test111.abc",
//				Type:       "ADD_SUB_DOMAIN",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("rr", this.Rr)
//			ctx.Export("value", this.Value)
//			return nil
//		})
//	}
//
// ```
func GetDomainTxtGuid(ctx *pulumi.Context, args *GetDomainTxtGuidArgs, opts ...pulumi.InvokeOption) (*GetDomainTxtGuidResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainTxtGuidResult
	err := ctx.Invoke("alicloud:dns/getDomainTxtGuid:getDomainTxtGuid", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomainTxtGuid.
type GetDomainTxtGuidArgs struct {
	// Verified domain name.
	DomainName string `pulumi:"domainName"`
	// User language.
	Lang *string `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Txt verification function. Value:`ADD_SUB_DOMAIN`, `RETRIEVAL`.
	Type string `pulumi:"type"`
}

// A collection of values returned by getDomainTxtGuid.
type GetDomainTxtGuidResult struct {
	DomainName string `pulumi:"domainName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	Lang       *string `pulumi:"lang"`
	OutputFile *string `pulumi:"outputFile"`
	// Host record.
	Rr   string `pulumi:"rr"`
	Type string `pulumi:"type"`
	// Record the value.
	Value string `pulumi:"value"`
}

func GetDomainTxtGuidOutput(ctx *pulumi.Context, args GetDomainTxtGuidOutputArgs, opts ...pulumi.InvokeOption) GetDomainTxtGuidResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainTxtGuidResult, error) {
			args := v.(GetDomainTxtGuidArgs)
			r, err := GetDomainTxtGuid(ctx, &args, opts...)
			var s GetDomainTxtGuidResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainTxtGuidResultOutput)
}

// A collection of arguments for invoking getDomainTxtGuid.
type GetDomainTxtGuidOutputArgs struct {
	// Verified domain name.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// User language.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Txt verification function. Value:`ADD_SUB_DOMAIN`, `RETRIEVAL`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetDomainTxtGuidOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTxtGuidArgs)(nil)).Elem()
}

// A collection of values returned by getDomainTxtGuid.
type GetDomainTxtGuidResultOutput struct{ *pulumi.OutputState }

func (GetDomainTxtGuidResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainTxtGuidResult)(nil)).Elem()
}

func (o GetDomainTxtGuidResultOutput) ToGetDomainTxtGuidResultOutput() GetDomainTxtGuidResultOutput {
	return o
}

func (o GetDomainTxtGuidResultOutput) ToGetDomainTxtGuidResultOutputWithContext(ctx context.Context) GetDomainTxtGuidResultOutput {
	return o
}

func (o GetDomainTxtGuidResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDomainTxtGuidResult] {
	return pulumix.Output[GetDomainTxtGuidResult]{
		OutputState: o.OutputState,
	}
}

func (o GetDomainTxtGuidResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainTxtGuidResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainTxtGuidResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetDomainTxtGuidResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Host record.
func (o GetDomainTxtGuidResultOutput) Rr() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) string { return v.Rr }).(pulumi.StringOutput)
}

func (o GetDomainTxtGuidResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) string { return v.Type }).(pulumi.StringOutput)
}

// Record the value.
func (o GetDomainTxtGuidResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainTxtGuidResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainTxtGuidResultOutput{})
}
