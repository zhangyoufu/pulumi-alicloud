// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eipanycast

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Eipanycast Anycast Eip Address Attachment resource.
//
// For information about Eipanycast Anycast Eip Address Attachment and how to use it, see [What is Anycast Eip Address Attachment](https://www.alibabacloud.com/help/en/anycast-eip/latest/api-eipanycast-2020-03-09-associateanycasteipaddress).
//
// > **NOTE:** Available since v1.113.0.
//
// > **NOTE:** The following regions support currently while Slb instance support bound.
// [eu-west-1-gb33-a01,cn-hongkong-am4-c04,ap-southeast-os30-a01,us-west-ot7-a01,ap-south-in73-a01,ap-southeast-my88-a01]
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eipanycast"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			_default, err := slb.GetZones(ctx, &slb.GetZonesArgs{
//				AvailableSlbAddressType: pulumi.StringRef("vpc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.0.0.0/8"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.1.0.0/16"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "default", &slb.ApplicationLoadBalancerArgs{
//				AddressType:      pulumi.String("intranet"),
//				VswitchId:        defaultSwitch.ID(),
//				LoadBalancerName: pulumi.String(name),
//				LoadBalancerSpec: pulumi.String("slb.s1.small"),
//				MasterZoneId:     pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultAnycastEipAddress, err := eipanycast.NewAnycastEipAddress(ctx, "default", &eipanycast.AnycastEipAddressArgs{
//				AnycastEipAddressName: pulumi.String(name),
//				ServiceLocation:       pulumi.String("ChineseMainland"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGetRegions, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = eipanycast.NewAnycastEipAddressAttachment(ctx, "default", &eipanycast.AnycastEipAddressAttachmentArgs{
//				BindInstanceId:       defaultApplicationLoadBalancer.ID(),
//				BindInstanceType:     pulumi.String("SlbInstance"),
//				BindInstanceRegionId: pulumi.String(defaultGetRegions.Regions[0].Id),
//				AnycastId:            defaultAnycastEipAddress.ID(),
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
// # Multiple Usage
//
// > **NOTE:**  Anycast EIP supports binding cloud resource instances in multiple regions. Only one cloud resource instance is supported as the default origin station, and the rest are normal origin stations. When no access point is specified or an access point is added, the access request is forwarded to the default origin by default.  If you are bound for the first time, the Default value of the binding mode is **Default * *. /li> li> If you are not binding for the first time, you can set the binding mode to **Default**, and the new Default origin will take effect. The original Default origin will be changed to a common origin.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eipanycast"
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
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(_default.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultVpc, err := vpc.NewNetwork(ctx, "defaultVpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("192.168.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVsw, err := vpc.NewSwitch(ctx, "defaultVsw", &vpc.SwitchArgs{
//				VpcId:     defaultVpc.ID(),
//				CidrBlock: pulumi.String("192.168.0.0/24"),
//				ZoneId:    pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultuBsECI, err := ecs.NewSecurityGroup(ctx, "defaultuBsECI", &ecs.SecurityGroupArgs{
//				VpcId: defaultVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			default9KDlN7, err := ecs.NewInstance(ctx, "default9KDlN7", &ecs.InstanceArgs{
//				ImageId:      pulumi.String(defaultGetImages.Images[0].Id),
//				InstanceType: pulumi.String(defaultGetInstanceTypes.InstanceTypes[0].Id),
//				InstanceName: pulumi.String(name),
//				SecurityGroups: pulumi.StringArray{
//					defaultuBsECI.ID(),
//				},
//				AvailabilityZone:   defaultVsw.ZoneId,
//				InstanceChargeType: pulumi.String("PostPaid"),
//				SystemDiskCategory: pulumi.String("cloud_efficiency"),
//				VswitchId:          defaultVsw.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultXkpFRs, err := eipanycast.NewAnycastEipAddress(ctx, "defaultXkpFRs", &eipanycast.AnycastEipAddressArgs{
//				ServiceLocation: pulumi.String("ChineseMainland"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultVpc2, err := vpc.NewNetwork(ctx, "defaultVpc2", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(fmt.Sprintf("%v6", name)),
//				CidrBlock: pulumi.String("192.168.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			default2, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			default2GetImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			default2GetInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(default2.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultdsVsw2, err := vpc.NewSwitch(ctx, "defaultdsVsw2", &vpc.SwitchArgs{
//				VpcId:     defaultVpc2.ID(),
//				CidrBlock: pulumi.String("192.168.0.0/24"),
//				ZoneId:    pulumi.String(default2.Zones[1].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultuBsECI2, err := ecs.NewSecurityGroup(ctx, "defaultuBsECI2", &ecs.SecurityGroupArgs{
//				VpcId: defaultVpc2.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEcs2, err := ecs.NewInstance(ctx, "defaultEcs2", &ecs.InstanceArgs{
//				ImageId:      pulumi.String(default2GetImages.Images[0].Id),
//				InstanceType: pulumi.String(default2GetInstanceTypes.InstanceTypes[0].Id),
//				InstanceName: pulumi.String(name),
//				SecurityGroups: pulumi.StringArray{
//					defaultuBsECI2.ID(),
//				},
//				AvailabilityZone:   defaultdsVsw2.ZoneId,
//				InstanceChargeType: pulumi.String("PostPaid"),
//				SystemDiskCategory: pulumi.String("cloud_efficiency"),
//				VswitchId:          defaultdsVsw2.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEfYBJY, err := eipanycast.NewAnycastEipAddressAttachment(ctx, "defaultEfYBJY", &eipanycast.AnycastEipAddressAttachmentArgs{
//				BindInstanceId:       default9KDlN7.NetworkInterfaceId,
//				BindInstanceType:     pulumi.String("NetworkInterface"),
//				BindInstanceRegionId: pulumi.String("cn-beijing"),
//				AnycastId:            defaultXkpFRs.ID(),
//				AssociationMode:      pulumi.String("Default"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eipanycast.NewAnycastEipAddressAttachment(ctx, "normal", &eipanycast.AnycastEipAddressAttachmentArgs{
//				BindInstanceId:       defaultEcs2.NetworkInterfaceId,
//				BindInstanceType:     pulumi.String("NetworkInterface"),
//				BindInstanceRegionId: pulumi.String("cn-hangzhou"),
//				AnycastId:            defaultEfYBJY.AnycastId,
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
// Eipanycast Anycast Eip Address Attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment example <anycast_id>:<bind_instance_id>:<bind_instance_region_id>:<bind_instance_type>
// ```
type AnycastEipAddressAttachment struct {
	pulumi.CustomResourceState

	// The ID of the Anycast EIP instance.
	AnycastId pulumi.StringOutput `pulumi:"anycastId"`
	// Binding mode, value:
	// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
	// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
	AssociationMode pulumi.StringOutput `pulumi:"associationMode"`
	// The ID of the cloud resource instance to be bound.
	BindInstanceId pulumi.StringOutput `pulumi:"bindInstanceId"`
	// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
	BindInstanceRegionId pulumi.StringOutput `pulumi:"bindInstanceRegionId"`
	// The type of the cloud resource instance to be bound. Value:
	// - **SlbInstance**: a private network SLB instance.
	// - **NetworkInterface**: ENI.
	BindInstanceType pulumi.StringOutput `pulumi:"bindInstanceType"`
	// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
	BindTime pulumi.StringOutput `pulumi:"bindTime"`
	// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
	PopLocations AnycastEipAddressAttachmentPopLocationArrayOutput `pulumi:"popLocations"`
	// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAnycastEipAddressAttachment registers a new resource with the given unique name, arguments, and options.
func NewAnycastEipAddressAttachment(ctx *pulumi.Context,
	name string, args *AnycastEipAddressAttachmentArgs, opts ...pulumi.ResourceOption) (*AnycastEipAddressAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnycastId == nil {
		return nil, errors.New("invalid value for required argument 'AnycastId'")
	}
	if args.BindInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceId'")
	}
	if args.BindInstanceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceRegionId'")
	}
	if args.BindInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'BindInstanceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnycastEipAddressAttachment
	err := ctx.RegisterResource("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnycastEipAddressAttachment gets an existing AnycastEipAddressAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnycastEipAddressAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnycastEipAddressAttachmentState, opts ...pulumi.ResourceOption) (*AnycastEipAddressAttachment, error) {
	var resource AnycastEipAddressAttachment
	err := ctx.ReadResource("alicloud:eipanycast/anycastEipAddressAttachment:AnycastEipAddressAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnycastEipAddressAttachment resources.
type anycastEipAddressAttachmentState struct {
	// The ID of the Anycast EIP instance.
	AnycastId *string `pulumi:"anycastId"`
	// Binding mode, value:
	// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
	// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
	AssociationMode *string `pulumi:"associationMode"`
	// The ID of the cloud resource instance to be bound.
	BindInstanceId *string `pulumi:"bindInstanceId"`
	// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
	BindInstanceRegionId *string `pulumi:"bindInstanceRegionId"`
	// The type of the cloud resource instance to be bound. Value:
	// - **SlbInstance**: a private network SLB instance.
	// - **NetworkInterface**: ENI.
	BindInstanceType *string `pulumi:"bindInstanceType"`
	// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
	BindTime *string `pulumi:"bindTime"`
	// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
	PopLocations []AnycastEipAddressAttachmentPopLocation `pulumi:"popLocations"`
	// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
	Status *string `pulumi:"status"`
}

type AnycastEipAddressAttachmentState struct {
	// The ID of the Anycast EIP instance.
	AnycastId pulumi.StringPtrInput
	// Binding mode, value:
	// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
	// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
	AssociationMode pulumi.StringPtrInput
	// The ID of the cloud resource instance to be bound.
	BindInstanceId pulumi.StringPtrInput
	// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
	BindInstanceRegionId pulumi.StringPtrInput
	// The type of the cloud resource instance to be bound. Value:
	// - **SlbInstance**: a private network SLB instance.
	// - **NetworkInterface**: ENI.
	BindInstanceType pulumi.StringPtrInput
	// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
	BindTime pulumi.StringPtrInput
	// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
	PopLocations AnycastEipAddressAttachmentPopLocationArrayInput
	// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
	PrivateIpAddress pulumi.StringPtrInput
	// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
	Status pulumi.StringPtrInput
}

func (AnycastEipAddressAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressAttachmentState)(nil)).Elem()
}

type anycastEipAddressAttachmentArgs struct {
	// The ID of the Anycast EIP instance.
	AnycastId string `pulumi:"anycastId"`
	// Binding mode, value:
	// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
	// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
	AssociationMode *string `pulumi:"associationMode"`
	// The ID of the cloud resource instance to be bound.
	BindInstanceId string `pulumi:"bindInstanceId"`
	// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
	BindInstanceRegionId string `pulumi:"bindInstanceRegionId"`
	// The type of the cloud resource instance to be bound. Value:
	// - **SlbInstance**: a private network SLB instance.
	// - **NetworkInterface**: ENI.
	BindInstanceType string `pulumi:"bindInstanceType"`
	// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
	PopLocations []AnycastEipAddressAttachmentPopLocation `pulumi:"popLocations"`
	// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
}

// The set of arguments for constructing a AnycastEipAddressAttachment resource.
type AnycastEipAddressAttachmentArgs struct {
	// The ID of the Anycast EIP instance.
	AnycastId pulumi.StringInput
	// Binding mode, value:
	// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
	// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
	AssociationMode pulumi.StringPtrInput
	// The ID of the cloud resource instance to be bound.
	BindInstanceId pulumi.StringInput
	// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
	BindInstanceRegionId pulumi.StringInput
	// The type of the cloud resource instance to be bound. Value:
	// - **SlbInstance**: a private network SLB instance.
	// - **NetworkInterface**: ENI.
	BindInstanceType pulumi.StringInput
	// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
	PopLocations AnycastEipAddressAttachmentPopLocationArrayInput
	// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
	PrivateIpAddress pulumi.StringPtrInput
}

func (AnycastEipAddressAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anycastEipAddressAttachmentArgs)(nil)).Elem()
}

type AnycastEipAddressAttachmentInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput
	ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput
}

func (*AnycastEipAddressAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddressAttachment)(nil)).Elem()
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput {
	return i.ToAnycastEipAddressAttachmentOutputWithContext(context.Background())
}

func (i *AnycastEipAddressAttachment) ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentOutput)
}

// AnycastEipAddressAttachmentArrayInput is an input type that accepts AnycastEipAddressAttachmentArray and AnycastEipAddressAttachmentArrayOutput values.
// You can construct a concrete instance of `AnycastEipAddressAttachmentArrayInput` via:
//
//	AnycastEipAddressAttachmentArray{ AnycastEipAddressAttachmentArgs{...} }
type AnycastEipAddressAttachmentArrayInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput
	ToAnycastEipAddressAttachmentArrayOutputWithContext(context.Context) AnycastEipAddressAttachmentArrayOutput
}

type AnycastEipAddressAttachmentArray []AnycastEipAddressAttachmentInput

func (AnycastEipAddressAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnycastEipAddressAttachment)(nil)).Elem()
}

func (i AnycastEipAddressAttachmentArray) ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput {
	return i.ToAnycastEipAddressAttachmentArrayOutputWithContext(context.Background())
}

func (i AnycastEipAddressAttachmentArray) ToAnycastEipAddressAttachmentArrayOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentArrayOutput)
}

// AnycastEipAddressAttachmentMapInput is an input type that accepts AnycastEipAddressAttachmentMap and AnycastEipAddressAttachmentMapOutput values.
// You can construct a concrete instance of `AnycastEipAddressAttachmentMapInput` via:
//
//	AnycastEipAddressAttachmentMap{ "key": AnycastEipAddressAttachmentArgs{...} }
type AnycastEipAddressAttachmentMapInput interface {
	pulumi.Input

	ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput
	ToAnycastEipAddressAttachmentMapOutputWithContext(context.Context) AnycastEipAddressAttachmentMapOutput
}

type AnycastEipAddressAttachmentMap map[string]AnycastEipAddressAttachmentInput

func (AnycastEipAddressAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnycastEipAddressAttachment)(nil)).Elem()
}

func (i AnycastEipAddressAttachmentMap) ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput {
	return i.ToAnycastEipAddressAttachmentMapOutputWithContext(context.Background())
}

func (i AnycastEipAddressAttachmentMap) ToAnycastEipAddressAttachmentMapOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnycastEipAddressAttachmentMapOutput)
}

type AnycastEipAddressAttachmentOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnycastEipAddressAttachment)(nil)).Elem()
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentOutput() AnycastEipAddressAttachmentOutput {
	return o
}

func (o AnycastEipAddressAttachmentOutput) ToAnycastEipAddressAttachmentOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentOutput {
	return o
}

// The ID of the Anycast EIP instance.
func (o AnycastEipAddressAttachmentOutput) AnycastId() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.AnycastId }).(pulumi.StringOutput)
}

// Binding mode, value:
// - **Default**: The Default mode. The cloud resource instance to be bound is set as the Default origin.
// - **Normal**: In Normal mode, the cloud resource instance to be bound is set to the common source station.
func (o AnycastEipAddressAttachmentOutput) AssociationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.AssociationMode }).(pulumi.StringOutput)
}

// The ID of the cloud resource instance to be bound.
func (o AnycastEipAddressAttachmentOutput) BindInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.BindInstanceId }).(pulumi.StringOutput)
}

// The region ID of the cloud resource instance to be bound.You can only bind cloud resource instances in some regions. You can call the describeanystserverregions operation to obtain the region ID of the cloud resource instances that can be bound.
func (o AnycastEipAddressAttachmentOutput) BindInstanceRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.BindInstanceRegionId }).(pulumi.StringOutput)
}

// The type of the cloud resource instance to be bound. Value:
// - **SlbInstance**: a private network SLB instance.
// - **NetworkInterface**: ENI.
func (o AnycastEipAddressAttachmentOutput) BindInstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.BindInstanceType }).(pulumi.StringOutput)
}

// Binding time.Time is expressed according to ISO8601 standard and UTC time is used. The format is: 'YYYY-MM-DDThh:mm:ssZ'.
func (o AnycastEipAddressAttachmentOutput) BindTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.BindTime }).(pulumi.StringOutput)
}

// The access point information of the associated access area when the cloud resource instance is bound.If you are binding for the first time, this parameter does not need to be configured, and the system automatically associates all access areas. See `popLocations` below.
func (o AnycastEipAddressAttachmentOutput) PopLocations() AnycastEipAddressAttachmentPopLocationArrayOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) AnycastEipAddressAttachmentPopLocationArrayOutput {
		return v.PopLocations
	}).(AnycastEipAddressAttachmentPopLocationArrayOutput)
}

// The secondary private IP address of the elastic network card to be bound.This parameter takes effect only when **BindInstanceType** is set to **NetworkInterface. When you do not enter, this parameter is the primary private IP of the ENI by default.
func (o AnycastEipAddressAttachmentOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringPtrOutput { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// The status of the bound cloud resource instance. Value:BINDING: BINDING.Bound: Bound.UNBINDING: UNBINDING.DELETED: DELETED.MODIFYING: being modified.
func (o AnycastEipAddressAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AnycastEipAddressAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AnycastEipAddressAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnycastEipAddressAttachment)(nil)).Elem()
}

func (o AnycastEipAddressAttachmentArrayOutput) ToAnycastEipAddressAttachmentArrayOutput() AnycastEipAddressAttachmentArrayOutput {
	return o
}

func (o AnycastEipAddressAttachmentArrayOutput) ToAnycastEipAddressAttachmentArrayOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentArrayOutput {
	return o
}

func (o AnycastEipAddressAttachmentArrayOutput) Index(i pulumi.IntInput) AnycastEipAddressAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnycastEipAddressAttachment {
		return vs[0].([]*AnycastEipAddressAttachment)[vs[1].(int)]
	}).(AnycastEipAddressAttachmentOutput)
}

type AnycastEipAddressAttachmentMapOutput struct{ *pulumi.OutputState }

func (AnycastEipAddressAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnycastEipAddressAttachment)(nil)).Elem()
}

func (o AnycastEipAddressAttachmentMapOutput) ToAnycastEipAddressAttachmentMapOutput() AnycastEipAddressAttachmentMapOutput {
	return o
}

func (o AnycastEipAddressAttachmentMapOutput) ToAnycastEipAddressAttachmentMapOutputWithContext(ctx context.Context) AnycastEipAddressAttachmentMapOutput {
	return o
}

func (o AnycastEipAddressAttachmentMapOutput) MapIndex(k pulumi.StringInput) AnycastEipAddressAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnycastEipAddressAttachment {
		return vs[0].(map[string]*AnycastEipAddressAttachment)[vs[1].(string)]
	}).(AnycastEipAddressAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressAttachmentInput)(nil)).Elem(), &AnycastEipAddressAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressAttachmentArrayInput)(nil)).Elem(), AnycastEipAddressAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnycastEipAddressAttachmentMapInput)(nil)).Elem(), AnycastEipAddressAttachmentMap{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AnycastEipAddressAttachmentMapOutput{})
}
