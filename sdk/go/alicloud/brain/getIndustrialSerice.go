// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package brain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open Brain Industrial service automatically. If the service has been opened, it will return opened.
//
// > **NOTE:** Available in v1.115.0+
//
// > **NOTE:** The Brain Industrial service is not support in the international site.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/brain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := brain.GetIndustrialSerice(ctx, &brain.GetIndustrialSericeArgs{
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
// <!--End PulumiCodeChooser -->
func GetIndustrialSerice(ctx *pulumi.Context, args *GetIndustrialSericeArgs, opts ...pulumi.InvokeOption) (*GetIndustrialSericeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIndustrialSericeResult
	err := ctx.Invoke("alicloud:brain/getIndustrialSerice:getIndustrialSerice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIndustrialSerice.
type GetIndustrialSericeArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	//
	// > **NOTE:** Setting `enable = "On"` to open the Brain Industrial service. The service can not closed once it is opened.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getIndustrialSerice.
type GetIndustrialSericeResult struct {
	Enable *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current service enable status.
	Status string `pulumi:"status"`
}

func GetIndustrialSericeOutput(ctx *pulumi.Context, args GetIndustrialSericeOutputArgs, opts ...pulumi.InvokeOption) GetIndustrialSericeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIndustrialSericeResult, error) {
			args := v.(GetIndustrialSericeArgs)
			r, err := GetIndustrialSerice(ctx, &args, opts...)
			var s GetIndustrialSericeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIndustrialSericeResultOutput)
}

// A collection of arguments for invoking getIndustrialSerice.
type GetIndustrialSericeOutputArgs struct {
	// Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: `On` or `Off`. Default to `Off`.
	//
	// > **NOTE:** Setting `enable = "On"` to open the Brain Industrial service. The service can not closed once it is opened.
	Enable pulumi.StringPtrInput `pulumi:"enable"`
}

func (GetIndustrialSericeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIndustrialSericeArgs)(nil)).Elem()
}

// A collection of values returned by getIndustrialSerice.
type GetIndustrialSericeResultOutput struct{ *pulumi.OutputState }

func (GetIndustrialSericeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIndustrialSericeResult)(nil)).Elem()
}

func (o GetIndustrialSericeResultOutput) ToGetIndustrialSericeResultOutput() GetIndustrialSericeResultOutput {
	return o
}

func (o GetIndustrialSericeResultOutput) ToGetIndustrialSericeResultOutputWithContext(ctx context.Context) GetIndustrialSericeResultOutput {
	return o
}

func (o GetIndustrialSericeResultOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIndustrialSericeResult) *string { return v.Enable }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIndustrialSericeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIndustrialSericeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The current service enable status.
func (o GetIndustrialSericeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetIndustrialSericeResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIndustrialSericeResultOutput{})
}
