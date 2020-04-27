// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides available scaling rule resources.
func GetScalingRules(ctx *pulumi.Context, args *GetScalingRulesArgs, opts ...pulumi.InvokeOption) (*GetScalingRulesResult, error) {
	var rv GetScalingRulesResult
	err := ctx.Invoke("alicloud:ess/getScalingRules:getScalingRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingRules.
type GetScalingRulesArgs struct {
	// A list of scaling rule IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting scaling rules by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Scaling group id the scaling rules belong to.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Type of scaling rule.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getScalingRules.
type GetScalingRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of scaling rule ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of scaling rule names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of scaling rules. Each element contains the following attributes:
	Rules []GetScalingRulesRule `pulumi:"rules"`
	// ID of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Type of the scaling rule.
	Type *string `pulumi:"type"`
}
