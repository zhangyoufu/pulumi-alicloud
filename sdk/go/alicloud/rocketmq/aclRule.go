// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rocketmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sag Acl Rule resource. This topic describes how to configure an access control list (ACL) rule for a target Smart Access Gateway instance to permit or deny access to or from specified IP addresses in the ACL rule.
//
// For information about Sag Acl Rule and how to use it, see [What is access control list (ACL) rule](https://www.alibabacloud.com/help/doc-detail/111483.htm).
//
// > **NOTE:** Available in 1.60.0+
//
// > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_acl_rule.html.markdown.
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
	if args == nil || args.AclId == nil {
		return nil, errors.New("missing required argument 'AclId'")
	}
	if args == nil || args.DestCidr == nil {
		return nil, errors.New("missing required argument 'DestCidr'")
	}
	if args == nil || args.DestPortRange == nil {
		return nil, errors.New("missing required argument 'DestPortRange'")
	}
	if args == nil || args.Direction == nil {
		return nil, errors.New("missing required argument 'Direction'")
	}
	if args == nil || args.IpProtocol == nil {
		return nil, errors.New("missing required argument 'IpProtocol'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.SourceCidr == nil {
		return nil, errors.New("missing required argument 'SourceCidr'")
	}
	if args == nil || args.SourcePortRange == nil {
		return nil, errors.New("missing required argument 'SourcePortRange'")
	}
	if args == nil {
		args = &AclRuleArgs{}
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
