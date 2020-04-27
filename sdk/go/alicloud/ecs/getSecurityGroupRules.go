// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `ecs.getSecurityGroupRules` data source provides a collection of security permissions of a specific security group.
// Each collection item represents a single `ingress` or `egress` permission rule.
// The ID of the security group can be provided via a variable or the result from the other data source `ecs.getSecurityGroups`.
func GetSecurityGroupRules(ctx *pulumi.Context, args *GetSecurityGroupRulesArgs, opts ...pulumi.InvokeOption) (*GetSecurityGroupRulesResult, error) {
	var rv GetSecurityGroupRulesResult
	err := ctx.Invoke("alicloud:ecs/getSecurityGroupRules:getSecurityGroupRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroupRules.
type GetSecurityGroupRulesArgs struct {
	// Authorization direction. Valid values are: `ingress` or `egress`.
	Direction *string `pulumi:"direction"`
	// The ID of the security group that owns the rules.
	GroupId string `pulumi:"groupId"`
	// The IP protocol. Valid values are: `tcp`, `udp`, `icmp`, `gre` and `all`.
	IpProtocol *string `pulumi:"ipProtocol"`
	// Refers to the network type. Can be either `internet` or `intranet`. The default value is `internet`.
	NicType    *string `pulumi:"nicType"`
	OutputFile *string `pulumi:"outputFile"`
	// Authorization policy. Can be either `accept` or `drop`. The default value is `accept`.
	Policy *string `pulumi:"policy"`
}

// A collection of values returned by getSecurityGroupRules.
type GetSecurityGroupRulesResult struct {
	// Authorization direction, `ingress` or `egress`.
	Direction *string `pulumi:"direction"`
	// The description of the security group that owns the rules.
	GroupDesc string `pulumi:"groupDesc"`
	GroupId   string `pulumi:"groupId"`
	// The name of the security group that owns the rules.
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The protocol. Can be `tcp`, `udp`, `icmp`, `gre` or `all`.
	IpProtocol *string `pulumi:"ipProtocol"`
	// Network type, `internet` or `intranet`.
	NicType    *string `pulumi:"nicType"`
	OutputFile *string `pulumi:"outputFile"`
	// Authorization policy. Can be either `accept` or `drop`.
	Policy *string `pulumi:"policy"`
	// A list of security group rules. Each element contains the following attributes:
	Rules []GetSecurityGroupRulesRule `pulumi:"rules"`
}
