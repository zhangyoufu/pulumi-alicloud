// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    /// <summary>
    /// Provides a Global Accelerator (GA) Bandwidth Package Attachment resource.
    /// 
    /// For information about Global Accelerator (GA) Bandwidth Package Attachment and how to use it, see [What is Bandwidth Package Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-bandwidthpackageaddaccelerator).
    /// 
    /// &gt; **NOTE:** Available since v1.113.0.
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
    ///     var exampleAccelerator = new AliCloud.Ga.Accelerator("exampleAccelerator", new()
    ///     {
    ///         Duration = 1,
    ///         AutoUseCoupon = true,
    ///         Spec = "1",
    ///     });
    /// 
    ///     var exampleBandwidthPackage = new AliCloud.Ga.BandwidthPackage("exampleBandwidthPackage", new()
    ///     {
    ///         Bandwidth = 20,
    ///         Type = "Basic",
    ///         BandwidthType = "Basic",
    ///         Duration = "1",
    ///         AutoPay = true,
    ///         Ratio = 30,
    ///     });
    /// 
    ///     var exampleBandwidthPackageAttachment = new AliCloud.Ga.BandwidthPackageAttachment("exampleBandwidthPackageAttachment", new()
    ///     {
    ///         AcceleratorId = exampleAccelerator.Id,
    ///         BandwidthPackageId = exampleBandwidthPackage.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Ga Bandwidth Package Attachment can be imported using the id. Format to `&lt;accelerator_id&gt;:&lt;bandwidth_package_id&gt;`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment example your_accelerator_id:your_bandwidth_package_id
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment")]
    public partial class BandwidthPackageAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// Accelerators bound with current Bandwidth Package.
        /// </summary>
        [Output("accelerators")]
        public Output<ImmutableArray<string>> Accelerators { get; private set; } = null!;

        /// <summary>
        /// The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
        /// </summary>
        [Output("bandwidthPackageId")]
        public Output<string> BandwidthPackageId { get; private set; } = null!;

        /// <summary>
        /// State of Bandwidth Package.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackageAttachment(string name, BandwidthPackageAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment", name, args ?? new BandwidthPackageAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackageAttachment(string name, Input<string> id, BandwidthPackageAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment", name, state, MakeResourceOptions(options, id))
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
        /// The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
        /// </summary>
        [Input("bandwidthPackageId", required: true)]
        public Input<string> BandwidthPackageId { get; set; } = null!;

        public BandwidthPackageAttachmentArgs()
        {
        }
        public static new BandwidthPackageAttachmentArgs Empty => new BandwidthPackageAttachmentArgs();
    }

    public sealed class BandwidthPackageAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance from which you want to disassociate the bandwidth plan.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        [Input("accelerators")]
        private InputList<string>? _accelerators;

        /// <summary>
        /// Accelerators bound with current Bandwidth Package.
        /// </summary>
        public InputList<string> Accelerators
        {
            get => _accelerators ?? (_accelerators = new InputList<string>());
            set => _accelerators = value;
        }

        /// <summary>
        /// The ID of the bandwidth plan to disassociate. **NOTE:** From version 1.192.0, `bandwidth_package_id` can be modified.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// State of Bandwidth Package.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BandwidthPackageAttachmentState()
        {
        }
        public static new BandwidthPackageAttachmentState Empty => new BandwidthPackageAttachmentState();
    }
}
