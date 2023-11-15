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

// Provides a ECS Network Interface resource.
//
// For information about ECS Network Interface and how to use it, see [What is Network Interface](https://www.alibabacloud.com/help/en/doc-detail/58504.htm).
//
// > **NOTE:** Available in v1.123.1+.
//
// > **NOTE** Only one of `privateIpAddresses` or `secondaryPrivateIpAddressCount` can be specified when assign private IPs.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
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
//			name := "tf-testAcc"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
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
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("192.168.0.0/24"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VpcId:       defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewEcsNetworkInterface(ctx, "defaultEcsNetworkInterface", &ecs.EcsNetworkInterfaceArgs{
//				NetworkInterfaceName: pulumi.String(name),
//				VswitchId:            defaultSwitch.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					defaultSecurityGroup.ID(),
//				},
//				Description:      pulumi.String("Basic test"),
//				PrimaryIpAddress: pulumi.String("192.168.0.2"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("Test"),
//				},
//				ResourceGroupId: *pulumi.String(defaultResourceGroups.Ids[0]),
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
// ECS Network Interface can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface example eni-abcd12345
//
// ```
type EcsNetworkInterface struct {
	pulumi.CustomResourceState

	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
	Ipv6AddressCount pulumi.IntOutput `pulumi:"ipv6AddressCount"`
	// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
	Ipv6Addresses pulumi.StringArrayOutput `pulumi:"ipv6Addresses"`
	// The MAC address of the ENI.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	NetworkInterfaceName pulumi.StringOutput `pulumi:"networkInterfaceName"`
	// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
	PrimaryIpAddress pulumi.StringOutput `pulumi:"primaryIpAddress"`
	// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
	PrivateIpAddresses pulumi.StringArrayOutput `pulumi:"privateIpAddresses"`
	// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayOutput `pulumi:"privateIps"`
	// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntOutput `pulumi:"privateIpsCount"`
	// The queue number of the ENI.
	QueueNumber pulumi.IntOutput `pulumi:"queueNumber"`
	// The resource group id.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The number of private IP addresses that can be automatically created by ECS.
	SecondaryPrivateIpAddressCount pulumi.IntOutput `pulumi:"secondaryPrivateIpAddressCount"`
	// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The status of the ENI.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
}

// NewEcsNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewEcsNetworkInterface(ctx *pulumi.Context,
	name string, args *EcsNetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*EcsNetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EcsNetworkInterface
	err := ctx.RegisterResource("alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcsNetworkInterface gets an existing EcsNetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcsNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcsNetworkInterfaceState, opts ...pulumi.ResourceOption) (*EcsNetworkInterface, error) {
	var resource EcsNetworkInterface
	err := ctx.ReadResource("alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcsNetworkInterface resources.
type ecsNetworkInterfaceState struct {
	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `pulumi:"description"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// The MAC address of the ENI.
	Mac *string `pulumi:"mac"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name *string `pulumi:"name"`
	// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
	PrimaryIpAddress *string `pulumi:"primaryIpAddress"`
	// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp *string `pulumi:"privateIp"`
	// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps []string `pulumi:"privateIps"`
	// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// The queue number of the ENI.
	QueueNumber *int `pulumi:"queueNumber"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The number of private IP addresses that can be automatically created by ECS.
	SecondaryPrivateIpAddressCount *int `pulumi:"secondaryPrivateIpAddressCount"`
	// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups []string `pulumi:"securityGroups"`
	// The status of the ENI.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type EcsNetworkInterfaceState struct {
	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description pulumi.StringPtrInput
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
	Ipv6AddressCount pulumi.IntPtrInput
	// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
	Ipv6Addresses pulumi.StringArrayInput
	// The MAC address of the ENI.
	Mac pulumi.StringPtrInput
	// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name pulumi.StringPtrInput
	// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	NetworkInterfaceName pulumi.StringPtrInput
	// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
	PrimaryIpAddress pulumi.StringPtrInput
	// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp pulumi.StringPtrInput
	// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
	PrivateIpAddresses pulumi.StringArrayInput
	// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayInput
	// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntPtrInput
	// The queue number of the ENI.
	QueueNumber pulumi.IntPtrInput
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput
	// The number of private IP addresses that can be automatically created by ECS.
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
	SecurityGroupIds pulumi.StringArrayInput
	// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayInput
	// The status of the ENI.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (EcsNetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsNetworkInterfaceState)(nil)).Elem()
}

type ecsNetworkInterfaceArgs struct {
	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description *string `pulumi:"description"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name *string `pulumi:"name"`
	// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
	PrimaryIpAddress *string `pulumi:"primaryIpAddress"`
	// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp *string `pulumi:"privateIp"`
	// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps []string `pulumi:"privateIps"`
	// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount *int `pulumi:"privateIpsCount"`
	// The queue number of the ENI.
	QueueNumber *int `pulumi:"queueNumber"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The number of private IP addresses that can be automatically created by ECS.
	SecondaryPrivateIpAddressCount *int `pulumi:"secondaryPrivateIpAddressCount"`
	// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups []string `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a EcsNetworkInterface resource.
type EcsNetworkInterfaceArgs struct {
	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	Description pulumi.StringPtrInput
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
	Ipv6AddressCount pulumi.IntPtrInput
	// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
	Ipv6Addresses pulumi.StringArrayInput
	// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
	Name pulumi.StringPtrInput
	// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
	NetworkInterfaceName pulumi.StringPtrInput
	// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
	PrimaryIpAddress pulumi.StringPtrInput
	// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
	//
	// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
	PrivateIp pulumi.StringPtrInput
	// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
	PrivateIpAddresses pulumi.StringArrayInput
	// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
	//
	// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
	PrivateIps pulumi.StringArrayInput
	// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
	//
	// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
	PrivateIpsCount pulumi.IntPtrInput
	// The queue number of the ENI.
	QueueNumber pulumi.IntPtrInput
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput
	// The number of private IP addresses that can be automatically created by ECS.
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
	SecurityGroupIds pulumi.StringArrayInput
	// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
	//
	// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
	SecurityGroups pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
	VswitchId pulumi.StringInput
}

func (EcsNetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecsNetworkInterfaceArgs)(nil)).Elem()
}

type EcsNetworkInterfaceInput interface {
	pulumi.Input

	ToEcsNetworkInterfaceOutput() EcsNetworkInterfaceOutput
	ToEcsNetworkInterfaceOutputWithContext(ctx context.Context) EcsNetworkInterfaceOutput
}

func (*EcsNetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsNetworkInterface)(nil)).Elem()
}

func (i *EcsNetworkInterface) ToEcsNetworkInterfaceOutput() EcsNetworkInterfaceOutput {
	return i.ToEcsNetworkInterfaceOutputWithContext(context.Background())
}

func (i *EcsNetworkInterface) ToEcsNetworkInterfaceOutputWithContext(ctx context.Context) EcsNetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsNetworkInterfaceOutput)
}

// EcsNetworkInterfaceArrayInput is an input type that accepts EcsNetworkInterfaceArray and EcsNetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `EcsNetworkInterfaceArrayInput` via:
//
//	EcsNetworkInterfaceArray{ EcsNetworkInterfaceArgs{...} }
type EcsNetworkInterfaceArrayInput interface {
	pulumi.Input

	ToEcsNetworkInterfaceArrayOutput() EcsNetworkInterfaceArrayOutput
	ToEcsNetworkInterfaceArrayOutputWithContext(context.Context) EcsNetworkInterfaceArrayOutput
}

type EcsNetworkInterfaceArray []EcsNetworkInterfaceInput

func (EcsNetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsNetworkInterface)(nil)).Elem()
}

func (i EcsNetworkInterfaceArray) ToEcsNetworkInterfaceArrayOutput() EcsNetworkInterfaceArrayOutput {
	return i.ToEcsNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i EcsNetworkInterfaceArray) ToEcsNetworkInterfaceArrayOutputWithContext(ctx context.Context) EcsNetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsNetworkInterfaceArrayOutput)
}

// EcsNetworkInterfaceMapInput is an input type that accepts EcsNetworkInterfaceMap and EcsNetworkInterfaceMapOutput values.
// You can construct a concrete instance of `EcsNetworkInterfaceMapInput` via:
//
//	EcsNetworkInterfaceMap{ "key": EcsNetworkInterfaceArgs{...} }
type EcsNetworkInterfaceMapInput interface {
	pulumi.Input

	ToEcsNetworkInterfaceMapOutput() EcsNetworkInterfaceMapOutput
	ToEcsNetworkInterfaceMapOutputWithContext(context.Context) EcsNetworkInterfaceMapOutput
}

type EcsNetworkInterfaceMap map[string]EcsNetworkInterfaceInput

func (EcsNetworkInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsNetworkInterface)(nil)).Elem()
}

func (i EcsNetworkInterfaceMap) ToEcsNetworkInterfaceMapOutput() EcsNetworkInterfaceMapOutput {
	return i.ToEcsNetworkInterfaceMapOutputWithContext(context.Background())
}

func (i EcsNetworkInterfaceMap) ToEcsNetworkInterfaceMapOutputWithContext(ctx context.Context) EcsNetworkInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcsNetworkInterfaceMapOutput)
}

type EcsNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (EcsNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcsNetworkInterface)(nil)).Elem()
}

func (o EcsNetworkInterfaceOutput) ToEcsNetworkInterfaceOutput() EcsNetworkInterfaceOutput {
	return o
}

func (o EcsNetworkInterfaceOutput) ToEcsNetworkInterfaceOutputWithContext(ctx context.Context) EcsNetworkInterfaceOutput {
	return o
}

// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
func (o EcsNetworkInterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6Addresses` and `ipv6AddressCount` parameters.
func (o EcsNetworkInterfaceOutput) Ipv6AddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.IntOutput { return v.Ipv6AddressCount }).(pulumi.IntOutput)
}

// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
func (o EcsNetworkInterfaceOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringArrayOutput { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

// The MAC address of the ENI.
func (o EcsNetworkInterfaceOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Field `name` has been deprecated from provider version 1.123.1. New field `networkInterfaceName` instead
//
// Deprecated: Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead
func (o EcsNetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
func (o EcsNetworkInterfaceOutput) NetworkInterfaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.NetworkInterfaceName }).(pulumi.StringOutput)
}

// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
func (o EcsNetworkInterfaceOutput) PrimaryIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.PrimaryIpAddress }).(pulumi.StringOutput)
}

// Field `privateIp` has been deprecated from provider version 1.123.1. New field `primaryIpAddress` instead
//
// Deprecated: Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead
func (o EcsNetworkInterfaceOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
func (o EcsNetworkInterfaceOutput) PrivateIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringArrayOutput { return v.PrivateIpAddresses }).(pulumi.StringArrayOutput)
}

// Field `privateIps` has been deprecated from provider version 1.123.1. New field `privateIpAddresses` instead
//
// Deprecated: Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead
func (o EcsNetworkInterfaceOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringArrayOutput { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

// Field `privateIpsCount` has been deprecated from provider version 1.123.1. New field `secondaryPrivateIpAddressCount` instead
//
// Deprecated: Field 'private_ips_count' has been deprecated from provider version 1.123.1. New field 'secondary_private_ip_address_count' instead
func (o EcsNetworkInterfaceOutput) PrivateIpsCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.IntOutput { return v.PrivateIpsCount }).(pulumi.IntOutput)
}

// The queue number of the ENI.
func (o EcsNetworkInterfaceOutput) QueueNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.IntOutput { return v.QueueNumber }).(pulumi.IntOutput)
}

// The resource group id.
func (o EcsNetworkInterfaceOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// The number of private IP addresses that can be automatically created by ECS.
func (o EcsNetworkInterfaceOutput) SecondaryPrivateIpAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.IntOutput { return v.SecondaryPrivateIpAddressCount }).(pulumi.IntOutput)
}

// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
func (o EcsNetworkInterfaceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Field `securityGroups` has been deprecated from provider version 1.123.1. New field `securityGroupIds` instead
//
// Deprecated: Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead
func (o EcsNetworkInterfaceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// The status of the ENI.
func (o EcsNetworkInterfaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A mapping of tags to assign to the resource.
func (o EcsNetworkInterfaceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
func (o EcsNetworkInterfaceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcsNetworkInterface) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

type EcsNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (EcsNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcsNetworkInterface)(nil)).Elem()
}

func (o EcsNetworkInterfaceArrayOutput) ToEcsNetworkInterfaceArrayOutput() EcsNetworkInterfaceArrayOutput {
	return o
}

func (o EcsNetworkInterfaceArrayOutput) ToEcsNetworkInterfaceArrayOutputWithContext(ctx context.Context) EcsNetworkInterfaceArrayOutput {
	return o
}

func (o EcsNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) EcsNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcsNetworkInterface {
		return vs[0].([]*EcsNetworkInterface)[vs[1].(int)]
	}).(EcsNetworkInterfaceOutput)
}

type EcsNetworkInterfaceMapOutput struct{ *pulumi.OutputState }

func (EcsNetworkInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcsNetworkInterface)(nil)).Elem()
}

func (o EcsNetworkInterfaceMapOutput) ToEcsNetworkInterfaceMapOutput() EcsNetworkInterfaceMapOutput {
	return o
}

func (o EcsNetworkInterfaceMapOutput) ToEcsNetworkInterfaceMapOutputWithContext(ctx context.Context) EcsNetworkInterfaceMapOutput {
	return o
}

func (o EcsNetworkInterfaceMapOutput) MapIndex(k pulumi.StringInput) EcsNetworkInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcsNetworkInterface {
		return vs[0].(map[string]*EcsNetworkInterface)[vs[1].(string)]
	}).(EcsNetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcsNetworkInterfaceInput)(nil)).Elem(), &EcsNetworkInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsNetworkInterfaceArrayInput)(nil)).Elem(), EcsNetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcsNetworkInterfaceMapInput)(nil)).Elem(), EcsNetworkInterfaceMap{})
	pulumi.RegisterOutputType(EcsNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(EcsNetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(EcsNetworkInterfaceMapOutput{})
}
