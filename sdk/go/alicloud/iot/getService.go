// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open IOT service automatically. If the service has been opened, it will return opened.
//
// For information about IOT and how to use it, see [What is IOT](https://www.alibabacloud.com/help/en/product/30520.htm).
//
// > **NOTE:** Available in v1.115.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/iot"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iot.GetService(ctx, &iot.GetServiceArgs{
//				Enable: pulumi.StringRef("On"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("alicloud:iot/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}

func GetServiceOutput(ctx *pulumi.Context, args GetServiceOutputArgs, opts ...pulumi.InvokeOption) GetServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceResult, error) {
			args := v.(GetServiceArgs)
			r, err := GetService(ctx, &args, opts...)
			var s GetServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceResultOutput)
}

// A collection of arguments for invoking getService.
type GetServiceOutputArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable pulumi.StringPtrInput `pulumi:"enable"`
}

func (GetServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceArgs)(nil)).Elem()
}

// A collection of values returned by getService.
type GetServiceResultOutput struct{ *pulumi.OutputState }

func (GetServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceResult)(nil)).Elem()
}

func (o GetServiceResultOutput) ToGetServiceResultOutput() GetServiceResultOutput {
	return o
}

func (o GetServiceResultOutput) ToGetServiceResultOutputWithContext(ctx context.Context) GetServiceResultOutput {
	return o
}

func (o GetServiceResultOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceResult) *string { return v.Enable }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current service enable status.
func (o GetServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceResultOutput{})
}
