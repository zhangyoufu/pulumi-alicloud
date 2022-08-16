// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud SSO User Attachment resource.
//
// For information about Cloud SSO User Attachment and how to use it, see [What is User Attachment](https://www.alibabacloud.com/help/en/doc-detail/264683.htm).
//
// > **NOTE:** Available in v1.141.0+.
//
// > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
//
// ## Import
//
// Cloud SSO User Attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudsso/userAttachment:UserAttachment example <directory_id>:<group_id>:<user_id>
//
// ```
type UserAttachment struct {
	pulumi.CustomResourceState

	// The ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// The Group ID.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The User ID.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserAttachment(ctx *pulumi.Context,
	name string, args *UserAttachmentArgs, opts ...pulumi.ResourceOption) (*UserAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource UserAttachment
	err := ctx.RegisterResource("alicloud:cloudsso/userAttachment:UserAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAttachment gets an existing UserAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAttachmentState, opts ...pulumi.ResourceOption) (*UserAttachment, error) {
	var resource UserAttachment
	err := ctx.ReadResource("alicloud:cloudsso/userAttachment:UserAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAttachment resources.
type userAttachmentState struct {
	// The ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// The Group ID.
	GroupId *string `pulumi:"groupId"`
	// The User ID.
	UserId *string `pulumi:"userId"`
}

type UserAttachmentState struct {
	// The ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// The Group ID.
	GroupId pulumi.StringPtrInput
	// The User ID.
	UserId pulumi.StringPtrInput
}

func (UserAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttachmentState)(nil)).Elem()
}

type userAttachmentArgs struct {
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// The Group ID.
	GroupId string `pulumi:"groupId"`
	// The User ID.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserAttachment resource.
type UserAttachmentArgs struct {
	// The ID of the Directory.
	DirectoryId pulumi.StringInput
	// The Group ID.
	GroupId pulumi.StringInput
	// The User ID.
	UserId pulumi.StringInput
}

func (UserAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttachmentArgs)(nil)).Elem()
}

type UserAttachmentInput interface {
	pulumi.Input

	ToUserAttachmentOutput() UserAttachmentOutput
	ToUserAttachmentOutputWithContext(ctx context.Context) UserAttachmentOutput
}

func (*UserAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAttachment)(nil)).Elem()
}

func (i *UserAttachment) ToUserAttachmentOutput() UserAttachmentOutput {
	return i.ToUserAttachmentOutputWithContext(context.Background())
}

func (i *UserAttachment) ToUserAttachmentOutputWithContext(ctx context.Context) UserAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttachmentOutput)
}

// UserAttachmentArrayInput is an input type that accepts UserAttachmentArray and UserAttachmentArrayOutput values.
// You can construct a concrete instance of `UserAttachmentArrayInput` via:
//
//	UserAttachmentArray{ UserAttachmentArgs{...} }
type UserAttachmentArrayInput interface {
	pulumi.Input

	ToUserAttachmentArrayOutput() UserAttachmentArrayOutput
	ToUserAttachmentArrayOutputWithContext(context.Context) UserAttachmentArrayOutput
}

type UserAttachmentArray []UserAttachmentInput

func (UserAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAttachment)(nil)).Elem()
}

func (i UserAttachmentArray) ToUserAttachmentArrayOutput() UserAttachmentArrayOutput {
	return i.ToUserAttachmentArrayOutputWithContext(context.Background())
}

func (i UserAttachmentArray) ToUserAttachmentArrayOutputWithContext(ctx context.Context) UserAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttachmentArrayOutput)
}

// UserAttachmentMapInput is an input type that accepts UserAttachmentMap and UserAttachmentMapOutput values.
// You can construct a concrete instance of `UserAttachmentMapInput` via:
//
//	UserAttachmentMap{ "key": UserAttachmentArgs{...} }
type UserAttachmentMapInput interface {
	pulumi.Input

	ToUserAttachmentMapOutput() UserAttachmentMapOutput
	ToUserAttachmentMapOutputWithContext(context.Context) UserAttachmentMapOutput
}

type UserAttachmentMap map[string]UserAttachmentInput

func (UserAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAttachment)(nil)).Elem()
}

func (i UserAttachmentMap) ToUserAttachmentMapOutput() UserAttachmentMapOutput {
	return i.ToUserAttachmentMapOutputWithContext(context.Background())
}

func (i UserAttachmentMap) ToUserAttachmentMapOutputWithContext(ctx context.Context) UserAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttachmentMapOutput)
}

type UserAttachmentOutput struct{ *pulumi.OutputState }

func (UserAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAttachment)(nil)).Elem()
}

func (o UserAttachmentOutput) ToUserAttachmentOutput() UserAttachmentOutput {
	return o
}

func (o UserAttachmentOutput) ToUserAttachmentOutputWithContext(ctx context.Context) UserAttachmentOutput {
	return o
}

// The ID of the Directory.
func (o UserAttachmentOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttachment) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// The Group ID.
func (o UserAttachmentOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttachment) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The User ID.
func (o UserAttachmentOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttachment) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UserAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAttachment)(nil)).Elem()
}

func (o UserAttachmentArrayOutput) ToUserAttachmentArrayOutput() UserAttachmentArrayOutput {
	return o
}

func (o UserAttachmentArrayOutput) ToUserAttachmentArrayOutputWithContext(ctx context.Context) UserAttachmentArrayOutput {
	return o
}

func (o UserAttachmentArrayOutput) Index(i pulumi.IntInput) UserAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserAttachment {
		return vs[0].([]*UserAttachment)[vs[1].(int)]
	}).(UserAttachmentOutput)
}

type UserAttachmentMapOutput struct{ *pulumi.OutputState }

func (UserAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAttachment)(nil)).Elem()
}

func (o UserAttachmentMapOutput) ToUserAttachmentMapOutput() UserAttachmentMapOutput {
	return o
}

func (o UserAttachmentMapOutput) ToUserAttachmentMapOutputWithContext(ctx context.Context) UserAttachmentMapOutput {
	return o
}

func (o UserAttachmentMapOutput) MapIndex(k pulumi.StringInput) UserAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserAttachment {
		return vs[0].(map[string]*UserAttachment)[vs[1].(string)]
	}).(UserAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttachmentInput)(nil)).Elem(), &UserAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttachmentArrayInput)(nil)).Elem(), UserAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttachmentMapInput)(nil)).Elem(), UserAttachmentMap{})
	pulumi.RegisterOutputType(UserAttachmentOutput{})
	pulumi.RegisterOutputType(UserAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UserAttachmentMapOutput{})
}
