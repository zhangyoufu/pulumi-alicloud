// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Api Gateway Backend resource.
//
// For information about Api Gateway Backend and how to use it, see [What is Backend](https://www.alibabacloud.com/help/en/api-gateway/developer-reference/api-cloudapi-2016-07-14-createbackend).
//
// > **NOTE:** Available since v1.181.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf_example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := apigateway.NewBackend(ctx, "default", &apigateway.BackendArgs{
//				BackendName: pulumi.String(name),
//				Description: pulumi.String(name),
//				BackendType: pulumi.String("HTTP"),
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
// Api Gateway Backend can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:apigateway/backend:Backend example <id>
// ```
type Backend struct {
	pulumi.CustomResourceState

	// The name of the Backend.
	BackendName pulumi.StringOutput `pulumi:"backendName"`
	// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
	BackendType pulumi.StringOutput `pulumi:"backendType"`
	// Whether to create an Event bus service association role.
	CreateEventBridgeServiceLinkedRole pulumi.BoolPtrOutput `pulumi:"createEventBridgeServiceLinkedRole"`
	// The description of the Backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendName == nil {
		return nil, errors.New("invalid value for required argument 'BackendName'")
	}
	if args.BackendType == nil {
		return nil, errors.New("invalid value for required argument 'BackendType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Backend
	err := ctx.RegisterResource("alicloud:apigateway/backend:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("alicloud:apigateway/backend:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
	// The name of the Backend.
	BackendName *string `pulumi:"backendName"`
	// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
	BackendType *string `pulumi:"backendType"`
	// Whether to create an Event bus service association role.
	CreateEventBridgeServiceLinkedRole *bool `pulumi:"createEventBridgeServiceLinkedRole"`
	// The description of the Backend.
	Description *string `pulumi:"description"`
}

type BackendState struct {
	// The name of the Backend.
	BackendName pulumi.StringPtrInput
	// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
	BackendType pulumi.StringPtrInput
	// Whether to create an Event bus service association role.
	CreateEventBridgeServiceLinkedRole pulumi.BoolPtrInput
	// The description of the Backend.
	Description pulumi.StringPtrInput
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// The name of the Backend.
	BackendName string `pulumi:"backendName"`
	// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
	BackendType string `pulumi:"backendType"`
	// Whether to create an Event bus service association role.
	CreateEventBridgeServiceLinkedRole *bool `pulumi:"createEventBridgeServiceLinkedRole"`
	// The description of the Backend.
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// The name of the Backend.
	BackendName pulumi.StringInput
	// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
	BackendType pulumi.StringInput
	// Whether to create an Event bus service association role.
	CreateEventBridgeServiceLinkedRole pulumi.BoolPtrInput
	// The description of the Backend.
	Description pulumi.StringPtrInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (*Backend) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (i *Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i *Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

// BackendArrayInput is an input type that accepts BackendArray and BackendArrayOutput values.
// You can construct a concrete instance of `BackendArrayInput` via:
//
//	BackendArray{ BackendArgs{...} }
type BackendArrayInput interface {
	pulumi.Input

	ToBackendArrayOutput() BackendArrayOutput
	ToBackendArrayOutputWithContext(context.Context) BackendArrayOutput
}

type BackendArray []BackendInput

func (BackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backend)(nil)).Elem()
}

func (i BackendArray) ToBackendArrayOutput() BackendArrayOutput {
	return i.ToBackendArrayOutputWithContext(context.Background())
}

func (i BackendArray) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendArrayOutput)
}

// BackendMapInput is an input type that accepts BackendMap and BackendMapOutput values.
// You can construct a concrete instance of `BackendMapInput` via:
//
//	BackendMap{ "key": BackendArgs{...} }
type BackendMapInput interface {
	pulumi.Input

	ToBackendMapOutput() BackendMapOutput
	ToBackendMapOutputWithContext(context.Context) BackendMapOutput
}

type BackendMap map[string]BackendInput

func (BackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backend)(nil)).Elem()
}

func (i BackendMap) ToBackendMapOutput() BackendMapOutput {
	return i.ToBackendMapOutputWithContext(context.Background())
}

func (i BackendMap) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendMapOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

// The name of the Backend.
func (o BackendOutput) BackendName() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.BackendName }).(pulumi.StringOutput)
}

// The type of the Backend. Valid values: `HTTP`, `VPC`, `FC_EVENT`, `FC_HTTP`, `OSS`, `MOCK`.
func (o BackendOutput) BackendType() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.BackendType }).(pulumi.StringOutput)
}

// Whether to create an Event bus service association role.
func (o BackendOutput) CreateEventBridgeServiceLinkedRole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.BoolPtrOutput { return v.CreateEventBridgeServiceLinkedRole }).(pulumi.BoolPtrOutput)
}

// The description of the Backend.
func (o BackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

type BackendArrayOutput struct{ *pulumi.OutputState }

func (BackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backend)(nil)).Elem()
}

func (o BackendArrayOutput) ToBackendArrayOutput() BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) Index(i pulumi.IntInput) BackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Backend {
		return vs[0].([]*Backend)[vs[1].(int)]
	}).(BackendOutput)
}

type BackendMapOutput struct{ *pulumi.OutputState }

func (BackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backend)(nil)).Elem()
}

func (o BackendMapOutput) ToBackendMapOutput() BackendMapOutput {
	return o
}

func (o BackendMapOutput) ToBackendMapOutputWithContext(ctx context.Context) BackendMapOutput {
	return o
}

func (o BackendMapOutput) MapIndex(k pulumi.StringInput) BackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Backend {
		return vs[0].(map[string]*Backend)[vs[1].(string)]
	}).(BackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendInput)(nil)).Elem(), &Backend{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendArrayInput)(nil)).Elem(), BackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendMapInput)(nil)).Elem(), BackendMap{})
	pulumi.RegisterOutputType(BackendOutput{})
	pulumi.RegisterOutputType(BackendArrayOutput{})
	pulumi.RegisterOutputType(BackendMapOutput{})
}
