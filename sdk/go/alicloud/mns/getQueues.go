// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mns

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of MNS queues in an Alibaba Cloud account according to the specified parameters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/mns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "tf-"
// 		queues, err := mns.GetQueues(ctx, &mns.GetQueuesArgs{
// 			NamePrefix: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstQueueId", queues.Queues[0].Id)
// 		return nil
// 	})
// }
// ```
func GetQueues(ctx *pulumi.Context, args *GetQueuesArgs, opts ...pulumi.InvokeOption) (*GetQueuesResult, error) {
	var rv GetQueuesResult
	err := ctx.Invoke("alicloud:mns/getQueues:getQueues", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueues.
type GetQueuesArgs struct {
	// A string to filter resulting queues by their name prefixs.
	NamePrefix *string `pulumi:"namePrefix"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getQueues.
type GetQueuesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of queue names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of queues. Each element contains the following attributes:
	Queues []GetQueuesQueue `pulumi:"queues"`
}
