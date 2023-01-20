// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Resource Manager Policy Versions of the current Alibaba Cloud user.
//
// > **NOTE:**  Available in 1.85.0+.
//
// ## Example Usage
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
//			_default, err := resourcemanager.GetPolicyVersions(ctx, &resourcemanager.GetPolicyVersionsArgs{
//				PolicyName: "tftest",
//				PolicyType: "Custom",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstPolicyVersionId", _default.Versions[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetPolicyVersions(ctx *pulumi.Context, args *GetPolicyVersionsArgs, opts ...pulumi.InvokeOption) (*GetPolicyVersionsResult, error) {
	var rv GetPolicyVersionsResult
	err := ctx.Invoke("alicloud:resourcemanager/getPolicyVersions:getPolicyVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyVersions.
type GetPolicyVersionsArgs struct {
	// Default to `false`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of policy version IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The name of the policy.
	PolicyName string `pulumi:"policyName"`
	// The type of the policy. Valid values:`Custom` and `System`.
	PolicyType string `pulumi:"policyType"`
}

// A collection of values returned by getPolicyVersions.
type GetPolicyVersionsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of policy version IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	PolicyName string   `pulumi:"policyName"`
	PolicyType string   `pulumi:"policyType"`
	// A list of policy versions. Each element contains the following attributes:
	Versions []GetPolicyVersionsVersion `pulumi:"versions"`
}

func GetPolicyVersionsOutput(ctx *pulumi.Context, args GetPolicyVersionsOutputArgs, opts ...pulumi.InvokeOption) GetPolicyVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyVersionsResult, error) {
			args := v.(GetPolicyVersionsArgs)
			r, err := GetPolicyVersions(ctx, &args, opts...)
			var s GetPolicyVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyVersionsResultOutput)
}

// A collection of arguments for invoking getPolicyVersions.
type GetPolicyVersionsOutputArgs struct {
	// Default to `false`. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of policy version IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The name of the policy.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The type of the policy. Valid values:`Custom` and `System`.
	PolicyType pulumi.StringInput `pulumi:"policyType"`
}

func (GetPolicyVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyVersions.
type GetPolicyVersionsResultOutput struct{ *pulumi.OutputState }

func (GetPolicyVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyVersionsResult)(nil)).Elem()
}

func (o GetPolicyVersionsResultOutput) ToGetPolicyVersionsResultOutput() GetPolicyVersionsResultOutput {
	return o
}

func (o GetPolicyVersionsResultOutput) ToGetPolicyVersionsResultOutputWithContext(ctx context.Context) GetPolicyVersionsResultOutput {
	return o
}

func (o GetPolicyVersionsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of policy version IDs.
func (o GetPolicyVersionsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPolicyVersionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetPolicyVersionsResultOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) string { return v.PolicyName }).(pulumi.StringOutput)
}

func (o GetPolicyVersionsResultOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) string { return v.PolicyType }).(pulumi.StringOutput)
}

// A list of policy versions. Each element contains the following attributes:
func (o GetPolicyVersionsResultOutput) Versions() GetPolicyVersionsVersionArrayOutput {
	return o.ApplyT(func(v GetPolicyVersionsResult) []GetPolicyVersionsVersion { return v.Versions }).(GetPolicyVersionsVersionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyVersionsResultOutput{})
}
