// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    /// <summary>
    /// Provides a ECD Desktop resource.
    /// 
    /// For information about ECD Desktop and how to use it, see [What is Desktop](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createdesktops)
    /// 
    /// &gt; **NOTE:** Available since v1.144.0.
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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultRandomInteger = new Random.RandomInteger("defaultRandomInteger", new()
    ///     {
    ///         Min = 10000,
    ///         Max = 99999,
    ///     });
    /// 
    ///     var defaultSimpleOfficeSite = new AliCloud.Eds.SimpleOfficeSite("defaultSimpleOfficeSite", new()
    ///     {
    ///         CidrBlock = "172.16.0.0/12",
    ///         EnableAdminAccess = true,
    ///         DesktopAccessType = "Internet",
    ///         OfficeSiteName = defaultRandomInteger.Result.Apply(result =&gt; $"{name}-{result}"),
    ///     });
    /// 
    ///     var defaultEcdPolicyGroup = new AliCloud.Eds.EcdPolicyGroup("defaultEcdPolicyGroup", new()
    ///     {
    ///         PolicyGroupName = name,
    ///         Clipboard = "read",
    ///         LocalDrive = "read",
    ///         UsbRedirect = "off",
    ///         Watermark = "off",
    ///         AuthorizeAccessPolicyRules = new[]
    ///         {
    ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs
    ///             {
    ///                 Description = name,
    ///                 CidrIp = "1.2.3.45/24",
    ///             },
    ///         },
    ///         AuthorizeSecurityPolicyRules = new[]
    ///         {
    ///             new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs
    ///             {
    ///                 Type = "inflow",
    ///                 Policy = "accept",
    ///                 Description = name,
    ///                 PortRange = "80/80",
    ///                 IpProtocol = "TCP",
    ///                 Priority = "1",
    ///                 CidrIp = "1.2.3.4/24",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultBundles = AliCloud.Eds.GetBundles.Invoke(new()
    ///     {
    ///         BundleType = "SYSTEM",
    ///     });
    /// 
    ///     var defaultDesktop = new AliCloud.Eds.Desktop("defaultDesktop", new()
    ///     {
    ///         OfficeSiteId = defaultSimpleOfficeSite.Id,
    ///         PolicyGroupId = defaultEcdPolicyGroup.Id,
    ///         BundleId = defaultBundles.Apply(getBundlesResult =&gt; getBundlesResult.Bundles[1]?.Id),
    ///         DesktopName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ECD Desktop can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:eds/desktop:Desktop example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/desktop:Desktop")]
    public partial class Desktop : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The amount of the Desktop.
        /// </summary>
        [Output("amount")]
        public Output<int?> Amount { get; private set; } = null!;

        /// <summary>
        /// The auto-pay of the Desktop whether to pay automatically. values: `true`, `false`.
        /// </summary>
        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// The auto-renewal of the Desktop whether to renew automatically. It takes effect only when the parameter ChargeType is set to PrePaid. values: `true`, `false`.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The bundle id of the Desktop.
        /// </summary>
        [Output("bundleId")]
        public Output<string> BundleId { get; private set; } = null!;

        /// <summary>
        /// The desktop name of the Desktop.
        /// </summary>
        [Output("desktopName")]
        public Output<string?> DesktopName { get; private set; } = null!;

        /// <summary>
        /// The desktop type of the Desktop.
        /// </summary>
        [Output("desktopType")]
        public Output<string> DesktopType { get; private set; } = null!;

        /// <summary>
        /// The desktop end user id of the Desktop.
        /// </summary>
        [Output("endUserIds")]
        public Output<ImmutableArray<string>> EndUserIds { get; private set; } = null!;

        /// <summary>
        /// The hostname of the Desktop.
        /// </summary>
        [Output("hostName")]
        public Output<string?> HostName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Simple Office Site.
        /// </summary>
        [Output("officeSiteId")]
        public Output<string> OfficeSiteId { get; private set; } = null!;

        /// <summary>
        /// The payment type of the Desktop. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The period of the Desktop.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The period unit of the Desktop.
        /// </summary>
        [Output("periodUnit")]
        public Output<string?> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The policy group id of the Desktop.
        /// </summary>
        [Output("policyGroupId")]
        public Output<string> PolicyGroupId { get; private set; } = null!;

        /// <summary>
        /// The root disk size gib of the Desktop.
        /// </summary>
        [Output("rootDiskSizeGib")]
        public Output<int?> RootDiskSizeGib { get; private set; } = null!;

        /// <summary>
        /// The status of the Desktop. Valid values: `Deleted`, `Expired`, `Pending`, `Running`, `Starting`, `Stopped`, `Stopping`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The stopped mode of the Desktop.
        /// </summary>
        [Output("stoppedMode")]
        public Output<string> StoppedMode { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The user assign mode of the Desktop. Valid values: `ALL`, `PER_USER`. Default to `ALL`.
        /// </summary>
        [Output("userAssignMode")]
        public Output<string> UserAssignMode { get; private set; } = null!;

        /// <summary>
        /// The user disk size gib of the Desktop.
        /// </summary>
        [Output("userDiskSizeGib")]
        public Output<int?> UserDiskSizeGib { get; private set; } = null!;


        /// <summary>
        /// Create a Desktop resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Desktop(string name, DesktopArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/desktop:Desktop", name, args ?? new DesktopArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Desktop(string name, Input<string> id, DesktopState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/desktop:Desktop", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Desktop resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Desktop Get(string name, Input<string> id, DesktopState? state = null, CustomResourceOptions? options = null)
        {
            return new Desktop(name, id, state, options);
        }
    }

    public sealed class DesktopArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of the Desktop.
        /// </summary>
        [Input("amount")]
        public Input<int>? Amount { get; set; }

        /// <summary>
        /// The auto-pay of the Desktop whether to pay automatically. values: `true`, `false`.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// The auto-renewal of the Desktop whether to renew automatically. It takes effect only when the parameter ChargeType is set to PrePaid. values: `true`, `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The bundle id of the Desktop.
        /// </summary>
        [Input("bundleId", required: true)]
        public Input<string> BundleId { get; set; } = null!;

        /// <summary>
        /// The desktop name of the Desktop.
        /// </summary>
        [Input("desktopName")]
        public Input<string>? DesktopName { get; set; }

        /// <summary>
        /// The desktop type of the Desktop.
        /// </summary>
        [Input("desktopType")]
        public Input<string>? DesktopType { get; set; }

        [Input("endUserIds")]
        private InputList<string>? _endUserIds;

        /// <summary>
        /// The desktop end user id of the Desktop.
        /// </summary>
        public InputList<string> EndUserIds
        {
            get => _endUserIds ?? (_endUserIds = new InputList<string>());
            set => _endUserIds = value;
        }

        /// <summary>
        /// The hostname of the Desktop.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The ID of the Simple Office Site.
        /// </summary>
        [Input("officeSiteId", required: true)]
        public Input<string> OfficeSiteId { get; set; } = null!;

        /// <summary>
        /// The payment type of the Desktop. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The period of the Desktop.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The period unit of the Desktop.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The policy group id of the Desktop.
        /// </summary>
        [Input("policyGroupId", required: true)]
        public Input<string> PolicyGroupId { get; set; } = null!;

        /// <summary>
        /// The root disk size gib of the Desktop.
        /// </summary>
        [Input("rootDiskSizeGib")]
        public Input<int>? RootDiskSizeGib { get; set; }

        /// <summary>
        /// The status of the Desktop. Valid values: `Deleted`, `Expired`, `Pending`, `Running`, `Starting`, `Stopped`, `Stopping`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The stopped mode of the Desktop.
        /// </summary>
        [Input("stoppedMode")]
        public Input<string>? StoppedMode { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The user assign mode of the Desktop. Valid values: `ALL`, `PER_USER`. Default to `ALL`.
        /// </summary>
        [Input("userAssignMode")]
        public Input<string>? UserAssignMode { get; set; }

        /// <summary>
        /// The user disk size gib of the Desktop.
        /// </summary>
        [Input("userDiskSizeGib")]
        public Input<int>? UserDiskSizeGib { get; set; }

        public DesktopArgs()
        {
        }
        public static new DesktopArgs Empty => new DesktopArgs();
    }

    public sealed class DesktopState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of the Desktop.
        /// </summary>
        [Input("amount")]
        public Input<int>? Amount { get; set; }

        /// <summary>
        /// The auto-pay of the Desktop whether to pay automatically. values: `true`, `false`.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// The auto-renewal of the Desktop whether to renew automatically. It takes effect only when the parameter ChargeType is set to PrePaid. values: `true`, `false`.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The bundle id of the Desktop.
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        /// <summary>
        /// The desktop name of the Desktop.
        /// </summary>
        [Input("desktopName")]
        public Input<string>? DesktopName { get; set; }

        /// <summary>
        /// The desktop type of the Desktop.
        /// </summary>
        [Input("desktopType")]
        public Input<string>? DesktopType { get; set; }

        [Input("endUserIds")]
        private InputList<string>? _endUserIds;

        /// <summary>
        /// The desktop end user id of the Desktop.
        /// </summary>
        public InputList<string> EndUserIds
        {
            get => _endUserIds ?? (_endUserIds = new InputList<string>());
            set => _endUserIds = value;
        }

        /// <summary>
        /// The hostname of the Desktop.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The ID of the Simple Office Site.
        /// </summary>
        [Input("officeSiteId")]
        public Input<string>? OfficeSiteId { get; set; }

        /// <summary>
        /// The payment type of the Desktop. Valid values: `PayAsYouGo`, `Subscription`. Default to `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The period of the Desktop.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The period unit of the Desktop.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The policy group id of the Desktop.
        /// </summary>
        [Input("policyGroupId")]
        public Input<string>? PolicyGroupId { get; set; }

        /// <summary>
        /// The root disk size gib of the Desktop.
        /// </summary>
        [Input("rootDiskSizeGib")]
        public Input<int>? RootDiskSizeGib { get; set; }

        /// <summary>
        /// The status of the Desktop. Valid values: `Deleted`, `Expired`, `Pending`, `Running`, `Starting`, `Stopped`, `Stopping`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The stopped mode of the Desktop.
        /// </summary>
        [Input("stoppedMode")]
        public Input<string>? StoppedMode { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The user assign mode of the Desktop. Valid values: `ALL`, `PER_USER`. Default to `ALL`.
        /// </summary>
        [Input("userAssignMode")]
        public Input<string>? UserAssignMode { get; set; }

        /// <summary>
        /// The user disk size gib of the Desktop.
        /// </summary>
        [Input("userDiskSizeGib")]
        public Input<int>? UserDiskSizeGib { get; set; }

        public DesktopState()
        {
        }
        public static new DesktopState Empty => new DesktopState();
    }
}
