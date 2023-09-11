// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Elastic Desktop Service (ECD) User resource.
//
// For information about Elastic Desktop Service (ECD) User and how to use it, see [What is User](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/api-doc-eds-user-2021-03-08-api-doc-createusers-desktop).
//
// > **NOTE:** Available since v1.142.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eds.NewUser(ctx, "default", &eds.UserArgs{
//				Email:     pulumi.String("tf.example@abc.com"),
//				EndUserId: pulumi.String("terraform_example123"),
//				Password:  pulumi.String("Example_123"),
//				Phone:     pulumi.String("18888888888"),
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
// ECD User can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:eds/user:User example <end_user_id>
//
// ```
type User struct {
	pulumi.CustomResourceState

	// The email of the user email.
	Email pulumi.StringOutput `pulumi:"email"`
	// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
	EndUserId pulumi.StringOutput `pulumi:"endUserId"`
	// The password of the user password.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The phone of the mobile phone number.
	Phone pulumi.StringPtrOutput `pulumi:"phone"`
	// The status of the resource. Valid values: `Unlocked`, `Locked`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.EndUserId == nil {
		return nil, errors.New("invalid value for required argument 'EndUserId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("alicloud:eds/user:User", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:eds/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The email of the user email.
	Email *string `pulumi:"email"`
	// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
	EndUserId *string `pulumi:"endUserId"`
	// The password of the user password.
	Password *string `pulumi:"password"`
	// The phone of the mobile phone number.
	Phone *string `pulumi:"phone"`
	// The status of the resource. Valid values: `Unlocked`, `Locked`.
	Status *string `pulumi:"status"`
}

type UserState struct {
	// The email of the user email.
	Email pulumi.StringPtrInput
	// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
	EndUserId pulumi.StringPtrInput
	// The password of the user password.
	Password pulumi.StringPtrInput
	// The phone of the mobile phone number.
	Phone pulumi.StringPtrInput
	// The status of the resource. Valid values: `Unlocked`, `Locked`.
	Status pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The email of the user email.
	Email string `pulumi:"email"`
	// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
	EndUserId string `pulumi:"endUserId"`
	// The password of the user password.
	Password *string `pulumi:"password"`
	// The phone of the mobile phone number.
	Phone *string `pulumi:"phone"`
	// The status of the resource. Valid values: `Unlocked`, `Locked`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The email of the user email.
	Email pulumi.StringInput
	// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
	EndUserId pulumi.StringInput
	// The password of the user password.
	Password pulumi.StringPtrInput
	// The phone of the mobile phone number.
	Phone pulumi.StringPtrInput
	// The status of the resource. Valid values: `Unlocked`, `Locked`.
	Status pulumi.StringPtrInput
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

func (i *User) ToOutput(ctx context.Context) pulumix.Output[*User] {
	return pulumix.Output[*User]{
		OutputState: i.ToUserOutputWithContext(ctx).OutputState,
	}
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

func (i UserArray) ToOutput(ctx context.Context) pulumix.Output[[]*User] {
	return pulumix.Output[[]*User]{
		OutputState: i.ToUserArrayOutputWithContext(ctx).OutputState,
	}
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

func (i UserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*User] {
	return pulumix.Output[map[string]*User]{
		OutputState: i.ToUserMapOutputWithContext(ctx).OutputState,
	}
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

func (o UserOutput) ToOutput(ctx context.Context) pulumix.Output[*User] {
	return pulumix.Output[*User]{
		OutputState: o.OutputState,
	}
}

// The email of the user email.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The Username. The custom setting is composed of lowercase letters, numbers and underscores, and the length is 3~24 characters.
func (o UserOutput) EndUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.EndUserId }).(pulumi.StringOutput)
}

// The password of the user password.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The phone of the mobile phone number.
func (o UserOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Phone }).(pulumi.StringPtrOutput)
}

// The status of the resource. Valid values: `Unlocked`, `Locked`.
func (o UserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
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

func (o UserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*User] {
	return pulumix.Output[[]*User]{
		OutputState: o.OutputState,
	}
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

func (o UserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*User] {
	return pulumix.Output[map[string]*User]{
		OutputState: o.OutputState,
	}
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
