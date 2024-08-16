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
    /// Provides a CEN Transit Router Ecr Attachment resource.
    /// 
    /// For information about CEN Transit Router Ecr Attachment and how to use it, see [What is Transit Router Ecr Attachment](https://www.alibabacloud.com/help/en/).
    /// 
    /// &gt; **NOTE:** Available since v1.226.0.
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
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var asn = config.Get("asn") ?? "4200000666";
    ///     var defaultO8Hcfx = new AliCloud.ExpressConnect.RouterExpressConnectRouter("defaultO8Hcfx", new()
    ///     {
    ///         AlibabaSideAsn = asn,
    ///         EcrName = name,
    ///     });
    /// 
    ///     var defaultQKBiay = new AliCloud.Cen.Instance("defaultQKBiay", new()
    ///     {
    ///         CenInstanceName = name,
    ///     });
    /// 
    ///     var defaultQa94Y1 = new AliCloud.Cen.TransitRouter("defaultQa94Y1", new()
    ///     {
    ///         CenId = defaultQKBiay.Id,
    ///         TransitRouterName = name,
    ///     });
    /// 
    ///     var current = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultedPu6c = new AliCloud.ExpressConnect.RouterTrAssociation("defaultedPu6c", new()
    ///     {
    ///         AssociationRegionId = "cn-hangzhou",
    ///         EcrId = defaultO8Hcfx.Id,
    ///         CenId = defaultQKBiay.Id,
    ///         TransitRouterId = defaultQa94Y1.TransitRouterId,
    ///         TransitRouterOwnerId = current.Apply(getAccountResult =&gt; getAccountResult.Id),
    ///     });
    /// 
    ///     var @default = new AliCloud.Cen.TransitRouterEcrAttachment("default", new()
    ///     {
    ///         EcrId = defaultO8Hcfx.Id,
    ///         CenId = defaultedPu6c.CenId,
    ///         TransitRouterEcrAttachmentName = name,
    ///         TransitRouterAttachmentDescription = name,
    ///         TransitRouterId = defaultQa94Y1.TransitRouterId,
    ///         EcrOwnerId = current.Apply(getAccountResult =&gt; getAccountResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN Transit Router Ecr Attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cen/transitRouterEcrAttachment:TransitRouterEcrAttachment example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/transitRouterEcrAttachment:TransitRouterEcrAttachment")]
    public partial class TransitRouterEcrAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CenId
        /// </summary>
        [Output("cenId")]
        public Output<string?> CenId { get; private set; } = null!;

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// EcrId
        /// </summary>
        [Output("ecrId")]
        public Output<string> EcrId { get; private set; } = null!;

        /// <summary>
        /// EcrOwnerId
        /// </summary>
        [Output("ecrOwnerId")]
        public Output<int?> EcrOwnerId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// TransitRouterAttachmentDescription
        /// </summary>
        [Output("transitRouterAttachmentDescription")]
        public Output<string?> TransitRouterAttachmentDescription { get; private set; } = null!;

        /// <summary>
        /// TransitRouterAttachmentName
        /// </summary>
        [Output("transitRouterEcrAttachmentName")]
        public Output<string?> TransitRouterEcrAttachmentName { get; private set; } = null!;

        /// <summary>
        /// TransitRouterId
        /// </summary>
        [Output("transitRouterId")]
        public Output<string?> TransitRouterId { get; private set; } = null!;


        /// <summary>
        /// Create a TransitRouterEcrAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransitRouterEcrAttachment(string name, TransitRouterEcrAttachmentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterEcrAttachment:TransitRouterEcrAttachment", name, args ?? new TransitRouterEcrAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransitRouterEcrAttachment(string name, Input<string> id, TransitRouterEcrAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/transitRouterEcrAttachment:TransitRouterEcrAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransitRouterEcrAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransitRouterEcrAttachment Get(string name, Input<string> id, TransitRouterEcrAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new TransitRouterEcrAttachment(name, id, state, options);
        }
    }

    public sealed class TransitRouterEcrAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CenId
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// EcrId
        /// </summary>
        [Input("ecrId", required: true)]
        public Input<string> EcrId { get; set; } = null!;

        /// <summary>
        /// EcrOwnerId
        /// </summary>
        [Input("ecrOwnerId")]
        public Input<int>? EcrOwnerId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// TransitRouterAttachmentDescription
        /// </summary>
        [Input("transitRouterAttachmentDescription")]
        public Input<string>? TransitRouterAttachmentDescription { get; set; }

        /// <summary>
        /// TransitRouterAttachmentName
        /// </summary>
        [Input("transitRouterEcrAttachmentName")]
        public Input<string>? TransitRouterEcrAttachmentName { get; set; }

        /// <summary>
        /// TransitRouterId
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public TransitRouterEcrAttachmentArgs()
        {
        }
        public static new TransitRouterEcrAttachmentArgs Empty => new TransitRouterEcrAttachmentArgs();
    }

    public sealed class TransitRouterEcrAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CenId
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The creation time of the resource
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// EcrId
        /// </summary>
        [Input("ecrId")]
        public Input<string>? EcrId { get; set; }

        /// <summary>
        /// EcrOwnerId
        /// </summary>
        [Input("ecrOwnerId")]
        public Input<int>? EcrOwnerId { get; set; }

        /// <summary>
        /// The status of the resource
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// TransitRouterAttachmentDescription
        /// </summary>
        [Input("transitRouterAttachmentDescription")]
        public Input<string>? TransitRouterAttachmentDescription { get; set; }

        /// <summary>
        /// TransitRouterAttachmentName
        /// </summary>
        [Input("transitRouterEcrAttachmentName")]
        public Input<string>? TransitRouterEcrAttachmentName { get; set; }

        /// <summary>
        /// TransitRouterId
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public TransitRouterEcrAttachmentState()
        {
        }
        public static new TransitRouterEcrAttachmentState Empty => new TransitRouterEcrAttachmentState();
    }
}
