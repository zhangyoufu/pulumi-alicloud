// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.
//
// > **NOTE:** One ECS instance can be added into multiple master slave server groups.
//
// > **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.
//
// > **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.
//
// > **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.
//
// > **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.
//
// > **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.
//
// > **NOTE:** Available in 1.54.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			msServerGroupZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableDiskCategory:     pulumi.StringRef("cloud_efficiency"),
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			msServerGroupInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				AvailabilityZone: pulumi.StringRef(msServerGroupZones.Zones[0].Id),
//				EniAmount:        pulumi.IntRef(2),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			image, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				NameRegex:  pulumi.StringRef("^ubuntu_18.*64"),
//				MostRecent: pulumi.BoolRef(true),
//				Owners:     pulumi.StringRef("system"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			slbMasterSlaveServerGroup := "forSlbMasterSlaveServerGroup"
//			if param := cfg.Get("slbMasterSlaveServerGroup"); param != "" {
//				slbMasterSlaveServerGroup = param
//			}
//			mainNetwork, err := vpc.NewNetwork(ctx, "mainNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(slbMasterSlaveServerGroup),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			mainSwitch, err := vpc.NewSwitch(ctx, "mainSwitch", &vpc.SwitchArgs{
//				VpcId:       mainNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/16"),
//				ZoneId:      *pulumi.String(msServerGroupZones.Zones[0].Id),
//				VswitchName: pulumi.String(slbMasterSlaveServerGroup),
//			})
//			if err != nil {
//				return err
//			}
//			groupSecurityGroup, err := ecs.NewSecurityGroup(ctx, "groupSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: mainNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			var msServerGroupInstance []*ecs.Instance
//			for index := 0; index < 2; index++ {
//				key0 := index
//				_ := index
//				__res, err := ecs.NewInstance(ctx, fmt.Sprintf("msServerGroupInstance-%v", key0), &ecs.InstanceArgs{
//					ImageId:      *pulumi.String(image.Images[0].Id),
//					InstanceType: *pulumi.String(msServerGroupInstanceTypes.InstanceTypes[0].Id),
//					InstanceName: pulumi.String(slbMasterSlaveServerGroup),
//					SecurityGroups: pulumi.StringArray{
//						groupSecurityGroup.ID(),
//					},
//					InternetChargeType:      pulumi.String("PayByTraffic"),
//					InternetMaxBandwidthOut: pulumi.Int(10),
//					AvailabilityZone:        *pulumi.String(msServerGroupZones.Zones[0].Id),
//					InstanceChargeType:      pulumi.String("PostPaid"),
//					SystemDiskCategory:      pulumi.String("cloud_efficiency"),
//					VswitchId:               mainSwitch.ID(),
//				})
//				if err != nil {
//					return err
//				}
//				msServerGroupInstance = append(msServerGroupInstance, __res)
//			}
//			msServerGroupApplicationLoadBalancer, err := slb.NewApplicationLoadBalancer(ctx, "msServerGroupApplicationLoadBalancer", &slb.ApplicationLoadBalancerArgs{
//				LoadBalancerName: pulumi.String(slbMasterSlaveServerGroup),
//				VswitchId:        mainSwitch.ID(),
//				LoadBalancerSpec: pulumi.String("slb.s2.small"),
//			})
//			if err != nil {
//				return err
//			}
//			msServerGroupEcsNetworkInterface, err := ecs.NewEcsNetworkInterface(ctx, "msServerGroupEcsNetworkInterface", &ecs.EcsNetworkInterfaceArgs{
//				NetworkInterfaceName: pulumi.String(slbMasterSlaveServerGroup),
//				VswitchId:            mainSwitch.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					groupSecurityGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsNetworkInterfaceAttachment(ctx, "msServerGroupEcsNetworkInterfaceAttachment", &ecs.EcsNetworkInterfaceAttachmentArgs{
//				InstanceId:         msServerGroupInstance[0].ID(),
//				NetworkInterfaceId: msServerGroupEcsNetworkInterface.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			groupMasterSlaveServerGroup, err := slb.NewMasterSlaveServerGroup(ctx, "groupMasterSlaveServerGroup", &slb.MasterSlaveServerGroupArgs{
//				LoadBalancerId: msServerGroupApplicationLoadBalancer.ID(),
//				Servers: slb.MasterSlaveServerGroupServerArray{
//					&slb.MasterSlaveServerGroupServerArgs{
//						ServerId:   msServerGroupInstance[0].ID(),
//						Port:       pulumi.Int(100),
//						Weight:     pulumi.Int(100),
//						ServerType: pulumi.String("Master"),
//					},
//					&slb.MasterSlaveServerGroupServerArgs{
//						ServerId:   msServerGroupInstance[1].ID(),
//						Port:       pulumi.Int(100),
//						Weight:     pulumi.Int(100),
//						ServerType: pulumi.String("Slave"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = slb.NewListener(ctx, "tcp", &slb.ListenerArgs{
//				LoadBalancerId:           msServerGroupApplicationLoadBalancer.ID(),
//				MasterSlaveServerGroupId: groupMasterSlaveServerGroup.ID(),
//				FrontendPort:             pulumi.Int(22),
//				Protocol:                 pulumi.String("tcp"),
//				Bandwidth:                pulumi.Int(10),
//				HealthCheckType:          pulumi.String("tcp"),
//				PersistenceTimeout:       pulumi.Int(3600),
//				HealthyThreshold:         pulumi.Int(8),
//				UnhealthyThreshold:       pulumi.Int(8),
//				HealthCheckTimeout:       pulumi.Int(8),
//				HealthCheckInterval:      pulumi.Int(5),
//				HealthCheckHttpCode:      pulumi.String("http_2xx"),
//				HealthCheckConnectPort:   pulumi.Int(20),
//				HealthCheckUri:           pulumi.String("/console"),
//				EstablishedTimeout:       pulumi.Int(600),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Block servers
//
// The servers mapping supports the following:
//
// * `serverIds` - (Required) A list backend server ID (ECS instance ID).
// * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
// * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
// * `serverType` - (Optional) The server type of the backend server. Valid value Master, Slave.
// * `isBackup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.
//
// ## Import
//
// Load balancer master slave server group can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
// ```
type MasterSlaveServerGroup struct {
	pulumi.CustomResourceState

	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayOutput `pulumi:"servers"`
}

// NewMasterSlaveServerGroup registers a new resource with the given unique name, arguments, and options.
func NewMasterSlaveServerGroup(ctx *pulumi.Context,
	name string, args *MasterSlaveServerGroupArgs, opts ...pulumi.ResourceOption) (*MasterSlaveServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MasterSlaveServerGroup
	err := ctx.RegisterResource("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMasterSlaveServerGroup gets an existing MasterSlaveServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMasterSlaveServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MasterSlaveServerGroupState, opts ...pulumi.ResourceOption) (*MasterSlaveServerGroup, error) {
	var resource MasterSlaveServerGroup
	err := ctx.ReadResource("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MasterSlaveServerGroup resources.
type masterSlaveServerGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers []MasterSlaveServerGroupServer `pulumi:"servers"`
}

type MasterSlaveServerGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringPtrInput
	// Name of the master slave server group.
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayInput
}

func (MasterSlaveServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSlaveServerGroupState)(nil)).Elem()
}

type masterSlaveServerGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers []MasterSlaveServerGroupServer `pulumi:"servers"`
}

// The set of arguments for constructing a MasterSlaveServerGroup resource.
type MasterSlaveServerGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringInput
	// Name of the master slave server group.
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayInput
}

func (MasterSlaveServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSlaveServerGroupArgs)(nil)).Elem()
}

type MasterSlaveServerGroupInput interface {
	pulumi.Input

	ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput
	ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput
}

func (*MasterSlaveServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterSlaveServerGroup)(nil)).Elem()
}

func (i *MasterSlaveServerGroup) ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput {
	return i.ToMasterSlaveServerGroupOutputWithContext(context.Background())
}

func (i *MasterSlaveServerGroup) ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSlaveServerGroupOutput)
}

// MasterSlaveServerGroupArrayInput is an input type that accepts MasterSlaveServerGroupArray and MasterSlaveServerGroupArrayOutput values.
// You can construct a concrete instance of `MasterSlaveServerGroupArrayInput` via:
//
//	MasterSlaveServerGroupArray{ MasterSlaveServerGroupArgs{...} }
type MasterSlaveServerGroupArrayInput interface {
	pulumi.Input

	ToMasterSlaveServerGroupArrayOutput() MasterSlaveServerGroupArrayOutput
	ToMasterSlaveServerGroupArrayOutputWithContext(context.Context) MasterSlaveServerGroupArrayOutput
}

type MasterSlaveServerGroupArray []MasterSlaveServerGroupInput

func (MasterSlaveServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MasterSlaveServerGroup)(nil)).Elem()
}

func (i MasterSlaveServerGroupArray) ToMasterSlaveServerGroupArrayOutput() MasterSlaveServerGroupArrayOutput {
	return i.ToMasterSlaveServerGroupArrayOutputWithContext(context.Background())
}

func (i MasterSlaveServerGroupArray) ToMasterSlaveServerGroupArrayOutputWithContext(ctx context.Context) MasterSlaveServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSlaveServerGroupArrayOutput)
}

// MasterSlaveServerGroupMapInput is an input type that accepts MasterSlaveServerGroupMap and MasterSlaveServerGroupMapOutput values.
// You can construct a concrete instance of `MasterSlaveServerGroupMapInput` via:
//
//	MasterSlaveServerGroupMap{ "key": MasterSlaveServerGroupArgs{...} }
type MasterSlaveServerGroupMapInput interface {
	pulumi.Input

	ToMasterSlaveServerGroupMapOutput() MasterSlaveServerGroupMapOutput
	ToMasterSlaveServerGroupMapOutputWithContext(context.Context) MasterSlaveServerGroupMapOutput
}

type MasterSlaveServerGroupMap map[string]MasterSlaveServerGroupInput

func (MasterSlaveServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MasterSlaveServerGroup)(nil)).Elem()
}

func (i MasterSlaveServerGroupMap) ToMasterSlaveServerGroupMapOutput() MasterSlaveServerGroupMapOutput {
	return i.ToMasterSlaveServerGroupMapOutputWithContext(context.Background())
}

func (i MasterSlaveServerGroupMap) ToMasterSlaveServerGroupMapOutputWithContext(ctx context.Context) MasterSlaveServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSlaveServerGroupMapOutput)
}

type MasterSlaveServerGroupOutput struct{ *pulumi.OutputState }

func (MasterSlaveServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MasterSlaveServerGroup)(nil)).Elem()
}

func (o MasterSlaveServerGroupOutput) ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput {
	return o
}

func (o MasterSlaveServerGroupOutput) ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput {
	return o
}

// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
func (o MasterSlaveServerGroupOutput) DeleteProtectionValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MasterSlaveServerGroup) pulumi.BoolPtrOutput { return v.DeleteProtectionValidation }).(pulumi.BoolPtrOutput)
}

// The Load Balancer ID which is used to launch a new master slave server group.
func (o MasterSlaveServerGroupOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSlaveServerGroup) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

// Name of the master slave server group.
func (o MasterSlaveServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MasterSlaveServerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
func (o MasterSlaveServerGroupOutput) Servers() MasterSlaveServerGroupServerArrayOutput {
	return o.ApplyT(func(v *MasterSlaveServerGroup) MasterSlaveServerGroupServerArrayOutput { return v.Servers }).(MasterSlaveServerGroupServerArrayOutput)
}

type MasterSlaveServerGroupArrayOutput struct{ *pulumi.OutputState }

func (MasterSlaveServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MasterSlaveServerGroup)(nil)).Elem()
}

func (o MasterSlaveServerGroupArrayOutput) ToMasterSlaveServerGroupArrayOutput() MasterSlaveServerGroupArrayOutput {
	return o
}

func (o MasterSlaveServerGroupArrayOutput) ToMasterSlaveServerGroupArrayOutputWithContext(ctx context.Context) MasterSlaveServerGroupArrayOutput {
	return o
}

func (o MasterSlaveServerGroupArrayOutput) Index(i pulumi.IntInput) MasterSlaveServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MasterSlaveServerGroup {
		return vs[0].([]*MasterSlaveServerGroup)[vs[1].(int)]
	}).(MasterSlaveServerGroupOutput)
}

type MasterSlaveServerGroupMapOutput struct{ *pulumi.OutputState }

func (MasterSlaveServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MasterSlaveServerGroup)(nil)).Elem()
}

func (o MasterSlaveServerGroupMapOutput) ToMasterSlaveServerGroupMapOutput() MasterSlaveServerGroupMapOutput {
	return o
}

func (o MasterSlaveServerGroupMapOutput) ToMasterSlaveServerGroupMapOutputWithContext(ctx context.Context) MasterSlaveServerGroupMapOutput {
	return o
}

func (o MasterSlaveServerGroupMapOutput) MapIndex(k pulumi.StringInput) MasterSlaveServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MasterSlaveServerGroup {
		return vs[0].(map[string]*MasterSlaveServerGroup)[vs[1].(string)]
	}).(MasterSlaveServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MasterSlaveServerGroupInput)(nil)).Elem(), &MasterSlaveServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*MasterSlaveServerGroupArrayInput)(nil)).Elem(), MasterSlaveServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MasterSlaveServerGroupMapInput)(nil)).Elem(), MasterSlaveServerGroupMap{})
	pulumi.RegisterOutputType(MasterSlaveServerGroupOutput{})
	pulumi.RegisterOutputType(MasterSlaveServerGroupArrayOutput{})
	pulumi.RegisterOutputType(MasterSlaveServerGroupMapOutput{})
}
