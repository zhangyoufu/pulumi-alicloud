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

// Provides a Oos Default Patch Baseline resource.
//
// For information about Oos Default Patch Baseline and how to use it, see [What is Default Patch Baseline](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/api-oos-2019-06-01-registerdefaultpatchbaseline).
//
// > **NOTE:** Available in v1.203.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultPatchBaseline, err := oos.NewPatchBaseline(ctx, "defaultPatchBaseline", &oos.PatchBaselineArgs{
//				OperationSystem:   pulumi.String("Windows"),
//				PatchBaselineName: pulumi.String("terraform-example"),
//				Description:       pulumi.String("terraform-example"),
//				ApprovalRules:     pulumi.String("{\"PatchRules\":[{\"PatchFilterGroup\":[{\"Key\":\"PatchSet\",\"Values\":[\"OS\"]},{\"Key\":\"ProductFamily\",\"Values\":[\"Windows\"]},{\"Key\":\"Product\",\"Values\":[\"Windows 10\",\"Windows 7\"]},{\"Key\":\"Classification\",\"Values\":[\"Security Updates\",\"Updates\",\"Update Rollups\",\"Critical Updates\"]},{\"Key\":\"Severity\",\"Values\":[\"Critical\",\"Important\",\"Moderate\"]}],\"ApproveAfterDays\":7,\"EnableNonSecurity\":true,\"ComplianceLevel\":\"Medium\"}]}"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oos.NewDefaultPatchBaseline(ctx, "defaultDefaultPatchBaseline", &oos.DefaultPatchBaselineArgs{
//				PatchBaselineName: defaultPatchBaseline.ID(),
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
// Oos Default Patch Baseline can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline example <id>
// ```
type DefaultPatchBaseline struct {
	pulumi.CustomResourceState

	// The ID of the patch baseline.
	PatchBaselineId pulumi.StringOutput `pulumi:"patchBaselineId"`
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringOutput `pulumi:"patchBaselineName"`
}

// NewDefaultPatchBaseline registers a new resource with the given unique name, arguments, and options.
func NewDefaultPatchBaseline(ctx *pulumi.Context,
	name string, args *DefaultPatchBaselineArgs, opts ...pulumi.ResourceOption) (*DefaultPatchBaseline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PatchBaselineName == nil {
		return nil, errors.New("invalid value for required argument 'PatchBaselineName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultPatchBaseline
	err := ctx.RegisterResource("alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultPatchBaseline gets an existing DefaultPatchBaseline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPatchBaseline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultPatchBaselineState, opts ...pulumi.ResourceOption) (*DefaultPatchBaseline, error) {
	var resource DefaultPatchBaseline
	err := ctx.ReadResource("alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultPatchBaseline resources.
type defaultPatchBaselineState struct {
	// The ID of the patch baseline.
	PatchBaselineId *string `pulumi:"patchBaselineId"`
	// The name of the patch baseline.
	PatchBaselineName *string `pulumi:"patchBaselineName"`
}

type DefaultPatchBaselineState struct {
	// The ID of the patch baseline.
	PatchBaselineId pulumi.StringPtrInput
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringPtrInput
}

func (DefaultPatchBaselineState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPatchBaselineState)(nil)).Elem()
}

type defaultPatchBaselineArgs struct {
	// The name of the patch baseline.
	PatchBaselineName string `pulumi:"patchBaselineName"`
}

// The set of arguments for constructing a DefaultPatchBaseline resource.
type DefaultPatchBaselineArgs struct {
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringInput
}

func (DefaultPatchBaselineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPatchBaselineArgs)(nil)).Elem()
}

type DefaultPatchBaselineInput interface {
	pulumi.Input

	ToDefaultPatchBaselineOutput() DefaultPatchBaselineOutput
	ToDefaultPatchBaselineOutputWithContext(ctx context.Context) DefaultPatchBaselineOutput
}

func (*DefaultPatchBaseline) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPatchBaseline)(nil)).Elem()
}

func (i *DefaultPatchBaseline) ToDefaultPatchBaselineOutput() DefaultPatchBaselineOutput {
	return i.ToDefaultPatchBaselineOutputWithContext(context.Background())
}

func (i *DefaultPatchBaseline) ToDefaultPatchBaselineOutputWithContext(ctx context.Context) DefaultPatchBaselineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPatchBaselineOutput)
}

// DefaultPatchBaselineArrayInput is an input type that accepts DefaultPatchBaselineArray and DefaultPatchBaselineArrayOutput values.
// You can construct a concrete instance of `DefaultPatchBaselineArrayInput` via:
//
//	DefaultPatchBaselineArray{ DefaultPatchBaselineArgs{...} }
type DefaultPatchBaselineArrayInput interface {
	pulumi.Input

	ToDefaultPatchBaselineArrayOutput() DefaultPatchBaselineArrayOutput
	ToDefaultPatchBaselineArrayOutputWithContext(context.Context) DefaultPatchBaselineArrayOutput
}

type DefaultPatchBaselineArray []DefaultPatchBaselineInput

func (DefaultPatchBaselineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPatchBaseline)(nil)).Elem()
}

func (i DefaultPatchBaselineArray) ToDefaultPatchBaselineArrayOutput() DefaultPatchBaselineArrayOutput {
	return i.ToDefaultPatchBaselineArrayOutputWithContext(context.Background())
}

func (i DefaultPatchBaselineArray) ToDefaultPatchBaselineArrayOutputWithContext(ctx context.Context) DefaultPatchBaselineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPatchBaselineArrayOutput)
}

// DefaultPatchBaselineMapInput is an input type that accepts DefaultPatchBaselineMap and DefaultPatchBaselineMapOutput values.
// You can construct a concrete instance of `DefaultPatchBaselineMapInput` via:
//
//	DefaultPatchBaselineMap{ "key": DefaultPatchBaselineArgs{...} }
type DefaultPatchBaselineMapInput interface {
	pulumi.Input

	ToDefaultPatchBaselineMapOutput() DefaultPatchBaselineMapOutput
	ToDefaultPatchBaselineMapOutputWithContext(context.Context) DefaultPatchBaselineMapOutput
}

type DefaultPatchBaselineMap map[string]DefaultPatchBaselineInput

func (DefaultPatchBaselineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPatchBaseline)(nil)).Elem()
}

func (i DefaultPatchBaselineMap) ToDefaultPatchBaselineMapOutput() DefaultPatchBaselineMapOutput {
	return i.ToDefaultPatchBaselineMapOutputWithContext(context.Background())
}

func (i DefaultPatchBaselineMap) ToDefaultPatchBaselineMapOutputWithContext(ctx context.Context) DefaultPatchBaselineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPatchBaselineMapOutput)
}

type DefaultPatchBaselineOutput struct{ *pulumi.OutputState }

func (DefaultPatchBaselineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPatchBaseline)(nil)).Elem()
}

func (o DefaultPatchBaselineOutput) ToDefaultPatchBaselineOutput() DefaultPatchBaselineOutput {
	return o
}

func (o DefaultPatchBaselineOutput) ToDefaultPatchBaselineOutputWithContext(ctx context.Context) DefaultPatchBaselineOutput {
	return o
}

// The ID of the patch baseline.
func (o DefaultPatchBaselineOutput) PatchBaselineId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPatchBaseline) pulumi.StringOutput { return v.PatchBaselineId }).(pulumi.StringOutput)
}

// The name of the patch baseline.
func (o DefaultPatchBaselineOutput) PatchBaselineName() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPatchBaseline) pulumi.StringOutput { return v.PatchBaselineName }).(pulumi.StringOutput)
}

type DefaultPatchBaselineArrayOutput struct{ *pulumi.OutputState }

func (DefaultPatchBaselineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPatchBaseline)(nil)).Elem()
}

func (o DefaultPatchBaselineArrayOutput) ToDefaultPatchBaselineArrayOutput() DefaultPatchBaselineArrayOutput {
	return o
}

func (o DefaultPatchBaselineArrayOutput) ToDefaultPatchBaselineArrayOutputWithContext(ctx context.Context) DefaultPatchBaselineArrayOutput {
	return o
}

func (o DefaultPatchBaselineArrayOutput) Index(i pulumi.IntInput) DefaultPatchBaselineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultPatchBaseline {
		return vs[0].([]*DefaultPatchBaseline)[vs[1].(int)]
	}).(DefaultPatchBaselineOutput)
}

type DefaultPatchBaselineMapOutput struct{ *pulumi.OutputState }

func (DefaultPatchBaselineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPatchBaseline)(nil)).Elem()
}

func (o DefaultPatchBaselineMapOutput) ToDefaultPatchBaselineMapOutput() DefaultPatchBaselineMapOutput {
	return o
}

func (o DefaultPatchBaselineMapOutput) ToDefaultPatchBaselineMapOutputWithContext(ctx context.Context) DefaultPatchBaselineMapOutput {
	return o
}

func (o DefaultPatchBaselineMapOutput) MapIndex(k pulumi.StringInput) DefaultPatchBaselineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultPatchBaseline {
		return vs[0].(map[string]*DefaultPatchBaseline)[vs[1].(string)]
	}).(DefaultPatchBaselineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPatchBaselineInput)(nil)).Elem(), &DefaultPatchBaseline{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPatchBaselineArrayInput)(nil)).Elem(), DefaultPatchBaselineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPatchBaselineMapInput)(nil)).Elem(), DefaultPatchBaselineMap{})
	pulumi.RegisterOutputType(DefaultPatchBaselineOutput{})
	pulumi.RegisterOutputType(DefaultPatchBaselineArrayOutput{})
	pulumi.RegisterOutputType(DefaultPatchBaselineMapOutput{})
}
