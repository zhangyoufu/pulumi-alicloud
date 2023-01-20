// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Architect Design Tools Application resource.
//
// For information about Cloud Architect Design Tools Application and how to use it, see [What is Application](https://help.aliyun.com/document_detail/428263.html).
//
// > **NOTE:** Available in v1.192.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/bp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bp.NewStudioApplication(ctx, "default", &bp.StudioApplicationArgs{
//				ApplicationName: pulumi.String("example_value"),
//				AreaId:          pulumi.String("example_value"),
//				Configuration: pulumi.AnyMap{
//					"enableMonitor": pulumi.Any("1"),
//				},
//				Instances: bp.StudioApplicationInstanceArray{
//					&bp.StudioApplicationInstanceArgs{
//						Id:       pulumi.String("example_value"),
//						NodeName: pulumi.String("example_value"),
//						NodeType: pulumi.String("ecs"),
//					},
//				},
//				ResourceGroupId: pulumi.String("example_value"),
//				TemplateId:      pulumi.String("example_value"),
//				Variables: pulumi.AnyMap{
//					"test": pulumi.Any("1"),
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
// ## Import
//
// Cloud Architect Design Tools Application can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:bp/studioApplication:StudioApplication example <id>
//
// ```
type StudioApplication struct {
	pulumi.CustomResourceState

	// The name of the application.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The id of the area.
	AreaId pulumi.StringPtrOutput `pulumi:"areaId"`
	// The configuration of the application.
	Configuration pulumi.MapOutput `pulumi:"configuration"`
	// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
	Instances StudioApplicationInstanceArrayOutput `pulumi:"instances"`
	// The id of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the Application.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the template.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// The variables of the application.
	Variables pulumi.MapOutput `pulumi:"variables"`
}

// NewStudioApplication registers a new resource with the given unique name, arguments, and options.
func NewStudioApplication(ctx *pulumi.Context,
	name string, args *StudioApplicationArgs, opts ...pulumi.ResourceOption) (*StudioApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	var resource StudioApplication
	err := ctx.RegisterResource("alicloud:bp/studioApplication:StudioApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudioApplication gets an existing StudioApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudioApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudioApplicationState, opts ...pulumi.ResourceOption) (*StudioApplication, error) {
	var resource StudioApplication
	err := ctx.ReadResource("alicloud:bp/studioApplication:StudioApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StudioApplication resources.
type studioApplicationState struct {
	// The name of the application.
	ApplicationName *string `pulumi:"applicationName"`
	// The id of the area.
	AreaId *string `pulumi:"areaId"`
	// The configuration of the application.
	Configuration map[string]interface{} `pulumi:"configuration"`
	// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
	Instances []StudioApplicationInstance `pulumi:"instances"`
	// The id of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the Application.
	Status *string `pulumi:"status"`
	// The id of the template.
	TemplateId *string `pulumi:"templateId"`
	// The variables of the application.
	Variables map[string]interface{} `pulumi:"variables"`
}

type StudioApplicationState struct {
	// The name of the application.
	ApplicationName pulumi.StringPtrInput
	// The id of the area.
	AreaId pulumi.StringPtrInput
	// The configuration of the application.
	Configuration pulumi.MapInput
	// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
	Instances StudioApplicationInstanceArrayInput
	// The id of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the Application.
	Status pulumi.StringPtrInput
	// The id of the template.
	TemplateId pulumi.StringPtrInput
	// The variables of the application.
	Variables pulumi.MapInput
}

func (StudioApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*studioApplicationState)(nil)).Elem()
}

type studioApplicationArgs struct {
	// The name of the application.
	ApplicationName string `pulumi:"applicationName"`
	// The id of the area.
	AreaId *string `pulumi:"areaId"`
	// The configuration of the application.
	Configuration map[string]interface{} `pulumi:"configuration"`
	// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
	Instances []StudioApplicationInstance `pulumi:"instances"`
	// The id of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The id of the template.
	TemplateId string `pulumi:"templateId"`
	// The variables of the application.
	Variables map[string]interface{} `pulumi:"variables"`
}

// The set of arguments for constructing a StudioApplication resource.
type StudioApplicationArgs struct {
	// The name of the application.
	ApplicationName pulumi.StringInput
	// The id of the area.
	AreaId pulumi.StringPtrInput
	// The configuration of the application.
	Configuration pulumi.MapInput
	// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
	Instances StudioApplicationInstanceArrayInput
	// The id of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The id of the template.
	TemplateId pulumi.StringInput
	// The variables of the application.
	Variables pulumi.MapInput
}

func (StudioApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studioApplicationArgs)(nil)).Elem()
}

type StudioApplicationInput interface {
	pulumi.Input

	ToStudioApplicationOutput() StudioApplicationOutput
	ToStudioApplicationOutputWithContext(ctx context.Context) StudioApplicationOutput
}

func (*StudioApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioApplication)(nil)).Elem()
}

func (i *StudioApplication) ToStudioApplicationOutput() StudioApplicationOutput {
	return i.ToStudioApplicationOutputWithContext(context.Background())
}

func (i *StudioApplication) ToStudioApplicationOutputWithContext(ctx context.Context) StudioApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioApplicationOutput)
}

// StudioApplicationArrayInput is an input type that accepts StudioApplicationArray and StudioApplicationArrayOutput values.
// You can construct a concrete instance of `StudioApplicationArrayInput` via:
//
//	StudioApplicationArray{ StudioApplicationArgs{...} }
type StudioApplicationArrayInput interface {
	pulumi.Input

	ToStudioApplicationArrayOutput() StudioApplicationArrayOutput
	ToStudioApplicationArrayOutputWithContext(context.Context) StudioApplicationArrayOutput
}

type StudioApplicationArray []StudioApplicationInput

func (StudioApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StudioApplication)(nil)).Elem()
}

func (i StudioApplicationArray) ToStudioApplicationArrayOutput() StudioApplicationArrayOutput {
	return i.ToStudioApplicationArrayOutputWithContext(context.Background())
}

func (i StudioApplicationArray) ToStudioApplicationArrayOutputWithContext(ctx context.Context) StudioApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioApplicationArrayOutput)
}

// StudioApplicationMapInput is an input type that accepts StudioApplicationMap and StudioApplicationMapOutput values.
// You can construct a concrete instance of `StudioApplicationMapInput` via:
//
//	StudioApplicationMap{ "key": StudioApplicationArgs{...} }
type StudioApplicationMapInput interface {
	pulumi.Input

	ToStudioApplicationMapOutput() StudioApplicationMapOutput
	ToStudioApplicationMapOutputWithContext(context.Context) StudioApplicationMapOutput
}

type StudioApplicationMap map[string]StudioApplicationInput

func (StudioApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StudioApplication)(nil)).Elem()
}

func (i StudioApplicationMap) ToStudioApplicationMapOutput() StudioApplicationMapOutput {
	return i.ToStudioApplicationMapOutputWithContext(context.Background())
}

func (i StudioApplicationMap) ToStudioApplicationMapOutputWithContext(ctx context.Context) StudioApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioApplicationMapOutput)
}

type StudioApplicationOutput struct{ *pulumi.OutputState }

func (StudioApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioApplication)(nil)).Elem()
}

func (o StudioApplicationOutput) ToStudioApplicationOutput() StudioApplicationOutput {
	return o
}

func (o StudioApplicationOutput) ToStudioApplicationOutputWithContext(ctx context.Context) StudioApplicationOutput {
	return o
}

// The name of the application.
func (o StudioApplicationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// The id of the area.
func (o StudioApplicationOutput) AreaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.StringPtrOutput { return v.AreaId }).(pulumi.StringPtrOutput)
}

// The configuration of the application.
func (o StudioApplicationOutput) Configuration() pulumi.MapOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.MapOutput { return v.Configuration }).(pulumi.MapOutput)
}

// The instance list. Support the creation of instances in the existing vpc under the application. See the following `Block instances`.
func (o StudioApplicationOutput) Instances() StudioApplicationInstanceArrayOutput {
	return o.ApplyT(func(v *StudioApplication) StudioApplicationInstanceArrayOutput { return v.Instances }).(StudioApplicationInstanceArrayOutput)
}

// The id of the resource group.
func (o StudioApplicationOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the Application.
func (o StudioApplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the template.
func (o StudioApplicationOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// The variables of the application.
func (o StudioApplicationOutput) Variables() pulumi.MapOutput {
	return o.ApplyT(func(v *StudioApplication) pulumi.MapOutput { return v.Variables }).(pulumi.MapOutput)
}

type StudioApplicationArrayOutput struct{ *pulumi.OutputState }

func (StudioApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StudioApplication)(nil)).Elem()
}

func (o StudioApplicationArrayOutput) ToStudioApplicationArrayOutput() StudioApplicationArrayOutput {
	return o
}

func (o StudioApplicationArrayOutput) ToStudioApplicationArrayOutputWithContext(ctx context.Context) StudioApplicationArrayOutput {
	return o
}

func (o StudioApplicationArrayOutput) Index(i pulumi.IntInput) StudioApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StudioApplication {
		return vs[0].([]*StudioApplication)[vs[1].(int)]
	}).(StudioApplicationOutput)
}

type StudioApplicationMapOutput struct{ *pulumi.OutputState }

func (StudioApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StudioApplication)(nil)).Elem()
}

func (o StudioApplicationMapOutput) ToStudioApplicationMapOutput() StudioApplicationMapOutput {
	return o
}

func (o StudioApplicationMapOutput) ToStudioApplicationMapOutputWithContext(ctx context.Context) StudioApplicationMapOutput {
	return o
}

func (o StudioApplicationMapOutput) MapIndex(k pulumi.StringInput) StudioApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StudioApplication {
		return vs[0].(map[string]*StudioApplication)[vs[1].(string)]
	}).(StudioApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioApplicationInput)(nil)).Elem(), &StudioApplication{})
	pulumi.RegisterInputType(reflect.TypeOf((*StudioApplicationArrayInput)(nil)).Elem(), StudioApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StudioApplicationMapInput)(nil)).Elem(), StudioApplicationMap{})
	pulumi.RegisterOutputType(StudioApplicationOutput{})
	pulumi.RegisterOutputType(StudioApplicationArrayOutput{})
	pulumi.RegisterOutputType(StudioApplicationMapOutput{})
}
