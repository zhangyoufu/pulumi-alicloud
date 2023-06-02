// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a security group resource.
//
// > **NOTE:** `ecs.SecurityGroup` is used to build and manage a security group, and `ecs.SecurityGroupRule` can define ingress or egress rules for it.
//
// > **NOTE:** From version 1.7.2, `ecs.SecurityGroup` has supported to segregate different ECS instance in which the same security group.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewSecurityGroup(ctx, "default", &ecs.SecurityGroupArgs{
//				Description: pulumi.String("New security group"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// Basic usage for vpc
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("10.1.0.0/21"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewSecurityGroup(ctx, "group", &ecs.SecurityGroupArgs{
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Module Support
//
// You can use the existing security-group module
// to create a security group and add several rules one-click.
//
// ## Import
//
// Security Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/securityGroup:SecurityGroup example sg-abc123456
//
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The security group description. Defaults to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	//
	// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	InnerAccess pulumi.BoolOutput `pulumi:"innerAccess"`
	// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
	InnerAccessPolicy pulumi.StringOutput `pulumi:"innerAccessPolicy"`
	// The name of the security group. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Id of resource group which the securityGroup belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The type of the security group. Valid values:
	// `normal`: basic security group.
	// `enterprise`: advanced security group For more information.
	SecurityGroupType pulumi.StringPtrOutput `pulumi:"securityGroupType"`
	// A mapping of tags to assign to the resource.
	//
	// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		args = &SecurityGroupArgs{}
	}

	var resource SecurityGroup
	err := ctx.RegisterResource("alicloud:ecs/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("alicloud:ecs/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The security group description. Defaults to null.
	Description *string `pulumi:"description"`
	// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	//
	// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	InnerAccess *bool `pulumi:"innerAccess"`
	// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
	InnerAccessPolicy *string `pulumi:"innerAccessPolicy"`
	// The name of the security group. Defaults to null.
	Name *string `pulumi:"name"`
	// The Id of resource group which the securityGroup belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The type of the security group. Valid values:
	// `normal`: basic security group.
	// `enterprise`: advanced security group For more information.
	SecurityGroupType *string `pulumi:"securityGroupType"`
	// A mapping of tags to assign to the resource.
	//
	// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type SecurityGroupState struct {
	// The security group description. Defaults to null.
	Description pulumi.StringPtrInput
	// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	//
	// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	InnerAccess pulumi.BoolPtrInput
	// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
	InnerAccessPolicy pulumi.StringPtrInput
	// The name of the security group. Defaults to null.
	Name pulumi.StringPtrInput
	// The Id of resource group which the securityGroup belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The type of the security group. Valid values:
	// `normal`: basic security group.
	// `enterprise`: advanced security group For more information.
	SecurityGroupType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	//
	// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The security group description. Defaults to null.
	Description *string `pulumi:"description"`
	// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	//
	// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	InnerAccess *bool `pulumi:"innerAccess"`
	// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
	InnerAccessPolicy *string `pulumi:"innerAccessPolicy"`
	// The name of the security group. Defaults to null.
	Name *string `pulumi:"name"`
	// The Id of resource group which the securityGroup belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The type of the security group. Valid values:
	// `normal`: basic security group.
	// `enterprise`: advanced security group For more information.
	SecurityGroupType *string `pulumi:"securityGroupType"`
	// A mapping of tags to assign to the resource.
	//
	// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The security group description. Defaults to null.
	Description pulumi.StringPtrInput
	// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	//
	// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
	InnerAccess pulumi.BoolPtrInput
	// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
	InnerAccessPolicy pulumi.StringPtrInput
	// The name of the security group. Defaults to null.
	Name pulumi.StringPtrInput
	// The Id of resource group which the securityGroup belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The type of the security group. Valid values:
	// `normal`: basic security group.
	// `enterprise`: advanced security group For more information.
	SecurityGroupType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	//
	// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
	Tags pulumi.MapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//	SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//	SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

// The security group description. Defaults to null.
func (o SecurityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
//
// Deprecated: Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.
func (o SecurityGroupOutput) InnerAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.BoolOutput { return v.InnerAccess }).(pulumi.BoolOutput)
}

// Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]
func (o SecurityGroupOutput) InnerAccessPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.InnerAccessPolicy }).(pulumi.StringOutput)
}

// The name of the security group. Defaults to null.
func (o SecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Id of resource group which the securityGroup belongs.
func (o SecurityGroupOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The type of the security group. Valid values:
// `normal`: basic security group.
// `enterprise`: advanced security group For more information.
func (o SecurityGroupOutput) SecurityGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.SecurityGroupType }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the resource.
//
// Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from version `1.7.2`.
func (o SecurityGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The VPC ID.
func (o SecurityGroupOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].([]*SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].(map[string]*SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupArrayInput)(nil)).Elem(), SecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupMapInput)(nil)).Elem(), SecurityGroupMap{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
