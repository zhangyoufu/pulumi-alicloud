// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos
{
    /// <summary>
    /// Provides a Ddos Bgp Ip resource.
    /// 
    /// For information about Ddos Bgp Ip and how to use it, see [What is Ip](https://www.alibabacloud.com/help/en/ddos-protection/latest/addip).
    /// 
    /// &gt; **NOTE:** Available since v1.180.0.
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
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var instance = new AliCloud.Ddos.DdosBgpInstance("instance", new()
    ///     {
    ///         BaseBandwidth = 20,
    ///         Bandwidth = -1,
    ///         IpCount = 100,
    ///         IpType = "IPv4",
    ///         NormalBandwidth = 100,
    ///         Type = "Enterprise",
    ///     });
    /// 
    ///     var defaultEipAddress = new AliCloud.Ecs.EipAddress("defaultEipAddress", new()
    ///     {
    ///         AddressName = name,
    ///     });
    /// 
    ///     var defaultBgpIp = new AliCloud.Ddos.BgpIp("defaultBgpIp", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Ip = defaultEipAddress.IpAddress,
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Ddos Bgp Ip can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ddos/bgpIp:BgpIp example &lt;instance_id&gt;:&lt;ip&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ddos/bgpIp:BgpIp")]
    public partial class BgpIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the native protection enterprise instance to be operated.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The IP address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string?> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The current state of the IP address. Valid Value: `normal`, `hole_begin`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BgpIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpIp(string name, BgpIpArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ddos/bgpIp:BgpIp", name, args ?? new BgpIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpIp(string name, Input<string> id, BgpIpState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ddos/bgpIp:BgpIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BgpIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpIp Get(string name, Input<string> id, BgpIpState? state = null, CustomResourceOptions? options = null)
        {
            return new BgpIp(name, id, state, options);
        }
    }

    public sealed class BgpIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the native protection enterprise instance to be operated.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        public BgpIpArgs()
        {
        }
        public static new BgpIpArgs Empty => new BgpIpArgs();
    }

    public sealed class BgpIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the native protection enterprise instance to be operated.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The current state of the IP address. Valid Value: `normal`, `hole_begin`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BgpIpState()
        {
        }
        public static new BgpIpState Empty => new BgpIpState();
    }
}
