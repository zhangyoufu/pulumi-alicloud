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
    /// Provides a Cloud Enterprise Network (CEN) Transit Router Cidr resource.
    /// 
    /// For information about Cloud Enterprise Network (CEN) Transit Router Cidr and how to use it, see [What is Transit Router Cidr](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/createtransitroutercidr).
    /// 
    /// &gt; **NOTE:** Available since v1.193.0.
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
    ///     var example = new AliCloud.Cen.Instance("example", new()
    ///     {
    ///         CenInstanceName = "tf_example",
    ///         Description = "an example for cen",
    ///     });
    /// 
    ///     var exampleTransitRouter = new AliCloud.Cen.TransitRouter("example", new()
    ///     {
    ///         TransitRouterName = "tf_example",
    ///         CenId = example.Id,
    ///     });
    /// 
    ///     var exampleTransitRouterCidr = new AliCloud.Cen.TransitRouterCidr("example", new()
    ///     {
    ///         TransitRouterId = exampleTransitRouter.TransitRouterId,
    ///         Cidr = "192.168.0.0/16",
    ///         TransitRouterCidrName = "tf_example",
    ///         Description = "tf_example",
    ///         PublishCidrRoute = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Cloud Enterprise Network (CEN) Transit Router Cidr can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/transitRouterCidr:TransitRouterCidr default &lt;transit_router_id&gt;:&lt;transit_router_cidr_id&gt;.
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/transitRouterCidr:TransitRouterCidr")]
    public partial class TransitRouterCidr : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The cidr of the transit router.
        /// </summary>
        [Output("cidr")]
        public Output<string> Cidr { get; private set; } = null!;

        /// <summary>
        /// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        /// </summary>
        [Output("publishCidrRoute")]
        public Output<bool> PublishCidrRoute { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit router cidr.
        /// </summary>
        [Output("transitRouterCidrId")]
        public Output<string> TransitRouterCidrId { get; private set; } = null!;

        /// <summary>
        /// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Output("transitRouterCidrName")]
        public Output<string?> TransitRouterCidrName { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Output("transitRouterId")]
        public Output<string> TransitRouterId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitRouterCidr resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitRouterCidr(string name, TransitRouterCidrArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterCidr:TransitRouterCidr", name, args ?? new TransitRouterCidrArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitRouterCidr(string name, Input<string> id, TransitRouterCidrState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterCidr:TransitRouterCidr", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitRouterCidr resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitRouterCidr Get(string name, Input<string> id, TransitRouterCidrState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitRouterCidr(name, id, state, options);
        }
    }

    public sealed class TransitRouterCidrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cidr of the transit router.
        /// </summary>
        [Input("cidr", required: true)]
        public Input<string> Cidr { get; set; } = null!;

        /// <summary>
        /// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        /// </summary>
        [Input("publishCidrRoute")]
        public Input<bool>? PublishCidrRoute { get; set; }

        /// <summary>
        /// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("transitRouterCidrName")]
        public Input<string>? TransitRouterCidrName { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        public TransitRouterCidrArgs()
        {
        }
        public static new TransitRouterCidrArgs Empty => new TransitRouterCidrArgs();
    }

    public sealed class TransitRouterCidrState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cidr of the transit router.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// The description of the transit router. The description must be `2` to `256` characters in length, and it must start with English letters, but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to allow automatically adding Transit Router Cidr in Transit Router Route Table. Valid values: `true` and `false`. Default value: `true`.
        /// </summary>
        [Input("publishCidrRoute")]
        public Input<bool>? PublishCidrRoute { get; set; }

        /// <summary>
        /// The ID of the transit router cidr.
        /// </summary>
        [Input("transitRouterCidrId")]
        public Input<string>? TransitRouterCidrId { get; set; }

        /// <summary>
        /// The name of the transit router. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
        /// </summary>
        [Input("transitRouterCidrName")]
        public Input<string>? TransitRouterCidrName { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public TransitRouterCidrState()
        {
        }
        public static new TransitRouterCidrState Empty => new TransitRouterCidrState();
    }
}
