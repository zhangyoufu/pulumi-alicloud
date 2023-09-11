// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides MountTargets available to the user.
//
// > **NOTE**: Available in 1.35.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := nas.GetMountTargets(ctx, &nas.GetMountTargetsArgs{
//				FileSystemId:    "1a2sc4d",
//				AccessGroupName: pulumi.StringRef("tf-testAccNasConfig"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("theFirstMountTargetDomain", example.Targets[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetMountTargets(ctx *pulumi.Context, args *GetMountTargetsArgs, opts ...pulumi.InvokeOption) (*GetMountTargetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMountTargetsResult
	err := ctx.Invoke("alicloud:nas/getMountTargets:getMountTargets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMountTargets.
type GetMountTargetsArgs struct {
	// Filter results by a specific AccessGroupName.
	AccessGroupName *string `pulumi:"accessGroupName"`
	// The ID of the FileSystem that owns the MountTarget.
	FileSystemId string `pulumi:"fileSystemId"`
	// A list of MountTargetDomain.
	Ids []string `pulumi:"ids"`
	// Field `mountTargetDomain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
	//
	// Deprecated: Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	// Filter results by a specific NetworkType.
	NetworkType *string `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
	Status *string `pulumi:"status"`
	// Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
	Type *string `pulumi:"type"`
	// Filter results by a specific VpcId.
	VpcId *string `pulumi:"vpcId"`
	// Filter results by a specific VSwitchId.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getMountTargets.
type GetMountTargetsResult struct {
	// AccessGroup of The MountTarget.
	AccessGroupName *string `pulumi:"accessGroupName"`
	FileSystemId    string  `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of MountTargetDomain.
	Ids []string `pulumi:"ids"`
	// MountTargetDomain of the MountTarget.
	//
	// Deprecated: Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	// (Available 1.95.0+) NetworkType of The MountTarget.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	// (Available 1.95.0+) The status of the mount target.
	Status *string `pulumi:"status"`
	// A list of MountTargetDomains. Each element contains the following attributes:
	Targets []GetMountTargetsTarget `pulumi:"targets"`
	// Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
	Type *string `pulumi:"type"`
	// VpcId of The MountTarget.
	VpcId *string `pulumi:"vpcId"`
	// VSwitchId of The MountTarget.
	VswitchId *string `pulumi:"vswitchId"`
}

func GetMountTargetsOutput(ctx *pulumi.Context, args GetMountTargetsOutputArgs, opts ...pulumi.InvokeOption) GetMountTargetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMountTargetsResult, error) {
			args := v.(GetMountTargetsArgs)
			r, err := GetMountTargets(ctx, &args, opts...)
			var s GetMountTargetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMountTargetsResultOutput)
}

// A collection of arguments for invoking getMountTargets.
type GetMountTargetsOutputArgs struct {
	// Filter results by a specific AccessGroupName.
	AccessGroupName pulumi.StringPtrInput `pulumi:"accessGroupName"`
	// The ID of the FileSystem that owns the MountTarget.
	FileSystemId pulumi.StringInput `pulumi:"fileSystemId"`
	// A list of MountTargetDomain.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Field `mountTargetDomain` has been deprecated from provider version 1.53.0. New field `ids` replaces it.
	//
	// Deprecated: Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
	MountTargetDomain pulumi.StringPtrInput `pulumi:"mountTargetDomain"`
	// Filter results by a specific NetworkType.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Filter results by the status of mount target. Valid values: `Active`, `Inactive` and `Pending`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Filter results by a specific VpcId.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// Filter results by a specific VSwitchId.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetMountTargetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMountTargetsArgs)(nil)).Elem()
}

// A collection of values returned by getMountTargets.
type GetMountTargetsResultOutput struct{ *pulumi.OutputState }

func (GetMountTargetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMountTargetsResult)(nil)).Elem()
}

func (o GetMountTargetsResultOutput) ToGetMountTargetsResultOutput() GetMountTargetsResultOutput {
	return o
}

func (o GetMountTargetsResultOutput) ToGetMountTargetsResultOutputWithContext(ctx context.Context) GetMountTargetsResultOutput {
	return o
}

func (o GetMountTargetsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetMountTargetsResult] {
	return pulumix.Output[GetMountTargetsResult]{
		OutputState: o.OutputState,
	}
}

// AccessGroup of The MountTarget.
func (o GetMountTargetsResultOutput) AccessGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.AccessGroupName }).(pulumi.StringPtrOutput)
}

func (o GetMountTargetsResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMountTargetsResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMountTargetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMountTargetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of MountTargetDomain.
func (o GetMountTargetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMountTargetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// MountTargetDomain of the MountTarget.
//
// Deprecated: Field 'mount_target_domain' has been deprecated from provider version 1.53.0. New field 'ids' replaces it.
func (o GetMountTargetsResultOutput) MountTargetDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.MountTargetDomain }).(pulumi.StringPtrOutput)
}

// (Available 1.95.0+) NetworkType of The MountTarget.
func (o GetMountTargetsResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetMountTargetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// (Available 1.95.0+) The status of the mount target.
func (o GetMountTargetsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of MountTargetDomains. Each element contains the following attributes:
func (o GetMountTargetsResultOutput) Targets() GetMountTargetsTargetArrayOutput {
	return o.ApplyT(func(v GetMountTargetsResult) []GetMountTargetsTarget { return v.Targets }).(GetMountTargetsTargetArrayOutput)
}

// Field `type` has been deprecated from provider version 1.95.0. New field `networkType` replaces it.
//
// Deprecated: Field 'type' has been deprecated from provider version 1.95.0. New field 'network_type' replaces it.
func (o GetMountTargetsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// VpcId of The MountTarget.
func (o GetMountTargetsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// VSwitchId of The MountTarget.
func (o GetMountTargetsResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMountTargetsResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMountTargetsResultOutput{})
}
