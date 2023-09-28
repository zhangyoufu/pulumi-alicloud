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
    /// Provides a VPC Dhcp Options Set resource. DHCP option set.
    /// 
    /// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/dhcp-options-sets-overview).
    /// 
    /// &gt; **NOTE:** Available since v1.134.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var domain = config.Get("domain") ?? "terraform-example.com";
    ///     var example = new AliCloud.Vpc.DhcpOptionsSet("example", new()
    ///     {
    ///         DhcpOptionsSetName = name,
    ///         DhcpOptionsSetDescription = name,
    ///         DomainName = domain,
    ///         DomainNameServers = "100.100.2.136",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Dhcp Options Set can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet")]
    public partial class DhcpOptionsSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
        /// </summary>
        [Output("associateVpcs")]
        public Output<ImmutableArray<Outputs.DhcpOptionsSetAssociateVpc>> AssociateVpcs { get; private set; } = null!;

        /// <summary>
        /// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
        /// </summary>
        [Output("dhcpOptionsSetDescription")]
        public Output<string?> DhcpOptionsSetDescription { get; private set; } = null!;

        /// <summary>
        /// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Output("dhcpOptionsSetName")]
        public Output<string?> DhcpOptionsSetName { get; private set; } = null!;

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
        /// </summary>
        [Output("domainNameServers")]
        public Output<string?> DomainNameServers { get; private set; } = null!;

        /// <summary>
        /// Whether to PreCheck only this request, value:
        /// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Output("ipv6LeaseTime")]
        public Output<string> Ipv6LeaseTime { get; private set; } = null!;

        /// <summary>
        /// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Output("leaseTime")]
        public Output<string> LeaseTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the account to which the DHCP options set belongs.
        /// </summary>
        [Output("ownerId")]
        public Output<int> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags of the current resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DhcpOptionsSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DhcpOptionsSet(string name, DhcpOptionsSetArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, args ?? new DhcpOptionsSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DhcpOptionsSet(string name, Input<string> id, DhcpOptionsSetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DhcpOptionsSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DhcpOptionsSet Get(string name, Input<string> id, DhcpOptionsSetState? state = null, CustomResourceOptions? options = null)
        {
            return new DhcpOptionsSet(name, id, state, options);
        }
    }

    public sealed class DhcpOptionsSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("associateVpcs")]
        private InputList<Inputs.DhcpOptionsSetAssociateVpcArgs>? _associateVpcs;

        /// <summary>
        /// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
        /// </summary>
        [Obsolete(@"Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.")]
        public InputList<Inputs.DhcpOptionsSetAssociateVpcArgs> AssociateVpcs
        {
            get => _associateVpcs ?? (_associateVpcs = new InputList<Inputs.DhcpOptionsSetAssociateVpcArgs>());
            set => _associateVpcs = value;
        }

        /// <summary>
        /// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
        /// </summary>
        [Input("dhcpOptionsSetDescription")]
        public Input<string>? DhcpOptionsSetDescription { get; set; }

        /// <summary>
        /// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Input("dhcpOptionsSetName")]
        public Input<string>? DhcpOptionsSetName { get; set; }

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
        /// </summary>
        [Input("domainNameServers")]
        public Input<string>? DomainNameServers { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request, value:
        /// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Input("ipv6LeaseTime")]
        public Input<string>? Ipv6LeaseTime { get; set; }

        /// <summary>
        /// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Input("leaseTime")]
        public Input<string>? LeaseTime { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the current resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DhcpOptionsSetArgs()
        {
        }
        public static new DhcpOptionsSetArgs Empty => new DhcpOptionsSetArgs();
    }

    public sealed class DhcpOptionsSetState : global::Pulumi.ResourceArgs
    {
        [Input("associateVpcs")]
        private InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs>? _associateVpcs;

        /// <summary>
        /// Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc. See `associate_vpcs` below.
        /// </summary>
        [Obsolete(@"Field 'associate_vpcs' has been deprecated from provider version 1.211.0. Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.")]
        public InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs> AssociateVpcs
        {
            get => _associateVpcs ?? (_associateVpcs = new InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs>());
            set => _associateVpcs = value;
        }

        /// <summary>
        /// The description can be blank or contain 1 to 256 characters. It must start with a letter or Chinese character but cannot start with http:// or https://.
        /// </summary>
        [Input("dhcpOptionsSetDescription")]
        public Input<string>? DhcpOptionsSetDescription { get; set; }

        /// <summary>
        /// The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Input("dhcpOptionsSetName")]
        public Input<string>? DhcpOptionsSetName { get; set; }

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
        /// </summary>
        [Input("domainNameServers")]
        public Input<string>? DomainNameServers { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request, value:
        /// - **true**: sends a check request and does not delete the DHCP option set. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request and directly deletes the DHCP option set after checking.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The lease time of the IPv6 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Input("ipv6LeaseTime")]
        public Input<string>? Ipv6LeaseTime { get; set; }

        /// <summary>
        /// The lease time of the IPv4 DHCP option set.When the lease time is set to hours: Unit: h. Value range: 24h ~ 1176h,87600h ~ 175200h. Default value: 87600h.When the lease time is set to day: Unit: d. Value range: 1d ~ 49d,3650d ~ 7300d. Default value: 3650d.
        /// </summary>
        [Input("leaseTime")]
        public Input<string>? LeaseTime { get; set; }

        /// <summary>
        /// The ID of the account to which the DHCP options set belongs.
        /// </summary>
        [Input("ownerId")]
        public Input<int>? OwnerId { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the current resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DhcpOptionsSetState()
        {
        }
        public static new DhcpOptionsSetState Empty => new DhcpOptionsSetState();
    }
}
