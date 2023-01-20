// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the ECS instance types of Alibaba Cloud.
//
// > **NOTE:** By default, only the upgraded instance types are returned. If you want to get outdated instance types, you must set `isOutdated` to true.
//
// > **NOTE:** If one instance type is sold out, it will not be exported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			typesDs, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				CpuCoreCount: pulumi.IntRef(1),
//				MemorySize:   pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewInstance(ctx, "instance", &ecs.InstanceArgs{
//				InstanceType: *pulumi.String(typesDs.InstanceTypes[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceTypes(ctx *pulumi.Context, args *GetInstanceTypesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypesResult, error) {
	var rv GetInstanceTypesResult
	err := ctx.Invoke("alicloud:ecs/getInstanceTypes:getInstanceTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesArgs struct {
	// The zone where instance types are supported.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Filter the results to a specific number of cpu cores.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// Filter the result whose network interface number is no more than `eniAmount`.
	EniAmount *int `pulumi:"eniAmount"`
	// The GPU amount of an instance type.
	GpuAmount *int `pulumi:"gpuAmount"`
	// The GPU spec of an instance type.
	GpuSpec *string `pulumi:"gpuSpec"`
	// The ID of the image.
	ImageId *string `pulumi:"imageId"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Filter the results based on their family name. For example: 'ecs.n4'.
	InstanceTypeFamily *string `pulumi:"instanceTypeFamily"`
	// If true, outdated instance types are included in the results. Default to false.
	IsOutdated         *bool   `pulumi:"isOutdated"`
	KubernetesNodeRole *string `pulumi:"kubernetesNodeRole"`
	// Filter the results to a specific memory size in GB.
	MemorySize *float64 `pulumi:"memorySize"`
	// The minimum number of IPv6 addresses per ENI. **Note:** If an instance type supports fewer IPv6 addresses per ENI than the specified value, information about the instance type is not queried.
	MinimumEniIpv6AddressQuantity *int `pulumi:"minimumEniIpv6AddressQuantity"`
	// Filter the results by network type. Valid values: `Classic` and `Vpc`.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	SortedBy    *string `pulumi:"sortedBy"`
	// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// Filter the results by system disk category. Valid values: `cloud`, `ephemeralSsd`, `cloudEssd`, `cloudEfficiency`, `cloudSsd`.
	// **NOTE**: Its default value `cloudEfficiency` has been removed from the version v1.150.0.
	SystemDiskCategory *string `pulumi:"systemDiskCategory"`
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Number of CPU cores.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// The maximum number of network interfaces that an instance type can be attached to.
	EniAmount *int    `pulumi:"eniAmount"`
	GpuAmount *int    `pulumi:"gpuAmount"`
	GpuSpec   *string `pulumi:"gpuSpec"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instance type IDs.
	Ids                []string `pulumi:"ids"`
	ImageId            *string  `pulumi:"imageId"`
	InstanceChargeType *string  `pulumi:"instanceChargeType"`
	InstanceTypeFamily *string  `pulumi:"instanceTypeFamily"`
	// A list of image types. Each element contains the following attributes:
	InstanceTypes      []GetInstanceTypesInstanceType `pulumi:"instanceTypes"`
	IsOutdated         *bool                          `pulumi:"isOutdated"`
	KubernetesNodeRole *string                        `pulumi:"kubernetesNodeRole"`
	// Size of memory, measured in GB.
	MemorySize                    *float64 `pulumi:"memorySize"`
	MinimumEniIpv6AddressQuantity *int     `pulumi:"minimumEniIpv6AddressQuantity"`
	NetworkType                   *string  `pulumi:"networkType"`
	OutputFile                    *string  `pulumi:"outputFile"`
	SortedBy                      *string  `pulumi:"sortedBy"`
	SpotStrategy                  *string  `pulumi:"spotStrategy"`
	SystemDiskCategory            *string  `pulumi:"systemDiskCategory"`
}

func GetInstanceTypesOutput(ctx *pulumi.Context, args GetInstanceTypesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceTypesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceTypesResult, error) {
			args := v.(GetInstanceTypesArgs)
			r, err := GetInstanceTypes(ctx, &args, opts...)
			var s GetInstanceTypesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceTypesResultOutput)
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesOutputArgs struct {
	// The zone where instance types are supported.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Filter the results to a specific number of cpu cores.
	CpuCoreCount pulumi.IntPtrInput `pulumi:"cpuCoreCount"`
	// Filter the result whose network interface number is no more than `eniAmount`.
	EniAmount pulumi.IntPtrInput `pulumi:"eniAmount"`
	// The GPU amount of an instance type.
	GpuAmount pulumi.IntPtrInput `pulumi:"gpuAmount"`
	// The GPU spec of an instance type.
	GpuSpec pulumi.StringPtrInput `pulumi:"gpuSpec"`
	// The ID of the image.
	ImageId pulumi.StringPtrInput `pulumi:"imageId"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput `pulumi:"instanceChargeType"`
	// Filter the results based on their family name. For example: 'ecs.n4'.
	InstanceTypeFamily pulumi.StringPtrInput `pulumi:"instanceTypeFamily"`
	// If true, outdated instance types are included in the results. Default to false.
	IsOutdated         pulumi.BoolPtrInput   `pulumi:"isOutdated"`
	KubernetesNodeRole pulumi.StringPtrInput `pulumi:"kubernetesNodeRole"`
	// Filter the results to a specific memory size in GB.
	MemorySize pulumi.Float64PtrInput `pulumi:"memorySize"`
	// The minimum number of IPv6 addresses per ENI. **Note:** If an instance type supports fewer IPv6 addresses per ENI than the specified value, information about the instance type is not queried.
	MinimumEniIpv6AddressQuantity pulumi.IntPtrInput `pulumi:"minimumEniIpv6AddressQuantity"`
	// Filter the results by network type. Valid values: `Classic` and `Vpc`.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
	SortedBy    pulumi.StringPtrInput `pulumi:"sortedBy"`
	// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
	SpotStrategy pulumi.StringPtrInput `pulumi:"spotStrategy"`
	// Filter the results by system disk category. Valid values: `cloud`, `ephemeralSsd`, `cloudEssd`, `cloudEfficiency`, `cloudSsd`.
	// **NOTE**: Its default value `cloudEfficiency` has been removed from the version v1.150.0.
	SystemDiskCategory pulumi.StringPtrInput `pulumi:"systemDiskCategory"`
}

func (GetInstanceTypesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceTypesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceTypesResult)(nil)).Elem()
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutput() GetInstanceTypesResultOutput {
	return o
}

func (o GetInstanceTypesResultOutput) ToGetInstanceTypesResultOutputWithContext(ctx context.Context) GetInstanceTypesResultOutput {
	return o
}

func (o GetInstanceTypesResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Number of CPU cores.
func (o GetInstanceTypesResultOutput) CpuCoreCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *int { return v.CpuCoreCount }).(pulumi.IntPtrOutput)
}

// The maximum number of network interfaces that an instance type can be attached to.
func (o GetInstanceTypesResultOutput) EniAmount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *int { return v.EniAmount }).(pulumi.IntPtrOutput)
}

func (o GetInstanceTypesResultOutput) GpuAmount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *int { return v.GpuAmount }).(pulumi.IntPtrOutput)
}

func (o GetInstanceTypesResultOutput) GpuSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.GpuSpec }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceTypesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of instance type IDs.
func (o GetInstanceTypesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstanceTypesResultOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) InstanceTypeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.InstanceTypeFamily }).(pulumi.StringPtrOutput)
}

// A list of image types. Each element contains the following attributes:
func (o GetInstanceTypesResultOutput) InstanceTypes() GetInstanceTypesInstanceTypeArrayOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) []GetInstanceTypesInstanceType { return v.InstanceTypes }).(GetInstanceTypesInstanceTypeArrayOutput)
}

func (o GetInstanceTypesResultOutput) IsOutdated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *bool { return v.IsOutdated }).(pulumi.BoolPtrOutput)
}

func (o GetInstanceTypesResultOutput) KubernetesNodeRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.KubernetesNodeRole }).(pulumi.StringPtrOutput)
}

// Size of memory, measured in GB.
func (o GetInstanceTypesResultOutput) MemorySize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *float64 { return v.MemorySize }).(pulumi.Float64PtrOutput)
}

func (o GetInstanceTypesResultOutput) MinimumEniIpv6AddressQuantity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *int { return v.MinimumEniIpv6AddressQuantity }).(pulumi.IntPtrOutput)
}

func (o GetInstanceTypesResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) SortedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.SortedBy }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) SpotStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.SpotStrategy }).(pulumi.StringPtrOutput)
}

func (o GetInstanceTypesResultOutput) SystemDiskCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceTypesResult) *string { return v.SystemDiskCategory }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceTypesResultOutput{})
}
