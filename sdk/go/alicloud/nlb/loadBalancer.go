// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nlb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NLB Load Balancer resource.
//
// For information about NLB Load Balancer and how to use it, see [What is Load Balancer](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer).
//
// > **NOTE:** Available since v1.191.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nlb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
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
//			_default, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetZones, err := nlb.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(defaultGetZones.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			default1, err := vpc.NewSwitch(ctx, "default1", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.1.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(defaultGetZones.Zones[1].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nlb.NewLoadBalancer(ctx, "default", &nlb.LoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				ResourceGroupId:  pulumi.String(_default.Ids[0]),
//				LoadBalancerType: pulumi.String("Network"),
//				AddressType:      pulumi.String("Internet"),
//				AddressIpVersion: pulumi.String("Ipv4"),
//				VpcId:            defaultNetwork.ID(),
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//				ZoneMappings: nlb.LoadBalancerZoneMappingArray{
//					&nlb.LoadBalancerZoneMappingArgs{
//						VswitchId: defaultSwitch.ID(),
//						ZoneId:    pulumi.String(defaultGetZones.Zones[0].Id),
//					},
//					&nlb.LoadBalancerZoneMappingArgs{
//						VswitchId: default1.ID(),
//						ZoneId:    pulumi.String(defaultGetZones.Zones[1].Id),
//					},
//				},
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
// NLB Load Balancer can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:nlb/loadBalancer:LoadBalancer example <id>
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// Protocol version. Value:
	// - **Ipv4**:IPv4 type.
	// - **DualStack**: Double Stack type.
	AddressIpVersion pulumi.StringOutput `pulumi:"addressIpVersion"`
	// The network address type of IPv4 for network load balancing. Value:
	// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
	// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// The ID of the shared bandwidth package associated with the public network instance.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
	// Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether cross-zone is enabled for a network-based load balancing instance. Value:
	// - **true**: on.
	// - **false**: closed.
	CrossZoneEnabled pulumi.BoolOutput `pulumi:"crossZoneEnabled"`
	// Delete protection. See `deletionProtectionConfig` below.
	DeletionProtectionConfig LoadBalancerDeletionProtectionConfigOutput `pulumi:"deletionProtectionConfig"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolOutput `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringOutput `pulumi:"deletionProtectionReason"`
	// The domain name of the NLB instance.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// The IPv6 address type of network load balancing. Value:
	// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
	// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
	Ipv6AddressType pulumi.StringOutput `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus pulumi.StringOutput `pulumi:"loadBalancerBusinessStatus"`
	// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
	LoadBalancerName pulumi.StringPtrOutput `pulumi:"loadBalancerName"`
	// Load balancing type. Only value: **network**, which indicates network-based load balancing.
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`
	// Modify protection. See `modificationProtectionConfig` below.
	ModificationProtectionConfig LoadBalancerModificationProtectionConfigOutput `pulumi:"modificationProtectionConfig"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringOutput `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
	// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
	ModificationProtectionStatus pulumi.StringOutput `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The security group to which the network-based SLB instance belongs.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// List of labels.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the network-based SLB instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings LoadBalancerZoneMappingArrayOutput `pulumi:"zoneMappings"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressType == nil {
		return nil, errors.New("invalid value for required argument 'AddressType'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.ZoneMappings == nil {
		return nil, errors.New("invalid value for required argument 'ZoneMappings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("alicloud:nlb/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("alicloud:nlb/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// Protocol version. Value:
	// - **Ipv4**:IPv4 type.
	// - **DualStack**: Double Stack type.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The network address type of IPv4 for network load balancing. Value:
	// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
	// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
	AddressType *string `pulumi:"addressType"`
	// The ID of the shared bandwidth package associated with the public network instance.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
	CreateTime *string `pulumi:"createTime"`
	// Whether cross-zone is enabled for a network-based load balancing instance. Value:
	// - **true**: on.
	// - **false**: closed.
	CrossZoneEnabled *bool `pulumi:"crossZoneEnabled"`
	// Delete protection. See `deletionProtectionConfig` below.
	DeletionProtectionConfig *LoadBalancerDeletionProtectionConfig `pulumi:"deletionProtectionConfig"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// The domain name of the NLB instance.
	DnsName *string `pulumi:"dnsName"`
	// The IPv6 address type of network load balancing. Value:
	// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
	// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus *string `pulumi:"loadBalancerBusinessStatus"`
	// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Load balancing type. Only value: **network**, which indicates network-based load balancing.
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// Modify protection. See `modificationProtectionConfig` below.
	ModificationProtectionConfig *LoadBalancerModificationProtectionConfig `pulumi:"modificationProtectionConfig"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
	// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security group to which the network-based SLB instance belongs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// List of labels.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the network-based SLB instance.
	VpcId *string `pulumi:"vpcId"`
	// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings []LoadBalancerZoneMapping `pulumi:"zoneMappings"`
}

type LoadBalancerState struct {
	// Protocol version. Value:
	// - **Ipv4**:IPv4 type.
	// - **DualStack**: Double Stack type.
	AddressIpVersion pulumi.StringPtrInput
	// The network address type of IPv4 for network load balancing. Value:
	// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
	// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
	AddressType pulumi.StringPtrInput
	// The ID of the shared bandwidth package associated with the public network instance.
	BandwidthPackageId pulumi.StringPtrInput
	// Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
	CreateTime pulumi.StringPtrInput
	// Whether cross-zone is enabled for a network-based load balancing instance. Value:
	// - **true**: on.
	// - **false**: closed.
	CrossZoneEnabled pulumi.BoolPtrInput
	// Delete protection. See `deletionProtectionConfig` below.
	DeletionProtectionConfig LoadBalancerDeletionProtectionConfigPtrInput
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringPtrInput
	// The domain name of the NLB instance.
	DnsName pulumi.StringPtrInput
	// The IPv6 address type of network load balancing. Value:
	// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
	// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
	Ipv6AddressType pulumi.StringPtrInput
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus pulumi.StringPtrInput
	// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
	LoadBalancerName pulumi.StringPtrInput
	// Load balancing type. Only value: **network**, which indicates network-based load balancing.
	LoadBalancerType pulumi.StringPtrInput
	// Modify protection. See `modificationProtectionConfig` below.
	ModificationProtectionConfig LoadBalancerModificationProtectionConfigPtrInput
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringPtrInput
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
	// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The security group to which the network-based SLB instance belongs.
	SecurityGroupIds pulumi.StringArrayInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// List of labels.
	Tags pulumi.StringMapInput
	// The ID of the network-based SLB instance.
	VpcId pulumi.StringPtrInput
	// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings LoadBalancerZoneMappingArrayInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// Protocol version. Value:
	// - **Ipv4**:IPv4 type.
	// - **DualStack**: Double Stack type.
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The network address type of IPv4 for network load balancing. Value:
	// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
	// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
	AddressType string `pulumi:"addressType"`
	// The ID of the shared bandwidth package associated with the public network instance.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Whether cross-zone is enabled for a network-based load balancing instance. Value:
	// - **true**: on.
	// - **false**: closed.
	CrossZoneEnabled *bool `pulumi:"crossZoneEnabled"`
	// Delete protection. See `deletionProtectionConfig` below.
	DeletionProtectionConfig *LoadBalancerDeletionProtectionConfig `pulumi:"deletionProtectionConfig"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// The IPv6 address type of network load balancing. Value:
	// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
	// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Load balancing type. Only value: **network**, which indicates network-based load balancing.
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// Modify protection. See `modificationProtectionConfig` below.
	ModificationProtectionConfig *LoadBalancerModificationProtectionConfig `pulumi:"modificationProtectionConfig"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
	// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security group to which the network-based SLB instance belongs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// List of labels.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the network-based SLB instance.
	VpcId string `pulumi:"vpcId"`
	// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings []LoadBalancerZoneMapping `pulumi:"zoneMappings"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Protocol version. Value:
	// - **Ipv4**:IPv4 type.
	// - **DualStack**: Double Stack type.
	AddressIpVersion pulumi.StringPtrInput
	// The network address type of IPv4 for network load balancing. Value:
	// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
	// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
	AddressType pulumi.StringInput
	// The ID of the shared bandwidth package associated with the public network instance.
	BandwidthPackageId pulumi.StringPtrInput
	// Whether cross-zone is enabled for a network-based load balancing instance. Value:
	// - **true**: on.
	// - **false**: closed.
	CrossZoneEnabled pulumi.BoolPtrInput
	// Delete protection. See `deletionProtectionConfig` below.
	DeletionProtectionConfig LoadBalancerDeletionProtectionConfigPtrInput
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringPtrInput
	// The IPv6 address type of network load balancing. Value:
	// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
	// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
	Ipv6AddressType pulumi.StringPtrInput
	// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
	LoadBalancerName pulumi.StringPtrInput
	// Load balancing type. Only value: **network**, which indicates network-based load balancing.
	LoadBalancerType pulumi.StringPtrInput
	// Modify protection. See `modificationProtectionConfig` below.
	ModificationProtectionConfig LoadBalancerModificationProtectionConfigPtrInput
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringPtrInput
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
	// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The security group to which the network-based SLB instance belongs.
	SecurityGroupIds pulumi.StringArrayInput
	// List of labels.
	Tags pulumi.StringMapInput
	// The ID of the network-based SLB instance.
	VpcId pulumi.StringInput
	// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings LoadBalancerZoneMappingArrayInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

// LoadBalancerArrayInput is an input type that accepts LoadBalancerArray and LoadBalancerArrayOutput values.
// You can construct a concrete instance of `LoadBalancerArrayInput` via:
//
//	LoadBalancerArray{ LoadBalancerArgs{...} }
type LoadBalancerArrayInput interface {
	pulumi.Input

	ToLoadBalancerArrayOutput() LoadBalancerArrayOutput
	ToLoadBalancerArrayOutputWithContext(context.Context) LoadBalancerArrayOutput
}

type LoadBalancerArray []LoadBalancerInput

func (LoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return i.ToLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerArrayOutput)
}

// LoadBalancerMapInput is an input type that accepts LoadBalancerMap and LoadBalancerMapOutput values.
// You can construct a concrete instance of `LoadBalancerMapInput` via:
//
//	LoadBalancerMap{ "key": LoadBalancerArgs{...} }
type LoadBalancerMapInput interface {
	pulumi.Input

	ToLoadBalancerMapOutput() LoadBalancerMapOutput
	ToLoadBalancerMapOutputWithContext(context.Context) LoadBalancerMapOutput
}

type LoadBalancerMap map[string]LoadBalancerInput

func (LoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerMap) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return i.ToLoadBalancerMapOutputWithContext(context.Background())
}

func (i LoadBalancerMap) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerMapOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// Protocol version. Value:
// - **Ipv4**:IPv4 type.
// - **DualStack**: Double Stack type.
func (o LoadBalancerOutput) AddressIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.AddressIpVersion }).(pulumi.StringOutput)
}

// The network address type of IPv4 for network load balancing. Value:
// - **Internet**: public network. Load balancer has a public network IP address, and the DNS domain name is resolved to a public network IP address, so it can be accessed in a public network environment.
// - **Intranet**: private network. The server load balancer only has a private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the intranet environment of the VPC where the server load balancer is located.
func (o LoadBalancerOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.AddressType }).(pulumi.StringOutput)
}

// The ID of the shared bandwidth package associated with the public network instance.
func (o LoadBalancerOutput) BandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.BandwidthPackageId }).(pulumi.StringOutput)
}

// Resource creation time, using Greenwich Mean Time, formating' yyyy-MM-ddTHH:mm:ssZ '.
func (o LoadBalancerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether cross-zone is enabled for a network-based load balancing instance. Value:
// - **true**: on.
// - **false**: closed.
func (o LoadBalancerOutput) CrossZoneEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolOutput { return v.CrossZoneEnabled }).(pulumi.BoolOutput)
}

// Delete protection. See `deletionProtectionConfig` below.
func (o LoadBalancerOutput) DeletionProtectionConfig() LoadBalancerDeletionProtectionConfigOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerDeletionProtectionConfigOutput { return v.DeletionProtectionConfig }).(LoadBalancerDeletionProtectionConfigOutput)
}

// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
func (o LoadBalancerOutput) DeletionProtectionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolOutput { return v.DeletionProtectionEnabled }).(pulumi.BoolOutput)
}

// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
func (o LoadBalancerOutput) DeletionProtectionReason() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.DeletionProtectionReason }).(pulumi.StringOutput)
}

// The domain name of the NLB instance.
func (o LoadBalancerOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// The IPv6 address type of network load balancing. Value:
// - **Internet**: Server Load Balancer has a public IP address, and the DNS domain name is resolved to a public IP address, so it can be accessed in a public network environment.
// - **Intranet**: SLB only has the private IP address, and the DNS domain name is resolved to the private IP address, so it can only be accessed by the Intranet environment of the VPC where SLB is located.
func (o LoadBalancerOutput) Ipv6AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv6AddressType }).(pulumi.StringOutput)
}

// The business status of the NLB instance.
func (o LoadBalancerOutput) LoadBalancerBusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerBusinessStatus }).(pulumi.StringOutput)
}

// The name of the network-based load balancing instance.  2 to 128 English or Chinese characters in length, which must start with a letter or Chinese, and can contain numbers, half-width periods (.), underscores (_), and dashes (-).
func (o LoadBalancerOutput) LoadBalancerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.LoadBalancerName }).(pulumi.StringPtrOutput)
}

// Load balancing type. Only value: **network**, which indicates network-based load balancing.
func (o LoadBalancerOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerType }).(pulumi.StringOutput)
}

// Modify protection. See `modificationProtectionConfig` below.
func (o LoadBalancerOutput) ModificationProtectionConfig() LoadBalancerModificationProtectionConfigOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerModificationProtectionConfigOutput {
		return v.ModificationProtectionConfig
	}).(LoadBalancerModificationProtectionConfigOutput)
}

// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
func (o LoadBalancerOutput) ModificationProtectionReason() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ModificationProtectionReason }).(pulumi.StringOutput)
}

// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
// - `NonProtection`: Does not enable the configuration read-only mode. You cannot set the `modificationProtectionReason`. If the `modificationProtectionReason` is set, the value is cleared.
// - `ConsoleProtection`: Enables the configuration read-only mode. You can set the `modificationProtectionReason`.
func (o LoadBalancerOutput) ModificationProtectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ModificationProtectionStatus }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o LoadBalancerOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The security group to which the network-based SLB instance belongs.
func (o LoadBalancerOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The status of the resource.
func (o LoadBalancerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// List of labels.
func (o LoadBalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the network-based SLB instance.
func (o LoadBalancerOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The list of zones and vSwitch mappings. You must add at least two zones and a maximum of 10 zones. See `zoneMappings` below.
func (o LoadBalancerOutput) ZoneMappings() LoadBalancerZoneMappingArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerZoneMappingArrayOutput { return v.ZoneMappings }).(LoadBalancerZoneMappingArrayOutput)
}

type LoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) Index(i pulumi.IntInput) LoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].([]*LoadBalancer)[vs[1].(int)]
	}).(LoadBalancerOutput)
}

type LoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].(map[string]*LoadBalancer)[vs[1].(string)]
	}).(LoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerArrayInput)(nil)).Elem(), LoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerMapInput)(nil)).Elem(), LoadBalancerMap{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
	pulumi.RegisterOutputType(LoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerMapOutput{})
}
