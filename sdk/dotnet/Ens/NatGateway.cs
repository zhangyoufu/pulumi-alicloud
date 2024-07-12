// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ens
{
    /// <summary>
    /// Provides a Ens Nat Gateway resource.
    /// 
    /// Nat gateway of ENS.
    /// 
    /// For information about Ens Nat Gateway and how to use it, see [What is Nat Gateway](https://www.alibabacloud.com/help/en/).
    /// 
    /// &gt; **NOTE:** Available since v1.227.0.
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
    ///     var ensRegionId = config.Get("ensRegionId") ?? "cn-chenzhou-telecom_unicom_cmcc";
    ///     var defaultObbrL7 = new AliCloud.Ens.Network("defaultObbrL7", new()
    ///     {
    ///         NetworkName = name,
    ///         Description = name,
    ///         CidrBlock = "10.0.0.0/8",
    ///         EnsRegionId = ensRegionId,
    ///     });
    /// 
    ///     var defaulteFw783 = new AliCloud.Ens.Vswitch("defaulteFw783", new()
    ///     {
    ///         CidrBlock = "10.0.8.0/24",
    ///         VswitchName = name,
    ///         EnsRegionId = defaultObbrL7.EnsRegionId,
    ///         NetworkId = defaultObbrL7.Id,
    ///     });
    /// 
    ///     var @default = new AliCloud.Ens.NatGateway("default", new()
    ///     {
    ///         VswitchId = defaulteFw783.Id,
    ///         EnsRegionId = defaulteFw783.EnsRegionId,
    ///         NetworkId = defaulteFw783.NetworkId,
    ///         InstanceType = "enat.default",
    ///         NatName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ens Nat Gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ens/natGateway:NatGateway example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ens/natGateway:NatGateway")]
    public partial class NatGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time. UTC time, in the format of YYYY-MM-DDThh:mm:ssZ.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the ENS node.
        /// </summary>
        [Output("ensRegionId")]
        public Output<string> EnsRegionId { get; private set; } = null!;

        /// <summary>
        /// NAT specifications. Value: `enat.default`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The name of the NAT gateway. The length is 1 to 128 characters, but it cannot start with 'http:// 'or 'https.
        /// </summary>
        [Output("natName")]
        public Output<string?> NatName { get; private set; } = null!;

        /// <summary>
        /// The network ID.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// The vSwitch ID.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a NatGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NatGateway(string name, NatGatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ens/natGateway:NatGateway", name, args ?? new NatGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NatGateway(string name, Input<string> id, NatGatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ens/natGateway:NatGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NatGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NatGateway Get(string name, Input<string> id, NatGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new NatGateway(name, id, state, options);
        }
    }

    public sealed class NatGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the ENS node.
        /// </summary>
        [Input("ensRegionId", required: true)]
        public Input<string> EnsRegionId { get; set; } = null!;

        /// <summary>
        /// NAT specifications. Value: `enat.default`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The name of the NAT gateway. The length is 1 to 128 characters, but it cannot start with 'http:// 'or 'https.
        /// </summary>
        [Input("natName")]
        public Input<string>? NatName { get; set; }

        /// <summary>
        /// The network ID.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// The vSwitch ID.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public NatGatewayArgs()
        {
        }
        public static new NatGatewayArgs Empty => new NatGatewayArgs();
    }

    public sealed class NatGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time. UTC time, in the format of YYYY-MM-DDThh:mm:ssZ.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The ID of the ENS node.
        /// </summary>
        [Input("ensRegionId")]
        public Input<string>? EnsRegionId { get; set; }

        /// <summary>
        /// NAT specifications. Value: `enat.default`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The name of the NAT gateway. The length is 1 to 128 characters, but it cannot start with 'http:// 'or 'https.
        /// </summary>
        [Input("natName")]
        public Input<string>? NatName { get; set; }

        /// <summary>
        /// The network ID.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The vSwitch ID.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public NatGatewayState()
        {
        }
        public static new NatGatewayState Empty => new NatGatewayState();
    }
}
