// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a AnalyticDB for PostgreSQL instance resource supports replica set instances only. the AnalyticDB for PostgreSQL provides stable, reliable, and automatic scalable database services.
// You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/35387.htm)
//
// > **NOTE:**  Available in 1.47.0+
//
// > **NOTE:**  The following regions don't support create Classic network Gpdb instance.
// [`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`,`us-east-1`,`eu-central-1`,`cn-shanghai-finance-1`,`cn-shenzhen-finance-1`,`cn-hangzhou-finance`]
//
// > **NOTE:**  Create instance or change instance would cost 10~15 minutes. Please make full preparation.
type Instance struct {
	pulumi.CustomResourceState

	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Database engine: gpdb. System Default value: gpdb.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringOutput `pulumi:"instanceChargeType"`
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// The number of groups. Valid values: [2,4,8,16,32]
	InstanceGroupCount pulumi.StringOutput `pulumi:"instanceGroupCount"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists pulumi.StringArrayOutput `pulumi:"securityIpLists"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.InstanceClass == nil {
		return nil, errors.New("missing required argument 'InstanceClass'")
	}
	if args == nil || args.InstanceGroupCount == nil {
		return nil, errors.New("missing required argument 'InstanceGroupCount'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:gpdb/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:gpdb/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Description *string `pulumi:"description"`
	// Database engine: gpdb. System Default value: gpdb.
	Engine *string `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
	InstanceClass *string `pulumi:"instanceClass"`
	// The number of groups. Valid values: [2,4,8,16,32]
	InstanceGroupCount *string `pulumi:"instanceGroupCount"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
}

type InstanceState struct {
	AvailabilityZone pulumi.StringPtrInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Description pulumi.StringPtrInput
	// Database engine: gpdb. System Default value: gpdb.
	Engine pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
	EngineVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
	InstanceClass pulumi.StringPtrInput
	// The number of groups. Valid values: [2,4,8,16,32]
	InstanceGroupCount pulumi.StringPtrInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The name of DB instance. It a string of 2 to 256 characters.
	Description *string `pulumi:"description"`
	// Database engine: gpdb. System Default value: gpdb.
	Engine *string `pulumi:"engine"`
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
	InstanceClass string `pulumi:"instanceClass"`
	// The number of groups. Valid values: [2,4,8,16,32]
	InstanceGroupCount string `pulumi:"instanceGroupCount"`
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists []string `pulumi:"securityIpLists"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	AvailabilityZone pulumi.StringPtrInput
	// The name of DB instance. It a string of 2 to 256 characters.
	Description pulumi.StringPtrInput
	// Database engine: gpdb. System Default value: gpdb.
	Engine pulumi.StringPtrInput
	// Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
	EngineVersion pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
	InstanceClass pulumi.StringInput
	// The number of groups. Valid values: [2,4,8,16,32]
	InstanceGroupCount pulumi.StringInput
	// List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
	SecurityIpLists pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The virtual switch ID to launch DB instances in one VPC.
	VswitchId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
