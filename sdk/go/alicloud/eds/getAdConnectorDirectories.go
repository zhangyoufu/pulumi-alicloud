// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Ad Connector Directories of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.174.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := eds.GetAdConnectorDirectories(ctx, &eds.GetAdConnectorDirectoriesArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdAdConnectorDirectoryId1", ids.Directories[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAdConnectorDirectories(ctx *pulumi.Context, args *GetAdConnectorDirectoriesArgs, opts ...pulumi.InvokeOption) (*GetAdConnectorDirectoriesResult, error) {
	var rv GetAdConnectorDirectoriesResult
	err := ctx.Invoke("alicloud:eds/getAdConnectorDirectories:getAdConnectorDirectories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAdConnectorDirectories.
type GetAdConnectorDirectoriesArgs struct {
	// A list of Ad Connector Directory IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Ad Connector Directory name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of directory.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAdConnectorDirectories.
type GetAdConnectorDirectoriesResult struct {
	Directories []GetAdConnectorDirectoriesDirectory `pulumi:"directories"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetAdConnectorDirectoriesOutput(ctx *pulumi.Context, args GetAdConnectorDirectoriesOutputArgs, opts ...pulumi.InvokeOption) GetAdConnectorDirectoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAdConnectorDirectoriesResult, error) {
			args := v.(GetAdConnectorDirectoriesArgs)
			r, err := GetAdConnectorDirectories(ctx, &args, opts...)
			var s GetAdConnectorDirectoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAdConnectorDirectoriesResultOutput)
}

// A collection of arguments for invoking getAdConnectorDirectories.
type GetAdConnectorDirectoriesOutputArgs struct {
	// A list of Ad Connector Directory IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Ad Connector Directory name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of directory.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetAdConnectorDirectoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdConnectorDirectoriesArgs)(nil)).Elem()
}

// A collection of values returned by getAdConnectorDirectories.
type GetAdConnectorDirectoriesResultOutput struct{ *pulumi.OutputState }

func (GetAdConnectorDirectoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdConnectorDirectoriesResult)(nil)).Elem()
}

func (o GetAdConnectorDirectoriesResultOutput) ToGetAdConnectorDirectoriesResultOutput() GetAdConnectorDirectoriesResultOutput {
	return o
}

func (o GetAdConnectorDirectoriesResultOutput) ToGetAdConnectorDirectoriesResultOutputWithContext(ctx context.Context) GetAdConnectorDirectoriesResultOutput {
	return o
}

func (o GetAdConnectorDirectoriesResultOutput) Directories() GetAdConnectorDirectoriesDirectoryArrayOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) []GetAdConnectorDirectoriesDirectory { return v.Directories }).(GetAdConnectorDirectoriesDirectoryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAdConnectorDirectoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAdConnectorDirectoriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAdConnectorDirectoriesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAdConnectorDirectoriesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAdConnectorDirectoriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAdConnectorDirectoriesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAdConnectorDirectoriesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAdConnectorDirectoriesResultOutput{})
}
