// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yundun

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BastionHostInstance struct {
	pulumi.CustomResourceState

	AdAuthServers      BastionHostInstanceAdAuthServerArrayOutput   `pulumi:"adAuthServers"`
	Bandwidth          pulumi.StringOutput                          `pulumi:"bandwidth"`
	Description        pulumi.StringOutput                          `pulumi:"description"`
	EnablePublicAccess pulumi.BoolOutput                            `pulumi:"enablePublicAccess"`
	LdapAuthServers    BastionHostInstanceLdapAuthServerArrayOutput `pulumi:"ldapAuthServers"`
	LicenseCode        pulumi.StringOutput                          `pulumi:"licenseCode"`
	Period             pulumi.IntPtrOutput                          `pulumi:"period"`
	PlanCode           pulumi.StringOutput                          `pulumi:"planCode"`
	RenewPeriod        pulumi.IntPtrOutput                          `pulumi:"renewPeriod"`
	RenewalPeriodUnit  pulumi.StringOutput                          `pulumi:"renewalPeriodUnit"`
	RenewalStatus      pulumi.StringOutput                          `pulumi:"renewalStatus"`
	ResourceGroupId    pulumi.StringOutput                          `pulumi:"resourceGroupId"`
	SecurityGroupIds   pulumi.StringArrayOutput                     `pulumi:"securityGroupIds"`
	Storage            pulumi.StringOutput                          `pulumi:"storage"`
	Tags               pulumi.MapOutput                             `pulumi:"tags"`
	VswitchId          pulumi.StringOutput                          `pulumi:"vswitchId"`
}

// NewBastionHostInstance registers a new resource with the given unique name, arguments, and options.
func NewBastionHostInstance(ctx *pulumi.Context,
	name string, args *BastionHostInstanceArgs, opts ...pulumi.ResourceOption) (*BastionHostInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.LicenseCode == nil {
		return nil, errors.New("invalid value for required argument 'LicenseCode'")
	}
	if args.PlanCode == nil {
		return nil, errors.New("invalid value for required argument 'PlanCode'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource BastionHostInstance
	err := ctx.RegisterResource("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBastionHostInstance gets an existing BastionHostInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBastionHostInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostInstanceState, opts ...pulumi.ResourceOption) (*BastionHostInstance, error) {
	var resource BastionHostInstance
	err := ctx.ReadResource("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BastionHostInstance resources.
type bastionHostInstanceState struct {
	AdAuthServers      []BastionHostInstanceAdAuthServer   `pulumi:"adAuthServers"`
	Bandwidth          *string                             `pulumi:"bandwidth"`
	Description        *string                             `pulumi:"description"`
	EnablePublicAccess *bool                               `pulumi:"enablePublicAccess"`
	LdapAuthServers    []BastionHostInstanceLdapAuthServer `pulumi:"ldapAuthServers"`
	LicenseCode        *string                             `pulumi:"licenseCode"`
	Period             *int                                `pulumi:"period"`
	PlanCode           *string                             `pulumi:"planCode"`
	RenewPeriod        *int                                `pulumi:"renewPeriod"`
	RenewalPeriodUnit  *string                             `pulumi:"renewalPeriodUnit"`
	RenewalStatus      *string                             `pulumi:"renewalStatus"`
	ResourceGroupId    *string                             `pulumi:"resourceGroupId"`
	SecurityGroupIds   []string                            `pulumi:"securityGroupIds"`
	Storage            *string                             `pulumi:"storage"`
	Tags               map[string]interface{}              `pulumi:"tags"`
	VswitchId          *string                             `pulumi:"vswitchId"`
}

type BastionHostInstanceState struct {
	AdAuthServers      BastionHostInstanceAdAuthServerArrayInput
	Bandwidth          pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	EnablePublicAccess pulumi.BoolPtrInput
	LdapAuthServers    BastionHostInstanceLdapAuthServerArrayInput
	LicenseCode        pulumi.StringPtrInput
	Period             pulumi.IntPtrInput
	PlanCode           pulumi.StringPtrInput
	RenewPeriod        pulumi.IntPtrInput
	RenewalPeriodUnit  pulumi.StringPtrInput
	RenewalStatus      pulumi.StringPtrInput
	ResourceGroupId    pulumi.StringPtrInput
	SecurityGroupIds   pulumi.StringArrayInput
	Storage            pulumi.StringPtrInput
	Tags               pulumi.MapInput
	VswitchId          pulumi.StringPtrInput
}

func (BastionHostInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostInstanceState)(nil)).Elem()
}

type bastionHostInstanceArgs struct {
	AdAuthServers      []BastionHostInstanceAdAuthServer   `pulumi:"adAuthServers"`
	Bandwidth          string                              `pulumi:"bandwidth"`
	Description        string                              `pulumi:"description"`
	EnablePublicAccess *bool                               `pulumi:"enablePublicAccess"`
	LdapAuthServers    []BastionHostInstanceLdapAuthServer `pulumi:"ldapAuthServers"`
	LicenseCode        string                              `pulumi:"licenseCode"`
	Period             *int                                `pulumi:"period"`
	PlanCode           string                              `pulumi:"planCode"`
	RenewPeriod        *int                                `pulumi:"renewPeriod"`
	RenewalPeriodUnit  *string                             `pulumi:"renewalPeriodUnit"`
	RenewalStatus      *string                             `pulumi:"renewalStatus"`
	ResourceGroupId    *string                             `pulumi:"resourceGroupId"`
	SecurityGroupIds   []string                            `pulumi:"securityGroupIds"`
	Storage            string                              `pulumi:"storage"`
	Tags               map[string]interface{}              `pulumi:"tags"`
	VswitchId          string                              `pulumi:"vswitchId"`
}

// The set of arguments for constructing a BastionHostInstance resource.
type BastionHostInstanceArgs struct {
	AdAuthServers      BastionHostInstanceAdAuthServerArrayInput
	Bandwidth          pulumi.StringInput
	Description        pulumi.StringInput
	EnablePublicAccess pulumi.BoolPtrInput
	LdapAuthServers    BastionHostInstanceLdapAuthServerArrayInput
	LicenseCode        pulumi.StringInput
	Period             pulumi.IntPtrInput
	PlanCode           pulumi.StringInput
	RenewPeriod        pulumi.IntPtrInput
	RenewalPeriodUnit  pulumi.StringPtrInput
	RenewalStatus      pulumi.StringPtrInput
	ResourceGroupId    pulumi.StringPtrInput
	SecurityGroupIds   pulumi.StringArrayInput
	Storage            pulumi.StringInput
	Tags               pulumi.MapInput
	VswitchId          pulumi.StringInput
}

func (BastionHostInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostInstanceArgs)(nil)).Elem()
}

type BastionHostInstanceInput interface {
	pulumi.Input

	ToBastionHostInstanceOutput() BastionHostInstanceOutput
	ToBastionHostInstanceOutputWithContext(ctx context.Context) BastionHostInstanceOutput
}

func (*BastionHostInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHostInstance)(nil)).Elem()
}

func (i *BastionHostInstance) ToBastionHostInstanceOutput() BastionHostInstanceOutput {
	return i.ToBastionHostInstanceOutputWithContext(context.Background())
}

func (i *BastionHostInstance) ToBastionHostInstanceOutputWithContext(ctx context.Context) BastionHostInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostInstanceOutput)
}

// BastionHostInstanceArrayInput is an input type that accepts BastionHostInstanceArray and BastionHostInstanceArrayOutput values.
// You can construct a concrete instance of `BastionHostInstanceArrayInput` via:
//
//	BastionHostInstanceArray{ BastionHostInstanceArgs{...} }
type BastionHostInstanceArrayInput interface {
	pulumi.Input

	ToBastionHostInstanceArrayOutput() BastionHostInstanceArrayOutput
	ToBastionHostInstanceArrayOutputWithContext(context.Context) BastionHostInstanceArrayOutput
}

type BastionHostInstanceArray []BastionHostInstanceInput

func (BastionHostInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BastionHostInstance)(nil)).Elem()
}

func (i BastionHostInstanceArray) ToBastionHostInstanceArrayOutput() BastionHostInstanceArrayOutput {
	return i.ToBastionHostInstanceArrayOutputWithContext(context.Background())
}

func (i BastionHostInstanceArray) ToBastionHostInstanceArrayOutputWithContext(ctx context.Context) BastionHostInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostInstanceArrayOutput)
}

// BastionHostInstanceMapInput is an input type that accepts BastionHostInstanceMap and BastionHostInstanceMapOutput values.
// You can construct a concrete instance of `BastionHostInstanceMapInput` via:
//
//	BastionHostInstanceMap{ "key": BastionHostInstanceArgs{...} }
type BastionHostInstanceMapInput interface {
	pulumi.Input

	ToBastionHostInstanceMapOutput() BastionHostInstanceMapOutput
	ToBastionHostInstanceMapOutputWithContext(context.Context) BastionHostInstanceMapOutput
}

type BastionHostInstanceMap map[string]BastionHostInstanceInput

func (BastionHostInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BastionHostInstance)(nil)).Elem()
}

func (i BastionHostInstanceMap) ToBastionHostInstanceMapOutput() BastionHostInstanceMapOutput {
	return i.ToBastionHostInstanceMapOutputWithContext(context.Background())
}

func (i BastionHostInstanceMap) ToBastionHostInstanceMapOutputWithContext(ctx context.Context) BastionHostInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostInstanceMapOutput)
}

type BastionHostInstanceOutput struct{ *pulumi.OutputState }

func (BastionHostInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHostInstance)(nil)).Elem()
}

func (o BastionHostInstanceOutput) ToBastionHostInstanceOutput() BastionHostInstanceOutput {
	return o
}

func (o BastionHostInstanceOutput) ToBastionHostInstanceOutputWithContext(ctx context.Context) BastionHostInstanceOutput {
	return o
}

func (o BastionHostInstanceOutput) AdAuthServers() BastionHostInstanceAdAuthServerArrayOutput {
	return o.ApplyT(func(v *BastionHostInstance) BastionHostInstanceAdAuthServerArrayOutput { return v.AdAuthServers }).(BastionHostInstanceAdAuthServerArrayOutput)
}

func (o BastionHostInstanceOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.Bandwidth }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) EnablePublicAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.BoolOutput { return v.EnablePublicAccess }).(pulumi.BoolOutput)
}

func (o BastionHostInstanceOutput) LdapAuthServers() BastionHostInstanceLdapAuthServerArrayOutput {
	return o.ApplyT(func(v *BastionHostInstance) BastionHostInstanceLdapAuthServerArrayOutput { return v.LdapAuthServers }).(BastionHostInstanceLdapAuthServerArrayOutput)
}

func (o BastionHostInstanceOutput) LicenseCode() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.LicenseCode }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

func (o BastionHostInstanceOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.PlanCode }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) RenewPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.IntPtrOutput { return v.RenewPeriod }).(pulumi.IntPtrOutput)
}

func (o BastionHostInstanceOutput) RenewalPeriodUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.RenewalPeriodUnit }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) RenewalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.RenewalStatus }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o BastionHostInstanceOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.Storage }).(pulumi.StringOutput)
}

func (o BastionHostInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o BastionHostInstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHostInstance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type BastionHostInstanceArrayOutput struct{ *pulumi.OutputState }

func (BastionHostInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BastionHostInstance)(nil)).Elem()
}

func (o BastionHostInstanceArrayOutput) ToBastionHostInstanceArrayOutput() BastionHostInstanceArrayOutput {
	return o
}

func (o BastionHostInstanceArrayOutput) ToBastionHostInstanceArrayOutputWithContext(ctx context.Context) BastionHostInstanceArrayOutput {
	return o
}

func (o BastionHostInstanceArrayOutput) Index(i pulumi.IntInput) BastionHostInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BastionHostInstance {
		return vs[0].([]*BastionHostInstance)[vs[1].(int)]
	}).(BastionHostInstanceOutput)
}

type BastionHostInstanceMapOutput struct{ *pulumi.OutputState }

func (BastionHostInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BastionHostInstance)(nil)).Elem()
}

func (o BastionHostInstanceMapOutput) ToBastionHostInstanceMapOutput() BastionHostInstanceMapOutput {
	return o
}

func (o BastionHostInstanceMapOutput) ToBastionHostInstanceMapOutputWithContext(ctx context.Context) BastionHostInstanceMapOutput {
	return o
}

func (o BastionHostInstanceMapOutput) MapIndex(k pulumi.StringInput) BastionHostInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BastionHostInstance {
		return vs[0].(map[string]*BastionHostInstance)[vs[1].(string)]
	}).(BastionHostInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostInstanceInput)(nil)).Elem(), &BastionHostInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostInstanceArrayInput)(nil)).Elem(), BastionHostInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BastionHostInstanceMapInput)(nil)).Elem(), BastionHostInstanceMap{})
	pulumi.RegisterOutputType(BastionHostInstanceOutput{})
	pulumi.RegisterOutputType(BastionHostInstanceArrayOutput{})
	pulumi.RegisterOutputType(BastionHostInstanceMapOutput{})
}
