// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Compute Nest Service Instance resource.
//
// For information about Compute Nest Service Instance and how to use it, see [What is Service Instance](https://help.aliyun.com/document_detail/396194.html).
//
// > **NOTE:** Available in v1.205.0+.
//
// ## Import
//
// Compute Nest Service Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:compute/nestServiceInstance:NestServiceInstance example <id>
//
// ```
type NestServiceInstance struct {
	pulumi.CustomResourceState

	// The order information of cloud market. See the following `Block commodity`.
	Commodity NestServiceInstanceCommodityPtrOutput `pulumi:"commodity"`
	// Whether the service instance has the O&M function. Default value: `false`. Valid values:
	EnableInstanceOps pulumi.BoolOutput `pulumi:"enableInstanceOps"`
	// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
	EnableUserPrometheus pulumi.BoolOutput `pulumi:"enableUserPrometheus"`
	// The configuration of O&M. See the following `Block operationMetadata`.
	OperationMetadata NestServiceInstanceOperationMetadataOutput `pulumi:"operationMetadata"`
	// The parameters entered by the deployment service instance.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The ID of the service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The name of the Service Instance.
	ServiceInstanceName pulumi.StringOutput `pulumi:"serviceInstanceName"`
	// The version of the service.
	ServiceVersion pulumi.StringOutput `pulumi:"serviceVersion"`
	// The name of the specification.
	SpecificationName pulumi.StringPtrOutput `pulumi:"specificationName"`
	// The status of the Service Instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The name of the template.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewNestServiceInstance registers a new resource with the given unique name, arguments, and options.
func NewNestServiceInstance(ctx *pulumi.Context,
	name string, args *NestServiceInstanceArgs, opts ...pulumi.ResourceOption) (*NestServiceInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.ServiceVersion == nil {
		return nil, errors.New("invalid value for required argument 'ServiceVersion'")
	}
	var resource NestServiceInstance
	err := ctx.RegisterResource("alicloud:compute/nestServiceInstance:NestServiceInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNestServiceInstance gets an existing NestServiceInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNestServiceInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NestServiceInstanceState, opts ...pulumi.ResourceOption) (*NestServiceInstance, error) {
	var resource NestServiceInstance
	err := ctx.ReadResource("alicloud:compute/nestServiceInstance:NestServiceInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NestServiceInstance resources.
type nestServiceInstanceState struct {
	// The order information of cloud market. See the following `Block commodity`.
	Commodity *NestServiceInstanceCommodity `pulumi:"commodity"`
	// Whether the service instance has the O&M function. Default value: `false`. Valid values:
	EnableInstanceOps *bool `pulumi:"enableInstanceOps"`
	// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
	EnableUserPrometheus *bool `pulumi:"enableUserPrometheus"`
	// The configuration of O&M. See the following `Block operationMetadata`.
	OperationMetadata *NestServiceInstanceOperationMetadata `pulumi:"operationMetadata"`
	// The parameters entered by the deployment service instance.
	Parameters *string `pulumi:"parameters"`
	// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
	PaymentType *string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the service.
	ServiceId *string `pulumi:"serviceId"`
	// The name of the Service Instance.
	ServiceInstanceName *string `pulumi:"serviceInstanceName"`
	// The version of the service.
	ServiceVersion *string `pulumi:"serviceVersion"`
	// The name of the specification.
	SpecificationName *string `pulumi:"specificationName"`
	// The status of the Service Instance.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the template.
	TemplateName *string `pulumi:"templateName"`
}

type NestServiceInstanceState struct {
	// The order information of cloud market. See the following `Block commodity`.
	Commodity NestServiceInstanceCommodityPtrInput
	// Whether the service instance has the O&M function. Default value: `false`. Valid values:
	EnableInstanceOps pulumi.BoolPtrInput
	// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
	EnableUserPrometheus pulumi.BoolPtrInput
	// The configuration of O&M. See the following `Block operationMetadata`.
	OperationMetadata NestServiceInstanceOperationMetadataPtrInput
	// The parameters entered by the deployment service instance.
	Parameters pulumi.StringPtrInput
	// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
	PaymentType pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the service.
	ServiceId pulumi.StringPtrInput
	// The name of the Service Instance.
	ServiceInstanceName pulumi.StringPtrInput
	// The version of the service.
	ServiceVersion pulumi.StringPtrInput
	// The name of the specification.
	SpecificationName pulumi.StringPtrInput
	// The status of the Service Instance.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The name of the template.
	TemplateName pulumi.StringPtrInput
}

func (NestServiceInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*nestServiceInstanceState)(nil)).Elem()
}

type nestServiceInstanceArgs struct {
	// The order information of cloud market. See the following `Block commodity`.
	Commodity *NestServiceInstanceCommodity `pulumi:"commodity"`
	// Whether the service instance has the O&M function. Default value: `false`. Valid values:
	EnableInstanceOps *bool `pulumi:"enableInstanceOps"`
	// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
	EnableUserPrometheus *bool `pulumi:"enableUserPrometheus"`
	// The configuration of O&M. See the following `Block operationMetadata`.
	OperationMetadata *NestServiceInstanceOperationMetadata `pulumi:"operationMetadata"`
	// The parameters entered by the deployment service instance.
	Parameters *string `pulumi:"parameters"`
	// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
	PaymentType *string `pulumi:"paymentType"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The ID of the service.
	ServiceId string `pulumi:"serviceId"`
	// The name of the Service Instance.
	ServiceInstanceName *string `pulumi:"serviceInstanceName"`
	// The version of the service.
	ServiceVersion string `pulumi:"serviceVersion"`
	// The name of the specification.
	SpecificationName *string `pulumi:"specificationName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the template.
	TemplateName *string `pulumi:"templateName"`
}

// The set of arguments for constructing a NestServiceInstance resource.
type NestServiceInstanceArgs struct {
	// The order information of cloud market. See the following `Block commodity`.
	Commodity NestServiceInstanceCommodityPtrInput
	// Whether the service instance has the O&M function. Default value: `false`. Valid values:
	EnableInstanceOps pulumi.BoolPtrInput
	// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
	EnableUserPrometheus pulumi.BoolPtrInput
	// The configuration of O&M. See the following `Block operationMetadata`.
	OperationMetadata NestServiceInstanceOperationMetadataPtrInput
	// The parameters entered by the deployment service instance.
	Parameters pulumi.StringPtrInput
	// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
	PaymentType pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The ID of the service.
	ServiceId pulumi.StringInput
	// The name of the Service Instance.
	ServiceInstanceName pulumi.StringPtrInput
	// The version of the service.
	ServiceVersion pulumi.StringInput
	// The name of the specification.
	SpecificationName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The name of the template.
	TemplateName pulumi.StringPtrInput
}

func (NestServiceInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nestServiceInstanceArgs)(nil)).Elem()
}

type NestServiceInstanceInput interface {
	pulumi.Input

	ToNestServiceInstanceOutput() NestServiceInstanceOutput
	ToNestServiceInstanceOutputWithContext(ctx context.Context) NestServiceInstanceOutput
}

func (*NestServiceInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**NestServiceInstance)(nil)).Elem()
}

func (i *NestServiceInstance) ToNestServiceInstanceOutput() NestServiceInstanceOutput {
	return i.ToNestServiceInstanceOutputWithContext(context.Background())
}

func (i *NestServiceInstance) ToNestServiceInstanceOutputWithContext(ctx context.Context) NestServiceInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NestServiceInstanceOutput)
}

// NestServiceInstanceArrayInput is an input type that accepts NestServiceInstanceArray and NestServiceInstanceArrayOutput values.
// You can construct a concrete instance of `NestServiceInstanceArrayInput` via:
//
//	NestServiceInstanceArray{ NestServiceInstanceArgs{...} }
type NestServiceInstanceArrayInput interface {
	pulumi.Input

	ToNestServiceInstanceArrayOutput() NestServiceInstanceArrayOutput
	ToNestServiceInstanceArrayOutputWithContext(context.Context) NestServiceInstanceArrayOutput
}

type NestServiceInstanceArray []NestServiceInstanceInput

func (NestServiceInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NestServiceInstance)(nil)).Elem()
}

func (i NestServiceInstanceArray) ToNestServiceInstanceArrayOutput() NestServiceInstanceArrayOutput {
	return i.ToNestServiceInstanceArrayOutputWithContext(context.Background())
}

func (i NestServiceInstanceArray) ToNestServiceInstanceArrayOutputWithContext(ctx context.Context) NestServiceInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NestServiceInstanceArrayOutput)
}

// NestServiceInstanceMapInput is an input type that accepts NestServiceInstanceMap and NestServiceInstanceMapOutput values.
// You can construct a concrete instance of `NestServiceInstanceMapInput` via:
//
//	NestServiceInstanceMap{ "key": NestServiceInstanceArgs{...} }
type NestServiceInstanceMapInput interface {
	pulumi.Input

	ToNestServiceInstanceMapOutput() NestServiceInstanceMapOutput
	ToNestServiceInstanceMapOutputWithContext(context.Context) NestServiceInstanceMapOutput
}

type NestServiceInstanceMap map[string]NestServiceInstanceInput

func (NestServiceInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NestServiceInstance)(nil)).Elem()
}

func (i NestServiceInstanceMap) ToNestServiceInstanceMapOutput() NestServiceInstanceMapOutput {
	return i.ToNestServiceInstanceMapOutputWithContext(context.Background())
}

func (i NestServiceInstanceMap) ToNestServiceInstanceMapOutputWithContext(ctx context.Context) NestServiceInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NestServiceInstanceMapOutput)
}

type NestServiceInstanceOutput struct{ *pulumi.OutputState }

func (NestServiceInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NestServiceInstance)(nil)).Elem()
}

func (o NestServiceInstanceOutput) ToNestServiceInstanceOutput() NestServiceInstanceOutput {
	return o
}

func (o NestServiceInstanceOutput) ToNestServiceInstanceOutputWithContext(ctx context.Context) NestServiceInstanceOutput {
	return o
}

// The order information of cloud market. See the following `Block commodity`.
func (o NestServiceInstanceOutput) Commodity() NestServiceInstanceCommodityPtrOutput {
	return o.ApplyT(func(v *NestServiceInstance) NestServiceInstanceCommodityPtrOutput { return v.Commodity }).(NestServiceInstanceCommodityPtrOutput)
}

// Whether the service instance has the O&M function. Default value: `false`. Valid values:
func (o NestServiceInstanceOutput) EnableInstanceOps() pulumi.BoolOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.BoolOutput { return v.EnableInstanceOps }).(pulumi.BoolOutput)
}

// Whether Prometheus monitoring is enabled. Default value: `false`. Valid values:
func (o NestServiceInstanceOutput) EnableUserPrometheus() pulumi.BoolOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.BoolOutput { return v.EnableUserPrometheus }).(pulumi.BoolOutput)
}

// The configuration of O&M. See the following `Block operationMetadata`.
func (o NestServiceInstanceOutput) OperationMetadata() NestServiceInstanceOperationMetadataOutput {
	return o.ApplyT(func(v *NestServiceInstance) NestServiceInstanceOperationMetadataOutput { return v.OperationMetadata }).(NestServiceInstanceOperationMetadataOutput)
}

// The parameters entered by the deployment service instance.
func (o NestServiceInstanceOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringPtrOutput { return v.Parameters }).(pulumi.StringPtrOutput)
}

// The type of payment. Valid values: `Permanent`, `Subscription`, `PayAsYouGo`, `CustomFixTime`.
func (o NestServiceInstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o NestServiceInstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The ID of the service.
func (o NestServiceInstanceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The name of the Service Instance.
func (o NestServiceInstanceOutput) ServiceInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.ServiceInstanceName }).(pulumi.StringOutput)
}

// The version of the service.
func (o NestServiceInstanceOutput) ServiceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.ServiceVersion }).(pulumi.StringOutput)
}

// The name of the specification.
func (o NestServiceInstanceOutput) SpecificationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringPtrOutput { return v.SpecificationName }).(pulumi.StringPtrOutput)
}

// The status of the Service Instance.
func (o NestServiceInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o NestServiceInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The name of the template.
func (o NestServiceInstanceOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *NestServiceInstance) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type NestServiceInstanceArrayOutput struct{ *pulumi.OutputState }

func (NestServiceInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NestServiceInstance)(nil)).Elem()
}

func (o NestServiceInstanceArrayOutput) ToNestServiceInstanceArrayOutput() NestServiceInstanceArrayOutput {
	return o
}

func (o NestServiceInstanceArrayOutput) ToNestServiceInstanceArrayOutputWithContext(ctx context.Context) NestServiceInstanceArrayOutput {
	return o
}

func (o NestServiceInstanceArrayOutput) Index(i pulumi.IntInput) NestServiceInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NestServiceInstance {
		return vs[0].([]*NestServiceInstance)[vs[1].(int)]
	}).(NestServiceInstanceOutput)
}

type NestServiceInstanceMapOutput struct{ *pulumi.OutputState }

func (NestServiceInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NestServiceInstance)(nil)).Elem()
}

func (o NestServiceInstanceMapOutput) ToNestServiceInstanceMapOutput() NestServiceInstanceMapOutput {
	return o
}

func (o NestServiceInstanceMapOutput) ToNestServiceInstanceMapOutputWithContext(ctx context.Context) NestServiceInstanceMapOutput {
	return o
}

func (o NestServiceInstanceMapOutput) MapIndex(k pulumi.StringInput) NestServiceInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NestServiceInstance {
		return vs[0].(map[string]*NestServiceInstance)[vs[1].(string)]
	}).(NestServiceInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NestServiceInstanceInput)(nil)).Elem(), &NestServiceInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*NestServiceInstanceArrayInput)(nil)).Elem(), NestServiceInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NestServiceInstanceMapInput)(nil)).Elem(), NestServiceInstanceMap{})
	pulumi.RegisterOutputType(NestServiceInstanceOutput{})
	pulumi.RegisterOutputType(NestServiceInstanceArrayOutput{})
	pulumi.RegisterOutputType(NestServiceInstanceMapOutput{})
}
