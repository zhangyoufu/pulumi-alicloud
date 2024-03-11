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
    /// Provides a ECS Hpc Cluster resource.
    /// 
    /// For information about ECS Hpc Cluster and how to use it, see [What is Hpc Cluster](https://www.alibabacloud.com/help/en/doc-detail/109138.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.116.0+.
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
    ///     var example = new AliCloud.Ecs.HpcCluster("example", new()
    ///     {
    ///         Description = "For Terraform Test",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ECS Hpc Cluster can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/hpcCluster:HpcCluster example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/hpcCluster:HpcCluster")]
    public partial class HpcCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of ECS Hpc Cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of ECS Hpc Cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a HpcCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HpcCluster(string name, HpcClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/hpcCluster:HpcCluster", name, args ?? new HpcClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HpcCluster(string name, Input<string> id, HpcClusterState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/hpcCluster:HpcCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HpcCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HpcCluster Get(string name, Input<string> id, HpcClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new HpcCluster(name, id, state, options);
        }
    }

    public sealed class HpcClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of ECS Hpc Cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of ECS Hpc Cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public HpcClusterArgs()
        {
        }
        public static new HpcClusterArgs Empty => new HpcClusterArgs();
    }

    public sealed class HpcClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of ECS Hpc Cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of ECS Hpc Cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public HpcClusterState()
        {
        }
        public static new HpcClusterState Empty => new HpcClusterState();
    }
}
