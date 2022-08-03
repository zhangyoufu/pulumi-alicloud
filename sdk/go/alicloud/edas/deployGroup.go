// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EDAS deploy group resource.
//
// > **NOTE:** Available in 1.82.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/edas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := edas.NewDeployGroup(ctx, "default", &edas.DeployGroupArgs{
// 			AppId:     pulumi.Any(_var.App_id),
// 			GroupName: pulumi.Any(_var.Group_name),
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
// EDAS deploy group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:edas/deployGroup:DeployGroup group app_id:group_name:group_id
// ```
type DeployGroup struct {
	pulumi.CustomResourceState

	// The ID of the application that you want to deploy.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType pulumi.IntOutput `pulumi:"groupType"`
}

// NewDeployGroup registers a new resource with the given unique name, arguments, and options.
func NewDeployGroup(ctx *pulumi.Context,
	name string, args *DeployGroupArgs, opts ...pulumi.ResourceOption) (*DeployGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	var resource DeployGroup
	err := ctx.RegisterResource("alicloud:edas/deployGroup:DeployGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployGroup gets an existing DeployGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployGroupState, opts ...pulumi.ResourceOption) (*DeployGroup, error) {
	var resource DeployGroup
	err := ctx.ReadResource("alicloud:edas/deployGroup:DeployGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployGroup resources.
type deployGroupState struct {
	// The ID of the application that you want to deploy.
	AppId *string `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName *string `pulumi:"groupName"`
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType *int `pulumi:"groupType"`
}

type DeployGroupState struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringPtrInput
	// The name of the instance group that you want to create.
	GroupName pulumi.StringPtrInput
	// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType pulumi.IntPtrInput
}

func (DeployGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployGroupState)(nil)).Elem()
}

type deployGroupArgs struct {
	// The ID of the application that you want to deploy.
	AppId string `pulumi:"appId"`
	// The name of the instance group that you want to create.
	GroupName string `pulumi:"groupName"`
}

// The set of arguments for constructing a DeployGroup resource.
type DeployGroupArgs struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringInput
	// The name of the instance group that you want to create.
	GroupName pulumi.StringInput
}

func (DeployGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployGroupArgs)(nil)).Elem()
}

type DeployGroupInput interface {
	pulumi.Input

	ToDeployGroupOutput() DeployGroupOutput
	ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput
}

func (*DeployGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployGroup)(nil)).Elem()
}

func (i *DeployGroup) ToDeployGroupOutput() DeployGroupOutput {
	return i.ToDeployGroupOutputWithContext(context.Background())
}

func (i *DeployGroup) ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployGroupOutput)
}

// DeployGroupArrayInput is an input type that accepts DeployGroupArray and DeployGroupArrayOutput values.
// You can construct a concrete instance of `DeployGroupArrayInput` via:
//
//          DeployGroupArray{ DeployGroupArgs{...} }
type DeployGroupArrayInput interface {
	pulumi.Input

	ToDeployGroupArrayOutput() DeployGroupArrayOutput
	ToDeployGroupArrayOutputWithContext(context.Context) DeployGroupArrayOutput
}

type DeployGroupArray []DeployGroupInput

func (DeployGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployGroup)(nil)).Elem()
}

func (i DeployGroupArray) ToDeployGroupArrayOutput() DeployGroupArrayOutput {
	return i.ToDeployGroupArrayOutputWithContext(context.Background())
}

func (i DeployGroupArray) ToDeployGroupArrayOutputWithContext(ctx context.Context) DeployGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployGroupArrayOutput)
}

// DeployGroupMapInput is an input type that accepts DeployGroupMap and DeployGroupMapOutput values.
// You can construct a concrete instance of `DeployGroupMapInput` via:
//
//          DeployGroupMap{ "key": DeployGroupArgs{...} }
type DeployGroupMapInput interface {
	pulumi.Input

	ToDeployGroupMapOutput() DeployGroupMapOutput
	ToDeployGroupMapOutputWithContext(context.Context) DeployGroupMapOutput
}

type DeployGroupMap map[string]DeployGroupInput

func (DeployGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployGroup)(nil)).Elem()
}

func (i DeployGroupMap) ToDeployGroupMapOutput() DeployGroupMapOutput {
	return i.ToDeployGroupMapOutputWithContext(context.Background())
}

func (i DeployGroupMap) ToDeployGroupMapOutputWithContext(ctx context.Context) DeployGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployGroupMapOutput)
}

type DeployGroupOutput struct{ *pulumi.OutputState }

func (DeployGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployGroup)(nil)).Elem()
}

func (o DeployGroupOutput) ToDeployGroupOutput() DeployGroupOutput {
	return o
}

func (o DeployGroupOutput) ToDeployGroupOutputWithContext(ctx context.Context) DeployGroupOutput {
	return o
}

// The ID of the application that you want to deploy.
func (o DeployGroupOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeployGroup) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The name of the instance group that you want to create.
func (o DeployGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DeployGroup) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
func (o DeployGroupOutput) GroupType() pulumi.IntOutput {
	return o.ApplyT(func(v *DeployGroup) pulumi.IntOutput { return v.GroupType }).(pulumi.IntOutput)
}

type DeployGroupArrayOutput struct{ *pulumi.OutputState }

func (DeployGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployGroup)(nil)).Elem()
}

func (o DeployGroupArrayOutput) ToDeployGroupArrayOutput() DeployGroupArrayOutput {
	return o
}

func (o DeployGroupArrayOutput) ToDeployGroupArrayOutputWithContext(ctx context.Context) DeployGroupArrayOutput {
	return o
}

func (o DeployGroupArrayOutput) Index(i pulumi.IntInput) DeployGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeployGroup {
		return vs[0].([]*DeployGroup)[vs[1].(int)]
	}).(DeployGroupOutput)
}

type DeployGroupMapOutput struct{ *pulumi.OutputState }

func (DeployGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployGroup)(nil)).Elem()
}

func (o DeployGroupMapOutput) ToDeployGroupMapOutput() DeployGroupMapOutput {
	return o
}

func (o DeployGroupMapOutput) ToDeployGroupMapOutputWithContext(ctx context.Context) DeployGroupMapOutput {
	return o
}

func (o DeployGroupMapOutput) MapIndex(k pulumi.StringInput) DeployGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeployGroup {
		return vs[0].(map[string]*DeployGroup)[vs[1].(string)]
	}).(DeployGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployGroupInput)(nil)).Elem(), &DeployGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployGroupArrayInput)(nil)).Elem(), DeployGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployGroupMapInput)(nil)).Elem(), DeployGroupMap{})
	pulumi.RegisterOutputType(DeployGroupOutput{})
	pulumi.RegisterOutputType(DeployGroupArrayOutput{})
	pulumi.RegisterOutputType(DeployGroupMapOutput{})
}
