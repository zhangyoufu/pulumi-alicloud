// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Storage Gateway Gateway Cache Disks of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.144.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudstoragegateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cloudstoragegateway.GetGatewayCacheDisks(ctx, &cloudstoragegateway.GetGatewayCacheDisksArgs{
//				GatewayId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cloudStorageGatewayGatewayCacheDiskId1", ids.Disks[0].Id)
//			status, err := cloudstoragegateway.GetGatewayCacheDisks(ctx, &cloudstoragegateway.GetGatewayCacheDisksArgs{
//				GatewayId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//				Status: pulumi.IntRef(0),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cloudStorageGatewayGatewayCacheDiskId2", status.Disks[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetGatewayCacheDisks(ctx *pulumi.Context, args *GetGatewayCacheDisksArgs, opts ...pulumi.InvokeOption) (*GetGatewayCacheDisksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGatewayCacheDisksResult
	err := ctx.Invoke("alicloud:cloudstoragegateway/getGatewayCacheDisks:getGatewayCacheDisks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayCacheDisks.
type GetGatewayCacheDisksArgs struct {
	// The ID of the gateway.
	GatewayId string `pulumi:"gatewayId"`
	// A list of Gateway Cache Disk IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource.
	Status *int `pulumi:"status"`
}

// A collection of values returned by getGatewayCacheDisks.
type GetGatewayCacheDisksResult struct {
	Disks     []GetGatewayCacheDisksDisk `pulumi:"disks"`
	GatewayId string                     `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *int     `pulumi:"status"`
}

func GetGatewayCacheDisksOutput(ctx *pulumi.Context, args GetGatewayCacheDisksOutputArgs, opts ...pulumi.InvokeOption) GetGatewayCacheDisksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewayCacheDisksResult, error) {
			args := v.(GetGatewayCacheDisksArgs)
			r, err := GetGatewayCacheDisks(ctx, &args, opts...)
			var s GetGatewayCacheDisksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewayCacheDisksResultOutput)
}

// A collection of arguments for invoking getGatewayCacheDisks.
type GetGatewayCacheDisksOutputArgs struct {
	// The ID of the gateway.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// A list of Gateway Cache Disk IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource.
	Status pulumi.IntPtrInput `pulumi:"status"`
}

func (GetGatewayCacheDisksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayCacheDisksArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayCacheDisks.
type GetGatewayCacheDisksResultOutput struct{ *pulumi.OutputState }

func (GetGatewayCacheDisksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayCacheDisksResult)(nil)).Elem()
}

func (o GetGatewayCacheDisksResultOutput) ToGetGatewayCacheDisksResultOutput() GetGatewayCacheDisksResultOutput {
	return o
}

func (o GetGatewayCacheDisksResultOutput) ToGetGatewayCacheDisksResultOutputWithContext(ctx context.Context) GetGatewayCacheDisksResultOutput {
	return o
}

func (o GetGatewayCacheDisksResultOutput) Disks() GetGatewayCacheDisksDiskArrayOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) []GetGatewayCacheDisksDisk { return v.Disks }).(GetGatewayCacheDisksDiskArrayOutput)
}

func (o GetGatewayCacheDisksResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewayCacheDisksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewayCacheDisksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGatewayCacheDisksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGatewayCacheDisksResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGatewayCacheDisksResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewayCacheDisksResultOutput{})
}
