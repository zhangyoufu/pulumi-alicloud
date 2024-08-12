// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudmonitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Monitor Service Hybrid Double Write resource.
//
// For information about Cloud Monitor Service Hybrid Double Write and how to use it, see [What is Hybrid Double Write](https://next.api.alibabacloud.com/document/Cms/2018-03-08/CreateHybridDoubleWrite).
//
// > **NOTE:** Available since v1.210.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudmonitor"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			source, err := cms.NewNamespace(ctx, "source", &cms.NamespaceArgs{
//				Namespace: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultNamespace, err := cms.NewNamespace(ctx, "default", &cms.NamespaceArgs{
//				Namespace: pulumi.Sprintf("%v-source", name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudmonitor.NewServiceHybridDoubleWrite(ctx, "default", &cloudmonitor.ServiceHybridDoubleWriteArgs{
//				SourceNamespace: source.ID(),
//				SourceUserId:    pulumi.String(_default.Id),
//				Namespace:       defaultNamespace.ID(),
//				UserId:          pulumi.String(_default.Id),
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
// Cloud Monitor Service Hybrid Double Write can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudmonitor/serviceHybridDoubleWrite:ServiceHybridDoubleWrite example <source_namespace>:<source_user_id>
// ```
type ServiceHybridDoubleWrite struct {
	pulumi.CustomResourceState

	// Target Namespace.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// Source Namespace.
	SourceNamespace pulumi.StringOutput `pulumi:"sourceNamespace"`
	// Source UserId.
	SourceUserId pulumi.StringOutput `pulumi:"sourceUserId"`
	// Target UserId.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewServiceHybridDoubleWrite registers a new resource with the given unique name, arguments, and options.
func NewServiceHybridDoubleWrite(ctx *pulumi.Context,
	name string, args *ServiceHybridDoubleWriteArgs, opts ...pulumi.ResourceOption) (*ServiceHybridDoubleWrite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.SourceNamespace == nil {
		return nil, errors.New("invalid value for required argument 'SourceNamespace'")
	}
	if args.SourceUserId == nil {
		return nil, errors.New("invalid value for required argument 'SourceUserId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceHybridDoubleWrite
	err := ctx.RegisterResource("alicloud:cloudmonitor/serviceHybridDoubleWrite:ServiceHybridDoubleWrite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceHybridDoubleWrite gets an existing ServiceHybridDoubleWrite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceHybridDoubleWrite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceHybridDoubleWriteState, opts ...pulumi.ResourceOption) (*ServiceHybridDoubleWrite, error) {
	var resource ServiceHybridDoubleWrite
	err := ctx.ReadResource("alicloud:cloudmonitor/serviceHybridDoubleWrite:ServiceHybridDoubleWrite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceHybridDoubleWrite resources.
type serviceHybridDoubleWriteState struct {
	// Target Namespace.
	Namespace *string `pulumi:"namespace"`
	// Source Namespace.
	SourceNamespace *string `pulumi:"sourceNamespace"`
	// Source UserId.
	SourceUserId *string `pulumi:"sourceUserId"`
	// Target UserId.
	UserId *string `pulumi:"userId"`
}

type ServiceHybridDoubleWriteState struct {
	// Target Namespace.
	Namespace pulumi.StringPtrInput
	// Source Namespace.
	SourceNamespace pulumi.StringPtrInput
	// Source UserId.
	SourceUserId pulumi.StringPtrInput
	// Target UserId.
	UserId pulumi.StringPtrInput
}

func (ServiceHybridDoubleWriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceHybridDoubleWriteState)(nil)).Elem()
}

type serviceHybridDoubleWriteArgs struct {
	// Target Namespace.
	Namespace string `pulumi:"namespace"`
	// Source Namespace.
	SourceNamespace string `pulumi:"sourceNamespace"`
	// Source UserId.
	SourceUserId string `pulumi:"sourceUserId"`
	// Target UserId.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ServiceHybridDoubleWrite resource.
type ServiceHybridDoubleWriteArgs struct {
	// Target Namespace.
	Namespace pulumi.StringInput
	// Source Namespace.
	SourceNamespace pulumi.StringInput
	// Source UserId.
	SourceUserId pulumi.StringInput
	// Target UserId.
	UserId pulumi.StringInput
}

func (ServiceHybridDoubleWriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceHybridDoubleWriteArgs)(nil)).Elem()
}

type ServiceHybridDoubleWriteInput interface {
	pulumi.Input

	ToServiceHybridDoubleWriteOutput() ServiceHybridDoubleWriteOutput
	ToServiceHybridDoubleWriteOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteOutput
}

func (*ServiceHybridDoubleWrite) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHybridDoubleWrite)(nil)).Elem()
}

func (i *ServiceHybridDoubleWrite) ToServiceHybridDoubleWriteOutput() ServiceHybridDoubleWriteOutput {
	return i.ToServiceHybridDoubleWriteOutputWithContext(context.Background())
}

func (i *ServiceHybridDoubleWrite) ToServiceHybridDoubleWriteOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHybridDoubleWriteOutput)
}

// ServiceHybridDoubleWriteArrayInput is an input type that accepts ServiceHybridDoubleWriteArray and ServiceHybridDoubleWriteArrayOutput values.
// You can construct a concrete instance of `ServiceHybridDoubleWriteArrayInput` via:
//
//	ServiceHybridDoubleWriteArray{ ServiceHybridDoubleWriteArgs{...} }
type ServiceHybridDoubleWriteArrayInput interface {
	pulumi.Input

	ToServiceHybridDoubleWriteArrayOutput() ServiceHybridDoubleWriteArrayOutput
	ToServiceHybridDoubleWriteArrayOutputWithContext(context.Context) ServiceHybridDoubleWriteArrayOutput
}

type ServiceHybridDoubleWriteArray []ServiceHybridDoubleWriteInput

func (ServiceHybridDoubleWriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceHybridDoubleWrite)(nil)).Elem()
}

func (i ServiceHybridDoubleWriteArray) ToServiceHybridDoubleWriteArrayOutput() ServiceHybridDoubleWriteArrayOutput {
	return i.ToServiceHybridDoubleWriteArrayOutputWithContext(context.Background())
}

func (i ServiceHybridDoubleWriteArray) ToServiceHybridDoubleWriteArrayOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHybridDoubleWriteArrayOutput)
}

// ServiceHybridDoubleWriteMapInput is an input type that accepts ServiceHybridDoubleWriteMap and ServiceHybridDoubleWriteMapOutput values.
// You can construct a concrete instance of `ServiceHybridDoubleWriteMapInput` via:
//
//	ServiceHybridDoubleWriteMap{ "key": ServiceHybridDoubleWriteArgs{...} }
type ServiceHybridDoubleWriteMapInput interface {
	pulumi.Input

	ToServiceHybridDoubleWriteMapOutput() ServiceHybridDoubleWriteMapOutput
	ToServiceHybridDoubleWriteMapOutputWithContext(context.Context) ServiceHybridDoubleWriteMapOutput
}

type ServiceHybridDoubleWriteMap map[string]ServiceHybridDoubleWriteInput

func (ServiceHybridDoubleWriteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceHybridDoubleWrite)(nil)).Elem()
}

func (i ServiceHybridDoubleWriteMap) ToServiceHybridDoubleWriteMapOutput() ServiceHybridDoubleWriteMapOutput {
	return i.ToServiceHybridDoubleWriteMapOutputWithContext(context.Background())
}

func (i ServiceHybridDoubleWriteMap) ToServiceHybridDoubleWriteMapOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceHybridDoubleWriteMapOutput)
}

type ServiceHybridDoubleWriteOutput struct{ *pulumi.OutputState }

func (ServiceHybridDoubleWriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceHybridDoubleWrite)(nil)).Elem()
}

func (o ServiceHybridDoubleWriteOutput) ToServiceHybridDoubleWriteOutput() ServiceHybridDoubleWriteOutput {
	return o
}

func (o ServiceHybridDoubleWriteOutput) ToServiceHybridDoubleWriteOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteOutput {
	return o
}

// Target Namespace.
func (o ServiceHybridDoubleWriteOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceHybridDoubleWrite) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// Source Namespace.
func (o ServiceHybridDoubleWriteOutput) SourceNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceHybridDoubleWrite) pulumi.StringOutput { return v.SourceNamespace }).(pulumi.StringOutput)
}

// Source UserId.
func (o ServiceHybridDoubleWriteOutput) SourceUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceHybridDoubleWrite) pulumi.StringOutput { return v.SourceUserId }).(pulumi.StringOutput)
}

// Target UserId.
func (o ServiceHybridDoubleWriteOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceHybridDoubleWrite) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type ServiceHybridDoubleWriteArrayOutput struct{ *pulumi.OutputState }

func (ServiceHybridDoubleWriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceHybridDoubleWrite)(nil)).Elem()
}

func (o ServiceHybridDoubleWriteArrayOutput) ToServiceHybridDoubleWriteArrayOutput() ServiceHybridDoubleWriteArrayOutput {
	return o
}

func (o ServiceHybridDoubleWriteArrayOutput) ToServiceHybridDoubleWriteArrayOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteArrayOutput {
	return o
}

func (o ServiceHybridDoubleWriteArrayOutput) Index(i pulumi.IntInput) ServiceHybridDoubleWriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceHybridDoubleWrite {
		return vs[0].([]*ServiceHybridDoubleWrite)[vs[1].(int)]
	}).(ServiceHybridDoubleWriteOutput)
}

type ServiceHybridDoubleWriteMapOutput struct{ *pulumi.OutputState }

func (ServiceHybridDoubleWriteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceHybridDoubleWrite)(nil)).Elem()
}

func (o ServiceHybridDoubleWriteMapOutput) ToServiceHybridDoubleWriteMapOutput() ServiceHybridDoubleWriteMapOutput {
	return o
}

func (o ServiceHybridDoubleWriteMapOutput) ToServiceHybridDoubleWriteMapOutputWithContext(ctx context.Context) ServiceHybridDoubleWriteMapOutput {
	return o
}

func (o ServiceHybridDoubleWriteMapOutput) MapIndex(k pulumi.StringInput) ServiceHybridDoubleWriteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceHybridDoubleWrite {
		return vs[0].(map[string]*ServiceHybridDoubleWrite)[vs[1].(string)]
	}).(ServiceHybridDoubleWriteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceHybridDoubleWriteInput)(nil)).Elem(), &ServiceHybridDoubleWrite{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceHybridDoubleWriteArrayInput)(nil)).Elem(), ServiceHybridDoubleWriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceHybridDoubleWriteMapInput)(nil)).Elem(), ServiceHybridDoubleWriteMap{})
	pulumi.RegisterOutputType(ServiceHybridDoubleWriteOutput{})
	pulumi.RegisterOutputType(ServiceHybridDoubleWriteArrayOutput{})
	pulumi.RegisterOutputType(ServiceHybridDoubleWriteMapOutput{})
}
