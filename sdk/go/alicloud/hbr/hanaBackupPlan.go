// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Hybrid Backup Recovery (HBR) Hana Backup Plan resource.
//
// For information about Hybrid Backup Recovery (HBR) Hana Backup Plan and how to use it, see [What is Hana Backup Plan](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/api-hbr-2017-09-08-createhanabackupplan).
//
// > **NOTE:** Available in v1.179.0+.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			exampleVault, err := hbr.NewVault(ctx, "example", &hbr.VaultArgs{
//				VaultName: pulumi.Sprintf("terraform-example-%v", _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			exampleHanaInstance, err := hbr.NewHanaInstance(ctx, "example", &hbr.HanaInstanceArgs{
//				AlertSetting:        pulumi.String("INHERITED"),
//				HanaName:            pulumi.Sprintf("terraform-example-%v", _default.Result),
//				Host:                pulumi.String("1.1.1.1"),
//				InstanceNumber:      pulumi.Int(1),
//				Password:            pulumi.String("YouPassword123"),
//				ResourceGroupId:     pulumi.String(example.Groups[0].Id),
//				Sid:                 pulumi.String("HXE"),
//				UseSsl:              pulumi.Bool(false),
//				UserName:            pulumi.String("admin"),
//				ValidateCertificate: pulumi.Bool(false),
//				VaultId:             exampleVault.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = hbr.NewHanaBackupPlan(ctx, "example", &hbr.HanaBackupPlanArgs{
//				BackupPrefix:    pulumi.String("DIFF_DATA_BACKUP"),
//				BackupType:      pulumi.String("COMPLETE"),
//				ClusterId:       exampleHanaInstance.HanaInstanceId,
//				DatabaseName:    pulumi.String("SYSTEMDB"),
//				PlanName:        pulumi.String("terraform-example"),
//				ResourceGroupId: pulumi.String(example.Groups[0].Id),
//				Schedule:        pulumi.String("I|1602673264|P1D"),
//				VaultId:         exampleHanaInstance.VaultId,
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
// Hybrid Backup Recovery (HBR) Hana Backup Plan can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:hbr/hanaBackupPlan:HanaBackupPlan example <plan_id>:<vault_id>:<cluster_id>
// ```
type HanaBackupPlan struct {
	pulumi.CustomResourceState

	// The backup prefix.
	BackupPrefix pulumi.StringPtrOutput `pulumi:"backupPrefix"`
	// The backup type. Valid values:
	// - `COMPLETE`: full backup.
	// - `INCREMENTAL`: incremental backup.
	// - `DIFFERENTIAL`: differential backup.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// The ID of the SAP HANA instance.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The name of the database.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The id of the plan.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// The name of the backup plan.
	PlanName pulumi.StringOutput `pulumi:"planName"`
	// The resource attribute field that represents the resource group ID.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The status of the resource. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the backup vault.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewHanaBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewHanaBackupPlan(ctx *pulumi.Context,
	name string, args *HanaBackupPlanArgs, opts ...pulumi.ResourceOption) (*HanaBackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupType == nil {
		return nil, errors.New("invalid value for required argument 'BackupType'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.PlanName == nil {
		return nil, errors.New("invalid value for required argument 'PlanName'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HanaBackupPlan
	err := ctx.RegisterResource("alicloud:hbr/hanaBackupPlan:HanaBackupPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHanaBackupPlan gets an existing HanaBackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHanaBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HanaBackupPlanState, opts ...pulumi.ResourceOption) (*HanaBackupPlan, error) {
	var resource HanaBackupPlan
	err := ctx.ReadResource("alicloud:hbr/hanaBackupPlan:HanaBackupPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HanaBackupPlan resources.
type hanaBackupPlanState struct {
	// The backup prefix.
	BackupPrefix *string `pulumi:"backupPrefix"`
	// The backup type. Valid values:
	// - `COMPLETE`: full backup.
	// - `INCREMENTAL`: incremental backup.
	// - `DIFFERENTIAL`: differential backup.
	BackupType *string `pulumi:"backupType"`
	// The ID of the SAP HANA instance.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the database.
	DatabaseName *string `pulumi:"databaseName"`
	// The id of the plan.
	PlanId *string `pulumi:"planId"`
	// The name of the backup plan.
	PlanName *string `pulumi:"planName"`
	// The resource attribute field that represents the resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	Schedule *string `pulumi:"schedule"`
	// The status of the resource. Valid values: `Enabled`, `Disabled`.
	Status *string `pulumi:"status"`
	// The ID of the backup vault.
	VaultId *string `pulumi:"vaultId"`
}

type HanaBackupPlanState struct {
	// The backup prefix.
	BackupPrefix pulumi.StringPtrInput
	// The backup type. Valid values:
	// - `COMPLETE`: full backup.
	// - `INCREMENTAL`: incremental backup.
	// - `DIFFERENTIAL`: differential backup.
	BackupType pulumi.StringPtrInput
	// The ID of the SAP HANA instance.
	ClusterId pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringPtrInput
	// The id of the plan.
	PlanId pulumi.StringPtrInput
	// The name of the backup plan.
	PlanName pulumi.StringPtrInput
	// The resource attribute field that represents the resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	Schedule pulumi.StringPtrInput
	// The status of the resource. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringPtrInput
	// The ID of the backup vault.
	VaultId pulumi.StringPtrInput
}

func (HanaBackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaBackupPlanState)(nil)).Elem()
}

type hanaBackupPlanArgs struct {
	// The backup prefix.
	BackupPrefix *string `pulumi:"backupPrefix"`
	// The backup type. Valid values:
	// - `COMPLETE`: full backup.
	// - `INCREMENTAL`: incremental backup.
	// - `DIFFERENTIAL`: differential backup.
	BackupType string `pulumi:"backupType"`
	// The ID of the SAP HANA instance.
	ClusterId string `pulumi:"clusterId"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the backup plan.
	PlanName string `pulumi:"planName"`
	// The resource attribute field that represents the resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	Schedule string `pulumi:"schedule"`
	// The status of the resource. Valid values: `Enabled`, `Disabled`.
	Status *string `pulumi:"status"`
	// The ID of the backup vault.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a HanaBackupPlan resource.
type HanaBackupPlanArgs struct {
	// The backup prefix.
	BackupPrefix pulumi.StringPtrInput
	// The backup type. Valid values:
	// - `COMPLETE`: full backup.
	// - `INCREMENTAL`: incremental backup.
	// - `DIFFERENTIAL`: differential backup.
	BackupType pulumi.StringInput
	// The ID of the SAP HANA instance.
	ClusterId pulumi.StringInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The name of the backup plan.
	PlanName pulumi.StringInput
	// The resource attribute field that represents the resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	Schedule pulumi.StringInput
	// The status of the resource. Valid values: `Enabled`, `Disabled`.
	Status pulumi.StringPtrInput
	// The ID of the backup vault.
	VaultId pulumi.StringInput
}

func (HanaBackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hanaBackupPlanArgs)(nil)).Elem()
}

type HanaBackupPlanInput interface {
	pulumi.Input

	ToHanaBackupPlanOutput() HanaBackupPlanOutput
	ToHanaBackupPlanOutputWithContext(ctx context.Context) HanaBackupPlanOutput
}

func (*HanaBackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaBackupPlan)(nil)).Elem()
}

func (i *HanaBackupPlan) ToHanaBackupPlanOutput() HanaBackupPlanOutput {
	return i.ToHanaBackupPlanOutputWithContext(context.Background())
}

func (i *HanaBackupPlan) ToHanaBackupPlanOutputWithContext(ctx context.Context) HanaBackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupPlanOutput)
}

// HanaBackupPlanArrayInput is an input type that accepts HanaBackupPlanArray and HanaBackupPlanArrayOutput values.
// You can construct a concrete instance of `HanaBackupPlanArrayInput` via:
//
//	HanaBackupPlanArray{ HanaBackupPlanArgs{...} }
type HanaBackupPlanArrayInput interface {
	pulumi.Input

	ToHanaBackupPlanArrayOutput() HanaBackupPlanArrayOutput
	ToHanaBackupPlanArrayOutputWithContext(context.Context) HanaBackupPlanArrayOutput
}

type HanaBackupPlanArray []HanaBackupPlanInput

func (HanaBackupPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HanaBackupPlan)(nil)).Elem()
}

func (i HanaBackupPlanArray) ToHanaBackupPlanArrayOutput() HanaBackupPlanArrayOutput {
	return i.ToHanaBackupPlanArrayOutputWithContext(context.Background())
}

func (i HanaBackupPlanArray) ToHanaBackupPlanArrayOutputWithContext(ctx context.Context) HanaBackupPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupPlanArrayOutput)
}

// HanaBackupPlanMapInput is an input type that accepts HanaBackupPlanMap and HanaBackupPlanMapOutput values.
// You can construct a concrete instance of `HanaBackupPlanMapInput` via:
//
//	HanaBackupPlanMap{ "key": HanaBackupPlanArgs{...} }
type HanaBackupPlanMapInput interface {
	pulumi.Input

	ToHanaBackupPlanMapOutput() HanaBackupPlanMapOutput
	ToHanaBackupPlanMapOutputWithContext(context.Context) HanaBackupPlanMapOutput
}

type HanaBackupPlanMap map[string]HanaBackupPlanInput

func (HanaBackupPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HanaBackupPlan)(nil)).Elem()
}

func (i HanaBackupPlanMap) ToHanaBackupPlanMapOutput() HanaBackupPlanMapOutput {
	return i.ToHanaBackupPlanMapOutputWithContext(context.Background())
}

func (i HanaBackupPlanMap) ToHanaBackupPlanMapOutputWithContext(ctx context.Context) HanaBackupPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HanaBackupPlanMapOutput)
}

type HanaBackupPlanOutput struct{ *pulumi.OutputState }

func (HanaBackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HanaBackupPlan)(nil)).Elem()
}

func (o HanaBackupPlanOutput) ToHanaBackupPlanOutput() HanaBackupPlanOutput {
	return o
}

func (o HanaBackupPlanOutput) ToHanaBackupPlanOutputWithContext(ctx context.Context) HanaBackupPlanOutput {
	return o
}

// The backup prefix.
func (o HanaBackupPlanOutput) BackupPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringPtrOutput { return v.BackupPrefix }).(pulumi.StringPtrOutput)
}

// The backup type. Valid values:
// - `COMPLETE`: full backup.
// - `INCREMENTAL`: incremental backup.
// - `DIFFERENTIAL`: differential backup.
func (o HanaBackupPlanOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// The ID of the SAP HANA instance.
func (o HanaBackupPlanOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of the database.
func (o HanaBackupPlanOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// The id of the plan.
func (o HanaBackupPlanOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// The name of the backup plan.
func (o HanaBackupPlanOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.PlanName }).(pulumi.StringOutput)
}

// The resource attribute field that represents the resource group ID.
func (o HanaBackupPlanOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
func (o HanaBackupPlanOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The status of the resource. Valid values: `Enabled`, `Disabled`.
func (o HanaBackupPlanOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the backup vault.
func (o HanaBackupPlanOutput) VaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *HanaBackupPlan) pulumi.StringOutput { return v.VaultId }).(pulumi.StringOutput)
}

type HanaBackupPlanArrayOutput struct{ *pulumi.OutputState }

func (HanaBackupPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HanaBackupPlan)(nil)).Elem()
}

func (o HanaBackupPlanArrayOutput) ToHanaBackupPlanArrayOutput() HanaBackupPlanArrayOutput {
	return o
}

func (o HanaBackupPlanArrayOutput) ToHanaBackupPlanArrayOutputWithContext(ctx context.Context) HanaBackupPlanArrayOutput {
	return o
}

func (o HanaBackupPlanArrayOutput) Index(i pulumi.IntInput) HanaBackupPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HanaBackupPlan {
		return vs[0].([]*HanaBackupPlan)[vs[1].(int)]
	}).(HanaBackupPlanOutput)
}

type HanaBackupPlanMapOutput struct{ *pulumi.OutputState }

func (HanaBackupPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HanaBackupPlan)(nil)).Elem()
}

func (o HanaBackupPlanMapOutput) ToHanaBackupPlanMapOutput() HanaBackupPlanMapOutput {
	return o
}

func (o HanaBackupPlanMapOutput) ToHanaBackupPlanMapOutputWithContext(ctx context.Context) HanaBackupPlanMapOutput {
	return o
}

func (o HanaBackupPlanMapOutput) MapIndex(k pulumi.StringInput) HanaBackupPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HanaBackupPlan {
		return vs[0].(map[string]*HanaBackupPlan)[vs[1].(string)]
	}).(HanaBackupPlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupPlanInput)(nil)).Elem(), &HanaBackupPlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupPlanArrayInput)(nil)).Elem(), HanaBackupPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HanaBackupPlanMapInput)(nil)).Elem(), HanaBackupPlanMap{})
	pulumi.RegisterOutputType(HanaBackupPlanOutput{})
	pulumi.RegisterOutputType(HanaBackupPlanArrayOutput{})
	pulumi.RegisterOutputType(HanaBackupPlanMapOutput{})
}
