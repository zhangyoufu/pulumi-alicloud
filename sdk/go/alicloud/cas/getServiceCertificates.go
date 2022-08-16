// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ssl Certificates Service Certificates of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.129.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			certs, err := cas.GetCertificates(ctx, &cas.GetCertificatesArgs{
//				NameRegex: pulumi.StringRef("^cas"),
//				Ids: []string{
//					"Certificate Id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cert", certs.Certificates[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetServiceCertificates(ctx *pulumi.Context, args *GetServiceCertificatesArgs, opts ...pulumi.InvokeOption) (*GetServiceCertificatesResult, error) {
	var rv GetServiceCertificatesResult
	err := ctx.Invoke("alicloud:cas/getServiceCertificates:getServiceCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceCertificates.
type GetServiceCertificatesArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Certificate IDs.
	Ids []string `pulumi:"ids"`
	// The lang.
	Lang *string `pulumi:"lang"`
	// A regex string to filter results by Certificate name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getServiceCertificates.
type GetServiceCertificatesResult struct {
	Certificates  []GetServiceCertificatesCertificate `pulumi:"certificates"`
	EnableDetails *bool                               `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	Lang       *string  `pulumi:"lang"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetServiceCertificatesOutput(ctx *pulumi.Context, args GetServiceCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetServiceCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceCertificatesResult, error) {
			args := v.(GetServiceCertificatesArgs)
			r, err := GetServiceCertificates(ctx, &args, opts...)
			var s GetServiceCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceCertificatesResultOutput)
}

// A collection of arguments for invoking getServiceCertificates.
type GetServiceCertificatesOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Certificate IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The lang.
	Lang pulumi.StringPtrInput `pulumi:"lang"`
	// A regex string to filter results by Certificate name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetServiceCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getServiceCertificates.
type GetServiceCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetServiceCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesResult)(nil)).Elem()
}

func (o GetServiceCertificatesResultOutput) ToGetServiceCertificatesResultOutput() GetServiceCertificatesResultOutput {
	return o
}

func (o GetServiceCertificatesResultOutput) ToGetServiceCertificatesResultOutputWithContext(ctx context.Context) GetServiceCertificatesResultOutput {
	return o
}

func (o GetServiceCertificatesResultOutput) Certificates() GetServiceCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) []GetServiceCertificatesCertificate { return v.Certificates }).(GetServiceCertificatesCertificateArrayOutput)
}

func (o GetServiceCertificatesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServiceCertificatesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetServiceCertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetServiceCertificatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServiceCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceCertificatesResultOutput{})
}
