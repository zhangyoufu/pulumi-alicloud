// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides CEN Bandwidth Limits available to the user.
func GetBandwidthLimits(ctx *pulumi.Context, args *GetBandwidthLimitsArgs, opts ...pulumi.InvokeOption) (*GetBandwidthLimitsResult, error) {
	var rv GetBandwidthLimitsResult
	err := ctx.Invoke("alicloud:cen/getBandwidthLimits:getBandwidthLimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBandwidthLimits.
type GetBandwidthLimitsArgs struct {
	// A list of CEN instances IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	OutputFile  *string  `pulumi:"outputFile"`
}

// A collection of values returned by getBandwidthLimits.
type GetBandwidthLimitsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	InstanceIds []string `pulumi:"instanceIds"`
	// A list of CEN Bandwidth Limits. Each element contains the following attributes:
	Limits     []GetBandwidthLimitsLimit `pulumi:"limits"`
	OutputFile *string                   `pulumi:"outputFile"`
}
