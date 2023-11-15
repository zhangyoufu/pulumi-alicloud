// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Resource Manager Shared Resources of the current Alibaba Cloud user.
//
// > **NOTE:** Available since v1.111.0.
func GetSharedResources(ctx *pulumi.Context, args *GetSharedResourcesArgs, opts ...pulumi.InvokeOption) (*GetSharedResourcesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSharedResourcesResult
	err := ctx.Invoke("alicloud:resourcemanager/getSharedResources:getSharedResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedResources.
type GetSharedResourcesArgs struct {
	// A list of shared resource IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The resource share ID of resource manager.
	ResourceShareId *string `pulumi:"resourceShareId"`
	// The status of share resource. Valid values: `Associated`, `Associating`, `Disassociated`, `Disassociating` and `Failed`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getSharedResources.
type GetSharedResourcesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The resource share ID of resource manager.
	ResourceShareId *string `pulumi:"resourceShareId"`
	// A list of Resource Manager Shared Resources. Each element contains the following attributes:
	Resources []GetSharedResourcesResource `pulumi:"resources"`
	// The status of shared resource.
	Status *string `pulumi:"status"`
}

func GetSharedResourcesOutput(ctx *pulumi.Context, args GetSharedResourcesOutputArgs, opts ...pulumi.InvokeOption) GetSharedResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSharedResourcesResult, error) {
			args := v.(GetSharedResourcesArgs)
			r, err := GetSharedResources(ctx, &args, opts...)
			var s GetSharedResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSharedResourcesResultOutput)
}

// A collection of arguments for invoking getSharedResources.
type GetSharedResourcesOutputArgs struct {
	// A list of shared resource IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The resource share ID of resource manager.
	ResourceShareId pulumi.StringPtrInput `pulumi:"resourceShareId"`
	// The status of share resource. Valid values: `Associated`, `Associating`, `Disassociated`, `Disassociating` and `Failed`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetSharedResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSharedResourcesArgs)(nil)).Elem()
}

// A collection of values returned by getSharedResources.
type GetSharedResourcesResultOutput struct{ *pulumi.OutputState }

func (GetSharedResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSharedResourcesResult)(nil)).Elem()
}

func (o GetSharedResourcesResultOutput) ToGetSharedResourcesResultOutput() GetSharedResourcesResultOutput {
	return o
}

func (o GetSharedResourcesResultOutput) ToGetSharedResourcesResultOutputWithContext(ctx context.Context) GetSharedResourcesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSharedResourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSharedResourcesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSharedResourcesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The resource share ID of resource manager.
func (o GetSharedResourcesResultOutput) ResourceShareId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) *string { return v.ResourceShareId }).(pulumi.StringPtrOutput)
}

// A list of Resource Manager Shared Resources. Each element contains the following attributes:
func (o GetSharedResourcesResultOutput) Resources() GetSharedResourcesResourceArrayOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) []GetSharedResourcesResource { return v.Resources }).(GetSharedResourcesResourceArrayOutput)
}

// The status of shared resource.
func (o GetSharedResourcesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSharedResourcesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSharedResourcesResultOutput{})
}
