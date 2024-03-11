// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cr Endpoint Acl Policies of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.139.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cr.GetEndpointAclPolicies(ctx, &cr.GetEndpointAclPoliciesArgs{
//				InstanceId:   "example_value",
//				EndpointType: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("crEndpointAclPolicyId1", ids.Policies[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEndpointAclPolicies(ctx *pulumi.Context, args *GetEndpointAclPoliciesArgs, opts ...pulumi.InvokeOption) (*GetEndpointAclPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEndpointAclPoliciesResult
	err := ctx.Invoke("alicloud:cr/getEndpointAclPolicies:getEndpointAclPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointAclPolicies.
type GetEndpointAclPoliciesArgs struct {
	// The type of endpoint.
	EndpointType string `pulumi:"endpointType"`
	// A list of Endpoint Acl Policy IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the CR Instance.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getEndpointAclPolicies.
type GetEndpointAclPoliciesResult struct {
	EndpointType string `pulumi:"endpointType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                         `pulumi:"id"`
	Ids        []string                       `pulumi:"ids"`
	InstanceId string                         `pulumi:"instanceId"`
	OutputFile *string                        `pulumi:"outputFile"`
	Policies   []GetEndpointAclPoliciesPolicy `pulumi:"policies"`
}

func GetEndpointAclPoliciesOutput(ctx *pulumi.Context, args GetEndpointAclPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetEndpointAclPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEndpointAclPoliciesResult, error) {
			args := v.(GetEndpointAclPoliciesArgs)
			r, err := GetEndpointAclPolicies(ctx, &args, opts...)
			var s GetEndpointAclPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEndpointAclPoliciesResultOutput)
}

// A collection of arguments for invoking getEndpointAclPolicies.
type GetEndpointAclPoliciesOutputArgs struct {
	// The type of endpoint.
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// A list of Endpoint Acl Policy IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the CR Instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetEndpointAclPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointAclPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getEndpointAclPolicies.
type GetEndpointAclPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetEndpointAclPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEndpointAclPoliciesResult)(nil)).Elem()
}

func (o GetEndpointAclPoliciesResultOutput) ToGetEndpointAclPoliciesResultOutput() GetEndpointAclPoliciesResultOutput {
	return o
}

func (o GetEndpointAclPoliciesResultOutput) ToGetEndpointAclPoliciesResultOutputWithContext(ctx context.Context) GetEndpointAclPoliciesResultOutput {
	return o
}

func (o GetEndpointAclPoliciesResultOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) string { return v.EndpointType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEndpointAclPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEndpointAclPoliciesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEndpointAclPoliciesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetEndpointAclPoliciesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEndpointAclPoliciesResultOutput) Policies() GetEndpointAclPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetEndpointAclPoliciesResult) []GetEndpointAclPoliciesPolicy { return v.Policies }).(GetEndpointAclPoliciesPolicyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEndpointAclPoliciesResultOutput{})
}
