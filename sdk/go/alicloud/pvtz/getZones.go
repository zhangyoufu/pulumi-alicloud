// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source lists a number of Private Zones resource information owned by an Alibaba Cloud account.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/pvtz"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pvtzZonesDs, err := pvtz.GetZones(ctx, &pvtz.GetZonesArgs{
//				Keyword: pulumi.StringRef(alicloud_pvtz_zone.Basic.Zone_name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstZoneId", pvtzZonesDs.Zones[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetZones(ctx *pulumi.Context, args *GetZonesArgs, opts ...pulumi.InvokeOption) (*GetZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetZonesResult
	err := ctx.Invoke("alicloud:pvtz/getZones:getZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZones.
type GetZonesArgs struct {
	// Default to `false`. Set it to true can output more details.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of zone IDs.
	Ids []string `pulumi:"ids"`
	// keyword for zone name.
	Keyword *string `pulumi:"keyword"`
	// User language.
	Lang      *string `pulumi:"lang"`
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// query_region_id for zone regionId.
	QueryRegionId *string `pulumi:"queryRegionId"`
	// query_vpc_id for zone vpcId.
	QueryVpcId *string `pulumi:"queryVpcId"`
	// resource_group_id for zone resourceGroupId.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Search mode. Value:
	// - LIKE: fuzzy search.
	// - EXACT: precise search. It is not filled in by default.
	SearchMode *string `pulumi:"searchMode"`
}

// A collection of values returned by getZones.
type GetZonesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone IDs.
	Ids       []string `pulumi:"ids"`
	Keyword   *string  `pulumi:"keyword"`
	Lang      *string  `pulumi:"lang"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of zone names.
	Names         []string `pulumi:"names"`
	OutputFile    *string  `pulumi:"outputFile"`
	QueryRegionId *string  `pulumi:"queryRegionId"`
	QueryVpcId    *string  `pulumi:"queryVpcId"`
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	SearchMode      *string `pulumi:"searchMode"`
	// A list of zones. Each element contains the following attributes:
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
	// Default to `false`. Set it to true can output more details.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of zone IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// keyword for zone name.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// User language.
	Lang      pulumi.StringPtrInput `pulumi:"lang"`
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// query_region_id for zone regionId.
	QueryRegionId pulumi.StringPtrInput `pulumi:"queryRegionId"`
	// query_vpc_id for zone vpcId.
	QueryVpcId pulumi.StringPtrInput `pulumi:"queryVpcId"`
	// resource_group_id for zone resourceGroupId.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// Search mode. Value:
	// - LIKE: fuzzy search.
	// - EXACT: precise search. It is not filled in by default.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
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

func (o GetZonesResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.Lang }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of zone names.
func (o GetZonesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) QueryRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.QueryRegionId }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) QueryVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.QueryVpcId }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the Private Zone belongs.
func (o GetZonesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetZonesResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZonesResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

// A list of zones. Each element contains the following attributes:
func (o GetZonesResultOutput) Zones() GetZonesZoneArrayOutput {
	return o.ApplyT(func(v GetZonesResult) []GetZonesZone { return v.Zones }).(GetZonesZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZonesResultOutput{})
}
