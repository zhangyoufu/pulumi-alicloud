// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Alicloud ECS Disk Attachment as a resource, to attach and detach disks from ECS Instances.
//
// For information about ECS Disk Attachment and how to use it, see [What is Disk Attachment](https://www.alibabacloud.com/help/en/doc-detail/25515.htm).
//
// > **NOTE:** Available since v1.122.0+.
//
// ## Example Usage
//
// # Basic usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
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
//				AvailabilityZone:   pulumi.StringRef(defaultZones.Zones[0].Id),
//				InstanceTypeFamily: pulumi.StringRef("ecs.sn1ne"),
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
//				VpcId:     defaultNetwork.ID(),
//				CidrBlock: pulumi.String("10.4.0.0/24"),
//				ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
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
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := ecs.NewInstance(ctx, "defaultInstance", &ecs.InstanceArgs{
//				AvailabilityZone: *pulumi.String(defaultZones.Zones[0].Id),
//				InstanceName:     pulumi.String(name),
//				HostName:         pulumi.String(name),
//				ImageId:          *pulumi.String(defaultImages.Images[0].Id),
//				InstanceType:     *pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				SecurityGroups: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//				VswitchId: defaultSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			disk, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultEcsDisk, err := ecs.NewEcsDisk(ctx, "defaultEcsDisk", &ecs.EcsDiskArgs{
//				ZoneId:             *pulumi.String(disk.Zones[0].Id),
//				Category:           pulumi.String("cloud_efficiency"),
//				DeleteAutoSnapshot: pulumi.Bool(true),
//				Description:        pulumi.String("Test For Terraform"),
//				DiskName:           pulumi.String(name),
//				EnableAutoSnapshot: pulumi.Bool(true),
//				Encrypted:          pulumi.Bool(true),
//				Size:               pulumi.Int(500),
//				Tags: pulumi.Map{
//					"Created":     pulumi.Any("TF"),
//					"Environment": pulumi.Any("Acceptance-test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsDiskAttachment(ctx, "defaultEcsDiskAttachment", &ecs.EcsDiskAttachmentArgs{
//				DiskId:     defaultEcsDisk.ID(),
//				InstanceId: defaultInstance.ID(),
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
// The disk attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment example d-abc12345678:i-abc12355
// ```
type EcsDiskAttachment struct {
	pulumi.CustomResourceState

	// Whether to mount as a system disk. Default to: `false`.
	Bootable pulumi.BoolPtrOutput `pulumi:"bootable"`
	// Indicates whether the disk is released together with the instance. Default to: `false`.
	DeleteWithInstance pulumi.BoolPtrOutput `pulumi:"deleteWithInstance"`
	// The name of the cloud disk device.
	Device pulumi.StringOutput `pulumi:"device"`
	// ID of the Disk to be attached.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of key pair
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
	Password pulumi.StringPtrOutput `pulumi:"password"`
}

// NewEcsDiskAttachment registers a new resource with the given unique name, arguments, and options.
func NewEcsDiskAttachment(ctx *pulumi.Context,
	name string, args *EcsDiskAttachmentArgs, opts ...pulumi.ResourceOption) (*EcsDiskAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EcsDiskAttachment
	err := ctx.RegisterResource("alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsDiskAttachment gets an existing EcsDiskAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsDiskAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsDiskAttachmentState, opts ...pulumi.ResourceOption) (*EcsDiskAttachment, error) {
	var resource EcsDiskAttachment
	err := ctx.ReadResource("alicloud:ecs/ecsDiskAttachment:EcsDiskAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsDiskAttachment resources.
type ecsDiskAttachmentState struct {
	// Whether to mount as a system disk. Default to: `false`.
	Bootable *bool `pulumi:"bootable"`
	// Indicates whether the disk is released together with the instance. Default to: `false`.
	DeleteWithInstance *bool `pulumi:"deleteWithInstance"`
	// The name of the cloud disk device.
	Device *string `pulumi:"device"`
	// ID of the Disk to be attached.
	DiskId *string `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId *string `pulumi:"instanceId"`
	// The name of key pair
	KeyPairName *string `pulumi:"keyPairName"`
	// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
	Password *string `pulumi:"password"`
}

type EcsDiskAttachmentState struct {
	// Whether to mount as a system disk. Default to: `false`.
	Bootable pulumi.BoolPtrInput
	// Indicates whether the disk is released together with the instance. Default to: `false`.
	DeleteWithInstance pulumi.BoolPtrInput
	// The name of the cloud disk device.
	Device pulumi.StringPtrInput
	// ID of the Disk to be attached.
	DiskId pulumi.StringPtrInput
	// ID of the Instance to attach to.
	InstanceId pulumi.StringPtrInput
	// The name of key pair
	KeyPairName pulumi.StringPtrInput
	// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
	Password pulumi.StringPtrInput
}

func (EcsDiskAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsDiskAttachmentState)(nil)).Elem()
}

type ecsDiskAttachmentArgs struct {
	// Whether to mount as a system disk. Default to: `false`.
	Bootable *bool `pulumi:"bootable"`
	// Indicates whether the disk is released together with the instance. Default to: `false`.
	DeleteWithInstance *bool `pulumi:"deleteWithInstance"`
	// ID of the Disk to be attached.
	DiskId string `pulumi:"diskId"`
	// ID of the Instance to attach to.
	InstanceId string `pulumi:"instanceId"`
	// The name of key pair
	KeyPairName *string `pulumi:"keyPairName"`
	// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
	Password *string `pulumi:"password"`
}

// The set of arguments for constructing a EcsDiskAttachment resource.
type EcsDiskAttachmentArgs struct {
	// Whether to mount as a system disk. Default to: `false`.
	Bootable pulumi.BoolPtrInput
	// Indicates whether the disk is released together with the instance. Default to: `false`.
	DeleteWithInstance pulumi.BoolPtrInput
	// ID of the Disk to be attached.
	DiskId pulumi.StringInput
	// ID of the Instance to attach to.
	InstanceId pulumi.StringInput
	// The name of key pair
	KeyPairName pulumi.StringPtrInput
	// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
	Password pulumi.StringPtrInput
}

func (EcsDiskAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsDiskAttachmentArgs)(nil)).Elem()
}

type EcsDiskAttachmentInput interface {
	pulumi.Input

	ToEcsDiskAttachmentOutput() EcsDiskAttachmentOutput
	ToEcsDiskAttachmentOutputWithContext(ctx context.Context) EcsDiskAttachmentOutput
}

func (*EcsDiskAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsDiskAttachment)(nil)).Elem()
}

func (i *EcsDiskAttachment) ToEcsDiskAttachmentOutput() EcsDiskAttachmentOutput {
	return i.ToEcsDiskAttachmentOutputWithContext(context.Background())
}

func (i *EcsDiskAttachment) ToEcsDiskAttachmentOutputWithContext(ctx context.Context) EcsDiskAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDiskAttachmentOutput)
}

// EcsDiskAttachmentArrayInput is an input type that accepts EcsDiskAttachmentArray and EcsDiskAttachmentArrayOutput values.
// You can construct a concrete instance of `EcsDiskAttachmentArrayInput` via:
//
//	EcsDiskAttachmentArray{ EcsDiskAttachmentArgs{...} }
type EcsDiskAttachmentArrayInput interface {
	pulumi.Input

	ToEcsDiskAttachmentArrayOutput() EcsDiskAttachmentArrayOutput
	ToEcsDiskAttachmentArrayOutputWithContext(context.Context) EcsDiskAttachmentArrayOutput
}

type EcsDiskAttachmentArray []EcsDiskAttachmentInput

func (EcsDiskAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsDiskAttachment)(nil)).Elem()
}

func (i EcsDiskAttachmentArray) ToEcsDiskAttachmentArrayOutput() EcsDiskAttachmentArrayOutput {
	return i.ToEcsDiskAttachmentArrayOutputWithContext(context.Background())
}

func (i EcsDiskAttachmentArray) ToEcsDiskAttachmentArrayOutputWithContext(ctx context.Context) EcsDiskAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDiskAttachmentArrayOutput)
}

// EcsDiskAttachmentMapInput is an input type that accepts EcsDiskAttachmentMap and EcsDiskAttachmentMapOutput values.
// You can construct a concrete instance of `EcsDiskAttachmentMapInput` via:
//
//	EcsDiskAttachmentMap{ "key": EcsDiskAttachmentArgs{...} }
type EcsDiskAttachmentMapInput interface {
	pulumi.Input

	ToEcsDiskAttachmentMapOutput() EcsDiskAttachmentMapOutput
	ToEcsDiskAttachmentMapOutputWithContext(context.Context) EcsDiskAttachmentMapOutput
}

type EcsDiskAttachmentMap map[string]EcsDiskAttachmentInput

func (EcsDiskAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsDiskAttachment)(nil)).Elem()
}

func (i EcsDiskAttachmentMap) ToEcsDiskAttachmentMapOutput() EcsDiskAttachmentMapOutput {
	return i.ToEcsDiskAttachmentMapOutputWithContext(context.Background())
}

func (i EcsDiskAttachmentMap) ToEcsDiskAttachmentMapOutputWithContext(ctx context.Context) EcsDiskAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsDiskAttachmentMapOutput)
}

type EcsDiskAttachmentOutput struct{ *pulumi.OutputState }

func (EcsDiskAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsDiskAttachment)(nil)).Elem()
}

func (o EcsDiskAttachmentOutput) ToEcsDiskAttachmentOutput() EcsDiskAttachmentOutput {
	return o
}

func (o EcsDiskAttachmentOutput) ToEcsDiskAttachmentOutputWithContext(ctx context.Context) EcsDiskAttachmentOutput {
	return o
}

// Whether to mount as a system disk. Default to: `false`.
func (o EcsDiskAttachmentOutput) Bootable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.BoolPtrOutput { return v.Bootable }).(pulumi.BoolPtrOutput)
}

// Indicates whether the disk is released together with the instance. Default to: `false`.
func (o EcsDiskAttachmentOutput) DeleteWithInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.BoolPtrOutput { return v.DeleteWithInstance }).(pulumi.BoolPtrOutput)
}

// The name of the cloud disk device.
func (o EcsDiskAttachmentOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// ID of the Disk to be attached.
func (o EcsDiskAttachmentOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

// ID of the Instance to attach to.
func (o EcsDiskAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of key pair
func (o EcsDiskAttachmentOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.StringPtrOutput { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

// When mounting the system disk, setting the user name and password of the instance is only effective for the administrator and root user names, and other user names are not effective.
func (o EcsDiskAttachmentOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsDiskAttachment) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

type EcsDiskAttachmentArrayOutput struct{ *pulumi.OutputState }

func (EcsDiskAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsDiskAttachment)(nil)).Elem()
}

func (o EcsDiskAttachmentArrayOutput) ToEcsDiskAttachmentArrayOutput() EcsDiskAttachmentArrayOutput {
	return o
}

func (o EcsDiskAttachmentArrayOutput) ToEcsDiskAttachmentArrayOutputWithContext(ctx context.Context) EcsDiskAttachmentArrayOutput {
	return o
}

func (o EcsDiskAttachmentArrayOutput) Index(i pulumi.IntInput) EcsDiskAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsDiskAttachment {
		return vs[0].([]*EcsDiskAttachment)[vs[1].(int)]
	}).(EcsDiskAttachmentOutput)
}

type EcsDiskAttachmentMapOutput struct{ *pulumi.OutputState }

func (EcsDiskAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsDiskAttachment)(nil)).Elem()
}

func (o EcsDiskAttachmentMapOutput) ToEcsDiskAttachmentMapOutput() EcsDiskAttachmentMapOutput {
	return o
}

func (o EcsDiskAttachmentMapOutput) ToEcsDiskAttachmentMapOutputWithContext(ctx context.Context) EcsDiskAttachmentMapOutput {
	return o
}

func (o EcsDiskAttachmentMapOutput) MapIndex(k pulumi.StringInput) EcsDiskAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsDiskAttachment {
		return vs[0].(map[string]*EcsDiskAttachment)[vs[1].(string)]
	}).(EcsDiskAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDiskAttachmentInput)(nil)).Elem(), &EcsDiskAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDiskAttachmentArrayInput)(nil)).Elem(), EcsDiskAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsDiskAttachmentMapInput)(nil)).Elem(), EcsDiskAttachmentMap{})
	pulumi.RegisterOutputType(EcsDiskAttachmentOutput{})
	pulumi.RegisterOutputType(EcsDiskAttachmentArrayOutput{})
	pulumi.RegisterOutputType(EcsDiskAttachmentMapOutput{})
}
