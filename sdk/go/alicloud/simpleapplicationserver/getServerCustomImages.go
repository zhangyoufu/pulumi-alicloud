// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleapplicationserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Simple Application Server Custom Images of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.143.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/simpleapplicationserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := simpleapplicationserver.GetServerCustomImages(ctx, &simpleapplicationserver.GetServerCustomImagesArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerCustomImageId1", ids.Images[0].Id)
//			nameRegex, err := simpleapplicationserver.GetServerCustomImages(ctx, &simpleapplicationserver.GetServerCustomImagesArgs{
//				NameRegex: pulumi.StringRef("^my-CustomImage"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerCustomImageId2", nameRegex.Images[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetServerCustomImages(ctx *pulumi.Context, args *GetServerCustomImagesArgs, opts ...pulumi.InvokeOption) (*GetServerCustomImagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerCustomImagesResult
	err := ctx.Invoke("alicloud:simpleapplicationserver/getServerCustomImages:getServerCustomImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerCustomImages.
type GetServerCustomImagesArgs struct {
	// A list of Custom Image IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Custom Image name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getServerCustomImages.
type GetServerCustomImagesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                       `pulumi:"id"`
	Ids        []string                     `pulumi:"ids"`
	Images     []GetServerCustomImagesImage `pulumi:"images"`
	NameRegex  *string                      `pulumi:"nameRegex"`
	Names      []string                     `pulumi:"names"`
	OutputFile *string                      `pulumi:"outputFile"`
}

func GetServerCustomImagesOutput(ctx *pulumi.Context, args GetServerCustomImagesOutputArgs, opts ...pulumi.InvokeOption) GetServerCustomImagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerCustomImagesResult, error) {
			args := v.(GetServerCustomImagesArgs)
			r, err := GetServerCustomImages(ctx, &args, opts...)
			var s GetServerCustomImagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerCustomImagesResultOutput)
}

// A collection of arguments for invoking getServerCustomImages.
type GetServerCustomImagesOutputArgs struct {
	// A list of Custom Image IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Custom Image name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetServerCustomImagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerCustomImagesArgs)(nil)).Elem()
}

// A collection of values returned by getServerCustomImages.
type GetServerCustomImagesResultOutput struct{ *pulumi.OutputState }

func (GetServerCustomImagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerCustomImagesResult)(nil)).Elem()
}

func (o GetServerCustomImagesResultOutput) ToGetServerCustomImagesResultOutput() GetServerCustomImagesResultOutput {
	return o
}

func (o GetServerCustomImagesResultOutput) ToGetServerCustomImagesResultOutputWithContext(ctx context.Context) GetServerCustomImagesResultOutput {
	return o
}

func (o GetServerCustomImagesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetServerCustomImagesResult] {
	return pulumix.Output[GetServerCustomImagesResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerCustomImagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerCustomImagesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServerCustomImagesResultOutput) Images() GetServerCustomImagesImageArrayOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) []GetServerCustomImagesImage { return v.Images }).(GetServerCustomImagesImageArrayOutput)
}

func (o GetServerCustomImagesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetServerCustomImagesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServerCustomImagesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerCustomImagesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerCustomImagesResultOutput{})
}
