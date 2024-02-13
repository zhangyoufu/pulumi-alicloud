// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.
//
// For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutermulticastdomain).
//
// > **NOTE:** Available since v1.195.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := cen.NewInstance(ctx, "exampleInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String("tf_example"),
//				Description:     pulumi.String("an example for cen"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouter, err := cen.NewTransitRouter(ctx, "exampleTransitRouter", &cen.TransitRouterArgs{
//				TransitRouterName: pulumi.String("tf_example"),
//				CenId:             exampleInstance.ID(),
//				SupportMulticast:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTransitRouterMulticastDomain(ctx, "exampleTransitRouterMulticastDomain", &cen.TransitRouterMulticastDomainArgs{
//				TransitRouterId:                         exampleTransitRouter.TransitRouterId,
//				TransitRouterMulticastDomainName:        pulumi.String("tf_example"),
//				TransitRouterMulticastDomainDescription: pulumi.String("tf_example"),
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
// Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain example <id>
// ```
type TransitRouterMulticastDomain struct {
	pulumi.CustomResourceState

	// The status of the Transit Router Multicast Domain.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
	// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainDescription pulumi.StringPtrOutput `pulumi:"transitRouterMulticastDomainDescription"`
	// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainName pulumi.StringPtrOutput `pulumi:"transitRouterMulticastDomainName"`
}

// NewTransitRouterMulticastDomain registers a new resource with the given unique name, arguments, and options.
func NewTransitRouterMulticastDomain(ctx *pulumi.Context,
	name string, args *TransitRouterMulticastDomainArgs, opts ...pulumi.ResourceOption) (*TransitRouterMulticastDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitRouterMulticastDomain
	err := ctx.RegisterResource("alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitRouterMulticastDomain gets an existing TransitRouterMulticastDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitRouterMulticastDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitRouterMulticastDomainState, opts ...pulumi.ResourceOption) (*TransitRouterMulticastDomain, error) {
	var resource TransitRouterMulticastDomain
	err := ctx.ReadResource("alicloud:cen/transitRouterMulticastDomain:TransitRouterMulticastDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitRouterMulticastDomain resources.
type transitRouterMulticastDomainState struct {
	// The status of the Transit Router Multicast Domain.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
	// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainDescription *string `pulumi:"transitRouterMulticastDomainDescription"`
	// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainName *string `pulumi:"transitRouterMulticastDomainName"`
}

type TransitRouterMulticastDomainState struct {
	// The status of the Transit Router Multicast Domain.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringPtrInput
	// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainDescription pulumi.StringPtrInput
	// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainName pulumi.StringPtrInput
}

func (TransitRouterMulticastDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterMulticastDomainState)(nil)).Elem()
}

type transitRouterMulticastDomainArgs struct {
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainDescription *string `pulumi:"transitRouterMulticastDomainDescription"`
	// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainName *string `pulumi:"transitRouterMulticastDomainName"`
}

// The set of arguments for constructing a TransitRouterMulticastDomain resource.
type TransitRouterMulticastDomainArgs struct {
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput
	// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainDescription pulumi.StringPtrInput
	// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	TransitRouterMulticastDomainName pulumi.StringPtrInput
}

func (TransitRouterMulticastDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitRouterMulticastDomainArgs)(nil)).Elem()
}

type TransitRouterMulticastDomainInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainOutput() TransitRouterMulticastDomainOutput
	ToTransitRouterMulticastDomainOutputWithContext(ctx context.Context) TransitRouterMulticastDomainOutput
}

func (*TransitRouterMulticastDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterMulticastDomain)(nil)).Elem()
}

func (i *TransitRouterMulticastDomain) ToTransitRouterMulticastDomainOutput() TransitRouterMulticastDomainOutput {
	return i.ToTransitRouterMulticastDomainOutputWithContext(context.Background())
}

func (i *TransitRouterMulticastDomain) ToTransitRouterMulticastDomainOutputWithContext(ctx context.Context) TransitRouterMulticastDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainOutput)
}

// TransitRouterMulticastDomainArrayInput is an input type that accepts TransitRouterMulticastDomainArray and TransitRouterMulticastDomainArrayOutput values.
// You can construct a concrete instance of `TransitRouterMulticastDomainArrayInput` via:
//
//	TransitRouterMulticastDomainArray{ TransitRouterMulticastDomainArgs{...} }
type TransitRouterMulticastDomainArrayInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainArrayOutput() TransitRouterMulticastDomainArrayOutput
	ToTransitRouterMulticastDomainArrayOutputWithContext(context.Context) TransitRouterMulticastDomainArrayOutput
}

type TransitRouterMulticastDomainArray []TransitRouterMulticastDomainInput

func (TransitRouterMulticastDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterMulticastDomain)(nil)).Elem()
}

func (i TransitRouterMulticastDomainArray) ToTransitRouterMulticastDomainArrayOutput() TransitRouterMulticastDomainArrayOutput {
	return i.ToTransitRouterMulticastDomainArrayOutputWithContext(context.Background())
}

func (i TransitRouterMulticastDomainArray) ToTransitRouterMulticastDomainArrayOutputWithContext(ctx context.Context) TransitRouterMulticastDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainArrayOutput)
}

// TransitRouterMulticastDomainMapInput is an input type that accepts TransitRouterMulticastDomainMap and TransitRouterMulticastDomainMapOutput values.
// You can construct a concrete instance of `TransitRouterMulticastDomainMapInput` via:
//
//	TransitRouterMulticastDomainMap{ "key": TransitRouterMulticastDomainArgs{...} }
type TransitRouterMulticastDomainMapInput interface {
	pulumi.Input

	ToTransitRouterMulticastDomainMapOutput() TransitRouterMulticastDomainMapOutput
	ToTransitRouterMulticastDomainMapOutputWithContext(context.Context) TransitRouterMulticastDomainMapOutput
}

type TransitRouterMulticastDomainMap map[string]TransitRouterMulticastDomainInput

func (TransitRouterMulticastDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterMulticastDomain)(nil)).Elem()
}

func (i TransitRouterMulticastDomainMap) ToTransitRouterMulticastDomainMapOutput() TransitRouterMulticastDomainMapOutput {
	return i.ToTransitRouterMulticastDomainMapOutputWithContext(context.Background())
}

func (i TransitRouterMulticastDomainMap) ToTransitRouterMulticastDomainMapOutputWithContext(ctx context.Context) TransitRouterMulticastDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitRouterMulticastDomainMapOutput)
}

type TransitRouterMulticastDomainOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitRouterMulticastDomain)(nil)).Elem()
}

func (o TransitRouterMulticastDomainOutput) ToTransitRouterMulticastDomainOutput() TransitRouterMulticastDomainOutput {
	return o
}

func (o TransitRouterMulticastDomainOutput) ToTransitRouterMulticastDomainOutputWithContext(ctx context.Context) TransitRouterMulticastDomainOutput {
	return o
}

// The status of the Transit Router Multicast Domain.
func (o TransitRouterMulticastDomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomain) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o TransitRouterMulticastDomainOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomain) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the transit router.
func (o TransitRouterMulticastDomainOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomain) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

// The description of the multicast domain. The description must be 0 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
func (o TransitRouterMulticastDomainOutput) TransitRouterMulticastDomainDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomain) pulumi.StringPtrOutput {
		return v.TransitRouterMulticastDomainDescription
	}).(pulumi.StringPtrOutput)
}

// The name of the multicast domain. The name must be 0 to 128 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
func (o TransitRouterMulticastDomainOutput) TransitRouterMulticastDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransitRouterMulticastDomain) pulumi.StringPtrOutput {
		return v.TransitRouterMulticastDomainName
	}).(pulumi.StringPtrOutput)
}

type TransitRouterMulticastDomainArrayOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransitRouterMulticastDomain)(nil)).Elem()
}

func (o TransitRouterMulticastDomainArrayOutput) ToTransitRouterMulticastDomainArrayOutput() TransitRouterMulticastDomainArrayOutput {
	return o
}

func (o TransitRouterMulticastDomainArrayOutput) ToTransitRouterMulticastDomainArrayOutputWithContext(ctx context.Context) TransitRouterMulticastDomainArrayOutput {
	return o
}

func (o TransitRouterMulticastDomainArrayOutput) Index(i pulumi.IntInput) TransitRouterMulticastDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransitRouterMulticastDomain {
		return vs[0].([]*TransitRouterMulticastDomain)[vs[1].(int)]
	}).(TransitRouterMulticastDomainOutput)
}

type TransitRouterMulticastDomainMapOutput struct{ *pulumi.OutputState }

func (TransitRouterMulticastDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransitRouterMulticastDomain)(nil)).Elem()
}

func (o TransitRouterMulticastDomainMapOutput) ToTransitRouterMulticastDomainMapOutput() TransitRouterMulticastDomainMapOutput {
	return o
}

func (o TransitRouterMulticastDomainMapOutput) ToTransitRouterMulticastDomainMapOutputWithContext(ctx context.Context) TransitRouterMulticastDomainMapOutput {
	return o
}

func (o TransitRouterMulticastDomainMapOutput) MapIndex(k pulumi.StringInput) TransitRouterMulticastDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransitRouterMulticastDomain {
		return vs[0].(map[string]*TransitRouterMulticastDomain)[vs[1].(string)]
	}).(TransitRouterMulticastDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainInput)(nil)).Elem(), &TransitRouterMulticastDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainArrayInput)(nil)).Elem(), TransitRouterMulticastDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransitRouterMulticastDomainMapInput)(nil)).Elem(), TransitRouterMulticastDomainMap{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainOutput{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainArrayOutput{})
	pulumi.RegisterOutputType(TransitRouterMulticastDomainMapOutput{})
}
