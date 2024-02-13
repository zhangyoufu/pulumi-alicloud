// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oos
{
    /// <summary>
    /// Provides a OOS Application resource.
    /// 
    /// For information about OOS Application and how to use it, see [What is Application](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/api-oos-2019-06-01-createapplication).
    /// 
    /// &gt; **NOTE:** Available since v1.145.0.
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
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultApplication = new AliCloud.Oos.Application("defaultApplication", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id),
    ///         ApplicationName = name,
    ///         Description = name,
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OOS Application can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:oos/application:Application example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:oos/application:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the application.
        /// </summary>
        [Output("applicationName")]
        public Output<string> ApplicationName { get; private set; } = null!;

        /// <summary>
        /// Application group description information.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("alicloud:oos/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:oos/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the application.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// Application group description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }

    public sealed class ApplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the application.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// Application group description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ApplicationState()
        {
        }
        public static new ApplicationState Empty => new ApplicationState();
    }
}
