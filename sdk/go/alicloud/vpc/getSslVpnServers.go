// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SSL-VPN servers data source lists lots of SSL-VPN servers resource information owned by an Alicloud account.
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
//			_, err := vpc.GetSslVpnServers(ctx, &vpc.GetSslVpnServersArgs{
//				Ids: []string{
//					"fake-server-id",
//				},
//				VpnGatewayId: pulumi.StringRef("fake-vpn-id"),
//				OutputFile:   pulumi.StringRef("/tmp/sslserver"),
//				NameRegex:    pulumi.StringRef("^foo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetSslVpnServers(ctx *pulumi.Context, args *GetSslVpnServersArgs, opts ...pulumi.InvokeOption) (*GetSslVpnServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSslVpnServersResult
	err := ctx.Invoke("alicloud:vpc/getSslVpnServers:getSslVpnServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSslVpnServers.
type GetSslVpnServersArgs struct {
	// IDs of the SSL-VPN servers.
	Ids []string `pulumi:"ids"`
	// A regex string of SSL-VPN server name.
	NameRegex *string `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile *string `pulumi:"outputFile"`
	// Use the VPN gateway ID as the search key.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// A collection of values returned by getSslVpnServers.
type GetSslVpnServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SSL-VPN server IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of SSL-VPN server names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of SSL-VPN servers. Each element contains the following attributes:
	Servers []GetSslVpnServersServer `pulumi:"servers"`
	// The ID of the VPN gateway instance.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

func GetSslVpnServersOutput(ctx *pulumi.Context, args GetSslVpnServersOutputArgs, opts ...pulumi.InvokeOption) GetSslVpnServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSslVpnServersResult, error) {
			args := v.(GetSslVpnServersArgs)
			r, err := GetSslVpnServers(ctx, &args, opts...)
			var s GetSslVpnServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSslVpnServersResultOutput)
}

// A collection of arguments for invoking getSslVpnServers.
type GetSslVpnServersOutputArgs struct {
	// IDs of the SSL-VPN servers.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string of SSL-VPN server name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Use the VPN gateway ID as the search key.
	VpnGatewayId pulumi.StringPtrInput `pulumi:"vpnGatewayId"`
}

func (GetSslVpnServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSslVpnServersArgs)(nil)).Elem()
}

// A collection of values returned by getSslVpnServers.
type GetSslVpnServersResultOutput struct{ *pulumi.OutputState }

func (GetSslVpnServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSslVpnServersResult)(nil)).Elem()
}

func (o GetSslVpnServersResultOutput) ToGetSslVpnServersResultOutput() GetSslVpnServersResultOutput {
	return o
}

func (o GetSslVpnServersResultOutput) ToGetSslVpnServersResultOutputWithContext(ctx context.Context) GetSslVpnServersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSslVpnServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SSL-VPN server IDs.
func (o GetSslVpnServersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSslVpnServersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of SSL-VPN server names.
func (o GetSslVpnServersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSslVpnServersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of SSL-VPN servers. Each element contains the following attributes:
func (o GetSslVpnServersResultOutput) Servers() GetSslVpnServersServerArrayOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) []GetSslVpnServersServer { return v.Servers }).(GetSslVpnServersServerArrayOutput)
}

// The ID of the VPN gateway instance.
func (o GetSslVpnServersResultOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSslVpnServersResult) *string { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSslVpnServersResultOutput{})
}
