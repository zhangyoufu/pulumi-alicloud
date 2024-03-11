// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cen Transit Router Multicast Domains of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cen.GetTransitRouterMulticastDomains(ctx, &cen.GetTransitRouterMulticastDomainsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				TransitRouterId: "your_transit_router_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTransitRouterMulticastDomainId0", ids.Domains[0].Id)
//			nameRegex, err := cen.GetTransitRouterMulticastDomains(ctx, &cen.GetTransitRouterMulticastDomainsArgs{
//				NameRegex:       pulumi.StringRef("^my-name"),
//				TransitRouterId: "your_transit_router_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTransitRouterMulticastDomainId1", nameRegex.Domains[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetTransitRouterMulticastDomains(ctx *pulumi.Context, args *GetTransitRouterMulticastDomainsArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterMulticastDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTransitRouterMulticastDomainsResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterMulticastDomains:getTransitRouterMulticastDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterMulticastDomains.
type GetTransitRouterMulticastDomainsArgs struct {
	// A list of Transit Router Multicast Domain IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Transit Router Multicast Domain name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the multicast domain. Valid Value: `Active`.
	Status *string `pulumi:"status"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The ID of the multicast domain.
	TransitRouterMulticastDomainId *string `pulumi:"transitRouterMulticastDomainId"`
}

// A collection of values returned by getTransitRouterMulticastDomains.
type GetTransitRouterMulticastDomainsResult struct {
	// A list of Cen Transit Router Multicast Domains. Each element contains the following attributes:
	Domains []GetTransitRouterMulticastDomainsDomain `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Transit Router Multicast Domain names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the Transit Router Multicast Domain.
	Status *string `pulumi:"status"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The ID of the Transit Router Multicast Domain.
	TransitRouterMulticastDomainId *string `pulumi:"transitRouterMulticastDomainId"`
}

func GetTransitRouterMulticastDomainsOutput(ctx *pulumi.Context, args GetTransitRouterMulticastDomainsOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterMulticastDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterMulticastDomainsResult, error) {
			args := v.(GetTransitRouterMulticastDomainsArgs)
			r, err := GetTransitRouterMulticastDomains(ctx, &args, opts...)
			var s GetTransitRouterMulticastDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterMulticastDomainsResultOutput)
}

// A collection of arguments for invoking getTransitRouterMulticastDomains.
type GetTransitRouterMulticastDomainsOutputArgs struct {
	// A list of Transit Router Multicast Domain IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Transit Router Multicast Domain name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the multicast domain. Valid Value: `Active`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput `pulumi:"transitRouterId"`
	// The ID of the multicast domain.
	TransitRouterMulticastDomainId pulumi.StringPtrInput `pulumi:"transitRouterMulticastDomainId"`
}

func (GetTransitRouterMulticastDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterMulticastDomains.
type GetTransitRouterMulticastDomainsResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterMulticastDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainsResult)(nil)).Elem()
}

func (o GetTransitRouterMulticastDomainsResultOutput) ToGetTransitRouterMulticastDomainsResultOutput() GetTransitRouterMulticastDomainsResultOutput {
	return o
}

func (o GetTransitRouterMulticastDomainsResultOutput) ToGetTransitRouterMulticastDomainsResultOutputWithContext(ctx context.Context) GetTransitRouterMulticastDomainsResultOutput {
	return o
}

// A list of Cen Transit Router Multicast Domains. Each element contains the following attributes:
func (o GetTransitRouterMulticastDomainsResultOutput) Domains() GetTransitRouterMulticastDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) []GetTransitRouterMulticastDomainsDomain {
		return v.Domains
	}).(GetTransitRouterMulticastDomainsDomainArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterMulticastDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTransitRouterMulticastDomainsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterMulticastDomainsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Transit Router Multicast Domain names.
func (o GetTransitRouterMulticastDomainsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterMulticastDomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The status of the Transit Router Multicast Domain.
func (o GetTransitRouterMulticastDomainsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The ID of the transit router.
func (o GetTransitRouterMulticastDomainsResultOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) string { return v.TransitRouterId }).(pulumi.StringOutput)
}

// The ID of the Transit Router Multicast Domain.
func (o GetTransitRouterMulticastDomainsResultOutput) TransitRouterMulticastDomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainsResult) *string { return v.TransitRouterMulticastDomainId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterMulticastDomainsResultOutput{})
}
