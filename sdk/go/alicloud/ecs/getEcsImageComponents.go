// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Image Components of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.159.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ecs.GetEcsImageComponents(ctx, &ecs.GetEcsImageComponentsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsImageComponentId1", ids.Components[0].Id)
//			nameRegex, err := ecs.GetEcsImageComponents(ctx, &ecs.GetEcsImageComponentsArgs{
//				NameRegex: pulumi.StringRef("^my-ImageComponent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsImageComponentId2", nameRegex.Components[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEcsImageComponents(ctx *pulumi.Context, args *GetEcsImageComponentsArgs, opts ...pulumi.InvokeOption) (*GetEcsImageComponentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEcsImageComponentsResult
	err := ctx.Invoke("alicloud:ecs/getEcsImageComponents:getEcsImageComponents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsImageComponents.
type GetEcsImageComponentsArgs struct {
	// A list of Image Component IDs.
	Ids []string `pulumi:"ids"`
	// The name of the image component.
	ImageComponentName *string `pulumi:"imageComponentName"`
	// A regex string to filter results by Image Component name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The type of the image component.
	Owner *string `pulumi:"owner"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// List of label key-value pairs.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getEcsImageComponents.
type GetEcsImageComponentsResult struct {
	Components []GetEcsImageComponentsComponent `pulumi:"components"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                 `pulumi:"id"`
	Ids                []string               `pulumi:"ids"`
	ImageComponentName *string                `pulumi:"imageComponentName"`
	NameRegex          *string                `pulumi:"nameRegex"`
	Names              []string               `pulumi:"names"`
	OutputFile         *string                `pulumi:"outputFile"`
	Owner              *string                `pulumi:"owner"`
	ResourceGroupId    *string                `pulumi:"resourceGroupId"`
	Tags               map[string]interface{} `pulumi:"tags"`
}

func GetEcsImageComponentsOutput(ctx *pulumi.Context, args GetEcsImageComponentsOutputArgs, opts ...pulumi.InvokeOption) GetEcsImageComponentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsImageComponentsResult, error) {
			args := v.(GetEcsImageComponentsArgs)
			r, err := GetEcsImageComponents(ctx, &args, opts...)
			var s GetEcsImageComponentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsImageComponentsResultOutput)
}

// A collection of arguments for invoking getEcsImageComponents.
type GetEcsImageComponentsOutputArgs struct {
	// A list of Image Component IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of the image component.
	ImageComponentName pulumi.StringPtrInput `pulumi:"imageComponentName"`
	// A regex string to filter results by Image Component name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The type of the image component.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// List of label key-value pairs.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetEcsImageComponentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsImageComponentsArgs)(nil)).Elem()
}

// A collection of values returned by getEcsImageComponents.
type GetEcsImageComponentsResultOutput struct{ *pulumi.OutputState }

func (GetEcsImageComponentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsImageComponentsResult)(nil)).Elem()
}

func (o GetEcsImageComponentsResultOutput) ToGetEcsImageComponentsResultOutput() GetEcsImageComponentsResultOutput {
	return o
}

func (o GetEcsImageComponentsResultOutput) ToGetEcsImageComponentsResultOutputWithContext(ctx context.Context) GetEcsImageComponentsResultOutput {
	return o
}

func (o GetEcsImageComponentsResultOutput) Components() GetEcsImageComponentsComponentArrayOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) []GetEcsImageComponentsComponent { return v.Components }).(GetEcsImageComponentsComponentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsImageComponentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsImageComponentsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsImageComponentsResultOutput) ImageComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) *string { return v.ImageComponentName }).(pulumi.StringPtrOutput)
}

func (o GetEcsImageComponentsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsImageComponentsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsImageComponentsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsImageComponentsResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o GetEcsImageComponentsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsImageComponentsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEcsImageComponentsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsImageComponentsResultOutput{})
}
