// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **DEPRECATED:** This resource has been renamed to ecs.EcsNetworkInterface from version 1.123.1.
//
// Provides an ECS Elastic Network Interface resource.
//
// For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
//
// > **NOTE** Only one of privateIps or privateIpsCount can be specified when assign private IPs.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "networkInterfaceName"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vswitch, err := vpc.NewSwitch(ctx, "vswitch", &vpc.SwitchArgs{
//				CidrBlock: pulumi.String("192.168.0.0/24"),
//				ZoneId:    *pulumi.String(defaultZones.Zones[0].Id),
//				VpcId:     vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			group, err := ecs.NewSecurityGroup(ctx, "group", &ecs.SecurityGroupArgs{
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewNetworkInterface(ctx, "defaultNetworkInterface", &vpc.NetworkInterfaceArgs{
//				NetworkInterfaceName: pulumi.String(name),
//				VswitchId:            vswitch.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					group.ID(),
//				},
//				PrivateIp:       pulumi.String("192.168.0.2"),
//				PrivateIpsCount: pulumi.Int(3),
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
// ENI can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
//
// ```
type NetworkInterface struct {
	pulumi.CustomResourceState

	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Ipv6AddressCount pulumi.IntOutput         `pulumi:"ipv6AddressCount"`
	Ipv6Addresses    pulumi.StringArrayOutput `pulumi:"ipv6Addresses"`
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 pulumi.StringOutput `pulumi:"name"`
	NetworkInterfaceName pulumi.StringOutput `pulumi:"networkInterfaceName"`
	PrimaryIpAddress     pulumi.StringOutput `pulumi:"primaryIpAddress"`
	// The primary private IP of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp          pulumi.StringOutput      `pulumi:"privateIp"`
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayOutput `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntOutput `pulumi:"privateIpsCount"`
	QueueNumber     pulumi.IntOutput `pulumi:"queueNumber"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId                pulumi.StringPtrOutput   `pulumi:"resourceGroupId"`
	SecondaryPrivateIpAddressCount pulumi.IntOutput         `pulumi:"secondaryPrivateIpAddressCount"`
	SecurityGroupIds               pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A list of security group ids to associate with.
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	Status         pulumi.StringOutput      `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("alicloud:vpc/networkInterface:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("alicloud:vpc/networkInterface:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description      *string  `pulumi:"description"`
	Ipv6AddressCount *int     `pulumi:"ipv6AddressCount"`
	Ipv6Addresses    []string `pulumi:"ipv6Addresses"`
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac *string `pulumi:"mac"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 *string `pulumi:"name"`
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	PrimaryIpAddress     *string `pulumi:"primaryIpAddress"`
	// The primary private IP of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp          *string  `pulumi:"privateIp"`
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	QueueNumber     *int `pulumi:"queueNumber"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId                *string  `pulumi:"resourceGroupId"`
	SecondaryPrivateIpAddressCount *int     `pulumi:"secondaryPrivateIpAddressCount"`
	SecurityGroupIds               []string `pulumi:"securityGroupIds"`
	// A list of security group ids to associate with.
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups []string `pulumi:"securityGroups"`
	Status         *string  `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId *string `pulumi:"vswitchId"`
}

type NetworkInterfaceState struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description      pulumi.StringPtrInput
	Ipv6AddressCount pulumi.IntPtrInput
	Ipv6Addresses    pulumi.StringArrayInput
	// (Available in 1.54.0+) The MAC address of an ENI.
	Mac pulumi.StringPtrInput
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 pulumi.StringPtrInput
	NetworkInterfaceName pulumi.StringPtrInput
	PrimaryIpAddress     pulumi.StringPtrInput
	// The primary private IP of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp          pulumi.StringPtrInput
	PrivateIpAddresses pulumi.StringArrayInput
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntPtrInput
	QueueNumber     pulumi.IntPtrInput
	// The Id of resource group which the network interface belongs.
	ResourceGroupId                pulumi.StringPtrInput
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	SecurityGroupIds               pulumi.StringArrayInput
	// A list of security group ids to associate with.
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayInput
	Status         pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringPtrInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description      *string  `pulumi:"description"`
	Ipv6AddressCount *int     `pulumi:"ipv6AddressCount"`
	Ipv6Addresses    []string `pulumi:"ipv6Addresses"`
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 *string `pulumi:"name"`
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	PrimaryIpAddress     *string `pulumi:"primaryIpAddress"`
	// The primary private IP of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp          *string  `pulumi:"privateIp"`
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps []string `pulumi:"privateIps"`
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	QueueNumber     *int `pulumi:"queueNumber"`
	// The Id of resource group which the network interface belongs.
	ResourceGroupId                *string  `pulumi:"resourceGroupId"`
	SecondaryPrivateIpAddressCount *int     `pulumi:"secondaryPrivateIpAddressCount"`
	SecurityGroupIds               []string `pulumi:"securityGroupIds"`
	// A list of security group ids to associate with.
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups []string `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VSwitch to create the ENI in.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
	Description      pulumi.StringPtrInput
	Ipv6AddressCount pulumi.IntPtrInput
	Ipv6Addresses    pulumi.StringArrayInput
	// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name                 pulumi.StringPtrInput
	NetworkInterfaceName pulumi.StringPtrInput
	PrimaryIpAddress     pulumi.StringPtrInput
	// The primary private IP of the ENI.
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp          pulumi.StringPtrInput
	PrivateIpAddresses pulumi.StringArrayInput
	// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayInput
	// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntPtrInput
	QueueNumber     pulumi.IntPtrInput
	// The Id of resource group which the network interface belongs.
	ResourceGroupId                pulumi.StringPtrInput
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	SecurityGroupIds               pulumi.StringArrayInput
	// A list of security group ids to associate with.
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The VSwitch to create the ENI in.
	VswitchId pulumi.StringInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

// NetworkInterfaceArrayInput is an input type that accepts NetworkInterfaceArray and NetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceArrayInput` via:
//
//	NetworkInterfaceArray{ NetworkInterfaceArgs{...} }
type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

// NetworkInterfaceMapInput is an input type that accepts NetworkInterfaceMap and NetworkInterfaceMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceMapInput` via:
//
//	NetworkInterfaceMap{ "key": NetworkInterfaceArgs{...} }
type NetworkInterfaceMapInput interface {
	pulumi.Input

	ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput
	ToNetworkInterfaceMapOutputWithContext(context.Context) NetworkInterfaceMapOutput
}

type NetworkInterfaceMap map[string]NetworkInterfaceInput

func (NetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return i.ToNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceMap) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceMapOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
func (o NetworkInterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) Ipv6AddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.IntOutput { return v.Ipv6AddressCount }).(pulumi.IntOutput)
}

func (o NetworkInterfaceOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

// (Available in 1.54.0+) The MAC address of an ENI.
func (o NetworkInterfaceOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) NetworkInterfaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.NetworkInterfaceName }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) PrimaryIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.PrimaryIpAddress }).(pulumi.StringOutput)
}

// The primary private IP of the ENI.
//
// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
func (o NetworkInterfaceOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) PrivateIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.PrivateIpAddresses }).(pulumi.StringArrayOutput)
}

// List of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
//
// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
func (o NetworkInterfaceOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

// Number of secondary private IPs to assign to the ENI. Don't use both privateIps and privateIpsCount in the same ENI resource block.
//
// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
func (o NetworkInterfaceOutput) PrivateIpsCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.IntOutput { return v.PrivateIpsCount }).(pulumi.IntOutput)
}

func (o NetworkInterfaceOutput) QueueNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.IntOutput { return v.QueueNumber }).(pulumi.IntOutput)
}

// The Id of resource group which the network interface belongs.
func (o NetworkInterfaceOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) SecondaryPrivateIpAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.IntOutput { return v.SecondaryPrivateIpAddressCount }).(pulumi.IntOutput)
}

func (o NetworkInterfaceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of security group ids to associate with.
//
// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
func (o NetworkInterfaceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o NetworkInterfaceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The VSwitch to create the ENI in.
func (o NetworkInterfaceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkInterface {
		return vs[0].([]*NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutput() NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) ToNetworkInterfaceMapOutputWithContext(ctx context.Context) NetworkInterfaceMapOutput {
	return o
}

func (o NetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkInterface {
		return vs[0].(map[string]*NetworkInterface)[vs[1].(string)]
	}).(NetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceInput)(nil)).Elem(), &NetworkInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceArrayInput)(nil)).Elem(), NetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceMapInput)(nil)).Elem(), NetworkInterfaceMap{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceMapOutput{})
}
