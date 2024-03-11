// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cddc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cddc Zones of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.147.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cddc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := cddc.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cddcZonesId1", example.Zones[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:cddc/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string         `pulumi:"id"`
	Ids        []string       `pulumi:"ids"`
	OutputFile *string        `pulumi:"outputFile"`
	Zones      []GetZonesZone `pulumi:"zones"`
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
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
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

// The provider-assigned unique ID for this managed resource.
func (o GetZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetZonesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) Zones() GetZonesZoneArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []GetZonesZone { return v.Zones }).(GetZonesZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZonesResultOutput{})
}
