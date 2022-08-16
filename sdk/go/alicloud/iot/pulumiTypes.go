// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetDeviceGroupsGroup struct {
	// The Group CreateTime.
	CreateTime string `pulumi:"createTime"`
	// The Group Number of activated devices.
	DeviceActive string `pulumi:"deviceActive"`
	// The Group Total number of devices.
	DeviceCount string `pulumi:"deviceCount"`
	// The Group Number of online devices.
	DeviceOnline string `pulumi:"deviceOnline"`
	// The Error_Message of the device group.
	ErrorMessage string `pulumi:"errorMessage"`
	// The GroupDesc of the device group.
	GroupDesc string `pulumi:"groupDesc"`
	// The GroupId of the device group.
	GroupId string `pulumi:"groupId"`
	// The GroupName of the device group.
	GroupName string `pulumi:"groupName"`
	// The ID of the device group.
	Id string `pulumi:"id"`
	// Whether the call is successful.
	Success bool `pulumi:"success"`
}

// GetDeviceGroupsGroupInput is an input type that accepts GetDeviceGroupsGroupArgs and GetDeviceGroupsGroupOutput values.
// You can construct a concrete instance of `GetDeviceGroupsGroupInput` via:
//
//	GetDeviceGroupsGroupArgs{...}
type GetDeviceGroupsGroupInput interface {
	pulumi.Input

	ToGetDeviceGroupsGroupOutput() GetDeviceGroupsGroupOutput
	ToGetDeviceGroupsGroupOutputWithContext(context.Context) GetDeviceGroupsGroupOutput
}

type GetDeviceGroupsGroupArgs struct {
	// The Group CreateTime.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The Group Number of activated devices.
	DeviceActive pulumi.StringInput `pulumi:"deviceActive"`
	// The Group Total number of devices.
	DeviceCount pulumi.StringInput `pulumi:"deviceCount"`
	// The Group Number of online devices.
	DeviceOnline pulumi.StringInput `pulumi:"deviceOnline"`
	// The Error_Message of the device group.
	ErrorMessage pulumi.StringInput `pulumi:"errorMessage"`
	// The GroupDesc of the device group.
	GroupDesc pulumi.StringInput `pulumi:"groupDesc"`
	// The GroupId of the device group.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The GroupName of the device group.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// The ID of the device group.
	Id pulumi.StringInput `pulumi:"id"`
	// Whether the call is successful.
	Success pulumi.BoolInput `pulumi:"success"`
}

func (GetDeviceGroupsGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceGroupsGroup)(nil)).Elem()
}

func (i GetDeviceGroupsGroupArgs) ToGetDeviceGroupsGroupOutput() GetDeviceGroupsGroupOutput {
	return i.ToGetDeviceGroupsGroupOutputWithContext(context.Background())
}

func (i GetDeviceGroupsGroupArgs) ToGetDeviceGroupsGroupOutputWithContext(ctx context.Context) GetDeviceGroupsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDeviceGroupsGroupOutput)
}

// GetDeviceGroupsGroupArrayInput is an input type that accepts GetDeviceGroupsGroupArray and GetDeviceGroupsGroupArrayOutput values.
// You can construct a concrete instance of `GetDeviceGroupsGroupArrayInput` via:
//
//	GetDeviceGroupsGroupArray{ GetDeviceGroupsGroupArgs{...} }
type GetDeviceGroupsGroupArrayInput interface {
	pulumi.Input

	ToGetDeviceGroupsGroupArrayOutput() GetDeviceGroupsGroupArrayOutput
	ToGetDeviceGroupsGroupArrayOutputWithContext(context.Context) GetDeviceGroupsGroupArrayOutput
}

type GetDeviceGroupsGroupArray []GetDeviceGroupsGroupInput

func (GetDeviceGroupsGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDeviceGroupsGroup)(nil)).Elem()
}

func (i GetDeviceGroupsGroupArray) ToGetDeviceGroupsGroupArrayOutput() GetDeviceGroupsGroupArrayOutput {
	return i.ToGetDeviceGroupsGroupArrayOutputWithContext(context.Background())
}

func (i GetDeviceGroupsGroupArray) ToGetDeviceGroupsGroupArrayOutputWithContext(ctx context.Context) GetDeviceGroupsGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDeviceGroupsGroupArrayOutput)
}

type GetDeviceGroupsGroupOutput struct{ *pulumi.OutputState }

func (GetDeviceGroupsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceGroupsGroup)(nil)).Elem()
}

func (o GetDeviceGroupsGroupOutput) ToGetDeviceGroupsGroupOutput() GetDeviceGroupsGroupOutput {
	return o
}

func (o GetDeviceGroupsGroupOutput) ToGetDeviceGroupsGroupOutputWithContext(ctx context.Context) GetDeviceGroupsGroupOutput {
	return o
}

// The Group CreateTime.
func (o GetDeviceGroupsGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The Group Number of activated devices.
func (o GetDeviceGroupsGroupOutput) DeviceActive() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.DeviceActive }).(pulumi.StringOutput)
}

// The Group Total number of devices.
func (o GetDeviceGroupsGroupOutput) DeviceCount() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.DeviceCount }).(pulumi.StringOutput)
}

// The Group Number of online devices.
func (o GetDeviceGroupsGroupOutput) DeviceOnline() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.DeviceOnline }).(pulumi.StringOutput)
}

// The Error_Message of the device group.
func (o GetDeviceGroupsGroupOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The GroupDesc of the device group.
func (o GetDeviceGroupsGroupOutput) GroupDesc() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.GroupDesc }).(pulumi.StringOutput)
}

// The GroupId of the device group.
func (o GetDeviceGroupsGroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.GroupId }).(pulumi.StringOutput)
}

// The GroupName of the device group.
func (o GetDeviceGroupsGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.GroupName }).(pulumi.StringOutput)
}

// The ID of the device group.
func (o GetDeviceGroupsGroupOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the call is successful.
func (o GetDeviceGroupsGroupOutput) Success() pulumi.BoolOutput {
	return o.ApplyT(func(v GetDeviceGroupsGroup) bool { return v.Success }).(pulumi.BoolOutput)
}

type GetDeviceGroupsGroupArrayOutput struct{ *pulumi.OutputState }

func (GetDeviceGroupsGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDeviceGroupsGroup)(nil)).Elem()
}

func (o GetDeviceGroupsGroupArrayOutput) ToGetDeviceGroupsGroupArrayOutput() GetDeviceGroupsGroupArrayOutput {
	return o
}

func (o GetDeviceGroupsGroupArrayOutput) ToGetDeviceGroupsGroupArrayOutputWithContext(ctx context.Context) GetDeviceGroupsGroupArrayOutput {
	return o
}

func (o GetDeviceGroupsGroupArrayOutput) Index(i pulumi.IntInput) GetDeviceGroupsGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDeviceGroupsGroup {
		return vs[0].([]GetDeviceGroupsGroup)[vs[1].(int)]
	}).(GetDeviceGroupsGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetDeviceGroupsGroupInput)(nil)).Elem(), GetDeviceGroupsGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDeviceGroupsGroupArrayInput)(nil)).Elem(), GetDeviceGroupsGroupArray{})
	pulumi.RegisterOutputType(GetDeviceGroupsGroupOutput{})
	pulumi.RegisterOutputType(GetDeviceGroupsGroupArrayOutput{})
}
