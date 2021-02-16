// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// ADB backup policy can be imported using the id or cluster id, e.g.
//
// ```sh
//  $ pulumi import alicloud:adb/backupPolicy:BackupPolicy example "am-12345678"
// ```
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Cluster backup retention days, Fixed for 7 days, not modified.
	BackupRetentionPeriod pulumi.StringOutput `pulumi:"backupRetentionPeriod"`
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
	PreferredBackupPeriods pulumi.StringArrayOutput `pulumi:"preferredBackupPeriods"`
	// ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime pulumi.StringOutput `pulumi:"preferredBackupTime"`
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
	var resource BackupPolicy
	err := ctx.RegisterResource("alicloud:adb/backupPolicy:BackupPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:adb/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Cluster backup retention days, Fixed for 7 days, not modified.
	BackupRetentionPeriod *string `pulumi:"backupRetentionPeriod"`
	// The Id of cluster that can run database.
	DbClusterId *string `pulumi:"dbClusterId"`
	// ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
	PreferredBackupPeriods []string `pulumi:"preferredBackupPeriods"`
	// ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime *string `pulumi:"preferredBackupTime"`
}

type BackupPolicyState struct {
	// Cluster backup retention days, Fixed for 7 days, not modified.
	BackupRetentionPeriod pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringPtrInput
	// ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
	PreferredBackupPeriods pulumi.StringArrayInput
	// ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// The Id of cluster that can run database.
	DbClusterId string `pulumi:"dbClusterId"`
	// ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
	PreferredBackupPeriods []string `pulumi:"preferredBackupPeriods"`
	// ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it.
	PreferredBackupTime string `pulumi:"preferredBackupTime"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringInput
	// ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].
	PreferredBackupPeriods pulumi.StringArrayInput
	// ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it.
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
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

func (i *BackupPolicy) ToBackupPolicyPtrOutput() BackupPolicyPtrOutput {
	return i.ToBackupPolicyPtrOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyPtrOutputWithContext(ctx context.Context) BackupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyPtrOutput)
}

type BackupPolicyPtrInput interface {
	pulumi.Input

	ToBackupPolicyPtrOutput() BackupPolicyPtrOutput
	ToBackupPolicyPtrOutputWithContext(ctx context.Context) BackupPolicyPtrOutput
}

type backupPolicyPtrType BackupPolicyArgs

func (*backupPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil))
}

func (i *backupPolicyPtrType) ToBackupPolicyPtrOutput() BackupPolicyPtrOutput {
	return i.ToBackupPolicyPtrOutputWithContext(context.Background())
}

func (i *backupPolicyPtrType) ToBackupPolicyPtrOutputWithContext(ctx context.Context) BackupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyPtrOutput)
}

// BackupPolicyArrayInput is an input type that accepts BackupPolicyArray and BackupPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPolicyArrayInput` via:
//
//          BackupPolicyArray{ BackupPolicyArgs{...} }
type BackupPolicyArrayInput interface {
	pulumi.Input

	ToBackupPolicyArrayOutput() BackupPolicyArrayOutput
	ToBackupPolicyArrayOutputWithContext(context.Context) BackupPolicyArrayOutput
}

type BackupPolicyArray []BackupPolicyInput

func (BackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BackupPolicy)(nil))
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
//          BackupPolicyMap{ "key": BackupPolicyArgs{...} }
type BackupPolicyMapInput interface {
	pulumi.Input

	ToBackupPolicyMapOutput() BackupPolicyMapOutput
	ToBackupPolicyMapOutputWithContext(context.Context) BackupPolicyMapOutput
}

type BackupPolicyMap map[string]BackupPolicyInput

func (BackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BackupPolicy)(nil))
}

func (i BackupPolicyMap) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return i.ToBackupPolicyMapOutputWithContext(context.Background())
}

func (i BackupPolicyMap) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMapOutput)
}

type BackupPolicyOutput struct {
	*pulumi.OutputState
}

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyPtrOutput() BackupPolicyPtrOutput {
	return o.ToBackupPolicyPtrOutputWithContext(context.Background())
}

func (o BackupPolicyOutput) ToBackupPolicyPtrOutputWithContext(ctx context.Context) BackupPolicyPtrOutput {
	return o.ApplyT(func(v BackupPolicy) *BackupPolicy {
		return &v
	}).(BackupPolicyPtrOutput)
}

type BackupPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (BackupPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil))
}

func (o BackupPolicyPtrOutput) ToBackupPolicyPtrOutput() BackupPolicyPtrOutput {
	return o
}

func (o BackupPolicyPtrOutput) ToBackupPolicyPtrOutputWithContext(ctx context.Context) BackupPolicyPtrOutput {
	return o
}

type BackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupPolicy)(nil))
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) Index(i pulumi.IntInput) BackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupPolicy {
		return vs[0].([]BackupPolicy)[vs[1].(int)]
	}).(BackupPolicyOutput)
}

type BackupPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackupPolicy)(nil))
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackupPolicy {
		return vs[0].(map[string]BackupPolicy)[vs[1].(string)]
	}).(BackupPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
	pulumi.RegisterOutputType(BackupPolicyPtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyMapOutput{})
}
