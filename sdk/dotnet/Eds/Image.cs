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
    /// Provides a ECD Image resource.
    /// 
    /// For information about ECD Image and how to use it, see [What is Image](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createimage).
    /// 
    /// &gt; **NOTE:** Available since v1.146.0.
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
    ///     var defaultImage = new AliCloud.Eds.Image("defaultImage", new()
    ///     {
    ///         ImageName = name,
    ///         DesktopId = defaultDesktop.Id,
    ///         Description = name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Image can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:eds/image:Image example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/image:Image")]
    public partial class Image : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the image.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The desktop id of the desktop.
        /// </summary>
        [Output("desktopId")]
        public Output<string> DesktopId { get; private set; } = null!;

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Output("imageName")]
        public Output<string?> ImageName { get; private set; } = null!;

        /// <summary>
        /// The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/image:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/image:Image", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
        {
            return new Image(name, id, state, options);
        }
    }

    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The desktop id of the desktop.
        /// </summary>
        [Input("desktopId", required: true)]
        public Input<string> DesktopId { get; set; } = null!;

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        public ImageArgs()
        {
        }
        public static new ImageArgs Empty => new ImageArgs();
    }

    public sealed class ImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the image.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The desktop id of the desktop.
        /// </summary>
        [Input("desktopId")]
        public Input<string>? DesktopId { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ImageState()
        {
        }
        public static new ImageState Empty => new ImageState();
    }
}
