// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dfs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DFS Mount Point resource.
//
// For information about DFS Mount Point and how to use it, see [What is Mount Point](https://www.alibabacloud.com/help/en/aibaba-cloud-storage-services/latest/apsara-file-storage-for-hdfs).
//
// > **NOTE:** Available since v1.140.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dfs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := dfs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultRandomInteger, err := random.NewRandomInteger(ctx, "defaultRandomInteger", &random.RandomIntegerArgs{
//				Min: pulumi.Int(10000),
//				Max: pulumi.Int(99999),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVPC, err := vpc.NewNetwork(ctx, "defaultVPC", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//				VpcName:   pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVSwitch, err := vpc.NewSwitch(ctx, "defaultVSwitch", &vpc.SwitchArgs{
//				Description: pulumi.String("example"),
//				VpcId:       defaultVPC.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				VswitchName: pulumi.String(name),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].ZoneId),
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccessGroup, err := dfs.NewAccessGroup(ctx, "defaultAccessGroup", &dfs.AccessGroupArgs{
//				Description: pulumi.String("AccessGroup resource manager center example"),
//				NetworkType: pulumi.String("VPC"),
//				AccessGroupName: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("%v-%v", name, result), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dfs.NewAccessGroup(ctx, "updateAccessGroup", &dfs.AccessGroupArgs{
//				Description: pulumi.String("Second AccessGroup resource manager center example"),
//				NetworkType: pulumi.String("VPC"),
//				AccessGroupName: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("%v-update-%v", name, result), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			defaultFs, err := dfs.NewFileSystem(ctx, "defaultFs", &dfs.FileSystemArgs{
//				SpaceCapacity:      pulumi.Int(1024),
//				Description:        pulumi.String("for mountpoint  example"),
//				StorageType:        pulumi.String("STANDARD"),
//				ZoneId:             *pulumi.String(defaultZones.Zones[0].ZoneId),
//				ProtocolType:       pulumi.String("HDFS"),
//				DataRedundancyType: pulumi.String("LRS"),
//				FileSystemName: defaultRandomInteger.Result.ApplyT(func(result int) (string, error) {
//					return fmt.Sprintf("%v-%v", name, result), nil
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dfs.NewMountPoint(ctx, "defaultMountPoint", &dfs.MountPointArgs{
//				VpcId:         defaultVPC.ID(),
//				Description:   pulumi.String("mountpoint example"),
//				NetworkType:   pulumi.String("VPC"),
//				VswitchId:     defaultVSwitch.ID(),
//				FileSystemId:  defaultFs.ID(),
//				AccessGroupId: defaultAccessGroup.ID(),
//				Status:        pulumi.String("Active"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// DFS Mount Point can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dfs/mountPoint:MountPoint example <file_system_id>:<mount_point_id>
// ```
type MountPoint struct {
	pulumi.CustomResourceState

	// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
	AccessGroupId pulumi.StringOutput `pulumi:"accessGroupId"`
	// The mount point alias prefix, which specifies the mount point alias prefix.
	AliasPrefix pulumi.StringPtrOutput `pulumi:"aliasPrefix"`
	// The creation time of the Mount point resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the Mount point.  No more than 32 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique file system identifier, used to retrieve specified file system resources.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
	MountPointId pulumi.StringOutput `pulumi:"mountPointId"`
	// The network type of the Mount point.  Only VPC (VPC) is supported.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewMountPoint registers a new resource with the given unique name, arguments, and options.
func NewMountPoint(ctx *pulumi.Context,
	name string, args *MountPointArgs, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'AccessGroupId'")
	}
	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MountPoint
	err := ctx.RegisterResource("alicloud:dfs/mountPoint:MountPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMountPoint gets an existing MountPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMountPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountPointState, opts ...pulumi.ResourceOption) (*MountPoint, error) {
	var resource MountPoint
	err := ctx.ReadResource("alicloud:dfs/mountPoint:MountPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MountPoint resources.
type mountPointState struct {
	// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
	AccessGroupId *string `pulumi:"accessGroupId"`
	// The mount point alias prefix, which specifies the mount point alias prefix.
	AliasPrefix *string `pulumi:"aliasPrefix"`
	// The creation time of the Mount point resource.
	CreateTime *string `pulumi:"createTime"`
	// The description of the Mount point.  No more than 32 characters in length.
	Description *string `pulumi:"description"`
	// Unique file system identifier, used to retrieve specified file system resources.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
	MountPointId *string `pulumi:"mountPointId"`
	// The network type of the Mount point.  Only VPC (VPC) is supported.
	NetworkType *string `pulumi:"networkType"`
	// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
	Status *string `pulumi:"status"`
	// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
	VpcId *string `pulumi:"vpcId"`
	// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
	VswitchId *string `pulumi:"vswitchId"`
}

type MountPointState struct {
	// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
	AccessGroupId pulumi.StringPtrInput
	// The mount point alias prefix, which specifies the mount point alias prefix.
	AliasPrefix pulumi.StringPtrInput
	// The creation time of the Mount point resource.
	CreateTime pulumi.StringPtrInput
	// The description of the Mount point.  No more than 32 characters in length.
	Description pulumi.StringPtrInput
	// Unique file system identifier, used to retrieve specified file system resources.
	FileSystemId pulumi.StringPtrInput
	// The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
	MountPointId pulumi.StringPtrInput
	// The network type of the Mount point.  Only VPC (VPC) is supported.
	NetworkType pulumi.StringPtrInput
	// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
	Status pulumi.StringPtrInput
	// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
	VpcId pulumi.StringPtrInput
	// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
	VswitchId pulumi.StringPtrInput
}

func (MountPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointState)(nil)).Elem()
}

type mountPointArgs struct {
	// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
	AccessGroupId string `pulumi:"accessGroupId"`
	// The mount point alias prefix, which specifies the mount point alias prefix.
	AliasPrefix *string `pulumi:"aliasPrefix"`
	// The description of the Mount point.  No more than 32 characters in length.
	Description *string `pulumi:"description"`
	// Unique file system identifier, used to retrieve specified file system resources.
	FileSystemId string `pulumi:"fileSystemId"`
	// The network type of the Mount point.  Only VPC (VPC) is supported.
	NetworkType string `pulumi:"networkType"`
	// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
	Status *string `pulumi:"status"`
	// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
	VpcId string `pulumi:"vpcId"`
	// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a MountPoint resource.
type MountPointArgs struct {
	// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
	AccessGroupId pulumi.StringInput
	// The mount point alias prefix, which specifies the mount point alias prefix.
	AliasPrefix pulumi.StringPtrInput
	// The description of the Mount point.  No more than 32 characters in length.
	Description pulumi.StringPtrInput
	// Unique file system identifier, used to retrieve specified file system resources.
	FileSystemId pulumi.StringInput
	// The network type of the Mount point.  Only VPC (VPC) is supported.
	NetworkType pulumi.StringInput
	// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
	Status pulumi.StringPtrInput
	// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
	VpcId pulumi.StringInput
	// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
	VswitchId pulumi.StringInput
}

func (MountPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountPointArgs)(nil)).Elem()
}

type MountPointInput interface {
	pulumi.Input

	ToMountPointOutput() MountPointOutput
	ToMountPointOutputWithContext(ctx context.Context) MountPointOutput
}

func (*MountPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (i *MountPoint) ToMountPointOutput() MountPointOutput {
	return i.ToMountPointOutputWithContext(context.Background())
}

func (i *MountPoint) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointOutput)
}

// MountPointArrayInput is an input type that accepts MountPointArray and MountPointArrayOutput values.
// You can construct a concrete instance of `MountPointArrayInput` via:
//
//	MountPointArray{ MountPointArgs{...} }
type MountPointArrayInput interface {
	pulumi.Input

	ToMountPointArrayOutput() MountPointArrayOutput
	ToMountPointArrayOutputWithContext(context.Context) MountPointArrayOutput
}

type MountPointArray []MountPointInput

func (MountPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (i MountPointArray) ToMountPointArrayOutput() MountPointArrayOutput {
	return i.ToMountPointArrayOutputWithContext(context.Background())
}

func (i MountPointArray) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointArrayOutput)
}

// MountPointMapInput is an input type that accepts MountPointMap and MountPointMapOutput values.
// You can construct a concrete instance of `MountPointMapInput` via:
//
//	MountPointMap{ "key": MountPointArgs{...} }
type MountPointMapInput interface {
	pulumi.Input

	ToMountPointMapOutput() MountPointMapOutput
	ToMountPointMapOutputWithContext(context.Context) MountPointMapOutput
}

type MountPointMap map[string]MountPointInput

func (MountPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (i MountPointMap) ToMountPointMapOutput() MountPointMapOutput {
	return i.ToMountPointMapOutputWithContext(context.Background())
}

func (i MountPointMap) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountPointMapOutput)
}

type MountPointOutput struct{ *pulumi.OutputState }

func (MountPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MountPoint)(nil)).Elem()
}

func (o MountPointOutput) ToMountPointOutput() MountPointOutput {
	return o
}

func (o MountPointOutput) ToMountPointOutputWithContext(ctx context.Context) MountPointOutput {
	return o
}

// The id of the permission group associated with the Mount point, which is used to set the access permissions of the Mount point.
func (o MountPointOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.AccessGroupId }).(pulumi.StringOutput)
}

// The mount point alias prefix, which specifies the mount point alias prefix.
func (o MountPointOutput) AliasPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringPtrOutput { return v.AliasPrefix }).(pulumi.StringPtrOutput)
}

// The creation time of the Mount point resource.
func (o MountPointOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the Mount point.  No more than 32 characters in length.
func (o MountPointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique file system identifier, used to retrieve specified file system resources.
func (o MountPointOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The unique identifier of the Mount point, which is used to retrieve the specified mount point resources.
func (o MountPointOutput) MountPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.MountPointId }).(pulumi.StringOutput)
}

// The network type of the Mount point.  Only VPC (VPC) is supported.
func (o MountPointOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// Mount point status. Value: Inactive: Disable mount points Active: Activate the mount point.
func (o MountPointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the VPC. Specifies the VPC environment to which the mount point belongs.
func (o MountPointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// VSwitch ID, which specifies the VSwitch resource used to create the mount point.
func (o MountPointOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *MountPoint) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type MountPointArrayOutput struct{ *pulumi.OutputState }

func (MountPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MountPoint)(nil)).Elem()
}

func (o MountPointArrayOutput) ToMountPointArrayOutput() MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) ToMountPointArrayOutputWithContext(ctx context.Context) MountPointArrayOutput {
	return o
}

func (o MountPointArrayOutput) Index(i pulumi.IntInput) MountPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].([]*MountPoint)[vs[1].(int)]
	}).(MountPointOutput)
}

type MountPointMapOutput struct{ *pulumi.OutputState }

func (MountPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MountPoint)(nil)).Elem()
}

func (o MountPointMapOutput) ToMountPointMapOutput() MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) ToMountPointMapOutputWithContext(ctx context.Context) MountPointMapOutput {
	return o
}

func (o MountPointMapOutput) MapIndex(k pulumi.StringInput) MountPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MountPoint {
		return vs[0].(map[string]*MountPoint)[vs[1].(string)]
	}).(MountPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointInput)(nil)).Elem(), &MountPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointArrayInput)(nil)).Elem(), MountPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountPointMapInput)(nil)).Elem(), MountPointMap{})
	pulumi.RegisterOutputType(MountPointOutput{})
	pulumi.RegisterOutputType(MountPointArrayOutput{})
	pulumi.RegisterOutputType(MountPointMapOutput{})
}
