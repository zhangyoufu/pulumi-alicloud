// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB cluster resource. A PolarDB cluster is an isolated database
// environment in the cloud. A PolarDB cluster can contain multiple user-created
// databases.
//
// > **NOTE:** Available in v1.66.0+.
//
// ## Example Usage
// ### Create a PolarDB MySQL cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "polardbClusterconfig"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		creation := "PolarDB"
// 		if param := cfg.Get("creation"); param != "" {
// 			creation = param
// 		}
// 		opt0 := creation
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/24"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = polardb.NewCluster(ctx, "defaultCluster", &polardb.ClusterArgs{
// 			DbType:      pulumi.String("MySQL"),
// 			DbVersion:   pulumi.String("5.6"),
// 			DbNodeClass: pulumi.String("polar.mysql.x4.medium"),
// 			PayType:     pulumi.String("PostPaid"),
// 			Description: pulumi.String(name),
// 			VswitchId:   defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// PolarDB cluster can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:polardb/cluster:Cluster example pc-abc12345678
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrOutput `pulumi:"autoRenewPeriod"`
	// Specifies whether to enable or disable SQL data collector. Valid values are `Enable`, `Disabled`.
	CollectorStatus pulumi.StringOutput `pulumi:"collectorStatus"`
	// (Available in 1.81.0+) PolarDB cluster connection string. When securityIps is configured, the address of cluster type endpoint will be returned, and if only "127.0.0.1" is configured, it will also be an empty string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringOutput `pulumi:"dbNodeClass"`
	// Number of the PolarDB cluster nodes, default is 2(Each cluster must contain at least a primary node and a read-only node). Add/remove nodes by modifying this parameter, valid values: [2~16].\
	// **NOTE:** To avoid adding or removing multiple read-only nodes by mistake, the system allows you to add or remove one read-only node at a time.
	DbNodeCount pulumi.IntPtrOutput `pulumi:"dbNodeCount"`
	// Database type. Value options: MySQL, Oracle, PostgreSQL.
	DbType pulumi.StringOutput `pulumi:"dbType"`
	// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
	DbVersion pulumi.StringOutput `pulumi:"dbVersion"`
	// The description of cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringOutput `pulumi:"maintainTime"`
	// Use as `dbNodeClass` change class, define upgrade or downgrade. Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
	ModifyType pulumi.StringPtrOutput `pulumi:"modifyType"`
	// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/98122.htm) .
	Parameters ClusterParameterArrayOutput `pulumi:"parameters"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrOutput `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrOutput `pulumi:"renewalStatus"`
	// The ID of resource group which the PolarDB cluster belongs. If not specified, then it belongs to the default resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayOutput `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.\
	// **NOTE:** If vswitchId is not specified, system will get a vswitch belongs to the user automatically.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster. it supports multiple zone.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbNodeClass == nil {
		return nil, errors.New("invalid value for required argument 'DbNodeClass'")
	}
	if args.DbType == nil {
		return nil, errors.New("invalid value for required argument 'DbType'")
	}
	if args.DbVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbVersion'")
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:polardb/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:polardb/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// Specifies whether to enable or disable SQL data collector. Valid values are `Enable`, `Disabled`.
	CollectorStatus *string `pulumi:"collectorStatus"`
	// (Available in 1.81.0+) PolarDB cluster connection string. When securityIps is configured, the address of cluster type endpoint will be returned, and if only "127.0.0.1" is configured, it will also be an empty string.
	ConnectionString *string `pulumi:"connectionString"`
	// The dbNodeClass of cluster node.
	DbNodeClass *string `pulumi:"dbNodeClass"`
	// Number of the PolarDB cluster nodes, default is 2(Each cluster must contain at least a primary node and a read-only node). Add/remove nodes by modifying this parameter, valid values: [2~16].\
	// **NOTE:** To avoid adding or removing multiple read-only nodes by mistake, the system allows you to add or remove one read-only node at a time.
	DbNodeCount *int `pulumi:"dbNodeCount"`
	// Database type. Value options: MySQL, Oracle, PostgreSQL.
	DbType *string `pulumi:"dbType"`
	// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
	DbVersion *string `pulumi:"dbVersion"`
	// The description of cluster.
	Description *string `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// Use as `dbNodeClass` change class, define upgrade or downgrade. Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
	ModifyType *string `pulumi:"modifyType"`
	// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/98122.htm) .
	Parameters []ClusterParameter `pulumi:"parameters"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType *string `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The ID of resource group which the PolarDB cluster belongs. If not specified, then it belongs to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.\
	// **NOTE:** If vswitchId is not specified, system will get a vswitch belongs to the user automatically.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster. it supports multiple zone.
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// Specifies whether to enable or disable SQL data collector. Valid values are `Enable`, `Disabled`.
	CollectorStatus pulumi.StringPtrInput
	// (Available in 1.81.0+) PolarDB cluster connection string. When securityIps is configured, the address of cluster type endpoint will be returned, and if only "127.0.0.1" is configured, it will also be an empty string.
	ConnectionString pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringPtrInput
	// Number of the PolarDB cluster nodes, default is 2(Each cluster must contain at least a primary node and a read-only node). Add/remove nodes by modifying this parameter, valid values: [2~16].\
	// **NOTE:** To avoid adding or removing multiple read-only nodes by mistake, the system allows you to add or remove one read-only node at a time.
	DbNodeCount pulumi.IntPtrInput
	// Database type. Value options: MySQL, Oracle, PostgreSQL.
	DbType pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
	DbVersion pulumi.StringPtrInput
	// The description of cluster.
	Description pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// Use as `dbNodeClass` change class, define upgrade or downgrade. Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
	ModifyType pulumi.StringPtrInput
	// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/98122.htm) .
	Parameters ClusterParameterArrayInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// The ID of resource group which the PolarDB cluster belongs. If not specified, then it belongs to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.\
	// **NOTE:** If vswitchId is not specified, system will get a vswitch belongs to the user automatically.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster. it supports multiple zone.
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// Specifies whether to enable or disable SQL data collector. Valid values are `Enable`, `Disabled`.
	CollectorStatus *string `pulumi:"collectorStatus"`
	// The dbNodeClass of cluster node.
	DbNodeClass string `pulumi:"dbNodeClass"`
	// Number of the PolarDB cluster nodes, default is 2(Each cluster must contain at least a primary node and a read-only node). Add/remove nodes by modifying this parameter, valid values: [2~16].\
	// **NOTE:** To avoid adding or removing multiple read-only nodes by mistake, the system allows you to add or remove one read-only node at a time.
	DbNodeCount *int `pulumi:"dbNodeCount"`
	// Database type. Value options: MySQL, Oracle, PostgreSQL.
	DbType string `pulumi:"dbType"`
	// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
	DbVersion string `pulumi:"dbVersion"`
	// The description of cluster.
	Description *string `pulumi:"description"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// Use as `dbNodeClass` change class, define upgrade or downgrade. Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
	ModifyType *string `pulumi:"modifyType"`
	// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/98122.htm) .
	Parameters []ClusterParameter `pulumi:"parameters"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType *string `pulumi:"payType"`
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// The ID of resource group which the PolarDB cluster belongs. If not specified, then it belongs to the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.\
	// **NOTE:** If vswitchId is not specified, system will get a vswitch belongs to the user automatically.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB cluster. it supports multiple zone.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Auto-renewal period of an cluster, in the unit of the month. It is valid when payType is `PrePaid`. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// Specifies whether to enable or disable SQL data collector. Valid values are `Enable`, `Disabled`.
	CollectorStatus pulumi.StringPtrInput
	// The dbNodeClass of cluster node.
	DbNodeClass pulumi.StringInput
	// Number of the PolarDB cluster nodes, default is 2(Each cluster must contain at least a primary node and a read-only node). Add/remove nodes by modifying this parameter, valid values: [2~16].\
	// **NOTE:** To avoid adding or removing multiple read-only nodes by mistake, the system allows you to add or remove one read-only node at a time.
	DbNodeCount pulumi.IntPtrInput
	// Database type. Value options: MySQL, Oracle, PostgreSQL.
	DbType pulumi.StringInput
	// Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) `DBVersion`.
	DbVersion pulumi.StringInput
	// The description of cluster.
	Description pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// Use as `dbNodeClass` change class, define upgrade or downgrade. Valid values are `Upgrade`, `Downgrade`, Default to `Upgrade`.
	ModifyType pulumi.StringPtrInput
	// Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/98122.htm) .
	Parameters ClusterParameterArrayInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`. Currently, the resource can not supports change pay type.
	PayType pulumi.StringPtrInput
	// The duration that you will buy DB cluster (in month). It is valid when payType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// Valid values are `AutoRenewal`, `Normal`, `NotRenewal`, Default to `NotRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// The ID of resource group which the PolarDB cluster belongs. If not specified, then it belongs to the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.\
	// **NOTE:** If vswitchId is not specified, system will get a vswitch belongs to the user automatically.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB cluster. it supports multiple zone.
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
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

type ClusterPtrInput interface {
	pulumi.Input

	ToClusterPtrOutput() ClusterPtrOutput
	ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput
}

type clusterPtrType ClusterArgs

func (*clusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (i *clusterPtrType) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *clusterPtrType) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Cluster)(nil))
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
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Cluster)(nil))
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o.ToClusterPtrOutputWithContext(context.Background())
}

func (o ClusterOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o.ApplyT(func(v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cluster)(nil))
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].([]Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cluster)(nil))
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].(map[string]Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
