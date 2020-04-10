// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS account resource and used to manage databases.
type Account struct {
	pulumi.CustomResourceState

	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Privilege type of account.
	// - Normal: Common privilege.
	// - Super: High privilege.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("alicloud:rds/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:rds/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance in which account belongs.
	InstanceId *string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	Name *string `pulumi:"name"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// Privilege type of account.
	// - Normal: Common privilege.
	// - Super: High privilege.
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	Name pulumi.StringPtrInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// Privilege type of account.
	// - Normal: Common privilege.
	// - Super: High privilege.
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance in which account belongs.
	InstanceId string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	Name *string `pulumi:"name"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password *string `pulumi:"password"`
	// Privilege type of account.
	// - Normal: Common privilege.
	// - Super: High privilege.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringInput
	// An KMS encrypts password used to a db account. If the `password` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.
	Name pulumi.StringPtrInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	Password pulumi.StringPtrInput
	// Privilege type of account.
	// - Normal: Common privilege.
	// - Super: High privilege.
	Type pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
