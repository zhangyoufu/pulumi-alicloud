// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kvstore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a KVStore Account resource.
//
// For information about KVStore Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/doc-detail/95973.htm).
//
// > **NOTE:** Available since v1.66.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kvstore"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := kvstore.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := kvstore.NewInstance(ctx, "default", &kvstore.InstanceArgs{
//				DbInstanceName:  pulumi.String(name),
//				VswitchId:       defaultSwitch.ID(),
//				ResourceGroupId: pulumi.String(defaultGetResourceGroups.Ids[0]),
//				ZoneId:          pulumi.String(_default.Zones[0].Id),
//				InstanceClass:   pulumi.String("redis.master.large.default"),
//				InstanceType:    pulumi.String("Redis"),
//				EngineVersion:   pulumi.String("5.0"),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("10.23.12.24"),
//				},
//				Config: pulumi.Map{
//					"appendonly":             pulumi.Any("yes"),
//					"lazyfree-lazy-eviction": pulumi.Any("yes"),
//				},
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kvstore.NewAccount(ctx, "default", &kvstore.AccountArgs{
//				AccountName:     pulumi.String("tfexamplename"),
//				AccountPassword: pulumi.String("YourPassword_123"),
//				InstanceId:      defaultInstance.ID(),
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
// KVStore account can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:kvstore/account:Account example <instance_id>:<account_name>
// ```
type Account struct {
	pulumi.CustomResourceState

	// The name of the account. The name must meet the following requirements:
	// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
	// * The name can be up to 100 characters in length.
	// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringPtrOutput `pulumi:"accountPassword"`
	// The privilege of account access database. Default value: `RoleReadWrite`
	AccountPrivilege pulumi.StringPtrOutput `pulumi:"accountPrivilege"`
	// Privilege type of account.
	// - Normal: Common privilege.
	//   Default to Normal.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The status of KVStore Account.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("alicloud:kvstore/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("alicloud:kvstore/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The name of the account. The name must meet the following requirements:
	// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
	// * The name can be up to 100 characters in length.
	// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
	AccountName *string `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
	AccountPassword *string `pulumi:"accountPassword"`
	// The privilege of account access database. Default value: `RoleReadWrite`
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// Privilege type of account.
	// - Normal: Common privilege.
	//   Default to Normal.
	AccountType *string `pulumi:"accountType"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
	InstanceId *string `pulumi:"instanceId"`
	// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The status of KVStore Account.
	Status *string `pulumi:"status"`
}

type AccountState struct {
	// The name of the account. The name must meet the following requirements:
	// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
	// * The name can be up to 100 characters in length.
	// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
	AccountName pulumi.StringPtrInput
	// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringPtrInput
	// The privilege of account access database. Default value: `RoleReadWrite`
	AccountPrivilege pulumi.StringPtrInput
	// Privilege type of account.
	// - Normal: Common privilege.
	//   Default to Normal.
	AccountType pulumi.StringPtrInput
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
	InstanceId pulumi.StringPtrInput
	// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The status of KVStore Account.
	Status pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the account. The name must meet the following requirements:
	// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
	// * The name can be up to 100 characters in length.
	// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
	AccountName string `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
	AccountPassword *string `pulumi:"accountPassword"`
	// The privilege of account access database. Default value: `RoleReadWrite`
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// Privilege type of account.
	// - Normal: Common privilege.
	//   Default to Normal.
	AccountType *string `pulumi:"accountType"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
	InstanceId string `pulumi:"instanceId"`
	// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the account. The name must meet the following requirements:
	// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
	// * The name can be up to 100 characters in length.
	// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
	AccountName pulumi.StringInput
	// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringPtrInput
	// The privilege of account access database. Default value: `RoleReadWrite`
	AccountPrivilege pulumi.StringPtrInput
	// Privilege type of account.
	// - Normal: Common privilege.
	//   Default to Normal.
	AccountType pulumi.StringPtrInput
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
	InstanceId pulumi.StringInput
	// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//	AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//	AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// The name of the account. The name must meet the following requirements:
// * The name can contain lowercase letters, digits, and hyphens (-), and must start with a lowercase letter.
// * The name can be up to 100 characters in length.
// * The name cannot be one of the reserved words in the [Reserved words for Redis account names](https://www.alibabacloud.com/help/zh/doc-detail/92665.htm) section.
func (o AccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// The password of the account. The password must be 8 to 32 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `!@ # $ % ^ & * ( ) _ + - =`. You have to specify one of `accountPassword` and `kmsEncryptedPassword` fields.
func (o AccountOutput) AccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountPassword }).(pulumi.StringPtrOutput)
}

// The privilege of account access database. Default value: `RoleReadWrite`
func (o AccountOutput) AccountPrivilege() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountPrivilege }).(pulumi.StringPtrOutput)
}

// Privilege type of account.
//   - Normal: Common privilege.
//     Default to Normal.
func (o AccountOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountType }).(pulumi.StringPtrOutput)
}

// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
func (o AccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Id of instance in which account belongs (The engine version of instance must be 4.0 or 4.0+).
func (o AccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// An KMS encrypts password used to a KVStore account. If the `accountPassword` is filled in, this field will be ignored.
func (o AccountOutput) KmsEncryptedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.KmsEncryptedPassword }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a KVStore account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (o AccountOutput) KmsEncryptionContext() pulumi.MapOutput {
	return o.ApplyT(func(v *Account) pulumi.MapOutput { return v.KmsEncryptionContext }).(pulumi.MapOutput)
}

// The status of KVStore Account.
func (o AccountOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
