// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Threat Detection Backup Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := threatdetection.GetBackupPolicies(ctx, &threatdetection.GetBackupPoliciesArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("threatDetectionBackupPoliciesId1", ids.Policies[0].Id)
//			nameRegex, err := threatdetection.GetBackupPolicies(ctx, &threatdetection.GetBackupPoliciesArgs{
//				NameRegex: pulumi.StringRef("tf-example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("threatDetectionBackupPoliciesId2", nameRegex.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetBackupPolicies(ctx *pulumi.Context, args *GetBackupPoliciesArgs, opts ...pulumi.InvokeOption) (*GetBackupPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupPoliciesResult
	err := ctx.Invoke("alicloud:threatdetection/getBackupPolicies:getBackupPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupPolicies.
type GetBackupPoliciesArgs struct {
	CurrentPage *int `pulumi:"currentPage"`
	// A list of Threat Detection Backup Policies IDs.
	Ids []string `pulumi:"ids"`
	// The information that you want to use to identify the servers protected by the anti-ransomware policy. You can enter the IP address or ID of a server.
	MachineRemark *string `pulumi:"machineRemark"`
	// The name of the anti-ransomware policy that you want to query.
	Name *string `pulumi:"name"`
	// A regex string to filter results by Threat Detection Backup Policies name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageSize   *int    `pulumi:"pageSize"`
	// The status of the anti-ransomware policy. Valid Value: `enabled`, `disabled`, `closed`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getBackupPolicies.
type GetBackupPoliciesResult struct {
	CurrentPage *int `pulumi:"currentPage"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	MachineRemark *string  `pulumi:"machineRemark"`
	Name          *string  `pulumi:"name"`
	NameRegex     *string  `pulumi:"nameRegex"`
	// A list of Threat Detection Backup Policy names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	PageSize   *int     `pulumi:"pageSize"`
	// A list of Threat Detection Backup policies. Each element contains the following attributes:
	Policies []GetBackupPoliciesPolicy `pulumi:"policies"`
	// The status of the anti-ransomware policy.
	Status *string `pulumi:"status"`
}

func GetBackupPoliciesOutput(ctx *pulumi.Context, args GetBackupPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetBackupPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupPoliciesResult, error) {
			args := v.(GetBackupPoliciesArgs)
			r, err := GetBackupPolicies(ctx, &args, opts...)
			var s GetBackupPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupPoliciesResultOutput)
}

// A collection of arguments for invoking getBackupPolicies.
type GetBackupPoliciesOutputArgs struct {
	CurrentPage pulumi.IntPtrInput `pulumi:"currentPage"`
	// A list of Threat Detection Backup Policies IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The information that you want to use to identify the servers protected by the anti-ransomware policy. You can enter the IP address or ID of a server.
	MachineRemark pulumi.StringPtrInput `pulumi:"machineRemark"`
	// The name of the anti-ransomware policy that you want to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A regex string to filter results by Threat Detection Backup Policies name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The status of the anti-ransomware policy. Valid Value: `enabled`, `disabled`, `closed`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetBackupPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getBackupPolicies.
type GetBackupPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetBackupPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupPoliciesResult)(nil)).Elem()
}

func (o GetBackupPoliciesResultOutput) ToGetBackupPoliciesResultOutput() GetBackupPoliciesResultOutput {
	return o
}

func (o GetBackupPoliciesResultOutput) ToGetBackupPoliciesResultOutputWithContext(ctx context.Context) GetBackupPoliciesResultOutput {
	return o
}

func (o GetBackupPoliciesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBackupPoliciesResult] {
	return pulumix.Output[GetBackupPoliciesResult]{
		OutputState: o.OutputState,
	}
}

func (o GetBackupPoliciesResultOutput) CurrentPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *int { return v.CurrentPage }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBackupPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetBackupPoliciesResultOutput) MachineRemark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *string { return v.MachineRemark }).(pulumi.StringPtrOutput)
}

func (o GetBackupPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetBackupPoliciesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Threat Detection Backup Policy names.
func (o GetBackupPoliciesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetBackupPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBackupPoliciesResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// A list of Threat Detection Backup policies. Each element contains the following attributes:
func (o GetBackupPoliciesResultOutput) Policies() GetBackupPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) []GetBackupPoliciesPolicy { return v.Policies }).(GetBackupPoliciesPolicyArrayOutput)
}

// The status of the anti-ransomware policy.
func (o GetBackupPoliciesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupPoliciesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupPoliciesResultOutput{})
}
