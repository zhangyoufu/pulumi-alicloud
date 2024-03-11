// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cddc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cddc Dedicated Host Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cddc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := cddc.GetDedicatedHostGroups(ctx, &cddc.GetDedicatedHostGroupsArgs{
//				Engine: pulumi.StringRef("MongoDB"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cddcDedicatedHostGroupId", _default.Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetDedicatedHostGroups(ctx *pulumi.Context, args *GetDedicatedHostGroupsArgs, opts ...pulumi.InvokeOption) (*GetDedicatedHostGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDedicatedHostGroupsResult
	err := ctx.Invoke("alicloud:cddc/getDedicatedHostGroups:getDedicatedHostGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedHostGroups.
type GetDedicatedHostGroupsArgs struct {
	// Database Engine Type.The database engine of the dedicated cluster. Valid values:`Redis`, `SQLServer`, `MySQL`, `PostgreSQL`, `MongoDB`
	Engine *string `pulumi:"engine"`
	// A list of Dedicated Host Group IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Dedicated Host Group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getDedicatedHostGroups.
type GetDedicatedHostGroupsResult struct {
	Engine *string                       `pulumi:"engine"`
	Groups []GetDedicatedHostGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetDedicatedHostGroupsOutput(ctx *pulumi.Context, args GetDedicatedHostGroupsOutputArgs, opts ...pulumi.InvokeOption) GetDedicatedHostGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDedicatedHostGroupsResult, error) {
			args := v.(GetDedicatedHostGroupsArgs)
			r, err := GetDedicatedHostGroups(ctx, &args, opts...)
			var s GetDedicatedHostGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDedicatedHostGroupsResultOutput)
}

// A collection of arguments for invoking getDedicatedHostGroups.
type GetDedicatedHostGroupsOutputArgs struct {
	// Database Engine Type.The database engine of the dedicated cluster. Valid values:`Redis`, `SQLServer`, `MySQL`, `PostgreSQL`, `MongoDB`
	Engine pulumi.StringPtrInput `pulumi:"engine"`
	// A list of Dedicated Host Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Dedicated Host Group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetDedicatedHostGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDedicatedHostGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getDedicatedHostGroups.
type GetDedicatedHostGroupsResultOutput struct{ *pulumi.OutputState }

func (GetDedicatedHostGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDedicatedHostGroupsResult)(nil)).Elem()
}

func (o GetDedicatedHostGroupsResultOutput) ToGetDedicatedHostGroupsResultOutput() GetDedicatedHostGroupsResultOutput {
	return o
}

func (o GetDedicatedHostGroupsResultOutput) ToGetDedicatedHostGroupsResultOutputWithContext(ctx context.Context) GetDedicatedHostGroupsResultOutput {
	return o
}

func (o GetDedicatedHostGroupsResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o GetDedicatedHostGroupsResultOutput) Groups() GetDedicatedHostGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) []GetDedicatedHostGroupsGroup { return v.Groups }).(GetDedicatedHostGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDedicatedHostGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDedicatedHostGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDedicatedHostGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetDedicatedHostGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetDedicatedHostGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDedicatedHostGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDedicatedHostGroupsResultOutput{})
}
