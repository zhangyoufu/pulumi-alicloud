// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Yundun_bastionhost instance can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:bastionhost/instance:Instance example bastionhost-exampe123456
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The AD auth server of the Instance. See the following `Block adAuthServer`.
	AdAuthServers InstanceAdAuthServerArrayOutput `pulumi:"adAuthServers"`
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
	EnablePublicAccess pulumi.BoolOutput `pulumi:"enablePublicAccess"`
	// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
	LdapAuthServers InstanceLdapAuthServerArrayOutput `pulumi:"ldapAuthServers"`
	// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
	LicenseCode pulumi.StringOutput `pulumi:"licenseCode"`
	Period      pulumi.IntPtrOutput `pulumi:"period"`
	// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
	ResourceGroupId  pulumi.StringPtrOutput   `pulumi:"resourceGroupId"`
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// VSwitch ID configured to Bastionhost.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.LicenseCode == nil {
		return nil, errors.New("invalid value for required argument 'LicenseCode'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:bastionhost/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:bastionhost/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The AD auth server of the Instance. See the following `Block adAuthServer`.
	AdAuthServers []InstanceAdAuthServer `pulumi:"adAuthServers"`
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description *string `pulumi:"description"`
	// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
	EnablePublicAccess *bool `pulumi:"enablePublicAccess"`
	// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
	LdapAuthServers []InstanceLdapAuthServer `pulumi:"ldapAuthServers"`
	// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
	LicenseCode *string `pulumi:"licenseCode"`
	Period      *int    `pulumi:"period"`
	// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
	ResourceGroupId  *string  `pulumi:"resourceGroupId"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// VSwitch ID configured to Bastionhost.
	VswitchId *string `pulumi:"vswitchId"`
}

type InstanceState struct {
	// The AD auth server of the Instance. See the following `Block adAuthServer`.
	AdAuthServers InstanceAdAuthServerArrayInput
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringPtrInput
	// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
	EnablePublicAccess pulumi.BoolPtrInput
	// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
	LdapAuthServers InstanceLdapAuthServerArrayInput
	// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
	LicenseCode pulumi.StringPtrInput
	Period      pulumi.IntPtrInput
	// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
	ResourceGroupId  pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// VSwitch ID configured to Bastionhost.
	VswitchId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The AD auth server of the Instance. See the following `Block adAuthServer`.
	AdAuthServers []InstanceAdAuthServer `pulumi:"adAuthServers"`
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description string `pulumi:"description"`
	// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
	EnablePublicAccess *bool `pulumi:"enablePublicAccess"`
	// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
	LdapAuthServers []InstanceLdapAuthServer `pulumi:"ldapAuthServers"`
	// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
	LicenseCode string `pulumi:"licenseCode"`
	Period      *int   `pulumi:"period"`
	// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
	ResourceGroupId  *string  `pulumi:"resourceGroupId"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// VSwitch ID configured to Bastionhost.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The AD auth server of the Instance. See the following `Block adAuthServer`.
	AdAuthServers InstanceAdAuthServerArrayInput
	// Description of the instance. This name can have a string of 1 to 63 characters.
	Description pulumi.StringInput
	// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
	EnablePublicAccess pulumi.BoolPtrInput
	// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
	LdapAuthServers InstanceLdapAuthServerArrayInput
	// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
	LicenseCode pulumi.StringInput
	Period      pulumi.IntPtrInput
	// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
	ResourceGroupId  pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// VSwitch ID configured to Bastionhost.
	VswitchId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// The AD auth server of the Instance. See the following `Block adAuthServer`.
func (o InstanceOutput) AdAuthServers() InstanceAdAuthServerArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceAdAuthServerArrayOutput { return v.AdAuthServers }).(InstanceAdAuthServerArrayOutput)
}

// Description of the instance. This name can have a string of 1 to 63 characters.
func (o InstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
func (o InstanceOutput) EnablePublicAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnablePublicAccess }).(pulumi.BoolOutput)
}

// The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
func (o InstanceOutput) LdapAuthServers() InstanceLdapAuthServerArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceLdapAuthServerArrayOutput { return v.LdapAuthServers }).(InstanceLdapAuthServerArrayOutput)
}

// The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
func (o InstanceOutput) LicenseCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.LicenseCode }).(pulumi.StringOutput)
}

func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
func (o InstanceOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o InstanceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A mapping of tags to assign to the resource.
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// VSwitch ID configured to Bastionhost.
func (o InstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
