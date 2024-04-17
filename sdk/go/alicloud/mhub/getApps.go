// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Mhub Apps of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.138.0+.
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
//			_, err := mhub.NewApp(ctx, "default", &mhub.AppArgs{
//				AppName:     pulumi.String(name),
//				ProductId:   pulumi.Any(defaultAlicloudMhubProduct.Id),
//				PackageName: pulumi.String("com.test.android"),
//				Type:        pulumi.String("2"),
//			})
//			if err != nil {
//				return err
//			}
//			ids, err := mhub.GetApps(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("mhubAppId1", ids.Apps[0].Id)
//			nameRegex, err := mhub.GetApps(ctx, &mhub.GetAppsArgs{
//				NameRegex: pulumi.StringRef("^my-App"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("mhubAppId2", nameRegex.Apps[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetApps(ctx *pulumi.Context, args *GetAppsArgs, opts ...pulumi.InvokeOption) (*GetAppsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppsResult
	err := ctx.Invoke("alicloud:mhub/getApps:getApps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApps.
type GetAppsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of App IDs. The value formats as `<product_id>:<app_key>`
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by App name.
	NameRegex *string `pulumi:"nameRegex"`
	// The os type. Valid values: `Android` and `iOS`.
	OsType *string `pulumi:"osType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the Product.
	ProductId string `pulumi:"productId"`
}

// A collection of values returned by getApps.
type GetAppsResult struct {
	Apps          []GetAppsApp `pulumi:"apps"`
	EnableDetails *bool        `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OsType     *string  `pulumi:"osType"`
	OutputFile *string  `pulumi:"outputFile"`
	ProductId  string   `pulumi:"productId"`
}

func GetAppsOutput(ctx *pulumi.Context, args GetAppsOutputArgs, opts ...pulumi.InvokeOption) GetAppsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppsResult, error) {
			args := v.(GetAppsArgs)
			r, err := GetApps(ctx, &args, opts...)
			var s GetAppsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppsResultOutput)
}

// A collection of arguments for invoking getApps.
type GetAppsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of App IDs. The value formats as `<product_id>:<app_key>`
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by App name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The os type. Valid values: `Android` and `iOS`.
	OsType pulumi.StringPtrInput `pulumi:"osType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the Product.
	ProductId pulumi.StringInput `pulumi:"productId"`
}

func (GetAppsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsArgs)(nil)).Elem()
}

// A collection of values returned by getApps.
type GetAppsResultOutput struct{ *pulumi.OutputState }

func (GetAppsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppsResult)(nil)).Elem()
}

func (o GetAppsResultOutput) ToGetAppsResultOutput() GetAppsResultOutput {
	return o
}

func (o GetAppsResultOutput) ToGetAppsResultOutputWithContext(ctx context.Context) GetAppsResultOutput {
	return o
}

func (o GetAppsResultOutput) Apps() GetAppsAppArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []GetAppsApp { return v.Apps }).(GetAppsAppArrayOutput)
}

func (o GetAppsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAppsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAppsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAppsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAppsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAppsResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetAppsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAppsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAppsResultOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppsResult) string { return v.ProductId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppsResultOutput{})
}
