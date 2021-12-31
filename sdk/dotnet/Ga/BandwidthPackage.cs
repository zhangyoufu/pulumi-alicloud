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
    /// Provides a Global Accelerator (GA) Bandwidth Package resource.
    /// 
    /// For information about Global Accelerator (GA) Bandwidth Package and how to use it, see [What is Bandwidth Package](https://www.alibabacloud.com/help/en/doc-detail/153241.htm).
    /// 
    /// &gt; **NOTE:** At present, The `alicloud.ga.BandwidthPackage` created with `Subscription` cannot be deleted. you need to wait until the resource is outdated and released automatically.
    /// 
    /// &gt; **NOTE:** Available in v1.112.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.Ga.BandwidthPackage("example", new AliCloud.Ga.BandwidthPackageArgs
    ///         {
    ///             AutoPay = true,
    ///             Bandwidth = 20,
    ///             BandwidthType = "Basic",
    ///             Duration = "1",
    ///             Ratio = 30,
    ///             Type = "Basic",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Ga Bandwidth Package can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ga/bandwidthPackage:BandwidthPackage example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ga/bandwidthPackage:BandwidthPackage")]
    public partial class BandwidthPackage : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to pay automatically. Valid values:
        /// `false`: If automatic payment is not enabled, you need to go to the order center to complete the payment after the order is generated.
        /// `true`: Enable automatic payment, automatic payment order.
        /// </summary>
        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// Whether use vouchers. Default value is `false`. Valid values: `false`: Not used, `true`: Use.
        /// </summary>
        [Output("autoUseCoupon")]
        public Output<bool?> AutoUseCoupon { get; private set; } = null!;

        /// <summary>
        /// The bandwidth value of bandwidth packet.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The name of the bandwidth packet.
        /// </summary>
        [Output("bandwidthPackageName")]
        public Output<string?> BandwidthPackageName { get; private set; } = null!;

        /// <summary>
        /// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
        /// </summary>
        [Output("bandwidthType")]
        public Output<string?> BandwidthType { get; private set; } = null!;

        /// <summary>
        /// The billing type. Valid values: `PayBy95`, `PayByTraffic`.
        /// </summary>
        [Output("billingType")]
        public Output<string?> BillingType { get; private set; } = null!;

        /// <summary>
        /// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value is `China-mainland`.
        /// </summary>
        [Output("cbnGeographicRegionIda")]
        public Output<string?> CbnGeographicRegionIda { get; private set; } = null!;

        /// <summary>
        /// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value is `Global`.
        /// </summary>
        [Output("cbnGeographicRegionIdb")]
        public Output<string?> CbnGeographicRegionIdb { get; private set; } = null!;

        /// <summary>
        /// The description of bandwidth package.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `payment_type` is set to `Subscription`, this parameter is required.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// The payment type of the bandwidth. Valid values: `PayAsYouGo`, `Subscription`. Default value is `Subscription`.
        /// </summary>
        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: 30 to 100.
        /// </summary>
        [Output("ratio")]
        public Output<int?> Ratio { get; private set; } = null!;

        /// <summary>
        /// The status of the bandwidth plan.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackage(string name, BandwidthPackageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ga/bandwidthPackage:BandwidthPackage", name, args ?? new BandwidthPackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackage(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ga/bandwidthPackage:BandwidthPackage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackage Get(string name, Input<string> id, BandwidthPackageState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackage(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to pay automatically. Valid values:
        /// `false`: If automatic payment is not enabled, you need to go to the order center to complete the payment after the order is generated.
        /// `true`: Enable automatic payment, automatic payment order.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Whether use vouchers. Default value is `false`. Valid values: `false`: Not used, `true`: Use.
        /// </summary>
        [Input("autoUseCoupon")]
        public Input<bool>? AutoUseCoupon { get; set; }

        /// <summary>
        /// The bandwidth value of bandwidth packet.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<int> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The name of the bandwidth packet.
        /// </summary>
        [Input("bandwidthPackageName")]
        public Input<string>? BandwidthPackageName { get; set; }

        /// <summary>
        /// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
        /// </summary>
        [Input("bandwidthType")]
        public Input<string>? BandwidthType { get; set; }

        /// <summary>
        /// The billing type. Valid values: `PayBy95`, `PayByTraffic`.
        /// </summary>
        [Input("billingType")]
        public Input<string>? BillingType { get; set; }

        /// <summary>
        /// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value is `China-mainland`.
        /// </summary>
        [Input("cbnGeographicRegionIda")]
        public Input<string>? CbnGeographicRegionIda { get; set; }

        /// <summary>
        /// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value is `Global`.
        /// </summary>
        [Input("cbnGeographicRegionIdb")]
        public Input<string>? CbnGeographicRegionIdb { get; set; }

        /// <summary>
        /// The description of bandwidth package.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `payment_type` is set to `Subscription`, this parameter is required.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// The payment type of the bandwidth. Valid values: `PayAsYouGo`, `Subscription`. Default value is `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: 30 to 100.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BandwidthPackageArgs()
        {
        }
    }

    public sealed class BandwidthPackageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to pay automatically. Valid values:
        /// `false`: If automatic payment is not enabled, you need to go to the order center to complete the payment after the order is generated.
        /// `true`: Enable automatic payment, automatic payment order.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// Whether use vouchers. Default value is `false`. Valid values: `false`: Not used, `true`: Use.
        /// </summary>
        [Input("autoUseCoupon")]
        public Input<bool>? AutoUseCoupon { get; set; }

        /// <summary>
        /// The bandwidth value of bandwidth packet.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The name of the bandwidth packet.
        /// </summary>
        [Input("bandwidthPackageName")]
        public Input<string>? BandwidthPackageName { get; set; }

        /// <summary>
        /// The bandwidth type of the bandwidth. Valid values: `Advanced`, `Basic`, `Enhanced`. If `type` is set to `Basic`, this parameter is required.
        /// </summary>
        [Input("bandwidthType")]
        public Input<string>? BandwidthType { get; set; }

        /// <summary>
        /// The billing type. Valid values: `PayBy95`, `PayByTraffic`.
        /// </summary>
        [Input("billingType")]
        public Input<string>? BillingType { get; set; }

        /// <summary>
        /// Interworking area A of cross domain acceleration package. Only international stations support returning this parameter. Default value is `China-mainland`.
        /// </summary>
        [Input("cbnGeographicRegionIda")]
        public Input<string>? CbnGeographicRegionIda { get; set; }

        /// <summary>
        /// Interworking area B of cross domain acceleration package. Only international stations support returning this parameter. Default value is `Global`.
        /// </summary>
        [Input("cbnGeographicRegionIdb")]
        public Input<string>? CbnGeographicRegionIdb { get; set; }

        /// <summary>
        /// The description of bandwidth package.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The subscription duration. **NOTE:** The ForceNew attribute has be removed from version 1.148.0. If `payment_type` is set to `Subscription`, this parameter is required.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// The payment type of the bandwidth. Valid values: `PayAsYouGo`, `Subscription`. Default value is `Subscription`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: 30 to 100.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// The status of the bandwidth plan.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the bandwidth packet. China station only supports return to basic. Valid values: `Basic`, `CrossDomain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BandwidthPackageState()
        {
        }
    }
}
