// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PolicyVersion struct {
	pulumi.CustomResourceState

	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrOutput `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewPolicyVersion registers a new resource with the given unique name, arguments, and options.
func NewPolicyVersion(ctx *pulumi.Context,
	name string, args *PolicyVersionArgs, opts ...pulumi.ResourceOption) (*PolicyVersion, error) {
	if args == nil || args.PolicyDocument == nil {
		return nil, errors.New("missing required argument 'PolicyDocument'")
	}
	if args == nil || args.PolicyName == nil {
		return nil, errors.New("missing required argument 'PolicyName'")
	}
	if args == nil {
		args = &PolicyVersionArgs{}
	}
	var resource PolicyVersion
	err := ctx.RegisterResource("alicloud:resourcemanager/policyVersion:PolicyVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyVersion gets an existing PolicyVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyVersionState, opts ...pulumi.ResourceOption) (*PolicyVersion, error) {
	var resource PolicyVersion
	err := ctx.ReadResource("alicloud:resourcemanager/policyVersion:PolicyVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyVersion resources.
type policyVersionState struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `pulumi:"policyName"`
}

type PolicyVersionState struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrInput
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringPtrInput
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringPtrInput
}

func (PolicyVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVersionState)(nil)).Elem()
}

type policyVersionArgs struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion *bool `pulumi:"isDefaultVersion"`
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument string `pulumi:"policyDocument"`
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a PolicyVersion resource.
type PolicyVersionArgs struct {
	// Specifies whether to set the policy version as the default version. Default to `false`.
	//
	// Deprecated: Field 'is_default_version' has been deprecated from provider version 1.90.0
	IsDefaultVersion pulumi.BoolPtrInput
	// The content of the policy. The content must be 1 to 2,048 characters in length.
	PolicyDocument pulumi.StringInput
	// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringInput
}

func (PolicyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVersionArgs)(nil)).Elem()
}
