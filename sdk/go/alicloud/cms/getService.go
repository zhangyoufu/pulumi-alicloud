// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Using this data source can open CMS service automatically. If the service has been opened, it will return opened.
//
// For information about CMS and how to use it, see [What is CMS](https://help.aliyun.com/product/28572.html).
//
// > **NOTE:** Available in v1.111.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/cms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "On"
// 		_, err := cms.GetService(ctx, &cms.GetServiceArgs{
// 			Enable: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("alicloud:cms/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}
