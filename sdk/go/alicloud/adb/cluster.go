// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package adb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ADB cluster resource. An ADB cluster is an isolated database
// environment in the cloud. An ADB cluster can contain multiple user-created
// databases.
//
// > **DEPRECATED:**  This resource  has been deprecated from version `1.121.0`. Please use new resource alicloud_adb_db_cluster.
//
// > **NOTE:** Available in v1.71.0+.
//
// ## Example Usage
//
// ### Create a ADB MySQL cluster
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/adb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "adbClusterconfig"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			creation := "ADB"
//			if param := cfg.Get("creation"); param != "" {
//				creation = param
//			}
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef(creation),
//			}, nil)
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
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = adb.NewCluster(ctx, "default", &adb.ClusterArgs{
//				DbClusterVersion:  pulumi.String("3.0"),
//				DbClusterCategory: pulumi.String("Cluster"),
//				DbNodeClass:       pulumi.String("C8"),
//				DbNodeCount:       pulumi.Int(2),
//				DbNodeStorage:     pulumi.Int(200),
//				PayType:           pulumi.String("PostPaid"),
//				Description:       pulumi.String(name),
//				VswitchId:         defaultSwitch.ID(),
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
// ADB cluster can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:adb/cluster:Cluster example am-abc12345678
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntOutput       `pulumi:"autoRenewPeriod"`
	ComputeResource pulumi.StringPtrOutput `pulumi:"computeResource"`
	// (Available in 1.93.0+) The connection string of the ADB cluster.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringOutput `pulumi:"dbClusterCategory"`
	// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
	DbClusterClass pulumi.StringPtrOutput `pulumi:"dbClusterClass"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrOutput `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringOutput `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntOutput `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntOutput `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description           pulumi.StringOutput    `pulumi:"description"`
	DiskEncryption        pulumi.BoolPtrOutput   `pulumi:"diskEncryption"`
	DiskPerformanceLevel  pulumi.StringOutput    `pulumi:"diskPerformanceLevel"`
	ElasticIoResource     pulumi.IntOutput       `pulumi:"elasticIoResource"`
	ElasticIoResourceSize pulumi.StringOutput    `pulumi:"elasticIoResourceSize"`
	KmsId                 pulumi.StringPtrOutput `pulumi:"kmsId"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringOutput    `pulumi:"maintainTime"`
	Mode         pulumi.StringOutput    `pulumi:"mode"`
	ModifyType   pulumi.StringPtrOutput `pulumi:"modifyType"`
	// Field `payType` has been deprecated. New field `paymentType` instead.
	//
	// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	PayType pulumi.StringOutput `pulumi:"payType"`
	// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// (Available in 1.196.0+) The connection port of the ADB cluster.
	Port pulumi.StringOutput `pulumi:"port"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus   pulumi.StringOutput `pulumi:"renewalStatus"`
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayOutput `pulumi:"securityIps"`
	Status      pulumi.StringOutput      `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	//
	// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
	Tags  pulumi.StringMapOutput `pulumi:"tags"`
	VpcId pulumi.StringOutput    `pulumi:"vpcId"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterCategory == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterCategory'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("alicloud:adb/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("alicloud:adb/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int    `pulumi:"autoRenewPeriod"`
	ComputeResource *string `pulumi:"computeResource"`
	// (Available in 1.93.0+) The connection string of the ADB cluster.
	ConnectionString *string `pulumi:"connectionString"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory *string `pulumi:"dbClusterCategory"`
	// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
	DbClusterClass *string `pulumi:"dbClusterClass"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount *int `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage *int `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description           *string `pulumi:"description"`
	DiskEncryption        *bool   `pulumi:"diskEncryption"`
	DiskPerformanceLevel  *string `pulumi:"diskPerformanceLevel"`
	ElasticIoResource     *int    `pulumi:"elasticIoResource"`
	ElasticIoResourceSize *string `pulumi:"elasticIoResourceSize"`
	KmsId                 *string `pulumi:"kmsId"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	Mode         *string `pulumi:"mode"`
	ModifyType   *string `pulumi:"modifyType"`
	// Field `payType` has been deprecated. New field `paymentType` instead.
	//
	// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	PayType *string `pulumi:"payType"`
	// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// (Available in 1.196.0+) The connection port of the ADB cluster.
	Port *string `pulumi:"port"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus   *string `pulumi:"renewalStatus"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	Status      *string  `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	//
	// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
	Tags  map[string]string `pulumi:"tags"`
	VpcId *string           `pulumi:"vpcId"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	ComputeResource pulumi.StringPtrInput
	// (Available in 1.93.0+) The connection string of the ADB cluster.
	ConnectionString pulumi.StringPtrInput
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringPtrInput
	// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
	DbClusterClass pulumi.StringPtrInput
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringPtrInput
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntPtrInput
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntPtrInput
	// The description of cluster.
	Description           pulumi.StringPtrInput
	DiskEncryption        pulumi.BoolPtrInput
	DiskPerformanceLevel  pulumi.StringPtrInput
	ElasticIoResource     pulumi.IntPtrInput
	ElasticIoResourceSize pulumi.StringPtrInput
	KmsId                 pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	Mode         pulumi.StringPtrInput
	ModifyType   pulumi.StringPtrInput
	// Field `payType` has been deprecated. New field `paymentType` instead.
	//
	// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	PayType pulumi.StringPtrInput
	// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// (Available in 1.196.0+) The connection port of the ADB cluster.
	Port pulumi.StringPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus   pulumi.StringPtrInput
	ResourceGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	Status      pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	//
	// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
	Tags  pulumi.StringMapInput
	VpcId pulumi.StringPtrInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int    `pulumi:"autoRenewPeriod"`
	ComputeResource *string `pulumi:"computeResource"`
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory string `pulumi:"dbClusterCategory"`
	// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
	DbClusterClass *string `pulumi:"dbClusterClass"`
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion *string `pulumi:"dbClusterVersion"`
	// The dbNodeClass of cluster node.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	// The dbNodeCount of cluster node.
	DbNodeCount *int `pulumi:"dbNodeCount"`
	// The dbNodeStorage of cluster node.
	DbNodeStorage *int `pulumi:"dbNodeStorage"`
	// The description of cluster.
	Description           *string `pulumi:"description"`
	DiskEncryption        *bool   `pulumi:"diskEncryption"`
	DiskPerformanceLevel  *string `pulumi:"diskPerformanceLevel"`
	ElasticIoResource     *int    `pulumi:"elasticIoResource"`
	ElasticIoResourceSize *string `pulumi:"elasticIoResourceSize"`
	KmsId                 *string `pulumi:"kmsId"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	Mode         string  `pulumi:"mode"`
	ModifyType   *string `pulumi:"modifyType"`
	// Field `payType` has been deprecated. New field `paymentType` instead.
	//
	// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	PayType *string `pulumi:"payType"`
	// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
	PaymentType *string `pulumi:"paymentType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus   *string `pulumi:"renewalStatus"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	//
	// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
	Tags  map[string]string `pulumi:"tags"`
	VpcId *string           `pulumi:"vpcId"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	ComputeResource pulumi.StringPtrInput
	// Cluster category. Value options: `Basic`, `Cluster`.
	DbClusterCategory pulumi.StringInput
	// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
	DbClusterClass pulumi.StringPtrInput
	// Cluster version. Value options: `3.0`, Default to `3.0`.
	DbClusterVersion pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringPtrInput
	// The dbNodeCount of cluster node.
	DbNodeCount pulumi.IntPtrInput
	// The dbNodeStorage of cluster node.
	DbNodeStorage pulumi.IntPtrInput
	// The description of cluster.
	Description           pulumi.StringPtrInput
	DiskEncryption        pulumi.BoolPtrInput
	DiskPerformanceLevel  pulumi.StringPtrInput
	ElasticIoResource     pulumi.IntPtrInput
	ElasticIoResourceSize pulumi.StringPtrInput
	KmsId                 pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	Mode         pulumi.StringInput
	ModifyType   pulumi.StringPtrInput
	// Field `payType` has been deprecated. New field `paymentType` instead.
	//
	// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
	PayType pulumi.StringPtrInput
	// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
	PaymentType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus   pulumi.StringPtrInput
	ResourceGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	//
	// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
	Tags  pulumi.StringMapInput
	VpcId pulumi.StringPtrInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster.
	ZoneId pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
func (o ClusterOutput) AutoRenewPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.AutoRenewPeriod }).(pulumi.IntOutput)
}

func (o ClusterOutput) ComputeResource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ComputeResource }).(pulumi.StringPtrOutput)
}

// (Available in 1.93.0+) The connection string of the ADB cluster.
func (o ClusterOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Cluster category. Value options: `Basic`, `Cluster`.
func (o ClusterOutput) DbClusterCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DbClusterCategory }).(pulumi.StringOutput)
}

// Deprecated: It duplicates with attribute dbNodeClass and is deprecated from 1.121.2.
func (o ClusterOutput) DbClusterClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.DbClusterClass }).(pulumi.StringPtrOutput)
}

// Cluster version. Value options: `3.0`, Default to `3.0`.
func (o ClusterOutput) DbClusterVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.DbClusterVersion }).(pulumi.StringPtrOutput)
}

// The dbNodeClass of cluster node.
func (o ClusterOutput) DbNodeClass() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DbNodeClass }).(pulumi.StringOutput)
}

// The dbNodeCount of cluster node.
func (o ClusterOutput) DbNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.DbNodeCount }).(pulumi.IntOutput)
}

// The dbNodeStorage of cluster node.
func (o ClusterOutput) DbNodeStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.DbNodeStorage }).(pulumi.IntOutput)
}

// The description of cluster.
func (o ClusterOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ClusterOutput) DiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.DiskEncryption }).(pulumi.BoolPtrOutput)
}

func (o ClusterOutput) DiskPerformanceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DiskPerformanceLevel }).(pulumi.StringOutput)
}

func (o ClusterOutput) ElasticIoResource() pulumi.IntOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntOutput { return v.ElasticIoResource }).(pulumi.IntOutput)
}

func (o ClusterOutput) ElasticIoResourceSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ElasticIoResourceSize }).(pulumi.StringOutput)
}

func (o ClusterOutput) KmsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.KmsId }).(pulumi.StringPtrOutput)
}

// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
func (o ClusterOutput) MaintainTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.MaintainTime }).(pulumi.StringOutput)
}

func (o ClusterOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

func (o ClusterOutput) ModifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ModifyType }).(pulumi.StringPtrOutput)
}

// Field `payType` has been deprecated. New field `paymentType` instead.
//
// Deprecated: Attribute 'pay_type' has been deprecated from the provider version 1.166.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead.
func (o ClusterOutput) PayType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PayType }).(pulumi.StringOutput)
}

// The payment type of the resource. Valid values are `PayAsYouGo` and `Subscription`. Default to `PayAsYouGo`. **Note:** The `paymentType` supports updating from v1.166.0+.
func (o ClusterOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
func (o ClusterOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// (Available in 1.196.0+) The connection port of the ADB cluster.
func (o ClusterOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
func (o ClusterOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

func (o ClusterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
func (o ClusterOutput) SecurityIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.SecurityIps }).(pulumi.StringArrayOutput)
}

func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
//
// > **NOTE:** Because of data backup and migration, change DB cluster type and storage would cost 15~30 minutes. Please make full preparation before changing them.
func (o ClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ClusterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The virtual switch ID to launch DB instances in one VPC.
func (o ClusterOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VswitchId }).(pulumi.StringPtrOutput)
}

// The Zone to launch the DB cluster.
func (o ClusterOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
