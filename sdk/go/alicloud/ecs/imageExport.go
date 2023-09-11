// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Export a custom image to the OSS bucket in the same region as the custom image.
//
// > **NOTE:** If you create an ECS instance using a mirror image and create a system disk snapshot again, exporting a custom image created from the system disk snapshot is not supported.
//
// > **NOTE:** Support for exporting custom images that include data disk snapshot information in the image. The number of data disks cannot exceed 4 and the maximum capacity of a single data disk cannot exceed 500 GiB.
//
// > **NOTE:** Before exporting the image, you must authorize the cloud server ECS official service account to write OSS permissions through RAM.
//
// > **NOTE:** Available in 1.68.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				InstanceTypeFamily: pulumi.StringRef("ecs.sn1ne"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex: pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//				Owners:    pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := ecs.NewInstance(ctx, "defaultInstance", &ecs.InstanceArgs{
//				AvailabilityZone: *pulumi.String(defaultZones.Zones[0].Id),
//				InstanceName:     pulumi.String("terraform-example"),
//				SecurityGroups: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//				VswitchId:               defaultSwitch.ID(),
//				InstanceType:            *pulumi.String(defaultInstanceTypes.Ids[0]),
//				ImageId:                 *pulumi.String(defaultImages.Ids[0]),
//				InternetMaxBandwidthOut: pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			defaultImage, err := ecs.NewImage(ctx, "defaultImage", &ecs.ImageArgs{
//				InstanceId:  defaultInstance.ID(),
//				ImageName:   pulumi.String("terraform-example"),
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRandomInteger, err := random.NewRandomInteger(ctx, "defaultRandomInteger", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBucket, err := oss.NewBucket(ctx, "defaultBucket", &oss.BucketArgs{
//				Bucket: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("example-value-%v", result), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewImageExport(ctx, "defaultImageExport", &ecs.ImageExportArgs{
//				ImageId:   defaultImage.ID(),
//				OssBucket: defaultBucket.ID(),
//				OssPrefix: pulumi.String("ecsExport"),
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
type ImageExport struct {
	pulumi.CustomResourceState

	// The source image ID.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// Save the exported OSS bucket.
	OssBucket pulumi.StringOutput `pulumi:"ossBucket"`
	// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
	OssPrefix pulumi.StringPtrOutput `pulumi:"ossPrefix"`
}

// NewImageExport registers a new resource with the given unique name, arguments, and options.
func NewImageExport(ctx *pulumi.Context,
	name string, args *ImageExportArgs, opts ...pulumi.ResourceOption) (*ImageExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.OssBucket == nil {
		return nil, errors.New("invalid value for required argument 'OssBucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageExport
	err := ctx.RegisterResource("alicloud:ecs/imageExport:ImageExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageExport gets an existing ImageExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageExportState, opts ...pulumi.ResourceOption) (*ImageExport, error) {
	var resource ImageExport
	err := ctx.ReadResource("alicloud:ecs/imageExport:ImageExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageExport resources.
type imageExportState struct {
	// The source image ID.
	ImageId *string `pulumi:"imageId"`
	// Save the exported OSS bucket.
	OssBucket *string `pulumi:"ossBucket"`
	// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
	OssPrefix *string `pulumi:"ossPrefix"`
}

type ImageExportState struct {
	// The source image ID.
	ImageId pulumi.StringPtrInput
	// Save the exported OSS bucket.
	OssBucket pulumi.StringPtrInput
	// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
	OssPrefix pulumi.StringPtrInput
}

func (ImageExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageExportState)(nil)).Elem()
}

type imageExportArgs struct {
	// The source image ID.
	ImageId string `pulumi:"imageId"`
	// Save the exported OSS bucket.
	OssBucket string `pulumi:"ossBucket"`
	// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
	OssPrefix *string `pulumi:"ossPrefix"`
}

// The set of arguments for constructing a ImageExport resource.
type ImageExportArgs struct {
	// The source image ID.
	ImageId pulumi.StringInput
	// Save the exported OSS bucket.
	OssBucket pulumi.StringInput
	// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
	OssPrefix pulumi.StringPtrInput
}

func (ImageExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageExportArgs)(nil)).Elem()
}

type ImageExportInput interface {
	pulumi.Input

	ToImageExportOutput() ImageExportOutput
	ToImageExportOutputWithContext(ctx context.Context) ImageExportOutput
}

func (*ImageExport) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageExport)(nil)).Elem()
}

func (i *ImageExport) ToImageExportOutput() ImageExportOutput {
	return i.ToImageExportOutputWithContext(context.Background())
}

func (i *ImageExport) ToImageExportOutputWithContext(ctx context.Context) ImageExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageExportOutput)
}

func (i *ImageExport) ToOutput(ctx context.Context) pulumix.Output[*ImageExport] {
	return pulumix.Output[*ImageExport]{
		OutputState: i.ToImageExportOutputWithContext(ctx).OutputState,
	}
}

// ImageExportArrayInput is an input type that accepts ImageExportArray and ImageExportArrayOutput values.
// You can construct a concrete instance of `ImageExportArrayInput` via:
//
//	ImageExportArray{ ImageExportArgs{...} }
type ImageExportArrayInput interface {
	pulumi.Input

	ToImageExportArrayOutput() ImageExportArrayOutput
	ToImageExportArrayOutputWithContext(context.Context) ImageExportArrayOutput
}

type ImageExportArray []ImageExportInput

func (ImageExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageExport)(nil)).Elem()
}

func (i ImageExportArray) ToImageExportArrayOutput() ImageExportArrayOutput {
	return i.ToImageExportArrayOutputWithContext(context.Background())
}

func (i ImageExportArray) ToImageExportArrayOutputWithContext(ctx context.Context) ImageExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageExportArrayOutput)
}

func (i ImageExportArray) ToOutput(ctx context.Context) pulumix.Output[[]*ImageExport] {
	return pulumix.Output[[]*ImageExport]{
		OutputState: i.ToImageExportArrayOutputWithContext(ctx).OutputState,
	}
}

// ImageExportMapInput is an input type that accepts ImageExportMap and ImageExportMapOutput values.
// You can construct a concrete instance of `ImageExportMapInput` via:
//
//	ImageExportMap{ "key": ImageExportArgs{...} }
type ImageExportMapInput interface {
	pulumi.Input

	ToImageExportMapOutput() ImageExportMapOutput
	ToImageExportMapOutputWithContext(context.Context) ImageExportMapOutput
}

type ImageExportMap map[string]ImageExportInput

func (ImageExportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageExport)(nil)).Elem()
}

func (i ImageExportMap) ToImageExportMapOutput() ImageExportMapOutput {
	return i.ToImageExportMapOutputWithContext(context.Background())
}

func (i ImageExportMap) ToImageExportMapOutputWithContext(ctx context.Context) ImageExportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageExportMapOutput)
}

func (i ImageExportMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ImageExport] {
	return pulumix.Output[map[string]*ImageExport]{
		OutputState: i.ToImageExportMapOutputWithContext(ctx).OutputState,
	}
}

type ImageExportOutput struct{ *pulumi.OutputState }

func (ImageExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageExport)(nil)).Elem()
}

func (o ImageExportOutput) ToImageExportOutput() ImageExportOutput {
	return o
}

func (o ImageExportOutput) ToImageExportOutputWithContext(ctx context.Context) ImageExportOutput {
	return o
}

func (o ImageExportOutput) ToOutput(ctx context.Context) pulumix.Output[*ImageExport] {
	return pulumix.Output[*ImageExport]{
		OutputState: o.OutputState,
	}
}

// The source image ID.
func (o ImageExportOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageExport) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// Save the exported OSS bucket.
func (o ImageExportOutput) OssBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageExport) pulumi.StringOutput { return v.OssBucket }).(pulumi.StringOutput)
}

// The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30.
func (o ImageExportOutput) OssPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageExport) pulumi.StringPtrOutput { return v.OssPrefix }).(pulumi.StringPtrOutput)
}

type ImageExportArrayOutput struct{ *pulumi.OutputState }

func (ImageExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageExport)(nil)).Elem()
}

func (o ImageExportArrayOutput) ToImageExportArrayOutput() ImageExportArrayOutput {
	return o
}

func (o ImageExportArrayOutput) ToImageExportArrayOutputWithContext(ctx context.Context) ImageExportArrayOutput {
	return o
}

func (o ImageExportArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ImageExport] {
	return pulumix.Output[[]*ImageExport]{
		OutputState: o.OutputState,
	}
}

func (o ImageExportArrayOutput) Index(i pulumi.IntInput) ImageExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageExport {
		return vs[0].([]*ImageExport)[vs[1].(int)]
	}).(ImageExportOutput)
}

type ImageExportMapOutput struct{ *pulumi.OutputState }

func (ImageExportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageExport)(nil)).Elem()
}

func (o ImageExportMapOutput) ToImageExportMapOutput() ImageExportMapOutput {
	return o
}

func (o ImageExportMapOutput) ToImageExportMapOutputWithContext(ctx context.Context) ImageExportMapOutput {
	return o
}

func (o ImageExportMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ImageExport] {
	return pulumix.Output[map[string]*ImageExport]{
		OutputState: o.OutputState,
	}
}

func (o ImageExportMapOutput) MapIndex(k pulumi.StringInput) ImageExportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageExport {
		return vs[0].(map[string]*ImageExport)[vs[1].(string)]
	}).(ImageExportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageExportInput)(nil)).Elem(), &ImageExport{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageExportArrayInput)(nil)).Elem(), ImageExportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageExportMapInput)(nil)).Elem(), ImageExportMap{})
	pulumi.RegisterOutputType(ImageExportOutput{})
	pulumi.RegisterOutputType(ImageExportArrayOutput{})
	pulumi.RegisterOutputType(ImageExportMapOutput{})
}
