// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DCDN Domain resource.
//
// Full station accelerated domain name.
//
// For information about DCDN Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/en/doc-detail/130628.htm).
//
// > **NOTE:** Available since v1.94.0.
//
// > **NOTE:** Field `forceSet`, `securityToken` has been removed from provider version 1.227.1.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			domainName := "tf-example.com"
//			if param := cfg.Get("domainName"); param != "" {
//				domainName = param
//			}
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcdn.NewDomain(ctx, "example", &dcdn.DomainArgs{
//				DomainName: pulumi.String(fmt.Sprintf("%v-%v", domainName, _default.Result)),
//				Scope:      pulumi.String("overseas"),
//				Sources: dcdn.DomainSourceArray{
//					&dcdn.DomainSourceArgs{
//						Content:  pulumi.String("1.1.1.1"),
//						Port:     pulumi.Int(80),
//						Priority: pulumi.String("20"),
//						Type:     pulumi.String("ipaddr"),
//						Weight:   pulumi.String("10"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// DCDN Domain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dcdn/domain:Domain example <id>
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
	CertId pulumi.StringOutput `pulumi:"certId"`
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
	CertName pulumi.StringOutput `pulumi:"certName"`
	// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
	CertRegion pulumi.StringOutput `pulumi:"certRegion"`
	// The certificate type.
	CertType pulumi.StringOutput `pulumi:"certType"`
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrOutput `pulumi:"checkUrl"`
	// The CNAME domain name corresponding to the accelerated domain name.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// The time when the accelerated domain name was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
	Env pulumi.StringPtrOutput `pulumi:"env"`
	// Computing service type. Valid values:
	FunctionType pulumi.StringPtrOutput `pulumi:"functionType"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The Acceleration scen. Supported:
	Scene pulumi.StringPtrOutput `pulumi:"scene"`
	// The region where the acceleration service is deployed. Valid values:
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Source  See `sources` below.
	Sources DomainSourceArrayOutput `pulumi:"sources"`
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	SslPri pulumi.StringPtrOutput `pulumi:"sslPri"`
	// Specifies whether to enable the SSL certificate. Valid values:
	SslProtocol pulumi.StringPtrOutput `pulumi:"sslProtocol"`
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	SslPub pulumi.StringOutput `pulumi:"sslPub"`
	// The status of the domain name. Valid values:
	Status pulumi.StringOutput `pulumi:"status"`
	// The tag of the resource
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The top-level domain.
	TopLevelDomain pulumi.StringPtrOutput `pulumi:"topLevelDomain"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.SslPri != nil {
		args.SslPri = pulumi.ToSecret(args.SslPri).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sslPri",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("alicloud:dcdn/domain:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("alicloud:dcdn/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
	CertId *string `pulumi:"certId"`
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
	CertName *string `pulumi:"certName"`
	// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
	CertRegion *string `pulumi:"certRegion"`
	// The certificate type.
	CertType *string `pulumi:"certType"`
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The CNAME domain name corresponding to the accelerated domain name.
	Cname *string `pulumi:"cname"`
	// The time when the accelerated domain name was created.
	CreateTime *string `pulumi:"createTime"`
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	DomainName *string `pulumi:"domainName"`
	// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
	Env *string `pulumi:"env"`
	// Computing service type. Valid values:
	FunctionType *string `pulumi:"functionType"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The Acceleration scen. Supported:
	Scene *string `pulumi:"scene"`
	// The region where the acceleration service is deployed. Valid values:
	Scope *string `pulumi:"scope"`
	// Source  See `sources` below.
	Sources []DomainSource `pulumi:"sources"`
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	SslPri *string `pulumi:"sslPri"`
	// Specifies whether to enable the SSL certificate. Valid values:
	SslProtocol *string `pulumi:"sslProtocol"`
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	SslPub *string `pulumi:"sslPub"`
	// The status of the domain name. Valid values:
	Status *string `pulumi:"status"`
	// The tag of the resource
	Tags map[string]interface{} `pulumi:"tags"`
	// The top-level domain.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
}

type DomainState struct {
	// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
	CertId pulumi.StringPtrInput
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
	CertName pulumi.StringPtrInput
	// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
	CertRegion pulumi.StringPtrInput
	// The certificate type.
	CertType pulumi.StringPtrInput
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The CNAME domain name corresponding to the accelerated domain name.
	Cname pulumi.StringPtrInput
	// The time when the accelerated domain name was created.
	CreateTime pulumi.StringPtrInput
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	DomainName pulumi.StringPtrInput
	// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
	Env pulumi.StringPtrInput
	// Computing service type. Valid values:
	FunctionType pulumi.StringPtrInput
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The Acceleration scen. Supported:
	Scene pulumi.StringPtrInput
	// The region where the acceleration service is deployed. Valid values:
	Scope pulumi.StringPtrInput
	// Source  See `sources` below.
	Sources DomainSourceArrayInput
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	SslPri pulumi.StringPtrInput
	// Specifies whether to enable the SSL certificate. Valid values:
	SslProtocol pulumi.StringPtrInput
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	SslPub pulumi.StringPtrInput
	// The status of the domain name. Valid values:
	Status pulumi.StringPtrInput
	// The tag of the resource
	Tags pulumi.MapInput
	// The top-level domain.
	TopLevelDomain pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
	CertId *string `pulumi:"certId"`
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
	CertName *string `pulumi:"certName"`
	// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
	CertRegion *string `pulumi:"certRegion"`
	// The certificate type.
	CertType *string `pulumi:"certType"`
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	DomainName string `pulumi:"domainName"`
	// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
	Env *string `pulumi:"env"`
	// Computing service type. Valid values:
	FunctionType *string `pulumi:"functionType"`
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The Acceleration scen. Supported:
	Scene *string `pulumi:"scene"`
	// The region where the acceleration service is deployed. Valid values:
	Scope *string `pulumi:"scope"`
	// Source  See `sources` below.
	Sources []DomainSource `pulumi:"sources"`
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	SslPri *string `pulumi:"sslPri"`
	// Specifies whether to enable the SSL certificate. Valid values:
	SslProtocol *string `pulumi:"sslProtocol"`
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	SslPub *string `pulumi:"sslPub"`
	// The status of the domain name. Valid values:
	Status *string `pulumi:"status"`
	// The tag of the resource
	Tags map[string]interface{} `pulumi:"tags"`
	// The top-level domain.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
	CertId pulumi.StringPtrInput
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
	CertName pulumi.StringPtrInput
	// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
	CertRegion pulumi.StringPtrInput
	// The certificate type.
	CertType pulumi.StringPtrInput
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
	DomainName pulumi.StringInput
	// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
	Env pulumi.StringPtrInput
	// Computing service type. Valid values:
	FunctionType pulumi.StringPtrInput
	// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The Acceleration scen. Supported:
	Scene pulumi.StringPtrInput
	// The region where the acceleration service is deployed. Valid values:
	Scope pulumi.StringPtrInput
	// Source  See `sources` below.
	Sources DomainSourceArrayInput
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	SslPri pulumi.StringPtrInput
	// Specifies whether to enable the SSL certificate. Valid values:
	SslProtocol pulumi.StringPtrInput
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	SslPub pulumi.StringPtrInput
	// The status of the domain name. Valid values:
	Status pulumi.StringPtrInput
	// The tag of the resource
	Tags pulumi.MapInput
	// The top-level domain.
	TopLevelDomain pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

// DomainArrayInput is an input type that accepts DomainArray and DomainArrayOutput values.
// You can construct a concrete instance of `DomainArrayInput` via:
//
//	DomainArray{ DomainArgs{...} }
type DomainArrayInput interface {
	pulumi.Input

	ToDomainArrayOutput() DomainArrayOutput
	ToDomainArrayOutputWithContext(context.Context) DomainArrayOutput
}

type DomainArray []DomainInput

func (DomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (i DomainArray) ToDomainArrayOutput() DomainArrayOutput {
	return i.ToDomainArrayOutputWithContext(context.Background())
}

func (i DomainArray) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainArrayOutput)
}

// DomainMapInput is an input type that accepts DomainMap and DomainMapOutput values.
// You can construct a concrete instance of `DomainMapInput` via:
//
//	DomainMap{ "key": DomainArgs{...} }
type DomainMapInput interface {
	pulumi.Input

	ToDomainMapOutput() DomainMapOutput
	ToDomainMapOutputWithContext(context.Context) DomainMapOutput
}

type DomainMap map[string]DomainInput

func (DomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (i DomainMap) ToDomainMapOutput() DomainMapOutput {
	return i.ToDomainMapOutputWithContext(context.Background())
}

func (i DomainMap) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMapOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// The certificate ID. This parameter is required and valid only when `CertType` is set to `cas`. If you specify this parameter, an existing certificate is used.
func (o DomainOutput) CertId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CertId }).(pulumi.StringOutput)
}

// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when `CertType` is set to `upload`.
func (o DomainOutput) CertName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CertName }).(pulumi.StringOutput)
}

// The region of the SSL certificate. This parameter takes effect only when `CertType` is set to `cas`. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou** and **ap-southeast-1**.
func (o DomainOutput) CertRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CertRegion }).(pulumi.StringOutput)
}

// The certificate type.
func (o DomainOutput) CertType() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CertType }).(pulumi.StringOutput)
}

// The URL that is used for health checks.
func (o DomainOutput) CheckUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.CheckUrl }).(pulumi.StringPtrOutput)
}

// The CNAME domain name corresponding to the accelerated domain name.
func (o DomainOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

// The time when the accelerated domain name was created.
func (o DomainOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The accelerated domain name. You can specify multiple domain names and separate them with commas (,). You can specify up to 500 domain names in each request. The query results of multiple domain names are aggregated. If you do not specify this parameter, data of all accelerated domain names under your account is queried.
func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Specifies whether the certificate is issued in canary releases. If you set this parameter to `staging`, the certificate is issued in canary releases. If you do not specify this parameter or set this parameter to other values, the certificate is officially issued.
func (o DomainOutput) Env() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Env }).(pulumi.StringPtrOutput)
}

// Computing service type. Valid values:
func (o DomainOutput) FunctionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.FunctionType }).(pulumi.StringPtrOutput)
}

// The ID of the resource group. If you do not specify a value for this parameter, the system automatically assigns the ID of the default resource group.
func (o DomainOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The Acceleration scen. Supported:
func (o DomainOutput) Scene() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Scene }).(pulumi.StringPtrOutput)
}

// The region where the acceleration service is deployed. Valid values:
func (o DomainOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Source  See `sources` below.
func (o DomainOutput) Sources() DomainSourceArrayOutput {
	return o.ApplyT(func(v *Domain) DomainSourceArrayOutput { return v.Sources }).(DomainSourceArrayOutput)
}

// The private key. Specify the private key only if you want to enable the SSL certificate.
func (o DomainOutput) SslPri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.SslPri }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable the SSL certificate. Valid values:
func (o DomainOutput) SslProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.SslProtocol }).(pulumi.StringPtrOutput)
}

// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
func (o DomainOutput) SslPub() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SslPub }).(pulumi.StringOutput)
}

// The status of the domain name. Valid values:
func (o DomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tag of the resource
func (o DomainOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Domain) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The top-level domain.
func (o DomainOutput) TopLevelDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.TopLevelDomain }).(pulumi.StringPtrOutput)
}

type DomainArrayOutput struct{ *pulumi.OutputState }

func (DomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Domain)(nil)).Elem()
}

func (o DomainArrayOutput) ToDomainArrayOutput() DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) ToDomainArrayOutputWithContext(ctx context.Context) DomainArrayOutput {
	return o
}

func (o DomainArrayOutput) Index(i pulumi.IntInput) DomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].([]*Domain)[vs[1].(int)]
	}).(DomainOutput)
}

type DomainMapOutput struct{ *pulumi.OutputState }

func (DomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Domain)(nil)).Elem()
}

func (o DomainMapOutput) ToDomainMapOutput() DomainMapOutput {
	return o
}

func (o DomainMapOutput) ToDomainMapOutputWithContext(ctx context.Context) DomainMapOutput {
	return o
}

func (o DomainMapOutput) MapIndex(k pulumi.StringInput) DomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Domain {
		return vs[0].(map[string]*Domain)[vs[1].(string)]
	}).(DomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainArrayInput)(nil)).Elem(), DomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainMapInput)(nil)).Elem(), DomainMap{})
	pulumi.RegisterOutputType(DomainOutput{})
	pulumi.RegisterOutputType(DomainArrayOutput{})
	pulumi.RegisterOutputType(DomainMapOutput{})
}
