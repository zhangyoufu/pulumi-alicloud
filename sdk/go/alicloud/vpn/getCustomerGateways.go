// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VPN customers gateways data source lists a number of VPN customer gateways resource information owned by an Alicloud account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.GetCustomerGateways(ctx, &vpn.GetCustomerGatewaysArgs{
//				Ids: []string{
//					"fake-id1",
//					"fake-id2",
//				},
//				NameRegex:  pulumi.StringRef("testAcc*"),
//				OutputFile: pulumi.StringRef("/tmp/cgws"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCustomerGateways(ctx *pulumi.Context, args *GetCustomerGatewaysArgs, opts ...pulumi.InvokeOption) (*GetCustomerGatewaysResult, error) {
	var rv GetCustomerGatewaysResult
	err := ctx.Invoke("alicloud:vpn/getCustomerGateways:getCustomerGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerGateways.
type GetCustomerGatewaysArgs struct {
	// ID of the VPN customer gateways.
	Ids []string `pulumi:"ids"`
	// A regex string of VPN customer gateways name.
	NameRegex *string `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getCustomerGateways.
type GetCustomerGatewaysResult struct {
	// A list of VPN customer gateways. Each element contains the following attributes:
	Gateways []GetCustomerGatewaysGateway `pulumi:"gateways"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IDs of VPN customer gateway.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// names of VPN customer gateway.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetCustomerGatewaysOutput(ctx *pulumi.Context, args GetCustomerGatewaysOutputArgs, opts ...pulumi.InvokeOption) GetCustomerGatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomerGatewaysResult, error) {
			args := v.(GetCustomerGatewaysArgs)
			r, err := GetCustomerGateways(ctx, &args, opts...)
			var s GetCustomerGatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomerGatewaysResultOutput)
}

// A collection of arguments for invoking getCustomerGateways.
type GetCustomerGatewaysOutputArgs struct {
	// ID of the VPN customer gateways.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string of VPN customer gateways name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Save the result to the file.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetCustomerGatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerGatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getCustomerGateways.
type GetCustomerGatewaysResultOutput struct{ *pulumi.OutputState }

func (GetCustomerGatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerGatewaysResult)(nil)).Elem()
}

func (o GetCustomerGatewaysResultOutput) ToGetCustomerGatewaysResultOutput() GetCustomerGatewaysResultOutput {
	return o
}

func (o GetCustomerGatewaysResultOutput) ToGetCustomerGatewaysResultOutputWithContext(ctx context.Context) GetCustomerGatewaysResultOutput {
	return o
}

// A list of VPN customer gateways. Each element contains the following attributes:
func (o GetCustomerGatewaysResultOutput) Gateways() GetCustomerGatewaysGatewayArrayOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) []GetCustomerGatewaysGateway { return v.Gateways }).(GetCustomerGatewaysGatewayArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomerGatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

// IDs of VPN customer gateway.
func (o GetCustomerGatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCustomerGatewaysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// names of VPN customer gateway.
func (o GetCustomerGatewaysResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetCustomerGatewaysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomerGatewaysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomerGatewaysResultOutput{})
}
