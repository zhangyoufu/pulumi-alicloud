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

// Provides an Alicloud EIP Association resource for associating Elastic IP to ECS Instance, SLB Instance or Nat Gateway.
//
// > **NOTE:** `ecs.EipAssociation` is useful in scenarios where EIPs are either
//
//	pre-existing or distributed to customers or users and therefore cannot be changed.
//
// > **NOTE:** From version 1.7.1, the resource support to associate EIP to SLB Instance or Nat Gateway.
//
// > **NOTE:** One EIP can only be associated with ECS or SLB instance which in the VPC.
//
// > **NOTE:** Available since v1.117.0.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			exampleZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("Instance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(exampleZones.Zones[0].Id),
//				CpuCoreCount:     pulumi.IntRef(1),
//				MemorySize:       pulumi.Float64Ref(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex: pulumi.StringRef("^ubuntu_[0-9]+_[0-9]+_x64*"),
//				Owners:    pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      *pulumi.String(exampleZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSecurityGroup, err := ecs.NewSecurityGroup(ctx, "exampleSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: exampleNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := ecs.NewInstance(ctx, "exampleInstance", &ecs.InstanceArgs{
//				AvailabilityZone: *pulumi.String(exampleZones.Zones[0].Id),
//				InstanceName:     pulumi.String(name),
//				ImageId:          *pulumi.String(exampleImages.Images[0].Id),
//				InstanceType:     *pulumi.String(exampleInstanceTypes.InstanceTypes[0].Id),
//				SecurityGroups: pulumi.StringArray{
//					exampleSecurityGroup.ID(),
//				},
//				VswitchId: exampleSwitch.ID(),
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleEipAddress, err := ecs.NewEipAddress(ctx, "exampleEipAddress", &ecs.EipAddressArgs{
//				AddressName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEipAssociation(ctx, "exampleEipAssociation", &ecs.EipAssociationArgs{
//				AllocationId: exampleEipAddress.ID(),
//				InstanceId:   exampleInstance.ID(),
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
// ## Module Support
//
// You can use the existing eip module
// to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.
//
// ## Import
//
// Elastic IP address association can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:ecs/eipAssociation:EipAssociation example <allocation_id>:<instance_id>
// ```
type EipAssociation struct {
	pulumi.CustomResourceState

	// The ID of the EIP that you want to associate with an instance.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The IP address in the CIDR block of the vSwitch.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewEipAssociation registers a new resource with the given unique name, arguments, and options.
func NewEipAssociation(ctx *pulumi.Context,
	name string, args *EipAssociationArgs, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationId == nil {
		return nil, errors.New("invalid value for required argument 'AllocationId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EipAssociation
	err := ctx.RegisterResource("alicloud:ecs/eipAssociation:EipAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEipAssociation gets an existing EipAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEipAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipAssociationState, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	var resource EipAssociation
	err := ctx.ReadResource("alicloud:ecs/eipAssociation:EipAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EipAssociation resources.
type eipAssociationState struct {
	// The ID of the EIP that you want to associate with an instance.
	AllocationId *string `pulumi:"allocationId"`
	// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
	Force *bool `pulumi:"force"`
	// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
	InstanceId *string `pulumi:"instanceId"`
	// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
	InstanceType *string `pulumi:"instanceType"`
	// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
	Mode *string `pulumi:"mode"`
	// The IP address in the CIDR block of the vSwitch.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
	VpcId *string `pulumi:"vpcId"`
}

type EipAssociationState struct {
	// The ID of the EIP that you want to associate with an instance.
	AllocationId pulumi.StringPtrInput
	// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
	Force pulumi.BoolPtrInput
	// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
	InstanceId pulumi.StringPtrInput
	// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
	InstanceType pulumi.StringPtrInput
	// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
	Mode pulumi.StringPtrInput
	// The IP address in the CIDR block of the vSwitch.
	PrivateIpAddress pulumi.StringPtrInput
	// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
	VpcId pulumi.StringPtrInput
}

func (EipAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationState)(nil)).Elem()
}

type eipAssociationArgs struct {
	// The ID of the EIP that you want to associate with an instance.
	AllocationId string `pulumi:"allocationId"`
	// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
	Force *bool `pulumi:"force"`
	// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
	InstanceId string `pulumi:"instanceId"`
	// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
	InstanceType *string `pulumi:"instanceType"`
	// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
	Mode *string `pulumi:"mode"`
	// The IP address in the CIDR block of the vSwitch.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a EipAssociation resource.
type EipAssociationArgs struct {
	// The ID of the EIP that you want to associate with an instance.
	AllocationId pulumi.StringInput
	// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
	Force pulumi.BoolPtrInput
	// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
	InstanceId pulumi.StringInput
	// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
	InstanceType pulumi.StringPtrInput
	// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
	Mode pulumi.StringPtrInput
	// The IP address in the CIDR block of the vSwitch.
	PrivateIpAddress pulumi.StringPtrInput
	// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
	VpcId pulumi.StringPtrInput
}

func (EipAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipAssociationArgs)(nil)).Elem()
}

type EipAssociationInput interface {
	pulumi.Input

	ToEipAssociationOutput() EipAssociationOutput
	ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput
}

func (*EipAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAssociation)(nil)).Elem()
}

func (i *EipAssociation) ToEipAssociationOutput() EipAssociationOutput {
	return i.ToEipAssociationOutputWithContext(context.Background())
}

func (i *EipAssociation) ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationOutput)
}

// EipAssociationArrayInput is an input type that accepts EipAssociationArray and EipAssociationArrayOutput values.
// You can construct a concrete instance of `EipAssociationArrayInput` via:
//
//	EipAssociationArray{ EipAssociationArgs{...} }
type EipAssociationArrayInput interface {
	pulumi.Input

	ToEipAssociationArrayOutput() EipAssociationArrayOutput
	ToEipAssociationArrayOutputWithContext(context.Context) EipAssociationArrayOutput
}

type EipAssociationArray []EipAssociationInput

func (EipAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EipAssociation)(nil)).Elem()
}

func (i EipAssociationArray) ToEipAssociationArrayOutput() EipAssociationArrayOutput {
	return i.ToEipAssociationArrayOutputWithContext(context.Background())
}

func (i EipAssociationArray) ToEipAssociationArrayOutputWithContext(ctx context.Context) EipAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationArrayOutput)
}

// EipAssociationMapInput is an input type that accepts EipAssociationMap and EipAssociationMapOutput values.
// You can construct a concrete instance of `EipAssociationMapInput` via:
//
//	EipAssociationMap{ "key": EipAssociationArgs{...} }
type EipAssociationMapInput interface {
	pulumi.Input

	ToEipAssociationMapOutput() EipAssociationMapOutput
	ToEipAssociationMapOutputWithContext(context.Context) EipAssociationMapOutput
}

type EipAssociationMap map[string]EipAssociationInput

func (EipAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EipAssociation)(nil)).Elem()
}

func (i EipAssociationMap) ToEipAssociationMapOutput() EipAssociationMapOutput {
	return i.ToEipAssociationMapOutputWithContext(context.Background())
}

func (i EipAssociationMap) ToEipAssociationMapOutputWithContext(ctx context.Context) EipAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipAssociationMapOutput)
}

type EipAssociationOutput struct{ *pulumi.OutputState }

func (EipAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EipAssociation)(nil)).Elem()
}

func (o EipAssociationOutput) ToEipAssociationOutput() EipAssociationOutput {
	return o
}

func (o EipAssociationOutput) ToEipAssociationOutputWithContext(ctx context.Context) EipAssociationOutput {
	return o
}

// The ID of the EIP that you want to associate with an instance.
func (o EipAssociationOutput) AllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringOutput { return v.AllocationId }).(pulumi.StringOutput)
}

// When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default value: `false`. Valid values: `true`, `false`.
func (o EipAssociationOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
func (o EipAssociationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of the instance with which you want to associate the EIP. Valid values: `Nat`, `SlbInstance`, `EcsInstance`, `NetworkInterface`, `HaVip` and `IpAddress`.
func (o EipAssociationOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The association mode. Default value: `NAT`. Valid values: `NAT`, `BINDED`, `MULTI_BINDED`. **Note:** This parameter is required only when `instanceType` is set to `NetworkInterface`.
func (o EipAssociationOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The IP address in the CIDR block of the vSwitch.
func (o EipAssociationOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringPtrOutput { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instanceType` is set to `IpAddress`.
func (o EipAssociationOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EipAssociation) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type EipAssociationArrayOutput struct{ *pulumi.OutputState }

func (EipAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EipAssociation)(nil)).Elem()
}

func (o EipAssociationArrayOutput) ToEipAssociationArrayOutput() EipAssociationArrayOutput {
	return o
}

func (o EipAssociationArrayOutput) ToEipAssociationArrayOutputWithContext(ctx context.Context) EipAssociationArrayOutput {
	return o
}

func (o EipAssociationArrayOutput) Index(i pulumi.IntInput) EipAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EipAssociation {
		return vs[0].([]*EipAssociation)[vs[1].(int)]
	}).(EipAssociationOutput)
}

type EipAssociationMapOutput struct{ *pulumi.OutputState }

func (EipAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EipAssociation)(nil)).Elem()
}

func (o EipAssociationMapOutput) ToEipAssociationMapOutput() EipAssociationMapOutput {
	return o
}

func (o EipAssociationMapOutput) ToEipAssociationMapOutputWithContext(ctx context.Context) EipAssociationMapOutput {
	return o
}

func (o EipAssociationMapOutput) MapIndex(k pulumi.StringInput) EipAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EipAssociation {
		return vs[0].(map[string]*EipAssociation)[vs[1].(string)]
	}).(EipAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipAssociationInput)(nil)).Elem(), &EipAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipAssociationArrayInput)(nil)).Elem(), EipAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipAssociationMapInput)(nil)).Elem(), EipAssociationMap{})
	pulumi.RegisterOutputType(EipAssociationOutput{})
	pulumi.RegisterOutputType(EipAssociationArrayOutput{})
	pulumi.RegisterOutputType(EipAssociationMapOutput{})
}
