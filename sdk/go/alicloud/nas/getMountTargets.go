// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides MountTargets available to the user.
//
// > NOTE: Available in 1.35.0+
func GetMountTargets(ctx *pulumi.Context, args *GetMountTargetsArgs, opts ...pulumi.InvokeOption) (*GetMountTargetsResult, error) {
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
	// Filter results by a specific MountTargetDomain.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	OutputFile        *string `pulumi:"outputFile"`
	// Filter results by a specific NetworkType.
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
	// * `type`- NetworkType of The MountTarget.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	OutputFile        *string `pulumi:"outputFile"`
	// A list of MountTargetDomains. Each element contains the following attributes:
	Targets []GetMountTargetsTarget `pulumi:"targets"`
	Type    *string                 `pulumi:"type"`
	// VpcId of The MountTarget.
	VpcId *string `pulumi:"vpcId"`
	// VSwitchId of The MountTarget.
	VswitchId *string `pulumi:"vswitchId"`
}
