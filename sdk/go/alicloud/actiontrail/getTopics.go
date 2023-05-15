// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of ALIKAFKA Topics in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.56.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/actiontrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			topicsDs, err := actiontrail.GetTopics(ctx, &actiontrail.GetTopicsArgs{
//				InstanceId: "xxx",
//				NameRegex:  pulumi.StringRef("alikafkaTopicName"),
//				OutputFile: pulumi.StringRef("topics.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstTopicName", topicsDs.Topics[0].Topic)
//			return nil
//		})
//	}
//
// ```
func GetTopics(ctx *pulumi.Context, args *GetTopicsArgs, opts ...pulumi.InvokeOption) (*GetTopicsResult, error) {
	var rv GetTopicsResult
	err := ctx.Invoke("alicloud:actiontrail/getTopics:getTopics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopics.
type GetTopicsArgs struct {
	// A list of ALIKAFKA Topics IDs, It is formatted to `<instance_id>:<topic>`.
	Ids []string `pulumi:"ids"`
	// ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by the topic name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// A topic to filter results by the topic name.
	Topic *string `pulumi:"topic"`
}

// A collection of values returned by getTopics.
type GetTopicsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The instanceId of the instance.
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of topic names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	// The name of the topic.
	Topic *string `pulumi:"topic"`
	// A list of topics. Each element contains the following attributes:
	Topics     []GetTopicsTopic `pulumi:"topics"`
	TotalCount int              `pulumi:"totalCount"`
}

func GetTopicsOutput(ctx *pulumi.Context, args GetTopicsOutputArgs, opts ...pulumi.InvokeOption) GetTopicsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTopicsResult, error) {
			args := v.(GetTopicsArgs)
			r, err := GetTopics(ctx, &args, opts...)
			var s GetTopicsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTopicsResultOutput)
}

// A collection of arguments for invoking getTopics.
type GetTopicsOutputArgs struct {
	// A list of ALIKAFKA Topics IDs, It is formatted to `<instance_id>:<topic>`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by the topic name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// A topic to filter results by the topic name.
	Topic pulumi.StringPtrInput `pulumi:"topic"`
}

func (GetTopicsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicsArgs)(nil)).Elem()
}

// A collection of values returned by getTopics.
type GetTopicsResultOutput struct{ *pulumi.OutputState }

func (GetTopicsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTopicsResult)(nil)).Elem()
}

func (o GetTopicsResultOutput) ToGetTopicsResultOutput() GetTopicsResultOutput {
	return o
}

func (o GetTopicsResultOutput) ToGetTopicsResultOutputWithContext(ctx context.Context) GetTopicsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTopicsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTopicsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTopicsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The instanceId of the instance.
func (o GetTopicsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetTopicsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of topic names.
func (o GetTopicsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTopicsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetTopicsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTopicsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetTopicsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// The name of the topic.
func (o GetTopicsResultOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *string { return v.Topic }).(pulumi.StringPtrOutput)
}

// A list of topics. Each element contains the following attributes:
func (o GetTopicsResultOutput) Topics() GetTopicsTopicArrayOutput {
	return o.ApplyT(func(v GetTopicsResult) []GetTopicsTopic { return v.Topics }).(GetTopicsTopicArrayOutput)
}

func (o GetTopicsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetTopicsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopicsResultOutput{})
}
