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
    /// Provides a Global Accelerator (GA) Basic Ip Set resource.
    /// 
    /// For information about Global Accelerator (GA) Basic Ip Set and how to use it, see [What is Basic Ip Set](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createbasicipset).
    /// 
    /// &gt; **NOTE:** Available since v1.194.0.
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
    ///     var region = config.Get("region") ?? "cn-hangzhou";
    ///     var defaultBasicAccelerator = new AliCloud.Ga.BasicAccelerator("defaultBasicAccelerator", new()
    ///     {
    ///         Duration = 1,
    ///         PricingCycle = "Month",
    ///         BandwidthBillingType = "CDT",
    ///         AutoPay = true,
    ///         AutoUseCoupon = "true",
    ///         AutoRenew = false,
    ///         AutoRenewDuration = 1,
    ///     });
    /// 
    ///     var defaultBasicIpSet = new AliCloud.Ga.BasicIpSet("defaultBasicIpSet", new()
    ///     {
    ///         AcceleratorId = defaultBasicAccelerator.Id,
    ///         AccelerateRegionId = region,
    ///         IspType = "BGP",
    ///         Bandwidth = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Global Accelerator (GA) Basic Ip Set can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/basicIpSet:BasicIpSet example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/basicIpSet:BasicIpSet")]
    public partial class BasicIpSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the acceleration region.
        /// </summary>
        [Output("accelerateRegionId")]
        public Output<string> AccelerateRegionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Output("acceleratorId")]
        public Output<string> AcceleratorId { get; private set; } = null!;

        /// <summary>
        /// The bandwidth of the acceleration region. Unit: Mbit/s.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The line type of the elastic IP address (EIP) in the acceleration region. Default value: `BGP`. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`.
        /// </summary>
        [Output("ispType")]
        public Output<string> IspType { get; private set; } = null!;

        /// <summary>
        /// The status of the Basic Ip Set instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BasicIpSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BasicIpSet(string name, BasicIpSetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicIpSet:BasicIpSet", name, args ?? new BasicIpSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BasicIpSet(string name, Input<string> id, BasicIpSetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/basicIpSet:BasicIpSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BasicIpSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BasicIpSet Get(string name, Input<string> id, BasicIpSetState? state = null, CustomResourceOptions? options = null)
        {
            return new BasicIpSet(name, id, state, options);
        }
    }

    public sealed class BasicIpSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the acceleration region.
        /// </summary>
        [Input("accelerateRegionId", required: true)]
        public Input<string> AccelerateRegionId { get; set; } = null!;

        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        /// <summary>
        /// The bandwidth of the acceleration region. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The line type of the elastic IP address (EIP) in the acceleration region. Default value: `BGP`. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`.
        /// </summary>
        [Input("ispType")]
        public Input<string>? IspType { get; set; }

        public BasicIpSetArgs()
        {
        }
        public static new BasicIpSetArgs Empty => new BasicIpSetArgs();
    }

    public sealed class BasicIpSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the acceleration region.
        /// </summary>
        [Input("accelerateRegionId")]
        public Input<string>? AccelerateRegionId { get; set; }

        /// <summary>
        /// The ID of the basic GA instance.
        /// </summary>
        [Input("acceleratorId")]
        public Input<string>? AcceleratorId { get; set; }

        /// <summary>
        /// The bandwidth of the acceleration region. Unit: Mbit/s.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The line type of the elastic IP address (EIP) in the acceleration region. Default value: `BGP`. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`.
        /// </summary>
        [Input("ispType")]
        public Input<string>? IspType { get; set; }

        /// <summary>
        /// The status of the Basic Ip Set instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BasicIpSetState()
        {
        }
        public static new BasicIpSetState Empty => new BasicIpSetState();
    }
}
