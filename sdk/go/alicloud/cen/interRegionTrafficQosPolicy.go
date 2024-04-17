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

// Provides a Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy resource.
//
// For information about Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy and how to use it, see [What is Inter Region Traffic Qos Policy](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createceninterregiontrafficqospolicy).
//
// > **NOTE:** Available since v1.195.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewInstance(ctx, "default", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String("tf-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackage, err := cen.NewBandwidthPackage(ctx, "default", &cen.BandwidthPackageArgs{
//				Bandwidth:           pulumi.Int(5),
//				GeographicRegionAId: pulumi.String("China"),
//				GeographicRegionBId: pulumi.String("China"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackageAttachment, err := cen.NewBandwidthPackageAttachment(ctx, "default", &cen.BandwidthPackageAttachmentArgs{
//				InstanceId:         _default.ID(),
//				BandwidthPackageId: defaultBandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			hz, err := cen.NewTransitRouter(ctx, "hz", &cen.TransitRouterArgs{
//				CenId: defaultBandwidthPackageAttachment.InstanceId,
//			})
//			if err != nil {
//				return err
//			}
//			bj, err := cen.NewTransitRouter(ctx, "bj", &cen.TransitRouterArgs{
//				CenId: hz.CenId,
//			})
//			if err != nil {
//				return err
//			}
//			defaultTransitRouterPeerAttachment, err := cen.NewTransitRouterPeerAttachment(ctx, "default", &cen.TransitRouterPeerAttachmentArgs{
//				CenId:                     _default.ID(),
//				TransitRouterId:           hz.TransitRouterId,
//				PeerTransitRouterRegionId: pulumi.String("cn-beijing"),
//				PeerTransitRouterId:       bj.TransitRouterId,
//				CenBandwidthPackageId:     defaultBandwidthPackageAttachment.BandwidthPackageId,
//				Bandwidth:                 pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewInterRegionTrafficQosPolicy(ctx, "default", &cen.InterRegionTrafficQosPolicyArgs{
//				TransitRouterId:                        hz.TransitRouterId,
//				TransitRouterAttachmentId:              defaultTransitRouterPeerAttachment.TransitRouterAttachmentId,
//				InterRegionTrafficQosPolicyName:        pulumi.String("tf-example-name"),
//				InterRegionTrafficQosPolicyDescription: pulumi.String("tf-example-description"),
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
// Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy example <id>
// ```
type InterRegionTrafficQosPolicy struct {
	pulumi.CustomResourceState

	// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	InterRegionTrafficQosPolicyDescription pulumi.StringPtrOutput `pulumi:"interRegionTrafficQosPolicyDescription"`
	// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	InterRegionTrafficQosPolicyName pulumi.StringPtrOutput `pulumi:"interRegionTrafficQosPolicyName"`
	// The status of the Inter Region Traffic Qos Policy.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the inter-region connection.
	TransitRouterAttachmentId pulumi.StringOutput `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
}

// NewInterRegionTrafficQosPolicy registers a new resource with the given unique name, arguments, and options.
func NewInterRegionTrafficQosPolicy(ctx *pulumi.Context,
	name string, args *InterRegionTrafficQosPolicyArgs, opts ...pulumi.ResourceOption) (*InterRegionTrafficQosPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitRouterAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterAttachmentId'")
	}
	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InterRegionTrafficQosPolicy
	err := ctx.RegisterResource("alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterRegionTrafficQosPolicy gets an existing InterRegionTrafficQosPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterRegionTrafficQosPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterRegionTrafficQosPolicyState, opts ...pulumi.ResourceOption) (*InterRegionTrafficQosPolicy, error) {
	var resource InterRegionTrafficQosPolicy
	err := ctx.ReadResource("alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterRegionTrafficQosPolicy resources.
type interRegionTrafficQosPolicyState struct {
	// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	InterRegionTrafficQosPolicyDescription *string `pulumi:"interRegionTrafficQosPolicyDescription"`
	// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	InterRegionTrafficQosPolicyName *string `pulumi:"interRegionTrafficQosPolicyName"`
	// The status of the Inter Region Traffic Qos Policy.
	Status *string `pulumi:"status"`
	// The ID of the inter-region connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

type InterRegionTrafficQosPolicyState struct {
	// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	InterRegionTrafficQosPolicyDescription pulumi.StringPtrInput
	// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	InterRegionTrafficQosPolicyName pulumi.StringPtrInput
	// The status of the Inter Region Traffic Qos Policy.
	Status pulumi.StringPtrInput
	// The ID of the inter-region connection.
	TransitRouterAttachmentId pulumi.StringPtrInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringPtrInput
}

func (InterRegionTrafficQosPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*interRegionTrafficQosPolicyState)(nil)).Elem()
}

type interRegionTrafficQosPolicyArgs struct {
	// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	InterRegionTrafficQosPolicyDescription *string `pulumi:"interRegionTrafficQosPolicyDescription"`
	// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	InterRegionTrafficQosPolicyName *string `pulumi:"interRegionTrafficQosPolicyName"`
	// The ID of the inter-region connection.
	TransitRouterAttachmentId string `pulumi:"transitRouterAttachmentId"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// The set of arguments for constructing a InterRegionTrafficQosPolicy resource.
type InterRegionTrafficQosPolicyArgs struct {
	// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
	InterRegionTrafficQosPolicyDescription pulumi.StringPtrInput
	// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	InterRegionTrafficQosPolicyName pulumi.StringPtrInput
	// The ID of the inter-region connection.
	TransitRouterAttachmentId pulumi.StringInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput
}

func (InterRegionTrafficQosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interRegionTrafficQosPolicyArgs)(nil)).Elem()
}

type InterRegionTrafficQosPolicyInput interface {
	pulumi.Input

	ToInterRegionTrafficQosPolicyOutput() InterRegionTrafficQosPolicyOutput
	ToInterRegionTrafficQosPolicyOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyOutput
}

func (*InterRegionTrafficQosPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (i *InterRegionTrafficQosPolicy) ToInterRegionTrafficQosPolicyOutput() InterRegionTrafficQosPolicyOutput {
	return i.ToInterRegionTrafficQosPolicyOutputWithContext(context.Background())
}

func (i *InterRegionTrafficQosPolicy) ToInterRegionTrafficQosPolicyOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterRegionTrafficQosPolicyOutput)
}

// InterRegionTrafficQosPolicyArrayInput is an input type that accepts InterRegionTrafficQosPolicyArray and InterRegionTrafficQosPolicyArrayOutput values.
// You can construct a concrete instance of `InterRegionTrafficQosPolicyArrayInput` via:
//
//	InterRegionTrafficQosPolicyArray{ InterRegionTrafficQosPolicyArgs{...} }
type InterRegionTrafficQosPolicyArrayInput interface {
	pulumi.Input

	ToInterRegionTrafficQosPolicyArrayOutput() InterRegionTrafficQosPolicyArrayOutput
	ToInterRegionTrafficQosPolicyArrayOutputWithContext(context.Context) InterRegionTrafficQosPolicyArrayOutput
}

type InterRegionTrafficQosPolicyArray []InterRegionTrafficQosPolicyInput

func (InterRegionTrafficQosPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (i InterRegionTrafficQosPolicyArray) ToInterRegionTrafficQosPolicyArrayOutput() InterRegionTrafficQosPolicyArrayOutput {
	return i.ToInterRegionTrafficQosPolicyArrayOutputWithContext(context.Background())
}

func (i InterRegionTrafficQosPolicyArray) ToInterRegionTrafficQosPolicyArrayOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterRegionTrafficQosPolicyArrayOutput)
}

// InterRegionTrafficQosPolicyMapInput is an input type that accepts InterRegionTrafficQosPolicyMap and InterRegionTrafficQosPolicyMapOutput values.
// You can construct a concrete instance of `InterRegionTrafficQosPolicyMapInput` via:
//
//	InterRegionTrafficQosPolicyMap{ "key": InterRegionTrafficQosPolicyArgs{...} }
type InterRegionTrafficQosPolicyMapInput interface {
	pulumi.Input

	ToInterRegionTrafficQosPolicyMapOutput() InterRegionTrafficQosPolicyMapOutput
	ToInterRegionTrafficQosPolicyMapOutputWithContext(context.Context) InterRegionTrafficQosPolicyMapOutput
}

type InterRegionTrafficQosPolicyMap map[string]InterRegionTrafficQosPolicyInput

func (InterRegionTrafficQosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (i InterRegionTrafficQosPolicyMap) ToInterRegionTrafficQosPolicyMapOutput() InterRegionTrafficQosPolicyMapOutput {
	return i.ToInterRegionTrafficQosPolicyMapOutputWithContext(context.Background())
}

func (i InterRegionTrafficQosPolicyMap) ToInterRegionTrafficQosPolicyMapOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterRegionTrafficQosPolicyMapOutput)
}

type InterRegionTrafficQosPolicyOutput struct{ *pulumi.OutputState }

func (InterRegionTrafficQosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (o InterRegionTrafficQosPolicyOutput) ToInterRegionTrafficQosPolicyOutput() InterRegionTrafficQosPolicyOutput {
	return o
}

func (o InterRegionTrafficQosPolicyOutput) ToInterRegionTrafficQosPolicyOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyOutput {
	return o
}

// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
func (o InterRegionTrafficQosPolicyOutput) InterRegionTrafficQosPolicyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterRegionTrafficQosPolicy) pulumi.StringPtrOutput {
		return v.InterRegionTrafficQosPolicyDescription
	}).(pulumi.StringPtrOutput)
}

// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
func (o InterRegionTrafficQosPolicyOutput) InterRegionTrafficQosPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterRegionTrafficQosPolicy) pulumi.StringPtrOutput { return v.InterRegionTrafficQosPolicyName }).(pulumi.StringPtrOutput)
}

// The status of the Inter Region Traffic Qos Policy.
func (o InterRegionTrafficQosPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *InterRegionTrafficQosPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the inter-region connection.
func (o InterRegionTrafficQosPolicyOutput) TransitRouterAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *InterRegionTrafficQosPolicy) pulumi.StringOutput { return v.TransitRouterAttachmentId }).(pulumi.StringOutput)
}

// The ID of the transit router.
func (o InterRegionTrafficQosPolicyOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *InterRegionTrafficQosPolicy) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

type InterRegionTrafficQosPolicyArrayOutput struct{ *pulumi.OutputState }

func (InterRegionTrafficQosPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (o InterRegionTrafficQosPolicyArrayOutput) ToInterRegionTrafficQosPolicyArrayOutput() InterRegionTrafficQosPolicyArrayOutput {
	return o
}

func (o InterRegionTrafficQosPolicyArrayOutput) ToInterRegionTrafficQosPolicyArrayOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyArrayOutput {
	return o
}

func (o InterRegionTrafficQosPolicyArrayOutput) Index(i pulumi.IntInput) InterRegionTrafficQosPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InterRegionTrafficQosPolicy {
		return vs[0].([]*InterRegionTrafficQosPolicy)[vs[1].(int)]
	}).(InterRegionTrafficQosPolicyOutput)
}

type InterRegionTrafficQosPolicyMapOutput struct{ *pulumi.OutputState }

func (InterRegionTrafficQosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterRegionTrafficQosPolicy)(nil)).Elem()
}

func (o InterRegionTrafficQosPolicyMapOutput) ToInterRegionTrafficQosPolicyMapOutput() InterRegionTrafficQosPolicyMapOutput {
	return o
}

func (o InterRegionTrafficQosPolicyMapOutput) ToInterRegionTrafficQosPolicyMapOutputWithContext(ctx context.Context) InterRegionTrafficQosPolicyMapOutput {
	return o
}

func (o InterRegionTrafficQosPolicyMapOutput) MapIndex(k pulumi.StringInput) InterRegionTrafficQosPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InterRegionTrafficQosPolicy {
		return vs[0].(map[string]*InterRegionTrafficQosPolicy)[vs[1].(string)]
	}).(InterRegionTrafficQosPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterRegionTrafficQosPolicyInput)(nil)).Elem(), &InterRegionTrafficQosPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterRegionTrafficQosPolicyArrayInput)(nil)).Elem(), InterRegionTrafficQosPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterRegionTrafficQosPolicyMapInput)(nil)).Elem(), InterRegionTrafficQosPolicyMap{})
	pulumi.RegisterOutputType(InterRegionTrafficQosPolicyOutput{})
	pulumi.RegisterOutputType(InterRegionTrafficQosPolicyArrayOutput{})
	pulumi.RegisterOutputType(InterRegionTrafficQosPolicyMapOutput{})
}
