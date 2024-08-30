// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Log Service manages all the ECS instances whose logs need to be collected by using the Logtail client in the form of machine groups.
//
//	[Refer to details](https://www.alibabacloud.com/help/doc-detail/28966.htm)
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			example, err := log.NewProject(ctx, "example", &log.ProjectArgs{
//				ProjectName: pulumi.Sprintf("terraform-example-%v", _default.Result),
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = log.NewMachineGroup(ctx, "example", &log.MachineGroupArgs{
//				Project:      example.ProjectName,
//				Name:         pulumi.String("terraform-example"),
//				IdentifyType: pulumi.String("ip"),
//				Topic:        pulumi.String("terraform"),
//				IdentifyLists: pulumi.StringArray{
//					pulumi.String("10.0.0.1"),
//					pulumi.String("10.0.0.2"),
//				},
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
// ## Module Support
//
// You can use the existing sls-logtail module
// to create logtail config, machine group, install logtail on ECS instances and join instances into machine group one-click.
//
// ## Import
//
// Log machine group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:log/machineGroup:MachineGroup example tf-log:tf-machine-group
// ```
type MachineGroup struct {
	pulumi.CustomResourceState

	// The specific machine identification, which can be an IP address or user-defined identity.
	IdentifyLists pulumi.StringArrayOutput `pulumi:"identifyLists"`
	// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
	IdentifyType pulumi.StringPtrOutput `pulumi:"identifyType"`
	// The machine group name, which is unique in the same project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project name to the machine group belongs.
	Project pulumi.StringOutput `pulumi:"project"`
	// The topic of a machine group.
	Topic pulumi.StringPtrOutput `pulumi:"topic"`
}

// NewMachineGroup registers a new resource with the given unique name, arguments, and options.
func NewMachineGroup(ctx *pulumi.Context,
	name string, args *MachineGroupArgs, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentifyLists == nil {
		return nil, errors.New("invalid value for required argument 'IdentifyLists'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MachineGroup
	err := ctx.RegisterResource("alicloud:log/machineGroup:MachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineGroup gets an existing MachineGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineGroupState, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	var resource MachineGroup
	err := ctx.ReadResource("alicloud:log/machineGroup:MachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineGroup resources.
type machineGroupState struct {
	// The specific machine identification, which can be an IP address or user-defined identity.
	IdentifyLists []string `pulumi:"identifyLists"`
	// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
	IdentifyType *string `pulumi:"identifyType"`
	// The machine group name, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The project name to the machine group belongs.
	Project *string `pulumi:"project"`
	// The topic of a machine group.
	Topic *string `pulumi:"topic"`
}

type MachineGroupState struct {
	// The specific machine identification, which can be an IP address or user-defined identity.
	IdentifyLists pulumi.StringArrayInput
	// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
	IdentifyType pulumi.StringPtrInput
	// The machine group name, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The project name to the machine group belongs.
	Project pulumi.StringPtrInput
	// The topic of a machine group.
	Topic pulumi.StringPtrInput
}

func (MachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupState)(nil)).Elem()
}

type machineGroupArgs struct {
	// The specific machine identification, which can be an IP address or user-defined identity.
	IdentifyLists []string `pulumi:"identifyLists"`
	// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
	IdentifyType *string `pulumi:"identifyType"`
	// The machine group name, which is unique in the same project.
	Name *string `pulumi:"name"`
	// The project name to the machine group belongs.
	Project string `pulumi:"project"`
	// The topic of a machine group.
	Topic *string `pulumi:"topic"`
}

// The set of arguments for constructing a MachineGroup resource.
type MachineGroupArgs struct {
	// The specific machine identification, which can be an IP address or user-defined identity.
	IdentifyLists pulumi.StringArrayInput
	// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
	IdentifyType pulumi.StringPtrInput
	// The machine group name, which is unique in the same project.
	Name pulumi.StringPtrInput
	// The project name to the machine group belongs.
	Project pulumi.StringInput
	// The topic of a machine group.
	Topic pulumi.StringPtrInput
}

func (MachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupArgs)(nil)).Elem()
}

type MachineGroupInput interface {
	pulumi.Input

	ToMachineGroupOutput() MachineGroupOutput
	ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput
}

func (*MachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (i *MachineGroup) ToMachineGroupOutput() MachineGroupOutput {
	return i.ToMachineGroupOutputWithContext(context.Background())
}

func (i *MachineGroup) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupOutput)
}

// MachineGroupArrayInput is an input type that accepts MachineGroupArray and MachineGroupArrayOutput values.
// You can construct a concrete instance of `MachineGroupArrayInput` via:
//
//	MachineGroupArray{ MachineGroupArgs{...} }
type MachineGroupArrayInput interface {
	pulumi.Input

	ToMachineGroupArrayOutput() MachineGroupArrayOutput
	ToMachineGroupArrayOutputWithContext(context.Context) MachineGroupArrayOutput
}

type MachineGroupArray []MachineGroupInput

func (MachineGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineGroup)(nil)).Elem()
}

func (i MachineGroupArray) ToMachineGroupArrayOutput() MachineGroupArrayOutput {
	return i.ToMachineGroupArrayOutputWithContext(context.Background())
}

func (i MachineGroupArray) ToMachineGroupArrayOutputWithContext(ctx context.Context) MachineGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupArrayOutput)
}

// MachineGroupMapInput is an input type that accepts MachineGroupMap and MachineGroupMapOutput values.
// You can construct a concrete instance of `MachineGroupMapInput` via:
//
//	MachineGroupMap{ "key": MachineGroupArgs{...} }
type MachineGroupMapInput interface {
	pulumi.Input

	ToMachineGroupMapOutput() MachineGroupMapOutput
	ToMachineGroupMapOutputWithContext(context.Context) MachineGroupMapOutput
}

type MachineGroupMap map[string]MachineGroupInput

func (MachineGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineGroup)(nil)).Elem()
}

func (i MachineGroupMap) ToMachineGroupMapOutput() MachineGroupMapOutput {
	return i.ToMachineGroupMapOutputWithContext(context.Background())
}

func (i MachineGroupMap) ToMachineGroupMapOutputWithContext(ctx context.Context) MachineGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupMapOutput)
}

type MachineGroupOutput struct{ *pulumi.OutputState }

func (MachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (o MachineGroupOutput) ToMachineGroupOutput() MachineGroupOutput {
	return o
}

func (o MachineGroupOutput) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return o
}

// The specific machine identification, which can be an IP address or user-defined identity.
func (o MachineGroupOutput) IdentifyLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringArrayOutput { return v.IdentifyLists }).(pulumi.StringArrayOutput)
}

// The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".
func (o MachineGroupOutput) IdentifyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.IdentifyType }).(pulumi.StringPtrOutput)
}

// The machine group name, which is unique in the same project.
func (o MachineGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project name to the machine group belongs.
func (o MachineGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The topic of a machine group.
func (o MachineGroupOutput) Topic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.Topic }).(pulumi.StringPtrOutput)
}

type MachineGroupArrayOutput struct{ *pulumi.OutputState }

func (MachineGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MachineGroup)(nil)).Elem()
}

func (o MachineGroupArrayOutput) ToMachineGroupArrayOutput() MachineGroupArrayOutput {
	return o
}

func (o MachineGroupArrayOutput) ToMachineGroupArrayOutputWithContext(ctx context.Context) MachineGroupArrayOutput {
	return o
}

func (o MachineGroupArrayOutput) Index(i pulumi.IntInput) MachineGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MachineGroup {
		return vs[0].([]*MachineGroup)[vs[1].(int)]
	}).(MachineGroupOutput)
}

type MachineGroupMapOutput struct{ *pulumi.OutputState }

func (MachineGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MachineGroup)(nil)).Elem()
}

func (o MachineGroupMapOutput) ToMachineGroupMapOutput() MachineGroupMapOutput {
	return o
}

func (o MachineGroupMapOutput) ToMachineGroupMapOutputWithContext(ctx context.Context) MachineGroupMapOutput {
	return o
}

func (o MachineGroupMapOutput) MapIndex(k pulumi.StringInput) MachineGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MachineGroup {
		return vs[0].(map[string]*MachineGroup)[vs[1].(string)]
	}).(MachineGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupInput)(nil)).Elem(), &MachineGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupArrayInput)(nil)).Elem(), MachineGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineGroupMapInput)(nil)).Elem(), MachineGroupMap{})
	pulumi.RegisterOutputType(MachineGroupOutput{})
	pulumi.RegisterOutputType(MachineGroupArrayOutput{})
	pulumi.RegisterOutputType(MachineGroupMapOutput{})
}
