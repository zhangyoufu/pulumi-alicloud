// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Anti-DDoS Pro Domain Resource resource.
//
// For information about Anti-DDoS Pro Domain Resource and how to use it, see [What is Domain Resource](https://www.alibabacloud.com/help/en/doc-detail/157463.htm).
//
// > **NOTE:** Available in v1.123.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ddos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ddos.NewDomainResource(ctx, "example", &ddos.DomainResourceArgs{
//				Domain:   pulumi.String("tftestacc1234.abc"),
//				HttpsExt: pulumi.String("{\"Http2\":1,\"Http2https\":0，\"Https2http\":0}"),
//				InstanceIds: pulumi.StringArray{
//					pulumi.String("ddoscoo-cn-6ja1rl4j****"),
//				},
//				ProxyTypes: ddos.DomainResourceProxyTypeArray{
//					&ddos.DomainResourceProxyTypeArgs{
//						ProxyPorts: pulumi.IntArray{
//							pulumi.Int(443),
//						},
//						ProxyType: pulumi.String("https"),
//					},
//				},
//				RealServers: pulumi.StringArray{
//					pulumi.String("177.167.32.11"),
//				},
//				RsType: pulumi.Int(0),
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
// Anti-DDoS Pro Domain Resource can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ddos/domainResource:DomainResource example <domain>
//
// ```
type DomainResource struct {
	pulumi.CustomResourceState

	// The domain name of the website that you want to add to the instance.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
	// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
	HttpsExt    pulumi.StringOutput      `pulumi:"httpsExt"`
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// Protocol type and port number information.
	ProxyTypes DomainResourceProxyTypeArrayOutput `pulumi:"proxyTypes"`
	// the IP address. This field is required and must be a string array.
	RealServers pulumi.StringArrayOutput `pulumi:"realServers"`
	// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	RsType pulumi.IntOutput `pulumi:"rsType"`
}

// NewDomainResource registers a new resource with the given unique name, arguments, and options.
func NewDomainResource(ctx *pulumi.Context,
	name string, args *DomainResourceArgs, opts ...pulumi.ResourceOption) (*DomainResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.InstanceIds == nil {
		return nil, errors.New("invalid value for required argument 'InstanceIds'")
	}
	if args.ProxyTypes == nil {
		return nil, errors.New("invalid value for required argument 'ProxyTypes'")
	}
	if args.RealServers == nil {
		return nil, errors.New("invalid value for required argument 'RealServers'")
	}
	if args.RsType == nil {
		return nil, errors.New("invalid value for required argument 'RsType'")
	}
	var resource DomainResource
	err := ctx.RegisterResource("alicloud:ddos/domainResource:DomainResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainResource gets an existing DomainResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainResourceState, opts ...pulumi.ResourceOption) (*DomainResource, error) {
	var resource DomainResource
	err := ctx.ReadResource("alicloud:ddos/domainResource:DomainResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainResource resources.
type domainResourceState struct {
	// The domain name of the website that you want to add to the instance.
	Domain *string `pulumi:"domain"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
	// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
	HttpsExt    *string  `pulumi:"httpsExt"`
	InstanceIds []string `pulumi:"instanceIds"`
	// Protocol type and port number information.
	ProxyTypes []DomainResourceProxyType `pulumi:"proxyTypes"`
	// the IP address. This field is required and must be a string array.
	RealServers []string `pulumi:"realServers"`
	// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	RsType *int `pulumi:"rsType"`
}

type DomainResourceState struct {
	// The domain name of the website that you want to add to the instance.
	Domain pulumi.StringPtrInput
	// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
	// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
	HttpsExt    pulumi.StringPtrInput
	InstanceIds pulumi.StringArrayInput
	// Protocol type and port number information.
	ProxyTypes DomainResourceProxyTypeArrayInput
	// the IP address. This field is required and must be a string array.
	RealServers pulumi.StringArrayInput
	// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	RsType pulumi.IntPtrInput
}

func (DomainResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainResourceState)(nil)).Elem()
}

type domainResourceArgs struct {
	// The domain name of the website that you want to add to the instance.
	Domain string `pulumi:"domain"`
	// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
	// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
	HttpsExt    *string  `pulumi:"httpsExt"`
	InstanceIds []string `pulumi:"instanceIds"`
	// Protocol type and port number information.
	ProxyTypes []DomainResourceProxyType `pulumi:"proxyTypes"`
	// the IP address. This field is required and must be a string array.
	RealServers []string `pulumi:"realServers"`
	// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	RsType int `pulumi:"rsType"`
}

// The set of arguments for constructing a DomainResource resource.
type DomainResourceArgs struct {
	// The domain name of the website that you want to add to the instance.
	Domain pulumi.StringInput
	// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
	// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
	// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
	// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
	HttpsExt    pulumi.StringPtrInput
	InstanceIds pulumi.StringArrayInput
	// Protocol type and port number information.
	ProxyTypes DomainResourceProxyTypeArrayInput
	// the IP address. This field is required and must be a string array.
	RealServers pulumi.StringArrayInput
	// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
	RsType pulumi.IntInput
}

func (DomainResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainResourceArgs)(nil)).Elem()
}

type DomainResourceInput interface {
	pulumi.Input

	ToDomainResourceOutput() DomainResourceOutput
	ToDomainResourceOutputWithContext(ctx context.Context) DomainResourceOutput
}

func (*DomainResource) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainResource)(nil)).Elem()
}

func (i *DomainResource) ToDomainResourceOutput() DomainResourceOutput {
	return i.ToDomainResourceOutputWithContext(context.Background())
}

func (i *DomainResource) ToDomainResourceOutputWithContext(ctx context.Context) DomainResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainResourceOutput)
}

// DomainResourceArrayInput is an input type that accepts DomainResourceArray and DomainResourceArrayOutput values.
// You can construct a concrete instance of `DomainResourceArrayInput` via:
//
//	DomainResourceArray{ DomainResourceArgs{...} }
type DomainResourceArrayInput interface {
	pulumi.Input

	ToDomainResourceArrayOutput() DomainResourceArrayOutput
	ToDomainResourceArrayOutputWithContext(context.Context) DomainResourceArrayOutput
}

type DomainResourceArray []DomainResourceInput

func (DomainResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainResource)(nil)).Elem()
}

func (i DomainResourceArray) ToDomainResourceArrayOutput() DomainResourceArrayOutput {
	return i.ToDomainResourceArrayOutputWithContext(context.Background())
}

func (i DomainResourceArray) ToDomainResourceArrayOutputWithContext(ctx context.Context) DomainResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainResourceArrayOutput)
}

// DomainResourceMapInput is an input type that accepts DomainResourceMap and DomainResourceMapOutput values.
// You can construct a concrete instance of `DomainResourceMapInput` via:
//
//	DomainResourceMap{ "key": DomainResourceArgs{...} }
type DomainResourceMapInput interface {
	pulumi.Input

	ToDomainResourceMapOutput() DomainResourceMapOutput
	ToDomainResourceMapOutputWithContext(context.Context) DomainResourceMapOutput
}

type DomainResourceMap map[string]DomainResourceInput

func (DomainResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainResource)(nil)).Elem()
}

func (i DomainResourceMap) ToDomainResourceMapOutput() DomainResourceMapOutput {
	return i.ToDomainResourceMapOutputWithContext(context.Background())
}

func (i DomainResourceMap) ToDomainResourceMapOutputWithContext(ctx context.Context) DomainResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainResourceMapOutput)
}

type DomainResourceOutput struct{ *pulumi.OutputState }

func (DomainResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainResource)(nil)).Elem()
}

func (o DomainResourceOutput) ToDomainResourceOutput() DomainResourceOutput {
	return o
}

func (o DomainResourceOutput) ToDomainResourceOutputWithContext(ctx context.Context) DomainResourceOutput {
	return o
}

// The domain name of the website that you want to add to the instance.
func (o DomainResourceOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainResource) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
// - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
// - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
// - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
func (o DomainResourceOutput) HttpsExt() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainResource) pulumi.StringOutput { return v.HttpsExt }).(pulumi.StringOutput)
}

func (o DomainResourceOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainResource) pulumi.StringArrayOutput { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// Protocol type and port number information.
func (o DomainResourceOutput) ProxyTypes() DomainResourceProxyTypeArrayOutput {
	return o.ApplyT(func(v *DomainResource) DomainResourceProxyTypeArrayOutput { return v.ProxyTypes }).(DomainResourceProxyTypeArrayOutput)
}

// the IP address. This field is required and must be a string array.
func (o DomainResourceOutput) RealServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainResource) pulumi.StringArrayOutput { return v.RealServers }).(pulumi.StringArrayOutput)
}

// The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
func (o DomainResourceOutput) RsType() pulumi.IntOutput {
	return o.ApplyT(func(v *DomainResource) pulumi.IntOutput { return v.RsType }).(pulumi.IntOutput)
}

type DomainResourceArrayOutput struct{ *pulumi.OutputState }

func (DomainResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainResource)(nil)).Elem()
}

func (o DomainResourceArrayOutput) ToDomainResourceArrayOutput() DomainResourceArrayOutput {
	return o
}

func (o DomainResourceArrayOutput) ToDomainResourceArrayOutputWithContext(ctx context.Context) DomainResourceArrayOutput {
	return o
}

func (o DomainResourceArrayOutput) Index(i pulumi.IntInput) DomainResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainResource {
		return vs[0].([]*DomainResource)[vs[1].(int)]
	}).(DomainResourceOutput)
}

type DomainResourceMapOutput struct{ *pulumi.OutputState }

func (DomainResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainResource)(nil)).Elem()
}

func (o DomainResourceMapOutput) ToDomainResourceMapOutput() DomainResourceMapOutput {
	return o
}

func (o DomainResourceMapOutput) ToDomainResourceMapOutputWithContext(ctx context.Context) DomainResourceMapOutput {
	return o
}

func (o DomainResourceMapOutput) MapIndex(k pulumi.StringInput) DomainResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainResource {
		return vs[0].(map[string]*DomainResource)[vs[1].(string)]
	}).(DomainResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainResourceInput)(nil)).Elem(), &DomainResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainResourceArrayInput)(nil)).Elem(), DomainResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainResourceMapInput)(nil)).Elem(), DomainResourceMap{})
	pulumi.RegisterOutputType(DomainResourceOutput{})
	pulumi.RegisterOutputType(DomainResourceArrayOutput{})
	pulumi.RegisterOutputType(DomainResourceMapOutput{})
}
