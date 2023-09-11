// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fc"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oss"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultRandomInteger, err := random.NewRandomInteger(ctx, "defaultRandomInteger", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBucket, err := oss.NewBucket(ctx, "defaultBucket", &oss.BucketArgs{
//				Bucket: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("terraform-example-%v", result), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBucketObject, err := oss.NewBucketObject(ctx, "defaultBucketObject", &oss.BucketObjectArgs{
//				Bucket:  defaultBucket.ID(),
//				Key:     pulumi.String("index.py"),
//				Content: pulumi.String("import logging \ndef handler(event, context): \nlogger = logging.getLogger() \nlogger.info('hello world') \nreturn 'hello world'"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fc.NewLayerVersion(ctx, "example", &fc.LayerVersionArgs{
//				LayerName: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("terraform-example-%v", result), nil
//				}).(pulumi.StringOutput),
//				CompatibleRuntimes: pulumi.StringArray{
//					pulumi.String("python2.7"),
//				},
//				OssBucketName: defaultBucket.Bucket,
//				OssObjectName: defaultBucketObject.Key,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Function Compute Layer Version can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:fc/layerVersion:LayerVersion example my_function
//
// ```
type LayerVersion struct {
	pulumi.CustomResourceState

	// The access mode of Layer Version.
	Acl pulumi.StringOutput `pulumi:"acl"`
	// The arn of Layer Version.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The checksum of the layer code package.
	CodeCheckSum pulumi.StringOutput `pulumi:"codeCheckSum"`
	// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
	CompatibleRuntimes pulumi.StringArrayOutput `pulumi:"compatibleRuntimes"`
	// The description of the layer version.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the layer.
	LayerName pulumi.StringOutput `pulumi:"layerName"`
	// The name of the OSS bucket that stores the ZIP package of the function code.
	OssBucketName pulumi.StringPtrOutput `pulumi:"ossBucketName"`
	// The name of the OSS object (ZIP package) that contains the function code.
	OssObjectName pulumi.StringPtrOutput `pulumi:"ossObjectName"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// The version of Layer Version.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ZIP package of the function code that is encoded in the Base64 format.
	//
	// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
	ZipFile pulumi.StringPtrOutput `pulumi:"zipFile"`
}

// NewLayerVersion registers a new resource with the given unique name, arguments, and options.
func NewLayerVersion(ctx *pulumi.Context,
	name string, args *LayerVersionArgs, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompatibleRuntimes == nil {
		return nil, errors.New("invalid value for required argument 'CompatibleRuntimes'")
	}
	if args.LayerName == nil {
		return nil, errors.New("invalid value for required argument 'LayerName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LayerVersion
	err := ctx.RegisterResource("alicloud:fc/layerVersion:LayerVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayerVersion gets an existing LayerVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayerVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LayerVersionState, opts ...pulumi.ResourceOption) (*LayerVersion, error) {
	var resource LayerVersion
	err := ctx.ReadResource("alicloud:fc/layerVersion:LayerVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LayerVersion resources.
type layerVersionState struct {
	// The access mode of Layer Version.
	Acl *string `pulumi:"acl"`
	// The arn of Layer Version.
	Arn *string `pulumi:"arn"`
	// The checksum of the layer code package.
	CodeCheckSum *string `pulumi:"codeCheckSum"`
	// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// The description of the layer version.
	Description *string `pulumi:"description"`
	// The name of the layer.
	LayerName *string `pulumi:"layerName"`
	// The name of the OSS bucket that stores the ZIP package of the function code.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The name of the OSS object (ZIP package) that contains the function code.
	OssObjectName *string `pulumi:"ossObjectName"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// The version of Layer Version.
	Version *string `pulumi:"version"`
	// The ZIP package of the function code that is encoded in the Base64 format.
	//
	// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
	ZipFile *string `pulumi:"zipFile"`
}

type LayerVersionState struct {
	// The access mode of Layer Version.
	Acl pulumi.StringPtrInput
	// The arn of Layer Version.
	Arn pulumi.StringPtrInput
	// The checksum of the layer code package.
	CodeCheckSum pulumi.StringPtrInput
	// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
	CompatibleRuntimes pulumi.StringArrayInput
	// The description of the layer version.
	Description pulumi.StringPtrInput
	// The name of the layer.
	LayerName pulumi.StringPtrInput
	// The name of the OSS bucket that stores the ZIP package of the function code.
	OssBucketName pulumi.StringPtrInput
	// The name of the OSS object (ZIP package) that contains the function code.
	OssObjectName pulumi.StringPtrInput
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrInput
	// The version of Layer Version.
	Version pulumi.StringPtrInput
	// The ZIP package of the function code that is encoded in the Base64 format.
	//
	// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
	ZipFile pulumi.StringPtrInput
}

func (LayerVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionState)(nil)).Elem()
}

type layerVersionArgs struct {
	// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// The description of the layer version.
	Description *string `pulumi:"description"`
	// The name of the layer.
	LayerName string `pulumi:"layerName"`
	// The name of the OSS bucket that stores the ZIP package of the function code.
	OssBucketName *string `pulumi:"ossBucketName"`
	// The name of the OSS object (ZIP package) that contains the function code.
	OssObjectName *string `pulumi:"ossObjectName"`
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// The ZIP package of the function code that is encoded in the Base64 format.
	//
	// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
	ZipFile *string `pulumi:"zipFile"`
}

// The set of arguments for constructing a LayerVersion resource.
type LayerVersionArgs struct {
	// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
	CompatibleRuntimes pulumi.StringArrayInput
	// The description of the layer version.
	Description pulumi.StringPtrInput
	// The name of the layer.
	LayerName pulumi.StringInput
	// The name of the OSS bucket that stores the ZIP package of the function code.
	OssBucketName pulumi.StringPtrInput
	// The name of the OSS object (ZIP package) that contains the function code.
	OssObjectName pulumi.StringPtrInput
	// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
	SkipDestroy pulumi.BoolPtrInput
	// The ZIP package of the function code that is encoded in the Base64 format.
	//
	// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
	ZipFile pulumi.StringPtrInput
}

func (LayerVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionArgs)(nil)).Elem()
}

type LayerVersionInput interface {
	pulumi.Input

	ToLayerVersionOutput() LayerVersionOutput
	ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput
}

func (*LayerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil)).Elem()
}

func (i *LayerVersion) ToLayerVersionOutput() LayerVersionOutput {
	return i.ToLayerVersionOutputWithContext(context.Background())
}

func (i *LayerVersion) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionOutput)
}

func (i *LayerVersion) ToOutput(ctx context.Context) pulumix.Output[*LayerVersion] {
	return pulumix.Output[*LayerVersion]{
		OutputState: i.ToLayerVersionOutputWithContext(ctx).OutputState,
	}
}

// LayerVersionArrayInput is an input type that accepts LayerVersionArray and LayerVersionArrayOutput values.
// You can construct a concrete instance of `LayerVersionArrayInput` via:
//
//	LayerVersionArray{ LayerVersionArgs{...} }
type LayerVersionArrayInput interface {
	pulumi.Input

	ToLayerVersionArrayOutput() LayerVersionArrayOutput
	ToLayerVersionArrayOutputWithContext(context.Context) LayerVersionArrayOutput
}

type LayerVersionArray []LayerVersionInput

func (LayerVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LayerVersion)(nil)).Elem()
}

func (i LayerVersionArray) ToLayerVersionArrayOutput() LayerVersionArrayOutput {
	return i.ToLayerVersionArrayOutputWithContext(context.Background())
}

func (i LayerVersionArray) ToLayerVersionArrayOutputWithContext(ctx context.Context) LayerVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionArrayOutput)
}

func (i LayerVersionArray) ToOutput(ctx context.Context) pulumix.Output[[]*LayerVersion] {
	return pulumix.Output[[]*LayerVersion]{
		OutputState: i.ToLayerVersionArrayOutputWithContext(ctx).OutputState,
	}
}

// LayerVersionMapInput is an input type that accepts LayerVersionMap and LayerVersionMapOutput values.
// You can construct a concrete instance of `LayerVersionMapInput` via:
//
//	LayerVersionMap{ "key": LayerVersionArgs{...} }
type LayerVersionMapInput interface {
	pulumi.Input

	ToLayerVersionMapOutput() LayerVersionMapOutput
	ToLayerVersionMapOutputWithContext(context.Context) LayerVersionMapOutput
}

type LayerVersionMap map[string]LayerVersionInput

func (LayerVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LayerVersion)(nil)).Elem()
}

func (i LayerVersionMap) ToLayerVersionMapOutput() LayerVersionMapOutput {
	return i.ToLayerVersionMapOutputWithContext(context.Background())
}

func (i LayerVersionMap) ToLayerVersionMapOutputWithContext(ctx context.Context) LayerVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionMapOutput)
}

func (i LayerVersionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LayerVersion] {
	return pulumix.Output[map[string]*LayerVersion]{
		OutputState: i.ToLayerVersionMapOutputWithContext(ctx).OutputState,
	}
}

type LayerVersionOutput struct{ *pulumi.OutputState }

func (LayerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LayerVersion)(nil)).Elem()
}

func (o LayerVersionOutput) ToLayerVersionOutput() LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToLayerVersionOutputWithContext(ctx context.Context) LayerVersionOutput {
	return o
}

func (o LayerVersionOutput) ToOutput(ctx context.Context) pulumix.Output[*LayerVersion] {
	return pulumix.Output[*LayerVersion]{
		OutputState: o.OutputState,
	}
}

// The access mode of Layer Version.
func (o LayerVersionOutput) Acl() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.Acl }).(pulumi.StringOutput)
}

// The arn of Layer Version.
func (o LayerVersionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The checksum of the layer code package.
func (o LayerVersionOutput) CodeCheckSum() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.CodeCheckSum }).(pulumi.StringOutput)
}

// The list of runtime environments that are supported by the layer. Valid values: `nodejs14`, `nodejs12`, `nodejs10`, `nodejs8`, `nodejs6`, `python3.9`, `python3`, `python2.7`, `java11`, `java8`, `php7.2`, `go1`,`dotnetcore2.1`, `custom`.
func (o LayerVersionOutput) CompatibleRuntimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringArrayOutput { return v.CompatibleRuntimes }).(pulumi.StringArrayOutput)
}

// The description of the layer version.
func (o LayerVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the layer.
func (o LayerVersionOutput) LayerName() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.LayerName }).(pulumi.StringOutput)
}

// The name of the OSS bucket that stores the ZIP package of the function code.
func (o LayerVersionOutput) OssBucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.OssBucketName }).(pulumi.StringPtrOutput)
}

// The name of the OSS object (ZIP package) that contains the function code.
func (o LayerVersionOutput) OssObjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.OssObjectName }).(pulumi.StringPtrOutput)
}

// Whether to retain the old version of a previously deployed Lambda Layer. Default is `false`. When this is not set to `true`, changing any of `compatibleRuntimes`, `description`, `layerName`, `ossBucketName`,  `ossObjectName`, or `zipFile` forces deletion of the existing layer version and creation of a new layer version.
func (o LayerVersionOutput) SkipDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.BoolPtrOutput { return v.SkipDestroy }).(pulumi.BoolPtrOutput)
}

// The version of Layer Version.
func (o LayerVersionOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The ZIP package of the function code that is encoded in the Base64 format.
//
// > **NOTE:** `zipFile` and `ossBucketName`, `ossObjectName` cannot be used together.
func (o LayerVersionOutput) ZipFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LayerVersion) pulumi.StringPtrOutput { return v.ZipFile }).(pulumi.StringPtrOutput)
}

type LayerVersionArrayOutput struct{ *pulumi.OutputState }

func (LayerVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LayerVersion)(nil)).Elem()
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutput() LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) ToLayerVersionArrayOutputWithContext(ctx context.Context) LayerVersionArrayOutput {
	return o
}

func (o LayerVersionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LayerVersion] {
	return pulumix.Output[[]*LayerVersion]{
		OutputState: o.OutputState,
	}
}

func (o LayerVersionArrayOutput) Index(i pulumi.IntInput) LayerVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LayerVersion {
		return vs[0].([]*LayerVersion)[vs[1].(int)]
	}).(LayerVersionOutput)
}

type LayerVersionMapOutput struct{ *pulumi.OutputState }

func (LayerVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LayerVersion)(nil)).Elem()
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutput() LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) ToLayerVersionMapOutputWithContext(ctx context.Context) LayerVersionMapOutput {
	return o
}

func (o LayerVersionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LayerVersion] {
	return pulumix.Output[map[string]*LayerVersion]{
		OutputState: o.OutputState,
	}
}

func (o LayerVersionMapOutput) MapIndex(k pulumi.StringInput) LayerVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LayerVersion {
		return vs[0].(map[string]*LayerVersion)[vs[1].(string)]
	}).(LayerVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionInput)(nil)).Elem(), &LayerVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionArrayInput)(nil)).Elem(), LayerVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LayerVersionMapInput)(nil)).Elem(), LayerVersionMap{})
	pulumi.RegisterOutputType(LayerVersionOutput{})
	pulumi.RegisterOutputType(LayerVersionArrayOutput{})
	pulumi.RegisterOutputType(LayerVersionMapOutput{})
}
