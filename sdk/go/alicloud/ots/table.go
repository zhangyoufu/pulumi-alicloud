// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an OTS table resource.
//
// > **NOTE:** From Provider version 1.10.0, the provider field 'ots_instance_name' has been deprecated and
// you should use resource alicloud_ots_table's new field 'instance_name' and 'table_name' to re-import this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraformtest"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			foo, err := ots.NewInstance(ctx, "foo", &ots.InstanceArgs{
//				Description: pulumi.String(name),
//				AccessedBy:  pulumi.String("Any"),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("acceptance test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ots.NewTable(ctx, "basic", &ots.TableArgs{
//				InstanceName: foo.Name,
//				TableName:    pulumi.String(name),
//				PrimaryKeys: ots.TablePrimaryKeyArray{
//					&ots.TablePrimaryKeyArgs{
//						Name: pulumi.String("pk1"),
//						Type: pulumi.String("Integer"),
//					},
//					&ots.TablePrimaryKeyArgs{
//						Name: pulumi.String("pk2"),
//						Type: pulumi.String("String"),
//					},
//					&ots.TablePrimaryKeyArgs{
//						Name: pulumi.String("pk3"),
//						Type: pulumi.String("Binary"),
//					},
//				},
//				DefinedColumns: ots.TableDefinedColumnArray{
//					&ots.TableDefinedColumnArgs{
//						Name: pulumi.String("col1"),
//						Type: pulumi.String("Integer"),
//					},
//					&ots.TableDefinedColumnArgs{
//						Name: pulumi.String("col2"),
//						Type: pulumi.String("String"),
//					},
//					&ots.TableDefinedColumnArgs{
//						Name: pulumi.String("col3"),
//						Type: pulumi.String("Binary"),
//					},
//				},
//				TimeToLive:                -1,
//				MaxVersion:                pulumi.Int(1),
//				DeviationCellVersionInSec: pulumi.String("1"),
//				EnableSse:                 pulumi.Bool(true),
//				SseKeyType:                pulumi.String("SSE_KMS_SERVICE"),
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
// OTS table can be imported using id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ots/table:Table table my-ots:ots_table
//
// ```
type Table struct {
	pulumi.CustomResourceState

	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
	DefinedColumns TableDefinedColumnArrayOutput `pulumi:"definedColumns"`
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrOutput `pulumi:"deviationCellVersionInSec"`
	// Whether enable OTS server side encryption. Default value is false.
	EnableSse pulumi.BoolPtrOutput `pulumi:"enableSse"`
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntOutput `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayOutput `pulumi:"primaryKeys"`
	// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
	SseKeyType pulumi.StringPtrOutput `pulumi:"sseKeyType"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntOutput `pulumi:"timeToLive"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.MaxVersion == nil {
		return nil, errors.New("invalid value for required argument 'MaxVersion'")
	}
	if args.PrimaryKeys == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryKeys'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	if args.TimeToLive == nil {
		return nil, errors.New("invalid value for required argument 'TimeToLive'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("alicloud:ots/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("alicloud:ots/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
	DefinedColumns []TableDefinedColumn `pulumi:"definedColumns"`
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec *string `pulumi:"deviationCellVersionInSec"`
	// Whether enable OTS server side encryption. Default value is false.
	EnableSse *bool `pulumi:"enableSse"`
	// The name of the OTS instance in which table will located.
	InstanceName *string `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion *int `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys []TablePrimaryKey `pulumi:"primaryKeys"`
	// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
	SseKeyType *string `pulumi:"sseKeyType"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName *string `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive *int `pulumi:"timeToLive"`
}

type TableState struct {
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
	DefinedColumns TableDefinedColumnArrayInput
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrInput
	// Whether enable OTS server side encryption. Default value is false.
	EnableSse pulumi.BoolPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringPtrInput
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntPtrInput
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayInput
	// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
	SseKeyType pulumi.StringPtrInput
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringPtrInput
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
	DefinedColumns []TableDefinedColumn `pulumi:"definedColumns"`
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec *string `pulumi:"deviationCellVersionInSec"`
	// Whether enable OTS server side encryption. Default value is false.
	EnableSse *bool `pulumi:"enableSse"`
	// The name of the OTS instance in which table will located.
	InstanceName string `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion int `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys []TablePrimaryKey `pulumi:"primaryKeys"`
	// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
	SseKeyType *string `pulumi:"sseKeyType"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName string `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive int `pulumi:"timeToLive"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
	DefinedColumns TableDefinedColumnArrayInput
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrInput
	// Whether enable OTS server side encryption. Default value is false.
	EnableSse pulumi.BoolPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringInput
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntInput
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayInput
	// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
	SseKeyType pulumi.StringPtrInput
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringInput
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: i.ToTableOutputWithContext(ctx).OutputState,
	}
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//	TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

func (i TableArray) ToOutput(ctx context.Context) pulumix.Output[[]*Table] {
	return pulumix.Output[[]*Table]{
		OutputState: i.ToTableArrayOutputWithContext(ctx).OutputState,
	}
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//	TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

func (i TableMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Table] {
	return pulumix.Output[map[string]*Table]{
		OutputState: i.ToTableMapOutputWithContext(ctx).OutputState,
	}
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: o.OutputState,
	}
}

// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of defined column. The number of `definedColumn` should not be more than 32.
func (o TableOutput) DefinedColumns() TableDefinedColumnArrayOutput {
	return o.ApplyT(func(v *Table) TableDefinedColumnArrayOutput { return v.DefinedColumns }).(TableDefinedColumnArrayOutput)
}

// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
func (o TableOutput) DeviationCellVersionInSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.DeviationCellVersionInSec }).(pulumi.StringPtrOutput)
}

// Whether enable OTS server side encryption. Default value is false.
func (o TableOutput) EnableSse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.BoolPtrOutput { return v.EnableSse }).(pulumi.BoolPtrOutput)
}

// The name of the OTS instance in which table will located.
func (o TableOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The maximum number of versions stored in this table. The valid value is 1-2147483647.
func (o TableOutput) MaxVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.MaxVersion }).(pulumi.IntOutput)
}

// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
func (o TableOutput) PrimaryKeys() TablePrimaryKeyArrayOutput {
	return o.ApplyT(func(v *Table) TablePrimaryKeyArrayOutput { return v.PrimaryKeys }).(TablePrimaryKeyArrayOutput)
}

// The key type of OTS server side encryption. Only `SSE_KMS_SERVICE` is allowed.
func (o TableOutput) SseKeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.SseKeyType }).(pulumi.StringPtrOutput)
}

// The table name of the OTS instance. If changed, a new table would be created.
func (o TableOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
func (o TableOutput) TimeToLive() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.TimeToLive }).(pulumi.IntOutput)
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Table] {
	return pulumix.Output[[]*Table]{
		OutputState: o.OutputState,
	}
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Table {
		return vs[0].([]*Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Table] {
	return pulumix.Output[map[string]*Table]{
		OutputState: o.OutputState,
	}
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Table {
		return vs[0].(map[string]*Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableArrayInput)(nil)).Elem(), TableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableMapInput)(nil)).Elem(), TableMap{})
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}
