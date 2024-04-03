// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudFirewall
{
    /// <summary>
    /// Provides a Cloud Firewall Instance resource.
    /// 
    /// For information about Cloud Firewall Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/product/90174.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.139.0.
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
    ///     var @default = new AliCloud.CloudFirewall.Instance("default", new()
    ///     {
    ///         BandWidth = 200,
    ///         CfwLog = true,
    ///         CfwLogStorage = 1000,
    ///         IpNumber = 400,
    ///         PaymentType = "PayAsYouGo",
    ///         Spec = "ultimate_version",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Firewall Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cloudfirewall/instance:Instance example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudfirewall/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The number of multi account. It will be ignored when `cfw_account = false`.
        /// </summary>
        [Output("accountNumber")]
        public Output<int?> AccountNumber { get; private set; } = null!;

        /// <summary>
        /// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
        /// </summary>
        [Output("bandWidth")]
        public Output<int> BandWidth { get; private set; } = null!;

        /// <summary>
        /// Whether to use multi-account. Valid values: `true`, `false`.
        /// </summary>
        [Output("cfwAccount")]
        public Output<bool?> CfwAccount { get; private set; } = null!;

        /// <summary>
        /// Whether to use log audit. Valid values: `true`, `false`.
        /// </summary>
        [Output("cfwLog")]
        public Output<bool> CfwLog { get; private set; } = null!;

        /// <summary>
        /// The log storage capacity. It will be ignored when `cfw_log = false`.
        /// </summary>
        [Output("cfwLogStorage")]
        public Output<int?> CfwLogStorage { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The end time.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The number of protected VPCs. It will be ignored when `spec = "premium_version"`. Valid values between 2 and 500.
        /// </summary>
        [Output("fwVpcNumber")]
        public Output<int?> FwVpcNumber { get; private set; } = null!;

        /// <summary>
        /// The number of assets.
        /// </summary>
        [Output("instanceCount")]
        public Output<int?> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// The number of public IPs that can be protected. Valid values: 20 to 4000.
        /// </summary>
        [Output("ipNumber")]
        public Output<int> IpNumber { get; private set; } = null!;

        /// <summary>
        /// The logistics.
        /// </summary>
        [Output("logistics")]
        public Output<string?> Logistics { get; private set; } = null!;

        /// <summary>
        /// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
        /// </summary>
        [Output("modifyType")]
        public Output<string?> ModifyType { get; private set; } = null!;

        /// <summary>
        /// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `payment_type` can be set to `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `payment_type` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The release time.
        /// </summary>
        [Output("releaseTime")]
        public Output<string> ReleaseTime { get; private set; } = null!;

        /// <summary>
        /// Automatic renewal period. Attribute `renew_period` has been deprecated since 1.209.1. Using `renewal_duration` instead.
        /// </summary>
        [Output("renewPeriod")]
        public Output<int> RenewPeriod { get; private set; } = null!;

        /// <summary>
        /// Auto-Renewal Duration. It is required under the condition that `renewal_status` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
        /// **NOTE:** `renewal_duration` takes effect only if `payment_type` is set to `Subscription`, and `renewal_status` is set to `AutoRenewal`.
        /// </summary>
        [Output("renewalDuration")]
        public Output<int> RenewalDuration { get; private set; } = null!;

        /// <summary>
        /// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
        /// </summary>
        [Output("renewalDurationUnit")]
        public Output<string?> RenewalDurationUnit { get; private set; } = null!;

        /// <summary>
        /// Whether to renew an instance automatically or not. Default to "ManualRenewal".
        /// </summary>
        [Output("renewalStatus")]
        public Output<string> RenewalStatus { get; private set; } = null!;

        /// <summary>
        /// Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// The status of Instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudfirewall/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudfirewall/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of multi account. It will be ignored when `cfw_account = false`.
        /// </summary>
        [Input("accountNumber")]
        public Input<int>? AccountNumber { get; set; }

        /// <summary>
        /// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
        /// </summary>
        [Input("bandWidth", required: true)]
        public Input<int> BandWidth { get; set; } = null!;

        /// <summary>
        /// Whether to use multi-account. Valid values: `true`, `false`.
        /// </summary>
        [Input("cfwAccount")]
        public Input<bool>? CfwAccount { get; set; }

        /// <summary>
        /// Whether to use log audit. Valid values: `true`, `false`.
        /// </summary>
        [Input("cfwLog", required: true)]
        public Input<bool> CfwLog { get; set; } = null!;

        /// <summary>
        /// The log storage capacity. It will be ignored when `cfw_log = false`.
        /// </summary>
        [Input("cfwLogStorage")]
        public Input<int>? CfwLogStorage { get; set; }

        /// <summary>
        /// The number of protected VPCs. It will be ignored when `spec = "premium_version"`. Valid values between 2 and 500.
        /// </summary>
        [Input("fwVpcNumber")]
        public Input<int>? FwVpcNumber { get; set; }

        /// <summary>
        /// The number of assets.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The number of public IPs that can be protected. Valid values: 20 to 4000.
        /// </summary>
        [Input("ipNumber", required: true)]
        public Input<int> IpNumber { get; set; } = null!;

        /// <summary>
        /// The logistics.
        /// </summary>
        [Input("logistics")]
        public Input<string>? Logistics { get; set; }

        /// <summary>
        /// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
        /// </summary>
        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `payment_type` can be set to `PayAsYouGo`.
        /// </summary>
        [Input("paymentType", required: true)]
        public Input<string> PaymentType { get; set; } = null!;

        /// <summary>
        /// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `payment_type` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Automatic renewal period. Attribute `renew_period` has been deprecated since 1.209.1. Using `renewal_duration` instead.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// Auto-Renewal Duration. It is required under the condition that `renewal_status` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
        /// **NOTE:** `renewal_duration` takes effect only if `payment_type` is set to `Subscription`, and `renewal_status` is set to `AutoRenewal`.
        /// </summary>
        [Input("renewalDuration")]
        public Input<int>? RenewalDuration { get; set; }

        /// <summary>
        /// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
        /// </summary>
        [Input("renewalDurationUnit")]
        public Input<string>? RenewalDurationUnit { get; set; }

        /// <summary>
        /// Whether to renew an instance automatically or not. Default to "ManualRenewal".
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of multi account. It will be ignored when `cfw_account = false`.
        /// </summary>
        [Input("accountNumber")]
        public Input<int>? AccountNumber { get; set; }

        /// <summary>
        /// Public network processing capability. Valid values: 10 to 15000. Unit: Mbps.
        /// </summary>
        [Input("bandWidth")]
        public Input<int>? BandWidth { get; set; }

        /// <summary>
        /// Whether to use multi-account. Valid values: `true`, `false`.
        /// </summary>
        [Input("cfwAccount")]
        public Input<bool>? CfwAccount { get; set; }

        /// <summary>
        /// Whether to use log audit. Valid values: `true`, `false`.
        /// </summary>
        [Input("cfwLog")]
        public Input<bool>? CfwLog { get; set; }

        /// <summary>
        /// The log storage capacity. It will be ignored when `cfw_log = false`.
        /// </summary>
        [Input("cfwLogStorage")]
        public Input<int>? CfwLogStorage { get; set; }

        /// <summary>
        /// The creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The number of protected VPCs. It will be ignored when `spec = "premium_version"`. Valid values between 2 and 500.
        /// </summary>
        [Input("fwVpcNumber")]
        public Input<int>? FwVpcNumber { get; set; }

        /// <summary>
        /// The number of assets.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The number of public IPs that can be protected. Valid values: 20 to 4000.
        /// </summary>
        [Input("ipNumber")]
        public Input<int>? IpNumber { get; set; }

        /// <summary>
        /// The logistics.
        /// </summary>
        [Input("logistics")]
        public Input<string>? Logistics { get; set; }

        /// <summary>
        /// The type of modification. Valid values: `Upgrade`, `Downgrade`.  **NOTE:** The `modify_type` is required when you execute an update operation.
        /// </summary>
        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        /// <summary>
        /// The payment type of the resource. Valid values: `Subscription`, `PayAsYouGo`. **NOTE:** From version 1.220.0, `payment_type` can be set to `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The prepaid period. Valid values: `1`, `3`, `6`, `12`, `24`, `36`. **NOTE:** 1 and 3 available since 1.204.1. If `payment_type` is set to `Subscription`, `period` is required. Otherwise, it will be ignored.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The release time.
        /// </summary>
        [Input("releaseTime")]
        public Input<string>? ReleaseTime { get; set; }

        /// <summary>
        /// Automatic renewal period. Attribute `renew_period` has been deprecated since 1.209.1. Using `renewal_duration` instead.
        /// </summary>
        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        /// <summary>
        /// Auto-Renewal Duration. It is required under the condition that `renewal_status` is `AutoRenewal`. Valid values: `1`, `2`, `3`, `6`, `12`.
        /// **NOTE:** `renewal_duration` takes effect only if `payment_type` is set to `Subscription`, and `renewal_status` is set to `AutoRenewal`.
        /// </summary>
        [Input("renewalDuration")]
        public Input<int>? RenewalDuration { get; set; }

        /// <summary>
        /// Auto-Renewal Cycle Unit Values Include: Month: Month. Year: Years. Valid values: `Month`, `Year`.
        /// </summary>
        [Input("renewalDurationUnit")]
        public Input<string>? RenewalDurationUnit { get; set; }

        /// <summary>
        /// Whether to renew an instance automatically or not. Default to "ManualRenewal".
        /// </summary>
        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        /// <summary>
        /// Current version. Valid values: `premium_version`, `enterprise_version`,`ultimate_version`.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The status of Instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
