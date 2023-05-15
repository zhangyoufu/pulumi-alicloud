// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `cassandra.getClusters` data source provides a collection of Cassandra clusters available in Alicloud account.
// Filters support regular expression for the cluster name, ids or tags.
//
// > **NOTE:**  Available in 1.88.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.GetClusters(ctx, &cassandra.GetClustersArgs{
//				NameRegex: pulumi.StringRef("tf_testAccCassandra"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	var rv GetClustersResult
	err := ctx.Invoke("alicloud:cassandra/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// The list of Cassandra cluster ids.
	Ids []string `pulumi:"ids"`
	// A regex string to apply to the cluster name.
	NameRegex *string `pulumi:"nameRegex"`
	// The name of file that can save the collection of clusters after running `pulumi preview`.
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// A list of Cassandra clusters. Its every element contains the following attributes:
	Clusters []GetClustersCluster `pulumi:"clusters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of Cassandra cluster ids.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// The name list of Cassandra clusters.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetClustersOutput(ctx *pulumi.Context, args GetClustersOutputArgs, opts ...pulumi.InvokeOption) GetClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClustersResult, error) {
			args := v.(GetClustersArgs)
			r, err := GetClusters(ctx, &args, opts...)
			var s GetClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClustersResultOutput)
}

// A collection of arguments for invoking getClusters.
type GetClustersOutputArgs struct {
	// The list of Cassandra cluster ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to apply to the cluster name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The name of file that can save the collection of clusters after running `pulumi preview`.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersArgs)(nil)).Elem()
}

// A collection of values returned by getClusters.
type GetClustersResultOutput struct{ *pulumi.OutputState }

func (GetClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersResult)(nil)).Elem()
}

func (o GetClustersResultOutput) ToGetClustersResultOutput() GetClustersResultOutput {
	return o
}

func (o GetClustersResultOutput) ToGetClustersResultOutputWithContext(ctx context.Context) GetClustersResultOutput {
	return o
}

// A list of Cassandra clusters. Its every element contains the following attributes:
func (o GetClustersResultOutput) Clusters() GetClustersClusterArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []GetClustersCluster { return v.Clusters }).(GetClustersClusterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of Cassandra cluster ids.
func (o GetClustersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The name list of Cassandra clusters.
func (o GetClustersResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClustersResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o GetClustersResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetClustersResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClustersResultOutput{})
}
