// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicemesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Service Mesh UserPermission resource.
//
// For information about Service Mesh User Permission and how to use it, see [What is User Permission](https://www.alibabacloud.com/help/en/alibaba-cloud-service-mesh/latest/api-servicemesh-2020-01-11-grantuserpermissions).
//
// > **NOTE:** Available since v1.174.0.
//
// ## Import
//
// Service Mesh User Permission can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:servicemesh/userPermission:UserPermission example <id>
//
// ```
type UserPermission struct {
	pulumi.CustomResourceState

	// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
	Permissions UserPermissionPermissionArrayOutput `pulumi:"permissions"`
	// The configuration of the Load Balancer. See the following `Block loadBalancer`.
	SubAccountUserId pulumi.StringOutput `pulumi:"subAccountUserId"`
}

// NewUserPermission registers a new resource with the given unique name, arguments, and options.
func NewUserPermission(ctx *pulumi.Context,
	name string, args *UserPermissionArgs, opts ...pulumi.ResourceOption) (*UserPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubAccountUserId == nil {
		return nil, errors.New("invalid value for required argument 'SubAccountUserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPermission
	err := ctx.RegisterResource("alicloud:servicemesh/userPermission:UserPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPermission gets an existing UserPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPermissionState, opts ...pulumi.ResourceOption) (*UserPermission, error) {
	var resource UserPermission
	err := ctx.ReadResource("alicloud:servicemesh/userPermission:UserPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPermission resources.
type userPermissionState struct {
	// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
	Permissions []UserPermissionPermission `pulumi:"permissions"`
	// The configuration of the Load Balancer. See the following `Block loadBalancer`.
	SubAccountUserId *string `pulumi:"subAccountUserId"`
}

type UserPermissionState struct {
	// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
	Permissions UserPermissionPermissionArrayInput
	// The configuration of the Load Balancer. See the following `Block loadBalancer`.
	SubAccountUserId pulumi.StringPtrInput
}

func (UserPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPermissionState)(nil)).Elem()
}

type userPermissionArgs struct {
	// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
	Permissions []UserPermissionPermission `pulumi:"permissions"`
	// The configuration of the Load Balancer. See the following `Block loadBalancer`.
	SubAccountUserId string `pulumi:"subAccountUserId"`
}

// The set of arguments for constructing a UserPermission resource.
type UserPermissionArgs struct {
	// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
	Permissions UserPermissionPermissionArrayInput
	// The configuration of the Load Balancer. See the following `Block loadBalancer`.
	SubAccountUserId pulumi.StringInput
}

func (UserPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPermissionArgs)(nil)).Elem()
}

type UserPermissionInput interface {
	pulumi.Input

	ToUserPermissionOutput() UserPermissionOutput
	ToUserPermissionOutputWithContext(ctx context.Context) UserPermissionOutput
}

func (*UserPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPermission)(nil)).Elem()
}

func (i *UserPermission) ToUserPermissionOutput() UserPermissionOutput {
	return i.ToUserPermissionOutputWithContext(context.Background())
}

func (i *UserPermission) ToUserPermissionOutputWithContext(ctx context.Context) UserPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionOutput)
}

func (i *UserPermission) ToOutput(ctx context.Context) pulumix.Output[*UserPermission] {
	return pulumix.Output[*UserPermission]{
		OutputState: i.ToUserPermissionOutputWithContext(ctx).OutputState,
	}
}

// UserPermissionArrayInput is an input type that accepts UserPermissionArray and UserPermissionArrayOutput values.
// You can construct a concrete instance of `UserPermissionArrayInput` via:
//
//	UserPermissionArray{ UserPermissionArgs{...} }
type UserPermissionArrayInput interface {
	pulumi.Input

	ToUserPermissionArrayOutput() UserPermissionArrayOutput
	ToUserPermissionArrayOutputWithContext(context.Context) UserPermissionArrayOutput
}

type UserPermissionArray []UserPermissionInput

func (UserPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPermission)(nil)).Elem()
}

func (i UserPermissionArray) ToUserPermissionArrayOutput() UserPermissionArrayOutput {
	return i.ToUserPermissionArrayOutputWithContext(context.Background())
}

func (i UserPermissionArray) ToUserPermissionArrayOutputWithContext(ctx context.Context) UserPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionArrayOutput)
}

func (i UserPermissionArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserPermission] {
	return pulumix.Output[[]*UserPermission]{
		OutputState: i.ToUserPermissionArrayOutputWithContext(ctx).OutputState,
	}
}

// UserPermissionMapInput is an input type that accepts UserPermissionMap and UserPermissionMapOutput values.
// You can construct a concrete instance of `UserPermissionMapInput` via:
//
//	UserPermissionMap{ "key": UserPermissionArgs{...} }
type UserPermissionMapInput interface {
	pulumi.Input

	ToUserPermissionMapOutput() UserPermissionMapOutput
	ToUserPermissionMapOutputWithContext(context.Context) UserPermissionMapOutput
}

type UserPermissionMap map[string]UserPermissionInput

func (UserPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPermission)(nil)).Elem()
}

func (i UserPermissionMap) ToUserPermissionMapOutput() UserPermissionMapOutput {
	return i.ToUserPermissionMapOutputWithContext(context.Background())
}

func (i UserPermissionMap) ToUserPermissionMapOutputWithContext(ctx context.Context) UserPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionMapOutput)
}

func (i UserPermissionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserPermission] {
	return pulumix.Output[map[string]*UserPermission]{
		OutputState: i.ToUserPermissionMapOutputWithContext(ctx).OutputState,
	}
}

type UserPermissionOutput struct{ *pulumi.OutputState }

func (UserPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPermission)(nil)).Elem()
}

func (o UserPermissionOutput) ToUserPermissionOutput() UserPermissionOutput {
	return o
}

func (o UserPermissionOutput) ToUserPermissionOutputWithContext(ctx context.Context) UserPermissionOutput {
	return o
}

func (o UserPermissionOutput) ToOutput(ctx context.Context) pulumix.Output[*UserPermission] {
	return pulumix.Output[*UserPermission]{
		OutputState: o.OutputState,
	}
}

// List of permissions. **Warning:** The list requires the full amount of permission information to be passed. Adding permissions means adding items to the list, and deleting them or inputting nothing means removing items. See `permissions` below.
func (o UserPermissionOutput) Permissions() UserPermissionPermissionArrayOutput {
	return o.ApplyT(func(v *UserPermission) UserPermissionPermissionArrayOutput { return v.Permissions }).(UserPermissionPermissionArrayOutput)
}

// The configuration of the Load Balancer. See the following `Block loadBalancer`.
func (o UserPermissionOutput) SubAccountUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPermission) pulumi.StringOutput { return v.SubAccountUserId }).(pulumi.StringOutput)
}

type UserPermissionArrayOutput struct{ *pulumi.OutputState }

func (UserPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPermission)(nil)).Elem()
}

func (o UserPermissionArrayOutput) ToUserPermissionArrayOutput() UserPermissionArrayOutput {
	return o
}

func (o UserPermissionArrayOutput) ToUserPermissionArrayOutputWithContext(ctx context.Context) UserPermissionArrayOutput {
	return o
}

func (o UserPermissionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserPermission] {
	return pulumix.Output[[]*UserPermission]{
		OutputState: o.OutputState,
	}
}

func (o UserPermissionArrayOutput) Index(i pulumi.IntInput) UserPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPermission {
		return vs[0].([]*UserPermission)[vs[1].(int)]
	}).(UserPermissionOutput)
}

type UserPermissionMapOutput struct{ *pulumi.OutputState }

func (UserPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPermission)(nil)).Elem()
}

func (o UserPermissionMapOutput) ToUserPermissionMapOutput() UserPermissionMapOutput {
	return o
}

func (o UserPermissionMapOutput) ToUserPermissionMapOutputWithContext(ctx context.Context) UserPermissionMapOutput {
	return o
}

func (o UserPermissionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserPermission] {
	return pulumix.Output[map[string]*UserPermission]{
		OutputState: o.OutputState,
	}
}

func (o UserPermissionMapOutput) MapIndex(k pulumi.StringInput) UserPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPermission {
		return vs[0].(map[string]*UserPermission)[vs[1].(string)]
	}).(UserPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionInput)(nil)).Elem(), &UserPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionArrayInput)(nil)).Elem(), UserPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionMapInput)(nil)).Elem(), UserPermissionMap{})
	pulumi.RegisterOutputType(UserPermissionOutput{})
	pulumi.RegisterOutputType(UserPermissionArrayOutput{})
	pulumi.RegisterOutputType(UserPermissionMapOutput{})
}
