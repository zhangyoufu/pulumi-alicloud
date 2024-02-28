// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FNF
{
    /// <summary>
    /// Provides a Serverless Workflow Execution resource.
    /// 
    /// For information about Serverless Workflow Execution and how to use it, see [What is Execution](https://www.alibabacloud.com/help/en/doc-detail/122628.html).
    /// 
    /// &gt; **NOTE:** Available since v1.149.0+.
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
    ///     var name = config.Get("name") ?? "tf-example-fnfflow";
    ///     var defaultRole = new AliCloud.Ram.Role("defaultRole", new()
    ///     {
    ///         Document = @"  {
    ///     ""Statement"": [
    ///       {
    ///         ""Action"": ""sts:AssumeRole"",
    ///         ""Effect"": ""Allow"",
    ///         ""Principal"": {
    ///           ""Service"": [
    ///             ""fnf.aliyuncs.com""
    ///           ]
    ///         }
    ///       }
    ///     ],
    ///     ""Version"": ""1""
    ///   }
    /// ",
    ///     });
    /// 
    ///     var defaultFlow = new AliCloud.FNF.Flow("defaultFlow", new()
    ///     {
    ///         Definition = @"  version: v1beta1
    ///   type: flow
    ///   steps:
    ///     - type: wait
    ///       name: custom_wait
    ///       duration: $.wait
    /// ",
    ///         RoleArn = defaultRole.Arn,
    ///         Description = "Test for terraform fnf_flow.",
    ///         Type = "FDL",
    ///     });
    /// 
    ///     var defaultExecution = new AliCloud.FNF.Execution("defaultExecution", new()
    ///     {
    ///         ExecutionName = name,
    ///         FlowName = defaultFlow.Name,
    ///         Input = "{\"wait\": 600}",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Serverless Workflow Execution can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:fnf/execution:Execution example &lt;flow_name&gt;:&lt;execution_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:fnf/execution:Execution")]
    public partial class Execution : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the execution.
        /// </summary>
        [Output("executionName")]
        public Output<string> ExecutionName { get; private set; } = null!;

        /// <summary>
        /// The name of the flow.
        /// </summary>
        [Output("flowName")]
        public Output<string> FlowName { get; private set; } = null!;

        /// <summary>
        /// The Input information for this execution.
        /// </summary>
        [Output("input")]
        public Output<string?> Input { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `Stopped`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Execution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Execution(string name, ExecutionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:fnf/execution:Execution", name, args ?? new ExecutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Execution(string name, Input<string> id, ExecutionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:fnf/execution:Execution", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Execution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Execution Get(string name, Input<string> id, ExecutionState? state = null, CustomResourceOptions? options = null)
        {
            return new Execution(name, id, state, options);
        }
    }

    public sealed class ExecutionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the execution.
        /// </summary>
        [Input("executionName", required: true)]
        public Input<string> ExecutionName { get; set; } = null!;

        /// <summary>
        /// The name of the flow.
        /// </summary>
        [Input("flowName", required: true)]
        public Input<string> FlowName { get; set; } = null!;

        /// <summary>
        /// The Input information for this execution.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExecutionArgs()
        {
        }
        public static new ExecutionArgs Empty => new ExecutionArgs();
    }

    public sealed class ExecutionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the execution.
        /// </summary>
        [Input("executionName")]
        public Input<string>? ExecutionName { get; set; }

        /// <summary>
        /// The name of the flow.
        /// </summary>
        [Input("flowName")]
        public Input<string>? FlowName { get; set; }

        /// <summary>
        /// The Input information for this execution.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ExecutionState()
        {
        }
        public static new ExecutionState Empty => new ExecutionState();
    }
}
