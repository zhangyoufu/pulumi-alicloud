// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Launch Templates of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.120.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ecs.GetEcsLaunchTemplates(ctx, &ecs.GetEcsLaunchTemplatesArgs{
//				Ids: []string{
//					"lt-bp1a469uxxxxxx",
//				},
//				NameRegex: pulumi.StringRef("your_launch_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcsLaunchTemplateId", example.Templates[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEcsLaunchTemplates(ctx *pulumi.Context, args *GetEcsLaunchTemplatesArgs, opts ...pulumi.InvokeOption) (*GetEcsLaunchTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEcsLaunchTemplatesResult
	err := ctx.Invoke("alicloud:ecs/getEcsLaunchTemplates:getEcsLaunchTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsLaunchTemplates.
type GetEcsLaunchTemplatesArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Launch Template IDs.
	Ids []string `pulumi:"ids"`
	// The Launch Template Name.
	LaunchTemplateName *string `pulumi:"launchTemplateName"`
	// A regex string to filter results by Launch Template name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The template resource group id.
	TemplateResourceGroupId *string `pulumi:"templateResourceGroupId"`
	// The template tags.
	TemplateTags map[string]interface{} `pulumi:"templateTags"`
}

// A collection of values returned by getEcsLaunchTemplates.
type GetEcsLaunchTemplatesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                          `pulumi:"id"`
	Ids                     []string                        `pulumi:"ids"`
	LaunchTemplateName      *string                         `pulumi:"launchTemplateName"`
	NameRegex               *string                         `pulumi:"nameRegex"`
	Names                   []string                        `pulumi:"names"`
	OutputFile              *string                         `pulumi:"outputFile"`
	TemplateResourceGroupId *string                         `pulumi:"templateResourceGroupId"`
	TemplateTags            map[string]interface{}          `pulumi:"templateTags"`
	Templates               []GetEcsLaunchTemplatesTemplate `pulumi:"templates"`
}

func GetEcsLaunchTemplatesOutput(ctx *pulumi.Context, args GetEcsLaunchTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetEcsLaunchTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsLaunchTemplatesResult, error) {
			args := v.(GetEcsLaunchTemplatesArgs)
			r, err := GetEcsLaunchTemplates(ctx, &args, opts...)
			var s GetEcsLaunchTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsLaunchTemplatesResultOutput)
}

// A collection of arguments for invoking getEcsLaunchTemplates.
type GetEcsLaunchTemplatesOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Launch Template IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Launch Template Name.
	LaunchTemplateName pulumi.StringPtrInput `pulumi:"launchTemplateName"`
	// A regex string to filter results by Launch Template name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The template resource group id.
	TemplateResourceGroupId pulumi.StringPtrInput `pulumi:"templateResourceGroupId"`
	// The template tags.
	TemplateTags pulumi.MapInput `pulumi:"templateTags"`
}

func (GetEcsLaunchTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsLaunchTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getEcsLaunchTemplates.
type GetEcsLaunchTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetEcsLaunchTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsLaunchTemplatesResult)(nil)).Elem()
}

func (o GetEcsLaunchTemplatesResultOutput) ToGetEcsLaunchTemplatesResultOutput() GetEcsLaunchTemplatesResultOutput {
	return o
}

func (o GetEcsLaunchTemplatesResultOutput) ToGetEcsLaunchTemplatesResultOutputWithContext(ctx context.Context) GetEcsLaunchTemplatesResultOutput {
	return o
}

func (o GetEcsLaunchTemplatesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsLaunchTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) LaunchTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) *string { return v.LaunchTemplateName }).(pulumi.StringPtrOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) TemplateResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) *string { return v.TemplateResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) TemplateTags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) map[string]interface{} { return v.TemplateTags }).(pulumi.MapOutput)
}

func (o GetEcsLaunchTemplatesResultOutput) Templates() GetEcsLaunchTemplatesTemplateArrayOutput {
	return o.ApplyT(func(v GetEcsLaunchTemplatesResult) []GetEcsLaunchTemplatesTemplate { return v.Templates }).(GetEcsLaunchTemplatesTemplateArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsLaunchTemplatesResultOutput{})
}
