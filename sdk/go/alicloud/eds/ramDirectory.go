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

// Provides a ECD Ram Directory resource.
//
// For information about ECD Ram Directory and how to use it, see [What is Ram Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createramdirectory).
//
// > **NOTE:** Available since v1.174.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			_default, err := eds.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      pulumi.String(_default.Ids[0]),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eds.NewRamDirectory(ctx, "default", &eds.RamDirectoryArgs{
//				DesktopAccessType: pulumi.String("INTERNET"),
//				EnableAdminAccess: pulumi.Bool(true),
//				RamDirectoryName:  pulumi.String(name),
//				VswitchIds: pulumi.StringArray{
//					defaultSwitch.ID(),
//				},
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
// ECD Ram Directory can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:eds/ramDirectory:RamDirectory example <id>
// ```
type RamDirectory struct {
	pulumi.CustomResourceState

	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringOutput `pulumi:"desktopAccessType"`
	// Whether to enable public network access.
	EnableAdminAccess pulumi.BoolOutput `pulumi:"enableAdminAccess"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableInternetAccess pulumi.BoolOutput `pulumi:"enableInternetAccess"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	RamDirectoryName pulumi.StringOutput `pulumi:"ramDirectoryName"`
	// The status of directory.
	Status pulumi.StringOutput `pulumi:"status"`
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayOutput `pulumi:"vswitchIds"`
}

// NewRamDirectory registers a new resource with the given unique name, arguments, and options.
func NewRamDirectory(ctx *pulumi.Context,
	name string, args *RamDirectoryArgs, opts ...pulumi.ResourceOption) (*RamDirectory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RamDirectoryName == nil {
		return nil, errors.New("invalid value for required argument 'RamDirectoryName'")
	}
	if args.VswitchIds == nil {
		return nil, errors.New("invalid value for required argument 'VswitchIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RamDirectory
	err := ctx.RegisterResource("alicloud:eds/ramDirectory:RamDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRamDirectory gets an existing RamDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRamDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RamDirectoryState, opts ...pulumi.ResourceOption) (*RamDirectory, error) {
	var resource RamDirectory
	err := ctx.ReadResource("alicloud:eds/ramDirectory:RamDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RamDirectory resources.
type ramDirectoryState struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// Whether to enable public network access.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableInternetAccess *bool `pulumi:"enableInternetAccess"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	RamDirectoryName *string `pulumi:"ramDirectoryName"`
	// The status of directory.
	Status *string `pulumi:"status"`
	// List of VSwitch IDs in the directory.
	VswitchIds []string `pulumi:"vswitchIds"`
}

type RamDirectoryState struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringPtrInput
	// Whether to enable public network access.
	EnableAdminAccess pulumi.BoolPtrInput
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableInternetAccess pulumi.BoolPtrInput
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	RamDirectoryName pulumi.StringPtrInput
	// The status of directory.
	Status pulumi.StringPtrInput
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayInput
}

func (RamDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*ramDirectoryState)(nil)).Elem()
}

type ramDirectoryArgs struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// Whether to enable public network access.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableInternetAccess *bool `pulumi:"enableInternetAccess"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	RamDirectoryName string `pulumi:"ramDirectoryName"`
	// List of VSwitch IDs in the directory.
	VswitchIds []string `pulumi:"vswitchIds"`
}

// The set of arguments for constructing a RamDirectory resource.
type RamDirectoryArgs struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringPtrInput
	// Whether to enable public network access.
	EnableAdminAccess pulumi.BoolPtrInput
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableInternetAccess pulumi.BoolPtrInput
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	RamDirectoryName pulumi.StringInput
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayInput
}

func (RamDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ramDirectoryArgs)(nil)).Elem()
}

type RamDirectoryInput interface {
	pulumi.Input

	ToRamDirectoryOutput() RamDirectoryOutput
	ToRamDirectoryOutputWithContext(ctx context.Context) RamDirectoryOutput
}

func (*RamDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((**RamDirectory)(nil)).Elem()
}

func (i *RamDirectory) ToRamDirectoryOutput() RamDirectoryOutput {
	return i.ToRamDirectoryOutputWithContext(context.Background())
}

func (i *RamDirectory) ToRamDirectoryOutputWithContext(ctx context.Context) RamDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RamDirectoryOutput)
}

// RamDirectoryArrayInput is an input type that accepts RamDirectoryArray and RamDirectoryArrayOutput values.
// You can construct a concrete instance of `RamDirectoryArrayInput` via:
//
//	RamDirectoryArray{ RamDirectoryArgs{...} }
type RamDirectoryArrayInput interface {
	pulumi.Input

	ToRamDirectoryArrayOutput() RamDirectoryArrayOutput
	ToRamDirectoryArrayOutputWithContext(context.Context) RamDirectoryArrayOutput
}

type RamDirectoryArray []RamDirectoryInput

func (RamDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RamDirectory)(nil)).Elem()
}

func (i RamDirectoryArray) ToRamDirectoryArrayOutput() RamDirectoryArrayOutput {
	return i.ToRamDirectoryArrayOutputWithContext(context.Background())
}

func (i RamDirectoryArray) ToRamDirectoryArrayOutputWithContext(ctx context.Context) RamDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RamDirectoryArrayOutput)
}

// RamDirectoryMapInput is an input type that accepts RamDirectoryMap and RamDirectoryMapOutput values.
// You can construct a concrete instance of `RamDirectoryMapInput` via:
//
//	RamDirectoryMap{ "key": RamDirectoryArgs{...} }
type RamDirectoryMapInput interface {
	pulumi.Input

	ToRamDirectoryMapOutput() RamDirectoryMapOutput
	ToRamDirectoryMapOutputWithContext(context.Context) RamDirectoryMapOutput
}

type RamDirectoryMap map[string]RamDirectoryInput

func (RamDirectoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RamDirectory)(nil)).Elem()
}

func (i RamDirectoryMap) ToRamDirectoryMapOutput() RamDirectoryMapOutput {
	return i.ToRamDirectoryMapOutputWithContext(context.Background())
}

func (i RamDirectoryMap) ToRamDirectoryMapOutputWithContext(ctx context.Context) RamDirectoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RamDirectoryMapOutput)
}

type RamDirectoryOutput struct{ *pulumi.OutputState }

func (RamDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RamDirectory)(nil)).Elem()
}

func (o RamDirectoryOutput) ToRamDirectoryOutput() RamDirectoryOutput {
	return o
}

func (o RamDirectoryOutput) ToRamDirectoryOutputWithContext(ctx context.Context) RamDirectoryOutput {
	return o
}

// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
func (o RamDirectoryOutput) DesktopAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.StringOutput { return v.DesktopAccessType }).(pulumi.StringOutput)
}

// Whether to enable public network access.
func (o RamDirectoryOutput) EnableAdminAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.BoolOutput { return v.EnableAdminAccess }).(pulumi.BoolOutput)
}

// Whether to grant local administrator rights to users who use cloud desktops.
func (o RamDirectoryOutput) EnableInternetAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.BoolOutput { return v.EnableInternetAccess }).(pulumi.BoolOutput)
}

// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
func (o RamDirectoryOutput) RamDirectoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.StringOutput { return v.RamDirectoryName }).(pulumi.StringOutput)
}

// The status of directory.
func (o RamDirectoryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// List of VSwitch IDs in the directory.
func (o RamDirectoryOutput) VswitchIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RamDirectory) pulumi.StringArrayOutput { return v.VswitchIds }).(pulumi.StringArrayOutput)
}

type RamDirectoryArrayOutput struct{ *pulumi.OutputState }

func (RamDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RamDirectory)(nil)).Elem()
}

func (o RamDirectoryArrayOutput) ToRamDirectoryArrayOutput() RamDirectoryArrayOutput {
	return o
}

func (o RamDirectoryArrayOutput) ToRamDirectoryArrayOutputWithContext(ctx context.Context) RamDirectoryArrayOutput {
	return o
}

func (o RamDirectoryArrayOutput) Index(i pulumi.IntInput) RamDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RamDirectory {
		return vs[0].([]*RamDirectory)[vs[1].(int)]
	}).(RamDirectoryOutput)
}

type RamDirectoryMapOutput struct{ *pulumi.OutputState }

func (RamDirectoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RamDirectory)(nil)).Elem()
}

func (o RamDirectoryMapOutput) ToRamDirectoryMapOutput() RamDirectoryMapOutput {
	return o
}

func (o RamDirectoryMapOutput) ToRamDirectoryMapOutputWithContext(ctx context.Context) RamDirectoryMapOutput {
	return o
}

func (o RamDirectoryMapOutput) MapIndex(k pulumi.StringInput) RamDirectoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RamDirectory {
		return vs[0].(map[string]*RamDirectory)[vs[1].(string)]
	}).(RamDirectoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RamDirectoryInput)(nil)).Elem(), &RamDirectory{})
	pulumi.RegisterInputType(reflect.TypeOf((*RamDirectoryArrayInput)(nil)).Elem(), RamDirectoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RamDirectoryMapInput)(nil)).Elem(), RamDirectoryMap{})
	pulumi.RegisterOutputType(RamDirectoryOutput{})
	pulumi.RegisterOutputType(RamDirectoryArrayOutput{})
	pulumi.RegisterOutputType(RamDirectoryMapOutput{})
}
