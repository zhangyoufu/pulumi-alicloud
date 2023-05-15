// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM Policy resource.
//
// > **NOTE:** When you want to destroy this resource forcefully(means remove all the relationships associated with it automatically and then destroy it) without set `force`  with `true` at beginning, you need add `force = true` to configuration file and run `pulumi preview`, then you can delete resource forcefully.
//
// > **NOTE:** Each policy can own at most 5 versions and the oldest version will be removed after its version achieves 5.
//
// > **NOTE:** If the policy has multiple versions, all non-default versions will be deleted first when deleting policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ram.NewPolicy(ctx, "policy", &ram.PolicyArgs{
//				Description:    pulumi.String("this is a policy test"),
//				Force:          pulumi.Bool(true),
//				PolicyDocument: pulumi.String("  {\n    \"Statement\": [\n      {\n        \"Action\": [\n          \"oss:ListObjects\",\n          \"oss:GetObject\"\n        ],\n        \"Effect\": \"Allow\",\n        \"Resource\": [\n          \"acs:oss:*:*:mybucket\",\n          \"acs:oss:*:*:mybucket/*\"\n        ]\n      }\n    ],\n      \"Version\": \"1\"\n  }\n  \n"),
//				PolicyName:     pulumi.String("policyName"),
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
// RAM policy can be imported using the id or name, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/policy:Policy example my-policy
//
// ```
type Policy struct {
	pulumi.CustomResourceState

	// The policy attachment count.
	AttachmentCount pulumi.IntOutput `pulumi:"attachmentCount"`
	// The default version of policy.
	DefaultVersion pulumi.StringOutput `pulumi:"defaultVersion"`
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
	//
	// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
	Document pulumi.StringOutput `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// It has been deprecated from provider version 1.114.0 and `policyName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
	RotateStrategy pulumi.StringPtrOutput `pulumi:"rotateStrategy"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayOutput `pulumi:"statements"`
	// The policy type.
	Type pulumi.StringOutput `pulumi:"type"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// The ID of default version policy.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	var resource Policy
	err := ctx.RegisterResource("alicloud:ram/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("alicloud:ram/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The policy attachment count.
	AttachmentCount *int `pulumi:"attachmentCount"`
	// The default version of policy.
	DefaultVersion *string `pulumi:"defaultVersion"`
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description *string `pulumi:"description"`
	// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
	//
	// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// It has been deprecated from provider version 1.114.0 and `policyName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
	Name *string `pulumi:"name"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName *string `pulumi:"policyName"`
	// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
	RotateStrategy *string `pulumi:"rotateStrategy"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements []PolicyStatement `pulumi:"statements"`
	// The policy type.
	Type *string `pulumi:"type"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
	// The ID of default version policy.
	VersionId *string `pulumi:"versionId"`
}

type PolicyState struct {
	// The policy attachment count.
	AttachmentCount pulumi.IntPtrInput
	// The default version of policy.
	DefaultVersion pulumi.StringPtrInput
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrInput
	// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
	//
	// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.114.0 and `policyName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
	Name pulumi.StringPtrInput
	// Document of the RAM policy. It is required when the `statement` is not specified.
	PolicyDocument pulumi.StringPtrInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringPtrInput
	// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
	RotateStrategy pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayInput
	// The policy type.
	Type pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
	// The ID of default version policy.
	VersionId pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description *string `pulumi:"description"`
	// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
	//
	// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// It has been deprecated from provider version 1.114.0 and `policyName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
	Name *string `pulumi:"name"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName *string `pulumi:"policyName"`
	// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
	RotateStrategy *string `pulumi:"rotateStrategy"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements []PolicyStatement `pulumi:"statements"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrInput
	// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
	//
	// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// It has been deprecated from provider version 1.114.0 and `policyName` instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
	Name pulumi.StringPtrInput
	// Document of the RAM policy. It is required when the `statement` is not specified.
	PolicyDocument pulumi.StringPtrInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringPtrInput
	// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
	RotateStrategy pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// The policy attachment count.
func (o PolicyOutput) AttachmentCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy) pulumi.IntOutput { return v.AttachmentCount }).(pulumi.IntOutput)
}

// The default version of policy.
func (o PolicyOutput) DefaultVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.DefaultVersion }).(pulumi.StringOutput)
}

// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
func (o PolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// It has been deprecated from provider version 1.114.0 and `policyDocument` instead.
//
// Deprecated: Field 'document' has been deprecated from provider version 1.114.0. New field 'policy_document' instead.
func (o PolicyOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Document }).(pulumi.StringOutput)
}

// This parameter is used for resource destroy. Default value is `false`.
func (o PolicyOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// It has been deprecated from provider version 1.114.0 and `policyName` instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.114.0. New field 'policy_name' instead.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Document of the RAM policy. It is required when the `statement` is not specified.
func (o PolicyOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
func (o PolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// The rotation strategy of the policy. You can use this parameter to delete an early policy version. Valid Values: `None`, `DeleteOldestNonDefaultVersionWhenLimitExceeded`. Default to `None`.
func (o PolicyOutput) RotateStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.RotateStrategy }).(pulumi.StringPtrOutput)
}

// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
//
// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
func (o PolicyOutput) Statements() PolicyStatementArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyStatementArrayOutput { return v.Statements }).(PolicyStatementArrayOutput)
}

// The policy type.
func (o PolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
//
// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
func (o PolicyOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

// The ID of default version policy.
func (o PolicyOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
