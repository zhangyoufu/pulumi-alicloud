// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sag

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides Sag Acls available to the user.
//
// > **NOTE:** Available in 1.60.0+
//
// > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
func GetAcls(ctx *pulumi.Context, args *GetAclsArgs, opts ...pulumi.InvokeOption) (*GetAclsResult, error) {
	var rv GetAclsResult
	err := ctx.Invoke("alicloud:sag/getAcls:getAcls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAcls.
type GetAclsArgs struct {
	// A list of Sag Acl IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter Sag Acl instances by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAcls.
type GetAclsResult struct {
	// A list of Sag Acls. Each element contains the following attributes:
	Acls []GetAclsAcl `pulumi:"acls"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of Sag Acl IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Sag Acls names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
