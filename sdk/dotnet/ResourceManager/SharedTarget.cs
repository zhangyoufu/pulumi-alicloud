// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    /// <summary>
    /// Provides a Resource Manager Shared Target resource.
    /// 
    /// For information about Resource Manager Shared Target and how to use it, see [What is Shared Target](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.111.0+.
    /// 
    /// ## Import
    /// 
    /// Resource Manager Shared Target can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:resourcemanager/sharedTarget:SharedTarget example &lt;resource_share_id&gt;:&lt;target_id&gt;
    /// ```
    /// </summary>
    public partial class SharedTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Output("resourceShareId")]
        public Output<string> ResourceShareId { get; private set; } = null!;

        /// <summary>
        /// The status of shared target.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The member account ID in resource directory.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;


        /// <summary>
        /// Create a SharedTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SharedTarget(string name, SharedTargetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/sharedTarget:SharedTarget", name, args ?? new SharedTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SharedTarget(string name, Input<string> id, SharedTargetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/sharedTarget:SharedTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SharedTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SharedTarget Get(string name, Input<string> id, SharedTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new SharedTarget(name, id, state, options);
        }
    }

    public sealed class SharedTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Input("resourceShareId", required: true)]
        public Input<string> ResourceShareId { get; set; } = null!;

        /// <summary>
        /// The member account ID in resource directory.
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        public SharedTargetArgs()
        {
        }
    }

    public sealed class SharedTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource share ID of resource manager.
        /// </summary>
        [Input("resourceShareId")]
        public Input<string>? ResourceShareId { get; set; }

        /// <summary>
        /// The status of shared target.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The member account ID in resource directory.
        /// </summary>
        [Input("targetId")]
        public Input<string>? TargetId { get; set; }

        public SharedTargetState()
        {
        }
    }
}
