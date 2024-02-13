// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudmonitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Basic Public resource.
//
// For information about Cloud Monitor Service Basic Public and how to use it, see [What is Basic Public](https://www.alibabacloud.com/help/en/cms/product-overview/what-is-cloudmonitor).
//
// > **NOTE:** Available since v1.215.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudmonitor"
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
//			_, err := cloudmonitor.NewServiceBasicPublic(ctx, "default", nil)
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
// Cloud Monitor Service Basic Public can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudmonitor/serviceBasicPublic:ServiceBasicPublic example <id>
// ```
type ServiceBasicPublic struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
}

// NewServiceBasicPublic registers a new resource with the given unique name, arguments, and options.
func NewServiceBasicPublic(ctx *pulumi.Context,
	name string, args *ServiceBasicPublicArgs, opts ...pulumi.ResourceOption) (*ServiceBasicPublic, error) {
	if args == nil {
		args = &ServiceBasicPublicArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceBasicPublic
	err := ctx.RegisterResource("alicloud:cloudmonitor/serviceBasicPublic:ServiceBasicPublic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceBasicPublic gets an existing ServiceBasicPublic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceBasicPublic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceBasicPublicState, opts ...pulumi.ResourceOption) (*ServiceBasicPublic, error) {
	var resource ServiceBasicPublic
	err := ctx.ReadResource("alicloud:cloudmonitor/serviceBasicPublic:ServiceBasicPublic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceBasicPublic resources.
type serviceBasicPublicState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
}

type ServiceBasicPublicState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
}

func (ServiceBasicPublicState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceBasicPublicState)(nil)).Elem()
}

type serviceBasicPublicArgs struct {
}

// The set of arguments for constructing a ServiceBasicPublic resource.
type ServiceBasicPublicArgs struct {
}

func (ServiceBasicPublicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceBasicPublicArgs)(nil)).Elem()
}

type ServiceBasicPublicInput interface {
	pulumi.Input

	ToServiceBasicPublicOutput() ServiceBasicPublicOutput
	ToServiceBasicPublicOutputWithContext(ctx context.Context) ServiceBasicPublicOutput
}

func (*ServiceBasicPublic) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBasicPublic)(nil)).Elem()
}

func (i *ServiceBasicPublic) ToServiceBasicPublicOutput() ServiceBasicPublicOutput {
	return i.ToServiceBasicPublicOutputWithContext(context.Background())
}

func (i *ServiceBasicPublic) ToServiceBasicPublicOutputWithContext(ctx context.Context) ServiceBasicPublicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBasicPublicOutput)
}

// ServiceBasicPublicArrayInput is an input type that accepts ServiceBasicPublicArray and ServiceBasicPublicArrayOutput values.
// You can construct a concrete instance of `ServiceBasicPublicArrayInput` via:
//
//	ServiceBasicPublicArray{ ServiceBasicPublicArgs{...} }
type ServiceBasicPublicArrayInput interface {
	pulumi.Input

	ToServiceBasicPublicArrayOutput() ServiceBasicPublicArrayOutput
	ToServiceBasicPublicArrayOutputWithContext(context.Context) ServiceBasicPublicArrayOutput
}

type ServiceBasicPublicArray []ServiceBasicPublicInput

func (ServiceBasicPublicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceBasicPublic)(nil)).Elem()
}

func (i ServiceBasicPublicArray) ToServiceBasicPublicArrayOutput() ServiceBasicPublicArrayOutput {
	return i.ToServiceBasicPublicArrayOutputWithContext(context.Background())
}

func (i ServiceBasicPublicArray) ToServiceBasicPublicArrayOutputWithContext(ctx context.Context) ServiceBasicPublicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBasicPublicArrayOutput)
}

// ServiceBasicPublicMapInput is an input type that accepts ServiceBasicPublicMap and ServiceBasicPublicMapOutput values.
// You can construct a concrete instance of `ServiceBasicPublicMapInput` via:
//
//	ServiceBasicPublicMap{ "key": ServiceBasicPublicArgs{...} }
type ServiceBasicPublicMapInput interface {
	pulumi.Input

	ToServiceBasicPublicMapOutput() ServiceBasicPublicMapOutput
	ToServiceBasicPublicMapOutputWithContext(context.Context) ServiceBasicPublicMapOutput
}

type ServiceBasicPublicMap map[string]ServiceBasicPublicInput

func (ServiceBasicPublicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceBasicPublic)(nil)).Elem()
}

func (i ServiceBasicPublicMap) ToServiceBasicPublicMapOutput() ServiceBasicPublicMapOutput {
	return i.ToServiceBasicPublicMapOutputWithContext(context.Background())
}

func (i ServiceBasicPublicMap) ToServiceBasicPublicMapOutputWithContext(ctx context.Context) ServiceBasicPublicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBasicPublicMapOutput)
}

type ServiceBasicPublicOutput struct{ *pulumi.OutputState }

func (ServiceBasicPublicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceBasicPublic)(nil)).Elem()
}

func (o ServiceBasicPublicOutput) ToServiceBasicPublicOutput() ServiceBasicPublicOutput {
	return o
}

func (o ServiceBasicPublicOutput) ToServiceBasicPublicOutputWithContext(ctx context.Context) ServiceBasicPublicOutput {
	return o
}

// The creation time of the resource.
func (o ServiceBasicPublicOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceBasicPublic) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

type ServiceBasicPublicArrayOutput struct{ *pulumi.OutputState }

func (ServiceBasicPublicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceBasicPublic)(nil)).Elem()
}

func (o ServiceBasicPublicArrayOutput) ToServiceBasicPublicArrayOutput() ServiceBasicPublicArrayOutput {
	return o
}

func (o ServiceBasicPublicArrayOutput) ToServiceBasicPublicArrayOutputWithContext(ctx context.Context) ServiceBasicPublicArrayOutput {
	return o
}

func (o ServiceBasicPublicArrayOutput) Index(i pulumi.IntInput) ServiceBasicPublicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceBasicPublic {
		return vs[0].([]*ServiceBasicPublic)[vs[1].(int)]
	}).(ServiceBasicPublicOutput)
}

type ServiceBasicPublicMapOutput struct{ *pulumi.OutputState }

func (ServiceBasicPublicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceBasicPublic)(nil)).Elem()
}

func (o ServiceBasicPublicMapOutput) ToServiceBasicPublicMapOutput() ServiceBasicPublicMapOutput {
	return o
}

func (o ServiceBasicPublicMapOutput) ToServiceBasicPublicMapOutputWithContext(ctx context.Context) ServiceBasicPublicMapOutput {
	return o
}

func (o ServiceBasicPublicMapOutput) MapIndex(k pulumi.StringInput) ServiceBasicPublicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceBasicPublic {
		return vs[0].(map[string]*ServiceBasicPublic)[vs[1].(string)]
	}).(ServiceBasicPublicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBasicPublicInput)(nil)).Elem(), &ServiceBasicPublic{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBasicPublicArrayInput)(nil)).Elem(), ServiceBasicPublicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceBasicPublicMapInput)(nil)).Elem(), ServiceBasicPublicMap{})
	pulumi.RegisterOutputType(ServiceBasicPublicOutput{})
	pulumi.RegisterOutputType(ServiceBasicPublicArrayOutput{})
	pulumi.RegisterOutputType(ServiceBasicPublicMapOutput{})
}
