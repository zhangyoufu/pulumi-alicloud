// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alikafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// AliKafka instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:alikafka/instance:Instance instance <id>
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringOutput `pulumi:"config"`
	// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
	// - 4: eip/vpc instance
	// - 5: vpc instance.
	DeployType pulumi.IntOutput `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntOutput `pulumi:"diskType"`
	// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
	EipMax pulumi.IntOutput `pulumi:"eipMax"`
	// The EndPoint to access the kafka instance.
	EndPoint pulumi.StringOutput `pulumi:"endPoint"`
	// (Available since v1.214.1) The number of available groups.
	GroupLeft pulumi.IntOutput `pulumi:"groupLeft"`
	// (Available since v1.214.1) The number of used groups.
	GroupUsed pulumi.IntOutput `pulumi:"groupUsed"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntOutput `pulumi:"ioMax"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
	// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
	IoMaxSpec pulumi.StringOutput `pulumi:"ioMaxSpec"`
	// (Available since v1.214.1) The method that you use to purchase partitions.
	IsPartitionBuy pulumi.IntOutput `pulumi:"isPartitionBuy"`
	// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrOutput `pulumi:"paidType"`
	// (Available since v1.214.1) The number of available partitions.
	PartitionLeft pulumi.IntOutput `pulumi:"partitionLeft"`
	// The number of partitions.
	PartitionNum pulumi.IntPtrOutput `pulumi:"partitionNum"`
	// (Available since v1.214.1) The number of used partitions.
	PartitionUsed pulumi.IntOutput `pulumi:"partitionUsed"`
	// The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringOutput `pulumi:"securityGroup"`
	// The zones among which you want to deploy the instance.
	//
	// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
	//
	// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
	// |------|-------------|:----:|:-----:|
	// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
	// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
	// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
	// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
	// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
	SelectedZones pulumi.StringArrayOutput `pulumi:"selectedZones"`
	// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringOutput `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrOutput `pulumi:"specType"`
	// The status of the instance.
	Status pulumi.IntOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// (Available since v1.214.1) The number of available topics.
	TopicLeft pulumi.IntOutput `pulumi:"topicLeft"`
	// (Available since v1.214.1) The number of purchased topics.
	TopicNumOfBuy pulumi.IntOutput `pulumi:"topicNumOfBuy"`
	// The max num of topic can be creation of the instance.
	// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
	// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
	// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
	//
	// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
	TopicQuota pulumi.IntOutput `pulumi:"topicQuota"`
	// (Available since v1.214.1) The number of used topics.
	TopicUsed pulumi.IntOutput `pulumi:"topicUsed"`
	// The VPC ID of the instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployType == nil {
		return nil, errors.New("invalid value for required argument 'DeployType'")
	}
	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.DiskType == nil {
		return nil, errors.New("invalid value for required argument 'DiskType'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:alikafka/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:alikafka/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config *string `pulumi:"config"`
	// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
	// - 4: eip/vpc instance
	// - 5: vpc instance.
	DeployType *int `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
	DiskSize *int `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType *int `pulumi:"diskType"`
	// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
	EipMax *int `pulumi:"eipMax"`
	// The EndPoint to access the kafka instance.
	EndPoint *string `pulumi:"endPoint"`
	// (Available since v1.214.1) The number of available groups.
	GroupLeft *int `pulumi:"groupLeft"`
	// (Available since v1.214.1) The number of used groups.
	GroupUsed *int `pulumi:"groupUsed"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax *int `pulumi:"ioMax"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
	// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
	IoMaxSpec *string `pulumi:"ioMaxSpec"`
	// (Available since v1.214.1) The method that you use to purchase partitions.
	IsPartitionBuy *int `pulumi:"isPartitionBuy"`
	// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name *string `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType *string `pulumi:"paidType"`
	// (Available since v1.214.1) The number of available partitions.
	PartitionLeft *int `pulumi:"partitionLeft"`
	// The number of partitions.
	PartitionNum *int `pulumi:"partitionNum"`
	// (Available since v1.214.1) The number of used partitions.
	PartitionUsed *int `pulumi:"partitionUsed"`
	// The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup *string `pulumi:"securityGroup"`
	// The zones among which you want to deploy the instance.
	//
	// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
	//
	// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
	// |------|-------------|:----:|:-----:|
	// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
	// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
	// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
	// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
	// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
	SelectedZones []string `pulumi:"selectedZones"`
	// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion *string `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType *string `pulumi:"specType"`
	// The status of the instance.
	Status *int `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// (Available since v1.214.1) The number of available topics.
	TopicLeft *int `pulumi:"topicLeft"`
	// (Available since v1.214.1) The number of purchased topics.
	TopicNumOfBuy *int `pulumi:"topicNumOfBuy"`
	// The max num of topic can be creation of the instance.
	// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
	// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
	// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
	//
	// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
	TopicQuota *int `pulumi:"topicQuota"`
	// (Available since v1.214.1) The number of used topics.
	TopicUsed *int `pulumi:"topicUsed"`
	// The VPC ID of the instance.
	VpcId *string `pulumi:"vpcId"`
	// The ID of attaching vswitch to instance.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringPtrInput
	// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
	// - 4: eip/vpc instance
	// - 5: vpc instance.
	DeployType pulumi.IntPtrInput
	// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
	DiskSize pulumi.IntPtrInput
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntPtrInput
	// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
	EipMax pulumi.IntPtrInput
	// The EndPoint to access the kafka instance.
	EndPoint pulumi.StringPtrInput
	// (Available since v1.214.1) The number of available groups.
	GroupLeft pulumi.IntPtrInput
	// (Available since v1.214.1) The number of used groups.
	GroupUsed pulumi.IntPtrInput
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntPtrInput
	// The traffic specification of the instance. We recommend that you configure this parameter.
	// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
	// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
	IoMaxSpec pulumi.StringPtrInput
	// (Available since v1.214.1) The method that you use to purchase partitions.
	IsPartitionBuy pulumi.IntPtrInput
	// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
	KmsKeyId pulumi.StringPtrInput
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringPtrInput
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrInput
	// (Available since v1.214.1) The number of available partitions.
	PartitionLeft pulumi.IntPtrInput
	// The number of partitions.
	PartitionNum pulumi.IntPtrInput
	// (Available since v1.214.1) The number of used partitions.
	PartitionUsed pulumi.IntPtrInput
	// The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringPtrInput
	// The zones among which you want to deploy the instance.
	//
	// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
	//
	// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
	// |------|-------------|:----:|:-----:|
	// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
	// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
	// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
	// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
	// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
	SelectedZones pulumi.StringArrayInput
	// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringPtrInput
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrInput
	// The status of the instance.
	Status pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// (Available since v1.214.1) The number of available topics.
	TopicLeft pulumi.IntPtrInput
	// (Available since v1.214.1) The number of purchased topics.
	TopicNumOfBuy pulumi.IntPtrInput
	// The max num of topic can be creation of the instance.
	// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
	// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
	// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
	//
	// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
	TopicQuota pulumi.IntPtrInput
	// (Available since v1.214.1) The number of used topics.
	TopicUsed pulumi.IntPtrInput
	// The VPC ID of the instance.
	VpcId pulumi.StringPtrInput
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringPtrInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config *string `pulumi:"config"`
	// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
	// - 4: eip/vpc instance
	// - 5: vpc instance.
	DeployType int `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
	DiskSize int `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType int `pulumi:"diskType"`
	// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
	EipMax *int `pulumi:"eipMax"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax *int `pulumi:"ioMax"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
	// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
	IoMaxSpec *string `pulumi:"ioMaxSpec"`
	// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name *string `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType *string `pulumi:"paidType"`
	// The number of partitions.
	PartitionNum *int `pulumi:"partitionNum"`
	// The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup *string `pulumi:"securityGroup"`
	// The zones among which you want to deploy the instance.
	//
	// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
	//
	// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
	// |------|-------------|:----:|:-----:|
	// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
	// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
	// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
	// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
	// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
	SelectedZones []string `pulumi:"selectedZones"`
	// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion *string `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType *string `pulumi:"specType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The max num of topic can be creation of the instance.
	// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
	// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
	// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
	//
	// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
	TopicQuota *int `pulumi:"topicQuota"`
	// The VPC ID of the instance.
	VpcId *string `pulumi:"vpcId"`
	// The ID of attaching vswitch to instance.
	VswitchId string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringPtrInput
	// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
	// - 4: eip/vpc instance
	// - 5: vpc instance.
	DeployType pulumi.IntInput
	// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
	DiskSize pulumi.IntInput
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntInput
	// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
	EipMax pulumi.IntPtrInput
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntPtrInput
	// The traffic specification of the instance. We recommend that you configure this parameter.
	// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
	// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
	IoMaxSpec pulumi.StringPtrInput
	// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
	KmsKeyId pulumi.StringPtrInput
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringPtrInput
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrInput
	// The number of partitions.
	PartitionNum pulumi.IntPtrInput
	// The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringPtrInput
	// The zones among which you want to deploy the instance.
	//
	// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
	//
	// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
	// |------|-------------|:----:|:-----:|
	// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
	// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
	// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
	// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
	// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
	SelectedZones pulumi.StringArrayInput
	// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringPtrInput
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The max num of topic can be creation of the instance.
	// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
	// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
	// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
	//
	// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
	TopicQuota pulumi.IntPtrInput
	// The VPC ID of the instance.
	VpcId pulumi.StringPtrInput
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
func (o InstanceOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

// The deployment type of the instance. **NOTE:** From version 1.161.0, this attribute supports to be updated. Valid values:
// - 4: eip/vpc instance
// - 5: vpc instance.
func (o InstanceOutput) DeployType() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.DeployType }).(pulumi.IntOutput)
}

// The disk size of the instance. When modify this value, it only supports adjust to a greater value.
func (o InstanceOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
func (o InstanceOutput) DiskType() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.DiskType }).(pulumi.IntOutput)
}

// The max bandwidth of the instance. It will be ignored when `deployType = 5`. When modify this value, it only supports adjust to a greater value.
func (o InstanceOutput) EipMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.EipMax }).(pulumi.IntOutput)
}

// The EndPoint to access the kafka instance.
func (o InstanceOutput) EndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EndPoint }).(pulumi.StringOutput)
}

// (Available since v1.214.1) The number of available groups.
func (o InstanceOutput) GroupLeft() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.GroupLeft }).(pulumi.IntOutput)
}

// (Available since v1.214.1) The number of used groups.
func (o InstanceOutput) GroupUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.GroupUsed }).(pulumi.IntOutput)
}

// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
func (o InstanceOutput) IoMax() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.IoMax }).(pulumi.IntOutput)
}

// The traffic specification of the instance. We recommend that you configure this parameter.
// - You should specify one of the `ioMax` and `ioMaxSpec` parameters, and `ioMaxSpec` is recommended.
// - For more information about the valid values, see [Billing](https://www.alibabacloud.com/help/en/message-queue-for-apache-kafka/latest/billing-overview).
func (o InstanceOutput) IoMaxSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.IoMaxSpec }).(pulumi.StringOutput)
}

// (Available since v1.214.1) The method that you use to purchase partitions.
func (o InstanceOutput) IsPartitionBuy() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.IsPartitionBuy }).(pulumi.IntOutput)
}

// The ID of the key that is used to encrypt data on standard SSDs in the region of the instance.
func (o InstanceOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
func (o InstanceOutput) PaidType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PaidType }).(pulumi.StringPtrOutput)
}

// (Available since v1.214.1) The number of available partitions.
func (o InstanceOutput) PartitionLeft() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.PartitionLeft }).(pulumi.IntOutput)
}

// The number of partitions.
func (o InstanceOutput) PartitionNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.PartitionNum }).(pulumi.IntPtrOutput)
}

// (Available since v1.214.1) The number of used partitions.
func (o InstanceOutput) PartitionUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.PartitionUsed }).(pulumi.IntOutput)
}

// The ID of security group for this instance. If the security group is empty, system will create a default one.
func (o InstanceOutput) SecurityGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SecurityGroup }).(pulumi.StringOutput)
}

// The zones among which you want to deploy the instance.
//
// > **NOTE:** Arguments io_max, disk_size, topic_quota, eipMax should follow the following constraints.
//
// | ioMax | disk_size(min-max:lag) | topic_quota(min-max:lag) | eip_max(min-max:lag) |
// |------|-------------|:----:|:-----:|
// |20          |  500-6100:100   |   50-450:1  |    1-160:1  |
// |30          |  800-6100:100   |   50-450:1  |    1-240:1  |
// |60          |  1400-6100:100  |   80-450:1  |    1-500:1  |
// |90          |  2100-6100:100  |   100-450:1 |    1-500:1  |
// |120         |  2700-6100:100  |   150-450:1 |    1-500:1  |
func (o InstanceOutput) SelectedZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SelectedZones }).(pulumi.StringArrayOutput)
}

// The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
func (o InstanceOutput) ServiceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceVersion }).(pulumi.StringOutput)
}

// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
func (o InstanceOutput) SpecType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SpecType }).(pulumi.StringPtrOutput)
}

// The status of the instance.
func (o InstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// A mapping of tags to assign to the resource.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// (Available since v1.214.1) The number of available topics.
func (o InstanceOutput) TopicLeft() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TopicLeft }).(pulumi.IntOutput)
}

// (Available since v1.214.1) The number of purchased topics.
func (o InstanceOutput) TopicNumOfBuy() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TopicNumOfBuy }).(pulumi.IntOutput)
}

// The max num of topic can be creation of the instance.
// It has been deprecated since version 1.194.0 and using `partitionNum` instead.
// Currently, its value only can be set to 50 when creating it, and finally depends on `partitionNum` value: <`topicQuota`> = 1000 + <`partitionNum`>.
// Therefore, you can update it by updating the `partitionNum`, and it is the only updating path.
//
// Deprecated: Attribute `topic_quota` has been deprecated since 1.194.0 and it will be removed in the next future. Using new attribute `partition_num` instead.
func (o InstanceOutput) TopicQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TopicQuota }).(pulumi.IntOutput)
}

// (Available since v1.214.1) The number of used topics.
func (o InstanceOutput) TopicUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.TopicUsed }).(pulumi.IntOutput)
}

// The VPC ID of the instance.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of attaching vswitch to instance.
func (o InstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The zone ID of the instance.
func (o InstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
