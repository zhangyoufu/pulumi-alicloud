// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of EIPs (Elastic IP address) owned by an Alibaba Cloud account.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/eips.html.markdown.
func GetEips(ctx *pulumi.Context, args *GetEipsArgs, opts ...pulumi.InvokeOption) (*GetEipsResult, error) {
	var rv GetEipsResult
	err := ctx.Invoke("alicloud:ecs/getEips:getEips", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEips.
type GetEipsArgs struct {
	// A list of EIP IDs.
	Ids []string `pulumi:"ids"`
	// Deprecated since the version 1.8.0 of this provider.
	InUse *bool `pulumi:"inUse"`
	// A list of EIP public IP addresses.
	IpAddresses []string `pulumi:"ipAddresses"`
	OutputFile  *string  `pulumi:"outputFile"`
	// The Id of resource group which the eips belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getEips.
type GetEipsResult struct {
	// A list of EIPs. Each element contains the following attributes:
	Eips []GetEipsEip `pulumi:"eips"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of EIP IDs.
	Ids         []string `pulumi:"ids"`
	InUse       *bool    `pulumi:"inUse"`
	IpAddresses []string `pulumi:"ipAddresses"`
	// (Optional) A list of EIP names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The Id of resource group which the eips belongs.
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
}
