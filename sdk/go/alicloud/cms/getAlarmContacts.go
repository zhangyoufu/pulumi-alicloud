// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of alarm contact owned by an Alibaba Cloud account.
//
// > **NOTE:** Available in v1.99.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cms.GetAlarmContacts(ctx, &cms.GetAlarmContactsArgs{
//				Ids: []string{
//					"tf-testAccCmsAlarmContact",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("first-contact", this.Contacts)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAlarmContacts(ctx *pulumi.Context, args *GetAlarmContactsArgs, opts ...pulumi.InvokeOption) (*GetAlarmContactsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAlarmContactsResult
	err := ctx.Invoke("alicloud:cms/getAlarmContacts:getAlarmContacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmContacts.
type GetAlarmContactsArgs struct {
	// The alarm notification method. Alarm notifications can be sent by using `Email` or `DingWebHook`.
	ChanelType *string `pulumi:"chanelType"`
	// The alarm notification target.
	ChanelValue *string `pulumi:"chanelValue"`
	// A list of alarm contact IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by alarm contact name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	//
	// > **NOTE:** Specify at least one of the following alarm notification targets: phone number, email address, webhook URL of the DingTalk chatbot, and TradeManager ID.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAlarmContacts.
type GetAlarmContactsResult struct {
	ChanelType  *string `pulumi:"chanelType"`
	ChanelValue *string `pulumi:"chanelValue"`
	// A list of alarm contacts. Each element contains the following attributes:
	Contacts []GetAlarmContactsContact `pulumi:"contacts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of alarm contact IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of alarm contact names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetAlarmContactsOutput(ctx *pulumi.Context, args GetAlarmContactsOutputArgs, opts ...pulumi.InvokeOption) GetAlarmContactsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAlarmContactsResult, error) {
			args := v.(GetAlarmContactsArgs)
			r, err := GetAlarmContacts(ctx, &args, opts...)
			var s GetAlarmContactsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAlarmContactsResultOutput)
}

// A collection of arguments for invoking getAlarmContacts.
type GetAlarmContactsOutputArgs struct {
	// The alarm notification method. Alarm notifications can be sent by using `Email` or `DingWebHook`.
	ChanelType pulumi.StringPtrInput `pulumi:"chanelType"`
	// The alarm notification target.
	ChanelValue pulumi.StringPtrInput `pulumi:"chanelValue"`
	// A list of alarm contact IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by alarm contact name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	//
	// > **NOTE:** Specify at least one of the following alarm notification targets: phone number, email address, webhook URL of the DingTalk chatbot, and TradeManager ID.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAlarmContactsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmContactsArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmContacts.
type GetAlarmContactsResultOutput struct{ *pulumi.OutputState }

func (GetAlarmContactsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAlarmContactsResult)(nil)).Elem()
}

func (o GetAlarmContactsResultOutput) ToGetAlarmContactsResultOutput() GetAlarmContactsResultOutput {
	return o
}

func (o GetAlarmContactsResultOutput) ToGetAlarmContactsResultOutputWithContext(ctx context.Context) GetAlarmContactsResultOutput {
	return o
}

func (o GetAlarmContactsResultOutput) ChanelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) *string { return v.ChanelType }).(pulumi.StringPtrOutput)
}

func (o GetAlarmContactsResultOutput) ChanelValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) *string { return v.ChanelValue }).(pulumi.StringPtrOutput)
}

// A list of alarm contacts. Each element contains the following attributes:
func (o GetAlarmContactsResultOutput) Contacts() GetAlarmContactsContactArrayOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) []GetAlarmContactsContact { return v.Contacts }).(GetAlarmContactsContactArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAlarmContactsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of alarm contact IDs.
func (o GetAlarmContactsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAlarmContactsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of alarm contact names.
func (o GetAlarmContactsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAlarmContactsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAlarmContactsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAlarmContactsResultOutput{})
}
