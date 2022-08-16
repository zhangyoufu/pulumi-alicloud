// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Firewall Address Books of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.178.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cloudfirewall.GetAddressBooks(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cloudFirewallAddressBookId1", ids.Books[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetAddressBooks(ctx *pulumi.Context, args *GetAddressBooksArgs, opts ...pulumi.InvokeOption) (*GetAddressBooksResult, error) {
	var rv GetAddressBooksResult
	err := ctx.Invoke("alicloud:cloudfirewall/getAddressBooks:getAddressBooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddressBooks.
type GetAddressBooksArgs struct {
	// The type of the Address Book.
	GroupType *string `pulumi:"groupType"`
	// A list of Address Book IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results Address Book name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAddressBooks.
type GetAddressBooksResult struct {
	Books     []GetAddressBooksBook `pulumi:"books"`
	GroupType *string               `pulumi:"groupType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetAddressBooksOutput(ctx *pulumi.Context, args GetAddressBooksOutputArgs, opts ...pulumi.InvokeOption) GetAddressBooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddressBooksResult, error) {
			args := v.(GetAddressBooksArgs)
			r, err := GetAddressBooks(ctx, &args, opts...)
			var s GetAddressBooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddressBooksResultOutput)
}

// A collection of arguments for invoking getAddressBooks.
type GetAddressBooksOutputArgs struct {
	// The type of the Address Book.
	GroupType pulumi.StringPtrInput `pulumi:"groupType"`
	// A list of Address Book IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results Address Book name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAddressBooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressBooksArgs)(nil)).Elem()
}

// A collection of values returned by getAddressBooks.
type GetAddressBooksResultOutput struct{ *pulumi.OutputState }

func (GetAddressBooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressBooksResult)(nil)).Elem()
}

func (o GetAddressBooksResultOutput) ToGetAddressBooksResultOutput() GetAddressBooksResultOutput {
	return o
}

func (o GetAddressBooksResultOutput) ToGetAddressBooksResultOutputWithContext(ctx context.Context) GetAddressBooksResultOutput {
	return o
}

func (o GetAddressBooksResultOutput) Books() GetAddressBooksBookArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []GetAddressBooksBook { return v.Books }).(GetAddressBooksBookArrayOutput)
}

func (o GetAddressBooksResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddressBooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddressBooksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAddressBooksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAddressBooksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAddressBooksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAddressBooksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddressBooksResultOutput{})
}
