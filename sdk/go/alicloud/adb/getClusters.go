// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package adb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `adb.getClusters` data source provides a collection of ADB clusters available in Alibaba Cloud account.
// Filters support regular expression for the cluster description, searches by tags, and other filters which are listed below.
//
// > **NOTE:** Available in v1.71.0+.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/adb_clusters.html.markdown.
func GetClusters(ctx *pulumi.Context, args *GetClustersArgs, opts ...pulumi.InvokeOption) (*GetClustersResult, error) {
	var rv GetClustersResult
	err := ctx.Invoke("alicloud:adb/getClusters:getClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusters.
type GetClustersArgs struct {
	// A regex string to filter results by cluster description.
	DescriptionRegex *string `pulumi:"descriptionRegex"`
	// A list of ADB cluster IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getClusters.
type GetClustersResult struct {
	// A list of ADB clusters. Each element contains the following attributes:
	Clusters         []GetClustersCluster `pulumi:"clusters"`
	DescriptionRegex *string              `pulumi:"descriptionRegex"`
	// A list of ADB cluster descriptions.
	Descriptions []string `pulumi:"descriptions"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of ADB cluster IDs.
	Ids        []string               `pulumi:"ids"`
	OutputFile *string                `pulumi:"outputFile"`
	Tags       map[string]interface{} `pulumi:"tags"`
}
