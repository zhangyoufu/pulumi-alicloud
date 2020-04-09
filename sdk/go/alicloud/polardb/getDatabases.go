// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package polardb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `polardb.getDatabases` data source provides a collection of PolarDB cluster database available in Alibaba Cloud account.
// Filters support regular expression for the database name, searches by clusterId.
//
// > **NOTE:** Available in v1.70.0+.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/polardb_databases.html.markdown.
func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	var rv GetDatabasesResult
	err := ctx.Invoke("alicloud:polardb/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	// The polarDB cluster ID.
	DbClusterId string `pulumi:"dbClusterId"`
	// A regex string to filter results by database name.
	NameRegex *string `pulumi:"nameRegex"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	// A list of PolarDB cluster databases. Each element contains the following attributes:
	Databases   []GetDatabasesDatabase `pulumi:"databases"`
	DbClusterId string                 `pulumi:"dbClusterId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	NameRegex *string `pulumi:"nameRegex"`
	// database name of the cluster.
	Names []string `pulumi:"names"`
}
