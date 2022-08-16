// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Rds Accounts of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.120.0+.
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
//			example, err := rds.GetAccounts(ctx, &rds.GetAccountsArgs{
//				DbInstanceId: "example_value",
//				NameRegex:    pulumi.StringRef("the_resource_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRdsAccountId", example.Accounts[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAccounts(ctx *pulumi.Context, args *GetAccountsArgs, opts ...pulumi.InvokeOption) (*GetAccountsResult, error) {
	var rv GetAccountsResult
	err := ctx.Invoke("alicloud:rds/getAccounts:getAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccounts.
type GetAccountsArgs struct {
	// The db instance id.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// A list of Account IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Account name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getAccounts.
type GetAccountsResult struct {
	Accounts     []GetAccountsAccount `pulumi:"accounts"`
	DbInstanceId string               `pulumi:"dbInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetAccountsOutput(ctx *pulumi.Context, args GetAccountsOutputArgs, opts ...pulumi.InvokeOption) GetAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountsResult, error) {
			args := v.(GetAccountsArgs)
			r, err := GetAccounts(ctx, &args, opts...)
			var s GetAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccountsResultOutput)
}

// A collection of arguments for invoking getAccounts.
type GetAccountsOutputArgs struct {
	// The db instance id.
	DbInstanceId pulumi.StringInput `pulumi:"dbInstanceId"`
	// A list of Account IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Account name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getAccounts.
type GetAccountsResultOutput struct{ *pulumi.OutputState }

func (GetAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountsResult)(nil)).Elem()
}

func (o GetAccountsResultOutput) ToGetAccountsResultOutput() GetAccountsResultOutput {
	return o
}

func (o GetAccountsResultOutput) ToGetAccountsResultOutputWithContext(ctx context.Context) GetAccountsResultOutput {
	return o
}

func (o GetAccountsResultOutput) Accounts() GetAccountsAccountArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []GetAccountsAccount { return v.Accounts }).(GetAccountsAccountArrayOutput)
}

func (o GetAccountsResultOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAccountsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAccountsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAccountsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAccountsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountsResultOutput{})
}
