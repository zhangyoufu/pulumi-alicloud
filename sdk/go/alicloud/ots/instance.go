// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource will help you to manager a [Table Store](https://www.alibabacloud.com/help/doc-detail/27280.htm) Instance.
// It is foundation of creating data table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ots"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ots.NewInstance(ctx, "foo", &ots.InstanceArgs{
// 			AccessedBy:  pulumi.String("Vpc"),
// 			Description: pulumi.String("for table"),
// 			Tags: pulumi.StringMap{
// 				"Created": pulumi.String("TF"),
// 				"For":     pulumi.String("Building table"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The network limitation of accessing instance. Valid values:
	AccessedBy pulumi.StringPtrOutput `pulumi:"accessedBy"`
	// The description of the instance. Currently, it does not support modifying.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// The name of the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of tags to assign to the instance.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:ots/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:ots/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The network limitation of accessing instance. Valid values:
	AccessedBy *string `pulumi:"accessedBy"`
	// The description of the instance. Currently, it does not support modifying.
	Description *string `pulumi:"description"`
	// The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
	InstanceType *string `pulumi:"instanceType"`
	// The name of the instance.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

type InstanceState struct {
	// The network limitation of accessing instance. Valid values:
	AccessedBy pulumi.StringPtrInput
	// The description of the instance. Currently, it does not support modifying.
	Description pulumi.StringPtrInput
	// The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
	InstanceType pulumi.StringPtrInput
	// The name of the instance.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the instance.
	Tags pulumi.MapInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The network limitation of accessing instance. Valid values:
	AccessedBy *string `pulumi:"accessedBy"`
	// The description of the instance. Currently, it does not support modifying.
	Description *string `pulumi:"description"`
	// The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
	InstanceType *string `pulumi:"instanceType"`
	// The name of the instance.
	Name *string `pulumi:"name"`
	// A mapping of tags to assign to the instance.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The network limitation of accessing instance. Valid values:
	AccessedBy pulumi.StringPtrInput
	// The description of the instance. Currently, it does not support modifying.
	Description pulumi.StringPtrInput
	// The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".
	InstanceType pulumi.StringPtrInput
	// The name of the instance.
	Name pulumi.StringPtrInput
	// A mapping of tags to assign to the instance.
	Tags pulumi.MapInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
