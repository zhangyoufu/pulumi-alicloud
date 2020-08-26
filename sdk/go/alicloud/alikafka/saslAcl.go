// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alikafka

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ALIKAFKA sasl acl resource.
//
// > **NOTE:** Available in 1.66.0+
//
// > **NOTE:**  Only the following regions support create alikafka sasl user.
// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`ap-southeast-1`,`ap-south-1`,`ap-southeast-5`]
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/alikafka"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "VSwitch"
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 			CidrBlock:        pulumi.String("172.16.0.0/24"),
// 			VpcId:            defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultInstance, err := alikafka.NewInstance(ctx, "defaultInstance", &alikafka.InstanceArgs{
// 			DeployType: pulumi.Int(5),
// 			DiskSize:   pulumi.Int(500),
// 			DiskType:   pulumi.Int(1),
// 			IoMax:      pulumi.Int(20),
// 			TopicQuota: pulumi.Int(50),
// 			VswitchId:  defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultTopic, err := alikafka.NewTopic(ctx, "defaultTopic", &alikafka.TopicArgs{
// 			InstanceId: defaultInstance.ID(),
// 			Remark:     pulumi.String("topic-remark"),
// 			Topic:      pulumi.String("test-topic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSaslUser, err := alikafka.NewSaslUser(ctx, "defaultSaslUser", &alikafka.SaslUserArgs{
// 			InstanceId: defaultInstance.ID(),
// 			Password:   pulumi.String(password),
// 			Username:   pulumi.String(username),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = alikafka.NewSaslAcl(ctx, "defaultSaslAcl", &alikafka.SaslAclArgs{
// 			AclOperationType:       pulumi.String("Write"),
// 			AclResourceName:        defaultTopic.Topic,
// 			AclResourcePatternType: pulumi.String("LITERAL"),
// 			AclResourceType:        pulumi.String("Topic"),
// 			InstanceId:             defaultInstance.ID(),
// 			Username:               defaultSaslUser.Username,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SaslAcl struct {
	pulumi.CustomResourceState

	// Operation type for this acl. The operation type can only be "Write" and "Read".
	AclOperationType pulumi.StringOutput `pulumi:"aclOperationType"`
	// Resource name for this acl. The resource name should be a topic or consumer group name.
	AclResourceName pulumi.StringOutput `pulumi:"aclResourceName"`
	// Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
	AclResourcePatternType pulumi.StringOutput `pulumi:"aclResourcePatternType"`
	// Resource type for this acl. The resource type can only be "Topic" and "Group".
	AclResourceType pulumi.StringOutput `pulumi:"aclResourceType"`
	// The host of the acl.
	Host pulumi.StringOutput `pulumi:"host"`
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSaslAcl registers a new resource with the given unique name, arguments, and options.
func NewSaslAcl(ctx *pulumi.Context,
	name string, args *SaslAclArgs, opts ...pulumi.ResourceOption) (*SaslAcl, error) {
	if args == nil || args.AclOperationType == nil {
		return nil, errors.New("missing required argument 'AclOperationType'")
	}
	if args == nil || args.AclResourceName == nil {
		return nil, errors.New("missing required argument 'AclResourceName'")
	}
	if args == nil || args.AclResourcePatternType == nil {
		return nil, errors.New("missing required argument 'AclResourcePatternType'")
	}
	if args == nil || args.AclResourceType == nil {
		return nil, errors.New("missing required argument 'AclResourceType'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &SaslAclArgs{}
	}
	var resource SaslAcl
	err := ctx.RegisterResource("alicloud:alikafka/saslAcl:SaslAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSaslAcl gets an existing SaslAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSaslAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SaslAclState, opts ...pulumi.ResourceOption) (*SaslAcl, error) {
	var resource SaslAcl
	err := ctx.ReadResource("alicloud:alikafka/saslAcl:SaslAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SaslAcl resources.
type saslAclState struct {
	// Operation type for this acl. The operation type can only be "Write" and "Read".
	AclOperationType *string `pulumi:"aclOperationType"`
	// Resource name for this acl. The resource name should be a topic or consumer group name.
	AclResourceName *string `pulumi:"aclResourceName"`
	// Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
	AclResourcePatternType *string `pulumi:"aclResourcePatternType"`
	// Resource type for this acl. The resource type can only be "Topic" and "Group".
	AclResourceType *string `pulumi:"aclResourceType"`
	// The host of the acl.
	Host *string `pulumi:"host"`
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId *string `pulumi:"instanceId"`
	// Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
	Username *string `pulumi:"username"`
}

type SaslAclState struct {
	// Operation type for this acl. The operation type can only be "Write" and "Read".
	AclOperationType pulumi.StringPtrInput
	// Resource name for this acl. The resource name should be a topic or consumer group name.
	AclResourceName pulumi.StringPtrInput
	// Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
	AclResourcePatternType pulumi.StringPtrInput
	// Resource type for this acl. The resource type can only be "Topic" and "Group".
	AclResourceType pulumi.StringPtrInput
	// The host of the acl.
	Host pulumi.StringPtrInput
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringPtrInput
	// Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
	Username pulumi.StringPtrInput
}

func (SaslAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*saslAclState)(nil)).Elem()
}

type saslAclArgs struct {
	// Operation type for this acl. The operation type can only be "Write" and "Read".
	AclOperationType string `pulumi:"aclOperationType"`
	// Resource name for this acl. The resource name should be a topic or consumer group name.
	AclResourceName string `pulumi:"aclResourceName"`
	// Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
	AclResourcePatternType string `pulumi:"aclResourcePatternType"`
	// Resource type for this acl. The resource type can only be "Topic" and "Group".
	AclResourceType string `pulumi:"aclResourceType"`
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId string `pulumi:"instanceId"`
	// Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SaslAcl resource.
type SaslAclArgs struct {
	// Operation type for this acl. The operation type can only be "Write" and "Read".
	AclOperationType pulumi.StringInput
	// Resource name for this acl. The resource name should be a topic or consumer group name.
	AclResourceName pulumi.StringInput
	// Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "*" can be used to represent a resource with any name. "PREFIXED": A prefixed name defines a prefix for a resource.
	AclResourcePatternType pulumi.StringInput
	// Resource type for this acl. The resource type can only be "Topic" and "Group".
	AclResourceType pulumi.StringInput
	// ID of the ALIKAFKA Instance that owns the groups.
	InstanceId pulumi.StringInput
	// Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.
	Username pulumi.StringInput
}

func (SaslAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*saslAclArgs)(nil)).Elem()
}
