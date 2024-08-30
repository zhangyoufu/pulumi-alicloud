// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package brain

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Brain Industrial Pid Project resource.
//
// > **NOTE:** Available since v1.113.0.
//
// > **DEPRECATED:**  This resource has been deprecated from version `1.222.0`.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/brain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := brain.NewIndustrialPidProject(ctx, "example", &brain.IndustrialPidProjectArgs{
//				PidOrganizationId: pulumi.String("3e74e684-cbb5-xxxx"),
//				PidProjectName:    pulumi.String("tf-testAcc"),
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
// Brain Industrial Pid Project can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:brain/industrialPidProject:IndustrialPidProject example <id>
// ```
type IndustrialPidProject struct {
	pulumi.CustomResourceState

	// The ID of Pid Organization.
	PidOrganizationId pulumi.StringOutput `pulumi:"pidOrganizationId"`
	// The description of Pid Project.
	PidProjectDesc pulumi.StringPtrOutput `pulumi:"pidProjectDesc"`
	// The name of Pid Project.
	PidProjectName pulumi.StringOutput `pulumi:"pidProjectName"`
}

// NewIndustrialPidProject registers a new resource with the given unique name, arguments, and options.
func NewIndustrialPidProject(ctx *pulumi.Context,
	name string, args *IndustrialPidProjectArgs, opts ...pulumi.ResourceOption) (*IndustrialPidProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PidOrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'PidOrganizationId'")
	}
	if args.PidProjectName == nil {
		return nil, errors.New("invalid value for required argument 'PidProjectName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IndustrialPidProject
	err := ctx.RegisterResource("alicloud:brain/industrialPidProject:IndustrialPidProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndustrialPidProject gets an existing IndustrialPidProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndustrialPidProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndustrialPidProjectState, opts ...pulumi.ResourceOption) (*IndustrialPidProject, error) {
	var resource IndustrialPidProject
	err := ctx.ReadResource("alicloud:brain/industrialPidProject:IndustrialPidProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IndustrialPidProject resources.
type industrialPidProjectState struct {
	// The ID of Pid Organization.
	PidOrganizationId *string `pulumi:"pidOrganizationId"`
	// The description of Pid Project.
	PidProjectDesc *string `pulumi:"pidProjectDesc"`
	// The name of Pid Project.
	PidProjectName *string `pulumi:"pidProjectName"`
}

type IndustrialPidProjectState struct {
	// The ID of Pid Organization.
	PidOrganizationId pulumi.StringPtrInput
	// The description of Pid Project.
	PidProjectDesc pulumi.StringPtrInput
	// The name of Pid Project.
	PidProjectName pulumi.StringPtrInput
}

func (IndustrialPidProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*industrialPidProjectState)(nil)).Elem()
}

type industrialPidProjectArgs struct {
	// The ID of Pid Organization.
	PidOrganizationId string `pulumi:"pidOrganizationId"`
	// The description of Pid Project.
	PidProjectDesc *string `pulumi:"pidProjectDesc"`
	// The name of Pid Project.
	PidProjectName string `pulumi:"pidProjectName"`
}

// The set of arguments for constructing a IndustrialPidProject resource.
type IndustrialPidProjectArgs struct {
	// The ID of Pid Organization.
	PidOrganizationId pulumi.StringInput
	// The description of Pid Project.
	PidProjectDesc pulumi.StringPtrInput
	// The name of Pid Project.
	PidProjectName pulumi.StringInput
}

func (IndustrialPidProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*industrialPidProjectArgs)(nil)).Elem()
}

type IndustrialPidProjectInput interface {
	pulumi.Input

	ToIndustrialPidProjectOutput() IndustrialPidProjectOutput
	ToIndustrialPidProjectOutputWithContext(ctx context.Context) IndustrialPidProjectOutput
}

func (*IndustrialPidProject) ElementType() reflect.Type {
	return reflect.TypeOf((**IndustrialPidProject)(nil)).Elem()
}

func (i *IndustrialPidProject) ToIndustrialPidProjectOutput() IndustrialPidProjectOutput {
	return i.ToIndustrialPidProjectOutputWithContext(context.Background())
}

func (i *IndustrialPidProject) ToIndustrialPidProjectOutputWithContext(ctx context.Context) IndustrialPidProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndustrialPidProjectOutput)
}

// IndustrialPidProjectArrayInput is an input type that accepts IndustrialPidProjectArray and IndustrialPidProjectArrayOutput values.
// You can construct a concrete instance of `IndustrialPidProjectArrayInput` via:
//
//	IndustrialPidProjectArray{ IndustrialPidProjectArgs{...} }
type IndustrialPidProjectArrayInput interface {
	pulumi.Input

	ToIndustrialPidProjectArrayOutput() IndustrialPidProjectArrayOutput
	ToIndustrialPidProjectArrayOutputWithContext(context.Context) IndustrialPidProjectArrayOutput
}

type IndustrialPidProjectArray []IndustrialPidProjectInput

func (IndustrialPidProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IndustrialPidProject)(nil)).Elem()
}

func (i IndustrialPidProjectArray) ToIndustrialPidProjectArrayOutput() IndustrialPidProjectArrayOutput {
	return i.ToIndustrialPidProjectArrayOutputWithContext(context.Background())
}

func (i IndustrialPidProjectArray) ToIndustrialPidProjectArrayOutputWithContext(ctx context.Context) IndustrialPidProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndustrialPidProjectArrayOutput)
}

// IndustrialPidProjectMapInput is an input type that accepts IndustrialPidProjectMap and IndustrialPidProjectMapOutput values.
// You can construct a concrete instance of `IndustrialPidProjectMapInput` via:
//
//	IndustrialPidProjectMap{ "key": IndustrialPidProjectArgs{...} }
type IndustrialPidProjectMapInput interface {
	pulumi.Input

	ToIndustrialPidProjectMapOutput() IndustrialPidProjectMapOutput
	ToIndustrialPidProjectMapOutputWithContext(context.Context) IndustrialPidProjectMapOutput
}

type IndustrialPidProjectMap map[string]IndustrialPidProjectInput

func (IndustrialPidProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IndustrialPidProject)(nil)).Elem()
}

func (i IndustrialPidProjectMap) ToIndustrialPidProjectMapOutput() IndustrialPidProjectMapOutput {
	return i.ToIndustrialPidProjectMapOutputWithContext(context.Background())
}

func (i IndustrialPidProjectMap) ToIndustrialPidProjectMapOutputWithContext(ctx context.Context) IndustrialPidProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndustrialPidProjectMapOutput)
}

type IndustrialPidProjectOutput struct{ *pulumi.OutputState }

func (IndustrialPidProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndustrialPidProject)(nil)).Elem()
}

func (o IndustrialPidProjectOutput) ToIndustrialPidProjectOutput() IndustrialPidProjectOutput {
	return o
}

func (o IndustrialPidProjectOutput) ToIndustrialPidProjectOutputWithContext(ctx context.Context) IndustrialPidProjectOutput {
	return o
}

// The ID of Pid Organization.
func (o IndustrialPidProjectOutput) PidOrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IndustrialPidProject) pulumi.StringOutput { return v.PidOrganizationId }).(pulumi.StringOutput)
}

// The description of Pid Project.
func (o IndustrialPidProjectOutput) PidProjectDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndustrialPidProject) pulumi.StringPtrOutput { return v.PidProjectDesc }).(pulumi.StringPtrOutput)
}

// The name of Pid Project.
func (o IndustrialPidProjectOutput) PidProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *IndustrialPidProject) pulumi.StringOutput { return v.PidProjectName }).(pulumi.StringOutput)
}

type IndustrialPidProjectArrayOutput struct{ *pulumi.OutputState }

func (IndustrialPidProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IndustrialPidProject)(nil)).Elem()
}

func (o IndustrialPidProjectArrayOutput) ToIndustrialPidProjectArrayOutput() IndustrialPidProjectArrayOutput {
	return o
}

func (o IndustrialPidProjectArrayOutput) ToIndustrialPidProjectArrayOutputWithContext(ctx context.Context) IndustrialPidProjectArrayOutput {
	return o
}

func (o IndustrialPidProjectArrayOutput) Index(i pulumi.IntInput) IndustrialPidProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IndustrialPidProject {
		return vs[0].([]*IndustrialPidProject)[vs[1].(int)]
	}).(IndustrialPidProjectOutput)
}

type IndustrialPidProjectMapOutput struct{ *pulumi.OutputState }

func (IndustrialPidProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IndustrialPidProject)(nil)).Elem()
}

func (o IndustrialPidProjectMapOutput) ToIndustrialPidProjectMapOutput() IndustrialPidProjectMapOutput {
	return o
}

func (o IndustrialPidProjectMapOutput) ToIndustrialPidProjectMapOutputWithContext(ctx context.Context) IndustrialPidProjectMapOutput {
	return o
}

func (o IndustrialPidProjectMapOutput) MapIndex(k pulumi.StringInput) IndustrialPidProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IndustrialPidProject {
		return vs[0].(map[string]*IndustrialPidProject)[vs[1].(string)]
	}).(IndustrialPidProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndustrialPidProjectInput)(nil)).Elem(), &IndustrialPidProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndustrialPidProjectArrayInput)(nil)).Elem(), IndustrialPidProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndustrialPidProjectMapInput)(nil)).Elem(), IndustrialPidProjectMap{})
	pulumi.RegisterOutputType(IndustrialPidProjectOutput{})
	pulumi.RegisterOutputType(IndustrialPidProjectArrayOutput{})
	pulumi.RegisterOutputType(IndustrialPidProjectMapOutput{})
}
