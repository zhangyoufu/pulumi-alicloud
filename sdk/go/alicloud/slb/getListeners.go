// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the listeners related to a server load balancer of the current Alibaba Cloud user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := slb.NewApplicationLoadBalancer(ctx, "_default", &slb.ApplicationLoadBalancerArgs{
// 			LoadBalancerName: pulumi.String("tf-testAccSlbListenertcp"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = slb.NewListener(ctx, "tcp", &slb.ListenerArgs{
// 			LoadBalancerId:         _default.ID(),
// 			BackendPort:            pulumi.Int(22),
// 			FrontendPort:           pulumi.Int(22),
// 			Protocol:               pulumi.String("tcp"),
// 			Bandwidth:              pulumi.Int(10),
// 			HealthCheckType:        pulumi.String("tcp"),
// 			PersistenceTimeout:     pulumi.Int(3600),
// 			HealthyThreshold:       pulumi.Int(8),
// 			UnhealthyThreshold:     pulumi.Int(8),
// 			HealthCheckTimeout:     pulumi.Int(8),
// 			HealthCheckInterval:    pulumi.Int(5),
// 			HealthCheckHttpCode:    pulumi.String("http_2xx"),
// 			HealthCheckConnectPort: pulumi.Int(20),
// 			HealthCheckUri:         pulumi.String("/console"),
// 			EstablishedTimeout:     pulumi.Int(600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstSlbListenerProtocol", sampleDs.ApplyT(func(sampleDs slb.GetListenersResult) (string, error) {
// 			return sampleDs.SlbListeners[0].Protocol, nil
// 		}).(pulumi.StringOutput))
// 		return nil
// 	})
// }
// ```
func GetListeners(ctx *pulumi.Context, args *GetListenersArgs, opts ...pulumi.InvokeOption) (*GetListenersResult, error) {
	var rv GetListenersResult
	err := ctx.Invoke("alicloud:slb/getListeners:getListeners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListeners.
type GetListenersArgs struct {
	// A regex string to filter results by SLB listener description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// Filter listeners by the specified frontend port.
	FrontendPort *int `pulumi:"frontendPort"`
	// ID of the SLB with listeners.
	LoadBalancerId string  `pulumi:"loadBalancerId"`
	OutputFile     *string `pulumi:"outputFile"`
	// Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
	Protocol *string `pulumi:"protocol"`
}

// A collection of values returned by getListeners.
type GetListenersResult struct {
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// Frontend port used to receive incoming traffic and distribute it to the backend servers.
	FrontendPort *int `pulumi:"frontendPort"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LoadBalancerId string  `pulumi:"loadBalancerId"`
	OutputFile     *string `pulumi:"outputFile"`
	// Listener protocol. Possible values: `http`, `https`, `tcp` and `udp`.
	Protocol *string `pulumi:"protocol"`
	// A list of SLB listeners. Each element contains the following attributes:
	SlbListeners []GetListenersSlbListener `pulumi:"slbListeners"`
}
