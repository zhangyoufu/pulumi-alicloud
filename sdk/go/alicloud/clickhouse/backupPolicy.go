// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clickhouse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Click House Backup Policy resource.
//
// For information about Click House Backup Policy and how to use it, see [What is Backup Policy](https://www.alibabacloud.com/help/zh/clickhouse/latest/api-clickhouse-2019-11-11-createbackuppolicy).
//
// > **NOTE:** Available since v1.147.0.
//
// > **NOTE:** Only the cloud database ClickHouse cluster version `20.3` supports data backup.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/clickhouse"
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
//			defaultRegions, err := clickhouse.GetRegions(ctx, &clickhouse.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(defaultRegions.Regions[0].ZoneIds[0].ZoneId),
//			})
//			if err != nil {
//				return err
//			}
//			defaultDbCluster, err := clickhouse.NewDbCluster(ctx, "defaultDbCluster", &clickhouse.DbClusterArgs{
//				DbClusterVersion:     pulumi.String("22.8.5.29"),
//				Status:               pulumi.String("Running"),
//				Category:             pulumi.String("Basic"),
//				DbClusterClass:       pulumi.String("S8"),
//				DbClusterNetworkType: pulumi.String("vpc"),
//				DbNodeGroupCount:     pulumi.Int(1),
//				PaymentType:          pulumi.String("PayAsYouGo"),
//				DbNodeStorage:        pulumi.String("500"),
//				StorageType:          pulumi.String("cloud_essd"),
//				VswitchId:            defaultSwitch.ID(),
//				VpcId:                defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = clickhouse.NewBackupPolicy(ctx, "defaultBackupPolicy", &clickhouse.BackupPolicyArgs{
//				DbClusterId: defaultDbCluster.ID(),
//				PreferredBackupPeriods: pulumi.StringArray{
//					pulumi.String("Monday"),
//					pulumi.String("Friday"),
//				},
//				PreferredBackupTime:   pulumi.String("00:00Z-01:00Z"),
//				BackupRetentionPeriod: pulumi.Int(7),
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
// Click House Backup Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:clickhouse/backupPolicy:BackupPolicy example <db_cluster_id>
// ```
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Data backup days. Valid values: `7` to `730`.
	BackupRetentionPeriod pulumi.IntPtrOutput `pulumi:"backupRetentionPeriod"`
	// The id of the DBCluster.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
	PreferredBackupPeriods pulumi.StringArrayOutput `pulumi:"preferredBackupPeriods"`
	// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime pulumi.StringOutput `pulumi:"preferredBackupTime"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	if args.PreferredBackupPeriods == nil {
		return nil, errors.New("invalid value for required argument 'PreferredBackupPeriods'")
	}
	if args.PreferredBackupTime == nil {
		return nil, errors.New("invalid value for required argument 'PreferredBackupTime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("alicloud:clickhouse/backupPolicy:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("alicloud:clickhouse/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Data backup days. Valid values: `7` to `730`.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The id of the DBCluster.
	DbClusterId *string `pulumi:"dbClusterId"`
	// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
	PreferredBackupPeriods []string `pulumi:"preferredBackupPeriods"`
	// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime *string `pulumi:"preferredBackupTime"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

type BackupPolicyState struct {
	// Data backup days. Valid values: `7` to `730`.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The id of the DBCluster.
	DbClusterId pulumi.StringPtrInput
	// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
	PreferredBackupPeriods pulumi.StringArrayInput
	// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// Data backup days. Valid values: `7` to `730`.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The id of the DBCluster.
	DbClusterId string `pulumi:"dbClusterId"`
	// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
	PreferredBackupPeriods []string `pulumi:"preferredBackupPeriods"`
	// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime string `pulumi:"preferredBackupTime"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// Data backup days. Valid values: `7` to `730`.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The id of the DBCluster.
	DbClusterId pulumi.StringInput
	// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
	PreferredBackupPeriods pulumi.StringArrayInput
	// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime pulumi.StringInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

// BackupPolicyArrayInput is an input type that accepts BackupPolicyArray and BackupPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPolicyArrayInput` via:
//
//	BackupPolicyArray{ BackupPolicyArgs{...} }
type BackupPolicyArrayInput interface {
	pulumi.Input

	ToBackupPolicyArrayOutput() BackupPolicyArrayOutput
	ToBackupPolicyArrayOutputWithContext(context.Context) BackupPolicyArrayOutput
}

type BackupPolicyArray []BackupPolicyInput

func (BackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return i.ToBackupPolicyArrayOutputWithContext(context.Background())
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyArrayOutput)
}

// BackupPolicyMapInput is an input type that accepts BackupPolicyMap and BackupPolicyMapOutput values.
// You can construct a concrete instance of `BackupPolicyMapInput` via:
//
//	BackupPolicyMap{ "key": BackupPolicyArgs{...} }
type BackupPolicyMapInput interface {
	pulumi.Input

	ToBackupPolicyMapOutput() BackupPolicyMapOutput
	ToBackupPolicyMapOutputWithContext(context.Context) BackupPolicyMapOutput
}

type BackupPolicyMap map[string]BackupPolicyInput

func (BackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyMap) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return i.ToBackupPolicyMapOutputWithContext(context.Background())
}

func (i BackupPolicyMap) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMapOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

// Data backup days. Valid values: `7` to `730`.
func (o BackupPolicyOutput) BackupRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntPtrOutput { return v.BackupRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The id of the DBCluster.
func (o BackupPolicyOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// DBCluster Backup period. A list of DBCluster Backup period. Valid values: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].
func (o BackupPolicyOutput) PreferredBackupPeriods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringArrayOutput { return v.PreferredBackupPeriods }).(pulumi.StringArrayOutput)
}

// DBCluster backup time, in the format of `HH:mmZ-HH:mmZ`. Time setting interval is one hour. China time is 8 hours behind it.
func (o BackupPolicyOutput) PreferredBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.PreferredBackupTime }).(pulumi.StringOutput)
}

// The status of the resource.
func (o BackupPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) Index(i pulumi.IntInput) BackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].([]*BackupPolicy)[vs[1].(int)]
	}).(BackupPolicyOutput)
}

type BackupPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].(map[string]*BackupPolicy)[vs[1].(string)]
	}).(BackupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyInput)(nil)).Elem(), &BackupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyArrayInput)(nil)).Elem(), BackupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyMapInput)(nil)).Elem(), BackupPolicyMap{})
	pulumi.RegisterOutputType(BackupPolicyOutput{})
	pulumi.RegisterOutputType(BackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyMapOutput{})
}
