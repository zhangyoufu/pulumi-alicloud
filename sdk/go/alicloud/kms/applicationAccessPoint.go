// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a KMS Application Access Point resource. An application access point (AAP) is used to implement fine-grained access control for Key Management Service (KMS) resources. An application can access a KMS instance only after an AAP is created for the application. .
//
// For information about KMS Application Access Point and how to use it, see [What is Application Access Point](https://www.alibabacloud.com/help/zh/key-management-service/latest/api-createapplicationaccesspoint).
//
// > **NOTE:** Available since v1.210.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/kms"
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
//			_, err := kms.NewApplicationAccessPoint(ctx, "default", &kms.ApplicationAccessPointArgs{
//				Description:                pulumi.String("example aap"),
//				ApplicationAccessPointName: pulumi.String(name),
//				Policies: pulumi.StringArray{
//					pulumi.String("abc"),
//					pulumi.String("efg"),
//					pulumi.String("hfc"),
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
// KMS Application Access Point can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:kms/applicationAccessPoint:ApplicationAccessPoint example <id>
// ```
type ApplicationAccessPoint struct {
	pulumi.CustomResourceState

	// Application Access Point Name.
	ApplicationAccessPointName pulumi.StringOutput `pulumi:"applicationAccessPointName"`
	// Description .
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The policies that have bound to the Application Access Point (AAP).
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewApplicationAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewApplicationAccessPoint(ctx *pulumi.Context,
	name string, args *ApplicationAccessPointArgs, opts ...pulumi.ResourceOption) (*ApplicationAccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationAccessPointName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationAccessPointName'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationAccessPoint
	err := ctx.RegisterResource("alicloud:kms/applicationAccessPoint:ApplicationAccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationAccessPoint gets an existing ApplicationAccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationAccessPointState, opts ...pulumi.ResourceOption) (*ApplicationAccessPoint, error) {
	var resource ApplicationAccessPoint
	err := ctx.ReadResource("alicloud:kms/applicationAccessPoint:ApplicationAccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationAccessPoint resources.
type applicationAccessPointState struct {
	// Application Access Point Name.
	ApplicationAccessPointName *string `pulumi:"applicationAccessPointName"`
	// Description .
	Description *string `pulumi:"description"`
	// The policies that have bound to the Application Access Point (AAP).
	Policies []string `pulumi:"policies"`
}

type ApplicationAccessPointState struct {
	// Application Access Point Name.
	ApplicationAccessPointName pulumi.StringPtrInput
	// Description .
	Description pulumi.StringPtrInput
	// The policies that have bound to the Application Access Point (AAP).
	Policies pulumi.StringArrayInput
}

func (ApplicationAccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAccessPointState)(nil)).Elem()
}

type applicationAccessPointArgs struct {
	// Application Access Point Name.
	ApplicationAccessPointName string `pulumi:"applicationAccessPointName"`
	// Description .
	Description *string `pulumi:"description"`
	// The policies that have bound to the Application Access Point (AAP).
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a ApplicationAccessPoint resource.
type ApplicationAccessPointArgs struct {
	// Application Access Point Name.
	ApplicationAccessPointName pulumi.StringInput
	// Description .
	Description pulumi.StringPtrInput
	// The policies that have bound to the Application Access Point (AAP).
	Policies pulumi.StringArrayInput
}

func (ApplicationAccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationAccessPointArgs)(nil)).Elem()
}

type ApplicationAccessPointInput interface {
	pulumi.Input

	ToApplicationAccessPointOutput() ApplicationAccessPointOutput
	ToApplicationAccessPointOutputWithContext(ctx context.Context) ApplicationAccessPointOutput
}

func (*ApplicationAccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccessPoint)(nil)).Elem()
}

func (i *ApplicationAccessPoint) ToApplicationAccessPointOutput() ApplicationAccessPointOutput {
	return i.ToApplicationAccessPointOutputWithContext(context.Background())
}

func (i *ApplicationAccessPoint) ToApplicationAccessPointOutputWithContext(ctx context.Context) ApplicationAccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAccessPointOutput)
}

// ApplicationAccessPointArrayInput is an input type that accepts ApplicationAccessPointArray and ApplicationAccessPointArrayOutput values.
// You can construct a concrete instance of `ApplicationAccessPointArrayInput` via:
//
//	ApplicationAccessPointArray{ ApplicationAccessPointArgs{...} }
type ApplicationAccessPointArrayInput interface {
	pulumi.Input

	ToApplicationAccessPointArrayOutput() ApplicationAccessPointArrayOutput
	ToApplicationAccessPointArrayOutputWithContext(context.Context) ApplicationAccessPointArrayOutput
}

type ApplicationAccessPointArray []ApplicationAccessPointInput

func (ApplicationAccessPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationAccessPoint)(nil)).Elem()
}

func (i ApplicationAccessPointArray) ToApplicationAccessPointArrayOutput() ApplicationAccessPointArrayOutput {
	return i.ToApplicationAccessPointArrayOutputWithContext(context.Background())
}

func (i ApplicationAccessPointArray) ToApplicationAccessPointArrayOutputWithContext(ctx context.Context) ApplicationAccessPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAccessPointArrayOutput)
}

// ApplicationAccessPointMapInput is an input type that accepts ApplicationAccessPointMap and ApplicationAccessPointMapOutput values.
// You can construct a concrete instance of `ApplicationAccessPointMapInput` via:
//
//	ApplicationAccessPointMap{ "key": ApplicationAccessPointArgs{...} }
type ApplicationAccessPointMapInput interface {
	pulumi.Input

	ToApplicationAccessPointMapOutput() ApplicationAccessPointMapOutput
	ToApplicationAccessPointMapOutputWithContext(context.Context) ApplicationAccessPointMapOutput
}

type ApplicationAccessPointMap map[string]ApplicationAccessPointInput

func (ApplicationAccessPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationAccessPoint)(nil)).Elem()
}

func (i ApplicationAccessPointMap) ToApplicationAccessPointMapOutput() ApplicationAccessPointMapOutput {
	return i.ToApplicationAccessPointMapOutputWithContext(context.Background())
}

func (i ApplicationAccessPointMap) ToApplicationAccessPointMapOutputWithContext(ctx context.Context) ApplicationAccessPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationAccessPointMapOutput)
}

type ApplicationAccessPointOutput struct{ *pulumi.OutputState }

func (ApplicationAccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationAccessPoint)(nil)).Elem()
}

func (o ApplicationAccessPointOutput) ToApplicationAccessPointOutput() ApplicationAccessPointOutput {
	return o
}

func (o ApplicationAccessPointOutput) ToApplicationAccessPointOutputWithContext(ctx context.Context) ApplicationAccessPointOutput {
	return o
}

// Application Access Point Name.
func (o ApplicationAccessPointOutput) ApplicationAccessPointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationAccessPoint) pulumi.StringOutput { return v.ApplicationAccessPointName }).(pulumi.StringOutput)
}

// Description .
func (o ApplicationAccessPointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationAccessPoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The policies that have bound to the Application Access Point (AAP).
func (o ApplicationAccessPointOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationAccessPoint) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

type ApplicationAccessPointArrayOutput struct{ *pulumi.OutputState }

func (ApplicationAccessPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationAccessPoint)(nil)).Elem()
}

func (o ApplicationAccessPointArrayOutput) ToApplicationAccessPointArrayOutput() ApplicationAccessPointArrayOutput {
	return o
}

func (o ApplicationAccessPointArrayOutput) ToApplicationAccessPointArrayOutputWithContext(ctx context.Context) ApplicationAccessPointArrayOutput {
	return o
}

func (o ApplicationAccessPointArrayOutput) Index(i pulumi.IntInput) ApplicationAccessPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationAccessPoint {
		return vs[0].([]*ApplicationAccessPoint)[vs[1].(int)]
	}).(ApplicationAccessPointOutput)
}

type ApplicationAccessPointMapOutput struct{ *pulumi.OutputState }

func (ApplicationAccessPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationAccessPoint)(nil)).Elem()
}

func (o ApplicationAccessPointMapOutput) ToApplicationAccessPointMapOutput() ApplicationAccessPointMapOutput {
	return o
}

func (o ApplicationAccessPointMapOutput) ToApplicationAccessPointMapOutputWithContext(ctx context.Context) ApplicationAccessPointMapOutput {
	return o
}

func (o ApplicationAccessPointMapOutput) MapIndex(k pulumi.StringInput) ApplicationAccessPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationAccessPoint {
		return vs[0].(map[string]*ApplicationAccessPoint)[vs[1].(string)]
	}).(ApplicationAccessPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAccessPointInput)(nil)).Elem(), &ApplicationAccessPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAccessPointArrayInput)(nil)).Elem(), ApplicationAccessPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationAccessPointMapInput)(nil)).Elem(), ApplicationAccessPointMap{})
	pulumi.RegisterOutputType(ApplicationAccessPointOutput{})
	pulumi.RegisterOutputType(ApplicationAccessPointArrayOutput{})
	pulumi.RegisterOutputType(ApplicationAccessPointMapOutput{})
}
