// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maxcompute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The project is the basic unit of operation in maxcompute. It is similar to the concept of Database or Schema in traditional databases, and sets the boundary for maxcompute multi-user isolation and access control. [Refer to details](https://www.alibabacloud.com/help/doc-detail/27818.html).
//
// ->**NOTE:** Available in 1.77.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/maxcompute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := maxcompute.NewProject(ctx, "example", &maxcompute.ProjectArgs{
//				OrderType:         pulumi.String("PayAsYouGo"),
//				ProjectName:       pulumi.String("tf_maxcompute_project"),
//				SpecificationType: pulumi.String("OdpsStandard"),
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
// MaxCompute project can be imported using the *name* or ID, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:maxcompute/project:Project example tf_maxcompute_project
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringOutput `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringOutput `pulumi:"specificationType"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrderType == nil {
		return nil, errors.New("invalid value for required argument 'OrderType'")
	}
	if args.SpecificationType == nil {
		return nil, errors.New("invalid value for required argument 'SpecificationType'")
	}
	var resource Project
	err := ctx.RegisterResource("alicloud:maxcompute/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("alicloud:maxcompute/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name *string `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType *string `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName *string `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType *string `pulumi:"specificationType"`
}

type ProjectState struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringPtrInput
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringPtrInput
	// The name of the maxcompute project.
	ProjectName pulumi.StringPtrInput
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name *string `pulumi:"name"`
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType string `pulumi:"orderType"`
	// The name of the maxcompute project.
	ProjectName *string `pulumi:"projectName"`
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType string `pulumi:"specificationType"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// It has been deprecated from provider version 1.110.0 and `projectName` instead.
	Name pulumi.StringPtrInput
	// The type of payment, only `PayAsYouGo` supported currently.
	OrderType pulumi.StringInput
	// The name of the maxcompute project.
	ProjectName pulumi.StringPtrInput
	// The type of resource Specification, only `OdpsStandard` supported currently.
	SpecificationType pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// It has been deprecated from provider version 1.110.0 and `projectName` instead.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of payment, only `PayAsYouGo` supported currently.
func (o ProjectOutput) OrderType() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.OrderType }).(pulumi.StringOutput)
}

// The name of the maxcompute project.
func (o ProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The type of resource Specification, only `OdpsStandard` supported currently.
func (o ProjectOutput) SpecificationType() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.SpecificationType }).(pulumi.StringOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
