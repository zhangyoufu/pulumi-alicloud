// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides the Cdn Real Time Log Deliveries of the current Alibaba Cloud user.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := cdn.GetRealTimeLogDeliveries(ctx, &cdn.GetRealTimeLogDeliveriesArgs{
//				Domain: "example_value",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cdnRealTimeLogDelivery1", example.Deliveries[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetRealTimeLogDeliveries(ctx *pulumi.Context, args *GetRealTimeLogDeliveriesArgs, opts ...pulumi.InvokeOption) (*GetRealTimeLogDeliveriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRealTimeLogDeliveriesResult
	err := ctx.Invoke("alicloud:cdn/getRealTimeLogDeliveries:getRealTimeLogDeliveries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealTimeLogDeliveries.
type GetRealTimeLogDeliveriesArgs struct {
	// Real-Time Log Service Domain.
	Domain string `pulumi:"domain"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getRealTimeLogDeliveries.
type GetRealTimeLogDeliveriesResult struct {
	Deliveries []GetRealTimeLogDeliveriesDelivery `pulumi:"deliveries"`
	Domain     string                             `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	Status     *string `pulumi:"status"`
}

func GetRealTimeLogDeliveriesOutput(ctx *pulumi.Context, args GetRealTimeLogDeliveriesOutputArgs, opts ...pulumi.InvokeOption) GetRealTimeLogDeliveriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRealTimeLogDeliveriesResult, error) {
			args := v.(GetRealTimeLogDeliveriesArgs)
			r, err := GetRealTimeLogDeliveries(ctx, &args, opts...)
			var s GetRealTimeLogDeliveriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRealTimeLogDeliveriesResultOutput)
}

// A collection of arguments for invoking getRealTimeLogDeliveries.
type GetRealTimeLogDeliveriesOutputArgs struct {
	// Real-Time Log Service Domain.
	Domain pulumi.StringInput `pulumi:"domain"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the real-time log delivery feature. Valid Values: `online` and `offline`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetRealTimeLogDeliveriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealTimeLogDeliveriesArgs)(nil)).Elem()
}

// A collection of values returned by getRealTimeLogDeliveries.
type GetRealTimeLogDeliveriesResultOutput struct{ *pulumi.OutputState }

func (GetRealTimeLogDeliveriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealTimeLogDeliveriesResult)(nil)).Elem()
}

func (o GetRealTimeLogDeliveriesResultOutput) ToGetRealTimeLogDeliveriesResultOutput() GetRealTimeLogDeliveriesResultOutput {
	return o
}

func (o GetRealTimeLogDeliveriesResultOutput) ToGetRealTimeLogDeliveriesResultOutputWithContext(ctx context.Context) GetRealTimeLogDeliveriesResultOutput {
	return o
}

func (o GetRealTimeLogDeliveriesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRealTimeLogDeliveriesResult] {
	return pulumix.Output[GetRealTimeLogDeliveriesResult]{
		OutputState: o.OutputState,
	}
}

func (o GetRealTimeLogDeliveriesResultOutput) Deliveries() GetRealTimeLogDeliveriesDeliveryArrayOutput {
	return o.ApplyT(func(v GetRealTimeLogDeliveriesResult) []GetRealTimeLogDeliveriesDelivery { return v.Deliveries }).(GetRealTimeLogDeliveriesDeliveryArrayOutput)
}

func (o GetRealTimeLogDeliveriesResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetRealTimeLogDeliveriesResult) string { return v.Domain }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRealTimeLogDeliveriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRealTimeLogDeliveriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRealTimeLogDeliveriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRealTimeLogDeliveriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRealTimeLogDeliveriesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRealTimeLogDeliveriesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRealTimeLogDeliveriesResultOutput{})
}
