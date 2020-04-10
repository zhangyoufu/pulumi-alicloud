// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides AccessRule available to the user.
//
// > NOTE: Available in 1.35.0+
func GetAccessRules(ctx *pulumi.Context, args *GetAccessRulesArgs, opts ...pulumi.InvokeOption) (*GetAccessRulesResult, error) {
	var rv GetAccessRulesResult
	err := ctx.Invoke("alicloud:nas/getAccessRules:getAccessRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessRules.
type GetAccessRulesArgs struct {
	// Filter results by a specific AccessGroupName.
	AccessGroupName string `pulumi:"accessGroupName"`
	// A list of rule IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// Filter results by a specific RWAccess.
	RwAccess *string `pulumi:"rwAccess"`
	// Filter results by a specific SourceCidrIp.
	SourceCidrIp *string `pulumi:"sourceCidrIp"`
	// Filter results by a specific UserAccess.
	UserAccess *string `pulumi:"userAccess"`
}

// A collection of values returned by getAccessRules.
type GetAccessRulesResult struct {
	AccessGroupName string `pulumi:"accessGroupName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of rule IDs, Each element set to `accessRuleId` (Each element formats as `<access_group_name>:<access rule id>` before 1.53.0).
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of AccessRules. Each element contains the following attributes:
	Rules []GetAccessRulesRule `pulumi:"rules"`
	// RWAccess of the AccessRule.
	RwAccess *string `pulumi:"rwAccess"`
	// SourceCidrIp of the AccessRule.
	SourceCidrIp *string `pulumi:"sourceCidrIp"`
	// UserAccess of the AccessRule
	UserAccess *string `pulumi:"userAccess"`
}
