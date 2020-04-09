// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS read write splitting connection resource to allocate an Intranet connection string for RDS instance.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_read_write_splitting_connection.html.markdown.
type ReadWriteSplittingConnection struct {
	pulumi.CustomResourceState

	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrOutput `pulumi:"connectionPrefix"`
	// Connection instance string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringOutput `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntOutput `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntOutput `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapOutput `pulumi:"weight"`
}

// NewReadWriteSplittingConnection registers a new resource with the given unique name, arguments, and options.
func NewReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, args *ReadWriteSplittingConnectionArgs, opts ...pulumi.ResourceOption) (*ReadWriteSplittingConnection, error) {
	if args == nil || args.DistributionType == nil {
		return nil, errors.New("missing required argument 'DistributionType'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &ReadWriteSplittingConnectionArgs{}
	}
	var resource ReadWriteSplittingConnection
	err := ctx.RegisterResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadWriteSplittingConnection gets an existing ReadWriteSplittingConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadWriteSplittingConnectionState, opts ...pulumi.ResourceOption) (*ReadWriteSplittingConnection, error) {
	var resource ReadWriteSplittingConnection
	err := ctx.ReadResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadWriteSplittingConnection resources.
type readWriteSplittingConnectionState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Connection instance string.
	ConnectionString *string `pulumi:"connectionString"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType *string `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId *string `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime *int `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port *int `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight map[string]interface{} `pulumi:"weight"`
}

type ReadWriteSplittingConnectionState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrInput
	// Connection instance string.
	ConnectionString pulumi.StringPtrInput
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringPtrInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringPtrInput
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntPtrInput
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntPtrInput
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapInput
}

func (ReadWriteSplittingConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteSplittingConnectionState)(nil)).Elem()
}

type readWriteSplittingConnectionArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType string `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId string `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime *int `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port *int `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight map[string]interface{} `pulumi:"weight"`
}

// The set of arguments for constructing a ReadWriteSplittingConnection resource.
type ReadWriteSplittingConnectionArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrInput
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringInput
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntPtrInput
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntPtrInput
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapInput
}

func (ReadWriteSplittingConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteSplittingConnectionArgs)(nil)).Elem()
}
