// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log store is a unit in Log Service to collect, store, and query the log data. Each log store belongs to a project,
// and each project can create multiple Logstores. [Refer to details](https://www.alibabacloud.com/help/doc-detail/48874.htm)
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewRandomInteger(ctx, "default", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			exampleProject, err := log.NewProject(ctx, "exampleProject", &log.ProjectArgs{
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = log.NewStore(ctx, "exampleStore", &log.StoreArgs{
//				Project:            exampleProject.Name,
//				ShardCount:         pulumi.Int(3),
//				AutoSplit:          pulumi.Bool(true),
//				MaxSplitShardCount: pulumi.Int(60),
//				AppendMeta:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// Encrypt Usage
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kms"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			region := "cn-hangzhou"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			exampleAccount, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = random.NewRandomInteger(ctx, "default", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
//				Description:         pulumi.String("terraform-example"),
//				PendingWindowInDays: pulumi.Int(7),
//				Status:              pulumi.String("Enabled"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleProject, err := log.NewProject(ctx, "exampleProject", &log.ProjectArgs{
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = log.NewStore(ctx, "exampleStore", &log.StoreArgs{
//				Project:            exampleProject.Name,
//				ShardCount:         pulumi.Int(1),
//				AutoSplit:          pulumi.Bool(true),
//				MaxSplitShardCount: pulumi.Int(60),
//				EncryptConf: &log.StoreEncryptConfArgs{
//					Enable:      pulumi.Bool(true),
//					EncryptType: pulumi.String("default"),
//					UserCmkInfo: &log.StoreEncryptConfUserCmkInfoArgs{
//						CmkKeyId: exampleKey.ID(),
//						Arn:      pulumi.String(fmt.Sprintf("acs:ram::%v:role/aliyunlogdefaultrole", exampleAccount.Id)),
//						RegionId: pulumi.String(region),
//					},
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
// ## Module Support
//
// You can use the existing sls module
// to create SLS project, store and store index one-click, like ECS instances.
//
// ## Import
//
// Log store can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:log/store:Store example tf-log:tf-log-store
//
// ```
type Store struct {
	pulumi.CustomResourceState

	// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
	AppendMeta pulumi.BoolPtrOutput `pulumi:"appendMeta"`
	// Determines whether to automatically split a shard. Default to `false`.
	AutoSplit pulumi.BoolPtrOutput `pulumi:"autoSplit"`
	// Determines whether to enable Web Tracking. Default `false`.
	EnableWebTracking pulumi.BoolPtrOutput `pulumi:"enableWebTracking"`
	// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
	EncryptConf StoreEncryptConfPtrOutput `pulumi:"encryptConf"`
	// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
	HotTtl pulumi.IntPtrOutput `pulumi:"hotTtl"`
	// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
	MaxSplitShardCount pulumi.IntPtrOutput `pulumi:"maxSplitShardCount"`
	// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The log store, which is unique in the same project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project name to the log store belongs.
	Project pulumi.StringOutput `pulumi:"project"`
	// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
	ShardCount pulumi.IntPtrOutput `pulumi:"shardCount"`
	// The shard attribute.
	Shards StoreShardArrayOutput `pulumi:"shards"`
	// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
	TelemetryType pulumi.StringPtrOutput `pulumi:"telemetryType"`
}

// NewStore registers a new resource with the given unique name, arguments, and options.
func NewStore(ctx *pulumi.Context,
	name string, args *StoreArgs, opts ...pulumi.ResourceOption) (*Store, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Store
	err := ctx.RegisterResource("alicloud:log/store:Store", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStore gets an existing Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoreState, opts ...pulumi.ResourceOption) (*Store, error) {
	var resource Store
	err := ctx.ReadResource("alicloud:log/store:Store", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Store resources.
type storeState struct {
	// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
	AppendMeta *bool `pulumi:"appendMeta"`
	// Determines whether to automatically split a shard. Default to `false`.
	AutoSplit *bool `pulumi:"autoSplit"`
	// Determines whether to enable Web Tracking. Default `false`.
	EnableWebTracking *bool `pulumi:"enableWebTracking"`
	// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
	EncryptConf *StoreEncryptConf `pulumi:"encryptConf"`
	// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
	HotTtl *int `pulumi:"hotTtl"`
	// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
	MaxSplitShardCount *int `pulumi:"maxSplitShardCount"`
	// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
	Mode *string `pulumi:"mode"`
	// The log store, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The project name to the log store belongs.
	Project *string `pulumi:"project"`
	// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
	ShardCount *int `pulumi:"shardCount"`
	// The shard attribute.
	Shards []StoreShard `pulumi:"shards"`
	// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
	TelemetryType *string `pulumi:"telemetryType"`
}

type StoreState struct {
	// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
	AppendMeta pulumi.BoolPtrInput
	// Determines whether to automatically split a shard. Default to `false`.
	AutoSplit pulumi.BoolPtrInput
	// Determines whether to enable Web Tracking. Default `false`.
	EnableWebTracking pulumi.BoolPtrInput
	// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
	EncryptConf StoreEncryptConfPtrInput
	// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
	HotTtl pulumi.IntPtrInput
	// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
	MaxSplitShardCount pulumi.IntPtrInput
	// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
	Mode pulumi.StringPtrInput
	// The log store, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The project name to the log store belongs.
	Project pulumi.StringPtrInput
	// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
	ShardCount pulumi.IntPtrInput
	// The shard attribute.
	Shards StoreShardArrayInput
	// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
	TelemetryType pulumi.StringPtrInput
}

func (StoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*storeState)(nil)).Elem()
}

type storeArgs struct {
	// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
	AppendMeta *bool `pulumi:"appendMeta"`
	// Determines whether to automatically split a shard. Default to `false`.
	AutoSplit *bool `pulumi:"autoSplit"`
	// Determines whether to enable Web Tracking. Default `false`.
	EnableWebTracking *bool `pulumi:"enableWebTracking"`
	// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
	EncryptConf *StoreEncryptConf `pulumi:"encryptConf"`
	// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
	HotTtl *int `pulumi:"hotTtl"`
	// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
	MaxSplitShardCount *int `pulumi:"maxSplitShardCount"`
	// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
	Mode *string `pulumi:"mode"`
	// The log store, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The project name to the log store belongs.
	Project string `pulumi:"project"`
	// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
	ShardCount *int `pulumi:"shardCount"`
	// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
	TelemetryType *string `pulumi:"telemetryType"`
}

// The set of arguments for constructing a Store resource.
type StoreArgs struct {
	// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
	AppendMeta pulumi.BoolPtrInput
	// Determines whether to automatically split a shard. Default to `false`.
	AutoSplit pulumi.BoolPtrInput
	// Determines whether to enable Web Tracking. Default `false`.
	EnableWebTracking pulumi.BoolPtrInput
	// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
	EncryptConf StoreEncryptConfPtrInput
	// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
	HotTtl pulumi.IntPtrInput
	// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
	MaxSplitShardCount pulumi.IntPtrInput
	// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
	Mode pulumi.StringPtrInput
	// The log store, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The project name to the log store belongs.
	Project pulumi.StringInput
	// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
	ShardCount pulumi.IntPtrInput
	// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
	TelemetryType pulumi.StringPtrInput
}

func (StoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storeArgs)(nil)).Elem()
}

type StoreInput interface {
	pulumi.Input

	ToStoreOutput() StoreOutput
	ToStoreOutputWithContext(ctx context.Context) StoreOutput
}

func (*Store) ElementType() reflect.Type {
	return reflect.TypeOf((**Store)(nil)).Elem()
}

func (i *Store) ToStoreOutput() StoreOutput {
	return i.ToStoreOutputWithContext(context.Background())
}

func (i *Store) ToStoreOutputWithContext(ctx context.Context) StoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreOutput)
}

// StoreArrayInput is an input type that accepts StoreArray and StoreArrayOutput values.
// You can construct a concrete instance of `StoreArrayInput` via:
//
//	StoreArray{ StoreArgs{...} }
type StoreArrayInput interface {
	pulumi.Input

	ToStoreArrayOutput() StoreArrayOutput
	ToStoreArrayOutputWithContext(context.Context) StoreArrayOutput
}

type StoreArray []StoreInput

func (StoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Store)(nil)).Elem()
}

func (i StoreArray) ToStoreArrayOutput() StoreArrayOutput {
	return i.ToStoreArrayOutputWithContext(context.Background())
}

func (i StoreArray) ToStoreArrayOutputWithContext(ctx context.Context) StoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreArrayOutput)
}

// StoreMapInput is an input type that accepts StoreMap and StoreMapOutput values.
// You can construct a concrete instance of `StoreMapInput` via:
//
//	StoreMap{ "key": StoreArgs{...} }
type StoreMapInput interface {
	pulumi.Input

	ToStoreMapOutput() StoreMapOutput
	ToStoreMapOutputWithContext(context.Context) StoreMapOutput
}

type StoreMap map[string]StoreInput

func (StoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Store)(nil)).Elem()
}

func (i StoreMap) ToStoreMapOutput() StoreMapOutput {
	return i.ToStoreMapOutputWithContext(context.Background())
}

func (i StoreMap) ToStoreMapOutputWithContext(ctx context.Context) StoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreMapOutput)
}

type StoreOutput struct{ *pulumi.OutputState }

func (StoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Store)(nil)).Elem()
}

func (o StoreOutput) ToStoreOutput() StoreOutput {
	return o
}

func (o StoreOutput) ToStoreOutputWithContext(ctx context.Context) StoreOutput {
	return o
}

// Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to `true`.
func (o StoreOutput) AppendMeta() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.BoolPtrOutput { return v.AppendMeta }).(pulumi.BoolPtrOutput)
}

// Determines whether to automatically split a shard. Default to `false`.
func (o StoreOutput) AutoSplit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.BoolPtrOutput { return v.AutoSplit }).(pulumi.BoolPtrOutput)
}

// Determines whether to enable Web Tracking. Default `false`.
func (o StoreOutput) EnableWebTracking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.BoolPtrOutput { return v.EnableWebTracking }).(pulumi.BoolPtrOutput)
}

// Encrypted storage of data, providing data static protection capability, `encryptConf` can be updated since 1.188.0+ (only `enable` change is supported when updating logstore)
func (o StoreOutput) EncryptConf() StoreEncryptConfPtrOutput {
	return o.ApplyT(func(v *Store) StoreEncryptConfPtrOutput { return v.EncryptConf }).(StoreEncryptConfPtrOutput)
}

// The ttl of hot storage. Default to `30`, at least `30`, hot storage ttl must be less than ttl.
func (o StoreOutput) HotTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.IntPtrOutput { return v.HotTtl }).(pulumi.IntPtrOutput)
}

// The maximum number of shards for automatic split, which is in the range of 1 to 256. You must specify this parameter when autoSplit is true.
func (o StoreOutput) MaxSplitShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.IntPtrOutput { return v.MaxSplitShardCount }).(pulumi.IntPtrOutput)
}

// The mode of storage. Default to `standard`, must be `standard` or `query`, `mode` is only valid when creating, can't be changed after created.
func (o StoreOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Store) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The log store, which is unique in the same project.
func (o StoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Store) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project name to the log store belongs.
func (o StoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Store) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The data retention time (in days). Valid values: [1-3650]. Default to `30`. Log store data will be stored permanently when the value is `3650`.
func (o StoreOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.IntPtrOutput { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)
func (o StoreOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.IntPtrOutput { return v.ShardCount }).(pulumi.IntPtrOutput)
}

// The shard attribute.
func (o StoreOutput) Shards() StoreShardArrayOutput {
	return o.ApplyT(func(v *Store) StoreShardArrayOutput { return v.Shards }).(StoreShardArrayOutput)
}

// Determines whether store type is metric. `Metrics` means metric store, empty means log store.
func (o StoreOutput) TelemetryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Store) pulumi.StringPtrOutput { return v.TelemetryType }).(pulumi.StringPtrOutput)
}

type StoreArrayOutput struct{ *pulumi.OutputState }

func (StoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Store)(nil)).Elem()
}

func (o StoreArrayOutput) ToStoreArrayOutput() StoreArrayOutput {
	return o
}

func (o StoreArrayOutput) ToStoreArrayOutputWithContext(ctx context.Context) StoreArrayOutput {
	return o
}

func (o StoreArrayOutput) Index(i pulumi.IntInput) StoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Store {
		return vs[0].([]*Store)[vs[1].(int)]
	}).(StoreOutput)
}

type StoreMapOutput struct{ *pulumi.OutputState }

func (StoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Store)(nil)).Elem()
}

func (o StoreMapOutput) ToStoreMapOutput() StoreMapOutput {
	return o
}

func (o StoreMapOutput) ToStoreMapOutputWithContext(ctx context.Context) StoreMapOutput {
	return o
}

func (o StoreMapOutput) MapIndex(k pulumi.StringInput) StoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Store {
		return vs[0].(map[string]*Store)[vs[1].(string)]
	}).(StoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StoreInput)(nil)).Elem(), &Store{})
	pulumi.RegisterInputType(reflect.TypeOf((*StoreArrayInput)(nil)).Elem(), StoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StoreMapInput)(nil)).Elem(), StoreMap{})
	pulumi.RegisterOutputType(StoreOutput{})
	pulumi.RegisterOutputType(StoreArrayOutput{})
	pulumi.RegisterOutputType(StoreMapOutput{})
}
