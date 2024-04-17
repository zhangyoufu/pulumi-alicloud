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
    /// Provides a CEN cross-regional interconnection bandwidth resource. To connect networks in different regions, you must set cross-region interconnection bandwidth after buying a bandwidth package. The total bandwidth set for all the interconnected regions of a bandwidth package cannot exceed the bandwidth of the bandwidth package. By default, 1 Kbps bandwidth is provided for connectivity test. To run normal business, you must buy a bandwidth package and set a proper interconnection bandwidth.
    /// 
    /// For example, a CEN instance is bound to a bandwidth package of 20 Mbps and  the interconnection areas are Mainland China and North America. You can set the cross-region interconnection bandwidth between US West 1 and China East 1, China East 2, China South 1, and so on. However, the total bandwidth set for all the interconnected regions cannot exceed 20  Mbps.
    /// 
    /// For information about CEN and how to use it, see [Cross-region interconnection bandwidth](https://www.alibabacloud.com/help/doc-detail/65983.htm)
    /// 
    /// &gt; **NOTE:** Available since v1.18.0.
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
    ///     var region1 = config.Get("region1") ?? "eu-central-1";
    ///     var region2 = config.Get("region2") ?? "ap-southeast-1";
    ///     var vpc1 = new AliCloud.Vpc.Network("vpc1", new()
    ///     {
    ///         VpcName = "tf-example",
    ///         CidrBlock = "192.168.0.0/16",
    ///     });
    /// 
    ///     var vpc2 = new AliCloud.Vpc.Network("vpc2", new()
    ///     {
    ///         VpcName = "tf-example",
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var example = new AliCloud.Cen.Instance("example", new()
    ///     {
    ///         CenInstanceName = "tf_example",
    ///         Description = "an example for cen",
    ///     });
    /// 
    ///     var example1 = new AliCloud.Cen.InstanceAttachment("example1", new()
    ///     {
    ///         InstanceId = example.Id,
    ///         ChildInstanceId = vpc1.Id,
    ///         ChildInstanceType = "VPC",
    ///         ChildInstanceRegionId = region1,
    ///     });
    /// 
    ///     var example2 = new AliCloud.Cen.InstanceAttachment("example2", new()
    ///     {
    ///         InstanceId = example.Id,
    ///         ChildInstanceId = vpc2.Id,
    ///         ChildInstanceType = "VPC",
    ///         ChildInstanceRegionId = region2,
    ///     });
    /// 
    ///     var exampleBandwidthPackage = new AliCloud.Cen.BandwidthPackage("example", new()
    ///     {
    ///         Bandwidth = 5,
    ///         CenBandwidthPackageName = "tf_example",
    ///         GeographicRegionAId = "Europe",
    ///         GeographicRegionBId = "Asia-Pacific",
    ///     });
    /// 
    ///     var exampleBandwidthPackageAttachment = new AliCloud.Cen.BandwidthPackageAttachment("example", new()
    ///     {
    ///         InstanceId = example.Id,
    ///         BandwidthPackageId = exampleBandwidthPackage.Id,
    ///     });
    /// 
    ///     var exampleBandwidthLimit = new AliCloud.Cen.BandwidthLimit("example", new()
    ///     {
    ///         InstanceId = exampleBandwidthPackageAttachment.InstanceId,
    ///         RegionIds = new[]
    ///         {
    ///             example1.ChildInstanceRegionId,
    ///             example2.ChildInstanceRegionId,
    ///         },
    ///         Limit = 4,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// CEN bandwidth limit can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/bandwidthLimit:BandwidthLimit example cen-abc123456:cn-beijing:eu-west-1
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/bandwidthLimit:BandwidthLimit")]
    public partial class BandwidthLimit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth configured for the interconnected regions communication.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_bandwidthlimit" resource depends on the related "alicloud.cen.BandwidthPackageAttachment" resource and "alicloud.cen.InstanceAttachment" resource.
        /// </summary>
        [Output("bandwidthLimit")]
        public Output<int> Limit { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// List of the two regions to interconnect. Must be two different regions.
        /// </summary>
        [Output("regionIds")]
        public Output<ImmutableArray<string>> RegionIds { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthLimit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthLimit(string name, BandwidthLimitArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthLimit:BandwidthLimit", name, args ?? new BandwidthLimitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthLimit(string name, Input<string> id, BandwidthLimitState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthLimit:BandwidthLimit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthLimit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthLimit Get(string name, Input<string> id, BandwidthLimitState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthLimit(name, id, state, options);
        }
    }

    public sealed class BandwidthLimitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth configured for the interconnected regions communication.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_bandwidthlimit" resource depends on the related "alicloud.cen.BandwidthPackageAttachment" resource and "alicloud.cen.InstanceAttachment" resource.
        /// </summary>
        [Input("bandwidthLimit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("regionIds", required: true)]
        private InputList<string>? _regionIds;

        /// <summary>
        /// List of the two regions to interconnect. Must be two different regions.
        /// </summary>
        public InputList<string> RegionIds
        {
            get => _regionIds ?? (_regionIds = new InputList<string>());
            set => _regionIds = value;
        }

        public BandwidthLimitArgs()
        {
        }
        public static new BandwidthLimitArgs Empty => new BandwidthLimitArgs();
    }

    public sealed class BandwidthLimitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth configured for the interconnected regions communication.
        /// 
        /// -&gt;**NOTE:** The "alicloud_cen_bandwidthlimit" resource depends on the related "alicloud.cen.BandwidthPackageAttachment" resource and "alicloud.cen.InstanceAttachment" resource.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("regionIds")]
        private InputList<string>? _regionIds;

        /// <summary>
        /// List of the two regions to interconnect. Must be two different regions.
        /// </summary>
        public InputList<string> RegionIds
        {
            get => _regionIds ?? (_regionIds = new InputList<string>());
            set => _regionIds = value;
        }

        public BandwidthLimitState()
        {
        }
        public static new BandwidthLimitState Empty => new BandwidthLimitState();
    }
}
