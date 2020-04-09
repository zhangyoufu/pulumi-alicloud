// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package polardb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a PolarDB account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
//
// > **NOTE:** Available in v1.67.0+.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account_privilege.html.markdown.
type AccountPrivilege struct {
	pulumi.CustomResourceState

	// A specified account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrOutput `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames pulumi.StringArrayOutput `pulumi:"dbNames"`
}

// NewAccountPrivilege registers a new resource with the given unique name, arguments, and options.
func NewAccountPrivilege(ctx *pulumi.Context,
	name string, args *AccountPrivilegeArgs, opts ...pulumi.ResourceOption) (*AccountPrivilege, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DbClusterId == nil {
		return nil, errors.New("missing required argument 'DbClusterId'")
	}
	if args == nil || args.DbNames == nil {
		return nil, errors.New("missing required argument 'DbNames'")
	}
	if args == nil {
		args = &AccountPrivilegeArgs{}
	}
	var resource AccountPrivilege
	err := ctx.RegisterResource("alicloud:polardb/accountPrivilege:AccountPrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPrivilege gets an existing AccountPrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPrivilegeState, opts ...pulumi.ResourceOption) (*AccountPrivilege, error) {
	var resource AccountPrivilege
	err := ctx.ReadResource("alicloud:polardb/accountPrivilege:AccountPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPrivilege resources.
type accountPrivilegeState struct {
	// A specified account name.
	AccountName *string `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId *string `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
}

type AccountPrivilegeState struct {
	// A specified account name.
	AccountName pulumi.StringPtrInput
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringPtrInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
}

func (AccountPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeState)(nil)).Elem()
}

type accountPrivilegeArgs struct {
	// A specified account name.
	AccountName string `pulumi:"accountName"`
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
	AccountPrivilege *string `pulumi:"accountPrivilege"`
	// The Id of cluster in which account belongs.
	DbClusterId string `pulumi:"dbClusterId"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
}

// The set of arguments for constructing a AccountPrivilege resource.
type AccountPrivilegeArgs struct {
	// A specified account name.
	AccountName pulumi.StringInput
	// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
	AccountPrivilege pulumi.StringPtrInput
	// The Id of cluster in which account belongs.
	DbClusterId pulumi.StringInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
}

func (AccountPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeArgs)(nil)).Elem()
}
