// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Activations of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.177.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ecs.GetActivations(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsActivationId1", ids.Activations[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetActivations(ctx *pulumi.Context, args *GetActivationsArgs, opts ...pulumi.InvokeOption) (*GetActivationsResult, error) {
	var rv GetActivationsResult
	err := ctx.Invoke("alicloud:ecs/getActivations:getActivations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActivations.
type GetActivationsArgs struct {
	// A list of Activation IDs.
	Ids []string `pulumi:"ids"`
	// The default prefix of the instance name.
	InstanceName *string `pulumi:"instanceName"`
	OutputFile   *string `pulumi:"outputFile"`
	PageNumber   *int    `pulumi:"pageNumber"`
	PageSize     *int    `pulumi:"pageSize"`
}

// A collection of values returned by getActivations.
type GetActivationsResult struct {
	Activations []GetActivationsActivation `pulumi:"activations"`
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	Ids          []string `pulumi:"ids"`
	InstanceName *string  `pulumi:"instanceName"`
	OutputFile   *string  `pulumi:"outputFile"`
	PageNumber   *int     `pulumi:"pageNumber"`
	PageSize     *int     `pulumi:"pageSize"`
	TotalCount   int      `pulumi:"totalCount"`
}

func GetActivationsOutput(ctx *pulumi.Context, args GetActivationsOutputArgs, opts ...pulumi.InvokeOption) GetActivationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActivationsResult, error) {
			args := v.(GetActivationsArgs)
			r, err := GetActivations(ctx, &args, opts...)
			var s GetActivationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActivationsResultOutput)
}

// A collection of arguments for invoking getActivations.
type GetActivationsOutputArgs struct {
	// A list of Activation IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The default prefix of the instance name.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	OutputFile   pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber   pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize     pulumi.IntPtrInput    `pulumi:"pageSize"`
}

func (GetActivationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActivationsArgs)(nil)).Elem()
}

// A collection of values returned by getActivations.
type GetActivationsResultOutput struct{ *pulumi.OutputState }

func (GetActivationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActivationsResult)(nil)).Elem()
}

func (o GetActivationsResultOutput) ToGetActivationsResultOutput() GetActivationsResultOutput {
	return o
}

func (o GetActivationsResultOutput) ToGetActivationsResultOutputWithContext(ctx context.Context) GetActivationsResultOutput {
	return o
}

func (o GetActivationsResultOutput) Activations() GetActivationsActivationArrayOutput {
	return o.ApplyT(func(v GetActivationsResult) []GetActivationsActivation { return v.Activations }).(GetActivationsActivationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetActivationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActivationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetActivationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetActivationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetActivationsResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActivationsResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

func (o GetActivationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActivationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetActivationsResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetActivationsResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetActivationsResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetActivationsResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetActivationsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetActivationsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActivationsResultOutput{})
}
