// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the ECS instance type families of Alibaba Cloud.
//
// > **NOTE:** Available in 1.54.0+
func GetInstanceTypeFamilies(ctx *pulumi.Context, args *GetInstanceTypeFamiliesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeFamiliesResult, error) {
	var rv GetInstanceTypeFamiliesResult
	err := ctx.Invoke("alicloud:ecs/getInstanceTypeFamilies:getInstanceTypeFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeFamilies.
type GetInstanceTypeFamiliesArgs struct {
	// The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3` and `ecs-4`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm).
	Generation *string `pulumi:"generation"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	OutputFile         *string `pulumi:"outputFile"`
	// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// The Zone to launch the instance.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getInstanceTypeFamilies.
type GetInstanceTypeFamiliesResult struct {
	Families []GetInstanceTypeFamiliesFamily `pulumi:"families"`
	// The generation of the instance type family.
	Generation *string `pulumi:"generation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance type family IDs.
	Ids                []string `pulumi:"ids"`
	InstanceChargeType *string  `pulumi:"instanceChargeType"`
	OutputFile         *string  `pulumi:"outputFile"`
	SpotStrategy       *string  `pulumi:"spotStrategy"`
	ZoneId             *string  `pulumi:"zoneId"`
}
