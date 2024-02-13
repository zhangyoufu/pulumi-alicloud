// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OTS secondary index resource.
//
// For information about OTS secondary index and how to use it, see [Secondary index overview](https://www.alibabacloud.com/help/en/tablestore/latest/secondary-index-overview).
//
// > **NOTE:** Available since v1.187.0.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewRandomInteger(ctx, "defaultRandomInteger", &random.RandomIntegerArgs{
//				Min: pulumi.Int(10000),
//				Max: pulumi.Int(99999),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := ots.NewInstance(ctx, "defaultInstance", &ots.InstanceArgs{
//				Description: pulumi.String(name),
//				AccessedBy:  pulumi.String("Any"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultTable, err := ots.NewTable(ctx, "defaultTable", &ots.TableArgs{
//				InstanceName: defaultInstance.Name,
//				TableName:    pulumi.String("tf_example"),
//				TimeToLive:   -1,
//				MaxVersion:   pulumi.Int(1),
//				EnableSse:    pulumi.Bool(true),
//				SseKeyType:   pulumi.String("SSE_KMS_SERVICE"),
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
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ots.NewSecondaryIndex(ctx, "defaultSecondaryIndex", &ots.SecondaryIndexArgs{
//				InstanceName:    defaultInstance.Name,
//				TableName:       defaultTable.TableName,
//				IndexName:       pulumi.String("example_index"),
//				IndexType:       pulumi.String("Global"),
//				IncludeBaseData: pulumi.Bool(true),
//				PrimaryKeys: pulumi.StringArray{
//					pulumi.String("pk1"),
//					pulumi.String("pk2"),
//					pulumi.String("pk3"),
//				},
//				DefinedColumns: pulumi.StringArray{
//					pulumi.String("col1"),
//					pulumi.String("col2"),
//					pulumi.String("col3"),
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
// OTS secondary index can be imported using id, e.g.
//
// ```sh
// $ pulumi import alicloud:ots/secondaryIndex:SecondaryIndex index1 <instance_name>:<table_name>:<index_name>:<index_type>
// ```
type SecondaryIndex struct {
	pulumi.CustomResourceState

	// A list of defined column for index, referenced from Table's primary keys or predefined columns.
	DefinedColumns pulumi.StringArrayOutput `pulumi:"definedColumns"`
	// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
	IncludeBaseData pulumi.BoolOutput `pulumi:"includeBaseData"`
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringOutput `pulumi:"indexName"`
	// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
	IndexType pulumi.StringOutput `pulumi:"indexType"`
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
	PrimaryKeys pulumi.StringArrayOutput `pulumi:"primaryKeys"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringOutput `pulumi:"tableName"`
}

// NewSecondaryIndex registers a new resource with the given unique name, arguments, and options.
func NewSecondaryIndex(ctx *pulumi.Context,
	name string, args *SecondaryIndexArgs, opts ...pulumi.ResourceOption) (*SecondaryIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncludeBaseData == nil {
		return nil, errors.New("invalid value for required argument 'IncludeBaseData'")
	}
	if args.IndexName == nil {
		return nil, errors.New("invalid value for required argument 'IndexName'")
	}
	if args.IndexType == nil {
		return nil, errors.New("invalid value for required argument 'IndexType'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.PrimaryKeys == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryKeys'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecondaryIndex
	err := ctx.RegisterResource("alicloud:ots/secondaryIndex:SecondaryIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecondaryIndex gets an existing SecondaryIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecondaryIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecondaryIndexState, opts ...pulumi.ResourceOption) (*SecondaryIndex, error) {
	var resource SecondaryIndex
	err := ctx.ReadResource("alicloud:ots/secondaryIndex:SecondaryIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecondaryIndex resources.
type secondaryIndexState struct {
	// A list of defined column for index, referenced from Table's primary keys or predefined columns.
	DefinedColumns []string `pulumi:"definedColumns"`
	// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
	IncludeBaseData *bool `pulumi:"includeBaseData"`
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName *string `pulumi:"indexName"`
	// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
	IndexType *string `pulumi:"indexType"`
	// The name of the OTS instance in which table will located.
	InstanceName *string `pulumi:"instanceName"`
	// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
	PrimaryKeys []string `pulumi:"primaryKeys"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName *string `pulumi:"tableName"`
}

type SecondaryIndexState struct {
	// A list of defined column for index, referenced from Table's primary keys or predefined columns.
	DefinedColumns pulumi.StringArrayInput
	// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
	IncludeBaseData pulumi.BoolPtrInput
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringPtrInput
	// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
	IndexType pulumi.StringPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringPtrInput
	// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
	PrimaryKeys pulumi.StringArrayInput
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringPtrInput
}

func (SecondaryIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*secondaryIndexState)(nil)).Elem()
}

type secondaryIndexArgs struct {
	// A list of defined column for index, referenced from Table's primary keys or predefined columns.
	DefinedColumns []string `pulumi:"definedColumns"`
	// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
	IncludeBaseData bool `pulumi:"includeBaseData"`
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName string `pulumi:"indexName"`
	// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
	IndexType string `pulumi:"indexType"`
	// The name of the OTS instance in which table will located.
	InstanceName string `pulumi:"instanceName"`
	// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
	PrimaryKeys []string `pulumi:"primaryKeys"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a SecondaryIndex resource.
type SecondaryIndexArgs struct {
	// A list of defined column for index, referenced from Table's primary keys or predefined columns.
	DefinedColumns pulumi.StringArrayInput
	// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
	IncludeBaseData pulumi.BoolInput
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringInput
	// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
	IndexType pulumi.StringInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringInput
	// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
	PrimaryKeys pulumi.StringArrayInput
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringInput
}

func (SecondaryIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secondaryIndexArgs)(nil)).Elem()
}

type SecondaryIndexInput interface {
	pulumi.Input

	ToSecondaryIndexOutput() SecondaryIndexOutput
	ToSecondaryIndexOutputWithContext(ctx context.Context) SecondaryIndexOutput
}

func (*SecondaryIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**SecondaryIndex)(nil)).Elem()
}

func (i *SecondaryIndex) ToSecondaryIndexOutput() SecondaryIndexOutput {
	return i.ToSecondaryIndexOutputWithContext(context.Background())
}

func (i *SecondaryIndex) ToSecondaryIndexOutputWithContext(ctx context.Context) SecondaryIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIndexOutput)
}

// SecondaryIndexArrayInput is an input type that accepts SecondaryIndexArray and SecondaryIndexArrayOutput values.
// You can construct a concrete instance of `SecondaryIndexArrayInput` via:
//
//	SecondaryIndexArray{ SecondaryIndexArgs{...} }
type SecondaryIndexArrayInput interface {
	pulumi.Input

	ToSecondaryIndexArrayOutput() SecondaryIndexArrayOutput
	ToSecondaryIndexArrayOutputWithContext(context.Context) SecondaryIndexArrayOutput
}

type SecondaryIndexArray []SecondaryIndexInput

func (SecondaryIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecondaryIndex)(nil)).Elem()
}

func (i SecondaryIndexArray) ToSecondaryIndexArrayOutput() SecondaryIndexArrayOutput {
	return i.ToSecondaryIndexArrayOutputWithContext(context.Background())
}

func (i SecondaryIndexArray) ToSecondaryIndexArrayOutputWithContext(ctx context.Context) SecondaryIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIndexArrayOutput)
}

// SecondaryIndexMapInput is an input type that accepts SecondaryIndexMap and SecondaryIndexMapOutput values.
// You can construct a concrete instance of `SecondaryIndexMapInput` via:
//
//	SecondaryIndexMap{ "key": SecondaryIndexArgs{...} }
type SecondaryIndexMapInput interface {
	pulumi.Input

	ToSecondaryIndexMapOutput() SecondaryIndexMapOutput
	ToSecondaryIndexMapOutputWithContext(context.Context) SecondaryIndexMapOutput
}

type SecondaryIndexMap map[string]SecondaryIndexInput

func (SecondaryIndexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecondaryIndex)(nil)).Elem()
}

func (i SecondaryIndexMap) ToSecondaryIndexMapOutput() SecondaryIndexMapOutput {
	return i.ToSecondaryIndexMapOutputWithContext(context.Background())
}

func (i SecondaryIndexMap) ToSecondaryIndexMapOutputWithContext(ctx context.Context) SecondaryIndexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecondaryIndexMapOutput)
}

type SecondaryIndexOutput struct{ *pulumi.OutputState }

func (SecondaryIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecondaryIndex)(nil)).Elem()
}

func (o SecondaryIndexOutput) ToSecondaryIndexOutput() SecondaryIndexOutput {
	return o
}

func (o SecondaryIndexOutput) ToSecondaryIndexOutputWithContext(ctx context.Context) SecondaryIndexOutput {
	return o
}

// A list of defined column for index, referenced from Table's primary keys or predefined columns.
func (o SecondaryIndexOutput) DefinedColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringArrayOutput { return v.DefinedColumns }).(pulumi.StringArrayOutput)
}

// whether the index contains data that already exists in the data table. When includeBaseData is set to true, it means that stock data is included.
func (o SecondaryIndexOutput) IncludeBaseData() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.BoolOutput { return v.IncludeBaseData }).(pulumi.BoolOutput)
}

// The index name of the OTS Table. If changed, a new index would be created.
func (o SecondaryIndexOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringOutput { return v.IndexName }).(pulumi.StringOutput)
}

// The index type of the OTS Table. If changed, a new index would be created, only `Global` or `Local` is allowed.
func (o SecondaryIndexOutput) IndexType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringOutput { return v.IndexType }).(pulumi.StringOutput)
}

// The name of the OTS instance in which table will located.
func (o SecondaryIndexOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// A list of primary keys for index, referenced from Table's primary keys or predefined columns.
func (o SecondaryIndexOutput) PrimaryKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringArrayOutput { return v.PrimaryKeys }).(pulumi.StringArrayOutput)
}

// The name of the OTS table. If changed, a new table would be created.
func (o SecondaryIndexOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecondaryIndex) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

type SecondaryIndexArrayOutput struct{ *pulumi.OutputState }

func (SecondaryIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecondaryIndex)(nil)).Elem()
}

func (o SecondaryIndexArrayOutput) ToSecondaryIndexArrayOutput() SecondaryIndexArrayOutput {
	return o
}

func (o SecondaryIndexArrayOutput) ToSecondaryIndexArrayOutputWithContext(ctx context.Context) SecondaryIndexArrayOutput {
	return o
}

func (o SecondaryIndexArrayOutput) Index(i pulumi.IntInput) SecondaryIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecondaryIndex {
		return vs[0].([]*SecondaryIndex)[vs[1].(int)]
	}).(SecondaryIndexOutput)
}

type SecondaryIndexMapOutput struct{ *pulumi.OutputState }

func (SecondaryIndexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecondaryIndex)(nil)).Elem()
}

func (o SecondaryIndexMapOutput) ToSecondaryIndexMapOutput() SecondaryIndexMapOutput {
	return o
}

func (o SecondaryIndexMapOutput) ToSecondaryIndexMapOutputWithContext(ctx context.Context) SecondaryIndexMapOutput {
	return o
}

func (o SecondaryIndexMapOutput) MapIndex(k pulumi.StringInput) SecondaryIndexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecondaryIndex {
		return vs[0].(map[string]*SecondaryIndex)[vs[1].(string)]
	}).(SecondaryIndexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIndexInput)(nil)).Elem(), &SecondaryIndex{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIndexArrayInput)(nil)).Elem(), SecondaryIndexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecondaryIndexMapInput)(nil)).Elem(), SecondaryIndexMap{})
	pulumi.RegisterOutputType(SecondaryIndexOutput{})
	pulumi.RegisterOutputType(SecondaryIndexArrayOutput{})
	pulumi.RegisterOutputType(SecondaryIndexMapOutput{})
}
