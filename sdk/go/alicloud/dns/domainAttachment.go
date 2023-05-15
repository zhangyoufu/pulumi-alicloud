// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides bind the domain name to the DNS instance resource.
//
// > **NOTE:** Available in v1.80.0+.
//
// > **DEPRECATED:**  This resource has been deprecated from version `1.99.0`. Please use new resource alicloud_alidns_domain_attachment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewDomainAttachment(ctx, "dns", &dns.DomainAttachmentArgs{
//				DomainNames: pulumi.StringArray{
//					pulumi.String("test111.abc"),
//					pulumi.String("test222.abc"),
//				},
//				InstanceId: pulumi.String("dns-cn-mp91lyq9xxxx"),
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
// DNS domain attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dns/domainAttachment:DomainAttachment example dns-cn-v0h1ldjhxxx
//
// ```
type DomainAttachment struct {
	pulumi.CustomResourceState

	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayOutput `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDomainAttachment registers a new resource with the given unique name, arguments, and options.
func NewDomainAttachment(ctx *pulumi.Context,
	name string, args *DomainAttachmentArgs, opts ...pulumi.ResourceOption) (*DomainAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainNames == nil {
		return nil, errors.New("invalid value for required argument 'DomainNames'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource DomainAttachment
	err := ctx.RegisterResource("alicloud:dns/domainAttachment:DomainAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainAttachment gets an existing DomainAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainAttachmentState, opts ...pulumi.ResourceOption) (*DomainAttachment, error) {
	var resource DomainAttachment
	err := ctx.ReadResource("alicloud:dns/domainAttachment:DomainAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainAttachment resources.
type domainAttachmentState struct {
	// The domain names bound to the DNS instance.
	DomainNames []string `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId *string `pulumi:"instanceId"`
}

type DomainAttachmentState struct {
	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayInput
	// The id of the DNS instance.
	InstanceId pulumi.StringPtrInput
}

func (DomainAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAttachmentState)(nil)).Elem()
}

type domainAttachmentArgs struct {
	// The domain names bound to the DNS instance.
	DomainNames []string `pulumi:"domainNames"`
	// The id of the DNS instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a DomainAttachment resource.
type DomainAttachmentArgs struct {
	// The domain names bound to the DNS instance.
	DomainNames pulumi.StringArrayInput
	// The id of the DNS instance.
	InstanceId pulumi.StringInput
}

func (DomainAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAttachmentArgs)(nil)).Elem()
}

type DomainAttachmentInput interface {
	pulumi.Input

	ToDomainAttachmentOutput() DomainAttachmentOutput
	ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput
}

func (*DomainAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAttachment)(nil)).Elem()
}

func (i *DomainAttachment) ToDomainAttachmentOutput() DomainAttachmentOutput {
	return i.ToDomainAttachmentOutputWithContext(context.Background())
}

func (i *DomainAttachment) ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentOutput)
}

// DomainAttachmentArrayInput is an input type that accepts DomainAttachmentArray and DomainAttachmentArrayOutput values.
// You can construct a concrete instance of `DomainAttachmentArrayInput` via:
//
//	DomainAttachmentArray{ DomainAttachmentArgs{...} }
type DomainAttachmentArrayInput interface {
	pulumi.Input

	ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput
	ToDomainAttachmentArrayOutputWithContext(context.Context) DomainAttachmentArrayOutput
}

type DomainAttachmentArray []DomainAttachmentInput

func (DomainAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainAttachment)(nil)).Elem()
}

func (i DomainAttachmentArray) ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput {
	return i.ToDomainAttachmentArrayOutputWithContext(context.Background())
}

func (i DomainAttachmentArray) ToDomainAttachmentArrayOutputWithContext(ctx context.Context) DomainAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentArrayOutput)
}

// DomainAttachmentMapInput is an input type that accepts DomainAttachmentMap and DomainAttachmentMapOutput values.
// You can construct a concrete instance of `DomainAttachmentMapInput` via:
//
//	DomainAttachmentMap{ "key": DomainAttachmentArgs{...} }
type DomainAttachmentMapInput interface {
	pulumi.Input

	ToDomainAttachmentMapOutput() DomainAttachmentMapOutput
	ToDomainAttachmentMapOutputWithContext(context.Context) DomainAttachmentMapOutput
}

type DomainAttachmentMap map[string]DomainAttachmentInput

func (DomainAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainAttachment)(nil)).Elem()
}

func (i DomainAttachmentMap) ToDomainAttachmentMapOutput() DomainAttachmentMapOutput {
	return i.ToDomainAttachmentMapOutputWithContext(context.Background())
}

func (i DomainAttachmentMap) ToDomainAttachmentMapOutputWithContext(ctx context.Context) DomainAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAttachmentMapOutput)
}

type DomainAttachmentOutput struct{ *pulumi.OutputState }

func (DomainAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAttachment)(nil)).Elem()
}

func (o DomainAttachmentOutput) ToDomainAttachmentOutput() DomainAttachmentOutput {
	return o
}

func (o DomainAttachmentOutput) ToDomainAttachmentOutputWithContext(ctx context.Context) DomainAttachmentOutput {
	return o
}

// The domain names bound to the DNS instance.
func (o DomainAttachmentOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainAttachment) pulumi.StringArrayOutput { return v.DomainNames }).(pulumi.StringArrayOutput)
}

// The id of the DNS instance.
func (o DomainAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type DomainAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DomainAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainAttachment)(nil)).Elem()
}

func (o DomainAttachmentArrayOutput) ToDomainAttachmentArrayOutput() DomainAttachmentArrayOutput {
	return o
}

func (o DomainAttachmentArrayOutput) ToDomainAttachmentArrayOutputWithContext(ctx context.Context) DomainAttachmentArrayOutput {
	return o
}

func (o DomainAttachmentArrayOutput) Index(i pulumi.IntInput) DomainAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainAttachment {
		return vs[0].([]*DomainAttachment)[vs[1].(int)]
	}).(DomainAttachmentOutput)
}

type DomainAttachmentMapOutput struct{ *pulumi.OutputState }

func (DomainAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainAttachment)(nil)).Elem()
}

func (o DomainAttachmentMapOutput) ToDomainAttachmentMapOutput() DomainAttachmentMapOutput {
	return o
}

func (o DomainAttachmentMapOutput) ToDomainAttachmentMapOutputWithContext(ctx context.Context) DomainAttachmentMapOutput {
	return o
}

func (o DomainAttachmentMapOutput) MapIndex(k pulumi.StringInput) DomainAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainAttachment {
		return vs[0].(map[string]*DomainAttachment)[vs[1].(string)]
	}).(DomainAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAttachmentInput)(nil)).Elem(), &DomainAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAttachmentArrayInput)(nil)).Elem(), DomainAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAttachmentMapInput)(nil)).Elem(), DomainAttachmentMap{})
	pulumi.RegisterOutputType(DomainAttachmentOutput{})
	pulumi.RegisterOutputType(DomainAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DomainAttachmentMapOutput{})
}
