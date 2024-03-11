// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Edas
{
    /// <summary>
    /// Provides a EDAS Namespace resource.
    /// 
    /// For information about EDAS Namespace and how to use it, see [What is Namespace](https://www.alibabacloud.com/help/en/enterprise-distributed-application-service/latest/insertorupdateregion).
    /// 
    /// &gt; **NOTE:** Available since v1.173.0.
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
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-example";
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var defaultNamespace = new AliCloud.Edas.Namespace("defaultNamespace", new()
    ///     {
    ///         DebugEnable = false,
    ///         Description = name,
    ///         NamespaceLogicalId = $"{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:example",
    ///         NamespaceName = name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// EDAS Namespace can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:edas/namespace:Namespace example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:edas/namespace:Namespace")]
    public partial class Namespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to enable remote debugging.
        /// </summary>
        [Output("debugEnable")]
        public Output<bool> DebugEnable { get; private set; } = null!;

        /// <summary>
        /// The description of the namespace, The description can be up to `128` characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the namespace.
        /// - The ID of a custom namespace is in the `region ID:namespace identifier` format. An example is `cn-beijing:tdy218`.
        /// - The ID of the default namespace is in the `region ID` format. An example is cn-beijing.
        /// </summary>
        [Output("namespaceLogicalId")]
        public Output<string> NamespaceLogicalId { get; private set; } = null!;

        /// <summary>
        /// The name of the namespace, The name can be up to `63` characters in length.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:edas/namespace:Namespace", name, args ?? new NamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:edas/namespace:Namespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable remote debugging.
        /// </summary>
        [Input("debugEnable")]
        public Input<bool>? DebugEnable { get; set; }

        /// <summary>
        /// The description of the namespace, The description can be up to `128` characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the namespace.
        /// - The ID of a custom namespace is in the `region ID:namespace identifier` format. An example is `cn-beijing:tdy218`.
        /// - The ID of the default namespace is in the `region ID` format. An example is cn-beijing.
        /// </summary>
        [Input("namespaceLogicalId", required: true)]
        public Input<string> NamespaceLogicalId { get; set; } = null!;

        /// <summary>
        /// The name of the namespace, The name can be up to `63` characters in length.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        public NamespaceArgs()
        {
        }
        public static new NamespaceArgs Empty => new NamespaceArgs();
    }

    public sealed class NamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable remote debugging.
        /// </summary>
        [Input("debugEnable")]
        public Input<bool>? DebugEnable { get; set; }

        /// <summary>
        /// The description of the namespace, The description can be up to `128` characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the namespace.
        /// - The ID of a custom namespace is in the `region ID:namespace identifier` format. An example is `cn-beijing:tdy218`.
        /// - The ID of the default namespace is in the `region ID` format. An example is cn-beijing.
        /// </summary>
        [Input("namespaceLogicalId")]
        public Input<string>? NamespaceLogicalId { get; set; }

        /// <summary>
        /// The name of the namespace, The name can be up to `63` characters in length.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        public NamespaceState()
        {
        }
        public static new NamespaceState Empty => new NamespaceState();
    }
}
