// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rdc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetOrganizationsOrganization struct {
	// The ID of the Organization.
	Id string `pulumi:"id"`
	// The first ID of the resource.
	OrganizationId string `pulumi:"organizationId"`
	// Company name.
	OrganizationName string `pulumi:"organizationName"`
}

// GetOrganizationsOrganizationInput is an input type that accepts GetOrganizationsOrganizationArgs and GetOrganizationsOrganizationOutput values.
// You can construct a concrete instance of `GetOrganizationsOrganizationInput` via:
//
//	GetOrganizationsOrganizationArgs{...}
type GetOrganizationsOrganizationInput interface {
	pulumi.Input

	ToGetOrganizationsOrganizationOutput() GetOrganizationsOrganizationOutput
	ToGetOrganizationsOrganizationOutputWithContext(context.Context) GetOrganizationsOrganizationOutput
}

type GetOrganizationsOrganizationArgs struct {
	// The ID of the Organization.
	Id pulumi.StringInput `pulumi:"id"`
	// The first ID of the resource.
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
	// Company name.
	OrganizationName pulumi.StringInput `pulumi:"organizationName"`
}

func (GetOrganizationsOrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationsOrganization)(nil)).Elem()
}

func (i GetOrganizationsOrganizationArgs) ToGetOrganizationsOrganizationOutput() GetOrganizationsOrganizationOutput {
	return i.ToGetOrganizationsOrganizationOutputWithContext(context.Background())
}

func (i GetOrganizationsOrganizationArgs) ToGetOrganizationsOrganizationOutputWithContext(ctx context.Context) GetOrganizationsOrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetOrganizationsOrganizationOutput)
}

// GetOrganizationsOrganizationArrayInput is an input type that accepts GetOrganizationsOrganizationArray and GetOrganizationsOrganizationArrayOutput values.
// You can construct a concrete instance of `GetOrganizationsOrganizationArrayInput` via:
//
//	GetOrganizationsOrganizationArray{ GetOrganizationsOrganizationArgs{...} }
type GetOrganizationsOrganizationArrayInput interface {
	pulumi.Input

	ToGetOrganizationsOrganizationArrayOutput() GetOrganizationsOrganizationArrayOutput
	ToGetOrganizationsOrganizationArrayOutputWithContext(context.Context) GetOrganizationsOrganizationArrayOutput
}

type GetOrganizationsOrganizationArray []GetOrganizationsOrganizationInput

func (GetOrganizationsOrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetOrganizationsOrganization)(nil)).Elem()
}

func (i GetOrganizationsOrganizationArray) ToGetOrganizationsOrganizationArrayOutput() GetOrganizationsOrganizationArrayOutput {
	return i.ToGetOrganizationsOrganizationArrayOutputWithContext(context.Background())
}

func (i GetOrganizationsOrganizationArray) ToGetOrganizationsOrganizationArrayOutputWithContext(ctx context.Context) GetOrganizationsOrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetOrganizationsOrganizationArrayOutput)
}

type GetOrganizationsOrganizationOutput struct{ *pulumi.OutputState }

func (GetOrganizationsOrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationsOrganization)(nil)).Elem()
}

func (o GetOrganizationsOrganizationOutput) ToGetOrganizationsOrganizationOutput() GetOrganizationsOrganizationOutput {
	return o
}

func (o GetOrganizationsOrganizationOutput) ToGetOrganizationsOrganizationOutputWithContext(ctx context.Context) GetOrganizationsOrganizationOutput {
	return o
}

// The ID of the Organization.
func (o GetOrganizationsOrganizationOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationsOrganization) string { return v.Id }).(pulumi.StringOutput)
}

// The first ID of the resource.
func (o GetOrganizationsOrganizationOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationsOrganization) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Company name.
func (o GetOrganizationsOrganizationOutput) OrganizationName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationsOrganization) string { return v.OrganizationName }).(pulumi.StringOutput)
}

type GetOrganizationsOrganizationArrayOutput struct{ *pulumi.OutputState }

func (GetOrganizationsOrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetOrganizationsOrganization)(nil)).Elem()
}

func (o GetOrganizationsOrganizationArrayOutput) ToGetOrganizationsOrganizationArrayOutput() GetOrganizationsOrganizationArrayOutput {
	return o
}

func (o GetOrganizationsOrganizationArrayOutput) ToGetOrganizationsOrganizationArrayOutputWithContext(ctx context.Context) GetOrganizationsOrganizationArrayOutput {
	return o
}

func (o GetOrganizationsOrganizationArrayOutput) Index(i pulumi.IntInput) GetOrganizationsOrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetOrganizationsOrganization {
		return vs[0].([]GetOrganizationsOrganization)[vs[1].(int)]
	}).(GetOrganizationsOrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetOrganizationsOrganizationInput)(nil)).Elem(), GetOrganizationsOrganizationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetOrganizationsOrganizationArrayInput)(nil)).Elem(), GetOrganizationsOrganizationArray{})
	pulumi.RegisterOutputType(GetOrganizationsOrganizationOutput{})
	pulumi.RegisterOutputType(GetOrganizationsOrganizationArrayOutput{})
}
