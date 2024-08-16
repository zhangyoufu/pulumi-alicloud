// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Dedicated Host Clusters of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.146.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := ecs.GetEcsDedicatedHostClusters(ctx, &ecs.GetEcsDedicatedHostClustersArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDedicatedHostClusterId1", ids.Clusters[0].Id)
//			nameRegex, err := ecs.GetEcsDedicatedHostClusters(ctx, &ecs.GetEcsDedicatedHostClustersArgs{
//				NameRegex: pulumi.StringRef("^my-DedicatedHostCluster"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDedicatedHostClusterId2", nameRegex.Clusters[0].Id)
//			zoneId, err := ecs.GetEcsDedicatedHostClusters(ctx, &ecs.GetEcsDedicatedHostClustersArgs{
//				ZoneId: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDedicatedHostClusterId3", zoneId.Clusters[0].Id)
//			clusterName, err := ecs.GetEcsDedicatedHostClusters(ctx, &ecs.GetEcsDedicatedHostClustersArgs{
//				DedicatedHostClusterName: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDedicatedHostClusterId4", clusterName.Clusters[0].Id)
//			clusterIds, err := ecs.GetEcsDedicatedHostClusters(ctx, &ecs.GetEcsDedicatedHostClustersArgs{
//				DedicatedHostClusterIds: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecsDedicatedHostClusterId5", clusterIds.Clusters[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEcsDedicatedHostClusters(ctx *pulumi.Context, args *GetEcsDedicatedHostClustersArgs, opts ...pulumi.InvokeOption) (*GetEcsDedicatedHostClustersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEcsDedicatedHostClustersResult
	err := ctx.Invoke("alicloud:ecs/getEcsDedicatedHostClusters:getEcsDedicatedHostClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsDedicatedHostClusters.
type GetEcsDedicatedHostClustersArgs struct {
	// The IDs of dedicated host clusters.
	DedicatedHostClusterIds []string `pulumi:"dedicatedHostClusterIds"`
	// The name of the dedicated host cluster.
	DedicatedHostClusterName *string `pulumi:"dedicatedHostClusterName"`
	// A list of Dedicated Host Cluster IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Dedicated Host Cluster name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The zone ID of the dedicated host cluster.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getEcsDedicatedHostClusters.
type GetEcsDedicatedHostClustersResult struct {
	Clusters                 []GetEcsDedicatedHostClustersCluster `pulumi:"clusters"`
	DedicatedHostClusterIds  []string                             `pulumi:"dedicatedHostClusterIds"`
	DedicatedHostClusterName *string                              `pulumi:"dedicatedHostClusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string            `pulumi:"id"`
	Ids        []string          `pulumi:"ids"`
	NameRegex  *string           `pulumi:"nameRegex"`
	Names      []string          `pulumi:"names"`
	OutputFile *string           `pulumi:"outputFile"`
	Tags       map[string]string `pulumi:"tags"`
	ZoneId     *string           `pulumi:"zoneId"`
}

func GetEcsDedicatedHostClustersOutput(ctx *pulumi.Context, args GetEcsDedicatedHostClustersOutputArgs, opts ...pulumi.InvokeOption) GetEcsDedicatedHostClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsDedicatedHostClustersResult, error) {
			args := v.(GetEcsDedicatedHostClustersArgs)
			r, err := GetEcsDedicatedHostClusters(ctx, &args, opts...)
			var s GetEcsDedicatedHostClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsDedicatedHostClustersResultOutput)
}

// A collection of arguments for invoking getEcsDedicatedHostClusters.
type GetEcsDedicatedHostClustersOutputArgs struct {
	// The IDs of dedicated host clusters.
	DedicatedHostClusterIds pulumi.StringArrayInput `pulumi:"dedicatedHostClusterIds"`
	// The name of the dedicated host cluster.
	DedicatedHostClusterName pulumi.StringPtrInput `pulumi:"dedicatedHostClusterName"`
	// A list of Dedicated Host Cluster IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Dedicated Host Cluster name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The zone ID of the dedicated host cluster.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetEcsDedicatedHostClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDedicatedHostClustersArgs)(nil)).Elem()
}

// A collection of values returned by getEcsDedicatedHostClusters.
type GetEcsDedicatedHostClustersResultOutput struct{ *pulumi.OutputState }

func (GetEcsDedicatedHostClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDedicatedHostClustersResult)(nil)).Elem()
}

func (o GetEcsDedicatedHostClustersResultOutput) ToGetEcsDedicatedHostClustersResultOutput() GetEcsDedicatedHostClustersResultOutput {
	return o
}

func (o GetEcsDedicatedHostClustersResultOutput) ToGetEcsDedicatedHostClustersResultOutputWithContext(ctx context.Context) GetEcsDedicatedHostClustersResultOutput {
	return o
}

func (o GetEcsDedicatedHostClustersResultOutput) Clusters() GetEcsDedicatedHostClustersClusterArrayOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) []GetEcsDedicatedHostClustersCluster { return v.Clusters }).(GetEcsDedicatedHostClustersClusterArrayOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) DedicatedHostClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) []string { return v.DedicatedHostClusterIds }).(pulumi.StringArrayOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) DedicatedHostClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) *string { return v.DedicatedHostClusterName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsDedicatedHostClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetEcsDedicatedHostClustersResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDedicatedHostClustersResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsDedicatedHostClustersResultOutput{})
}
