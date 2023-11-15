// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Bastionhost User Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := bastionhost.GetUserGroups(ctx, &bastionhost.GetUserGroupsArgs{
//				InstanceId: "bastionhost-cn-xxxx",
//				Ids: []string{
//					"1",
//					"2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bastionhostUserGroupId1", ids.Groups[0].Id)
//			nameRegex, err := bastionhost.GetUserGroups(ctx, &bastionhost.GetUserGroupsArgs{
//				InstanceId: "bastionhost-cn-xxxx",
//				NameRegex:  pulumi.StringRef("^my-UserGroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bastionhostUserGroupId2", nameRegex.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetUserGroups(ctx *pulumi.Context, args *GetUserGroupsArgs, opts ...pulumi.InvokeOption) (*GetUserGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserGroupsResult
	err := ctx.Invoke("alicloud:bastionhost/getUserGroups:getUserGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserGroups.
type GetUserGroupsArgs struct {
	// A list of User Group self IDs.
	Ids []string `pulumi:"ids"`
	// Specify the New Group of the Bastion Host of Instance Id.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by User Group name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Specify the New Group Name. Supports up to 128 Characters.
	UserGroupName *string `pulumi:"userGroupName"`
}

// A collection of values returned by getUserGroups.
type GetUserGroupsResult struct {
	Groups []GetUserGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	InstanceId    string   `pulumi:"instanceId"`
	NameRegex     *string  `pulumi:"nameRegex"`
	Names         []string `pulumi:"names"`
	OutputFile    *string  `pulumi:"outputFile"`
	UserGroupName *string  `pulumi:"userGroupName"`
}

func GetUserGroupsOutput(ctx *pulumi.Context, args GetUserGroupsOutputArgs, opts ...pulumi.InvokeOption) GetUserGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserGroupsResult, error) {
			args := v.(GetUserGroupsArgs)
			r, err := GetUserGroups(ctx, &args, opts...)
			var s GetUserGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserGroupsResultOutput)
}

// A collection of arguments for invoking getUserGroups.
type GetUserGroupsOutputArgs struct {
	// A list of User Group self IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Specify the New Group of the Bastion Host of Instance Id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by User Group name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Specify the New Group Name. Supports up to 128 Characters.
	UserGroupName pulumi.StringPtrInput `pulumi:"userGroupName"`
}

func (GetUserGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getUserGroups.
type GetUserGroupsResultOutput struct{ *pulumi.OutputState }

func (GetUserGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserGroupsResult)(nil)).Elem()
}

func (o GetUserGroupsResultOutput) ToGetUserGroupsResultOutput() GetUserGroupsResultOutput {
	return o
}

func (o GetUserGroupsResultOutput) ToGetUserGroupsResultOutputWithContext(ctx context.Context) GetUserGroupsResultOutput {
	return o
}

func (o GetUserGroupsResultOutput) Groups() GetUserGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetUserGroupsResult) []GetUserGroupsGroup { return v.Groups }).(GetUserGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetUserGroupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserGroupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetUserGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetUserGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetUserGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUserGroupsResultOutput) UserGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserGroupsResult) *string { return v.UserGroupName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserGroupsResultOutput{})
}
