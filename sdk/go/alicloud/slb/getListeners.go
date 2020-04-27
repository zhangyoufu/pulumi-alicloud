// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the listeners related to a server load balancer of the current Alibaba Cloud user.
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
