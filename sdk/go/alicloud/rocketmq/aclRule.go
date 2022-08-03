// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sag Acl Rule resource. This topic describes how to configure an access control list (ACL) rule for a target Smart Access Gateway instance to permit or deny access to or from specified IP addresses in the ACL rule.
//
// For information about Sag Acl Rule and how to use it, see [What is access control list (ACL) rule](https://www.alibabacloud.com/help/doc-detail/111483.htm).
//
// > **NOTE:** Available in 1.60.0+
//
// > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
// ## Import
//
// The Sag Acl Rule can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:rocketmq/aclRule:AclRule example acr-abc123456
// ```
type AclRule struct {
	pulumi.CustomResourceState

	// The ID of the ACL.
	AclId pulumi.StringOutput `pulumi:"aclId"`
	// The description of the ACL rule. It must be 1 to 512 characters in length.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
	DestCidr pulumi.StringOutput `pulumi:"destCidr"`
	// The range of the destination port. Valid value: 80/80.
	DestPortRange pulumi.StringOutput `pulumi:"destPortRange"`
	// The direction of the ACL rule. Valid values: in|out.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The protocol used by the ACL rule. The value is not case sensitive.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The policy used by the ACL rule. Valid values: accept|drop.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The priority of the ACL rule. Value range: 1 to 100.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
	SourceCidr pulumi.StringOutput `pulumi:"sourceCidr"`
	// The range of the source port. Valid value: 80/80.
	SourcePortRange pulumi.StringOutput `pulumi:"sourcePortRange"`
}

// NewAclRule registers a new resource with the given unique name, arguments, and options.
func NewAclRule(ctx *pulumi.Context,
	name string, args *AclRuleArgs, opts ...pulumi.ResourceOption) (*AclRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclId == nil {
		return nil, errors.New("invalid value for required argument 'AclId'")
	}
	if args.DestCidr == nil {
		return nil, errors.New("invalid value for required argument 'DestCidr'")
	}
	if args.DestPortRange == nil {
		return nil, errors.New("invalid value for required argument 'DestPortRange'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.IpProtocol == nil {
		return nil, errors.New("invalid value for required argument 'IpProtocol'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.SourceCidr == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidr'")
	}
	if args.SourcePortRange == nil {
		return nil, errors.New("invalid value for required argument 'SourcePortRange'")
	}
	var resource AclRule
	err := ctx.RegisterResource("alicloud:rocketmq/aclRule:AclRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAclRule gets an existing AclRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAclRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclRuleState, opts ...pulumi.ResourceOption) (*AclRule, error) {
	var resource AclRule
	err := ctx.ReadResource("alicloud:rocketmq/aclRule:AclRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AclRule resources.
type aclRuleState struct {
	// The ID of the ACL.
	AclId *string `pulumi:"aclId"`
	// The description of the ACL rule. It must be 1 to 512 characters in length.
	Description *string `pulumi:"description"`
	// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
	DestCidr *string `pulumi:"destCidr"`
	// The range of the destination port. Valid value: 80/80.
	DestPortRange *string `pulumi:"destPortRange"`
	// The direction of the ACL rule. Valid values: in|out.
	Direction *string `pulumi:"direction"`
	// The protocol used by the ACL rule. The value is not case sensitive.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The policy used by the ACL rule. Valid values: accept|drop.
	Policy *string `pulumi:"policy"`
	// The priority of the ACL rule. Value range: 1 to 100.
	Priority *int `pulumi:"priority"`
	// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
	SourceCidr *string `pulumi:"sourceCidr"`
	// The range of the source port. Valid value: 80/80.
	SourcePortRange *string `pulumi:"sourcePortRange"`
}

type AclRuleState struct {
	// The ID of the ACL.
	AclId pulumi.StringPtrInput
	// The description of the ACL rule. It must be 1 to 512 characters in length.
	Description pulumi.StringPtrInput
	// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
	DestCidr pulumi.StringPtrInput
	// The range of the destination port. Valid value: 80/80.
	DestPortRange pulumi.StringPtrInput
	// The direction of the ACL rule. Valid values: in|out.
	Direction pulumi.StringPtrInput
	// The protocol used by the ACL rule. The value is not case sensitive.
	IpProtocol pulumi.StringPtrInput
	// The policy used by the ACL rule. Valid values: accept|drop.
	Policy pulumi.StringPtrInput
	// The priority of the ACL rule. Value range: 1 to 100.
	Priority pulumi.IntPtrInput
	// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
	SourceCidr pulumi.StringPtrInput
	// The range of the source port. Valid value: 80/80.
	SourcePortRange pulumi.StringPtrInput
}

func (AclRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleState)(nil)).Elem()
}

type aclRuleArgs struct {
	// The ID of the ACL.
	AclId string `pulumi:"aclId"`
	// The description of the ACL rule. It must be 1 to 512 characters in length.
	Description *string `pulumi:"description"`
	// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
	DestCidr string `pulumi:"destCidr"`
	// The range of the destination port. Valid value: 80/80.
	DestPortRange string `pulumi:"destPortRange"`
	// The direction of the ACL rule. Valid values: in|out.
	Direction string `pulumi:"direction"`
	// The protocol used by the ACL rule. The value is not case sensitive.
	IpProtocol string `pulumi:"ipProtocol"`
	// The policy used by the ACL rule. Valid values: accept|drop.
	Policy string `pulumi:"policy"`
	// The priority of the ACL rule. Value range: 1 to 100.
	Priority *int `pulumi:"priority"`
	// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
	SourceCidr string `pulumi:"sourceCidr"`
	// The range of the source port. Valid value: 80/80.
	SourcePortRange string `pulumi:"sourcePortRange"`
}

// The set of arguments for constructing a AclRule resource.
type AclRuleArgs struct {
	// The ID of the ACL.
	AclId pulumi.StringInput
	// The description of the ACL rule. It must be 1 to 512 characters in length.
	Description pulumi.StringPtrInput
	// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
	DestCidr pulumi.StringInput
	// The range of the destination port. Valid value: 80/80.
	DestPortRange pulumi.StringInput
	// The direction of the ACL rule. Valid values: in|out.
	Direction pulumi.StringInput
	// The protocol used by the ACL rule. The value is not case sensitive.
	IpProtocol pulumi.StringInput
	// The policy used by the ACL rule. Valid values: accept|drop.
	Policy pulumi.StringInput
	// The priority of the ACL rule. Value range: 1 to 100.
	Priority pulumi.IntPtrInput
	// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
	SourceCidr pulumi.StringInput
	// The range of the source port. Valid value: 80/80.
	SourcePortRange pulumi.StringInput
}

func (AclRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclRuleArgs)(nil)).Elem()
}

type AclRuleInput interface {
	pulumi.Input

	ToAclRuleOutput() AclRuleOutput
	ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput
}

func (*AclRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (i *AclRule) ToAclRuleOutput() AclRuleOutput {
	return i.ToAclRuleOutputWithContext(context.Background())
}

func (i *AclRule) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleOutput)
}

// AclRuleArrayInput is an input type that accepts AclRuleArray and AclRuleArrayOutput values.
// You can construct a concrete instance of `AclRuleArrayInput` via:
//
//          AclRuleArray{ AclRuleArgs{...} }
type AclRuleArrayInput interface {
	pulumi.Input

	ToAclRuleArrayOutput() AclRuleArrayOutput
	ToAclRuleArrayOutputWithContext(context.Context) AclRuleArrayOutput
}

type AclRuleArray []AclRuleInput

func (AclRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (i AclRuleArray) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return i.ToAclRuleArrayOutputWithContext(context.Background())
}

func (i AclRuleArray) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleArrayOutput)
}

// AclRuleMapInput is an input type that accepts AclRuleMap and AclRuleMapOutput values.
// You can construct a concrete instance of `AclRuleMapInput` via:
//
//          AclRuleMap{ "key": AclRuleArgs{...} }
type AclRuleMapInput interface {
	pulumi.Input

	ToAclRuleMapOutput() AclRuleMapOutput
	ToAclRuleMapOutputWithContext(context.Context) AclRuleMapOutput
}

type AclRuleMap map[string]AclRuleInput

func (AclRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (i AclRuleMap) ToAclRuleMapOutput() AclRuleMapOutput {
	return i.ToAclRuleMapOutputWithContext(context.Background())
}

func (i AclRuleMap) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclRuleMapOutput)
}

type AclRuleOutput struct{ *pulumi.OutputState }

func (AclRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AclRule)(nil)).Elem()
}

func (o AclRuleOutput) ToAclRuleOutput() AclRuleOutput {
	return o
}

func (o AclRuleOutput) ToAclRuleOutputWithContext(ctx context.Context) AclRuleOutput {
	return o
}

// The ID of the ACL.
func (o AclRuleOutput) AclId() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.AclId }).(pulumi.StringOutput)
}

// The description of the ACL rule. It must be 1 to 512 characters in length.
func (o AclRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.
func (o AclRuleOutput) DestCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.DestCidr }).(pulumi.StringOutput)
}

// The range of the destination port. Valid value: 80/80.
func (o AclRuleOutput) DestPortRange() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.DestPortRange }).(pulumi.StringOutput)
}

// The direction of the ACL rule. Valid values: in|out.
func (o AclRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// The protocol used by the ACL rule. The value is not case sensitive.
func (o AclRuleOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.IpProtocol }).(pulumi.StringOutput)
}

// The policy used by the ACL rule. Valid values: accept|drop.
func (o AclRuleOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The priority of the ACL rule. Value range: 1 to 100.
func (o AclRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AclRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.
func (o AclRuleOutput) SourceCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.SourceCidr }).(pulumi.StringOutput)
}

// The range of the source port. Valid value: 80/80.
func (o AclRuleOutput) SourcePortRange() pulumi.StringOutput {
	return o.ApplyT(func(v *AclRule) pulumi.StringOutput { return v.SourcePortRange }).(pulumi.StringOutput)
}

type AclRuleArrayOutput struct{ *pulumi.OutputState }

func (AclRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AclRule)(nil)).Elem()
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutput() AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) ToAclRuleArrayOutputWithContext(ctx context.Context) AclRuleArrayOutput {
	return o
}

func (o AclRuleArrayOutput) Index(i pulumi.IntInput) AclRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].([]*AclRule)[vs[1].(int)]
	}).(AclRuleOutput)
}

type AclRuleMapOutput struct{ *pulumi.OutputState }

func (AclRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AclRule)(nil)).Elem()
}

func (o AclRuleMapOutput) ToAclRuleMapOutput() AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) ToAclRuleMapOutputWithContext(ctx context.Context) AclRuleMapOutput {
	return o
}

func (o AclRuleMapOutput) MapIndex(k pulumi.StringInput) AclRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AclRule {
		return vs[0].(map[string]*AclRule)[vs[1].(string)]
	}).(AclRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleInput)(nil)).Elem(), &AclRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleArrayInput)(nil)).Elem(), AclRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclRuleMapInput)(nil)).Elem(), AclRuleMap{})
	pulumi.RegisterOutputType(AclRuleOutput{})
	pulumi.RegisterOutputType(AclRuleArrayOutput{})
	pulumi.RegisterOutputType(AclRuleMapOutput{})
}
