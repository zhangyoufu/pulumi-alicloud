// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package brain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Brain Industrial Pid Loops of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.117.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/brain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := brain.GetIndustrialPidLoops(ctx, &brain.GetIndustrialPidLoopsArgs{
//				PidProjectId: "856c6b8f-ca63-40a4-xxxx-xxxx",
//				Ids: []string{
//					"742a3d4e-d8b0-47c8-xxxx-xxxx",
//				},
//				NameRegex: pulumi.StringRef("tf-testACC"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstBrainIndustrialPidLoopId", example.Loops[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetIndustrialPidLoops(ctx *pulumi.Context, args *GetIndustrialPidLoopsArgs, opts ...pulumi.InvokeOption) (*GetIndustrialPidLoopsResult, error) {
	var rv GetIndustrialPidLoopsResult
	err := ctx.Invoke("alicloud:brain/getIndustrialPidLoops:getIndustrialPidLoops", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIndustrialPidLoops.
type GetIndustrialPidLoopsArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Pid Loop IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Pid Loop name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of Pid Loop.
	PidLoopName *string `pulumi:"pidLoopName"`
	// The pid project id.
	PidProjectId string `pulumi:"pidProjectId"`
	// The status of Pid Loop.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getIndustrialPidLoops.
type GetIndustrialPidLoopsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id           string                      `pulumi:"id"`
	Ids          []string                    `pulumi:"ids"`
	Loops        []GetIndustrialPidLoopsLoop `pulumi:"loops"`
	NameRegex    *string                     `pulumi:"nameRegex"`
	Names        []string                    `pulumi:"names"`
	OutputFile   *string                     `pulumi:"outputFile"`
	PidLoopName  *string                     `pulumi:"pidLoopName"`
	PidProjectId string                      `pulumi:"pidProjectId"`
	Status       *string                     `pulumi:"status"`
}

func GetIndustrialPidLoopsOutput(ctx *pulumi.Context, args GetIndustrialPidLoopsOutputArgs, opts ...pulumi.InvokeOption) GetIndustrialPidLoopsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIndustrialPidLoopsResult, error) {
			args := v.(GetIndustrialPidLoopsArgs)
			r, err := GetIndustrialPidLoops(ctx, &args, opts...)
			var s GetIndustrialPidLoopsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIndustrialPidLoopsResultOutput)
}

// A collection of arguments for invoking getIndustrialPidLoops.
type GetIndustrialPidLoopsOutputArgs struct {
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Pid Loop IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Pid Loop name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of Pid Loop.
	PidLoopName pulumi.StringPtrInput `pulumi:"pidLoopName"`
	// The pid project id.
	PidProjectId pulumi.StringInput `pulumi:"pidProjectId"`
	// The status of Pid Loop.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetIndustrialPidLoopsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIndustrialPidLoopsArgs)(nil)).Elem()
}

// A collection of values returned by getIndustrialPidLoops.
type GetIndustrialPidLoopsResultOutput struct{ *pulumi.OutputState }

func (GetIndustrialPidLoopsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIndustrialPidLoopsResult)(nil)).Elem()
}

func (o GetIndustrialPidLoopsResultOutput) ToGetIndustrialPidLoopsResultOutput() GetIndustrialPidLoopsResultOutput {
	return o
}

func (o GetIndustrialPidLoopsResultOutput) ToGetIndustrialPidLoopsResultOutputWithContext(ctx context.Context) GetIndustrialPidLoopsResultOutput {
	return o
}

func (o GetIndustrialPidLoopsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIndustrialPidLoopsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIndustrialPidLoopsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetIndustrialPidLoopsResultOutput) Loops() GetIndustrialPidLoopsLoopArrayOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) []GetIndustrialPidLoopsLoop { return v.Loops }).(GetIndustrialPidLoopsLoopArrayOutput)
}

func (o GetIndustrialPidLoopsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetIndustrialPidLoopsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetIndustrialPidLoopsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetIndustrialPidLoopsResultOutput) PidLoopName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) *string { return v.PidLoopName }).(pulumi.StringPtrOutput)
}

func (o GetIndustrialPidLoopsResultOutput) PidProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) string { return v.PidProjectId }).(pulumi.StringOutput)
}

func (o GetIndustrialPidLoopsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIndustrialPidLoopsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIndustrialPidLoopsResultOutput{})
}
