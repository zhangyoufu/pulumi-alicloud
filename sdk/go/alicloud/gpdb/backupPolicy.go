// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Import
//
// GPDB Backup Policy can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:gpdb/backupPolicy:BackupPolicy example <id>
//
// ```
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Data backup retention days.
	BackupRetentionPeriod pulumi.IntOutput `pulumi:"backupRetentionPeriod"`
	// The instance ID.
	// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// Whether to enable automatic recovery points. Value Description:
	// - **true**: enabled.
	// - **false**: Closed.
	EnableRecoveryPoint pulumi.BoolOutput `pulumi:"enableRecoveryPoint"`
	// Data backup cycle. Separate multiple values with commas (,). Value Description:
	// - **Monday**: Monday.
	// - **Tuesday**: Tuesday.
	// - **Wednesday**: Wednesday.
	// - **Thursday**: Thursday.
	// - **Friday**: Friday.
	// - **Saturday**: Saturday.
	// - **Sunday**: Sunday.
	PreferredBackupPeriod pulumi.StringOutput `pulumi:"preferredBackupPeriod"`
	// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
	PreferredBackupTime pulumi.StringOutput `pulumi:"preferredBackupTime"`
	// Recovery point frequency. Value Description:
	// - **1**: Hourly.
	// - **2**: Every two hours.
	// - **4**: Every four hours.
	// - **8**: Every eight hours.
	RecoveryPointPeriod pulumi.StringOutput `pulumi:"recoveryPointPeriod"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.PreferredBackupPeriod == nil {
		return nil, errors.New("invalid value for required argument 'PreferredBackupPeriod'")
	}
	if args.PreferredBackupTime == nil {
		return nil, errors.New("invalid value for required argument 'PreferredBackupTime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("alicloud:gpdb/backupPolicy:BackupPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:gpdb/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Data backup retention days.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The instance ID.
	// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// Whether to enable automatic recovery points. Value Description:
	// - **true**: enabled.
	// - **false**: Closed.
	EnableRecoveryPoint *bool `pulumi:"enableRecoveryPoint"`
	// Data backup cycle. Separate multiple values with commas (,). Value Description:
	// - **Monday**: Monday.
	// - **Tuesday**: Tuesday.
	// - **Wednesday**: Wednesday.
	// - **Thursday**: Thursday.
	// - **Friday**: Friday.
	// - **Saturday**: Saturday.
	// - **Sunday**: Sunday.
	PreferredBackupPeriod *string `pulumi:"preferredBackupPeriod"`
	// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
	PreferredBackupTime *string `pulumi:"preferredBackupTime"`
	// Recovery point frequency. Value Description:
	// - **1**: Hourly.
	// - **2**: Every two hours.
	// - **4**: Every four hours.
	// - **8**: Every eight hours.
	RecoveryPointPeriod *string `pulumi:"recoveryPointPeriod"`
}

type BackupPolicyState struct {
	// Data backup retention days.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The instance ID.
	// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	DbInstanceId pulumi.StringPtrInput
	// Whether to enable automatic recovery points. Value Description:
	// - **true**: enabled.
	// - **false**: Closed.
	EnableRecoveryPoint pulumi.BoolPtrInput
	// Data backup cycle. Separate multiple values with commas (,). Value Description:
	// - **Monday**: Monday.
	// - **Tuesday**: Tuesday.
	// - **Wednesday**: Wednesday.
	// - **Thursday**: Thursday.
	// - **Friday**: Friday.
	// - **Saturday**: Saturday.
	// - **Sunday**: Sunday.
	PreferredBackupPeriod pulumi.StringPtrInput
	// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
	PreferredBackupTime pulumi.StringPtrInput
	// Recovery point frequency. Value Description:
	// - **1**: Hourly.
	// - **2**: Every two hours.
	// - **4**: Every four hours.
	// - **8**: Every eight hours.
	RecoveryPointPeriod pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// Data backup retention days.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The instance ID.
	// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// Whether to enable automatic recovery points. Value Description:
	// - **true**: enabled.
	// - **false**: Closed.
	EnableRecoveryPoint *bool `pulumi:"enableRecoveryPoint"`
	// Data backup cycle. Separate multiple values with commas (,). Value Description:
	// - **Monday**: Monday.
	// - **Tuesday**: Tuesday.
	// - **Wednesday**: Wednesday.
	// - **Thursday**: Thursday.
	// - **Friday**: Friday.
	// - **Saturday**: Saturday.
	// - **Sunday**: Sunday.
	PreferredBackupPeriod string `pulumi:"preferredBackupPeriod"`
	// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
	PreferredBackupTime string `pulumi:"preferredBackupTime"`
	// Recovery point frequency. Value Description:
	// - **1**: Hourly.
	// - **2**: Every two hours.
	// - **4**: Every four hours.
	// - **8**: Every eight hours.
	RecoveryPointPeriod *string `pulumi:"recoveryPointPeriod"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// Data backup retention days.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The instance ID.
	// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	DbInstanceId pulumi.StringInput
	// Whether to enable automatic recovery points. Value Description:
	// - **true**: enabled.
	// - **false**: Closed.
	EnableRecoveryPoint pulumi.BoolPtrInput
	// Data backup cycle. Separate multiple values with commas (,). Value Description:
	// - **Monday**: Monday.
	// - **Tuesday**: Tuesday.
	// - **Wednesday**: Wednesday.
	// - **Thursday**: Thursday.
	// - **Friday**: Friday.
	// - **Saturday**: Saturday.
	// - **Sunday**: Sunday.
	PreferredBackupPeriod pulumi.StringInput
	// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
	PreferredBackupTime pulumi.StringInput
	// Recovery point frequency. Value Description:
	// - **1**: Hourly.
	// - **2**: Every two hours.
	// - **4**: Every four hours.
	// - **8**: Every eight hours.
	RecoveryPointPeriod pulumi.StringPtrInput
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

func (i *BackupPolicy) ToOutput(ctx context.Context) pulumix.Output[*BackupPolicy] {
	return pulumix.Output[*BackupPolicy]{
		OutputState: i.ToBackupPolicyOutputWithContext(ctx).OutputState,
	}
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

func (i BackupPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*BackupPolicy] {
	return pulumix.Output[[]*BackupPolicy]{
		OutputState: i.ToBackupPolicyArrayOutputWithContext(ctx).OutputState,
	}
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

func (i BackupPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BackupPolicy] {
	return pulumix.Output[map[string]*BackupPolicy]{
		OutputState: i.ToBackupPolicyMapOutputWithContext(ctx).OutputState,
	}
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

func (o BackupPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*BackupPolicy] {
	return pulumix.Output[*BackupPolicy]{
		OutputState: o.OutputState,
	}
}

// Data backup retention days.
func (o BackupPolicyOutput) BackupRetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.IntOutput { return v.BackupRetentionPeriod }).(pulumi.IntOutput)
}

// The instance ID.
// > **NOTE:**  You can call the DescribeDBInstances operation to view the details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
func (o BackupPolicyOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// Whether to enable automatic recovery points. Value Description:
// - **true**: enabled.
// - **false**: Closed.
func (o BackupPolicyOutput) EnableRecoveryPoint() pulumi.BoolOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolOutput { return v.EnableRecoveryPoint }).(pulumi.BoolOutput)
}

// Data backup cycle. Separate multiple values with commas (,). Value Description:
// - **Monday**: Monday.
// - **Tuesday**: Tuesday.
// - **Wednesday**: Wednesday.
// - **Thursday**: Thursday.
// - **Friday**: Friday.
// - **Saturday**: Saturday.
// - **Sunday**: Sunday.
func (o BackupPolicyOutput) PreferredBackupPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.PreferredBackupPeriod }).(pulumi.StringOutput)
}

// Data backup time. Format: HH:mmZ-HH:mmZ(UTC time).
func (o BackupPolicyOutput) PreferredBackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.PreferredBackupTime }).(pulumi.StringOutput)
}

// Recovery point frequency. Value Description:
// - **1**: Hourly.
// - **2**: Every two hours.
// - **4**: Every four hours.
// - **8**: Every eight hours.
func (o BackupPolicyOutput) RecoveryPointPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.RecoveryPointPeriod }).(pulumi.StringOutput)
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

func (o BackupPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BackupPolicy] {
	return pulumix.Output[[]*BackupPolicy]{
		OutputState: o.OutputState,
	}
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

func (o BackupPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BackupPolicy] {
	return pulumix.Output[map[string]*BackupPolicy]{
		OutputState: o.OutputState,
	}
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
