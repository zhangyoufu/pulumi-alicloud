// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the ots instance attachments of the current Alibaba Cloud user.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			attachmentsDs, err := ots.GetInstanceAttachments(ctx, &ots.GetInstanceAttachmentsArgs{
//				InstanceName: "sample-instance",
//				NameRegex:    pulumi.StringRef("testvpc"),
//				OutputFile:   pulumi.StringRef("attachments.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstOtsAttachmentId", attachmentsDs.Attachments[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// Deprecated: alicloud.oss.getInstanceAttachments has been deprecated in favor of alicloud.ots.getInstanceAttachments
func GetInstanceAttachments(ctx *pulumi.Context, args *GetInstanceAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetInstanceAttachmentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceAttachmentsResult
	err := ctx.Invoke("alicloud:oss/getInstanceAttachments:getInstanceAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceAttachments.
type GetInstanceAttachmentsArgs struct {
	// The name of OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// A regex string to filter results by vpc name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInstanceAttachments.
type GetInstanceAttachmentsResult struct {
	// A list of instance attachments. Each element contains the following attributes:
	Attachments []GetInstanceAttachmentsAttachment `pulumi:"attachments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instance name.
	InstanceName string  `pulumi:"instanceName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of vpc names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of vpc ids.
	VpcIds []string `pulumi:"vpcIds"`
}

func GetInstanceAttachmentsOutput(ctx *pulumi.Context, args GetInstanceAttachmentsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceAttachmentsResult, error) {
			args := v.(GetInstanceAttachmentsArgs)
			r, err := GetInstanceAttachments(ctx, &args, opts...)
			var s GetInstanceAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceAttachmentsResultOutput)
}

// A collection of arguments for invoking getInstanceAttachments.
type GetInstanceAttachmentsOutputArgs struct {
	// The name of OTS instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// A regex string to filter results by vpc name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetInstanceAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceAttachments.
type GetInstanceAttachmentsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceAttachmentsResult)(nil)).Elem()
}

func (o GetInstanceAttachmentsResultOutput) ToGetInstanceAttachmentsResultOutput() GetInstanceAttachmentsResultOutput {
	return o
}

func (o GetInstanceAttachmentsResultOutput) ToGetInstanceAttachmentsResultOutputWithContext(ctx context.Context) GetInstanceAttachmentsResultOutput {
	return o
}

// A list of instance attachments. Each element contains the following attributes:
func (o GetInstanceAttachmentsResultOutput) Attachments() GetInstanceAttachmentsAttachmentArrayOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) []GetInstanceAttachmentsAttachment { return v.Attachments }).(GetInstanceAttachmentsAttachmentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The instance name.
func (o GetInstanceAttachmentsResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetInstanceAttachmentsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of vpc names.
func (o GetInstanceAttachmentsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetInstanceAttachmentsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of vpc ids.
func (o GetInstanceAttachmentsResultOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceAttachmentsResult) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceAttachmentsResultOutput{})
}
