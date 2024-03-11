// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:** This datasource has been renamed to slb.getApplicationLoadBalancers from version 1.123.1.
//
// This data source provides the server load balancers of the current Alibaba Cloud user.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := slb.NewLoadBalancer(ctx, "default", nil)
//			if err != nil {
//				return err
//			}
//			slbsDs, err := slb.GetLoadBalancers(ctx, &slb.GetLoadBalancersArgs{
//				NameRegex: pulumi.StringRef("sample_slb"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstSlbId", slbsDs.Slbs[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetLoadBalancers(ctx *pulumi.Context, args *GetLoadBalancersArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLoadBalancersResult
	err := ctx.Invoke("alicloud:slb/getLoadBalancers:getLoadBalancers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancers.
type GetLoadBalancersArgs struct {
	// Service address of the SLBs.
	Address          *string `pulumi:"address"`
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	AddressType      *string `pulumi:"addressType"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	// A list of SLBs IDs.
	Ids                []string `pulumi:"ids"`
	InternetChargeType *string  `pulumi:"internetChargeType"`
	LoadBalancerName   *string  `pulumi:"loadBalancerName"`
	MasterZoneId       *string  `pulumi:"masterZoneId"`
	// A regex string to filter results by SLB name.
	NameRegex *string `pulumi:"nameRegex"`
	// Network type of the SLBs. Valid values: `vpc` and `classic`.
	NetworkType *string `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  *string `pulumi:"outputFile"`
	PageNumber  *int    `pulumi:"pageNumber"`
	PageSize    *int    `pulumi:"pageSize"`
	PaymentType *string `pulumi:"paymentType"`
	// The Id of resource group which SLB belongs.
	ResourceGroupId       *string `pulumi:"resourceGroupId"`
	ServerId              *string `pulumi:"serverId"`
	ServerIntranetAddress *string `pulumi:"serverIntranetAddress"`
	SlaveZoneId           *string `pulumi:"slaveZoneId"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
	// <!--Start PulumiCodeChooser -->
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
	// 		_, err := slb.GetLoadBalancers(ctx, &slb.GetLoadBalancersArgs{
	// 			Tags: map[string]interface{}{
	// 				"tagKey1": "tagValue1",
	// 				"tagKey2": "tagValue2",
	// 			},
	// 		}, nil)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	})
	// }
	// ```
	// <!--End PulumiCodeChooser -->
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC linked to the SLBs.
	VpcId *string `pulumi:"vpcId"`
	// ID of the VSwitch linked to the SLBs.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getLoadBalancers.
type GetLoadBalancersResult struct {
	// Service address of the SLB.
	Address          *string                    `pulumi:"address"`
	AddressIpVersion *string                    `pulumi:"addressIpVersion"`
	AddressType      *string                    `pulumi:"addressType"`
	Balancers        []GetLoadBalancersBalancer `pulumi:"balancers"`
	EnableDetails    *bool                      `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of slb IDs.
	Ids                []string `pulumi:"ids"`
	InternetChargeType *string  `pulumi:"internetChargeType"`
	LoadBalancerName   *string  `pulumi:"loadBalancerName"`
	MasterZoneId       *string  `pulumi:"masterZoneId"`
	NameRegex          *string  `pulumi:"nameRegex"`
	// A list of slb names.
	Names []string `pulumi:"names"`
	// Network type of the SLB. Possible values: `vpc` and `classic`.
	NetworkType           *string `pulumi:"networkType"`
	OutputFile            *string `pulumi:"outputFile"`
	PageNumber            *int    `pulumi:"pageNumber"`
	PageSize              *int    `pulumi:"pageSize"`
	PaymentType           *string `pulumi:"paymentType"`
	ResourceGroupId       *string `pulumi:"resourceGroupId"`
	ServerId              *string `pulumi:"serverId"`
	ServerIntranetAddress *string `pulumi:"serverIntranetAddress"`
	SlaveZoneId           *string `pulumi:"slaveZoneId"`
	// A list of SLBs. Each element contains the following attributes:
	//
	// Deprecated: Field 'slbs' has deprecated from v1.123.1 and replace by 'balancers'.
	Slbs []GetLoadBalancersSlb `pulumi:"slbs"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the SLB instance.
	Tags       map[string]interface{} `pulumi:"tags"`
	TotalCount int                    `pulumi:"totalCount"`
	// ID of the VPC the SLB belongs to.
	VpcId *string `pulumi:"vpcId"`
	// ID of the VSwitch the SLB belongs to.
	VswitchId *string `pulumi:"vswitchId"`
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
	// Service address of the SLBs.
	Address          pulumi.StringPtrInput `pulumi:"address"`
	AddressIpVersion pulumi.StringPtrInput `pulumi:"addressIpVersion"`
	AddressType      pulumi.StringPtrInput `pulumi:"addressType"`
	EnableDetails    pulumi.BoolPtrInput   `pulumi:"enableDetails"`
	// A list of SLBs IDs.
	Ids                pulumi.StringArrayInput `pulumi:"ids"`
	InternetChargeType pulumi.StringPtrInput   `pulumi:"internetChargeType"`
	LoadBalancerName   pulumi.StringPtrInput   `pulumi:"loadBalancerName"`
	MasterZoneId       pulumi.StringPtrInput   `pulumi:"masterZoneId"`
	// A regex string to filter results by SLB name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Network type of the SLBs. Valid values: `vpc` and `classic`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber  pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize    pulumi.IntPtrInput    `pulumi:"pageSize"`
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// The Id of resource group which SLB belongs.
	ResourceGroupId       pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	ServerId              pulumi.StringPtrInput `pulumi:"serverId"`
	ServerIntranetAddress pulumi.StringPtrInput `pulumi:"serverIntranetAddress"`
	SlaveZoneId           pulumi.StringPtrInput `pulumi:"slaveZoneId"`
	// SLB current status. Possible values: `inactive`, `active` and `locked`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A map of tags assigned to the SLB instances. The `tags` can have a maximum of 5 tag. It must be in the format:
	// <!--Start PulumiCodeChooser -->
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
	// 		_, err := slb.GetLoadBalancers(ctx, &slb.GetLoadBalancersArgs{
	// 			Tags: map[string]interface{}{
	// 				"tagKey1": "tagValue1",
	// 				"tagKey2": "tagValue2",
	// 			},
	// 		}, nil)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	})
	// }
	// ```
	// <!--End PulumiCodeChooser -->
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the VPC linked to the SLBs.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// ID of the VSwitch linked to the SLBs.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
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

// Service address of the SLB.
func (o GetLoadBalancersResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.Address }).(pulumi.StringPtrOutput)
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

func (o GetLoadBalancersResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLoadBalancersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of slb IDs.
func (o GetLoadBalancersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetLoadBalancersResultOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) MasterZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.MasterZoneId }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of slb names.
func (o GetLoadBalancersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Network type of the SLB. Possible values: `vpc` and `classic`.
func (o GetLoadBalancersResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetLoadBalancersResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetLoadBalancersResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) ServerIntranetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.ServerIntranetAddress }).(pulumi.StringPtrOutput)
}

func (o GetLoadBalancersResultOutput) SlaveZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.SlaveZoneId }).(pulumi.StringPtrOutput)
}

// A list of SLBs. Each element contains the following attributes:
//
// Deprecated: Field 'slbs' has deprecated from v1.123.1 and replace by 'balancers'.
func (o GetLoadBalancersResultOutput) Slbs() GetLoadBalancersSlbArrayOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) []GetLoadBalancersSlb { return v.Slbs }).(GetLoadBalancersSlbArrayOutput)
}

// SLB current status. Possible values: `inactive`, `active` and `locked`.
func (o GetLoadBalancersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A map of tags assigned to the SLB instance.
func (o GetLoadBalancersResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetLoadBalancersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// ID of the VPC the SLB belongs to.
func (o GetLoadBalancersResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// ID of the VSwitch the SLB belongs to.
func (o GetLoadBalancersResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLoadBalancersResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLoadBalancersResultOutput{})
}
