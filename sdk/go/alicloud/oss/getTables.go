// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oss

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the ots tables of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.40.0+.
func GetTables(ctx *pulumi.Context, args *GetTablesArgs, opts ...pulumi.InvokeOption) (*GetTablesResult, error) {
	var rv GetTablesResult
	err := ctx.Invoke("alicloud:oss/getTables:getTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTables.
type GetTablesArgs struct {
	// A list of table IDs.
	Ids []string `pulumi:"ids"`
	// The name of OTS instance.
	InstanceName string `pulumi:"instanceName"`
	// A regex string to filter results by table name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getTables.
type GetTablesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of table IDs.
	Ids []string `pulumi:"ids"`
	// The OTS instance name.
	InstanceName string  `pulumi:"instanceName"`
	NameRegex    *string `pulumi:"nameRegex"`
	// A list of table names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of tables. Each element contains the following attributes:
	Tables []GetTablesTable `pulumi:"tables"`
}
