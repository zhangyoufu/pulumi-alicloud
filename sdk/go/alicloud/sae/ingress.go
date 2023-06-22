// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Serverless App Engine (SAE) Ingress resource.
//
// For information about Serverless App Engine (SAE) Ingress and how to use it, see [What is Ingress](https://www.alibabacloud.com/help/en/sae/latest/createingress).
//
// > **NOTE:** Available since v1.137.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sae"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "example_value"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/21"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultLoadBalancer, err := slb.NewLoadBalancer(ctx, "defaultLoadBalancer", &slb.LoadBalancerArgs{
//				Specification: pulumi.String("slb.s2.small"),
//				VswitchId:     pulumi.Any(data.Alicloud_vswitches.Default.Ids[0]),
//			})
//			if err != nil {
//				return err
//			}
//			namespaceId := "cn-hangzhou:yourname"
//			if param := cfg.Get("namespaceId"); param != "" {
//				namespaceId = param
//			}
//			defaultNamespace, err := sae.NewNamespace(ctx, "defaultNamespace", &sae.NamespaceArgs{
//				NamespaceId:          pulumi.String(namespaceId),
//				NamespaceName:        pulumi.String(name),
//				NamespaceDescription: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultApplication, err := sae.NewApplication(ctx, "defaultApplication", &sae.ApplicationArgs{
//				AppDescription: pulumi.String("your_app_description"),
//				AppName:        pulumi.String("your_app_name"),
//				NamespaceId:    pulumi.String("your_namespace_id"),
//				PackageUrl:     pulumi.String("your_package_url"),
//				PackageType:    pulumi.String("your_package_url"),
//				Jdk:            pulumi.String("jdk_specifications"),
//				VswitchId:      pulumi.Any(data.Alicloud_vswitches.Default.Ids[0]),
//				Replicas:       pulumi.Int("your_replicas"),
//				Cpu:            pulumi.Int("cpu_specifications"),
//				Memory:         pulumi.Int("memory_specifications"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sae.NewIngress(ctx, "defaultIngress", &sae.IngressArgs{
//				SlbId:        defaultLoadBalancer.ID(),
//				NamespaceId:  defaultNamespace.ID(),
//				ListenerPort: pulumi.Int("your_listener_port"),
//				Rules: sae.IngressRuleArray{
//					&sae.IngressRuleArgs{
//						AppId:         defaultApplication.ID(),
//						ContainerPort: pulumi.Int("your_container_port"),
//						Domain:        pulumi.String("your_domain"),
//						AppName:       pulumi.String("your_name"),
//						Path:          pulumi.String("your_path"),
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
// Serverless App Engine (SAE) Ingress can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sae/ingress:Ingress example <id>
//
// ```
type Ingress struct {
	pulumi.CustomResourceState

	// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
	CertId pulumi.StringPtrOutput `pulumi:"certId"`
	// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
	CertIds pulumi.StringPtrOutput `pulumi:"certIds"`
	// Default Rule. See `defaultRule` below.
	DefaultRule IngressDefaultRulePtrOutput `pulumi:"defaultRule"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// SLB listening port.
	ListenerPort pulumi.IntOutput `pulumi:"listenerPort"`
	// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
	ListenerProtocol pulumi.StringOutput `pulumi:"listenerProtocol"`
	// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
	LoadBalanceType pulumi.StringOutput `pulumi:"loadBalanceType"`
	// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
	Rules IngressRuleArrayOutput `pulumi:"rules"`
	// SLB ID.
	SlbId pulumi.StringOutput `pulumi:"slbId"`
}

// NewIngress registers a new resource with the given unique name, arguments, and options.
func NewIngress(ctx *pulumi.Context,
	name string, args *IngressArgs, opts ...pulumi.ResourceOption) (*Ingress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenerPort == nil {
		return nil, errors.New("invalid value for required argument 'ListenerPort'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.SlbId == nil {
		return nil, errors.New("invalid value for required argument 'SlbId'")
	}
	var resource Ingress
	err := ctx.RegisterResource("alicloud:sae/ingress:Ingress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngress gets an existing Ingress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngressState, opts ...pulumi.ResourceOption) (*Ingress, error) {
	var resource Ingress
	err := ctx.ReadResource("alicloud:sae/ingress:Ingress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ingress resources.
type ingressState struct {
	// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
	CertId *string `pulumi:"certId"`
	// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
	CertIds *string `pulumi:"certIds"`
	// Default Rule. See `defaultRule` below.
	DefaultRule *IngressDefaultRule `pulumi:"defaultRule"`
	// Description.
	Description *string `pulumi:"description"`
	// SLB listening port.
	ListenerPort *int `pulumi:"listenerPort"`
	// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
	ListenerProtocol *string `pulumi:"listenerProtocol"`
	// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
	LoadBalanceType *string `pulumi:"loadBalanceType"`
	// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
	NamespaceId *string `pulumi:"namespaceId"`
	// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
	Rules []IngressRule `pulumi:"rules"`
	// SLB ID.
	SlbId *string `pulumi:"slbId"`
}

type IngressState struct {
	// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
	CertId pulumi.StringPtrInput
	// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
	CertIds pulumi.StringPtrInput
	// Default Rule. See `defaultRule` below.
	DefaultRule IngressDefaultRulePtrInput
	// Description.
	Description pulumi.StringPtrInput
	// SLB listening port.
	ListenerPort pulumi.IntPtrInput
	// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
	ListenerProtocol pulumi.StringPtrInput
	// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
	LoadBalanceType pulumi.StringPtrInput
	// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
	NamespaceId pulumi.StringPtrInput
	// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
	Rules IngressRuleArrayInput
	// SLB ID.
	SlbId pulumi.StringPtrInput
}

func (IngressState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressState)(nil)).Elem()
}

type ingressArgs struct {
	// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
	CertId *string `pulumi:"certId"`
	// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
	CertIds *string `pulumi:"certIds"`
	// Default Rule. See `defaultRule` below.
	DefaultRule *IngressDefaultRule `pulumi:"defaultRule"`
	// Description.
	Description *string `pulumi:"description"`
	// SLB listening port.
	ListenerPort int `pulumi:"listenerPort"`
	// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
	ListenerProtocol *string `pulumi:"listenerProtocol"`
	// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
	LoadBalanceType *string `pulumi:"loadBalanceType"`
	// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
	NamespaceId string `pulumi:"namespaceId"`
	// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
	Rules []IngressRule `pulumi:"rules"`
	// SLB ID.
	SlbId string `pulumi:"slbId"`
}

// The set of arguments for constructing a Ingress resource.
type IngressArgs struct {
	// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
	CertId pulumi.StringPtrInput
	// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
	CertIds pulumi.StringPtrInput
	// Default Rule. See `defaultRule` below.
	DefaultRule IngressDefaultRulePtrInput
	// Description.
	Description pulumi.StringPtrInput
	// SLB listening port.
	ListenerPort pulumi.IntInput
	// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
	ListenerProtocol pulumi.StringPtrInput
	// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
	LoadBalanceType pulumi.StringPtrInput
	// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
	NamespaceId pulumi.StringInput
	// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
	Rules IngressRuleArrayInput
	// SLB ID.
	SlbId pulumi.StringInput
}

func (IngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressArgs)(nil)).Elem()
}

type IngressInput interface {
	pulumi.Input

	ToIngressOutput() IngressOutput
	ToIngressOutputWithContext(ctx context.Context) IngressOutput
}

func (*Ingress) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (i *Ingress) ToIngressOutput() IngressOutput {
	return i.ToIngressOutputWithContext(context.Background())
}

func (i *Ingress) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressOutput)
}

// IngressArrayInput is an input type that accepts IngressArray and IngressArrayOutput values.
// You can construct a concrete instance of `IngressArrayInput` via:
//
//	IngressArray{ IngressArgs{...} }
type IngressArrayInput interface {
	pulumi.Input

	ToIngressArrayOutput() IngressArrayOutput
	ToIngressArrayOutputWithContext(context.Context) IngressArrayOutput
}

type IngressArray []IngressInput

func (IngressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingress)(nil)).Elem()
}

func (i IngressArray) ToIngressArrayOutput() IngressArrayOutput {
	return i.ToIngressArrayOutputWithContext(context.Background())
}

func (i IngressArray) ToIngressArrayOutputWithContext(ctx context.Context) IngressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressArrayOutput)
}

// IngressMapInput is an input type that accepts IngressMap and IngressMapOutput values.
// You can construct a concrete instance of `IngressMapInput` via:
//
//	IngressMap{ "key": IngressArgs{...} }
type IngressMapInput interface {
	pulumi.Input

	ToIngressMapOutput() IngressMapOutput
	ToIngressMapOutputWithContext(context.Context) IngressMapOutput
}

type IngressMap map[string]IngressInput

func (IngressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingress)(nil)).Elem()
}

func (i IngressMap) ToIngressMapOutput() IngressMapOutput {
	return i.ToIngressMapOutputWithContext(context.Background())
}

func (i IngressMap) ToIngressMapOutputWithContext(ctx context.Context) IngressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressMapOutput)
}

type IngressOutput struct{ *pulumi.OutputState }

func (IngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (o IngressOutput) ToIngressOutput() IngressOutput {
	return o
}

func (o IngressOutput) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return o
}

// The certificate ID of the HTTPS listener. The `certId` takes effect only when `loadBalanceType` is set to `clb`.
func (o IngressOutput) CertId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringPtrOutput { return v.CertId }).(pulumi.StringPtrOutput)
}

// The certificate IDs of the HTTPS listener, and multiple certificate IDs are separated by commas. The `certIds` takes effect only when `loadBalanceType` is set to `alb`.
func (o IngressOutput) CertIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringPtrOutput { return v.CertIds }).(pulumi.StringPtrOutput)
}

// Default Rule. See `defaultRule` below.
func (o IngressOutput) DefaultRule() IngressDefaultRulePtrOutput {
	return o.ApplyT(func(v *Ingress) IngressDefaultRulePtrOutput { return v.DefaultRule }).(IngressDefaultRulePtrOutput)
}

// Description.
func (o IngressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// SLB listening port.
func (o IngressOutput) ListenerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Ingress) pulumi.IntOutput { return v.ListenerPort }).(pulumi.IntOutput)
}

// The protocol that is used to forward requests. Default value: `HTTP`. Valid values: `HTTP`, `HTTPS`.
func (o IngressOutput) ListenerProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringOutput { return v.ListenerProtocol }).(pulumi.StringOutput)
}

// The type of the SLB instance. Default value: `clb`. Valid values: `clb`, `alb`.
func (o IngressOutput) LoadBalanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringOutput { return v.LoadBalanceType }).(pulumi.StringOutput)
}

// The ID of Namespace. It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`.
func (o IngressOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Forwarding rules. Forward traffic to the specified application according to the domain name and path. See `rules` below.
func (o IngressOutput) Rules() IngressRuleArrayOutput {
	return o.ApplyT(func(v *Ingress) IngressRuleArrayOutput { return v.Rules }).(IngressRuleArrayOutput)
}

// SLB ID.
func (o IngressOutput) SlbId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ingress) pulumi.StringOutput { return v.SlbId }).(pulumi.StringOutput)
}

type IngressArrayOutput struct{ *pulumi.OutputState }

func (IngressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ingress)(nil)).Elem()
}

func (o IngressArrayOutput) ToIngressArrayOutput() IngressArrayOutput {
	return o
}

func (o IngressArrayOutput) ToIngressArrayOutputWithContext(ctx context.Context) IngressArrayOutput {
	return o
}

func (o IngressArrayOutput) Index(i pulumi.IntInput) IngressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ingress {
		return vs[0].([]*Ingress)[vs[1].(int)]
	}).(IngressOutput)
}

type IngressMapOutput struct{ *pulumi.OutputState }

func (IngressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ingress)(nil)).Elem()
}

func (o IngressMapOutput) ToIngressMapOutput() IngressMapOutput {
	return o
}

func (o IngressMapOutput) ToIngressMapOutputWithContext(ctx context.Context) IngressMapOutput {
	return o
}

func (o IngressMapOutput) MapIndex(k pulumi.StringInput) IngressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ingress {
		return vs[0].(map[string]*Ingress)[vs[1].(string)]
	}).(IngressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngressInput)(nil)).Elem(), &Ingress{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressArrayInput)(nil)).Elem(), IngressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressMapInput)(nil)).Elem(), IngressMap{})
	pulumi.RegisterOutputType(IngressOutput{})
	pulumi.RegisterOutputType(IngressArrayOutput{})
	pulumi.RegisterOutputType(IngressMapOutput{})
}
