// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides DMS Enterprise Database available to the user.[What is Database](https://www.alibabacloud.com/help/zh/data-management-service/latest/api-doc-dms-enterprise-2018-11-01-api-doc-listdatabases)
//
// > **NOTE:** Available in 1.195.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := dms.GetEnterpriseDatabases(ctx, &dms.GetEnterpriseDatabasesArgs{
//				NameRegex:  pulumi.StringRef("test2"),
//				InstanceId: "2195118",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudDmsEnterpriseDatabaseExampleId", _default.Databases[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEnterpriseDatabases(ctx *pulumi.Context, args *GetEnterpriseDatabasesArgs, opts ...pulumi.InvokeOption) (*GetEnterpriseDatabasesResult, error) {
	var rv GetEnterpriseDatabasesResult
	err := ctx.Invoke("alicloud:dms/getEnterpriseDatabases:getEnterpriseDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnterpriseDatabases.
type GetEnterpriseDatabasesArgs struct {
	// A list of Database IDs.
	Ids []string `pulumi:"ids"`
	// The instance ID of the target database.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter the results by the database Schema Name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEnterpriseDatabases.
type GetEnterpriseDatabasesResult struct {
	// A list of Database Entries. Each element contains the following attributes:
	Databases []GetEnterpriseDatabasesDatabase `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Database IDs.
	Ids []string `pulumi:"ids"`
	// The instance ID of the target database.
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

func GetEnterpriseDatabasesOutput(ctx *pulumi.Context, args GetEnterpriseDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetEnterpriseDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEnterpriseDatabasesResult, error) {
			args := v.(GetEnterpriseDatabasesArgs)
			r, err := GetEnterpriseDatabases(ctx, &args, opts...)
			var s GetEnterpriseDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEnterpriseDatabasesResultOutput)
}

// A collection of arguments for invoking getEnterpriseDatabases.
type GetEnterpriseDatabasesOutputArgs struct {
	// A list of Database IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The instance ID of the target database.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter the results by the database Schema Name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetEnterpriseDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnterpriseDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getEnterpriseDatabases.
type GetEnterpriseDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetEnterpriseDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnterpriseDatabasesResult)(nil)).Elem()
}

func (o GetEnterpriseDatabasesResultOutput) ToGetEnterpriseDatabasesResultOutput() GetEnterpriseDatabasesResultOutput {
	return o
}

func (o GetEnterpriseDatabasesResultOutput) ToGetEnterpriseDatabasesResultOutputWithContext(ctx context.Context) GetEnterpriseDatabasesResultOutput {
	return o
}

// A list of Database Entries. Each element contains the following attributes:
func (o GetEnterpriseDatabasesResultOutput) Databases() GetEnterpriseDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) []GetEnterpriseDatabasesDatabase { return v.Databases }).(GetEnterpriseDatabasesDatabaseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEnterpriseDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Database IDs.
func (o GetEnterpriseDatabasesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The instance ID of the target database.
func (o GetEnterpriseDatabasesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetEnterpriseDatabasesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEnterpriseDatabasesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnterpriseDatabasesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEnterpriseDatabasesResultOutput{})
}
