// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides a list of DMS Enterprise Users in an Alibaba Cloud account according to the specified filters.
//
// > **NOTE:** Available in 1.90.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/dms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "USER"
// 		opt1 := "NORMAL"
// 		dmsEnterpriseUsersDs, err := dms.GetEnterpriseUsers(ctx, &dms.GetEnterpriseUsersArgs{
// 			Ids: []string{
// 				"uid",
// 			},
// 			Role:   &opt0,
// 			Status: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstUserId", dmsEnterpriseUsersDs.Users[0].Id)
// 		return nil
// 	})
// }
// ```
func GetEnterpriseUsers(ctx *pulumi.Context, args *GetEnterpriseUsersArgs, opts ...pulumi.InvokeOption) (*GetEnterpriseUsersResult, error) {
	var rv GetEnterpriseUsersResult
	err := ctx.Invoke("alicloud:dms/getEnterpriseUsers:getEnterpriseUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnterpriseUsers.
type GetEnterpriseUsersArgs struct {
	// A list of DMS Enterprise User IDs (UID).
	Ids []string `pulumi:"ids"`
	// A regex string to filter the results by the DMS Enterprise User nick_name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The role of the user to query.
	Role *string `pulumi:"role"`
	// The keyword used to query users.
	SearchKey *string `pulumi:"searchKey"`
	// The status of the user.
	Status *string `pulumi:"status"`
	// The ID of the tenant in DMS Enterprise.
	Tid *int `pulumi:"tid"`
}

// A collection of values returned by getEnterpriseUsers.
type GetEnterpriseUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of DMS Enterprise User IDs (UID).
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of DMS Enterprise User names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	Role       *string  `pulumi:"role"`
	SearchKey  *string  `pulumi:"searchKey"`
	// The status of the user.
	Status *string `pulumi:"status"`
	Tid    *int    `pulumi:"tid"`
	// A list of DMS Enterprise Users. Each element contains the following attributes:
	Users []GetEnterpriseUsersUser `pulumi:"users"`
}
