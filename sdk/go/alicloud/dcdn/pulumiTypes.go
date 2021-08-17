// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainConfigFunctionArg struct {
	// The name of arg.
	ArgName string `pulumi:"argName"`
	// The value of arg.
	ArgValue string `pulumi:"argValue"`
}

// DomainConfigFunctionArgInput is an input type that accepts DomainConfigFunctionArgArgs and DomainConfigFunctionArgOutput values.
// You can construct a concrete instance of `DomainConfigFunctionArgInput` via:
//
//          DomainConfigFunctionArgArgs{...}
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
//          DomainConfigFunctionArgArray{ DomainConfigFunctionArgArgs{...} }
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
	// The origin address.
	Content string `pulumi:"content"`
	// The port number. Valid values: `443` and `80`. Default to `80`.
	Port *int `pulumi:"port"`
	// The priority of the origin if multiple origins are specified. Default to `20`.
	Priority *string `pulumi:"priority"`
	// The type of the origin. Valid values:
	// `ipaddr`: The origin is configured using an IP address.
	// `domain`: The origin is configured using a domain name.
	// `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
	Type string `pulumi:"type"`
	// The weight of the origin if multiple origins are specified. Default to `10`.
	Weight *string `pulumi:"weight"`
}

// DomainSourceInput is an input type that accepts DomainSourceArgs and DomainSourceOutput values.
// You can construct a concrete instance of `DomainSourceInput` via:
//
//          DomainSourceArgs{...}
type DomainSourceInput interface {
	pulumi.Input

	ToDomainSourceOutput() DomainSourceOutput
	ToDomainSourceOutputWithContext(context.Context) DomainSourceOutput
}

type DomainSourceArgs struct {
	// The origin address.
	Content pulumi.StringInput `pulumi:"content"`
	// The port number. Valid values: `443` and `80`. Default to `80`.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// The priority of the origin if multiple origins are specified. Default to `20`.
	Priority pulumi.StringPtrInput `pulumi:"priority"`
	// The type of the origin. Valid values:
	// `ipaddr`: The origin is configured using an IP address.
	// `domain`: The origin is configured using a domain name.
	// `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
	Type pulumi.StringInput `pulumi:"type"`
	// The weight of the origin if multiple origins are specified. Default to `10`.
	Weight pulumi.StringPtrInput `pulumi:"weight"`
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
//          DomainSourceArray{ DomainSourceArgs{...} }
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

// The origin address.
func (o DomainSourceOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v DomainSource) string { return v.Content }).(pulumi.StringOutput)
}

// The port number. Valid values: `443` and `80`. Default to `80`.
func (o DomainSourceOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainSource) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The priority of the origin if multiple origins are specified. Default to `20`.
func (o DomainSourceOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSource) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// The type of the origin. Valid values:
// `ipaddr`: The origin is configured using an IP address.
// `domain`: The origin is configured using a domain name.
// `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
func (o DomainSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DomainSource) string { return v.Type }).(pulumi.StringOutput)
}

// The weight of the origin if multiple origins are specified. Default to `10`.
func (o DomainSourceOutput) Weight() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainSource) *string { return v.Weight }).(pulumi.StringPtrOutput)
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
	// Indicates the name of the certificate.
	CertName string `pulumi:"certName"`
	// The canonical name (CNAME) of the accelerated domain.
	Cname string `pulumi:"cname"`
	// The reason that causes the review failure.
	Description string `pulumi:"description"`
	// The name of the DCDN Domain.
	DomainName string `pulumi:"domainName"`
	// The time when the accelerated domain was last modified.
	GmtModified string `pulumi:"gmtModified"`
	// The ID of the DCDN Domain.
	Id string `pulumi:"id"`
	// The ID of the resource group.
	ResourceGroupId string `pulumi:"resourceGroupId"`
	// The acceleration region.
	Scope string `pulumi:"scope"`
	// The origin information.
	Sources []GetDomainsDomainSource `pulumi:"sources"`
	// Indicates whether the SSL certificate is enabled.
	SslProtocol string `pulumi:"sslProtocol"`
	// Indicates the public key of the certificate.
	SslPub string `pulumi:"sslPub"`
	// The status of DCDN Domain.
	Status string `pulumi:"status"`
}

// GetDomainsDomainInput is an input type that accepts GetDomainsDomainArgs and GetDomainsDomainOutput values.
// You can construct a concrete instance of `GetDomainsDomainInput` via:
//
//          GetDomainsDomainArgs{...}
type GetDomainsDomainInput interface {
	pulumi.Input

	ToGetDomainsDomainOutput() GetDomainsDomainOutput
	ToGetDomainsDomainOutputWithContext(context.Context) GetDomainsDomainOutput
}

type GetDomainsDomainArgs struct {
	// Indicates the name of the certificate.
	CertName pulumi.StringInput `pulumi:"certName"`
	// The canonical name (CNAME) of the accelerated domain.
	Cname pulumi.StringInput `pulumi:"cname"`
	// The reason that causes the review failure.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the DCDN Domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The time when the accelerated domain was last modified.
	GmtModified pulumi.StringInput `pulumi:"gmtModified"`
	// The ID of the DCDN Domain.
	Id pulumi.StringInput `pulumi:"id"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringInput `pulumi:"resourceGroupId"`
	// The acceleration region.
	Scope pulumi.StringInput `pulumi:"scope"`
	// The origin information.
	Sources GetDomainsDomainSourceArrayInput `pulumi:"sources"`
	// Indicates whether the SSL certificate is enabled.
	SslProtocol pulumi.StringInput `pulumi:"sslProtocol"`
	// Indicates the public key of the certificate.
	SslPub pulumi.StringInput `pulumi:"sslPub"`
	// The status of DCDN Domain.
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
//          GetDomainsDomainArray{ GetDomainsDomainArgs{...} }
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

// Indicates the name of the certificate.
func (o GetDomainsDomainOutput) CertName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CertName }).(pulumi.StringOutput)
}

// The canonical name (CNAME) of the accelerated domain.
func (o GetDomainsDomainOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Cname }).(pulumi.StringOutput)
}

// The reason that causes the review failure.
func (o GetDomainsDomainOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Description }).(pulumi.StringOutput)
}

// The name of the DCDN Domain.
func (o GetDomainsDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DomainName }).(pulumi.StringOutput)
}

// The time when the accelerated domain was last modified.
func (o GetDomainsDomainOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.GmtModified }).(pulumi.StringOutput)
}

// The ID of the DCDN Domain.
func (o GetDomainsDomainOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o GetDomainsDomainOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The acceleration region.
func (o GetDomainsDomainOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Scope }).(pulumi.StringOutput)
}

// The origin information.
func (o GetDomainsDomainOutput) Sources() GetDomainsDomainSourceArrayOutput {
	return o.ApplyT(func(v GetDomainsDomain) []GetDomainsDomainSource { return v.Sources }).(GetDomainsDomainSourceArrayOutput)
}

// Indicates whether the SSL certificate is enabled.
func (o GetDomainsDomainOutput) SslProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.SslProtocol }).(pulumi.StringOutput)
}

// Indicates the public key of the certificate.
func (o GetDomainsDomainOutput) SslPub() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.SslPub }).(pulumi.StringOutput)
}

// The status of DCDN Domain.
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

type GetDomainsDomainSource struct {
	// The origin address.
	Content string `pulumi:"content"`
	// The status of the origin.
	Enabled string `pulumi:"enabled"`
	// The port number.
	Port int `pulumi:"port"`
	// The priority of the origin if multiple origins are specified.
	Priority string `pulumi:"priority"`
	// The type of the origin. Valid values:
	Type string `pulumi:"type"`
	// The weight of the origin if multiple origins are specified.
	Weight string `pulumi:"weight"`
}

// GetDomainsDomainSourceInput is an input type that accepts GetDomainsDomainSourceArgs and GetDomainsDomainSourceOutput values.
// You can construct a concrete instance of `GetDomainsDomainSourceInput` via:
//
//          GetDomainsDomainSourceArgs{...}
type GetDomainsDomainSourceInput interface {
	pulumi.Input

	ToGetDomainsDomainSourceOutput() GetDomainsDomainSourceOutput
	ToGetDomainsDomainSourceOutputWithContext(context.Context) GetDomainsDomainSourceOutput
}

type GetDomainsDomainSourceArgs struct {
	// The origin address.
	Content pulumi.StringInput `pulumi:"content"`
	// The status of the origin.
	Enabled pulumi.StringInput `pulumi:"enabled"`
	// The port number.
	Port pulumi.IntInput `pulumi:"port"`
	// The priority of the origin if multiple origins are specified.
	Priority pulumi.StringInput `pulumi:"priority"`
	// The type of the origin. Valid values:
	Type pulumi.StringInput `pulumi:"type"`
	// The weight of the origin if multiple origins are specified.
	Weight pulumi.StringInput `pulumi:"weight"`
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
//          GetDomainsDomainSourceArray{ GetDomainsDomainSourceArgs{...} }
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

// The origin address.
func (o GetDomainsDomainSourceOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Content }).(pulumi.StringOutput)
}

// The status of the origin.
func (o GetDomainsDomainSourceOutput) Enabled() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Enabled }).(pulumi.StringOutput)
}

// The port number.
func (o GetDomainsDomainSourceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) int { return v.Port }).(pulumi.IntOutput)
}

// The priority of the origin if multiple origins are specified.
func (o GetDomainsDomainSourceOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Priority }).(pulumi.StringOutput)
}

// The type of the origin. Valid values:
func (o GetDomainsDomainSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Type }).(pulumi.StringOutput)
}

// The weight of the origin if multiple origins are specified.
func (o GetDomainsDomainSourceOutput) Weight() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomainSource) string { return v.Weight }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(DomainConfigFunctionArgOutput{})
	pulumi.RegisterOutputType(DomainConfigFunctionArgArrayOutput{})
	pulumi.RegisterOutputType(DomainSourceOutput{})
	pulumi.RegisterOutputType(DomainSourceArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainArrayOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainSourceOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainSourceArrayOutput{})
}
