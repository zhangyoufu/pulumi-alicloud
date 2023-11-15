// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eflo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetSubnetsSubnet struct {
	// Network segment
	Cidr string `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime string `pulumi:"createTime"`
	// Modification time
	GmtModified string `pulumi:"gmtModified"`
	// The ID of the resource.
	Id string `pulumi:"id"`
	// Error message
	Message string `pulumi:"message"`
	// Resource Group ID.
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status string `pulumi:"status"`
	// Primary key ID.
	SubnetId string `pulumi:"subnetId"`
	// The Subnet name.
	SubnetName string `pulumi:"subnetName"`
	// Eflo subnet usage type, optional value:
	// - General type is not filled in
	// - OOB:OOB type
	// - LB: LB type
	Type string `pulumi:"type"`
	// The Eflo VPD ID.
	VpdId string `pulumi:"vpdId"`
	// The zone ID of the resource.
	ZoneId string `pulumi:"zoneId"`
}

// GetSubnetsSubnetInput is an input type that accepts GetSubnetsSubnetArgs and GetSubnetsSubnetOutput values.
// You can construct a concrete instance of `GetSubnetsSubnetInput` via:
//
//	GetSubnetsSubnetArgs{...}
type GetSubnetsSubnetInput interface {
	pulumi.Input

	ToGetSubnetsSubnetOutput() GetSubnetsSubnetOutput
	ToGetSubnetsSubnetOutputWithContext(context.Context) GetSubnetsSubnetOutput
}

type GetSubnetsSubnetArgs struct {
	// Network segment
	Cidr pulumi.StringInput `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Modification time
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The ID of the resource.
	Id pulumi.StringInput `pulumi:"id"`
	// Error message
	Message pulumi.StringInput `pulumi:"message"`
	// Resource Group ID.
	ResourceGroupId pulumi.StringInput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringInput `pulumi:"status"`
	// Primary key ID.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	// The Subnet name.
	SubnetName pulumi.StringInput `pulumi:"subnetName"`
	// Eflo subnet usage type, optional value:
	// - General type is not filled in
	// - OOB:OOB type
	// - LB: LB type
	Type pulumi.StringInput `pulumi:"type"`
	// The Eflo VPD ID.
	VpdId pulumi.StringInput `pulumi:"vpdId"`
	// The zone ID of the resource.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (GetSubnetsSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsSubnet)(nil)).Elem()
}

func (i GetSubnetsSubnetArgs) ToGetSubnetsSubnetOutput() GetSubnetsSubnetOutput {
	return i.ToGetSubnetsSubnetOutputWithContext(context.Background())
}

func (i GetSubnetsSubnetArgs) ToGetSubnetsSubnetOutputWithContext(ctx context.Context) GetSubnetsSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubnetsSubnetOutput)
}

// GetSubnetsSubnetArrayInput is an input type that accepts GetSubnetsSubnetArray and GetSubnetsSubnetArrayOutput values.
// You can construct a concrete instance of `GetSubnetsSubnetArrayInput` via:
//
//	GetSubnetsSubnetArray{ GetSubnetsSubnetArgs{...} }
type GetSubnetsSubnetArrayInput interface {
	pulumi.Input

	ToGetSubnetsSubnetArrayOutput() GetSubnetsSubnetArrayOutput
	ToGetSubnetsSubnetArrayOutputWithContext(context.Context) GetSubnetsSubnetArrayOutput
}

type GetSubnetsSubnetArray []GetSubnetsSubnetInput

func (GetSubnetsSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubnetsSubnet)(nil)).Elem()
}

func (i GetSubnetsSubnetArray) ToGetSubnetsSubnetArrayOutput() GetSubnetsSubnetArrayOutput {
	return i.ToGetSubnetsSubnetArrayOutputWithContext(context.Background())
}

func (i GetSubnetsSubnetArray) ToGetSubnetsSubnetArrayOutputWithContext(ctx context.Context) GetSubnetsSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubnetsSubnetArrayOutput)
}

type GetSubnetsSubnetOutput struct{ *pulumi.OutputState }

func (GetSubnetsSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsSubnet)(nil)).Elem()
}

func (o GetSubnetsSubnetOutput) ToGetSubnetsSubnetOutput() GetSubnetsSubnetOutput {
	return o
}

func (o GetSubnetsSubnetOutput) ToGetSubnetsSubnetOutputWithContext(ctx context.Context) GetSubnetsSubnetOutput {
	return o
}

// Network segment
func (o GetSubnetsSubnetOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.Cidr }).(pulumi.StringOutput)
}

// The creation time of the resource
func (o GetSubnetsSubnetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Modification time
func (o GetSubnetsSubnetOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The ID of the resource.
func (o GetSubnetsSubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.Id }).(pulumi.StringOutput)
}

// Error message
func (o GetSubnetsSubnetOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.Message }).(pulumi.StringOutput)
}

// Resource Group ID.
func (o GetSubnetsSubnetOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o GetSubnetsSubnetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.Status }).(pulumi.StringOutput)
}

// Primary key ID.
func (o GetSubnetsSubnetOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.SubnetId }).(pulumi.StringOutput)
}

// The Subnet name.
func (o GetSubnetsSubnetOutput) SubnetName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.SubnetName }).(pulumi.StringOutput)
}

// Eflo subnet usage type, optional value:
// - General type is not filled in
// - OOB:OOB type
// - LB: LB type
func (o GetSubnetsSubnetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.Type }).(pulumi.StringOutput)
}

// The Eflo VPD ID.
func (o GetSubnetsSubnetOutput) VpdId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.VpdId }).(pulumi.StringOutput)
}

// The zone ID of the resource.
func (o GetSubnetsSubnetOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsSubnet) string { return v.ZoneId }).(pulumi.StringOutput)
}

type GetSubnetsSubnetArrayOutput struct{ *pulumi.OutputState }

func (GetSubnetsSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSubnetsSubnet)(nil)).Elem()
}

func (o GetSubnetsSubnetArrayOutput) ToGetSubnetsSubnetArrayOutput() GetSubnetsSubnetArrayOutput {
	return o
}

func (o GetSubnetsSubnetArrayOutput) ToGetSubnetsSubnetArrayOutputWithContext(ctx context.Context) GetSubnetsSubnetArrayOutput {
	return o
}

func (o GetSubnetsSubnetArrayOutput) Index(i pulumi.IntInput) GetSubnetsSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSubnetsSubnet {
		return vs[0].([]GetSubnetsSubnet)[vs[1].(int)]
	}).(GetSubnetsSubnetOutput)
}

type GetVpdsVpd struct {
	// CIDR network segment
	Cidr string `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime string `pulumi:"createTime"`
	// Modification time
	GmtModified string `pulumi:"gmtModified"`
	// The id of the vpd.
	Id string `pulumi:"id"`
	// The Resource group id
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
	Status string `pulumi:"status"`
	// The id of the vpd.
	VpdId string `pulumi:"vpdId"`
	// The Name of the VPD.
	VpdName string `pulumi:"vpdName"`
}

// GetVpdsVpdInput is an input type that accepts GetVpdsVpdArgs and GetVpdsVpdOutput values.
// You can construct a concrete instance of `GetVpdsVpdInput` via:
//
//	GetVpdsVpdArgs{...}
type GetVpdsVpdInput interface {
	pulumi.Input

	ToGetVpdsVpdOutput() GetVpdsVpdOutput
	ToGetVpdsVpdOutputWithContext(context.Context) GetVpdsVpdOutput
}

type GetVpdsVpdArgs struct {
	// CIDR network segment
	Cidr pulumi.StringInput `pulumi:"cidr"`
	// The creation time of the resource
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Modification time
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The id of the vpd.
	Id pulumi.StringInput `pulumi:"id"`
	// The Resource group id
	ResourceGroupId pulumi.StringInput `pulumi:"resourceGroupId"`
	// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
	Status pulumi.StringInput `pulumi:"status"`
	// The id of the vpd.
	VpdId pulumi.StringInput `pulumi:"vpdId"`
	// The Name of the VPD.
	VpdName pulumi.StringInput `pulumi:"vpdName"`
}

func (GetVpdsVpdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpdsVpd)(nil)).Elem()
}

func (i GetVpdsVpdArgs) ToGetVpdsVpdOutput() GetVpdsVpdOutput {
	return i.ToGetVpdsVpdOutputWithContext(context.Background())
}

func (i GetVpdsVpdArgs) ToGetVpdsVpdOutputWithContext(ctx context.Context) GetVpdsVpdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpdsVpdOutput)
}

// GetVpdsVpdArrayInput is an input type that accepts GetVpdsVpdArray and GetVpdsVpdArrayOutput values.
// You can construct a concrete instance of `GetVpdsVpdArrayInput` via:
//
//	GetVpdsVpdArray{ GetVpdsVpdArgs{...} }
type GetVpdsVpdArrayInput interface {
	pulumi.Input

	ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput
	ToGetVpdsVpdArrayOutputWithContext(context.Context) GetVpdsVpdArrayOutput
}

type GetVpdsVpdArray []GetVpdsVpdInput

func (GetVpdsVpdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpdsVpd)(nil)).Elem()
}

func (i GetVpdsVpdArray) ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput {
	return i.ToGetVpdsVpdArrayOutputWithContext(context.Background())
}

func (i GetVpdsVpdArray) ToGetVpdsVpdArrayOutputWithContext(ctx context.Context) GetVpdsVpdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpdsVpdArrayOutput)
}

type GetVpdsVpdOutput struct{ *pulumi.OutputState }

func (GetVpdsVpdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpdsVpd)(nil)).Elem()
}

func (o GetVpdsVpdOutput) ToGetVpdsVpdOutput() GetVpdsVpdOutput {
	return o
}

func (o GetVpdsVpdOutput) ToGetVpdsVpdOutputWithContext(ctx context.Context) GetVpdsVpdOutput {
	return o
}

// CIDR network segment
func (o GetVpdsVpdOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Cidr }).(pulumi.StringOutput)
}

// The creation time of the resource
func (o GetVpdsVpdOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Modification time
func (o GetVpdsVpdOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The id of the vpd.
func (o GetVpdsVpdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Id }).(pulumi.StringOutput)
}

// The Resource group id
func (o GetVpdsVpdOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The Vpd status. Valid values: `Available`, `Not Available`, `Executing`, `Deleting`,
func (o GetVpdsVpdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.Status }).(pulumi.StringOutput)
}

// The id of the vpd.
func (o GetVpdsVpdOutput) VpdId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.VpdId }).(pulumi.StringOutput)
}

// The Name of the VPD.
func (o GetVpdsVpdOutput) VpdName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpdsVpd) string { return v.VpdName }).(pulumi.StringOutput)
}

type GetVpdsVpdArrayOutput struct{ *pulumi.OutputState }

func (GetVpdsVpdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpdsVpd)(nil)).Elem()
}

func (o GetVpdsVpdArrayOutput) ToGetVpdsVpdArrayOutput() GetVpdsVpdArrayOutput {
	return o
}

func (o GetVpdsVpdArrayOutput) ToGetVpdsVpdArrayOutputWithContext(ctx context.Context) GetVpdsVpdArrayOutput {
	return o
}

func (o GetVpdsVpdArrayOutput) Index(i pulumi.IntInput) GetVpdsVpdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpdsVpd {
		return vs[0].([]GetVpdsVpd)[vs[1].(int)]
	}).(GetVpdsVpdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetSubnetsSubnetInput)(nil)).Elem(), GetSubnetsSubnetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSubnetsSubnetArrayInput)(nil)).Elem(), GetSubnetsSubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpdsVpdInput)(nil)).Elem(), GetVpdsVpdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpdsVpdArrayInput)(nil)).Elem(), GetVpdsVpdArray{})
	pulumi.RegisterOutputType(GetSubnetsSubnetOutput{})
	pulumi.RegisterOutputType(GetSubnetsSubnetArrayOutput{})
	pulumi.RegisterOutputType(GetVpdsVpdOutput{})
	pulumi.RegisterOutputType(GetVpdsVpdArrayOutput{})
}
