// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open Resource Manager Account Deletion Check Task.
//
// For information about Resource Manager Account Deletion Check Task and how to use it, see [What is Resource Manager Account Deletion Check Task](https://www.alibabacloud.com/help/en/resource-management/latest/check-account-delete).
//
// > **NOTE:** Available in v1.187.0+.
//
// > **NOTE:** The member deletion feature is in invitational preview. You can contact the service manager of Alibaba Cloud to apply for a trial.
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
//			task, err := resourcemanager.GetAccountDeletionCheckTask(ctx, &resourcemanager.GetAccountDeletionCheckTaskArgs{
//				AccountId: "your_account_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range task.AbandonAbleChecks {
//				splat0 = append(splat0, val0.CheckId)
//			}
//			ctx.Export("abandonAbleChecksIds", splat0)
//			return nil
//		})
//	}
//
// ```
func GetAccountDeletionCheckTask(ctx *pulumi.Context, args *GetAccountDeletionCheckTaskArgs, opts ...pulumi.InvokeOption) (*GetAccountDeletionCheckTaskResult, error) {
	var rv GetAccountDeletionCheckTaskResult
	err := ctx.Invoke("alicloud:resourcemanager/getAccountDeletionCheckTask:getAccountDeletionCheckTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountDeletionCheckTask.
type GetAccountDeletionCheckTaskArgs struct {
	// The ID of the member that you want to delete.
	AccountId string `pulumi:"accountId"`
}

// A collection of values returned by getAccountDeletionCheckTask.
type GetAccountDeletionCheckTaskResult struct {
	// The check items that you can choose to ignore for the member deletion. Each element contains the following attributes:
	AbandonAbleChecks []GetAccountDeletionCheckTaskAbandonAbleCheck `pulumi:"abandonAbleChecks"`
	AccountId         string                                        `pulumi:"accountId"`
	// Indicates whether the member can be deleted.
	AllowDelete bool `pulumi:"allowDelete"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The reasons why the member cannot be deleted. Each element contains the following attributes:
	NotAllowReasons []GetAccountDeletionCheckTaskNotAllowReason `pulumi:"notAllowReasons"`
	// The status of the check.
	Status string `pulumi:"status"`
}

func GetAccountDeletionCheckTaskOutput(ctx *pulumi.Context, args GetAccountDeletionCheckTaskOutputArgs, opts ...pulumi.InvokeOption) GetAccountDeletionCheckTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountDeletionCheckTaskResult, error) {
			args := v.(GetAccountDeletionCheckTaskArgs)
			r, err := GetAccountDeletionCheckTask(ctx, &args, opts...)
			var s GetAccountDeletionCheckTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccountDeletionCheckTaskResultOutput)
}

// A collection of arguments for invoking getAccountDeletionCheckTask.
type GetAccountDeletionCheckTaskOutputArgs struct {
	// The ID of the member that you want to delete.
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (GetAccountDeletionCheckTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountDeletionCheckTaskArgs)(nil)).Elem()
}

// A collection of values returned by getAccountDeletionCheckTask.
type GetAccountDeletionCheckTaskResultOutput struct{ *pulumi.OutputState }

func (GetAccountDeletionCheckTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountDeletionCheckTaskResult)(nil)).Elem()
}

func (o GetAccountDeletionCheckTaskResultOutput) ToGetAccountDeletionCheckTaskResultOutput() GetAccountDeletionCheckTaskResultOutput {
	return o
}

func (o GetAccountDeletionCheckTaskResultOutput) ToGetAccountDeletionCheckTaskResultOutputWithContext(ctx context.Context) GetAccountDeletionCheckTaskResultOutput {
	return o
}

// The check items that you can choose to ignore for the member deletion. Each element contains the following attributes:
func (o GetAccountDeletionCheckTaskResultOutput) AbandonAbleChecks() GetAccountDeletionCheckTaskAbandonAbleCheckArrayOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) []GetAccountDeletionCheckTaskAbandonAbleCheck {
		return v.AbandonAbleChecks
	}).(GetAccountDeletionCheckTaskAbandonAbleCheckArrayOutput)
}

func (o GetAccountDeletionCheckTaskResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// Indicates whether the member can be deleted.
func (o GetAccountDeletionCheckTaskResultOutput) AllowDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) bool { return v.AllowDelete }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountDeletionCheckTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

// The reasons why the member cannot be deleted. Each element contains the following attributes:
func (o GetAccountDeletionCheckTaskResultOutput) NotAllowReasons() GetAccountDeletionCheckTaskNotAllowReasonArrayOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) []GetAccountDeletionCheckTaskNotAllowReason {
		return v.NotAllowReasons
	}).(GetAccountDeletionCheckTaskNotAllowReasonArrayOutput)
}

// The status of the check.
func (o GetAccountDeletionCheckTaskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountDeletionCheckTaskResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountDeletionCheckTaskResultOutput{})
}
