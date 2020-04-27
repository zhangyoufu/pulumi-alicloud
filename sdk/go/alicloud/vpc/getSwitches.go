// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of VSwitches owned by an Alibaba Cloud account.
func GetSwitches(ctx *pulumi.Context, args *GetSwitchesArgs, opts ...pulumi.InvokeOption) (*GetSwitchesResult, error) {
	var rv GetSwitchesResult
	err := ctx.Invoke("alicloud:vpc/getSwitches:getSwitches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSwitches.
type GetSwitchesArgs struct {
	// Filter results by a specific CIDR block. For example: "172.16.0.0/12".
	CidrBlock *string `pulumi:"cidrBlock"`
	// A list of VSwitch IDs.
	Ids []string `pulumi:"ids"`
	// Indicate whether the VSwitch is created by the system.
	IsDefault *bool `pulumi:"isDefault"`
	// A regex string to filter results by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The Id of resource group which VSWitch belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC that owns the VSwitch.
	VpcId *string `pulumi:"vpcId"`
	// The availability zone of the VSwitch.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getSwitches.
type GetSwitchesResult struct {
	// CIDR block of the VSwitch.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of VSwitch IDs.
	Ids []string `pulumi:"ids"`
	// Whether the VSwitch is the default one in the region.
	IsDefault *bool   `pulumi:"isDefault"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of VSwitch names.
	Names           []string               `pulumi:"names"`
	OutputFile      *string                `pulumi:"outputFile"`
	ResourceGroupId *string                `pulumi:"resourceGroupId"`
	Tags            map[string]interface{} `pulumi:"tags"`
	// ID of the VPC that owns the VSwitch.
	VpcId *string `pulumi:"vpcId"`
	// A list of VSwitches. Each element contains the following attributes:
	Vswitches []GetSwitchesVswitch `pulumi:"vswitches"`
	// ID of the availability zone where the VSwitch is located.
	ZoneId *string `pulumi:"zoneId"`
}
