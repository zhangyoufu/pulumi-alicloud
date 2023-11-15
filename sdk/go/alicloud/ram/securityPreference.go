// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM Security Preference resource.
//
// For information about RAM Security Preference and how to use it, see [What is Security Preference](https://www.alibabacloud.com/help/en/doc-detail/186694.htm).
//
// > **NOTE:** Available since v1.152.0.
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
//			_, err := ram.NewSecurityPreference(ctx, "example", &ram.SecurityPreferenceArgs{
//				AllowUserToChangePassword: pulumi.Bool(true),
//				EnableSaveMfaTicket:       pulumi.Bool(false),
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
// RAM Security Preference can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/securityPreference:SecurityPreference example <id>
//
// ```
type SecurityPreference struct {
	pulumi.CustomResourceState

	// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
	AllowUserToChangePassword pulumi.BoolOutput `pulumi:"allowUserToChangePassword"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
	AllowUserToManageAccessKeys pulumi.BoolOutput `pulumi:"allowUserToManageAccessKeys"`
	// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
	AllowUserToManageMfaDevices pulumi.BoolOutput `pulumi:"allowUserToManageMfaDevices"`
	// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
	EnableSaveMfaTicket pulumi.BoolOutput `pulumi:"enableSaveMfaTicket"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
	EnforceMfaForLogin pulumi.BoolOutput `pulumi:"enforceMfaForLogin"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	LoginNetworkMasks pulumi.StringPtrOutput `pulumi:"loginNetworkMasks"`
	// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
	LoginSessionDuration pulumi.IntOutput `pulumi:"loginSessionDuration"`
}

// NewSecurityPreference registers a new resource with the given unique name, arguments, and options.
func NewSecurityPreference(ctx *pulumi.Context,
	name string, args *SecurityPreferenceArgs, opts ...pulumi.ResourceOption) (*SecurityPreference, error) {
	if args == nil {
		args = &SecurityPreferenceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityPreference
	err := ctx.RegisterResource("alicloud:ram/securityPreference:SecurityPreference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPreference gets an existing SecurityPreference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPreference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPreferenceState, opts ...pulumi.ResourceOption) (*SecurityPreference, error) {
	var resource SecurityPreference
	err := ctx.ReadResource("alicloud:ram/securityPreference:SecurityPreference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPreference resources.
type securityPreferenceState struct {
	// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
	AllowUserToChangePassword *bool `pulumi:"allowUserToChangePassword"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
	AllowUserToManageAccessKeys *bool `pulumi:"allowUserToManageAccessKeys"`
	// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
	AllowUserToManageMfaDevices *bool `pulumi:"allowUserToManageMfaDevices"`
	// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
	EnableSaveMfaTicket *bool `pulumi:"enableSaveMfaTicket"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
	EnforceMfaForLogin *bool `pulumi:"enforceMfaForLogin"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	LoginNetworkMasks *string `pulumi:"loginNetworkMasks"`
	// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
	LoginSessionDuration *int `pulumi:"loginSessionDuration"`
}

type SecurityPreferenceState struct {
	// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
	AllowUserToChangePassword pulumi.BoolPtrInput
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
	AllowUserToManageAccessKeys pulumi.BoolPtrInput
	// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
	AllowUserToManageMfaDevices pulumi.BoolPtrInput
	// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
	EnableSaveMfaTicket pulumi.BoolPtrInput
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
	EnforceMfaForLogin pulumi.BoolPtrInput
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	LoginNetworkMasks pulumi.StringPtrInput
	// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
	LoginSessionDuration pulumi.IntPtrInput
}

func (SecurityPreferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPreferenceState)(nil)).Elem()
}

type securityPreferenceArgs struct {
	// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
	AllowUserToChangePassword *bool `pulumi:"allowUserToChangePassword"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
	AllowUserToManageAccessKeys *bool `pulumi:"allowUserToManageAccessKeys"`
	// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
	AllowUserToManageMfaDevices *bool `pulumi:"allowUserToManageMfaDevices"`
	// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
	EnableSaveMfaTicket *bool `pulumi:"enableSaveMfaTicket"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
	EnforceMfaForLogin *bool `pulumi:"enforceMfaForLogin"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	LoginNetworkMasks *string `pulumi:"loginNetworkMasks"`
	// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
	LoginSessionDuration *int `pulumi:"loginSessionDuration"`
}

// The set of arguments for constructing a SecurityPreference resource.
type SecurityPreferenceArgs struct {
	// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
	AllowUserToChangePassword pulumi.BoolPtrInput
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
	AllowUserToManageAccessKeys pulumi.BoolPtrInput
	// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
	AllowUserToManageMfaDevices pulumi.BoolPtrInput
	// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
	EnableSaveMfaTicket pulumi.BoolPtrInput
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
	EnforceMfaForLogin pulumi.BoolPtrInput
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	LoginNetworkMasks pulumi.StringPtrInput
	// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
	LoginSessionDuration pulumi.IntPtrInput
}

func (SecurityPreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPreferenceArgs)(nil)).Elem()
}

type SecurityPreferenceInput interface {
	pulumi.Input

	ToSecurityPreferenceOutput() SecurityPreferenceOutput
	ToSecurityPreferenceOutputWithContext(ctx context.Context) SecurityPreferenceOutput
}

func (*SecurityPreference) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPreference)(nil)).Elem()
}

func (i *SecurityPreference) ToSecurityPreferenceOutput() SecurityPreferenceOutput {
	return i.ToSecurityPreferenceOutputWithContext(context.Background())
}

func (i *SecurityPreference) ToSecurityPreferenceOutputWithContext(ctx context.Context) SecurityPreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPreferenceOutput)
}

// SecurityPreferenceArrayInput is an input type that accepts SecurityPreferenceArray and SecurityPreferenceArrayOutput values.
// You can construct a concrete instance of `SecurityPreferenceArrayInput` via:
//
//	SecurityPreferenceArray{ SecurityPreferenceArgs{...} }
type SecurityPreferenceArrayInput interface {
	pulumi.Input

	ToSecurityPreferenceArrayOutput() SecurityPreferenceArrayOutput
	ToSecurityPreferenceArrayOutputWithContext(context.Context) SecurityPreferenceArrayOutput
}

type SecurityPreferenceArray []SecurityPreferenceInput

func (SecurityPreferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityPreference)(nil)).Elem()
}

func (i SecurityPreferenceArray) ToSecurityPreferenceArrayOutput() SecurityPreferenceArrayOutput {
	return i.ToSecurityPreferenceArrayOutputWithContext(context.Background())
}

func (i SecurityPreferenceArray) ToSecurityPreferenceArrayOutputWithContext(ctx context.Context) SecurityPreferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPreferenceArrayOutput)
}

// SecurityPreferenceMapInput is an input type that accepts SecurityPreferenceMap and SecurityPreferenceMapOutput values.
// You can construct a concrete instance of `SecurityPreferenceMapInput` via:
//
//	SecurityPreferenceMap{ "key": SecurityPreferenceArgs{...} }
type SecurityPreferenceMapInput interface {
	pulumi.Input

	ToSecurityPreferenceMapOutput() SecurityPreferenceMapOutput
	ToSecurityPreferenceMapOutputWithContext(context.Context) SecurityPreferenceMapOutput
}

type SecurityPreferenceMap map[string]SecurityPreferenceInput

func (SecurityPreferenceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityPreference)(nil)).Elem()
}

func (i SecurityPreferenceMap) ToSecurityPreferenceMapOutput() SecurityPreferenceMapOutput {
	return i.ToSecurityPreferenceMapOutputWithContext(context.Background())
}

func (i SecurityPreferenceMap) ToSecurityPreferenceMapOutputWithContext(ctx context.Context) SecurityPreferenceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPreferenceMapOutput)
}

type SecurityPreferenceOutput struct{ *pulumi.OutputState }

func (SecurityPreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPreference)(nil)).Elem()
}

func (o SecurityPreferenceOutput) ToSecurityPreferenceOutput() SecurityPreferenceOutput {
	return o
}

func (o SecurityPreferenceOutput) ToSecurityPreferenceOutputWithContext(ctx context.Context) SecurityPreferenceOutput {
	return o
}

// Specifies whether RAM users can change their passwords. Valid values: `true` and `false`
func (o SecurityPreferenceOutput) AllowUserToChangePassword() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.BoolOutput { return v.AllowUserToChangePassword }).(pulumi.BoolOutput)
}

// Specifies whether RAM users can manage their AccessKey pairs. Valid values: `true` and `false`
func (o SecurityPreferenceOutput) AllowUserToManageAccessKeys() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.BoolOutput { return v.AllowUserToManageAccessKeys }).(pulumi.BoolOutput)
}

// Specifies whether RAM users can manage their MFA devices. Valid values: `true` and `false`
func (o SecurityPreferenceOutput) AllowUserToManageMfaDevices() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.BoolOutput { return v.AllowUserToManageMfaDevices }).(pulumi.BoolOutput)
}

// Specifies whether to remember the MFA devices for seven days. Valid values: `true` and `false`
func (o SecurityPreferenceOutput) EnableSaveMfaTicket() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.BoolOutput { return v.EnableSaveMfaTicket }).(pulumi.BoolOutput)
}

// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console by using usernames and passwords. Valid values: `true` and `false`
func (o SecurityPreferenceOutput) EnforceMfaForLogin() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.BoolOutput { return v.EnforceMfaForLogin }).(pulumi.BoolOutput)
}

// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). However, this parameter does not take effect on API calls that are authenticated by using AccessKey pairs.**NOTE:** You can specify up to 25 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
// * If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
// * If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
// * If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
func (o SecurityPreferenceOutput) LoginNetworkMasks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.StringPtrOutput { return v.LoginNetworkMasks }).(pulumi.StringPtrOutput)
}

// The validity period of the logon session of RAM users. Valid values: 6 to 24. Unit: hours. Default value: 6.
func (o SecurityPreferenceOutput) LoginSessionDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityPreference) pulumi.IntOutput { return v.LoginSessionDuration }).(pulumi.IntOutput)
}

type SecurityPreferenceArrayOutput struct{ *pulumi.OutputState }

func (SecurityPreferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityPreference)(nil)).Elem()
}

func (o SecurityPreferenceArrayOutput) ToSecurityPreferenceArrayOutput() SecurityPreferenceArrayOutput {
	return o
}

func (o SecurityPreferenceArrayOutput) ToSecurityPreferenceArrayOutputWithContext(ctx context.Context) SecurityPreferenceArrayOutput {
	return o
}

func (o SecurityPreferenceArrayOutput) Index(i pulumi.IntInput) SecurityPreferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityPreference {
		return vs[0].([]*SecurityPreference)[vs[1].(int)]
	}).(SecurityPreferenceOutput)
}

type SecurityPreferenceMapOutput struct{ *pulumi.OutputState }

func (SecurityPreferenceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityPreference)(nil)).Elem()
}

func (o SecurityPreferenceMapOutput) ToSecurityPreferenceMapOutput() SecurityPreferenceMapOutput {
	return o
}

func (o SecurityPreferenceMapOutput) ToSecurityPreferenceMapOutputWithContext(ctx context.Context) SecurityPreferenceMapOutput {
	return o
}

func (o SecurityPreferenceMapOutput) MapIndex(k pulumi.StringInput) SecurityPreferenceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityPreference {
		return vs[0].(map[string]*SecurityPreference)[vs[1].(string)]
	}).(SecurityPreferenceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPreferenceInput)(nil)).Elem(), &SecurityPreference{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPreferenceArrayInput)(nil)).Elem(), SecurityPreferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPreferenceMapInput)(nil)).Elem(), SecurityPreferenceMap{})
	pulumi.RegisterOutputType(SecurityPreferenceOutput{})
	pulumi.RegisterOutputType(SecurityPreferenceArrayOutput{})
	pulumi.RegisterOutputType(SecurityPreferenceMapOutput{})
}
