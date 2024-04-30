// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SLS Project resource.
//
// For information about SLS Project and how to use it, see [What is Project](https://www.alibabacloud.com/help/en/sls/developer-reference/api-createproject).
//
// > **NOTE:** Available since v1.9.5.
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
//			_, err = log.NewProject(ctx, "example", &log.ProjectArgs{
//				Name:        pulumi.String(fmt.Sprintf("terraform-example-%v", _default.Result)),
//				Description: pulumi.String("terraform-example"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
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
// # Project With Policy Usage
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
//			_, err = log.NewProject(ctx, "example_policy", &log.ProjectArgs{
//				Name:        pulumi.String(fmt.Sprintf("terraform-example-%v", _default.Result)),
//				Description: pulumi.String("terraform-example"),
//				Policy: pulumi.String(`{
//	  "Statement": [
//	    {
//	      "Action": [
//	        "log:PostLogStoreLogs"
//	      ],
//	      "Condition": {
//	        "StringNotLike": {
//	          "acs:SourceVpc": [
//	            "vpc-*"
//	          ]
//	        }
//	      },
//	      "Effect": "Deny",
//	      "Resource": "acs:log:*:*:project/tf-log/*"
//	    }
//	  ],
//	  "Version": "1"
//	}
//
// `),
//
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
// You can use the existing sls module
// to create SLS project, store and store index one-click, like ECS instances.
//
// ## Import
//
// SLS Project can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:log/project:Project example <id>
// ```
type Project struct {
	pulumi.CustomResourceState

	// CreateTime.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// Log project policy, used to set a policy for a project.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
	// - The project name must be globally unique.
	// - The name can contain only lowercase letters, digits, and hyphens (-).
	// - It must start and end with a lowercase letter or number.
	// - The value contains 3 to 63 characters.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tag.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	// CreateTime.
	CreateTime *string `pulumi:"createTime"`
	// Description.
	Description *string `pulumi:"description"`
	// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
	Name *string `pulumi:"name"`
	// Log project policy, used to set a policy for a project.
	Policy *string `pulumi:"policy"`
	// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
	// - The project name must be globally unique.
	// - The name can contain only lowercase letters, digits, and hyphens (-).
	// - It must start and end with a lowercase letter or number.
	// - The value contains 3 to 63 characters.
	ProjectName *string `pulumi:"projectName"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// Tag.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Tags map[string]interface{} `pulumi:"tags"`
}

type ProjectState struct {
	// CreateTime.
	CreateTime pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
	Name pulumi.StringPtrInput
	// Log project policy, used to set a policy for a project.
	Policy pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
	// - The project name must be globally unique.
	// - The name can contain only lowercase letters, digits, and hyphens (-).
	// - It must start and end with a lowercase letter or number.
	// - The value contains 3 to 63 characters.
	ProjectName pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// Tag.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Tags pulumi.MapInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
	Name *string `pulumi:"name"`
	// Log project policy, used to set a policy for a project.
	Policy *string `pulumi:"policy"`
	// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
	// - The project name must be globally unique.
	// - The name can contain only lowercase letters, digits, and hyphens (-).
	// - It must start and end with a lowercase letter or number.
	// - The value contains 3 to 63 characters.
	ProjectName *string `pulumi:"projectName"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Tag.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
	Name pulumi.StringPtrInput
	// Log project policy, used to set a policy for a project.
	Policy pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
	// - The project name must be globally unique.
	// - The name can contain only lowercase letters, digits, and hyphens (-).
	// - It must start and end with a lowercase letter or number.
	// - The value contains 3 to 63 characters.
	ProjectName pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Tag.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
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

// CreateTime.
func (o ProjectOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// . Field 'name' has been deprecated from provider version 1.223.0. New field 'project_name' instead.
//
// Deprecated: Field 'name' has been deprecated since provider version 1.223.0. New field 'project_name' instead.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Log project policy, used to set a policy for a project.
func (o ProjectOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// The name of the log project. It is the only in one Alicloud account. The project name is globally unique in Alibaba Cloud and cannot be modified after it is created. The naming rules are as follows:
// - The project name must be globally unique.
// - The name can contain only lowercase letters, digits, and hyphens (-).
// - It must start and end with a lowercase letter or number.
// - The value contains 3 to 63 characters.
func (o ProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o ProjectOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the resource.
func (o ProjectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tag.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o ProjectOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Project) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
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
