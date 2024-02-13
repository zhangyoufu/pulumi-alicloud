// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RDS Backup resource.
//
// For information about RDS Backup and how to use it, see [What is Backup](https://www.alibabacloud.com/help/en/rds/developer-reference/api-rds-2014-08-15-createbackup).
//
// > **NOTE:** Available since v1.149.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
//				Engine:                pulumi.String("MySQL"),
//				EngineVersion:         pulumi.String("5.6"),
//				InstanceType:          pulumi.String("rds.mysql.t1.small"),
//				InstanceStorage:       pulumi.Int(30),
//				InstanceChargeType:    pulumi.String("Postpaid"),
//				DbInstanceStorageType: pulumi.String("local_ssd"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewRdsBackup(ctx, "exampleRdsBackup", &rds.RdsBackupArgs{
//				DbInstanceId: exampleInstance.ID(),
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
// RDS Backup can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:rds/rdsBackup:RdsBackup example <db_instance_id>:<backup_id>
// ```
type RdsBackup struct {
	pulumi.CustomResourceState

	// The backup id.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
	BackupMethod pulumi.StringOutput `pulumi:"backupMethod"`
	// The policy that you want to use for the backup task. Valid values:
	// * **db**: specifies to perform a database-level backup.
	// * **instance**: specifies to perform an instance-level backup.
	BackupStrategy pulumi.StringPtrOutput `pulumi:"backupStrategy"`
	// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
	// * **Auto**: specifies to automatically perform a full or incremental backup.
	// * **FullBackup**: specifies to perform a full backup.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// The db instance id.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	DbName pulumi.StringPtrOutput `pulumi:"dbName"`
	// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
	RemoveFromState pulumi.BoolPtrOutput `pulumi:"removeFromState"`
	// Indicates whether the data backup file can be deleted. Valid values: `Enabled` and `Disabled`.
	StoreStatus pulumi.StringOutput `pulumi:"storeStatus"`
}

// NewRdsBackup registers a new resource with the given unique name, arguments, and options.
func NewRdsBackup(ctx *pulumi.Context,
	name string, args *RdsBackupArgs, opts ...pulumi.ResourceOption) (*RdsBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsBackup
	err := ctx.RegisterResource("alicloud:rds/rdsBackup:RdsBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsBackup gets an existing RdsBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsBackupState, opts ...pulumi.ResourceOption) (*RdsBackup, error) {
	var resource RdsBackup
	err := ctx.ReadResource("alicloud:rds/rdsBackup:RdsBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsBackup resources.
type rdsBackupState struct {
	// The backup id.
	BackupId *string `pulumi:"backupId"`
	// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
	BackupMethod *string `pulumi:"backupMethod"`
	// The policy that you want to use for the backup task. Valid values:
	// * **db**: specifies to perform a database-level backup.
	// * **instance**: specifies to perform an instance-level backup.
	BackupStrategy *string `pulumi:"backupStrategy"`
	// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
	// * **Auto**: specifies to automatically perform a full or incremental backup.
	// * **FullBackup**: specifies to perform a full backup.
	BackupType *string `pulumi:"backupType"`
	// The db instance id.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	DbName *string `pulumi:"dbName"`
	// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
	RemoveFromState *bool `pulumi:"removeFromState"`
	// Indicates whether the data backup file can be deleted. Valid values: `Enabled` and `Disabled`.
	StoreStatus *string `pulumi:"storeStatus"`
}

type RdsBackupState struct {
	// The backup id.
	BackupId pulumi.StringPtrInput
	// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
	BackupMethod pulumi.StringPtrInput
	// The policy that you want to use for the backup task. Valid values:
	// * **db**: specifies to perform a database-level backup.
	// * **instance**: specifies to perform an instance-level backup.
	BackupStrategy pulumi.StringPtrInput
	// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
	// * **Auto**: specifies to automatically perform a full or incremental backup.
	// * **FullBackup**: specifies to perform a full backup.
	BackupType pulumi.StringPtrInput
	// The db instance id.
	DbInstanceId pulumi.StringPtrInput
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	DbName pulumi.StringPtrInput
	// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
	RemoveFromState pulumi.BoolPtrInput
	// Indicates whether the data backup file can be deleted. Valid values: `Enabled` and `Disabled`.
	StoreStatus pulumi.StringPtrInput
}

func (RdsBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsBackupState)(nil)).Elem()
}

type rdsBackupArgs struct {
	// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
	BackupMethod *string `pulumi:"backupMethod"`
	// The policy that you want to use for the backup task. Valid values:
	// * **db**: specifies to perform a database-level backup.
	// * **instance**: specifies to perform an instance-level backup.
	BackupStrategy *string `pulumi:"backupStrategy"`
	// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
	// * **Auto**: specifies to automatically perform a full or incremental backup.
	// * **FullBackup**: specifies to perform a full backup.
	BackupType *string `pulumi:"backupType"`
	// The db instance id.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	DbName *string `pulumi:"dbName"`
	// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
	RemoveFromState *bool `pulumi:"removeFromState"`
}

// The set of arguments for constructing a RdsBackup resource.
type RdsBackupArgs struct {
	// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
	BackupMethod pulumi.StringPtrInput
	// The policy that you want to use for the backup task. Valid values:
	// * **db**: specifies to perform a database-level backup.
	// * **instance**: specifies to perform an instance-level backup.
	BackupStrategy pulumi.StringPtrInput
	// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
	// * **Auto**: specifies to automatically perform a full or incremental backup.
	// * **FullBackup**: specifies to perform a full backup.
	BackupType pulumi.StringPtrInput
	// The db instance id.
	DbInstanceId pulumi.StringInput
	// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
	DbName pulumi.StringPtrInput
	// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
	RemoveFromState pulumi.BoolPtrInput
}

func (RdsBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsBackupArgs)(nil)).Elem()
}

type RdsBackupInput interface {
	pulumi.Input

	ToRdsBackupOutput() RdsBackupOutput
	ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput
}

func (*RdsBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsBackup)(nil)).Elem()
}

func (i *RdsBackup) ToRdsBackupOutput() RdsBackupOutput {
	return i.ToRdsBackupOutputWithContext(context.Background())
}

func (i *RdsBackup) ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupOutput)
}

// RdsBackupArrayInput is an input type that accepts RdsBackupArray and RdsBackupArrayOutput values.
// You can construct a concrete instance of `RdsBackupArrayInput` via:
//
//	RdsBackupArray{ RdsBackupArgs{...} }
type RdsBackupArrayInput interface {
	pulumi.Input

	ToRdsBackupArrayOutput() RdsBackupArrayOutput
	ToRdsBackupArrayOutputWithContext(context.Context) RdsBackupArrayOutput
}

type RdsBackupArray []RdsBackupInput

func (RdsBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsBackup)(nil)).Elem()
}

func (i RdsBackupArray) ToRdsBackupArrayOutput() RdsBackupArrayOutput {
	return i.ToRdsBackupArrayOutputWithContext(context.Background())
}

func (i RdsBackupArray) ToRdsBackupArrayOutputWithContext(ctx context.Context) RdsBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupArrayOutput)
}

// RdsBackupMapInput is an input type that accepts RdsBackupMap and RdsBackupMapOutput values.
// You can construct a concrete instance of `RdsBackupMapInput` via:
//
//	RdsBackupMap{ "key": RdsBackupArgs{...} }
type RdsBackupMapInput interface {
	pulumi.Input

	ToRdsBackupMapOutput() RdsBackupMapOutput
	ToRdsBackupMapOutputWithContext(context.Context) RdsBackupMapOutput
}

type RdsBackupMap map[string]RdsBackupInput

func (RdsBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsBackup)(nil)).Elem()
}

func (i RdsBackupMap) ToRdsBackupMapOutput() RdsBackupMapOutput {
	return i.ToRdsBackupMapOutputWithContext(context.Background())
}

func (i RdsBackupMap) ToRdsBackupMapOutputWithContext(ctx context.Context) RdsBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsBackupMapOutput)
}

type RdsBackupOutput struct{ *pulumi.OutputState }

func (RdsBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsBackup)(nil)).Elem()
}

func (o RdsBackupOutput) ToRdsBackupOutput() RdsBackupOutput {
	return o
}

func (o RdsBackupOutput) ToRdsBackupOutputWithContext(ctx context.Context) RdsBackupOutput {
	return o
}

// The backup id.
func (o RdsBackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// The type of backup that you want to perform. Default value: `Physical`. Valid values: `Logical`, `Physical` and `Snapshot`.
func (o RdsBackupOutput) BackupMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.BackupMethod }).(pulumi.StringOutput)
}

// The policy that you want to use for the backup task. Valid values:
// * **db**: specifies to perform a database-level backup.
// * **instance**: specifies to perform an instance-level backup.
func (o RdsBackupOutput) BackupStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringPtrOutput { return v.BackupStrategy }).(pulumi.StringPtrOutput)
}

// The method that you want to use for the backup task. Default value: `Auto`. Valid values:
// * **Auto**: specifies to automatically perform a full or incremental backup.
// * **FullBackup**: specifies to perform a full backup.
func (o RdsBackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// The db instance id.
func (o RdsBackupOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The names of the databases whose data you want to back up. Separate the names of the databases with commas (,).
func (o RdsBackupOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringPtrOutput { return v.DbName }).(pulumi.StringPtrOutput)
}

// Remove form state when resource cannot be deleted. Valid values: `true` and `false`.
func (o RdsBackupOutput) RemoveFromState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.BoolPtrOutput { return v.RemoveFromState }).(pulumi.BoolPtrOutput)
}

// Indicates whether the data backup file can be deleted. Valid values: `Enabled` and `Disabled`.
func (o RdsBackupOutput) StoreStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsBackup) pulumi.StringOutput { return v.StoreStatus }).(pulumi.StringOutput)
}

type RdsBackupArrayOutput struct{ *pulumi.OutputState }

func (RdsBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsBackup)(nil)).Elem()
}

func (o RdsBackupArrayOutput) ToRdsBackupArrayOutput() RdsBackupArrayOutput {
	return o
}

func (o RdsBackupArrayOutput) ToRdsBackupArrayOutputWithContext(ctx context.Context) RdsBackupArrayOutput {
	return o
}

func (o RdsBackupArrayOutput) Index(i pulumi.IntInput) RdsBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsBackup {
		return vs[0].([]*RdsBackup)[vs[1].(int)]
	}).(RdsBackupOutput)
}

type RdsBackupMapOutput struct{ *pulumi.OutputState }

func (RdsBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsBackup)(nil)).Elem()
}

func (o RdsBackupMapOutput) ToRdsBackupMapOutput() RdsBackupMapOutput {
	return o
}

func (o RdsBackupMapOutput) ToRdsBackupMapOutputWithContext(ctx context.Context) RdsBackupMapOutput {
	return o
}

func (o RdsBackupMapOutput) MapIndex(k pulumi.StringInput) RdsBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsBackup {
		return vs[0].(map[string]*RdsBackup)[vs[1].(string)]
	}).(RdsBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupInput)(nil)).Elem(), &RdsBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupArrayInput)(nil)).Elem(), RdsBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsBackupMapInput)(nil)).Elem(), RdsBackupMap{})
	pulumi.RegisterOutputType(RdsBackupOutput{})
	pulumi.RegisterOutputType(RdsBackupArrayOutput{})
	pulumi.RegisterOutputType(RdsBackupMapOutput{})
}
