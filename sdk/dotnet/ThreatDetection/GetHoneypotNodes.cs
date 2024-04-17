// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ThreatDetection
{
    public static class GetHoneypotNodes
    {
        /// <summary>
        /// This data source provides Threat Detection Honeypot Node available to the user.[What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotnode)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// 
        /// ## Example Usage
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
        ///     var @default = AliCloud.ThreatDetection.GetHoneypotNodes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAlicloudThreatDetectionHoneypotNode.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionHoneypotNodeExampleId"] = @default.Apply(@default =&gt; @default.Apply(getHoneypotNodesResult =&gt; getHoneypotNodesResult.Nodes[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetHoneypotNodesResult> InvokeAsync(GetHoneypotNodesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHoneypotNodesResult>("alicloud:threatdetection/getHoneypotNodes:getHoneypotNodes", args ?? new GetHoneypotNodesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Threat Detection Honeypot Node available to the user.[What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotnode)
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
        /// 
        /// ## Example Usage
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
        ///     var @default = AliCloud.ThreatDetection.GetHoneypotNodes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             defaultAlicloudThreatDetectionHoneypotNode.Id,
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudThreatDetectionHoneypotNodeExampleId"] = @default.Apply(@default =&gt; @default.Apply(getHoneypotNodesResult =&gt; getHoneypotNodesResult.Nodes[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetHoneypotNodesResult> Invoke(GetHoneypotNodesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHoneypotNodesResult>("alicloud:threatdetection/getHoneypotNodes:getHoneypotNodes", args ?? new GetHoneypotNodesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHoneypotNodesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Honeypot Node IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Honeypot Node name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Honeypot management node id.
        /// </summary>
        [Input("nodeId")]
        public string? NodeId { get; set; }

        /// <summary>
        /// The name of the management node.
        /// </summary>
        [Input("nodeName")]
        public string? NodeName { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        public GetHoneypotNodesArgs()
        {
        }
        public static new GetHoneypotNodesArgs Empty => new GetHoneypotNodesArgs();
    }

    public sealed class GetHoneypotNodesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Honeypot Node IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Honeypot Node name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// Honeypot management node id.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// The name of the management node.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        public GetHoneypotNodesInvokeArgs()
        {
        }
        public static new GetHoneypotNodesInvokeArgs Empty => new GetHoneypotNodesInvokeArgs();
    }


    [OutputType]
    public sealed class GetHoneypotNodesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Honeypot Node IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Honeypot Node names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Honeypot management node id.
        /// </summary>
        public readonly string? NodeId;
        /// <summary>
        /// Management node name.
        /// </summary>
        public readonly string? NodeName;
        /// <summary>
        /// A list of Honeypot Node Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHoneypotNodesNodeResult> Nodes;
        public readonly string? OutputFile;
        public readonly int? PageNumber;
        public readonly int? PageSize;

        [OutputConstructor]
        private GetHoneypotNodesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? nodeId,

            string? nodeName,

            ImmutableArray<Outputs.GetHoneypotNodesNodeResult> nodes,

            string? outputFile,

            int? pageNumber,

            int? pageSize)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            NodeId = nodeId;
            NodeName = nodeName;
            Nodes = nodes;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
        }
    }
}
