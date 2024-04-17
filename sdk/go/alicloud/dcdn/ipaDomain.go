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

// Provides a DCDN Ipa Domain resource.
//
// For information about DCDN Ipa Domain and how to use it, see [What is Ipa Domain](https://www.alibabacloud.com/help/en/doc-detail/130634.html).
//
// > **NOTE:** Available since v1.158.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultInteger, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			_default, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = dcdn.NewIpaDomain(ctx, "example", &dcdn.IpaDomainArgs{
//				DomainName:      pulumi.String(fmt.Sprintf("example-%v.com", defaultInteger.Result)),
//				ResourceGroupId: pulumi.String(_default.Groups[0].Id),
//				Scope:           pulumi.String("overseas"),
//				Status:          pulumi.String("online"),
//				Sources: dcdn.IpaDomainSourceArray{
//					&dcdn.IpaDomainSourceArgs{
//						Content:  pulumi.String("www.alicloud-provider.cn"),
//						Port:     pulumi.Int(8898),
//						Priority: pulumi.String("20"),
//						Type:     pulumi.String("domain"),
//						Weight:   pulumi.Int(10),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// DCDN Ipa Domain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dcdn/ipaDomain:IpaDomain example <domain_name>
// ```
type IpaDomain struct {
	pulumi.CustomResourceState

	// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Sources. See `sources` below.
	Sources IpaDomainSourceArrayOutput `pulumi:"sources"`
	// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewIpaDomain registers a new resource with the given unique name, arguments, and options.
func NewIpaDomain(ctx *pulumi.Context,
	name string, args *IpaDomainArgs, opts ...pulumi.ResourceOption) (*IpaDomain, error) {
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
	var resource IpaDomain
	err := ctx.RegisterResource("alicloud:dcdn/ipaDomain:IpaDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpaDomain gets an existing IpaDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpaDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpaDomainState, opts ...pulumi.ResourceOption) (*IpaDomain, error) {
	var resource IpaDomain
	err := ctx.ReadResource("alicloud:dcdn/ipaDomain:IpaDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpaDomain resources.
type ipaDomainState struct {
	// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
	DomainName *string `pulumi:"domainName"`
	// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
	Scope *string `pulumi:"scope"`
	// Sources. See `sources` below.
	Sources []IpaDomainSource `pulumi:"sources"`
	// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
	Status *string `pulumi:"status"`
}

type IpaDomainState struct {
	// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
	DomainName pulumi.StringPtrInput
	// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
	Scope pulumi.StringPtrInput
	// Sources. See `sources` below.
	Sources IpaDomainSourceArrayInput
	// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
	Status pulumi.StringPtrInput
}

func (IpaDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipaDomainState)(nil)).Elem()
}

type ipaDomainArgs struct {
	// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
	DomainName string `pulumi:"domainName"`
	// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
	Scope *string `pulumi:"scope"`
	// Sources. See `sources` below.
	Sources []IpaDomainSource `pulumi:"sources"`
	// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a IpaDomain resource.
type IpaDomainArgs struct {
	// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
	DomainName pulumi.StringInput
	// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
	Scope pulumi.StringPtrInput
	// Sources. See `sources` below.
	Sources IpaDomainSourceArrayInput
	// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
	Status pulumi.StringPtrInput
}

func (IpaDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipaDomainArgs)(nil)).Elem()
}

type IpaDomainInput interface {
	pulumi.Input

	ToIpaDomainOutput() IpaDomainOutput
	ToIpaDomainOutputWithContext(ctx context.Context) IpaDomainOutput
}

func (*IpaDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**IpaDomain)(nil)).Elem()
}

func (i *IpaDomain) ToIpaDomainOutput() IpaDomainOutput {
	return i.ToIpaDomainOutputWithContext(context.Background())
}

func (i *IpaDomain) ToIpaDomainOutputWithContext(ctx context.Context) IpaDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpaDomainOutput)
}

// IpaDomainArrayInput is an input type that accepts IpaDomainArray and IpaDomainArrayOutput values.
// You can construct a concrete instance of `IpaDomainArrayInput` via:
//
//	IpaDomainArray{ IpaDomainArgs{...} }
type IpaDomainArrayInput interface {
	pulumi.Input

	ToIpaDomainArrayOutput() IpaDomainArrayOutput
	ToIpaDomainArrayOutputWithContext(context.Context) IpaDomainArrayOutput
}

type IpaDomainArray []IpaDomainInput

func (IpaDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpaDomain)(nil)).Elem()
}

func (i IpaDomainArray) ToIpaDomainArrayOutput() IpaDomainArrayOutput {
	return i.ToIpaDomainArrayOutputWithContext(context.Background())
}

func (i IpaDomainArray) ToIpaDomainArrayOutputWithContext(ctx context.Context) IpaDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpaDomainArrayOutput)
}

// IpaDomainMapInput is an input type that accepts IpaDomainMap and IpaDomainMapOutput values.
// You can construct a concrete instance of `IpaDomainMapInput` via:
//
//	IpaDomainMap{ "key": IpaDomainArgs{...} }
type IpaDomainMapInput interface {
	pulumi.Input

	ToIpaDomainMapOutput() IpaDomainMapOutput
	ToIpaDomainMapOutputWithContext(context.Context) IpaDomainMapOutput
}

type IpaDomainMap map[string]IpaDomainInput

func (IpaDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpaDomain)(nil)).Elem()
}

func (i IpaDomainMap) ToIpaDomainMapOutput() IpaDomainMapOutput {
	return i.ToIpaDomainMapOutputWithContext(context.Background())
}

func (i IpaDomainMap) ToIpaDomainMapOutputWithContext(ctx context.Context) IpaDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpaDomainMapOutput)
}

type IpaDomainOutput struct{ *pulumi.OutputState }

func (IpaDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpaDomain)(nil)).Elem()
}

func (o IpaDomainOutput) ToIpaDomainOutput() IpaDomainOutput {
	return o
}

func (o IpaDomainOutput) ToIpaDomainOutputWithContext(ctx context.Context) IpaDomainOutput {
	return o
}

// The domain name to be added to IPA. Wildcard domain names are supported. A wildcard domain name must start with a period (.).
func (o IpaDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpaDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The ID of the resource group. If you do not set this parameter, the system automatically assigns the ID of the default resource group.
func (o IpaDomainOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpaDomain) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The accelerated region. Valid values: `domestic`, `global`, `overseas`.
func (o IpaDomainOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *IpaDomain) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Sources. See `sources` below.
func (o IpaDomainOutput) Sources() IpaDomainSourceArrayOutput {
	return o.ApplyT(func(v *IpaDomain) IpaDomainSourceArrayOutput { return v.Sources }).(IpaDomainSourceArrayOutput)
}

// The status of DCDN Ipa Domain. Valid values: `online`, `offline`. Default to `online`.
func (o IpaDomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IpaDomain) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type IpaDomainArrayOutput struct{ *pulumi.OutputState }

func (IpaDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpaDomain)(nil)).Elem()
}

func (o IpaDomainArrayOutput) ToIpaDomainArrayOutput() IpaDomainArrayOutput {
	return o
}

func (o IpaDomainArrayOutput) ToIpaDomainArrayOutputWithContext(ctx context.Context) IpaDomainArrayOutput {
	return o
}

func (o IpaDomainArrayOutput) Index(i pulumi.IntInput) IpaDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpaDomain {
		return vs[0].([]*IpaDomain)[vs[1].(int)]
	}).(IpaDomainOutput)
}

type IpaDomainMapOutput struct{ *pulumi.OutputState }

func (IpaDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpaDomain)(nil)).Elem()
}

func (o IpaDomainMapOutput) ToIpaDomainMapOutput() IpaDomainMapOutput {
	return o
}

func (o IpaDomainMapOutput) ToIpaDomainMapOutputWithContext(ctx context.Context) IpaDomainMapOutput {
	return o
}

func (o IpaDomainMapOutput) MapIndex(k pulumi.StringInput) IpaDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpaDomain {
		return vs[0].(map[string]*IpaDomain)[vs[1].(string)]
	}).(IpaDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpaDomainInput)(nil)).Elem(), &IpaDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpaDomainArrayInput)(nil)).Elem(), IpaDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpaDomainMapInput)(nil)).Elem(), IpaDomainMap{})
	pulumi.RegisterOutputType(IpaDomainOutput{})
	pulumi.RegisterOutputType(IpaDomainArrayOutput{})
	pulumi.RegisterOutputType(IpaDomainMapOutput{})
}
