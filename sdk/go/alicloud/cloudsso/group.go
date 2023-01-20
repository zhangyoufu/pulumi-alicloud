// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud SSO Group resource.
//
// For information about Cloud SSO Group and how to use it, see [What is Group](https://www.alibabacloud.com/help/doc-detail/264683.html).
//
// > **NOTE:** Available in v1.138.0+.
//
// ## Import
//
// Cloud SSO Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudsso/group:Group example <directory_id>:<group_id>
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// The Description of the group. The description can be up to `1024` characters long.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The GroupId of the group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	GroupName pulumi.StringOutput `pulumi:"groupName"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	var resource Group
	err := ctx.RegisterResource("alicloud:cloudsso/group:Group", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:cloudsso/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The Description of the group. The description can be up to `1024` characters long.
	Description *string `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// The GroupId of the group.
	GroupId *string `pulumi:"groupId"`
	// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	GroupName *string `pulumi:"groupName"`
}

type GroupState struct {
	// The Description of the group. The description can be up to `1024` characters long.
	Description pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// The GroupId of the group.
	GroupId pulumi.StringPtrInput
	// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	GroupName pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The Description of the group. The description can be up to `1024` characters long.
	Description *string `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	GroupName string `pulumi:"groupName"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The Description of the group. The description can be up to `1024` characters long.
	Description pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringInput
	// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	GroupName pulumi.StringInput
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

// The Description of the group. The description can be up to `1024` characters long.
func (o GroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Directory.
func (o GroupOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The GroupId of the group.
func (o GroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The Name of the group. The name must be `1` to `128` characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
func (o GroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
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
