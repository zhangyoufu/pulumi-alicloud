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
    /// Provides a Ecs Image Component resource.
    /// 
    /// For information about Ecs Image Component and how to use it, see [What is Image Component](https://www.alibabacloud.com/help/en/doc-detail/200424.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.159.0.
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
    ///     var @default = AliCloud.ResourceManager.GetResourceGroups.Invoke(new()
    ///     {
    ///         NameRegex = "default",
    ///     });
    /// 
    ///     var example = new AliCloud.Ecs.EcsImageComponent("example", new()
    ///     {
    ///         ComponentType = "Build",
    ///         Content = "RUN yum update -y",
    ///         Description = "example_value",
    ///         ImageComponentName = "example_value",
    ///         ResourceGroupId = @default.Apply(@default =&gt; @default.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Groups[0]?.Id)),
    ///         SystemType = "Linux",
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
    /// Ecs Image Component can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ecs/ecsImageComponent:EcsImageComponent example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/ecsImageComponent:EcsImageComponent")]
    public partial class EcsImageComponent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The component type. Currently, only mirror build components are supported. Value: Build.  Default value: Build.
        /// </summary>
        [Output("componentType")]
        public Output<string> ComponentType { get; private set; } = null!;

        /// <summary>
        /// Component content.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Component creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describe the information.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The component name. The name must be 2 to 128 characters in length and must start with an uppercase letter or a Chinese character. It cannot start with http:// or https. Can contain Chinese, English, numbers, half-length colons (:), underscores (_), half-length periods (.), or dashes (-).  Note: If Name is not set, the return value of ImageComponentId is used by default.
        /// </summary>
        [Output("imageComponentName")]
        public Output<string> ImageComponentName { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The operating system supported by the component. Currently, only Linux systems are supported. Value: Linux.  Default value: Linux.
        /// </summary>
        [Output("systemType")]
        public Output<string> SystemType { get; private set; } = null!;

        /// <summary>
        /// List of label key-value pairs.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EcsImageComponent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EcsImageComponent(string name, EcsImageComponentArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, args ?? new EcsImageComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EcsImageComponent(string name, Input<string> id, EcsImageComponentState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/ecsImageComponent:EcsImageComponent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EcsImageComponent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EcsImageComponent Get(string name, Input<string> id, EcsImageComponentState? state = null, CustomResourceOptions? options = null)
        {
            return new EcsImageComponent(name, id, state, options);
        }
    }

    public sealed class EcsImageComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The component type. Currently, only mirror build components are supported. Value: Build.  Default value: Build.
        /// </summary>
        [Input("componentType")]
        public Input<string>? ComponentType { get; set; }

        /// <summary>
        /// Component content.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Describe the information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The component name. The name must be 2 to 128 characters in length and must start with an uppercase letter or a Chinese character. It cannot start with http:// or https. Can contain Chinese, English, numbers, half-length colons (:), underscores (_), half-length periods (.), or dashes (-).  Note: If Name is not set, the return value of ImageComponentId is used by default.
        /// </summary>
        [Input("imageComponentName")]
        public Input<string>? ImageComponentName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The operating system supported by the component. Currently, only Linux systems are supported. Value: Linux.  Default value: Linux.
        /// </summary>
        [Input("systemType")]
        public Input<string>? SystemType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// List of label key-value pairs.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsImageComponentArgs()
        {
        }
        public static new EcsImageComponentArgs Empty => new EcsImageComponentArgs();
    }

    public sealed class EcsImageComponentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The component type. Currently, only mirror build components are supported. Value: Build.  Default value: Build.
        /// </summary>
        [Input("componentType")]
        public Input<string>? ComponentType { get; set; }

        /// <summary>
        /// Component content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Component creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Describe the information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The component name. The name must be 2 to 128 characters in length and must start with an uppercase letter or a Chinese character. It cannot start with http:// or https. Can contain Chinese, English, numbers, half-length colons (:), underscores (_), half-length periods (.), or dashes (-).  Note: If Name is not set, the return value of ImageComponentId is used by default.
        /// </summary>
        [Input("imageComponentName")]
        public Input<string>? ImageComponentName { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The operating system supported by the component. Currently, only Linux systems are supported. Value: Linux.  Default value: Linux.
        /// </summary>
        [Input("systemType")]
        public Input<string>? SystemType { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// List of label key-value pairs.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EcsImageComponentState()
        {
        }
        public static new EcsImageComponentState Empty => new EcsImageComponentState();
    }
}
