// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Vpc Nat Ip Cidrs of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.136.0+.
func GetNatIpCidrs(ctx *pulumi.Context, args *GetNatIpCidrsArgs, opts ...pulumi.InvokeOption) (*GetNatIpCidrsResult, error) {
	var rv GetNatIpCidrsResult
	err := ctx.Invoke("alicloud:vpc/getNatIpCidrs:getNatIpCidrs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatIpCidrs.
type GetNatIpCidrsArgs struct {
	// A list of Nat Ip Cidr IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Nat Ip Cidr name.
	NameRegex *string `pulumi:"nameRegex"`
	// The ID of the VPC NAT gateway.
	NatGatewayId string `pulumi:"natGatewayId"`
	// NAT IP ADDRESS the name of the root directory. Length is from `2` to `128` characters, must start with a letter or the Chinese at the beginning can contain numbers, half a period (.), underscore (_) and dash (-). But do not start with `http://` or `https://` at the beginning.
	NatIpCidrNames []string `pulumi:"natIpCidrNames"`
	// The NAT CIDR block to be created. Support up to `20`. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
	NatIpCidrs []string `pulumi:"natIpCidrs"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the CIDR block of the NAT gateway. If the value is `Available`, the CIDR block is available.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getNatIpCidrs.
type GetNatIpCidrsResult struct {
	Cidrs []GetNatIpCidrsCidr `pulumi:"cidrs"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	Ids            []string `pulumi:"ids"`
	NameRegex      *string  `pulumi:"nameRegex"`
	Names          []string `pulumi:"names"`
	NatGatewayId   string   `pulumi:"natGatewayId"`
	NatIpCidrNames []string `pulumi:"natIpCidrNames"`
	NatIpCidrs     []string `pulumi:"natIpCidrs"`
	OutputFile     *string  `pulumi:"outputFile"`
	Status         *string  `pulumi:"status"`
}
