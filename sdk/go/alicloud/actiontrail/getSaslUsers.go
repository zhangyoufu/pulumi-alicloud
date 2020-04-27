// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package actiontrail

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of ALIKAFKA Sasl users in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.66.0+
func GetSaslUsers(ctx *pulumi.Context, args *GetSaslUsersArgs, opts ...pulumi.InvokeOption) (*GetSaslUsersResult, error) {
	var rv GetSaslUsersResult
	err := ctx.Invoke("alicloud:actiontrail/getSaslUsers:getSaslUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSaslUsers.
type GetSaslUsersArgs struct {
	// ID of the ALIKAFKA Instance that owns the sasl users.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by the username.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getSaslUsers.
type GetSaslUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	// A list of sasl usernames.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of sasl users. Each element contains the following attributes:
	Users []GetSaslUsersUser `pulumi:"users"`
}
