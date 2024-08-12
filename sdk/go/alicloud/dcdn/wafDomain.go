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

// Provides a DCDN Waf Domain resource.
//
// For information about DCDN Waf Domain and how to use it, see [What is Waf Domain](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-batchsetdcdnwafdomainconfigs).
//
// > **NOTE:** Available since v1.185.0.
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
//			example, err := dcdn.NewDomain(ctx, "example", &dcdn.DomainArgs{
//				DomainName: pulumi.Sprintf("%v-%v", domainName, _default.Result),
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
//			_, err = dcdn.NewWafDomain(ctx, "example", &dcdn.WafDomainArgs{
//				DomainName:  example.DomainName,
//				ClientIpTag: pulumi.String("X-Forwarded-For"),
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
// DCDN Waf Domain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dcdn/wafDomain:WafDomain example <domain_name>
// ```
type WafDomain struct {
	pulumi.CustomResourceState

	// The client ip tag.
	ClientIpTag pulumi.StringPtrOutput `pulumi:"clientIpTag"`
	// The accelerated domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewWafDomain registers a new resource with the given unique name, arguments, and options.
func NewWafDomain(ctx *pulumi.Context,
	name string, args *WafDomainArgs, opts ...pulumi.ResourceOption) (*WafDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WafDomain
	err := ctx.RegisterResource("alicloud:dcdn/wafDomain:WafDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafDomain gets an existing WafDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafDomainState, opts ...pulumi.ResourceOption) (*WafDomain, error) {
	var resource WafDomain
	err := ctx.ReadResource("alicloud:dcdn/wafDomain:WafDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafDomain resources.
type wafDomainState struct {
	// The client ip tag.
	ClientIpTag *string `pulumi:"clientIpTag"`
	// The accelerated domain name.
	DomainName *string `pulumi:"domainName"`
}

type WafDomainState struct {
	// The client ip tag.
	ClientIpTag pulumi.StringPtrInput
	// The accelerated domain name.
	DomainName pulumi.StringPtrInput
}

func (WafDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafDomainState)(nil)).Elem()
}

type wafDomainArgs struct {
	// The client ip tag.
	ClientIpTag *string `pulumi:"clientIpTag"`
	// The accelerated domain name.
	DomainName string `pulumi:"domainName"`
}

// The set of arguments for constructing a WafDomain resource.
type WafDomainArgs struct {
	// The client ip tag.
	ClientIpTag pulumi.StringPtrInput
	// The accelerated domain name.
	DomainName pulumi.StringInput
}

func (WafDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafDomainArgs)(nil)).Elem()
}

type WafDomainInput interface {
	pulumi.Input

	ToWafDomainOutput() WafDomainOutput
	ToWafDomainOutputWithContext(ctx context.Context) WafDomainOutput
}

func (*WafDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**WafDomain)(nil)).Elem()
}

func (i *WafDomain) ToWafDomainOutput() WafDomainOutput {
	return i.ToWafDomainOutputWithContext(context.Background())
}

func (i *WafDomain) ToWafDomainOutputWithContext(ctx context.Context) WafDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafDomainOutput)
}

// WafDomainArrayInput is an input type that accepts WafDomainArray and WafDomainArrayOutput values.
// You can construct a concrete instance of `WafDomainArrayInput` via:
//
//	WafDomainArray{ WafDomainArgs{...} }
type WafDomainArrayInput interface {
	pulumi.Input

	ToWafDomainArrayOutput() WafDomainArrayOutput
	ToWafDomainArrayOutputWithContext(context.Context) WafDomainArrayOutput
}

type WafDomainArray []WafDomainInput

func (WafDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafDomain)(nil)).Elem()
}

func (i WafDomainArray) ToWafDomainArrayOutput() WafDomainArrayOutput {
	return i.ToWafDomainArrayOutputWithContext(context.Background())
}

func (i WafDomainArray) ToWafDomainArrayOutputWithContext(ctx context.Context) WafDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafDomainArrayOutput)
}

// WafDomainMapInput is an input type that accepts WafDomainMap and WafDomainMapOutput values.
// You can construct a concrete instance of `WafDomainMapInput` via:
//
//	WafDomainMap{ "key": WafDomainArgs{...} }
type WafDomainMapInput interface {
	pulumi.Input

	ToWafDomainMapOutput() WafDomainMapOutput
	ToWafDomainMapOutputWithContext(context.Context) WafDomainMapOutput
}

type WafDomainMap map[string]WafDomainInput

func (WafDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafDomain)(nil)).Elem()
}

func (i WafDomainMap) ToWafDomainMapOutput() WafDomainMapOutput {
	return i.ToWafDomainMapOutputWithContext(context.Background())
}

func (i WafDomainMap) ToWafDomainMapOutputWithContext(ctx context.Context) WafDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafDomainMapOutput)
}

type WafDomainOutput struct{ *pulumi.OutputState }

func (WafDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafDomain)(nil)).Elem()
}

func (o WafDomainOutput) ToWafDomainOutput() WafDomainOutput {
	return o
}

func (o WafDomainOutput) ToWafDomainOutputWithContext(ctx context.Context) WafDomainOutput {
	return o
}

// The client ip tag.
func (o WafDomainOutput) ClientIpTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WafDomain) pulumi.StringPtrOutput { return v.ClientIpTag }).(pulumi.StringPtrOutput)
}

// The accelerated domain name.
func (o WafDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *WafDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

type WafDomainArrayOutput struct{ *pulumi.OutputState }

func (WafDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafDomain)(nil)).Elem()
}

func (o WafDomainArrayOutput) ToWafDomainArrayOutput() WafDomainArrayOutput {
	return o
}

func (o WafDomainArrayOutput) ToWafDomainArrayOutputWithContext(ctx context.Context) WafDomainArrayOutput {
	return o
}

func (o WafDomainArrayOutput) Index(i pulumi.IntInput) WafDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafDomain {
		return vs[0].([]*WafDomain)[vs[1].(int)]
	}).(WafDomainOutput)
}

type WafDomainMapOutput struct{ *pulumi.OutputState }

func (WafDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafDomain)(nil)).Elem()
}

func (o WafDomainMapOutput) ToWafDomainMapOutput() WafDomainMapOutput {
	return o
}

func (o WafDomainMapOutput) ToWafDomainMapOutputWithContext(ctx context.Context) WafDomainMapOutput {
	return o
}

func (o WafDomainMapOutput) MapIndex(k pulumi.StringInput) WafDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafDomain {
		return vs[0].(map[string]*WafDomain)[vs[1].(string)]
	}).(WafDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafDomainInput)(nil)).Elem(), &WafDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafDomainArrayInput)(nil)).Elem(), WafDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafDomainMapInput)(nil)).Elem(), WafDomainMap{})
	pulumi.RegisterOutputType(WafDomainOutput{})
	pulumi.RegisterOutputType(WafDomainArrayOutput{})
	pulumi.RegisterOutputType(WafDomainMapOutput{})
}
