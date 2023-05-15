// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Resource Manager Control Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.120.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := resourcemanager.GetControlPolicies(ctx, &resourcemanager.GetControlPoliciesArgs{
//				Ids: []string{
//					"example_value",
//				},
//				NameRegex: pulumi.StringRef("the_resource_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstResourceManagerControlPolicyId", example.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetControlPolicies(ctx *pulumi.Context, args *GetControlPoliciesArgs, opts ...pulumi.InvokeOption) (*GetControlPoliciesResult, error) {
	var rv GetControlPoliciesResult
	err := ctx.Invoke("alicloud:resourcemanager/getControlPolicies:getControlPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Control Policy IDs.
	Ids []string `pulumi:"ids"`
	// The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`.
	Language *string `pulumi:"language"`
	// A regex string to filter results by Control Policy name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The type of policy.
	PolicyType *string `pulumi:"policyType"`
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                     `pulumi:"id"`
	Ids        []string                   `pulumi:"ids"`
	Language   *string                    `pulumi:"language"`
	NameRegex  *string                    `pulumi:"nameRegex"`
	Names      []string                   `pulumi:"names"`
	OutputFile *string                    `pulumi:"outputFile"`
	Policies   []GetControlPoliciesPolicy `pulumi:"policies"`
	PolicyType *string                    `pulumi:"policyType"`
}

func GetControlPoliciesOutput(ctx *pulumi.Context, args GetControlPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetControlPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetControlPoliciesResult, error) {
			args := v.(GetControlPoliciesArgs)
			r, err := GetControlPolicies(ctx, &args, opts...)
			var s GetControlPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetControlPoliciesResultOutput)
}

// A collection of arguments for invoking getControlPolicies.
type GetControlPoliciesOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Control Policy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`.
	Language pulumi.StringPtrInput `pulumi:"language"`
	// A regex string to filter results by Control Policy name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The type of policy.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
}

func (GetControlPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getControlPolicies.
type GetControlPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetControlPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetControlPoliciesResult)(nil)).Elem()
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutput() GetControlPoliciesResultOutput {
	return o
}

func (o GetControlPoliciesResultOutput) ToGetControlPoliciesResultOutputWithContext(ctx context.Context) GetControlPoliciesResultOutput {
	return o
}

func (o GetControlPoliciesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetControlPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetControlPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetControlPoliciesResultOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetControlPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetControlPoliciesResultOutput) Policies() GetControlPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) []GetControlPoliciesPolicy { return v.Policies }).(GetControlPoliciesPolicyArrayOutput)
}

func (o GetControlPoliciesResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetControlPoliciesResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetControlPoliciesResultOutput{})
}
