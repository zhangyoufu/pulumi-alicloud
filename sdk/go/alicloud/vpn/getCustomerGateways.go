// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The VPN customers gateways data source lists a number of VPN customer gateways resource information owned by an Alicloud account.
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
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
