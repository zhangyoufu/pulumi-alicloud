// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list of KMS keys in an Alibaba Cloud account according to the specified filters.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Declare the data source
//			kmsKeysDs, err := kms.GetKeys(ctx, &kms.GetKeysArgs{
//				DescriptionRegex: pulumi.StringRef("Hello KMS"),
//				OutputFile:       pulumi.StringRef("kms_keys.json"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstKeyId", kmsKeysDs.Keys[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetKeys(ctx *pulumi.Context, args *GetKeysArgs, opts ...pulumi.InvokeOption) (*GetKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKeysResult
	err := ctx.Invoke("alicloud:kms/getKeys:getKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeys.
type GetKeysArgs struct {
	// A regex string to filter the results by the KMS key description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	// The CMK filter. The filter consists of one or more key-value pairs.
	// You can specify a maximum of 10 key-value pairs. More details see API [ListKeys](https://www.alibabacloud.com/help/en/key-management-service/latest/listkeys).
	Filters *string `pulumi:"filters"`
	// A list of KMS key IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getKeys.
type GetKeysResult struct {
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	Filters          *string `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of KMS key IDs.
	Ids []string `pulumi:"ids"`
	// A list of KMS keys. Each element contains the following attributes:
	Keys       []GetKeysKey `pulumi:"keys"`
	OutputFile *string      `pulumi:"outputFile"`
	// Status of the key. Possible values: `Enabled`, `Disabled` and `PendingDeletion`.
	Status *string `pulumi:"status"`
}

func GetKeysOutput(ctx *pulumi.Context, args GetKeysOutputArgs, opts ...pulumi.InvokeOption) GetKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKeysResult, error) {
			args := v.(GetKeysArgs)
			r, err := GetKeys(ctx, &args, opts...)
			var s GetKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKeysResultOutput)
}

// A collection of arguments for invoking getKeys.
type GetKeysOutputArgs struct {
	// A regex string to filter the results by the KMS key description.
	DescriptionRegex pulumi.StringPtrInput `pulumi:"descriptionRegex"`
	EnableDetails    pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// The CMK filter. The filter consists of one or more key-value pairs.
	// You can specify a maximum of 10 key-value pairs. More details see API [ListKeys](https://www.alibabacloud.com/help/en/key-management-service/latest/listkeys).
	Filters pulumi.StringPtrInput `pulumi:"filters"`
	// A list of KMS key IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Filter the results by status of the KMS keys. Valid values: `Enabled`, `Disabled`, `PendingDeletion`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeysArgs)(nil)).Elem()
}

// A collection of values returned by getKeys.
type GetKeysResultOutput struct{ *pulumi.OutputState }

func (GetKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeysResult)(nil)).Elem()
}

func (o GetKeysResultOutput) ToGetKeysResultOutput() GetKeysResultOutput {
	return o
}

func (o GetKeysResultOutput) ToGetKeysResultOutputWithContext(ctx context.Context) GetKeysResultOutput {
	return o
}

func (o GetKeysResultOutput) DescriptionRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeysResult) *string { return v.DescriptionRegex }).(pulumi.StringPtrOutput)
}

func (o GetKeysResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetKeysResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetKeysResultOutput) Filters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeysResult) *string { return v.Filters }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of KMS key IDs.
func (o GetKeysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKeysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of KMS keys. Each element contains the following attributes:
func (o GetKeysResultOutput) Keys() GetKeysKeyArrayOutput {
	return o.ApplyT(func(v GetKeysResult) []GetKeysKey { return v.Keys }).(GetKeysKeyArrayOutput)
}

func (o GetKeysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Status of the key. Possible values: `Enabled`, `Disabled` and `PendingDeletion`.
func (o GetKeysResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKeysResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKeysResultOutput{})
}
