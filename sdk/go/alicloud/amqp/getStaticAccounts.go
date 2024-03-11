// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amqp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Amqp Static Account available to the user.[What is Static Account](https://help.aliyun.com/document_detail/184399.html)
//
// > **NOTE:** Available in 1.195.0+
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := amqp.GetStaticAccounts(ctx, &amqp.GetStaticAccountsArgs{
//				InstanceId: pulumi.StringRef("amqp-cn-0ju2y01zs001"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudAmqpStaticAccountExampleId", _default.Accounts[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetStaticAccounts(ctx *pulumi.Context, args *GetStaticAccountsArgs, opts ...pulumi.InvokeOption) (*GetStaticAccountsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStaticAccountsResult
	err := ctx.Invoke("alicloud:amqp/getStaticAccounts:getStaticAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStaticAccounts.
type GetStaticAccountsArgs struct {
	// The `key` of the resource supplied above.The value is formulated as `<instance_id>:<access_key>`.
	Ids []string `pulumi:"ids"`
	// InstanceId
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getStaticAccounts.
type GetStaticAccountsResult struct {
	// A list of Static Account Entries. Each element contains the following attributes:
	Accounts []GetStaticAccountsAccount `pulumi:"accounts"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// Amqp instance ID.
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
}

func GetStaticAccountsOutput(ctx *pulumi.Context, args GetStaticAccountsOutputArgs, opts ...pulumi.InvokeOption) GetStaticAccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStaticAccountsResult, error) {
			args := v.(GetStaticAccountsArgs)
			r, err := GetStaticAccounts(ctx, &args, opts...)
			var s GetStaticAccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStaticAccountsResultOutput)
}

// A collection of arguments for invoking getStaticAccounts.
type GetStaticAccountsOutputArgs struct {
	// The `key` of the resource supplied above.The value is formulated as `<instance_id>:<access_key>`.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// InstanceId
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetStaticAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getStaticAccounts.
type GetStaticAccountsResultOutput struct{ *pulumi.OutputState }

func (GetStaticAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticAccountsResult)(nil)).Elem()
}

func (o GetStaticAccountsResultOutput) ToGetStaticAccountsResultOutput() GetStaticAccountsResultOutput {
	return o
}

func (o GetStaticAccountsResultOutput) ToGetStaticAccountsResultOutputWithContext(ctx context.Context) GetStaticAccountsResultOutput {
	return o
}

// A list of Static Account Entries. Each element contains the following attributes:
func (o GetStaticAccountsResultOutput) Accounts() GetStaticAccountsAccountArrayOutput {
	return o.ApplyT(func(v GetStaticAccountsResult) []GetStaticAccountsAccount { return v.Accounts }).(GetStaticAccountsAccountArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStaticAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStaticAccountsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStaticAccountsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Amqp instance ID.
func (o GetStaticAccountsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticAccountsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetStaticAccountsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticAccountsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStaticAccountsResultOutput{})
}
