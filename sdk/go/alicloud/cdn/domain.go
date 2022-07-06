// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Domain struct {
	pulumi.CustomResourceState

	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	AuthConfig DomainAuthConfigPtrOutput `pulumi:"authConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	BlockIps pulumi.StringArrayOutput `pulumi:"blockIps"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CacheConfigs DomainCacheConfigArrayOutput `pulumi:"cacheConfigs"`
	CdnType      pulumi.StringOutput          `pulumi:"cdnType"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CertificateConfig DomainCertificateConfigPtrOutput `pulumi:"certificateConfig"`
	DomainName        pulumi.StringOutput              `pulumi:"domainName"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	HttpHeaderConfigs DomainHttpHeaderConfigArrayOutput `pulumi:"httpHeaderConfigs"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	OptimizeEnable pulumi.StringPtrOutput `pulumi:"optimizeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	Page404Config DomainPage404ConfigPtrOutput `pulumi:"page404Config"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	PageCompressEnable pulumi.StringPtrOutput `pulumi:"pageCompressEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ParameterFilterConfig DomainParameterFilterConfigPtrOutput `pulumi:"parameterFilterConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	RangeEnable pulumi.StringPtrOutput `pulumi:"rangeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ReferConfig DomainReferConfigPtrOutput `pulumi:"referConfig"`
	Scope       pulumi.StringOutput        `pulumi:"scope"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `port` argument instead.
	SourcePort pulumi.IntPtrOutput `pulumi:"sourcePort"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `type` argument instead.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` argument instead.
	Sources pulumi.StringArrayOutput `pulumi:"sources"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	VideoSeekEnable pulumi.StringPtrOutput `pulumi:"videoSeekEnable"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CdnType == nil {
		return nil, errors.New("invalid value for required argument 'CdnType'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource Domain
	err := ctx.RegisterResource("alicloud:cdn/domain:Domain", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:cdn/domain:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	AuthConfig *DomainAuthConfig `pulumi:"authConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	BlockIps []string `pulumi:"blockIps"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CacheConfigs []DomainCacheConfig `pulumi:"cacheConfigs"`
	CdnType      *string             `pulumi:"cdnType"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CertificateConfig *DomainCertificateConfig `pulumi:"certificateConfig"`
	DomainName        *string                  `pulumi:"domainName"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	HttpHeaderConfigs []DomainHttpHeaderConfig `pulumi:"httpHeaderConfigs"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	OptimizeEnable *string `pulumi:"optimizeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	Page404Config *DomainPage404Config `pulumi:"page404Config"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	PageCompressEnable *string `pulumi:"pageCompressEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ParameterFilterConfig *DomainParameterFilterConfig `pulumi:"parameterFilterConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	RangeEnable *string `pulumi:"rangeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ReferConfig *DomainReferConfig `pulumi:"referConfig"`
	Scope       *string            `pulumi:"scope"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `port` argument instead.
	SourcePort *int `pulumi:"sourcePort"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `type` argument instead.
	SourceType *string `pulumi:"sourceType"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` argument instead.
	Sources []string `pulumi:"sources"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	VideoSeekEnable *string `pulumi:"videoSeekEnable"`
}

type DomainState struct {
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	AuthConfig DomainAuthConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	BlockIps pulumi.StringArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CacheConfigs DomainCacheConfigArrayInput
	CdnType      pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CertificateConfig DomainCertificateConfigPtrInput
	DomainName        pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	HttpHeaderConfigs DomainHttpHeaderConfigArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	OptimizeEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	Page404Config DomainPage404ConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	PageCompressEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ParameterFilterConfig DomainParameterFilterConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	RangeEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ReferConfig DomainReferConfigPtrInput
	Scope       pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `port` argument instead.
	SourcePort pulumi.IntPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `type` argument instead.
	SourceType pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` argument instead.
	Sources pulumi.StringArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	VideoSeekEnable pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	AuthConfig *DomainAuthConfig `pulumi:"authConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	BlockIps []string `pulumi:"blockIps"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CacheConfigs []DomainCacheConfig `pulumi:"cacheConfigs"`
	CdnType      string              `pulumi:"cdnType"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CertificateConfig *DomainCertificateConfig `pulumi:"certificateConfig"`
	DomainName        string                   `pulumi:"domainName"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	HttpHeaderConfigs []DomainHttpHeaderConfig `pulumi:"httpHeaderConfigs"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	OptimizeEnable *string `pulumi:"optimizeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	Page404Config *DomainPage404Config `pulumi:"page404Config"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	PageCompressEnable *string `pulumi:"pageCompressEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ParameterFilterConfig *DomainParameterFilterConfig `pulumi:"parameterFilterConfig"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	RangeEnable *string `pulumi:"rangeEnable"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ReferConfig *DomainReferConfig `pulumi:"referConfig"`
	Scope       *string            `pulumi:"scope"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `port` argument instead.
	SourcePort *int `pulumi:"sourcePort"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `type` argument instead.
	SourceType *string `pulumi:"sourceType"`
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` argument instead.
	Sources []string `pulumi:"sources"`
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	VideoSeekEnable *string `pulumi:"videoSeekEnable"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	AuthConfig DomainAuthConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	BlockIps pulumi.StringArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CacheConfigs DomainCacheConfigArrayInput
	CdnType      pulumi.StringInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	CertificateConfig DomainCertificateConfigPtrInput
	DomainName        pulumi.StringInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	HttpHeaderConfigs DomainHttpHeaderConfigArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	OptimizeEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	Page404Config DomainPage404ConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	PageCompressEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ParameterFilterConfig DomainParameterFilterConfigPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	RangeEnable pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	ReferConfig DomainReferConfigPtrInput
	Scope       pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `port` argument instead.
	SourcePort pulumi.IntPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` block `type` argument instead.
	SourceType pulumi.StringPtrInput
	// Deprecated: Use `alicloud_cdn_domain_new` configuration `sources` argument instead.
	Sources pulumi.StringArrayInput
	// Deprecated: Use `alicloud_cdn_domain_config` configuration `function_name` and `function_args` arguments instead.
	VideoSeekEnable pulumi.StringPtrInput
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
//          DomainArray{ DomainArgs{...} }
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
//          DomainMap{ "key": DomainArgs{...} }
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
