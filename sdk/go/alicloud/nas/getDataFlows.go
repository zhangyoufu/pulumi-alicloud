// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Nas Data Flows of the current Alibaba Cloud user.
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
//			ids, err := nas.GetDataFlows(ctx, &nas.GetDataFlowsArgs{
//				FileSystemId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nasDataFlowId1", ids.Flows[0].Id)
//			status, err := nas.GetDataFlows(ctx, &nas.GetDataFlowsArgs{
//				FileSystemId: "example_value",
//				Status:       pulumi.StringRef("Running"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nasDataFlowId2", status.Flows[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetDataFlows(ctx *pulumi.Context, args *GetDataFlowsArgs, opts ...pulumi.InvokeOption) (*GetDataFlowsResult, error) {
	var rv GetDataFlowsResult
	err := ctx.Invoke("alicloud:nas/getDataFlows:getDataFlows", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataFlows.
type GetDataFlowsArgs struct {
	// The ID of the file system.
	FileSystemId string `pulumi:"fileSystemId"`
	// A list of Data Flow IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the Data flow.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDataFlows.
type GetDataFlowsResult struct {
	FileSystemId string             `pulumi:"fileSystemId"`
	Flows        []GetDataFlowsFlow `pulumi:"flows"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetDataFlowsOutput(ctx *pulumi.Context, args GetDataFlowsOutputArgs, opts ...pulumi.InvokeOption) GetDataFlowsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDataFlowsResult, error) {
			args := v.(GetDataFlowsArgs)
			r, err := GetDataFlows(ctx, &args, opts...)
			var s GetDataFlowsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDataFlowsResultOutput)
}

// A collection of arguments for invoking getDataFlows.
type GetDataFlowsOutputArgs struct {
	// The ID of the file system.
	FileSystemId pulumi.StringInput `pulumi:"fileSystemId"`
	// A list of Data Flow IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The status of the Data flow.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetDataFlowsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataFlowsArgs)(nil)).Elem()
}

// A collection of values returned by getDataFlows.
type GetDataFlowsResultOutput struct{ *pulumi.OutputState }

func (GetDataFlowsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataFlowsResult)(nil)).Elem()
}

func (o GetDataFlowsResultOutput) ToGetDataFlowsResultOutput() GetDataFlowsResultOutput {
	return o
}

func (o GetDataFlowsResultOutput) ToGetDataFlowsResultOutputWithContext(ctx context.Context) GetDataFlowsResultOutput {
	return o
}

func (o GetDataFlowsResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataFlowsResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o GetDataFlowsResultOutput) Flows() GetDataFlowsFlowArrayOutput {
	return o.ApplyT(func(v GetDataFlowsResult) []GetDataFlowsFlow { return v.Flows }).(GetDataFlowsFlowArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDataFlowsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataFlowsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDataFlowsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDataFlowsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDataFlowsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDataFlowsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDataFlowsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDataFlowsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDataFlowsResultOutput{})
}
