// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yundun

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBastionHostInstances(ctx *pulumi.Context, args *GetBastionHostInstancesArgs, opts ...pulumi.InvokeOption) (*GetBastionHostInstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBastionHostInstancesResult
	err := ctx.Invoke("alicloud:yundun/getBastionHostInstances:getBastionHostInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBastionHostInstances.
type GetBastionHostInstancesArgs struct {
	DescriptionRegex *string                `pulumi:"descriptionRegex"`
	Ids              []string               `pulumi:"ids"`
	OutputFile       *string                `pulumi:"outputFile"`
	Tags             map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getBastionHostInstances.
type GetBastionHostInstancesResult struct {
	DescriptionRegex *string  `pulumi:"descriptionRegex"`
	Descriptions     []string `pulumi:"descriptions"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                            `pulumi:"id"`
	Ids        []string                          `pulumi:"ids"`
	Instances  []GetBastionHostInstancesInstance `pulumi:"instances"`
	OutputFile *string                           `pulumi:"outputFile"`
	Tags       map[string]interface{}            `pulumi:"tags"`
}

func GetBastionHostInstancesOutput(ctx *pulumi.Context, args GetBastionHostInstancesOutputArgs, opts ...pulumi.InvokeOption) GetBastionHostInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBastionHostInstancesResult, error) {
			args := v.(GetBastionHostInstancesArgs)
			r, err := GetBastionHostInstances(ctx, &args, opts...)
			var s GetBastionHostInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBastionHostInstancesResultOutput)
}

// A collection of arguments for invoking getBastionHostInstances.
type GetBastionHostInstancesOutputArgs struct {
	DescriptionRegex pulumi.StringPtrInput   `pulumi:"descriptionRegex"`
	Ids              pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile       pulumi.StringPtrInput   `pulumi:"outputFile"`
	Tags             pulumi.MapInput         `pulumi:"tags"`
}

func (GetBastionHostInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionHostInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getBastionHostInstances.
type GetBastionHostInstancesResultOutput struct{ *pulumi.OutputState }

func (GetBastionHostInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionHostInstancesResult)(nil)).Elem()
}

func (o GetBastionHostInstancesResultOutput) ToGetBastionHostInstancesResultOutput() GetBastionHostInstancesResultOutput {
	return o
}

func (o GetBastionHostInstancesResultOutput) ToGetBastionHostInstancesResultOutputWithContext(ctx context.Context) GetBastionHostInstancesResultOutput {
	return o
}

func (o GetBastionHostInstancesResultOutput) DescriptionRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) *string { return v.DescriptionRegex }).(pulumi.StringPtrOutput)
}

func (o GetBastionHostInstancesResultOutput) Descriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) []string { return v.Descriptions }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBastionHostInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBastionHostInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetBastionHostInstancesResultOutput) Instances() GetBastionHostInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) []GetBastionHostInstancesInstance { return v.Instances }).(GetBastionHostInstancesInstanceArrayOutput)
}

func (o GetBastionHostInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBastionHostInstancesResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetBastionHostInstancesResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBastionHostInstancesResultOutput{})
}
