// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the server load balancers of the current Alibaba Cloud user.
//
// > **NOTE:** Available in 1.123.1+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := slb.GetApplicationLoadBalancers(ctx, &slb.GetApplicationLoadBalancersArgs{
//				NameRegex: pulumi.StringRef("sample_slb"),
//				Tags: map[string]interface{}{
//					"tagKey1": "tagValue1",
//					"tagKey2": "tagValue2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstSlbId", example.Balancers[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetApplicationLoadBalancers(ctx *pulumi.Context, args *GetApplicationLoadBalancersArgs, opts ...pulumi.InvokeOption) (*GetApplicationLoadBalancersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationLoadBalancersResult
	err := ctx.Invoke("alicloud:slb/getApplicationLoadBalancers:getApplicationLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationLoadBalancers.
type GetApplicationLoadBalancersArgs struct {
	// Service address of the SLBs.
	Address *string `pulumi:"address"`
	// The address ip version. Valid values `ipv4` and `ipv6`.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The address type of the SLB. Valid values `internet` and `intranet`.
	AddressType   *string `pulumi:"addressType"`
	EnableDetails *bool   `pulumi:"enableDetails"`
	// A list of SLBs IDs.
	Ids []string `pulumi:"ids"`
	// The internet charge type. Valid values `PayByBandwidth` and `PayByTraffic`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The name of the SLB.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The master zone id of the SLB.
	MasterZoneId *string `pulumi:"masterZoneId"`
	// A regex string to filter results by SLB name.
	NameRegex *string `pulumi:"nameRegex"`
	// Network type of the SLBs. Valid values: `vpc` and `classic`.
	NetworkType *string `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The payment type of SLB. Valid values `PayAsYouGo` and `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The Id of resource group which SLB belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The server ID.
	ServerId *string `pulumi:"serverId"`
	// The server intranet address.
	ServerIntranetAddress *string `pulumi:"serverIntranetAddress"`
	// The slave zone id of the SLB.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
	Tags map[string]string `pulumi:"tags"`
	// ID of the VPC linked to the SLBs.
	VpcId *string `pulumi:"vpcId"`
	// ID of the vSwitch linked to the SLBs.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getApplicationLoadBalancers.
type GetApplicationLoadBalancersResult struct {
	// The IP address that the SLB instance uses to provide services.
	Address *string `pulumi:"address"`
	// The address ip version.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The address type.
	AddressType *string `pulumi:"addressType"`
	// A list of SLBs. Each element contains the following attributes:
	Balancers     []GetApplicationLoadBalancersBalancer `pulumi:"balancers"`
	EnableDetails *bool                                 `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of slb IDs.
	Ids []string `pulumi:"ids"`
	// The billing method of the Internet-facing SLB instance.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The name of the SLB.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Master availability zone of the SLBs.
	MasterZoneId *string `pulumi:"masterZoneId"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of slb names.
	Names []string `pulumi:"names"`
	// Network type of the SLB. Possible values: `vpc` and `classic`.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	PageNumber  *int    `pulumi:"pageNumber"`
	PageSize    *int    `pulumi:"pageSize"`
	PaymentType *string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the Elastic Compute Service (ECS) instance that is specified as a backend server of the CLB instance.
	ServerId              *string `pulumi:"serverId"`
	ServerIntranetAddress *string `pulumi:"serverIntranetAddress"`
	// Slave availability zone of the SLBs.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// Deprecated: Field 'slbs' has deprecated from v1.123.1 and replace by 'balancers'.
	Slbs []GetApplicationLoadBalancersSlb `pulumi:"slbs"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status *string `pulumi:"status"`
	// The tags of the SLB.
	Tags       map[string]string `pulumi:"tags"`
	TotalCount int               `pulumi:"totalCount"`
	// ID of the VPC the SLB belongs to.
	VpcId *string `pulumi:"vpcId"`
	// ID of the vSwitch the SLB belongs to.
	VswitchId *string `pulumi:"vswitchId"`
}

func GetApplicationLoadBalancersOutput(ctx *pulumi.Context, args GetApplicationLoadBalancersOutputArgs, opts ...pulumi.InvokeOption) GetApplicationLoadBalancersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationLoadBalancersResult, error) {
			args := v.(GetApplicationLoadBalancersArgs)
			r, err := GetApplicationLoadBalancers(ctx, &args, opts...)
			var s GetApplicationLoadBalancersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationLoadBalancersResultOutput)
}

// A collection of arguments for invoking getApplicationLoadBalancers.
type GetApplicationLoadBalancersOutputArgs struct {
	// Service address of the SLBs.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The address ip version. Valid values `ipv4` and `ipv6`.
	AddressIpVersion pulumi.StringPtrInput `pulumi:"addressIpVersion"`
	// The address type of the SLB. Valid values `internet` and `intranet`.
	AddressType   pulumi.StringPtrInput `pulumi:"addressType"`
	EnableDetails pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// A list of SLBs IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The internet charge type. Valid values `PayByBandwidth` and `PayByTraffic`.
	InternetChargeType pulumi.StringPtrInput `pulumi:"internetChargeType"`
	// The name of the SLB.
	LoadBalancerName pulumi.StringPtrInput `pulumi:"loadBalancerName"`
	// The master zone id of the SLB.
	MasterZoneId pulumi.StringPtrInput `pulumi:"masterZoneId"`
	// A regex string to filter results by SLB name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Network type of the SLBs. Valid values: `vpc` and `classic`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The payment type of SLB. Valid values `PayAsYouGo` and `Subscription`.
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// The Id of resource group which SLB belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The server ID.
	ServerId pulumi.StringPtrInput `pulumi:"serverId"`
	// The server intranet address.
	ServerIntranetAddress pulumi.StringPtrInput `pulumi:"serverIntranetAddress"`
	// The slave zone id of the SLB.
	SlaveZoneId pulumi.StringPtrInput `pulumi:"slaveZoneId"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// ID of the VPC linked to the SLBs.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// ID of the vSwitch linked to the SLBs.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetApplicationLoadBalancersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationLoadBalancersArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationLoadBalancers.
type GetApplicationLoadBalancersResultOutput struct{ *pulumi.OutputState }

func (GetApplicationLoadBalancersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationLoadBalancersResult)(nil)).Elem()
}

func (o GetApplicationLoadBalancersResultOutput) ToGetApplicationLoadBalancersResultOutput() GetApplicationLoadBalancersResultOutput {
	return o
}

func (o GetApplicationLoadBalancersResultOutput) ToGetApplicationLoadBalancersResultOutputWithContext(ctx context.Context) GetApplicationLoadBalancersResultOutput {
	return o
}

// The IP address that the SLB instance uses to provide services.
func (o GetApplicationLoadBalancersResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The address ip version.
func (o GetApplicationLoadBalancersResultOutput) AddressIpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.AddressIpVersion }).(pulumi.StringPtrOutput)
}

// The address type.
func (o GetApplicationLoadBalancersResultOutput) AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.AddressType }).(pulumi.StringPtrOutput)
}

// A list of SLBs. Each element contains the following attributes:
func (o GetApplicationLoadBalancersResultOutput) Balancers() GetApplicationLoadBalancersBalancerArrayOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) []GetApplicationLoadBalancersBalancer { return v.Balancers }).(GetApplicationLoadBalancersBalancerArrayOutput)
}

func (o GetApplicationLoadBalancersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationLoadBalancersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of slb IDs.
func (o GetApplicationLoadBalancersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The billing method of the Internet-facing SLB instance.
func (o GetApplicationLoadBalancersResultOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

// The name of the SLB.
func (o GetApplicationLoadBalancersResultOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

// Master availability zone of the SLBs.
func (o GetApplicationLoadBalancersResultOutput) MasterZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.MasterZoneId }).(pulumi.StringPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of slb names.
func (o GetApplicationLoadBalancersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Network type of the SLB. Possible values: `vpc` and `classic`.
func (o GetApplicationLoadBalancersResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

// The ID of the resource group.
func (o GetApplicationLoadBalancersResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The ID of the Elastic Compute Service (ECS) instance that is specified as a backend server of the CLB instance.
func (o GetApplicationLoadBalancersResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

func (o GetApplicationLoadBalancersResultOutput) ServerIntranetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.ServerIntranetAddress }).(pulumi.StringPtrOutput)
}

// Slave availability zone of the SLBs.
func (o GetApplicationLoadBalancersResultOutput) SlaveZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.SlaveZoneId }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'slbs' has deprecated from v1.123.1 and replace by 'balancers'.
func (o GetApplicationLoadBalancersResultOutput) Slbs() GetApplicationLoadBalancersSlbArrayOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) []GetApplicationLoadBalancersSlb { return v.Slbs }).(GetApplicationLoadBalancersSlbArrayOutput)
}

// SLB current status. Possible values: `inactive`, `active` and `locked`.
func (o GetApplicationLoadBalancersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the SLB.
func (o GetApplicationLoadBalancersResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetApplicationLoadBalancersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// ID of the VPC the SLB belongs to.
func (o GetApplicationLoadBalancersResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// ID of the vSwitch the SLB belongs to.
func (o GetApplicationLoadBalancersResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationLoadBalancersResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationLoadBalancersResultOutput{})
}
