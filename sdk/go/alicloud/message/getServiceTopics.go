// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package message

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Message Notification Service Topics of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.188.0+.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/message"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := message.GetServiceTopics(ctx, &message.GetServiceTopicsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("topicId1", ids.Topics[0].Id)
//			name, err := message.GetServiceTopics(ctx, &message.GetServiceTopicsArgs{
//				TopicName: pulumi.StringRef("tf-example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("topicId2", name.Topics[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetServiceTopics(ctx *pulumi.Context, args *GetServiceTopicsArgs, opts ...pulumi.InvokeOption) (*GetServiceTopicsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceTopicsResult
	err := ctx.Invoke("alicloud:message/getServiceTopics:getServiceTopics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceTopics.
type GetServiceTopicsArgs struct {
	// A list of Topic IDs. Its element value is same as Topic Name.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Topic name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The name of the topic.
	TopicName *string `pulumi:"topicName"`
}

// A collection of values returned by getServiceTopics.
type GetServiceTopicsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Topic names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
	// The name of the topic.
	TopicName *string `pulumi:"topicName"`
	// A list of Topics. Each element contains the following attributes:
	Topics []GetServiceTopicsTopic `pulumi:"topics"`
}

func GetServiceTopicsOutput(ctx *pulumi.Context, args GetServiceTopicsOutputArgs, opts ...pulumi.InvokeOption) GetServiceTopicsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceTopicsResult, error) {
			args := v.(GetServiceTopicsArgs)
			r, err := GetServiceTopics(ctx, &args, opts...)
			var s GetServiceTopicsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceTopicsResultOutput)
}

// A collection of arguments for invoking getServiceTopics.
type GetServiceTopicsOutputArgs struct {
	// A list of Topic IDs. Its element value is same as Topic Name.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Topic name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The name of the topic.
	TopicName pulumi.StringPtrInput `pulumi:"topicName"`
}

func (GetServiceTopicsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceTopicsArgs)(nil)).Elem()
}

// A collection of values returned by getServiceTopics.
type GetServiceTopicsResultOutput struct{ *pulumi.OutputState }

func (GetServiceTopicsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceTopicsResult)(nil)).Elem()
}

func (o GetServiceTopicsResultOutput) ToGetServiceTopicsResultOutput() GetServiceTopicsResultOutput {
	return o
}

func (o GetServiceTopicsResultOutput) ToGetServiceTopicsResultOutputWithContext(ctx context.Context) GetServiceTopicsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceTopicsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceTopicsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServiceTopicsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Topic names.
func (o GetServiceTopicsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServiceTopicsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetServiceTopicsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetServiceTopicsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// The name of the topic.
func (o GetServiceTopicsResultOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

// A list of Topics. Each element contains the following attributes:
func (o GetServiceTopicsResultOutput) Topics() GetServiceTopicsTopicArrayOutput {
	return o.ApplyT(func(v GetServiceTopicsResult) []GetServiceTopicsTopic { return v.Topics }).(GetServiceTopicsTopicArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceTopicsResultOutput{})
}
