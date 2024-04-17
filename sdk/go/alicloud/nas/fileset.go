// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Network Attached Storage (NAS) Fileset resource.
//
// For information about Network Attached Storage (NAS) Fileset and how to use it, see [What is Fileset](https://www.alibabacloud.com/help/en/doc-detail/27530.html).
//
// > **NOTE:** Available in v1.153.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := nas.GetZones(ctx, &nas.GetZonesArgs{
//				FileSystemType: pulumi.StringRef("cpfs"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "example", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.17.3.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "example", &vpc.SwitchArgs{
//				VswitchName: pulumi.String("terraform-example"),
//				CidrBlock:   pulumi.String("172.17.3.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      pulumi.String(example.Zones[1].ZoneId),
//			})
//			if err != nil {
//				return err
//			}
//			exampleFileSystem, err := nas.NewFileSystem(ctx, "example", &nas.FileSystemArgs{
//				ProtocolType:   pulumi.String("cpfs"),
//				StorageType:    pulumi.String("advance_200"),
//				FileSystemType: pulumi.String("cpfs"),
//				Capacity:       pulumi.Int(3600),
//				ZoneId:         pulumi.String(example.Zones[1].ZoneId),
//				VpcId:          exampleNetwork.ID(),
//				VswitchId:      exampleSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nas.NewFileset(ctx, "example", &nas.FilesetArgs{
//				FileSystemId:   exampleFileSystem.ID(),
//				Description:    pulumi.String("terraform-example"),
//				FileSystemPath: pulumi.String("/example_path/"),
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
// Network Attached Storage (NAS) Fileset can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:nas/fileset:Fileset example <file_system_id>:<fileset_id>
// ```
type Fileset struct {
	pulumi.CustomResourceState

	// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The path of the fileset.
	FileSystemPath pulumi.StringOutput `pulumi:"fileSystemPath"`
	// The first ID of the resource.
	FilesetId pulumi.StringOutput `pulumi:"filesetId"`
	// The status of the fileset.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewFileset registers a new resource with the given unique name, arguments, and options.
func NewFileset(ctx *pulumi.Context,
	name string, args *FilesetArgs, opts ...pulumi.ResourceOption) (*Fileset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.FileSystemPath == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemPath'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fileset
	err := ctx.RegisterResource("alicloud:nas/fileset:Fileset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileset gets an existing Fileset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilesetState, opts ...pulumi.ResourceOption) (*Fileset, error) {
	var resource Fileset
	err := ctx.ReadResource("alicloud:nas/fileset:Fileset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fileset resources.
type filesetState struct {
	// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The path of the fileset.
	FileSystemPath *string `pulumi:"fileSystemPath"`
	// The first ID of the resource.
	FilesetId *string `pulumi:"filesetId"`
	// The status of the fileset.
	Status *string `pulumi:"status"`
}

type FilesetState struct {
	// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the file system.
	FileSystemId pulumi.StringPtrInput
	// The path of the fileset.
	FileSystemPath pulumi.StringPtrInput
	// The first ID of the resource.
	FilesetId pulumi.StringPtrInput
	// The status of the fileset.
	Status pulumi.StringPtrInput
}

func (FilesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*filesetState)(nil)).Elem()
}

type filesetArgs struct {
	// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId string `pulumi:"fileSystemId"`
	// The path of the fileset.
	FileSystemPath string `pulumi:"fileSystemPath"`
}

// The set of arguments for constructing a Fileset resource.
type FilesetArgs struct {
	// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the file system.
	FileSystemId pulumi.StringInput
	// The path of the fileset.
	FileSystemPath pulumi.StringInput
}

func (FilesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filesetArgs)(nil)).Elem()
}

type FilesetInput interface {
	pulumi.Input

	ToFilesetOutput() FilesetOutput
	ToFilesetOutputWithContext(ctx context.Context) FilesetOutput
}

func (*Fileset) ElementType() reflect.Type {
	return reflect.TypeOf((**Fileset)(nil)).Elem()
}

func (i *Fileset) ToFilesetOutput() FilesetOutput {
	return i.ToFilesetOutputWithContext(context.Background())
}

func (i *Fileset) ToFilesetOutputWithContext(ctx context.Context) FilesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesetOutput)
}

// FilesetArrayInput is an input type that accepts FilesetArray and FilesetArrayOutput values.
// You can construct a concrete instance of `FilesetArrayInput` via:
//
//	FilesetArray{ FilesetArgs{...} }
type FilesetArrayInput interface {
	pulumi.Input

	ToFilesetArrayOutput() FilesetArrayOutput
	ToFilesetArrayOutputWithContext(context.Context) FilesetArrayOutput
}

type FilesetArray []FilesetInput

func (FilesetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fileset)(nil)).Elem()
}

func (i FilesetArray) ToFilesetArrayOutput() FilesetArrayOutput {
	return i.ToFilesetArrayOutputWithContext(context.Background())
}

func (i FilesetArray) ToFilesetArrayOutputWithContext(ctx context.Context) FilesetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesetArrayOutput)
}

// FilesetMapInput is an input type that accepts FilesetMap and FilesetMapOutput values.
// You can construct a concrete instance of `FilesetMapInput` via:
//
//	FilesetMap{ "key": FilesetArgs{...} }
type FilesetMapInput interface {
	pulumi.Input

	ToFilesetMapOutput() FilesetMapOutput
	ToFilesetMapOutputWithContext(context.Context) FilesetMapOutput
}

type FilesetMap map[string]FilesetInput

func (FilesetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fileset)(nil)).Elem()
}

func (i FilesetMap) ToFilesetMapOutput() FilesetMapOutput {
	return i.ToFilesetMapOutputWithContext(context.Background())
}

func (i FilesetMap) ToFilesetMapOutputWithContext(ctx context.Context) FilesetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesetMapOutput)
}

type FilesetOutput struct{ *pulumi.OutputState }

func (FilesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fileset)(nil)).Elem()
}

func (o FilesetOutput) ToFilesetOutput() FilesetOutput {
	return o
}

func (o FilesetOutput) ToFilesetOutputWithContext(ctx context.Context) FilesetOutput {
	return o
}

// The description of the Fileset. It must be `2` to `128` characters in length and must start with a letter or Chinese, but cannot start with `https://` or `https://`.
func (o FilesetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fileset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dry run.
func (o FilesetOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Fileset) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the file system.
func (o FilesetOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Fileset) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The path of the fileset.
func (o FilesetOutput) FileSystemPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Fileset) pulumi.StringOutput { return v.FileSystemPath }).(pulumi.StringOutput)
}

// The first ID of the resource.
func (o FilesetOutput) FilesetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Fileset) pulumi.StringOutput { return v.FilesetId }).(pulumi.StringOutput)
}

// The status of the fileset.
func (o FilesetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Fileset) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type FilesetArrayOutput struct{ *pulumi.OutputState }

func (FilesetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fileset)(nil)).Elem()
}

func (o FilesetArrayOutput) ToFilesetArrayOutput() FilesetArrayOutput {
	return o
}

func (o FilesetArrayOutput) ToFilesetArrayOutputWithContext(ctx context.Context) FilesetArrayOutput {
	return o
}

func (o FilesetArrayOutput) Index(i pulumi.IntInput) FilesetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fileset {
		return vs[0].([]*Fileset)[vs[1].(int)]
	}).(FilesetOutput)
}

type FilesetMapOutput struct{ *pulumi.OutputState }

func (FilesetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fileset)(nil)).Elem()
}

func (o FilesetMapOutput) ToFilesetMapOutput() FilesetMapOutput {
	return o
}

func (o FilesetMapOutput) ToFilesetMapOutputWithContext(ctx context.Context) FilesetMapOutput {
	return o
}

func (o FilesetMapOutput) MapIndex(k pulumi.StringInput) FilesetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fileset {
		return vs[0].(map[string]*Fileset)[vs[1].(string)]
	}).(FilesetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilesetInput)(nil)).Elem(), &Fileset{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilesetArrayInput)(nil)).Elem(), FilesetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilesetMapInput)(nil)).Elem(), FilesetMap{})
	pulumi.RegisterOutputType(FilesetOutput{})
	pulumi.RegisterOutputType(FilesetArrayOutput{})
	pulumi.RegisterOutputType(FilesetMapOutput{})
}
