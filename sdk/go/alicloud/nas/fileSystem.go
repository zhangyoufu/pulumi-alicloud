// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package nas

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Nas File System resource.
//
// After activating NAS, you can create a file system and purchase a storage package for it in the NAS console. The NAS console also enables you to view the file system details and remove unnecessary file systems.
//
// For information about NAS file system and how to use it, see [Manage file systems](https://www.alibabacloud.com/help/doc-detail/27530.htm)
//
// > **NOTE:** Available in v1.33.0+.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/nas_file_system.html.markdown.
type FileSystem struct {
	pulumi.CustomResourceState

	// The File System description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType pulumi.StringOutput `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil || args.ProtocolType == nil {
		return nil, errors.New("missing required argument 'ProtocolType'")
	}
	if args == nil || args.StorageType == nil {
		return nil, errors.New("missing required argument 'StorageType'")
	}
	if args == nil {
		args = &FileSystemArgs{}
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
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType *string `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType *string `pulumi:"storageType"`
}

type FileSystemState struct {
	// The File System description.
	Description pulumi.StringPtrInput
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
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType string `pulumi:"protocolType"`
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType string `pulumi:"storageType"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// The File System description.
	Description pulumi.StringPtrInput
	// The Protocol Type of a File System. Valid values: `NFS` and `SMB`.
	ProtocolType pulumi.StringInput
	// The Storage Type of a File System. Valid values: `Capacity` and `Performance`.
	StorageType pulumi.StringInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}
