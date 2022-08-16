// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Invocations of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.168.0+.
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
//			ids, err := ecs.GetEcsInvocations(ctx, &ecs.GetEcsInvocationsArgs{
//				Ids: []string{
//					"example-id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsInvocationId1", ids.Invocations[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEcsInvocations(ctx *pulumi.Context, args *GetEcsInvocationsArgs, opts ...pulumi.InvokeOption) (*GetEcsInvocationsResult, error) {
	var rv GetEcsInvocationsResult
	err := ctx.Invoke("alicloud:ecs/getEcsInvocations:getEcsInvocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsInvocations.
type GetEcsInvocationsArgs struct {
	// The ID of the command.
	CommandId *string `pulumi:"commandId"`
	// The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// A list of Invocation IDs.
	Ids []string `pulumi:"ids"`
	// The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocationStatus` response parameter for the overall execution state.
	InvokeStatus *string `pulumi:"invokeStatus"`
	OutputFile   *string `pulumi:"outputFile"`
	PageNumber   *int    `pulumi:"pageNumber"`
	PageSize     *int    `pulumi:"pageSize"`
}

// A collection of values returned by getEcsInvocations.
type GetEcsInvocationsResult struct {
	CommandId       *string `pulumi:"commandId"`
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The provider-assigned unique ID for this managed resource.
	Id           string                        `pulumi:"id"`
	Ids          []string                      `pulumi:"ids"`
	Invocations  []GetEcsInvocationsInvocation `pulumi:"invocations"`
	InvokeStatus *string                       `pulumi:"invokeStatus"`
	OutputFile   *string                       `pulumi:"outputFile"`
	PageNumber   *int                          `pulumi:"pageNumber"`
	PageSize     *int                          `pulumi:"pageSize"`
}

func GetEcsInvocationsOutput(ctx *pulumi.Context, args GetEcsInvocationsOutputArgs, opts ...pulumi.InvokeOption) GetEcsInvocationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsInvocationsResult, error) {
			args := v.(GetEcsInvocationsArgs)
			r, err := GetEcsInvocations(ctx, &args, opts...)
			var s GetEcsInvocationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsInvocationsResultOutput)
}

// A collection of arguments for invoking getEcsInvocations.
type GetEcsInvocationsOutputArgs struct {
	// The ID of the command.
	CommandId pulumi.StringPtrInput `pulumi:"commandId"`
	// The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
	ContentEncoding pulumi.StringPtrInput `pulumi:"contentEncoding"`
	// A list of Invocation IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocationStatus` response parameter for the overall execution state.
	InvokeStatus pulumi.StringPtrInput `pulumi:"invokeStatus"`
	OutputFile   pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber   pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize     pulumi.IntPtrInput    `pulumi:"pageSize"`
}

func (GetEcsInvocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsInvocationsArgs)(nil)).Elem()
}

// A collection of values returned by getEcsInvocations.
type GetEcsInvocationsResultOutput struct{ *pulumi.OutputState }

func (GetEcsInvocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsInvocationsResult)(nil)).Elem()
}

func (o GetEcsInvocationsResultOutput) ToGetEcsInvocationsResultOutput() GetEcsInvocationsResultOutput {
	return o
}

func (o GetEcsInvocationsResultOutput) ToGetEcsInvocationsResultOutputWithContext(ctx context.Context) GetEcsInvocationsResultOutput {
	return o
}

func (o GetEcsInvocationsResultOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

func (o GetEcsInvocationsResultOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *string { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsInvocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsInvocationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsInvocationsResultOutput) Invocations() GetEcsInvocationsInvocationArrayOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) []GetEcsInvocationsInvocation { return v.Invocations }).(GetEcsInvocationsInvocationArrayOutput)
}

func (o GetEcsInvocationsResultOutput) InvokeStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *string { return v.InvokeStatus }).(pulumi.StringPtrOutput)
}

func (o GetEcsInvocationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsInvocationsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetEcsInvocationsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetEcsInvocationsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsInvocationsResultOutput{})
}
