// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sddp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Data Security Center Data Limit resource.
//
// For information about Data Security Center Data Limit and how to use it, see [What is Data Limit](https://www.alibabacloud.com/help/en/doc-detail/158987.html).
//
// > **NOTE:** Available in v1.159.0+.
//
// ## Import
//
// Data Security Center Data Limit can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sddp/dataLimit:DataLimit example <id>
//
// ```
type DataLimit struct {
	pulumi.CustomResourceState

	// Whether to enable the log auditing feature. Valid values: `0`, `1`.
	AuditStatus pulumi.IntOutput `pulumi:"auditStatus"`
	// The type of the database. Valid values: `MySQL`, `SQLServer`.
	EngineType pulumi.StringPtrOutput `pulumi:"engineType"`
	// The lang.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
	LogStoreDay pulumi.IntPtrOutput `pulumi:"logStoreDay"`
	// The ID of the data asset.
	ParentId pulumi.StringPtrOutput `pulumi:"parentId"`
	// The password that is used to connect to the database.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The port that is used to connect to the database.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The region ID of the data asset.
	ServiceRegionId pulumi.StringPtrOutput `pulumi:"serviceRegionId"`
	// The name of the service to which the data asset belongs.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewDataLimit registers a new resource with the given unique name, arguments, and options.
func NewDataLimit(ctx *pulumi.Context,
	name string, args *DataLimitArgs, opts ...pulumi.ResourceOption) (*DataLimit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	var resource DataLimit
	err := ctx.RegisterResource("alicloud:sddp/dataLimit:DataLimit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLimit gets an existing DataLimit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLimit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLimitState, opts ...pulumi.ResourceOption) (*DataLimit, error) {
	var resource DataLimit
	err := ctx.ReadResource("alicloud:sddp/dataLimit:DataLimit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLimit resources.
type dataLimitState struct {
	// Whether to enable the log auditing feature. Valid values: `0`, `1`.
	AuditStatus *int `pulumi:"auditStatus"`
	// The type of the database. Valid values: `MySQL`, `SQLServer`.
	EngineType *string `pulumi:"engineType"`
	// The lang.
	Lang *string `pulumi:"lang"`
	// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
	LogStoreDay *int `pulumi:"logStoreDay"`
	// The ID of the data asset.
	ParentId *string `pulumi:"parentId"`
	// The password that is used to connect to the database.
	Password *string `pulumi:"password"`
	// The port that is used to connect to the database.
	Port *int `pulumi:"port"`
	// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
	ResourceType *string `pulumi:"resourceType"`
	// The region ID of the data asset.
	ServiceRegionId *string `pulumi:"serviceRegionId"`
	// The name of the service to which the data asset belongs.
	UserName *string `pulumi:"userName"`
}

type DataLimitState struct {
	// Whether to enable the log auditing feature. Valid values: `0`, `1`.
	AuditStatus pulumi.IntPtrInput
	// The type of the database. Valid values: `MySQL`, `SQLServer`.
	EngineType pulumi.StringPtrInput
	// The lang.
	Lang pulumi.StringPtrInput
	// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
	LogStoreDay pulumi.IntPtrInput
	// The ID of the data asset.
	ParentId pulumi.StringPtrInput
	// The password that is used to connect to the database.
	Password pulumi.StringPtrInput
	// The port that is used to connect to the database.
	Port pulumi.IntPtrInput
	// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
	ResourceType pulumi.StringPtrInput
	// The region ID of the data asset.
	ServiceRegionId pulumi.StringPtrInput
	// The name of the service to which the data asset belongs.
	UserName pulumi.StringPtrInput
}

func (DataLimitState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLimitState)(nil)).Elem()
}

type dataLimitArgs struct {
	// Whether to enable the log auditing feature. Valid values: `0`, `1`.
	AuditStatus *int `pulumi:"auditStatus"`
	// The type of the database. Valid values: `MySQL`, `SQLServer`.
	EngineType *string `pulumi:"engineType"`
	// The lang.
	Lang *string `pulumi:"lang"`
	// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
	LogStoreDay *int `pulumi:"logStoreDay"`
	// The ID of the data asset.
	ParentId *string `pulumi:"parentId"`
	// The password that is used to connect to the database.
	Password *string `pulumi:"password"`
	// The port that is used to connect to the database.
	Port *int `pulumi:"port"`
	// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
	ResourceType string `pulumi:"resourceType"`
	// The region ID of the data asset.
	ServiceRegionId *string `pulumi:"serviceRegionId"`
	// The name of the service to which the data asset belongs.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a DataLimit resource.
type DataLimitArgs struct {
	// Whether to enable the log auditing feature. Valid values: `0`, `1`.
	AuditStatus pulumi.IntPtrInput
	// The type of the database. Valid values: `MySQL`, `SQLServer`.
	EngineType pulumi.StringPtrInput
	// The lang.
	Lang pulumi.StringPtrInput
	// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
	LogStoreDay pulumi.IntPtrInput
	// The ID of the data asset.
	ParentId pulumi.StringPtrInput
	// The password that is used to connect to the database.
	Password pulumi.StringPtrInput
	// The port that is used to connect to the database.
	Port pulumi.IntPtrInput
	// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
	ResourceType pulumi.StringInput
	// The region ID of the data asset.
	ServiceRegionId pulumi.StringPtrInput
	// The name of the service to which the data asset belongs.
	UserName pulumi.StringPtrInput
}

func (DataLimitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLimitArgs)(nil)).Elem()
}

type DataLimitInput interface {
	pulumi.Input

	ToDataLimitOutput() DataLimitOutput
	ToDataLimitOutputWithContext(ctx context.Context) DataLimitOutput
}

func (*DataLimit) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLimit)(nil)).Elem()
}

func (i *DataLimit) ToDataLimitOutput() DataLimitOutput {
	return i.ToDataLimitOutputWithContext(context.Background())
}

func (i *DataLimit) ToDataLimitOutputWithContext(ctx context.Context) DataLimitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLimitOutput)
}

// DataLimitArrayInput is an input type that accepts DataLimitArray and DataLimitArrayOutput values.
// You can construct a concrete instance of `DataLimitArrayInput` via:
//
//	DataLimitArray{ DataLimitArgs{...} }
type DataLimitArrayInput interface {
	pulumi.Input

	ToDataLimitArrayOutput() DataLimitArrayOutput
	ToDataLimitArrayOutputWithContext(context.Context) DataLimitArrayOutput
}

type DataLimitArray []DataLimitInput

func (DataLimitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLimit)(nil)).Elem()
}

func (i DataLimitArray) ToDataLimitArrayOutput() DataLimitArrayOutput {
	return i.ToDataLimitArrayOutputWithContext(context.Background())
}

func (i DataLimitArray) ToDataLimitArrayOutputWithContext(ctx context.Context) DataLimitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLimitArrayOutput)
}

// DataLimitMapInput is an input type that accepts DataLimitMap and DataLimitMapOutput values.
// You can construct a concrete instance of `DataLimitMapInput` via:
//
//	DataLimitMap{ "key": DataLimitArgs{...} }
type DataLimitMapInput interface {
	pulumi.Input

	ToDataLimitMapOutput() DataLimitMapOutput
	ToDataLimitMapOutputWithContext(context.Context) DataLimitMapOutput
}

type DataLimitMap map[string]DataLimitInput

func (DataLimitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLimit)(nil)).Elem()
}

func (i DataLimitMap) ToDataLimitMapOutput() DataLimitMapOutput {
	return i.ToDataLimitMapOutputWithContext(context.Background())
}

func (i DataLimitMap) ToDataLimitMapOutputWithContext(ctx context.Context) DataLimitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLimitMapOutput)
}

type DataLimitOutput struct{ *pulumi.OutputState }

func (DataLimitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLimit)(nil)).Elem()
}

func (o DataLimitOutput) ToDataLimitOutput() DataLimitOutput {
	return o
}

func (o DataLimitOutput) ToDataLimitOutputWithContext(ctx context.Context) DataLimitOutput {
	return o
}

// Whether to enable the log auditing feature. Valid values: `0`, `1`.
func (o DataLimitOutput) AuditStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.IntOutput { return v.AuditStatus }).(pulumi.IntOutput)
}

// The type of the database. Valid values: `MySQL`, `SQLServer`.
func (o DataLimitOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.EngineType }).(pulumi.StringPtrOutput)
}

// The lang.
func (o DataLimitOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.Lang }).(pulumi.StringPtrOutput)
}

// The retention period of raw logs after you enable the log auditing feature. Unit: day. Valid values: `180`, `30`, `365`, `90`. **NOTE:** The`logStoreDay` is valid when the `auditStatus` is `1`.
func (o DataLimitOutput) LogStoreDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.IntPtrOutput { return v.LogStoreDay }).(pulumi.IntPtrOutput)
}

// The ID of the data asset.
func (o DataLimitOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.ParentId }).(pulumi.StringPtrOutput)
}

// The password that is used to connect to the database.
func (o DataLimitOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The port that is used to connect to the database.
func (o DataLimitOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The type of the service to which the data asset belongs. Valid values: `MaxCompute`, `OSS`, `RDS`.
func (o DataLimitOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The region ID of the data asset.
func (o DataLimitOutput) ServiceRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.ServiceRegionId }).(pulumi.StringPtrOutput)
}

// The name of the service to which the data asset belongs.
func (o DataLimitOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLimit) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

type DataLimitArrayOutput struct{ *pulumi.OutputState }

func (DataLimitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLimit)(nil)).Elem()
}

func (o DataLimitArrayOutput) ToDataLimitArrayOutput() DataLimitArrayOutput {
	return o
}

func (o DataLimitArrayOutput) ToDataLimitArrayOutputWithContext(ctx context.Context) DataLimitArrayOutput {
	return o
}

func (o DataLimitArrayOutput) Index(i pulumi.IntInput) DataLimitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataLimit {
		return vs[0].([]*DataLimit)[vs[1].(int)]
	}).(DataLimitOutput)
}

type DataLimitMapOutput struct{ *pulumi.OutputState }

func (DataLimitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLimit)(nil)).Elem()
}

func (o DataLimitMapOutput) ToDataLimitMapOutput() DataLimitMapOutput {
	return o
}

func (o DataLimitMapOutput) ToDataLimitMapOutputWithContext(ctx context.Context) DataLimitMapOutput {
	return o
}

func (o DataLimitMapOutput) MapIndex(k pulumi.StringInput) DataLimitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataLimit {
		return vs[0].(map[string]*DataLimit)[vs[1].(string)]
	}).(DataLimitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataLimitInput)(nil)).Elem(), &DataLimit{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLimitArrayInput)(nil)).Elem(), DataLimitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLimitMapInput)(nil)).Elem(), DataLimitMap{})
	pulumi.RegisterOutputType(DataLimitOutput{})
	pulumi.RegisterOutputType(DataLimitArrayOutput{})
	pulumi.RegisterOutputType(DataLimitMapOutput{})
}
