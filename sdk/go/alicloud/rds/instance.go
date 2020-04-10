// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS instance resource. A DB instance is an isolated database
// environment in the cloud. A DB instance can contain multiple user-created
// databases.
type Instance struct {
	pulumi.CustomResourceState

	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrOutput `pulumi:"autoRenewPeriod"`
	// The upgrade method to use. Valid values:
	// - Auto: Instances are automatically upgraded to a higher minor version.
	// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
	AutoUpgradeMinorVersion pulumi.StringOutput `pulumi:"autoUpgradeMinorVersion"`
	// RDS database connection string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The storage type of the instance. Valid values:
	// - local_ssd: specifies to use local SSDs. This value is recommended.
	// - cloud_ssd: specifies to use standard SSDs.
	// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
	DbInstanceStorageType pulumi.StringOutput `pulumi:"dbInstanceStorageType"`
	// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Set it to true to make some parameter efficient when modifying them. Default to false.
	ForceRestart pulumi.BoolPtrOutput `pulumi:"forceRestart"`
	// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// The name of DB instance. It a string of 2 to 256 characters.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// User-defined DB instance storage space. Value range:
	// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
	// - [20,1000] for MySQL 5.7 basic single node edition;
	// - [10, 2000] for SQL Server 2008R2;
	// - [20,2000] for SQL Server 2012 basic single node edition
	// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
	InstanceStorage pulumi.IntOutput `pulumi:"instanceStorage"`
	// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringOutput `pulumi:"maintainTime"`
	// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
	MonitoringPeriod pulumi.IntOutput `pulumi:"monitoringPeriod"`
	// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
	Parameters InstanceParameterArrayOutput `pulumi:"parameters"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// RDS database connection port.
	Port pulumi.StringOutput `pulumi:"port"`
	// It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
	SecurityIpMode pulumi.StringPtrOutput `pulumi:"securityIpMode"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayOutput `pulumi:"securityIps"`
	// The sql collector keep time of the instance. Valid values are `1`, `30`, `180`, `365`, `1095`, `1825`, `1` is the initial value, and can't change it to `1`.
	SqlCollectorConfigValue pulumi.IntPtrOutput `pulumi:"sqlCollectorConfigValue"`
	// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
	SqlCollectorStatus pulumi.StringPtrOutput `pulumi:"sqlCollectorStatus"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.Engine == nil {
		return nil, errors.New("missing required argument 'Engine'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	if args == nil || args.InstanceStorage == nil {
		return nil, errors.New("missing required argument 'InstanceStorage'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:rds/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:rds/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The upgrade method to use. Valid values:
	// - Auto: Instances are automatically upgraded to a higher minor version.
	// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
	AutoUpgradeMinorVersion *string `pulumi:"autoUpgradeMinorVersion"`
	// RDS database connection string.
	ConnectionString *string `pulumi:"connectionString"`
	// The storage type of the instance. Valid values:
	// - local_ssd: specifies to use local SSDs. This value is recommended.
	// - cloud_ssd: specifies to use standard SSDs.
	// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
	DbInstanceStorageType *string `pulumi:"dbInstanceStorageType"`
	// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
	Engine *string `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Set it to true to make some parameter efficient when modifying them. Default to false.
	ForceRestart *bool `pulumi:"forceRestart"`
	// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The name of DB instance. It a string of 2 to 256 characters.
	InstanceName *string `pulumi:"instanceName"`
	// User-defined DB instance storage space. Value range:
	// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
	// - [20,1000] for MySQL 5.7 basic single node edition;
	// - [10, 2000] for SQL Server 2008R2;
	// - [20,2000] for SQL Server 2012 basic single node edition
	// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
	InstanceStorage *int `pulumi:"instanceStorage"`
	// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	InstanceType *string `pulumi:"instanceType"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
	MonitoringPeriod *int `pulumi:"monitoringPeriod"`
	// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
	Parameters []InstanceParameter `pulumi:"parameters"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// RDS database connection port.
	Port *string `pulumi:"port"`
	// It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
	SecurityIpMode *string `pulumi:"securityIpMode"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// The sql collector keep time of the instance. Valid values are `1`, `30`, `180`, `365`, `1095`, `1825`, `1` is the initial value, and can't change it to `1`.
	SqlCollectorConfigValue *int `pulumi:"sqlCollectorConfigValue"`
	// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
	SqlCollectorStatus *string `pulumi:"sqlCollectorStatus"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrInput
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// The upgrade method to use. Valid values:
	// - Auto: Instances are automatically upgraded to a higher minor version.
	// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
	AutoUpgradeMinorVersion pulumi.StringPtrInput
	// RDS database connection string.
	ConnectionString pulumi.StringPtrInput
	// The storage type of the instance. Valid values:
	// - local_ssd: specifies to use local SSDs. This value is recommended.
	// - cloud_ssd: specifies to use standard SSDs.
	// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
	DbInstanceStorageType pulumi.StringPtrInput
	// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
	Engine pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion pulumi.StringPtrInput
	// Set it to true to make some parameter efficient when modifying them. Default to false.
	ForceRestart pulumi.BoolPtrInput
	// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
	InstanceChargeType pulumi.StringPtrInput
	// The name of DB instance. It a string of 2 to 256 characters.
	InstanceName pulumi.StringPtrInput
	// User-defined DB instance storage space. Value range:
	// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
	// - [20,1000] for MySQL 5.7 basic single node edition;
	// - [10, 2000] for SQL Server 2008R2;
	// - [20,2000] for SQL Server 2012 basic single node edition
	// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
	InstanceStorage pulumi.IntPtrInput
	// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	InstanceType pulumi.StringPtrInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
	MonitoringPeriod pulumi.IntPtrInput
	// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
	Parameters InstanceParameterArrayInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// RDS database connection port.
	Port pulumi.StringPtrInput
	// It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
	SecurityGroupId pulumi.StringPtrInput
	// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
	SecurityGroupIds pulumi.StringArrayInput
	// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
	SecurityIpMode pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// The sql collector keep time of the instance. Valid values are `1`, `30`, `180`, `365`, `1095`, `1825`, `1` is the initial value, and can't change it to `1`.
	SqlCollectorConfigValue pulumi.IntPtrInput
	// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
	SqlCollectorStatus pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew *bool `pulumi:"autoRenew"`
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod *int `pulumi:"autoRenewPeriod"`
	// The upgrade method to use. Valid values:
	// - Auto: Instances are automatically upgraded to a higher minor version.
	// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
	AutoUpgradeMinorVersion *string `pulumi:"autoUpgradeMinorVersion"`
	// The storage type of the instance. Valid values:
	// - local_ssd: specifies to use local SSDs. This value is recommended.
	// - cloud_ssd: specifies to use standard SSDs.
	// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
	DbInstanceStorageType *string `pulumi:"dbInstanceStorageType"`
	// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
	Engine string `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion string `pulumi:"engineVersion"`
	// Set it to true to make some parameter efficient when modifying them. Default to false.
	ForceRestart *bool `pulumi:"forceRestart"`
	// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The name of DB instance. It a string of 2 to 256 characters.
	InstanceName *string `pulumi:"instanceName"`
	// User-defined DB instance storage space. Value range:
	// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
	// - [20,1000] for MySQL 5.7 basic single node edition;
	// - [10, 2000] for SQL Server 2008R2;
	// - [20,2000] for SQL Server 2012 basic single node edition
	// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
	InstanceStorage int `pulumi:"instanceStorage"`
	// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	InstanceType string `pulumi:"instanceType"`
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime *string `pulumi:"maintainTime"`
	// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
	MonitoringPeriod *int `pulumi:"monitoringPeriod"`
	// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
	Parameters []InstanceParameter `pulumi:"parameters"`
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period *int `pulumi:"period"`
	// It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
	SecurityIpMode *string `pulumi:"securityIpMode"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps []string `pulumi:"securityIps"`
	// The sql collector keep time of the instance. Valid values are `1`, `30`, `180`, `365`, `1095`, `1825`, `1` is the initial value, and can't change it to `1`.
	SqlCollectorConfigValue *int `pulumi:"sqlCollectorConfigValue"`
	// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
	SqlCollectorStatus *string `pulumi:"sqlCollectorStatus"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
	AutoRenew pulumi.BoolPtrInput
	// Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
	AutoRenewPeriod pulumi.IntPtrInput
	// The upgrade method to use. Valid values:
	// - Auto: Instances are automatically upgraded to a higher minor version.
	// - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
	AutoUpgradeMinorVersion pulumi.StringPtrInput
	// The storage type of the instance. Valid values:
	// - local_ssd: specifies to use local SSDs. This value is recommended.
	// - cloud_ssd: specifies to use standard SSDs.
	// - cloud_essd: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
	// - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
	DbInstanceStorageType pulumi.StringPtrInput
	// Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
	Engine pulumi.StringInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
	EngineVersion pulumi.StringInput
	// Set it to true to make some parameter efficient when modifying them. Default to false.
	ForceRestart pulumi.BoolPtrInput
	// Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
	InstanceChargeType pulumi.StringPtrInput
	// The name of DB instance. It a string of 2 to 256 characters.
	InstanceName pulumi.StringPtrInput
	// User-defined DB instance storage space. Value range:
	// - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
	// - [20,1000] for MySQL 5.7 basic single node edition;
	// - [10, 2000] for SQL Server 2008R2;
	// - [20,2000] for SQL Server 2012 basic single node edition
	// Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	// Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
	InstanceStorage pulumi.IntInput
	// DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
	InstanceType pulumi.StringInput
	// Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
	MaintainTime pulumi.StringPtrInput
	// The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.
	MonitoringPeriod pulumi.IntPtrInput
	// Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
	Parameters InstanceParameterArrayInput
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
	Period pulumi.IntPtrInput
	// It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
	SecurityGroupId pulumi.StringPtrInput
	// , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
	SecurityGroupIds pulumi.StringArrayInput
	// Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode
	SecurityIpMode pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIps pulumi.StringArrayInput
	// The sql collector keep time of the instance. Valid values are `1`, `30`, `180`, `365`, `1095`, `1825`, `1` is the initial value, and can't change it to `1`.
	SqlCollectorConfigValue pulumi.IntPtrInput
	// The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
	SqlCollectorStatus pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
