// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of ONS Topics in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.53.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "onsInstanceName"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			topic := "onsTopicDatasourceName"
//			if param := cfg.Get("topic"); param != "" {
//				topic = param
//			}
//			_, err := rocketmq.NewInstance(ctx, "default", &rocketmq.InstanceArgs{
//				InstanceName: pulumi.String(name),
//				Remark:       pulumi.String("default_ons_instance_remark"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultTopic, err := rocketmq.NewTopic(ctx, "default", &rocketmq.TopicArgs{
//				TopicName:   pulumi.String(topic),
//				InstanceId:  _default.ID(),
//				MessageType: pulumi.Int(0),
//				Remark:      pulumi.String("dafault_ons_topic_remark"),
//			})
//			if err != nil {
//				return err
//			}
//			topicsDs := rocketmq.GetTopicsOutput(ctx, rocketmq.GetTopicsOutputArgs{
//				InstanceId: defaultTopic.InstanceId,
//				NameRegex:  pulumi.String(topic),
//				OutputFile: pulumi.String("topics.txt"),
//			}, nil)
//			ctx.Export("firstTopicName", topicsDs.ApplyT(func(topicsDs rocketmq.GetTopicsResult) (*string, error) {
//				return &topicsDs.Topics[0].TopicName, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
func GetTopics(ctx *pulumi.Context, args *GetTopicsArgs, opts ...pulumi.InvokeOption) (*GetTopicsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTopicsResult
	err := ctx.Invoke("alicloud:rocketmq/getTopics:getTopics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTopics.
type GetTopicsArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of topic IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the ONS Instance that owns the topics.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by the topic name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// A map of tags assigned to the Ons instance.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getTopics.
type GetTopicsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	InstanceId string   `pulumi:"instanceId"`
	NameRegex  *string  `pulumi:"nameRegex"`
	// A list of topic names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A map of tags assigned to the Ons instance.
	Tags map[string]string `pulumi:"tags"`
	// A list of topics. Each element contains the following attributes:
	Topics []GetTopicsTopic `pulumi:"topics"`
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
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of topic IDs to filter results.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the ONS Instance that owns the topics.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by the topic name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A map of tags assigned to the Ons instance.
	Tags pulumi.StringMapInput `pulumi:"tags"`
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

func (o GetTopicsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetTopicsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTopicsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTopicsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTopicsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTopicsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

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

// A map of tags assigned to the Ons instance.
func (o GetTopicsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetTopicsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A list of topics. Each element contains the following attributes:
func (o GetTopicsResultOutput) Topics() GetTopicsTopicArrayOutput {
	return o.ApplyT(func(v GetTopicsResult) []GetTopicsTopic { return v.Topics }).(GetTopicsTopicArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTopicsResultOutput{})
}
