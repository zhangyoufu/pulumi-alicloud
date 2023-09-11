// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides VPCs available to the user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpcsDs, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				CidrBlock: pulumi.StringRef("172.16.0.0/12"),
//				NameRegex: pulumi.StringRef("^foo"),
//				Status:    pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstVpcId", vpcsDs.Vpcs[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetNetworks(ctx *pulumi.Context, args *GetNetworksArgs, opts ...pulumi.InvokeOption) (*GetNetworksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNetworksResult
	err := ctx.Invoke("alicloud:vpc/getNetworks:getNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworks.
type GetNetworksArgs struct {
	// Filter results by a specific CIDR block. For example: "172.16.0.0/12".
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of dhcp options set.
	DhcpOptionsSetId *string `pulumi:"dhcpOptionsSetId"`
	// Indicates whether to check this request only. Valid values: `true` and `false`.
	DryRun *bool `pulumi:"dryRun"`
	// Default to `true`. Set it to true can output the `routeTableId`.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of VPC IDs.
	Ids []string `pulumi:"ids"`
	// Indicate whether the VPC is the default one in the specified region.
	IsDefault *bool `pulumi:"isDefault"`
	// A regex string to filter VPCs by name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The Id of resource group which VPC belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Filter results by a specific status. Valid value are `Pending` and `Available`.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the VPC.
	VpcName *string `pulumi:"vpcName"`
	// The owner ID of VPC.
	VpcOwnerId *int `pulumi:"vpcOwnerId"`
	// Filter results by the specified VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getNetworks.
type GetNetworksResult struct {
	// CIDR block of the VPC.
	CidrBlock        *string `pulumi:"cidrBlock"`
	DhcpOptionsSetId *string `pulumi:"dhcpOptionsSetId"`
	DryRun           *bool   `pulumi:"dryRun"`
	EnableDetails    *bool   `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of VPC IDs.
	Ids []string `pulumi:"ids"`
	// Whether the VPC is the default VPC in the region.
	IsDefault *bool   `pulumi:"isDefault"`
	NameRegex *string `pulumi:"nameRegex"`
	// A list of VPC names.
	Names           []string `pulumi:"names"`
	OutputFile      *string  `pulumi:"outputFile"`
	PageNumber      *int     `pulumi:"pageNumber"`
	PageSize        *int     `pulumi:"pageSize"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	// Status of the VPC.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the VPC.
	Tags       map[string]interface{} `pulumi:"tags"`
	TotalCount int                    `pulumi:"totalCount"`
	// Name of the VPC.
	VpcName    *string `pulumi:"vpcName"`
	VpcOwnerId *int    `pulumi:"vpcOwnerId"`
	// A list of VPCs. Each element contains the following attributes:
	Vpcs      []GetNetworksVpc `pulumi:"vpcs"`
	VswitchId *string          `pulumi:"vswitchId"`
}

func GetNetworksOutput(ctx *pulumi.Context, args GetNetworksOutputArgs, opts ...pulumi.InvokeOption) GetNetworksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNetworksResult, error) {
			args := v.(GetNetworksArgs)
			r, err := GetNetworks(ctx, &args, opts...)
			var s GetNetworksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNetworksResultOutput)
}

// A collection of arguments for invoking getNetworks.
type GetNetworksOutputArgs struct {
	// Filter results by a specific CIDR block. For example: "172.16.0.0/12".
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
	// The ID of dhcp options set.
	DhcpOptionsSetId pulumi.StringPtrInput `pulumi:"dhcpOptionsSetId"`
	// Indicates whether to check this request only. Valid values: `true` and `false`.
	DryRun pulumi.BoolPtrInput `pulumi:"dryRun"`
	// Default to `true`. Set it to true can output the `routeTableId`.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of VPC IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Indicate whether the VPC is the default one in the specified region.
	IsDefault pulumi.BoolPtrInput `pulumi:"isDefault"`
	// A regex string to filter VPCs by name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The Id of resource group which VPC belongs.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// Filter results by a specific status. Valid value are `Pending` and `Available`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The name of the VPC.
	VpcName pulumi.StringPtrInput `pulumi:"vpcName"`
	// The owner ID of VPC.
	VpcOwnerId pulumi.IntPtrInput `pulumi:"vpcOwnerId"`
	// Filter results by the specified VSwitch.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetNetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworksArgs)(nil)).Elem()
}

// A collection of values returned by getNetworks.
type GetNetworksResultOutput struct{ *pulumi.OutputState }

func (GetNetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworksResult)(nil)).Elem()
}

func (o GetNetworksResultOutput) ToGetNetworksResultOutput() GetNetworksResultOutput {
	return o
}

func (o GetNetworksResultOutput) ToGetNetworksResultOutputWithContext(ctx context.Context) GetNetworksResultOutput {
	return o
}

func (o GetNetworksResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetNetworksResult] {
	return pulumix.Output[GetNetworksResult]{
		OutputState: o.OutputState,
	}
}

// CIDR block of the VPC.
func (o GetNetworksResultOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

func (o GetNetworksResultOutput) DhcpOptionsSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.DhcpOptionsSetId }).(pulumi.StringPtrOutput)
}

func (o GetNetworksResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

func (o GetNetworksResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of VPC IDs.
func (o GetNetworksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Whether the VPC is the default VPC in the region.
func (o GetNetworksResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o GetNetworksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of VPC names.
func (o GetNetworksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetNetworksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetNetworksResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetNetworksResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetNetworksResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// Status of the VPC.
func (o GetNetworksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A map of tags assigned to the VPC.
func (o GetNetworksResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetNetworksResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetNetworksResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetNetworksResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// Name of the VPC.
func (o GetNetworksResultOutput) VpcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.VpcName }).(pulumi.StringPtrOutput)
}

func (o GetNetworksResultOutput) VpcOwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *int { return v.VpcOwnerId }).(pulumi.IntPtrOutput)
}

// A list of VPCs. Each element contains the following attributes:
func (o GetNetworksResultOutput) Vpcs() GetNetworksVpcArrayOutput {
	return o.ApplyT(func(v GetNetworksResult) []GetNetworksVpc { return v.Vpcs }).(GetNetworksVpcArrayOutput)
}

func (o GetNetworksResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworksResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworksResultOutput{})
}
