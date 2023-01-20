// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a ECS Session Manager Status resource.
    /// 
    /// For information about ECS Session Manager Status and how to use it, see [What is Session Manager Status](https://www.alibabacloud.com/help/zh/doc-detail/337915.html).
    /// 
    /// &gt; **NOTE:** Available in v1.148.0+.
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
    ///     var @default = new AliCloud.Ecs.EcsSessionManagerStatus("default", new()
    ///     {
    ///         SessionManagerStatusName = "sessionManagerStatus",
    ///         Status = "Disabled",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS Session Manager Status can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/ecsSessionManagerStatus:EcsSessionManagerStatus example &lt;session_manager_status_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/ecsSessionManagerStatus:EcsSessionManagerStatus")]
    public partial class EcsSessionManagerStatus : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource. Valid values: `sessionManagerStatus`.
        /// </summary>
        [Output("sessionManagerStatusName")]
        public Output<string> SessionManagerStatusName { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `Disabled`, `Enabled`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a EcsSessionManagerStatus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EcsSessionManagerStatus(string name, EcsSessionManagerStatusArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsSessionManagerStatus:EcsSessionManagerStatus", name, args ?? new EcsSessionManagerStatusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EcsSessionManagerStatus(string name, Input<string> id, EcsSessionManagerStatusState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsSessionManagerStatus:EcsSessionManagerStatus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EcsSessionManagerStatus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EcsSessionManagerStatus Get(string name, Input<string> id, EcsSessionManagerStatusState? state = null, CustomResourceOptions? options = null)
        {
            return new EcsSessionManagerStatus(name, id, state, options);
        }
    }

    public sealed class EcsSessionManagerStatusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource. Valid values: `sessionManagerStatus`.
        /// </summary>
        [Input("sessionManagerStatusName", required: true)]
        public Input<string> SessionManagerStatusName { get; set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `Disabled`, `Enabled`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public EcsSessionManagerStatusArgs()
        {
        }
        public static new EcsSessionManagerStatusArgs Empty => new EcsSessionManagerStatusArgs();
    }

    public sealed class EcsSessionManagerStatusState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the resource. Valid values: `sessionManagerStatus`.
        /// </summary>
        [Input("sessionManagerStatusName")]
        public Input<string>? SessionManagerStatusName { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Disabled`, `Enabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public EcsSessionManagerStatusState()
        {
        }
        public static new EcsSessionManagerStatusState Empty => new EcsSessionManagerStatusState();
    }
}
