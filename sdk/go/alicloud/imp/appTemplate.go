// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Apsara Agile Live (IMP) App Template resource.
//
// For information about Apsara Agile Live (IMP) App Template and how to use it, see [What is App Template](https://help.aliyun.com/document_detail/270121.html).
//
// > **NOTE:** Available in v1.137.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/imp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imp.NewAppTemplate(ctx, "example", &imp.AppTemplateArgs{
//				AppTemplateName: pulumi.String("example_value"),
//				ComponentLists: pulumi.StringArray{
//					pulumi.String("component.live"),
//					pulumi.String("component.liveRecord"),
//				},
//				IntegrationMode: pulumi.String("paasSDK"),
//				Scene:           pulumi.String("business"),
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
// Apsara Agile Live (IMP) App Template can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:imp/appTemplate:AppTemplate example <id>
// ```
type AppTemplate struct {
	pulumi.CustomResourceState

	// The name of the resource.
	AppTemplateName pulumi.StringOutput `pulumi:"appTemplateName"`
	// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
	ComponentLists pulumi.StringArrayOutput `pulumi:"componentLists"`
	// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
	ConfigLists AppTemplateConfigListArrayOutput `pulumi:"configLists"`
	// Integration mode. Valid values:
	// * paasSDK: Integrated SDK.
	// * standardRoom: Model Room.
	IntegrationMode pulumi.StringPtrOutput `pulumi:"integrationMode"`
	// Application Template scenario. Valid values: ["business", "classroom"].
	Scene pulumi.StringPtrOutput `pulumi:"scene"`
	// Application template usage status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAppTemplate registers a new resource with the given unique name, arguments, and options.
func NewAppTemplate(ctx *pulumi.Context,
	name string, args *AppTemplateArgs, opts ...pulumi.ResourceOption) (*AppTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'AppTemplateName'")
	}
	if args.ComponentLists == nil {
		return nil, errors.New("invalid value for required argument 'ComponentLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppTemplate
	err := ctx.RegisterResource("alicloud:imp/appTemplate:AppTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppTemplate gets an existing AppTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppTemplateState, opts ...pulumi.ResourceOption) (*AppTemplate, error) {
	var resource AppTemplate
	err := ctx.ReadResource("alicloud:imp/appTemplate:AppTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppTemplate resources.
type appTemplateState struct {
	// The name of the resource.
	AppTemplateName *string `pulumi:"appTemplateName"`
	// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
	ComponentLists []string `pulumi:"componentLists"`
	// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
	ConfigLists []AppTemplateConfigList `pulumi:"configLists"`
	// Integration mode. Valid values:
	// * paasSDK: Integrated SDK.
	// * standardRoom: Model Room.
	IntegrationMode *string `pulumi:"integrationMode"`
	// Application Template scenario. Valid values: ["business", "classroom"].
	Scene *string `pulumi:"scene"`
	// Application template usage status.
	Status *string `pulumi:"status"`
}

type AppTemplateState struct {
	// The name of the resource.
	AppTemplateName pulumi.StringPtrInput
	// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
	ComponentLists pulumi.StringArrayInput
	// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
	ConfigLists AppTemplateConfigListArrayInput
	// Integration mode. Valid values:
	// * paasSDK: Integrated SDK.
	// * standardRoom: Model Room.
	IntegrationMode pulumi.StringPtrInput
	// Application Template scenario. Valid values: ["business", "classroom"].
	Scene pulumi.StringPtrInput
	// Application template usage status.
	Status pulumi.StringPtrInput
}

func (AppTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*appTemplateState)(nil)).Elem()
}

type appTemplateArgs struct {
	// The name of the resource.
	AppTemplateName string `pulumi:"appTemplateName"`
	// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
	ComponentLists []string `pulumi:"componentLists"`
	// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
	ConfigLists []AppTemplateConfigList `pulumi:"configLists"`
	// Integration mode. Valid values:
	// * paasSDK: Integrated SDK.
	// * standardRoom: Model Room.
	IntegrationMode *string `pulumi:"integrationMode"`
	// Application Template scenario. Valid values: ["business", "classroom"].
	Scene *string `pulumi:"scene"`
}

// The set of arguments for constructing a AppTemplate resource.
type AppTemplateArgs struct {
	// The name of the resource.
	AppTemplateName pulumi.StringInput
	// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
	ComponentLists pulumi.StringArrayInput
	// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
	ConfigLists AppTemplateConfigListArrayInput
	// Integration mode. Valid values:
	// * paasSDK: Integrated SDK.
	// * standardRoom: Model Room.
	IntegrationMode pulumi.StringPtrInput
	// Application Template scenario. Valid values: ["business", "classroom"].
	Scene pulumi.StringPtrInput
}

func (AppTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appTemplateArgs)(nil)).Elem()
}

type AppTemplateInput interface {
	pulumi.Input

	ToAppTemplateOutput() AppTemplateOutput
	ToAppTemplateOutputWithContext(ctx context.Context) AppTemplateOutput
}

func (*AppTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**AppTemplate)(nil)).Elem()
}

func (i *AppTemplate) ToAppTemplateOutput() AppTemplateOutput {
	return i.ToAppTemplateOutputWithContext(context.Background())
}

func (i *AppTemplate) ToAppTemplateOutputWithContext(ctx context.Context) AppTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTemplateOutput)
}

// AppTemplateArrayInput is an input type that accepts AppTemplateArray and AppTemplateArrayOutput values.
// You can construct a concrete instance of `AppTemplateArrayInput` via:
//
//	AppTemplateArray{ AppTemplateArgs{...} }
type AppTemplateArrayInput interface {
	pulumi.Input

	ToAppTemplateArrayOutput() AppTemplateArrayOutput
	ToAppTemplateArrayOutputWithContext(context.Context) AppTemplateArrayOutput
}

type AppTemplateArray []AppTemplateInput

func (AppTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppTemplate)(nil)).Elem()
}

func (i AppTemplateArray) ToAppTemplateArrayOutput() AppTemplateArrayOutput {
	return i.ToAppTemplateArrayOutputWithContext(context.Background())
}

func (i AppTemplateArray) ToAppTemplateArrayOutputWithContext(ctx context.Context) AppTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTemplateArrayOutput)
}

// AppTemplateMapInput is an input type that accepts AppTemplateMap and AppTemplateMapOutput values.
// You can construct a concrete instance of `AppTemplateMapInput` via:
//
//	AppTemplateMap{ "key": AppTemplateArgs{...} }
type AppTemplateMapInput interface {
	pulumi.Input

	ToAppTemplateMapOutput() AppTemplateMapOutput
	ToAppTemplateMapOutputWithContext(context.Context) AppTemplateMapOutput
}

type AppTemplateMap map[string]AppTemplateInput

func (AppTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppTemplate)(nil)).Elem()
}

func (i AppTemplateMap) ToAppTemplateMapOutput() AppTemplateMapOutput {
	return i.ToAppTemplateMapOutputWithContext(context.Background())
}

func (i AppTemplateMap) ToAppTemplateMapOutputWithContext(ctx context.Context) AppTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTemplateMapOutput)
}

type AppTemplateOutput struct{ *pulumi.OutputState }

func (AppTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppTemplate)(nil)).Elem()
}

func (o AppTemplateOutput) ToAppTemplateOutput() AppTemplateOutput {
	return o
}

func (o AppTemplateOutput) ToAppTemplateOutputWithContext(ctx context.Context) AppTemplateOutput {
	return o
}

// The name of the resource.
func (o AppTemplateOutput) AppTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppTemplate) pulumi.StringOutput { return v.AppTemplateName }).(pulumi.StringOutput)
}

// List of components. Its element valid values: ["component.live","component.liveRecord","component.liveBeauty","component.rtc","component.rtcRecord","component.im","component.whiteboard","component.liveSecurity","component.chatSecurity"].
func (o AppTemplateOutput) ComponentLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppTemplate) pulumi.StringArrayOutput { return v.ComponentLists }).(pulumi.StringArrayOutput)
}

// Configuration list. It have several default configs after the resource is created. See the following `Block configList`.
func (o AppTemplateOutput) ConfigLists() AppTemplateConfigListArrayOutput {
	return o.ApplyT(func(v *AppTemplate) AppTemplateConfigListArrayOutput { return v.ConfigLists }).(AppTemplateConfigListArrayOutput)
}

// Integration mode. Valid values:
// * paasSDK: Integrated SDK.
// * standardRoom: Model Room.
func (o AppTemplateOutput) IntegrationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppTemplate) pulumi.StringPtrOutput { return v.IntegrationMode }).(pulumi.StringPtrOutput)
}

// Application Template scenario. Valid values: ["business", "classroom"].
func (o AppTemplateOutput) Scene() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppTemplate) pulumi.StringPtrOutput { return v.Scene }).(pulumi.StringPtrOutput)
}

// Application template usage status.
func (o AppTemplateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppTemplate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AppTemplateArrayOutput struct{ *pulumi.OutputState }

func (AppTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppTemplate)(nil)).Elem()
}

func (o AppTemplateArrayOutput) ToAppTemplateArrayOutput() AppTemplateArrayOutput {
	return o
}

func (o AppTemplateArrayOutput) ToAppTemplateArrayOutputWithContext(ctx context.Context) AppTemplateArrayOutput {
	return o
}

func (o AppTemplateArrayOutput) Index(i pulumi.IntInput) AppTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppTemplate {
		return vs[0].([]*AppTemplate)[vs[1].(int)]
	}).(AppTemplateOutput)
}

type AppTemplateMapOutput struct{ *pulumi.OutputState }

func (AppTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppTemplate)(nil)).Elem()
}

func (o AppTemplateMapOutput) ToAppTemplateMapOutput() AppTemplateMapOutput {
	return o
}

func (o AppTemplateMapOutput) ToAppTemplateMapOutputWithContext(ctx context.Context) AppTemplateMapOutput {
	return o
}

func (o AppTemplateMapOutput) MapIndex(k pulumi.StringInput) AppTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppTemplate {
		return vs[0].(map[string]*AppTemplate)[vs[1].(string)]
	}).(AppTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppTemplateInput)(nil)).Elem(), &AppTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppTemplateArrayInput)(nil)).Elem(), AppTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppTemplateMapInput)(nil)).Elem(), AppTemplateMap{})
	pulumi.RegisterOutputType(AppTemplateOutput{})
	pulumi.RegisterOutputType(AppTemplateArrayOutput{})
	pulumi.RegisterOutputType(AppTemplateMapOutput{})
}
