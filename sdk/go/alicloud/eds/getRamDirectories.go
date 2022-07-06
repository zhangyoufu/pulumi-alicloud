// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Ram Directories of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.174.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := eds.GetRamDirectories(ctx, &eds.GetRamDirectoriesArgs{
// 			Ids: []string{
// 				"example_id",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ecdRamDirectoryId1", ids.Directories[0].Id)
// 		nameRegex, err := eds.GetRamDirectories(ctx, &eds.GetRamDirectoriesArgs{
// 			NameRegex: pulumi.StringRef("^my-RamDirectory"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ecdRamDirectoryId2", nameRegex.Directories[0].Id)
// 		return nil
// 	})
// }
// ```
func GetRamDirectories(ctx *pulumi.Context, args *GetRamDirectoriesArgs, opts ...pulumi.InvokeOption) (*GetRamDirectoriesResult, error) {
	var rv GetRamDirectoriesResult
	err := ctx.Invoke("alicloud:eds/getRamDirectories:getRamDirectories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRamDirectories.
type GetRamDirectoriesArgs struct {
	// A list of Ram Directory IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Ram Directory name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of directory.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getRamDirectories.
type GetRamDirectoriesResult struct {
	Directories []GetRamDirectoriesDirectory `pulumi:"directories"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetRamDirectoriesOutput(ctx *pulumi.Context, args GetRamDirectoriesOutputArgs, opts ...pulumi.InvokeOption) GetRamDirectoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRamDirectoriesResult, error) {
			args := v.(GetRamDirectoriesArgs)
			r, err := GetRamDirectories(ctx, &args, opts...)
			var s GetRamDirectoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRamDirectoriesResultOutput)
}

// A collection of arguments for invoking getRamDirectories.
type GetRamDirectoriesOutputArgs struct {
	// A list of Ram Directory IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Ram Directory name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of directory.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetRamDirectoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRamDirectoriesArgs)(nil)).Elem()
}

// A collection of values returned by getRamDirectories.
type GetRamDirectoriesResultOutput struct{ *pulumi.OutputState }

func (GetRamDirectoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRamDirectoriesResult)(nil)).Elem()
}

func (o GetRamDirectoriesResultOutput) ToGetRamDirectoriesResultOutput() GetRamDirectoriesResultOutput {
	return o
}

func (o GetRamDirectoriesResultOutput) ToGetRamDirectoriesResultOutputWithContext(ctx context.Context) GetRamDirectoriesResultOutput {
	return o
}

func (o GetRamDirectoriesResultOutput) Directories() GetRamDirectoriesDirectoryArrayOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) []GetRamDirectoriesDirectory { return v.Directories }).(GetRamDirectoriesDirectoryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRamDirectoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRamDirectoriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRamDirectoriesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetRamDirectoriesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRamDirectoriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRamDirectoriesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRamDirectoriesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRamDirectoriesResultOutput{})
}
