// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Copies a custom image from one region to another. You can use copied images to perform operations in the target region, such as creating instances (RunInstances) and replacing system disks (ReplaceSystemDisk).
//
// > **NOTE:** You can only copy the custom image when it is in the Available state.
//
// > **NOTE:** You can only copy the image belonging to your Alibaba Cloud account. Images cannot be copied from one account to another.
//
// > **NOTE:** If the copying is not completed, you cannot call DeleteImage to delete the image but you can call CancelCopyImage to cancel the copying.
//
// > **NOTE:** Available in 1.66.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewImageCopy(ctx, "_default", &ecs.ImageCopyArgs{
// 			Description:    pulumi.String("test-image"),
// 			ImageName:      pulumi.String("test-image"),
// 			SourceImageId:  pulumi.String("m-bp1gxyhdswlsn18tu***"),
// 			SourceRegionId: pulumi.String("cn-hangzhou"),
// 			Tags: pulumi.StringMap{
// 				"FinanceDept": pulumi.String("FinanceDeptJoshua"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attributes Reference0
//
//  The following attributes are exported:
//
// * `id` - ID of the image.
type ImageCopy struct {
	pulumi.CustomResourceState

	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The source image ID.
	SourceImageId pulumi.StringOutput `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringOutput `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewImageCopy registers a new resource with the given unique name, arguments, and options.
func NewImageCopy(ctx *pulumi.Context,
	name string, args *ImageCopyArgs, opts ...pulumi.ResourceOption) (*ImageCopy, error) {
	if args == nil || args.SourceImageId == nil {
		return nil, errors.New("missing required argument 'SourceImageId'")
	}
	if args == nil || args.SourceRegionId == nil {
		return nil, errors.New("missing required argument 'SourceRegionId'")
	}
	if args == nil {
		args = &ImageCopyArgs{}
	}
	var resource ImageCopy
	err := ctx.RegisterResource("alicloud:ecs/imageCopy:ImageCopy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageCopy gets an existing ImageCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageCopy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageCopyState, opts ...pulumi.ResourceOption) (*ImageCopy, error) {
	var resource ImageCopy
	err := ctx.ReadResource("alicloud:ecs/imageCopy:ImageCopy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageCopy resources.
type imageCopyState struct {
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// The source image ID.
	SourceImageId *string `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId *string `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ImageCopyState struct {
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// The source image ID.
	SourceImageId pulumi.StringPtrInput
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringPtrInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageCopyState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCopyState)(nil)).Elem()
}

type imageCopyArgs struct {
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Indicates whether to encrypt the image.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// Key ID used to encrypt the image.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// The source image ID.
	SourceImageId string `pulumi:"sourceImageId"`
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId string `pulumi:"sourceRegionId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ImageCopy resource.
type ImageCopyArgs struct {
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Indicates whether to encrypt the image.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// Key ID used to encrypt the image.
	KmsKeyId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// The source image ID.
	SourceImageId pulumi.StringInput
	// The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
	SourceRegionId pulumi.StringInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageCopyArgs)(nil)).Elem()
}
