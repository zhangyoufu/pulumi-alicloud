// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Resource Manager Control Policy resource.
//
// For information about Resource Manager Control Policy and how to use it, see [What is Control Policy](https://help.aliyun.com/document_detail/208287.html).
//
// > **NOTE:** Available in v1.120.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcemanager.NewControlPolicy(ctx, "example", &resourcemanager.ControlPolicyArgs{
// 			ControlPolicyName: pulumi.String("tf-testAccRDControlPolicy"),
// 			Description:       pulumi.String("tf-testAccRDControlPolicy"),
// 			EffectScope:       pulumi.String("RAM"),
// 			PolicyDocument: pulumi.String(fmt.Sprintf(`  {
//     "Version": "1",
//     "Statement": [
//       {
//         "Effect": "Deny",
//         "Action": [
//           "ram:UpdateRole",
//           "ram:DeleteRole",
//           "ram:AttachPolicyToRole",
//           "ram:DetachPolicyFromRole"
//         ],
//         "Resource": "acs:ram:*:*:role/ResourceDirectoryAccountAccessRole"
//       }
//     ]
//   }
//
// `)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Resource Manager Control Policy can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:resourcemanager/controlPolicy:ControlPolicy example <id>
// ```
type ControlPolicy struct {
	pulumi.CustomResourceState

	// The name of control policy.
	ControlPolicyName pulumi.StringOutput `pulumi:"controlPolicyName"`
	// The description of control policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The effect scope. Valid values `RAM`.
	EffectScope pulumi.StringOutput `pulumi:"effectScope"`
	// The policy document of control policy.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
}

// NewControlPolicy registers a new resource with the given unique name, arguments, and options.
func NewControlPolicy(ctx *pulumi.Context,
	name string, args *ControlPolicyArgs, opts ...pulumi.ResourceOption) (*ControlPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'ControlPolicyName'")
	}
	if args.EffectScope == nil {
		return nil, errors.New("invalid value for required argument 'EffectScope'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	var resource ControlPolicy
	err := ctx.RegisterResource("alicloud:resourcemanager/controlPolicy:ControlPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetControlPolicy gets an existing ControlPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetControlPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ControlPolicyState, opts ...pulumi.ResourceOption) (*ControlPolicy, error) {
	var resource ControlPolicy
	err := ctx.ReadResource("alicloud:resourcemanager/controlPolicy:ControlPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ControlPolicy resources.
type controlPolicyState struct {
	// The name of control policy.
	ControlPolicyName *string `pulumi:"controlPolicyName"`
	// The description of control policy.
	Description *string `pulumi:"description"`
	// The effect scope. Valid values `RAM`.
	EffectScope *string `pulumi:"effectScope"`
	// The policy document of control policy.
	PolicyDocument *string `pulumi:"policyDocument"`
}

type ControlPolicyState struct {
	// The name of control policy.
	ControlPolicyName pulumi.StringPtrInput
	// The description of control policy.
	Description pulumi.StringPtrInput
	// The effect scope. Valid values `RAM`.
	EffectScope pulumi.StringPtrInput
	// The policy document of control policy.
	PolicyDocument pulumi.StringPtrInput
}

func (ControlPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyState)(nil)).Elem()
}

type controlPolicyArgs struct {
	// The name of control policy.
	ControlPolicyName string `pulumi:"controlPolicyName"`
	// The description of control policy.
	Description *string `pulumi:"description"`
	// The effect scope. Valid values `RAM`.
	EffectScope string `pulumi:"effectScope"`
	// The policy document of control policy.
	PolicyDocument string `pulumi:"policyDocument"`
}

// The set of arguments for constructing a ControlPolicy resource.
type ControlPolicyArgs struct {
	// The name of control policy.
	ControlPolicyName pulumi.StringInput
	// The description of control policy.
	Description pulumi.StringPtrInput
	// The effect scope. Valid values `RAM`.
	EffectScope pulumi.StringInput
	// The policy document of control policy.
	PolicyDocument pulumi.StringInput
}

func (ControlPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*controlPolicyArgs)(nil)).Elem()
}

type ControlPolicyInput interface {
	pulumi.Input

	ToControlPolicyOutput() ControlPolicyOutput
	ToControlPolicyOutputWithContext(ctx context.Context) ControlPolicyOutput
}

func (*ControlPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicy)(nil)).Elem()
}

func (i *ControlPolicy) ToControlPolicyOutput() ControlPolicyOutput {
	return i.ToControlPolicyOutputWithContext(context.Background())
}

func (i *ControlPolicy) ToControlPolicyOutputWithContext(ctx context.Context) ControlPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyOutput)
}

// ControlPolicyArrayInput is an input type that accepts ControlPolicyArray and ControlPolicyArrayOutput values.
// You can construct a concrete instance of `ControlPolicyArrayInput` via:
//
//          ControlPolicyArray{ ControlPolicyArgs{...} }
type ControlPolicyArrayInput interface {
	pulumi.Input

	ToControlPolicyArrayOutput() ControlPolicyArrayOutput
	ToControlPolicyArrayOutputWithContext(context.Context) ControlPolicyArrayOutput
}

type ControlPolicyArray []ControlPolicyInput

func (ControlPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPolicy)(nil)).Elem()
}

func (i ControlPolicyArray) ToControlPolicyArrayOutput() ControlPolicyArrayOutput {
	return i.ToControlPolicyArrayOutputWithContext(context.Background())
}

func (i ControlPolicyArray) ToControlPolicyArrayOutputWithContext(ctx context.Context) ControlPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyArrayOutput)
}

// ControlPolicyMapInput is an input type that accepts ControlPolicyMap and ControlPolicyMapOutput values.
// You can construct a concrete instance of `ControlPolicyMapInput` via:
//
//          ControlPolicyMap{ "key": ControlPolicyArgs{...} }
type ControlPolicyMapInput interface {
	pulumi.Input

	ToControlPolicyMapOutput() ControlPolicyMapOutput
	ToControlPolicyMapOutputWithContext(context.Context) ControlPolicyMapOutput
}

type ControlPolicyMap map[string]ControlPolicyInput

func (ControlPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPolicy)(nil)).Elem()
}

func (i ControlPolicyMap) ToControlPolicyMapOutput() ControlPolicyMapOutput {
	return i.ToControlPolicyMapOutputWithContext(context.Background())
}

func (i ControlPolicyMap) ToControlPolicyMapOutputWithContext(ctx context.Context) ControlPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControlPolicyMapOutput)
}

type ControlPolicyOutput struct{ *pulumi.OutputState }

func (ControlPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControlPolicy)(nil)).Elem()
}

func (o ControlPolicyOutput) ToControlPolicyOutput() ControlPolicyOutput {
	return o
}

func (o ControlPolicyOutput) ToControlPolicyOutputWithContext(ctx context.Context) ControlPolicyOutput {
	return o
}

// The name of control policy.
func (o ControlPolicyOutput) ControlPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPolicy) pulumi.StringOutput { return v.ControlPolicyName }).(pulumi.StringOutput)
}

// The description of control policy.
func (o ControlPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControlPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The effect scope. Valid values `RAM`.
func (o ControlPolicyOutput) EffectScope() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPolicy) pulumi.StringOutput { return v.EffectScope }).(pulumi.StringOutput)
}

// The policy document of control policy.
func (o ControlPolicyOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *ControlPolicy) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

type ControlPolicyArrayOutput struct{ *pulumi.OutputState }

func (ControlPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ControlPolicy)(nil)).Elem()
}

func (o ControlPolicyArrayOutput) ToControlPolicyArrayOutput() ControlPolicyArrayOutput {
	return o
}

func (o ControlPolicyArrayOutput) ToControlPolicyArrayOutputWithContext(ctx context.Context) ControlPolicyArrayOutput {
	return o
}

func (o ControlPolicyArrayOutput) Index(i pulumi.IntInput) ControlPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ControlPolicy {
		return vs[0].([]*ControlPolicy)[vs[1].(int)]
	}).(ControlPolicyOutput)
}

type ControlPolicyMapOutput struct{ *pulumi.OutputState }

func (ControlPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ControlPolicy)(nil)).Elem()
}

func (o ControlPolicyMapOutput) ToControlPolicyMapOutput() ControlPolicyMapOutput {
	return o
}

func (o ControlPolicyMapOutput) ToControlPolicyMapOutputWithContext(ctx context.Context) ControlPolicyMapOutput {
	return o
}

func (o ControlPolicyMapOutput) MapIndex(k pulumi.StringInput) ControlPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ControlPolicy {
		return vs[0].(map[string]*ControlPolicy)[vs[1].(string)]
	}).(ControlPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyInput)(nil)).Elem(), &ControlPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyArrayInput)(nil)).Elem(), ControlPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ControlPolicyMapInput)(nil)).Elem(), ControlPolicyMap{})
	pulumi.RegisterOutputType(ControlPolicyOutput{})
	pulumi.RegisterOutputType(ControlPolicyArrayOutput{})
	pulumi.RegisterOutputType(ControlPolicyMapOutput{})
}
