// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms
{
    /// <summary>
    /// Provides a ARMS Grafana Workspace resource.
    /// 
    /// For information about ARMS Grafana Workspace and how to use it, see [What is Grafana Workspace](https://next.api.alibabacloud.com/document/ARMS/2019-08-08/ListGrafanaWorkspace).
    /// 
    /// &gt; **NOTE:** Available since v1.215.0.
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
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultGrafanaWorkspace = new AliCloud.Arms.GrafanaWorkspace("default", new()
    ///     {
    ///         GrafanaVersion = "9.0.x",
    ///         Description = name,
    ///         ResourceGroupId = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0])),
    ///         GrafanaWorkspaceEdition = "standard",
    ///         GrafanaWorkspaceName = name,
    ///         Tags = 
    ///         {
    ///             { "Created", "tf" },
    ///             { "For", "example" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ARMS Grafana Workspace can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:arms/grafanaWorkspace:GrafanaWorkspace example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:arms/grafanaWorkspace:GrafanaWorkspace")]
    public partial class GrafanaWorkspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The version of the grafana.
        /// </summary>
        [Output("grafanaVersion")]
        public Output<string?> GrafanaVersion { get; private set; } = null!;

        /// <summary>
        /// The edition of the grafana.
        /// </summary>
        [Output("grafanaWorkspaceEdition")]
        public Output<string?> GrafanaWorkspaceEdition { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("grafanaWorkspaceName")]
        public Output<string?> GrafanaWorkspaceName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaWorkspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaWorkspace(string name, GrafanaWorkspaceArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:arms/grafanaWorkspace:GrafanaWorkspace", name, args ?? new GrafanaWorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaWorkspace(string name, Input<string> id, GrafanaWorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:arms/grafanaWorkspace:GrafanaWorkspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaWorkspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaWorkspace Get(string name, Input<string> id, GrafanaWorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaWorkspace(name, id, state, options);
        }
    }

    public sealed class GrafanaWorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the grafana.
        /// </summary>
        [Input("grafanaVersion")]
        public Input<string>? GrafanaVersion { get; set; }

        /// <summary>
        /// The edition of the grafana.
        /// </summary>
        [Input("grafanaWorkspaceEdition")]
        public Input<string>? GrafanaWorkspaceEdition { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("grafanaWorkspaceName")]
        public Input<string>? GrafanaWorkspaceName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GrafanaWorkspaceArgs()
        {
        }
        public static new GrafanaWorkspaceArgs Empty => new GrafanaWorkspaceArgs();
    }

    public sealed class GrafanaWorkspaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the resource.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the grafana.
        /// </summary>
        [Input("grafanaVersion")]
        public Input<string>? GrafanaVersion { get; set; }

        /// <summary>
        /// The edition of the grafana.
        /// </summary>
        [Input("grafanaWorkspaceEdition")]
        public Input<string>? GrafanaWorkspaceEdition { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("grafanaWorkspaceName")]
        public Input<string>? GrafanaWorkspaceName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GrafanaWorkspaceState()
        {
        }
        public static new GrafanaWorkspaceState Empty => new GrafanaWorkspaceState();
    }
}
