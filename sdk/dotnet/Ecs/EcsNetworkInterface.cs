// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a ECS Network Interface resource.
    /// 
    /// For information about ECS Network Interface and how to use it, see [What is Network Interface](https://www.alibabacloud.com/help/en/doc-detail/58504.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.123.1.
    /// 
    /// &gt; **NOTE** Only one of `private_ip_addresses` or `secondary_private_ip_address_count` can be specified when assign private IPs.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultNetwork = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "192.168.0.0/24",
    ///     });
    /// 
    ///     var @default = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("default", new()
    ///     {
    ///         VswitchName = name,
    ///         CidrBlock = "192.168.0.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultSecurityGroup = new AliCloud.Ecs.SecurityGroup("default", new()
    ///     {
    ///         Name = name,
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    ///     var defaultGetResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke(new()
    ///     {
    ///         Status = "OK",
    ///     });
    /// 
    ///     var defaultEcsNetworkInterface = new AliCloud.Ecs.EcsNetworkInterface("default", new()
    ///     {
    ///         NetworkInterfaceName = name,
    ///         VswitchId = defaultSwitch.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             defaultSecurityGroup.Id,
    ///         },
    ///         Description = "Basic test",
    ///         PrimaryIpAddress = "192.168.0.2",
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///             { "For", "Test" },
    ///         },
    ///         ResourceGroupId = defaultGetResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Network Interface can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface example eni-abcd12345
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface")]
    public partial class EcsNetworkInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The type of the ENI. Default value: `Secondary`. Valid values: `Secondary`, `Trunk`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The number of IPv4 prefixes that can be automatically created by ECS. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv4_prefixes` and `ipv4_prefix_count` parameters.
        /// </summary>
        [Output("ipv4PrefixCount")]
        public Output<int> Ipv4PrefixCount { get; private set; } = null!;

        /// <summary>
        /// A list of IPv4 prefixes to be assigned to the ENI. Support up to 10.
        /// </summary>
        [Output("ipv4Prefixes")]
        public Output<ImmutableArray<string>> Ipv4Prefixes { get; private set; } = null!;

        /// <summary>
        /// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6_addresses` and `ipv6_address_count` parameters.
        /// </summary>
        [Output("ipv6AddressCount")]
        public Output<int> Ipv6AddressCount { get; private set; } = null!;

        /// <summary>
        /// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<string>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// The MAC address of the ENI.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
        /// </summary>
        [Output("networkInterfaceName")]
        public Output<string> NetworkInterfaceName { get; private set; } = null!;

        /// <summary>
        /// The communication mode of the ENI. Default value: `Standard`. Valid values: `Standard`, `HighPerformance`.
        /// </summary>
        [Output("networkInterfaceTrafficMode")]
        public Output<string> NetworkInterfaceTrafficMode { get; private set; } = null!;

        /// <summary>
        /// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
        /// </summary>
        [Output("primaryIpAddress")]
        public Output<string> PrimaryIpAddress { get; private set; } = null!;

        /// <summary>
        /// Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Field `private_ips` has been deprecated from provider version 1.123.1. New field `private_ip_addresses` instead
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// Field `private_ips_count` has been deprecated from provider version 1.123.1. New field `secondary_private_ip_address_count` instead
        /// </summary>
        [Output("privateIpsCount")]
        public Output<int> PrivateIpsCount { get; private set; } = null!;

        /// <summary>
        /// The queue number of the ENI.
        /// </summary>
        [Output("queueNumber")]
        public Output<int> QueueNumber { get; private set; } = null!;

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The number of private IP addresses that can be automatically created by ECS.
        /// </summary>
        [Output("secondaryPrivateIpAddressCount")]
        public Output<int> SecondaryPrivateIpAddressCount { get; private set; } = null!;

        /// <summary>
        /// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Field `security_groups` has been deprecated from provider version 1.123.1. New field `security_group_ids` instead
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The status of the ENI.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a EcsNetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EcsNetworkInterface(string name, EcsNetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface", name, args ?? new EcsNetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EcsNetworkInterface(string name, Input<string> id, EcsNetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsNetworkInterface:EcsNetworkInterface", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EcsNetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EcsNetworkInterface Get(string name, Input<string> id, EcsNetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new EcsNetworkInterface(name, id, state, options);
        }
    }

    public sealed class EcsNetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of the ENI. Default value: `Secondary`. Valid values: `Secondary`, `Trunk`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The number of IPv4 prefixes that can be automatically created by ECS. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv4_prefixes` and `ipv4_prefix_count` parameters.
        /// </summary>
        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;

        /// <summary>
        /// A list of IPv4 prefixes to be assigned to the ENI. Support up to 10.
        /// </summary>
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        /// <summary>
        /// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6_addresses` and `ipv6_address_count` parameters.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// The communication mode of the ENI. Default value: `Standard`. Valid values: `Standard`, `HighPerformance`.
        /// </summary>
        [Input("networkInterfaceTrafficMode")]
        public Input<string>? NetworkInterfaceTrafficMode { get; set; }

        /// <summary>
        /// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
        /// </summary>
        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        /// <summary>
        /// Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// Field `private_ips` has been deprecated from provider version 1.123.1. New field `private_ip_addresses` instead
        /// </summary>
        [Obsolete(@"Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead")]
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Field `private_ips_count` has been deprecated from provider version 1.123.1. New field `secondary_private_ip_address_count` instead
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        /// <summary>
        /// The queue number of the ENI.
        /// </summary>
        [Input("queueNumber")]
        public Input<int>? QueueNumber { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The number of private IP addresses that can be automatically created by ECS.
        /// </summary>
        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Field `security_groups` has been deprecated from provider version 1.123.1. New field `security_group_ids` instead
        /// </summary>
        [Obsolete(@"Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead")]
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public EcsNetworkInterfaceArgs()
        {
        }
        public static new EcsNetworkInterfaceArgs Empty => new EcsNetworkInterfaceArgs();
    }

    public sealed class EcsNetworkInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of the ENI. Default value: `Secondary`. Valid values: `Secondary`, `Trunk`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The number of IPv4 prefixes that can be automatically created by ECS. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv4_prefixes` and `ipv4_prefix_count` parameters.
        /// </summary>
        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;

        /// <summary>
        /// A list of IPv4 prefixes to be assigned to the ENI. Support up to 10.
        /// </summary>
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        /// <summary>
        /// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10. **NOTE:** You cannot specify both the `ipv6_addresses` and `ipv6_address_count` parameters.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;

        /// <summary>
        /// A list of IPv6 address to be assigned to the primary ENI. Support up to 10.
        /// </summary>
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// The MAC address of the ENI.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.123.1. New field `network_interface_name` instead
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// The communication mode of the ENI. Default value: `Standard`. Valid values: `Standard`, `HighPerformance`.
        /// </summary>
        [Input("networkInterfaceTrafficMode")]
        public Input<string>? NetworkInterfaceTrafficMode { get; set; }

        /// <summary>
        /// The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.
        /// </summary>
        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        /// <summary>
        /// Field `private_ip` has been deprecated from provider version 1.123.1. New field `primary_ip_address` instead
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// Field `private_ips` has been deprecated from provider version 1.123.1. New field `private_ip_addresses` instead
        /// </summary>
        [Obsolete(@"Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead")]
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Field `private_ips_count` has been deprecated from provider version 1.123.1. New field `secondary_private_ip_address_count` instead
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        /// <summary>
        /// The queue number of the ENI.
        /// </summary>
        [Input("queueNumber")]
        public Input<int>? QueueNumber { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The number of private IP addresses that can be automatically created by ECS.
        /// </summary>
        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// Field `security_groups` has been deprecated from provider version 1.123.1. New field `security_group_ids` instead
        /// </summary>
        [Obsolete(@"Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead")]
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        /// <summary>
        /// The status of the ENI.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public EcsNetworkInterfaceState()
        {
        }
        public static new EcsNetworkInterfaceState Empty => new EcsNetworkInterfaceState();
    }
}
