// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.
//
// > **NOTE:** You can only share your own custom images to other Alibaba Cloud users.
//
// > **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.
//
// > **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.
//
// > **NOTE:** Available in 1.68.0+.
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
// 		_, err := ecs.NewImageSharePermission(ctx, "_default", &ecs.ImageSharePermissionArgs{
// 			AccountId: pulumi.String("1234567890"),
// 			ImageId:   pulumi.String("m-bp1gxyh***"),
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
// * `id` - ID of the image. It formats as `<image_id>:<account_id>`
//
// ## Import
//
// image can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:ecs/imageSharePermission:ImageSharePermission default m-uf66yg1q:123456789
// ```
type ImageSharePermission struct {
	pulumi.CustomResourceState

	// Alibaba Cloud Account ID. It is used to share images.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The source image ID.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
}

// NewImageSharePermission registers a new resource with the given unique name, arguments, and options.
func NewImageSharePermission(ctx *pulumi.Context,
	name string, args *ImageSharePermissionArgs, opts ...pulumi.ResourceOption) (*ImageSharePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	var resource ImageSharePermission
	err := ctx.RegisterResource("alicloud:ecs/imageSharePermission:ImageSharePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageSharePermission gets an existing ImageSharePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageSharePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageSharePermissionState, opts ...pulumi.ResourceOption) (*ImageSharePermission, error) {
	var resource ImageSharePermission
	err := ctx.ReadResource("alicloud:ecs/imageSharePermission:ImageSharePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageSharePermission resources.
type imageSharePermissionState struct {
	// Alibaba Cloud Account ID. It is used to share images.
	AccountId *string `pulumi:"accountId"`
	// The source image ID.
	ImageId *string `pulumi:"imageId"`
}

type ImageSharePermissionState struct {
	// Alibaba Cloud Account ID. It is used to share images.
	AccountId pulumi.StringPtrInput
	// The source image ID.
	ImageId pulumi.StringPtrInput
}

func (ImageSharePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSharePermissionState)(nil)).Elem()
}

type imageSharePermissionArgs struct {
	// Alibaba Cloud Account ID. It is used to share images.
	AccountId string `pulumi:"accountId"`
	// The source image ID.
	ImageId string `pulumi:"imageId"`
}

// The set of arguments for constructing a ImageSharePermission resource.
type ImageSharePermissionArgs struct {
	// Alibaba Cloud Account ID. It is used to share images.
	AccountId pulumi.StringInput
	// The source image ID.
	ImageId pulumi.StringInput
}

func (ImageSharePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSharePermissionArgs)(nil)).Elem()
}

type ImageSharePermissionInput interface {
	pulumi.Input

	ToImageSharePermissionOutput() ImageSharePermissionOutput
	ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput
}

func (*ImageSharePermission) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageSharePermission)(nil))
}

func (i *ImageSharePermission) ToImageSharePermissionOutput() ImageSharePermissionOutput {
	return i.ToImageSharePermissionOutputWithContext(context.Background())
}

func (i *ImageSharePermission) ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionOutput)
}

func (i *ImageSharePermission) ToImageSharePermissionPtrOutput() ImageSharePermissionPtrOutput {
	return i.ToImageSharePermissionPtrOutputWithContext(context.Background())
}

func (i *ImageSharePermission) ToImageSharePermissionPtrOutputWithContext(ctx context.Context) ImageSharePermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionPtrOutput)
}

type ImageSharePermissionPtrInput interface {
	pulumi.Input

	ToImageSharePermissionPtrOutput() ImageSharePermissionPtrOutput
	ToImageSharePermissionPtrOutputWithContext(ctx context.Context) ImageSharePermissionPtrOutput
}

type imageSharePermissionPtrType ImageSharePermissionArgs

func (*imageSharePermissionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSharePermission)(nil))
}

func (i *imageSharePermissionPtrType) ToImageSharePermissionPtrOutput() ImageSharePermissionPtrOutput {
	return i.ToImageSharePermissionPtrOutputWithContext(context.Background())
}

func (i *imageSharePermissionPtrType) ToImageSharePermissionPtrOutputWithContext(ctx context.Context) ImageSharePermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionPtrOutput)
}

// ImageSharePermissionArrayInput is an input type that accepts ImageSharePermissionArray and ImageSharePermissionArrayOutput values.
// You can construct a concrete instance of `ImageSharePermissionArrayInput` via:
//
//          ImageSharePermissionArray{ ImageSharePermissionArgs{...} }
type ImageSharePermissionArrayInput interface {
	pulumi.Input

	ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput
	ToImageSharePermissionArrayOutputWithContext(context.Context) ImageSharePermissionArrayOutput
}

type ImageSharePermissionArray []ImageSharePermissionInput

func (ImageSharePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ImageSharePermission)(nil))
}

func (i ImageSharePermissionArray) ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput {
	return i.ToImageSharePermissionArrayOutputWithContext(context.Background())
}

func (i ImageSharePermissionArray) ToImageSharePermissionArrayOutputWithContext(ctx context.Context) ImageSharePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionArrayOutput)
}

// ImageSharePermissionMapInput is an input type that accepts ImageSharePermissionMap and ImageSharePermissionMapOutput values.
// You can construct a concrete instance of `ImageSharePermissionMapInput` via:
//
//          ImageSharePermissionMap{ "key": ImageSharePermissionArgs{...} }
type ImageSharePermissionMapInput interface {
	pulumi.Input

	ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput
	ToImageSharePermissionMapOutputWithContext(context.Context) ImageSharePermissionMapOutput
}

type ImageSharePermissionMap map[string]ImageSharePermissionInput

func (ImageSharePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ImageSharePermission)(nil))
}

func (i ImageSharePermissionMap) ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput {
	return i.ToImageSharePermissionMapOutputWithContext(context.Background())
}

func (i ImageSharePermissionMap) ToImageSharePermissionMapOutputWithContext(ctx context.Context) ImageSharePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionMapOutput)
}

type ImageSharePermissionOutput struct {
	*pulumi.OutputState
}

func (ImageSharePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageSharePermission)(nil))
}

func (o ImageSharePermissionOutput) ToImageSharePermissionOutput() ImageSharePermissionOutput {
	return o
}

func (o ImageSharePermissionOutput) ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput {
	return o
}

func (o ImageSharePermissionOutput) ToImageSharePermissionPtrOutput() ImageSharePermissionPtrOutput {
	return o.ToImageSharePermissionPtrOutputWithContext(context.Background())
}

func (o ImageSharePermissionOutput) ToImageSharePermissionPtrOutputWithContext(ctx context.Context) ImageSharePermissionPtrOutput {
	return o.ApplyT(func(v ImageSharePermission) *ImageSharePermission {
		return &v
	}).(ImageSharePermissionPtrOutput)
}

type ImageSharePermissionPtrOutput struct {
	*pulumi.OutputState
}

func (ImageSharePermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSharePermission)(nil))
}

func (o ImageSharePermissionPtrOutput) ToImageSharePermissionPtrOutput() ImageSharePermissionPtrOutput {
	return o
}

func (o ImageSharePermissionPtrOutput) ToImageSharePermissionPtrOutputWithContext(ctx context.Context) ImageSharePermissionPtrOutput {
	return o
}

type ImageSharePermissionArrayOutput struct{ *pulumi.OutputState }

func (ImageSharePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageSharePermission)(nil))
}

func (o ImageSharePermissionArrayOutput) ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput {
	return o
}

func (o ImageSharePermissionArrayOutput) ToImageSharePermissionArrayOutputWithContext(ctx context.Context) ImageSharePermissionArrayOutput {
	return o
}

func (o ImageSharePermissionArrayOutput) Index(i pulumi.IntInput) ImageSharePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageSharePermission {
		return vs[0].([]ImageSharePermission)[vs[1].(int)]
	}).(ImageSharePermissionOutput)
}

type ImageSharePermissionMapOutput struct{ *pulumi.OutputState }

func (ImageSharePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageSharePermission)(nil))
}

func (o ImageSharePermissionMapOutput) ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput {
	return o
}

func (o ImageSharePermissionMapOutput) ToImageSharePermissionMapOutputWithContext(ctx context.Context) ImageSharePermissionMapOutput {
	return o
}

func (o ImageSharePermissionMapOutput) MapIndex(k pulumi.StringInput) ImageSharePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageSharePermission {
		return vs[0].(map[string]ImageSharePermission)[vs[1].(string)]
	}).(ImageSharePermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageSharePermissionOutput{})
	pulumi.RegisterOutputType(ImageSharePermissionPtrOutput{})
	pulumi.RegisterOutputType(ImageSharePermissionArrayOutput{})
	pulumi.RegisterOutputType(ImageSharePermissionMapOutput{})
}
