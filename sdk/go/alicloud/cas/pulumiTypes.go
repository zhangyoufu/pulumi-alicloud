// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type GetCertificatesCertificate struct {
	// The cert is buy from aliyun or not.
	BuyInAliyun     bool   `pulumi:"buyInAliyun"`
	Cert            string `pulumi:"cert"`
	CertId          string `pulumi:"certId"`
	CertificateName string `pulumi:"certificateName"`
	// The cert's city.
	City string `pulumi:"city"`
	// The cert's common name.
	Common string `pulumi:"common"`
	// The cert's country.
	Country string `pulumi:"country"`
	// The cert's not valid after time.
	EndDate string `pulumi:"endDate"`
	// The cert is expired or not.
	Expired     bool   `pulumi:"expired"`
	Fingerprint string `pulumi:"fingerprint"`
	// The cert's id.
	Id string `pulumi:"id"`
	// The cert's .
	Issuer string `pulumi:"issuer"`
	Key    string `pulumi:"key"`
	// The cert's name.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
	Name string `pulumi:"name"`
	// The cert's organization.
	OrgName string `pulumi:"orgName"`
	// The cert's province.
	Province string `pulumi:"province"`
	// The cert's subject alternative name.
	Sans string `pulumi:"sans"`
	// The cert's not valid before time.
	StartDate string `pulumi:"startDate"`
}

// GetCertificatesCertificateInput is an input type that accepts GetCertificatesCertificateArgs and GetCertificatesCertificateOutput values.
// You can construct a concrete instance of `GetCertificatesCertificateInput` via:
//
//	GetCertificatesCertificateArgs{...}
type GetCertificatesCertificateInput interface {
	pulumi.Input

	ToGetCertificatesCertificateOutput() GetCertificatesCertificateOutput
	ToGetCertificatesCertificateOutputWithContext(context.Context) GetCertificatesCertificateOutput
}

type GetCertificatesCertificateArgs struct {
	// The cert is buy from aliyun or not.
	BuyInAliyun     pulumi.BoolInput   `pulumi:"buyInAliyun"`
	Cert            pulumi.StringInput `pulumi:"cert"`
	CertId          pulumi.StringInput `pulumi:"certId"`
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The cert's city.
	City pulumi.StringInput `pulumi:"city"`
	// The cert's common name.
	Common pulumi.StringInput `pulumi:"common"`
	// The cert's country.
	Country pulumi.StringInput `pulumi:"country"`
	// The cert's not valid after time.
	EndDate pulumi.StringInput `pulumi:"endDate"`
	// The cert is expired or not.
	Expired     pulumi.BoolInput   `pulumi:"expired"`
	Fingerprint pulumi.StringInput `pulumi:"fingerprint"`
	// The cert's id.
	Id pulumi.StringInput `pulumi:"id"`
	// The cert's .
	Issuer pulumi.StringInput `pulumi:"issuer"`
	Key    pulumi.StringInput `pulumi:"key"`
	// The cert's name.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
	Name pulumi.StringInput `pulumi:"name"`
	// The cert's organization.
	OrgName pulumi.StringInput `pulumi:"orgName"`
	// The cert's province.
	Province pulumi.StringInput `pulumi:"province"`
	// The cert's subject alternative name.
	Sans pulumi.StringInput `pulumi:"sans"`
	// The cert's not valid before time.
	StartDate pulumi.StringInput `pulumi:"startDate"`
}

func (GetCertificatesCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesCertificate)(nil)).Elem()
}

func (i GetCertificatesCertificateArgs) ToGetCertificatesCertificateOutput() GetCertificatesCertificateOutput {
	return i.ToGetCertificatesCertificateOutputWithContext(context.Background())
}

func (i GetCertificatesCertificateArgs) ToGetCertificatesCertificateOutputWithContext(ctx context.Context) GetCertificatesCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCertificatesCertificateOutput)
}

func (i GetCertificatesCertificateArgs) ToOutput(ctx context.Context) pulumix.Output[GetCertificatesCertificate] {
	return pulumix.Output[GetCertificatesCertificate]{
		OutputState: i.ToGetCertificatesCertificateOutputWithContext(ctx).OutputState,
	}
}

// GetCertificatesCertificateArrayInput is an input type that accepts GetCertificatesCertificateArray and GetCertificatesCertificateArrayOutput values.
// You can construct a concrete instance of `GetCertificatesCertificateArrayInput` via:
//
//	GetCertificatesCertificateArray{ GetCertificatesCertificateArgs{...} }
type GetCertificatesCertificateArrayInput interface {
	pulumi.Input

	ToGetCertificatesCertificateArrayOutput() GetCertificatesCertificateArrayOutput
	ToGetCertificatesCertificateArrayOutputWithContext(context.Context) GetCertificatesCertificateArrayOutput
}

type GetCertificatesCertificateArray []GetCertificatesCertificateInput

func (GetCertificatesCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCertificatesCertificate)(nil)).Elem()
}

func (i GetCertificatesCertificateArray) ToGetCertificatesCertificateArrayOutput() GetCertificatesCertificateArrayOutput {
	return i.ToGetCertificatesCertificateArrayOutputWithContext(context.Background())
}

func (i GetCertificatesCertificateArray) ToGetCertificatesCertificateArrayOutputWithContext(ctx context.Context) GetCertificatesCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCertificatesCertificateArrayOutput)
}

func (i GetCertificatesCertificateArray) ToOutput(ctx context.Context) pulumix.Output[[]GetCertificatesCertificate] {
	return pulumix.Output[[]GetCertificatesCertificate]{
		OutputState: i.ToGetCertificatesCertificateArrayOutputWithContext(ctx).OutputState,
	}
}

type GetCertificatesCertificateOutput struct{ *pulumi.OutputState }

func (GetCertificatesCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesCertificate)(nil)).Elem()
}

func (o GetCertificatesCertificateOutput) ToGetCertificatesCertificateOutput() GetCertificatesCertificateOutput {
	return o
}

func (o GetCertificatesCertificateOutput) ToGetCertificatesCertificateOutputWithContext(ctx context.Context) GetCertificatesCertificateOutput {
	return o
}

func (o GetCertificatesCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[GetCertificatesCertificate] {
	return pulumix.Output[GetCertificatesCertificate]{
		OutputState: o.OutputState,
	}
}

// The cert is buy from aliyun or not.
func (o GetCertificatesCertificateOutput) BuyInAliyun() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) bool { return v.BuyInAliyun }).(pulumi.BoolOutput)
}

func (o GetCertificatesCertificateOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Cert }).(pulumi.StringOutput)
}

func (o GetCertificatesCertificateOutput) CertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.CertId }).(pulumi.StringOutput)
}

func (o GetCertificatesCertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.CertificateName }).(pulumi.StringOutput)
}

// The cert's city.
func (o GetCertificatesCertificateOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.City }).(pulumi.StringOutput)
}

// The cert's common name.
func (o GetCertificatesCertificateOutput) Common() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Common }).(pulumi.StringOutput)
}

// The cert's country.
func (o GetCertificatesCertificateOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Country }).(pulumi.StringOutput)
}

// The cert's not valid after time.
func (o GetCertificatesCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.EndDate }).(pulumi.StringOutput)
}

// The cert is expired or not.
func (o GetCertificatesCertificateOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) bool { return v.Expired }).(pulumi.BoolOutput)
}

func (o GetCertificatesCertificateOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The cert's id.
func (o GetCertificatesCertificateOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Id }).(pulumi.StringOutput)
}

// The cert's .
func (o GetCertificatesCertificateOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o GetCertificatesCertificateOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Key }).(pulumi.StringOutput)
}

// The cert's name.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
func (o GetCertificatesCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Name }).(pulumi.StringOutput)
}

// The cert's organization.
func (o GetCertificatesCertificateOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.OrgName }).(pulumi.StringOutput)
}

// The cert's province.
func (o GetCertificatesCertificateOutput) Province() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Province }).(pulumi.StringOutput)
}

// The cert's subject alternative name.
func (o GetCertificatesCertificateOutput) Sans() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.Sans }).(pulumi.StringOutput)
}

// The cert's not valid before time.
func (o GetCertificatesCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesCertificate) string { return v.StartDate }).(pulumi.StringOutput)
}

type GetCertificatesCertificateArrayOutput struct{ *pulumi.OutputState }

func (GetCertificatesCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCertificatesCertificate)(nil)).Elem()
}

func (o GetCertificatesCertificateArrayOutput) ToGetCertificatesCertificateArrayOutput() GetCertificatesCertificateArrayOutput {
	return o
}

func (o GetCertificatesCertificateArrayOutput) ToGetCertificatesCertificateArrayOutputWithContext(ctx context.Context) GetCertificatesCertificateArrayOutput {
	return o
}

func (o GetCertificatesCertificateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GetCertificatesCertificate] {
	return pulumix.Output[[]GetCertificatesCertificate]{
		OutputState: o.OutputState,
	}
}

func (o GetCertificatesCertificateArrayOutput) Index(i pulumi.IntInput) GetCertificatesCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCertificatesCertificate {
		return vs[0].([]GetCertificatesCertificate)[vs[1].(int)]
	}).(GetCertificatesCertificateOutput)
}

type GetServiceCertificatesCertificate struct {
	// The cert is buy from aliyun or not.
	BuyInAliyun bool `pulumi:"buyInAliyun"`
	// The cert's Cert.
	Cert string `pulumi:"cert"`
	// The cert's id.
	CertId string `pulumi:"certId"`
	// The cert's name.
	CertificateName string `pulumi:"certificateName"`
	// The cert's city.
	City string `pulumi:"city"`
	// The cert's common name.
	Common string `pulumi:"common"`
	// The cert's country.
	Country string `pulumi:"country"`
	// The cert's not valid after time.
	EndDate string `pulumi:"endDate"`
	// The cert is expired or not.
	Expired bool `pulumi:"expired"`
	// The cert's finger.
	Fingerprint string `pulumi:"fingerprint"`
	// The cert's id.
	Id string `pulumi:"id"`
	// The cert's Issuer.
	Issuer string `pulumi:"issuer"`
	// The cert's Keye.
	Key string `pulumi:"key"`
	// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
	Name string `pulumi:"name"`
	// The cert's organization.
	OrgName string `pulumi:"orgName"`
	// The cert's province.
	Province string `pulumi:"province"`
	// The cert's subject alternative name.
	Sans string `pulumi:"sans"`
	// The cert's not valid before time.
	StartDate string `pulumi:"startDate"`
}

// GetServiceCertificatesCertificateInput is an input type that accepts GetServiceCertificatesCertificateArgs and GetServiceCertificatesCertificateOutput values.
// You can construct a concrete instance of `GetServiceCertificatesCertificateInput` via:
//
//	GetServiceCertificatesCertificateArgs{...}
type GetServiceCertificatesCertificateInput interface {
	pulumi.Input

	ToGetServiceCertificatesCertificateOutput() GetServiceCertificatesCertificateOutput
	ToGetServiceCertificatesCertificateOutputWithContext(context.Context) GetServiceCertificatesCertificateOutput
}

type GetServiceCertificatesCertificateArgs struct {
	// The cert is buy from aliyun or not.
	BuyInAliyun pulumi.BoolInput `pulumi:"buyInAliyun"`
	// The cert's Cert.
	Cert pulumi.StringInput `pulumi:"cert"`
	// The cert's id.
	CertId pulumi.StringInput `pulumi:"certId"`
	// The cert's name.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The cert's city.
	City pulumi.StringInput `pulumi:"city"`
	// The cert's common name.
	Common pulumi.StringInput `pulumi:"common"`
	// The cert's country.
	Country pulumi.StringInput `pulumi:"country"`
	// The cert's not valid after time.
	EndDate pulumi.StringInput `pulumi:"endDate"`
	// The cert is expired or not.
	Expired pulumi.BoolInput `pulumi:"expired"`
	// The cert's finger.
	Fingerprint pulumi.StringInput `pulumi:"fingerprint"`
	// The cert's id.
	Id pulumi.StringInput `pulumi:"id"`
	// The cert's Issuer.
	Issuer pulumi.StringInput `pulumi:"issuer"`
	// The cert's Keye.
	Key pulumi.StringInput `pulumi:"key"`
	// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
	Name pulumi.StringInput `pulumi:"name"`
	// The cert's organization.
	OrgName pulumi.StringInput `pulumi:"orgName"`
	// The cert's province.
	Province pulumi.StringInput `pulumi:"province"`
	// The cert's subject alternative name.
	Sans pulumi.StringInput `pulumi:"sans"`
	// The cert's not valid before time.
	StartDate pulumi.StringInput `pulumi:"startDate"`
}

func (GetServiceCertificatesCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesCertificate)(nil)).Elem()
}

func (i GetServiceCertificatesCertificateArgs) ToGetServiceCertificatesCertificateOutput() GetServiceCertificatesCertificateOutput {
	return i.ToGetServiceCertificatesCertificateOutputWithContext(context.Background())
}

func (i GetServiceCertificatesCertificateArgs) ToGetServiceCertificatesCertificateOutputWithContext(ctx context.Context) GetServiceCertificatesCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceCertificatesCertificateOutput)
}

func (i GetServiceCertificatesCertificateArgs) ToOutput(ctx context.Context) pulumix.Output[GetServiceCertificatesCertificate] {
	return pulumix.Output[GetServiceCertificatesCertificate]{
		OutputState: i.ToGetServiceCertificatesCertificateOutputWithContext(ctx).OutputState,
	}
}

// GetServiceCertificatesCertificateArrayInput is an input type that accepts GetServiceCertificatesCertificateArray and GetServiceCertificatesCertificateArrayOutput values.
// You can construct a concrete instance of `GetServiceCertificatesCertificateArrayInput` via:
//
//	GetServiceCertificatesCertificateArray{ GetServiceCertificatesCertificateArgs{...} }
type GetServiceCertificatesCertificateArrayInput interface {
	pulumi.Input

	ToGetServiceCertificatesCertificateArrayOutput() GetServiceCertificatesCertificateArrayOutput
	ToGetServiceCertificatesCertificateArrayOutputWithContext(context.Context) GetServiceCertificatesCertificateArrayOutput
}

type GetServiceCertificatesCertificateArray []GetServiceCertificatesCertificateInput

func (GetServiceCertificatesCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceCertificatesCertificate)(nil)).Elem()
}

func (i GetServiceCertificatesCertificateArray) ToGetServiceCertificatesCertificateArrayOutput() GetServiceCertificatesCertificateArrayOutput {
	return i.ToGetServiceCertificatesCertificateArrayOutputWithContext(context.Background())
}

func (i GetServiceCertificatesCertificateArray) ToGetServiceCertificatesCertificateArrayOutputWithContext(ctx context.Context) GetServiceCertificatesCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceCertificatesCertificateArrayOutput)
}

func (i GetServiceCertificatesCertificateArray) ToOutput(ctx context.Context) pulumix.Output[[]GetServiceCertificatesCertificate] {
	return pulumix.Output[[]GetServiceCertificatesCertificate]{
		OutputState: i.ToGetServiceCertificatesCertificateArrayOutputWithContext(ctx).OutputState,
	}
}

type GetServiceCertificatesCertificateOutput struct{ *pulumi.OutputState }

func (GetServiceCertificatesCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesCertificate)(nil)).Elem()
}

func (o GetServiceCertificatesCertificateOutput) ToGetServiceCertificatesCertificateOutput() GetServiceCertificatesCertificateOutput {
	return o
}

func (o GetServiceCertificatesCertificateOutput) ToGetServiceCertificatesCertificateOutputWithContext(ctx context.Context) GetServiceCertificatesCertificateOutput {
	return o
}

func (o GetServiceCertificatesCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[GetServiceCertificatesCertificate] {
	return pulumix.Output[GetServiceCertificatesCertificate]{
		OutputState: o.OutputState,
	}
}

// The cert is buy from aliyun or not.
func (o GetServiceCertificatesCertificateOutput) BuyInAliyun() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) bool { return v.BuyInAliyun }).(pulumi.BoolOutput)
}

// The cert's Cert.
func (o GetServiceCertificatesCertificateOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Cert }).(pulumi.StringOutput)
}

// The cert's id.
func (o GetServiceCertificatesCertificateOutput) CertId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.CertId }).(pulumi.StringOutput)
}

// The cert's name.
func (o GetServiceCertificatesCertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.CertificateName }).(pulumi.StringOutput)
}

// The cert's city.
func (o GetServiceCertificatesCertificateOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.City }).(pulumi.StringOutput)
}

// The cert's common name.
func (o GetServiceCertificatesCertificateOutput) Common() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Common }).(pulumi.StringOutput)
}

// The cert's country.
func (o GetServiceCertificatesCertificateOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Country }).(pulumi.StringOutput)
}

// The cert's not valid after time.
func (o GetServiceCertificatesCertificateOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.EndDate }).(pulumi.StringOutput)
}

// The cert is expired or not.
func (o GetServiceCertificatesCertificateOutput) Expired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) bool { return v.Expired }).(pulumi.BoolOutput)
}

// The cert's finger.
func (o GetServiceCertificatesCertificateOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The cert's id.
func (o GetServiceCertificatesCertificateOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Id }).(pulumi.StringOutput)
}

// The cert's Issuer.
func (o GetServiceCertificatesCertificateOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Issuer }).(pulumi.StringOutput)
}

// The cert's Keye.
func (o GetServiceCertificatesCertificateOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Key }).(pulumi.StringOutput)
}

// Deprecated: Field 'name' has been deprecated from provider version 1.129.0 and it will be removed in the future version. Please use the new attribute 'certificate_name' instead.
func (o GetServiceCertificatesCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Name }).(pulumi.StringOutput)
}

// The cert's organization.
func (o GetServiceCertificatesCertificateOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.OrgName }).(pulumi.StringOutput)
}

// The cert's province.
func (o GetServiceCertificatesCertificateOutput) Province() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Province }).(pulumi.StringOutput)
}

// The cert's subject alternative name.
func (o GetServiceCertificatesCertificateOutput) Sans() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.Sans }).(pulumi.StringOutput)
}

// The cert's not valid before time.
func (o GetServiceCertificatesCertificateOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesCertificate) string { return v.StartDate }).(pulumi.StringOutput)
}

type GetServiceCertificatesCertificateArrayOutput struct{ *pulumi.OutputState }

func (GetServiceCertificatesCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceCertificatesCertificate)(nil)).Elem()
}

func (o GetServiceCertificatesCertificateArrayOutput) ToGetServiceCertificatesCertificateArrayOutput() GetServiceCertificatesCertificateArrayOutput {
	return o
}

func (o GetServiceCertificatesCertificateArrayOutput) ToGetServiceCertificatesCertificateArrayOutputWithContext(ctx context.Context) GetServiceCertificatesCertificateArrayOutput {
	return o
}

func (o GetServiceCertificatesCertificateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GetServiceCertificatesCertificate] {
	return pulumix.Output[[]GetServiceCertificatesCertificate]{
		OutputState: o.OutputState,
	}
}

func (o GetServiceCertificatesCertificateArrayOutput) Index(i pulumi.IntInput) GetServiceCertificatesCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceCertificatesCertificate {
		return vs[0].([]GetServiceCertificatesCertificate)[vs[1].(int)]
	}).(GetServiceCertificatesCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetCertificatesCertificateInput)(nil)).Elem(), GetCertificatesCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCertificatesCertificateArrayInput)(nil)).Elem(), GetCertificatesCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceCertificatesCertificateInput)(nil)).Elem(), GetServiceCertificatesCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceCertificatesCertificateArrayInput)(nil)).Elem(), GetServiceCertificatesCertificateArray{})
	pulumi.RegisterOutputType(GetCertificatesCertificateOutput{})
	pulumi.RegisterOutputType(GetCertificatesCertificateArrayOutput{})
	pulumi.RegisterOutputType(GetServiceCertificatesCertificateOutput{})
	pulumi.RegisterOutputType(GetServiceCertificatesCertificateArrayOutput{})
}
