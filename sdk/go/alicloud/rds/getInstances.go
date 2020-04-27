// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
// Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:rds/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// `Standard` for standard access mode and `Safe` for high security access mode.
	ConnectionMode *string `pulumi:"connectionMode"`
	// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
	DbType *string `pulumi:"dbType"`
	// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
	Engine *string `pulumi:"engine"`
	// A list of RDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// Status of the instance.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the DB instances.
	// Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
	Tags map[string]interface{} `pulumi:"tags"`
	// Used to retrieve instances belong to specified VPC.
	VpcId *string `pulumi:"vpcId"`
	// Used to retrieve instances belong to specified `vswitch` resources.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// `Standard` for standard access mode and `Safe` for high security access mode.
	ConnectionMode *string `pulumi:"connectionMode"`
	// `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
	DbType *string `pulumi:"dbType"`
	// Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
	Engine *string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of RDS instance IDs.
	Ids []string `pulumi:"ids"`
	// A list of RDS instances. Each element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// A list of RDS instance names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// Status of the instance.
	Status *string                `pulumi:"status"`
	Tags   map[string]interface{} `pulumi:"tags"`
	// ID of the VPC the instance belongs to.
	VpcId *string `pulumi:"vpcId"`
	// ID of the VSwitch the instance belongs to.
	VswitchId *string `pulumi:"vswitchId"`
}
