// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open CEN Transit Router Service automatically. If the service has been opened, it will return opened.
//
// For information about CEN and how to use it, see [What is CEN](https://www.alibabacloud.com/help/en/doc-detail/59870.htm).
//
// > **NOTE:** Available in v1.139.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.GetTransitRouterService(ctx, &cen.GetTransitRouterServiceArgs{
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
func GetTransitRouterService(ctx *pulumi.Context, args *GetTransitRouterServiceArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterServiceResult, error) {
	var rv GetTransitRouterServiceResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterService:getTransitRouterService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterService.
type GetTransitRouterServiceArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getTransitRouterService.
type GetTransitRouterServiceResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}

func GetTransitRouterServiceOutput(ctx *pulumi.Context, args GetTransitRouterServiceOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterServiceResult, error) {
			args := v.(GetTransitRouterServiceArgs)
			r, err := GetTransitRouterService(ctx, &args, opts...)
			var s GetTransitRouterServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterServiceResultOutput)
}

// A collection of arguments for invoking getTransitRouterService.
type GetTransitRouterServiceOutputArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	Enable pulumi.StringPtrInput `pulumi:"enable"`
}

func (GetTransitRouterServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterServiceArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterService.
type GetTransitRouterServiceResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterServiceResult)(nil)).Elem()
}

func (o GetTransitRouterServiceResultOutput) ToGetTransitRouterServiceResultOutput() GetTransitRouterServiceResultOutput {
	return o
}

func (o GetTransitRouterServiceResultOutput) ToGetTransitRouterServiceResultOutputWithContext(ctx context.Context) GetTransitRouterServiceResultOutput {
	return o
}

func (o GetTransitRouterServiceResultOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterServiceResult) *string { return v.Enable }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current service enable status.
func (o GetTransitRouterServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterServiceResultOutput{})
}
