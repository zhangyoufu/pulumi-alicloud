// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databasegateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Database Gateway Gateways of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.135.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/databasegateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := databasegateway.GetGateways(ctx, &databasegateway.GetGatewaysArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("databaseGatewayGatewayId1", ids.Gateways[0].Id)
//			nameRegex, err := databasegateway.GetGateways(ctx, &databasegateway.GetGatewaysArgs{
//				NameRegex: pulumi.StringRef("^my-Gateway"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("databaseGatewayGatewayId2", nameRegex.Gateways[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetGateways(ctx *pulumi.Context, args *GetGatewaysArgs, opts ...pulumi.InvokeOption) (*GetGatewaysResult, error) {
	var rv GetGatewaysResult
	err := ctx.Invoke("alicloud:databasegateway/getGateways:getGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGateways.
type GetGatewaysArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Gateway IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Gateway name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The search key.
	SearchKey *string `pulumi:"searchKey"`
	// The status of gateway. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getGateways.
type GetGatewaysResult struct {
	EnableDetails *bool                `pulumi:"enableDetails"`
	Gateways      []GetGatewaysGateway `pulumi:"gateways"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	SearchKey  *string  `pulumi:"searchKey"`
	Status     *string  `pulumi:"status"`
}

func GetGatewaysOutput(ctx *pulumi.Context, args GetGatewaysOutputArgs, opts ...pulumi.InvokeOption) GetGatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewaysResult, error) {
			args := v.(GetGatewaysArgs)
			r, err := GetGateways(ctx, &args, opts...)
			var s GetGatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewaysResultOutput)
}

// A collection of arguments for invoking getGateways.
type GetGatewaysOutputArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails pulumi.BoolPtrInput `pulumi:"enableDetails"`
	// A list of Gateway IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Gateway name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The search key.
	SearchKey pulumi.StringPtrInput `pulumi:"searchKey"`
	// The status of gateway. Valid values: `EXCEPTION`, `NEW`, `RUNNING`, `STOPPED`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetGatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getGateways.
type GetGatewaysResultOutput struct{ *pulumi.OutputState }

func (GetGatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysResult)(nil)).Elem()
}

func (o GetGatewaysResultOutput) ToGetGatewaysResultOutput() GetGatewaysResultOutput {
	return o
}

func (o GetGatewaysResultOutput) ToGetGatewaysResultOutputWithContext(ctx context.Context) GetGatewaysResultOutput {
	return o
}

func (o GetGatewaysResultOutput) EnableDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *bool { return v.EnableDetails }).(pulumi.BoolPtrOutput)
}

func (o GetGatewaysResultOutput) Gateways() GetGatewaysGatewayArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []GetGatewaysGateway { return v.Gateways }).(GetGatewaysGatewayArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGatewaysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetGatewaysResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetGatewaysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGatewaysResultOutput) SearchKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.SearchKey }).(pulumi.StringPtrOutput)
}

func (o GetGatewaysResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewaysResultOutput{})
}
