// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bastionhost

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Bastion Host User resource.
//
// For information about Bastion Host User and how to use it, see [What is User](https://www.alibabacloud.com/help/en/bastion-host/latest/api-yundun-bastionhost-2019-12-09-createuser).
//
// > **NOTE:** Available since v1.133.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bastionhost"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("^default-NODELETING$"),
//				CidrBlock: pulumi.StringRef("10.4.0.0/16"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSwitches, err := vpc.GetSwitches(ctx, &vpc.GetSwitchesArgs{
//				CidrBlock: pulumi.StringRef("10.4.0.0/24"),
//				VpcId:     pulumi.StringRef(defaultNetworks.Ids[0]),
//				ZoneId:    pulumi.StringRef(defaultZones.Zones[0].Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: *pulumi.String(defaultNetworks.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := bastionhost.NewInstance(ctx, "defaultInstance", &bastionhost.InstanceArgs{
//				Description: pulumi.String(name),
//				LicenseCode: pulumi.String("bhah_ent_50_asset"),
//				PlanCode:    pulumi.String("cloudbastion"),
//				Storage:     pulumi.String("5"),
//				Bandwidth:   pulumi.String("5"),
//				Period:      pulumi.Int(1),
//				VswitchId:   *pulumi.String(defaultSwitches.Ids[0]),
//				SecurityGroupIds: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bastionhost.NewUser(ctx, "localUser", &bastionhost.UserArgs{
//				InstanceId:        defaultInstance.ID(),
//				MobileCountryCode: pulumi.String("CN"),
//				Mobile:            pulumi.String("13312345678"),
//				Password:          pulumi.String("YourPassword-123"),
//				Source:            pulumi.String("Local"),
//				UserName:          pulumi.String(fmt.Sprintf("%v_local_user", name)),
//			})
//			if err != nil {
//				return err
//			}
//			user, err := ram.NewUser(ctx, "user", &ram.UserArgs{
//				DisplayName: pulumi.String(fmt.Sprintf("%v_bastionhost_user", name)),
//				Mobile:      pulumi.String("86-18688888888"),
//				Email:       pulumi.String("hello.uuu@aaa.com"),
//				Comments:    pulumi.String("yoyoyo"),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccount, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = bastionhost.NewUser(ctx, "ramUser", &bastionhost.UserArgs{
//				InstanceId:   defaultInstance.ID(),
//				Source:       pulumi.String("Ram"),
//				SourceUserId: *pulumi.String(defaultAccount.Id),
//				UserName:     user.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Bastion Host User can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:bastionhost/user:User example <instance_id>:<user_id>
// ```
type User struct {
	pulumi.CustomResourceState

	// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specify the New User's Mailbox.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// You Want to Query the User the Bastion Host ID of.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile pulumi.StringPtrOutput `pulumi:"mobile"`
	// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
	// * CN: Mainland China (+86)
	// * HK: hong Kong, China (+852)
	// * MO: Macau, China (+853)
	// * TW: Taiwan, China (+886)
	// * RU: Russian (+7)
	// * SG: Singapore (+65)
	// * MY: malaysia (+60)
	// * ID: Indonesia (+62)
	// * DE: Germany (+49)
	// * AU: Australia (+61)
	// * US: United States (+1)
	// * AE: dubai (+971)
	// * JP: Japan (+81) Introducing the Long-Range
	// * GB: United Kingdom (+44)
	// * IN: India (+91)
	// * KR: South Korea (+82)
	// * PH: philippines (+63)
	// * CH: Switzerland (+41)
	// * SE: Sweden (+46)
	MobileCountryCode pulumi.StringOutput `pulumi:"mobileCountryCode"`
	// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Specify the New of the User That Created the Source. Valid Values:
	// * Local: Local User
	// * Ram: Ram User
	// * AD: AD-authenticated User
	// * LDAP: LDAP-authenticated User
	// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
	Source pulumi.StringOutput `pulumi:"source"`
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId pulumi.StringPtrOutput `pulumi:"sourceUserId"`
	// The status of the resource. Valid values: `Frozen`, `Normal`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The User ID.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
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
	err := ctx.RegisterResource("alicloud:bastionhost/user:User", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:bastionhost/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
	Comment *string `pulumi:"comment"`
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName *string `pulumi:"displayName"`
	// Specify the New User's Mailbox.
	Email *string `pulumi:"email"`
	// You Want to Query the User the Bastion Host ID of.
	InstanceId *string `pulumi:"instanceId"`
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile *string `pulumi:"mobile"`
	// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
	// * CN: Mainland China (+86)
	// * HK: hong Kong, China (+852)
	// * MO: Macau, China (+853)
	// * TW: Taiwan, China (+886)
	// * RU: Russian (+7)
	// * SG: Singapore (+65)
	// * MY: malaysia (+60)
	// * ID: Indonesia (+62)
	// * DE: Germany (+49)
	// * AU: Australia (+61)
	// * US: United States (+1)
	// * AE: dubai (+971)
	// * JP: Japan (+81) Introducing the Long-Range
	// * GB: United Kingdom (+44)
	// * IN: India (+91)
	// * KR: South Korea (+82)
	// * PH: philippines (+63)
	// * CH: Switzerland (+41)
	// * SE: Sweden (+46)
	MobileCountryCode *string `pulumi:"mobileCountryCode"`
	// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
	Password *string `pulumi:"password"`
	// Specify the New of the User That Created the Source. Valid Values:
	// * Local: Local User
	// * Ram: Ram User
	// * AD: AD-authenticated User
	// * LDAP: LDAP-authenticated User
	// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
	Source *string `pulumi:"source"`
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId *string `pulumi:"sourceUserId"`
	// The status of the resource. Valid values: `Frozen`, `Normal`.
	Status *string `pulumi:"status"`
	// The User ID.
	UserId *string `pulumi:"userId"`
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName *string `pulumi:"userName"`
}

type UserState struct {
	// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
	Comment pulumi.StringPtrInput
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName pulumi.StringPtrInput
	// Specify the New User's Mailbox.
	Email pulumi.StringPtrInput
	// You Want to Query the User the Bastion Host ID of.
	InstanceId pulumi.StringPtrInput
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile pulumi.StringPtrInput
	// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
	// * CN: Mainland China (+86)
	// * HK: hong Kong, China (+852)
	// * MO: Macau, China (+853)
	// * TW: Taiwan, China (+886)
	// * RU: Russian (+7)
	// * SG: Singapore (+65)
	// * MY: malaysia (+60)
	// * ID: Indonesia (+62)
	// * DE: Germany (+49)
	// * AU: Australia (+61)
	// * US: United States (+1)
	// * AE: dubai (+971)
	// * JP: Japan (+81) Introducing the Long-Range
	// * GB: United Kingdom (+44)
	// * IN: India (+91)
	// * KR: South Korea (+82)
	// * PH: philippines (+63)
	// * CH: Switzerland (+41)
	// * SE: Sweden (+46)
	MobileCountryCode pulumi.StringPtrInput
	// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
	Password pulumi.StringPtrInput
	// Specify the New of the User That Created the Source. Valid Values:
	// * Local: Local User
	// * Ram: Ram User
	// * AD: AD-authenticated User
	// * LDAP: LDAP-authenticated User
	// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
	Source pulumi.StringPtrInput
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId pulumi.StringPtrInput
	// The status of the resource. Valid values: `Frozen`, `Normal`.
	Status pulumi.StringPtrInput
	// The User ID.
	UserId pulumi.StringPtrInput
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
	Comment *string `pulumi:"comment"`
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName *string `pulumi:"displayName"`
	// Specify the New User's Mailbox.
	Email *string `pulumi:"email"`
	// You Want to Query the User the Bastion Host ID of.
	InstanceId string `pulumi:"instanceId"`
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile *string `pulumi:"mobile"`
	// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
	// * CN: Mainland China (+86)
	// * HK: hong Kong, China (+852)
	// * MO: Macau, China (+853)
	// * TW: Taiwan, China (+886)
	// * RU: Russian (+7)
	// * SG: Singapore (+65)
	// * MY: malaysia (+60)
	// * ID: Indonesia (+62)
	// * DE: Germany (+49)
	// * AU: Australia (+61)
	// * US: United States (+1)
	// * AE: dubai (+971)
	// * JP: Japan (+81) Introducing the Long-Range
	// * GB: United Kingdom (+44)
	// * IN: India (+91)
	// * KR: South Korea (+82)
	// * PH: philippines (+63)
	// * CH: Switzerland (+41)
	// * SE: Sweden (+46)
	MobileCountryCode *string `pulumi:"mobileCountryCode"`
	// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
	Password *string `pulumi:"password"`
	// Specify the New of the User That Created the Source. Valid Values:
	// * Local: Local User
	// * Ram: Ram User
	// * AD: AD-authenticated User
	// * LDAP: LDAP-authenticated User
	// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
	Source string `pulumi:"source"`
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId *string `pulumi:"sourceUserId"`
	// The status of the resource. Valid values: `Frozen`, `Normal`.
	Status *string `pulumi:"status"`
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
	Comment pulumi.StringPtrInput
	// Specify the New Created the User's Display Name. Supports up to 128 Characters.
	DisplayName pulumi.StringPtrInput
	// Specify the New User's Mailbox.
	Email pulumi.StringPtrInput
	// You Want to Query the User the Bastion Host ID of.
	InstanceId pulumi.StringInput
	// Specify the New of the User That Created a Different Mobile Phone Number from Your.
	Mobile pulumi.StringPtrInput
	// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
	// * CN: Mainland China (+86)
	// * HK: hong Kong, China (+852)
	// * MO: Macau, China (+853)
	// * TW: Taiwan, China (+886)
	// * RU: Russian (+7)
	// * SG: Singapore (+65)
	// * MY: malaysia (+60)
	// * ID: Indonesia (+62)
	// * DE: Germany (+49)
	// * AU: Australia (+61)
	// * US: United States (+1)
	// * AE: dubai (+971)
	// * JP: Japan (+81) Introducing the Long-Range
	// * GB: United Kingdom (+44)
	// * IN: India (+91)
	// * KR: South Korea (+82)
	// * PH: philippines (+63)
	// * CH: Switzerland (+41)
	// * SE: Sweden (+46)
	MobileCountryCode pulumi.StringPtrInput
	// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
	Password pulumi.StringPtrInput
	// Specify the New of the User That Created the Source. Valid Values:
	// * Local: Local User
	// * Ram: Ram User
	// * AD: AD-authenticated User
	// * LDAP: LDAP-authenticated User
	// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
	Source pulumi.StringInput
	// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
	SourceUserId pulumi.StringPtrInput
	// The status of the resource. Valid values: `Frozen`, `Normal`.
	Status pulumi.StringPtrInput
	// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
	UserName pulumi.StringInput
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

// Specify the New of the User That Created the Remark Information. Supports up to 500 Characters.
func (o UserOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Specify the New Created the User's Display Name. Supports up to 128 Characters.
func (o UserOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Specify the New User's Mailbox.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// You Want to Query the User the Bastion Host ID of.
func (o UserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specify the New of the User That Created a Different Mobile Phone Number from Your.
func (o UserOutput) Mobile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Mobile }).(pulumi.StringPtrOutput)
}

// Specify the New Create User Mobile Phone Number of the International Domain Name. The Default Value Is the CN. Valid Values:
// * CN: Mainland China (+86)
// * HK: hong Kong, China (+852)
// * MO: Macau, China (+853)
// * TW: Taiwan, China (+886)
// * RU: Russian (+7)
// * SG: Singapore (+65)
// * MY: malaysia (+60)
// * ID: Indonesia (+62)
// * DE: Germany (+49)
// * AU: Australia (+61)
// * US: United States (+1)
// * AE: dubai (+971)
// * JP: Japan (+81) Introducing the Long-Range
// * GB: United Kingdom (+44)
// * IN: India (+91)
// * KR: South Korea (+82)
// * PH: philippines (+63)
// * CH: Switzerland (+41)
// * SE: Sweden (+46)
func (o UserOutput) MobileCountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.MobileCountryCode }).(pulumi.StringOutput)
}

// Specify the New User's Password. Supports up to 128 Characters. Description of the New User as the Source of the Local User That Is, Source Value for Local, this Parameter Is Required.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Specify the New of the User That Created the Source. Valid Values:
// * Local: Local User
// * Ram: Ram User
// * AD: AD-authenticated User
// * LDAP: LDAP-authenticated User
// > **NOTE:** From version 1.199.0, `source` can be set to `AD`, `LDAP`.
func (o UserOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Specify the Newly Created User Is Uniquely Identified. Indicates That the Parameter Is a Bastion Host Corresponding to the User with the Ram User's Unique Identifier. The Newly Created User Source Grant Permission to a RAM User (That Is, Source Used as a Ram), this Parameter Is Required. You Can Call Access Control of Listusers Interface from the Return Data Userid to Obtain the Parameters.
func (o UserOutput) SourceUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.SourceUserId }).(pulumi.StringPtrOutput)
}

// The status of the resource. Valid values: `Frozen`, `Normal`.
func (o UserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The User ID.
func (o UserOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// Specify the New User Name. This Parameter Is Only by Letters, Lowercase Letters, Numbers, and Underscores (_), Supports up to 128 Characters.
func (o UserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
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
