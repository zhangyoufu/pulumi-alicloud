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
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultZones, err := nlb.GetZones(ctx, nil, nil)
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
//			default1, err := vpc.NewSwitch(ctx, "default1", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.1.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      *pulumi.String(defaultZones.Zones[1].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nlb.NewLoadBalancer(ctx, "defaultLoadBalancer", &nlb.LoadBalancerArgs{
//				LoadBalancerName: pulumi.String(name),
//				ResourceGroupId:  *pulumi.String(defaultResourceGroups.Ids[0]),
//				LoadBalancerType: pulumi.String("Network"),
//				AddressType:      pulumi.String("Internet"),
//				AddressIpVersion: pulumi.String("Ipv4"),
//				VpcId:            defaultNetwork.ID(),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("example"),
//				},
//				ZoneMappings: nlb.LoadBalancerZoneMappingArray{
//					&nlb.LoadBalancerZoneMappingArgs{
//						VswitchId: defaultSwitch.ID(),
//						ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
//					},
//					&nlb.LoadBalancerZoneMappingArgs{
//						VswitchId: default1.ID(),
//						ZoneId:    *pulumi.String(defaultZones.Zones[1].Id),
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

	// The protocol version. Valid values:
	// - ipv4 (default): IPv4
	// - DualStack: dual stack
	AddressIpVersion pulumi.StringOutput `pulumi:"addressIpVersion"`
	// The type of IPv4 address used by the NLB instance. Valid values:
	// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
	// The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Specifies whether to enable cross-zone load balancing for the NLB instance.
	CrossZoneEnabled pulumi.BoolOutput `pulumi:"crossZoneEnabled"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolOutput `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringOutput `pulumi:"deletionProtectionReason"`
	// The domain name of the NLB instance.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// The type of IPv6 address used by the NLB instance.
	Ipv6AddressType pulumi.StringOutput `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus pulumi.StringOutput `pulumi:"loadBalancerBusinessStatus"`
	// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringOutput `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	ModificationProtectionStatus pulumi.StringOutput `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The status of the NLB instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VPC where the NLB instance is deployed.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
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
	// The protocol version. Valid values:
	// - ipv4 (default): IPv4
	// - DualStack: dual stack
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The type of IPv4 address used by the NLB instance. Valid values:
	// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType *string `pulumi:"addressType"`
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
	CreateTime *string `pulumi:"createTime"`
	// Specifies whether to enable cross-zone load balancing for the NLB instance.
	CrossZoneEnabled *bool `pulumi:"crossZoneEnabled"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// The domain name of the NLB instance.
	DnsName *string `pulumi:"dnsName"`
	// The type of IPv6 address used by the NLB instance.
	Ipv6AddressType *string `pulumi:"ipv6AddressType"`
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus *string `pulumi:"loadBalancerBusinessStatus"`
	// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The status of the NLB instance.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC where the NLB instance is deployed.
	VpcId *string `pulumi:"vpcId"`
	// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings []LoadBalancerZoneMapping `pulumi:"zoneMappings"`
}

type LoadBalancerState struct {
	// The protocol version. Valid values:
	// - ipv4 (default): IPv4
	// - DualStack: dual stack
	AddressIpVersion pulumi.StringPtrInput
	// The type of IPv4 address used by the NLB instance. Valid values:
	// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType pulumi.StringPtrInput
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId pulumi.StringPtrInput
	// The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
	CreateTime pulumi.StringPtrInput
	// Specifies whether to enable cross-zone load balancing for the NLB instance.
	CrossZoneEnabled pulumi.BoolPtrInput
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringPtrInput
	// The domain name of the NLB instance.
	DnsName pulumi.StringPtrInput
	// The type of IPv6 address used by the NLB instance.
	Ipv6AddressType pulumi.StringPtrInput
	// The business status of the NLB instance.
	LoadBalancerBusinessStatus pulumi.StringPtrInput
	// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	LoadBalancerName pulumi.StringPtrInput
	// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
	LoadBalancerType pulumi.StringPtrInput
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringPtrInput
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// The status of the NLB instance.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC where the NLB instance is deployed.
	VpcId pulumi.StringPtrInput
	// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings LoadBalancerZoneMappingArrayInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// The protocol version. Valid values:
	// - ipv4 (default): IPv4
	// - DualStack: dual stack
	AddressIpVersion *string `pulumi:"addressIpVersion"`
	// The type of IPv4 address used by the NLB instance. Valid values:
	// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType string `pulumi:"addressType"`
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
	// Specifies whether to enable cross-zone load balancing for the NLB instance.
	CrossZoneEnabled *bool `pulumi:"crossZoneEnabled"`
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
	LoadBalancerType *string `pulumi:"loadBalancerType"`
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC where the NLB instance is deployed.
	VpcId string `pulumi:"vpcId"`
	// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
	ZoneMappings []LoadBalancerZoneMapping `pulumi:"zoneMappings"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// The protocol version. Valid values:
	// - ipv4 (default): IPv4
	// - DualStack: dual stack
	AddressIpVersion pulumi.StringPtrInput
	// The type of IPv4 address used by the NLB instance. Valid values:
	// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType pulumi.StringInput
	// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
	BandwidthPackageId pulumi.StringPtrInput
	// Specifies whether to enable cross-zone load balancing for the NLB instance.
	CrossZoneEnabled pulumi.BoolPtrInput
	// Specifies whether to enable deletion protection. Default value: `false`. Valid values:
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason why the deletion protection feature is enabled or disabled. The `deletionProtectionReason` takes effect only when `deletionProtectionEnabled` is set to `true`.
	DeletionProtectionReason pulumi.StringPtrInput
	// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	LoadBalancerName pulumi.StringPtrInput
	// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
	LoadBalancerType pulumi.StringPtrInput
	// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
	ModificationProtectionReason pulumi.StringPtrInput
	// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC where the NLB instance is deployed.
	VpcId pulumi.StringInput
	// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
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

// The protocol version. Valid values:
// - ipv4 (default): IPv4
// - DualStack: dual stack
func (o LoadBalancerOutput) AddressIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.AddressIpVersion }).(pulumi.StringOutput)
}

// The type of IPv4 address used by the NLB instance. Valid values:
// - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
// - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
func (o LoadBalancerOutput) AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.AddressType }).(pulumi.StringOutput)
}

// The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
func (o LoadBalancerOutput) BandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.BandwidthPackageId }).(pulumi.StringOutput)
}

// The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
func (o LoadBalancerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Specifies whether to enable cross-zone load balancing for the NLB instance.
func (o LoadBalancerOutput) CrossZoneEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolOutput { return v.CrossZoneEnabled }).(pulumi.BoolOutput)
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

// The type of IPv6 address used by the NLB instance.
func (o LoadBalancerOutput) Ipv6AddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv6AddressType }).(pulumi.StringOutput)
}

// The business status of the NLB instance.
func (o LoadBalancerOutput) LoadBalancerBusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerBusinessStatus }).(pulumi.StringOutput)
}

// The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
func (o LoadBalancerOutput) LoadBalancerName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerName }).(pulumi.StringOutput)
}

// The type of the instance. Set the value to `Network`, which specifies an NLB instance.
func (o LoadBalancerOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerType }).(pulumi.StringOutput)
}

// The reason why the configuration read-only mode is enabled. The `modificationProtectionReason` takes effect only when `modificationProtectionStatus` is set to `ConsoleProtection`.
func (o LoadBalancerOutput) ModificationProtectionReason() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ModificationProtectionReason }).(pulumi.StringOutput)
}

// Specifies whether to enable the configuration read-only mode. Default value: `NonProtection`. Valid values:
func (o LoadBalancerOutput) ModificationProtectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ModificationProtectionStatus }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o LoadBalancerOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The status of the NLB instance.
func (o LoadBalancerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o LoadBalancerOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the VPC where the NLB instance is deployed.
func (o LoadBalancerOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See `zoneMappings` below.
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
