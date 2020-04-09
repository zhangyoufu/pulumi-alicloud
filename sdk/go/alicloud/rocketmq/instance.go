// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rocketmq

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ONS instance resource.
//
// For more information about how to use it, see [RocketMQ Instance Management API](https://www.alibabacloud.com/help/doc-detail/106354.html).
//
// > **NOTE:** The number of instances in the same region cannot exceed 8. At present, the resource does not support region "mq-internet-access" and "ap-southeast-5".
//
// > **NOTE:** Available in 1.51.0+
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ons_instance.html.markdown.
type Instance struct {
	pulumi.CustomResourceState

	// The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
	InstanceStatus pulumi.IntOutput `pulumi:"instanceStatus"`
	// The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
	InstanceType pulumi.IntOutput `pulumi:"instanceType"`
	// Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Platinum edition instance expiration time.
	ReleaseTime pulumi.StringOutput `pulumi:"releaseTime"`
	// This attribute is a concise description of instance. The length cannot exceed 128.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:rocketmq/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:rocketmq/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
	InstanceStatus *int `pulumi:"instanceStatus"`
	// The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
	InstanceType *int `pulumi:"instanceType"`
	// Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
	Name *string `pulumi:"name"`
	// Platinum edition instance expiration time.
	ReleaseTime *string `pulumi:"releaseTime"`
	// This attribute is a concise description of instance. The length cannot exceed 128.
	Remark *string `pulumi:"remark"`
}

type InstanceState struct {
	// The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.
	InstanceStatus pulumi.IntPtrInput
	// The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.
	InstanceType pulumi.IntPtrInput
	// Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
	Name pulumi.StringPtrInput
	// Platinum edition instance expiration time.
	ReleaseTime pulumi.StringPtrInput
	// This attribute is a concise description of instance. The length cannot exceed 128.
	Remark pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
	Name *string `pulumi:"name"`
	// This attribute is a concise description of instance. The length cannot exceed 128.
	Remark *string `pulumi:"remark"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.
	Name pulumi.StringPtrInput
	// This attribute is a concise description of instance. The length cannot exceed 128.
	Remark pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
