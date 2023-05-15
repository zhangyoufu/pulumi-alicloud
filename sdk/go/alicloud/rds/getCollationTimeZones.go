// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Operation to query the character set collations and time zones available for use in ApsaraDB RDS.
//
// > **NOTE:** Available in v1.198.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := rds.GetCollationTimeZones(ctx, &rds.GetCollationTimeZonesArgs{
//				OutputFile: pulumi.StringRef("./classes.txt"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRdsCollationTimeZones", zones.CollationTimeZones[0])
//			return nil
//		})
//	}
//
// ```
func GetCollationTimeZones(ctx *pulumi.Context, args *GetCollationTimeZonesArgs, opts ...pulumi.InvokeOption) (*GetCollationTimeZonesResult, error) {
	var rv GetCollationTimeZonesResult
	err := ctx.Invoke("alicloud:rds/getCollationTimeZones:getCollationTimeZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCollationTimeZones.
type GetCollationTimeZonesArgs struct {
	// An array that consists of the character set collations and time zones that are available for
	// use in ApsaraDB RDS.
	CollationTimeZones []GetCollationTimeZonesCollationTimeZone `pulumi:"collationTimeZones"`
	// File name where to save data source results (after running `pulumi up`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getCollationTimeZones.
type GetCollationTimeZonesResult struct {
	CollationTimeZones []GetCollationTimeZonesCollationTimeZone `pulumi:"collationTimeZones"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetCollationTimeZonesOutput(ctx *pulumi.Context, args GetCollationTimeZonesOutputArgs, opts ...pulumi.InvokeOption) GetCollationTimeZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCollationTimeZonesResult, error) {
			args := v.(GetCollationTimeZonesArgs)
			r, err := GetCollationTimeZones(ctx, &args, opts...)
			var s GetCollationTimeZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCollationTimeZonesResultOutput)
}

// A collection of arguments for invoking getCollationTimeZones.
type GetCollationTimeZonesOutputArgs struct {
	// An array that consists of the character set collations and time zones that are available for
	// use in ApsaraDB RDS.
	CollationTimeZones GetCollationTimeZonesCollationTimeZoneArrayInput `pulumi:"collationTimeZones"`
	// File name where to save data source results (after running `pulumi up`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetCollationTimeZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCollationTimeZonesArgs)(nil)).Elem()
}

// A collection of values returned by getCollationTimeZones.
type GetCollationTimeZonesResultOutput struct{ *pulumi.OutputState }

func (GetCollationTimeZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCollationTimeZonesResult)(nil)).Elem()
}

func (o GetCollationTimeZonesResultOutput) ToGetCollationTimeZonesResultOutput() GetCollationTimeZonesResultOutput {
	return o
}

func (o GetCollationTimeZonesResultOutput) ToGetCollationTimeZonesResultOutputWithContext(ctx context.Context) GetCollationTimeZonesResultOutput {
	return o
}

func (o GetCollationTimeZonesResultOutput) CollationTimeZones() GetCollationTimeZonesCollationTimeZoneArrayOutput {
	return o.ApplyT(func(v GetCollationTimeZonesResult) []GetCollationTimeZonesCollationTimeZone {
		return v.CollationTimeZones
	}).(GetCollationTimeZonesCollationTimeZoneArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCollationTimeZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCollationTimeZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCollationTimeZonesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCollationTimeZonesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCollationTimeZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCollationTimeZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCollationTimeZonesResultOutput{})
}
