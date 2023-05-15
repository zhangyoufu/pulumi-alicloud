// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Vpc Dhcp Options Sets of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.134.0+.
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
//			ids, err := vpc.GetDhcpOptionsSets(ctx, &vpc.GetDhcpOptionsSetsArgs{
//				Ids: []string{
//					"example_value",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcDhcpOptionsSetId1", ids.Sets[0].Id)
//			nameRegex, err := vpc.GetDhcpOptionsSets(ctx, &vpc.GetDhcpOptionsSetsArgs{
//				NameRegex: pulumi.StringRef("^my-DhcpOptionsSet"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcDhcpOptionsSetId2", nameRegex.Sets[0].Id)
//			dhcpOptionsSetName, err := vpc.GetDhcpOptionsSets(ctx, &vpc.GetDhcpOptionsSetsArgs{
//				DhcpOptionsSetName: pulumi.StringRef("my-DhcpOptionsSet"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcDhcpOptionsSetId3", dhcpOptionsSetName.Sets[0].Id)
//			domainName, err := vpc.GetDhcpOptionsSets(ctx, &vpc.GetDhcpOptionsSetsArgs{
//				Ids: []string{
//					"example_value",
//				},
//				DomainName: pulumi.StringRef("example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcDhcpOptionsSetId4", domainName.Sets[0].Id)
//			status, err := vpc.GetDhcpOptionsSets(ctx, &vpc.GetDhcpOptionsSetsArgs{
//				Ids: []string{
//					"example_value",
//				},
//				Status: pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpcDhcpOptionsSetId5", status.Sets[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetDhcpOptionsSets(ctx *pulumi.Context, args *GetDhcpOptionsSetsArgs, opts ...pulumi.InvokeOption) (*GetDhcpOptionsSetsResult, error) {
	var rv GetDhcpOptionsSetsResult
	err := ctx.Invoke("alicloud:vpc/getDhcpOptionsSets:getDhcpOptionsSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDhcpOptionsSets.
type GetDhcpOptionsSetsArgs struct {
	// The root domain, for example, example.com. After a DHCP options set is associated with a
	// Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the
	// ECS instances in the VPC network.
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual
	// Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS
	// instances in the VPC network.
	DomainName *string `pulumi:"domainName"`
	// A list of Dhcp Options Set IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Dhcp Options Set name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDhcpOptionsSets.
type GetDhcpOptionsSetsResult struct {
	DhcpOptionsSetName *string `pulumi:"dhcpOptionsSetName"`
	DomainName         *string `pulumi:"domainName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                  `pulumi:"id"`
	Ids        []string                `pulumi:"ids"`
	NameRegex  *string                 `pulumi:"nameRegex"`
	Names      []string                `pulumi:"names"`
	OutputFile *string                 `pulumi:"outputFile"`
	Sets       []GetDhcpOptionsSetsSet `pulumi:"sets"`
	Status     *string                 `pulumi:"status"`
}

func GetDhcpOptionsSetsOutput(ctx *pulumi.Context, args GetDhcpOptionsSetsOutputArgs, opts ...pulumi.InvokeOption) GetDhcpOptionsSetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDhcpOptionsSetsResult, error) {
			args := v.(GetDhcpOptionsSetsArgs)
			r, err := GetDhcpOptionsSets(ctx, &args, opts...)
			var s GetDhcpOptionsSetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDhcpOptionsSetsResultOutput)
}

// A collection of arguments for invoking getDhcpOptionsSets.
type GetDhcpOptionsSetsOutputArgs struct {
	// The root domain, for example, example.com. After a DHCP options set is associated with a
	// Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the
	// ECS instances in the VPC network.
	DhcpOptionsSetName pulumi.StringPtrInput `pulumi:"dhcpOptionsSetName"`
	// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual
	// Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS
	// instances in the VPC network.
	DomainName pulumi.StringPtrInput `pulumi:"domainName"`
	// A list of Dhcp Options Set IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Dhcp Options Set name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetDhcpOptionsSetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDhcpOptionsSetsArgs)(nil)).Elem()
}

// A collection of values returned by getDhcpOptionsSets.
type GetDhcpOptionsSetsResultOutput struct{ *pulumi.OutputState }

func (GetDhcpOptionsSetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDhcpOptionsSetsResult)(nil)).Elem()
}

func (o GetDhcpOptionsSetsResultOutput) ToGetDhcpOptionsSetsResultOutput() GetDhcpOptionsSetsResultOutput {
	return o
}

func (o GetDhcpOptionsSetsResultOutput) ToGetDhcpOptionsSetsResultOutputWithContext(ctx context.Context) GetDhcpOptionsSetsResultOutput {
	return o
}

func (o GetDhcpOptionsSetsResultOutput) DhcpOptionsSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) *string { return v.DhcpOptionsSetName }).(pulumi.StringPtrOutput)
}

func (o GetDhcpOptionsSetsResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDhcpOptionsSetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDhcpOptionsSetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetDhcpOptionsSetsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetDhcpOptionsSetsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetDhcpOptionsSetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetDhcpOptionsSetsResultOutput) Sets() GetDhcpOptionsSetsSetArrayOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) []GetDhcpOptionsSetsSet { return v.Sets }).(GetDhcpOptionsSetsSetArrayOutput)
}

func (o GetDhcpOptionsSetsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDhcpOptionsSetsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDhcpOptionsSetsResultOutput{})
}
