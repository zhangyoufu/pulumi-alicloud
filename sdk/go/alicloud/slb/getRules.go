// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the rules associated with a server load balancer listener.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/slb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "slbrulebasicconfig"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := "cloud_efficiency"
// 		opt1 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableDiskCategory:     &opt0,
// 			AvailableResourceCreation: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/16"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultLoadBalancer, err := slb.NewLoadBalancer(ctx, "defaultLoadBalancer", &slb.LoadBalancerArgs{
// 			VswitchId: defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultListener, err := slb.NewListener(ctx, "defaultListener", &slb.ListenerArgs{
// 			LoadBalancerId:         defaultLoadBalancer.ID(),
// 			BackendPort:            pulumi.Int(22),
// 			FrontendPort:           pulumi.Int(22),
// 			Protocol:               pulumi.String("http"),
// 			Bandwidth:              pulumi.Int(5),
// 			HealthCheckConnectPort: pulumi.Int(20),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultServerGroup, err := slb.NewServerGroup(ctx, "defaultServerGroup", &slb.ServerGroupArgs{
// 			LoadBalancerId: defaultLoadBalancer.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = slb.NewRule(ctx, "defaultRule", &slb.RuleArgs{
// 			LoadBalancerId: defaultLoadBalancer.ID(),
// 			FrontendPort:   defaultListener.FrontendPort,
// 			Domain:         pulumi.String("*.aliyun.com"),
// 			Url:            pulumi.String("/image"),
// 			ServerGroupId:  defaultServerGroup.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstSlbRuleId", sampleDs.ApplyT(func(sampleDs slb.GetRulesResult) (string, error) {
// 			return sampleDs.SlbRules[0].Id, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetRules(ctx *pulumi.Context, args *GetRulesArgs, opts ...pulumi.InvokeOption) (*GetRulesResult, error) {
	var rv GetRulesResult
	err := ctx.Invoke("alicloud:slb/getRules:getRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRules.
type GetRulesArgs struct {
	// SLB listener port.
	FrontendPort int `pulumi:"frontendPort"`
	// A list of rules IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB with listener rules.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by rule name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRules.
type GetRulesResult struct {
	FrontendPort int `pulumi:"frontendPort"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB listener rules IDs.
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	// A list of SLB listener rules names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of SLB listener rules. Each element contains the following attributes:
	SlbRules []GetRulesSlbRule `pulumi:"slbRules"`
}
