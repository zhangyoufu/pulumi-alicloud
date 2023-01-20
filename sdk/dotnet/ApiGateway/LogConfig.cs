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
    /// Provides a Api Gateway Log Config resource.
    /// 
    /// For information about Api Gateway Log Config and how to use it, see [What is Log Config](https://help.aliyun.com/document_detail/400392.html).
    /// 
    /// &gt; **NOTE:** Available in v1.185.0+.
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
    ///     var @default = new AliCloud.ApiGateway.LogConfig("default", new()
    ///     {
    ///         LogType = "PROVIDER",
    ///         SlsLogStore = "example_value",
    ///         SlsProject = "example_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api Gateway Log Config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:apigateway/logConfig:LogConfig example &lt;log_type&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/logConfig:LogConfig")]
    public partial class LogConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type the of log. Valid values: `PROVIDER`.
        /// </summary>
        [Output("logType")]
        public Output<string> LogType { get; private set; } = null!;

        /// <summary>
        /// The name of the Log Store.
        /// </summary>
        [Output("slsLogStore")]
        public Output<string> SlsLogStore { get; private set; } = null!;

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Output("slsProject")]
        public Output<string> SlsProject { get; private set; } = null!;


        /// <summary>
        /// Create a LogConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogConfig(string name, LogConfigArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/logConfig:LogConfig", name, args ?? new LogConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogConfig(string name, Input<string> id, LogConfigState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/logConfig:LogConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogConfig Get(string name, Input<string> id, LogConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new LogConfig(name, id, state, options);
        }
    }

    public sealed class LogConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type the of log. Valid values: `PROVIDER`.
        /// </summary>
        [Input("logType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        /// <summary>
        /// The name of the Log Store.
        /// </summary>
        [Input("slsLogStore", required: true)]
        public Input<string> SlsLogStore { get; set; } = null!;

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Input("slsProject", required: true)]
        public Input<string> SlsProject { get; set; } = null!;

        public LogConfigArgs()
        {
        }
        public static new LogConfigArgs Empty => new LogConfigArgs();
    }

    public sealed class LogConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type the of log. Valid values: `PROVIDER`.
        /// </summary>
        [Input("logType")]
        public Input<string>? LogType { get; set; }

        /// <summary>
        /// The name of the Log Store.
        /// </summary>
        [Input("slsLogStore")]
        public Input<string>? SlsLogStore { get; set; }

        /// <summary>
        /// The name of the Project.
        /// </summary>
        [Input("slsProject")]
        public Input<string>? SlsProject { get; set; }

        public LogConfigState()
        {
        }
        public static new LogConfigState Empty => new LogConfigState();
    }
}
