// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Actiontrail History Delivery Jobs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.139.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/actiontrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := actiontrail.GetHistoryDeliveryJobs(ctx, &actiontrail.GetHistoryDeliveryJobsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("actiontrailHistoryDeliveryJobId1", ids.Jobs[0].Id)
//			status, err := actiontrail.GetHistoryDeliveryJobs(ctx, &actiontrail.GetHistoryDeliveryJobsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				Status: pulumi.IntRef(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("actiontrailHistoryDeliveryJobId2", status.Jobs[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetHistoryDeliveryJobs(ctx *pulumi.Context, args *GetHistoryDeliveryJobsArgs, opts ...pulumi.InvokeOption) (*GetHistoryDeliveryJobsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHistoryDeliveryJobsResult
	err := ctx.Invoke("alicloud:actiontrail/getHistoryDeliveryJobs:getHistoryDeliveryJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHistoryDeliveryJobs.
type GetHistoryDeliveryJobsArgs struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of History Delivery Job IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getHistoryDeliveryJobs.
type GetHistoryDeliveryJobsResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                      `pulumi:"id"`
	Ids        []string                    `pulumi:"ids"`
	Jobs       []GetHistoryDeliveryJobsJob `pulumi:"jobs"`
	OutputFile *string                     `pulumi:"outputFile"`
	Status     *int                        `pulumi:"status"`
}

func GetHistoryDeliveryJobsOutput(ctx *pulumi.Context, args GetHistoryDeliveryJobsOutputArgs, opts ...pulumi.InvokeOption) GetHistoryDeliveryJobsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHistoryDeliveryJobsResult, error) {
			args := v.(GetHistoryDeliveryJobsArgs)
			r, err := GetHistoryDeliveryJobs(ctx, &args, opts...)
			var s GetHistoryDeliveryJobsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHistoryDeliveryJobsResultOutput)
}

// A collection of arguments for invoking getHistoryDeliveryJobs.
type GetHistoryDeliveryJobsOutputArgs struct {
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of History Delivery Job IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
	Status pulumi.IntPtrInput `pulumi:"status"`
}

func (GetHistoryDeliveryJobsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHistoryDeliveryJobsArgs)(nil)).Elem()
}

// A collection of values returned by getHistoryDeliveryJobs.
type GetHistoryDeliveryJobsResultOutput struct{ *pulumi.OutputState }

func (GetHistoryDeliveryJobsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHistoryDeliveryJobsResult)(nil)).Elem()
}

func (o GetHistoryDeliveryJobsResultOutput) ToGetHistoryDeliveryJobsResultOutput() GetHistoryDeliveryJobsResultOutput {
	return o
}

func (o GetHistoryDeliveryJobsResultOutput) ToGetHistoryDeliveryJobsResultOutputWithContext(ctx context.Context) GetHistoryDeliveryJobsResultOutput {
	return o
}

func (o GetHistoryDeliveryJobsResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHistoryDeliveryJobsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetHistoryDeliveryJobsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetHistoryDeliveryJobsResultOutput) Jobs() GetHistoryDeliveryJobsJobArrayOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) []GetHistoryDeliveryJobsJob { return v.Jobs }).(GetHistoryDeliveryJobsJobArrayOutput)
}

func (o GetHistoryDeliveryJobsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetHistoryDeliveryJobsResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetHistoryDeliveryJobsResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHistoryDeliveryJobsResultOutput{})
}
