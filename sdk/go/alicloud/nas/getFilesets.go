// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Nas Filesets of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.153.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := nas.GetFilesets(ctx, &nas.GetFilesetsArgs{
//				FileSystemId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nasFilesetId1", ids.Filesets[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetFilesets(ctx *pulumi.Context, args *GetFilesetsArgs, opts ...pulumi.InvokeOption) (*GetFilesetsResult, error) {
	var rv GetFilesetsResult
	err := ctx.Invoke("alicloud:nas/getFilesets:getFilesets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFilesets.
type GetFilesetsArgs struct {
	// The ID of the file system.
	FileSystemId string `pulumi:"fileSystemId"`
	// A list of Fileset IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the fileset.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getFilesets.
type GetFilesetsResult struct {
	FileSystemId string               `pulumi:"fileSystemId"`
	Filesets     []GetFilesetsFileset `pulumi:"filesets"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetFilesetsOutput(ctx *pulumi.Context, args GetFilesetsOutputArgs, opts ...pulumi.InvokeOption) GetFilesetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFilesetsResult, error) {
			args := v.(GetFilesetsArgs)
			r, err := GetFilesets(ctx, &args, opts...)
			var s GetFilesetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFilesetsResultOutput)
}

// A collection of arguments for invoking getFilesets.
type GetFilesetsOutputArgs struct {
	// The ID of the file system.
	FileSystemId pulumi.StringInput `pulumi:"fileSystemId"`
	// A list of Fileset IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the fileset.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetFilesetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFilesetsArgs)(nil)).Elem()
}

// A collection of values returned by getFilesets.
type GetFilesetsResultOutput struct{ *pulumi.OutputState }

func (GetFilesetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFilesetsResult)(nil)).Elem()
}

func (o GetFilesetsResultOutput) ToGetFilesetsResultOutput() GetFilesetsResultOutput {
	return o
}

func (o GetFilesetsResultOutput) ToGetFilesetsResultOutputWithContext(ctx context.Context) GetFilesetsResultOutput {
	return o
}

func (o GetFilesetsResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFilesetsResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o GetFilesetsResultOutput) Filesets() GetFilesetsFilesetArrayOutput {
	return o.ApplyT(func(v GetFilesetsResult) []GetFilesetsFileset { return v.Filesets }).(GetFilesetsFilesetArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFilesetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFilesetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFilesetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFilesetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetFilesetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFilesetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetFilesetsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFilesetsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFilesetsResultOutput{})
}
