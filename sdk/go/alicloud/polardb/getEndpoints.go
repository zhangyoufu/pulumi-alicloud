// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `polardb.getEndpoints` data source provides a collection of PolarDB endpoints available in Alibaba Cloud account.
// Filters support regular expression for the cluster name, searches by clusterId, and other filters which are listed below.
//
// > **NOTE:** Available in v1.68.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/polardb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "pc-\\w+"
// 		opt1 := "Running"
// 		polardbClustersDs, err := polardb.GetClusters(ctx, &polardb.GetClustersArgs{
// 			DescriptionRegex: &opt0,
// 			Status:           &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_default, err := polardb.GetEndpoints(ctx, &polardb.GetEndpointsArgs{
// 			DbClusterId: polardbClustersDs.Clusters[0].Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("ends", _default.Endpoints[0].DbEndpointId)
// 		return nil
// 	})
// }
// ```
func GetEndpoints(ctx *pulumi.Context, args *GetEndpointsArgs, opts ...pulumi.InvokeOption) (*GetEndpointsResult, error) {
	var rv GetEndpointsResult
	err := ctx.Invoke("alicloud:polardb/getEndpoints:getEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpoints.
type GetEndpointsArgs struct {
	// PolarDB cluster ID.
	DbClusterId string `pulumi:"dbClusterId"`
	// endpoint of the cluster.
	DbEndpointId *string `pulumi:"dbEndpointId"`
}

// A collection of values returned by getEndpoints.
type GetEndpointsResult struct {
	DbClusterId string `pulumi:"dbClusterId"`
	// The endpoint ID.
	DbEndpointId *string `pulumi:"dbEndpointId"`
	// A list of PolarDB cluster endpoints. Each element contains the following attributes:
	Endpoints []GetEndpointsEndpoint `pulumi:"endpoints"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
