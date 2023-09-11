// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directmail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Direct Mail Receiverses of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.125.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/directmail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := directmail.LookupReceivers(ctx, &directmail.LookupReceiversArgs{
//				Ids: []string{
//					"ca73b1e4fb0df7c935a5097a****",
//				},
//				NameRegex: pulumi.StringRef("the_resource_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstDirectMailReceiversId", example.Receiverses[0].Id)
//			return nil
//		})
//	}
//
// ```
func LookupReceivers(ctx *pulumi.Context, args *LookupReceiversArgs, opts ...pulumi.InvokeOption) (*LookupReceiversResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReceiversResult
	err := ctx.Invoke("alicloud:directmail/getReceivers:getReceivers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReceivers.
type LookupReceiversArgs struct {
	// A list of Receivers IDs.
	Ids []string `pulumi:"ids"`
	// The key word.
	KeyWord *string `pulumi:"keyWord"`
	// A regex string to filter results by Receivers name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getReceivers.
type LookupReceiversResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string                   `pulumi:"id"`
	Ids         []string                 `pulumi:"ids"`
	KeyWord     *string                  `pulumi:"keyWord"`
	NameRegex   *string                  `pulumi:"nameRegex"`
	Names       []string                 `pulumi:"names"`
	OutputFile  *string                  `pulumi:"outputFile"`
	Receiverses []GetReceiversReceiverse `pulumi:"receiverses"`
	Status      *int                     `pulumi:"status"`
}

func LookupReceiversOutput(ctx *pulumi.Context, args LookupReceiversOutputArgs, opts ...pulumi.InvokeOption) LookupReceiversResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReceiversResult, error) {
			args := v.(LookupReceiversArgs)
			r, err := LookupReceivers(ctx, &args, opts...)
			var s LookupReceiversResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReceiversResultOutput)
}

// A collection of arguments for invoking getReceivers.
type LookupReceiversOutputArgs struct {
	// A list of Receivers IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The key word.
	KeyWord pulumi.StringPtrInput `pulumi:"keyWord"`
	// A regex string to filter results by Receivers name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.IntPtrInput `pulumi:"status"`
}

func (LookupReceiversOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiversArgs)(nil)).Elem()
}

// A collection of values returned by getReceivers.
type LookupReceiversResultOutput struct{ *pulumi.OutputState }

func (LookupReceiversResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiversResult)(nil)).Elem()
}

func (o LookupReceiversResultOutput) ToLookupReceiversResultOutput() LookupReceiversResultOutput {
	return o
}

func (o LookupReceiversResultOutput) ToLookupReceiversResultOutputWithContext(ctx context.Context) LookupReceiversResultOutput {
	return o
}

func (o LookupReceiversResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReceiversResult] {
	return pulumix.Output[LookupReceiversResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReceiversResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReceiversResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReceiversResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupReceiversResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o LookupReceiversResultOutput) KeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiversResult) *string { return v.KeyWord }).(pulumi.StringPtrOutput)
}

func (o LookupReceiversResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiversResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o LookupReceiversResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupReceiversResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o LookupReceiversResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiversResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupReceiversResultOutput) Receiverses() GetReceiversReceiverseArrayOutput {
	return o.ApplyT(func(v LookupReceiversResult) []GetReceiversReceiverse { return v.Receiverses }).(GetReceiversReceiverseArrayOutput)
}

func (o LookupReceiversResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupReceiversResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReceiversResultOutput{})
}
