// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides available scaling configuration resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			scalingconfigurationsDs, err := ess.GetScalingConfigurations(ctx, &ess.GetScalingConfigurationsArgs{
//				Ids: []string{
//					"scaling_configuration_id1",
//					"scaling_configuration_id2",
//				},
//				NameRegex:      pulumi.StringRef("scaling_configuration_name"),
//				ScalingGroupId: pulumi.StringRef("scaling_group_id"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstScalingRule", scalingconfigurationsDs.Configurations[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetScalingConfigurations(ctx *pulumi.Context, args *GetScalingConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetScalingConfigurationsResult, error) {
	var rv GetScalingConfigurationsResult
	err := ctx.Invoke("alicloud:ess/getScalingConfigurations:getScalingConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingConfigurations.
type GetScalingConfigurationsArgs struct {
	// A list of scaling configuration IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter resulting scaling configurations by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Scaling group id the scaling configurations belong to.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

// A collection of values returned by getScalingConfigurations.
type GetScalingConfigurationsResult struct {
	// A list of scaling rules. Each element contains the following attributes:
	Configurations []GetScalingConfigurationsConfiguration `pulumi:"configurations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of scaling configuration ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of scaling configuration names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// ID of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

func GetScalingConfigurationsOutput(ctx *pulumi.Context, args GetScalingConfigurationsOutputArgs, opts ...pulumi.InvokeOption) GetScalingConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScalingConfigurationsResult, error) {
			args := v.(GetScalingConfigurationsArgs)
			r, err := GetScalingConfigurations(ctx, &args, opts...)
			var s GetScalingConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScalingConfigurationsResultOutput)
}

// A collection of arguments for invoking getScalingConfigurations.
type GetScalingConfigurationsOutputArgs struct {
	// A list of scaling configuration IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter resulting scaling configurations by name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Scaling group id the scaling configurations belong to.
	ScalingGroupId pulumi.StringPtrInput `pulumi:"scalingGroupId"`
}

func (GetScalingConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingConfigurationsArgs)(nil)).Elem()
}

// A collection of values returned by getScalingConfigurations.
type GetScalingConfigurationsResultOutput struct{ *pulumi.OutputState }

func (GetScalingConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingConfigurationsResult)(nil)).Elem()
}

func (o GetScalingConfigurationsResultOutput) ToGetScalingConfigurationsResultOutput() GetScalingConfigurationsResultOutput {
	return o
}

func (o GetScalingConfigurationsResultOutput) ToGetScalingConfigurationsResultOutputWithContext(ctx context.Context) GetScalingConfigurationsResultOutput {
	return o
}

// A list of scaling rules. Each element contains the following attributes:
func (o GetScalingConfigurationsResultOutput) Configurations() GetScalingConfigurationsConfigurationArrayOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) []GetScalingConfigurationsConfiguration {
		return v.Configurations
	}).(GetScalingConfigurationsConfigurationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetScalingConfigurationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of scaling configuration ids.
func (o GetScalingConfigurationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetScalingConfigurationsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of scaling configuration names.
func (o GetScalingConfigurationsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetScalingConfigurationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// ID of the scaling group.
func (o GetScalingConfigurationsResultOutput) ScalingGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigurationsResult) *string { return v.ScalingGroupId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScalingConfigurationsResultOutput{})
}
