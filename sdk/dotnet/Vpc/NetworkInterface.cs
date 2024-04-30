// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// &gt; **DEPRECATED:** This resource has been renamed to alicloud.ecs.EcsNetworkInterface from version 1.123.1.
    /// 
    /// Provides an ECS Elastic Network Interface resource.
    /// 
    /// For information about Elastic Network Interface and how to use it, see [Elastic Network Interface](https://www.alibabacloud.com/help/doc-detail/58496.html).
    /// 
    /// &gt; **NOTE** Only one of private_ips or private_ips_count can be specified when assign private IPs.
    /// 
    /// ## Example Usage
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
    ///     var name = config.Get("name") ?? "networkInterfaceName";
    ///     var vpc = new AliCloud.Vpc.Network("vpc", new()
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
    ///     var vswitch = new AliCloud.Vpc.Switch("vswitch", new()
    ///     {
    ///         Name = name,
    ///         CidrBlock = "192.168.0.0/24",
    ///         ZoneId = @default.Apply(@default =&gt; @default.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id)),
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var @group = new AliCloud.Ecs.SecurityGroup("group", new()
    ///     {
    ///         Name = name,
    ///         VpcId = vpc.Id,
    ///     });
    /// 
    ///     var defaultNetworkInterface = new AliCloud.Vpc.NetworkInterface("default", new()
    ///     {
    ///         NetworkInterfaceName = name,
    ///         VswitchId = vswitch.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             @group.Id,
    ///         },
    ///         PrivateIp = "192.168.0.2",
    ///         PrivateIpsCount = 3,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ENI can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/networkInterface:NetworkInterface eni eni-abc1234567890000
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/networkInterface:NetworkInterface")]
    public partial class NetworkInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        [Output("ipv4PrefixCount")]
        public Output<int> Ipv4PrefixCount { get; private set; } = null!;

        [Output("ipv4Prefixes")]
        public Output<ImmutableArray<string>> Ipv4Prefixes { get; private set; } = null!;

        [Output("ipv6AddressCount")]
        public Output<int> Ipv6AddressCount { get; private set; } = null!;

        [Output("ipv6Addresses")]
        public Output<ImmutableArray<string>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// (Available in 1.54.0+) The MAC address of an ENI.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkInterfaceName")]
        public Output<string> NetworkInterfaceName { get; private set; } = null!;

        [Output("networkInterfaceTrafficMode")]
        public Output<string> NetworkInterfaceTrafficMode { get; private set; } = null!;

        [Output("primaryIpAddress")]
        public Output<string> PrimaryIpAddress { get; private set; } = null!;

        /// <summary>
        /// The primary private IP of the ENI.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Output("privateIps")]
        public Output<ImmutableArray<string>> PrivateIps { get; private set; } = null!;

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Output("privateIpsCount")]
        public Output<int> PrivateIpsCount { get; private set; } = null!;

        [Output("queueNumber")]
        public Output<int> QueueNumber { get; private set; } = null!;

        /// <summary>
        /// The Id of resource group which the network interface belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        [Output("secondaryPrivateIpAddressCount")]
        public Output<int> SecondaryPrivateIpAddressCount { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A list of security group ids to associate with.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The VSwitch to create the ENI in.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkInterface:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkInterface:NetworkInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        [Input("networkInterfaceTrafficMode")]
        public Input<string>? NetworkInterfaceTrafficMode { get; set; }

        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        /// <summary>
        /// The primary private IP of the ENI.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Obsolete(@"Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead")]
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("queueNumber")]
        public Input<int>? QueueNumber { get; set; }

        /// <summary>
        /// The Id of resource group which the network interface belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group ids to associate with.
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
        /// The VSwitch to create the ENI in.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public NetworkInterfaceArgs()
        {
        }
        public static new NetworkInterfaceArgs Empty => new NetworkInterfaceArgs();
    }

    public sealed class NetworkInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<string>? _ipv4Prefixes;
        public InputList<string> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<string>());
            set => _ipv4Prefixes = value;
        }

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<string>? _ipv6Addresses;
        public InputList<string> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<string>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// (Available in 1.54.0+) The MAC address of an ENI.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        [Input("networkInterfaceTrafficMode")]
        public Input<string>? NetworkInterfaceTrafficMode { get; set; }

        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        /// <summary>
        /// The primary private IP of the ENI.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        [Input("privateIps")]
        private InputList<string>? _privateIps;

        /// <summary>
        /// List of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Obsolete(@"Field 'private_ips' has been deprecated from provider version 1.123.1. New field 'private_ip_addresses' instead")]
        public InputList<string> PrivateIps
        {
            get => _privateIps ?? (_privateIps = new InputList<string>());
            set => _privateIps = value;
        }

        /// <summary>
        /// Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.
        /// </summary>
        [Input("privateIpsCount")]
        public Input<int>? PrivateIpsCount { get; set; }

        [Input("queueNumber")]
        public Input<int>? QueueNumber { get; set; }

        /// <summary>
        /// The Id of resource group which the network interface belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// A list of security group ids to associate with.
        /// </summary>
        [Obsolete(@"Field 'security_groups' has been deprecated from provider version 1.123.1. New field 'security_group_ids' instead")]
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

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
        /// The VSwitch to create the ENI in.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public NetworkInterfaceState()
        {
        }
        public static new NetworkInterfaceState Empty => new NetworkInterfaceState();
    }
}
