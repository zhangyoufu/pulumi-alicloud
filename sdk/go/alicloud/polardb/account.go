// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB account resource and used to manage databases.
//
// > **NOTE:** Available in v1.67.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultNodeClasses, err := polardb.GetNodeClasses(ctx, &polardb.GetNodeClassesArgs{
//				DbType:    pulumi.StringRef("MySQL"),
//				DbVersion: pulumi.StringRef("8.0"),
//				PayType:   "PostPaid",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(defaultNodeClasses.Classes[0].ZoneId),
//				VswitchName: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultCluster, err := polardb.NewCluster(ctx, "defaultCluster", &polardb.ClusterArgs{
//				DbType:      pulumi.String("MySQL"),
//				DbVersion:   pulumi.String("8.0"),
//				DbNodeClass: *pulumi.String(defaultNodeClasses.Classes[0].SupportedEngines[0].AvailableResources[0].DbNodeClass),
//				PayType:     pulumi.String("PostPaid"),
//				VswitchId:   defaultSwitch.ID(),
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = polardb.NewAccount(ctx, "defaultAccount", &polardb.AccountArgs{
//				DbClusterId:        defaultCluster.ID(),
//				AccountName:        pulumi.String("terraform_example"),
//				AccountPassword:    pulumi.String("Example1234"),
//				AccountDescription: pulumi.String("terraform-example"),
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
// PolarDB account can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:polardb/account:Account example "pc-12345:tf_account"
//
// ```
type Account struct {
	pulumi.CustomResourceState

	// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringPtrOutput `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
	AccountType pulumi.StringPtrOutput `pulumi:"accountType"`
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
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
	if args.AccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'AccountPassword'")
	}
	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
	})
	opts = append(opts, secrets)
	var resource Account
	err := ctx.RegisterResource("alicloud:polardb/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:polardb/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	AccountName *string `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
	AccountPassword *string `pulumi:"accountPassword"`
	// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
	AccountType *string `pulumi:"accountType"`
	// The Id of cluster in which account belongs.
	DbClusterId *string `pulumi:"dbClusterId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
}

type AccountState struct {
	// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringPtrInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	AccountName pulumi.StringPtrInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
	AccountPassword pulumi.StringPtrInput
	// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
	AccountType pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	AccountName string `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
	AccountPassword string `pulumi:"accountPassword"`
	// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
	AccountType *string `pulumi:"accountType"`
	// The Id of cluster in which account belongs.
	DbClusterId string `pulumi:"dbClusterId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringPtrInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	AccountName pulumi.StringInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
	AccountPassword pulumi.StringInput
	// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
	AccountType pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringInput
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
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

// Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
func (o AccountOutput) AccountDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountDescription }).(pulumi.StringPtrOutput)
}

// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
func (o AccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.
func (o AccountOutput) AccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountPassword }).(pulumi.StringOutput)
}

// Account type, Valid values are `Normal`, `Super`, Default to `Normal`.
func (o AccountOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountType }).(pulumi.StringPtrOutput)
}

// The Id of cluster in which account belongs.
func (o AccountOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
func (o AccountOutput) KmsEncryptedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.KmsEncryptedPassword }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (o AccountOutput) KmsEncryptionContext() pulumi.MapOutput {
	return o.ApplyT(func(v *Account) pulumi.MapOutput { return v.KmsEncryptionContext }).(pulumi.MapOutput)
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
