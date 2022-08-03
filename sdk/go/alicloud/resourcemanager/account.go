// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Manager Account resource. Member accounts are containers for resources in a resource directory. These accounts isolate resources and serve as organizational units in the resource directory. You can create member accounts in a folder and then manage them in a unified manner.
// For information about Resource Manager Account and how to use it, see [What is Resource Manager Account](https://www.alibabacloud.com/help/en/doc-detail/111231.htm).
//
// > **NOTE:** Available in v1.83.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		f1, err := resourcemanager.NewFolder(ctx, "f1", &resourcemanager.FolderArgs{
// 			FolderName: pulumi.String("test1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = resourcemanager.NewAccount(ctx, "example", &resourcemanager.AccountArgs{
// 			DisplayName: pulumi.String("RDAccount"),
// 			FolderId:    f1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Resource Manager Account can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/account:Account example 13148890145*****
// ```
type Account struct {
	pulumi.CustomResourceState

	// The name prefix of account.
	AccountNamePrefix pulumi.StringPtrOutput `pulumi:"accountNamePrefix"`
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod pulumi.StringOutput `pulumi:"joinMethod"`
	// The time when the member joined the resource directory.
	JoinTime pulumi.StringOutput `pulumi:"joinTime"`
	// The modification time of the invitation.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
	PayerAccountId pulumi.StringPtrOutput `pulumi:"payerAccountId"`
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringOutput `pulumi:"resourceDirectoryId"`
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Account
	err := ctx.RegisterResource("alicloud:resourcemanager/account:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:resourcemanager/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The name prefix of account.
	AccountNamePrefix *string `pulumi:"accountNamePrefix"`
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName *string `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId *string `pulumi:"folderId"`
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod *string `pulumi:"joinMethod"`
	// The time when the member joined the resource directory.
	JoinTime *string `pulumi:"joinTime"`
	// The modification time of the invitation.
	ModifyTime *string `pulumi:"modifyTime"`
	// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
	PayerAccountId *string `pulumi:"payerAccountId"`
	// Resource directory ID.
	ResourceDirectoryId *string `pulumi:"resourceDirectoryId"`
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status *string `pulumi:"status"`
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// The name prefix of account.
	AccountNamePrefix pulumi.StringPtrInput
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringPtrInput
	// The ID of the parent folder.
	FolderId pulumi.StringPtrInput
	// Ways for members to join the resource directory. Valid values: `invited`, `created`.
	JoinMethod pulumi.StringPtrInput
	// The time when the member joined the resource directory.
	JoinTime pulumi.StringPtrInput
	// The modification time of the invitation.
	ModifyTime pulumi.StringPtrInput
	// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
	PayerAccountId pulumi.StringPtrInput
	// Resource directory ID.
	ResourceDirectoryId pulumi.StringPtrInput
	// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
	Status pulumi.StringPtrInput
	// Member type. The value of `ResourceAccount` indicates the resource account.
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name prefix of account.
	AccountNamePrefix *string `pulumi:"accountNamePrefix"`
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName string `pulumi:"displayName"`
	// The ID of the parent folder.
	FolderId *string `pulumi:"folderId"`
	// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
	PayerAccountId *string `pulumi:"payerAccountId"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name prefix of account.
	AccountNamePrefix pulumi.StringPtrInput
	// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
	DisplayName pulumi.StringInput
	// The ID of the parent folder.
	FolderId pulumi.StringPtrInput
	// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
	PayerAccountId pulumi.StringPtrInput
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
//          AccountArray{ AccountArgs{...} }
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
//          AccountMap{ "key": AccountArgs{...} }
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

// The name prefix of account.
func (o AccountOutput) AccountNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountNamePrefix }).(pulumi.StringPtrOutput)
}

// Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).
func (o AccountOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The ID of the parent folder.
func (o AccountOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Ways for members to join the resource directory. Valid values: `invited`, `created`.
func (o AccountOutput) JoinMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.JoinMethod }).(pulumi.StringOutput)
}

// The time when the member joined the resource directory.
func (o AccountOutput) JoinTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.JoinTime }).(pulumi.StringOutput)
}

// The modification time of the invitation.
func (o AccountOutput) ModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ModifyTime }).(pulumi.StringOutput)
}

// The ID of the billing account. If you leave this parameter empty, the current account is used as the billing account.
func (o AccountOutput) PayerAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.PayerAccountId }).(pulumi.StringPtrOutput)
}

// Resource directory ID.
func (o AccountOutput) ResourceDirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ResourceDirectoryId }).(pulumi.StringOutput)
}

// Member joining status. Valid values: `CreateSuccess`,`CreateVerifying`,`CreateFailed`,`CreateExpired`,`CreateCancelled`,`PromoteVerifying`,`PromoteFailed`,`PromoteExpired`,`PromoteCancelled`,`PromoteSuccess`,`InviteSuccess`,`Removed`.
func (o AccountOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Member type. The value of `ResourceAccount` indicates the resource account.
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
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
