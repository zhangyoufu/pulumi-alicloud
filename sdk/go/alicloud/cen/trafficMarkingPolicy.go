// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Enterprise Network (CEN) Traffic Marking Policy resource.
//
// For information about Cloud Enterprise Network (CEN) Traffic Marking Policy and how to use it, see [What is Traffic Marking Policy](https://help.aliyun.com/document_detail/419025.html).
//
// > **NOTE:** Available in v1.173.0+.
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
//				CenInstanceName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleTransitRouter, err := cen.NewTransitRouter(ctx, "exampleTransitRouter", &cen.TransitRouterArgs{
//				CenId:             exampleInstance.ID(),
//				TransitRouterName: pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cen.NewTrafficMarkingPolicy(ctx, "exampleTrafficMarkingPolicy", &cen.TrafficMarkingPolicyArgs{
//				MarkingDscp:              pulumi.Int(1),
//				Priority:                 pulumi.Int(1),
//				TrafficMarkingPolicyName: pulumi.String("example_value"),
//				TransitRouterId:          exampleTransitRouter.TransitRouterId,
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
// Cloud Enterprise Network (CEN) Traffic Marking Policy can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy example <transit_router_id>:<traffic_marking_policy_id>
//
// ```
type TrafficMarkingPolicy struct {
	pulumi.CustomResourceState

	// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
	MarkingDscp pulumi.IntOutput `pulumi:"markingDscp"`
	// The Priority of the Traffic Marking Policy. Value range: 1~100.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the Traffic Marking Policy.
	TrafficMarkingPolicyId pulumi.StringOutput `pulumi:"trafficMarkingPolicyId"`
	// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	TrafficMarkingPolicyName pulumi.StringPtrOutput `pulumi:"trafficMarkingPolicyName"`
	// The ID of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
}

// NewTrafficMarkingPolicy registers a new resource with the given unique name, arguments, and options.
func NewTrafficMarkingPolicy(ctx *pulumi.Context,
	name string, args *TrafficMarkingPolicyArgs, opts ...pulumi.ResourceOption) (*TrafficMarkingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MarkingDscp == nil {
		return nil, errors.New("invalid value for required argument 'MarkingDscp'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	var resource TrafficMarkingPolicy
	err := ctx.RegisterResource("alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMarkingPolicy gets an existing TrafficMarkingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMarkingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMarkingPolicyState, opts ...pulumi.ResourceOption) (*TrafficMarkingPolicy, error) {
	var resource TrafficMarkingPolicy
	err := ctx.ReadResource("alicloud:cen/trafficMarkingPolicy:TrafficMarkingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMarkingPolicy resources.
type trafficMarkingPolicyState struct {
	// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
	MarkingDscp *int `pulumi:"markingDscp"`
	// The Priority of the Traffic Marking Policy. Value range: 1~100.
	Priority *int `pulumi:"priority"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the Traffic Marking Policy.
	TrafficMarkingPolicyId *string `pulumi:"trafficMarkingPolicyId"`
	// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	TrafficMarkingPolicyName *string `pulumi:"trafficMarkingPolicyName"`
	// The ID of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

type TrafficMarkingPolicyState struct {
	// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
	MarkingDscp pulumi.IntPtrInput
	// The Priority of the Traffic Marking Policy. Value range: 1~100.
	Priority pulumi.IntPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The ID of the Traffic Marking Policy.
	TrafficMarkingPolicyId pulumi.StringPtrInput
	// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	TrafficMarkingPolicyName pulumi.StringPtrInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringPtrInput
}

func (TrafficMarkingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMarkingPolicyState)(nil)).Elem()
}

type trafficMarkingPolicyArgs struct {
	// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	Description *string `pulumi:"description"`
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
	MarkingDscp int `pulumi:"markingDscp"`
	// The Priority of the Traffic Marking Policy. Value range: 1~100.
	Priority int `pulumi:"priority"`
	// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	TrafficMarkingPolicyName *string `pulumi:"trafficMarkingPolicyName"`
	// The ID of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// The set of arguments for constructing a TrafficMarkingPolicy resource.
type TrafficMarkingPolicyArgs struct {
	// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	Description pulumi.StringPtrInput
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
	MarkingDscp pulumi.IntInput
	// The Priority of the Traffic Marking Policy. Value range: 1~100.
	Priority pulumi.IntInput
	// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
	TrafficMarkingPolicyName pulumi.StringPtrInput
	// The ID of the transit router.
	TransitRouterId pulumi.StringInput
}

func (TrafficMarkingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMarkingPolicyArgs)(nil)).Elem()
}

type TrafficMarkingPolicyInput interface {
	pulumi.Input

	ToTrafficMarkingPolicyOutput() TrafficMarkingPolicyOutput
	ToTrafficMarkingPolicyOutputWithContext(ctx context.Context) TrafficMarkingPolicyOutput
}

func (*TrafficMarkingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMarkingPolicy)(nil)).Elem()
}

func (i *TrafficMarkingPolicy) ToTrafficMarkingPolicyOutput() TrafficMarkingPolicyOutput {
	return i.ToTrafficMarkingPolicyOutputWithContext(context.Background())
}

func (i *TrafficMarkingPolicy) ToTrafficMarkingPolicyOutputWithContext(ctx context.Context) TrafficMarkingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMarkingPolicyOutput)
}

// TrafficMarkingPolicyArrayInput is an input type that accepts TrafficMarkingPolicyArray and TrafficMarkingPolicyArrayOutput values.
// You can construct a concrete instance of `TrafficMarkingPolicyArrayInput` via:
//
//	TrafficMarkingPolicyArray{ TrafficMarkingPolicyArgs{...} }
type TrafficMarkingPolicyArrayInput interface {
	pulumi.Input

	ToTrafficMarkingPolicyArrayOutput() TrafficMarkingPolicyArrayOutput
	ToTrafficMarkingPolicyArrayOutputWithContext(context.Context) TrafficMarkingPolicyArrayOutput
}

type TrafficMarkingPolicyArray []TrafficMarkingPolicyInput

func (TrafficMarkingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMarkingPolicy)(nil)).Elem()
}

func (i TrafficMarkingPolicyArray) ToTrafficMarkingPolicyArrayOutput() TrafficMarkingPolicyArrayOutput {
	return i.ToTrafficMarkingPolicyArrayOutputWithContext(context.Background())
}

func (i TrafficMarkingPolicyArray) ToTrafficMarkingPolicyArrayOutputWithContext(ctx context.Context) TrafficMarkingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMarkingPolicyArrayOutput)
}

// TrafficMarkingPolicyMapInput is an input type that accepts TrafficMarkingPolicyMap and TrafficMarkingPolicyMapOutput values.
// You can construct a concrete instance of `TrafficMarkingPolicyMapInput` via:
//
//	TrafficMarkingPolicyMap{ "key": TrafficMarkingPolicyArgs{...} }
type TrafficMarkingPolicyMapInput interface {
	pulumi.Input

	ToTrafficMarkingPolicyMapOutput() TrafficMarkingPolicyMapOutput
	ToTrafficMarkingPolicyMapOutputWithContext(context.Context) TrafficMarkingPolicyMapOutput
}

type TrafficMarkingPolicyMap map[string]TrafficMarkingPolicyInput

func (TrafficMarkingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMarkingPolicy)(nil)).Elem()
}

func (i TrafficMarkingPolicyMap) ToTrafficMarkingPolicyMapOutput() TrafficMarkingPolicyMapOutput {
	return i.ToTrafficMarkingPolicyMapOutputWithContext(context.Background())
}

func (i TrafficMarkingPolicyMap) ToTrafficMarkingPolicyMapOutputWithContext(ctx context.Context) TrafficMarkingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMarkingPolicyMapOutput)
}

type TrafficMarkingPolicyOutput struct{ *pulumi.OutputState }

func (TrafficMarkingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMarkingPolicy)(nil)).Elem()
}

func (o TrafficMarkingPolicyOutput) ToTrafficMarkingPolicyOutput() TrafficMarkingPolicyOutput {
	return o
}

func (o TrafficMarkingPolicyOutput) ToTrafficMarkingPolicyOutputWithContext(ctx context.Context) TrafficMarkingPolicyOutput {
	return o
}

// The description of the Traffic Marking Policy. The description must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
func (o TrafficMarkingPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dry run.
func (o TrafficMarkingPolicyOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The DSCP(Differentiated Services Code Point) of the Traffic Marking Policy. Value range: 0~63.
func (o TrafficMarkingPolicyOutput) MarkingDscp() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.IntOutput { return v.MarkingDscp }).(pulumi.IntOutput)
}

// The Priority of the Traffic Marking Policy. Value range: 1~100.
func (o TrafficMarkingPolicyOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The status of the resource.
func (o TrafficMarkingPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the Traffic Marking Policy.
func (o TrafficMarkingPolicyOutput) TrafficMarkingPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.StringOutput { return v.TrafficMarkingPolicyId }).(pulumi.StringOutput)
}

// The name of the Traffic Marking Policy. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
func (o TrafficMarkingPolicyOutput) TrafficMarkingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.StringPtrOutput { return v.TrafficMarkingPolicyName }).(pulumi.StringPtrOutput)
}

// The ID of the transit router.
func (o TrafficMarkingPolicyOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMarkingPolicy) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

type TrafficMarkingPolicyArrayOutput struct{ *pulumi.OutputState }

func (TrafficMarkingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMarkingPolicy)(nil)).Elem()
}

func (o TrafficMarkingPolicyArrayOutput) ToTrafficMarkingPolicyArrayOutput() TrafficMarkingPolicyArrayOutput {
	return o
}

func (o TrafficMarkingPolicyArrayOutput) ToTrafficMarkingPolicyArrayOutputWithContext(ctx context.Context) TrafficMarkingPolicyArrayOutput {
	return o
}

func (o TrafficMarkingPolicyArrayOutput) Index(i pulumi.IntInput) TrafficMarkingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMarkingPolicy {
		return vs[0].([]*TrafficMarkingPolicy)[vs[1].(int)]
	}).(TrafficMarkingPolicyOutput)
}

type TrafficMarkingPolicyMapOutput struct{ *pulumi.OutputState }

func (TrafficMarkingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMarkingPolicy)(nil)).Elem()
}

func (o TrafficMarkingPolicyMapOutput) ToTrafficMarkingPolicyMapOutput() TrafficMarkingPolicyMapOutput {
	return o
}

func (o TrafficMarkingPolicyMapOutput) ToTrafficMarkingPolicyMapOutputWithContext(ctx context.Context) TrafficMarkingPolicyMapOutput {
	return o
}

func (o TrafficMarkingPolicyMapOutput) MapIndex(k pulumi.StringInput) TrafficMarkingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMarkingPolicy {
		return vs[0].(map[string]*TrafficMarkingPolicy)[vs[1].(string)]
	}).(TrafficMarkingPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMarkingPolicyInput)(nil)).Elem(), &TrafficMarkingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMarkingPolicyArrayInput)(nil)).Elem(), TrafficMarkingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMarkingPolicyMapInput)(nil)).Elem(), TrafficMarkingPolicyMap{})
	pulumi.RegisterOutputType(TrafficMarkingPolicyOutput{})
	pulumi.RegisterOutputType(TrafficMarkingPolicyArrayOutput{})
	pulumi.RegisterOutputType(TrafficMarkingPolicyMapOutput{})
}
