// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `emr.getMainVersions` data source provides a collection of emr
// main versions available in Alibaba Cloud account when create a emr cluster.
//
// > **NOTE:** Available in 1.59.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/emr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := emr.GetMainVersions(ctx, &emr.GetMainVersionsArgs{
//				ClusterTypes: []string{
//					"HADOOP",
//					"ZOOKEEPER",
//				},
//				EmrVersion: pulumi.StringRef("EMR-3.22.0"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstMainVersion", _default.MainVersions[0].EmrVersion)
//			ctx.Export("thisClusterTypes", _default.MainVersions[0].ClusterTypes)
//			return nil
//		})
//	}
//
// ```
func GetMainVersions(ctx *pulumi.Context, args *GetMainVersionsArgs, opts ...pulumi.InvokeOption) (*GetMainVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMainVersionsResult
	err := ctx.Invoke("alicloud:emr/getMainVersions:getMainVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMainVersions.
type GetMainVersionsArgs struct {
	// The supported clusterType of this emr version.
	// Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]
	ClusterTypes []string `pulumi:"clusterTypes"`
	// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
	EmrVersion *string `pulumi:"emrVersion"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getMainVersions.
type GetMainVersionsResult struct {
	ClusterTypes []string `pulumi:"clusterTypes"`
	// The version of the emr cluster instance.
	EmrVersion *string `pulumi:"emrVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of emr instance types IDs.
	Ids []string `pulumi:"ids"`
	// A list of versions of the emr cluster instance. Each element contains the following attributes:
	MainVersions []GetMainVersionsMainVersion `pulumi:"mainVersions"`
	OutputFile   *string                      `pulumi:"outputFile"`
}

func GetMainVersionsOutput(ctx *pulumi.Context, args GetMainVersionsOutputArgs, opts ...pulumi.InvokeOption) GetMainVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMainVersionsResult, error) {
			args := v.(GetMainVersionsArgs)
			r, err := GetMainVersions(ctx, &args, opts...)
			var s GetMainVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMainVersionsResultOutput)
}

// A collection of arguments for invoking getMainVersions.
type GetMainVersionsOutputArgs struct {
	// The supported clusterType of this emr version.
	// Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]
	ClusterTypes pulumi.StringArrayInput `pulumi:"clusterTypes"`
	// The version of the emr cluster instance. Possible values: `EMR-4.0.0`, `EMR-3.23.0`, `EMR-3.22.0`.
	EmrVersion pulumi.StringPtrInput `pulumi:"emrVersion"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetMainVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMainVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getMainVersions.
type GetMainVersionsResultOutput struct{ *pulumi.OutputState }

func (GetMainVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMainVersionsResult)(nil)).Elem()
}

func (o GetMainVersionsResultOutput) ToGetMainVersionsResultOutput() GetMainVersionsResultOutput {
	return o
}

func (o GetMainVersionsResultOutput) ToGetMainVersionsResultOutputWithContext(ctx context.Context) GetMainVersionsResultOutput {
	return o
}

func (o GetMainVersionsResultOutput) ClusterTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMainVersionsResult) []string { return v.ClusterTypes }).(pulumi.StringArrayOutput)
}

// The version of the emr cluster instance.
func (o GetMainVersionsResultOutput) EmrVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMainVersionsResult) *string { return v.EmrVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMainVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMainVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of emr instance types IDs.
func (o GetMainVersionsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMainVersionsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of versions of the emr cluster instance. Each element contains the following attributes:
func (o GetMainVersionsResultOutput) MainVersions() GetMainVersionsMainVersionArrayOutput {
	return o.ApplyT(func(v GetMainVersionsResult) []GetMainVersionsMainVersion { return v.MainVersions }).(GetMainVersionsMainVersionArrayOutput)
}

func (o GetMainVersionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMainVersionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMainVersionsResultOutput{})
}
