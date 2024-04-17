// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alicloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides availability zones that can be accessed by an Alibaba Cloud account within the region configured in the provider.
//
// > **NOTE:** If one zone is sold out, it will not be exported.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Declare the data source
//			_, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableInstanceType: pulumi.StringRef("ecs.n4.large"),
//				AvailableDiskCategory: pulumi.StringRef("cloud_ssd"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:index/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
	AvailableDiskCategory *string `pulumi:"availableDiskCategory"`
	// Filter the results by a specific instance type.
	AvailableInstanceType *string `pulumi:"availableInstanceType"`
	// Filter the results by a specific resource type.
	// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
	//
	// > **NOTE:** From version 1.134.0, the `availableResourceCreation` value "Rds" has been deprecated.
	// If you want to fetch the available zones for RDS instance, you can use datasource alicloud_db_zones
	AvailableResourceCreation *string `pulumi:"availableResourceCreation"`
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	//
	// > **NOTE:** The disk category `cloud` has been outdated and can only be used by non-I/O Optimized ECS instances. Many availability zones don't support it. It is recommended to use `cloudEfficiency` or `cloudSsd`.
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType *string `pulumi:"availableSlbAddressType"`
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
	Multi *bool `pulumi:"multi"`
	// Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
	NetworkType *string `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy *string `pulumi:"spotStrategy"`
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	AvailableDiskCategory *string `pulumi:"availableDiskCategory"`
	AvailableInstanceType *string `pulumi:"availableInstanceType"`
	// Type of resources that can be created.
	AvailableResourceCreation    *string `pulumi:"availableResourceCreation"`
	AvailableSlbAddressIpVersion *string `pulumi:"availableSlbAddressIpVersion"`
	AvailableSlbAddressType      *string `pulumi:"availableSlbAddressType"`
	EnableDetails                *bool   `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone IDs.
	Ids                []string `pulumi:"ids"`
	InstanceChargeType *string  `pulumi:"instanceChargeType"`
	Multi              *bool    `pulumi:"multi"`
	NetworkType        *string  `pulumi:"networkType"`
	OutputFile         *string  `pulumi:"outputFile"`
	SpotStrategy       *string  `pulumi:"spotStrategy"`
	// A list of availability zones. Each element contains the following attributes:
	Zones []GetZonesZone `pulumi:"zones"`
}

func GetZonesOutput(ctx *pulumi.Context, args GetZonesOutputArgs, opts ...pulumi.InvokeOption) GetZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZonesResult, error) {
			args := v.(GetZonesArgs)
			r, err := GetZones(ctx, &args, opts...)
			var s GetZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZonesResultOutput)
}

// A collection of arguments for invoking getZones.
type GetZonesOutputArgs struct {
	// Filter the results by a specific disk category. Can be either `cloud`, `cloudEfficiency`, `cloudSsd`, `ephemeralSsd`.
	AvailableDiskCategory pulumi.StringPtrInput `pulumi:"availableDiskCategory"`
	// Filter the results by a specific instance type.
	AvailableInstanceType pulumi.StringPtrInput `pulumi:"availableInstanceType"`
	// Filter the results by a specific resource type.
	// Valid values: `Instance`, `Disk`, `VSwitch`, `Rds`, `KVStore`, `FunctionCompute`, `Elasticsearch`, `Slb`.
	//
	// > **NOTE:** From version 1.134.0, the `availableResourceCreation` value "Rds" has been deprecated.
	// If you want to fetch the available zones for RDS instance, you can use datasource alicloud_db_zones
	AvailableResourceCreation pulumi.StringPtrInput `pulumi:"availableResourceCreation"`
	// Filter the results by a slb instance address version. Can be either `ipv4`, or `ipv6`.
	//
	// > **NOTE:** The disk category `cloud` has been outdated and can only be used by non-I/O Optimized ECS instances. Many availability zones don't support it. It is recommended to use `cloudEfficiency` or `cloudSsd`.
	AvailableSlbAddressIpVersion pulumi.StringPtrInput `pulumi:"availableSlbAddressIpVersion"`
	// Filter the results by a slb instance address type. Can be either `Vpc`, `classicInternet` or `classicIntranet`
	AvailableSlbAddressType pulumi.StringPtrInput `pulumi:"availableSlbAddressType"`
	// Default to false and only output `id` in the `zones` block. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// Filter the results by a specific ECS instance charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput `pulumi:"instanceChargeType"`
	// Indicate whether the zones can be used in a multi AZ configuration. Default to `false`. Multi AZ is usually used to launch RDS instances.
	Multi pulumi.BoolPtrInput `pulumi:"multi"`
	// Filter the results by a specific network type. Valid values: `Classic` and `Vpc`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// - (Optional) Filter the results by a specific ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy pulumi.StringPtrInput `pulumi:"spotStrategy"`
}

func (GetZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesArgs)(nil)).Elem()
}

// A collection of values returned by getZones.
type GetZonesResultOutput struct{ *pulumi.OutputState }

func (GetZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZonesResult)(nil)).Elem()
}

func (o GetZonesResultOutput) ToGetZonesResultOutput() GetZonesResultOutput {
	return o
}

func (o GetZonesResultOutput) ToGetZonesResultOutputWithContext(ctx context.Context) GetZonesResultOutput {
	return o
}

func (o GetZonesResultOutput) AvailableDiskCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableDiskCategory }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) AvailableInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableInstanceType }).(pulumi.StringPtrOutput)
}

// Type of resources that can be created.
func (o GetZonesResultOutput) AvailableResourceCreation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableResourceCreation }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) AvailableSlbAddressIpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableSlbAddressIpVersion }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) AvailableSlbAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.AvailableSlbAddressType }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zone IDs.
func (o GetZonesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetZonesResultOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) Multi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *bool { return v.Multi }).(pulumi.BoolPtrOutput)
}

func (o GetZonesResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) SpotStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.SpotStrategy }).(pulumi.StringPtrOutput)
}

// A list of availability zones. Each element contains the following attributes:
func (o GetZonesResultOutput) Zones() GetZonesZoneArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []GetZonesZone { return v.Zones }).(GetZonesZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZonesResultOutput{})
}
