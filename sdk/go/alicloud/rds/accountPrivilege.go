// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an RDS account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account, see [What is DB Account Privilege](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-grantaccountprivilege).
//
// > **NOTE:** At present, a database can only have one database owner.
//
// > **NOTE:** Available since v1.5.0.
//
// ## Example Usage
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
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := rds.NewInstance(ctx, "instance", &rds.InstanceArgs{
//				Engine:          pulumi.String("MySQL"),
//				EngineVersion:   pulumi.String("5.6"),
//				InstanceType:    pulumi.String("rds.mysql.s1.small"),
//				InstanceStorage: pulumi.Int(10),
//				VswitchId:       defaultSwitch.ID(),
//				InstanceName:    pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			var db []*rds.Database
//			for index := 0; index < 2; index++ {
//				key0 := index
//				_ := index
//				__res, err := rds.NewDatabase(ctx, fmt.Sprintf("db-%v", key0), &rds.DatabaseArgs{
//					InstanceId:  instance.ID(),
//					Description: pulumi.String("from terraform"),
//				})
//				if err != nil {
//					return err
//				}
//				db = append(db, __res)
//			}
//			account, err := rds.NewAccount(ctx, "account", &rds.AccountArgs{
//				DbInstanceId:       instance.ID(),
//				AccountName:        pulumi.String("tfexample"),
//				AccountPassword:    pulumi.String("Example12345"),
//				AccountDescription: pulumi.String("from terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			var splat0 pulumi.StringArray
//			for _, val0 := range db {
//				splat0 = append(splat0, val0.Name)
//			}
//			_, err = rds.NewAccountPrivilege(ctx, "privilege", &rds.AccountPrivilegeArgs{
//				InstanceId:  instance.ID(),
//				AccountName: account.Name,
//				Privilege:   pulumi.String("ReadOnly"),
//				DbNames:     splat0,
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
// RDS account privilege can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rds/accountPrivilege:AccountPrivilege example "rm-12345:tf_account:ReadOnly"
//
// ```
type AccountPrivilege struct {
	pulumi.CustomResourceState

	// A specified account name.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// List of specified database name.
	DbNames pulumi.StringArrayOutput `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	//   Default to "ReadOnly".
	Privilege pulumi.StringPtrOutput `pulumi:"privilege"`
}

// NewAccountPrivilege registers a new resource with the given unique name, arguments, and options.
func NewAccountPrivilege(ctx *pulumi.Context,
	name string, args *AccountPrivilegeArgs, opts ...pulumi.ResourceOption) (*AccountPrivilege, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DbNames == nil {
		return nil, errors.New("invalid value for required argument 'DbNames'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountPrivilege
	err := ctx.RegisterResource("alicloud:rds/accountPrivilege:AccountPrivilege", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:rds/accountPrivilege:AccountPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPrivilege resources.
type accountPrivilegeState struct {
	// A specified account name.
	AccountName *string `pulumi:"accountName"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId *string `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	//   Default to "ReadOnly".
	Privilege *string `pulumi:"privilege"`
}

type AccountPrivilegeState struct {
	// A specified account name.
	AccountName pulumi.StringPtrInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringPtrInput
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	//   Default to "ReadOnly".
	Privilege pulumi.StringPtrInput
}

func (AccountPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeState)(nil)).Elem()
}

type accountPrivilegeArgs struct {
	// A specified account name.
	AccountName string `pulumi:"accountName"`
	// List of specified database name.
	DbNames []string `pulumi:"dbNames"`
	// The Id of instance in which account belongs.
	InstanceId string `pulumi:"instanceId"`
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	//   Default to "ReadOnly".
	Privilege *string `pulumi:"privilege"`
}

// The set of arguments for constructing a AccountPrivilege resource.
type AccountPrivilegeArgs struct {
	// A specified account name.
	AccountName pulumi.StringInput
	// List of specified database name.
	DbNames pulumi.StringArrayInput
	// The Id of instance in which account belongs.
	InstanceId pulumi.StringInput
	// The privilege of one account access database. Valid values:
	// - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
	// - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
	// - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
	// - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
	//   Default to "ReadOnly".
	Privilege pulumi.StringPtrInput
}

func (AccountPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPrivilegeArgs)(nil)).Elem()
}

type AccountPrivilegeInput interface {
	pulumi.Input

	ToAccountPrivilegeOutput() AccountPrivilegeOutput
	ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput
}

func (*AccountPrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPrivilege)(nil)).Elem()
}

func (i *AccountPrivilege) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return i.ToAccountPrivilegeOutputWithContext(context.Background())
}

func (i *AccountPrivilege) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeOutput)
}

// AccountPrivilegeArrayInput is an input type that accepts AccountPrivilegeArray and AccountPrivilegeArrayOutput values.
// You can construct a concrete instance of `AccountPrivilegeArrayInput` via:
//
//	AccountPrivilegeArray{ AccountPrivilegeArgs{...} }
type AccountPrivilegeArrayInput interface {
	pulumi.Input

	ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput
	ToAccountPrivilegeArrayOutputWithContext(context.Context) AccountPrivilegeArrayOutput
}

type AccountPrivilegeArray []AccountPrivilegeInput

func (AccountPrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPrivilege)(nil)).Elem()
}

func (i AccountPrivilegeArray) ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput {
	return i.ToAccountPrivilegeArrayOutputWithContext(context.Background())
}

func (i AccountPrivilegeArray) ToAccountPrivilegeArrayOutputWithContext(ctx context.Context) AccountPrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeArrayOutput)
}

// AccountPrivilegeMapInput is an input type that accepts AccountPrivilegeMap and AccountPrivilegeMapOutput values.
// You can construct a concrete instance of `AccountPrivilegeMapInput` via:
//
//	AccountPrivilegeMap{ "key": AccountPrivilegeArgs{...} }
type AccountPrivilegeMapInput interface {
	pulumi.Input

	ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput
	ToAccountPrivilegeMapOutputWithContext(context.Context) AccountPrivilegeMapOutput
}

type AccountPrivilegeMap map[string]AccountPrivilegeInput

func (AccountPrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPrivilege)(nil)).Elem()
}

func (i AccountPrivilegeMap) ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput {
	return i.ToAccountPrivilegeMapOutputWithContext(context.Background())
}

func (i AccountPrivilegeMap) ToAccountPrivilegeMapOutputWithContext(ctx context.Context) AccountPrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPrivilegeMapOutput)
}

type AccountPrivilegeOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutput() AccountPrivilegeOutput {
	return o
}

func (o AccountPrivilegeOutput) ToAccountPrivilegeOutputWithContext(ctx context.Context) AccountPrivilegeOutput {
	return o
}

// A specified account name.
func (o AccountPrivilegeOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountPrivilege) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

// List of specified database name.
func (o AccountPrivilegeOutput) DbNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountPrivilege) pulumi.StringArrayOutput { return v.DbNames }).(pulumi.StringArrayOutput)
}

// The Id of instance in which account belongs.
func (o AccountPrivilegeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountPrivilege) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The privilege of one account access database. Valid values:
//   - ReadOnly: This value is only for MySQL, MariaDB and SQL Server
//   - ReadWrite: This value is only for MySQL, MariaDB and SQL Server
//   - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
//   - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB
//   - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL.
//     Default to "ReadOnly".
func (o AccountPrivilegeOutput) Privilege() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPrivilege) pulumi.StringPtrOutput { return v.Privilege }).(pulumi.StringPtrOutput)
}

type AccountPrivilegeArrayOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeArrayOutput) ToAccountPrivilegeArrayOutput() AccountPrivilegeArrayOutput {
	return o
}

func (o AccountPrivilegeArrayOutput) ToAccountPrivilegeArrayOutputWithContext(ctx context.Context) AccountPrivilegeArrayOutput {
	return o
}

func (o AccountPrivilegeArrayOutput) Index(i pulumi.IntInput) AccountPrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountPrivilege {
		return vs[0].([]*AccountPrivilege)[vs[1].(int)]
	}).(AccountPrivilegeOutput)
}

type AccountPrivilegeMapOutput struct{ *pulumi.OutputState }

func (AccountPrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPrivilege)(nil)).Elem()
}

func (o AccountPrivilegeMapOutput) ToAccountPrivilegeMapOutput() AccountPrivilegeMapOutput {
	return o
}

func (o AccountPrivilegeMapOutput) ToAccountPrivilegeMapOutputWithContext(ctx context.Context) AccountPrivilegeMapOutput {
	return o
}

func (o AccountPrivilegeMapOutput) MapIndex(k pulumi.StringInput) AccountPrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountPrivilege {
		return vs[0].(map[string]*AccountPrivilege)[vs[1].(string)]
	}).(AccountPrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeInput)(nil)).Elem(), &AccountPrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeArrayInput)(nil)).Elem(), AccountPrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPrivilegeMapInput)(nil)).Elem(), AccountPrivilegeMap{})
	pulumi.RegisterOutputType(AccountPrivilegeOutput{})
	pulumi.RegisterOutputType(AccountPrivilegeArrayOutput{})
	pulumi.RegisterOutputType(AccountPrivilegeMapOutput{})
}
