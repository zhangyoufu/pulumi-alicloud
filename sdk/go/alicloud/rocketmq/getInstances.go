// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of ONS Instances in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.52.0+
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:rocketmq/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// A list of instance IDs to filter results.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by the instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of instances. Each element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// A list of instance names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
