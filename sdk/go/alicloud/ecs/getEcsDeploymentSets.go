// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Deployment Sets of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.140.0+.
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
//			ids, err := ecs.GetEcsDeploymentSets(ctx, &ecs.GetEcsDeploymentSetsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDeploymentSetId1", ids.Sets[0].Id)
//			nameRegex, err := ecs.GetEcsDeploymentSets(ctx, &ecs.GetEcsDeploymentSetsArgs{
//				NameRegex: pulumi.StringRef("^my-DeploymentSet"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDeploymentSetId2", nameRegex.Sets[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEcsDeploymentSets(ctx *pulumi.Context, args *GetEcsDeploymentSetsArgs, opts ...pulumi.InvokeOption) (*GetEcsDeploymentSetsResult, error) {
	var rv GetEcsDeploymentSetsResult
	err := ctx.Invoke("alicloud:ecs/getEcsDeploymentSets:getEcsDeploymentSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsDeploymentSets.
type GetEcsDeploymentSetsArgs struct {
	// The name of the deployment set.
	DeploymentSetName *string `pulumi:"deploymentSetName"`
	// A list of Deployment Set IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Deployment Set name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The deployment strategy.
	Strategy *string `pulumi:"strategy"`
}

// A collection of values returned by getEcsDeploymentSets.
type GetEcsDeploymentSetsResult struct {
	DeploymentSetName *string `pulumi:"deploymentSetName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                    `pulumi:"id"`
	Ids        []string                  `pulumi:"ids"`
	NameRegex  *string                   `pulumi:"nameRegex"`
	Names      []string                  `pulumi:"names"`
	OutputFile *string                   `pulumi:"outputFile"`
	Sets       []GetEcsDeploymentSetsSet `pulumi:"sets"`
	Strategy   *string                   `pulumi:"strategy"`
}

func GetEcsDeploymentSetsOutput(ctx *pulumi.Context, args GetEcsDeploymentSetsOutputArgs, opts ...pulumi.InvokeOption) GetEcsDeploymentSetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsDeploymentSetsResult, error) {
			args := v.(GetEcsDeploymentSetsArgs)
			r, err := GetEcsDeploymentSets(ctx, &args, opts...)
			var s GetEcsDeploymentSetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsDeploymentSetsResultOutput)
}

// A collection of arguments for invoking getEcsDeploymentSets.
type GetEcsDeploymentSetsOutputArgs struct {
	// The name of the deployment set.
	DeploymentSetName pulumi.StringPtrInput `pulumi:"deploymentSetName"`
	// A list of Deployment Set IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Deployment Set name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The deployment strategy.
	Strategy pulumi.StringPtrInput `pulumi:"strategy"`
}

func (GetEcsDeploymentSetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDeploymentSetsArgs)(nil)).Elem()
}

// A collection of values returned by getEcsDeploymentSets.
type GetEcsDeploymentSetsResultOutput struct{ *pulumi.OutputState }

func (GetEcsDeploymentSetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDeploymentSetsResult)(nil)).Elem()
}

func (o GetEcsDeploymentSetsResultOutput) ToGetEcsDeploymentSetsResultOutput() GetEcsDeploymentSetsResultOutput {
	return o
}

func (o GetEcsDeploymentSetsResultOutput) ToGetEcsDeploymentSetsResultOutputWithContext(ctx context.Context) GetEcsDeploymentSetsResultOutput {
	return o
}

func (o GetEcsDeploymentSetsResultOutput) DeploymentSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) *string { return v.DeploymentSetName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsDeploymentSetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsDeploymentSetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsDeploymentSetsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsDeploymentSetsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsDeploymentSetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsDeploymentSetsResultOutput) Sets() GetEcsDeploymentSetsSetArrayOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) []GetEcsDeploymentSetsSet { return v.Sets }).(GetEcsDeploymentSetsSetArrayOutput)
}

func (o GetEcsDeploymentSetsResultOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDeploymentSetsResult) *string { return v.Strategy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsDeploymentSetsResultOutput{})
}
