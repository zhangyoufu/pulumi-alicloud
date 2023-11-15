// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Dts Migration Jobs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.157.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := dts.GetMigrationJobs(ctx, &dts.GetMigrationJobsArgs{
//				Ids: []string{
//					"dts_job_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("dtsMigrationJobId1", ids.Jobs[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetMigrationJobs(ctx *pulumi.Context, args *GetMigrationJobsArgs, opts ...pulumi.InvokeOption) (*GetMigrationJobsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMigrationJobsResult
	err := ctx.Invoke("alicloud:dts/getMigrationJobs:getMigrationJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMigrationJobs.
type GetMigrationJobsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Synchronization Job IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Migration Job name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getMigrationJobs.
type GetMigrationJobsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                `pulumi:"id"`
	Ids        []string              `pulumi:"ids"`
	Jobs       []GetMigrationJobsJob `pulumi:"jobs"`
	NameRegex  *string               `pulumi:"nameRegex"`
	Names      []string              `pulumi:"names"`
	OutputFile *string               `pulumi:"outputFile"`
}

func GetMigrationJobsOutput(ctx *pulumi.Context, args GetMigrationJobsOutputArgs, opts ...pulumi.InvokeOption) GetMigrationJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMigrationJobsResult, error) {
			args := v.(GetMigrationJobsArgs)
			r, err := GetMigrationJobs(ctx, &args, opts...)
			var s GetMigrationJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMigrationJobsResultOutput)
}

// A collection of arguments for invoking getMigrationJobs.
type GetMigrationJobsOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Synchronization Job IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Migration Job name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetMigrationJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMigrationJobsArgs)(nil)).Elem()
}

// A collection of values returned by getMigrationJobs.
type GetMigrationJobsResultOutput struct{ *pulumi.OutputState }

func (GetMigrationJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMigrationJobsResult)(nil)).Elem()
}

func (o GetMigrationJobsResultOutput) ToGetMigrationJobsResultOutput() GetMigrationJobsResultOutput {
	return o
}

func (o GetMigrationJobsResultOutput) ToGetMigrationJobsResultOutputWithContext(ctx context.Context) GetMigrationJobsResultOutput {
	return o
}

func (o GetMigrationJobsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMigrationJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMigrationJobsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetMigrationJobsResultOutput) Jobs() GetMigrationJobsJobArrayOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) []GetMigrationJobsJob { return v.Jobs }).(GetMigrationJobsJobArrayOutput)
}

func (o GetMigrationJobsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetMigrationJobsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetMigrationJobsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrationJobsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMigrationJobsResultOutput{})
}
