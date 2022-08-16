// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cen Traffic Marking Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.173.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cen.GetTrafficMarkingPolicies(ctx, &cen.GetTrafficMarkingPoliciesArgs{
//				TransitRouterId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTrafficMarkingPolicyId1", ids.Policies[0].Id)
//			nameRegex, err := cen.GetTrafficMarkingPolicies(ctx, &cen.GetTrafficMarkingPoliciesArgs{
//				TransitRouterId: "example_value",
//				NameRegex:       pulumi.StringRef("^my-TrafficMarkingPolicy"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTrafficMarkingPolicyId2", nameRegex.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetTrafficMarkingPolicies(ctx *pulumi.Context, args *GetTrafficMarkingPoliciesArgs, opts ...pulumi.InvokeOption) (*GetTrafficMarkingPoliciesResult, error) {
	var rv GetTrafficMarkingPoliciesResult
	err := ctx.Invoke("alicloud:cen/getTrafficMarkingPolicies:getTrafficMarkingPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrafficMarkingPolicies.
type GetTrafficMarkingPoliciesArgs struct {
	// The description of the Traffic Marking Policy.
	Description *string `pulumi:"description"`
	// A list of Traffic Marking Policy IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Traffic Marking Policy name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// A collection of values returned by getTrafficMarkingPolicies.
type GetTrafficMarkingPoliciesResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                            `pulumi:"id"`
	Ids             []string                          `pulumi:"ids"`
	NameRegex       *string                           `pulumi:"nameRegex"`
	Names           []string                          `pulumi:"names"`
	OutputFile      *string                           `pulumi:"outputFile"`
	Policies        []GetTrafficMarkingPoliciesPolicy `pulumi:"policies"`
	Status          *string                           `pulumi:"status"`
	TransitRouterId string                            `pulumi:"transitRouterId"`
}

func GetTrafficMarkingPoliciesOutput(ctx *pulumi.Context, args GetTrafficMarkingPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetTrafficMarkingPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTrafficMarkingPoliciesResult, error) {
			args := v.(GetTrafficMarkingPoliciesArgs)
			r, err := GetTrafficMarkingPolicies(ctx, &args, opts...)
			var s GetTrafficMarkingPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTrafficMarkingPoliciesResultOutput)
}

// A collection of arguments for invoking getTrafficMarkingPolicies.
type GetTrafficMarkingPoliciesOutputArgs struct {
	// The description of the Traffic Marking Policy.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A list of Traffic Marking Policy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Traffic Marking Policy name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput `pulumi:"transitRouterId"`
}

func (GetTrafficMarkingPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficMarkingPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getTrafficMarkingPolicies.
type GetTrafficMarkingPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetTrafficMarkingPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTrafficMarkingPoliciesResult)(nil)).Elem()
}

func (o GetTrafficMarkingPoliciesResultOutput) ToGetTrafficMarkingPoliciesResultOutput() GetTrafficMarkingPoliciesResultOutput {
	return o
}

func (o GetTrafficMarkingPoliciesResultOutput) ToGetTrafficMarkingPoliciesResultOutputWithContext(ctx context.Context) GetTrafficMarkingPoliciesResultOutput {
	return o
}

func (o GetTrafficMarkingPoliciesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTrafficMarkingPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) Policies() GetTrafficMarkingPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) []GetTrafficMarkingPoliciesPolicy { return v.Policies }).(GetTrafficMarkingPoliciesPolicyArrayOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetTrafficMarkingPoliciesResultOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTrafficMarkingPoliciesResult) string { return v.TransitRouterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTrafficMarkingPoliciesResultOutput{})
}
