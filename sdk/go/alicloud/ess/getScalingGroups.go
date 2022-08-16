// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides available scaling group resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			scalinggroupsDs, err := ess.GetScalingGroups(ctx, &ess.GetScalingGroupsArgs{
//				Ids: []string{
//					"scaling_group_id1",
//					"scaling_group_id2",
//				},
//				NameRegex: pulumi.StringRef("scaling_group_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstScalingGroup", scalinggroupsDs.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetScalingGroups(ctx *pulumi.Context, args *GetScalingGroupsArgs, opts ...pulumi.InvokeOption) (*GetScalingGroupsResult, error) {
	var rv GetScalingGroupsResult
	err := ctx.Invoke("alicloud:ess/getScalingGroups:getScalingGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingGroups.
type GetScalingGroupsArgs struct {
	// A list of scaling group IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting scaling groups by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getScalingGroups.
type GetScalingGroupsResult struct {
	// A list of scaling groups. Each element contains the following attributes:
	Groups []GetScalingGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of scaling group ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of scaling group names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetScalingGroupsOutput(ctx *pulumi.Context, args GetScalingGroupsOutputArgs, opts ...pulumi.InvokeOption) GetScalingGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScalingGroupsResult, error) {
			args := v.(GetScalingGroupsArgs)
			r, err := GetScalingGroups(ctx, &args, opts...)
			var s GetScalingGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScalingGroupsResultOutput)
}

// A collection of arguments for invoking getScalingGroups.
type GetScalingGroupsOutputArgs struct {
	// A list of scaling group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter resulting scaling groups by name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetScalingGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getScalingGroups.
type GetScalingGroupsResultOutput struct{ *pulumi.OutputState }

func (GetScalingGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingGroupsResult)(nil)).Elem()
}

func (o GetScalingGroupsResultOutput) ToGetScalingGroupsResultOutput() GetScalingGroupsResultOutput {
	return o
}

func (o GetScalingGroupsResultOutput) ToGetScalingGroupsResultOutputWithContext(ctx context.Context) GetScalingGroupsResultOutput {
	return o
}

// A list of scaling groups. Each element contains the following attributes:
func (o GetScalingGroupsResultOutput) Groups() GetScalingGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) []GetScalingGroupsGroup { return v.Groups }).(GetScalingGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetScalingGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of scaling group ids.
func (o GetScalingGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetScalingGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of scaling group names.
func (o GetScalingGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetScalingGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScalingGroupsResultOutput{})
}
