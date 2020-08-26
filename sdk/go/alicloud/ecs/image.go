// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a custom image. You can then use a custom image to create ECS instances (RunInstances) or change the system disk for an existing instance (ReplaceSystemDisk).
//
// > **NOTE:**  If you want to create a template from an ECS instance, you can specify the instance ID (InstanceId) to create a custom image. You must make sure that the status of the specified instance is Running or Stopped. After a successful invocation, each disk of the specified instance has a new snapshot created.
//
// > **NOTE:**  If you want to create a custom image based on the system disk of your ECS instance, you can specify one of the system disk snapshots (SnapshotId) to create a custom image. However, the specified snapshot cannot be created on or before July 15, 2013.
//
// > **NOTE:**  If you want to combine snapshots of multiple disks into an image template, you can specify DiskDeviceMapping to create a custom image.
//
// > **NOTE:**  Available in 1.64.0+
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
// 		_, err := ecs.NewImage(ctx, "_default", &ecs.ImageArgs{
// 			Architecture:    pulumi.String("x86_64"),
// 			Description:     pulumi.String("test-image"),
// 			ImageName:       pulumi.String("test-image"),
// 			InstanceId:      pulumi.String("i-bp1g6zv0ce8oghu7k***"),
// 			Platform:        pulumi.String("CentOS"),
// 			ResourceGroupId: pulumi.String("rg-bp67acfmxazb4ph***"),
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
type Image struct {
	pulumi.CustomResourceState

	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageDiskDeviceMappingArrayOutput `pulumi:"diskDeviceMappings"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The instance ID.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
	Platform pulumi.StringPtrOutput `pulumi:"platform"`
	// The ID of the enterprise resource group to which a custom image belongs
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// Specifies a snapshot that is used to create a combined custom image.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}
	var resource Image
	err := ctx.RegisterResource("alicloud:ecs/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("alicloud:ecs/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture *string `pulumi:"architecture"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings []ImageDiskDeviceMapping `pulumi:"diskDeviceMappings"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// The instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
	Platform *string `pulumi:"platform"`
	// The ID of the enterprise resource group to which a custom image belongs
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Specifies a snapshot that is used to create a combined custom image.
	SnapshotId *string `pulumi:"snapshotId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ImageState struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrInput
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageDiskDeviceMappingArrayInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// The instance ID.
	InstanceId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
	Platform pulumi.StringPtrInput
	// The ID of the enterprise resource group to which a custom image belongs
	ResourceGroupId pulumi.StringPtrInput
	// Specifies a snapshot that is used to create a combined custom image.
	SnapshotId pulumi.StringPtrInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture *string `pulumi:"architecture"`
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description *string `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings []ImageDiskDeviceMapping `pulumi:"diskDeviceMappings"`
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force *bool `pulumi:"force"`
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName *string `pulumi:"imageName"`
	// The instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name *string `pulumi:"name"`
	// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
	Platform *string `pulumi:"platform"`
	// The ID of the enterprise resource group to which a custom image belongs
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Specifies a snapshot that is used to create a combined custom image.
	SnapshotId *string `pulumi:"snapshotId"`
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrInput
	// The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
	Description pulumi.StringPtrInput
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageDiskDeviceMappingArrayInput
	// Indicates whether to force delete the custom image, Default is `false`.
	// - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
	// - false：Verifies that the image is not currently in use by any other instances before deleting the image.
	Force pulumi.BoolPtrInput
	// The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
	ImageName pulumi.StringPtrInput
	// The instance ID.
	InstanceId pulumi.StringPtrInput
	// Deprecated: Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead.
	Name pulumi.StringPtrInput
	// Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `CentOS`, `Ubuntu`, `SUSE`, `OpenSUSE`, `RedHat`, `Debian`, `CoreOS`, `Aliyun Linux`, `Windows Server 2003`, `Windows Server 2008`, `Windows Server 2012`, `Windows 7`, Default is `Others Linux`, `Customized Linux`.
	Platform pulumi.StringPtrInput
	// The ID of the enterprise resource group to which a custom image belongs
	ResourceGroupId pulumi.StringPtrInput
	// Specifies a snapshot that is used to create a combined custom image.
	SnapshotId pulumi.StringPtrInput
	// The tag value of an image. The value of N ranges from 1 to 20.
	Tags pulumi.MapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}
