// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci
{
    public static class GetVirtualNodes
    {
        /// <summary>
        /// This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
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
        ///     var ids = AliCloud.Eci.GetVirtualNodes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Eci.GetVirtualNodes.Invoke(new()
        ///     {
        ///         NameRegex = "^my-VirtualNode",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["eciVirtualNodeId1"] = ids.Apply(getVirtualNodesResult =&gt; getVirtualNodesResult.Nodes[0]?.Id),
        ///         ["eciVirtualNodeId2"] = nameRegex.Apply(getVirtualNodesResult =&gt; getVirtualNodesResult.Nodes[0]?.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetVirtualNodesResult> InvokeAsync(GetVirtualNodesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNodesResult>("alicloud:eci/getVirtualNodes:getVirtualNodes", args ?? new GetVirtualNodesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.145.0+.
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
        ///     var ids = AliCloud.Eci.GetVirtualNodes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     var nameRegex = AliCloud.Eci.GetVirtualNodes.Invoke(new()
        ///     {
        ///         NameRegex = "^my-VirtualNode",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["eciVirtualNodeId1"] = ids.Apply(getVirtualNodesResult =&gt; getVirtualNodesResult.Nodes[0]?.Id),
        ///         ["eciVirtualNodeId2"] = nameRegex.Apply(getVirtualNodesResult =&gt; getVirtualNodesResult.Nodes[0]?.Id),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualNodesResult> Invoke(GetVirtualNodesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNodesResult>("alicloud:eci/getVirtualNodes:getVirtualNodes", args ?? new GetVirtualNodesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNodesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Virtual Node IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Virtual Node name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The resource group ID. If when you create a GPU does not specify a resource group instance will automatically add the account's default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// VNode itself and by VNode created (ECI) the security group used by.
        /// </summary>
        [Input("securityGroupId")]
        public string? SecurityGroupId { get; set; }

        /// <summary>
        /// The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the virtual node.
        /// </summary>
        [Input("virtualNodeName")]
        public string? VirtualNodeName { get; set; }

        [Input("vswitchId")]
        public string? VswitchId { get; set; }

        public GetVirtualNodesArgs()
        {
        }
        public static new GetVirtualNodesArgs Empty => new GetVirtualNodesArgs();
    }

    public sealed class GetVirtualNodesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Virtual Node IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Virtual Node name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The resource group ID. If when you create a GPU does not specify a resource group instance will automatically add the account's default resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// VNode itself and by VNode created (ECI) the security group used by.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the virtual node.
        /// </summary>
        [Input("virtualNodeName")]
        public Input<string>? VirtualNodeName { get; set; }

        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GetVirtualNodesInvokeArgs()
        {
        }
        public static new GetVirtualNodesInvokeArgs Empty => new GetVirtualNodesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualNodesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly ImmutableArray<Outputs.GetVirtualNodesNodeResult> Nodes;
        public readonly string? OutputFile;
        public readonly string? ResourceGroupId;
        public readonly string? SecurityGroupId;
        public readonly string? Status;
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly string? VirtualNodeName;
        public readonly string? VswitchId;

        [OutputConstructor]
        private GetVirtualNodesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            ImmutableArray<Outputs.GetVirtualNodesNodeResult> nodes,

            string? outputFile,

            string? resourceGroupId,

            string? securityGroupId,

            string? status,

            ImmutableDictionary<string, string>? tags,

            string? virtualNodeName,

            string? vswitchId)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            Nodes = nodes;
            OutputFile = outputFile;
            ResourceGroupId = resourceGroupId;
            SecurityGroupId = securityGroupId;
            Status = status;
            Tags = tags;
            VirtualNodeName = virtualNodeName;
            VswitchId = vswitchId;
        }
    }
}
