// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS Enterprise Logic Database resource.
//
// For information about DMS Enterprise Logic Database and how to use it, see [What is Logic Database](https://www.alibabacloud.com/help/zh/data-management-service/latest/api-doc-dms-enterprise-2018-11-01-api-doc-createlogicdatabase).
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dms.NewEnterpriseLogicDatabase(ctx, "default", &dms.EnterpriseLogicDatabaseArgs{
//				Alias: pulumi.String("TF_logic_db_test"),
//				DatabaseIds: pulumi.StringArray{
//					pulumi.String("35617919"),
//					pulumi.String("35617920"),
//				},
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
// DMS Enterprise Logic Database can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase example <id>
//
// ```
type EnterpriseLogicDatabase struct {
	pulumi.CustomResourceState

	// Logical Library alias.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Sub-Database ID
	DatabaseIds pulumi.StringArrayOutput `pulumi:"databaseIds"`
	// Database type.
	DbType pulumi.StringOutput `pulumi:"dbType"`
	// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
	EnvType pulumi.StringOutput `pulumi:"envType"`
	// Whether it is a logical Library, the return value is true.
	Logic pulumi.BoolOutput `pulumi:"logic"`
	// The ID of the logical Library.
	LogicDatabaseId pulumi.StringOutput `pulumi:"logicDatabaseId"`
	// The user ID list of the logical library Owner.
	OwnerIdLists pulumi.StringArrayOutput `pulumi:"ownerIdLists"`
	// The nickname list of the logical library Owner.
	OwnerNameLists pulumi.StringArrayOutput `pulumi:"ownerNameLists"`
	// Logical Library name.
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// Logical library search name.
	SearchName pulumi.StringOutput `pulumi:"searchName"`
}

// NewEnterpriseLogicDatabase registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseLogicDatabase(ctx *pulumi.Context,
	name string, args *EnterpriseLogicDatabaseArgs, opts ...pulumi.ResourceOption) (*EnterpriseLogicDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.DatabaseIds == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseIds'")
	}
	var resource EnterpriseLogicDatabase
	err := ctx.RegisterResource("alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseLogicDatabase gets an existing EnterpriseLogicDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseLogicDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseLogicDatabaseState, opts ...pulumi.ResourceOption) (*EnterpriseLogicDatabase, error) {
	var resource EnterpriseLogicDatabase
	err := ctx.ReadResource("alicloud:dms/enterpriseLogicDatabase:EnterpriseLogicDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseLogicDatabase resources.
type enterpriseLogicDatabaseState struct {
	// Logical Library alias.
	Alias *string `pulumi:"alias"`
	// Sub-Database ID
	DatabaseIds []string `pulumi:"databaseIds"`
	// Database type.
	DbType *string `pulumi:"dbType"`
	// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
	EnvType *string `pulumi:"envType"`
	// Whether it is a logical Library, the return value is true.
	Logic *bool `pulumi:"logic"`
	// The ID of the logical Library.
	LogicDatabaseId *string `pulumi:"logicDatabaseId"`
	// The user ID list of the logical library Owner.
	OwnerIdLists []string `pulumi:"ownerIdLists"`
	// The nickname list of the logical library Owner.
	OwnerNameLists []string `pulumi:"ownerNameLists"`
	// Logical Library name.
	SchemaName *string `pulumi:"schemaName"`
	// Logical library search name.
	SearchName *string `pulumi:"searchName"`
}

type EnterpriseLogicDatabaseState struct {
	// Logical Library alias.
	Alias pulumi.StringPtrInput
	// Sub-Database ID
	DatabaseIds pulumi.StringArrayInput
	// Database type.
	DbType pulumi.StringPtrInput
	// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
	EnvType pulumi.StringPtrInput
	// Whether it is a logical Library, the return value is true.
	Logic pulumi.BoolPtrInput
	// The ID of the logical Library.
	LogicDatabaseId pulumi.StringPtrInput
	// The user ID list of the logical library Owner.
	OwnerIdLists pulumi.StringArrayInput
	// The nickname list of the logical library Owner.
	OwnerNameLists pulumi.StringArrayInput
	// Logical Library name.
	SchemaName pulumi.StringPtrInput
	// Logical library search name.
	SearchName pulumi.StringPtrInput
}

func (EnterpriseLogicDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseLogicDatabaseState)(nil)).Elem()
}

type enterpriseLogicDatabaseArgs struct {
	// Logical Library alias.
	Alias string `pulumi:"alias"`
	// Sub-Database ID
	DatabaseIds []string `pulumi:"databaseIds"`
	// The ID of the logical Library.
	LogicDatabaseId *string `pulumi:"logicDatabaseId"`
}

// The set of arguments for constructing a EnterpriseLogicDatabase resource.
type EnterpriseLogicDatabaseArgs struct {
	// Logical Library alias.
	Alias pulumi.StringInput
	// Sub-Database ID
	DatabaseIds pulumi.StringArrayInput
	// The ID of the logical Library.
	LogicDatabaseId pulumi.StringPtrInput
}

func (EnterpriseLogicDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseLogicDatabaseArgs)(nil)).Elem()
}

type EnterpriseLogicDatabaseInput interface {
	pulumi.Input

	ToEnterpriseLogicDatabaseOutput() EnterpriseLogicDatabaseOutput
	ToEnterpriseLogicDatabaseOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseOutput
}

func (*EnterpriseLogicDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseLogicDatabase)(nil)).Elem()
}

func (i *EnterpriseLogicDatabase) ToEnterpriseLogicDatabaseOutput() EnterpriseLogicDatabaseOutput {
	return i.ToEnterpriseLogicDatabaseOutputWithContext(context.Background())
}

func (i *EnterpriseLogicDatabase) ToEnterpriseLogicDatabaseOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLogicDatabaseOutput)
}

// EnterpriseLogicDatabaseArrayInput is an input type that accepts EnterpriseLogicDatabaseArray and EnterpriseLogicDatabaseArrayOutput values.
// You can construct a concrete instance of `EnterpriseLogicDatabaseArrayInput` via:
//
//	EnterpriseLogicDatabaseArray{ EnterpriseLogicDatabaseArgs{...} }
type EnterpriseLogicDatabaseArrayInput interface {
	pulumi.Input

	ToEnterpriseLogicDatabaseArrayOutput() EnterpriseLogicDatabaseArrayOutput
	ToEnterpriseLogicDatabaseArrayOutputWithContext(context.Context) EnterpriseLogicDatabaseArrayOutput
}

type EnterpriseLogicDatabaseArray []EnterpriseLogicDatabaseInput

func (EnterpriseLogicDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseLogicDatabase)(nil)).Elem()
}

func (i EnterpriseLogicDatabaseArray) ToEnterpriseLogicDatabaseArrayOutput() EnterpriseLogicDatabaseArrayOutput {
	return i.ToEnterpriseLogicDatabaseArrayOutputWithContext(context.Background())
}

func (i EnterpriseLogicDatabaseArray) ToEnterpriseLogicDatabaseArrayOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLogicDatabaseArrayOutput)
}

// EnterpriseLogicDatabaseMapInput is an input type that accepts EnterpriseLogicDatabaseMap and EnterpriseLogicDatabaseMapOutput values.
// You can construct a concrete instance of `EnterpriseLogicDatabaseMapInput` via:
//
//	EnterpriseLogicDatabaseMap{ "key": EnterpriseLogicDatabaseArgs{...} }
type EnterpriseLogicDatabaseMapInput interface {
	pulumi.Input

	ToEnterpriseLogicDatabaseMapOutput() EnterpriseLogicDatabaseMapOutput
	ToEnterpriseLogicDatabaseMapOutputWithContext(context.Context) EnterpriseLogicDatabaseMapOutput
}

type EnterpriseLogicDatabaseMap map[string]EnterpriseLogicDatabaseInput

func (EnterpriseLogicDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseLogicDatabase)(nil)).Elem()
}

func (i EnterpriseLogicDatabaseMap) ToEnterpriseLogicDatabaseMapOutput() EnterpriseLogicDatabaseMapOutput {
	return i.ToEnterpriseLogicDatabaseMapOutputWithContext(context.Background())
}

func (i EnterpriseLogicDatabaseMap) ToEnterpriseLogicDatabaseMapOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseLogicDatabaseMapOutput)
}

type EnterpriseLogicDatabaseOutput struct{ *pulumi.OutputState }

func (EnterpriseLogicDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseLogicDatabase)(nil)).Elem()
}

func (o EnterpriseLogicDatabaseOutput) ToEnterpriseLogicDatabaseOutput() EnterpriseLogicDatabaseOutput {
	return o
}

func (o EnterpriseLogicDatabaseOutput) ToEnterpriseLogicDatabaseOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseOutput {
	return o
}

// Logical Library alias.
func (o EnterpriseLogicDatabaseOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Sub-Database ID
func (o EnterpriseLogicDatabaseOutput) DatabaseIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringArrayOutput { return v.DatabaseIds }).(pulumi.StringArrayOutput)
}

// Database type.
func (o EnterpriseLogicDatabaseOutput) DbType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.DbType }).(pulumi.StringOutput)
}

// Environment type, return value is as follows:-product: production environment-dev: development environment-pre: Advance Environment-test: test environment-sit:SIT environment-uat:UAT environment-pet: Pressure measurement environment-stag:STAG environment
func (o EnterpriseLogicDatabaseOutput) EnvType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.EnvType }).(pulumi.StringOutput)
}

// Whether it is a logical Library, the return value is true.
func (o EnterpriseLogicDatabaseOutput) Logic() pulumi.BoolOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.BoolOutput { return v.Logic }).(pulumi.BoolOutput)
}

// The ID of the logical Library.
func (o EnterpriseLogicDatabaseOutput) LogicDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.LogicDatabaseId }).(pulumi.StringOutput)
}

// The user ID list of the logical library Owner.
func (o EnterpriseLogicDatabaseOutput) OwnerIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringArrayOutput { return v.OwnerIdLists }).(pulumi.StringArrayOutput)
}

// The nickname list of the logical library Owner.
func (o EnterpriseLogicDatabaseOutput) OwnerNameLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringArrayOutput { return v.OwnerNameLists }).(pulumi.StringArrayOutput)
}

// Logical Library name.
func (o EnterpriseLogicDatabaseOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// Logical library search name.
func (o EnterpriseLogicDatabaseOutput) SearchName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseLogicDatabase) pulumi.StringOutput { return v.SearchName }).(pulumi.StringOutput)
}

type EnterpriseLogicDatabaseArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseLogicDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseLogicDatabase)(nil)).Elem()
}

func (o EnterpriseLogicDatabaseArrayOutput) ToEnterpriseLogicDatabaseArrayOutput() EnterpriseLogicDatabaseArrayOutput {
	return o
}

func (o EnterpriseLogicDatabaseArrayOutput) ToEnterpriseLogicDatabaseArrayOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseArrayOutput {
	return o
}

func (o EnterpriseLogicDatabaseArrayOutput) Index(i pulumi.IntInput) EnterpriseLogicDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseLogicDatabase {
		return vs[0].([]*EnterpriseLogicDatabase)[vs[1].(int)]
	}).(EnterpriseLogicDatabaseOutput)
}

type EnterpriseLogicDatabaseMapOutput struct{ *pulumi.OutputState }

func (EnterpriseLogicDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseLogicDatabase)(nil)).Elem()
}

func (o EnterpriseLogicDatabaseMapOutput) ToEnterpriseLogicDatabaseMapOutput() EnterpriseLogicDatabaseMapOutput {
	return o
}

func (o EnterpriseLogicDatabaseMapOutput) ToEnterpriseLogicDatabaseMapOutputWithContext(ctx context.Context) EnterpriseLogicDatabaseMapOutput {
	return o
}

func (o EnterpriseLogicDatabaseMapOutput) MapIndex(k pulumi.StringInput) EnterpriseLogicDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseLogicDatabase {
		return vs[0].(map[string]*EnterpriseLogicDatabase)[vs[1].(string)]
	}).(EnterpriseLogicDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLogicDatabaseInput)(nil)).Elem(), &EnterpriseLogicDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLogicDatabaseArrayInput)(nil)).Elem(), EnterpriseLogicDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseLogicDatabaseMapInput)(nil)).Elem(), EnterpriseLogicDatabaseMap{})
	pulumi.RegisterOutputType(EnterpriseLogicDatabaseOutput{})
	pulumi.RegisterOutputType(EnterpriseLogicDatabaseArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseLogicDatabaseMapOutput{})
}
