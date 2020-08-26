// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF datasource to retrieve domains.
//
// For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
//
// > **NOTE:** Available in 1.86.0+ .
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/waf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.GetDomains(ctx, &waf.GetDomainsArgs{
// 			InstanceId: "waf-cf-xxxxx",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	var rv GetDomainsResult
	err := ctx.Invoke("alicloud:waf/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// A list of WAF domain names. Each item is domain name.
	Ids []string `pulumi:"ids"`
	// The Id of waf instance to which waf domain belongs.
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	// A list of Domains. Each element contains the following attributes:
	Domains []GetDomainsDomain `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A list of WAF domain names. Each item is domain name.
	Ids        []string `pulumi:"ids"`
	InstanceId string   `pulumi:"instanceId"`
	OutputFile *string  `pulumi:"outputFile"`
}
