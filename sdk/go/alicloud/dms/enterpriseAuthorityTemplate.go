// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS Enterprise Authority Template resource.
//
// For information about DMS Enterprise Authority Template and how to use it, see [What is Authority Template](https://www.alibabacloud.com/help/en/dms/developer-reference/api-dms-enterprise-2018-11-01-createauthoritytemplate).
//
// > **NOTE:** Available since v1.212.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultUserTenants, err := dms.GetUserTenants(ctx, &dms.GetUserTenantsArgs{
//				Status: pulumi.StringRef("ACTIVE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = dms.NewEnterpriseAuthorityTemplate(ctx, "defaultEnterpriseAuthorityTemplate", &dms.EnterpriseAuthorityTemplateArgs{
//				Tid:                   *pulumi.String(defaultUserTenants.Ids[0]),
//				AuthorityTemplateName: pulumi.String(name),
//				Description:           pulumi.String(name),
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
// DMS Enterprise Authority Template can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate example <tid>:<authority_template_id>
// ```
type EnterpriseAuthorityTemplate struct {
	pulumi.CustomResourceState

	// Permission template ID.
	AuthorityTemplateId pulumi.IntOutput `pulumi:"authorityTemplateId"`
	// Permission Template name.
	AuthorityTemplateName pulumi.StringOutput `pulumi:"authorityTemplateName"`
	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Permission template description information.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Tenant ID.
	Tid pulumi.IntOutput `pulumi:"tid"`
}

// NewEnterpriseAuthorityTemplate registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseAuthorityTemplate(ctx *pulumi.Context,
	name string, args *EnterpriseAuthorityTemplateArgs, opts ...pulumi.ResourceOption) (*EnterpriseAuthorityTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorityTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'AuthorityTemplateName'")
	}
	if args.Tid == nil {
		return nil, errors.New("invalid value for required argument 'Tid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseAuthorityTemplate
	err := ctx.RegisterResource("alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseAuthorityTemplate gets an existing EnterpriseAuthorityTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseAuthorityTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseAuthorityTemplateState, opts ...pulumi.ResourceOption) (*EnterpriseAuthorityTemplate, error) {
	var resource EnterpriseAuthorityTemplate
	err := ctx.ReadResource("alicloud:dms/enterpriseAuthorityTemplate:EnterpriseAuthorityTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseAuthorityTemplate resources.
type enterpriseAuthorityTemplateState struct {
	// Permission template ID.
	AuthorityTemplateId *int `pulumi:"authorityTemplateId"`
	// Permission Template name.
	AuthorityTemplateName *string `pulumi:"authorityTemplateName"`
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// Permission template description information.
	Description *string `pulumi:"description"`
	// Tenant ID.
	Tid *int `pulumi:"tid"`
}

type EnterpriseAuthorityTemplateState struct {
	// Permission template ID.
	AuthorityTemplateId pulumi.IntPtrInput
	// Permission Template name.
	AuthorityTemplateName pulumi.StringPtrInput
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// Permission template description information.
	Description pulumi.StringPtrInput
	// Tenant ID.
	Tid pulumi.IntPtrInput
}

func (EnterpriseAuthorityTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseAuthorityTemplateState)(nil)).Elem()
}

type enterpriseAuthorityTemplateArgs struct {
	// Permission Template name.
	AuthorityTemplateName string `pulumi:"authorityTemplateName"`
	// Permission template description information.
	Description *string `pulumi:"description"`
	// Tenant ID.
	Tid int `pulumi:"tid"`
}

// The set of arguments for constructing a EnterpriseAuthorityTemplate resource.
type EnterpriseAuthorityTemplateArgs struct {
	// Permission Template name.
	AuthorityTemplateName pulumi.StringInput
	// Permission template description information.
	Description pulumi.StringPtrInput
	// Tenant ID.
	Tid pulumi.IntInput
}

func (EnterpriseAuthorityTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseAuthorityTemplateArgs)(nil)).Elem()
}

type EnterpriseAuthorityTemplateInput interface {
	pulumi.Input

	ToEnterpriseAuthorityTemplateOutput() EnterpriseAuthorityTemplateOutput
	ToEnterpriseAuthorityTemplateOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateOutput
}

func (*EnterpriseAuthorityTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (i *EnterpriseAuthorityTemplate) ToEnterpriseAuthorityTemplateOutput() EnterpriseAuthorityTemplateOutput {
	return i.ToEnterpriseAuthorityTemplateOutputWithContext(context.Background())
}

func (i *EnterpriseAuthorityTemplate) ToEnterpriseAuthorityTemplateOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseAuthorityTemplateOutput)
}

// EnterpriseAuthorityTemplateArrayInput is an input type that accepts EnterpriseAuthorityTemplateArray and EnterpriseAuthorityTemplateArrayOutput values.
// You can construct a concrete instance of `EnterpriseAuthorityTemplateArrayInput` via:
//
//	EnterpriseAuthorityTemplateArray{ EnterpriseAuthorityTemplateArgs{...} }
type EnterpriseAuthorityTemplateArrayInput interface {
	pulumi.Input

	ToEnterpriseAuthorityTemplateArrayOutput() EnterpriseAuthorityTemplateArrayOutput
	ToEnterpriseAuthorityTemplateArrayOutputWithContext(context.Context) EnterpriseAuthorityTemplateArrayOutput
}

type EnterpriseAuthorityTemplateArray []EnterpriseAuthorityTemplateInput

func (EnterpriseAuthorityTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (i EnterpriseAuthorityTemplateArray) ToEnterpriseAuthorityTemplateArrayOutput() EnterpriseAuthorityTemplateArrayOutput {
	return i.ToEnterpriseAuthorityTemplateArrayOutputWithContext(context.Background())
}

func (i EnterpriseAuthorityTemplateArray) ToEnterpriseAuthorityTemplateArrayOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseAuthorityTemplateArrayOutput)
}

// EnterpriseAuthorityTemplateMapInput is an input type that accepts EnterpriseAuthorityTemplateMap and EnterpriseAuthorityTemplateMapOutput values.
// You can construct a concrete instance of `EnterpriseAuthorityTemplateMapInput` via:
//
//	EnterpriseAuthorityTemplateMap{ "key": EnterpriseAuthorityTemplateArgs{...} }
type EnterpriseAuthorityTemplateMapInput interface {
	pulumi.Input

	ToEnterpriseAuthorityTemplateMapOutput() EnterpriseAuthorityTemplateMapOutput
	ToEnterpriseAuthorityTemplateMapOutputWithContext(context.Context) EnterpriseAuthorityTemplateMapOutput
}

type EnterpriseAuthorityTemplateMap map[string]EnterpriseAuthorityTemplateInput

func (EnterpriseAuthorityTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (i EnterpriseAuthorityTemplateMap) ToEnterpriseAuthorityTemplateMapOutput() EnterpriseAuthorityTemplateMapOutput {
	return i.ToEnterpriseAuthorityTemplateMapOutputWithContext(context.Background())
}

func (i EnterpriseAuthorityTemplateMap) ToEnterpriseAuthorityTemplateMapOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseAuthorityTemplateMapOutput)
}

type EnterpriseAuthorityTemplateOutput struct{ *pulumi.OutputState }

func (EnterpriseAuthorityTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (o EnterpriseAuthorityTemplateOutput) ToEnterpriseAuthorityTemplateOutput() EnterpriseAuthorityTemplateOutput {
	return o
}

func (o EnterpriseAuthorityTemplateOutput) ToEnterpriseAuthorityTemplateOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateOutput {
	return o
}

// Permission template ID.
func (o EnterpriseAuthorityTemplateOutput) AuthorityTemplateId() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseAuthorityTemplate) pulumi.IntOutput { return v.AuthorityTemplateId }).(pulumi.IntOutput)
}

// Permission Template name.
func (o EnterpriseAuthorityTemplateOutput) AuthorityTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseAuthorityTemplate) pulumi.StringOutput { return v.AuthorityTemplateName }).(pulumi.StringOutput)
}

// The creation time of the resource.
func (o EnterpriseAuthorityTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseAuthorityTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Permission template description information.
func (o EnterpriseAuthorityTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseAuthorityTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Tenant ID.
func (o EnterpriseAuthorityTemplateOutput) Tid() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseAuthorityTemplate) pulumi.IntOutput { return v.Tid }).(pulumi.IntOutput)
}

type EnterpriseAuthorityTemplateArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseAuthorityTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (o EnterpriseAuthorityTemplateArrayOutput) ToEnterpriseAuthorityTemplateArrayOutput() EnterpriseAuthorityTemplateArrayOutput {
	return o
}

func (o EnterpriseAuthorityTemplateArrayOutput) ToEnterpriseAuthorityTemplateArrayOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateArrayOutput {
	return o
}

func (o EnterpriseAuthorityTemplateArrayOutput) Index(i pulumi.IntInput) EnterpriseAuthorityTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseAuthorityTemplate {
		return vs[0].([]*EnterpriseAuthorityTemplate)[vs[1].(int)]
	}).(EnterpriseAuthorityTemplateOutput)
}

type EnterpriseAuthorityTemplateMapOutput struct{ *pulumi.OutputState }

func (EnterpriseAuthorityTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseAuthorityTemplate)(nil)).Elem()
}

func (o EnterpriseAuthorityTemplateMapOutput) ToEnterpriseAuthorityTemplateMapOutput() EnterpriseAuthorityTemplateMapOutput {
	return o
}

func (o EnterpriseAuthorityTemplateMapOutput) ToEnterpriseAuthorityTemplateMapOutputWithContext(ctx context.Context) EnterpriseAuthorityTemplateMapOutput {
	return o
}

func (o EnterpriseAuthorityTemplateMapOutput) MapIndex(k pulumi.StringInput) EnterpriseAuthorityTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseAuthorityTemplate {
		return vs[0].(map[string]*EnterpriseAuthorityTemplate)[vs[1].(string)]
	}).(EnterpriseAuthorityTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseAuthorityTemplateInput)(nil)).Elem(), &EnterpriseAuthorityTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseAuthorityTemplateArrayInput)(nil)).Elem(), EnterpriseAuthorityTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseAuthorityTemplateMapInput)(nil)).Elem(), EnterpriseAuthorityTemplateMap{})
	pulumi.RegisterOutputType(EnterpriseAuthorityTemplateOutput{})
	pulumi.RegisterOutputType(EnterpriseAuthorityTemplateArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseAuthorityTemplateMapOutput{})
}
