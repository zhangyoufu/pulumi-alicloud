// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the master slave server groups related to a server load balancer.
//
// > **NOTE:** Available in 1.54.0+
func GetMasterSlaveServerGroups(ctx *pulumi.Context, args *GetMasterSlaveServerGroupsArgs, opts ...pulumi.InvokeOption) (*GetMasterSlaveServerGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMasterSlaveServerGroupsResult
	err := ctx.Invoke("alicloud:slb/getMasterSlaveServerGroups:getMasterSlaveServerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMasterSlaveServerGroups.
type GetMasterSlaveServerGroupsArgs struct {
	// A list of master slave server group IDs to filter results.
	Ids []string `pulumi:"ids"`
	// ID of the SLB.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A regex string to filter results by master slave server group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getMasterSlaveServerGroups.
type GetMasterSlaveServerGroupsResult struct {
	// A list of SLB master slave server groups. Each element contains the following attributes:
	Groups []GetMasterSlaveServerGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of SLB master slave server groups IDs.
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	// A list of SLB master slave server groups names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetMasterSlaveServerGroupsOutput(ctx *pulumi.Context, args GetMasterSlaveServerGroupsOutputArgs, opts ...pulumi.InvokeOption) GetMasterSlaveServerGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMasterSlaveServerGroupsResult, error) {
			args := v.(GetMasterSlaveServerGroupsArgs)
			r, err := GetMasterSlaveServerGroups(ctx, &args, opts...)
			var s GetMasterSlaveServerGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMasterSlaveServerGroupsResultOutput)
}

// A collection of arguments for invoking getMasterSlaveServerGroups.
type GetMasterSlaveServerGroupsOutputArgs struct {
	// A list of master slave server group IDs to filter results.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// ID of the SLB.
	LoadBalancerId pulumi.StringInput `pulumi:"loadBalancerId"`
	// A regex string to filter results by master slave server group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetMasterSlaveServerGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMasterSlaveServerGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getMasterSlaveServerGroups.
type GetMasterSlaveServerGroupsResultOutput struct{ *pulumi.OutputState }

func (GetMasterSlaveServerGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMasterSlaveServerGroupsResult)(nil)).Elem()
}

func (o GetMasterSlaveServerGroupsResultOutput) ToGetMasterSlaveServerGroupsResultOutput() GetMasterSlaveServerGroupsResultOutput {
	return o
}

func (o GetMasterSlaveServerGroupsResultOutput) ToGetMasterSlaveServerGroupsResultOutputWithContext(ctx context.Context) GetMasterSlaveServerGroupsResultOutput {
	return o
}

func (o GetMasterSlaveServerGroupsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetMasterSlaveServerGroupsResult] {
	return pulumix.Output[GetMasterSlaveServerGroupsResult]{
		OutputState: o.OutputState,
	}
}

// A list of SLB master slave server groups. Each element contains the following attributes:
func (o GetMasterSlaveServerGroupsResultOutput) Groups() GetMasterSlaveServerGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) []GetMasterSlaveServerGroupsGroup { return v.Groups }).(GetMasterSlaveServerGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMasterSlaveServerGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of SLB master slave server groups IDs.
func (o GetMasterSlaveServerGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetMasterSlaveServerGroupsResultOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) string { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o GetMasterSlaveServerGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of SLB master slave server groups names.
func (o GetMasterSlaveServerGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetMasterSlaveServerGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMasterSlaveServerGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMasterSlaveServerGroupsResultOutput{})
}
