// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directmail

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Direct Mail Domains of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/directmail"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ids, err := directmail.GetDomains(ctx, &directmail.GetDomainsArgs{
// 			Ids: []string{
// 				"example_id",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("directMailDomainId1", ids.Domains[0].Id)
// 		opt0 := "^my-Domain"
// 		nameRegex, err := directmail.GetDomains(ctx, &directmail.GetDomainsArgs{
// 			NameRegex: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("directMailDomainId2", nameRegex.Domains[0].Id)
// 		opt1 := "1"
// 		opt2 := "^my-Domain"
// 		example, err := directmail.GetDomains(ctx, &directmail.GetDomainsArgs{
// 			Status:  &opt1,
// 			KeyWord: &opt2,
// 			Ids: []string{
// 				"example_id",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("directMailDomainId3", example.Domains[0].Id)
// 		return nil
// 	})
// }
// ```
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	var rv GetDomainsResult
	err := ctx.Invoke("alicloud:directmail/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// Default to `false`. Set it to `true` can output more details about resource attributes.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of Domain IDs.
	Ids []string `pulumi:"ids"`
	// domain, length `1` to `50`, including numbers or capitals or lowercase letters or `.` or `-`
	KeyWord *string `pulumi:"keyWord"`
	// A regex string to filter results by Domain name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the domain name. Valid values:`0` to `4`. `0`:Available, Passed. `1`: Unavailable, No passed. `2`: Available, cname no passed, icp no passed. `3`: Available, icp no passed. `4`: Available, cname no passed.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	Domains       []GetDomainsDomain `pulumi:"domains"`
	EnableDetails *bool              `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	KeyWord    *string  `pulumi:"keyWord"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}
