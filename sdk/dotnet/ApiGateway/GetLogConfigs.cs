// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    public static class GetLogConfigs
    {
        /// <summary>
        /// This data source provides the Api Gateway Log Configs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
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
        ///     var ids = AliCloud.ApiGateway.GetLogConfigs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var logType = AliCloud.ApiGateway.GetLogConfigs.Invoke(new()
        ///     {
        ///         LogType = "PROVIDER",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["apiGatewayLogConfigId1"] = ids.Apply(getLogConfigsResult =&gt; getLogConfigsResult.Configs[0]?.Id),
        ///         ["apiGatewayLogConfigId2"] = logType.Apply(getLogConfigsResult =&gt; getLogConfigsResult.Configs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLogConfigsResult> InvokeAsync(GetLogConfigsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogConfigsResult>("alicloud:apigateway/getLogConfigs:getLogConfigs", args ?? new GetLogConfigsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Api Gateway Log Configs of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.185.0+.
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
        ///     var ids = AliCloud.ApiGateway.GetLogConfigs.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///     });
        /// 
        ///     var logType = AliCloud.ApiGateway.GetLogConfigs.Invoke(new()
        ///     {
        ///         LogType = "PROVIDER",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["apiGatewayLogConfigId1"] = ids.Apply(getLogConfigsResult =&gt; getLogConfigsResult.Configs[0]?.Id),
        ///         ["apiGatewayLogConfigId2"] = logType.Apply(getLogConfigsResult =&gt; getLogConfigsResult.Configs[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLogConfigsResult> Invoke(GetLogConfigsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogConfigsResult>("alicloud:apigateway/getLogConfigs:getLogConfigs", args ?? new GetLogConfigsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogConfigsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Log Config IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The type the of log.
        /// </summary>
        [Input("logType")]
        public string? LogType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetLogConfigsArgs()
        {
        }
        public static new GetLogConfigsArgs Empty => new GetLogConfigsArgs();
    }

    public sealed class GetLogConfigsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Log Config IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The type the of log.
        /// </summary>
        [Input("logType")]
        public Input<string>? LogType { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetLogConfigsInvokeArgs()
        {
        }
        public static new GetLogConfigsInvokeArgs Empty => new GetLogConfigsInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogConfigsResult
    {
        public readonly ImmutableArray<Outputs.GetLogConfigsConfigResult> Configs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? LogType;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetLogConfigsResult(
            ImmutableArray<Outputs.GetLogConfigsConfigResult> configs,

            string id,

            ImmutableArray<string> ids,

            string? logType,

            string? outputFile)
        {
            Configs = configs;
            Id = id;
            Ids = ids;
            LogType = logType;
            OutputFile = outputFile;
        }
    }
}
