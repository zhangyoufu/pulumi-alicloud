// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Manager Control Policy Attachment resource.
//
// For information about Resource Manager Control Policy Attachment and how to use it, see [What is Control Policy Attachment](https://www.alibabacloud.com/help/en/resource-management/latest/api-resourcedirectorymaster-2022-04-19-attachcontrolpolicy).
//
// > **NOTE:** Available since v1.120.0.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			example, err := resourcemanager.NewControlPolicy(ctx, "example", &resourcemanager.ControlPolicyArgs{
//				ControlPolicyName: pulumi.String(name),
//				Description:       pulumi.String(name),
//				EffectScope:       pulumi.String("RAM"),
//				PolicyDocument: pulumi.String(`  {
//	    "Version": "1",
//	    "Statement": [
//	      {
//	        "Effect": "Deny",
//	        "Action": [
//	          "ram:UpdateRole",
//	          "ram:DeleteRole",
//	          "ram:AttachPolicyToRole",
//	          "ram:DetachPolicyFromRole"
//	        ],
//	        "Resource": "acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"
//	      }
//	    ]
//	  }
//
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			exampleFolder, err := resourcemanager.NewFolder(ctx, "example", &resourcemanager.FolderArgs{
//				FolderName: pulumi.Sprintf("%v-%v", name, _default.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourcemanager.NewControlPolicyAttachment(ctx, "example", &resourcemanager.ControlPolicyAttachmentArgs{
//				PolicyId: example.ID(),
//				TargetId: exampleFolder.ID(),
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
// Resource Manager Control Policy Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment example <policy_id>:<target_id>
// ```
type ControlPolicyAttachment struct {
	pulumi.CustomResourceState

	// The ID of control policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The ID of target.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewControlPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewControlPolicyAttachment(ctx *pulumi.Context,
	name string, args *ControlPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*ControlPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ControlPolicyAttachment
	err := ctx.RegisterResource("alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlPolicyAttachment gets an existing ControlPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlPolicyAttachmentState, opts ...pulumi.ResourceOption) (*ControlPolicyAttachment, error) {
	var resource ControlPolicyAttachment
	err := ctx.ReadResource("alicloud:resourcemanager/controlPolicyAttachment:ControlPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlPolicyAttachment resources.
type controlPolicyAttachmentState struct {
	// The ID of control policy.
	PolicyId *string `pulumi:"policyId"`
	// The ID of target.
	TargetId *string `pulumi:"targetId"`
}

type ControlPolicyAttachmentState struct {
	// The ID of control policy.
	PolicyId pulumi.StringPtrInput
	// The ID of target.
	TargetId pulumi.StringPtrInput
}

func (ControlPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyAttachmentState)(nil)).Elem()
}

type controlPolicyAttachmentArgs struct {
	// The ID of control policy.
	PolicyId string `pulumi:"policyId"`
	// The ID of target.
	TargetId string `pulumi:"targetId"`
}

// The set of arguments for constructing a ControlPolicyAttachment resource.
type ControlPolicyAttachmentArgs struct {
	// The ID of control policy.
	PolicyId pulumi.StringInput
	// The ID of target.
	TargetId pulumi.StringInput
}

func (ControlPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyAttachmentArgs)(nil)).Elem()
}

type ControlPolicyAttachmentInput interface {
	pulumi.Input

	ToControlPolicyAttachmentOutput() ControlPolicyAttachmentOutput
	ToControlPolicyAttachmentOutputWithContext(ctx context.Context) ControlPolicyAttachmentOutput
}

func (*ControlPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicyAttachment)(nil)).Elem()
}

func (i *ControlPolicyAttachment) ToControlPolicyAttachmentOutput() ControlPolicyAttachmentOutput {
	return i.ToControlPolicyAttachmentOutputWithContext(context.Background())
}

func (i *ControlPolicyAttachment) ToControlPolicyAttachmentOutputWithContext(ctx context.Context) ControlPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyAttachmentOutput)
}

// ControlPolicyAttachmentArrayInput is an input type that accepts ControlPolicyAttachmentArray and ControlPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `ControlPolicyAttachmentArrayInput` via:
//
//	ControlPolicyAttachmentArray{ ControlPolicyAttachmentArgs{...} }
type ControlPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToControlPolicyAttachmentArrayOutput() ControlPolicyAttachmentArrayOutput
	ToControlPolicyAttachmentArrayOutputWithContext(context.Context) ControlPolicyAttachmentArrayOutput
}

type ControlPolicyAttachmentArray []ControlPolicyAttachmentInput

func (ControlPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPolicyAttachment)(nil)).Elem()
}

func (i ControlPolicyAttachmentArray) ToControlPolicyAttachmentArrayOutput() ControlPolicyAttachmentArrayOutput {
	return i.ToControlPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i ControlPolicyAttachmentArray) ToControlPolicyAttachmentArrayOutputWithContext(ctx context.Context) ControlPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyAttachmentArrayOutput)
}

// ControlPolicyAttachmentMapInput is an input type that accepts ControlPolicyAttachmentMap and ControlPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `ControlPolicyAttachmentMapInput` via:
//
//	ControlPolicyAttachmentMap{ "key": ControlPolicyAttachmentArgs{...} }
type ControlPolicyAttachmentMapInput interface {
	pulumi.Input

	ToControlPolicyAttachmentMapOutput() ControlPolicyAttachmentMapOutput
	ToControlPolicyAttachmentMapOutputWithContext(context.Context) ControlPolicyAttachmentMapOutput
}

type ControlPolicyAttachmentMap map[string]ControlPolicyAttachmentInput

func (ControlPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPolicyAttachment)(nil)).Elem()
}

func (i ControlPolicyAttachmentMap) ToControlPolicyAttachmentMapOutput() ControlPolicyAttachmentMapOutput {
	return i.ToControlPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i ControlPolicyAttachmentMap) ToControlPolicyAttachmentMapOutputWithContext(ctx context.Context) ControlPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyAttachmentMapOutput)
}

type ControlPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (ControlPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicyAttachment)(nil)).Elem()
}

func (o ControlPolicyAttachmentOutput) ToControlPolicyAttachmentOutput() ControlPolicyAttachmentOutput {
	return o
}

func (o ControlPolicyAttachmentOutput) ToControlPolicyAttachmentOutputWithContext(ctx context.Context) ControlPolicyAttachmentOutput {
	return o
}

// The ID of control policy.
func (o ControlPolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The ID of target.
func (o ControlPolicyAttachmentOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPolicyAttachment) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

type ControlPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ControlPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPolicyAttachment)(nil)).Elem()
}

func (o ControlPolicyAttachmentArrayOutput) ToControlPolicyAttachmentArrayOutput() ControlPolicyAttachmentArrayOutput {
	return o
}

func (o ControlPolicyAttachmentArrayOutput) ToControlPolicyAttachmentArrayOutputWithContext(ctx context.Context) ControlPolicyAttachmentArrayOutput {
	return o
}

func (o ControlPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) ControlPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ControlPolicyAttachment {
		return vs[0].([]*ControlPolicyAttachment)[vs[1].(int)]
	}).(ControlPolicyAttachmentOutput)
}

type ControlPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (ControlPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPolicyAttachment)(nil)).Elem()
}

func (o ControlPolicyAttachmentMapOutput) ToControlPolicyAttachmentMapOutput() ControlPolicyAttachmentMapOutput {
	return o
}

func (o ControlPolicyAttachmentMapOutput) ToControlPolicyAttachmentMapOutputWithContext(ctx context.Context) ControlPolicyAttachmentMapOutput {
	return o
}

func (o ControlPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) ControlPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ControlPolicyAttachment {
		return vs[0].(map[string]*ControlPolicyAttachment)[vs[1].(string)]
	}).(ControlPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyAttachmentInput)(nil)).Elem(), &ControlPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyAttachmentArrayInput)(nil)).Elem(), ControlPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyAttachmentMapInput)(nil)).Elem(), ControlPolicyAttachmentMap{})
	pulumi.RegisterOutputType(ControlPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(ControlPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ControlPolicyAttachmentMapOutput{})
}
