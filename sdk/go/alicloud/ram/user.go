// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM User resource.
//
// For information about RAM User and how to use it, see [What is User](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ram-2015-05-01-createuser).
//
// > **NOTE:** When you want to destroy this resource forcefully(means release all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
//
// > **NOTE:** Available since v1.0.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ram.NewUser(ctx, "user", &ram.UserArgs{
//				Comments:    pulumi.String("yoyoyo"),
//				DisplayName: pulumi.String("user_display_name"),
//				Email:       pulumi.String("hello.uuu@aaa.com"),
//				Mobile:      pulumi.String("86-18688888888"),
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
// RAM User can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ram/user:User example 123456789xxx
// ```
type User struct {
	pulumi.CustomResourceState

	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Email of the RAM user.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// This parameter is used for resource destroy. Default value: `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile pulumi.StringPtrOutput `pulumi:"mobile"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("alicloud:ram/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("alicloud:ram/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments *string `pulumi:"comments"`
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName *string `pulumi:"displayName"`
	// Email of the RAM user.
	Email *string `pulumi:"email"`
	// This parameter is used for resource destroy. Default value: `false`.
	Force *bool `pulumi:"force"`
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile *string `pulumi:"mobile"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
}

type UserState struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrInput
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName pulumi.StringPtrInput
	// Email of the RAM user.
	Email pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value: `false`.
	Force pulumi.BoolPtrInput
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile pulumi.StringPtrInput
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments *string `pulumi:"comments"`
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName *string `pulumi:"displayName"`
	// Email of the RAM user.
	Email *string `pulumi:"email"`
	// This parameter is used for resource destroy. Default value: `false`.
	Force *bool `pulumi:"force"`
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile *string `pulumi:"mobile"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
	Comments pulumi.StringPtrInput
	// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
	DisplayName pulumi.StringPtrInput
	// Email of the RAM user.
	Email pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value: `false`.
	Force pulumi.BoolPtrInput
	// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
	Mobile pulumi.StringPtrInput
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Comment of the RAM user. This parameter can have a string of 1 to 128 characters.
func (o UserOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.
func (o UserOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Email of the RAM user.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// This parameter is used for resource destroy. Default value: `false`.
func (o UserOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.
func (o UserOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Mobile }).(pulumi.StringPtrOutput)
}

// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
