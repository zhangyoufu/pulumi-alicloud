// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS Enterprise Instance resource.
//
// > **NOTE:** API users must first register in DMS.
//
// > **NOTE:** Available since v1.81.0.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
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
//			current, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_default, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetUserTenants, err := dms.GetUserTenants(ctx, &dms.GetUserTenantsArgs{
//				Status: pulumi.StringRef("ACTIVE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetZones, err := rds.GetZones(ctx, &rds.GetZonesArgs{
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//				Category:              pulumi.StringRef("HighAvailability"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetInstanceClasses, err := rds.GetInstanceClasses(ctx, &rds.GetInstanceClassesArgs{
//				ZoneId:                pulumi.StringRef(defaultGetZones.Zones[0].Id),
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				Category:              pulumi.StringRef("HighAvailability"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(defaultGetZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewSecurityGroup(ctx, "default", &ecs.SecurityGroupArgs{
//				Name:  pulumi.String(name),
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
//				Engine:                pulumi.String("MySQL"),
//				EngineVersion:         pulumi.String("8.0"),
//				DbInstanceStorageType: pulumi.String("cloud_essd"),
//				InstanceType:          pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].InstanceClass),
//				InstanceStorage:       pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].StorageRange.Min),
//				VswitchId:             defaultSwitch.ID(),
//				InstanceName:          pulumi.String(name),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("100.104.5.0/24"),
//					pulumi.String("192.168.0.6"),
//				},
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccount, err := rds.NewAccount(ctx, "default", &rds.AccountArgs{
//				DbInstanceId:    defaultInstance.ID(),
//				AccountName:     pulumi.String("tfexamplename"),
//				AccountPassword: pulumi.String("Example12345"),
//				AccountType:     pulumi.String("Normal"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dms.NewEnterpriseInstance(ctx, "default", &dms.EnterpriseInstanceArgs{
//				Tid:              pulumi.String(defaultGetUserTenants.Ids[0]),
//				InstanceType:     pulumi.String("mysql"),
//				InstanceSource:   pulumi.String("RDS"),
//				NetworkType:      pulumi.String("VPC"),
//				EnvType:          pulumi.String("dev"),
//				Host:             defaultInstance.ConnectionString,
//				Port:             pulumi.Int(3306),
//				DatabaseUser:     defaultAccount.AccountName,
//				DatabasePassword: defaultAccount.AccountPassword,
//				InstanceName:     pulumi.String(name),
//				DbaUid:           pulumi.String(current.Id),
//				SafeRule:         pulumi.String("904496"),
//				UseDsql:          pulumi.Int(1),
//				QueryTimeout:     pulumi.Int(60),
//				ExportTimeout:    pulumi.Int(600),
//				EcsRegion:        pulumi.String(_default.Regions[0].Id),
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
// DMS Enterprise can be imported using host and port, e.g.
//
// ```sh
// $ pulumi import alicloud:dms/enterpriseInstance:EnterpriseInstance example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306
// ```
type EnterpriseInstance struct {
	pulumi.CustomResourceState

	// Cross-database query datalink name.
	DataLinkName pulumi.StringOutput `pulumi:"dataLinkName"`
	// Database access password.
	DatabasePassword pulumi.StringOutput `pulumi:"databasePassword"`
	// Database access account.
	DatabaseUser pulumi.StringOutput `pulumi:"databaseUser"`
	// The dba id of the database instance.
	DbaId pulumi.StringOutput `pulumi:"dbaId"`
	// The instance dba nickname.
	DbaNickName pulumi.StringOutput `pulumi:"dbaNickName"`
	// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
	DbaUid pulumi.IntOutput `pulumi:"dbaUid"`
	// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
	DdlOnline pulumi.IntPtrOutput `pulumi:"ddlOnline"`
	// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
	EcsInstanceId pulumi.StringOutput `pulumi:"ecsInstanceId"`
	// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
	EcsRegion pulumi.StringPtrOutput `pulumi:"ecsRegion"`
	// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
	EnvType pulumi.StringOutput `pulumi:"envType"`
	// Export timeout, unit: s (seconds).
	ExportTimeout pulumi.IntOutput `pulumi:"exportTimeout"`
	// Host address of the target database.
	Host pulumi.StringOutput `pulumi:"host"`
	// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
	//
	// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
	InstanceAlias pulumi.StringOutput `pulumi:"instanceAlias"`
	// The instance id of the database instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Instance name, to help users quickly distinguish positioning.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
	InstanceSource pulumi.StringOutput `pulumi:"instanceSource"`
	// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Network type. Valid values: `CLASSIC`, `VPC`.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// Access port of the target database.
	Port pulumi.IntOutput `pulumi:"port"`
	// Query timeout time, unit: s (seconds).
	QueryTimeout pulumi.IntOutput `pulumi:"queryTimeout"`
	// The security rule of the instance is passed into the name of the security rule in the enterprise.
	SafeRule pulumi.StringOutput `pulumi:"safeRule"`
	// The safe rule id of the database instance.
	SafeRuleId pulumi.StringOutput `pulumi:"safeRuleId"`
	// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
	Sid pulumi.StringPtrOutput `pulumi:"sid"`
	// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
	SkipTest pulumi.BoolPtrOutput `pulumi:"skipTest"`
	// It has been deprecated from provider version 1.100.0 and 'status' instead.
	//
	// Deprecated: Field 'state' has been deprecated from version 1.100.0. Use 'status' instead.
	State pulumi.StringOutput `pulumi:"state"`
	// The instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tenant ID.
	Tid pulumi.IntPtrOutput `pulumi:"tid"`
	// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
	UseDsql pulumi.IntOutput `pulumi:"useDsql"`
	// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewEnterpriseInstance registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseInstance(ctx *pulumi.Context,
	name string, args *EnterpriseInstanceArgs, opts ...pulumi.ResourceOption) (*EnterpriseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabasePassword == nil {
		return nil, errors.New("invalid value for required argument 'DatabasePassword'")
	}
	if args.DatabaseUser == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseUser'")
	}
	if args.DbaUid == nil {
		return nil, errors.New("invalid value for required argument 'DbaUid'")
	}
	if args.EnvType == nil {
		return nil, errors.New("invalid value for required argument 'EnvType'")
	}
	if args.ExportTimeout == nil {
		return nil, errors.New("invalid value for required argument 'ExportTimeout'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.InstanceSource == nil {
		return nil, errors.New("invalid value for required argument 'InstanceSource'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.NetworkType == nil {
		return nil, errors.New("invalid value for required argument 'NetworkType'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.QueryTimeout == nil {
		return nil, errors.New("invalid value for required argument 'QueryTimeout'")
	}
	if args.SafeRule == nil {
		return nil, errors.New("invalid value for required argument 'SafeRule'")
	}
	if args.DatabasePassword != nil {
		args.DatabasePassword = pulumi.ToSecret(args.DatabasePassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"databasePassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseInstance
	err := ctx.RegisterResource("alicloud:dms/enterpriseInstance:EnterpriseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseInstance gets an existing EnterpriseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseInstanceState, opts ...pulumi.ResourceOption) (*EnterpriseInstance, error) {
	var resource EnterpriseInstance
	err := ctx.ReadResource("alicloud:dms/enterpriseInstance:EnterpriseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseInstance resources.
type enterpriseInstanceState struct {
	// Cross-database query datalink name.
	DataLinkName *string `pulumi:"dataLinkName"`
	// Database access password.
	DatabasePassword *string `pulumi:"databasePassword"`
	// Database access account.
	DatabaseUser *string `pulumi:"databaseUser"`
	// The dba id of the database instance.
	DbaId *string `pulumi:"dbaId"`
	// The instance dba nickname.
	DbaNickName *string `pulumi:"dbaNickName"`
	// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
	DbaUid *int `pulumi:"dbaUid"`
	// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
	DdlOnline *int `pulumi:"ddlOnline"`
	// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
	EcsInstanceId *string `pulumi:"ecsInstanceId"`
	// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
	EcsRegion *string `pulumi:"ecsRegion"`
	// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
	EnvType *string `pulumi:"envType"`
	// Export timeout, unit: s (seconds).
	ExportTimeout *int `pulumi:"exportTimeout"`
	// Host address of the target database.
	Host *string `pulumi:"host"`
	// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
	//
	// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
	InstanceAlias *string `pulumi:"instanceAlias"`
	// The instance id of the database instance.
	InstanceId *string `pulumi:"instanceId"`
	// Instance name, to help users quickly distinguish positioning.
	InstanceName *string `pulumi:"instanceName"`
	// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
	InstanceSource *string `pulumi:"instanceSource"`
	// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
	InstanceType *string `pulumi:"instanceType"`
	// Network type. Valid values: `CLASSIC`, `VPC`.
	NetworkType *string `pulumi:"networkType"`
	// Access port of the target database.
	Port *int `pulumi:"port"`
	// Query timeout time, unit: s (seconds).
	QueryTimeout *int `pulumi:"queryTimeout"`
	// The security rule of the instance is passed into the name of the security rule in the enterprise.
	SafeRule *string `pulumi:"safeRule"`
	// The safe rule id of the database instance.
	SafeRuleId *string `pulumi:"safeRuleId"`
	// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
	Sid *string `pulumi:"sid"`
	// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
	SkipTest *bool `pulumi:"skipTest"`
	// It has been deprecated from provider version 1.100.0 and 'status' instead.
	//
	// Deprecated: Field 'state' has been deprecated from version 1.100.0. Use 'status' instead.
	State *string `pulumi:"state"`
	// The instance status.
	Status *string `pulumi:"status"`
	// The tenant ID.
	Tid *int `pulumi:"tid"`
	// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
	UseDsql *int `pulumi:"useDsql"`
	// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
	VpcId *string `pulumi:"vpcId"`
}

type EnterpriseInstanceState struct {
	// Cross-database query datalink name.
	DataLinkName pulumi.StringPtrInput
	// Database access password.
	DatabasePassword pulumi.StringPtrInput
	// Database access account.
	DatabaseUser pulumi.StringPtrInput
	// The dba id of the database instance.
	DbaId pulumi.StringPtrInput
	// The instance dba nickname.
	DbaNickName pulumi.StringPtrInput
	// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
	DbaUid pulumi.IntPtrInput
	// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
	DdlOnline pulumi.IntPtrInput
	// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
	EcsInstanceId pulumi.StringPtrInput
	// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
	EcsRegion pulumi.StringPtrInput
	// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
	EnvType pulumi.StringPtrInput
	// Export timeout, unit: s (seconds).
	ExportTimeout pulumi.IntPtrInput
	// Host address of the target database.
	Host pulumi.StringPtrInput
	// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
	//
	// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
	InstanceAlias pulumi.StringPtrInput
	// The instance id of the database instance.
	InstanceId pulumi.StringPtrInput
	// Instance name, to help users quickly distinguish positioning.
	InstanceName pulumi.StringPtrInput
	// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
	InstanceSource pulumi.StringPtrInput
	// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
	InstanceType pulumi.StringPtrInput
	// Network type. Valid values: `CLASSIC`, `VPC`.
	NetworkType pulumi.StringPtrInput
	// Access port of the target database.
	Port pulumi.IntPtrInput
	// Query timeout time, unit: s (seconds).
	QueryTimeout pulumi.IntPtrInput
	// The security rule of the instance is passed into the name of the security rule in the enterprise.
	SafeRule pulumi.StringPtrInput
	// The safe rule id of the database instance.
	SafeRuleId pulumi.StringPtrInput
	// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
	Sid pulumi.StringPtrInput
	// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
	SkipTest pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.100.0 and 'status' instead.
	//
	// Deprecated: Field 'state' has been deprecated from version 1.100.0. Use 'status' instead.
	State pulumi.StringPtrInput
	// The instance status.
	Status pulumi.StringPtrInput
	// The tenant ID.
	Tid pulumi.IntPtrInput
	// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
	UseDsql pulumi.IntPtrInput
	// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
	VpcId pulumi.StringPtrInput
}

func (EnterpriseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseInstanceState)(nil)).Elem()
}

type enterpriseInstanceArgs struct {
	// Cross-database query datalink name.
	DataLinkName *string `pulumi:"dataLinkName"`
	// Database access password.
	DatabasePassword string `pulumi:"databasePassword"`
	// Database access account.
	DatabaseUser string `pulumi:"databaseUser"`
	// The dba id of the database instance.
	DbaId *string `pulumi:"dbaId"`
	// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
	DbaUid int `pulumi:"dbaUid"`
	// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
	DdlOnline *int `pulumi:"ddlOnline"`
	// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
	EcsInstanceId *string `pulumi:"ecsInstanceId"`
	// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
	EcsRegion *string `pulumi:"ecsRegion"`
	// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
	EnvType string `pulumi:"envType"`
	// Export timeout, unit: s (seconds).
	ExportTimeout int `pulumi:"exportTimeout"`
	// Host address of the target database.
	Host string `pulumi:"host"`
	// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
	//
	// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
	InstanceAlias *string `pulumi:"instanceAlias"`
	// The instance id of the database instance.
	InstanceId *string `pulumi:"instanceId"`
	// Instance name, to help users quickly distinguish positioning.
	InstanceName *string `pulumi:"instanceName"`
	// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
	InstanceSource string `pulumi:"instanceSource"`
	// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
	InstanceType string `pulumi:"instanceType"`
	// Network type. Valid values: `CLASSIC`, `VPC`.
	NetworkType string `pulumi:"networkType"`
	// Access port of the target database.
	Port int `pulumi:"port"`
	// Query timeout time, unit: s (seconds).
	QueryTimeout int `pulumi:"queryTimeout"`
	// The security rule of the instance is passed into the name of the security rule in the enterprise.
	SafeRule string `pulumi:"safeRule"`
	// The safe rule id of the database instance.
	SafeRuleId *string `pulumi:"safeRuleId"`
	// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
	Sid *string `pulumi:"sid"`
	// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
	SkipTest *bool `pulumi:"skipTest"`
	// The tenant ID.
	Tid *int `pulumi:"tid"`
	// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
	UseDsql *int `pulumi:"useDsql"`
	// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a EnterpriseInstance resource.
type EnterpriseInstanceArgs struct {
	// Cross-database query datalink name.
	DataLinkName pulumi.StringPtrInput
	// Database access password.
	DatabasePassword pulumi.StringInput
	// Database access account.
	DatabaseUser pulumi.StringInput
	// The dba id of the database instance.
	DbaId pulumi.StringPtrInput
	// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
	DbaUid pulumi.IntInput
	// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
	DdlOnline pulumi.IntPtrInput
	// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
	EcsInstanceId pulumi.StringPtrInput
	// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
	EcsRegion pulumi.StringPtrInput
	// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
	EnvType pulumi.StringInput
	// Export timeout, unit: s (seconds).
	ExportTimeout pulumi.IntInput
	// Host address of the target database.
	Host pulumi.StringInput
	// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
	//
	// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
	InstanceAlias pulumi.StringPtrInput
	// The instance id of the database instance.
	InstanceId pulumi.StringPtrInput
	// Instance name, to help users quickly distinguish positioning.
	InstanceName pulumi.StringPtrInput
	// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
	InstanceSource pulumi.StringInput
	// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
	InstanceType pulumi.StringInput
	// Network type. Valid values: `CLASSIC`, `VPC`.
	NetworkType pulumi.StringInput
	// Access port of the target database.
	Port pulumi.IntInput
	// Query timeout time, unit: s (seconds).
	QueryTimeout pulumi.IntInput
	// The security rule of the instance is passed into the name of the security rule in the enterprise.
	SafeRule pulumi.StringInput
	// The safe rule id of the database instance.
	SafeRuleId pulumi.StringPtrInput
	// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
	Sid pulumi.StringPtrInput
	// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
	SkipTest pulumi.BoolPtrInput
	// The tenant ID.
	Tid pulumi.IntPtrInput
	// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
	UseDsql pulumi.IntPtrInput
	// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
	VpcId pulumi.StringPtrInput
}

func (EnterpriseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseInstanceArgs)(nil)).Elem()
}

type EnterpriseInstanceInput interface {
	pulumi.Input

	ToEnterpriseInstanceOutput() EnterpriseInstanceOutput
	ToEnterpriseInstanceOutputWithContext(ctx context.Context) EnterpriseInstanceOutput
}

func (*EnterpriseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseInstance)(nil)).Elem()
}

func (i *EnterpriseInstance) ToEnterpriseInstanceOutput() EnterpriseInstanceOutput {
	return i.ToEnterpriseInstanceOutputWithContext(context.Background())
}

func (i *EnterpriseInstance) ToEnterpriseInstanceOutputWithContext(ctx context.Context) EnterpriseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseInstanceOutput)
}

// EnterpriseInstanceArrayInput is an input type that accepts EnterpriseInstanceArray and EnterpriseInstanceArrayOutput values.
// You can construct a concrete instance of `EnterpriseInstanceArrayInput` via:
//
//	EnterpriseInstanceArray{ EnterpriseInstanceArgs{...} }
type EnterpriseInstanceArrayInput interface {
	pulumi.Input

	ToEnterpriseInstanceArrayOutput() EnterpriseInstanceArrayOutput
	ToEnterpriseInstanceArrayOutputWithContext(context.Context) EnterpriseInstanceArrayOutput
}

type EnterpriseInstanceArray []EnterpriseInstanceInput

func (EnterpriseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseInstance)(nil)).Elem()
}

func (i EnterpriseInstanceArray) ToEnterpriseInstanceArrayOutput() EnterpriseInstanceArrayOutput {
	return i.ToEnterpriseInstanceArrayOutputWithContext(context.Background())
}

func (i EnterpriseInstanceArray) ToEnterpriseInstanceArrayOutputWithContext(ctx context.Context) EnterpriseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseInstanceArrayOutput)
}

// EnterpriseInstanceMapInput is an input type that accepts EnterpriseInstanceMap and EnterpriseInstanceMapOutput values.
// You can construct a concrete instance of `EnterpriseInstanceMapInput` via:
//
//	EnterpriseInstanceMap{ "key": EnterpriseInstanceArgs{...} }
type EnterpriseInstanceMapInput interface {
	pulumi.Input

	ToEnterpriseInstanceMapOutput() EnterpriseInstanceMapOutput
	ToEnterpriseInstanceMapOutputWithContext(context.Context) EnterpriseInstanceMapOutput
}

type EnterpriseInstanceMap map[string]EnterpriseInstanceInput

func (EnterpriseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseInstance)(nil)).Elem()
}

func (i EnterpriseInstanceMap) ToEnterpriseInstanceMapOutput() EnterpriseInstanceMapOutput {
	return i.ToEnterpriseInstanceMapOutputWithContext(context.Background())
}

func (i EnterpriseInstanceMap) ToEnterpriseInstanceMapOutputWithContext(ctx context.Context) EnterpriseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseInstanceMapOutput)
}

type EnterpriseInstanceOutput struct{ *pulumi.OutputState }

func (EnterpriseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseInstance)(nil)).Elem()
}

func (o EnterpriseInstanceOutput) ToEnterpriseInstanceOutput() EnterpriseInstanceOutput {
	return o
}

func (o EnterpriseInstanceOutput) ToEnterpriseInstanceOutputWithContext(ctx context.Context) EnterpriseInstanceOutput {
	return o
}

// Cross-database query datalink name.
func (o EnterpriseInstanceOutput) DataLinkName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.DataLinkName }).(pulumi.StringOutput)
}

// Database access password.
func (o EnterpriseInstanceOutput) DatabasePassword() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.DatabasePassword }).(pulumi.StringOutput)
}

// Database access account.
func (o EnterpriseInstanceOutput) DatabaseUser() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.DatabaseUser }).(pulumi.StringOutput)
}

// The dba id of the database instance.
func (o EnterpriseInstanceOutput) DbaId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.DbaId }).(pulumi.StringOutput)
}

// The instance dba nickname.
func (o EnterpriseInstanceOutput) DbaNickName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.DbaNickName }).(pulumi.StringOutput)
}

// The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.
func (o EnterpriseInstanceOutput) DbaUid() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntOutput { return v.DbaUid }).(pulumi.IntOutput)
}

// Whether to use online services, currently only supports MySQL and PolarDB. Valid values: `0` Not used, `1` Native online DDL priority, `2` DMS lock-free table structure change priority.
func (o EnterpriseInstanceOutput) DdlOnline() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntPtrOutput { return v.DdlOnline }).(pulumi.IntPtrOutput)
}

// ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.
func (o EnterpriseInstanceOutput) EcsInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.EcsInstanceId }).(pulumi.StringOutput)
}

// The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.
func (o EnterpriseInstanceOutput) EcsRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringPtrOutput { return v.EcsRegion }).(pulumi.StringPtrOutput)
}

// Environment type. Valid values: `product` production environment, `dev` development environment, `pre` pre-release environment, `test` test environment, `sit` SIT environment, `uat` UAT environment, `pet` pressure test environment, `stag` STAG environment.
func (o EnterpriseInstanceOutput) EnvType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.EnvType }).(pulumi.StringOutput)
}

// Export timeout, unit: s (seconds).
func (o EnterpriseInstanceOutput) ExportTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntOutput { return v.ExportTimeout }).(pulumi.IntOutput)
}

// Host address of the target database.
func (o EnterpriseInstanceOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Field `instanceAlias` has been deprecated from version 1.100.0. Use `instanceName` instead.
//
// Deprecated: Field 'instance_alias' has been deprecated from version 1.100.0. Use 'instance_name' instead.
func (o EnterpriseInstanceOutput) InstanceAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.InstanceAlias }).(pulumi.StringOutput)
}

// The instance id of the database instance.
func (o EnterpriseInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance name, to help users quickly distinguish positioning.
func (o EnterpriseInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The source of the database instance. Valid values: `PUBLIC_OWN`, `RDS`, `ECS_OWN`, `VPC_IDC`.
func (o EnterpriseInstanceOutput) InstanceSource() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.InstanceSource }).(pulumi.StringOutput)
}

// Database type. Valid values: `MySQL`, `SQLServer`, `PostgreSQL`, `Oracle,` `DRDS`, `OceanBase`, `Mongo`, `Redis`.
func (o EnterpriseInstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// Network type. Valid values: `CLASSIC`, `VPC`.
func (o EnterpriseInstanceOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// Access port of the target database.
func (o EnterpriseInstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Query timeout time, unit: s (seconds).
func (o EnterpriseInstanceOutput) QueryTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntOutput { return v.QueryTimeout }).(pulumi.IntOutput)
}

// The security rule of the instance is passed into the name of the security rule in the enterprise.
func (o EnterpriseInstanceOutput) SafeRule() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.SafeRule }).(pulumi.StringOutput)
}

// The safe rule id of the database instance.
func (o EnterpriseInstanceOutput) SafeRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.SafeRuleId }).(pulumi.StringOutput)
}

// The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.
func (o EnterpriseInstanceOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringPtrOutput { return v.Sid }).(pulumi.StringPtrOutput)
}

// Whether the instance ignores test connectivity. Valid values: `true`, `false`.
func (o EnterpriseInstanceOutput) SkipTest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.BoolPtrOutput { return v.SkipTest }).(pulumi.BoolPtrOutput)
}

// It has been deprecated from provider version 1.100.0 and 'status' instead.
//
// Deprecated: Field 'state' has been deprecated from version 1.100.0. Use 'status' instead.
func (o EnterpriseInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The instance status.
func (o EnterpriseInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tenant ID.
func (o EnterpriseInstanceOutput) Tid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntPtrOutput { return v.Tid }).(pulumi.IntPtrOutput)
}

// Whether to enable cross-instance query. Valid values: `0` not open, `1` open.
func (o EnterpriseInstanceOutput) UseDsql() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.IntOutput { return v.UseDsql }).(pulumi.IntOutput)
}

// VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.
func (o EnterpriseInstanceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseInstance) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type EnterpriseInstanceArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseInstance)(nil)).Elem()
}

func (o EnterpriseInstanceArrayOutput) ToEnterpriseInstanceArrayOutput() EnterpriseInstanceArrayOutput {
	return o
}

func (o EnterpriseInstanceArrayOutput) ToEnterpriseInstanceArrayOutputWithContext(ctx context.Context) EnterpriseInstanceArrayOutput {
	return o
}

func (o EnterpriseInstanceArrayOutput) Index(i pulumi.IntInput) EnterpriseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseInstance {
		return vs[0].([]*EnterpriseInstance)[vs[1].(int)]
	}).(EnterpriseInstanceOutput)
}

type EnterpriseInstanceMapOutput struct{ *pulumi.OutputState }

func (EnterpriseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseInstance)(nil)).Elem()
}

func (o EnterpriseInstanceMapOutput) ToEnterpriseInstanceMapOutput() EnterpriseInstanceMapOutput {
	return o
}

func (o EnterpriseInstanceMapOutput) ToEnterpriseInstanceMapOutputWithContext(ctx context.Context) EnterpriseInstanceMapOutput {
	return o
}

func (o EnterpriseInstanceMapOutput) MapIndex(k pulumi.StringInput) EnterpriseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseInstance {
		return vs[0].(map[string]*EnterpriseInstance)[vs[1].(string)]
	}).(EnterpriseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseInstanceInput)(nil)).Elem(), &EnterpriseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseInstanceArrayInput)(nil)).Elem(), EnterpriseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseInstanceMapInput)(nil)).Elem(), EnterpriseInstanceMap{})
	pulumi.RegisterOutputType(EnterpriseInstanceOutput{})
	pulumi.RegisterOutputType(EnterpriseInstanceArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseInstanceMapOutput{})
}
