// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ga

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Global Accelerator (GA) Endpoint Group resource.
//
// For information about Global Accelerator (GA) Endpoint Group and how to use it, see [What is Endpoint Group](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createendpointgroup).
//
// > **NOTE:** Available since v1.113.0.
//
// > **NOTE:** Listeners that use different protocols support different types of endpoint groups:
// * For a TCP or UDP listener, you can create only one default endpoint group.
// * For an HTTP or HTTPS listener, you can create one default endpoint group and one virtual endpoint group. By default, you can create only one virtual endpoint group.
//   - A default endpoint group refers to the endpoint group that you configure when you create an HTTP or HTTPS listener.
//   - A virtual endpoint group refers to the endpoint group that you can create on the Endpoint Group page after you create a listener.
//
// * After you create a virtual endpoint group for an HTTP or HTTPS listener, you can create a forwarding rule and associate the forwarding rule with the virtual endpoint group. Then, the HTTP or HTTPS listener forwards requests with different destination domain names or paths to the default or virtual endpoint group based on the forwarding rule. This way, you can use one Global Accelerator (GA) instance to accelerate access to multiple domain names or paths. For more information about how to create a forwarding rule, see [Manage forwarding rules](https://www.alibabacloud.com/help/en/doc-detail/204224.htm).
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
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
//				PortRanges: ga.ListenerPortRangeArray{
//					&ga.ListenerPortRangeArgs{
//						FromPort: pulumi.Int(60),
//						ToPort:   pulumi.Int(70),
//					},
//				},
//				ClientAffinity: pulumi.String("SOURCE_IP"),
//				Protocol:       pulumi.String("UDP"),
//			})
//			if err != nil {
//				return err
//			}
//			var defaultEipAddress []*ecs.EipAddress
//			for index := 0; index < 2; index++ {
//				key0 := index
//				_ := index
//				__res, err := ecs.NewEipAddress(ctx, fmt.Sprintf("defaultEipAddress-%v", key0), &ecs.EipAddressArgs{
//					Bandwidth:          pulumi.String("10"),
//					InternetChargeType: pulumi.String("PayByBandwidth"),
//					AddressName:        pulumi.String("terraform-example"),
//				})
//				if err != nil {
//					return err
//				}
//				defaultEipAddress = append(defaultEipAddress, __res)
//			}
//			_, err = ga.NewEndpointGroup(ctx, "defaultEndpointGroup", &ga.EndpointGroupArgs{
//				AcceleratorId: defaultAccelerator.ID(),
//				EndpointConfigurations: ga.EndpointGroupEndpointConfigurationArray{
//					&ga.EndpointGroupEndpointConfigurationArgs{
//						Endpoint: defaultEipAddress[0].IpAddress,
//						Type:     pulumi.String("PublicIp"),
//						Weight:   pulumi.Int(20),
//					},
//					&ga.EndpointGroupEndpointConfigurationArgs{
//						Endpoint: defaultEipAddress[1].IpAddress,
//						Type:     pulumi.String("PublicIp"),
//						Weight:   pulumi.Int(20),
//					},
//				},
//				EndpointGroupRegion: pulumi.String(region),
//				ListenerId:          defaultListener.ID(),
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
// Ga Endpoint Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ga/endpointGroup:EndpointGroup example <id>
//
// ```
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The description of the endpoint group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayOutput `pulumi:"endpointConfigurations"`
	// The ID of the region where the endpoint group is deployed.
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
	// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
	EndpointGroupType pulumi.StringOutput `pulumi:"endpointGroupType"`
	// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
	// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
	EndpointRequestProtocol pulumi.StringPtrOutput `pulumi:"endpointRequestProtocol"`
	// The interval between two consecutive health checks. Unit: seconds.
	HealthCheckIntervalSeconds pulumi.IntPtrOutput `pulumi:"healthCheckIntervalSeconds"`
	// The path specified as the destination of the targets for health checks.
	HealthCheckPath pulumi.StringPtrOutput `pulumi:"healthCheckPath"`
	// The port that is used for health checks.
	HealthCheckPort pulumi.IntPtrOutput `pulumi:"healthCheckPort"`
	// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
	HealthCheckProtocol pulumi.StringPtrOutput `pulumi:"healthCheckProtocol"`
	// The ID of the listener that is associated with the endpoint group.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The name of the endpoint group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
	// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
	PortOverrides EndpointGroupPortOverridesPtrOutput `pulumi:"portOverrides"`
	// The status of the endpoint group.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
	ThresholdCount pulumi.IntOutput `pulumi:"thresholdCount"`
	// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
	TrafficPercentage pulumi.IntPtrOutput `pulumi:"trafficPercentage"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'EndpointConfigurations'")
	}
	if args.EndpointGroupRegion == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupRegion'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointGroup
	err := ctx.RegisterResource("alicloud:ga/endpointGroup:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("alicloud:ga/endpointGroup:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The description of the endpoint group.
	Description *string `pulumi:"description"`
	// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The ID of the region where the endpoint group is deployed.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
	// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
	EndpointGroupType *string `pulumi:"endpointGroupType"`
	// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
	// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
	EndpointRequestProtocol *string `pulumi:"endpointRequestProtocol"`
	// The interval between two consecutive health checks. Unit: seconds.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// The path specified as the destination of the targets for health checks.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that is used for health checks.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The ID of the listener that is associated with the endpoint group.
	ListenerId *string `pulumi:"listenerId"`
	// The name of the endpoint group.
	Name *string `pulumi:"name"`
	// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
	// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
	PortOverrides *EndpointGroupPortOverrides `pulumi:"portOverrides"`
	// The status of the endpoint group.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
	TrafficPercentage *int `pulumi:"trafficPercentage"`
}

type EndpointGroupState struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId pulumi.StringPtrInput
	// The description of the endpoint group.
	Description pulumi.StringPtrInput
	// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The ID of the region where the endpoint group is deployed.
	EndpointGroupRegion pulumi.StringPtrInput
	// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
	// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
	EndpointGroupType pulumi.StringPtrInput
	// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
	// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
	EndpointRequestProtocol pulumi.StringPtrInput
	// The interval between two consecutive health checks. Unit: seconds.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// The path specified as the destination of the targets for health checks.
	HealthCheckPath pulumi.StringPtrInput
	// The port that is used for health checks.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
	HealthCheckProtocol pulumi.StringPtrInput
	// The ID of the listener that is associated with the endpoint group.
	ListenerId pulumi.StringPtrInput
	// The name of the endpoint group.
	Name pulumi.StringPtrInput
	// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
	// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
	PortOverrides EndpointGroupPortOverridesPtrInput
	// The status of the endpoint group.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
	ThresholdCount pulumi.IntPtrInput
	// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
	TrafficPercentage pulumi.IntPtrInput
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The description of the endpoint group.
	Description *string `pulumi:"description"`
	// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The ID of the region where the endpoint group is deployed.
	EndpointGroupRegion string `pulumi:"endpointGroupRegion"`
	// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
	// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
	EndpointGroupType *string `pulumi:"endpointGroupType"`
	// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
	// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
	EndpointRequestProtocol *string `pulumi:"endpointRequestProtocol"`
	// The interval between two consecutive health checks. Unit: seconds.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// The path specified as the destination of the targets for health checks.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that is used for health checks.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The ID of the listener that is associated with the endpoint group.
	ListenerId string `pulumi:"listenerId"`
	// The name of the endpoint group.
	Name *string `pulumi:"name"`
	// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
	// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
	PortOverrides *EndpointGroupPortOverrides `pulumi:"portOverrides"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
	TrafficPercentage *int `pulumi:"trafficPercentage"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The ID of the Global Accelerator instance to which the endpoint group will be added.
	AcceleratorId pulumi.StringInput
	// The description of the endpoint group.
	Description pulumi.StringPtrInput
	// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The ID of the region where the endpoint group is deployed.
	EndpointGroupRegion pulumi.StringInput
	// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
	// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
	EndpointGroupType pulumi.StringPtrInput
	// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
	// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
	EndpointRequestProtocol pulumi.StringPtrInput
	// The interval between two consecutive health checks. Unit: seconds.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// The path specified as the destination of the targets for health checks.
	HealthCheckPath pulumi.StringPtrInput
	// The port that is used for health checks.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
	HealthCheckProtocol pulumi.StringPtrInput
	// The ID of the listener that is associated with the endpoint group.
	ListenerId pulumi.StringInput
	// The name of the endpoint group.
	Name pulumi.StringPtrInput
	// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
	// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
	PortOverrides EndpointGroupPortOverridesPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
	ThresholdCount pulumi.IntPtrInput
	// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
	TrafficPercentage pulumi.IntPtrInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}

type EndpointGroupInput interface {
	pulumi.Input

	ToEndpointGroupOutput() EndpointGroupOutput
	ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput
}

func (*EndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (i *EndpointGroup) ToEndpointGroupOutput() EndpointGroupOutput {
	return i.ToEndpointGroupOutputWithContext(context.Background())
}

func (i *EndpointGroup) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupOutput)
}

func (i *EndpointGroup) ToOutput(ctx context.Context) pulumix.Output[*EndpointGroup] {
	return pulumix.Output[*EndpointGroup]{
		OutputState: i.ToEndpointGroupOutputWithContext(ctx).OutputState,
	}
}

// EndpointGroupArrayInput is an input type that accepts EndpointGroupArray and EndpointGroupArrayOutput values.
// You can construct a concrete instance of `EndpointGroupArrayInput` via:
//
//	EndpointGroupArray{ EndpointGroupArgs{...} }
type EndpointGroupArrayInput interface {
	pulumi.Input

	ToEndpointGroupArrayOutput() EndpointGroupArrayOutput
	ToEndpointGroupArrayOutputWithContext(context.Context) EndpointGroupArrayOutput
}

type EndpointGroupArray []EndpointGroupInput

func (EndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointGroup)(nil)).Elem()
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return i.ToEndpointGroupArrayOutputWithContext(context.Background())
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupArrayOutput)
}

func (i EndpointGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointGroup] {
	return pulumix.Output[[]*EndpointGroup]{
		OutputState: i.ToEndpointGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// EndpointGroupMapInput is an input type that accepts EndpointGroupMap and EndpointGroupMapOutput values.
// You can construct a concrete instance of `EndpointGroupMapInput` via:
//
//	EndpointGroupMap{ "key": EndpointGroupArgs{...} }
type EndpointGroupMapInput interface {
	pulumi.Input

	ToEndpointGroupMapOutput() EndpointGroupMapOutput
	ToEndpointGroupMapOutputWithContext(context.Context) EndpointGroupMapOutput
}

type EndpointGroupMap map[string]EndpointGroupInput

func (EndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointGroup)(nil)).Elem()
}

func (i EndpointGroupMap) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return i.ToEndpointGroupMapOutputWithContext(context.Background())
}

func (i EndpointGroupMap) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupMapOutput)
}

func (i EndpointGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointGroup] {
	return pulumix.Output[map[string]*EndpointGroup]{
		OutputState: i.ToEndpointGroupMapOutputWithContext(ctx).OutputState,
	}
}

type EndpointGroupOutput struct{ *pulumi.OutputState }

func (EndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupOutput) ToEndpointGroupOutput() EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*EndpointGroup] {
	return pulumix.Output[*EndpointGroup]{
		OutputState: o.OutputState,
	}
}

// The ID of the Global Accelerator instance to which the endpoint group will be added.
func (o EndpointGroupOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The description of the endpoint group.
func (o EndpointGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The endpointConfigurations of the endpoint group. See `endpointConfigurations` below.
func (o EndpointGroupOutput) EndpointConfigurations() EndpointGroupEndpointConfigurationArrayOutput {
	return o.ApplyT(func(v *EndpointGroup) EndpointGroupEndpointConfigurationArrayOutput { return v.EndpointConfigurations }).(EndpointGroupEndpointConfigurationArrayOutput)
}

// The ID of the region where the endpoint group is deployed.
func (o EndpointGroupOutput) EndpointGroupRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.EndpointGroupRegion }).(pulumi.StringOutput)
}

// The endpoint group type. Default value: `default`. Valid values: `default`, `virtual`.
// > **NOTE:** Only the listening instance of HTTP or HTTPS protocol supports the creation of virtual terminal node group.
func (o EndpointGroupOutput) EndpointGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.EndpointGroupType }).(pulumi.StringOutput)
}

// The endpoint request protocol. Valid values: `HTTP`, `HTTPS`.
// > **NOTE:** This item is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. For the listening instance of HTTP protocol, the back-end service protocol supports and only supports HTTP.
func (o EndpointGroupOutput) EndpointRequestProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringPtrOutput { return v.EndpointRequestProtocol }).(pulumi.StringPtrOutput)
}

// The interval between two consecutive health checks. Unit: seconds.
func (o EndpointGroupOutput) HealthCheckIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.IntPtrOutput { return v.HealthCheckIntervalSeconds }).(pulumi.IntPtrOutput)
}

// The path specified as the destination of the targets for health checks.
func (o EndpointGroupOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringPtrOutput { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

// The port that is used for health checks.
func (o EndpointGroupOutput) HealthCheckPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.IntPtrOutput { return v.HealthCheckPort }).(pulumi.IntPtrOutput)
}

// The protocol that is used to connect to the targets for health checks. Valid values: `http`, `https`, `tcp`.
func (o EndpointGroupOutput) HealthCheckProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringPtrOutput { return v.HealthCheckProtocol }).(pulumi.StringPtrOutput)
}

// The ID of the listener that is associated with the endpoint group.
func (o EndpointGroupOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The name of the endpoint group.
func (o EndpointGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Mapping between listening port and forwarding port of boarding point. See `portOverrides` below.
// > **NOTE:** Port mapping is only supported when creating terminal node group for listening instance of HTTP or HTTPS protocol. The listening port in the port map must be consistent with the listening port of the current listening instance.
func (o EndpointGroupOutput) PortOverrides() EndpointGroupPortOverridesPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) EndpointGroupPortOverridesPtrOutput { return v.PortOverrides }).(EndpointGroupPortOverridesPtrOutput)
}

// The status of the endpoint group.
func (o EndpointGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o EndpointGroupOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy. Default value: `3`.
func (o EndpointGroupOutput) ThresholdCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.IntOutput { return v.ThresholdCount }).(pulumi.IntOutput)
}

// The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
func (o EndpointGroupOutput) TrafficPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.IntPtrOutput { return v.TrafficPercentage }).(pulumi.IntPtrOutput)
}

type EndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (EndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EndpointGroup] {
	return pulumix.Output[[]*EndpointGroup]{
		OutputState: o.OutputState,
	}
}

func (o EndpointGroupArrayOutput) Index(i pulumi.IntInput) EndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointGroup {
		return vs[0].([]*EndpointGroup)[vs[1].(int)]
	}).(EndpointGroupOutput)
}

type EndpointGroupMapOutput struct{ *pulumi.OutputState }

func (EndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EndpointGroup] {
	return pulumix.Output[map[string]*EndpointGroup]{
		OutputState: o.OutputState,
	}
}

func (o EndpointGroupMapOutput) MapIndex(k pulumi.StringInput) EndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointGroup {
		return vs[0].(map[string]*EndpointGroup)[vs[1].(string)]
	}).(EndpointGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupInput)(nil)).Elem(), &EndpointGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupArrayInput)(nil)).Elem(), EndpointGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupMapInput)(nil)).Elem(), EndpointGroupMap{})
	pulumi.RegisterOutputType(EndpointGroupOutput{})
	pulumi.RegisterOutputType(EndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(EndpointGroupMapOutput{})
}
