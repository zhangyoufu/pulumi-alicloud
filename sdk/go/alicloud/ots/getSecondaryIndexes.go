// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the ots secondary index of the current Alibaba Cloud user.
//
// For information about OTS secondary index and how to use it, see [Secondary index overview](https://www.alibabacloud.com/help/en/tablestore/latest/secondary-index-overview).
//
// > **NOTE:** Available in v1.187.0+.
func GetSecondaryIndexes(ctx *pulumi.Context, args *GetSecondaryIndexesArgs, opts ...pulumi.InvokeOption) (*GetSecondaryIndexesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSecondaryIndexesResult
	err := ctx.Invoke("alicloud:ots/getSecondaryIndexes:getSecondaryIndexes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecondaryIndexes.
type GetSecondaryIndexesArgs struct {
	// A list of secondary index IDs.
	Ids []string `pulumi:"ids"`
	// The name of OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// A regex string to filter results by secondary index name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The name of OTS table.
	TableName string `pulumi:"tableName"`
}

// A collection of values returned by getSecondaryIndexes.
type GetSecondaryIndexesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of secondary index IDs.
	Ids []string `pulumi:"ids"`
	// A list of indexes. Each element contains the following attributes:
	Indexes []GetSecondaryIndexesIndex `pulumi:"indexes"`
	// The OTS instance name.
	InstanceName string  `pulumi:"instanceName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of secondary index  names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The table name of the OTS which could not be changed.
	TableName string `pulumi:"tableName"`
}

func GetSecondaryIndexesOutput(ctx *pulumi.Context, args GetSecondaryIndexesOutputArgs, opts ...pulumi.InvokeOption) GetSecondaryIndexesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecondaryIndexesResult, error) {
			args := v.(GetSecondaryIndexesArgs)
			r, err := GetSecondaryIndexes(ctx, &args, opts...)
			var s GetSecondaryIndexesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecondaryIndexesResultOutput)
}

// A collection of arguments for invoking getSecondaryIndexes.
type GetSecondaryIndexesOutputArgs struct {
	// A list of secondary index IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of OTS instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// A regex string to filter results by secondary index name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of OTS table.
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (GetSecondaryIndexesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecondaryIndexesArgs)(nil)).Elem()
}

// A collection of values returned by getSecondaryIndexes.
type GetSecondaryIndexesResultOutput struct{ *pulumi.OutputState }

func (GetSecondaryIndexesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecondaryIndexesResult)(nil)).Elem()
}

func (o GetSecondaryIndexesResultOutput) ToGetSecondaryIndexesResultOutput() GetSecondaryIndexesResultOutput {
	return o
}

func (o GetSecondaryIndexesResultOutput) ToGetSecondaryIndexesResultOutputWithContext(ctx context.Context) GetSecondaryIndexesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecondaryIndexesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of secondary index IDs.
func (o GetSecondaryIndexesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of indexes. Each element contains the following attributes:
func (o GetSecondaryIndexesResultOutput) Indexes() GetSecondaryIndexesIndexArrayOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) []GetSecondaryIndexesIndex { return v.Indexes }).(GetSecondaryIndexesIndexArrayOutput)
}

// The OTS instance name.
func (o GetSecondaryIndexesResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetSecondaryIndexesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of secondary index  names.
func (o GetSecondaryIndexesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSecondaryIndexesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The table name of the OTS which could not be changed.
func (o GetSecondaryIndexesResultOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecondaryIndexesResult) string { return v.TableName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecondaryIndexesResultOutput{})
}
