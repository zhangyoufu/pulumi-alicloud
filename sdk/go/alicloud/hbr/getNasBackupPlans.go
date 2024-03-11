// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Hbr NasBackupPlans of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := hbr.GetNasBackupPlans(ctx, &hbr.GetNasBackupPlansArgs{
//				NameRegex: pulumi.StringRef("^my-NasBackupPlan"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("hbrNasBackupPlanId", ids.Plans[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetNasBackupPlans(ctx *pulumi.Context, args *GetNasBackupPlansArgs, opts ...pulumi.InvokeOption) (*GetNasBackupPlansResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNasBackupPlansResult
	err := ctx.Invoke("alicloud:hbr/getNasBackupPlans:getNasBackupPlans", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNasBackupPlans.
type GetNasBackupPlansArgs struct {
	// The File System ID of Nas.
	FileSystemId *string `pulumi:"fileSystemId"`
	// A list of NasBackupPlan IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by NasBackupPlan name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of backup vault.
	VaultId *string `pulumi:"vaultId"`
}

// A collection of values returned by getNasBackupPlans.
type GetNasBackupPlansResult struct {
	FileSystemId *string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                  `pulumi:"id"`
	Ids        []string                `pulumi:"ids"`
	NameRegex  *string                 `pulumi:"nameRegex"`
	Names      []string                `pulumi:"names"`
	OutputFile *string                 `pulumi:"outputFile"`
	Plans      []GetNasBackupPlansPlan `pulumi:"plans"`
	VaultId    *string                 `pulumi:"vaultId"`
}

func GetNasBackupPlansOutput(ctx *pulumi.Context, args GetNasBackupPlansOutputArgs, opts ...pulumi.InvokeOption) GetNasBackupPlansResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNasBackupPlansResult, error) {
			args := v.(GetNasBackupPlansArgs)
			r, err := GetNasBackupPlans(ctx, &args, opts...)
			var s GetNasBackupPlansResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNasBackupPlansResultOutput)
}

// A collection of arguments for invoking getNasBackupPlans.
type GetNasBackupPlansOutputArgs struct {
	// The File System ID of Nas.
	FileSystemId pulumi.StringPtrInput `pulumi:"fileSystemId"`
	// A list of NasBackupPlan IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by NasBackupPlan name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of backup vault.
	VaultId pulumi.StringPtrInput `pulumi:"vaultId"`
}

func (GetNasBackupPlansOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasBackupPlansArgs)(nil)).Elem()
}

// A collection of values returned by getNasBackupPlans.
type GetNasBackupPlansResultOutput struct{ *pulumi.OutputState }

func (GetNasBackupPlansResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasBackupPlansResult)(nil)).Elem()
}

func (o GetNasBackupPlansResultOutput) ToGetNasBackupPlansResultOutput() GetNasBackupPlansResultOutput {
	return o
}

func (o GetNasBackupPlansResultOutput) ToGetNasBackupPlansResultOutputWithContext(ctx context.Context) GetNasBackupPlansResultOutput {
	return o
}

func (o GetNasBackupPlansResultOutput) FileSystemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) *string { return v.FileSystemId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNasBackupPlansResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNasBackupPlansResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetNasBackupPlansResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetNasBackupPlansResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetNasBackupPlansResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetNasBackupPlansResultOutput) Plans() GetNasBackupPlansPlanArrayOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) []GetNasBackupPlansPlan { return v.Plans }).(GetNasBackupPlansPlanArrayOutput)
}

func (o GetNasBackupPlansResultOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasBackupPlansResult) *string { return v.VaultId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNasBackupPlansResultOutput{})
}
