// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an ONS group resource.
//
// For more information about how to use it, see [RocketMQ Group Management API](https://www.alibabacloud.com/help/doc-detail/29616.html).
//
// > **NOTE:** Available in 1.53.0+
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "onsInstanceName"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			groupName := "GID-onsGroupDatasourceName"
//			if param := cfg.Get("groupName"); param != "" {
//				groupName = param
//			}
//			defaultInstance, err := rocketmq.NewInstance(ctx, "defaultInstance", &rocketmq.InstanceArgs{
//				Remark: pulumi.String("default_ons_instance_remark"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rocketmq.NewGroup(ctx, "defaultGroup", &rocketmq.GroupArgs{
//				GroupName:  pulumi.String(groupName),
//				InstanceId: defaultInstance.ID(),
//				Remark:     pulumi.String("dafault_ons_group_remark"),
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
// ONS GROUP can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rocketmq/group:Group group MQ_INST_1234567890_Baso1234567:GID-onsGroupDemo
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// Replaced by `groupName` after version 1.98.0.
	//
	// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType pulumi.StringPtrOutput `pulumi:"groupType"`
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrOutput `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("alicloud:rocketmq/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("alicloud:rocketmq/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Replaced by `groupName` after version 1.98.0.
	//
	// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
	GroupId *string `pulumi:"groupId"`
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupName *string `pulumi:"groupName"`
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType *string `pulumi:"groupType"`
	// ID of the ONS Instance that owns the groups.
	InstanceId *string `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable *bool `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark *string `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

type GroupState struct {
	// Replaced by `groupName` after version 1.98.0.
	//
	// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
	GroupId pulumi.StringPtrInput
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupName pulumi.StringPtrInput
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType pulumi.StringPtrInput
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringPtrInput
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrInput
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Replaced by `groupName` after version 1.98.0.
	//
	// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
	GroupId *string `pulumi:"groupId"`
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupName *string `pulumi:"groupName"`
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType *string `pulumi:"groupType"`
	// ID of the ONS Instance that owns the groups.
	InstanceId string `pulumi:"instanceId"`
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable *bool `pulumi:"readEnable"`
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark *string `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Replaced by `groupName` after version 1.98.0.
	//
	// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
	GroupId pulumi.StringPtrInput
	// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
	GroupName pulumi.StringPtrInput
	// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
	GroupType pulumi.StringPtrInput
	// ID of the ONS Instance that owns the groups.
	InstanceId pulumi.StringInput
	// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
	ReadEnable pulumi.BoolPtrInput
	// This attribute is a concise description of group. The length cannot exceed 256.
	Remark pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: i.ToGroupOutputWithContext(ctx).OutputState,
	}
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

func (i GroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*Group] {
	return pulumix.Output[[]*Group]{
		OutputState: i.ToGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

func (i GroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Group] {
	return pulumix.Output[map[string]*Group]{
		OutputState: i.ToGroupMapOutputWithContext(ctx).OutputState,
	}
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: o.OutputState,
	}
}

// Replaced by `groupName` after version 1.98.0.
//
// Deprecated: Field 'group_id' has been deprecated from version 1.98.0. Use 'group_name' instead.
func (o GroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Name of the group. Two groups on a single instance cannot have the same name. A `groupName` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).
func (o GroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
func (o GroupOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.GroupType }).(pulumi.StringPtrOutput)
}

// ID of the ONS Instance that owns the groups.
func (o GroupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client.
func (o GroupOutput) ReadEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.ReadEnable }).(pulumi.BoolPtrOutput)
}

// This attribute is a concise description of group. The length cannot exceed 256.
func (o GroupOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the resource.
// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
func (o GroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Group) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Group] {
	return pulumix.Output[[]*Group]{
		OutputState: o.OutputState,
	}
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Group] {
	return pulumix.Output[map[string]*Group]{
		OutputState: o.OutputState,
	}
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
