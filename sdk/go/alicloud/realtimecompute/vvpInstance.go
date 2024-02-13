// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package realtimecompute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Realtime Compute Vvp Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:realtimecompute/vvpInstance:VvpInstance example <id>
// ```
type VvpInstance struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The number of subscription periods. If the payment type is PRE, this parameter is required.
	Duration pulumi.IntPtrOutput `pulumi:"duration"`
	// The payment type of the resource.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The subscription period. If the payment type is PRE, this parameter is required.
	PricingCycle pulumi.StringPtrOutput `pulumi:"pricingCycle"`
	// The resource group to which the newly purchased instance belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Resource specifications. See `resourceSpec` below.
	ResourceSpec VvpInstanceResourceSpecOutput `pulumi:"resourceSpec"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Store information. See `storage` below.
	Storage VvpInstanceStorageOutput `pulumi:"storage"`
	// The tags of the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID of the user.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Virtual Switch ID.
	VswitchIds pulumi.StringArrayOutput `pulumi:"vswitchIds"`
	// The name of the workspace.
	VvpInstanceName pulumi.StringOutput `pulumi:"vvpInstanceName"`
	// The zone ID of the resource.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewVvpInstance registers a new resource with the given unique name, arguments, and options.
func NewVvpInstance(ctx *pulumi.Context,
	name string, args *VvpInstanceArgs, opts ...pulumi.ResourceOption) (*VvpInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PaymentType == nil {
		return nil, errors.New("invalid value for required argument 'PaymentType'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VswitchIds == nil {
		return nil, errors.New("invalid value for required argument 'VswitchIds'")
	}
	if args.VvpInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'VvpInstanceName'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VvpInstance
	err := ctx.RegisterResource("alicloud:realtimecompute/vvpInstance:VvpInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVvpInstance gets an existing VvpInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVvpInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VvpInstanceState, opts ...pulumi.ResourceOption) (*VvpInstance, error) {
	var resource VvpInstance
	err := ctx.ReadResource("alicloud:realtimecompute/vvpInstance:VvpInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VvpInstance resources.
type vvpInstanceState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The number of subscription periods. If the payment type is PRE, this parameter is required.
	Duration *int `pulumi:"duration"`
	// The payment type of the resource.
	PaymentType *string `pulumi:"paymentType"`
	// The subscription period. If the payment type is PRE, this parameter is required.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The resource group to which the newly purchased instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Resource specifications. See `resourceSpec` below.
	ResourceSpec *VvpInstanceResourceSpec `pulumi:"resourceSpec"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Store information. See `storage` below.
	Storage *VvpInstanceStorage `pulumi:"storage"`
	// The tags of the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID of the user.
	VpcId *string `pulumi:"vpcId"`
	// Virtual Switch ID.
	VswitchIds []string `pulumi:"vswitchIds"`
	// The name of the workspace.
	VvpInstanceName *string `pulumi:"vvpInstanceName"`
	// The zone ID of the resource.
	ZoneId *string `pulumi:"zoneId"`
}

type VvpInstanceState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The number of subscription periods. If the payment type is PRE, this parameter is required.
	Duration pulumi.IntPtrInput
	// The payment type of the resource.
	PaymentType pulumi.StringPtrInput
	// The subscription period. If the payment type is PRE, this parameter is required.
	PricingCycle pulumi.StringPtrInput
	// The resource group to which the newly purchased instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// Resource specifications. See `resourceSpec` below.
	ResourceSpec VvpInstanceResourceSpecPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Store information. See `storage` below.
	Storage VvpInstanceStoragePtrInput
	// The tags of the resource.
	Tags pulumi.MapInput
	// The VPC ID of the user.
	VpcId pulumi.StringPtrInput
	// Virtual Switch ID.
	VswitchIds pulumi.StringArrayInput
	// The name of the workspace.
	VvpInstanceName pulumi.StringPtrInput
	// The zone ID of the resource.
	ZoneId pulumi.StringPtrInput
}

func (VvpInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vvpInstanceState)(nil)).Elem()
}

type vvpInstanceArgs struct {
	// The number of subscription periods. If the payment type is PRE, this parameter is required.
	Duration *int `pulumi:"duration"`
	// The payment type of the resource.
	PaymentType string `pulumi:"paymentType"`
	// The subscription period. If the payment type is PRE, this parameter is required.
	PricingCycle *string `pulumi:"pricingCycle"`
	// The resource group to which the newly purchased instance belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Resource specifications. See `resourceSpec` below.
	ResourceSpec *VvpInstanceResourceSpec `pulumi:"resourceSpec"`
	// Store information. See `storage` below.
	Storage VvpInstanceStorage `pulumi:"storage"`
	// The tags of the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID of the user.
	VpcId string `pulumi:"vpcId"`
	// Virtual Switch ID.
	VswitchIds []string `pulumi:"vswitchIds"`
	// The name of the workspace.
	VvpInstanceName string `pulumi:"vvpInstanceName"`
	// The zone ID of the resource.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a VvpInstance resource.
type VvpInstanceArgs struct {
	// The number of subscription periods. If the payment type is PRE, this parameter is required.
	Duration pulumi.IntPtrInput
	// The payment type of the resource.
	PaymentType pulumi.StringInput
	// The subscription period. If the payment type is PRE, this parameter is required.
	PricingCycle pulumi.StringPtrInput
	// The resource group to which the newly purchased instance belongs.
	ResourceGroupId pulumi.StringPtrInput
	// Resource specifications. See `resourceSpec` below.
	ResourceSpec VvpInstanceResourceSpecPtrInput
	// Store information. See `storage` below.
	Storage VvpInstanceStorageInput
	// The tags of the resource.
	Tags pulumi.MapInput
	// The VPC ID of the user.
	VpcId pulumi.StringInput
	// Virtual Switch ID.
	VswitchIds pulumi.StringArrayInput
	// The name of the workspace.
	VvpInstanceName pulumi.StringInput
	// The zone ID of the resource.
	ZoneId pulumi.StringInput
}

func (VvpInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vvpInstanceArgs)(nil)).Elem()
}

type VvpInstanceInput interface {
	pulumi.Input

	ToVvpInstanceOutput() VvpInstanceOutput
	ToVvpInstanceOutputWithContext(ctx context.Context) VvpInstanceOutput
}

func (*VvpInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**VvpInstance)(nil)).Elem()
}

func (i *VvpInstance) ToVvpInstanceOutput() VvpInstanceOutput {
	return i.ToVvpInstanceOutputWithContext(context.Background())
}

func (i *VvpInstance) ToVvpInstanceOutputWithContext(ctx context.Context) VvpInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VvpInstanceOutput)
}

// VvpInstanceArrayInput is an input type that accepts VvpInstanceArray and VvpInstanceArrayOutput values.
// You can construct a concrete instance of `VvpInstanceArrayInput` via:
//
//	VvpInstanceArray{ VvpInstanceArgs{...} }
type VvpInstanceArrayInput interface {
	pulumi.Input

	ToVvpInstanceArrayOutput() VvpInstanceArrayOutput
	ToVvpInstanceArrayOutputWithContext(context.Context) VvpInstanceArrayOutput
}

type VvpInstanceArray []VvpInstanceInput

func (VvpInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VvpInstance)(nil)).Elem()
}

func (i VvpInstanceArray) ToVvpInstanceArrayOutput() VvpInstanceArrayOutput {
	return i.ToVvpInstanceArrayOutputWithContext(context.Background())
}

func (i VvpInstanceArray) ToVvpInstanceArrayOutputWithContext(ctx context.Context) VvpInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VvpInstanceArrayOutput)
}

// VvpInstanceMapInput is an input type that accepts VvpInstanceMap and VvpInstanceMapOutput values.
// You can construct a concrete instance of `VvpInstanceMapInput` via:
//
//	VvpInstanceMap{ "key": VvpInstanceArgs{...} }
type VvpInstanceMapInput interface {
	pulumi.Input

	ToVvpInstanceMapOutput() VvpInstanceMapOutput
	ToVvpInstanceMapOutputWithContext(context.Context) VvpInstanceMapOutput
}

type VvpInstanceMap map[string]VvpInstanceInput

func (VvpInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VvpInstance)(nil)).Elem()
}

func (i VvpInstanceMap) ToVvpInstanceMapOutput() VvpInstanceMapOutput {
	return i.ToVvpInstanceMapOutputWithContext(context.Background())
}

func (i VvpInstanceMap) ToVvpInstanceMapOutputWithContext(ctx context.Context) VvpInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VvpInstanceMapOutput)
}

type VvpInstanceOutput struct{ *pulumi.OutputState }

func (VvpInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VvpInstance)(nil)).Elem()
}

func (o VvpInstanceOutput) ToVvpInstanceOutput() VvpInstanceOutput {
	return o
}

func (o VvpInstanceOutput) ToVvpInstanceOutputWithContext(ctx context.Context) VvpInstanceOutput {
	return o
}

// The creation time of the resource.
func (o VvpInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The number of subscription periods. If the payment type is PRE, this parameter is required.
func (o VvpInstanceOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.IntPtrOutput { return v.Duration }).(pulumi.IntPtrOutput)
}

// The payment type of the resource.
func (o VvpInstanceOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The subscription period. If the payment type is PRE, this parameter is required.
func (o VvpInstanceOutput) PricingCycle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringPtrOutput { return v.PricingCycle }).(pulumi.StringPtrOutput)
}

// The resource group to which the newly purchased instance belongs.
func (o VvpInstanceOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Resource specifications. See `resourceSpec` below.
func (o VvpInstanceOutput) ResourceSpec() VvpInstanceResourceSpecOutput {
	return o.ApplyT(func(v *VvpInstance) VvpInstanceResourceSpecOutput { return v.ResourceSpec }).(VvpInstanceResourceSpecOutput)
}

// The status of the resource.
func (o VvpInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Store information. See `storage` below.
func (o VvpInstanceOutput) Storage() VvpInstanceStorageOutput {
	return o.ApplyT(func(v *VvpInstance) VvpInstanceStorageOutput { return v.Storage }).(VvpInstanceStorageOutput)
}

// The tags of the resource.
func (o VvpInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The VPC ID of the user.
func (o VvpInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Virtual Switch ID.
func (o VvpInstanceOutput) VswitchIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringArrayOutput { return v.VswitchIds }).(pulumi.StringArrayOutput)
}

// The name of the workspace.
func (o VvpInstanceOutput) VvpInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.VvpInstanceName }).(pulumi.StringOutput)
}

// The zone ID of the resource.
func (o VvpInstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *VvpInstance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type VvpInstanceArrayOutput struct{ *pulumi.OutputState }

func (VvpInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VvpInstance)(nil)).Elem()
}

func (o VvpInstanceArrayOutput) ToVvpInstanceArrayOutput() VvpInstanceArrayOutput {
	return o
}

func (o VvpInstanceArrayOutput) ToVvpInstanceArrayOutputWithContext(ctx context.Context) VvpInstanceArrayOutput {
	return o
}

func (o VvpInstanceArrayOutput) Index(i pulumi.IntInput) VvpInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VvpInstance {
		return vs[0].([]*VvpInstance)[vs[1].(int)]
	}).(VvpInstanceOutput)
}

type VvpInstanceMapOutput struct{ *pulumi.OutputState }

func (VvpInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VvpInstance)(nil)).Elem()
}

func (o VvpInstanceMapOutput) ToVvpInstanceMapOutput() VvpInstanceMapOutput {
	return o
}

func (o VvpInstanceMapOutput) ToVvpInstanceMapOutputWithContext(ctx context.Context) VvpInstanceMapOutput {
	return o
}

func (o VvpInstanceMapOutput) MapIndex(k pulumi.StringInput) VvpInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VvpInstance {
		return vs[0].(map[string]*VvpInstance)[vs[1].(string)]
	}).(VvpInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VvpInstanceInput)(nil)).Elem(), &VvpInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*VvpInstanceArrayInput)(nil)).Elem(), VvpInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VvpInstanceMapInput)(nil)).Elem(), VvpInstanceMap{})
	pulumi.RegisterOutputType(VvpInstanceOutput{})
	pulumi.RegisterOutputType(VvpInstanceArrayOutput{})
	pulumi.RegisterOutputType(VvpInstanceMapOutput{})
}
