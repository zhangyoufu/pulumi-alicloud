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
    /// Provides a Vpc Vpc resource. A VPC instance creates a VPC. You can fully control your own VPC, such as selecting IP address ranges, configuring routing tables, and gateways. You can use Alibaba cloud resources such as cloud servers, apsaradb for RDS, and load balancer in your own VPC.
    /// 
    /// &gt; **NOTE:** Available since v1.0.0.
    /// 
    /// &gt; **NOTE:** This resource will auto build a router and a route table while it uses `alicloud.vpc.Network` to build a vpc resource.
    /// 
    /// &gt; **NOTE:** Currently, the IPv4 / IPv6 dual-stack VPC function is under public testing. Only the following regions support IPv4 / IPv6 dual-stack VPC: `cn-hangzhou`, `cn-shanghai`, `cn-shenzhen`, `cn-beijing`, `cn-huhehaote`, `cn-hongkong` and `ap-southeast-1`, and need to apply for public beta qualification. To use, please [submit an application](https://www.alibabacloud.com/help/en/vpc/getting-started/create-a-vpc-with-an-ipv6-cidr-block).
    /// 
    /// ## Module Support
    /// 
    /// You can use the existing vpc module
    /// to create a VPC and several VSwitches one-click.
    /// 
    /// For information about Vpc Vpc and how to use it, see [What is Vpc](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/what-is-a-vpc).
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var @default = new AliCloud.Vpc.Network("default", new()
    ///     {
    ///         Ipv6Isp = "BGP",
    ///         Description = "test",
    ///         CidrBlock = "10.0.0.0/8",
    ///         VpcName = name,
    ///         EnableIpv6 = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Vpc Vpc can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/network:Network example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/network:Network")]
    public partial class Network : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The status of ClassicLink function.
        /// </summary>
        [Output("classicLinkEnabled")]
        public Output<bool?> ClassicLinkEnabled { get; private set; } = null!;

        /// <summary>
        /// The creation time of the VPC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The VPC description. Defaults to null.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to PreCheck this request only. Value:
        /// - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the IPv6 network segment. Value:
        /// - **false** (default): not enabled.
        /// - **true**: on.
        /// </summary>
        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR block of the VPC.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR block information of the VPC.
        /// </summary>
        [Output("ipv6CidrBlocks")]
        public Output<ImmutableArray<Outputs.NetworkIpv6CidrBlock>> Ipv6CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// The IPv6 address segment type of the VPC. Value:
        /// - **BGP** (default): Alibaba Cloud BGP IPv6.
        /// - **ChinaMobile**: China Mobile (single line).
        /// - **ChinaUnicom**: China Unicom (single line).
        /// - **ChinaTelecom**: China Telecom (single line).
        /// &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
        /// </summary>
        [Output("ipv6Isp")]
        public Output<string?> Ipv6Isp { get; private set; } = null!;

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the VPC belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The route table ID of the router created by default on VPC creation.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the router created by default on VPC creation.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;

        /// <summary>
        /// (Deprecated since v1.206.0+) Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
        /// </summary>
        [Output("routerTableId")]
        public Output<string> RouterTableId { get; private set; } = null!;

        /// <summary>
        /// Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
        /// </summary>
        [Output("secondaryCidrBlocks")]
        public Output<ImmutableArray<string>> SecondaryCidrBlocks { get; private set; } = null!;

        /// <summary>
        /// The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of Vpc.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of user CIDRs.
        /// </summary>
        [Output("userCidrs")]
        public Output<ImmutableArray<string>> UserCidrs { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC. Defaults to null.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Output("vpcName")]
        public Output<string> VpcName { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/network:Network", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The status of ClassicLink function.
        /// </summary>
        [Input("classicLinkEnabled")]
        public Input<bool>? ClassicLinkEnabled { get; set; }

        /// <summary>
        /// The VPC description. Defaults to null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to PreCheck this request only. Value:
        /// - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether to enable the IPv6 network segment. Value:
        /// - **false** (default): not enabled.
        /// - **true**: on.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The IPv6 address segment type of the VPC. Value:
        /// - **BGP** (default): Alibaba Cloud BGP IPv6.
        /// - **ChinaMobile**: China Mobile (single line).
        /// - **ChinaUnicom**: China Unicom (single line).
        /// - **ChinaTelecom**: China Telecom (single line).
        /// &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
        /// </summary>
        [Input("ipv6Isp")]
        public Input<string>? Ipv6Isp { get; set; }

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the resource group to which the VPC belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("secondaryCidrBlocks")]
        private InputList<string>? _secondaryCidrBlocks;

        /// <summary>
        /// Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
        /// </summary>
        [Obsolete(@"Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.")]
        public InputList<string> SecondaryCidrBlocks
        {
            get => _secondaryCidrBlocks ?? (_secondaryCidrBlocks = new InputList<string>());
            set => _secondaryCidrBlocks = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of Vpc.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("userCidrs")]
        private InputList<string>? _userCidrs;

        /// <summary>
        /// A list of user CIDRs.
        /// </summary>
        public InputList<string> UserCidrs
        {
            get => _userCidrs ?? (_userCidrs = new InputList<string>());
            set => _userCidrs = value;
        }

        /// <summary>
        /// The name of the VPC. Defaults to null.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }

    public sealed class NetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block for the VPC. The `cidr_block` is Optional and default value is `172.16.0.0/12` after v1.119.0+.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The status of ClassicLink function.
        /// </summary>
        [Input("classicLinkEnabled")]
        public Input<bool>? ClassicLinkEnabled { get; set; }

        /// <summary>
        /// The creation time of the VPC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The VPC description. Defaults to null.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to PreCheck this request only. Value:
        /// - **true**: sends a check request and does not create a VPC. Check items include whether required parameters, request format, and business restrictions have been filled in. If the check fails, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.
        /// - **false** (default): Sends a normal request, returns the HTTP 2xx status code after the check, and directly creates a VPC.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether to enable the IPv6 network segment. Value:
        /// - **false** (default): not enabled.
        /// - **true**: on.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The IPv6 CIDR block of the VPC.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        [Input("ipv6CidrBlocks")]
        private InputList<Inputs.NetworkIpv6CidrBlockGetArgs>? _ipv6CidrBlocks;

        /// <summary>
        /// The IPv6 CIDR block information of the VPC.
        /// </summary>
        public InputList<Inputs.NetworkIpv6CidrBlockGetArgs> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<Inputs.NetworkIpv6CidrBlockGetArgs>());
            set => _ipv6CidrBlocks = value;
        }

        /// <summary>
        /// The IPv6 address segment type of the VPC. Value:
        /// - **BGP** (default): Alibaba Cloud BGP IPv6.
        /// - **ChinaMobile**: China Mobile (single line).
        /// - **ChinaUnicom**: China Unicom (single line).
        /// - **ChinaTelecom**: China Telecom (single line).
        /// &gt; **NOTE:**  If a single-line bandwidth whitelist is enabled, this field can be set to **ChinaTelecom** (China Telecom), **ChinaUnicom** (China Unicom), or **ChinaMobile** (China Mobile).
        /// </summary>
        [Input("ipv6Isp")]
        public Input<string>? Ipv6Isp { get; set; }

        /// <summary>
        /// Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the resource group to which the VPC belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The route table ID of the router created by default on VPC creation.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The ID of the router created by default on VPC creation.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        /// <summary>
        /// (Deprecated since v1.206.0+) Field 'router_table_id' has been deprecated from provider version 1.206.0. New field 'route_table_id' instead.
        /// </summary>
        [Input("routerTableId")]
        public Input<string>? RouterTableId { get; set; }

        [Input("secondaryCidrBlocks")]
        private InputList<string>? _secondaryCidrBlocks;

        /// <summary>
        /// Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud.vpc.Ipv4CidrBlock` resource cannot be used at the same time.
        /// </summary>
        [Obsolete(@"Field 'SecondaryCidrBlocks' has been deprecated from provider version 1.206.0. Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.")]
        public InputList<string> SecondaryCidrBlocks
        {
            get => _secondaryCidrBlocks ?? (_secondaryCidrBlocks = new InputList<string>());
            set => _secondaryCidrBlocks = value;
        }

        /// <summary>
        /// The status of the VPC. Valid values:  **Pending**: The VPC is being configured. **Available**: The VPC is available.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of Vpc.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("userCidrs")]
        private InputList<string>? _userCidrs;

        /// <summary>
        /// A list of user CIDRs.
        /// </summary>
        public InputList<string> UserCidrs
        {
            get => _userCidrs ?? (_userCidrs = new InputList<string>());
            set => _userCidrs = value;
        }

        /// <summary>
        /// The name of the VPC. Defaults to null.
        /// 
        /// The following arguments will be discarded. Please use new fields as soon as possible:
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public NetworkState()
        {
        }
        public static new NetworkState Empty => new NetworkState();
    }
}
