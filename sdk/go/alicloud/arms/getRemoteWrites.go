// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Arms Remote Writes of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.204.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/arms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := arms.GetRemoteWrites(ctx, &arms.GetRemoteWritesArgs{
//				Ids: []string{
//					"example_id",
//				},
//				ClusterId: "your_cluster_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("armsRemoteWritesId1", ids.RemoteWrites[0].Id)
//			nameRegex, err := arms.GetRemoteWrites(ctx, &arms.GetRemoteWritesArgs{
//				NameRegex: pulumi.StringRef("tf-example"),
//				ClusterId: "your_cluster_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("armsRemoteWritesId2", nameRegex.RemoteWrites[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetRemoteWrites(ctx *pulumi.Context, args *GetRemoteWritesArgs, opts ...pulumi.InvokeOption) (*GetRemoteWritesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRemoteWritesResult
	err := ctx.Invoke("alicloud:arms/getRemoteWrites:getRemoteWrites", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteWrites.
type GetRemoteWritesArgs struct {
	// The ID of the Prometheus instance.
	ClusterId string `pulumi:"clusterId"`
	// A list of Remote Write IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Remote Write name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRemoteWrites.
type GetRemoteWritesResult struct {
	// The ID of the Prometheus instance.
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Remote Write names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of Remote Writes. Each element contains the following attributes:
	RemoteWrites []GetRemoteWritesRemoteWrite `pulumi:"remoteWrites"`
}

func GetRemoteWritesOutput(ctx *pulumi.Context, args GetRemoteWritesOutputArgs, opts ...pulumi.InvokeOption) GetRemoteWritesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRemoteWritesResult, error) {
			args := v.(GetRemoteWritesArgs)
			r, err := GetRemoteWrites(ctx, &args, opts...)
			var s GetRemoteWritesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRemoteWritesResultOutput)
}

// A collection of arguments for invoking getRemoteWrites.
type GetRemoteWritesOutputArgs struct {
	// The ID of the Prometheus instance.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// A list of Remote Write IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Remote Write name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetRemoteWritesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRemoteWritesArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteWrites.
type GetRemoteWritesResultOutput struct{ *pulumi.OutputState }

func (GetRemoteWritesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRemoteWritesResult)(nil)).Elem()
}

func (o GetRemoteWritesResultOutput) ToGetRemoteWritesResultOutput() GetRemoteWritesResultOutput {
	return o
}

func (o GetRemoteWritesResultOutput) ToGetRemoteWritesResultOutputWithContext(ctx context.Context) GetRemoteWritesResultOutput {
	return o
}

// The ID of the Prometheus instance.
func (o GetRemoteWritesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRemoteWritesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRemoteWritesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRemoteWritesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Remote Write names.
func (o GetRemoteWritesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRemoteWritesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A list of Remote Writes. Each element contains the following attributes:
func (o GetRemoteWritesResultOutput) RemoteWrites() GetRemoteWritesRemoteWriteArrayOutput {
	return o.ApplyT(func(v GetRemoteWritesResult) []GetRemoteWritesRemoteWrite { return v.RemoteWrites }).(GetRemoteWritesRemoteWriteArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRemoteWritesResultOutput{})
}
