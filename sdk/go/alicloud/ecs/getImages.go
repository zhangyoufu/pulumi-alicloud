// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides available image resources. It contains user's private images, system images provided by Alibaba Cloud,
// other public images and the ones available on the image market.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/images.html.markdown.
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("alicloud:ecs/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	// If more than one result are returned, select the most recent one.
	MostRecent *bool `pulumi:"mostRecent"`
	// A regex string to filter resulting images by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Filter results by a specific image owner. Valid items are `system`, `self`, `others`, `marketplace`.
	Owners *string `pulumi:"owners"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of image IDs.
	Ids []string `pulumi:"ids"`
	// A list of images. Each element contains the following attributes:
	Images     []GetImagesImage `pulumi:"images"`
	MostRecent *bool            `pulumi:"mostRecent"`
	NameRegex  *string          `pulumi:"nameRegex"`
	OutputFile *string          `pulumi:"outputFile"`
	Owners     *string          `pulumi:"owners"`
}
