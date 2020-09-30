// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the Config Delivery Channels of the current Alibaba Cloud user.
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
// 		opt0 := "tftest"
// 		example, err := cfg.GetDeliveryChannels(ctx, &cfg.GetDeliveryChannelsArgs{
// 			Ids: []string{
// 				"cdc-49a2ad756057********",
// 			},
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstConfigDeliveryChannelId", example.Channels[0].Id)
// 		return nil
// 	})
// }
// ```
func GetDeliveryChannels(ctx *pulumi.Context, args *GetDeliveryChannelsArgs, opts ...pulumi.InvokeOption) (*GetDeliveryChannelsResult, error) {
	var rv GetDeliveryChannelsResult
	err := ctx.Invoke("alicloud:cfg/getDeliveryChannels:getDeliveryChannels", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeliveryChannels.
type GetDeliveryChannelsArgs struct {
	// A list of Config Delivery Channel IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by delivery channel name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the config delivery channel. Valid values `0`: Disable delivery channel, `1`: Enable delivery channel.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getDeliveryChannels.
type GetDeliveryChannelsResult struct {
	// A list of Config Delivery Channels. Each element contains the following attributes:
	Channels []GetDeliveryChannelsChannel `pulumi:"channels"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Config Delivery Channel IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Config Delivery Channel names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// The status of the delivery method.
	Status *int `pulumi:"status"`
}
