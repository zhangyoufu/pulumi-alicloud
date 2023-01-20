// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FNF
{
    public static class GetExecutions
    {
        /// <summary>
        /// This data source provides the FnF Executions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.149.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.FNF.GetExecutions.Invoke(new()
        ///     {
        ///         FlowName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-Execution-1",
        ///             "my-Execution-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["fnfExecutionId1"] = data.Alicloud_fn_f_executions.Ids.Executions[0].Id,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetExecutionsResult> InvokeAsync(GetExecutionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExecutionsResult>("alicloud:fnf/getExecutions:getExecutions", args ?? new GetExecutionsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the FnF Executions of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.149.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var ids = AliCloud.FNF.GetExecutions.Invoke(new()
        ///     {
        ///         FlowName = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "my-Execution-1",
        ///             "my-Execution-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["fnfExecutionId1"] = data.Alicloud_fn_f_executions.Ids.Executions[0].Id,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetExecutionsResult> Invoke(GetExecutionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExecutionsResult>("alicloud:fnf/getExecutions:getExecutions", args ?? new GetExecutionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExecutionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        /// <summary>
        /// The name of the flow.
        /// </summary>
        [Input("flowName", required: true)]
        public string FlowName { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Execution IDs. The value formats as `&lt;flow_name&gt;:&lt;execution_name&gt;`.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Execution name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetExecutionsArgs()
        {
        }
        public static new GetExecutionsArgs Empty => new GetExecutionsArgs();
    }

    public sealed class GetExecutionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        /// <summary>
        /// The name of the flow.
        /// </summary>
        [Input("flowName", required: true)]
        public Input<string> FlowName { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Execution IDs. The value formats as `&lt;flow_name&gt;:&lt;execution_name&gt;`.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Execution name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetExecutionsInvokeArgs()
        {
        }
        public static new GetExecutionsInvokeArgs Empty => new GetExecutionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetExecutionsResult
    {
        public readonly bool? EnableDetails;
        public readonly ImmutableArray<Outputs.GetExecutionsExecutionResult> Executions;
        public readonly string FlowName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetExecutionsResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetExecutionsExecutionResult> executions,

            string flowName,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            EnableDetails = enableDetails;
            Executions = executions;
            FlowName = flowName;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
