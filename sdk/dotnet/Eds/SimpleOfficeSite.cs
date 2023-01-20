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
    /// Provides a ECD Simple Office Site resource.
    /// 
    /// For information about ECD Simple Office Site and how to use it, see [What is Simple Office Site](https://help.aliyun.com/document_detail/188382.html).
    /// 
    /// &gt; **NOTE:** Available in v1.140.0+.
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
    ///     var @default = new AliCloud.Eds.SimpleOfficeSite("default", new()
    ///     {
    ///         Bandwidth = 5,
    ///         CidrBlock = "172.16.0.0/12",
    ///         DesktopAccessType = "Internet",
    ///         OfficeSiteName = "site_name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Simple Office Site can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:eds/simpleOfficeSite:SimpleOfficeSite example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/simpleOfficeSite:SimpleOfficeSite")]
    public partial class SimpleOfficeSite : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Cloud Enterprise Network Instance ID.
        /// </summary>
        [Output("cenId")]
        public Output<string?> CenId { get; private set; } = null!;

        /// <summary>
        /// The cen owner id.
        /// </summary>
        [Output("cenOwnerId")]
        public Output<string?> CenOwnerId { get; private set; } = null!;

        /// <summary>
        /// Workspace Corresponds to the Security Office Network of IPv4 Segment.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
        /// </summary>
        [Output("desktopAccessType")]
        public Output<string> DesktopAccessType { get; private set; } = null!;

        /// <summary>
        /// Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
        /// </summary>
        [Output("enableAdminAccess")]
        public Output<bool> EnableAdminAccess { get; private set; } = null!;

        /// <summary>
        /// Enable Cross-Desktop Access.
        /// </summary>
        [Output("enableCrossDesktopAccess")]
        public Output<bool> EnableCrossDesktopAccess { get; private set; } = null!;

        /// <summary>
        /// Whether the Open Internet Access Function.
        /// </summary>
        [Output("enableInternetAccess")]
        public Output<bool> EnableInternetAccess { get; private set; } = null!;

        /// <summary>
        /// Whether to Enable Multi-Factor Authentication MFA.
        /// </summary>
        [Output("mfaEnabled")]
        public Output<bool> MfaEnabled { get; private set; } = null!;

        /// <summary>
        /// The office site name.
        /// </summary>
        [Output("officeSiteName")]
        public Output<string?> OfficeSiteName { get; private set; } = null!;

        /// <summary>
        /// Whether to Enable Single Sign-on (SSO) for User-Based SSO.
        /// </summary>
        [Output("ssoEnabled")]
        public Output<bool> SsoEnabled { get; private set; } = null!;

        /// <summary>
        /// Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SimpleOfficeSite resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SimpleOfficeSite(string name, SimpleOfficeSiteArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/simpleOfficeSite:SimpleOfficeSite", name, args ?? new SimpleOfficeSiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SimpleOfficeSite(string name, Input<string> id, SimpleOfficeSiteState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/simpleOfficeSite:SimpleOfficeSite", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SimpleOfficeSite resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SimpleOfficeSite Get(string name, Input<string> id, SimpleOfficeSiteState? state = null, CustomResourceOptions? options = null)
        {
            return new SimpleOfficeSite(name, id, state, options);
        }
    }

    public sealed class SimpleOfficeSiteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Cloud Enterprise Network Instance ID.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The cen owner id.
        /// </summary>
        [Input("cenOwnerId")]
        public Input<string>? CenOwnerId { get; set; }

        /// <summary>
        /// Workspace Corresponds to the Security Office Network of IPv4 Segment.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Enable Cross-Desktop Access.
        /// </summary>
        [Input("enableCrossDesktopAccess")]
        public Input<bool>? EnableCrossDesktopAccess { get; set; }

        /// <summary>
        /// Whether the Open Internet Access Function.
        /// </summary>
        [Input("enableInternetAccess")]
        public Input<bool>? EnableInternetAccess { get; set; }

        /// <summary>
        /// Whether to Enable Multi-Factor Authentication MFA.
        /// </summary>
        [Input("mfaEnabled")]
        public Input<bool>? MfaEnabled { get; set; }

        /// <summary>
        /// The office site name.
        /// </summary>
        [Input("officeSiteName")]
        public Input<string>? OfficeSiteName { get; set; }

        /// <summary>
        /// Whether to Enable Single Sign-on (SSO) for User-Based SSO.
        /// </summary>
        [Input("ssoEnabled")]
        public Input<bool>? SsoEnabled { get; set; }

        public SimpleOfficeSiteArgs()
        {
        }
        public static new SimpleOfficeSiteArgs Empty => new SimpleOfficeSiteArgs();
    }

    public sealed class SimpleOfficeSiteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Cloud Enterprise Network Instance ID.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The cen owner id.
        /// </summary>
        [Input("cenOwnerId")]
        public Input<string>? CenOwnerId { get; set; }

        /// <summary>
        /// Workspace Corresponds to the Security Office Network of IPv4 Segment.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Enable Cross-Desktop Access.
        /// </summary>
        [Input("enableCrossDesktopAccess")]
        public Input<bool>? EnableCrossDesktopAccess { get; set; }

        /// <summary>
        /// Whether the Open Internet Access Function.
        /// </summary>
        [Input("enableInternetAccess")]
        public Input<bool>? EnableInternetAccess { get; set; }

        /// <summary>
        /// Whether to Enable Multi-Factor Authentication MFA.
        /// </summary>
        [Input("mfaEnabled")]
        public Input<bool>? MfaEnabled { get; set; }

        /// <summary>
        /// The office site name.
        /// </summary>
        [Input("officeSiteName")]
        public Input<string>? OfficeSiteName { get; set; }

        /// <summary>
        /// Whether to Enable Single Sign-on (SSO) for User-Based SSO.
        /// </summary>
        [Input("ssoEnabled")]
        public Input<bool>? SsoEnabled { get; set; }

        /// <summary>
        /// Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SimpleOfficeSiteState()
        {
        }
        public static new SimpleOfficeSiteState Empty => new SimpleOfficeSiteState();
    }
}
