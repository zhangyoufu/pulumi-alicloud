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

// Provides a Global Accelerator (GA) Access Log resource.
//
// For information about Global Accelerator (GA) Access Log and how to use it, see [What is Access Log](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-attachlogstoretoendpointgroup).
//
// > **NOTE:** Available since v1.187.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ga"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_, err := random.NewRandomInteger(ctx, "defaultRandomInteger", &random.RandomIntegerArgs{
//				Max: pulumi.Int(99999),
//				Min: pulumi.Int(10000),
//			})
//			if err != nil {
//				return err
//			}
//			defaultProject, err := log.NewProject(ctx, "defaultProject", nil)
//			if err != nil {
//				return err
//			}
//			defaultStore, err := log.NewStore(ctx, "defaultStore", &log.StoreArgs{
//				Project: defaultProject.Name,
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccelerator, err := ga.NewAccelerator(ctx, "defaultAccelerator", &ga.AcceleratorArgs{
//				Duration:      pulumi.Int(1),
//				AutoUseCoupon: pulumi.Bool(true),
//				Spec:          pulumi.String("2"),
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
//				AcceleratorId:  defaultBandwidthPackageAttachment.AcceleratorId,
//				ClientAffinity: pulumi.String("SOURCE_IP"),
//				Protocol:       pulumi.String("HTTP"),
//				PortRanges: ga.ListenerPortRangeArray{
//					&ga.ListenerPortRangeArgs{
//						FromPort: pulumi.Int(70),
//						ToPort:   pulumi.Int(70),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultEipAddress, err := ecs.NewEipAddress(ctx, "defaultEipAddress", &ecs.EipAddressArgs{
//				Bandwidth:          pulumi.String("10"),
//				InternetChargeType: pulumi.String("PayByBandwidth"),
//				AddressName:        pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEndpointGroup, err := ga.NewEndpointGroup(ctx, "defaultEndpointGroup", &ga.EndpointGroupArgs{
//				AcceleratorId: defaultListener.AcceleratorId,
//				EndpointConfigurations: ga.EndpointGroupEndpointConfigurationArray{
//					&ga.EndpointGroupEndpointConfigurationArgs{
//						Endpoint: defaultEipAddress.IpAddress,
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
//			_, err = ga.NewAccessLog(ctx, "defaultAccessLog", &ga.AccessLogArgs{
//				AcceleratorId:   defaultAccelerator.ID(),
//				ListenerId:      defaultListener.ID(),
//				EndpointGroupId: defaultEndpointGroup.ID(),
//				SlsProjectName:  defaultProject.Name,
//				SlsLogStoreName: defaultStore.Name,
//				SlsRegionId:     pulumi.String(region),
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
// Global Accelerator (GA) Access Log can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ga/accessLog:AccessLog example <accelerator_id>:<listener_id>:<endpoint_group_id>
// ```
type AccessLog struct {
	pulumi.CustomResourceState

	// The ID of the global acceleration instance.
	AcceleratorId pulumi.StringOutput `pulumi:"acceleratorId"`
	// The ID of the endpoint group instance.
	EndpointGroupId pulumi.StringOutput `pulumi:"endpointGroupId"`
	// The ID of the listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The name of the Log Store.
	SlsLogStoreName pulumi.StringOutput `pulumi:"slsLogStoreName"`
	// The name of the Log Service project.
	SlsProjectName pulumi.StringOutput `pulumi:"slsProjectName"`
	// The region ID of the Log Service project.
	SlsRegionId pulumi.StringOutput `pulumi:"slsRegionId"`
	// Whether access log is enabled.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAccessLog registers a new resource with the given unique name, arguments, and options.
func NewAccessLog(ctx *pulumi.Context,
	name string, args *AccessLogArgs, opts ...pulumi.ResourceOption) (*AccessLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceleratorId == nil {
		return nil, errors.New("invalid value for required argument 'AcceleratorId'")
	}
	if args.EndpointGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointGroupId'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.SlsLogStoreName == nil {
		return nil, errors.New("invalid value for required argument 'SlsLogStoreName'")
	}
	if args.SlsProjectName == nil {
		return nil, errors.New("invalid value for required argument 'SlsProjectName'")
	}
	if args.SlsRegionId == nil {
		return nil, errors.New("invalid value for required argument 'SlsRegionId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessLog
	err := ctx.RegisterResource("alicloud:ga/accessLog:AccessLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLog gets an existing AccessLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLogState, opts ...pulumi.ResourceOption) (*AccessLog, error) {
	var resource AccessLog
	err := ctx.ReadResource("alicloud:ga/accessLog:AccessLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLog resources.
type accessLogState struct {
	// The ID of the global acceleration instance.
	AcceleratorId *string `pulumi:"acceleratorId"`
	// The ID of the endpoint group instance.
	EndpointGroupId *string `pulumi:"endpointGroupId"`
	// The ID of the listener.
	ListenerId *string `pulumi:"listenerId"`
	// The name of the Log Store.
	SlsLogStoreName *string `pulumi:"slsLogStoreName"`
	// The name of the Log Service project.
	SlsProjectName *string `pulumi:"slsProjectName"`
	// The region ID of the Log Service project.
	SlsRegionId *string `pulumi:"slsRegionId"`
	// Whether access log is enabled.
	Status *string `pulumi:"status"`
}

type AccessLogState struct {
	// The ID of the global acceleration instance.
	AcceleratorId pulumi.StringPtrInput
	// The ID of the endpoint group instance.
	EndpointGroupId pulumi.StringPtrInput
	// The ID of the listener.
	ListenerId pulumi.StringPtrInput
	// The name of the Log Store.
	SlsLogStoreName pulumi.StringPtrInput
	// The name of the Log Service project.
	SlsProjectName pulumi.StringPtrInput
	// The region ID of the Log Service project.
	SlsRegionId pulumi.StringPtrInput
	// Whether access log is enabled.
	Status pulumi.StringPtrInput
}

func (AccessLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLogState)(nil)).Elem()
}

type accessLogArgs struct {
	// The ID of the global acceleration instance.
	AcceleratorId string `pulumi:"acceleratorId"`
	// The ID of the endpoint group instance.
	EndpointGroupId string `pulumi:"endpointGroupId"`
	// The ID of the listener.
	ListenerId string `pulumi:"listenerId"`
	// The name of the Log Store.
	SlsLogStoreName string `pulumi:"slsLogStoreName"`
	// The name of the Log Service project.
	SlsProjectName string `pulumi:"slsProjectName"`
	// The region ID of the Log Service project.
	SlsRegionId string `pulumi:"slsRegionId"`
}

// The set of arguments for constructing a AccessLog resource.
type AccessLogArgs struct {
	// The ID of the global acceleration instance.
	AcceleratorId pulumi.StringInput
	// The ID of the endpoint group instance.
	EndpointGroupId pulumi.StringInput
	// The ID of the listener.
	ListenerId pulumi.StringInput
	// The name of the Log Store.
	SlsLogStoreName pulumi.StringInput
	// The name of the Log Service project.
	SlsProjectName pulumi.StringInput
	// The region ID of the Log Service project.
	SlsRegionId pulumi.StringInput
}

func (AccessLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLogArgs)(nil)).Elem()
}

type AccessLogInput interface {
	pulumi.Input

	ToAccessLogOutput() AccessLogOutput
	ToAccessLogOutputWithContext(ctx context.Context) AccessLogOutput
}

func (*AccessLog) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLog)(nil)).Elem()
}

func (i *AccessLog) ToAccessLogOutput() AccessLogOutput {
	return i.ToAccessLogOutputWithContext(context.Background())
}

func (i *AccessLog) ToAccessLogOutputWithContext(ctx context.Context) AccessLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLogOutput)
}

// AccessLogArrayInput is an input type that accepts AccessLogArray and AccessLogArrayOutput values.
// You can construct a concrete instance of `AccessLogArrayInput` via:
//
//	AccessLogArray{ AccessLogArgs{...} }
type AccessLogArrayInput interface {
	pulumi.Input

	ToAccessLogArrayOutput() AccessLogArrayOutput
	ToAccessLogArrayOutputWithContext(context.Context) AccessLogArrayOutput
}

type AccessLogArray []AccessLogInput

func (AccessLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLog)(nil)).Elem()
}

func (i AccessLogArray) ToAccessLogArrayOutput() AccessLogArrayOutput {
	return i.ToAccessLogArrayOutputWithContext(context.Background())
}

func (i AccessLogArray) ToAccessLogArrayOutputWithContext(ctx context.Context) AccessLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLogArrayOutput)
}

// AccessLogMapInput is an input type that accepts AccessLogMap and AccessLogMapOutput values.
// You can construct a concrete instance of `AccessLogMapInput` via:
//
//	AccessLogMap{ "key": AccessLogArgs{...} }
type AccessLogMapInput interface {
	pulumi.Input

	ToAccessLogMapOutput() AccessLogMapOutput
	ToAccessLogMapOutputWithContext(context.Context) AccessLogMapOutput
}

type AccessLogMap map[string]AccessLogInput

func (AccessLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLog)(nil)).Elem()
}

func (i AccessLogMap) ToAccessLogMapOutput() AccessLogMapOutput {
	return i.ToAccessLogMapOutputWithContext(context.Background())
}

func (i AccessLogMap) ToAccessLogMapOutputWithContext(ctx context.Context) AccessLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLogMapOutput)
}

type AccessLogOutput struct{ *pulumi.OutputState }

func (AccessLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLog)(nil)).Elem()
}

func (o AccessLogOutput) ToAccessLogOutput() AccessLogOutput {
	return o
}

func (o AccessLogOutput) ToAccessLogOutputWithContext(ctx context.Context) AccessLogOutput {
	return o
}

// The ID of the global acceleration instance.
func (o AccessLogOutput) AcceleratorId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.AcceleratorId }).(pulumi.StringOutput)
}

// The ID of the endpoint group instance.
func (o AccessLogOutput) EndpointGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.EndpointGroupId }).(pulumi.StringOutput)
}

// The ID of the listener.
func (o AccessLogOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The name of the Log Store.
func (o AccessLogOutput) SlsLogStoreName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.SlsLogStoreName }).(pulumi.StringOutput)
}

// The name of the Log Service project.
func (o AccessLogOutput) SlsProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.SlsProjectName }).(pulumi.StringOutput)
}

// The region ID of the Log Service project.
func (o AccessLogOutput) SlsRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.SlsRegionId }).(pulumi.StringOutput)
}

// Whether access log is enabled.
func (o AccessLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AccessLogArrayOutput struct{ *pulumi.OutputState }

func (AccessLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessLog)(nil)).Elem()
}

func (o AccessLogArrayOutput) ToAccessLogArrayOutput() AccessLogArrayOutput {
	return o
}

func (o AccessLogArrayOutput) ToAccessLogArrayOutputWithContext(ctx context.Context) AccessLogArrayOutput {
	return o
}

func (o AccessLogArrayOutput) Index(i pulumi.IntInput) AccessLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessLog {
		return vs[0].([]*AccessLog)[vs[1].(int)]
	}).(AccessLogOutput)
}

type AccessLogMapOutput struct{ *pulumi.OutputState }

func (AccessLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessLog)(nil)).Elem()
}

func (o AccessLogMapOutput) ToAccessLogMapOutput() AccessLogMapOutput {
	return o
}

func (o AccessLogMapOutput) ToAccessLogMapOutputWithContext(ctx context.Context) AccessLogMapOutput {
	return o
}

func (o AccessLogMapOutput) MapIndex(k pulumi.StringInput) AccessLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessLog {
		return vs[0].(map[string]*AccessLog)[vs[1].(string)]
	}).(AccessLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLogInput)(nil)).Elem(), &AccessLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLogArrayInput)(nil)).Elem(), AccessLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLogMapInput)(nil)).Elem(), AccessLogMap{})
	pulumi.RegisterOutputType(AccessLogOutput{})
	pulumi.RegisterOutputType(AccessLogArrayOutput{})
	pulumi.RegisterOutputType(AccessLogMapOutput{})
}
