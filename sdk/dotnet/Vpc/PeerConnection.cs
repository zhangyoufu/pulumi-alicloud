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
    /// Provides a VPC Peer Connection resource.
    /// 
    /// For information about VPC Peer Connection and how to use it, see [What is Peer Connection](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createvpcpeer).
    /// 
    /// &gt; **NOTE:** Available since v1.186.0.
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
    ///     var @default = AliCloud.GetAccount.Invoke();
    /// 
    ///     var config = new Config();
    ///     var acceptingRegion = config.Get("acceptingRegion") ?? "cn-beijing";
    ///     var localVpc = new AliCloud.Vpc.Network("local_vpc", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var acceptingVpc = new AliCloud.Vpc.Network("accepting_vpc", new()
    ///     {
    ///         VpcName = "terraform-example",
    ///         CidrBlock = "172.17.3.0/24",
    ///     });
    /// 
    ///     var defaultPeerConnection = new AliCloud.Vpc.PeerConnection("default", new()
    ///     {
    ///         PeerConnectionName = "terraform-example",
    ///         VpcId = localVpc.Id,
    ///         AcceptingAliUid = @default.Apply(@default =&gt; @default.Apply(getAccountResult =&gt; getAccountResult.Id)),
    ///         AcceptingRegionId = acceptingRegion,
    ///         AcceptingVpcId = acceptingVpc.Id,
    ///         Description = "terraform-example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Peer Connection can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:vpc/peerConnection:PeerConnection example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/peerConnection:PeerConnection")]
    public partial class PeerConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
        /// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
        /// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
        /// &gt; **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
        /// </summary>
        [Output("acceptingAliUid")]
        public Output<int?> AcceptingAliUid { get; private set; } = null!;

        /// <summary>
        /// The region ID of the recipient of the VPC peering connection to be created.
        /// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
        /// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
        /// </summary>
        [Output("acceptingRegionId")]
        public Output<string> AcceptingRegionId { get; private set; } = null!;

        /// <summary>
        /// The VPC ID of the receiving end of the VPC peer connection.
        /// </summary>
        [Output("acceptingVpcId")]
        public Output<string> AcceptingVpcId { get; private set; } = null!;

        /// <summary>
        /// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to PreCheck only this request. Default value: `false`. Valid values:
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Output("peerConnectionName")]
        public Output<string?> PeerConnectionName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of the VPC peer connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a PeerConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PeerConnection(string name, PeerConnectionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/peerConnection:PeerConnection", name, args ?? new PeerConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PeerConnection(string name, Input<string> id, PeerConnectionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/peerConnection:PeerConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PeerConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PeerConnection Get(string name, Input<string> id, PeerConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new PeerConnection(name, id, state, options);
        }
    }

    public sealed class PeerConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
        /// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
        /// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
        /// &gt; **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
        /// </summary>
        [Input("acceptingAliUid")]
        public Input<int>? AcceptingAliUid { get; set; }

        /// <summary>
        /// The region ID of the recipient of the VPC peering connection to be created.
        /// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
        /// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
        /// </summary>
        [Input("acceptingRegionId", required: true)]
        public Input<string> AcceptingRegionId { get; set; } = null!;

        /// <summary>
        /// The VPC ID of the receiving end of the VPC peer connection.
        /// </summary>
        [Input("acceptingVpcId", required: true)]
        public Input<string> AcceptingVpcId { get; set; } = null!;

        /// <summary>
        /// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request. Default value: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("peerConnectionName")]
        public Input<string>? PeerConnectionName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the VPC peer connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public PeerConnectionArgs()
        {
        }
        public static new PeerConnectionArgs Empty => new PeerConnectionArgs();
    }

    public sealed class PeerConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
        /// - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
        /// - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
        /// &gt; **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
        /// </summary>
        [Input("acceptingAliUid")]
        public Input<int>? AcceptingAliUid { get; set; }

        /// <summary>
        /// The region ID of the recipient of the VPC peering connection to be created.
        /// - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
        /// - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
        /// </summary>
        [Input("acceptingRegionId")]
        public Input<string>? AcceptingRegionId { get; set; }

        /// <summary>
        /// The VPC ID of the receiving end of the VPC peer connection.
        /// </summary>
        [Input("acceptingVpcId")]
        public Input<string>? AcceptingVpcId { get; set; }

        /// <summary>
        /// The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to PreCheck only this request. Default value: `false`. Valid values:
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
        /// </summary>
        [Input("peerConnectionName")]
        public Input<string>? PeerConnectionName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the VPC peer connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the requester VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public PeerConnectionState()
        {
        }
        public static new PeerConnectionState Empty => new PeerConnectionState();
    }
}
