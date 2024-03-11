// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Mongodb Accounts of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.148.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := mongodb.GetAccounts(ctx, &mongodb.GetAccountsArgs{
//				InstanceId:  "example_value",
//				AccountName: pulumi.StringRef("root"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("mongodbAccountId1", example.Accounts[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAccounts(ctx *pulumi.Context, args *GetAccountsArgs, opts ...pulumi.InvokeOption) (*GetAccountsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountsResult
	err := ctx.Invoke("alicloud:mongodb/getAccounts:getAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccounts.
type GetAccountsArgs struct {
	// The name of the account.
	AccountName *string `pulumi:"accountName"`
	// The id of the instance to which the account belongs.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAccounts.
type GetAccountsResult struct {
	AccountName *string              `pulumi:"accountName"`
	Accounts    []GetAccountsAccount `pulumi:"accounts"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
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
	// The name of the account.
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	// The id of the instance to which the account belongs.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
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

func (o GetAccountsResultOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o GetAccountsResultOutput) Accounts() GetAccountsAccountArrayOutput {
	return o.ApplyT(func(v GetAccountsResult) []GetAccountsAccount { return v.Accounts }).(GetAccountsAccountArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccountsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetAccountsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccountsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountsResultOutput{})
}
