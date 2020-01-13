// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a MongoDB instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services. 
// It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
// You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/26558.htm)
// 
// > **NOTE:**  Available in 1.37.0+
// 
// > **NOTE:**  The following regions don't support create Classic network MongoDB instance.
// [`cn-zhangjiakou`,`cn-huhehaote`,`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`] 
// 
// > **NOTE:**  Create MongoDB instance or change instance type and storage would cost 5~10 minutes. Please make full preparation
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/mongodb_instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.DbInstanceClass == nil {
		return nil, errors.New("missing required argument 'DbInstanceClass'")
	}
	if args == nil || args.DbInstanceStorage == nil {
		return nil, errors.New("missing required argument 'DbInstanceStorage'")
	}
	if args == nil || args.EngineVersion == nil {
		return nil, errors.New("missing required argument 'EngineVersion'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountPassword"] = nil
		inputs["backupPeriods"] = nil
		inputs["backupTime"] = nil
		inputs["dbInstanceClass"] = nil
		inputs["dbInstanceStorage"] = nil
		inputs["engineVersion"] = nil
		inputs["instanceChargeType"] = nil
		inputs["kmsEncryptedPassword"] = nil
		inputs["kmsEncryptionContext"] = nil
		inputs["maintainEndTime"] = nil
		inputs["maintainStartTime"] = nil
		inputs["name"] = nil
		inputs["period"] = nil
		inputs["replicationFactor"] = nil
		inputs["securityIpLists"] = nil
		inputs["storageEngine"] = nil
		inputs["tags"] = nil
		inputs["vswitchId"] = nil
		inputs["zoneId"] = nil
	} else {
		inputs["accountPassword"] = args.AccountPassword
		inputs["backupPeriods"] = args.BackupPeriods
		inputs["backupTime"] = args.BackupTime
		inputs["dbInstanceClass"] = args.DbInstanceClass
		inputs["dbInstanceStorage"] = args.DbInstanceStorage
		inputs["engineVersion"] = args.EngineVersion
		inputs["instanceChargeType"] = args.InstanceChargeType
		inputs["kmsEncryptedPassword"] = args.KmsEncryptedPassword
		inputs["kmsEncryptionContext"] = args.KmsEncryptionContext
		inputs["maintainEndTime"] = args.MaintainEndTime
		inputs["maintainStartTime"] = args.MaintainStartTime
		inputs["name"] = args.Name
		inputs["period"] = args.Period
		inputs["replicationFactor"] = args.ReplicationFactor
		inputs["securityIpLists"] = args.SecurityIpLists
		inputs["storageEngine"] = args.StorageEngine
		inputs["tags"] = args.Tags
		inputs["vswitchId"] = args.VswitchId
		inputs["zoneId"] = args.ZoneId
	}
	inputs["replicaSetName"] = nil
	inputs["retentionPeriod"] = nil
	s, err := ctx.RegisterResource("alicloud:mongodb/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountPassword"] = state.AccountPassword
		inputs["backupPeriods"] = state.BackupPeriods
		inputs["backupTime"] = state.BackupTime
		inputs["dbInstanceClass"] = state.DbInstanceClass
		inputs["dbInstanceStorage"] = state.DbInstanceStorage
		inputs["engineVersion"] = state.EngineVersion
		inputs["instanceChargeType"] = state.InstanceChargeType
		inputs["kmsEncryptedPassword"] = state.KmsEncryptedPassword
		inputs["kmsEncryptionContext"] = state.KmsEncryptionContext
		inputs["maintainEndTime"] = state.MaintainEndTime
		inputs["maintainStartTime"] = state.MaintainStartTime
		inputs["name"] = state.Name
		inputs["period"] = state.Period
		inputs["replicaSetName"] = state.ReplicaSetName
		inputs["replicationFactor"] = state.ReplicationFactor
		inputs["retentionPeriod"] = state.RetentionPeriod
		inputs["securityIpLists"] = state.SecurityIpLists
		inputs["storageEngine"] = state.StorageEngine
		inputs["tags"] = state.Tags
		inputs["vswitchId"] = state.VswitchId
		inputs["zoneId"] = state.ZoneId
	}
	s, err := ctx.ReadResource("alicloud:mongodb/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
func (r *Instance) AccountPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accountPassword"])
}

// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
func (r *Instance) BackupPeriods() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["backupPeriods"])
}

// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to a random time, like "23:00Z-24:00Z".
func (r *Instance) BackupTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backupTime"])
}

// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
func (r *Instance) DbInstanceClass() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dbInstanceClass"])
}

// User-defined DB instance storage space.Unit: GB. Value range:
// - Custom storage space; value range: [10,2000]
// - 10-GB increments.
func (r *Instance) DbInstanceStorage() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["dbInstanceStorage"])
}

// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
func (r *Instance) EngineVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["engineVersion"])
}

// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
func (r *Instance) InstanceChargeType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceChargeType"])
}

// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
func (r *Instance) KmsEncryptedPassword() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kmsEncryptedPassword"])
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (r *Instance) KmsEncryptionContext() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["kmsEncryptionContext"])
}

// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
func (r *Instance) MaintainEndTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["maintainEndTime"])
}

// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
func (r *Instance) MaintainStartTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["maintainStartTime"])
}

// The name of DB instance. It a string of 2 to 256 characters.
func (r *Instance) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
func (r *Instance) Period() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["period"])
}

// The name of the mongo replica set
func (r *Instance) ReplicaSetName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["replicaSetName"])
}

// Number of replica set nodes. Valid values: [3, 5, 7]
// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
func (r *Instance) ReplicationFactor() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["replicationFactor"])
}

// Instance log backup retention days. Available in 1.42.0+.
func (r *Instance) RetentionPeriod() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["retentionPeriod"])
}

// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
func (r *Instance) SecurityIpLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["securityIpLists"])
}

func (r *Instance) StorageEngine() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageEngine"])
}

// A mapping of tags to assign to the resource.
func (r *Instance) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The virtual switch ID to launch DB instances in one VPC.
func (r *Instance) VswitchId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vswitchId"])
}

// The Zone to launch the DB instance. it supports multiple zone.
// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
func (r *Instance) ZoneId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zoneId"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword interface{}
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods interface{}
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to a random time, like "23:00Z-24:00Z".
	BackupTime interface{}
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	DbInstanceClass interface{}
	// User-defined DB instance storage space.Unit: GB. Value range:
	// - Custom storage space; value range: [10,2000]
	// - 10-GB increments.
	DbInstanceStorage interface{}
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
	EngineVersion interface{}
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
	InstanceChargeType interface{}
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword interface{}
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext interface{}
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime interface{}
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime interface{}
	// The name of DB instance. It a string of 2 to 256 characters.
	Name interface{}
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period interface{}
	// The name of the mongo replica set
	ReplicaSetName interface{}
	// Number of replica set nodes. Valid values: [3, 5, 7]
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	ReplicationFactor interface{}
	// Instance log backup retention days. Available in 1.42.0+.
	RetentionPeriod interface{}
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists interface{}
	StorageEngine interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId interface{}
	// The Zone to launch the DB instance. it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
	AccountPassword interface{}
	// MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
	BackupPeriods interface{}
	// MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to a random time, like "23:00Z-24:00Z".
	BackupTime interface{}
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
	DbInstanceClass interface{}
	// User-defined DB instance storage space.Unit: GB. Value range:
	// - Custom storage space; value range: [10,2000]
	// - 10-GB increments.
	DbInstanceStorage interface{}
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) `EngineVersion`.
	EngineVersion interface{}
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. It can be modified from `PostPaid` to `PrePaid` after version 1.63.0.
	InstanceChargeType interface{}
	// An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword interface{}
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext interface{}
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainEndTime interface{}
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).
	MaintainStartTime interface{}
	// The name of DB instance. It a string of 2 to 256 characters.
	Name interface{}
	// The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
	Period interface{}
	// Number of replica set nodes. Valid values: [3, 5, 7]
	// * `storageEngine` (Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
	ReplicationFactor interface{}
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists interface{}
	StorageEngine interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId interface{}
	// The Zone to launch the DB instance. it supports multiple zone.
	// If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
	// The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `.getZones`.
	ZoneId interface{}
}
