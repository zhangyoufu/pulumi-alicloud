// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECS disk resource.
//
// > **DEPRECATED:** This resource has been renamed to ecs.EcsDisk from version 1.122.0.
//
// > **NOTE:** One of `size` or `snapshotId` is required when specifying an ECS disk. If all of them be specified, `size` must more than the size of snapshot which `snapshotId` represents. Currently, `ecs.Disk` doesn't resize disk.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Create a new ECS disk.
//			_, err := ecs.NewDisk(ctx, "ecs_disk", &ecs.DiskArgs{
//				AvailabilityZone: pulumi.String("cn-beijing-b"),
//				Name:             pulumi.String("New-disk"),
//				Description:      pulumi.String("Hello ecs disk."),
//				Category:         pulumi.String("cloud_efficiency"),
//				Size:             pulumi.Int(30),
//				Encrypted:        pulumi.Bool(true),
//				KmsKeyId:         pulumi.String("2a6767f0-a16c-4679-a60f-13bf*****"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("TerraformTest"),
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
//
// ## Import
//
// Cloud disk can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/disk:Disk example d-abc12345678
// ```
type Disk struct {
	pulumi.CustomResourceState

	AdvancedFeatures pulumi.StringPtrOutput `pulumi:"advancedFeatures"`
	// The Zone to create the disk in.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
	DeleteAutoSnapshot pulumi.BoolPtrOutput `pulumi:"deleteAutoSnapshot"`
	// Indicates whether the disk is released together with the instance: Default value: false.
	DeleteWithInstance pulumi.BoolOutput `pulumi:"deleteWithInstance"`
	// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	DiskName    pulumi.StringOutput    `pulumi:"diskName"`
	DryRun      pulumi.BoolPtrOutput   `pulumi:"dryRun"`
	// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
	EnableAutoSnapshot pulumi.BoolOutput      `pulumi:"enableAutoSnapshot"`
	EncryptAlgorithm   pulumi.StringPtrOutput `pulumi:"encryptAlgorithm"`
	// If true, the disk will be encrypted, conflict with `snapshotId`.
	Encrypted  pulumi.BoolPtrOutput `pulumi:"encrypted"`
	InstanceId pulumi.StringOutput  `pulumi:"instanceId"`
	// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
	Name        pulumi.StringOutput `pulumi:"name"`
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
	// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
	// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
	// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	PerformanceLevel pulumi.StringOutput `pulumi:"performanceLevel"`
	// The Id of resource group which the disk belongs.
	// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
	Size pulumi.IntOutput `pulumi:"size"`
	// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// The disk status.
	Status                    pulumi.StringOutput    `pulumi:"status"`
	StorageSetId              pulumi.StringPtrOutput `pulumi:"storageSetId"`
	StorageSetPartitionNumber pulumi.IntPtrOutput    `pulumi:"storageSetPartitionNumber"`
	// A mapping of tags to assign to the resource.
	Tags   pulumi.StringMapOutput `pulumi:"tags"`
	Type   pulumi.StringPtrOutput `pulumi:"type"`
	ZoneId pulumi.StringOutput    `pulumi:"zoneId"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		args = &DiskArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Disk
	err := ctx.RegisterResource("alicloud:ecs/disk:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("alicloud:ecs/disk:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
	AdvancedFeatures *string `pulumi:"advancedFeatures"`
	// The Zone to create the disk in.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
	Category *string `pulumi:"category"`
	// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
	DeleteAutoSnapshot *bool `pulumi:"deleteAutoSnapshot"`
	// Indicates whether the disk is released together with the instance: Default value: false.
	DeleteWithInstance *bool `pulumi:"deleteWithInstance"`
	// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	DiskName    *string `pulumi:"diskName"`
	DryRun      *bool   `pulumi:"dryRun"`
	// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
	EnableAutoSnapshot *bool   `pulumi:"enableAutoSnapshot"`
	EncryptAlgorithm   *string `pulumi:"encryptAlgorithm"`
	// If true, the disk will be encrypted, conflict with `snapshotId`.
	Encrypted  *bool   `pulumi:"encrypted"`
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
	Name        *string `pulumi:"name"`
	PaymentType *string `pulumi:"paymentType"`
	// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
	// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
	// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
	// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	PerformanceLevel *string `pulumi:"performanceLevel"`
	// The Id of resource group which the disk belongs.
	// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
	Size *int `pulumi:"size"`
	// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
	SnapshotId *string `pulumi:"snapshotId"`
	// The disk status.
	Status                    *string `pulumi:"status"`
	StorageSetId              *string `pulumi:"storageSetId"`
	StorageSetPartitionNumber *int    `pulumi:"storageSetPartitionNumber"`
	// A mapping of tags to assign to the resource.
	Tags   map[string]string `pulumi:"tags"`
	Type   *string           `pulumi:"type"`
	ZoneId *string           `pulumi:"zoneId"`
}

type DiskState struct {
	AdvancedFeatures pulumi.StringPtrInput
	// The Zone to create the disk in.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone pulumi.StringPtrInput
	// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
	Category pulumi.StringPtrInput
	// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
	DeleteAutoSnapshot pulumi.BoolPtrInput
	// Indicates whether the disk is released together with the instance: Default value: false.
	DeleteWithInstance pulumi.BoolPtrInput
	// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	DiskName    pulumi.StringPtrInput
	DryRun      pulumi.BoolPtrInput
	// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
	EnableAutoSnapshot pulumi.BoolPtrInput
	EncryptAlgorithm   pulumi.StringPtrInput
	// If true, the disk will be encrypted, conflict with `snapshotId`.
	Encrypted  pulumi.BoolPtrInput
	InstanceId pulumi.StringPtrInput
	// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
	KmsKeyId pulumi.StringPtrInput
	// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
	Name        pulumi.StringPtrInput
	PaymentType pulumi.StringPtrInput
	// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
	// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
	// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
	// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	PerformanceLevel pulumi.StringPtrInput
	// The Id of resource group which the disk belongs.
	// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
	ResourceGroupId pulumi.StringPtrInput
	// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
	Size pulumi.IntPtrInput
	// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
	SnapshotId pulumi.StringPtrInput
	// The disk status.
	Status                    pulumi.StringPtrInput
	StorageSetId              pulumi.StringPtrInput
	StorageSetPartitionNumber pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags   pulumi.StringMapInput
	Type   pulumi.StringPtrInput
	ZoneId pulumi.StringPtrInput
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	AdvancedFeatures *string `pulumi:"advancedFeatures"`
	// The Zone to create the disk in.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
	Category *string `pulumi:"category"`
	// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
	DeleteAutoSnapshot *bool `pulumi:"deleteAutoSnapshot"`
	// Indicates whether the disk is released together with the instance: Default value: false.
	DeleteWithInstance *bool `pulumi:"deleteWithInstance"`
	// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description *string `pulumi:"description"`
	DiskName    *string `pulumi:"diskName"`
	DryRun      *bool   `pulumi:"dryRun"`
	// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
	EnableAutoSnapshot *bool   `pulumi:"enableAutoSnapshot"`
	EncryptAlgorithm   *string `pulumi:"encryptAlgorithm"`
	// If true, the disk will be encrypted, conflict with `snapshotId`.
	Encrypted  *bool   `pulumi:"encrypted"`
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
	Name        *string `pulumi:"name"`
	PaymentType *string `pulumi:"paymentType"`
	// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
	// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
	// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
	// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	PerformanceLevel *string `pulumi:"performanceLevel"`
	// The Id of resource group which the disk belongs.
	// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
	Size *int `pulumi:"size"`
	// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
	SnapshotId                *string `pulumi:"snapshotId"`
	StorageSetId              *string `pulumi:"storageSetId"`
	StorageSetPartitionNumber *int    `pulumi:"storageSetPartitionNumber"`
	// A mapping of tags to assign to the resource.
	Tags   map[string]string `pulumi:"tags"`
	Type   *string           `pulumi:"type"`
	ZoneId *string           `pulumi:"zoneId"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	AdvancedFeatures pulumi.StringPtrInput
	// The Zone to create the disk in.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone pulumi.StringPtrInput
	// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
	Category pulumi.StringPtrInput
	// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
	DeleteAutoSnapshot pulumi.BoolPtrInput
	// Indicates whether the disk is released together with the instance: Default value: false.
	DeleteWithInstance pulumi.BoolPtrInput
	// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description pulumi.StringPtrInput
	DiskName    pulumi.StringPtrInput
	DryRun      pulumi.BoolPtrInput
	// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
	EnableAutoSnapshot pulumi.BoolPtrInput
	EncryptAlgorithm   pulumi.StringPtrInput
	// If true, the disk will be encrypted, conflict with `snapshotId`.
	Encrypted  pulumi.BoolPtrInput
	InstanceId pulumi.StringPtrInput
	// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
	KmsKeyId pulumi.StringPtrInput
	// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
	Name        pulumi.StringPtrInput
	PaymentType pulumi.StringPtrInput
	// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
	// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
	// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
	// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
	PerformanceLevel pulumi.StringPtrInput
	// The Id of resource group which the disk belongs.
	// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
	ResourceGroupId pulumi.StringPtrInput
	// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
	Size pulumi.IntPtrInput
	// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
	SnapshotId                pulumi.StringPtrInput
	StorageSetId              pulumi.StringPtrInput
	StorageSetPartitionNumber pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags   pulumi.StringMapInput
	Type   pulumi.StringPtrInput
	ZoneId pulumi.StringPtrInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

// DiskArrayInput is an input type that accepts DiskArray and DiskArrayOutput values.
// You can construct a concrete instance of `DiskArrayInput` via:
//
//	DiskArray{ DiskArgs{...} }
type DiskArrayInput interface {
	pulumi.Input

	ToDiskArrayOutput() DiskArrayOutput
	ToDiskArrayOutputWithContext(context.Context) DiskArrayOutput
}

type DiskArray []DiskInput

func (DiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Disk)(nil)).Elem()
}

func (i DiskArray) ToDiskArrayOutput() DiskArrayOutput {
	return i.ToDiskArrayOutputWithContext(context.Background())
}

func (i DiskArray) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskArrayOutput)
}

// DiskMapInput is an input type that accepts DiskMap and DiskMapOutput values.
// You can construct a concrete instance of `DiskMapInput` via:
//
//	DiskMap{ "key": DiskArgs{...} }
type DiskMapInput interface {
	pulumi.Input

	ToDiskMapOutput() DiskMapOutput
	ToDiskMapOutputWithContext(context.Context) DiskMapOutput
}

type DiskMap map[string]DiskInput

func (DiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Disk)(nil)).Elem()
}

func (i DiskMap) ToDiskMapOutput() DiskMapOutput {
	return i.ToDiskMapOutputWithContext(context.Background())
}

func (i DiskMap) ToDiskMapOutputWithContext(ctx context.Context) DiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskMapOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) AdvancedFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.AdvancedFeatures }).(pulumi.StringPtrOutput)
}

// The Zone to create the disk in.
//
// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
func (o DiskOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Category of the disk. Valid values are `cloud`, `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloudEssdEntry`. Default is `cloudEfficiency`.
func (o DiskOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Category }).(pulumi.StringPtrOutput)
}

// Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.
func (o DiskOutput) DeleteAutoSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.DeleteAutoSnapshot }).(pulumi.BoolPtrOutput)
}

// Indicates whether the disk is released together with the instance: Default value: false.
func (o DiskOutput) DeleteWithInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolOutput { return v.DeleteWithInstance }).(pulumi.BoolOutput)
}

// Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
func (o DiskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) DiskName() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.DiskName }).(pulumi.StringOutput)
}

func (o DiskOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.
func (o DiskOutput) EnableAutoSnapshot() pulumi.BoolOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolOutput { return v.EnableAutoSnapshot }).(pulumi.BoolOutput)
}

func (o DiskOutput) EncryptAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.EncryptAlgorithm }).(pulumi.StringPtrOutput)
}

// If true, the disk will be encrypted, conflict with `snapshotId`.
func (o DiskOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

func (o DiskOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the KMS key corresponding to the data disk, The specified parameter `Encrypted` must be `true` when KmsKeyId is not empty.
func (o DiskOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'disk_name' instead.
func (o DiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// Specifies the performance level of an ESSD when you create the ESSD. Default value: `PL1`. Valid values:
// * `PL1`: A single ESSD delivers up to 50,000 random read/write IOPS.
// * `PL2`: A single ESSD delivers up to 100,000 random read/write IOPS.
// * `PL3`: A single ESSD delivers up to 1,000,000 random read/write IOPS.
func (o DiskOutput) PerformanceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.PerformanceLevel }).(pulumi.StringOutput)
}

// The Id of resource group which the disk belongs.
// > **NOTE:** Disk category `cloud` has been outdated and it only can be used none I/O Optimized ECS instances. Recommend `cloudEfficiency` and `cloudSsd` disk.
func (o DiskOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error `InvalidDiskSize.TooSmall`.
func (o DiskOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// A snapshot to base the disk off of. If the disk size required by snapshot is greater than `size`, the `size` will be ignored, conflict with `encrypted`.
func (o DiskOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// The disk status.
func (o DiskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DiskOutput) StorageSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.StorageSetId }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) StorageSetPartitionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.StorageSetPartitionNumber }).(pulumi.IntPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o DiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type DiskArrayOutput struct{ *pulumi.OutputState }

func (DiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Disk)(nil)).Elem()
}

func (o DiskArrayOutput) ToDiskArrayOutput() DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) ToDiskArrayOutputWithContext(ctx context.Context) DiskArrayOutput {
	return o
}

func (o DiskArrayOutput) Index(i pulumi.IntInput) DiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Disk {
		return vs[0].([]*Disk)[vs[1].(int)]
	}).(DiskOutput)
}

type DiskMapOutput struct{ *pulumi.OutputState }

func (DiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Disk)(nil)).Elem()
}

func (o DiskMapOutput) ToDiskMapOutput() DiskMapOutput {
	return o
}

func (o DiskMapOutput) ToDiskMapOutputWithContext(ctx context.Context) DiskMapOutput {
	return o
}

func (o DiskMapOutput) MapIndex(k pulumi.StringInput) DiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Disk {
		return vs[0].(map[string]*Disk)[vs[1].(string)]
	}).(DiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskInput)(nil)).Elem(), &Disk{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskArrayInput)(nil)).Elem(), DiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskMapInput)(nil)).Elem(), DiskMap{})
	pulumi.RegisterOutputType(DiskOutput{})
	pulumi.RegisterOutputType(DiskArrayOutput{})
	pulumi.RegisterOutputType(DiskMapOutput{})
}
