// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tag

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Tag Policy Attachment resource to attaches a policy to an object. After you attach a tag policy to an object.
// For information about Tag Policy Attachment and how to use it,
// see [What is Policy Attachment](https://www.alibabacloud.com/help/en/resource-management/latest/attach-policy).
//
// > **NOTE:** Available since v1.204.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/tag"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			examplePolicy, err := tag.NewPolicy(ctx, "examplePolicy", &tag.PolicyArgs{
//				PolicyName:    pulumi.String(name),
//				PolicyDesc:    pulumi.String(name),
//				UserType:      pulumi.String("USER"),
//				PolicyContent: pulumi.String("		{\"tags\":{\"CostCenter\":{\"tag_value\":{\"@@assign\":[\"Beijing\",\"Shanghai\"]},\"tag_key\":{\"@@assign\":\"CostCenter\"}}}}\n"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = tag.NewPolicyAttachment(ctx, "examplePolicyAttachment", &tag.PolicyAttachmentArgs{
//				PolicyId:   examplePolicy.ID(),
//				TargetId:   *pulumi.String(_default.Id),
//				TargetType: pulumi.String("USER"),
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
// Tag Policy Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:tag/policyAttachment:PolicyAttachment example <policy_id>:<target_id>:<target_type>
// ```
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The ID of the tag policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The ID of the object.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
	TargetType pulumi.StringOutput `pulumi:"targetType"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicyAttachment
	err := ctx.RegisterResource("alicloud:tag/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:tag/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The ID of the tag policy.
	PolicyId *string `pulumi:"policyId"`
	// The ID of the object.
	TargetId *string `pulumi:"targetId"`
	// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
	TargetType *string `pulumi:"targetType"`
}

type PolicyAttachmentState struct {
	// The ID of the tag policy.
	PolicyId pulumi.StringPtrInput
	// The ID of the object.
	TargetId pulumi.StringPtrInput
	// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
	TargetType pulumi.StringPtrInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The ID of the tag policy.
	PolicyId string `pulumi:"policyId"`
	// The ID of the object.
	TargetId string `pulumi:"targetId"`
	// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
	TargetType string `pulumi:"targetType"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The ID of the tag policy.
	PolicyId pulumi.StringInput
	// The ID of the object.
	TargetId pulumi.StringInput
	// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
	TargetType pulumi.StringInput
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

// The ID of the tag policy.
func (o PolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The ID of the object.
func (o PolicyAttachmentOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// The type of the object. Valid values: `USER`, `ROOT`, `FOLDER`, `ACCOUNT`.
func (o PolicyAttachmentOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAttachment) pulumi.StringOutput { return v.TargetType }).(pulumi.StringOutput)
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
