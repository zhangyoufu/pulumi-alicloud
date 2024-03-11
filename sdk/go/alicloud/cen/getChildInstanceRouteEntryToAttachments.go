// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Cen Child Instance Route Entry To Attachment available to the user.[What is Child Instance Route Entry To Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createcenchildinstancerouteentrytoattachment)
//
// > **NOTE:** Available in 1.195.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := cen.GetChildInstanceRouteEntryToAttachments(ctx, &cen.GetChildInstanceRouteEntryToAttachmentsArgs{
//				ChildInstanceRouteTableId: "vtb-t4nt0z5xxbti85c78nkzy",
//				TransitRouterAttachmentId: "tr-attach-f1fd1y50rql00emvej",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudCenChildInstanceRouteEntryToAttachmentExampleId", _default.Attachments[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetChildInstanceRouteEntryToAttachments(ctx *pulumi.Context, args *GetChildInstanceRouteEntryToAttachmentsArgs, opts ...pulumi.InvokeOption) (*GetChildInstanceRouteEntryToAttachmentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetChildInstanceRouteEntryToAttachmentsResult
	err := ctx.Invoke("alicloud:cen/getChildInstanceRouteEntryToAttachments:getChildInstanceRouteEntryToAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getChildInstanceRouteEntryToAttachments.
type GetChildInstanceRouteEntryToAttachmentsArgs struct {
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// The first ID of the resource
	ChildInstanceRouteTableId string `pulumi:"childInstanceRouteTableId"`
	// Limit search to a list of specific IDs.The value is formulated as `<cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>`.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// ServiceType
	ServiceType *string `pulumi:"serviceType"`
	// TransitRouterAttachmentId
	TransitRouterAttachmentId string `pulumi:"transitRouterAttachmentId"`
}

// A collection of values returned by getChildInstanceRouteEntryToAttachments.
type GetChildInstanceRouteEntryToAttachmentsResult struct {
	// A list of Child Instance Route Entry To Attachment Entries. Each element contains the following attributes:
	Attachments []GetChildInstanceRouteEntryToAttachmentsAttachment `pulumi:"attachments"`
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// The first ID of the resource
	ChildInstanceRouteTableId string `pulumi:"childInstanceRouteTableId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Limit search to a list of specific IDs.The value is formulated as `<cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>`.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// ServiceType
	ServiceType *string `pulumi:"serviceType"`
	// TransitRouterAttachmentId
	TransitRouterAttachmentId string `pulumi:"transitRouterAttachmentId"`
}

func GetChildInstanceRouteEntryToAttachmentsOutput(ctx *pulumi.Context, args GetChildInstanceRouteEntryToAttachmentsOutputArgs, opts ...pulumi.InvokeOption) GetChildInstanceRouteEntryToAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetChildInstanceRouteEntryToAttachmentsResult, error) {
			args := v.(GetChildInstanceRouteEntryToAttachmentsArgs)
			r, err := GetChildInstanceRouteEntryToAttachments(ctx, &args, opts...)
			var s GetChildInstanceRouteEntryToAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetChildInstanceRouteEntryToAttachmentsResultOutput)
}

// A collection of arguments for invoking getChildInstanceRouteEntryToAttachments.
type GetChildInstanceRouteEntryToAttachmentsOutputArgs struct {
	// The ID of the CEN instance.
	CenId pulumi.StringPtrInput `pulumi:"cenId"`
	// The first ID of the resource
	ChildInstanceRouteTableId pulumi.StringInput `pulumi:"childInstanceRouteTableId"`
	// Limit search to a list of specific IDs.The value is formulated as `<cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// ServiceType
	ServiceType pulumi.StringPtrInput `pulumi:"serviceType"`
	// TransitRouterAttachmentId
	TransitRouterAttachmentId pulumi.StringInput `pulumi:"transitRouterAttachmentId"`
}

func (GetChildInstanceRouteEntryToAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChildInstanceRouteEntryToAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by getChildInstanceRouteEntryToAttachments.
type GetChildInstanceRouteEntryToAttachmentsResultOutput struct{ *pulumi.OutputState }

func (GetChildInstanceRouteEntryToAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChildInstanceRouteEntryToAttachmentsResult)(nil)).Elem()
}

func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) ToGetChildInstanceRouteEntryToAttachmentsResultOutput() GetChildInstanceRouteEntryToAttachmentsResultOutput {
	return o
}

func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) ToGetChildInstanceRouteEntryToAttachmentsResultOutputWithContext(ctx context.Context) GetChildInstanceRouteEntryToAttachmentsResultOutput {
	return o
}

// A list of Child Instance Route Entry To Attachment Entries. Each element contains the following attributes:
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) Attachments() GetChildInstanceRouteEntryToAttachmentsAttachmentArrayOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) []GetChildInstanceRouteEntryToAttachmentsAttachment {
		return v.Attachments
	}).(GetChildInstanceRouteEntryToAttachmentsAttachmentArrayOutput)
}

// The ID of the CEN instance.
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) CenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) *string { return v.CenId }).(pulumi.StringPtrOutput)
}

// The first ID of the resource
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) ChildInstanceRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) string { return v.ChildInstanceRouteTableId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Limit search to a list of specific IDs.The value is formulated as `<cen_id>:<child_instance_route_table_id>:<transit_router_attachment_id>:<destination_cidr_block>`.
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// ServiceType
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) *string { return v.ServiceType }).(pulumi.StringPtrOutput)
}

// TransitRouterAttachmentId
func (o GetChildInstanceRouteEntryToAttachmentsResultOutput) TransitRouterAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetChildInstanceRouteEntryToAttachmentsResult) string { return v.TransitRouterAttachmentId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChildInstanceRouteEntryToAttachmentsResultOutput{})
}
