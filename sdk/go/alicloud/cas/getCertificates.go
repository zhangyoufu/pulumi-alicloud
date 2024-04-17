// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:**  This datasource has been deprecated from version `1.129.0`. Please use new datasource alicloud_ssl_certificates_service_certificates.
//
// This data source provides a list of CAS Certificates in an Alibaba Cloud account according to the specified filters.
//
// Deprecated: This resource has been deprecated in favour of getServiceCertificates
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCertificatesResult
	err := ctx.Invoke("alicloud:cas/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of cert IDs.
	Ids  []string `pulumi:"ids"`
	Lang *string  `pulumi:"lang"`
	// A regex string to filter results by the certificate name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// A list of apis. Each element contains the following attributes:
	Certificates  []GetCertificatesCertificate `pulumi:"certificates"`
	EnableDetails *bool                        `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of cert IDs.
	Ids       []string `pulumi:"ids"`
	Lang      *string  `pulumi:"lang"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of cert names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetCertificatesOutput(ctx *pulumi.Context, args GetCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCertificatesResult, error) {
			args := v.(GetCertificatesArgs)
			r, err := GetCertificates(ctx, &args, opts...)
			var s GetCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCertificatesResultOutput)
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesOutputArgs struct {
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of cert IDs.
	Ids  pulumi.StringArrayInput `pulumi:"ids"`
	Lang pulumi.StringPtrInput   `pulumi:"lang"`
	// A regex string to filter results by the certificate name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCertificates.
type GetCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesResult)(nil)).Elem()
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutput() GetCertificatesResultOutput {
	return o
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutputWithContext(ctx context.Context) GetCertificatesResultOutput {
	return o
}

// A list of apis. Each element contains the following attributes:
func (o GetCertificatesResultOutput) Certificates() GetCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []GetCertificatesCertificate { return v.Certificates }).(GetCertificatesCertificateArrayOutput)
}

func (o GetCertificatesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of cert IDs.
func (o GetCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCertificatesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetCertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of cert names.
func (o GetCertificatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificatesResultOutput{})
}
