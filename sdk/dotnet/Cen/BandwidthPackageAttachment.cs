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
    /// Provides a CEN bandwidth package attachment resource. The resource can be used to bind a bandwidth package to a specified CEN instance.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new bandwidth package attachment and use it to attach a bandwidth package to a new CEN
    ///     var cen = new AliCloud.Cen.Instance("cen", new()
    ///     {
    ///         Description = "tf-testAccCenBandwidthPackageAttachmentDescription",
    ///     });
    /// 
    ///     var bwp = new AliCloud.Cen.BandwidthPackage("bwp", new()
    ///     {
    ///         Bandwidth = 20,
    ///         GeographicRegionIds = new[]
    ///         {
    ///             "China",
    ///             "Asia-Pacific",
    ///         },
    ///     });
    /// 
    ///     var foo = new AliCloud.Cen.BandwidthPackageAttachment("foo", new()
    ///     {
    ///         InstanceId = cen.Id,
    ///         BandwidthPackageId = bwp.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN bandwidth package attachment resource can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment example bwp-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment")]
    public partial class BandwidthPackageAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the bandwidth package.
        /// </summary>
        [Output("bandwidthPackageId")]
        public Output<string> BandwidthPackageId { get; private set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackageAttachment(string name, BandwidthPackageAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment", name, args ?? new BandwidthPackageAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackageAttachment(string name, Input<string> id, BandwidthPackageAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/bandwidthPackageAttachment:BandwidthPackageAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthPackageAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackageAttachment Get(string name, Input<string> id, BandwidthPackageAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackageAttachment(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the bandwidth package.
        /// </summary>
        [Input("bandwidthPackageId", required: true)]
        public Input<string> BandwidthPackageId { get; set; } = null!;

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public BandwidthPackageAttachmentArgs()
        {
        }
        public static new BandwidthPackageAttachmentArgs Empty => new BandwidthPackageAttachmentArgs();
    }

    public sealed class BandwidthPackageAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the bandwidth package.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public BandwidthPackageAttachmentState()
        {
        }
        public static new BandwidthPackageAttachmentState Empty => new BandwidthPackageAttachmentState();
    }
}
