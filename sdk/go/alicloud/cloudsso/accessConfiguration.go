// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudsso

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud SSO Access Configuration resource.
//
// For information about Cloud SSO Access Configuration and how to use it, see [What is Access Configuration](https://www.alibabacloud.com/help/en/doc-detail/266737.html).
//
// > **NOTE:** Available in v1.145.0+.
//
// > **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
//
// ## Import
//
// Cloud SSO Access Configuration can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cloudsso/accessConfiguration:AccessConfiguration example <directory_id>:<access_configuration_id>
//
// ```
type AccessConfiguration struct {
	pulumi.CustomResourceState

	// The AccessConfigurationId of the Access Configuration.
	AccessConfigurationId pulumi.StringOutput `pulumi:"accessConfigurationId"`
	// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
	AccessConfigurationName pulumi.StringOutput `pulumi:"accessConfigurationName"`
	// The Description of the  Access Configuration. The description can be up to `1024` characters long.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
	ForceRemovePermissionPolicies pulumi.BoolPtrOutput `pulumi:"forceRemovePermissionPolicies"`
	// The Policy List. See the following `Block permissionPolicies`.
	PermissionPolicies AccessConfigurationPermissionPolicyArrayOutput `pulumi:"permissionPolicies"`
	// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
	RelayState pulumi.StringPtrOutput `pulumi:"relayState"`
	// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
	SessionDuration pulumi.IntOutput `pulumi:"sessionDuration"`
}

// NewAccessConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAccessConfiguration(ctx *pulumi.Context,
	name string, args *AccessConfigurationArgs, opts ...pulumi.ResourceOption) (*AccessConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'AccessConfigurationName'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	var resource AccessConfiguration
	err := ctx.RegisterResource("alicloud:cloudsso/accessConfiguration:AccessConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessConfiguration gets an existing AccessConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessConfigurationState, opts ...pulumi.ResourceOption) (*AccessConfiguration, error) {
	var resource AccessConfiguration
	err := ctx.ReadResource("alicloud:cloudsso/accessConfiguration:AccessConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessConfiguration resources.
type accessConfigurationState struct {
	// The AccessConfigurationId of the Access Configuration.
	AccessConfigurationId *string `pulumi:"accessConfigurationId"`
	// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
	AccessConfigurationName *string `pulumi:"accessConfigurationName"`
	// The Description of the  Access Configuration. The description can be up to `1024` characters long.
	Description *string `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId *string `pulumi:"directoryId"`
	// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
	ForceRemovePermissionPolicies *bool `pulumi:"forceRemovePermissionPolicies"`
	// The Policy List. See the following `Block permissionPolicies`.
	PermissionPolicies []AccessConfigurationPermissionPolicy `pulumi:"permissionPolicies"`
	// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
	RelayState *string `pulumi:"relayState"`
	// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
	SessionDuration *int `pulumi:"sessionDuration"`
}

type AccessConfigurationState struct {
	// The AccessConfigurationId of the Access Configuration.
	AccessConfigurationId pulumi.StringPtrInput
	// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
	AccessConfigurationName pulumi.StringPtrInput
	// The Description of the  Access Configuration. The description can be up to `1024` characters long.
	Description pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringPtrInput
	// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
	ForceRemovePermissionPolicies pulumi.BoolPtrInput
	// The Policy List. See the following `Block permissionPolicies`.
	PermissionPolicies AccessConfigurationPermissionPolicyArrayInput
	// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
	RelayState pulumi.StringPtrInput
	// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
	SessionDuration pulumi.IntPtrInput
}

func (AccessConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConfigurationState)(nil)).Elem()
}

type accessConfigurationArgs struct {
	// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
	AccessConfigurationName string `pulumi:"accessConfigurationName"`
	// The Description of the  Access Configuration. The description can be up to `1024` characters long.
	Description *string `pulumi:"description"`
	// The ID of the Directory.
	DirectoryId string `pulumi:"directoryId"`
	// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
	ForceRemovePermissionPolicies *bool `pulumi:"forceRemovePermissionPolicies"`
	// The Policy List. See the following `Block permissionPolicies`.
	PermissionPolicies []AccessConfigurationPermissionPolicy `pulumi:"permissionPolicies"`
	// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
	RelayState *string `pulumi:"relayState"`
	// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
	SessionDuration *int `pulumi:"sessionDuration"`
}

// The set of arguments for constructing a AccessConfiguration resource.
type AccessConfigurationArgs struct {
	// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
	AccessConfigurationName pulumi.StringInput
	// The Description of the  Access Configuration. The description can be up to `1024` characters long.
	Description pulumi.StringPtrInput
	// The ID of the Directory.
	DirectoryId pulumi.StringInput
	// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
	ForceRemovePermissionPolicies pulumi.BoolPtrInput
	// The Policy List. See the following `Block permissionPolicies`.
	PermissionPolicies AccessConfigurationPermissionPolicyArrayInput
	// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
	RelayState pulumi.StringPtrInput
	// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
	SessionDuration pulumi.IntPtrInput
}

func (AccessConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessConfigurationArgs)(nil)).Elem()
}

type AccessConfigurationInput interface {
	pulumi.Input

	ToAccessConfigurationOutput() AccessConfigurationOutput
	ToAccessConfigurationOutputWithContext(ctx context.Context) AccessConfigurationOutput
}

func (*AccessConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConfiguration)(nil)).Elem()
}

func (i *AccessConfiguration) ToAccessConfigurationOutput() AccessConfigurationOutput {
	return i.ToAccessConfigurationOutputWithContext(context.Background())
}

func (i *AccessConfiguration) ToAccessConfigurationOutputWithContext(ctx context.Context) AccessConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationOutput)
}

// AccessConfigurationArrayInput is an input type that accepts AccessConfigurationArray and AccessConfigurationArrayOutput values.
// You can construct a concrete instance of `AccessConfigurationArrayInput` via:
//
//	AccessConfigurationArray{ AccessConfigurationArgs{...} }
type AccessConfigurationArrayInput interface {
	pulumi.Input

	ToAccessConfigurationArrayOutput() AccessConfigurationArrayOutput
	ToAccessConfigurationArrayOutputWithContext(context.Context) AccessConfigurationArrayOutput
}

type AccessConfigurationArray []AccessConfigurationInput

func (AccessConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessConfiguration)(nil)).Elem()
}

func (i AccessConfigurationArray) ToAccessConfigurationArrayOutput() AccessConfigurationArrayOutput {
	return i.ToAccessConfigurationArrayOutputWithContext(context.Background())
}

func (i AccessConfigurationArray) ToAccessConfigurationArrayOutputWithContext(ctx context.Context) AccessConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationArrayOutput)
}

// AccessConfigurationMapInput is an input type that accepts AccessConfigurationMap and AccessConfigurationMapOutput values.
// You can construct a concrete instance of `AccessConfigurationMapInput` via:
//
//	AccessConfigurationMap{ "key": AccessConfigurationArgs{...} }
type AccessConfigurationMapInput interface {
	pulumi.Input

	ToAccessConfigurationMapOutput() AccessConfigurationMapOutput
	ToAccessConfigurationMapOutputWithContext(context.Context) AccessConfigurationMapOutput
}

type AccessConfigurationMap map[string]AccessConfigurationInput

func (AccessConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessConfiguration)(nil)).Elem()
}

func (i AccessConfigurationMap) ToAccessConfigurationMapOutput() AccessConfigurationMapOutput {
	return i.ToAccessConfigurationMapOutputWithContext(context.Background())
}

func (i AccessConfigurationMap) ToAccessConfigurationMapOutputWithContext(ctx context.Context) AccessConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessConfigurationMapOutput)
}

type AccessConfigurationOutput struct{ *pulumi.OutputState }

func (AccessConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessConfiguration)(nil)).Elem()
}

func (o AccessConfigurationOutput) ToAccessConfigurationOutput() AccessConfigurationOutput {
	return o
}

func (o AccessConfigurationOutput) ToAccessConfigurationOutputWithContext(ctx context.Context) AccessConfigurationOutput {
	return o
}

// The AccessConfigurationId of the Access Configuration.
func (o AccessConfigurationOutput) AccessConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.StringOutput { return v.AccessConfigurationId }).(pulumi.StringOutput)
}

// The AccessConfigurationName of the Access Configuration. The name of the resource. The name can be up to `32` characters long and can contain letters, digits, and hyphens (-).
func (o AccessConfigurationOutput) AccessConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.StringOutput { return v.AccessConfigurationName }).(pulumi.StringOutput)
}

// The Description of the  Access Configuration. The description can be up to `1024` characters long.
func (o AccessConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Directory.
func (o AccessConfigurationOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

// This parameter is used to force deletion `permissionPolicies`. Valid Value: `true` and `false`.
func (o AccessConfigurationOutput) ForceRemovePermissionPolicies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.BoolPtrOutput { return v.ForceRemovePermissionPolicies }).(pulumi.BoolPtrOutput)
}

// The Policy List. See the following `Block permissionPolicies`.
func (o AccessConfigurationOutput) PermissionPolicies() AccessConfigurationPermissionPolicyArrayOutput {
	return o.ApplyT(func(v *AccessConfiguration) AccessConfigurationPermissionPolicyArrayOutput {
		return v.PermissionPolicies
	}).(AccessConfigurationPermissionPolicyArrayOutput)
}

// The RelayState of the Access Configuration, Cloud SSO users use this access configuration to access the RD account, the initial access page address. Must be the Alibaba Cloud console page, the default is the console home page.
func (o AccessConfigurationOutput) RelayState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.StringPtrOutput { return v.RelayState }).(pulumi.StringPtrOutput)
}

// The SessionDuration of the Access Configuration. Valid Value: `900` to `43200`. Unit: Seconds.
func (o AccessConfigurationOutput) SessionDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *AccessConfiguration) pulumi.IntOutput { return v.SessionDuration }).(pulumi.IntOutput)
}

type AccessConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AccessConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessConfiguration)(nil)).Elem()
}

func (o AccessConfigurationArrayOutput) ToAccessConfigurationArrayOutput() AccessConfigurationArrayOutput {
	return o
}

func (o AccessConfigurationArrayOutput) ToAccessConfigurationArrayOutputWithContext(ctx context.Context) AccessConfigurationArrayOutput {
	return o
}

func (o AccessConfigurationArrayOutput) Index(i pulumi.IntInput) AccessConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessConfiguration {
		return vs[0].([]*AccessConfiguration)[vs[1].(int)]
	}).(AccessConfigurationOutput)
}

type AccessConfigurationMapOutput struct{ *pulumi.OutputState }

func (AccessConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessConfiguration)(nil)).Elem()
}

func (o AccessConfigurationMapOutput) ToAccessConfigurationMapOutput() AccessConfigurationMapOutput {
	return o
}

func (o AccessConfigurationMapOutput) ToAccessConfigurationMapOutputWithContext(ctx context.Context) AccessConfigurationMapOutput {
	return o
}

func (o AccessConfigurationMapOutput) MapIndex(k pulumi.StringInput) AccessConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessConfiguration {
		return vs[0].(map[string]*AccessConfiguration)[vs[1].(string)]
	}).(AccessConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationInput)(nil)).Elem(), &AccessConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationArrayInput)(nil)).Elem(), AccessConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessConfigurationMapInput)(nil)).Elem(), AccessConfigurationMap{})
	pulumi.RegisterOutputType(AccessConfigurationOutput{})
	pulumi.RegisterOutputType(AccessConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AccessConfigurationMapOutput{})
}
