// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SLB Tls Cipher Policy resource.
//
// For information about SLB Tls Cipher Policy and how to use it, see [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
//
// > **NOTE:** Available in v1.135.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := slb.NewTlsCipherPolicy(ctx, "example", &slb.TlsCipherPolicyArgs{
//				Ciphers: pulumi.StringArray{
//					pulumi.String("AES256-SHA256"),
//					pulumi.String("AES128-GCM-SHA256"),
//				},
//				TlsCipherPolicyName: pulumi.String("Test-example_value"),
//				TlsVersions: pulumi.StringArray{
//					pulumi.String("TLSv1.2"),
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
// SLB Tls Cipher Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:slb/tlsCipherPolicy:TlsCipherPolicy example <id>
// ```
type TlsCipherPolicy struct {
	pulumi.CustomResourceState

	// The encryption algorithms supported. It depends on the value of `tlsVersions`.
	Ciphers pulumi.StringArrayOutput `pulumi:"ciphers"`
	// TLS policy instance state.
	Status pulumi.StringOutput `pulumi:"status"`
	// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
	TlsCipherPolicyName pulumi.StringOutput `pulumi:"tlsCipherPolicyName"`
	// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
	TlsVersions pulumi.StringArrayOutput `pulumi:"tlsVersions"`
}

// NewTlsCipherPolicy registers a new resource with the given unique name, arguments, and options.
func NewTlsCipherPolicy(ctx *pulumi.Context,
	name string, args *TlsCipherPolicyArgs, opts ...pulumi.ResourceOption) (*TlsCipherPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ciphers == nil {
		return nil, errors.New("invalid value for required argument 'Ciphers'")
	}
	if args.TlsCipherPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'TlsCipherPolicyName'")
	}
	if args.TlsVersions == nil {
		return nil, errors.New("invalid value for required argument 'TlsVersions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TlsCipherPolicy
	err := ctx.RegisterResource("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTlsCipherPolicy gets an existing TlsCipherPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTlsCipherPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TlsCipherPolicyState, opts ...pulumi.ResourceOption) (*TlsCipherPolicy, error) {
	var resource TlsCipherPolicy
	err := ctx.ReadResource("alicloud:slb/tlsCipherPolicy:TlsCipherPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TlsCipherPolicy resources.
type tlsCipherPolicyState struct {
	// The encryption algorithms supported. It depends on the value of `tlsVersions`.
	Ciphers []string `pulumi:"ciphers"`
	// TLS policy instance state.
	Status *string `pulumi:"status"`
	// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
	TlsCipherPolicyName *string `pulumi:"tlsCipherPolicyName"`
	// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
	TlsVersions []string `pulumi:"tlsVersions"`
}

type TlsCipherPolicyState struct {
	// The encryption algorithms supported. It depends on the value of `tlsVersions`.
	Ciphers pulumi.StringArrayInput
	// TLS policy instance state.
	Status pulumi.StringPtrInput
	// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
	TlsCipherPolicyName pulumi.StringPtrInput
	// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
	TlsVersions pulumi.StringArrayInput
}

func (TlsCipherPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsCipherPolicyState)(nil)).Elem()
}

type tlsCipherPolicyArgs struct {
	// The encryption algorithms supported. It depends on the value of `tlsVersions`.
	Ciphers []string `pulumi:"ciphers"`
	// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
	TlsCipherPolicyName string `pulumi:"tlsCipherPolicyName"`
	// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
	TlsVersions []string `pulumi:"tlsVersions"`
}

// The set of arguments for constructing a TlsCipherPolicy resource.
type TlsCipherPolicyArgs struct {
	// The encryption algorithms supported. It depends on the value of `tlsVersions`.
	Ciphers pulumi.StringArrayInput
	// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
	TlsCipherPolicyName pulumi.StringInput
	// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
	TlsVersions pulumi.StringArrayInput
}

func (TlsCipherPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsCipherPolicyArgs)(nil)).Elem()
}

type TlsCipherPolicyInput interface {
	pulumi.Input

	ToTlsCipherPolicyOutput() TlsCipherPolicyOutput
	ToTlsCipherPolicyOutputWithContext(ctx context.Context) TlsCipherPolicyOutput
}

func (*TlsCipherPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsCipherPolicy)(nil)).Elem()
}

func (i *TlsCipherPolicy) ToTlsCipherPolicyOutput() TlsCipherPolicyOutput {
	return i.ToTlsCipherPolicyOutputWithContext(context.Background())
}

func (i *TlsCipherPolicy) ToTlsCipherPolicyOutputWithContext(ctx context.Context) TlsCipherPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCipherPolicyOutput)
}

// TlsCipherPolicyArrayInput is an input type that accepts TlsCipherPolicyArray and TlsCipherPolicyArrayOutput values.
// You can construct a concrete instance of `TlsCipherPolicyArrayInput` via:
//
//	TlsCipherPolicyArray{ TlsCipherPolicyArgs{...} }
type TlsCipherPolicyArrayInput interface {
	pulumi.Input

	ToTlsCipherPolicyArrayOutput() TlsCipherPolicyArrayOutput
	ToTlsCipherPolicyArrayOutputWithContext(context.Context) TlsCipherPolicyArrayOutput
}

type TlsCipherPolicyArray []TlsCipherPolicyInput

func (TlsCipherPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsCipherPolicy)(nil)).Elem()
}

func (i TlsCipherPolicyArray) ToTlsCipherPolicyArrayOutput() TlsCipherPolicyArrayOutput {
	return i.ToTlsCipherPolicyArrayOutputWithContext(context.Background())
}

func (i TlsCipherPolicyArray) ToTlsCipherPolicyArrayOutputWithContext(ctx context.Context) TlsCipherPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCipherPolicyArrayOutput)
}

// TlsCipherPolicyMapInput is an input type that accepts TlsCipherPolicyMap and TlsCipherPolicyMapOutput values.
// You can construct a concrete instance of `TlsCipherPolicyMapInput` via:
//
//	TlsCipherPolicyMap{ "key": TlsCipherPolicyArgs{...} }
type TlsCipherPolicyMapInput interface {
	pulumi.Input

	ToTlsCipherPolicyMapOutput() TlsCipherPolicyMapOutput
	ToTlsCipherPolicyMapOutputWithContext(context.Context) TlsCipherPolicyMapOutput
}

type TlsCipherPolicyMap map[string]TlsCipherPolicyInput

func (TlsCipherPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsCipherPolicy)(nil)).Elem()
}

func (i TlsCipherPolicyMap) ToTlsCipherPolicyMapOutput() TlsCipherPolicyMapOutput {
	return i.ToTlsCipherPolicyMapOutputWithContext(context.Background())
}

func (i TlsCipherPolicyMap) ToTlsCipherPolicyMapOutputWithContext(ctx context.Context) TlsCipherPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCipherPolicyMapOutput)
}

type TlsCipherPolicyOutput struct{ *pulumi.OutputState }

func (TlsCipherPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsCipherPolicy)(nil)).Elem()
}

func (o TlsCipherPolicyOutput) ToTlsCipherPolicyOutput() TlsCipherPolicyOutput {
	return o
}

func (o TlsCipherPolicyOutput) ToTlsCipherPolicyOutputWithContext(ctx context.Context) TlsCipherPolicyOutput {
	return o
}

// The encryption algorithms supported. It depends on the value of `tlsVersions`.
func (o TlsCipherPolicyOutput) Ciphers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TlsCipherPolicy) pulumi.StringArrayOutput { return v.Ciphers }).(pulumi.StringArrayOutput)
}

// TLS policy instance state.
func (o TlsCipherPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsCipherPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
func (o TlsCipherPolicyOutput) TlsCipherPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsCipherPolicy) pulumi.StringOutput { return v.TlsCipherPolicyName }).(pulumi.StringOutput)
}

// The version of TLS protocol. You can find the corresponding value description in the document center [What is Tls Cipher Policy](https://www.alibabacloud.com/help/doc-detail/196714.htm).
func (o TlsCipherPolicyOutput) TlsVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TlsCipherPolicy) pulumi.StringArrayOutput { return v.TlsVersions }).(pulumi.StringArrayOutput)
}

type TlsCipherPolicyArrayOutput struct{ *pulumi.OutputState }

func (TlsCipherPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsCipherPolicy)(nil)).Elem()
}

func (o TlsCipherPolicyArrayOutput) ToTlsCipherPolicyArrayOutput() TlsCipherPolicyArrayOutput {
	return o
}

func (o TlsCipherPolicyArrayOutput) ToTlsCipherPolicyArrayOutputWithContext(ctx context.Context) TlsCipherPolicyArrayOutput {
	return o
}

func (o TlsCipherPolicyArrayOutput) Index(i pulumi.IntInput) TlsCipherPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TlsCipherPolicy {
		return vs[0].([]*TlsCipherPolicy)[vs[1].(int)]
	}).(TlsCipherPolicyOutput)
}

type TlsCipherPolicyMapOutput struct{ *pulumi.OutputState }

func (TlsCipherPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsCipherPolicy)(nil)).Elem()
}

func (o TlsCipherPolicyMapOutput) ToTlsCipherPolicyMapOutput() TlsCipherPolicyMapOutput {
	return o
}

func (o TlsCipherPolicyMapOutput) ToTlsCipherPolicyMapOutputWithContext(ctx context.Context) TlsCipherPolicyMapOutput {
	return o
}

func (o TlsCipherPolicyMapOutput) MapIndex(k pulumi.StringInput) TlsCipherPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TlsCipherPolicy {
		return vs[0].(map[string]*TlsCipherPolicy)[vs[1].(string)]
	}).(TlsCipherPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TlsCipherPolicyInput)(nil)).Elem(), &TlsCipherPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsCipherPolicyArrayInput)(nil)).Elem(), TlsCipherPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsCipherPolicyMapInput)(nil)).Elem(), TlsCipherPolicyMap{})
	pulumi.RegisterOutputType(TlsCipherPolicyOutput{})
	pulumi.RegisterOutputType(TlsCipherPolicyArrayOutput{})
	pulumi.RegisterOutputType(TlsCipherPolicyMapOutput{})
}
