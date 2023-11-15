// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Vpc Ipv6 Egress Rules of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.142.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := vpc.GetIpv6EgressRules(ctx, &vpc.GetIpv6EgressRulesArgs{
//				Ipv6GatewayId: "example_value",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6EgressRuleId1", ids.Rules[0].Id)
//			nameRegex, err := vpc.GetIpv6EgressRules(ctx, &vpc.GetIpv6EgressRulesArgs{
//				Ipv6GatewayId: "example_value",
//				NameRegex:     pulumi.StringRef("^my-Ipv6EgressRule"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6EgressRuleId2", nameRegex.Rules[0].Id)
//			status, err := vpc.GetIpv6EgressRules(ctx, &vpc.GetIpv6EgressRulesArgs{
//				Ipv6GatewayId: "example_value",
//				Status:        pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6EgressRuleId3", status.Rules[0].Id)
//			ipv6EgressRuleName, err := vpc.GetIpv6EgressRules(ctx, &vpc.GetIpv6EgressRulesArgs{
//				Ipv6GatewayId:      "example_value",
//				Ipv6EgressRuleName: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcIpv6EgressRuleId4", ipv6EgressRuleName.Rules[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetIpv6EgressRules(ctx *pulumi.Context, args *GetIpv6EgressRulesArgs, opts ...pulumi.InvokeOption) (*GetIpv6EgressRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpv6EgressRulesResult
	err := ctx.Invoke("alicloud:vpc/getIpv6EgressRules:getIpv6EgressRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv6EgressRules.
type GetIpv6EgressRulesArgs struct {
	// A list of Ipv6 Egress Rule IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the instance to which the egress-only rule is applied.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the resource.
	Ipv6EgressRuleName *string `pulumi:"ipv6EgressRuleName"`
	// The ID of the IPv6 gateway.
	Ipv6GatewayId string `pulumi:"ipv6GatewayId"`
	// A regex string to filter results by Ipv6 Egress Rule name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getIpv6EgressRules.
type GetIpv6EgressRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                   `pulumi:"id"`
	Ids                []string                 `pulumi:"ids"`
	InstanceId         *string                  `pulumi:"instanceId"`
	Ipv6EgressRuleName *string                  `pulumi:"ipv6EgressRuleName"`
	Ipv6GatewayId      string                   `pulumi:"ipv6GatewayId"`
	NameRegex          *string                  `pulumi:"nameRegex"`
	Names              []string                 `pulumi:"names"`
	OutputFile         *string                  `pulumi:"outputFile"`
	Rules              []GetIpv6EgressRulesRule `pulumi:"rules"`
	Status             *string                  `pulumi:"status"`
}

func GetIpv6EgressRulesOutput(ctx *pulumi.Context, args GetIpv6EgressRulesOutputArgs, opts ...pulumi.InvokeOption) GetIpv6EgressRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpv6EgressRulesResult, error) {
			args := v.(GetIpv6EgressRulesArgs)
			r, err := GetIpv6EgressRules(ctx, &args, opts...)
			var s GetIpv6EgressRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpv6EgressRulesResultOutput)
}

// A collection of arguments for invoking getIpv6EgressRules.
type GetIpv6EgressRulesOutputArgs struct {
	// A list of Ipv6 Egress Rule IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the instance to which the egress-only rule is applied.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the resource.
	Ipv6EgressRuleName pulumi.StringPtrInput `pulumi:"ipv6EgressRuleName"`
	// The ID of the IPv6 gateway.
	Ipv6GatewayId pulumi.StringInput `pulumi:"ipv6GatewayId"`
	// A regex string to filter results by Ipv6 Egress Rule name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetIpv6EgressRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpv6EgressRulesArgs)(nil)).Elem()
}

// A collection of values returned by getIpv6EgressRules.
type GetIpv6EgressRulesResultOutput struct{ *pulumi.OutputState }

func (GetIpv6EgressRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpv6EgressRulesResult)(nil)).Elem()
}

func (o GetIpv6EgressRulesResultOutput) ToGetIpv6EgressRulesResultOutput() GetIpv6EgressRulesResultOutput {
	return o
}

func (o GetIpv6EgressRulesResultOutput) ToGetIpv6EgressRulesResultOutputWithContext(ctx context.Context) GetIpv6EgressRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpv6EgressRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpv6EgressRulesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetIpv6EgressRulesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetIpv6EgressRulesResultOutput) Ipv6EgressRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) *string { return v.Ipv6EgressRuleName }).(pulumi.StringPtrOutput)
}

func (o GetIpv6EgressRulesResultOutput) Ipv6GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) string { return v.Ipv6GatewayId }).(pulumi.StringOutput)
}

func (o GetIpv6EgressRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetIpv6EgressRulesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetIpv6EgressRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetIpv6EgressRulesResultOutput) Rules() GetIpv6EgressRulesRuleArrayOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) []GetIpv6EgressRulesRule { return v.Rules }).(GetIpv6EgressRulesRuleArrayOutput)
}

func (o GetIpv6EgressRulesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIpv6EgressRulesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpv6EgressRulesResultOutput{})
}
