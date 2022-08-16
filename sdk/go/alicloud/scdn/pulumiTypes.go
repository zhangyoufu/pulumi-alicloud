// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainCertInfo struct {
	// If You Enable HTTPS Here Certificate Name.
	CertName *string `pulumi:"certName"`
	// Certificate Type. Value Range:
	// * upload: Certificate
	// * cas: Certificate Authority Certificate.
	// * free: Free Certificate.
	CertType *string `pulumi:"certType"`
	// Private Key. Do Not Enable Certificate without Entering a User Name and Configure Certificates Enter Private Key.
	SslPri *string `pulumi:"sslPri"`
	// Whether to Enable SSL Certificate. Valid Values: on, off. Valid values: `on`, `off`.
	SslProtocol *string `pulumi:"sslProtocol"`
	// If You Enable HTTPS Here Key.
	SslPub *string `pulumi:"sslPub"`
}

// DomainCertInfoInput is an input type that accepts DomainCertInfoArgs and DomainCertInfoOutput values.
// You can construct a concrete instance of `DomainCertInfoInput` via:
//
//	DomainCertInfoArgs{...}
type DomainCertInfoInput interface {
	pulumi.Input

	ToDomainCertInfoOutput() DomainCertInfoOutput
	ToDomainCertInfoOutputWithContext(context.Context) DomainCertInfoOutput
}

type DomainCertInfoArgs struct {
	// If You Enable HTTPS Here Certificate Name.
	CertName pulumi.StringPtrInput `pulumi:"certName"`
	// Certificate Type. Value Range:
	// * upload: Certificate
	// * cas: Certificate Authority Certificate.
	// * free: Free Certificate.
	CertType pulumi.StringPtrInput `pulumi:"certType"`
	// Private Key. Do Not Enable Certificate without Entering a User Name and Configure Certificates Enter Private Key.
	SslPri pulumi.StringPtrInput `pulumi:"sslPri"`
	// Whether to Enable SSL Certificate. Valid Values: on, off. Valid values: `on`, `off`.
	SslProtocol pulumi.StringPtrInput `pulumi:"sslProtocol"`
	// If You Enable HTTPS Here Key.
	SslPub pulumi.StringPtrInput `pulumi:"sslPub"`
}

func (DomainCertInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainCertInfo)(nil)).Elem()
}

func (i DomainCertInfoArgs) ToDomainCertInfoOutput() DomainCertInfoOutput {
	return i.ToDomainCertInfoOutputWithContext(context.Background())
}

func (i DomainCertInfoArgs) ToDomainCertInfoOutputWithContext(ctx context.Context) DomainCertInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainCertInfoOutput)
}

// DomainCertInfoArrayInput is an input type that accepts DomainCertInfoArray and DomainCertInfoArrayOutput values.
// You can construct a concrete instance of `DomainCertInfoArrayInput` via:
//
//	DomainCertInfoArray{ DomainCertInfoArgs{...} }
type DomainCertInfoArrayInput interface {
	pulumi.Input

	ToDomainCertInfoArrayOutput() DomainCertInfoArrayOutput
	ToDomainCertInfoArrayOutputWithContext(context.Context) DomainCertInfoArrayOutput
}

type DomainCertInfoArray []DomainCertInfoInput

func (DomainCertInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainCertInfo)(nil)).Elem()
}

func (i DomainCertInfoArray) ToDomainCertInfoArrayOutput() DomainCertInfoArrayOutput {
	return i.ToDomainCertInfoArrayOutputWithContext(context.Background())
}

func (i DomainCertInfoArray) ToDomainCertInfoArrayOutputWithContext(ctx context.Context) DomainCertInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainCertInfoArrayOutput)
}

type DomainCertInfoOutput struct{ *pulumi.OutputState }

func (DomainCertInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainCertInfo)(nil)).Elem()
}

func (o DomainCertInfoOutput) ToDomainCertInfoOutput() DomainCertInfoOutput {
	return o
}

func (o DomainCertInfoOutput) ToDomainCertInfoOutputWithContext(ctx context.Context) DomainCertInfoOutput {
	return o
}

// If You Enable HTTPS Here Certificate Name.
func (o DomainCertInfoOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainCertInfo) *string { return v.CertName }).(pulumi.StringPtrOutput)
}

// Certificate Type. Value Range:
// * upload: Certificate
// * cas: Certificate Authority Certificate.
// * free: Free Certificate.
func (o DomainCertInfoOutput) CertType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainCertInfo) *string { return v.CertType }).(pulumi.StringPtrOutput)
}

// Private Key. Do Not Enable Certificate without Entering a User Name and Configure Certificates Enter Private Key.
func (o DomainCertInfoOutput) SslPri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainCertInfo) *string { return v.SslPri }).(pulumi.StringPtrOutput)
}

// Whether to Enable SSL Certificate. Valid Values: on, off. Valid values: `on`, `off`.
func (o DomainCertInfoOutput) SslProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainCertInfo) *string { return v.SslProtocol }).(pulumi.StringPtrOutput)
}

// If You Enable HTTPS Here Key.
func (o DomainCertInfoOutput) SslPub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainCertInfo) *string { return v.SslPub }).(pulumi.StringPtrOutput)
}

type DomainCertInfoArrayOutput struct{ *pulumi.OutputState }

func (DomainCertInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainCertInfo)(nil)).Elem()
}

func (o DomainCertInfoArrayOutput) ToDomainCertInfoArrayOutput() DomainCertInfoArrayOutput {
	return o
}

func (o DomainCertInfoArrayOutput) ToDomainCertInfoArrayOutputWithContext(ctx context.Context) DomainCertInfoArrayOutput {
	return o
}

func (o DomainCertInfoArrayOutput) Index(i pulumi.IntInput) DomainCertInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainCertInfo {
		return vs[0].([]DomainCertInfo)[vs[1].(int)]
	}).(DomainCertInfoOutput)
}

type DomainConfigFunctionArg struct {
	// The name of arg.
	ArgName string `pulumi:"argName"`
	// The value of arg.
	ArgValue string `pulumi:"argValue"`
}

// DomainConfigFunctionArgInput is an input type that accepts DomainConfigFunctionArgArgs and DomainConfigFunctionArgOutput values.
// You can construct a concrete instance of `DomainConfigFunctionArgInput` via:
//
//	DomainConfigFunctionArgArgs{...}
type DomainConfigFunctionArgInput interface {
	pulumi.Input

	ToDomainConfigFunctionArgOutput() DomainConfigFunctionArgOutput
	ToDomainConfigFunctionArgOutputWithContext(context.Context) DomainConfigFunctionArgOutput
}

type DomainConfigFunctionArgArgs struct {
	// The name of arg.
	ArgName pulumi.StringInput `pulumi:"argName"`
	// The value of arg.
	ArgValue pulumi.StringInput `pulumi:"argValue"`
}

func (DomainConfigFunctionArgArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainConfigFunctionArg)(nil)).Elem()
}

func (i DomainConfigFunctionArgArgs) ToDomainConfigFunctionArgOutput() DomainConfigFunctionArgOutput {
	return i.ToDomainConfigFunctionArgOutputWithContext(context.Background())
}

func (i DomainConfigFunctionArgArgs) ToDomainConfigFunctionArgOutputWithContext(ctx context.Context) DomainConfigFunctionArgOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigFunctionArgOutput)
}

// DomainConfigFunctionArgArrayInput is an input type that accepts DomainConfigFunctionArgArray and DomainConfigFunctionArgArrayOutput values.
// You can construct a concrete instance of `DomainConfigFunctionArgArrayInput` via:
//
//	DomainConfigFunctionArgArray{ DomainConfigFunctionArgArgs{...} }
type DomainConfigFunctionArgArrayInput interface {
	pulumi.Input

	ToDomainConfigFunctionArgArrayOutput() DomainConfigFunctionArgArrayOutput
	ToDomainConfigFunctionArgArrayOutputWithContext(context.Context) DomainConfigFunctionArgArrayOutput
}

type DomainConfigFunctionArgArray []DomainConfigFunctionArgInput

func (DomainConfigFunctionArgArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainConfigFunctionArg)(nil)).Elem()
}

func (i DomainConfigFunctionArgArray) ToDomainConfigFunctionArgArrayOutput() DomainConfigFunctionArgArrayOutput {
	return i.ToDomainConfigFunctionArgArrayOutputWithContext(context.Background())
}

func (i DomainConfigFunctionArgArray) ToDomainConfigFunctionArgArrayOutputWithContext(ctx context.Context) DomainConfigFunctionArgArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainConfigFunctionArgArrayOutput)
}

type DomainConfigFunctionArgOutput struct{ *pulumi.OutputState }

func (DomainConfigFunctionArgOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainConfigFunctionArg)(nil)).Elem()
}

func (o DomainConfigFunctionArgOutput) ToDomainConfigFunctionArgOutput() DomainConfigFunctionArgOutput {
	return o
}

func (o DomainConfigFunctionArgOutput) ToDomainConfigFunctionArgOutputWithContext(ctx context.Context) DomainConfigFunctionArgOutput {
	return o
}

// The name of arg.
func (o DomainConfigFunctionArgOutput) ArgName() pulumi.StringOutput {
	return o.ApplyT(func(v DomainConfigFunctionArg) string { return v.ArgName }).(pulumi.StringOutput)
}

// The value of arg.
func (o DomainConfigFunctionArgOutput) ArgValue() pulumi.StringOutput {
	return o.ApplyT(func(v DomainConfigFunctionArg) string { return v.ArgValue }).(pulumi.StringOutput)
}

type DomainConfigFunctionArgArrayOutput struct{ *pulumi.OutputState }

func (DomainConfigFunctionArgArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainConfigFunctionArg)(nil)).Elem()
}

func (o DomainConfigFunctionArgArrayOutput) ToDomainConfigFunctionArgArrayOutput() DomainConfigFunctionArgArrayOutput {
	return o
}

func (o DomainConfigFunctionArgArrayOutput) ToDomainConfigFunctionArgArrayOutputWithContext(ctx context.Context) DomainConfigFunctionArgArrayOutput {
	return o
}

func (o DomainConfigFunctionArgArrayOutput) Index(i pulumi.IntInput) DomainConfigFunctionArgOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainConfigFunctionArg {
		return vs[0].([]DomainConfigFunctionArg)[vs[1].(int)]
	}).(DomainConfigFunctionArgOutput)
}

type DomainSource struct {
	// The Back-to-Source Address.
	Content string `pulumi:"content"`
	// The source status. Valid values: online, offline.
	Enabled *string `pulumi:"enabled"`
	// Port.
	Port int `pulumi:"port"`
	// Priority.
	Priority string `pulumi:"priority"`
	// The Origin Server Type. Valid Values:
	// * ipaddr: IP Source Station
	// * domain: the Domain Name
	// * oss: OSS Bucket as a Source Station.
	Type string `pulumi:"type"`
}

// DomainSourceInput is an input type that accepts DomainSourceArgs and DomainSourceOutput values.
// You can construct a concrete instance of `DomainSourceInput` via:
//
//	DomainSourceArgs{...}
type DomainSourceInput interface {
	pulumi.Input

	ToDomainSourceOutput() DomainSourceOutput
	ToDomainSourceOutputWithContext(context.Context) DomainSourceOutput
}

type DomainSourceArgs struct {
	// The Back-to-Source Address.
	Content pulumi.StringInput `pulumi:"content"`
	// The source status. Valid values: online, offline.
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
	// Port.
	Port pulumi.IntInput `pulumi:"port"`
	// Priority.
	Priority pulumi.StringInput `pulumi:"priority"`
	// The Origin Server Type. Valid Values:
	// * ipaddr: IP Source Station
	// * domain: the Domain Name
	// * oss: OSS Bucket as a Source Station.
	Type pulumi.StringInput `pulumi:"type"`
}

func (DomainSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSource)(nil)).Elem()
}

func (i DomainSourceArgs) ToDomainSourceOutput() DomainSourceOutput {
	return i.ToDomainSourceOutputWithContext(context.Background())
}

func (i DomainSourceArgs) ToDomainSourceOutputWithContext(ctx context.Context) DomainSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSourceOutput)
}

// DomainSourceArrayInput is an input type that accepts DomainSourceArray and DomainSourceArrayOutput values.
// You can construct a concrete instance of `DomainSourceArrayInput` via:
//
//	DomainSourceArray{ DomainSourceArgs{...} }
type DomainSourceArrayInput interface {
	pulumi.Input

	ToDomainSourceArrayOutput() DomainSourceArrayOutput
	ToDomainSourceArrayOutputWithContext(context.Context) DomainSourceArrayOutput
}

type DomainSourceArray []DomainSourceInput

func (DomainSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainSource)(nil)).Elem()
}

func (i DomainSourceArray) ToDomainSourceArrayOutput() DomainSourceArrayOutput {
	return i.ToDomainSourceArrayOutputWithContext(context.Background())
}

func (i DomainSourceArray) ToDomainSourceArrayOutputWithContext(ctx context.Context) DomainSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainSourceArrayOutput)
}

type DomainSourceOutput struct{ *pulumi.OutputState }

func (DomainSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainSource)(nil)).Elem()
}

func (o DomainSourceOutput) ToDomainSourceOutput() DomainSourceOutput {
	return o
}

func (o DomainSourceOutput) ToDomainSourceOutputWithContext(ctx context.Context) DomainSourceOutput {
	return o
}

// The Back-to-Source Address.
func (o DomainSourceOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v DomainSource) string { return v.Content }).(pulumi.StringOutput)
}

// The source status. Valid values: online, offline.
func (o DomainSourceOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSource) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

// Port.
func (o DomainSourceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v DomainSource) int { return v.Port }).(pulumi.IntOutput)
}

// Priority.
func (o DomainSourceOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v DomainSource) string { return v.Priority }).(pulumi.StringOutput)
}

// The Origin Server Type. Valid Values:
// * ipaddr: IP Source Station
// * domain: the Domain Name
// * oss: OSS Bucket as a Source Station.
func (o DomainSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DomainSource) string { return v.Type }).(pulumi.StringOutput)
}

type DomainSourceArrayOutput struct{ *pulumi.OutputState }

func (DomainSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainSource)(nil)).Elem()
}

func (o DomainSourceArrayOutput) ToDomainSourceArrayOutput() DomainSourceArrayOutput {
	return o
}

func (o DomainSourceArrayOutput) ToDomainSourceArrayOutputWithContext(ctx context.Context) DomainSourceArrayOutput {
	return o
}

func (o DomainSourceArrayOutput) Index(i pulumi.IntInput) DomainSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainSource {
		return vs[0].([]DomainSource)[vs[1].(int)]
	}).(DomainSourceOutput)
}

type GetDomainsDomain struct {
	// Certificate Information.
	CertInfos []GetDomainsDomainCertInfo `pulumi:"certInfos"`
	// In Order to Link the CDN Domain Name to Generate a CNAME Domain Name, in the Domain Name Resolution Service Provider at the Acceleration Domain Name CNAME Resolution to the Domain.
	Cname string `pulumi:"cname"`
	// Creation Time.
	CreateTime string `pulumi:"createTime"`
	// Review the Reason for the Failure Is Displayed.
	Description string `pulumi:"description"`
	// Your Domain Name.
	DomainName string `pulumi:"domainName"`
	// Last Modified Date.
	GmtModified string `pulumi:"gmtModified"`
	// The ID of the Domain. Its value is same as Queue Name.
	Id string `pulumi:"id"`
	// The Resource Group ID.
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// the Origin Server Information.
	Sources []GetDomainsDomainSource `pulumi:"sources"`
	// The status of the resource.
	Status string `pulumi:"status"`
}

// GetDomainsDomainInput is an input type that accepts GetDomainsDomainArgs and GetDomainsDomainOutput values.
// You can construct a concrete instance of `GetDomainsDomainInput` via:
//
//	GetDomainsDomainArgs{...}
type GetDomainsDomainInput interface {
	pulumi.Input

	ToGetDomainsDomainOutput() GetDomainsDomainOutput
	ToGetDomainsDomainOutputWithContext(context.Context) GetDomainsDomainOutput
}

type GetDomainsDomainArgs struct {
	// Certificate Information.
	CertInfos GetDomainsDomainCertInfoArrayInput `pulumi:"certInfos"`
	// In Order to Link the CDN Domain Name to Generate a CNAME Domain Name, in the Domain Name Resolution Service Provider at the Acceleration Domain Name CNAME Resolution to the Domain.
	Cname pulumi.StringInput `pulumi:"cname"`
	// Creation Time.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Review the Reason for the Failure Is Displayed.
	Description pulumi.StringInput `pulumi:"description"`
	// Your Domain Name.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Last Modified Date.
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The ID of the Domain. Its value is same as Queue Name.
	Id pulumi.StringInput `pulumi:"id"`
	// The Resource Group ID.
	ResourceGroupId pulumi.StringInput `pulumi:"resourceGroupId"`
	// the Origin Server Information.
	Sources GetDomainsDomainSourceArrayInput `pulumi:"sources"`
	// The status of the resource.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetDomainsDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return i.ToGetDomainsDomainOutputWithContext(context.Background())
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainOutput)
}

// GetDomainsDomainArrayInput is an input type that accepts GetDomainsDomainArray and GetDomainsDomainArrayOutput values.
// You can construct a concrete instance of `GetDomainsDomainArrayInput` via:
//
//	GetDomainsDomainArray{ GetDomainsDomainArgs{...} }
type GetDomainsDomainArrayInput interface {
	pulumi.Input

	ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput
	ToGetDomainsDomainArrayOutputWithContext(context.Context) GetDomainsDomainArrayOutput
}

type GetDomainsDomainArray []GetDomainsDomainInput

func (GetDomainsDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return i.ToGetDomainsDomainArrayOutputWithContext(context.Background())
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainArrayOutput)
}

type GetDomainsDomainOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return o
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return o
}

// Certificate Information.
func (o GetDomainsDomainOutput) CertInfos() GetDomainsDomainCertInfoArrayOutput {
	return o.ApplyT(func(v GetDomainsDomain) []GetDomainsDomainCertInfo { return v.CertInfos }).(GetDomainsDomainCertInfoArrayOutput)
}

// In Order to Link the CDN Domain Name to Generate a CNAME Domain Name, in the Domain Name Resolution Service Provider at the Acceleration Domain Name CNAME Resolution to the Domain.
func (o GetDomainsDomainOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Cname }).(pulumi.StringOutput)
}

// Creation Time.
func (o GetDomainsDomainOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Review the Reason for the Failure Is Displayed.
func (o GetDomainsDomainOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Description }).(pulumi.StringOutput)
}

// Your Domain Name.
func (o GetDomainsDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DomainName }).(pulumi.StringOutput)
}

// Last Modified Date.
func (o GetDomainsDomainOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The ID of the Domain. Its value is same as Queue Name.
func (o GetDomainsDomainOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Id }).(pulumi.StringOutput)
}

// The Resource Group ID.
func (o GetDomainsDomainOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// the Origin Server Information.
func (o GetDomainsDomainOutput) Sources() GetDomainsDomainSourceArrayOutput {
	return o.ApplyT(func(v GetDomainsDomain) []GetDomainsDomainSource { return v.Sources }).(GetDomainsDomainSourceArrayOutput)
}

// The status of the resource.
func (o GetDomainsDomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Status }).(pulumi.StringOutput)
}

type GetDomainsDomainArrayOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) Index(i pulumi.IntInput) GetDomainsDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDomainsDomain {
		return vs[0].([]GetDomainsDomain)[vs[1].(int)]
	}).(GetDomainsDomainOutput)
}

type GetDomainsDomainCertInfo struct {
	// If You Enable HTTPS Here Certificate Name.
	CertName string `pulumi:"certName"`
	// Certificate Type. Value Range: Upload: Certificate. CAS: Certificate Authority Certificate. Free: Free Certificate.
	CertType string `pulumi:"certType"`
	// Whether to Enable SSL Certificate. Valid Values: on, off.
	SslProtocol string `pulumi:"sslProtocol"`
	// If You Enable HTTPS Here Key.
	SslPub string `pulumi:"sslPub"`
}

// GetDomainsDomainCertInfoInput is an input type that accepts GetDomainsDomainCertInfoArgs and GetDomainsDomainCertInfoOutput values.
// You can construct a concrete instance of `GetDomainsDomainCertInfoInput` via:
//
//	GetDomainsDomainCertInfoArgs{...}
type GetDomainsDomainCertInfoInput interface {
	pulumi.Input

	ToGetDomainsDomainCertInfoOutput() GetDomainsDomainCertInfoOutput
	ToGetDomainsDomainCertInfoOutputWithContext(context.Context) GetDomainsDomainCertInfoOutput
}

type GetDomainsDomainCertInfoArgs struct {
	// If You Enable HTTPS Here Certificate Name.
	CertName pulumi.StringInput `pulumi:"certName"`
	// Certificate Type. Value Range: Upload: Certificate. CAS: Certificate Authority Certificate. Free: Free Certificate.
	CertType pulumi.StringInput `pulumi:"certType"`
	// Whether to Enable SSL Certificate. Valid Values: on, off.
	SslProtocol pulumi.StringInput `pulumi:"sslProtocol"`
	// If You Enable HTTPS Here Key.
	SslPub pulumi.StringInput `pulumi:"sslPub"`
}

func (GetDomainsDomainCertInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomainCertInfo)(nil)).Elem()
}

func (i GetDomainsDomainCertInfoArgs) ToGetDomainsDomainCertInfoOutput() GetDomainsDomainCertInfoOutput {
	return i.ToGetDomainsDomainCertInfoOutputWithContext(context.Background())
}

func (i GetDomainsDomainCertInfoArgs) ToGetDomainsDomainCertInfoOutputWithContext(ctx context.Context) GetDomainsDomainCertInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainCertInfoOutput)
}

// GetDomainsDomainCertInfoArrayInput is an input type that accepts GetDomainsDomainCertInfoArray and GetDomainsDomainCertInfoArrayOutput values.
// You can construct a concrete instance of `GetDomainsDomainCertInfoArrayInput` via:
//
//	GetDomainsDomainCertInfoArray{ GetDomainsDomainCertInfoArgs{...} }
type GetDomainsDomainCertInfoArrayInput interface {
	pulumi.Input

	ToGetDomainsDomainCertInfoArrayOutput() GetDomainsDomainCertInfoArrayOutput
	ToGetDomainsDomainCertInfoArrayOutputWithContext(context.Context) GetDomainsDomainCertInfoArrayOutput
}

type GetDomainsDomainCertInfoArray []GetDomainsDomainCertInfoInput

func (GetDomainsDomainCertInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomainCertInfo)(nil)).Elem()
}

func (i GetDomainsDomainCertInfoArray) ToGetDomainsDomainCertInfoArrayOutput() GetDomainsDomainCertInfoArrayOutput {
	return i.ToGetDomainsDomainCertInfoArrayOutputWithContext(context.Background())
}

func (i GetDomainsDomainCertInfoArray) ToGetDomainsDomainCertInfoArrayOutputWithContext(ctx context.Context) GetDomainsDomainCertInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainCertInfoArrayOutput)
}

type GetDomainsDomainCertInfoOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainCertInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomainCertInfo)(nil)).Elem()
}

func (o GetDomainsDomainCertInfoOutput) ToGetDomainsDomainCertInfoOutput() GetDomainsDomainCertInfoOutput {
	return o
}

func (o GetDomainsDomainCertInfoOutput) ToGetDomainsDomainCertInfoOutputWithContext(ctx context.Context) GetDomainsDomainCertInfoOutput {
	return o
}

// If You Enable HTTPS Here Certificate Name.
func (o GetDomainsDomainCertInfoOutput) CertName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainCertInfo) string { return v.CertName }).(pulumi.StringOutput)
}

// Certificate Type. Value Range: Upload: Certificate. CAS: Certificate Authority Certificate. Free: Free Certificate.
func (o GetDomainsDomainCertInfoOutput) CertType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainCertInfo) string { return v.CertType }).(pulumi.StringOutput)
}

// Whether to Enable SSL Certificate. Valid Values: on, off.
func (o GetDomainsDomainCertInfoOutput) SslProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainCertInfo) string { return v.SslProtocol }).(pulumi.StringOutput)
}

// If You Enable HTTPS Here Key.
func (o GetDomainsDomainCertInfoOutput) SslPub() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainCertInfo) string { return v.SslPub }).(pulumi.StringOutput)
}

type GetDomainsDomainCertInfoArrayOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainCertInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomainCertInfo)(nil)).Elem()
}

func (o GetDomainsDomainCertInfoArrayOutput) ToGetDomainsDomainCertInfoArrayOutput() GetDomainsDomainCertInfoArrayOutput {
	return o
}

func (o GetDomainsDomainCertInfoArrayOutput) ToGetDomainsDomainCertInfoArrayOutputWithContext(ctx context.Context) GetDomainsDomainCertInfoArrayOutput {
	return o
}

func (o GetDomainsDomainCertInfoArrayOutput) Index(i pulumi.IntInput) GetDomainsDomainCertInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDomainsDomainCertInfo {
		return vs[0].([]GetDomainsDomainCertInfo)[vs[1].(int)]
	}).(GetDomainsDomainCertInfoOutput)
}

type GetDomainsDomainSource struct {
	// The Back-to-Source Address.
	Content string `pulumi:"content"`
	// State.
	Enabled string `pulumi:"enabled"`
	// Port.
	Port int `pulumi:"port"`
	// Priority.
	Priority string `pulumi:"priority"`
	// the Origin Server Type. Valid Values: Ipaddr: IP Source Station Domain: the Domain Name, See Extra Domain Quota OSS: OSS Bucket as a Source Station.
	Type string `pulumi:"type"`
}

// GetDomainsDomainSourceInput is an input type that accepts GetDomainsDomainSourceArgs and GetDomainsDomainSourceOutput values.
// You can construct a concrete instance of `GetDomainsDomainSourceInput` via:
//
//	GetDomainsDomainSourceArgs{...}
type GetDomainsDomainSourceInput interface {
	pulumi.Input

	ToGetDomainsDomainSourceOutput() GetDomainsDomainSourceOutput
	ToGetDomainsDomainSourceOutputWithContext(context.Context) GetDomainsDomainSourceOutput
}

type GetDomainsDomainSourceArgs struct {
	// The Back-to-Source Address.
	Content pulumi.StringInput `pulumi:"content"`
	// State.
	Enabled pulumi.StringInput `pulumi:"enabled"`
	// Port.
	Port pulumi.IntInput `pulumi:"port"`
	// Priority.
	Priority pulumi.StringInput `pulumi:"priority"`
	// the Origin Server Type. Valid Values: Ipaddr: IP Source Station Domain: the Domain Name, See Extra Domain Quota OSS: OSS Bucket as a Source Station.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetDomainsDomainSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomainSource)(nil)).Elem()
}

func (i GetDomainsDomainSourceArgs) ToGetDomainsDomainSourceOutput() GetDomainsDomainSourceOutput {
	return i.ToGetDomainsDomainSourceOutputWithContext(context.Background())
}

func (i GetDomainsDomainSourceArgs) ToGetDomainsDomainSourceOutputWithContext(ctx context.Context) GetDomainsDomainSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainSourceOutput)
}

// GetDomainsDomainSourceArrayInput is an input type that accepts GetDomainsDomainSourceArray and GetDomainsDomainSourceArrayOutput values.
// You can construct a concrete instance of `GetDomainsDomainSourceArrayInput` via:
//
//	GetDomainsDomainSourceArray{ GetDomainsDomainSourceArgs{...} }
type GetDomainsDomainSourceArrayInput interface {
	pulumi.Input

	ToGetDomainsDomainSourceArrayOutput() GetDomainsDomainSourceArrayOutput
	ToGetDomainsDomainSourceArrayOutputWithContext(context.Context) GetDomainsDomainSourceArrayOutput
}

type GetDomainsDomainSourceArray []GetDomainsDomainSourceInput

func (GetDomainsDomainSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomainSource)(nil)).Elem()
}

func (i GetDomainsDomainSourceArray) ToGetDomainsDomainSourceArrayOutput() GetDomainsDomainSourceArrayOutput {
	return i.ToGetDomainsDomainSourceArrayOutputWithContext(context.Background())
}

func (i GetDomainsDomainSourceArray) ToGetDomainsDomainSourceArrayOutputWithContext(ctx context.Context) GetDomainsDomainSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainSourceArrayOutput)
}

type GetDomainsDomainSourceOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomainSource)(nil)).Elem()
}

func (o GetDomainsDomainSourceOutput) ToGetDomainsDomainSourceOutput() GetDomainsDomainSourceOutput {
	return o
}

func (o GetDomainsDomainSourceOutput) ToGetDomainsDomainSourceOutputWithContext(ctx context.Context) GetDomainsDomainSourceOutput {
	return o
}

// The Back-to-Source Address.
func (o GetDomainsDomainSourceOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Content }).(pulumi.StringOutput)
}

// State.
func (o GetDomainsDomainSourceOutput) Enabled() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Enabled }).(pulumi.StringOutput)
}

// Port.
func (o GetDomainsDomainSourceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) int { return v.Port }).(pulumi.IntOutput)
}

// Priority.
func (o GetDomainsDomainSourceOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Priority }).(pulumi.StringOutput)
}

// the Origin Server Type. Valid Values: Ipaddr: IP Source Station Domain: the Domain Name, See Extra Domain Quota OSS: OSS Bucket as a Source Station.
func (o GetDomainsDomainSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Type }).(pulumi.StringOutput)
}

type GetDomainsDomainSourceArrayOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomainSource)(nil)).Elem()
}

func (o GetDomainsDomainSourceArrayOutput) ToGetDomainsDomainSourceArrayOutput() GetDomainsDomainSourceArrayOutput {
	return o
}

func (o GetDomainsDomainSourceArrayOutput) ToGetDomainsDomainSourceArrayOutputWithContext(ctx context.Context) GetDomainsDomainSourceArrayOutput {
	return o
}

func (o GetDomainsDomainSourceArrayOutput) Index(i pulumi.IntInput) GetDomainsDomainSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDomainsDomainSource {
		return vs[0].([]GetDomainsDomainSource)[vs[1].(int)]
	}).(GetDomainsDomainSourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainCertInfoInput)(nil)).Elem(), DomainCertInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainCertInfoArrayInput)(nil)).Elem(), DomainCertInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigFunctionArgInput)(nil)).Elem(), DomainConfigFunctionArgArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainConfigFunctionArgArrayInput)(nil)).Elem(), DomainConfigFunctionArgArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainSourceInput)(nil)).Elem(), DomainSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainSourceArrayInput)(nil)).Elem(), DomainSourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainInput)(nil)).Elem(), GetDomainsDomainArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainArrayInput)(nil)).Elem(), GetDomainsDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainCertInfoInput)(nil)).Elem(), GetDomainsDomainCertInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainCertInfoArrayInput)(nil)).Elem(), GetDomainsDomainCertInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainSourceInput)(nil)).Elem(), GetDomainsDomainSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainSourceArrayInput)(nil)).Elem(), GetDomainsDomainSourceArray{})
	pulumi.RegisterOutputType(DomainCertInfoOutput{})
	pulumi.RegisterOutputType(DomainCertInfoArrayOutput{})
	pulumi.RegisterOutputType(DomainConfigFunctionArgOutput{})
	pulumi.RegisterOutputType(DomainConfigFunctionArgArrayOutput{})
	pulumi.RegisterOutputType(DomainSourceOutput{})
	pulumi.RegisterOutputType(DomainSourceArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainCertInfoOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainCertInfoArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainSourceOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainSourceArrayOutput{})
}
