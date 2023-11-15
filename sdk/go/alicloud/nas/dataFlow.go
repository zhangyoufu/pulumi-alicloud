// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Network Attached Storage (NAS) Data Flow resource.
//
// For information about Network Attached Storage (NAS) Data Flow and how to use it, see [What is Data Flow](https://www.alibabacloud.com/help/en/doc-detail/27530.html).
//
// > **NOTE:** Available since v1.153.0.
//
// ## Import
//
// Network Attached Storage (NAS) Data Flow can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:nas/dataFlow:DataFlow example <file_system_id>:<data_flow_id>
//
// ```
type DataFlow struct {
	pulumi.CustomResourceState

	// The ID of the Data flow.
	DataFlowId pulumi.StringOutput `pulumi:"dataFlowId"`
	// The Description of the data flow. Restrictions:
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The ID of the Fileset.
	FsetId pulumi.StringOutput `pulumi:"fsetId"`
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
	SourceSecurityType pulumi.StringOutput `pulumi:"sourceSecurityType"`
	// The access path of the source store. Format: `<storage type>://<path>`. Among them:
	// - storage type: currently only OSS is supported.
	// - path: the bucket name of OSS.
	// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
	SourceStorage pulumi.StringOutput `pulumi:"sourceStorage"`
	// The status of the Data flow. Valid values: `Running`, `Stopped`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewDataFlow registers a new resource with the given unique name, arguments, and options.
func NewDataFlow(ctx *pulumi.Context,
	name string, args *DataFlowArgs, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.FsetId == nil {
		return nil, errors.New("invalid value for required argument 'FsetId'")
	}
	if args.SourceStorage == nil {
		return nil, errors.New("invalid value for required argument 'SourceStorage'")
	}
	if args.Throughput == nil {
		return nil, errors.New("invalid value for required argument 'Throughput'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataFlow
	err := ctx.RegisterResource("alicloud:nas/dataFlow:DataFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataFlow gets an existing DataFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFlowState, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	var resource DataFlow
	err := ctx.ReadResource("alicloud:nas/dataFlow:DataFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataFlow resources.
type dataFlowState struct {
	// The ID of the Data flow.
	DataFlowId *string `pulumi:"dataFlowId"`
	// The Description of the data flow. Restrictions:
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The ID of the Fileset.
	FsetId *string `pulumi:"fsetId"`
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
	SourceSecurityType *string `pulumi:"sourceSecurityType"`
	// The access path of the source store. Format: `<storage type>://<path>`. Among them:
	// - storage type: currently only OSS is supported.
	// - path: the bucket name of OSS.
	// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
	SourceStorage *string `pulumi:"sourceStorage"`
	// The status of the Data flow. Valid values: `Running`, `Stopped`.
	Status *string `pulumi:"status"`
	// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
	Throughput *int `pulumi:"throughput"`
}

type DataFlowState struct {
	// The ID of the Data flow.
	DataFlowId pulumi.StringPtrInput
	// The Description of the data flow. Restrictions:
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the file system.
	FileSystemId pulumi.StringPtrInput
	// The ID of the Fileset.
	FsetId pulumi.StringPtrInput
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
	SourceSecurityType pulumi.StringPtrInput
	// The access path of the source store. Format: `<storage type>://<path>`. Among them:
	// - storage type: currently only OSS is supported.
	// - path: the bucket name of OSS.
	// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
	SourceStorage pulumi.StringPtrInput
	// The status of the Data flow. Valid values: `Running`, `Stopped`.
	Status pulumi.StringPtrInput
	// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
	Throughput pulumi.IntPtrInput
}

func (DataFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowState)(nil)).Elem()
}

type dataFlowArgs struct {
	// The Description of the data flow. Restrictions:
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The ID of the file system.
	FileSystemId string `pulumi:"fileSystemId"`
	// The ID of the Fileset.
	FsetId string `pulumi:"fsetId"`
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
	SourceSecurityType *string `pulumi:"sourceSecurityType"`
	// The access path of the source store. Format: `<storage type>://<path>`. Among them:
	// - storage type: currently only OSS is supported.
	// - path: the bucket name of OSS.
	// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
	SourceStorage string `pulumi:"sourceStorage"`
	// The status of the Data flow. Valid values: `Running`, `Stopped`.
	Status *string `pulumi:"status"`
	// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
	Throughput int `pulumi:"throughput"`
}

// The set of arguments for constructing a DataFlow resource.
type DataFlowArgs struct {
	// The Description of the data flow. Restrictions:
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The ID of the file system.
	FileSystemId pulumi.StringInput
	// The ID of the Fileset.
	FsetId pulumi.StringInput
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
	SourceSecurityType pulumi.StringPtrInput
	// The access path of the source store. Format: `<storage type>://<path>`. Among them:
	// - storage type: currently only OSS is supported.
	// - path: the bucket name of OSS.
	// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
	SourceStorage pulumi.StringInput
	// The status of the Data flow. Valid values: `Running`, `Stopped`.
	Status pulumi.StringPtrInput
	// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
	Throughput pulumi.IntInput
}

func (DataFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowArgs)(nil)).Elem()
}

type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput
}

func (*DataFlow) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlow)(nil)).Elem()
}

func (i *DataFlow) ToDataFlowOutput() DataFlowOutput {
	return i.ToDataFlowOutputWithContext(context.Background())
}

func (i *DataFlow) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowOutput)
}

// DataFlowArrayInput is an input type that accepts DataFlowArray and DataFlowArrayOutput values.
// You can construct a concrete instance of `DataFlowArrayInput` via:
//
//	DataFlowArray{ DataFlowArgs{...} }
type DataFlowArrayInput interface {
	pulumi.Input

	ToDataFlowArrayOutput() DataFlowArrayOutput
	ToDataFlowArrayOutputWithContext(context.Context) DataFlowArrayOutput
}

type DataFlowArray []DataFlowInput

func (DataFlowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataFlow)(nil)).Elem()
}

func (i DataFlowArray) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return i.ToDataFlowArrayOutputWithContext(context.Background())
}

func (i DataFlowArray) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowArrayOutput)
}

// DataFlowMapInput is an input type that accepts DataFlowMap and DataFlowMapOutput values.
// You can construct a concrete instance of `DataFlowMapInput` via:
//
//	DataFlowMap{ "key": DataFlowArgs{...} }
type DataFlowMapInput interface {
	pulumi.Input

	ToDataFlowMapOutput() DataFlowMapOutput
	ToDataFlowMapOutputWithContext(context.Context) DataFlowMapOutput
}

type DataFlowMap map[string]DataFlowInput

func (DataFlowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataFlow)(nil)).Elem()
}

func (i DataFlowMap) ToDataFlowMapOutput() DataFlowMapOutput {
	return i.ToDataFlowMapOutputWithContext(context.Background())
}

func (i DataFlowMap) ToDataFlowMapOutputWithContext(ctx context.Context) DataFlowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowMapOutput)
}

type DataFlowOutput struct{ *pulumi.OutputState }

func (DataFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlow)(nil)).Elem()
}

func (o DataFlowOutput) ToDataFlowOutput() DataFlowOutput {
	return o
}

func (o DataFlowOutput) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return o
}

// The ID of the Data flow.
func (o DataFlowOutput) DataFlowId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.DataFlowId }).(pulumi.StringOutput)
}

// The Description of the data flow. Restrictions:
func (o DataFlowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dry run.
func (o DataFlowOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The ID of the file system.
func (o DataFlowOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The ID of the Fileset.
func (o DataFlowOutput) FsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.FsetId }).(pulumi.StringOutput)
}

// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type of the source storage. Value:
func (o DataFlowOutput) SourceSecurityType() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.SourceSecurityType }).(pulumi.StringOutput)
}

// The access path of the source store. Format: `<storage type>://<path>`. Among them:
// - storage type: currently only OSS is supported.
// - path: the bucket name of OSS.
// - Only lowercase letters, numbers, and dashes (-) are supported and must start and end with lowercase letters or numbers.
func (o DataFlowOutput) SourceStorage() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.SourceStorage }).(pulumi.StringOutput)
}

// The status of the Data flow. Valid values: `Running`, `Stopped`.
func (o DataFlowOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The maximum transmission bandwidth of data flow, unit: `MB/s`. Valid values: `1200`, `1500`, `600`. **NOTE:** The transmission bandwidth of data flow must be less than the IO bandwidth of the file system.
func (o DataFlowOutput) Throughput() pulumi.IntOutput {
	return o.ApplyT(func(v *DataFlow) pulumi.IntOutput { return v.Throughput }).(pulumi.IntOutput)
}

type DataFlowArrayOutput struct{ *pulumi.OutputState }

func (DataFlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataFlow)(nil)).Elem()
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) Index(i pulumi.IntInput) DataFlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataFlow {
		return vs[0].([]*DataFlow)[vs[1].(int)]
	}).(DataFlowOutput)
}

type DataFlowMapOutput struct{ *pulumi.OutputState }

func (DataFlowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataFlow)(nil)).Elem()
}

func (o DataFlowMapOutput) ToDataFlowMapOutput() DataFlowMapOutput {
	return o
}

func (o DataFlowMapOutput) ToDataFlowMapOutputWithContext(ctx context.Context) DataFlowMapOutput {
	return o
}

func (o DataFlowMapOutput) MapIndex(k pulumi.StringInput) DataFlowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataFlow {
		return vs[0].(map[string]*DataFlow)[vs[1].(string)]
	}).(DataFlowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowInput)(nil)).Elem(), &DataFlow{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowArrayInput)(nil)).Elem(), DataFlowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowMapInput)(nil)).Elem(), DataFlowMap{})
	pulumi.RegisterOutputType(DataFlowOutput{})
	pulumi.RegisterOutputType(DataFlowArrayOutput{})
	pulumi.RegisterOutputType(DataFlowMapOutput{})
}
