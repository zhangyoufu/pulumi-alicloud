// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nlb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Nlb Load Balancers of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.191.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nlb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := nlb.GetLoadBalancers(ctx, &nlb.GetLoadBalancersArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nlbLoadBalancerId1", ids.Balancers[0].Id)
//			nameRegex, err := nlb.GetLoadBalancers(ctx, &nlb.GetLoadBalancersArgs{
//				NameRegex: pulumi.StringRef("^my-LoadBalancer"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nlbLoadBalancerId2", nameRegex.Balancers[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetLoadBalancers(ctx *pulumi.Context, args *GetLoadBalancersArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLoadBalancersResult
	err := ctx.Invoke("alicloud:nlb/getLoadBalancers:getLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersArgs struct {
	// The IP version.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The type of IPv4 address used by the NLB instance.
	AddressType *string `pulumi:"addressType"`
	// The domain name of the NLB instance.
	DnsName *string `pulumi:"dnsName"`
	// A list of Load Balancer IDs.
	Ids []string `pulumi:"ids"`
	// The type of IPv6 address used by the NLB instance.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus *string `pulumi:"loadBalancerBusinessStatus"`
	// The name of the NLB instance. You can specify at most 10 names.
	LoadBalancerNames []string `pulumi:"loadBalancerNames"`
	// A regex string to filter results by Load Balancer name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the NLB instance.
	Status *string `pulumi:"status"`
	// The tag of the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where the NLB instance is deployed. You can specify at most 10 IDs.
	VpcIds []string `pulumi:"vpcIds"`
	// The name of the zone.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResult struct {
	AddressIpVersion *string                    `pulumi:"addressIpVersion"`
	AddressType      *string                    `pulumi:"addressType"`
	Balancers        []GetLoadBalancersBalancer `pulumi:"balancers"`
	DnsName          *string                    `pulumi:"dnsName"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string                 `pulumi:"id"`
	Ids                        []string               `pulumi:"ids"`
	Ipv6AddressType            *string                `pulumi:"ipv6AddressType"`
	LoadBalancerBusinessStatus *string                `pulumi:"loadBalancerBusinessStatus"`
	LoadBalancerNames          []string               `pulumi:"loadBalancerNames"`
	NameRegex                  *string                `pulumi:"nameRegex"`
	Names                      []string               `pulumi:"names"`
	OutputFile                 *string                `pulumi:"outputFile"`
	ResourceGroupId            *string                `pulumi:"resourceGroupId"`
	Status                     *string                `pulumi:"status"`
	Tags                       map[string]interface{} `pulumi:"tags"`
	VpcIds                     []string               `pulumi:"vpcIds"`
	ZoneId                     *string                `pulumi:"zoneId"`
}

func GetLoadBalancersOutput(ctx *pulumi.Context, args GetLoadBalancersOutputArgs, opts ...pulumi.InvokeOption) GetLoadBalancersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLoadBalancersResult, error) {
			args := v.(GetLoadBalancersArgs)
			r, err := GetLoadBalancers(ctx, &args, opts...)
			var s GetLoadBalancersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLoadBalancersResultOutput)
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersOutputArgs struct {
	// The IP version.
	AddressIpVersion pulumi.StringPtrInput `pulumi:"addressIpVersion"`
	// The type of IPv4 address used by the NLB instance.
	AddressType pulumi.StringPtrInput `pulumi:"addressType"`
	// The domain name of the NLB instance.
	DnsName pulumi.StringPtrInput `pulumi:"dnsName"`
	// A list of Load Balancer IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The type of IPv6 address used by the NLB instance.
	Ipv6AddressType pulumi.StringPtrInput `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus pulumi.StringPtrInput `pulumi:"loadBalancerBusinessStatus"`
	// The name of the NLB instance. You can specify at most 10 names.
	LoadBalancerNames pulumi.StringArrayInput `pulumi:"loadBalancerNames"`
	// A regex string to filter results by Load Balancer name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status of the NLB instance.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The tag of the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the virtual private cloud (VPC) where the NLB instance is deployed. You can specify at most 10 IDs.
	VpcIds pulumi.StringArrayInput `pulumi:"vpcIds"`
	// The name of the zone.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetLoadBalancersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancersArgs)(nil)).Elem()
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResultOutput struct{ *pulumi.OutputState }

func (GetLoadBalancersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLoadBalancersResult)(nil)).Elem()
}

func (o GetLoadBalancersResultOutput) ToGetLoadBalancersResultOutput() GetLoadBalancersResultOutput {
	return o
}

func (o GetLoadBalancersResultOutput) ToGetLoadBalancersResultOutputWithContext(ctx context.Context) GetLoadBalancersResultOutput {
	return o
}

func (o GetLoadBalancersResultOutput) AddressIpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.AddressIpVersion }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) Balancers() GetLoadBalancersBalancerArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []GetLoadBalancersBalancer { return v.Balancers }).(GetLoadBalancersBalancerArrayOutput)
}

func (o GetLoadBalancersResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadBalancersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLoadBalancersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancersResultOutput) Ipv6AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.Ipv6AddressType }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) LoadBalancerBusinessStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.LoadBalancerBusinessStatus }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) LoadBalancerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.LoadBalancerNames }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetLoadBalancersResultOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancersResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadBalancersResultOutput{})
}
