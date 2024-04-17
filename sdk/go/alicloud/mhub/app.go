// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mhub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MHUB App resource.
//
// For information about MHUB App and how to use it, see [What is App](https://help.aliyun.com/product/65109.html).
//
// > **NOTE:** Available since v1.138.0+.
//
// > **NOTE:** At present, the resource only supports cn-shanghai region.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mhub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "example_value"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := mhub.NewProduct(ctx, "default", &mhub.ProductArgs{
//				ProductName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mhub.NewApp(ctx, "default", &mhub.AppArgs{
//				AppName:     pulumi.String(name),
//				ProductId:   _default.ID(),
//				PackageName: pulumi.String("com.example.android"),
//				Type:        pulumi.String("Android"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// MHUB App can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:mhub/app:App example <product_id>:<app_key>
// ```
type App struct {
	pulumi.CustomResourceState

	// AppName.
	AppName pulumi.StringOutput `pulumi:"appName"`
	// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
	BundleId pulumi.StringPtrOutput `pulumi:"bundleId"`
	// Base64 string of picture.
	EncodedIcon pulumi.StringPtrOutput `pulumi:"encodedIcon"`
	// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
	IndustryId pulumi.StringOutput `pulumi:"industryId"`
	// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
	PackageName pulumi.StringPtrOutput `pulumi:"packageName"`
	// The ID of the Product.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The type of the Product. Valid values: `Android` and `iOS`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("alicloud:mhub/app:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("alicloud:mhub/app:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	// AppName.
	AppName *string `pulumi:"appName"`
	// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
	BundleId *string `pulumi:"bundleId"`
	// Base64 string of picture.
	EncodedIcon *string `pulumi:"encodedIcon"`
	// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
	IndustryId *string `pulumi:"industryId"`
	// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
	PackageName *string `pulumi:"packageName"`
	// The ID of the Product.
	ProductId *string `pulumi:"productId"`
	// The type of the Product. Valid values: `Android` and `iOS`.
	Type *string `pulumi:"type"`
}

type AppState struct {
	// AppName.
	AppName pulumi.StringPtrInput
	// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
	BundleId pulumi.StringPtrInput
	// Base64 string of picture.
	EncodedIcon pulumi.StringPtrInput
	// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
	IndustryId pulumi.StringPtrInput
	// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
	PackageName pulumi.StringPtrInput
	// The ID of the Product.
	ProductId pulumi.StringPtrInput
	// The type of the Product. Valid values: `Android` and `iOS`.
	Type pulumi.StringPtrInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// AppName.
	AppName string `pulumi:"appName"`
	// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
	BundleId *string `pulumi:"bundleId"`
	// Base64 string of picture.
	EncodedIcon *string `pulumi:"encodedIcon"`
	// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
	IndustryId *string `pulumi:"industryId"`
	// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
	PackageName *string `pulumi:"packageName"`
	// The ID of the Product.
	ProductId string `pulumi:"productId"`
	// The type of the Product. Valid values: `Android` and `iOS`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// AppName.
	AppName pulumi.StringInput
	// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
	BundleId pulumi.StringPtrInput
	// Base64 string of picture.
	EncodedIcon pulumi.StringPtrInput
	// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
	IndustryId pulumi.StringPtrInput
	// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
	PackageName pulumi.StringPtrInput
	// The ID of the Product.
	ProductId pulumi.StringInput
	// The type of the Product. Valid values: `Android` and `iOS`.
	Type pulumi.StringInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

// AppArrayInput is an input type that accepts AppArray and AppArrayOutput values.
// You can construct a concrete instance of `AppArrayInput` via:
//
//	AppArray{ AppArgs{...} }
type AppArrayInput interface {
	pulumi.Input

	ToAppArrayOutput() AppArrayOutput
	ToAppArrayOutputWithContext(context.Context) AppArrayOutput
}

type AppArray []AppInput

func (AppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (i AppArray) ToAppArrayOutput() AppArrayOutput {
	return i.ToAppArrayOutputWithContext(context.Background())
}

func (i AppArray) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppArrayOutput)
}

// AppMapInput is an input type that accepts AppMap and AppMapOutput values.
// You can construct a concrete instance of `AppMapInput` via:
//
//	AppMap{ "key": AppArgs{...} }
type AppMapInput interface {
	pulumi.Input

	ToAppMapOutput() AppMapOutput
	ToAppMapOutputWithContext(context.Context) AppMapOutput
}

type AppMap map[string]AppInput

func (AppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (i AppMap) ToAppMapOutput() AppMapOutput {
	return i.ToAppMapOutputWithContext(context.Background())
}

func (i AppMap) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppMapOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

// AppName.
func (o AppOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.AppName }).(pulumi.StringOutput)
}

// The app id of iOS. **NOTE:** Either `bundleId` or `packageName` must be set.
func (o AppOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.BundleId }).(pulumi.StringPtrOutput)
}

// Base64 string of picture.
func (o AppOutput) EncodedIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.EncodedIcon }).(pulumi.StringPtrOutput)
}

// The Industry ID of the app. For information about Industry and how to use it, MHUB[Industry](https://help.aliyun.com/document_detail/201638.html).
func (o AppOutput) IndustryId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.IndustryId }).(pulumi.StringOutput)
}

// Android App package name. **NOTE:** Either `bundleId` or `packageName` must be set.
func (o AppOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.PackageName }).(pulumi.StringPtrOutput)
}

// The ID of the Product.
func (o AppOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

// The type of the Product. Valid values: `Android` and `iOS`.
func (o AppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AppArrayOutput struct{ *pulumi.OutputState }

func (AppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (o AppArrayOutput) ToAppArrayOutput() AppArrayOutput {
	return o
}

func (o AppArrayOutput) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return o
}

func (o AppArrayOutput) Index(i pulumi.IntInput) AppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *App {
		return vs[0].([]*App)[vs[1].(int)]
	}).(AppOutput)
}

type AppMapOutput struct{ *pulumi.OutputState }

func (AppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (o AppMapOutput) ToAppMapOutput() AppMapOutput {
	return o
}

func (o AppMapOutput) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return o
}

func (o AppMapOutput) MapIndex(k pulumi.StringInput) AppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *App {
		return vs[0].(map[string]*App)[vs[1].(string)]
	}).(AppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInput)(nil)).Elem(), &App{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppArrayInput)(nil)).Elem(), AppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppMapInput)(nil)).Elem(), AppMap{})
	pulumi.RegisterOutputType(AppOutput{})
	pulumi.RegisterOutputType(AppArrayOutput{})
	pulumi.RegisterOutputType(AppMapOutput{})
}
