// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Nas File System resource.
//
// After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
//
// For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
//
// > **NOTE:** Available in v1.33.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := nas.NewFileSystem(ctx, "foo", &nas.FileSystemArgs{
// 			Description:  pulumi.String("tf-testAccNasConfig"),
// 			EncryptType:  pulumi.Int(1),
// 			ProtocolType: pulumi.String("NFS"),
// 			StorageType:  pulumi.String("Performance"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Nas File System can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:nas/fileSystem:FileSystem foo 1337849c59
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// The File System description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the file system is encrypted.Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
	// Valid values:
	// 0: The file system is not encrypted.
	// 1: The file system is encrypted with a managed secret key.
	EncryptType pulumi.IntPtrOutput `pulumi:"encryptType"`
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType pulumi.StringOutput `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProtocolType == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolType'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	var resource FileSystem
	err := ctx.RegisterResource("alicloud:nas/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("alicloud:nas/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// The File System description.
	Description *string `pulumi:"description"`
	// Whether the file system is encrypted.Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
	// Valid values:
	// 0: The file system is not encrypted.
	// 1: The file system is encrypted with a managed secret key.
	EncryptType *int `pulumi:"encryptType"`
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType *string `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType *string `pulumi:"storageType"`
}

type FileSystemState struct {
	// The File System description.
	Description pulumi.StringPtrInput
	// Whether the file system is encrypted.Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
	// Valid values:
	// 0: The file system is not encrypted.
	// 1: The file system is encrypted with a managed secret key.
	EncryptType pulumi.IntPtrInput
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType pulumi.StringPtrInput
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// The File System description.
	Description *string `pulumi:"description"`
	// Whether the file system is encrypted.Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
	// Valid values:
	// 0: The file system is not encrypted.
	// 1: The file system is encrypted with a managed secret key.
	EncryptType *int `pulumi:"encryptType"`
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType string `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType string `pulumi:"storageType"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// The File System description.
	Description pulumi.StringPtrInput
	// Whether the file system is encrypted.Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.
	// Valid values:
	// 0: The file system is not encrypted.
	// 1: The file system is encrypted with a managed secret key.
	EncryptType pulumi.IntPtrInput
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType pulumi.StringInput
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType pulumi.StringInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

func (i *FileSystem) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

type FileSystemPtrInput interface {
	pulumi.Input

	ToFileSystemPtrOutput() FileSystemPtrOutput
	ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput
}

type fileSystemPtrType FileSystemArgs

func (*fileSystemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (i *fileSystemPtrType) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return i.ToFileSystemPtrOutputWithContext(context.Background())
}

func (i *fileSystemPtrType) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemPtrOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//          FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FileSystem)(nil))
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//          FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FileSystem)(nil))
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct {
	*pulumi.OutputState
}

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystem)(nil))
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o.ToFileSystemPtrOutputWithContext(context.Background())
}

func (o FileSystemOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o.ApplyT(func(v FileSystem) *FileSystem {
		return &v
	}).(FileSystemPtrOutput)
}

type FileSystemPtrOutput struct {
	*pulumi.OutputState
}

func (FileSystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil))
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutput() FileSystemPtrOutput {
	return o
}

func (o FileSystemPtrOutput) ToFileSystemPtrOutputWithContext(ctx context.Context) FileSystemPtrOutput {
	return o
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileSystem)(nil))
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].([]FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FileSystem)(nil))
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FileSystem {
		return vs[0].(map[string]FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemPtrOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
