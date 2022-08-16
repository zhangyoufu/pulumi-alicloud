// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gpdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GPDB Account resource.
//
// For information about GPDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/doc-detail/86924.htm).
//
// > **NOTE:** Available in v1.142.0+.
//
// ## Import
//
// GPDB Account can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:gpdb/account:Account example <db_instance_id>:<account_name>
//
// ```
type Account struct {
	pulumi.CustomResourceState

	// The description of the account.
	// * Starts with a letter.
	// * Does not start with `http://` or `https://`.
	// * Contains letters, underscores (_), hyphens (-), or digits.
	// * Be 2 to 256 characters in length.
	AccountDescription pulumi.StringPtrOutput `pulumi:"accountDescription"`
	// The name of the account. The account name must be unique and meet the following requirements:
	// * Starts with a letter.
	// * Contains only lowercase letters, digits, or underscores (_).
	// * Be up to 16 characters in length.
	// * Contains no reserved keywords.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
	AccountPassword pulumi.StringOutput `pulumi:"accountPassword"`
	// The ID of the instance.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
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
	if args.AccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'AccountPassword'")
	}
	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	var resource Account
	err := ctx.RegisterResource("alicloud:gpdb/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:gpdb/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The description of the account.
	// * Starts with a letter.
	// * Does not start with `http://` or `https://`.
	// * Contains letters, underscores (_), hyphens (-), or digits.
	// * Be 2 to 256 characters in length.
	AccountDescription *string `pulumi:"accountDescription"`
	// The name of the account. The account name must be unique and meet the following requirements:
	// * Starts with a letter.
	// * Contains only lowercase letters, digits, or underscores (_).
	// * Be up to 16 characters in length.
	// * Contains no reserved keywords.
	AccountName *string `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
	AccountPassword *string `pulumi:"accountPassword"`
	// The ID of the instance.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
	Status *string `pulumi:"status"`
}

type AccountState struct {
	// The description of the account.
	// * Starts with a letter.
	// * Does not start with `http://` or `https://`.
	// * Contains letters, underscores (_), hyphens (-), or digits.
	// * Be 2 to 256 characters in length.
	AccountDescription pulumi.StringPtrInput
	// The name of the account. The account name must be unique and meet the following requirements:
	// * Starts with a letter.
	// * Contains only lowercase letters, digits, or underscores (_).
	// * Be up to 16 characters in length.
	// * Contains no reserved keywords.
	AccountName pulumi.StringPtrInput
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
	AccountPassword pulumi.StringPtrInput
	// The ID of the instance.
	DbInstanceId pulumi.StringPtrInput
	// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
	Status pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The description of the account.
	// * Starts with a letter.
	// * Does not start with `http://` or `https://`.
	// * Contains letters, underscores (_), hyphens (-), or digits.
	// * Be 2 to 256 characters in length.
	AccountDescription *string `pulumi:"accountDescription"`
	// The name of the account. The account name must be unique and meet the following requirements:
	// * Starts with a letter.
	// * Contains only lowercase letters, digits, or underscores (_).
	// * Be up to 16 characters in length.
	// * Contains no reserved keywords.
	AccountName string `pulumi:"accountName"`
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
	AccountPassword string `pulumi:"accountPassword"`
	// The ID of the instance.
	DbInstanceId string `pulumi:"dbInstanceId"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The description of the account.
	// * Starts with a letter.
	// * Does not start with `http://` or `https://`.
	// * Contains letters, underscores (_), hyphens (-), or digits.
	// * Be 2 to 256 characters in length.
	AccountDescription pulumi.StringPtrInput
	// The name of the account. The account name must be unique and meet the following requirements:
	// * Starts with a letter.
	// * Contains only lowercase letters, digits, or underscores (_).
	// * Be up to 16 characters in length.
	// * Contains no reserved keywords.
	AccountName pulumi.StringInput
	// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
	AccountPassword pulumi.StringInput
	// The ID of the instance.
	DbInstanceId pulumi.StringInput
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

// The description of the account.
// * Starts with a letter.
// * Does not start with `http://` or `https://`.
// * Contains letters, underscores (_), hyphens (-), or digits.
// * Be 2 to 256 characters in length.
func (o AccountOutput) AccountDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountDescription }).(pulumi.StringPtrOutput)
}

// The name of the account. The account name must be unique and meet the following requirements:
// * Starts with a letter.
// * Contains only lowercase letters, digits, or underscores (_).
// * Be up to 16 characters in length.
// * Contains no reserved keywords.
func (o AccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & * ( ) _ + - =`.
func (o AccountOutput) AccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountPassword }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o AccountOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
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
