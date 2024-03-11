// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE:** Available in v1.162.0+.
//
// The data source lists a number of VPN Pbr Route Entries resource information owned by an Alicloud account.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := vpc.GetPbrRouteEntries(ctx, &vpc.GetPbrRouteEntriesArgs{
//				VpnGatewayId: "example_vpn_gateway_id",
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpnIpsecServerId1", ids.Entries[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetPbrRouteEntries(ctx *pulumi.Context, args *GetPbrRouteEntriesArgs, opts ...pulumi.InvokeOption) (*GetPbrRouteEntriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPbrRouteEntriesResult
	err := ctx.Invoke("alicloud:vpc/getPbrRouteEntries:getPbrRouteEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPbrRouteEntries.
type GetPbrRouteEntriesArgs struct {
	// A list of VPN Pbr Route Entries IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the VPN gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// A collection of values returned by getPbrRouteEntries.
type GetPbrRouteEntriesResult struct {
	// A list of VPN Pbr Route Entries. Each element contains the following attributes:
	Entries []GetPbrRouteEntriesEntry `pulumi:"entries"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the vpn gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

func GetPbrRouteEntriesOutput(ctx *pulumi.Context, args GetPbrRouteEntriesOutputArgs, opts ...pulumi.InvokeOption) GetPbrRouteEntriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPbrRouteEntriesResult, error) {
			args := v.(GetPbrRouteEntriesArgs)
			r, err := GetPbrRouteEntries(ctx, &args, opts...)
			var s GetPbrRouteEntriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPbrRouteEntriesResultOutput)
}

// A collection of arguments for invoking getPbrRouteEntries.
type GetPbrRouteEntriesOutputArgs struct {
	// A list of VPN Pbr Route Entries IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

func (GetPbrRouteEntriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPbrRouteEntriesArgs)(nil)).Elem()
}

// A collection of values returned by getPbrRouteEntries.
type GetPbrRouteEntriesResultOutput struct{ *pulumi.OutputState }

func (GetPbrRouteEntriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPbrRouteEntriesResult)(nil)).Elem()
}

func (o GetPbrRouteEntriesResultOutput) ToGetPbrRouteEntriesResultOutput() GetPbrRouteEntriesResultOutput {
	return o
}

func (o GetPbrRouteEntriesResultOutput) ToGetPbrRouteEntriesResultOutputWithContext(ctx context.Context) GetPbrRouteEntriesResultOutput {
	return o
}

// A list of VPN Pbr Route Entries. Each element contains the following attributes:
func (o GetPbrRouteEntriesResultOutput) Entries() GetPbrRouteEntriesEntryArrayOutput {
	return o.ApplyT(func(v GetPbrRouteEntriesResult) []GetPbrRouteEntriesEntry { return v.Entries }).(GetPbrRouteEntriesEntryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPbrRouteEntriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPbrRouteEntriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPbrRouteEntriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPbrRouteEntriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPbrRouteEntriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPbrRouteEntriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the vpn gateway.
func (o GetPbrRouteEntriesResultOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPbrRouteEntriesResult) string { return v.VpnGatewayId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPbrRouteEntriesResultOutput{})
}
