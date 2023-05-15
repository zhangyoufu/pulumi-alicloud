// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Rds Replication Slots of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.204.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := rds.GetSlots(ctx, &rds.GetSlotsArgs{
//				DbInstanceId: "example_value",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRdsSlotsName", example.Slots[0].SlotName)
//			return nil
//		})
//	}
//
// ```
func GetSlots(ctx *pulumi.Context, args *GetSlotsArgs, opts ...pulumi.InvokeOption) (*GetSlotsResult, error) {
	var rv GetSlotsResult
	err := ctx.Invoke("alicloud:rds/getSlots:getSlots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSlots.
type GetSlotsArgs struct {
	// The db instance id.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

// A collection of values returned by getSlots.
type GetSlotsResult struct {
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id              string         `pulumi:"id"`
	OutputFile      *string        `pulumi:"outputFile"`
	ResourceGroupId *string        `pulumi:"resourceGroupId"`
	Slots           []GetSlotsSlot `pulumi:"slots"`
}

func GetSlotsOutput(ctx *pulumi.Context, args GetSlotsOutputArgs, opts ...pulumi.InvokeOption) GetSlotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSlotsResult, error) {
			args := v.(GetSlotsArgs)
			r, err := GetSlots(ctx, &args, opts...)
			var s GetSlotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSlotsResultOutput)
}

// A collection of arguments for invoking getSlots.
type GetSlotsOutputArgs struct {
	// The db instance id.
	DbInstanceId pulumi.StringInput `pulumi:"dbInstanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
}

func (GetSlotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlotsArgs)(nil)).Elem()
}

// A collection of values returned by getSlots.
type GetSlotsResultOutput struct{ *pulumi.OutputState }

func (GetSlotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSlotsResult)(nil)).Elem()
}

func (o GetSlotsResultOutput) ToGetSlotsResultOutput() GetSlotsResultOutput {
	return o
}

func (o GetSlotsResultOutput) ToGetSlotsResultOutputWithContext(ctx context.Context) GetSlotsResultOutput {
	return o
}

func (o GetSlotsResultOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlotsResult) string { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSlotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSlotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSlotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSlotsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSlotsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetSlotsResultOutput) Slots() GetSlotsSlotArrayOutput {
	return o.ApplyT(func(v GetSlotsResult) []GetSlotsSlot { return v.Slots }).(GetSlotsSlotArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSlotsResultOutput{})
}
