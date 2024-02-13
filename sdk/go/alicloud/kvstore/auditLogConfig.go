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

// Provides a Redis And Memcache (KVStore) Audit Log Config resource.
//
// > **NOTE:** Available since v1.130.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kvstore"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
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
//			defaultZones, err := kvstore.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
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
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := kvstore.NewInstance(ctx, "defaultInstance", &kvstore.InstanceArgs{
//				DbInstanceName:  pulumi.String(name),
//				VswitchId:       defaultSwitch.ID(),
//				ResourceGroupId: *pulumi.String(defaultResourceGroups.Ids[0]),
//				ZoneId:          *pulumi.String(defaultZones.Zones[0].Id),
//				InstanceClass:   pulumi.String("redis.master.large.default"),
//				InstanceType:    pulumi.String("Redis"),
//				EngineVersion:   pulumi.String("5.0"),
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
//			_, err = kvstore.NewAuditLogConfig(ctx, "example", &kvstore.AuditLogConfigArgs{
//				InstanceId: defaultInstance.ID(),
//				DbAudit:    pulumi.Bool(true),
//				Retention:  pulumi.Int(1),
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
// Redis And Memcache (KVStore) Audit Log Config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:kvstore/auditLogConfig:AuditLogConfig example <instance_id>
// ```
type AuditLogConfig struct {
	pulumi.CustomResourceState

	// Instance Creation Time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates Whether to Enable the Audit Log.  Valid value:
	// * true: Default Value, Open.
	// * false: Closed.
	//
	// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
	DbAudit pulumi.BoolPtrOutput `pulumi:"dbAudit"`
	// Instance ID, Call the Describeinstances Get.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Audit Log Retention Period Value: 1~365.
	//
	// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
	Retention pulumi.IntPtrOutput `pulumi:"retention"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAuditLogConfig registers a new resource with the given unique name, arguments, and options.
func NewAuditLogConfig(ctx *pulumi.Context,
	name string, args *AuditLogConfigArgs, opts ...pulumi.ResourceOption) (*AuditLogConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuditLogConfig
	err := ctx.RegisterResource("alicloud:kvstore/auditLogConfig:AuditLogConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditLogConfig gets an existing AuditLogConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditLogConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditLogConfigState, opts ...pulumi.ResourceOption) (*AuditLogConfig, error) {
	var resource AuditLogConfig
	err := ctx.ReadResource("alicloud:kvstore/auditLogConfig:AuditLogConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditLogConfig resources.
type auditLogConfigState struct {
	// Instance Creation Time.
	CreateTime *string `pulumi:"createTime"`
	// Indicates Whether to Enable the Audit Log.  Valid value:
	// * true: Default Value, Open.
	// * false: Closed.
	//
	// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
	DbAudit *bool `pulumi:"dbAudit"`
	// Instance ID, Call the Describeinstances Get.
	InstanceId *string `pulumi:"instanceId"`
	// Audit Log Retention Period Value: 1~365.
	//
	// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
	Retention *int `pulumi:"retention"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

type AuditLogConfigState struct {
	// Instance Creation Time.
	CreateTime pulumi.StringPtrInput
	// Indicates Whether to Enable the Audit Log.  Valid value:
	// * true: Default Value, Open.
	// * false: Closed.
	//
	// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
	DbAudit pulumi.BoolPtrInput
	// Instance ID, Call the Describeinstances Get.
	InstanceId pulumi.StringPtrInput
	// Audit Log Retention Period Value: 1~365.
	//
	// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
	Retention pulumi.IntPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
}

func (AuditLogConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditLogConfigState)(nil)).Elem()
}

type auditLogConfigArgs struct {
	// Indicates Whether to Enable the Audit Log.  Valid value:
	// * true: Default Value, Open.
	// * false: Closed.
	//
	// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
	DbAudit *bool `pulumi:"dbAudit"`
	// Instance ID, Call the Describeinstances Get.
	InstanceId string `pulumi:"instanceId"`
	// Audit Log Retention Period Value: 1~365.
	//
	// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
	Retention *int `pulumi:"retention"`
}

// The set of arguments for constructing a AuditLogConfig resource.
type AuditLogConfigArgs struct {
	// Indicates Whether to Enable the Audit Log.  Valid value:
	// * true: Default Value, Open.
	// * false: Closed.
	//
	// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
	DbAudit pulumi.BoolPtrInput
	// Instance ID, Call the Describeinstances Get.
	InstanceId pulumi.StringInput
	// Audit Log Retention Period Value: 1~365.
	//
	// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
	Retention pulumi.IntPtrInput
}

func (AuditLogConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditLogConfigArgs)(nil)).Elem()
}

type AuditLogConfigInput interface {
	pulumi.Input

	ToAuditLogConfigOutput() AuditLogConfigOutput
	ToAuditLogConfigOutputWithContext(ctx context.Context) AuditLogConfigOutput
}

func (*AuditLogConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfig)(nil)).Elem()
}

func (i *AuditLogConfig) ToAuditLogConfigOutput() AuditLogConfigOutput {
	return i.ToAuditLogConfigOutputWithContext(context.Background())
}

func (i *AuditLogConfig) ToAuditLogConfigOutputWithContext(ctx context.Context) AuditLogConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogConfigOutput)
}

// AuditLogConfigArrayInput is an input type that accepts AuditLogConfigArray and AuditLogConfigArrayOutput values.
// You can construct a concrete instance of `AuditLogConfigArrayInput` via:
//
//	AuditLogConfigArray{ AuditLogConfigArgs{...} }
type AuditLogConfigArrayInput interface {
	pulumi.Input

	ToAuditLogConfigArrayOutput() AuditLogConfigArrayOutput
	ToAuditLogConfigArrayOutputWithContext(context.Context) AuditLogConfigArrayOutput
}

type AuditLogConfigArray []AuditLogConfigInput

func (AuditLogConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditLogConfig)(nil)).Elem()
}

func (i AuditLogConfigArray) ToAuditLogConfigArrayOutput() AuditLogConfigArrayOutput {
	return i.ToAuditLogConfigArrayOutputWithContext(context.Background())
}

func (i AuditLogConfigArray) ToAuditLogConfigArrayOutputWithContext(ctx context.Context) AuditLogConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogConfigArrayOutput)
}

// AuditLogConfigMapInput is an input type that accepts AuditLogConfigMap and AuditLogConfigMapOutput values.
// You can construct a concrete instance of `AuditLogConfigMapInput` via:
//
//	AuditLogConfigMap{ "key": AuditLogConfigArgs{...} }
type AuditLogConfigMapInput interface {
	pulumi.Input

	ToAuditLogConfigMapOutput() AuditLogConfigMapOutput
	ToAuditLogConfigMapOutputWithContext(context.Context) AuditLogConfigMapOutput
}

type AuditLogConfigMap map[string]AuditLogConfigInput

func (AuditLogConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditLogConfig)(nil)).Elem()
}

func (i AuditLogConfigMap) ToAuditLogConfigMapOutput() AuditLogConfigMapOutput {
	return i.ToAuditLogConfigMapOutputWithContext(context.Background())
}

func (i AuditLogConfigMap) ToAuditLogConfigMapOutputWithContext(ctx context.Context) AuditLogConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogConfigMapOutput)
}

type AuditLogConfigOutput struct{ *pulumi.OutputState }

func (AuditLogConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfig)(nil)).Elem()
}

func (o AuditLogConfigOutput) ToAuditLogConfigOutput() AuditLogConfigOutput {
	return o
}

func (o AuditLogConfigOutput) ToAuditLogConfigOutputWithContext(ctx context.Context) AuditLogConfigOutput {
	return o
}

// Instance Creation Time.
func (o AuditLogConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates Whether to Enable the Audit Log.  Valid value:
// * true: Default Value, Open.
// * false: Closed.
//
// Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
func (o AuditLogConfigOutput) DbAudit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuditLogConfig) pulumi.BoolPtrOutput { return v.DbAudit }).(pulumi.BoolPtrOutput)
}

// Instance ID, Call the Describeinstances Get.
func (o AuditLogConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Audit Log Retention Period Value: 1~365.
//
// > **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
func (o AuditLogConfigOutput) Retention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuditLogConfig) pulumi.IntPtrOutput { return v.Retention }).(pulumi.IntPtrOutput)
}

// The status of the resource.
func (o AuditLogConfigOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogConfig) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AuditLogConfigArrayOutput struct{ *pulumi.OutputState }

func (AuditLogConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditLogConfig)(nil)).Elem()
}

func (o AuditLogConfigArrayOutput) ToAuditLogConfigArrayOutput() AuditLogConfigArrayOutput {
	return o
}

func (o AuditLogConfigArrayOutput) ToAuditLogConfigArrayOutputWithContext(ctx context.Context) AuditLogConfigArrayOutput {
	return o
}

func (o AuditLogConfigArrayOutput) Index(i pulumi.IntInput) AuditLogConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuditLogConfig {
		return vs[0].([]*AuditLogConfig)[vs[1].(int)]
	}).(AuditLogConfigOutput)
}

type AuditLogConfigMapOutput struct{ *pulumi.OutputState }

func (AuditLogConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditLogConfig)(nil)).Elem()
}

func (o AuditLogConfigMapOutput) ToAuditLogConfigMapOutput() AuditLogConfigMapOutput {
	return o
}

func (o AuditLogConfigMapOutput) ToAuditLogConfigMapOutputWithContext(ctx context.Context) AuditLogConfigMapOutput {
	return o
}

func (o AuditLogConfigMapOutput) MapIndex(k pulumi.StringInput) AuditLogConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuditLogConfig {
		return vs[0].(map[string]*AuditLogConfig)[vs[1].(string)]
	}).(AuditLogConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigInput)(nil)).Elem(), &AuditLogConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigArrayInput)(nil)).Elem(), AuditLogConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigMapInput)(nil)).Elem(), AuditLogConfigMap{})
	pulumi.RegisterOutputType(AuditLogConfigOutput{})
	pulumi.RegisterOutputType(AuditLogConfigArrayOutput{})
	pulumi.RegisterOutputType(AuditLogConfigMapOutput{})
}
