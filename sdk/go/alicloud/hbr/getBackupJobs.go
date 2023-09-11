// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Hbr Backup Jobs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.138.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// defaultEcsBackupPlans, err := hbr.GetEcsBackupPlans(ctx, &hbr.GetEcsBackupPlansArgs{
// NameRegex: pulumi.StringRef("plan-name"),
// }, nil);
// if err != nil {
// return err
// }
// defaultBackupJobs, err := hbr.GetBackupJobs(ctx, &hbr.GetBackupJobsArgs{
// SourceType: "ECS_FILE",
// Filters: []hbr.GetBackupJobsFilter{
// {
// Key: pulumi.StringRef("VaultId"),
// Operator: pulumi.StringRef("IN"),
// Values: interface{}{
// defaultEcsBackupPlans.Plans[0].VaultId,
// },
// },
// {
// Key: pulumi.StringRef("InstanceId"),
// Operator: pulumi.StringRef("IN"),
// Values: interface{}{
// defaultEcsBackupPlans.Plans[0].InstanceId,
// },
// },
// {
// Key: pulumi.StringRef("CompleteTime"),
// Operator: pulumi.StringRef("BETWEEN"),
// Values: []string{
// "2021-08-23T14:17:15CST",
// "2021-08-24T14:17:15CST",
// },
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// example, err := hbr.GetBackupJobs(ctx, &hbr.GetBackupJobsArgs{
// SourceType: "ECS_FILE",
// Status: pulumi.StringRef("COMPLETE"),
// Filters: []hbr.GetBackupJobsFilter{
// {
// Key: pulumi.StringRef("VaultId"),
// Operator: pulumi.StringRef("IN"),
// Values: interface{}{
// defaultEcsBackupPlans.Plans[0].VaultId,
// },
// },
// {
// Key: pulumi.StringRef("InstanceId"),
// Operator: pulumi.StringRef("IN"),
// Values: interface{}{
// defaultEcsBackupPlans.Plans[0].InstanceId,
// },
// },
// {
// Key: pulumi.StringRef("CompleteTime"),
// Operator: pulumi.StringRef("LESS_THAN"),
// Values: []string{
// "2021-10-20T20:20:20CST",
// },
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("alicloudHbrBackupJobsDefault1", defaultBackupJobs.Jobs[0].Id)
// ctx.Export("alicloudHbrBackupJobsExample1", example.Jobs[0].Id)
// return nil
// })
// }
// ```
func GetBackupJobs(ctx *pulumi.Context, args *GetBackupJobsArgs, opts ...pulumi.InvokeOption) (*GetBackupJobsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupJobsResult
	err := ctx.Invoke("alicloud:hbr/getBackupJobs:getBackupJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupJobs.
type GetBackupJobsArgs struct {
	Filters []GetBackupJobsFilter `pulumi:"filters"`
	// A list of Backup Job IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
	SortDirection *string `pulumi:"sortDirection"`
	// The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
	SourceType string `pulumi:"sourceType"`
	// The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getBackupJobs.
type GetBackupJobsResult struct {
	Filters []GetBackupJobsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id            string             `pulumi:"id"`
	Ids           []string           `pulumi:"ids"`
	Jobs          []GetBackupJobsJob `pulumi:"jobs"`
	OutputFile    *string            `pulumi:"outputFile"`
	SortDirection *string            `pulumi:"sortDirection"`
	SourceType    string             `pulumi:"sourceType"`
	Status        *string            `pulumi:"status"`
}

func GetBackupJobsOutput(ctx *pulumi.Context, args GetBackupJobsOutputArgs, opts ...pulumi.InvokeOption) GetBackupJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupJobsResult, error) {
			args := v.(GetBackupJobsArgs)
			r, err := GetBackupJobs(ctx, &args, opts...)
			var s GetBackupJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupJobsResultOutput)
}

// A collection of arguments for invoking getBackupJobs.
type GetBackupJobsOutputArgs struct {
	Filters GetBackupJobsFilterArrayInput `pulumi:"filters"`
	// A list of Backup Job IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The sort direction, sort results by ascending or descending order based on the value jobs id. Valid values: `ASCEND`, `DESCEND`.
	SortDirection pulumi.StringPtrInput `pulumi:"sortDirection"`
	// The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
	SourceType pulumi.StringInput `pulumi:"sourceType"`
	// The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetBackupJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupJobsArgs)(nil)).Elem()
}

// A collection of values returned by getBackupJobs.
type GetBackupJobsResultOutput struct{ *pulumi.OutputState }

func (GetBackupJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupJobsResult)(nil)).Elem()
}

func (o GetBackupJobsResultOutput) ToGetBackupJobsResultOutput() GetBackupJobsResultOutput {
	return o
}

func (o GetBackupJobsResultOutput) ToGetBackupJobsResultOutputWithContext(ctx context.Context) GetBackupJobsResultOutput {
	return o
}

func (o GetBackupJobsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBackupJobsResult] {
	return pulumix.Output[GetBackupJobsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetBackupJobsResultOutput) Filters() GetBackupJobsFilterArrayOutput {
	return o.ApplyT(func(v GetBackupJobsResult) []GetBackupJobsFilter { return v.Filters }).(GetBackupJobsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBackupJobsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBackupJobsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetBackupJobsResultOutput) Jobs() GetBackupJobsJobArrayOutput {
	return o.ApplyT(func(v GetBackupJobsResult) []GetBackupJobsJob { return v.Jobs }).(GetBackupJobsJobArrayOutput)
}

func (o GetBackupJobsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupJobsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBackupJobsResultOutput) SortDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupJobsResult) *string { return v.SortDirection }).(pulumi.StringPtrOutput)
}

func (o GetBackupJobsResultOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupJobsResult) string { return v.SourceType }).(pulumi.StringOutput)
}

func (o GetBackupJobsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupJobsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupJobsResultOutput{})
}
