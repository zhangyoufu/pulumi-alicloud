// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Additional Certificate resource.
//
// For information about Global Accelerator (GA) Additional Certificate and how to use it, see [What is Additional Certificate](https://www.alibabacloud.com/help/en/doc-detail/302356.html).
//
// > **NOTE:** Available in v1.150.0+.
//
// ## Import
//
// Global Accelerator (GA) Additional Certificate can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/additionalCertificate:AdditionalCertificate example <accelerator_id>:<listener_id>:<domain>
//
// ```
type AdditionalCertificate struct {
	pulumi.CustomResourceState

	// The ID of the GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The Certificate ID.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
}

// NewAdditionalCertificate registers a new resource with the given unique name, arguments, and options.
func NewAdditionalCertificate(ctx *pulumi.Context,
	name string, args *AdditionalCertificateArgs, opts ...pulumi.ResourceOption) (*AdditionalCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.CertificateId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateId'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	var resource AdditionalCertificate
	err := ctx.RegisterResource("alicloud:ga/additionalCertificate:AdditionalCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdditionalCertificate gets an existing AdditionalCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdditionalCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdditionalCertificateState, opts ...pulumi.ResourceOption) (*AdditionalCertificate, error) {
	var resource AdditionalCertificate
	err := ctx.ReadResource("alicloud:ga/additionalCertificate:AdditionalCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdditionalCertificate resources.
type additionalCertificateState struct {
	// The ID of the GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The Certificate ID.
	CertificateId *string `pulumi:"certificateId"`
	// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
	Domain *string `pulumi:"domain"`
	// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
	ListenerId *string `pulumi:"listenerId"`
}

type AdditionalCertificateState struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The Certificate ID.
	CertificateId pulumi.StringPtrInput
	// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
	Domain pulumi.StringPtrInput
	// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
	ListenerId pulumi.StringPtrInput
}

func (AdditionalCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*additionalCertificateState)(nil)).Elem()
}

type additionalCertificateArgs struct {
	// The ID of the GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The Certificate ID.
	CertificateId string `pulumi:"certificateId"`
	// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
	Domain string `pulumi:"domain"`
	// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
	ListenerId string `pulumi:"listenerId"`
}

// The set of arguments for constructing a AdditionalCertificate resource.
type AdditionalCertificateArgs struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringInput
	// The Certificate ID.
	CertificateId pulumi.StringInput
	// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
	Domain pulumi.StringInput
	// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
	ListenerId pulumi.StringInput
}

func (AdditionalCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*additionalCertificateArgs)(nil)).Elem()
}

type AdditionalCertificateInput interface {
	pulumi.Input

	ToAdditionalCertificateOutput() AdditionalCertificateOutput
	ToAdditionalCertificateOutputWithContext(ctx context.Context) AdditionalCertificateOutput
}

func (*AdditionalCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalCertificate)(nil)).Elem()
}

func (i *AdditionalCertificate) ToAdditionalCertificateOutput() AdditionalCertificateOutput {
	return i.ToAdditionalCertificateOutputWithContext(context.Background())
}

func (i *AdditionalCertificate) ToAdditionalCertificateOutputWithContext(ctx context.Context) AdditionalCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCertificateOutput)
}

// AdditionalCertificateArrayInput is an input type that accepts AdditionalCertificateArray and AdditionalCertificateArrayOutput values.
// You can construct a concrete instance of `AdditionalCertificateArrayInput` via:
//
//	AdditionalCertificateArray{ AdditionalCertificateArgs{...} }
type AdditionalCertificateArrayInput interface {
	pulumi.Input

	ToAdditionalCertificateArrayOutput() AdditionalCertificateArrayOutput
	ToAdditionalCertificateArrayOutputWithContext(context.Context) AdditionalCertificateArrayOutput
}

type AdditionalCertificateArray []AdditionalCertificateInput

func (AdditionalCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdditionalCertificate)(nil)).Elem()
}

func (i AdditionalCertificateArray) ToAdditionalCertificateArrayOutput() AdditionalCertificateArrayOutput {
	return i.ToAdditionalCertificateArrayOutputWithContext(context.Background())
}

func (i AdditionalCertificateArray) ToAdditionalCertificateArrayOutputWithContext(ctx context.Context) AdditionalCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCertificateArrayOutput)
}

// AdditionalCertificateMapInput is an input type that accepts AdditionalCertificateMap and AdditionalCertificateMapOutput values.
// You can construct a concrete instance of `AdditionalCertificateMapInput` via:
//
//	AdditionalCertificateMap{ "key": AdditionalCertificateArgs{...} }
type AdditionalCertificateMapInput interface {
	pulumi.Input

	ToAdditionalCertificateMapOutput() AdditionalCertificateMapOutput
	ToAdditionalCertificateMapOutputWithContext(context.Context) AdditionalCertificateMapOutput
}

type AdditionalCertificateMap map[string]AdditionalCertificateInput

func (AdditionalCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdditionalCertificate)(nil)).Elem()
}

func (i AdditionalCertificateMap) ToAdditionalCertificateMapOutput() AdditionalCertificateMapOutput {
	return i.ToAdditionalCertificateMapOutputWithContext(context.Background())
}

func (i AdditionalCertificateMap) ToAdditionalCertificateMapOutputWithContext(ctx context.Context) AdditionalCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCertificateMapOutput)
}

type AdditionalCertificateOutput struct{ *pulumi.OutputState }

func (AdditionalCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalCertificate)(nil)).Elem()
}

func (o AdditionalCertificateOutput) ToAdditionalCertificateOutput() AdditionalCertificateOutput {
	return o
}

func (o AdditionalCertificateOutput) ToAdditionalCertificateOutputWithContext(ctx context.Context) AdditionalCertificateOutput {
	return o
}

// The ID of the GA instance.
func (o AdditionalCertificateOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdditionalCertificate) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The Certificate ID.
func (o AdditionalCertificateOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdditionalCertificate) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
func (o AdditionalCertificateOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *AdditionalCertificate) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
func (o AdditionalCertificateOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdditionalCertificate) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

type AdditionalCertificateArrayOutput struct{ *pulumi.OutputState }

func (AdditionalCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdditionalCertificate)(nil)).Elem()
}

func (o AdditionalCertificateArrayOutput) ToAdditionalCertificateArrayOutput() AdditionalCertificateArrayOutput {
	return o
}

func (o AdditionalCertificateArrayOutput) ToAdditionalCertificateArrayOutputWithContext(ctx context.Context) AdditionalCertificateArrayOutput {
	return o
}

func (o AdditionalCertificateArrayOutput) Index(i pulumi.IntInput) AdditionalCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdditionalCertificate {
		return vs[0].([]*AdditionalCertificate)[vs[1].(int)]
	}).(AdditionalCertificateOutput)
}

type AdditionalCertificateMapOutput struct{ *pulumi.OutputState }

func (AdditionalCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdditionalCertificate)(nil)).Elem()
}

func (o AdditionalCertificateMapOutput) ToAdditionalCertificateMapOutput() AdditionalCertificateMapOutput {
	return o
}

func (o AdditionalCertificateMapOutput) ToAdditionalCertificateMapOutputWithContext(ctx context.Context) AdditionalCertificateMapOutput {
	return o
}

func (o AdditionalCertificateMapOutput) MapIndex(k pulumi.StringInput) AdditionalCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdditionalCertificate {
		return vs[0].(map[string]*AdditionalCertificate)[vs[1].(string)]
	}).(AdditionalCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdditionalCertificateInput)(nil)).Elem(), &AdditionalCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdditionalCertificateArrayInput)(nil)).Elem(), AdditionalCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdditionalCertificateMapInput)(nil)).Elem(), AdditionalCertificateMap{})
	pulumi.RegisterOutputType(AdditionalCertificateOutput{})
	pulumi.RegisterOutputType(AdditionalCertificateArrayOutput{})
	pulumi.RegisterOutputType(AdditionalCertificateMapOutput{})
}
