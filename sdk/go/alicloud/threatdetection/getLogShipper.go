// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Using this data source can open Threat Detection Log Shipper automatically. If the service has been enabled, it will return `Opened`.
//
// For information about Threat Detection Log Shipper and how to use it, see [What is Log Shipper](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifyopenlogshipper).
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := threatdetection.GetLogShipper(ctx, &threatdetection.GetLogShipperArgs{
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
func GetLogShipper(ctx *pulumi.Context, args *GetLogShipperArgs, opts ...pulumi.InvokeOption) (*GetLogShipperResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogShipperResult
	err := ctx.Invoke("alicloud:threatdetection/getLogShipper:getLogShipper", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogShipper.
type GetLogShipperArgs struct {
	// Setting the value to `On` to enable the service. Valid values: `On` or `Off`. Default to `Off`.
	//
	// > **NOTE:** Setting `enable = "On"` to open the Threat Detection Log Shipper that means you have read and agreed the [Threat Detection Log Shipper Terms of Service](https://help.aliyun.com/document_detail/170157.html). The service can not closed once it is opened.
	Enable *string `pulumi:"enable"`
}

// A collection of values returned by getLogShipper.
type GetLogShipperResult struct {
	// Log Analysis Service authorization status.
	AuthStatus string `pulumi:"authStatus"`
	// Cloud Security Center purchase status.
	BuyStatus string  `pulumi:"buyStatus"`
	Enable    *string `pulumi:"enable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Log analysis shipping activation status.
	OpenStatus string `pulumi:"openStatus"`
	// Log analysis project status.
	SlsProjectStatus string `pulumi:"slsProjectStatus"`
	// Log Analysis Service is activated.
	SlsServiceStatus string `pulumi:"slsServiceStatus"`
	// The current service enable status.
	Status string `pulumi:"status"`
}

func GetLogShipperOutput(ctx *pulumi.Context, args GetLogShipperOutputArgs, opts ...pulumi.InvokeOption) GetLogShipperResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogShipperResult, error) {
			args := v.(GetLogShipperArgs)
			r, err := GetLogShipper(ctx, &args, opts...)
			var s GetLogShipperResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogShipperResultOutput)
}

// A collection of arguments for invoking getLogShipper.
type GetLogShipperOutputArgs struct {
	// Setting the value to `On` to enable the service. Valid values: `On` or `Off`. Default to `Off`.
	//
	// > **NOTE:** Setting `enable = "On"` to open the Threat Detection Log Shipper that means you have read and agreed the [Threat Detection Log Shipper Terms of Service](https://help.aliyun.com/document_detail/170157.html). The service can not closed once it is opened.
	Enable pulumi.StringPtrInput `pulumi:"enable"`
}

func (GetLogShipperOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogShipperArgs)(nil)).Elem()
}

// A collection of values returned by getLogShipper.
type GetLogShipperResultOutput struct{ *pulumi.OutputState }

func (GetLogShipperResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogShipperResult)(nil)).Elem()
}

func (o GetLogShipperResultOutput) ToGetLogShipperResultOutput() GetLogShipperResultOutput {
	return o
}

func (o GetLogShipperResultOutput) ToGetLogShipperResultOutputWithContext(ctx context.Context) GetLogShipperResultOutput {
	return o
}

// Log Analysis Service authorization status.
func (o GetLogShipperResultOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.AuthStatus }).(pulumi.StringOutput)
}

// Cloud Security Center purchase status.
func (o GetLogShipperResultOutput) BuyStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.BuyStatus }).(pulumi.StringOutput)
}

func (o GetLogShipperResultOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogShipperResult) *string { return v.Enable }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogShipperResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.Id }).(pulumi.StringOutput)
}

// Log analysis shipping activation status.
func (o GetLogShipperResultOutput) OpenStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.OpenStatus }).(pulumi.StringOutput)
}

// Log analysis project status.
func (o GetLogShipperResultOutput) SlsProjectStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.SlsProjectStatus }).(pulumi.StringOutput)
}

// Log Analysis Service is activated.
func (o GetLogShipperResultOutput) SlsServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.SlsServiceStatus }).(pulumi.StringOutput)
}

// The current service enable status.
func (o GetLogShipperResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogShipperResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogShipperResultOutput{})
}
