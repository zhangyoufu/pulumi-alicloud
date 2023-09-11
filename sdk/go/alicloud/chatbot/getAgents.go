// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chatbot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Chatbot Agents of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.203.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/chatbot"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nameRegex, err := chatbot.GetAgents(ctx, &chatbot.GetAgentsArgs{
//				NameRegex: pulumi.StringRef("^my-Agent"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudChatbotAgentsId1", nameRegex.Agents[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAgents(ctx *pulumi.Context, args *GetAgentsArgs, opts ...pulumi.InvokeOption) (*GetAgentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAgentsResult
	err := ctx.Invoke("alicloud:chatbot/getAgents:getAgents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAgents.
type GetAgentsArgs struct {
	// The name of the agent.
	AgentName *string `pulumi:"agentName"`
	// A regex string to filter resulting chatbot agents by name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
}

// A collection of values returned by getAgents.
type GetAgentsResult struct {
	// The agent Name.
	AgentName string `pulumi:"agentName"`
	// A list of availability zones. Each element contains the following attributes:
	Agents []GetAgentsAgent `pulumi:"agents"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of chatbot agents names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
}

func GetAgentsOutput(ctx *pulumi.Context, args GetAgentsOutputArgs, opts ...pulumi.InvokeOption) GetAgentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAgentsResult, error) {
			args := v.(GetAgentsArgs)
			r, err := GetAgents(ctx, &args, opts...)
			var s GetAgentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAgentsResultOutput)
}

// A collection of arguments for invoking getAgents.
type GetAgentsOutputArgs struct {
	// The name of the agent.
	AgentName pulumi.StringPtrInput `pulumi:"agentName"`
	// A regex string to filter resulting chatbot agents by name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
}

func (GetAgentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentsArgs)(nil)).Elem()
}

// A collection of values returned by getAgents.
type GetAgentsResultOutput struct{ *pulumi.OutputState }

func (GetAgentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAgentsResult)(nil)).Elem()
}

func (o GetAgentsResultOutput) ToGetAgentsResultOutput() GetAgentsResultOutput {
	return o
}

func (o GetAgentsResultOutput) ToGetAgentsResultOutputWithContext(ctx context.Context) GetAgentsResultOutput {
	return o
}

func (o GetAgentsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetAgentsResult] {
	return pulumix.Output[GetAgentsResult]{
		OutputState: o.OutputState,
	}
}

// The agent Name.
func (o GetAgentsResultOutput) AgentName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsResult) string { return v.AgentName }).(pulumi.StringOutput)
}

// A list of availability zones. Each element contains the following attributes:
func (o GetAgentsResultOutput) Agents() GetAgentsAgentArrayOutput {
	return o.ApplyT(func(v GetAgentsResult) []GetAgentsAgent { return v.Agents }).(GetAgentsAgentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAgentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAgentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAgentsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAgentsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAgentsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAgentsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of chatbot agents names.
func (o GetAgentsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAgentsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAgentsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAgentsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAgentsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAgentsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetAgentsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAgentsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAgentsResultOutput{})
}
