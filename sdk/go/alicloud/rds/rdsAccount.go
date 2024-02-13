// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RDS Account resource.
//
// For information about RDS Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createaccount).
//
// > **NOTE:** Available since v1.120.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
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
//			defaultZones, err := rds.GetZones(ctx, &rds.GetZonesArgs{
//				Engine:        pulumi.StringRef("MySQL"),
//				EngineVersion: pulumi.StringRef("5.6"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceClasses, err := rds.GetInstanceClasses(ctx, &rds.GetInstanceClassesArgs{
//				ZoneId:        pulumi.StringRef(defaultZones.Ids[0]),
//				Engine:        pulumi.StringRef("MySQL"),
//				EngineVersion: pulumi.StringRef("5.6"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Ids[0]),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
//				Engine:          pulumi.String("MySQL"),
//				EngineVersion:   pulumi.String("5.6"),
//				InstanceType:    *pulumi.String(defaultInstanceClasses.InstanceClasses[0].InstanceClass),
//				InstanceStorage: pulumi.Int(10),
//				VswitchId:       defaultSwitch.ID(),
//				InstanceName:    pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewRdsAccount(ctx, "defaultRdsAccount", &rds.RdsAccountArgs{
//				DbInstanceId:    defaultInstance.ID(),
//				AccountName:     pulumi.String(name),
//				AccountPassword: pulumi.String("Example1234"),
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
// RDS Account can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:rds/rdsAccount:RdsAccount example <db_instance_id>:<account_name>
// ```
type RdsAccount struct {
	pulumi.CustomResourceState

	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringOutput `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// Privilege type of account. Default to `Normal`.
	// `Normal`: Common privilege.
	// `Super`: High privilege.
	AccountType pulumi.StringOutput `pulumi:"accountType"`
	// The Id of instance in which account belongs.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
	//
	// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
	Description pulumi.StringOutput `pulumi:"description"`
	// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
	//
	// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrOutput `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapOutput `pulumi:"kmsEncryptionContext"`
	// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
	//
	// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
	Password pulumi.StringOutput `pulumi:"password"`
	// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
	ResetPermissionFlag pulumi.BoolPtrOutput `pulumi:"resetPermissionFlag"`
	// The status of the resource. Valid values: `Available`, `Unavailable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
	//
	// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRdsAccount registers a new resource with the given unique name, arguments, and options.
func NewRdsAccount(ctx *pulumi.Context,
	name string, args *RdsAccountArgs, opts ...pulumi.ResourceOption) (*RdsAccount, error) {
	if args == nil {
		args = &RdsAccountArgs{}
	}

	if args.AccountPassword != nil {
		args.AccountPassword = pulumi.ToSecret(args.AccountPassword).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accountPassword",
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsAccount
	err := ctx.RegisterResource("alicloud:rds/rdsAccount:RdsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsAccount gets an existing RdsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsAccountState, opts ...pulumi.ResourceOption) (*RdsAccount, error) {
	var resource RdsAccount
	err := ctx.ReadResource("alicloud:rds/rdsAccount:RdsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsAccount resources.
type rdsAccountState struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
	AccountName *string `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	AccountPassword *string `pulumi:"accountPassword"`
	// Privilege type of account. Default to `Normal`.
	// `Normal`: Common privilege.
	// `Super`: High privilege.
	AccountType *string `pulumi:"accountType"`
	// The Id of instance in which account belongs.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
	//
	// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
	Description *string `pulumi:"description"`
	// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
	//
	// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
	InstanceId *string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
	Name *string `pulumi:"name"`
	// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
	//
	// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
	Password *string `pulumi:"password"`
	// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
	ResetPermissionFlag *bool `pulumi:"resetPermissionFlag"`
	// The status of the resource. Valid values: `Available`, `Unavailable`.
	Status *string `pulumi:"status"`
	// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
	//
	// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
	Type *string `pulumi:"type"`
}

type RdsAccountState struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringPtrInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
	AccountName pulumi.StringPtrInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringPtrInput
	// Privilege type of account. Default to `Normal`.
	// `Normal`: Common privilege.
	// `Super`: High privilege.
	AccountType pulumi.StringPtrInput
	// The Id of instance in which account belongs.
	DbInstanceId pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
	//
	// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
	Description pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
	//
	// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
	InstanceId pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
	Name pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
	//
	// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
	Password pulumi.StringPtrInput
	// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
	ResetPermissionFlag pulumi.BoolPtrInput
	// The status of the resource. Valid values: `Available`, `Unavailable`.
	Status pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
	//
	// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
	Type pulumi.StringPtrInput
}

func (RdsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsAccountState)(nil)).Elem()
}

type rdsAccountArgs struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription *string `pulumi:"accountDescription"`
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
	AccountName *string `pulumi:"accountName"`
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	AccountPassword *string `pulumi:"accountPassword"`
	// Privilege type of account. Default to `Normal`.
	// `Normal`: Common privilege.
	// `Super`: High privilege.
	AccountType *string `pulumi:"accountType"`
	// The Id of instance in which account belongs.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
	//
	// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
	Description *string `pulumi:"description"`
	// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
	//
	// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
	InstanceId *string `pulumi:"instanceId"`
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword *string `pulumi:"kmsEncryptedPassword"`
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext map[string]interface{} `pulumi:"kmsEncryptionContext"`
	// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
	Name *string `pulumi:"name"`
	// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
	//
	// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
	Password *string `pulumi:"password"`
	// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
	ResetPermissionFlag *bool `pulumi:"resetPermissionFlag"`
	// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
	//
	// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a RdsAccount resource.
type RdsAccountArgs struct {
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	AccountDescription pulumi.StringPtrInput
	// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
	AccountName pulumi.StringPtrInput
	// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
	AccountPassword pulumi.StringPtrInput
	// Privilege type of account. Default to `Normal`.
	// `Normal`: Common privilege.
	// `Super`: High privilege.
	AccountType pulumi.StringPtrInput
	// The Id of instance in which account belongs.
	DbInstanceId pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
	//
	// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
	Description pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
	//
	// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
	InstanceId pulumi.StringPtrInput
	// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
	KmsEncryptedPassword pulumi.StringPtrInput
	// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
	KmsEncryptionContext pulumi.MapInput
	// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
	Name pulumi.StringPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
	//
	// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
	Password pulumi.StringPtrInput
	// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
	ResetPermissionFlag pulumi.BoolPtrInput
	// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
	//
	// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
	//
	// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
	Type pulumi.StringPtrInput
}

func (RdsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsAccountArgs)(nil)).Elem()
}

type RdsAccountInput interface {
	pulumi.Input

	ToRdsAccountOutput() RdsAccountOutput
	ToRdsAccountOutputWithContext(ctx context.Context) RdsAccountOutput
}

func (*RdsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsAccount)(nil)).Elem()
}

func (i *RdsAccount) ToRdsAccountOutput() RdsAccountOutput {
	return i.ToRdsAccountOutputWithContext(context.Background())
}

func (i *RdsAccount) ToRdsAccountOutputWithContext(ctx context.Context) RdsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsAccountOutput)
}

// RdsAccountArrayInput is an input type that accepts RdsAccountArray and RdsAccountArrayOutput values.
// You can construct a concrete instance of `RdsAccountArrayInput` via:
//
//	RdsAccountArray{ RdsAccountArgs{...} }
type RdsAccountArrayInput interface {
	pulumi.Input

	ToRdsAccountArrayOutput() RdsAccountArrayOutput
	ToRdsAccountArrayOutputWithContext(context.Context) RdsAccountArrayOutput
}

type RdsAccountArray []RdsAccountInput

func (RdsAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsAccount)(nil)).Elem()
}

func (i RdsAccountArray) ToRdsAccountArrayOutput() RdsAccountArrayOutput {
	return i.ToRdsAccountArrayOutputWithContext(context.Background())
}

func (i RdsAccountArray) ToRdsAccountArrayOutputWithContext(ctx context.Context) RdsAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsAccountArrayOutput)
}

// RdsAccountMapInput is an input type that accepts RdsAccountMap and RdsAccountMapOutput values.
// You can construct a concrete instance of `RdsAccountMapInput` via:
//
//	RdsAccountMap{ "key": RdsAccountArgs{...} }
type RdsAccountMapInput interface {
	pulumi.Input

	ToRdsAccountMapOutput() RdsAccountMapOutput
	ToRdsAccountMapOutputWithContext(context.Context) RdsAccountMapOutput
}

type RdsAccountMap map[string]RdsAccountInput

func (RdsAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsAccount)(nil)).Elem()
}

func (i RdsAccountMap) ToRdsAccountMapOutput() RdsAccountMapOutput {
	return i.ToRdsAccountMapOutputWithContext(context.Background())
}

func (i RdsAccountMap) ToRdsAccountMapOutputWithContext(ctx context.Context) RdsAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsAccountMapOutput)
}

type RdsAccountOutput struct{ *pulumi.OutputState }

func (RdsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsAccount)(nil)).Elem()
}

func (o RdsAccountOutput) ToRdsAccountOutput() RdsAccountOutput {
	return o
}

func (o RdsAccountOutput) ToRdsAccountOutputWithContext(ctx context.Context) RdsAccountOutput {
	return o
}

// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
func (o RdsAccountOutput) AccountDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.AccountDescription }).(pulumi.StringOutput)
}

// Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and end with letters or numbers, The length must be 2-63 characters for PostgreSQL, otherwise the length must be 2-32 characters.
func (o RdsAccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of `password` and `kmsEncryptedPassword` fields.
func (o RdsAccountOutput) AccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.AccountPassword }).(pulumi.StringOutput)
}

// Privilege type of account. Default to `Normal`.
// `Normal`: Common privilege.
// `Super`: High privilege.
func (o RdsAccountOutput) AccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.AccountType }).(pulumi.StringOutput)
}

// The Id of instance in which account belongs.
func (o RdsAccountOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The attribute has been deprecated from 1.120.0 and using `accountDescription` instead.
//
// Deprecated: Field 'description' has been deprecated from provider version 1.120.0. New field 'account_description' instead.
func (o RdsAccountOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The attribute has been deprecated from 1.120.0 and using `dbInstanceId` instead.
//
// Deprecated: Field 'instance_id' has been deprecated from provider version 1.120.0. New field 'db_instance_id' instead.
func (o RdsAccountOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// An KMS encrypts password used to a db account. If the `accountPassword` is filled in, this field will be ignored.
func (o RdsAccountOutput) KmsEncryptedPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringPtrOutput { return v.KmsEncryptedPassword }).(pulumi.StringPtrOutput)
}

// An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating a db account with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
func (o RdsAccountOutput) KmsEncryptionContext() pulumi.MapOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.MapOutput { return v.KmsEncryptionContext }).(pulumi.MapOutput)
}

// The attribute has been deprecated from 1.120.0 and using `accountName` instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.120.0. New field 'account_name' instead.
func (o RdsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The attribute has been deprecated from 1.120.0 and using `accountPassword` instead.
//
// Deprecated: Field 'password' has been deprecated from provider version 1.120.0. New field 'account_password' instead.
func (o RdsAccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Resets permissions flag of the privileged account. Default to `false`. Set it to `true` can resets permissions of the privileged account.
func (o RdsAccountOutput) ResetPermissionFlag() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.BoolPtrOutput { return v.ResetPermissionFlag }).(pulumi.BoolPtrOutput)
}

// The status of the resource. Valid values: `Available`, `Unavailable`.
func (o RdsAccountOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The attribute has been deprecated from 1.120.0 and using `accountType` instead.
//
// > **NOTE**: Only MySQL engine is supported resets permissions of the privileged account.
//
// Deprecated: Field 'type' has been deprecated from provider version 1.120.0. New field 'account_type' instead.
func (o RdsAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type RdsAccountArrayOutput struct{ *pulumi.OutputState }

func (RdsAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsAccount)(nil)).Elem()
}

func (o RdsAccountArrayOutput) ToRdsAccountArrayOutput() RdsAccountArrayOutput {
	return o
}

func (o RdsAccountArrayOutput) ToRdsAccountArrayOutputWithContext(ctx context.Context) RdsAccountArrayOutput {
	return o
}

func (o RdsAccountArrayOutput) Index(i pulumi.IntInput) RdsAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsAccount {
		return vs[0].([]*RdsAccount)[vs[1].(int)]
	}).(RdsAccountOutput)
}

type RdsAccountMapOutput struct{ *pulumi.OutputState }

func (RdsAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsAccount)(nil)).Elem()
}

func (o RdsAccountMapOutput) ToRdsAccountMapOutput() RdsAccountMapOutput {
	return o
}

func (o RdsAccountMapOutput) ToRdsAccountMapOutputWithContext(ctx context.Context) RdsAccountMapOutput {
	return o
}

func (o RdsAccountMapOutput) MapIndex(k pulumi.StringInput) RdsAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsAccount {
		return vs[0].(map[string]*RdsAccount)[vs[1].(string)]
	}).(RdsAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsAccountInput)(nil)).Elem(), &RdsAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsAccountArrayInput)(nil)).Elem(), RdsAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsAccountMapInput)(nil)).Elem(), RdsAccountMap{})
	pulumi.RegisterOutputType(RdsAccountOutput{})
	pulumi.RegisterOutputType(RdsAccountArrayOutput{})
	pulumi.RegisterOutputType(RdsAccountMapOutput{})
}
