// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Click House Regions of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.138.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/clickhouse"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clickhouse.GetRegions(ctx, &clickhouse.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = clickhouse.GetRegions(ctx, &clickhouse.GetRegionsArgs{
//				RegionId: pulumi.StringRef("cn-hangzhou"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRegions(ctx *pulumi.Context, args *GetRegionsArgs, opts ...pulumi.InvokeOption) (*GetRegionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRegionsResult
	err := ctx.Invoke("alicloud:clickhouse/getRegions:getRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegions.
type GetRegionsArgs struct {
	// Set to true to match only the region configured in the provider. Default value: `true`.
	Current *bool `pulumi:"current"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// You can use specified regionId to find the region and available zones information that supports ClickHouse.
	RegionId *string `pulumi:"regionId"`
}

// A collection of values returned by getRegions.
type GetRegionsResult struct {
	Current *bool `pulumi:"current"`
	// The provider-assigned unique ID for this managed resource.
	Id         string             `pulumi:"id"`
	OutputFile *string            `pulumi:"outputFile"`
	RegionId   *string            `pulumi:"regionId"`
	Regions    []GetRegionsRegion `pulumi:"regions"`
}

func GetRegionsOutput(ctx *pulumi.Context, args GetRegionsOutputArgs, opts ...pulumi.InvokeOption) GetRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegionsResult, error) {
			args := v.(GetRegionsArgs)
			r, err := GetRegions(ctx, &args, opts...)
			var s GetRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegionsResultOutput)
}

// A collection of arguments for invoking getRegions.
type GetRegionsOutputArgs struct {
	// Set to true to match only the region configured in the provider. Default value: `true`.
	Current pulumi.BoolPtrInput `pulumi:"current"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// You can use specified regionId to find the region and available zones information that supports ClickHouse.
	RegionId pulumi.StringPtrInput `pulumi:"regionId"`
}

func (GetRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionsArgs)(nil)).Elem()
}

// A collection of values returned by getRegions.
type GetRegionsResultOutput struct{ *pulumi.OutputState }

func (GetRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionsResult)(nil)).Elem()
}

func (o GetRegionsResultOutput) ToGetRegionsResultOutput() GetRegionsResultOutput {
	return o
}

func (o GetRegionsResultOutput) ToGetRegionsResultOutputWithContext(ctx context.Context) GetRegionsResultOutput {
	return o
}

func (o GetRegionsResultOutput) Current() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRegionsResult) *bool { return v.Current }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRegionsResultOutput) RegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionsResult) *string { return v.RegionId }).(pulumi.StringPtrOutput)
}

func (o GetRegionsResultOutput) Regions() GetRegionsRegionArrayOutput {
	return o.ApplyT(func(v GetRegionsResult) []GetRegionsRegion { return v.Regions }).(GetRegionsRegionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegionsResultOutput{})
}
