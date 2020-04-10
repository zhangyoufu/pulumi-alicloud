// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a EMR Cluster resource. With this you can create, read, and release  EMR Cluster.
//
// > **NOTE:** Available in 1.57.0+.
type Cluster struct {
	pulumi.CustomResourceState

	BootstrapActions ClusterBootstrapActionArrayOutput `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringOutput    `pulumi:"clusterType"`
	DepositType pulumi.StringPtrOutput `pulumi:"depositType"`
	EasEnable   pulumi.BoolPtrOutput   `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer                 pulumi.StringOutput  `pulumi:"emrVer"`
	HighAvailabilityEnable pulumi.BoolPtrOutput `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups     ClusterHostGroupArrayOutput `pulumi:"hostGroups"`
	IsOpenPublicIp pulumi.BoolPtrOutput        `pulumi:"isOpenPublicIp"`
	KeyPairName    pulumi.StringPtrOutput      `pulumi:"keyPairName"`
	MasterPwd      pulumi.StringPtrOutput      `pulumi:"masterPwd"`
	// bootstrap action name.
	Name                pulumi.StringOutput      `pulumi:"name"`
	OptionSoftwareLists pulumi.StringArrayOutput `pulumi:"optionSoftwareLists"`
	RelatedClusterId    pulumi.StringPtrOutput   `pulumi:"relatedClusterId"`
	SecurityGroupId     pulumi.StringPtrOutput   `pulumi:"securityGroupId"`
	SshEnable           pulumi.BoolPtrOutput     `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags                  pulumi.MapOutput       `pulumi:"tags"`
	UseLocalMetadb        pulumi.BoolPtrOutput   `pulumi:"useLocalMetadb"`
	UserDefinedEmrEcsRole pulumi.StringPtrOutput `pulumi:"userDefinedEmrEcsRole"`
	VswitchId             pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	// * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
	// * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
	// * `optionSoftwareList` (Optional, ForceNew) Optional software list.
	// * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	// * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
	// * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
	// * `masterPwd` (Optional, ForceNew) Master ssh password.
	// * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
	// * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	// * `keyPairName` (Optional, ForceNew) Ssh key pair.
	// * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	// * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ClusterType == nil {
		return nil, errors.New("missing required argument 'ClusterType'")
	}
	if args == nil || args.EmrVer == nil {
		return nil, errors.New("missing required argument 'EmrVer'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:emr/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:emr/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType *string `pulumi:"clusterType"`
	DepositType *string `pulumi:"depositType"`
	EasEnable   *bool   `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer                 *string `pulumi:"emrVer"`
	HighAvailabilityEnable *bool   `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups     []ClusterHostGroup `pulumi:"hostGroups"`
	IsOpenPublicIp *bool              `pulumi:"isOpenPublicIp"`
	KeyPairName    *string            `pulumi:"keyPairName"`
	MasterPwd      *string            `pulumi:"masterPwd"`
	// bootstrap action name.
	Name                *string  `pulumi:"name"`
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	RelatedClusterId    *string  `pulumi:"relatedClusterId"`
	SecurityGroupId     *string  `pulumi:"securityGroupId"`
	SshEnable           *bool    `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags                  map[string]interface{} `pulumi:"tags"`
	UseLocalMetadb        *bool                  `pulumi:"useLocalMetadb"`
	UserDefinedEmrEcsRole *string                `pulumi:"userDefinedEmrEcsRole"`
	VswitchId             *string                `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	// * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
	// * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
	// * `optionSoftwareList` (Optional, ForceNew) Optional software list.
	// * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	// * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
	// * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
	// * `masterPwd` (Optional, ForceNew) Master ssh password.
	// * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
	// * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	// * `keyPairName` (Optional, ForceNew) Ssh key pair.
	// * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	// * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringPtrInput
	DepositType pulumi.StringPtrInput
	EasEnable   pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer                 pulumi.StringPtrInput
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups     ClusterHostGroupArrayInput
	IsOpenPublicIp pulumi.BoolPtrInput
	KeyPairName    pulumi.StringPtrInput
	MasterPwd      pulumi.StringPtrInput
	// bootstrap action name.
	Name                pulumi.StringPtrInput
	OptionSoftwareLists pulumi.StringArrayInput
	RelatedClusterId    pulumi.StringPtrInput
	SecurityGroupId     pulumi.StringPtrInput
	SshEnable           pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags                  pulumi.MapInput
	UseLocalMetadb        pulumi.BoolPtrInput
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	VswitchId             pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	// * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
	// * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
	// * `optionSoftwareList` (Optional, ForceNew) Optional software list.
	// * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	// * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
	// * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
	// * `masterPwd` (Optional, ForceNew) Master ssh password.
	// * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
	// * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	// * `keyPairName` (Optional, ForceNew) Ssh key pair.
	// * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	// * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType string  `pulumi:"clusterType"`
	DepositType *string `pulumi:"depositType"`
	EasEnable   *bool   `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer                 string `pulumi:"emrVer"`
	HighAvailabilityEnable *bool  `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups     []ClusterHostGroup `pulumi:"hostGroups"`
	IsOpenPublicIp *bool              `pulumi:"isOpenPublicIp"`
	KeyPairName    *string            `pulumi:"keyPairName"`
	MasterPwd      *string            `pulumi:"masterPwd"`
	// bootstrap action name.
	Name                *string  `pulumi:"name"`
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	RelatedClusterId    *string  `pulumi:"relatedClusterId"`
	SecurityGroupId     *string  `pulumi:"securityGroupId"`
	SshEnable           *bool    `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags                  map[string]interface{} `pulumi:"tags"`
	UseLocalMetadb        *bool                  `pulumi:"useLocalMetadb"`
	UserDefinedEmrEcsRole *string                `pulumi:"userDefinedEmrEcsRole"`
	VswitchId             *string                `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	// * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
	// * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
	// * `optionSoftwareList` (Optional, ForceNew) Optional software list.
	// * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	// * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
	// * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
	// * `masterPwd` (Optional, ForceNew) Master ssh password.
	// * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
	// * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	// * `keyPairName` (Optional, ForceNew) Ssh key pair.
	// * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	// * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringInput
	DepositType pulumi.StringPtrInput
	EasEnable   pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer                 pulumi.StringInput
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups     ClusterHostGroupArrayInput
	IsOpenPublicIp pulumi.BoolPtrInput
	KeyPairName    pulumi.StringPtrInput
	MasterPwd      pulumi.StringPtrInput
	// bootstrap action name.
	Name                pulumi.StringPtrInput
	OptionSoftwareLists pulumi.StringArrayInput
	RelatedClusterId    pulumi.StringPtrInput
	SecurityGroupId     pulumi.StringPtrInput
	SshEnable           pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags                  pulumi.MapInput
	UseLocalMetadb        pulumi.BoolPtrInput
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	VswitchId             pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	// * `securityGroupId` (Optional, ForceNew) Security Group ID for Cluster, you can also specify this key for each host group.
	// * `vswitchId` (Optional, ForceNew) Global vswitch id, you can also specify it in host group.
	// * `optionSoftwareList` (Optional, ForceNew) Optional software list.
	// * `highAvailabilityEnable` (Optional, ForceNew) High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	// * `useLocalMetadb` (Optional, ForceNew) Use local metadb. Default is false.
	// * `sshEnable` (Optional, ForceNew) If this is set true, we can ssh into cluster. Default value is false.
	// * `masterPwd` (Optional, ForceNew) Master ssh password.
	// * `easEnable` (Optional, ForceNew) High security cluster (true) or not. Default value is false.
	// * `userDefinedEmrEcsRole` (Optional, ForceNew) Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	// * `keyPairName` (Optional, ForceNew) Ssh key pair.
	// * `depositType` (Optional, ForceNew) Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	// * `relatedClusterId` (Optional, ForceNew) This specify the related cluster id, if this cluster is a Gateway.
	ZoneId pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
