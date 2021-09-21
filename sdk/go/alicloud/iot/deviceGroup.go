// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Iot Device Group resource.
//
// For information about Iot Device Group and how to use it, see [What is Device Group](https://www.alibabacloud.com/help/product/30520.htm).
//
// > **NOTE:** Available in v1.134.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/iot"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iot.NewDeviceGroup(ctx, "example", &iot.DeviceGroupArgs{
// 			GroupName: pulumi.String("example_value"),
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
// Iot Device Group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:iot/deviceGroup:DeviceGroup example <id>
// ```
type DeviceGroup struct {
	pulumi.CustomResourceState

	// The GroupDesc of the device group.
	GroupDesc pulumi.StringPtrOutput `pulumi:"groupDesc"`
	// The GroupName of the device group.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The id of the Iot Instance.
	IotInstanceId pulumi.StringPtrOutput `pulumi:"iotInstanceId"`
	// The id of the SuperGroup.
	SuperGroupId pulumi.StringPtrOutput `pulumi:"superGroupId"`
}

// NewDeviceGroup registers a new resource with the given unique name, arguments, and options.
func NewDeviceGroup(ctx *pulumi.Context,
	name string, args *DeviceGroupArgs, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	var resource DeviceGroup
	err := ctx.RegisterResource("alicloud:iot/deviceGroup:DeviceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceGroup gets an existing DeviceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceGroupState, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	var resource DeviceGroup
	err := ctx.ReadResource("alicloud:iot/deviceGroup:DeviceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceGroup resources.
type deviceGroupState struct {
	// The GroupDesc of the device group.
	GroupDesc *string `pulumi:"groupDesc"`
	// The GroupName of the device group.
	GroupName *string `pulumi:"groupName"`
	// The id of the Iot Instance.
	IotInstanceId *string `pulumi:"iotInstanceId"`
	// The id of the SuperGroup.
	SuperGroupId *string `pulumi:"superGroupId"`
}

type DeviceGroupState struct {
	// The GroupDesc of the device group.
	GroupDesc pulumi.StringPtrInput
	// The GroupName of the device group.
	GroupName pulumi.StringPtrInput
	// The id of the Iot Instance.
	IotInstanceId pulumi.StringPtrInput
	// The id of the SuperGroup.
	SuperGroupId pulumi.StringPtrInput
}

func (DeviceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupState)(nil)).Elem()
}

type deviceGroupArgs struct {
	// The GroupDesc of the device group.
	GroupDesc *string `pulumi:"groupDesc"`
	// The GroupName of the device group.
	GroupName string `pulumi:"groupName"`
	// The id of the Iot Instance.
	IotInstanceId *string `pulumi:"iotInstanceId"`
	// The id of the SuperGroup.
	SuperGroupId *string `pulumi:"superGroupId"`
}

// The set of arguments for constructing a DeviceGroup resource.
type DeviceGroupArgs struct {
	// The GroupDesc of the device group.
	GroupDesc pulumi.StringPtrInput
	// The GroupName of the device group.
	GroupName pulumi.StringInput
	// The id of the Iot Instance.
	IotInstanceId pulumi.StringPtrInput
	// The id of the SuperGroup.
	SuperGroupId pulumi.StringPtrInput
}

func (DeviceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupArgs)(nil)).Elem()
}

type DeviceGroupInput interface {
	pulumi.Input

	ToDeviceGroupOutput() DeviceGroupOutput
	ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput
}

func (*DeviceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceGroup)(nil))
}

func (i *DeviceGroup) ToDeviceGroupOutput() DeviceGroupOutput {
	return i.ToDeviceGroupOutputWithContext(context.Background())
}

func (i *DeviceGroup) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupOutput)
}

func (i *DeviceGroup) ToDeviceGroupPtrOutput() DeviceGroupPtrOutput {
	return i.ToDeviceGroupPtrOutputWithContext(context.Background())
}

func (i *DeviceGroup) ToDeviceGroupPtrOutputWithContext(ctx context.Context) DeviceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupPtrOutput)
}

type DeviceGroupPtrInput interface {
	pulumi.Input

	ToDeviceGroupPtrOutput() DeviceGroupPtrOutput
	ToDeviceGroupPtrOutputWithContext(ctx context.Context) DeviceGroupPtrOutput
}

type deviceGroupPtrType DeviceGroupArgs

func (*deviceGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil))
}

func (i *deviceGroupPtrType) ToDeviceGroupPtrOutput() DeviceGroupPtrOutput {
	return i.ToDeviceGroupPtrOutputWithContext(context.Background())
}

func (i *deviceGroupPtrType) ToDeviceGroupPtrOutputWithContext(ctx context.Context) DeviceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupPtrOutput)
}

// DeviceGroupArrayInput is an input type that accepts DeviceGroupArray and DeviceGroupArrayOutput values.
// You can construct a concrete instance of `DeviceGroupArrayInput` via:
//
//          DeviceGroupArray{ DeviceGroupArgs{...} }
type DeviceGroupArrayInput interface {
	pulumi.Input

	ToDeviceGroupArrayOutput() DeviceGroupArrayOutput
	ToDeviceGroupArrayOutputWithContext(context.Context) DeviceGroupArrayOutput
}

type DeviceGroupArray []DeviceGroupInput

func (DeviceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DeviceGroup)(nil))
}

func (i DeviceGroupArray) ToDeviceGroupArrayOutput() DeviceGroupArrayOutput {
	return i.ToDeviceGroupArrayOutputWithContext(context.Background())
}

func (i DeviceGroupArray) ToDeviceGroupArrayOutputWithContext(ctx context.Context) DeviceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupArrayOutput)
}

// DeviceGroupMapInput is an input type that accepts DeviceGroupMap and DeviceGroupMapOutput values.
// You can construct a concrete instance of `DeviceGroupMapInput` via:
//
//          DeviceGroupMap{ "key": DeviceGroupArgs{...} }
type DeviceGroupMapInput interface {
	pulumi.Input

	ToDeviceGroupMapOutput() DeviceGroupMapOutput
	ToDeviceGroupMapOutputWithContext(context.Context) DeviceGroupMapOutput
}

type DeviceGroupMap map[string]DeviceGroupInput

func (DeviceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DeviceGroup)(nil))
}

func (i DeviceGroupMap) ToDeviceGroupMapOutput() DeviceGroupMapOutput {
	return i.ToDeviceGroupMapOutputWithContext(context.Background())
}

func (i DeviceGroupMap) ToDeviceGroupMapOutputWithContext(ctx context.Context) DeviceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupMapOutput)
}

type DeviceGroupOutput struct {
	*pulumi.OutputState
}

func (DeviceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceGroup)(nil))
}

func (o DeviceGroupOutput) ToDeviceGroupOutput() DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) ToDeviceGroupPtrOutput() DeviceGroupPtrOutput {
	return o.ToDeviceGroupPtrOutputWithContext(context.Background())
}

func (o DeviceGroupOutput) ToDeviceGroupPtrOutputWithContext(ctx context.Context) DeviceGroupPtrOutput {
	return o.ApplyT(func(v DeviceGroup) *DeviceGroup {
		return &v
	}).(DeviceGroupPtrOutput)
}

type DeviceGroupPtrOutput struct {
	*pulumi.OutputState
}

func (DeviceGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil))
}

func (o DeviceGroupPtrOutput) ToDeviceGroupPtrOutput() DeviceGroupPtrOutput {
	return o
}

func (o DeviceGroupPtrOutput) ToDeviceGroupPtrOutputWithContext(ctx context.Context) DeviceGroupPtrOutput {
	return o
}

type DeviceGroupArrayOutput struct{ *pulumi.OutputState }

func (DeviceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceGroup)(nil))
}

func (o DeviceGroupArrayOutput) ToDeviceGroupArrayOutput() DeviceGroupArrayOutput {
	return o
}

func (o DeviceGroupArrayOutput) ToDeviceGroupArrayOutputWithContext(ctx context.Context) DeviceGroupArrayOutput {
	return o
}

func (o DeviceGroupArrayOutput) Index(i pulumi.IntInput) DeviceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeviceGroup {
		return vs[0].([]DeviceGroup)[vs[1].(int)]
	}).(DeviceGroupOutput)
}

type DeviceGroupMapOutput struct{ *pulumi.OutputState }

func (DeviceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeviceGroup)(nil))
}

func (o DeviceGroupMapOutput) ToDeviceGroupMapOutput() DeviceGroupMapOutput {
	return o
}

func (o DeviceGroupMapOutput) ToDeviceGroupMapOutputWithContext(ctx context.Context) DeviceGroupMapOutput {
	return o
}

func (o DeviceGroupMapOutput) MapIndex(k pulumi.StringInput) DeviceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeviceGroup {
		return vs[0].(map[string]DeviceGroup)[vs[1].(string)]
	}).(DeviceGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceGroupOutput{})
	pulumi.RegisterOutputType(DeviceGroupPtrOutput{})
	pulumi.RegisterOutputType(DeviceGroupArrayOutput{})
	pulumi.RegisterOutputType(DeviceGroupMapOutput{})
}
