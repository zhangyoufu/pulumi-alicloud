// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Cen Transit Router Cidrs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.193.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cen.GetTransitRouterCidrs(ctx, &cen.GetTransitRouterCidrsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				TransitRouterId: "tr-6ehx7q2jze8ch5ji0****",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTransitRouterCidrId0", ids.Cidrs[0].Id)
//			nameRegex, err := cen.GetTransitRouterCidrs(ctx, &cen.GetTransitRouterCidrsArgs{
//				NameRegex:       pulumi.StringRef("^my-name"),
//				TransitRouterId: "tr-6ehx7q2jze8ch5ji0****",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTransitRouterCidrId1", nameRegex.Cidrs[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetTransitRouterCidrs(ctx *pulumi.Context, args *GetTransitRouterCidrsArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterCidrsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTransitRouterCidrsResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterCidrs:getTransitRouterCidrs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterCidrs.
type GetTransitRouterCidrsArgs struct {
	// A list of Cen Transit Router Cidr IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Transit Router Cidr name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the transit router cidr.
	TransitRouterCidrId *string `pulumi:"transitRouterCidrId"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// A collection of values returned by getTransitRouterCidrs.
type GetTransitRouterCidrsResult struct {
	Cidrs []GetTransitRouterCidrsCidr `pulumi:"cidrs"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	Ids                 []string `pulumi:"ids"`
	NameRegex           *string  `pulumi:"nameRegex"`
	Names               []string `pulumi:"names"`
	OutputFile          *string  `pulumi:"outputFile"`
	TransitRouterCidrId *string  `pulumi:"transitRouterCidrId"`
	TransitRouterId     string   `pulumi:"transitRouterId"`
}

func GetTransitRouterCidrsOutput(ctx *pulumi.Context, args GetTransitRouterCidrsOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterCidrsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterCidrsResult, error) {
			args := v.(GetTransitRouterCidrsArgs)
			r, err := GetTransitRouterCidrs(ctx, &args, opts...)
			var s GetTransitRouterCidrsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterCidrsResultOutput)
}

// A collection of arguments for invoking getTransitRouterCidrs.
type GetTransitRouterCidrsOutputArgs struct {
	// A list of Cen Transit Router Cidr IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Transit Router Cidr name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the transit router cidr.
	TransitRouterCidrId pulumi.StringPtrInput `pulumi:"transitRouterCidrId"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput `pulumi:"transitRouterId"`
}

func (GetTransitRouterCidrsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterCidrsArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterCidrs.
type GetTransitRouterCidrsResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterCidrsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterCidrsResult)(nil)).Elem()
}

func (o GetTransitRouterCidrsResultOutput) ToGetTransitRouterCidrsResultOutput() GetTransitRouterCidrsResultOutput {
	return o
}

func (o GetTransitRouterCidrsResultOutput) ToGetTransitRouterCidrsResultOutputWithContext(ctx context.Context) GetTransitRouterCidrsResultOutput {
	return o
}

func (o GetTransitRouterCidrsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetTransitRouterCidrsResult] {
	return pulumix.Output[GetTransitRouterCidrsResult]{
		OutputState: o.OutputState,
	}
}

func (o GetTransitRouterCidrsResultOutput) Cidrs() GetTransitRouterCidrsCidrArrayOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) []GetTransitRouterCidrsCidr { return v.Cidrs }).(GetTransitRouterCidrsCidrArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterCidrsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTransitRouterCidrsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterCidrsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterCidrsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterCidrsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterCidrsResultOutput) TransitRouterCidrId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) *string { return v.TransitRouterCidrId }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterCidrsResultOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterCidrsResult) string { return v.TransitRouterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterCidrsResultOutput{})
}
