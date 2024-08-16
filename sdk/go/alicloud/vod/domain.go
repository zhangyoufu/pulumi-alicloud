// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VOD Domain resource.
//
// For information about VOD Domain and how to use it, see [What is Domain](https://www.alibabacloud.com/help/product/29932.html).
//
// > **NOTE:** Available since v1.136.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vod"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vod.NewDomain(ctx, "default", &vod.DomainArgs{
//				DomainName: pulumi.Sprintf("example-%v.com", _default.Result),
//				Scope:      pulumi.String("domestic"),
//				Sources: vod.DomainSourceArray{
//					&vod.DomainSourceArgs{
//						SourceType:    pulumi.String("domain"),
//						SourceContent: pulumi.String("outin-c7405446108111ec9a7100163e0eb78b.oss-cn-beijing.aliyuncs.com"),
//						SourcePort:    pulumi.String("443"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("terraform"),
//					"For":     pulumi.String("example"),
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
// VOD Domain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vod/domain:Domain example <domain_name>
// ```
type Domain struct {
	pulumi.CustomResourceState

	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName pulumi.StringOutput `pulumi:"certName"`
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrOutput `pulumi:"checkUrl"`
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// The description of the domain name for CDN.
	Description pulumi.StringOutput `pulumi:"description"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated pulumi.StringOutput `pulumi:"gmtCreated"`
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified pulumi.StringOutput `pulumi:"gmtModified"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayOutput `pulumi:"sources"`
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol pulumi.StringOutput `pulumi:"sslProtocol"`
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub pulumi.StringOutput `pulumi:"sslPub"`
	// The status of the domain name for CDN. Valid values:
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain pulumi.StringPtrOutput `pulumi:"topLevelDomain"`
	// The weight of the origin server.
	Weight pulumi.StringOutput `pulumi:"weight"`
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
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("alicloud:vod/domain:Domain", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:vod/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName *string `pulumi:"certName"`
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname *string `pulumi:"cname"`
	// The description of the domain name for CDN.
	Description *string `pulumi:"description"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName *string `pulumi:"domainName"`
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated *string `pulumi:"gmtCreated"`
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified *string `pulumi:"gmtModified"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope *string `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources []DomainSource `pulumi:"sources"`
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol *string `pulumi:"sslProtocol"`
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub *string `pulumi:"sslPub"`
	// The status of the domain name for CDN. Valid values:
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags map[string]string `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
	// The weight of the origin server.
	Weight *string `pulumi:"weight"`
}

type DomainState struct {
	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	CertName pulumi.StringPtrInput
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	Cname pulumi.StringPtrInput
	// The description of the domain name for CDN.
	Description pulumi.StringPtrInput
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringPtrInput
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtCreated pulumi.StringPtrInput
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	GmtModified pulumi.StringPtrInput
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrInput
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayInput
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
	SslProtocol pulumi.StringPtrInput
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	SslPub pulumi.StringPtrInput
	// The status of the domain name for CDN. Valid values:
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.StringMapInput
	// The top-level domain name.
	TopLevelDomain pulumi.StringPtrInput
	// The weight of the origin server.
	Weight pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// The URL that is used for health checks.
	CheckUrl *string `pulumi:"checkUrl"`
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName string `pulumi:"domainName"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope *string `pulumi:"scope"`
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources []DomainSource `pulumi:"sources"`
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags map[string]string `pulumi:"tags"`
	// The top-level domain name.
	TopLevelDomain *string `pulumi:"topLevelDomain"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The URL that is used for health checks.
	CheckUrl pulumi.StringPtrInput
	// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
	DomainName pulumi.StringInput
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	Scope pulumi.StringPtrInput
	// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
	Sources DomainSourceArrayInput
	// A mapping of tags to assign to the resource.
	// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
	// * `Value`: It can be up to 128 characters in length. It can be a null string.
	Tags pulumi.StringMapInput
	// The top-level domain name.
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

// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
func (o DomainOutput) CertName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CertName }).(pulumi.StringOutput)
}

// The URL that is used for health checks.
func (o DomainOutput) CheckUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.CheckUrl }).(pulumi.StringPtrOutput)
}

// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
func (o DomainOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

// The description of the domain name for CDN.
func (o DomainOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The domain name for CDN that you want to add to ApsaraVideo VOD. Wildcard domain names are supported. Start the domain name with a period (.). Example: `.example.com.`.
func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o DomainOutput) GmtCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.GmtCreated }).(pulumi.StringOutput)
}

// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o DomainOutput) GmtModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.GmtModified }).(pulumi.StringOutput)
}

// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
func (o DomainOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The information about the address of the origin server. For more information about the Sources parameter, See the following `Block sources`.
func (o DomainOutput) Sources() DomainSourceArrayOutput {
	return o.ApplyT(func(v *Domain) DomainSourceArrayOutput { return v.Sources }).(DomainSourceArrayOutput)
}

// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values: `on`,`off`.
func (o DomainOutput) SslProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SslProtocol }).(pulumi.StringOutput)
}

// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
func (o DomainOutput) SslPub() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.SslPub }).(pulumi.StringOutput)
}

// The status of the domain name for CDN. Valid values:
func (o DomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
// * `Key`: It can be up to 64 characters in length. It cannot be a null string.
// * `Value`: It can be up to 128 characters in length. It can be a null string.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The top-level domain name.
func (o DomainOutput) TopLevelDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.TopLevelDomain }).(pulumi.StringPtrOutput)
}

// The weight of the origin server.
func (o DomainOutput) Weight() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Weight }).(pulumi.StringOutput)
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
