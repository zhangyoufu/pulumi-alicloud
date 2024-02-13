// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECD Nas File System resource.
//
// For information about ECD Nas File System and how to use it, see [What is Nas File System](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-reference-for-easy-use-1).
//
// > **NOTE:** Available since v1.141.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
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
//			_, err := eds.NewSimpleOfficeSite(ctx, "default", &eds.SimpleOfficeSiteArgs{
//				CidrBlock:         pulumi.String("172.16.0.0/12"),
//				EnableAdminAccess: pulumi.Bool(false),
//				DesktopAccessType: pulumi.String("Internet"),
//				OfficeSiteName:    pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eds.NewNasFileSystem(ctx, "example", &eds.NasFileSystemArgs{
//				NasFileSystemName: pulumi.String(name),
//				OfficeSiteId:      _default.ID(),
//				Description:       pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ECD Nas File System can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:eds/nasFileSystem:NasFileSystem example <id>
// ```
type NasFileSystem struct {
	pulumi.CustomResourceState

	// The description of nas file system.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The filesystem id of nas file system.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The domain of mount target.
	MountTargetDomain pulumi.StringOutput `pulumi:"mountTargetDomain"`
	// The name of nas file system.
	NasFileSystemName pulumi.StringPtrOutput `pulumi:"nasFileSystemName"`
	// The ID of office site.
	OfficeSiteId pulumi.StringOutput `pulumi:"officeSiteId"`
	// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
	Reset pulumi.BoolPtrOutput `pulumi:"reset"`
	// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNasFileSystem registers a new resource with the given unique name, arguments, and options.
func NewNasFileSystem(ctx *pulumi.Context,
	name string, args *NasFileSystemArgs, opts ...pulumi.ResourceOption) (*NasFileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OfficeSiteId == nil {
		return nil, errors.New("invalid value for required argument 'OfficeSiteId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NasFileSystem
	err := ctx.RegisterResource("alicloud:eds/nasFileSystem:NasFileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNasFileSystem gets an existing NasFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNasFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NasFileSystemState, opts ...pulumi.ResourceOption) (*NasFileSystem, error) {
	var resource NasFileSystem
	err := ctx.ReadResource("alicloud:eds/nasFileSystem:NasFileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NasFileSystem resources.
type nasFileSystemState struct {
	// The description of nas file system.
	Description *string `pulumi:"description"`
	// The filesystem id of nas file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The domain of mount target.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	// The name of nas file system.
	NasFileSystemName *string `pulumi:"nasFileSystemName"`
	// The ID of office site.
	OfficeSiteId *string `pulumi:"officeSiteId"`
	// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
	Reset *bool `pulumi:"reset"`
	// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
	Status *string `pulumi:"status"`
}

type NasFileSystemState struct {
	// The description of nas file system.
	Description pulumi.StringPtrInput
	// The filesystem id of nas file system.
	FileSystemId pulumi.StringPtrInput
	// The domain of mount target.
	MountTargetDomain pulumi.StringPtrInput
	// The name of nas file system.
	NasFileSystemName pulumi.StringPtrInput
	// The ID of office site.
	OfficeSiteId pulumi.StringPtrInput
	// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
	Reset pulumi.BoolPtrInput
	// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
	Status pulumi.StringPtrInput
}

func (NasFileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*nasFileSystemState)(nil)).Elem()
}

type nasFileSystemArgs struct {
	// The description of nas file system.
	Description *string `pulumi:"description"`
	// The filesystem id of nas file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The domain of mount target.
	MountTargetDomain *string `pulumi:"mountTargetDomain"`
	// The name of nas file system.
	NasFileSystemName *string `pulumi:"nasFileSystemName"`
	// The ID of office site.
	OfficeSiteId string `pulumi:"officeSiteId"`
	// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
	Reset *bool `pulumi:"reset"`
}

// The set of arguments for constructing a NasFileSystem resource.
type NasFileSystemArgs struct {
	// The description of nas file system.
	Description pulumi.StringPtrInput
	// The filesystem id of nas file system.
	FileSystemId pulumi.StringPtrInput
	// The domain of mount target.
	MountTargetDomain pulumi.StringPtrInput
	// The name of nas file system.
	NasFileSystemName pulumi.StringPtrInput
	// The ID of office site.
	OfficeSiteId pulumi.StringInput
	// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
	Reset pulumi.BoolPtrInput
}

func (NasFileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nasFileSystemArgs)(nil)).Elem()
}

type NasFileSystemInput interface {
	pulumi.Input

	ToNasFileSystemOutput() NasFileSystemOutput
	ToNasFileSystemOutputWithContext(ctx context.Context) NasFileSystemOutput
}

func (*NasFileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**NasFileSystem)(nil)).Elem()
}

func (i *NasFileSystem) ToNasFileSystemOutput() NasFileSystemOutput {
	return i.ToNasFileSystemOutputWithContext(context.Background())
}

func (i *NasFileSystem) ToNasFileSystemOutputWithContext(ctx context.Context) NasFileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasFileSystemOutput)
}

// NasFileSystemArrayInput is an input type that accepts NasFileSystemArray and NasFileSystemArrayOutput values.
// You can construct a concrete instance of `NasFileSystemArrayInput` via:
//
//	NasFileSystemArray{ NasFileSystemArgs{...} }
type NasFileSystemArrayInput interface {
	pulumi.Input

	ToNasFileSystemArrayOutput() NasFileSystemArrayOutput
	ToNasFileSystemArrayOutputWithContext(context.Context) NasFileSystemArrayOutput
}

type NasFileSystemArray []NasFileSystemInput

func (NasFileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasFileSystem)(nil)).Elem()
}

func (i NasFileSystemArray) ToNasFileSystemArrayOutput() NasFileSystemArrayOutput {
	return i.ToNasFileSystemArrayOutputWithContext(context.Background())
}

func (i NasFileSystemArray) ToNasFileSystemArrayOutputWithContext(ctx context.Context) NasFileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasFileSystemArrayOutput)
}

// NasFileSystemMapInput is an input type that accepts NasFileSystemMap and NasFileSystemMapOutput values.
// You can construct a concrete instance of `NasFileSystemMapInput` via:
//
//	NasFileSystemMap{ "key": NasFileSystemArgs{...} }
type NasFileSystemMapInput interface {
	pulumi.Input

	ToNasFileSystemMapOutput() NasFileSystemMapOutput
	ToNasFileSystemMapOutputWithContext(context.Context) NasFileSystemMapOutput
}

type NasFileSystemMap map[string]NasFileSystemInput

func (NasFileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasFileSystem)(nil)).Elem()
}

func (i NasFileSystemMap) ToNasFileSystemMapOutput() NasFileSystemMapOutput {
	return i.ToNasFileSystemMapOutputWithContext(context.Background())
}

func (i NasFileSystemMap) ToNasFileSystemMapOutputWithContext(ctx context.Context) NasFileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasFileSystemMapOutput)
}

type NasFileSystemOutput struct{ *pulumi.OutputState }

func (NasFileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NasFileSystem)(nil)).Elem()
}

func (o NasFileSystemOutput) ToNasFileSystemOutput() NasFileSystemOutput {
	return o
}

func (o NasFileSystemOutput) ToNasFileSystemOutputWithContext(ctx context.Context) NasFileSystemOutput {
	return o
}

// The description of nas file system.
func (o NasFileSystemOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The filesystem id of nas file system.
func (o NasFileSystemOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The domain of mount target.
func (o NasFileSystemOutput) MountTargetDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringOutput { return v.MountTargetDomain }).(pulumi.StringOutput)
}

// The name of nas file system.
func (o NasFileSystemOutput) NasFileSystemName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringPtrOutput { return v.NasFileSystemName }).(pulumi.StringPtrOutput)
}

// The ID of office site.
func (o NasFileSystemOutput) OfficeSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringOutput { return v.OfficeSiteId }).(pulumi.StringOutput)
}

// The mount point is in an inactive state, reset the mount point of the NAS file system. Default to `false`.
func (o NasFileSystemOutput) Reset() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.BoolPtrOutput { return v.Reset }).(pulumi.BoolPtrOutput)
}

// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
func (o NasFileSystemOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NasFileSystem) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type NasFileSystemArrayOutput struct{ *pulumi.OutputState }

func (NasFileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasFileSystem)(nil)).Elem()
}

func (o NasFileSystemArrayOutput) ToNasFileSystemArrayOutput() NasFileSystemArrayOutput {
	return o
}

func (o NasFileSystemArrayOutput) ToNasFileSystemArrayOutputWithContext(ctx context.Context) NasFileSystemArrayOutput {
	return o
}

func (o NasFileSystemArrayOutput) Index(i pulumi.IntInput) NasFileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NasFileSystem {
		return vs[0].([]*NasFileSystem)[vs[1].(int)]
	}).(NasFileSystemOutput)
}

type NasFileSystemMapOutput struct{ *pulumi.OutputState }

func (NasFileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasFileSystem)(nil)).Elem()
}

func (o NasFileSystemMapOutput) ToNasFileSystemMapOutput() NasFileSystemMapOutput {
	return o
}

func (o NasFileSystemMapOutput) ToNasFileSystemMapOutputWithContext(ctx context.Context) NasFileSystemMapOutput {
	return o
}

func (o NasFileSystemMapOutput) MapIndex(k pulumi.StringInput) NasFileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NasFileSystem {
		return vs[0].(map[string]*NasFileSystem)[vs[1].(string)]
	}).(NasFileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NasFileSystemInput)(nil)).Elem(), &NasFileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasFileSystemArrayInput)(nil)).Elem(), NasFileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasFileSystemMapInput)(nil)).Elem(), NasFileSystemMap{})
	pulumi.RegisterOutputType(NasFileSystemOutput{})
	pulumi.RegisterOutputType(NasFileSystemArrayOutput{})
	pulumi.RegisterOutputType(NasFileSystemMapOutput{})
}
