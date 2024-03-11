// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:**  This resource  has been deprecated from version `1.104.0`. Please use resource alicloud_kvstore_instance.
//
// Provides a backup policy for ApsaraDB Redis / Memcache instance resource.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kvstore"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "kvstorebackuppolicyvpc"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := kvstore.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := kvstore.NewInstance(ctx, "defaultInstance", &kvstore.InstanceArgs{
//				DbInstanceName: pulumi.String(name),
//				VswitchId:      defaultSwitch.ID(),
//				ZoneId:         *pulumi.String(defaultZones.Zones[0].Id),
//				InstanceClass:  pulumi.String("redis.master.large.default"),
//				InstanceType:   pulumi.String("Redis"),
//				EngineVersion:  pulumi.String("5.0"),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("10.23.12.24"),
//				},
//				Config: pulumi.Map{
//					"appendonly":             pulumi.Any("yes"),
//					"lazyfree-lazy-eviction": pulumi.Any("yes"),
//				},
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kvstore.NewBackupPolicy(ctx, "defaultBackupPolicy", &kvstore.BackupPolicyArgs{
//				InstanceId: defaultInstance.ID(),
//				BackupPeriods: pulumi.StringArray{
//					pulumi.String("Tuesday"),
//					pulumi.String("Wednesday"),
//				},
//				BackupTime: pulumi.String("10:00Z-11:00Z"),
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
// KVStore backup policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:kvstore/backupPolicy:BackupPolicy example r-abc12345678
// ```
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	BackupPeriods pulumi.StringArrayOutput `pulumi:"backupPeriods"`
	// Backup time, in the format of HH:mmZ- HH:mm Z
	BackupTime pulumi.StringPtrOutput `pulumi:"backupTime"`
	// The id of ApsaraDB for Redis or Memcache intance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("alicloud:kvstore/backupPolicy:BackupPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:kvstore/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	BackupPeriods []string `pulumi:"backupPeriods"`
	// Backup time, in the format of HH:mmZ- HH:mm Z
	BackupTime *string `pulumi:"backupTime"`
	// The id of ApsaraDB for Redis or Memcache intance.
	InstanceId *string `pulumi:"instanceId"`
}

type BackupPolicyState struct {
	// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	BackupPeriods pulumi.StringArrayInput
	// Backup time, in the format of HH:mmZ- HH:mm Z
	BackupTime pulumi.StringPtrInput
	// The id of ApsaraDB for Redis or Memcache intance.
	InstanceId pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	BackupPeriods []string `pulumi:"backupPeriods"`
	// Backup time, in the format of HH:mmZ- HH:mm Z
	BackupTime *string `pulumi:"backupTime"`
	// The id of ApsaraDB for Redis or Memcache intance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	BackupPeriods pulumi.StringArrayInput
	// Backup time, in the format of HH:mmZ- HH:mm Z
	BackupTime pulumi.StringPtrInput
	// The id of ApsaraDB for Redis or Memcache intance.
	InstanceId pulumi.StringInput
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

// Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
func (o BackupPolicyOutput) BackupPeriods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringArrayOutput { return v.BackupPeriods }).(pulumi.StringArrayOutput)
}

// Backup time, in the format of HH:mmZ- HH:mm Z
func (o BackupPolicyOutput) BackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.BackupTime }).(pulumi.StringPtrOutput)
}

// The id of ApsaraDB for Redis or Memcache intance.
func (o BackupPolicyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
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
