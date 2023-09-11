// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Ddos Basic Threshold resource.
//
// For information about Ddos Basic Threshold and how to use it, see [What is Threshold](https://www.alibabacloud.com/help/en/ddos-protection/latest/describe-ip-ddosthreshold).
//
// > **NOTE:** Available since v1.183.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ddos"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(defaultZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				Owners:    pulumi.StringRef("system"),
//				NameRegex: pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				Description: pulumi.String("New security group"),
//				VpcId:       defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := ecs.NewInstance(ctx, "defaultInstance", &ecs.InstanceArgs{
//				AvailabilityZone:        *pulumi.String(defaultZones.Zones[0].Id),
//				InstanceName:            pulumi.String(name),
//				HostName:                pulumi.String(name),
//				InternetMaxBandwidthOut: pulumi.Int(10),
//				ImageId:                 *pulumi.String(defaultImages.Images[0].Id),
//				InstanceType:            *pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				SecurityGroups: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//				VswitchId: defaultSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ddos.NewBasicThreshold(ctx, "example", &ddos.BasicThresholdArgs{
//				Pps:          pulumi.Int(60000),
//				Bps:          pulumi.Int(100),
//				InternetIp:   defaultInstance.PublicIp,
//				InstanceId:   defaultInstance.ID(),
//				InstanceType: pulumi.String("ecs"),
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
// Ddos Basic Threshold can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ddos/basicThreshold:BasicThreshold example <instance_type>:<instance_id>:<internet_ip>
//
// ```
type BasicThreshold struct {
	pulumi.CustomResourceState

	// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
	Bps pulumi.IntOutput `pulumi:"bps"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The IP address of the public IP address asset.
	InternetIp pulumi.StringOutput `pulumi:"internetIp"`
	// Maximum flow cleaning threshold. Unit: Mbps.
	MaxBps pulumi.IntOutput `pulumi:"maxBps"`
	// The maximum number of messages cleaning threshold. Unit: pps.
	MaxPps pulumi.IntOutput `pulumi:"maxPps"`
	// The current message number cleaning threshold. Unit: pps.
	Pps pulumi.IntOutput `pulumi:"pps"`
}

// NewBasicThreshold registers a new resource with the given unique name, arguments, and options.
func NewBasicThreshold(ctx *pulumi.Context,
	name string, args *BasicThresholdArgs, opts ...pulumi.ResourceOption) (*BasicThreshold, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bps == nil {
		return nil, errors.New("invalid value for required argument 'Bps'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.InternetIp == nil {
		return nil, errors.New("invalid value for required argument 'InternetIp'")
	}
	if args.Pps == nil {
		return nil, errors.New("invalid value for required argument 'Pps'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BasicThreshold
	err := ctx.RegisterResource("alicloud:ddos/basicThreshold:BasicThreshold", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasicThreshold gets an existing BasicThreshold resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasicThreshold(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasicThresholdState, opts ...pulumi.ResourceOption) (*BasicThreshold, error) {
	var resource BasicThreshold
	err := ctx.ReadResource("alicloud:ddos/basicThreshold:BasicThreshold", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasicThreshold resources.
type basicThresholdState struct {
	// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
	Bps *int `pulumi:"bps"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
	InstanceType *string `pulumi:"instanceType"`
	// The IP address of the public IP address asset.
	InternetIp *string `pulumi:"internetIp"`
	// Maximum flow cleaning threshold. Unit: Mbps.
	MaxBps *int `pulumi:"maxBps"`
	// The maximum number of messages cleaning threshold. Unit: pps.
	MaxPps *int `pulumi:"maxPps"`
	// The current message number cleaning threshold. Unit: pps.
	Pps *int `pulumi:"pps"`
}

type BasicThresholdState struct {
	// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
	Bps pulumi.IntPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
	InstanceType pulumi.StringPtrInput
	// The IP address of the public IP address asset.
	InternetIp pulumi.StringPtrInput
	// Maximum flow cleaning threshold. Unit: Mbps.
	MaxBps pulumi.IntPtrInput
	// The maximum number of messages cleaning threshold. Unit: pps.
	MaxPps pulumi.IntPtrInput
	// The current message number cleaning threshold. Unit: pps.
	Pps pulumi.IntPtrInput
}

func (BasicThresholdState) ElementType() reflect.Type {
	return reflect.TypeOf((*basicThresholdState)(nil)).Elem()
}

type basicThresholdArgs struct {
	// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
	Bps int `pulumi:"bps"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
	InstanceType string `pulumi:"instanceType"`
	// The IP address of the public IP address asset.
	InternetIp string `pulumi:"internetIp"`
	// The current message number cleaning threshold. Unit: pps.
	Pps int `pulumi:"pps"`
}

// The set of arguments for constructing a BasicThreshold resource.
type BasicThresholdArgs struct {
	// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
	Bps pulumi.IntInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
	InstanceType pulumi.StringInput
	// The IP address of the public IP address asset.
	InternetIp pulumi.StringInput
	// The current message number cleaning threshold. Unit: pps.
	Pps pulumi.IntInput
}

func (BasicThresholdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basicThresholdArgs)(nil)).Elem()
}

type BasicThresholdInput interface {
	pulumi.Input

	ToBasicThresholdOutput() BasicThresholdOutput
	ToBasicThresholdOutputWithContext(ctx context.Context) BasicThresholdOutput
}

func (*BasicThreshold) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicThreshold)(nil)).Elem()
}

func (i *BasicThreshold) ToBasicThresholdOutput() BasicThresholdOutput {
	return i.ToBasicThresholdOutputWithContext(context.Background())
}

func (i *BasicThreshold) ToBasicThresholdOutputWithContext(ctx context.Context) BasicThresholdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicThresholdOutput)
}

func (i *BasicThreshold) ToOutput(ctx context.Context) pulumix.Output[*BasicThreshold] {
	return pulumix.Output[*BasicThreshold]{
		OutputState: i.ToBasicThresholdOutputWithContext(ctx).OutputState,
	}
}

// BasicThresholdArrayInput is an input type that accepts BasicThresholdArray and BasicThresholdArrayOutput values.
// You can construct a concrete instance of `BasicThresholdArrayInput` via:
//
//	BasicThresholdArray{ BasicThresholdArgs{...} }
type BasicThresholdArrayInput interface {
	pulumi.Input

	ToBasicThresholdArrayOutput() BasicThresholdArrayOutput
	ToBasicThresholdArrayOutputWithContext(context.Context) BasicThresholdArrayOutput
}

type BasicThresholdArray []BasicThresholdInput

func (BasicThresholdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicThreshold)(nil)).Elem()
}

func (i BasicThresholdArray) ToBasicThresholdArrayOutput() BasicThresholdArrayOutput {
	return i.ToBasicThresholdArrayOutputWithContext(context.Background())
}

func (i BasicThresholdArray) ToBasicThresholdArrayOutputWithContext(ctx context.Context) BasicThresholdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicThresholdArrayOutput)
}

func (i BasicThresholdArray) ToOutput(ctx context.Context) pulumix.Output[[]*BasicThreshold] {
	return pulumix.Output[[]*BasicThreshold]{
		OutputState: i.ToBasicThresholdArrayOutputWithContext(ctx).OutputState,
	}
}

// BasicThresholdMapInput is an input type that accepts BasicThresholdMap and BasicThresholdMapOutput values.
// You can construct a concrete instance of `BasicThresholdMapInput` via:
//
//	BasicThresholdMap{ "key": BasicThresholdArgs{...} }
type BasicThresholdMapInput interface {
	pulumi.Input

	ToBasicThresholdMapOutput() BasicThresholdMapOutput
	ToBasicThresholdMapOutputWithContext(context.Context) BasicThresholdMapOutput
}

type BasicThresholdMap map[string]BasicThresholdInput

func (BasicThresholdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicThreshold)(nil)).Elem()
}

func (i BasicThresholdMap) ToBasicThresholdMapOutput() BasicThresholdMapOutput {
	return i.ToBasicThresholdMapOutputWithContext(context.Background())
}

func (i BasicThresholdMap) ToBasicThresholdMapOutputWithContext(ctx context.Context) BasicThresholdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasicThresholdMapOutput)
}

func (i BasicThresholdMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BasicThreshold] {
	return pulumix.Output[map[string]*BasicThreshold]{
		OutputState: i.ToBasicThresholdMapOutputWithContext(ctx).OutputState,
	}
}

type BasicThresholdOutput struct{ *pulumi.OutputState }

func (BasicThresholdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasicThreshold)(nil)).Elem()
}

func (o BasicThresholdOutput) ToBasicThresholdOutput() BasicThresholdOutput {
	return o
}

func (o BasicThresholdOutput) ToBasicThresholdOutputWithContext(ctx context.Context) BasicThresholdOutput {
	return o
}

func (o BasicThresholdOutput) ToOutput(ctx context.Context) pulumix.Output[*BasicThreshold] {
	return pulumix.Output[*BasicThreshold]{
		OutputState: o.OutputState,
	}
}

// Specifies the traffic scrubbing threshold. Unit: Mbit/s. The traffic scrubbing threshold cannot exceed the peak inbound or outbound Internet traffic, whichever is larger, of the asset.
func (o BasicThresholdOutput) Bps() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.IntOutput { return v.Bps }).(pulumi.IntOutput)
}

// The ID of the instance.
func (o BasicThresholdOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of the Instance. Valid values: `ecs`,`slb`,`eip`.
func (o BasicThresholdOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The IP address of the public IP address asset.
func (o BasicThresholdOutput) InternetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.StringOutput { return v.InternetIp }).(pulumi.StringOutput)
}

// Maximum flow cleaning threshold. Unit: Mbps.
func (o BasicThresholdOutput) MaxBps() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.IntOutput { return v.MaxBps }).(pulumi.IntOutput)
}

// The maximum number of messages cleaning threshold. Unit: pps.
func (o BasicThresholdOutput) MaxPps() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.IntOutput { return v.MaxPps }).(pulumi.IntOutput)
}

// The current message number cleaning threshold. Unit: pps.
func (o BasicThresholdOutput) Pps() pulumi.IntOutput {
	return o.ApplyT(func(v *BasicThreshold) pulumi.IntOutput { return v.Pps }).(pulumi.IntOutput)
}

type BasicThresholdArrayOutput struct{ *pulumi.OutputState }

func (BasicThresholdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BasicThreshold)(nil)).Elem()
}

func (o BasicThresholdArrayOutput) ToBasicThresholdArrayOutput() BasicThresholdArrayOutput {
	return o
}

func (o BasicThresholdArrayOutput) ToBasicThresholdArrayOutputWithContext(ctx context.Context) BasicThresholdArrayOutput {
	return o
}

func (o BasicThresholdArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BasicThreshold] {
	return pulumix.Output[[]*BasicThreshold]{
		OutputState: o.OutputState,
	}
}

func (o BasicThresholdArrayOutput) Index(i pulumi.IntInput) BasicThresholdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BasicThreshold {
		return vs[0].([]*BasicThreshold)[vs[1].(int)]
	}).(BasicThresholdOutput)
}

type BasicThresholdMapOutput struct{ *pulumi.OutputState }

func (BasicThresholdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BasicThreshold)(nil)).Elem()
}

func (o BasicThresholdMapOutput) ToBasicThresholdMapOutput() BasicThresholdMapOutput {
	return o
}

func (o BasicThresholdMapOutput) ToBasicThresholdMapOutputWithContext(ctx context.Context) BasicThresholdMapOutput {
	return o
}

func (o BasicThresholdMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BasicThreshold] {
	return pulumix.Output[map[string]*BasicThreshold]{
		OutputState: o.OutputState,
	}
}

func (o BasicThresholdMapOutput) MapIndex(k pulumi.StringInput) BasicThresholdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BasicThreshold {
		return vs[0].(map[string]*BasicThreshold)[vs[1].(string)]
	}).(BasicThresholdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BasicThresholdInput)(nil)).Elem(), &BasicThreshold{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicThresholdArrayInput)(nil)).Elem(), BasicThresholdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BasicThresholdMapInput)(nil)).Elem(), BasicThresholdMap{})
	pulumi.RegisterOutputType(BasicThresholdOutput{})
	pulumi.RegisterOutputType(BasicThresholdArrayOutput{})
	pulumi.RegisterOutputType(BasicThresholdMapOutput{})
}
