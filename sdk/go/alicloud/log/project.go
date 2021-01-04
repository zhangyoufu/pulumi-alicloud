// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Log project can be imported using the id or name, e.g.
//
// ```sh
//  $ pulumi import alicloud:log/project:Project example tf-log
// ```
type Project struct {
	pulumi.CustomResourceState

	// Description of the log project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Log project tags.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("alicloud:log/project:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:log/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Description of the log project.
	Description *string `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
	// Log project tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ProjectState struct {
	// Description of the log project.
	Description pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
	// Log project tags.
	Tags pulumi.MapInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Description of the log project.
	Description *string `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
	// Log project tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Description of the log project.
	Description pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
	// Log project tags.
	Tags pulumi.MapInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (Project) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil)).Elem()
}

func (i Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct {
	*pulumi.OutputState
}

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectOutput)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
