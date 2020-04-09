// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package yundun

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDBAuditInstance(ctx *pulumi.Context, args *LookupDBAuditInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDBAuditInstanceResult, error) {
	var rv LookupDBAuditInstanceResult
	err := ctx.Invoke("alicloud:yundun/getDBAuditInstance:getDBAuditInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDBAuditInstance.
type LookupDBAuditInstanceArgs struct {
	DescriptionRegex *string                `pulumi:"descriptionRegex"`
	Ids              []string               `pulumi:"ids"`
	OutputFile       *string                `pulumi:"outputFile"`
	Tags             map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getDBAuditInstance.
type LookupDBAuditInstanceResult struct {
	DescriptionRegex *string  `pulumi:"descriptionRegex"`
	Descriptions     []string `pulumi:"descriptions"`
	// id is the provider-assigned unique ID for this managed resource.
	Id         string                       `pulumi:"id"`
	Ids        []string                     `pulumi:"ids"`
	Instances  []GetDBAuditInstanceInstance `pulumi:"instances"`
	OutputFile *string                      `pulumi:"outputFile"`
	Tags       map[string]interface{}       `pulumi:"tags"`
}
