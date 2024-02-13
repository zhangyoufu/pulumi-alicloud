// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh
{
    /// <summary>
    /// Provides a Service Mesh Extension Provider resource.
    /// 
    /// For information about Service Mesh Extension Provider and how to use it, see [What is Extension Provider](https://help.aliyun.com/document_detail/461549.html).
    /// 
    /// &gt; **NOTE:** Available in v1.191.0+.
    /// 
    /// ## Import
    /// 
    /// Service Mesh Extension Provider can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:servicemesh/extensionProvider:ExtensionProvider example &lt;service_mesh_id&gt;:&lt;type&gt;:&lt;extension_provider_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:servicemesh/extensionProvider:ExtensionProvider")]
    public partial class ExtensionProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The config of the Service Mesh Extension Provider. The `config` format is json.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
        /// </summary>
        [Output("extensionProviderName")]
        public Output<string> ExtensionProviderName { get; private set; } = null!;

        /// <summary>
        /// The ID of the Service Mesh.
        /// </summary>
        [Output("serviceMeshId")]
        public Output<string> ServiceMeshId { get; private set; } = null!;

        /// <summary>
        /// The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ExtensionProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExtensionProvider(string name, ExtensionProviderArgs args, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/extensionProvider:ExtensionProvider", name, args ?? new ExtensionProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExtensionProvider(string name, Input<string> id, ExtensionProviderState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:servicemesh/extensionProvider:ExtensionProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExtensionProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExtensionProvider Get(string name, Input<string> id, ExtensionProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new ExtensionProvider(name, id, state, options);
        }
    }

    public sealed class ExtensionProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of the Service Mesh Extension Provider. The `config` format is json.
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
        /// </summary>
        [Input("extensionProviderName", required: true)]
        public Input<string> ExtensionProviderName { get; set; } = null!;

        /// <summary>
        /// The ID of the Service Mesh.
        /// </summary>
        [Input("serviceMeshId", required: true)]
        public Input<string> ServiceMeshId { get; set; } = null!;

        /// <summary>
        /// The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ExtensionProviderArgs()
        {
        }
        public static new ExtensionProviderArgs Empty => new ExtensionProviderArgs();
    }

    public sealed class ExtensionProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config of the Service Mesh Extension Provider. The `config` format is json.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The name of the Service Mesh Extension Provider. It must be prefixed with `$type-`, for example `httpextauth-xxx`, `grpcextauth-xxx`.
        /// </summary>
        [Input("extensionProviderName")]
        public Input<string>? ExtensionProviderName { get; set; }

        /// <summary>
        /// The ID of the Service Mesh.
        /// </summary>
        [Input("serviceMeshId")]
        public Input<string>? ServiceMeshId { get; set; }

        /// <summary>
        /// The type of the Service Mesh Extension Provider. Valid values: `httpextauth`, `grpcextauth`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ExtensionProviderState()
        {
        }
        public static new ExtensionProviderState Empty => new ExtensionProviderState();
    }
}
