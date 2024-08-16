// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    /// <summary>
    /// Provides a Api Gateway Plugin resource.
    /// 
    /// For information about Api Gateway Plugin and how to use it, see [What is Plugin](https://www.alibabacloud.com/help/en/api-gateway/developer-reference/api-cloudapi-2016-07-14-createplugin).
    /// 
    /// &gt; **NOTE:** Available since v1.187.0.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform_example";
    ///     var @default = new AliCloud.ApiGateway.Plugin("default", new()
    ///     {
    ///         Description = name,
    ///         PluginName = name,
    ///         PluginData = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["routes"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "Vip",
    ///                     ["condition"] = "$CaAppId = 123456",
    ///                     ["backend"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["type"] = "HTTP-VPC",
    ///                         ["vpcAccessName"] = "slbAccessForVip",
    ///                     },
    ///                 },
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "MockForOldClient",
    ///                     ["condition"] = "$ClientVersion &lt; '2.0.5'",
    ///                     ["backend"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["type"] = "MOCK",
    ///                         ["statusCode"] = 400,
    ///                         ["mockBody"] = "This version is not supported!!!",
    ///                     },
    ///                 },
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "BlueGreenPercent05",
    ///                     ["condition"] = "1 = 1",
    ///                     ["backend"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["type"] = "HTTP",
    ///                         ["address"] = "https://beta-version.api.foo.com",
    ///                     },
    ///                     ["constant-parameters"] = new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["name"] = "x-route-blue-green",
    ///                             ["location"] = "header",
    ///                             ["value"] = "route-blue-green",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///         PluginType = "routing",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api Gateway Plugin can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:apigateway/plugin:Plugin example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/plugin:Plugin")]
    public partial class Plugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the plug-in, which cannot exceed 200 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The definition statement of the plug-in. Plug-in definition statements in the JSON and YAML formats are supported.
        /// </summary>
        [Output("pluginData")]
        public Output<string> PluginData { get; private set; } = null!;

        /// <summary>
        /// The name of the plug-in that you want to create. It can contain uppercase English letters, lowercase English letters, Chinese characters, numbers, and underscores (_). It must be 4 to 50 characters in length and cannot start with an underscore (_).
        /// </summary>
        [Output("pluginName")]
        public Output<string> PluginName { get; private set; } = null!;

        /// <summary>
        /// The type of the plug-in. Valid values:
        /// - "trafficControl"
        /// - "ipControl"
        /// - "backendSignature"
        /// - "jwtAuth"
        /// - "basicAuth"
        /// - "cors"
        /// - "caching"
        /// - "routing"
        /// - "accessControl"
        /// - "errorMapping"
        /// - "circuitBreaker"
        /// - "remoteAuth"
        /// - "logMask"
        /// - "transformer".
        /// </summary>
        [Output("pluginType")]
        public Output<string> PluginType { get; private set; } = null!;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Plugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Plugin(string name, PluginArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/plugin:Plugin", name, args ?? new PluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Plugin(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/plugin:Plugin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Plugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Plugin Get(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
        {
            return new Plugin(name, id, state, options);
        }
    }

    public sealed class PluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the plug-in, which cannot exceed 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The definition statement of the plug-in. Plug-in definition statements in the JSON and YAML formats are supported.
        /// </summary>
        [Input("pluginData", required: true)]
        public Input<string> PluginData { get; set; } = null!;

        /// <summary>
        /// The name of the plug-in that you want to create. It can contain uppercase English letters, lowercase English letters, Chinese characters, numbers, and underscores (_). It must be 4 to 50 characters in length and cannot start with an underscore (_).
        /// </summary>
        [Input("pluginName", required: true)]
        public Input<string> PluginName { get; set; } = null!;

        /// <summary>
        /// The type of the plug-in. Valid values:
        /// - "trafficControl"
        /// - "ipControl"
        /// - "backendSignature"
        /// - "jwtAuth"
        /// - "basicAuth"
        /// - "cors"
        /// - "caching"
        /// - "routing"
        /// - "accessControl"
        /// - "errorMapping"
        /// - "circuitBreaker"
        /// - "remoteAuth"
        /// - "logMask"
        /// - "transformer".
        /// </summary>
        [Input("pluginType", required: true)]
        public Input<string> PluginType { get; set; } = null!;

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

        public PluginArgs()
        {
        }
        public static new PluginArgs Empty => new PluginArgs();
    }

    public sealed class PluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the plug-in, which cannot exceed 200 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The definition statement of the plug-in. Plug-in definition statements in the JSON and YAML formats are supported.
        /// </summary>
        [Input("pluginData")]
        public Input<string>? PluginData { get; set; }

        /// <summary>
        /// The name of the plug-in that you want to create. It can contain uppercase English letters, lowercase English letters, Chinese characters, numbers, and underscores (_). It must be 4 to 50 characters in length and cannot start with an underscore (_).
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// The type of the plug-in. Valid values:
        /// - "trafficControl"
        /// - "ipControl"
        /// - "backendSignature"
        /// - "jwtAuth"
        /// - "basicAuth"
        /// - "cors"
        /// - "caching"
        /// - "routing"
        /// - "accessControl"
        /// - "errorMapping"
        /// - "circuitBreaker"
        /// - "remoteAuth"
        /// - "logMask"
        /// - "transformer".
        /// </summary>
        [Input("pluginType")]
        public Input<string>? PluginType { get; set; }

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

        public PluginState()
        {
        }
        public static new PluginState Empty => new PluginState();
    }
}
