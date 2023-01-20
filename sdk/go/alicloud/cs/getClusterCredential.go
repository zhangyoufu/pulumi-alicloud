// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Container Service cluster credential on Alibaba Cloud.
//
// > **NOTE:** Available in v1.187.0+
//
// > **NOTE:** This datasource can be used on all kinds of ACK clusters, including managed clusters, imported kubernetes clusters, serverless clusters and edge clusters. Please make sure that the target cluster is not in the failed state before using this datasource, since the api server of clusters in the failed state cannot be accessed.
func GetClusterCredential(ctx *pulumi.Context, args *GetClusterCredentialArgs, opts ...pulumi.InvokeOption) (*GetClusterCredentialResult, error) {
	var rv GetClusterCredentialResult
	err := ctx.Invoke("alicloud:cs/getClusterCredential:getClusterCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterCredential.
type GetClusterCredentialArgs struct {
	// The id of target cluster.
	ClusterId  string  `pulumi:"clusterId"`
	OutputFile *string `pulumi:"outputFile"`
	// Automatic expiration time of the returned credential. The valid value between `15` and `4320`, in minutes. When this field is omitted, the expiration time will be determined by the system automatically and the result will be in the attributed field `expiration`.
	TemporaryDurationMinutes *int `pulumi:"temporaryDurationMinutes"`
}

// A collection of values returned by getClusterCredential.
type GetClusterCredentialResult struct {
	// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
	CertificateAuthority GetClusterCredentialCertificateAuthority `pulumi:"certificateAuthority"`
	// The id of target cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name of target cluster.
	ClusterName string `pulumi:"clusterName"`
	// Expiration time of kube config. Format: UTC time in rfc3339.
	Expiration string `pulumi:"expiration"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Sensitive) The kube config to use to authenticate with the cluster.
	KubeConfig               string  `pulumi:"kubeConfig"`
	OutputFile               *string `pulumi:"outputFile"`
	TemporaryDurationMinutes *int    `pulumi:"temporaryDurationMinutes"`
}

func GetClusterCredentialOutput(ctx *pulumi.Context, args GetClusterCredentialOutputArgs, opts ...pulumi.InvokeOption) GetClusterCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterCredentialResult, error) {
			args := v.(GetClusterCredentialArgs)
			r, err := GetClusterCredential(ctx, &args, opts...)
			var s GetClusterCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterCredentialResultOutput)
}

// A collection of arguments for invoking getClusterCredential.
type GetClusterCredentialOutputArgs struct {
	// The id of target cluster.
	ClusterId  pulumi.StringInput    `pulumi:"clusterId"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Automatic expiration time of the returned credential. The valid value between `15` and `4320`, in minutes. When this field is omitted, the expiration time will be determined by the system automatically and the result will be in the attributed field `expiration`.
	TemporaryDurationMinutes pulumi.IntPtrInput `pulumi:"temporaryDurationMinutes"`
}

func (GetClusterCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCredentialArgs)(nil)).Elem()
}

// A collection of values returned by getClusterCredential.
type GetClusterCredentialResultOutput struct{ *pulumi.OutputState }

func (GetClusterCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCredentialResult)(nil)).Elem()
}

func (o GetClusterCredentialResultOutput) ToGetClusterCredentialResultOutput() GetClusterCredentialResultOutput {
	return o
}

func (o GetClusterCredentialResultOutput) ToGetClusterCredentialResultOutputWithContext(ctx context.Context) GetClusterCredentialResultOutput {
	return o
}

// (Available in 1.105.0+) Nested attribute containing certificate authority data for your cluster.
func (o GetClusterCredentialResultOutput) CertificateAuthority() GetClusterCredentialCertificateAuthorityOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) GetClusterCredentialCertificateAuthority {
		return v.CertificateAuthority
	}).(GetClusterCredentialCertificateAuthorityOutput)
}

// The id of target cluster.
func (o GetClusterCredentialResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of target cluster.
func (o GetClusterCredentialResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// Expiration time of kube config. Format: UTC time in rfc3339.
func (o GetClusterCredentialResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Sensitive) The kube config to use to authenticate with the cluster.
func (o GetClusterCredentialResultOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) string { return v.KubeConfig }).(pulumi.StringOutput)
}

func (o GetClusterCredentialResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetClusterCredentialResultOutput) TemporaryDurationMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetClusterCredentialResult) *int { return v.TemporaryDurationMinutes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterCredentialResultOutput{})
}
