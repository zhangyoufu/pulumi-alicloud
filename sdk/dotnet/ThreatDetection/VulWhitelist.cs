// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    /// <summary>
    /// Provides a Threat Detection Vul Whitelist resource.
    /// 
    /// For information about Threat Detection Vul Whitelist and how to use it, see [What is Vul Whitelist](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
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
    ///     var @default = new AliCloud.ThreatDetection.VulWhitelist("default", new()
    ///     {
    ///         Whitelist = "[{\"aliasName\":\"RHSA-2021:2260: libwebp 安全更新\",\"name\":\"RHSA-2021:2260: libwebp 安全更新\",\"type\":\"cve\"}]",
    ///         TargetInfo = "{\"type\":\"GroupId\",\"uuids\":[],\"groupIds\":[10782678]}",
    ///         Reason = "tf-example-reason",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Threat Detection Vul Whitelist can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:threatdetection/vulWhitelist:VulWhitelist example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:threatdetection/vulWhitelist:VulWhitelist")]
    public partial class VulWhitelist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Reason for adding whitelist.
        /// </summary>
        [Output("reason")]
        public Output<string?> Reason { get; private set; } = null!;

        /// <summary>
        /// Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Output("targetInfo")]
        public Output<string?> TargetInfo { get; private set; } = null!;

        /// <summary>
        /// Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Output("whitelist")]
        public Output<string> Whitelist { get; private set; } = null!;


        /// <summary>
        /// Create a VulWhitelist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VulWhitelist(string name, VulWhitelistArgs args, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/vulWhitelist:VulWhitelist", name, args ?? new VulWhitelistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VulWhitelist(string name, Input<string> id, VulWhitelistState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:threatdetection/vulWhitelist:VulWhitelist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VulWhitelist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VulWhitelist Get(string name, Input<string> id, VulWhitelistState? state = null, CustomResourceOptions? options = null)
        {
            return new VulWhitelist(name, id, state, options);
        }
    }

    public sealed class VulWhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reason for adding whitelist.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Input("targetInfo")]
        public Input<string>? TargetInfo { get; set; }

        /// <summary>
        /// Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Input("whitelist", required: true)]
        public Input<string> Whitelist { get; set; } = null!;

        public VulWhitelistArgs()
        {
        }
        public static new VulWhitelistArgs Empty => new VulWhitelistArgs();
    }

    public sealed class VulWhitelistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reason for adding whitelist.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Set the effective range of the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Input("targetInfo")]
        public Input<string>? TargetInfo { get; set; }

        /// <summary>
        /// Information about the vulnerability to be added to the whitelist. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifycreatevulwhitelist).
        /// </summary>
        [Input("whitelist")]
        public Input<string>? Whitelist { get; set; }

        public VulWhitelistState()
        {
        }
        public static new VulWhitelistState Empty => new VulWhitelistState();
    }
}
