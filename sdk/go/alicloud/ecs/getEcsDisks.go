// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Disks of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.122.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ecs.GetEcsDisks(ctx, &ecs.GetEcsDisksArgs{
//				Ids: []string{
//					"d-artgdsvdvxxxx",
//				},
//				NameRegex: pulumi.StringRef("tf-test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcsDiskId", example.Disks[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEcsDisks(ctx *pulumi.Context, args *GetEcsDisksArgs, opts ...pulumi.InvokeOption) (*GetEcsDisksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEcsDisksResult
	err := ctx.Invoke("alicloud:ecs/getEcsDisks:getEcsDisks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsDisks.
type GetEcsDisksArgs struct {
	// Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
	AdditionalAttributes []string `pulumi:"additionalAttributes"`
	// Query cloud disks based on the automatic snapshot policy ID.
	AutoSnapshotPolicyId *string `pulumi:"autoSnapshotPolicyId"`
	// Availability zone of the disk.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Disk category.
	Category *string `pulumi:"category"`
	// Indicates whether the automatic snapshot is deleted when the disk is released.
	DeleteAutoSnapshot *bool `pulumi:"deleteAutoSnapshot"`
	// Indicates whether the disk is released together with the instance.
	DeleteWithInstance *bool `pulumi:"deleteWithInstance"`
	// The disk name.
	DiskName *string `pulumi:"diskName"`
	// The disk type.
	DiskType *string `pulumi:"diskType"`
	// Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
	DryRun *bool `pulumi:"dryRun"`
	// Whether the disk implements an automatic snapshot policy.
	EnableAutoSnapshot *bool `pulumi:"enableAutoSnapshot"`
	// Whether the disk implements an automatic snapshot policy.
	EnableAutomatedSnapshotPolicy *bool `pulumi:"enableAutomatedSnapshotPolicy"`
	// Whether it is shared block storage.
	EnableShared *bool `pulumi:"enableShared"`
	// Indicate whether the disk is encrypted or not.
	Encrypted *string `pulumi:"encrypted"`
	// A list of Disk IDs.
	Ids []string `pulumi:"ids"`
	// The instance ID of the disk mount.
	InstanceId *string `pulumi:"instanceId"`
	// The kms key id.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A regex string to filter results by Disk name.
	NameRegex      *string                    `pulumi:"nameRegex"`
	OperationLocks []GetEcsDisksOperationLock `pulumi:"operationLocks"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// Payment method for disk.
	PaymentType *string `pulumi:"paymentType"`
	// Whether the disk is unmountable.
	Portable *bool `pulumi:"portable"`
	// The Id of resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
	SnapshotId *string `pulumi:"snapshotId"`
	// Current status.
	Status *string `pulumi:"status"`
	// A map of tags assigned to the disk.
	Tags map[string]interface{} `pulumi:"tags"`
	// Disk type.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
	Type *string `pulumi:"type"`
	// The zone id.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getEcsDisks.
type GetEcsDisksResult struct {
	AdditionalAttributes []string `pulumi:"additionalAttributes"`
	AutoSnapshotPolicyId *string  `pulumi:"autoSnapshotPolicyId"`
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone              *string           `pulumi:"availabilityZone"`
	Category                      *string           `pulumi:"category"`
	DeleteAutoSnapshot            *bool             `pulumi:"deleteAutoSnapshot"`
	DeleteWithInstance            *bool             `pulumi:"deleteWithInstance"`
	DiskName                      *string           `pulumi:"diskName"`
	DiskType                      *string           `pulumi:"diskType"`
	Disks                         []GetEcsDisksDisk `pulumi:"disks"`
	DryRun                        *bool             `pulumi:"dryRun"`
	EnableAutoSnapshot            *bool             `pulumi:"enableAutoSnapshot"`
	EnableAutomatedSnapshotPolicy *bool             `pulumi:"enableAutomatedSnapshotPolicy"`
	EnableShared                  *bool             `pulumi:"enableShared"`
	Encrypted                     *string           `pulumi:"encrypted"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                     `pulumi:"id"`
	Ids             []string                   `pulumi:"ids"`
	InstanceId      *string                    `pulumi:"instanceId"`
	KmsKeyId        *string                    `pulumi:"kmsKeyId"`
	NameRegex       *string                    `pulumi:"nameRegex"`
	Names           []string                   `pulumi:"names"`
	OperationLocks  []GetEcsDisksOperationLock `pulumi:"operationLocks"`
	OutputFile      *string                    `pulumi:"outputFile"`
	PageNumber      *int                       `pulumi:"pageNumber"`
	PageSize        *int                       `pulumi:"pageSize"`
	PaymentType     *string                    `pulumi:"paymentType"`
	Portable        *bool                      `pulumi:"portable"`
	ResourceGroupId *string                    `pulumi:"resourceGroupId"`
	SnapshotId      *string                    `pulumi:"snapshotId"`
	Status          *string                    `pulumi:"status"`
	Tags            map[string]interface{}     `pulumi:"tags"`
	TotalCount      int                        `pulumi:"totalCount"`
	// Deprecated: Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
	Type   *string `pulumi:"type"`
	ZoneId *string `pulumi:"zoneId"`
}

func GetEcsDisksOutput(ctx *pulumi.Context, args GetEcsDisksOutputArgs, opts ...pulumi.InvokeOption) GetEcsDisksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsDisksResult, error) {
			args := v.(GetEcsDisksArgs)
			r, err := GetEcsDisks(ctx, &args, opts...)
			var s GetEcsDisksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsDisksResultOutput)
}

// A collection of arguments for invoking getEcsDisks.
type GetEcsDisksOutputArgs struct {
	// Other attribute values. Currently, only the incoming value of IOPS is supported, which means to query the IOPS upper limit of the current disk.
	AdditionalAttributes pulumi.StringArrayInput `pulumi:"additionalAttributes"`
	// Query cloud disks based on the automatic snapshot policy ID.
	AutoSnapshotPolicyId pulumi.StringPtrInput `pulumi:"autoSnapshotPolicyId"`
	// Availability zone of the disk.
	//
	// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Disk category.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// Indicates whether the automatic snapshot is deleted when the disk is released.
	DeleteAutoSnapshot pulumi.BoolPtrInput `pulumi:"deleteAutoSnapshot"`
	// Indicates whether the disk is released together with the instance.
	DeleteWithInstance pulumi.BoolPtrInput `pulumi:"deleteWithInstance"`
	// The disk name.
	DiskName pulumi.StringPtrInput `pulumi:"diskName"`
	// The disk type.
	DiskType pulumi.StringPtrInput `pulumi:"diskType"`
	// Specifies whether to check the validity of the request without actually making the request.request Default value: false. Valid values:
	DryRun pulumi.BoolPtrInput `pulumi:"dryRun"`
	// Whether the disk implements an automatic snapshot policy.
	EnableAutoSnapshot pulumi.BoolPtrInput `pulumi:"enableAutoSnapshot"`
	// Whether the disk implements an automatic snapshot policy.
	EnableAutomatedSnapshotPolicy pulumi.BoolPtrInput `pulumi:"enableAutomatedSnapshotPolicy"`
	// Whether it is shared block storage.
	EnableShared pulumi.BoolPtrInput `pulumi:"enableShared"`
	// Indicate whether the disk is encrypted or not.
	Encrypted pulumi.StringPtrInput `pulumi:"encrypted"`
	// A list of Disk IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The instance ID of the disk mount.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The kms key id.
	KmsKeyId pulumi.StringPtrInput `pulumi:"kmsKeyId"`
	// A regex string to filter results by Disk name.
	NameRegex      pulumi.StringPtrInput              `pulumi:"nameRegex"`
	OperationLocks GetEcsDisksOperationLockArrayInput `pulumi:"operationLocks"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// Payment method for disk.
	PaymentType pulumi.StringPtrInput `pulumi:"paymentType"`
	// Whether the disk is unmountable.
	Portable pulumi.BoolPtrInput `pulumi:"portable"`
	// The Id of resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// Snapshot used to create the disk. It is null if no snapshot is used to create the disk.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	// Current status.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A map of tags assigned to the disk.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Disk type.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The zone id.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (GetEcsDisksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDisksArgs)(nil)).Elem()
}

// A collection of values returned by getEcsDisks.
type GetEcsDisksResultOutput struct{ *pulumi.OutputState }

func (GetEcsDisksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsDisksResult)(nil)).Elem()
}

func (o GetEcsDisksResultOutput) ToGetEcsDisksResultOutput() GetEcsDisksResultOutput {
	return o
}

func (o GetEcsDisksResultOutput) ToGetEcsDisksResultOutputWithContext(ctx context.Context) GetEcsDisksResultOutput {
	return o
}

func (o GetEcsDisksResultOutput) AdditionalAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDisksResult) []string { return v.AdditionalAttributes }).(pulumi.StringArrayOutput)
}

func (o GetEcsDisksResultOutput) AutoSnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.AutoSnapshotPolicyId }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'availability_zone' has been deprecated from provider version 1.122.0. New field 'zone_id' instead
func (o GetEcsDisksResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) DeleteAutoSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.DeleteAutoSnapshot }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) DeleteWithInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.DeleteWithInstance }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Disks() GetEcsDisksDiskArrayOutput {
	return o.ApplyT(func(v GetEcsDisksResult) []GetEcsDisksDisk { return v.Disks }).(GetEcsDisksDiskArrayOutput)
}

func (o GetEcsDisksResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) EnableAutoSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.EnableAutoSnapshot }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) EnableAutomatedSnapshotPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.EnableAutomatedSnapshotPolicy }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) EnableShared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.EnableShared }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) Encrypted() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.Encrypted }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsDisksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsDisksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsDisksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDisksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsDisksResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsDisksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsDisksResultOutput) OperationLocks() GetEcsDisksOperationLockArrayOutput {
	return o.ApplyT(func(v GetEcsDisksResult) []GetEcsDisksOperationLock { return v.OperationLocks }).(GetEcsDisksOperationLockArrayOutput)
}

func (o GetEcsDisksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetEcsDisksResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o GetEcsDisksResultOutput) PaymentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.PaymentType }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Portable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *bool { return v.Portable }).(pulumi.BoolPtrOutput)
}

func (o GetEcsDisksResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEcsDisksResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetEcsDisksResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetEcsDisksResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// Deprecated: Field 'type' has been deprecated from provider version 1.122.0. New field 'disk_type' instead.
func (o GetEcsDisksResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetEcsDisksResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsDisksResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsDisksResultOutput{})
}
