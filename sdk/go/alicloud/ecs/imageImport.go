// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Import a copy of your local on-premise file to ECS, and appear as a custom replacement in the corresponding domain.
//
// > **NOTE:** You must upload the image file to the object storage OSS in advance.
//
// > **NOTE:** The region where the image is imported must be the same region as the OSS bucket where the image file is uploaded.
//
// > **NOTE:** Available in 1.69.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewImageImport(ctx, "this", &ecs.ImageImportArgs{
//				Architecture: pulumi.String("x86_64"),
//				Description:  pulumi.String("test import image"),
//				DiskDeviceMappings: ecs.ImageImportDiskDeviceMappingArray{
//					&ecs.ImageImportDiskDeviceMappingArgs{
//						DiskImageSize: pulumi.Int(5),
//						OssBucket:     pulumi.String("testimportimage"),
//						OssObject:     pulumi.String("root.img"),
//					},
//				},
//				ImageName:   pulumi.String("test-import-image"),
//				LicenseType: pulumi.String("Auto"),
//				OsType:      pulumi.String("linux"),
//				Platform:    pulumi.String("Ubuntu"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Attributes Reference0
//
//	The following attributes are exported:
//
// * `id` - ID of the image.
//
// ## Import
//
// image can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/imageImport:ImageImport default m-uf66871ape***yg1q***
//
// ```
type ImageImport struct {
	pulumi.CustomResourceState

	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageImportDiskDeviceMappingArrayOutput `pulumi:"diskDeviceMappings"`
	// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
	ImageName pulumi.StringPtrOutput `pulumi:"imageName"`
	// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Operating system platform type. Valid values: `windows`, Default is `linux`.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The operating system distribution. Default value: Others Linux.
	// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
	// **NOTE**: It's default value is Ubuntu before version 1.197.0.
	Platform pulumi.StringOutput `pulumi:"platform"`
}

// NewImageImport registers a new resource with the given unique name, arguments, and options.
func NewImageImport(ctx *pulumi.Context,
	name string, args *ImageImportArgs, opts ...pulumi.ResourceOption) (*ImageImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskDeviceMappings == nil {
		return nil, errors.New("invalid value for required argument 'DiskDeviceMappings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageImport
	err := ctx.RegisterResource("alicloud:ecs/imageImport:ImageImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageImport gets an existing ImageImport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageImportState, opts ...pulumi.ResourceOption) (*ImageImport, error) {
	var resource ImageImport
	err := ctx.ReadResource("alicloud:ecs/imageImport:ImageImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageImport resources.
type imageImportState struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture *string `pulumi:"architecture"`
	// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
	Description *string `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings []ImageImportDiskDeviceMapping `pulumi:"diskDeviceMappings"`
	// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
	ImageName *string `pulumi:"imageName"`
	// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
	LicenseType *string `pulumi:"licenseType"`
	// Operating system platform type. Valid values: `windows`, Default is `linux`.
	OsType *string `pulumi:"osType"`
	// The operating system distribution. Default value: Others Linux.
	// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
	// **NOTE**: It's default value is Ubuntu before version 1.197.0.
	Platform *string `pulumi:"platform"`
}

type ImageImportState struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrInput
	// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
	Description pulumi.StringPtrInput
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageImportDiskDeviceMappingArrayInput
	// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
	ImageName pulumi.StringPtrInput
	// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
	LicenseType pulumi.StringPtrInput
	// Operating system platform type. Valid values: `windows`, Default is `linux`.
	OsType pulumi.StringPtrInput
	// The operating system distribution. Default value: Others Linux.
	// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
	// **NOTE**: It's default value is Ubuntu before version 1.197.0.
	Platform pulumi.StringPtrInput
}

func (ImageImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageImportState)(nil)).Elem()
}

type imageImportArgs struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture *string `pulumi:"architecture"`
	// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
	Description *string `pulumi:"description"`
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings []ImageImportDiskDeviceMapping `pulumi:"diskDeviceMappings"`
	// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
	ImageName *string `pulumi:"imageName"`
	// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
	LicenseType *string `pulumi:"licenseType"`
	// Operating system platform type. Valid values: `windows`, Default is `linux`.
	OsType *string `pulumi:"osType"`
	// The operating system distribution. Default value: Others Linux.
	// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
	// **NOTE**: It's default value is Ubuntu before version 1.197.0.
	Platform *string `pulumi:"platform"`
}

// The set of arguments for constructing a ImageImport resource.
type ImageImportArgs struct {
	// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
	Architecture pulumi.StringPtrInput
	// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
	Description pulumi.StringPtrInput
	// Description of the system with disks and snapshots under the image.
	DiskDeviceMappings ImageImportDiskDeviceMappingArrayInput
	// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
	ImageName pulumi.StringPtrInput
	// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
	LicenseType pulumi.StringPtrInput
	// Operating system platform type. Valid values: `windows`, Default is `linux`.
	OsType pulumi.StringPtrInput
	// The operating system distribution. Default value: Others Linux.
	// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
	// **NOTE**: It's default value is Ubuntu before version 1.197.0.
	Platform pulumi.StringPtrInput
}

func (ImageImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageImportArgs)(nil)).Elem()
}

type ImageImportInput interface {
	pulumi.Input

	ToImageImportOutput() ImageImportOutput
	ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput
}

func (*ImageImport) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageImport)(nil)).Elem()
}

func (i *ImageImport) ToImageImportOutput() ImageImportOutput {
	return i.ToImageImportOutputWithContext(context.Background())
}

func (i *ImageImport) ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportOutput)
}

// ImageImportArrayInput is an input type that accepts ImageImportArray and ImageImportArrayOutput values.
// You can construct a concrete instance of `ImageImportArrayInput` via:
//
//	ImageImportArray{ ImageImportArgs{...} }
type ImageImportArrayInput interface {
	pulumi.Input

	ToImageImportArrayOutput() ImageImportArrayOutput
	ToImageImportArrayOutputWithContext(context.Context) ImageImportArrayOutput
}

type ImageImportArray []ImageImportInput

func (ImageImportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageImport)(nil)).Elem()
}

func (i ImageImportArray) ToImageImportArrayOutput() ImageImportArrayOutput {
	return i.ToImageImportArrayOutputWithContext(context.Background())
}

func (i ImageImportArray) ToImageImportArrayOutputWithContext(ctx context.Context) ImageImportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportArrayOutput)
}

// ImageImportMapInput is an input type that accepts ImageImportMap and ImageImportMapOutput values.
// You can construct a concrete instance of `ImageImportMapInput` via:
//
//	ImageImportMap{ "key": ImageImportArgs{...} }
type ImageImportMapInput interface {
	pulumi.Input

	ToImageImportMapOutput() ImageImportMapOutput
	ToImageImportMapOutputWithContext(context.Context) ImageImportMapOutput
}

type ImageImportMap map[string]ImageImportInput

func (ImageImportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageImport)(nil)).Elem()
}

func (i ImageImportMap) ToImageImportMapOutput() ImageImportMapOutput {
	return i.ToImageImportMapOutputWithContext(context.Background())
}

func (i ImageImportMap) ToImageImportMapOutputWithContext(ctx context.Context) ImageImportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportMapOutput)
}

type ImageImportOutput struct{ *pulumi.OutputState }

func (ImageImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageImport)(nil)).Elem()
}

func (o ImageImportOutput) ToImageImportOutput() ImageImportOutput {
	return o
}

func (o ImageImportOutput) ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput {
	return o
}

// Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: `i386` , Default is `x8664`.
func (o ImageImportOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.Architecture }).(pulumi.StringPtrOutput)
}

// Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.
func (o ImageImportOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Description of the system with disks and snapshots under the image.
func (o ImageImportOutput) DiskDeviceMappings() ImageImportDiskDeviceMappingArrayOutput {
	return o.ApplyT(func(v *ImageImport) ImageImportDiskDeviceMappingArrayOutput { return v.DiskDeviceMappings }).(ImageImportDiskDeviceMappingArrayOutput)
}

// The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).
func (o ImageImportOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.ImageName }).(pulumi.StringPtrOutput)
}

// The type of the license used to activate the operating system after the image is imported. Default value: `Auto`. Valid values: `Auto`,`Aliyun`,`BYOL`.
func (o ImageImportOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Operating system platform type. Valid values: `windows`, Default is `linux`.
func (o ImageImportOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The operating system distribution. Default value: Others Linux.
// More valid values refer to [ImportImage OpenAPI](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/importimage).
// **NOTE**: It's default value is Ubuntu before version 1.197.0.
func (o ImageImportOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

type ImageImportArrayOutput struct{ *pulumi.OutputState }

func (ImageImportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageImport)(nil)).Elem()
}

func (o ImageImportArrayOutput) ToImageImportArrayOutput() ImageImportArrayOutput {
	return o
}

func (o ImageImportArrayOutput) ToImageImportArrayOutputWithContext(ctx context.Context) ImageImportArrayOutput {
	return o
}

func (o ImageImportArrayOutput) Index(i pulumi.IntInput) ImageImportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageImport {
		return vs[0].([]*ImageImport)[vs[1].(int)]
	}).(ImageImportOutput)
}

type ImageImportMapOutput struct{ *pulumi.OutputState }

func (ImageImportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageImport)(nil)).Elem()
}

func (o ImageImportMapOutput) ToImageImportMapOutput() ImageImportMapOutput {
	return o
}

func (o ImageImportMapOutput) ToImageImportMapOutputWithContext(ctx context.Context) ImageImportMapOutput {
	return o
}

func (o ImageImportMapOutput) MapIndex(k pulumi.StringInput) ImageImportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageImport {
		return vs[0].(map[string]*ImageImport)[vs[1].(string)]
	}).(ImageImportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportInput)(nil)).Elem(), &ImageImport{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportArrayInput)(nil)).Elem(), ImageImportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportMapInput)(nil)).Elem(), ImageImportMap{})
	pulumi.RegisterOutputType(ImageImportOutput{})
	pulumi.RegisterOutputType(ImageImportArrayOutput{})
	pulumi.RegisterOutputType(ImageImportMapOutput{})
}
