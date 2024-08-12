// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Dcdn Kv Namespace resource.
//
// For information about Dcdn Kv Namespace and how to use it, see [What is Kv Namespace](https://www.alibabacloud.com/help/en/dcdn/developer-reference/api-dcdn-2018-01-15-putdcdnkvnamespace).
//
// > **NOTE:** Available since v1.198.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dcdn"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dcdn.NewKvNamespace(ctx, "default", &dcdn.KvNamespaceArgs{
//				Description: pulumi.String(name),
//				Namespace:   pulumi.Sprintf("%v-%v", name, _default.Result),
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
// Dcdn Kv Namespace can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dcdn/kvNamespace:KvNamespace example
// ```
type KvNamespace struct {
	pulumi.CustomResourceState

	// Namespace description information
	Description pulumi.StringOutput `pulumi:"description"`
	// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The status of the resource
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewKvNamespace registers a new resource with the given unique name, arguments, and options.
func NewKvNamespace(ctx *pulumi.Context,
	name string, args *KvNamespaceArgs, opts ...pulumi.ResourceOption) (*KvNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KvNamespace
	err := ctx.RegisterResource("alicloud:dcdn/kvNamespace:KvNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKvNamespace gets an existing KvNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKvNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KvNamespaceState, opts ...pulumi.ResourceOption) (*KvNamespace, error) {
	var resource KvNamespace
	err := ctx.ReadResource("alicloud:dcdn/kvNamespace:KvNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KvNamespace resources.
type kvNamespaceState struct {
	// Namespace description information
	Description *string `pulumi:"description"`
	// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
	Namespace *string `pulumi:"namespace"`
	// The status of the resource
	Status *string `pulumi:"status"`
}

type KvNamespaceState struct {
	// Namespace description information
	Description pulumi.StringPtrInput
	// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
	Namespace pulumi.StringPtrInput
	// The status of the resource
	Status pulumi.StringPtrInput
}

func (KvNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*kvNamespaceState)(nil)).Elem()
}

type kvNamespaceArgs struct {
	// Namespace description information
	Description string `pulumi:"description"`
	// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
	Namespace string `pulumi:"namespace"`
}

// The set of arguments for constructing a KvNamespace resource.
type KvNamespaceArgs struct {
	// Namespace description information
	Description pulumi.StringInput
	// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
	Namespace pulumi.StringInput
}

func (KvNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kvNamespaceArgs)(nil)).Elem()
}

type KvNamespaceInput interface {
	pulumi.Input

	ToKvNamespaceOutput() KvNamespaceOutput
	ToKvNamespaceOutputWithContext(ctx context.Context) KvNamespaceOutput
}

func (*KvNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**KvNamespace)(nil)).Elem()
}

func (i *KvNamespace) ToKvNamespaceOutput() KvNamespaceOutput {
	return i.ToKvNamespaceOutputWithContext(context.Background())
}

func (i *KvNamespace) ToKvNamespaceOutputWithContext(ctx context.Context) KvNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KvNamespaceOutput)
}

// KvNamespaceArrayInput is an input type that accepts KvNamespaceArray and KvNamespaceArrayOutput values.
// You can construct a concrete instance of `KvNamespaceArrayInput` via:
//
//	KvNamespaceArray{ KvNamespaceArgs{...} }
type KvNamespaceArrayInput interface {
	pulumi.Input

	ToKvNamespaceArrayOutput() KvNamespaceArrayOutput
	ToKvNamespaceArrayOutputWithContext(context.Context) KvNamespaceArrayOutput
}

type KvNamespaceArray []KvNamespaceInput

func (KvNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KvNamespace)(nil)).Elem()
}

func (i KvNamespaceArray) ToKvNamespaceArrayOutput() KvNamespaceArrayOutput {
	return i.ToKvNamespaceArrayOutputWithContext(context.Background())
}

func (i KvNamespaceArray) ToKvNamespaceArrayOutputWithContext(ctx context.Context) KvNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KvNamespaceArrayOutput)
}

// KvNamespaceMapInput is an input type that accepts KvNamespaceMap and KvNamespaceMapOutput values.
// You can construct a concrete instance of `KvNamespaceMapInput` via:
//
//	KvNamespaceMap{ "key": KvNamespaceArgs{...} }
type KvNamespaceMapInput interface {
	pulumi.Input

	ToKvNamespaceMapOutput() KvNamespaceMapOutput
	ToKvNamespaceMapOutputWithContext(context.Context) KvNamespaceMapOutput
}

type KvNamespaceMap map[string]KvNamespaceInput

func (KvNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KvNamespace)(nil)).Elem()
}

func (i KvNamespaceMap) ToKvNamespaceMapOutput() KvNamespaceMapOutput {
	return i.ToKvNamespaceMapOutputWithContext(context.Background())
}

func (i KvNamespaceMap) ToKvNamespaceMapOutputWithContext(ctx context.Context) KvNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KvNamespaceMapOutput)
}

type KvNamespaceOutput struct{ *pulumi.OutputState }

func (KvNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KvNamespace)(nil)).Elem()
}

func (o KvNamespaceOutput) ToKvNamespaceOutput() KvNamespaceOutput {
	return o
}

func (o KvNamespaceOutput) ToKvNamespaceOutputWithContext(ctx context.Context) KvNamespaceOutput {
	return o
}

// Namespace description information
func (o KvNamespaceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *KvNamespace) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
func (o KvNamespaceOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *KvNamespace) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The status of the resource
func (o KvNamespaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *KvNamespace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type KvNamespaceArrayOutput struct{ *pulumi.OutputState }

func (KvNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KvNamespace)(nil)).Elem()
}

func (o KvNamespaceArrayOutput) ToKvNamespaceArrayOutput() KvNamespaceArrayOutput {
	return o
}

func (o KvNamespaceArrayOutput) ToKvNamespaceArrayOutputWithContext(ctx context.Context) KvNamespaceArrayOutput {
	return o
}

func (o KvNamespaceArrayOutput) Index(i pulumi.IntInput) KvNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KvNamespace {
		return vs[0].([]*KvNamespace)[vs[1].(int)]
	}).(KvNamespaceOutput)
}

type KvNamespaceMapOutput struct{ *pulumi.OutputState }

func (KvNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KvNamespace)(nil)).Elem()
}

func (o KvNamespaceMapOutput) ToKvNamespaceMapOutput() KvNamespaceMapOutput {
	return o
}

func (o KvNamespaceMapOutput) ToKvNamespaceMapOutputWithContext(ctx context.Context) KvNamespaceMapOutput {
	return o
}

func (o KvNamespaceMapOutput) MapIndex(k pulumi.StringInput) KvNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KvNamespace {
		return vs[0].(map[string]*KvNamespace)[vs[1].(string)]
	}).(KvNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KvNamespaceInput)(nil)).Elem(), &KvNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*KvNamespaceArrayInput)(nil)).Elem(), KvNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KvNamespaceMapInput)(nil)).Elem(), KvNamespaceMap{})
	pulumi.RegisterOutputType(KvNamespaceOutput{})
	pulumi.RegisterOutputType(KvNamespaceArrayOutput{})
	pulumi.RegisterOutputType(KvNamespaceMapOutput{})
}
