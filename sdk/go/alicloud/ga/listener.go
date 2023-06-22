// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Listener resource.
//
// For information about Global Accelerator (GA) Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-doc-ga-2019-11-20-api-doc-createlistener).
//
// > **NOTE:** Available since v1.111.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAccelerator, err := ga.NewAccelerator(ctx, "defaultAccelerator", &ga.AcceleratorArgs{
//				Duration:      pulumi.Int(1),
//				AutoUseCoupon: pulumi.Bool(true),
//				Spec:          pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "defaultBandwidthPackage", &ga.BandwidthPackageArgs{
//				Bandwidth:     pulumi.Int(100),
//				Type:          pulumi.String("Basic"),
//				BandwidthType: pulumi.String("Basic"),
//				PaymentType:   pulumi.String("PayAsYouGo"),
//				BillingType:   pulumi.String("PayBy95"),
//				Ratio:         pulumi.Int(30),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "defaultBandwidthPackageAttachment", &ga.BandwidthPackageAttachmentArgs{
//				AcceleratorId:      defaultAccelerator.ID(),
//				BandwidthPackageId: defaultBandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewListener(ctx, "defaultListener", &ga.ListenerArgs{
//				AcceleratorId: defaultBandwidthPackageAttachment.AcceleratorId,
//				PortRanges: ga.ListenerPortRangeArray{
//					&ga.ListenerPortRangeArgs{
//						FromPort: pulumi.Int(80),
//						ToPort:   pulumi.Int(80),
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
// Ga Listener can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/listener:Listener example <id>
//
// ```
type Listener struct {
	pulumi.CustomResourceState

	// The accelerator id.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The certificates of the listener. See `certificates` below.
	// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
	Certificates ListenerCertificateArrayOutput `pulumi:"certificates"`
	// The clientAffinity of the listener. Default value: `NONE`. Valid values:
	ClientAffinity pulumi.StringPtrOutput `pulumi:"clientAffinity"`
	// The description of the listener.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The XForward headers. See `forwardedForConfig` below.
	ForwardedForConfig ListenerForwardedForConfigPtrOutput `pulumi:"forwardedForConfig"`
	// The routing type of the listener. Default Value: `Standard`. Valid values:
	ListenerType pulumi.StringOutput `pulumi:"listenerType"`
	// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
	Name pulumi.StringOutput `pulumi:"name"`
	// The portRanges of the listener. See `portRanges` below.
	// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
	PortRanges ListenerPortRangeArrayOutput `pulumi:"portRanges"`
	// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The proxy protocol of the listener. Default value: `false`. Valid values:
	ProxyProtocol pulumi.BoolPtrOutput `pulumi:"proxyProtocol"`
	// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
	// The status of the listener.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.PortRanges == nil {
		return nil, errors.New("invalid value for required argument 'PortRanges'")
	}
	var resource Listener
	err := ctx.RegisterResource("alicloud:ga/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("alicloud:ga/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The accelerator id.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The certificates of the listener. See `certificates` below.
	// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
	Certificates []ListenerCertificate `pulumi:"certificates"`
	// The clientAffinity of the listener. Default value: `NONE`. Valid values:
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The description of the listener.
	Description *string `pulumi:"description"`
	// The XForward headers. See `forwardedForConfig` below.
	ForwardedForConfig *ListenerForwardedForConfig `pulumi:"forwardedForConfig"`
	// The routing type of the listener. Default Value: `Standard`. Valid values:
	ListenerType *string `pulumi:"listenerType"`
	// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
	Name *string `pulumi:"name"`
	// The portRanges of the listener. See `portRanges` below.
	// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol *string `pulumi:"protocol"`
	// The proxy protocol of the listener. Default value: `false`. Valid values:
	ProxyProtocol *bool `pulumi:"proxyProtocol"`
	// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
	// The status of the listener.
	Status *string `pulumi:"status"`
}

type ListenerState struct {
	// The accelerator id.
	AcceleratorId pulumi.StringPtrInput
	// The certificates of the listener. See `certificates` below.
	// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
	Certificates ListenerCertificateArrayInput
	// The clientAffinity of the listener. Default value: `NONE`. Valid values:
	ClientAffinity pulumi.StringPtrInput
	// The description of the listener.
	Description pulumi.StringPtrInput
	// The XForward headers. See `forwardedForConfig` below.
	ForwardedForConfig ListenerForwardedForConfigPtrInput
	// The routing type of the listener. Default Value: `Standard`. Valid values:
	ListenerType pulumi.StringPtrInput
	// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
	Name pulumi.StringPtrInput
	// The portRanges of the listener. See `portRanges` below.
	// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
	PortRanges ListenerPortRangeArrayInput
	// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringPtrInput
	// The proxy protocol of the listener. Default value: `false`. Valid values:
	ProxyProtocol pulumi.BoolPtrInput
	// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
	SecurityPolicyId pulumi.StringPtrInput
	// The status of the listener.
	Status pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The accelerator id.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The certificates of the listener. See `certificates` below.
	// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
	Certificates []ListenerCertificate `pulumi:"certificates"`
	// The clientAffinity of the listener. Default value: `NONE`. Valid values:
	ClientAffinity *string `pulumi:"clientAffinity"`
	// The description of the listener.
	Description *string `pulumi:"description"`
	// The XForward headers. See `forwardedForConfig` below.
	ForwardedForConfig *ListenerForwardedForConfig `pulumi:"forwardedForConfig"`
	// The routing type of the listener. Default Value: `Standard`. Valid values:
	ListenerType *string `pulumi:"listenerType"`
	// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
	Name *string `pulumi:"name"`
	// The portRanges of the listener. See `portRanges` below.
	// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
	PortRanges []ListenerPortRange `pulumi:"portRanges"`
	// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol *string `pulumi:"protocol"`
	// The proxy protocol of the listener. Default value: `false`. Valid values:
	ProxyProtocol *bool `pulumi:"proxyProtocol"`
	// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The accelerator id.
	AcceleratorId pulumi.StringInput
	// The certificates of the listener. See `certificates` below.
	// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
	Certificates ListenerCertificateArrayInput
	// The clientAffinity of the listener. Default value: `NONE`. Valid values:
	ClientAffinity pulumi.StringPtrInput
	// The description of the listener.
	Description pulumi.StringPtrInput
	// The XForward headers. See `forwardedForConfig` below.
	ForwardedForConfig ListenerForwardedForConfigPtrInput
	// The routing type of the listener. Default Value: `Standard`. Valid values:
	ListenerType pulumi.StringPtrInput
	// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
	Name pulumi.StringPtrInput
	// The portRanges of the listener. See `portRanges` below.
	// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
	PortRanges ListenerPortRangeArrayInput
	// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringPtrInput
	// The proxy protocol of the listener. Default value: `false`. Valid values:
	ProxyProtocol pulumi.BoolPtrInput
	// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
	SecurityPolicyId pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//	ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//	ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

// The accelerator id.
func (o ListenerOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The certificates of the listener. See `certificates` below.
// > **NOTE:** This parameter needs to be configured only for monitoring of the `HTTPS` protocol.
func (o ListenerOutput) Certificates() ListenerCertificateArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerCertificateArrayOutput { return v.Certificates }).(ListenerCertificateArrayOutput)
}

// The clientAffinity of the listener. Default value: `NONE`. Valid values:
func (o ListenerOutput) ClientAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.ClientAffinity }).(pulumi.StringPtrOutput)
}

// The description of the listener.
func (o ListenerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The XForward headers. See `forwardedForConfig` below.
func (o ListenerOutput) ForwardedForConfig() ListenerForwardedForConfigPtrOutput {
	return o.ApplyT(func(v *Listener) ListenerForwardedForConfigPtrOutput { return v.ForwardedForConfig }).(ListenerForwardedForConfigPtrOutput)
}

// The routing type of the listener. Default Value: `Standard`. Valid values:
func (o ListenerOutput) ListenerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ListenerType }).(pulumi.StringOutput)
}

// The name of the listener. The length of the name is 2-128 characters. It starts with uppercase and lowercase letters or Chinese characters. It can contain numbers and underscores and dashes.
func (o ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The portRanges of the listener. See `portRanges` below.
// > **NOTE:** For `HTTP` or `HTTPS` protocol monitoring, only one monitoring port can be configured, that is, the start monitoring port and end monitoring port should be the same.
func (o ListenerOutput) PortRanges() ListenerPortRangeArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerPortRangeArrayOutput { return v.PortRanges }).(ListenerPortRangeArrayOutput)
}

// Type of network transport protocol monitored. Default value: `TCP`. Valid values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
func (o ListenerOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The proxy protocol of the listener. Default value: `false`. Valid values:
func (o ListenerOutput) ProxyProtocol() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.ProxyProtocol }).(pulumi.BoolPtrOutput)
}

// The ID of the security policy. **NOTE:** Only `HTTPS` listeners support this parameter. Valid values:
func (o ListenerOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

// The status of the listener.
func (o ListenerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].([]*Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].(map[string]*Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerArrayInput)(nil)).Elem(), ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerMapInput)(nil)).Elem(), ListenerMap{})
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
