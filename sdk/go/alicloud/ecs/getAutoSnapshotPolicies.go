// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Auto Snapshot Policies of the current Alibaba Cloud user.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ecs.GetAutoSnapshotPolicies(ctx, &ecs.GetAutoSnapshotPoliciesArgs{
//				Ids: []string{
//					"sp-bp14e66xxxxxxxx",
//				},
//				NameRegex: pulumi.StringRef("tf-testAcc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcsAutoSnapshotPolicyId", example.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAutoSnapshotPolicies(ctx *pulumi.Context, args *GetAutoSnapshotPoliciesArgs, opts ...pulumi.InvokeOption) (*GetAutoSnapshotPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAutoSnapshotPoliciesResult
	err := ctx.Invoke("alicloud:ecs/getAutoSnapshotPolicies:getAutoSnapshotPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoSnapshotPolicies.
type GetAutoSnapshotPoliciesArgs struct {
	// A list of Auto Snapshot Policy IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Auto Snapshot Policy name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of Auto Snapshot Policy. Valid Values: `Expire`, `Normal`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getAutoSnapshotPolicies.
type GetAutoSnapshotPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                          `pulumi:"id"`
	Ids        []string                        `pulumi:"ids"`
	NameRegex  *string                         `pulumi:"nameRegex"`
	Names      []string                        `pulumi:"names"`
	OutputFile *string                         `pulumi:"outputFile"`
	Policies   []GetAutoSnapshotPoliciesPolicy `pulumi:"policies"`
	Status     *string                         `pulumi:"status"`
	Tags       map[string]string               `pulumi:"tags"`
}

func GetAutoSnapshotPoliciesOutput(ctx *pulumi.Context, args GetAutoSnapshotPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetAutoSnapshotPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAutoSnapshotPoliciesResult, error) {
			args := v.(GetAutoSnapshotPoliciesArgs)
			r, err := GetAutoSnapshotPolicies(ctx, &args, opts...)
			var s GetAutoSnapshotPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAutoSnapshotPoliciesResultOutput)
}

// A collection of arguments for invoking getAutoSnapshotPolicies.
type GetAutoSnapshotPoliciesOutputArgs struct {
	// A list of Auto Snapshot Policy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Auto Snapshot Policy name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of Auto Snapshot Policy. Valid Values: `Expire`, `Normal`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetAutoSnapshotPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutoSnapshotPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getAutoSnapshotPolicies.
type GetAutoSnapshotPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetAutoSnapshotPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutoSnapshotPoliciesResult)(nil)).Elem()
}

func (o GetAutoSnapshotPoliciesResultOutput) ToGetAutoSnapshotPoliciesResultOutput() GetAutoSnapshotPoliciesResultOutput {
	return o
}

func (o GetAutoSnapshotPoliciesResultOutput) ToGetAutoSnapshotPoliciesResultOutputWithContext(ctx context.Context) GetAutoSnapshotPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAutoSnapshotPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) Policies() GetAutoSnapshotPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) []GetAutoSnapshotPoliciesPolicy { return v.Policies }).(GetAutoSnapshotPoliciesPolicyArrayOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetAutoSnapshotPoliciesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAutoSnapshotPoliciesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAutoSnapshotPoliciesResultOutput{})
}
