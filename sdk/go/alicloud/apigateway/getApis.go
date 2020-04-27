// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the apis of the current Alibaba Cloud user.
func GetApis(ctx *pulumi.Context, args *GetApisArgs, opts ...pulumi.InvokeOption) (*GetApisResult, error) {
	var rv GetApisResult
	err := ctx.Invoke("alicloud:apigateway/getApis:getApis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApis.
type GetApisArgs struct {
	// (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.
	ApiId *string `pulumi:"apiId"`
	// ID of the specified group.
	GroupId *string `pulumi:"groupId"`
	// A list of api IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter api gateway apis by name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getApis.
type GetApisResult struct {
	ApiId *string `pulumi:"apiId"`
	// A list of apis. Each element contains the following attributes:
	Apis []GetApisApi `pulumi:"apis"`
	// The group id that the apis belong to.
	GroupId *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of api IDs.
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of api names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
