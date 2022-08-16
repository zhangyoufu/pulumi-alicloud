// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of RAM Roles in an Alibaba Cloud account according to the specified filters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rolesDs, err := ram.GetRoles(ctx, &ram.GetRolesArgs{
//				NameRegex:  pulumi.StringRef(".*test.*"),
//				OutputFile: pulumi.StringRef("roles.txt"),
//				PolicyName: pulumi.StringRef("AliyunACSDefaultAccess"),
//				PolicyType: pulumi.StringRef("Custom"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRoleId", rolesDs.Roles[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetRoles(ctx *pulumi.Context, args *GetRolesArgs, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	var rv GetRolesResult
	err := ctx.Invoke("alicloud:ram/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type GetRolesArgs struct {
	// - A list of ram role IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by the role name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Filter results by a specific policy name. If you set this parameter without setting `policyType`, the later will be automatically set to `System`. The resulting roles will be attached to the specified policy.
	PolicyName *string `pulumi:"policyName"`
	// Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policyName` as well.
	PolicyType *string `pulumi:"policyType"`
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of ram role IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of ram role names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PolicyName *string  `pulumi:"policyName"`
	PolicyType *string  `pulumi:"policyType"`
	// A list of roles. Each element contains the following attributes:
	Roles []GetRolesRole `pulumi:"roles"`
}

func GetRolesOutput(ctx *pulumi.Context, args GetRolesOutputArgs, opts ...pulumi.InvokeOption) GetRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolesResult, error) {
			args := v.(GetRolesArgs)
			r, err := GetRoles(ctx, &args, opts...)
			var s GetRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type GetRolesOutputArgs struct {
	// - A list of ram role IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by the role name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Filter results by a specific policy name. If you set this parameter without setting `policyType`, the later will be automatically set to `System`. The resulting roles will be attached to the specified policy.
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
	// Filter results by a specific policy type. Valid values are `Custom` and `System`. If you set this parameter, you must set `policyName` as well.
	PolicyType pulumi.StringPtrInput `pulumi:"policyType"`
}

func (GetRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type GetRolesResultOutput struct{ *pulumi.OutputState }

func (GetRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesResult)(nil)).Elem()
}

func (o GetRolesResultOutput) ToGetRolesResultOutput() GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) ToGetRolesResultOutputWithContext(ctx context.Context) GetRolesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of ram role IDs.
func (o GetRolesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRolesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of ram role names.
func (o GetRolesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRolesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

// A list of roles. Each element contains the following attributes:
func (o GetRolesResultOutput) Roles() GetRolesRoleArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []GetRolesRole { return v.Roles }).(GetRolesRoleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesResultOutput{})
}
