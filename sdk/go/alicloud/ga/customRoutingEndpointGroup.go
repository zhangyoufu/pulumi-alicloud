// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Global Accelerator (GA) Custom Routing Endpoint Group resource.
//
// For information about Global Accelerator (GA) Custom Routing Endpoint Group and how to use it, see [What is Custom Routing Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createcustomroutingendpointgroups).
//
// > **NOTE:** Available since v1.197.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			region := "cn-hangzhou"
//			if param := cfg.Get("region"); param != "" {
//				region = param
//			}
//			_, err := ga.NewAccelerator(ctx, "default", &ga.AcceleratorArgs{
//				Duration:      pulumi.Int(1),
//				AutoUseCoupon: pulumi.Bool(true),
//				Spec:          pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "default", &ga.BandwidthPackageArgs{
//				Bandwidth:     pulumi.Int(100),
//				Type:          pulumi.String("Basic"),
//				BandwidthType: pulumi.String("Basic"),
//				PaymentType:   pulumi.String("PayAsYouGo"),
//				BillingType:   pulumi.String("PayBy95"),
//				Ratio:         pulumi.Int(30),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "default", &ga.BandwidthPackageAttachmentArgs{
//				AcceleratorId:      _default.ID(),
//				BandwidthPackageId: defaultBandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultListener, err := ga.NewListener(ctx, "default", &ga.ListenerArgs{
//				AcceleratorId: defaultBandwidthPackageAttachment.AcceleratorId,
//				ListenerType:  pulumi.String("CustomRouting"),
//				PortRanges: ga.ListenerPortRangeArray{
//					&ga.ListenerPortRangeArgs{
//						FromPort: pulumi.Int(10000),
//						ToPort:   pulumi.Int(16000),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewCustomRoutingEndpointGroup(ctx, "default", &ga.CustomRoutingEndpointGroupArgs{
//				AcceleratorId:                  defaultListener.AcceleratorId,
//				ListenerId:                     defaultListener.ID(),
//				EndpointGroupRegion:            pulumi.String(region),
//				CustomRoutingEndpointGroupName: pulumi.String("terraform-example"),
//				Description:                    pulumi.String("terraform-example"),
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
// Global Accelerator (GA) Custom Routing Endpoint Group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/customRoutingEndpointGroup:CustomRoutingEndpointGroup example <id>
// ```
type CustomRoutingEndpointGroup struct {
	pulumi.CustomResourceState

	// The ID of the GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The name of the endpoint group.
	CustomRoutingEndpointGroupName pulumi.StringPtrOutput `pulumi:"customRoutingEndpointGroupName"`
	// The description of the endpoint group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the region in which to create the endpoint group.
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The ID of the custom routing listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The status of the Custom Routing Endpoint Group.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewCustomRoutingEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewCustomRoutingEndpointGroup(ctx *pulumi.Context,
	name string, args *CustomRoutingEndpointGroupArgs, opts ...pulumi.ResourceOption) (*CustomRoutingEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointGroupRegion == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupRegion'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomRoutingEndpointGroup
	err := ctx.RegisterResource("alicloud:ga/customRoutingEndpointGroup:CustomRoutingEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomRoutingEndpointGroup gets an existing CustomRoutingEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomRoutingEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomRoutingEndpointGroupState, opts ...pulumi.ResourceOption) (*CustomRoutingEndpointGroup, error) {
	var resource CustomRoutingEndpointGroup
	err := ctx.ReadResource("alicloud:ga/customRoutingEndpointGroup:CustomRoutingEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomRoutingEndpointGroup resources.
type customRoutingEndpointGroupState struct {
	// The ID of the GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The name of the endpoint group.
	CustomRoutingEndpointGroupName *string `pulumi:"customRoutingEndpointGroupName"`
	// The description of the endpoint group.
	Description *string `pulumi:"description"`
	// The ID of the region in which to create the endpoint group.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The ID of the custom routing listener.
	ListenerId *string `pulumi:"listenerId"`
	// The status of the Custom Routing Endpoint Group.
	Status *string `pulumi:"status"`
}

type CustomRoutingEndpointGroupState struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The name of the endpoint group.
	CustomRoutingEndpointGroupName pulumi.StringPtrInput
	// The description of the endpoint group.
	Description pulumi.StringPtrInput
	// The ID of the region in which to create the endpoint group.
	EndpointGroupRegion pulumi.StringPtrInput
	// The ID of the custom routing listener.
	ListenerId pulumi.StringPtrInput
	// The status of the Custom Routing Endpoint Group.
	Status pulumi.StringPtrInput
}

func (CustomRoutingEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*customRoutingEndpointGroupState)(nil)).Elem()
}

type customRoutingEndpointGroupArgs struct {
	// The ID of the GA instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The name of the endpoint group.
	CustomRoutingEndpointGroupName *string `pulumi:"customRoutingEndpointGroupName"`
	// The description of the endpoint group.
	Description *string `pulumi:"description"`
	// The ID of the region in which to create the endpoint group.
	EndpointGroupRegion string `pulumi:"endpointGroupRegion"`
	// The ID of the custom routing listener.
	ListenerId string `pulumi:"listenerId"`
}

// The set of arguments for constructing a CustomRoutingEndpointGroup resource.
type CustomRoutingEndpointGroupArgs struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringInput
	// The name of the endpoint group.
	CustomRoutingEndpointGroupName pulumi.StringPtrInput
	// The description of the endpoint group.
	Description pulumi.StringPtrInput
	// The ID of the region in which to create the endpoint group.
	EndpointGroupRegion pulumi.StringInput
	// The ID of the custom routing listener.
	ListenerId pulumi.StringInput
}

func (CustomRoutingEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customRoutingEndpointGroupArgs)(nil)).Elem()
}

type CustomRoutingEndpointGroupInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupOutput() CustomRoutingEndpointGroupOutput
	ToCustomRoutingEndpointGroupOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupOutput
}

func (*CustomRoutingEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRoutingEndpointGroup)(nil)).Elem()
}

func (i *CustomRoutingEndpointGroup) ToCustomRoutingEndpointGroupOutput() CustomRoutingEndpointGroupOutput {
	return i.ToCustomRoutingEndpointGroupOutputWithContext(context.Background())
}

func (i *CustomRoutingEndpointGroup) ToCustomRoutingEndpointGroupOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupOutput)
}

// CustomRoutingEndpointGroupArrayInput is an input type that accepts CustomRoutingEndpointGroupArray and CustomRoutingEndpointGroupArrayOutput values.
// You can construct a concrete instance of `CustomRoutingEndpointGroupArrayInput` via:
//
//	CustomRoutingEndpointGroupArray{ CustomRoutingEndpointGroupArgs{...} }
type CustomRoutingEndpointGroupArrayInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupArrayOutput() CustomRoutingEndpointGroupArrayOutput
	ToCustomRoutingEndpointGroupArrayOutputWithContext(context.Context) CustomRoutingEndpointGroupArrayOutput
}

type CustomRoutingEndpointGroupArray []CustomRoutingEndpointGroupInput

func (CustomRoutingEndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomRoutingEndpointGroup)(nil)).Elem()
}

func (i CustomRoutingEndpointGroupArray) ToCustomRoutingEndpointGroupArrayOutput() CustomRoutingEndpointGroupArrayOutput {
	return i.ToCustomRoutingEndpointGroupArrayOutputWithContext(context.Background())
}

func (i CustomRoutingEndpointGroupArray) ToCustomRoutingEndpointGroupArrayOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupArrayOutput)
}

// CustomRoutingEndpointGroupMapInput is an input type that accepts CustomRoutingEndpointGroupMap and CustomRoutingEndpointGroupMapOutput values.
// You can construct a concrete instance of `CustomRoutingEndpointGroupMapInput` via:
//
//	CustomRoutingEndpointGroupMap{ "key": CustomRoutingEndpointGroupArgs{...} }
type CustomRoutingEndpointGroupMapInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupMapOutput() CustomRoutingEndpointGroupMapOutput
	ToCustomRoutingEndpointGroupMapOutputWithContext(context.Context) CustomRoutingEndpointGroupMapOutput
}

type CustomRoutingEndpointGroupMap map[string]CustomRoutingEndpointGroupInput

func (CustomRoutingEndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomRoutingEndpointGroup)(nil)).Elem()
}

func (i CustomRoutingEndpointGroupMap) ToCustomRoutingEndpointGroupMapOutput() CustomRoutingEndpointGroupMapOutput {
	return i.ToCustomRoutingEndpointGroupMapOutputWithContext(context.Background())
}

func (i CustomRoutingEndpointGroupMap) ToCustomRoutingEndpointGroupMapOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupMapOutput)
}

type CustomRoutingEndpointGroupOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRoutingEndpointGroup)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupOutput) ToCustomRoutingEndpointGroupOutput() CustomRoutingEndpointGroupOutput {
	return o
}

func (o CustomRoutingEndpointGroupOutput) ToCustomRoutingEndpointGroupOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupOutput {
	return o
}

// The ID of the GA instance.
func (o CustomRoutingEndpointGroupOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The name of the endpoint group.
func (o CustomRoutingEndpointGroupOutput) CustomRoutingEndpointGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringPtrOutput { return v.CustomRoutingEndpointGroupName }).(pulumi.StringPtrOutput)
}

// The description of the endpoint group.
func (o CustomRoutingEndpointGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the region in which to create the endpoint group.
func (o CustomRoutingEndpointGroupOutput) EndpointGroupRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringOutput { return v.EndpointGroupRegion }).(pulumi.StringOutput)
}

// The ID of the custom routing listener.
func (o CustomRoutingEndpointGroupOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The status of the Custom Routing Endpoint Group.
func (o CustomRoutingEndpointGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type CustomRoutingEndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomRoutingEndpointGroup)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupArrayOutput) ToCustomRoutingEndpointGroupArrayOutput() CustomRoutingEndpointGroupArrayOutput {
	return o
}

func (o CustomRoutingEndpointGroupArrayOutput) ToCustomRoutingEndpointGroupArrayOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupArrayOutput {
	return o
}

func (o CustomRoutingEndpointGroupArrayOutput) Index(i pulumi.IntInput) CustomRoutingEndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomRoutingEndpointGroup {
		return vs[0].([]*CustomRoutingEndpointGroup)[vs[1].(int)]
	}).(CustomRoutingEndpointGroupOutput)
}

type CustomRoutingEndpointGroupMapOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomRoutingEndpointGroup)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupMapOutput) ToCustomRoutingEndpointGroupMapOutput() CustomRoutingEndpointGroupMapOutput {
	return o
}

func (o CustomRoutingEndpointGroupMapOutput) ToCustomRoutingEndpointGroupMapOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupMapOutput {
	return o
}

func (o CustomRoutingEndpointGroupMapOutput) MapIndex(k pulumi.StringInput) CustomRoutingEndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomRoutingEndpointGroup {
		return vs[0].(map[string]*CustomRoutingEndpointGroup)[vs[1].(string)]
	}).(CustomRoutingEndpointGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupInput)(nil)).Elem(), &CustomRoutingEndpointGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupArrayInput)(nil)).Elem(), CustomRoutingEndpointGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupMapInput)(nil)).Elem(), CustomRoutingEndpointGroupMap{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupOutput{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupMapOutput{})
}
