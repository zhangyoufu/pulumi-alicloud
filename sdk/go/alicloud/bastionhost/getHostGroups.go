// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Bastionhost Host Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.134.0+.
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
//			ids, err := bastionhost.GetHostGroups(ctx, &bastionhost.GetHostGroupsArgs{
//				InstanceId: "bastionhost-cn-tl3xxxxxxx",
//				Ids: []string{
//					"example_value-1",
//					"example_value-2",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bastionhostHostGroupId1", ids.Groups[0].Id)
//			nameRegex, err := bastionhost.GetHostGroups(ctx, &bastionhost.GetHostGroupsArgs{
//				InstanceId: "bastionhost-cn-tl3xxxxxxx",
//				NameRegex:  pulumi.StringRef("^my-HostGroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bastionhostHostGroupId2", nameRegex.Groups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetHostGroups(ctx *pulumi.Context, args *GetHostGroupsArgs, opts ...pulumi.InvokeOption) (*GetHostGroupsResult, error) {
	var rv GetHostGroupsResult
	err := ctx.Invoke("alicloud:bastionhost/getHostGroups:getHostGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostGroups.
type GetHostGroupsArgs struct {
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName *string `pulumi:"hostGroupName"`
	// A list of Host Group IDs.
	Ids []string `pulumi:"ids"`
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId string `pulumi:"instanceId"`
	// A regex string to filter results by Host Group name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getHostGroups.
type GetHostGroupsResult struct {
	Groups        []GetHostGroupsGroup `pulumi:"groups"`
	HostGroupName *string              `pulumi:"hostGroupName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	InstanceId string   `pulumi:"instanceId"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetHostGroupsOutput(ctx *pulumi.Context, args GetHostGroupsOutputArgs, opts ...pulumi.InvokeOption) GetHostGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHostGroupsResult, error) {
			args := v.(GetHostGroupsArgs)
			r, err := GetHostGroups(ctx, &args, opts...)
			var s GetHostGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHostGroupsResultOutput)
}

// A collection of arguments for invoking getHostGroups.
type GetHostGroupsOutputArgs struct {
	// Specify the New Host Group Name, Supports up to 128 Characters.
	HostGroupName pulumi.StringPtrInput `pulumi:"hostGroupName"`
	// A list of Host Group IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Specify the New Host Group Where the Bastion Host ID of.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A regex string to filter results by Host Group name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetHostGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getHostGroups.
type GetHostGroupsResultOutput struct{ *pulumi.OutputState }

func (GetHostGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostGroupsResult)(nil)).Elem()
}

func (o GetHostGroupsResultOutput) ToGetHostGroupsResultOutput() GetHostGroupsResultOutput {
	return o
}

func (o GetHostGroupsResultOutput) ToGetHostGroupsResultOutputWithContext(ctx context.Context) GetHostGroupsResultOutput {
	return o
}

func (o GetHostGroupsResultOutput) Groups() GetHostGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetHostGroupsResult) []GetHostGroupsGroup { return v.Groups }).(GetHostGroupsGroupArrayOutput)
}

func (o GetHostGroupsResultOutput) HostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.HostGroupName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetHostGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHostGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetHostGroupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostGroupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetHostGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetHostGroupsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHostGroupsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetHostGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostGroupsResultOutput{})
}
