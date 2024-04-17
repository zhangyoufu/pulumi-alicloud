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

// Provides an RDS instance emote disaster recovery strategy policy resource and used to configure instance emote disaster recovery strategy policy.
//
// For information about RDS cross region backup settings and how to use them, see [What is cross region backup](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/modify-cross-region-backup-settings).
//
// > **NOTE:** Available since v1.195.0.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := rds.GetZones(ctx, &rds.GetZonesArgs{
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				DbInstanceStorageType: pulumi.StringRef("local_ssd"),
//				Category:              pulumi.StringRef("HighAvailability"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetInstanceClasses, err := rds.GetInstanceClasses(ctx, &rds.GetInstanceClassesArgs{
//				ZoneId:                pulumi.StringRef(_default.Ids[0]),
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				DbInstanceStorageType: pulumi.StringRef("local_ssd"),
//				Category:              pulumi.StringRef("HighAvailability"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			regions, err := rds.GetCrossRegions(ctx, nil, nil)
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
//			defaultInstance, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
//				Engine:                pulumi.String("MySQL"),
//				EngineVersion:         pulumi.String("8.0"),
//				InstanceType:          pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].InstanceClass),
//				InstanceStorage:       pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].StorageRange.Min),
//				InstanceChargeType:    pulumi.String("Postpaid"),
//				Category:              pulumi.String("HighAvailability"),
//				InstanceName:          pulumi.String(name),
//				VswitchId:             defaultSwitch.ID(),
//				DbInstanceStorageType: pulumi.String("local_ssd"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewRdsInstanceCrossBackupPolicy(ctx, "default", &rds.RdsInstanceCrossBackupPolicyArgs{
//				InstanceId:        defaultInstance.ID(),
//				CrossBackupRegion: pulumi.String(regions.Ids[0]),
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
// RDS remote disaster recovery policies can be imported using id or instance id, e.g.
//
// ```sh
// $ pulumi import alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy example "rm-12345678"
// ```
type RdsInstanceCrossBackupPolicy struct {
	pulumi.CustomResourceState

	// The status of the overall cross-region backup switch on the instance. Valid values:
	// - Disabled
	// - Enable
	BackupEnabled pulumi.StringOutput `pulumi:"backupEnabled"`
	// The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	BackupEnabledTime pulumi.StringOutput `pulumi:"backupEnabledTime"`
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	CrossBackupRegion pulumi.StringOutput `pulumi:"crossBackupRegion"`
	// The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
	CrossBackupType pulumi.StringOutput `pulumi:"crossBackupType"`
	// The state of the instance. For more information, see Instance status.
	DbInstanceStatus pulumi.StringOutput `pulumi:"dbInstanceStatus"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The lock status of the instance. Valid values:
	// - Unlock: The instance is not locked.
	// - ManualLock: The instance is manually locked.
	// - LockByExpiration: The instance is locked upon expiration.
	// - LockByRestoration: The instance is automatically locked before a rollback.
	// - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
	LockMode pulumi.StringOutput `pulumi:"lockMode"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	// - Enable: Enables the feature.
	// - Disabled: Disables the feature.
	LogBackupEnabled pulumi.StringOutput `pulumi:"logBackupEnabled"`
	// The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	LogBackupEnabledTime pulumi.StringOutput `pulumi:"logBackupEnabledTime"`
	// The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
	RetentType pulumi.StringOutput `pulumi:"retentType"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
	Retention pulumi.IntOutput `pulumi:"retention"`
}

// NewRdsInstanceCrossBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewRdsInstanceCrossBackupPolicy(ctx *pulumi.Context,
	name string, args *RdsInstanceCrossBackupPolicyArgs, opts ...pulumi.ResourceOption) (*RdsInstanceCrossBackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CrossBackupRegion == nil {
		return nil, errors.New("invalid value for required argument 'CrossBackupRegion'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsInstanceCrossBackupPolicy
	err := ctx.RegisterResource("alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsInstanceCrossBackupPolicy gets an existing RdsInstanceCrossBackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsInstanceCrossBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsInstanceCrossBackupPolicyState, opts ...pulumi.ResourceOption) (*RdsInstanceCrossBackupPolicy, error) {
	var resource RdsInstanceCrossBackupPolicy
	err := ctx.ReadResource("alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsInstanceCrossBackupPolicy resources.
type rdsInstanceCrossBackupPolicyState struct {
	// The status of the overall cross-region backup switch on the instance. Valid values:
	// - Disabled
	// - Enable
	BackupEnabled *string `pulumi:"backupEnabled"`
	// The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	BackupEnabledTime *string `pulumi:"backupEnabledTime"`
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	CrossBackupRegion *string `pulumi:"crossBackupRegion"`
	// The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
	CrossBackupType *string `pulumi:"crossBackupType"`
	// The state of the instance. For more information, see Instance status.
	DbInstanceStatus *string `pulumi:"dbInstanceStatus"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The lock status of the instance. Valid values:
	// - Unlock: The instance is not locked.
	// - ManualLock: The instance is manually locked.
	// - LockByExpiration: The instance is locked upon expiration.
	// - LockByRestoration: The instance is automatically locked before a rollback.
	// - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
	LockMode *string `pulumi:"lockMode"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	// - Enable: Enables the feature.
	// - Disabled: Disables the feature.
	LogBackupEnabled *string `pulumi:"logBackupEnabled"`
	// The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	LogBackupEnabledTime *string `pulumi:"logBackupEnabledTime"`
	// The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
	RetentType *string `pulumi:"retentType"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
	Retention *int `pulumi:"retention"`
}

type RdsInstanceCrossBackupPolicyState struct {
	// The status of the overall cross-region backup switch on the instance. Valid values:
	// - Disabled
	// - Enable
	BackupEnabled pulumi.StringPtrInput
	// The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	BackupEnabledTime pulumi.StringPtrInput
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	CrossBackupRegion pulumi.StringPtrInput
	// The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
	CrossBackupType pulumi.StringPtrInput
	// The state of the instance. For more information, see Instance status.
	DbInstanceStatus pulumi.StringPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The lock status of the instance. Valid values:
	// - Unlock: The instance is not locked.
	// - ManualLock: The instance is manually locked.
	// - LockByExpiration: The instance is locked upon expiration.
	// - LockByRestoration: The instance is automatically locked before a rollback.
	// - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
	LockMode pulumi.StringPtrInput
	// The status of the cross-region log backup feature on the instance. Valid values:
	// - Enable: Enables the feature.
	// - Disabled: Disables the feature.
	LogBackupEnabled pulumi.StringPtrInput
	// The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	LogBackupEnabledTime pulumi.StringPtrInput
	// The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
	RetentType pulumi.StringPtrInput
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
	Retention pulumi.IntPtrInput
}

func (RdsInstanceCrossBackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceCrossBackupPolicyState)(nil)).Elem()
}

type rdsInstanceCrossBackupPolicyArgs struct {
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	CrossBackupRegion string `pulumi:"crossBackupRegion"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	// - Enable: Enables the feature.
	// - Disabled: Disables the feature.
	LogBackupEnabled *string `pulumi:"logBackupEnabled"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
	Retention *int `pulumi:"retention"`
}

// The set of arguments for constructing a RdsInstanceCrossBackupPolicy resource.
type RdsInstanceCrossBackupPolicyArgs struct {
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	CrossBackupRegion pulumi.StringInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The status of the cross-region log backup feature on the instance. Valid values:
	// - Enable: Enables the feature.
	// - Disabled: Disables the feature.
	LogBackupEnabled pulumi.StringPtrInput
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
	Retention pulumi.IntPtrInput
}

func (RdsInstanceCrossBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceCrossBackupPolicyArgs)(nil)).Elem()
}

type RdsInstanceCrossBackupPolicyInput interface {
	pulumi.Input

	ToRdsInstanceCrossBackupPolicyOutput() RdsInstanceCrossBackupPolicyOutput
	ToRdsInstanceCrossBackupPolicyOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyOutput
}

func (*RdsInstanceCrossBackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (i *RdsInstanceCrossBackupPolicy) ToRdsInstanceCrossBackupPolicyOutput() RdsInstanceCrossBackupPolicyOutput {
	return i.ToRdsInstanceCrossBackupPolicyOutputWithContext(context.Background())
}

func (i *RdsInstanceCrossBackupPolicy) ToRdsInstanceCrossBackupPolicyOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceCrossBackupPolicyOutput)
}

// RdsInstanceCrossBackupPolicyArrayInput is an input type that accepts RdsInstanceCrossBackupPolicyArray and RdsInstanceCrossBackupPolicyArrayOutput values.
// You can construct a concrete instance of `RdsInstanceCrossBackupPolicyArrayInput` via:
//
//	RdsInstanceCrossBackupPolicyArray{ RdsInstanceCrossBackupPolicyArgs{...} }
type RdsInstanceCrossBackupPolicyArrayInput interface {
	pulumi.Input

	ToRdsInstanceCrossBackupPolicyArrayOutput() RdsInstanceCrossBackupPolicyArrayOutput
	ToRdsInstanceCrossBackupPolicyArrayOutputWithContext(context.Context) RdsInstanceCrossBackupPolicyArrayOutput
}

type RdsInstanceCrossBackupPolicyArray []RdsInstanceCrossBackupPolicyInput

func (RdsInstanceCrossBackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (i RdsInstanceCrossBackupPolicyArray) ToRdsInstanceCrossBackupPolicyArrayOutput() RdsInstanceCrossBackupPolicyArrayOutput {
	return i.ToRdsInstanceCrossBackupPolicyArrayOutputWithContext(context.Background())
}

func (i RdsInstanceCrossBackupPolicyArray) ToRdsInstanceCrossBackupPolicyArrayOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceCrossBackupPolicyArrayOutput)
}

// RdsInstanceCrossBackupPolicyMapInput is an input type that accepts RdsInstanceCrossBackupPolicyMap and RdsInstanceCrossBackupPolicyMapOutput values.
// You can construct a concrete instance of `RdsInstanceCrossBackupPolicyMapInput` via:
//
//	RdsInstanceCrossBackupPolicyMap{ "key": RdsInstanceCrossBackupPolicyArgs{...} }
type RdsInstanceCrossBackupPolicyMapInput interface {
	pulumi.Input

	ToRdsInstanceCrossBackupPolicyMapOutput() RdsInstanceCrossBackupPolicyMapOutput
	ToRdsInstanceCrossBackupPolicyMapOutputWithContext(context.Context) RdsInstanceCrossBackupPolicyMapOutput
}

type RdsInstanceCrossBackupPolicyMap map[string]RdsInstanceCrossBackupPolicyInput

func (RdsInstanceCrossBackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (i RdsInstanceCrossBackupPolicyMap) ToRdsInstanceCrossBackupPolicyMapOutput() RdsInstanceCrossBackupPolicyMapOutput {
	return i.ToRdsInstanceCrossBackupPolicyMapOutputWithContext(context.Background())
}

func (i RdsInstanceCrossBackupPolicyMap) ToRdsInstanceCrossBackupPolicyMapOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceCrossBackupPolicyMapOutput)
}

type RdsInstanceCrossBackupPolicyOutput struct{ *pulumi.OutputState }

func (RdsInstanceCrossBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (o RdsInstanceCrossBackupPolicyOutput) ToRdsInstanceCrossBackupPolicyOutput() RdsInstanceCrossBackupPolicyOutput {
	return o
}

func (o RdsInstanceCrossBackupPolicyOutput) ToRdsInstanceCrossBackupPolicyOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyOutput {
	return o
}

// The status of the overall cross-region backup switch on the instance. Valid values:
// - Disabled
// - Enable
func (o RdsInstanceCrossBackupPolicyOutput) BackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.BackupEnabled }).(pulumi.StringOutput)
}

// The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o RdsInstanceCrossBackupPolicyOutput) BackupEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.BackupEnabledTime }).(pulumi.StringOutput)
}

// The ID of the destination region where the cross-region backup files of the instance are stored.
func (o RdsInstanceCrossBackupPolicyOutput) CrossBackupRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.CrossBackupRegion }).(pulumi.StringOutput)
}

// The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
func (o RdsInstanceCrossBackupPolicyOutput) CrossBackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.CrossBackupType }).(pulumi.StringOutput)
}

// The state of the instance. For more information, see Instance status.
func (o RdsInstanceCrossBackupPolicyOutput) DbInstanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.DbInstanceStatus }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o RdsInstanceCrossBackupPolicyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The lock status of the instance. Valid values:
// - Unlock: The instance is not locked.
// - ManualLock: The instance is manually locked.
// - LockByExpiration: The instance is locked upon expiration.
// - LockByRestoration: The instance is automatically locked before a rollback.
// - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
func (o RdsInstanceCrossBackupPolicyOutput) LockMode() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.LockMode }).(pulumi.StringOutput)
}

// The status of the cross-region log backup feature on the instance. Valid values:
// - Enable: Enables the feature.
// - Disabled: Disables the feature.
func (o RdsInstanceCrossBackupPolicyOutput) LogBackupEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.LogBackupEnabled }).(pulumi.StringOutput)
}

// The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o RdsInstanceCrossBackupPolicyOutput) LogBackupEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.LogBackupEnabledTime }).(pulumi.StringOutput)
}

// The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
func (o RdsInstanceCrossBackupPolicyOutput) RetentType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.StringOutput { return v.RetentType }).(pulumi.StringOutput)
}

// The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
func (o RdsInstanceCrossBackupPolicyOutput) Retention() pulumi.IntOutput {
	return o.ApplyT(func(v *RdsInstanceCrossBackupPolicy) pulumi.IntOutput { return v.Retention }).(pulumi.IntOutput)
}

type RdsInstanceCrossBackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (RdsInstanceCrossBackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (o RdsInstanceCrossBackupPolicyArrayOutput) ToRdsInstanceCrossBackupPolicyArrayOutput() RdsInstanceCrossBackupPolicyArrayOutput {
	return o
}

func (o RdsInstanceCrossBackupPolicyArrayOutput) ToRdsInstanceCrossBackupPolicyArrayOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyArrayOutput {
	return o
}

func (o RdsInstanceCrossBackupPolicyArrayOutput) Index(i pulumi.IntInput) RdsInstanceCrossBackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsInstanceCrossBackupPolicy {
		return vs[0].([]*RdsInstanceCrossBackupPolicy)[vs[1].(int)]
	}).(RdsInstanceCrossBackupPolicyOutput)
}

type RdsInstanceCrossBackupPolicyMapOutput struct{ *pulumi.OutputState }

func (RdsInstanceCrossBackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstanceCrossBackupPolicy)(nil)).Elem()
}

func (o RdsInstanceCrossBackupPolicyMapOutput) ToRdsInstanceCrossBackupPolicyMapOutput() RdsInstanceCrossBackupPolicyMapOutput {
	return o
}

func (o RdsInstanceCrossBackupPolicyMapOutput) ToRdsInstanceCrossBackupPolicyMapOutputWithContext(ctx context.Context) RdsInstanceCrossBackupPolicyMapOutput {
	return o
}

func (o RdsInstanceCrossBackupPolicyMapOutput) MapIndex(k pulumi.StringInput) RdsInstanceCrossBackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsInstanceCrossBackupPolicy {
		return vs[0].(map[string]*RdsInstanceCrossBackupPolicy)[vs[1].(string)]
	}).(RdsInstanceCrossBackupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceCrossBackupPolicyInput)(nil)).Elem(), &RdsInstanceCrossBackupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceCrossBackupPolicyArrayInput)(nil)).Elem(), RdsInstanceCrossBackupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceCrossBackupPolicyMapInput)(nil)).Elem(), RdsInstanceCrossBackupPolicyMap{})
	pulumi.RegisterOutputType(RdsInstanceCrossBackupPolicyOutput{})
	pulumi.RegisterOutputType(RdsInstanceCrossBackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(RdsInstanceCrossBackupPolicyMapOutput{})
}
