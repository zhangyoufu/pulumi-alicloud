// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Actiontrail Global Events Storage Region of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.201.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/actiontrail"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := actiontrail.LookupGlobalEventsStorageRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudActiontrailGlobalEventsStorageRegion1", _default.StorageRegion)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupGlobalEventsStorageRegion(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupGlobalEventsStorageRegionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalEventsStorageRegionResult
	err := ctx.Invoke("alicloud:actiontrail/getGlobalEventsStorageRegion:getGlobalEventsStorageRegion", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getGlobalEventsStorageRegion.
type LookupGlobalEventsStorageRegionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	StorageRegion string `pulumi:"storageRegion"`
}

func LookupGlobalEventsStorageRegionOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupGlobalEventsStorageRegionResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupGlobalEventsStorageRegionResult, error) {
		r, err := LookupGlobalEventsStorageRegion(ctx, opts...)
		var s LookupGlobalEventsStorageRegionResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(LookupGlobalEventsStorageRegionResultOutput)
}

// A collection of values returned by getGlobalEventsStorageRegion.
type LookupGlobalEventsStorageRegionResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalEventsStorageRegionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalEventsStorageRegionResult)(nil)).Elem()
}

func (o LookupGlobalEventsStorageRegionResultOutput) ToLookupGlobalEventsStorageRegionResultOutput() LookupGlobalEventsStorageRegionResultOutput {
	return o
}

func (o LookupGlobalEventsStorageRegionResultOutput) ToLookupGlobalEventsStorageRegionResultOutputWithContext(ctx context.Context) LookupGlobalEventsStorageRegionResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGlobalEventsStorageRegionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalEventsStorageRegionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalEventsStorageRegionResultOutput) StorageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalEventsStorageRegionResult) string { return v.StorageRegion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalEventsStorageRegionResultOutput{})
}
