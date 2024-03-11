// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Vpc Ipv6 Gateways of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.142.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := vpc.GetIpv6Gateways(ctx, &vpc.GetIpv6GatewaysArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6GatewayId1", ids.Gateways[0].Id)
//			nameRegex, err := vpc.GetIpv6Gateways(ctx, &vpc.GetIpv6GatewaysArgs{
//				NameRegex: pulumi.StringRef("^my-Ipv6Gateway"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6GatewayId2", nameRegex.Gateways[0].Id)
//			vpcId, err := vpc.GetIpv6Gateways(ctx, &vpc.GetIpv6GatewaysArgs{
//				Ids: []string{
//					"example_id",
//				},
//				VpcId: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6GatewayId3", vpcId.Gateways[0].Id)
//			status, err := vpc.GetIpv6Gateways(ctx, &vpc.GetIpv6GatewaysArgs{
//				Ids: []string{
//					"example_id",
//				},
//				Status: pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6GatewayId4", status.Gateways[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetIpv6Gateways(ctx *pulumi.Context, args *GetIpv6GatewaysArgs, opts ...pulumi.InvokeOption) (*GetIpv6GatewaysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpv6GatewaysResult
	err := ctx.Invoke("alicloud:vpc/getIpv6Gateways:getIpv6Gateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv6Gateways.
type GetIpv6GatewaysArgs struct {
	// A list of Ipv6 Gateway IDs.
	Ids []string `pulumi:"ids"`
	// The name of the IPv6 gateway.
	Ipv6GatewayName *string `pulumi:"ipv6GatewayName"`
	// A regex string to filter results by Ipv6 Gateway name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
	Status *string `pulumi:"status"`
	// The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getIpv6Gateways.
type GetIpv6GatewaysResult struct {
	Gateways []GetIpv6GatewaysGateway `pulumi:"gateways"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	Ids             []string `pulumi:"ids"`
	Ipv6GatewayName *string  `pulumi:"ipv6GatewayName"`
	NameRegex       *string  `pulumi:"nameRegex"`
	Names           []string `pulumi:"names"`
	OutputFile      *string  `pulumi:"outputFile"`
	Status          *string  `pulumi:"status"`
	VpcId           *string  `pulumi:"vpcId"`
}

func GetIpv6GatewaysOutput(ctx *pulumi.Context, args GetIpv6GatewaysOutputArgs, opts ...pulumi.InvokeOption) GetIpv6GatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpv6GatewaysResult, error) {
			args := v.(GetIpv6GatewaysArgs)
			r, err := GetIpv6Gateways(ctx, &args, opts...)
			var s GetIpv6GatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpv6GatewaysResultOutput)
}

// A collection of arguments for invoking getIpv6Gateways.
type GetIpv6GatewaysOutputArgs struct {
	// A list of Ipv6 Gateway IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of the IPv6 gateway.
	Ipv6GatewayName pulumi.StringPtrInput `pulumi:"ipv6GatewayName"`
	// A regex string to filter results by Ipv6 Gateway name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetIpv6GatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpv6GatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getIpv6Gateways.
type GetIpv6GatewaysResultOutput struct{ *pulumi.OutputState }

func (GetIpv6GatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpv6GatewaysResult)(nil)).Elem()
}

func (o GetIpv6GatewaysResultOutput) ToGetIpv6GatewaysResultOutput() GetIpv6GatewaysResultOutput {
	return o
}

func (o GetIpv6GatewaysResultOutput) ToGetIpv6GatewaysResultOutputWithContext(ctx context.Context) GetIpv6GatewaysResultOutput {
	return o
}

func (o GetIpv6GatewaysResultOutput) Gateways() GetIpv6GatewaysGatewayArrayOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) []GetIpv6GatewaysGateway { return v.Gateways }).(GetIpv6GatewaysGatewayArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpv6GatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpv6GatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetIpv6GatewaysResultOutput) Ipv6GatewayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) *string { return v.Ipv6GatewayName }).(pulumi.StringPtrOutput)
}

func (o GetIpv6GatewaysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetIpv6GatewaysResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetIpv6GatewaysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetIpv6GatewaysResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetIpv6GatewaysResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6GatewaysResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpv6GatewaysResultOutput{})
}
