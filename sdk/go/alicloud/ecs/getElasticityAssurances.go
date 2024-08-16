// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Ecs Elasticity Assurance available to the user.
//
// > **NOTE:** Available in 1.196.0+
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _default, err := ecs.GetElasticityAssurances(ctx, &ecs.GetElasticityAssurancesArgs{
// Ids: interface{}{
// defaultAlicloudEcsElasticityAssurance.Id,
// },
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("alicloudEcsElasticityAssuranceExampleId", _default.Assurances[0].Id)
// return nil
// })
// }
// ```
func GetElasticityAssurances(ctx *pulumi.Context, args *GetElasticityAssurancesArgs, opts ...pulumi.InvokeOption) (*GetElasticityAssurancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetElasticityAssurancesResult
	err := ctx.Invoke("alicloud:ecs/getElasticityAssurances:getElasticityAssurances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getElasticityAssurances.
type GetElasticityAssurancesArgs struct {
	// A list of Elasticity Assurance IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the elastic protection service.
	PrivatePoolOptionsIds []string `pulumi:"privatePoolOptionsIds"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of flexible guarantee services. Possible values: `All`, `Preparing`, `Prepared`, `Active`, `Released`.
	Status *string `pulumi:"status"`
	// The tag key-value pair information bound by the elastic guarantee service.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getElasticityAssurances.
type GetElasticityAssurancesResult struct {
	// A list of Elasticity Assurance Entries. Each element contains the following attributes:
	Assurances []GetElasticityAssurancesAssurance `pulumi:"assurances"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Elasticity Assurance IDs.
	Ids                   []string `pulumi:"ids"`
	OutputFile            *string  `pulumi:"outputFile"`
	PrivatePoolOptionsIds []string `pulumi:"privatePoolOptionsIds"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of flexible guarantee services. Possible values:-Preparing: in preparation.-Prepared: to take effect.-Active: in effect.-Released: Released.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the Capacity Reservation.
	Tags map[string]string `pulumi:"tags"`
}

func GetElasticityAssurancesOutput(ctx *pulumi.Context, args GetElasticityAssurancesOutputArgs, opts ...pulumi.InvokeOption) GetElasticityAssurancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetElasticityAssurancesResult, error) {
			args := v.(GetElasticityAssurancesArgs)
			r, err := GetElasticityAssurances(ctx, &args, opts...)
			var s GetElasticityAssurancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetElasticityAssurancesResultOutput)
}

// A collection of arguments for invoking getElasticityAssurances.
type GetElasticityAssurancesOutputArgs struct {
	// A list of Elasticity Assurance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the elastic protection service.
	PrivatePoolOptionsIds pulumi.StringArrayInput `pulumi:"privatePoolOptionsIds"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The status of flexible guarantee services. Possible values: `All`, `Preparing`, `Prepared`, `Active`, `Released`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The tag key-value pair information bound by the elastic guarantee service.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetElasticityAssurancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElasticityAssurancesArgs)(nil)).Elem()
}

// A collection of values returned by getElasticityAssurances.
type GetElasticityAssurancesResultOutput struct{ *pulumi.OutputState }

func (GetElasticityAssurancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetElasticityAssurancesResult)(nil)).Elem()
}

func (o GetElasticityAssurancesResultOutput) ToGetElasticityAssurancesResultOutput() GetElasticityAssurancesResultOutput {
	return o
}

func (o GetElasticityAssurancesResultOutput) ToGetElasticityAssurancesResultOutputWithContext(ctx context.Context) GetElasticityAssurancesResultOutput {
	return o
}

// A list of Elasticity Assurance Entries. Each element contains the following attributes:
func (o GetElasticityAssurancesResultOutput) Assurances() GetElasticityAssurancesAssuranceArrayOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) []GetElasticityAssurancesAssurance { return v.Assurances }).(GetElasticityAssurancesAssuranceArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetElasticityAssurancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of Elasticity Assurance IDs.
func (o GetElasticityAssurancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetElasticityAssurancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetElasticityAssurancesResultOutput) PrivatePoolOptionsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) []string { return v.PrivatePoolOptionsIds }).(pulumi.StringArrayOutput)
}

// The ID of the resource group.
func (o GetElasticityAssurancesResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The status of flexible guarantee services. Possible values:-Preparing: in preparation.-Prepared: to take effect.-Active: in effect.-Released: Released.
func (o GetElasticityAssurancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the Capacity Reservation.
func (o GetElasticityAssurancesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetElasticityAssurancesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetElasticityAssurancesResultOutput{})
}
