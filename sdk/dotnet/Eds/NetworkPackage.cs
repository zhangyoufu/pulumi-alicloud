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
    /// Provides a ECD Network Package resource.
    /// 
    /// For information about ECD Network Package and how to use it, see [What is Network Package](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createnetworkpackage).
    /// 
    /// &gt; **NOTE:** Available since v1.142.0.
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
    ///         EnableAdminAccess = false,
    ///         DesktopAccessType = "Internet",
    ///         OfficeSiteName = defaultRandomInteger.Result.Apply(result =&gt; $"{name}-{result}"),
    ///     });
    /// 
    ///     var defaultNetworkPackage = new AliCloud.Eds.NetworkPackage("defaultNetworkPackage", new()
    ///     {
    ///         Bandwidth = 10,
    ///         OfficeSiteId = defaultSimpleOfficeSite.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ECD Network Package can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:eds/networkPackage:NetworkPackage example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/networkPackage:NetworkPackage")]
    public partial class NetworkPackage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// The internet charge type  of  package.
        /// </summary>
        [Output("internetChargeType")]
        public Output<string> InternetChargeType { get; private set; } = null!;

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Output("officeSiteId")]
        public Output<string> OfficeSiteId { get; private set; } = null!;

        /// <summary>
        /// The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkPackage(string name, NetworkPackageArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/networkPackage:NetworkPackage", name, args ?? new NetworkPackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkPackage(string name, Input<string> id, NetworkPackageState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/networkPackage:NetworkPackage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkPackage Get(string name, Input<string> id, NetworkPackageState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkPackage(name, id, state, options);
        }
    }

    public sealed class NetworkPackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<int> Bandwidth { get; set; } = null!;

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Input("officeSiteId", required: true)]
        public Input<string> OfficeSiteId { get; set; } = null!;

        public NetworkPackageArgs()
        {
        }
        public static new NetworkPackageArgs Empty => new NetworkPackageArgs();
    }

    public sealed class NetworkPackageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The internet charge type  of  package.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The ID of office site.
        /// </summary>
        [Input("officeSiteId")]
        public Input<string>? OfficeSiteId { get; set; }

        /// <summary>
        /// The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public NetworkPackageState()
        {
        }
        public static new NetworkPackageState Empty => new NetworkPackageState();
    }
}
