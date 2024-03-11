// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Oos Application Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.146.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := oos.GetApplicationGroups(ctx, &oos.GetApplicationGroupsArgs{
//				ApplicationName: "example_value",
//				Ids: []string{
//					"my-ApplicationGroup-1",
//					"my-ApplicationGroup-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("oosApplicationGroupId1", ids.Groups[0].Id)
//			nameRegex, err := oos.GetApplicationGroups(ctx, &oos.GetApplicationGroupsArgs{
//				ApplicationName: "example_value",
//				NameRegex:       pulumi.StringRef("^my-ApplicationGroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("oosApplicationGroupId2", nameRegex.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetApplicationGroups(ctx *pulumi.Context, args *GetApplicationGroupsArgs, opts ...pulumi.InvokeOption) (*GetApplicationGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationGroupsResult
	err := ctx.Invoke("alicloud:oos/getApplicationGroups:getApplicationGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationGroups.
type GetApplicationGroupsArgs struct {
	// The name of the Application.
	ApplicationName string `pulumi:"applicationName"`
	// The region ID of the deployment.
	DeployRegionId *string `pulumi:"deployRegionId"`
	// A list of Application Group IDs. Its element value is same as Application Group Name.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Application Group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getApplicationGroups.
type GetApplicationGroupsResult struct {
	ApplicationName string                      `pulumi:"applicationName"`
	DeployRegionId  *string                     `pulumi:"deployRegionId"`
	Groups          []GetApplicationGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetApplicationGroupsOutput(ctx *pulumi.Context, args GetApplicationGroupsOutputArgs, opts ...pulumi.InvokeOption) GetApplicationGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationGroupsResult, error) {
			args := v.(GetApplicationGroupsArgs)
			r, err := GetApplicationGroups(ctx, &args, opts...)
			var s GetApplicationGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationGroupsResultOutput)
}

// A collection of arguments for invoking getApplicationGroups.
type GetApplicationGroupsOutputArgs struct {
	// The name of the Application.
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The region ID of the deployment.
	DeployRegionId pulumi.StringPtrInput `pulumi:"deployRegionId"`
	// A list of Application Group IDs. Its element value is same as Application Group Name.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Application Group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetApplicationGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationGroups.
type GetApplicationGroupsResultOutput struct{ *pulumi.OutputState }

func (GetApplicationGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGroupsResult)(nil)).Elem()
}

func (o GetApplicationGroupsResultOutput) ToGetApplicationGroupsResultOutput() GetApplicationGroupsResultOutput {
	return o
}

func (o GetApplicationGroupsResultOutput) ToGetApplicationGroupsResultOutputWithContext(ctx context.Context) GetApplicationGroupsResultOutput {
	return o
}

func (o GetApplicationGroupsResultOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o GetApplicationGroupsResultOutput) DeployRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) *string { return v.DeployRegionId }).(pulumi.StringPtrOutput)
}

func (o GetApplicationGroupsResultOutput) Groups() GetApplicationGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) []GetApplicationGroupsGroup { return v.Groups }).(GetApplicationGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApplicationGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetApplicationGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetApplicationGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetApplicationGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationGroupsResultOutput{})
}
