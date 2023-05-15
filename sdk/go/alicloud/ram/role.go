// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM Role resource.
//
// > **NOTE:** When you want to destroy this resource forcefully(means remove all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
//
// ## Example Usage
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
//			_, err := ram.NewRole(ctx, "role", &ram.RoleArgs{
//				Description: pulumi.String("this is a role test."),
//				Document:    pulumi.String("  {\n    \"Statement\": [\n      {\n        \"Action\": \"sts:AssumeRole\",\n        \"Effect\": \"Allow\",\n        \"Principal\": {\n          \"Service\": [\n            \"apigateway.aliyuncs.com\", \n            \"ecs.aliyuncs.com\"\n          ]\n        }\n      }\n    ],\n    \"Version\": \"1\"\n  }\n  \n"),
//				Force:       pulumi.Bool(true),
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
// RAM role can be imported using the id or name, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/role:Role example my-role
//
// ```
type Role struct {
	pulumi.CustomResourceState

	// The role arn.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
	Document pulumi.StringOutput `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
	MaxSessionDuration pulumi.IntPtrOutput `pulumi:"maxSessionDuration"`
	// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	Name pulumi.StringOutput `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
	//
	// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
	RamUsers pulumi.StringArrayOutput `pulumi:"ramUsers"`
	// The role ID.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
	//
	// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		args = &RoleArgs{}
	}

	var resource Role
	err := ctx.RegisterResource("alicloud:ram/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("alicloud:ram/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The role arn.
	Arn *string `pulumi:"arn"`
	// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
	Description *string `pulumi:"description"`
	// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
	//
	// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
	RamUsers []string `pulumi:"ramUsers"`
	// The role ID.
	RoleId *string `pulumi:"roleId"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
	//
	// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Services []string `pulumi:"services"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
}

type RoleState struct {
	// The role arn.
	Arn pulumi.StringPtrInput
	// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
	Description pulumi.StringPtrInput
	// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
	MaxSessionDuration pulumi.IntPtrInput
	// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
	//
	// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
	RamUsers pulumi.StringArrayInput
	// The role ID.
	RoleId pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
	//
	// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Services pulumi.StringArrayInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
	Description *string `pulumi:"description"`
	// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
	//
	// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
	RamUsers []string `pulumi:"ramUsers"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
	//
	// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Services []string `pulumi:"services"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
	Description pulumi.StringPtrInput
	// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
	MaxSessionDuration pulumi.IntPtrInput
	// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
	//
	// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
	RamUsers pulumi.StringArrayInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
	//
	// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Services pulumi.StringArrayInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// The role arn.
func (o RoleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Authorization strategy of the RAM role. It is required when the `services` and `ramUsers` are not specified.
func (o RoleOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Document }).(pulumi.StringOutput)
}

// This parameter is used for resource destroy. Default value is `false`.
func (o RoleOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
func (o RoleOutput) MaxSessionDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.IntPtrOutput { return v.MaxSessionDuration }).(pulumi.IntPtrOutput)
}

// Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
//
// Deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace.
func (o RoleOutput) RamUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.RamUsers }).(pulumi.StringArrayOutput)
}

// The role ID.
func (o RoleOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
//
// Deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace.
func (o RoleOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
//
// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
func (o RoleOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
