// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Message Center Subscriptions of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.135.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := alicloud.GetMscSubSubscriptions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("mscSubSubscriptionId1", _default.Subscriptions[0].Id)
// 		return nil
// 	})
// }
// ```
func GetMscSubSubscriptions(ctx *pulumi.Context, args *GetMscSubSubscriptionsArgs, opts ...pulumi.InvokeOption) (*GetMscSubSubscriptionsResult, error) {
	var rv GetMscSubSubscriptionsResult
	err := ctx.Invoke("alicloud:index/getMscSubSubscriptions:getMscSubSubscriptions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMscSubSubscriptions.
type GetMscSubSubscriptionsArgs struct {
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getMscSubSubscriptions.
type GetMscSubSubscriptionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                               `pulumi:"id"`
	OutputFile    *string                              `pulumi:"outputFile"`
	Subscriptions []GetMscSubSubscriptionsSubscription `pulumi:"subscriptions"`
}
