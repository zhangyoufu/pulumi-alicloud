// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a Cen Transit Router Multicast Domain Peer Member resource.
    /// 
    /// For information about Cen Transit Router Multicast Domain Peer Member and how to use it, see [What is Transit Router Multicast Domain Peer Member](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-cbn-2017-09-12-deregistertransitroutermulticastgroupmembers).
    /// 
    /// &gt; **NOTE:** Available since v1.195.0.
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
    ///     var name = config.Get("name") ?? "tf_example";
    ///     var hz = new AliCloud.Provider("hz", new()
    ///     {
    ///         Region = "cn-hangzhou",
    ///     });
    /// 
    ///     var qd = new AliCloud.Provider("qd", new()
    ///     {
    ///         Region = "cn-qingdao",
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new()
    ///     {
    ///         CenInstanceName = name,
    ///     });
    /// 
    ///     var defaultBandwidthPackage = new AliCloud.Cen.BandwidthPackage("defaultBandwidthPackage", new()
    ///     {
    ///         Bandwidth = 5,
    ///         CenBandwidthPackageName = name,
    ///         GeographicRegionAId = "China",
    ///         GeographicRegionBId = "China",
    ///     });
    /// 
    ///     var defaultBandwidthPackageAttachment = new AliCloud.Cen.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         BandwidthPackageId = defaultBandwidthPackage.Id,
    ///     });
    /// 
    ///     var defaultTransitRouter = new AliCloud.Cen.TransitRouter("defaultTransitRouter", new()
    ///     {
    ///         CenId = defaultBandwidthPackageAttachment.InstanceId,
    ///         SupportMulticast = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var peerTransitRouter = new AliCloud.Cen.TransitRouter("peerTransitRouter", new()
    ///     {
    ///         CenId = defaultBandwidthPackageAttachment.InstanceId,
    ///         SupportMulticast = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Qd,
    ///     });
    /// 
    ///     var defaultTransitRouterPeerAttachment = new AliCloud.Cen.TransitRouterPeerAttachment("defaultTransitRouterPeerAttachment", new()
    ///     {
    ///         CenId = defaultBandwidthPackageAttachment.InstanceId,
    ///         TransitRouterId = defaultTransitRouter.TransitRouterId,
    ///         PeerTransitRouterId = peerTransitRouter.TransitRouterId,
    ///         PeerTransitRouterRegionId = "cn-qingdao",
    ///         CenBandwidthPackageId = defaultBandwidthPackageAttachment.BandwidthPackageId,
    ///         Bandwidth = 5,
    ///         TransitRouterAttachmentDescription = name,
    ///         TransitRouterAttachmentName = name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var defaultTransitRouterMulticastDomain = new AliCloud.Cen.TransitRouterMulticastDomain("defaultTransitRouterMulticastDomain", new()
    ///     {
    ///         TransitRouterId = defaultTransitRouterPeerAttachment.TransitRouterId,
    ///         TransitRouterMulticastDomainName = name,
    ///         TransitRouterMulticastDomainDescription = name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var peerTransitRouterMulticastDomain = new AliCloud.Cen.TransitRouterMulticastDomain("peerTransitRouterMulticastDomain", new()
    ///     {
    ///         TransitRouterId = defaultTransitRouterPeerAttachment.PeerTransitRouterId,
    ///         TransitRouterMulticastDomainName = name,
    ///         TransitRouterMulticastDomainDescription = name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Qd,
    ///     });
    /// 
    ///     var defaultTransitRouterMulticastDomainPeerMember = new AliCloud.Cen.TransitRouterMulticastDomainPeerMember("defaultTransitRouterMulticastDomainPeerMember", new()
    ///     {
    ///         TransitRouterMulticastDomainId = defaultTransitRouterMulticastDomain.Id,
    ///         PeerTransitRouterMulticastDomainId = peerTransitRouterMulticastDomain.Id,
    ///         GroupIpAddress = "224.0.0.1",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cen Transit Router Multicast Domain Peer Member can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember example &lt;transit_router_multicast_domain_id&gt;:&lt;group_ip_address&gt;:&lt;peer_transit_router_multicast_domain_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember")]
    public partial class TransitRouterMulticastDomainPeerMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether only to precheck the request.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
        /// </summary>
        [Output("groupIpAddress")]
        public Output<string> GroupIpAddress { get; private set; } = null!;

        /// <summary>
        /// The IDs of the inter-region multicast domains.
        /// </summary>
        [Output("peerTransitRouterMulticastDomainId")]
        public Output<string> PeerTransitRouterMulticastDomainId { get; private set; } = null!;

        /// <summary>
        /// The status of the multicast resource. Valid values:
        /// - Registering: being created
        /// - Registered: available
        /// - Deregistering: being deleted
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the multicast domain to which the multicast member belongs.
        /// </summary>
        [Output("transitRouterMulticastDomainId")]
        public Output<string> TransitRouterMulticastDomainId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitRouterMulticastDomainPeerMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitRouterMulticastDomainPeerMember(string name, TransitRouterMulticastDomainPeerMemberArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember", name, args ?? new TransitRouterMulticastDomainPeerMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitRouterMulticastDomainPeerMember(string name, Input<string> id, TransitRouterMulticastDomainPeerMemberState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterMulticastDomainPeerMember:TransitRouterMulticastDomainPeerMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitRouterMulticastDomainPeerMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitRouterMulticastDomainPeerMember Get(string name, Input<string> id, TransitRouterMulticastDomainPeerMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitRouterMulticastDomainPeerMember(name, id, state, options);
        }
    }

    public sealed class TransitRouterMulticastDomainPeerMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether only to precheck the request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
        /// </summary>
        [Input("groupIpAddress", required: true)]
        public Input<string> GroupIpAddress { get; set; } = null!;

        /// <summary>
        /// The IDs of the inter-region multicast domains.
        /// </summary>
        [Input("peerTransitRouterMulticastDomainId", required: true)]
        public Input<string> PeerTransitRouterMulticastDomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the multicast domain to which the multicast member belongs.
        /// </summary>
        [Input("transitRouterMulticastDomainId", required: true)]
        public Input<string> TransitRouterMulticastDomainId { get; set; } = null!;

        public TransitRouterMulticastDomainPeerMemberArgs()
        {
        }
        public static new TransitRouterMulticastDomainPeerMemberArgs Empty => new TransitRouterMulticastDomainPeerMemberArgs();
    }

    public sealed class TransitRouterMulticastDomainPeerMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether only to precheck the request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The IP address of the multicast group to which the multicast member belongs. Value range: **224.0.0.1** to **239.255.255.254**.If the multicast group you specified does not exist in the current multicast domain, the system will automatically create a new multicast group for you in the current multicast domain.
        /// </summary>
        [Input("groupIpAddress")]
        public Input<string>? GroupIpAddress { get; set; }

        /// <summary>
        /// The IDs of the inter-region multicast domains.
        /// </summary>
        [Input("peerTransitRouterMulticastDomainId")]
        public Input<string>? PeerTransitRouterMulticastDomainId { get; set; }

        /// <summary>
        /// The status of the multicast resource. Valid values:
        /// - Registering: being created
        /// - Registered: available
        /// - Deregistering: being deleted
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the multicast domain to which the multicast member belongs.
        /// </summary>
        [Input("transitRouterMulticastDomainId")]
        public Input<string>? TransitRouterMulticastDomainId { get; set; }

        public TransitRouterMulticastDomainPeerMemberState()
        {
        }
        public static new TransitRouterMulticastDomainPeerMemberState Empty => new TransitRouterMulticastDomainPeerMemberState();
    }
}
