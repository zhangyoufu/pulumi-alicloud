// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package expressconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Express Connect Grant Rule To Cens of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.196.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/expressconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := expressconnect.GetGrantRuleToCens(ctx, &expressconnect.GetGrantRuleToCensArgs{
//				Ids: []string{
//					"example_id",
//				},
//				InstanceId: "your_vbr_instance_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("expressConnectGrantRuleToCenId0", ids.Cens[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetGrantRuleToCens(ctx *pulumi.Context, args *GetGrantRuleToCensArgs, opts ...pulumi.InvokeOption) (*GetGrantRuleToCensResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGrantRuleToCensResult
	err := ctx.Invoke("alicloud:expressconnect/getGrantRuleToCens:getGrantRuleToCens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGrantRuleToCens.
type GetGrantRuleToCensArgs struct {
	// A list of Grant Rule To Cen IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the VBR.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
}

// A collection of values returned by getGrantRuleToCens.
type GetGrantRuleToCensResult struct {
	// A list of Express Connect Grant Rule To Cens. Each element contains the following attributes:
	Cens []GetGrantRuleToCensCen `pulumi:"cens"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	InstanceId string   `pulumi:"instanceId"`
	OutputFile *string  `pulumi:"outputFile"`
	PageNumber *int     `pulumi:"pageNumber"`
	PageSize   *int     `pulumi:"pageSize"`
}

func GetGrantRuleToCensOutput(ctx *pulumi.Context, args GetGrantRuleToCensOutputArgs, opts ...pulumi.InvokeOption) GetGrantRuleToCensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGrantRuleToCensResult, error) {
			args := v.(GetGrantRuleToCensArgs)
			r, err := GetGrantRuleToCens(ctx, &args, opts...)
			var s GetGrantRuleToCensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGrantRuleToCensResultOutput)
}

// A collection of arguments for invoking getGrantRuleToCens.
type GetGrantRuleToCensOutputArgs struct {
	// A list of Grant Rule To Cen IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the VBR.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
}

func (GetGrantRuleToCensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGrantRuleToCensArgs)(nil)).Elem()
}

// A collection of values returned by getGrantRuleToCens.
type GetGrantRuleToCensResultOutput struct{ *pulumi.OutputState }

func (GetGrantRuleToCensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGrantRuleToCensResult)(nil)).Elem()
}

func (o GetGrantRuleToCensResultOutput) ToGetGrantRuleToCensResultOutput() GetGrantRuleToCensResultOutput {
	return o
}

func (o GetGrantRuleToCensResultOutput) ToGetGrantRuleToCensResultOutputWithContext(ctx context.Context) GetGrantRuleToCensResultOutput {
	return o
}

// A list of Express Connect Grant Rule To Cens. Each element contains the following attributes:
func (o GetGrantRuleToCensResultOutput) Cens() GetGrantRuleToCensCenArrayOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) []GetGrantRuleToCensCen { return v.Cens }).(GetGrantRuleToCensCenArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGrantRuleToCensResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGrantRuleToCensResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGrantRuleToCensResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetGrantRuleToCensResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGrantRuleToCensResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetGrantRuleToCensResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGrantRuleToCensResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGrantRuleToCensResultOutput{})
}
