// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Alb Load Balancers of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := alb.GetLoadBalancers(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albLoadBalancerId1", ids.Balancers[0].Id)
// 		opt0 := "^my-LoadBalancer"
// 		nameRegex, err := alb.GetLoadBalancers(ctx, &alb.GetLoadBalancersArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("albLoadBalancerId2", nameRegex.Balancers[0].Id)
// 		return nil
// 	})
// }
// ```
func GetLoadBalancers(ctx *pulumi.Context, args *GetLoadBalancersArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancersResult, error) {
	var rv GetLoadBalancersResult
	err := ctx.Invoke("alicloud:alb/getLoadBalancers:getLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersArgs struct {
	// The type of IP address that the ALB instance uses to provide services.
	AddressType *string `pulumi:"addressType"`
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Load Balancer IDs.
	Ids []string `pulumi:"ids"`
	// Load Balancing of the Service Status. Valid Values: `Abnormal` and `Normal`.
	LoadBalancerBussinessStatus *string `pulumi:"loadBalancerBussinessStatus"`
	// The load balancer ids.
	LoadBalancerIds []string `pulumi:"loadBalancerIds"`
	// The name of the resource.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// A regex string to filter results by Load Balancer name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The The load balancer status. Valid values: `Active`, `Configuring`, `CreateFailed`, `Inactive` and `Provisioning`.
	Status *string `pulumi:"status"`
	// The tag of the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where the ALB instance is deployed.
	VpcId *string `pulumi:"vpcId"`
	// The vpc ids.
	VpcIds []string `pulumi:"vpcIds"`
	// The ID of the zone to which the ALB instance belongs.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResult struct {
	AddressType   *string                    `pulumi:"addressType"`
	Balancers     []GetLoadBalancersBalancer `pulumi:"balancers"`
	EnableDetails *bool                      `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id                          string                 `pulumi:"id"`
	Ids                         []string               `pulumi:"ids"`
	LoadBalancerBussinessStatus *string                `pulumi:"loadBalancerBussinessStatus"`
	LoadBalancerIds             []string               `pulumi:"loadBalancerIds"`
	LoadBalancerName            *string                `pulumi:"loadBalancerName"`
	NameRegex                   *string                `pulumi:"nameRegex"`
	Names                       []string               `pulumi:"names"`
	OutputFile                  *string                `pulumi:"outputFile"`
	ResourceGroupId             *string                `pulumi:"resourceGroupId"`
	Status                      *string                `pulumi:"status"`
	Tags                        map[string]interface{} `pulumi:"tags"`
	VpcId                       *string                `pulumi:"vpcId"`
	VpcIds                      []string               `pulumi:"vpcIds"`
	ZoneId                      *string                `pulumi:"zoneId"`
}
