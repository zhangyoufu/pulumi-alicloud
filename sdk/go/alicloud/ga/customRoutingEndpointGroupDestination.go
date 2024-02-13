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

// Provides a Global Accelerator (GA) Custom Routing Endpoint Group Destination resource.
//
// For information about Global Accelerator (GA) Custom Routing Endpoint Group Destination and how to use it, see [What is Custom Routing Endpoint Group Destination](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createcustomroutingendpointgroupdestinations).
//
// > **NOTE:** Available since v1.197.0.
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
//			defaultAccelerator, err := ga.NewAccelerator(ctx, "defaultAccelerator", &ga.AcceleratorArgs{
//				Duration:      pulumi.Int(1),
//				AutoUseCoupon: pulumi.Bool(true),
//				Spec:          pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBandwidthPackage, err := ga.NewBandwidthPackage(ctx, "defaultBandwidthPackage", &ga.BandwidthPackageArgs{
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
//			defaultBandwidthPackageAttachment, err := ga.NewBandwidthPackageAttachment(ctx, "defaultBandwidthPackageAttachment", &ga.BandwidthPackageAttachmentArgs{
//				AcceleratorId:      defaultAccelerator.ID(),
//				BandwidthPackageId: defaultBandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultListener, err := ga.NewListener(ctx, "defaultListener", &ga.ListenerArgs{
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
//			defaultCustomRoutingEndpointGroup, err := ga.NewCustomRoutingEndpointGroup(ctx, "defaultCustomRoutingEndpointGroup", &ga.CustomRoutingEndpointGroupArgs{
//				AcceleratorId:                  defaultListener.AcceleratorId,
//				ListenerId:                     defaultListener.ID(),
//				EndpointGroupRegion:            pulumi.String(region),
//				CustomRoutingEndpointGroupName: pulumi.String("terraform-example"),
//				Description:                    pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ga.NewCustomRoutingEndpointGroupDestination(ctx, "defaultCustomRoutingEndpointGroupDestination", &ga.CustomRoutingEndpointGroupDestinationArgs{
//				EndpointGroupId: defaultCustomRoutingEndpointGroup.ID(),
//				Protocols: pulumi.StringArray{
//					pulumi.String("TCP"),
//				},
//				FromPort: pulumi.Int(1),
//				ToPort:   pulumi.Int(2),
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
// Global Accelerator (GA) Custom Routing Endpoint Group Destination can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination example <endpoint_group_id>:<custom_routing_endpoint_group_destination_id>
// ```
type CustomRoutingEndpointGroupDestination struct {
	pulumi.CustomResourceState

	// The ID of the GA instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The ID of the Custom Routing Endpoint Group Destination.
	CustomRoutingEndpointGroupDestinationId pulumi.StringOutput `pulumi:"customRoutingEndpointGroupDestinationId"`
	// The ID of the endpoint group.
	EndpointGroupId pulumi.StringOutput `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort pulumi.IntOutput `pulumi:"fromPort"`
	// The ID of the listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols pulumi.StringArrayOutput `pulumi:"protocols"`
	// The status of the Custom Routing Endpoint Group Destination.
	Status pulumi.StringOutput `pulumi:"status"`
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort pulumi.IntOutput `pulumi:"toPort"`
}

// NewCustomRoutingEndpointGroupDestination registers a new resource with the given unique name, arguments, and options.
func NewCustomRoutingEndpointGroupDestination(ctx *pulumi.Context,
	name string, args *CustomRoutingEndpointGroupDestinationArgs, opts ...pulumi.ResourceOption) (*CustomRoutingEndpointGroupDestination, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupId'")
	}
	if args.FromPort == nil {
		return nil, errors.New("invalid value for required argument 'FromPort'")
	}
	if args.Protocols == nil {
		return nil, errors.New("invalid value for required argument 'Protocols'")
	}
	if args.ToPort == nil {
		return nil, errors.New("invalid value for required argument 'ToPort'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomRoutingEndpointGroupDestination
	err := ctx.RegisterResource("alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomRoutingEndpointGroupDestination gets an existing CustomRoutingEndpointGroupDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomRoutingEndpointGroupDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomRoutingEndpointGroupDestinationState, opts ...pulumi.ResourceOption) (*CustomRoutingEndpointGroupDestination, error) {
	var resource CustomRoutingEndpointGroupDestination
	err := ctx.ReadResource("alicloud:ga/customRoutingEndpointGroupDestination:CustomRoutingEndpointGroupDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomRoutingEndpointGroupDestination resources.
type customRoutingEndpointGroupDestinationState struct {
	// The ID of the GA instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The ID of the Custom Routing Endpoint Group Destination.
	CustomRoutingEndpointGroupDestinationId *string `pulumi:"customRoutingEndpointGroupDestinationId"`
	// The ID of the endpoint group.
	EndpointGroupId *string `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort *int `pulumi:"fromPort"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols []string `pulumi:"protocols"`
	// The status of the Custom Routing Endpoint Group Destination.
	Status *string `pulumi:"status"`
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort *int `pulumi:"toPort"`
}

type CustomRoutingEndpointGroupDestinationState struct {
	// The ID of the GA instance.
	AcceleratorId pulumi.StringPtrInput
	// The ID of the Custom Routing Endpoint Group Destination.
	CustomRoutingEndpointGroupDestinationId pulumi.StringPtrInput
	// The ID of the endpoint group.
	EndpointGroupId pulumi.StringPtrInput
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort pulumi.IntPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringPtrInput
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols pulumi.StringArrayInput
	// The status of the Custom Routing Endpoint Group Destination.
	Status pulumi.StringPtrInput
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort pulumi.IntPtrInput
}

func (CustomRoutingEndpointGroupDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customRoutingEndpointGroupDestinationState)(nil)).Elem()
}

type customRoutingEndpointGroupDestinationArgs struct {
	// The ID of the endpoint group.
	EndpointGroupId string `pulumi:"endpointGroupId"`
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort int `pulumi:"fromPort"`
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols []string `pulumi:"protocols"`
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort int `pulumi:"toPort"`
}

// The set of arguments for constructing a CustomRoutingEndpointGroupDestination resource.
type CustomRoutingEndpointGroupDestinationArgs struct {
	// The ID of the endpoint group.
	EndpointGroupId pulumi.StringInput
	// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	FromPort pulumi.IntInput
	// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
	Protocols pulumi.StringArrayInput
	// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
	ToPort pulumi.IntInput
}

func (CustomRoutingEndpointGroupDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customRoutingEndpointGroupDestinationArgs)(nil)).Elem()
}

type CustomRoutingEndpointGroupDestinationInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupDestinationOutput() CustomRoutingEndpointGroupDestinationOutput
	ToCustomRoutingEndpointGroupDestinationOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationOutput
}

func (*CustomRoutingEndpointGroupDestination) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (i *CustomRoutingEndpointGroupDestination) ToCustomRoutingEndpointGroupDestinationOutput() CustomRoutingEndpointGroupDestinationOutput {
	return i.ToCustomRoutingEndpointGroupDestinationOutputWithContext(context.Background())
}

func (i *CustomRoutingEndpointGroupDestination) ToCustomRoutingEndpointGroupDestinationOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupDestinationOutput)
}

// CustomRoutingEndpointGroupDestinationArrayInput is an input type that accepts CustomRoutingEndpointGroupDestinationArray and CustomRoutingEndpointGroupDestinationArrayOutput values.
// You can construct a concrete instance of `CustomRoutingEndpointGroupDestinationArrayInput` via:
//
//	CustomRoutingEndpointGroupDestinationArray{ CustomRoutingEndpointGroupDestinationArgs{...} }
type CustomRoutingEndpointGroupDestinationArrayInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupDestinationArrayOutput() CustomRoutingEndpointGroupDestinationArrayOutput
	ToCustomRoutingEndpointGroupDestinationArrayOutputWithContext(context.Context) CustomRoutingEndpointGroupDestinationArrayOutput
}

type CustomRoutingEndpointGroupDestinationArray []CustomRoutingEndpointGroupDestinationInput

func (CustomRoutingEndpointGroupDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (i CustomRoutingEndpointGroupDestinationArray) ToCustomRoutingEndpointGroupDestinationArrayOutput() CustomRoutingEndpointGroupDestinationArrayOutput {
	return i.ToCustomRoutingEndpointGroupDestinationArrayOutputWithContext(context.Background())
}

func (i CustomRoutingEndpointGroupDestinationArray) ToCustomRoutingEndpointGroupDestinationArrayOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupDestinationArrayOutput)
}

// CustomRoutingEndpointGroupDestinationMapInput is an input type that accepts CustomRoutingEndpointGroupDestinationMap and CustomRoutingEndpointGroupDestinationMapOutput values.
// You can construct a concrete instance of `CustomRoutingEndpointGroupDestinationMapInput` via:
//
//	CustomRoutingEndpointGroupDestinationMap{ "key": CustomRoutingEndpointGroupDestinationArgs{...} }
type CustomRoutingEndpointGroupDestinationMapInput interface {
	pulumi.Input

	ToCustomRoutingEndpointGroupDestinationMapOutput() CustomRoutingEndpointGroupDestinationMapOutput
	ToCustomRoutingEndpointGroupDestinationMapOutputWithContext(context.Context) CustomRoutingEndpointGroupDestinationMapOutput
}

type CustomRoutingEndpointGroupDestinationMap map[string]CustomRoutingEndpointGroupDestinationInput

func (CustomRoutingEndpointGroupDestinationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (i CustomRoutingEndpointGroupDestinationMap) ToCustomRoutingEndpointGroupDestinationMapOutput() CustomRoutingEndpointGroupDestinationMapOutput {
	return i.ToCustomRoutingEndpointGroupDestinationMapOutputWithContext(context.Background())
}

func (i CustomRoutingEndpointGroupDestinationMap) ToCustomRoutingEndpointGroupDestinationMapOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRoutingEndpointGroupDestinationMapOutput)
}

type CustomRoutingEndpointGroupDestinationOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupDestinationOutput) ToCustomRoutingEndpointGroupDestinationOutput() CustomRoutingEndpointGroupDestinationOutput {
	return o
}

func (o CustomRoutingEndpointGroupDestinationOutput) ToCustomRoutingEndpointGroupDestinationOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationOutput {
	return o
}

// The ID of the GA instance.
func (o CustomRoutingEndpointGroupDestinationOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The ID of the Custom Routing Endpoint Group Destination.
func (o CustomRoutingEndpointGroupDestinationOutput) CustomRoutingEndpointGroupDestinationId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringOutput {
		return v.CustomRoutingEndpointGroupDestinationId
	}).(pulumi.StringOutput)
}

// The ID of the endpoint group.
func (o CustomRoutingEndpointGroupDestinationOutput) EndpointGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringOutput { return v.EndpointGroupId }).(pulumi.StringOutput)
}

// The start port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
func (o CustomRoutingEndpointGroupDestinationOutput) FromPort() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.IntOutput { return v.FromPort }).(pulumi.IntOutput)
}

// The ID of the listener.
func (o CustomRoutingEndpointGroupDestinationOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The backend service protocol of the endpoint group. Valid values: `TCP`, `UDP`, `TCP, UDP`.
func (o CustomRoutingEndpointGroupDestinationOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

// The status of the Custom Routing Endpoint Group Destination.
func (o CustomRoutingEndpointGroupDestinationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The end port of the backend service port range of the endpoint group. The `fromPort` value must be smaller than or equal to the `toPort` value. Valid values: `1` to `65499`.
func (o CustomRoutingEndpointGroupDestinationOutput) ToPort() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomRoutingEndpointGroupDestination) pulumi.IntOutput { return v.ToPort }).(pulumi.IntOutput)
}

type CustomRoutingEndpointGroupDestinationArrayOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupDestinationArrayOutput) ToCustomRoutingEndpointGroupDestinationArrayOutput() CustomRoutingEndpointGroupDestinationArrayOutput {
	return o
}

func (o CustomRoutingEndpointGroupDestinationArrayOutput) ToCustomRoutingEndpointGroupDestinationArrayOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationArrayOutput {
	return o
}

func (o CustomRoutingEndpointGroupDestinationArrayOutput) Index(i pulumi.IntInput) CustomRoutingEndpointGroupDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomRoutingEndpointGroupDestination {
		return vs[0].([]*CustomRoutingEndpointGroupDestination)[vs[1].(int)]
	}).(CustomRoutingEndpointGroupDestinationOutput)
}

type CustomRoutingEndpointGroupDestinationMapOutput struct{ *pulumi.OutputState }

func (CustomRoutingEndpointGroupDestinationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomRoutingEndpointGroupDestination)(nil)).Elem()
}

func (o CustomRoutingEndpointGroupDestinationMapOutput) ToCustomRoutingEndpointGroupDestinationMapOutput() CustomRoutingEndpointGroupDestinationMapOutput {
	return o
}

func (o CustomRoutingEndpointGroupDestinationMapOutput) ToCustomRoutingEndpointGroupDestinationMapOutputWithContext(ctx context.Context) CustomRoutingEndpointGroupDestinationMapOutput {
	return o
}

func (o CustomRoutingEndpointGroupDestinationMapOutput) MapIndex(k pulumi.StringInput) CustomRoutingEndpointGroupDestinationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomRoutingEndpointGroupDestination {
		return vs[0].(map[string]*CustomRoutingEndpointGroupDestination)[vs[1].(string)]
	}).(CustomRoutingEndpointGroupDestinationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupDestinationInput)(nil)).Elem(), &CustomRoutingEndpointGroupDestination{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupDestinationArrayInput)(nil)).Elem(), CustomRoutingEndpointGroupDestinationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomRoutingEndpointGroupDestinationMapInput)(nil)).Elem(), CustomRoutingEndpointGroupDestinationMap{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupDestinationOutput{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupDestinationArrayOutput{})
	pulumi.RegisterOutputType(CustomRoutingEndpointGroupDestinationMapOutput{})
}
