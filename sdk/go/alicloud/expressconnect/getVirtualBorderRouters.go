// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Express Connect Virtual Border Routers of the current Alibaba Cloud user.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := expressconnect.GetVirtualBorderRouters(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("expressConnectVirtualBorderRouterId1", ids.Routers[0].Id)
//			nameRegex, err := expressconnect.GetVirtualBorderRouters(ctx, &expressconnect.GetVirtualBorderRoutersArgs{
//				NameRegex: pulumi.StringRef("^my-VirtualBorderRouter"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("expressConnectVirtualBorderRouterId2", nameRegex.Routers[0].Id)
//			filter, err := expressconnect.GetVirtualBorderRouters(ctx, &expressconnect.GetVirtualBorderRoutersArgs{
//				Filters: []expressconnect.GetVirtualBorderRoutersFilter{
//					expressconnect.GetVirtualBorderRoutersFilter{
//						Key: pulumi.StringRef("PhysicalConnectionId"),
//						Values: []string{
//							"pc-xxxx1",
//						},
//					},
//					expressconnect.GetVirtualBorderRoutersFilter{
//						Key: pulumi.StringRef("VbrId"),
//						Values: []string{
//							"vbr-xxxx1",
//							"vbr-xxxx2",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("expressConnectVirtualBorderRouterId3", filter.Routers[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetVirtualBorderRouters(ctx *pulumi.Context, args *GetVirtualBorderRoutersArgs, opts ...pulumi.InvokeOption) (*GetVirtualBorderRoutersResult, error) {
	var rv GetVirtualBorderRoutersResult
	err := ctx.Invoke("alicloud:expressconnect/getVirtualBorderRouters:getVirtualBorderRouters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualBorderRouters.
type GetVirtualBorderRoutersArgs struct {
	// Custom filter block as described below.
	Filters []GetVirtualBorderRoutersFilter `pulumi:"filters"`
	// A list of Virtual Border Router IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Virtual Border Router name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The VBR state.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getVirtualBorderRouters.
type GetVirtualBorderRoutersResult struct {
	Filters []GetVirtualBorderRoutersFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                          `pulumi:"id"`
	Ids        []string                        `pulumi:"ids"`
	NameRegex  *string                         `pulumi:"nameRegex"`
	Names      []string                        `pulumi:"names"`
	OutputFile *string                         `pulumi:"outputFile"`
	Routers    []GetVirtualBorderRoutersRouter `pulumi:"routers"`
	Status     *string                         `pulumi:"status"`
}

func GetVirtualBorderRoutersOutput(ctx *pulumi.Context, args GetVirtualBorderRoutersOutputArgs, opts ...pulumi.InvokeOption) GetVirtualBorderRoutersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualBorderRoutersResult, error) {
			args := v.(GetVirtualBorderRoutersArgs)
			r, err := GetVirtualBorderRouters(ctx, &args, opts...)
			var s GetVirtualBorderRoutersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualBorderRoutersResultOutput)
}

// A collection of arguments for invoking getVirtualBorderRouters.
type GetVirtualBorderRoutersOutputArgs struct {
	// Custom filter block as described below.
	Filters GetVirtualBorderRoutersFilterArrayInput `pulumi:"filters"`
	// A list of Virtual Border Router IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Virtual Border Router name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The VBR state.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetVirtualBorderRoutersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualBorderRoutersArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualBorderRouters.
type GetVirtualBorderRoutersResultOutput struct{ *pulumi.OutputState }

func (GetVirtualBorderRoutersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualBorderRoutersResult)(nil)).Elem()
}

func (o GetVirtualBorderRoutersResultOutput) ToGetVirtualBorderRoutersResultOutput() GetVirtualBorderRoutersResultOutput {
	return o
}

func (o GetVirtualBorderRoutersResultOutput) ToGetVirtualBorderRoutersResultOutputWithContext(ctx context.Context) GetVirtualBorderRoutersResultOutput {
	return o
}

func (o GetVirtualBorderRoutersResultOutput) Filters() GetVirtualBorderRoutersFilterArrayOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) []GetVirtualBorderRoutersFilter { return v.Filters }).(GetVirtualBorderRoutersFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualBorderRoutersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVirtualBorderRoutersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVirtualBorderRoutersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetVirtualBorderRoutersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetVirtualBorderRoutersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVirtualBorderRoutersResultOutput) Routers() GetVirtualBorderRoutersRouterArrayOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) []GetVirtualBorderRoutersRouter { return v.Routers }).(GetVirtualBorderRoutersRouterArrayOutput)
}

func (o GetVirtualBorderRoutersResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualBorderRoutersResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualBorderRoutersResultOutput{})
}
