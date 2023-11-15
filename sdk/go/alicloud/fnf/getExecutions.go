// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fnf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the FnF Executions of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.149.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fnf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fnf.GetExecutions(ctx, &fnf.GetExecutionsArgs{
//				FlowName: "example_value",
//				Ids: []string{
//					"my-Execution-1",
//					"my-Execution-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("fnfExecutionId1", data.Alicloud_fn_f_executions.Ids.Executions[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetExecutions(ctx *pulumi.Context, args *GetExecutionsArgs, opts ...pulumi.InvokeOption) (*GetExecutionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetExecutionsResult
	err := ctx.Invoke("alicloud:fnf/getExecutions:getExecutions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExecutions.
type GetExecutionsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// The name of the flow.
	FlowName string `pulumi:"flowName"`
	// A list of Execution IDs. The value formats as `<flow_name>:<execution_name>`.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Execution name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getExecutions.
type GetExecutionsResult struct {
	EnableDetails *bool                    `pulumi:"enableDetails"`
	Executions    []GetExecutionsExecution `pulumi:"executions"`
	FlowName      string                   `pulumi:"flowName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetExecutionsOutput(ctx *pulumi.Context, args GetExecutionsOutputArgs, opts ...pulumi.InvokeOption) GetExecutionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetExecutionsResult, error) {
			args := v.(GetExecutionsArgs)
			r, err := GetExecutions(ctx, &args, opts...)
			var s GetExecutionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetExecutionsResultOutput)
}

// A collection of arguments for invoking getExecutions.
type GetExecutionsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// The name of the flow.
	FlowName pulumi.StringInput `pulumi:"flowName"`
	// A list of Execution IDs. The value formats as `<flow_name>:<execution_name>`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Execution name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetExecutionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExecutionsArgs)(nil)).Elem()
}

// A collection of values returned by getExecutions.
type GetExecutionsResultOutput struct{ *pulumi.OutputState }

func (GetExecutionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExecutionsResult)(nil)).Elem()
}

func (o GetExecutionsResultOutput) ToGetExecutionsResultOutput() GetExecutionsResultOutput {
	return o
}

func (o GetExecutionsResultOutput) ToGetExecutionsResultOutputWithContext(ctx context.Context) GetExecutionsResultOutput {
	return o
}

func (o GetExecutionsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetExecutionsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetExecutionsResultOutput) Executions() GetExecutionsExecutionArrayOutput {
	return o.ApplyT(func(v GetExecutionsResult) []GetExecutionsExecution { return v.Executions }).(GetExecutionsExecutionArrayOutput)
}

func (o GetExecutionsResultOutput) FlowName() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsResult) string { return v.FlowName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetExecutionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExecutionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetExecutionsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetExecutionsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetExecutionsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExecutionsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetExecutionsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetExecutionsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetExecutionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExecutionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetExecutionsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetExecutionsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExecutionsResultOutput{})
}
