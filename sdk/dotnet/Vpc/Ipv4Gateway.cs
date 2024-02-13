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
    /// Provides a Vpc Ipv4 Gateway resource.
    /// 
    /// For information about Vpc Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv4gateway).
    /// 
    /// &gt; **NOTE:** Available in v1.181.0+.
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
    ///     var name = config.Get("name") ?? "tf-testacc-example";
    ///     var defaultResourceGroup = new AliCloud.ResourceManager.ResourceGroup("defaultResourceGroup", new()
    ///     {
    ///         DisplayName = "tf-testAcc-rg665",
    ///         ResourceGroupName = name,
    ///     });
    /// 
    ///     var modify = new AliCloud.ResourceManager.ResourceGroup("modify", new()
    ///     {
    ///         DisplayName = "tf-testAcc-rg298",
    ///         ResourceGroupName = $"{name}1",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = $"{name}2",
    ///         CidrBlock = "10.0.0.0/8",
    ///     });
    /// 
    ///     var defaultIpv4Gateway = new AliCloud.Vpc.Ipv4Gateway("defaultIpv4Gateway", new()
    ///     {
    ///         Ipv4GatewayName = name,
    ///         Ipv4GatewayDescription = "tf-testAcc-Ipv4Gateway",
    ///         ResourceGroupId = defaultResourceGroup.Id,
    ///         VpcId = defaultNetwork.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Vpc Ipv4 Gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/ipv4Gateway:Ipv4Gateway example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/ipv4Gateway:Ipv4Gateway")]
    public partial class Ipv4Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        /// </summary>
        [Output("ipv4GatewayDescription")]
        public Output<string?> Ipv4GatewayDescription { get; private set; } = null!;

        /// <summary>
        /// Resource primary key field.
        /// </summary>
        [Output("ipv4GatewayId")]
        public Output<string> Ipv4GatewayId { get; private set; } = null!;

        /// <summary>
        /// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Output("ipv4GatewayName")]
        public Output<string?> Ipv4GatewayName { get; private set; } = null!;

        /// <summary>
        /// ID of the route table associated with IPv4 Gateway.
        /// </summary>
        [Output("ipv4GatewayRouteTableId")]
        public Output<string> Ipv4GatewayRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group to which the instance belongs.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of the current resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv4Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv4Gateway(string name, Ipv4GatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, args ?? new Ipv4GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv4Gateway(string name, Input<string> id, Ipv4GatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv4Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv4Gateway Get(string name, Input<string> id, Ipv4GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv4Gateway(name, id, state, options);
        }
    }

    public sealed class Ipv4GatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        /// </summary>
        [Input("ipv4GatewayDescription")]
        public Input<string>? Ipv4GatewayDescription { get; set; }

        /// <summary>
        /// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Input("ipv4GatewayName")]
        public Input<string>? Ipv4GatewayName { get; set; }

        /// <summary>
        /// The ID of the resource group to which the instance belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of the current resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public Ipv4GatewayArgs()
        {
        }
        public static new Ipv4GatewayArgs Empty => new Ipv4GatewayArgs();
    }

    public sealed class Ipv4GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request. Value:-**true**: The check request is sent without creating an IPv4 Gateway. Check items include whether required parameters, request format, and business restrictions are filled in. If the check does not pass, the corresponding error is returned. If the check passes, the error code 'DryRunOperation' is returned '.-**false** (default): Sends a normal request, returns an HTTP 2xx status code and directly creates an IPv4 Gateway.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether the IPv4 gateway is active or not. Valid values are **true** and **false**.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The description of the IPv4 gateway. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with http:// or https://.
        /// </summary>
        [Input("ipv4GatewayDescription")]
        public Input<string>? Ipv4GatewayDescription { get; set; }

        /// <summary>
        /// Resource primary key field.
        /// </summary>
        [Input("ipv4GatewayId")]
        public Input<string>? Ipv4GatewayId { get; set; }

        /// <summary>
        /// The name of the IPv4 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Input("ipv4GatewayName")]
        public Input<string>? Ipv4GatewayName { get; set; }

        /// <summary>
        /// ID of the route table associated with IPv4 Gateway.
        /// </summary>
        [Input("ipv4GatewayRouteTableId")]
        public Input<string>? Ipv4GatewayRouteTableId { get; set; }

        /// <summary>
        /// The ID of the resource group to which the instance belongs.
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
        /// The tags of the current resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public Ipv4GatewayState()
        {
        }
        public static new Ipv4GatewayState Empty => new Ipv4GatewayState();
    }
}
