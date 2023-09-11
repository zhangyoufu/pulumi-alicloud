// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Resource Manager Policy Attachment resource to attaches a policy to an object. After you attach a policy to an object, the object has the operation permissions on the current resource group or the resources under the current account.
// For information about Resource Manager Policy Attachment and how to use it, see [How to authorize and manage resource groups](https://www.alibabacloud.com/help/en/doc-detail/94490.htm).
//
// > **NOTE:** Available in v1.93.0+.
//
// ## Import
//
// Resource Manager Policy Attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:resourcemanager/policyAttachment:PolicyAttachment example tf-testaccrdpolicy:Custom:tf-testaccrdpolicy@11827252********.onaliyun.com:IMSUser:rg******
//
// ```
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// The name of the object to which you want to attach the policy.
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	if args.PrincipalName == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalName'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyAttachment
	err := ctx.RegisterResource("alicloud:resourcemanager/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAttachmentState, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	var resource PolicyAttachment
	err := ctx.ReadResource("alicloud:resourcemanager/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName *string `pulumi:"policyName"`
	// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
	PolicyType *string `pulumi:"policyType"`
	// The name of the object to which you want to attach the policy.
	PrincipalName *string `pulumi:"principalName"`
	// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
	PrincipalType *string `pulumi:"principalType"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
}

type PolicyAttachmentState struct {
	// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringPtrInput
	// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
	PolicyType pulumi.StringPtrInput
	// The name of the object to which you want to attach the policy.
	PrincipalName pulumi.StringPtrInput
	// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
	PrincipalType pulumi.StringPtrInput
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId pulumi.StringPtrInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName string `pulumi:"policyName"`
	// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
	PolicyType string `pulumi:"policyType"`
	// The name of the object to which you want to attach the policy.
	PrincipalName string `pulumi:"principalName"`
	// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
	PrincipalType string `pulumi:"principalType"`
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	PolicyName pulumi.StringInput
	// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
	PolicyType pulumi.StringInput
	// The name of the object to which you want to attach the policy.
	PrincipalName pulumi.StringInput
	// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
	PrincipalType pulumi.StringInput
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	ResourceGroupId pulumi.StringInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

type PolicyAttachmentInput interface {
	pulumi.Input

	ToPolicyAttachmentOutput() PolicyAttachmentOutput
	ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput
}

func (*PolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil)).Elem()
}

func (i *PolicyAttachment) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return i.ToPolicyAttachmentOutputWithContext(context.Background())
}

func (i *PolicyAttachment) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentOutput)
}

func (i *PolicyAttachment) ToOutput(ctx context.Context) pulumix.Output[*PolicyAttachment] {
	return pulumix.Output[*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentOutputWithContext(ctx).OutputState,
	}
}

// PolicyAttachmentArrayInput is an input type that accepts PolicyAttachmentArray and PolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `PolicyAttachmentArrayInput` via:
//
//	PolicyAttachmentArray{ PolicyAttachmentArgs{...} }
type PolicyAttachmentArrayInput interface {
	pulumi.Input

	ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput
	ToPolicyAttachmentArrayOutputWithContext(context.Context) PolicyAttachmentArrayOutput
}

type PolicyAttachmentArray []PolicyAttachmentInput

func (PolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAttachment)(nil)).Elem()
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return i.ToPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentArrayOutput)
}

func (i PolicyAttachmentArray) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyAttachment] {
	return pulumix.Output[[]*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentArrayOutputWithContext(ctx).OutputState,
	}
}

// PolicyAttachmentMapInput is an input type that accepts PolicyAttachmentMap and PolicyAttachmentMapOutput values.
// You can construct a concrete instance of `PolicyAttachmentMapInput` via:
//
//	PolicyAttachmentMap{ "key": PolicyAttachmentArgs{...} }
type PolicyAttachmentMapInput interface {
	pulumi.Input

	ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput
	ToPolicyAttachmentMapOutputWithContext(context.Context) PolicyAttachmentMapOutput
}

type PolicyAttachmentMap map[string]PolicyAttachmentInput

func (PolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAttachment)(nil)).Elem()
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return i.ToPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentMapOutput)
}

func (i PolicyAttachmentMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyAttachment] {
	return pulumix.Output[map[string]*PolicyAttachment]{
		OutputState: i.ToPolicyAttachmentMapOutputWithContext(ctx).OutputState,
	}
}

type PolicyAttachmentOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToOutput(ctx context.Context) pulumix.Output[*PolicyAttachment] {
	return pulumix.Output[*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

// The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
func (o PolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// - (Required, ForceNew) The type of the policy. Valid values: `Custom`, `System`.
func (o PolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// The name of the object to which you want to attach the policy.
func (o PolicyAttachmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// The type of the object to which you want to attach the policy. Valid values: `IMSUser`: RAM user, `IMSGroup`: RAM user group, `ServiceRole`: RAM role.
func (o PolicyAttachmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
func (o PolicyAttachmentOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

type PolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PolicyAttachment] {
	return pulumix.Output[[]*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o PolicyAttachmentArrayOutput) Index(i pulumi.IntInput) PolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyAttachment {
		return vs[0].([]*PolicyAttachment)[vs[1].(int)]
	}).(PolicyAttachmentOutput)
}

type PolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyAttachment)(nil)).Elem()
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PolicyAttachment] {
	return pulumix.Output[map[string]*PolicyAttachment]{
		OutputState: o.OutputState,
	}
}

func (o PolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) PolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyAttachment {
		return vs[0].(map[string]*PolicyAttachment)[vs[1].(string)]
	}).(PolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentInput)(nil)).Elem(), &PolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentArrayInput)(nil)).Elem(), PolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAttachmentMapInput)(nil)).Elem(), PolicyAttachmentMap{})
	pulumi.RegisterOutputType(PolicyAttachmentOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentMapOutput{})
}
