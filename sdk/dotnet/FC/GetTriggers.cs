// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC
{
    public static class GetTriggers
    {
        /// <summary>
        /// This data source provides the Function Compute triggers of the current Alibaba Cloud user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fcTriggersDs = AliCloud.FC.GetTriggers.Invoke(new()
        ///     {
        ///         FunctionName = "sample_function",
        ///         NameRegex = "sample_fc_trigger",
        ///         ServiceName = "sample_service",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstFcTriggerName"] = fcTriggersDs.Apply(getTriggersResult =&gt; getTriggersResult.Triggers[0]?.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTriggersResult> InvokeAsync(GetTriggersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTriggersResult>("alicloud:fc/getTriggers:getTriggers", args ?? new GetTriggersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Function Compute triggers of the current Alibaba Cloud user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fcTriggersDs = AliCloud.FC.GetTriggers.Invoke(new()
        ///     {
        ///         FunctionName = "sample_function",
        ///         NameRegex = "sample_fc_trigger",
        ///         ServiceName = "sample_service",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstFcTriggerName"] = fcTriggersDs.Apply(getTriggersResult =&gt; getTriggersResult.Triggers[0]?.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTriggersResult> Invoke(GetTriggersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTriggersResult>("alicloud:fc/getTriggers:getTriggers", args ?? new GetTriggersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTriggersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// FC function name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of FC triggers ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by FC trigger name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// FC service name.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetTriggersArgs()
        {
        }
        public static new GetTriggersArgs Empty => new GetTriggersArgs();
    }

    public sealed class GetTriggersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// FC function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of FC triggers ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by FC trigger name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// FC service name.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetTriggersInvokeArgs()
        {
        }
        public static new GetTriggersInvokeArgs Empty => new GetTriggersInvokeArgs();
    }


    [OutputType]
    public sealed class GetTriggersResult
    {
        public readonly string FunctionName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of FC triggers ids.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of FC triggers names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string ServiceName;
        /// <summary>
        /// A list of FC triggers. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTriggersTriggerResult> Triggers;

        [OutputConstructor]
        private GetTriggersResult(
            string functionName,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string serviceName,

            ImmutableArray<Outputs.GetTriggersTriggerResult> triggers)
        {
            FunctionName = functionName;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            ServiceName = serviceName;
            Triggers = triggers;
        }
    }
}
