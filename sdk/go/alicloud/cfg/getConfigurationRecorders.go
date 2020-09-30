// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Config Configuration Recorders of the current Alibaba Cloud user.
//
// > **NOTE:**  Available in 1.99.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/cfg"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cfg.GetConfigurationRecorders(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("listOfResourceTypes", data.Alicloud_config_configuration_recorders.This.Recorders[0].Resource_types)
// 		return nil
// 	})
// }
// ```
func GetConfigurationRecorders(ctx *pulumi.Context, args *GetConfigurationRecordersArgs, opts ...pulumi.InvokeOption) (*GetConfigurationRecordersResult, error) {
	var rv GetConfigurationRecordersResult
	err := ctx.Invoke("alicloud:cfg/getConfigurationRecorders:getConfigurationRecorders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfigurationRecorders.
type GetConfigurationRecordersArgs struct {
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getConfigurationRecorders.
type GetConfigurationRecordersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// A list of Config Configuration Recorders. Each element contains the following attributes:
	Recorders []GetConfigurationRecordersRecorder `pulumi:"recorders"`
}
