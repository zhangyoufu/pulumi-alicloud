// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OOS Template resource. For information about Alicloud OOS Template and how to use it, see [What is Resource Alicloud OOS Template](https://www.alibabacloud.com/help/doc-detail/120761.htm).
//
// > **NOTE:** Available in 1.92.0+.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := oos.NewTemplate(ctx, "example", &oos.TemplateArgs{
//				Content: pulumi.String(`  {
//	    "FormatVersion": "OOS-2019-06-01",
//	    "Description": "Update Describe instances of given status",
//	    "Parameters":{
//	      "Status":{
//	        "Type": "String",
//	        "Description": "(Required) The status of the Ecs instance."
//	      }
//	    },
//	    "Tasks": [
//	      {
//	        "Properties" :{
//	          "Parameters":{
//	            "Status": "{{ Status }}"
//	          },
//	          "API": "DescribeInstances",
//	          "Service": "Ecs"
//	        },
//	        "Name": "foo",
//	        "Action": "ACS::ExecuteApi"
//	      }]
//	  }
//
// `),
//
//				TemplateName: pulumi.String("test-name"),
//				VersionName:  pulumi.String("test"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("acceptance Test"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// OOS Template can be imported using the id or template_name, e.g.
//
// ```sh
// $ pulumi import alicloud:oos/template:Template example template_name
// ```
type Template struct {
	pulumi.CustomResourceState

	// When deleting a template, whether to delete its related executions. Default to `false`.
	AutoDeleteExecutions pulumi.BoolPtrOutput `pulumi:"autoDeleteExecutions"`
	// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
	Content pulumi.StringOutput `pulumi:"content"`
	// The creator of the template.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The time when the template is created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the template.
	Description pulumi.StringOutput `pulumi:"description"`
	// Is it triggered successfully.
	HasTrigger pulumi.BoolOutput `pulumi:"hasTrigger"`
	// The ID of resource group which the template belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
	ShareType pulumi.StringOutput `pulumi:"shareType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
	TemplateFormat pulumi.StringOutput `pulumi:"templateFormat"`
	// The id of OOS Template.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
	TemplateType pulumi.StringOutput `pulumi:"templateType"`
	// The version of OOS Template.
	TemplateVersion pulumi.StringOutput `pulumi:"templateVersion"`
	// The user who updated the template.
	UpdatedBy pulumi.StringOutput `pulumi:"updatedBy"`
	// The time when the template was updated.
	UpdatedDate pulumi.StringOutput `pulumi:"updatedDate"`
	// The name of template version.
	VersionName pulumi.StringPtrOutput `pulumi:"versionName"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("alicloud:oos/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("alicloud:oos/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// When deleting a template, whether to delete its related executions. Default to `false`.
	AutoDeleteExecutions *bool `pulumi:"autoDeleteExecutions"`
	// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
	Content *string `pulumi:"content"`
	// The creator of the template.
	CreatedBy *string `pulumi:"createdBy"`
	// The time when the template is created.
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the template.
	Description *string `pulumi:"description"`
	// Is it triggered successfully.
	HasTrigger *bool `pulumi:"hasTrigger"`
	// The ID of resource group which the template belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
	ShareType *string `pulumi:"shareType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
	TemplateFormat *string `pulumi:"templateFormat"`
	// The id of OOS Template.
	TemplateId *string `pulumi:"templateId"`
	// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
	TemplateName *string `pulumi:"templateName"`
	// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
	TemplateType *string `pulumi:"templateType"`
	// The version of OOS Template.
	TemplateVersion *string `pulumi:"templateVersion"`
	// The user who updated the template.
	UpdatedBy *string `pulumi:"updatedBy"`
	// The time when the template was updated.
	UpdatedDate *string `pulumi:"updatedDate"`
	// The name of template version.
	VersionName *string `pulumi:"versionName"`
}

type TemplateState struct {
	// When deleting a template, whether to delete its related executions. Default to `false`.
	AutoDeleteExecutions pulumi.BoolPtrInput
	// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
	Content pulumi.StringPtrInput
	// The creator of the template.
	CreatedBy pulumi.StringPtrInput
	// The time when the template is created.
	CreatedDate pulumi.StringPtrInput
	// The description of the template.
	Description pulumi.StringPtrInput
	// Is it triggered successfully.
	HasTrigger pulumi.BoolPtrInput
	// The ID of resource group which the template belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
	ShareType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
	TemplateFormat pulumi.StringPtrInput
	// The id of OOS Template.
	TemplateId pulumi.StringPtrInput
	// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
	TemplateName pulumi.StringPtrInput
	// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
	TemplateType pulumi.StringPtrInput
	// The version of OOS Template.
	TemplateVersion pulumi.StringPtrInput
	// The user who updated the template.
	UpdatedBy pulumi.StringPtrInput
	// The time when the template was updated.
	UpdatedDate pulumi.StringPtrInput
	// The name of template version.
	VersionName pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// When deleting a template, whether to delete its related executions. Default to `false`.
	AutoDeleteExecutions *bool `pulumi:"autoDeleteExecutions"`
	// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
	Content string `pulumi:"content"`
	// The ID of resource group which the template belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
	TemplateName string `pulumi:"templateName"`
	// The name of template version.
	VersionName *string `pulumi:"versionName"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// When deleting a template, whether to delete its related executions. Default to `false`.
	AutoDeleteExecutions pulumi.BoolPtrInput
	// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
	Content pulumi.StringInput
	// The ID of resource group which the template belongs.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
	TemplateName pulumi.StringInput
	// The name of template version.
	VersionName pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//	TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//	TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// When deleting a template, whether to delete its related executions. Default to `false`.
func (o TemplateOutput) AutoDeleteExecutions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.BoolPtrOutput { return v.AutoDeleteExecutions }).(pulumi.BoolPtrOutput)
}

// The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.
func (o TemplateOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The creator of the template.
func (o TemplateOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The time when the template is created.
func (o TemplateOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// The description of the template.
func (o TemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Is it triggered successfully.
func (o TemplateOutput) HasTrigger() pulumi.BoolOutput {
	return o.ApplyT(func(v *Template) pulumi.BoolOutput { return v.HasTrigger }).(pulumi.BoolOutput)
}

// The ID of resource group which the template belongs.
func (o TemplateOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.
func (o TemplateOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.ShareType }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o TemplateOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Template) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The format of the template. The format can be JSON or YAML. The system automatically identifies the format.
func (o TemplateOutput) TemplateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateFormat }).(pulumi.StringOutput)
}

// The id of OOS Template.
func (o TemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with `ALIYUN`, `ACS`, `ALIBABA`, or `ALICLOUD`.
func (o TemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// The type of OOS Template. `Automation` means the implementation of Alibaba Cloud API template, `Package` means represents a template for installing software.
func (o TemplateOutput) TemplateType() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateType }).(pulumi.StringOutput)
}

// The version of OOS Template.
func (o TemplateOutput) TemplateVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateVersion }).(pulumi.StringOutput)
}

// The user who updated the template.
func (o TemplateOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.UpdatedBy }).(pulumi.StringOutput)
}

// The time when the template was updated.
func (o TemplateOutput) UpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.UpdatedDate }).(pulumi.StringOutput)
}

// The name of template version.
func (o TemplateOutput) VersionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.VersionName }).(pulumi.StringPtrOutput)
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Template {
		return vs[0].([]*Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Template {
		return vs[0].(map[string]*Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateArrayInput)(nil)).Elem(), TemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateMapInput)(nil)).Elem(), TemplateMap{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
