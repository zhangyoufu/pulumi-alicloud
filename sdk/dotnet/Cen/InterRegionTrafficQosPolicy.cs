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
    /// Provides a Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy resource.
    /// 
    /// For information about Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy and how to use it, see [What is Inter Region Traffic Qos Policy](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createceninterregiontrafficqospolicy).
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
    ///     var bj = new AliCloud.Provider("bj", new()
    ///     {
    ///         Region = "cn-beijing",
    ///     });
    /// 
    ///     var hz = new AliCloud.Provider("hz", new()
    ///     {
    ///         Region = "cn-hangzhou",
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new()
    ///     {
    ///         CenInstanceName = "tf-example",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var defaultBandwidthPackage = new AliCloud.Cen.BandwidthPackage("defaultBandwidthPackage", new()
    ///     {
    ///         Bandwidth = 5,
    ///         GeographicRegionAId = "China",
    ///         GeographicRegionBId = "China",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var defaultBandwidthPackageAttachment = new AliCloud.Cen.BandwidthPackageAttachment("defaultBandwidthPackageAttachment", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         BandwidthPackageId = defaultBandwidthPackage.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var hzTransitRouter = new AliCloud.Cen.TransitRouter("hzTransitRouter", new()
    ///     {
    ///         CenId = defaultBandwidthPackageAttachment.InstanceId,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var bjTransitRouter = new AliCloud.Cen.TransitRouter("bjTransitRouter", new()
    ///     {
    ///         CenId = hzTransitRouter.CenId,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Bj,
    ///     });
    /// 
    ///     var defaultTransitRouterPeerAttachment = new AliCloud.Cen.TransitRouterPeerAttachment("defaultTransitRouterPeerAttachment", new()
    ///     {
    ///         CenId = defaultInstance.Id,
    ///         TransitRouterId = hzTransitRouter.TransitRouterId,
    ///         PeerTransitRouterRegionId = "cn-beijing",
    ///         PeerTransitRouterId = bjTransitRouter.TransitRouterId,
    ///         CenBandwidthPackageId = defaultBandwidthPackageAttachment.BandwidthPackageId,
    ///         Bandwidth = 5,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Hz,
    ///     });
    /// 
    ///     var defaultInterRegionTrafficQosPolicy = new AliCloud.Cen.InterRegionTrafficQosPolicy("defaultInterRegionTrafficQosPolicy", new()
    ///     {
    ///         TransitRouterId = hzTransitRouter.TransitRouterId,
    ///         TransitRouterAttachmentId = defaultTransitRouterPeerAttachment.TransitRouterAttachmentId,
    ///         InterRegionTrafficQosPolicyName = "tf-example-name",
    ///         InterRegionTrafficQosPolicyDescription = "tf-example-description",
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
    /// Cloud Enterprise Network (CEN) Inter Region Traffic Qos Policy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy")]
    public partial class InterRegionTrafficQosPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
        /// </summary>
        [Output("interRegionTrafficQosPolicyDescription")]
        public Output<string?> InterRegionTrafficQosPolicyDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Output("interRegionTrafficQosPolicyName")]
        public Output<string?> InterRegionTrafficQosPolicyName { get; private set; } = null!;

        /// <summary>
        /// The status of the Inter Region Traffic Qos Policy.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the inter-region connection.
        /// </summary>
        [Output("transitRouterAttachmentId")]
        public Output<string> TransitRouterAttachmentId { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Output("transitRouterId")]
        public Output<string> TransitRouterId { get; private set; } = null!;


        /// <summary>
        /// Create a InterRegionTrafficQosPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterRegionTrafficQosPolicy(string name, InterRegionTrafficQosPolicyArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy", name, args ?? new InterRegionTrafficQosPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterRegionTrafficQosPolicy(string name, Input<string> id, InterRegionTrafficQosPolicyState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/interRegionTrafficQosPolicy:InterRegionTrafficQosPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterRegionTrafficQosPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterRegionTrafficQosPolicy Get(string name, Input<string> id, InterRegionTrafficQosPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new InterRegionTrafficQosPolicy(name, id, state, options);
        }
    }

    public sealed class InterRegionTrafficQosPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
        /// </summary>
        [Input("interRegionTrafficQosPolicyDescription")]
        public Input<string>? InterRegionTrafficQosPolicyDescription { get; set; }

        /// <summary>
        /// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Input("interRegionTrafficQosPolicyName")]
        public Input<string>? InterRegionTrafficQosPolicyName { get; set; }

        /// <summary>
        /// The ID of the inter-region connection.
        /// </summary>
        [Input("transitRouterAttachmentId", required: true)]
        public Input<string> TransitRouterAttachmentId { get; set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        public InterRegionTrafficQosPolicyArgs()
        {
        }
        public static new InterRegionTrafficQosPolicyArgs Empty => new InterRegionTrafficQosPolicyArgs();
    }

    public sealed class InterRegionTrafficQosPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the QoS policy. The description must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The description must start with a letter.
        /// </summary>
        [Input("interRegionTrafficQosPolicyDescription")]
        public Input<string>? InterRegionTrafficQosPolicyDescription { get; set; }

        /// <summary>
        /// The name of the QoS policy. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
        /// </summary>
        [Input("interRegionTrafficQosPolicyName")]
        public Input<string>? InterRegionTrafficQosPolicyName { get; set; }

        /// <summary>
        /// The status of the Inter Region Traffic Qos Policy.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the inter-region connection.
        /// </summary>
        [Input("transitRouterAttachmentId")]
        public Input<string>? TransitRouterAttachmentId { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public InterRegionTrafficQosPolicyState()
        {
        }
        public static new InterRegionTrafficQosPolicyState Empty => new InterRegionTrafficQosPolicyState();
    }
}
