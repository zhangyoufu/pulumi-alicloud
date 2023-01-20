// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DCDN Waf Policy Domain Attachment resource.
//
// For information about DCDN Waf Policy Domain Attachment and how to use it, see [What is Waf Policy Domain Attachment](https://www.alibabacloud.com/help/en/dynamic-route-for-cdn/latest/modify-the-domain-name-bound-to-a-protection-policies).
//
// > **NOTE:** Available in v1.186.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultDomain, err := dcdn.NewDomain(ctx, "defaultDomain", &dcdn.DomainArgs{
//				DomainName: pulumi.String("example_domain_name"),
//				Sources: dcdn.DomainSourceArray{
//					&dcdn.DomainSourceArgs{
//						Content:  pulumi.String("1.1.1.1"),
//						Port:     pulumi.Int(80),
//						Priority: pulumi.String("20"),
//						Type:     pulumi.String("ipaddr"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultWafDomain, err := dcdn.NewWafDomain(ctx, "defaultWafDomain", &dcdn.WafDomainArgs{
//				DomainName:  defaultDomain.DomainName,
//				ClientIpTag: pulumi.String("X-Forwarded-For"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultWafPolicy, err := dcdn.NewWafPolicy(ctx, "defaultWafPolicy", &dcdn.WafPolicyArgs{
//				PolicyType:   pulumi.String("custom"),
//				PolicyName:   pulumi.String("example_value"),
//				DefenseScene: pulumi.String("waf_group"),
//				Status:       pulumi.String("on"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcdn.NewWafPolicyDomainAttachment(ctx, "example", &dcdn.WafPolicyDomainAttachmentArgs{
//				DomainName: defaultWafDomain.DomainName,
//				PolicyId:   defaultWafPolicy.ID(),
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
// DCDN Waf Policy Domain Attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment example policy_id:domain_name
//
// ```
type WafPolicyDomainAttachment struct {
	pulumi.CustomResourceState

	// Access the accelerated domain name of the specified protection policy.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The protection policy ID. Only one input is supported.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
}

// NewWafPolicyDomainAttachment registers a new resource with the given unique name, arguments, and options.
func NewWafPolicyDomainAttachment(ctx *pulumi.Context,
	name string, args *WafPolicyDomainAttachmentArgs, opts ...pulumi.ResourceOption) (*WafPolicyDomainAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	var resource WafPolicyDomainAttachment
	err := ctx.RegisterResource("alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWafPolicyDomainAttachment gets an existing WafPolicyDomainAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWafPolicyDomainAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WafPolicyDomainAttachmentState, opts ...pulumi.ResourceOption) (*WafPolicyDomainAttachment, error) {
	var resource WafPolicyDomainAttachment
	err := ctx.ReadResource("alicloud:dcdn/wafPolicyDomainAttachment:WafPolicyDomainAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WafPolicyDomainAttachment resources.
type wafPolicyDomainAttachmentState struct {
	// Access the accelerated domain name of the specified protection policy.
	DomainName *string `pulumi:"domainName"`
	// The protection policy ID. Only one input is supported.
	PolicyId *string `pulumi:"policyId"`
}

type WafPolicyDomainAttachmentState struct {
	// Access the accelerated domain name of the specified protection policy.
	DomainName pulumi.StringPtrInput
	// The protection policy ID. Only one input is supported.
	PolicyId pulumi.StringPtrInput
}

func (WafPolicyDomainAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*wafPolicyDomainAttachmentState)(nil)).Elem()
}

type wafPolicyDomainAttachmentArgs struct {
	// Access the accelerated domain name of the specified protection policy.
	DomainName string `pulumi:"domainName"`
	// The protection policy ID. Only one input is supported.
	PolicyId string `pulumi:"policyId"`
}

// The set of arguments for constructing a WafPolicyDomainAttachment resource.
type WafPolicyDomainAttachmentArgs struct {
	// Access the accelerated domain name of the specified protection policy.
	DomainName pulumi.StringInput
	// The protection policy ID. Only one input is supported.
	PolicyId pulumi.StringInput
}

func (WafPolicyDomainAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wafPolicyDomainAttachmentArgs)(nil)).Elem()
}

type WafPolicyDomainAttachmentInput interface {
	pulumi.Input

	ToWafPolicyDomainAttachmentOutput() WafPolicyDomainAttachmentOutput
	ToWafPolicyDomainAttachmentOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentOutput
}

func (*WafPolicyDomainAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**WafPolicyDomainAttachment)(nil)).Elem()
}

func (i *WafPolicyDomainAttachment) ToWafPolicyDomainAttachmentOutput() WafPolicyDomainAttachmentOutput {
	return i.ToWafPolicyDomainAttachmentOutputWithContext(context.Background())
}

func (i *WafPolicyDomainAttachment) ToWafPolicyDomainAttachmentOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafPolicyDomainAttachmentOutput)
}

// WafPolicyDomainAttachmentArrayInput is an input type that accepts WafPolicyDomainAttachmentArray and WafPolicyDomainAttachmentArrayOutput values.
// You can construct a concrete instance of `WafPolicyDomainAttachmentArrayInput` via:
//
//	WafPolicyDomainAttachmentArray{ WafPolicyDomainAttachmentArgs{...} }
type WafPolicyDomainAttachmentArrayInput interface {
	pulumi.Input

	ToWafPolicyDomainAttachmentArrayOutput() WafPolicyDomainAttachmentArrayOutput
	ToWafPolicyDomainAttachmentArrayOutputWithContext(context.Context) WafPolicyDomainAttachmentArrayOutput
}

type WafPolicyDomainAttachmentArray []WafPolicyDomainAttachmentInput

func (WafPolicyDomainAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafPolicyDomainAttachment)(nil)).Elem()
}

func (i WafPolicyDomainAttachmentArray) ToWafPolicyDomainAttachmentArrayOutput() WafPolicyDomainAttachmentArrayOutput {
	return i.ToWafPolicyDomainAttachmentArrayOutputWithContext(context.Background())
}

func (i WafPolicyDomainAttachmentArray) ToWafPolicyDomainAttachmentArrayOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafPolicyDomainAttachmentArrayOutput)
}

// WafPolicyDomainAttachmentMapInput is an input type that accepts WafPolicyDomainAttachmentMap and WafPolicyDomainAttachmentMapOutput values.
// You can construct a concrete instance of `WafPolicyDomainAttachmentMapInput` via:
//
//	WafPolicyDomainAttachmentMap{ "key": WafPolicyDomainAttachmentArgs{...} }
type WafPolicyDomainAttachmentMapInput interface {
	pulumi.Input

	ToWafPolicyDomainAttachmentMapOutput() WafPolicyDomainAttachmentMapOutput
	ToWafPolicyDomainAttachmentMapOutputWithContext(context.Context) WafPolicyDomainAttachmentMapOutput
}

type WafPolicyDomainAttachmentMap map[string]WafPolicyDomainAttachmentInput

func (WafPolicyDomainAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafPolicyDomainAttachment)(nil)).Elem()
}

func (i WafPolicyDomainAttachmentMap) ToWafPolicyDomainAttachmentMapOutput() WafPolicyDomainAttachmentMapOutput {
	return i.ToWafPolicyDomainAttachmentMapOutputWithContext(context.Background())
}

func (i WafPolicyDomainAttachmentMap) ToWafPolicyDomainAttachmentMapOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WafPolicyDomainAttachmentMapOutput)
}

type WafPolicyDomainAttachmentOutput struct{ *pulumi.OutputState }

func (WafPolicyDomainAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WafPolicyDomainAttachment)(nil)).Elem()
}

func (o WafPolicyDomainAttachmentOutput) ToWafPolicyDomainAttachmentOutput() WafPolicyDomainAttachmentOutput {
	return o
}

func (o WafPolicyDomainAttachmentOutput) ToWafPolicyDomainAttachmentOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentOutput {
	return o
}

// Access the accelerated domain name of the specified protection policy.
func (o WafPolicyDomainAttachmentOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *WafPolicyDomainAttachment) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The protection policy ID. Only one input is supported.
func (o WafPolicyDomainAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *WafPolicyDomainAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

type WafPolicyDomainAttachmentArrayOutput struct{ *pulumi.OutputState }

func (WafPolicyDomainAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WafPolicyDomainAttachment)(nil)).Elem()
}

func (o WafPolicyDomainAttachmentArrayOutput) ToWafPolicyDomainAttachmentArrayOutput() WafPolicyDomainAttachmentArrayOutput {
	return o
}

func (o WafPolicyDomainAttachmentArrayOutput) ToWafPolicyDomainAttachmentArrayOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentArrayOutput {
	return o
}

func (o WafPolicyDomainAttachmentArrayOutput) Index(i pulumi.IntInput) WafPolicyDomainAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WafPolicyDomainAttachment {
		return vs[0].([]*WafPolicyDomainAttachment)[vs[1].(int)]
	}).(WafPolicyDomainAttachmentOutput)
}

type WafPolicyDomainAttachmentMapOutput struct{ *pulumi.OutputState }

func (WafPolicyDomainAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WafPolicyDomainAttachment)(nil)).Elem()
}

func (o WafPolicyDomainAttachmentMapOutput) ToWafPolicyDomainAttachmentMapOutput() WafPolicyDomainAttachmentMapOutput {
	return o
}

func (o WafPolicyDomainAttachmentMapOutput) ToWafPolicyDomainAttachmentMapOutputWithContext(ctx context.Context) WafPolicyDomainAttachmentMapOutput {
	return o
}

func (o WafPolicyDomainAttachmentMapOutput) MapIndex(k pulumi.StringInput) WafPolicyDomainAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WafPolicyDomainAttachment {
		return vs[0].(map[string]*WafPolicyDomainAttachment)[vs[1].(string)]
	}).(WafPolicyDomainAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WafPolicyDomainAttachmentInput)(nil)).Elem(), &WafPolicyDomainAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafPolicyDomainAttachmentArrayInput)(nil)).Elem(), WafPolicyDomainAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WafPolicyDomainAttachmentMapInput)(nil)).Elem(), WafPolicyDomainAttachmentMap{})
	pulumi.RegisterOutputType(WafPolicyDomainAttachmentOutput{})
	pulumi.RegisterOutputType(WafPolicyDomainAttachmentArrayOutput{})
	pulumi.RegisterOutputType(WafPolicyDomainAttachmentMapOutput{})
}
