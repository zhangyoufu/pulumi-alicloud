// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DirectMail
{
    /// <summary>
    /// Provides a Direct Mail Receivers resource.
    /// 
    /// For information about Direct Mail Receivers and how to use it, see [What is Direct Mail Receivers](https://www.alibabacloud.com/help/en/doc-detail/29414.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.125.0+.
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
    ///     var example = new AliCloud.DirectMail.Receivers("example", new()
    ///     {
    ///         ReceiversAlias = "tf-vme8@onaliyun.com",
    ///         ReceiversName = "vme8",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Direct Mail Receivers can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:directmail/receivers:Receivers example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:directmail/receivers:Receivers")]
    public partial class Receivers : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of receivers and 1-50 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The alias of receivers. Must email address and less than 30 characters in length.
        /// </summary>
        [Output("receiversAlias")]
        public Output<string> ReceiversAlias { get; private set; } = null!;

        /// <summary>
        /// The name of the resource. The length that cannot be repeated is 1-30 characters.
        /// </summary>
        [Output("receiversName")]
        public Output<string> ReceiversName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. `0` means uploading, `1` means upload completed.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Receivers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Receivers(string name, ReceiversArgs args, CustomResourceOptions? options = null)
            : base("alicloud:directmail/receivers:Receivers", name, args ?? new ReceiversArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Receivers(string name, Input<string> id, ReceiversState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:directmail/receivers:Receivers", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Receivers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Receivers Get(string name, Input<string> id, ReceiversState? state = null, CustomResourceOptions? options = null)
        {
            return new Receivers(name, id, state, options);
        }
    }

    public sealed class ReceiversArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of receivers and 1-50 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The alias of receivers. Must email address and less than 30 characters in length.
        /// </summary>
        [Input("receiversAlias", required: true)]
        public Input<string> ReceiversAlias { get; set; } = null!;

        /// <summary>
        /// The name of the resource. The length that cannot be repeated is 1-30 characters.
        /// </summary>
        [Input("receiversName", required: true)]
        public Input<string> ReceiversName { get; set; } = null!;

        public ReceiversArgs()
        {
        }
        public static new ReceiversArgs Empty => new ReceiversArgs();
    }

    public sealed class ReceiversState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of receivers and 1-50 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The alias of receivers. Must email address and less than 30 characters in length.
        /// </summary>
        [Input("receiversAlias")]
        public Input<string>? ReceiversAlias { get; set; }

        /// <summary>
        /// The name of the resource. The length that cannot be repeated is 1-30 characters.
        /// </summary>
        [Input("receiversName")]
        public Input<string>? ReceiversName { get; set; }

        /// <summary>
        /// The status of the resource. `0` means uploading, `1` means upload completed.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        public ReceiversState()
        {
        }
        public static new ReceiversState Empty => new ReceiversState();
    }
}
