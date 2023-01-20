// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterRouteTablePropagations
    {
        /// <summary>
        /// This data source provides CEN Transit Router Route Table Propagations available to the user.[What is Cen Transit Router Route Table Propagations](https://help.aliyun.com/document_detail/261245.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
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
        ///     var @default = AliCloud.Cen.GetTransitRouterRouteTablePropagations.Invoke(new()
        ///     {
        ///         TransitRouterRouteTableId = "rtb-id1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType"] = @default.Apply(getTransitRouterRouteTablePropagationsResult =&gt; getTransitRouterRouteTablePropagationsResult).Apply(@default =&gt; @default.Apply(getTransitRouterRouteTablePropagationsResult =&gt; getTransitRouterRouteTablePropagationsResult.Propagations[0]?.ResourceType)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTransitRouterRouteTablePropagationsResult> InvokeAsync(GetTransitRouterRouteTablePropagationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterRouteTablePropagationsResult>("alicloud:cen/getTransitRouterRouteTablePropagations:getTransitRouterRouteTablePropagations", args ?? new GetTransitRouterRouteTablePropagationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides CEN Transit Router Route Table Propagations available to the user.[What is Cen Transit Router Route Table Propagations](https://help.aliyun.com/document_detail/261245.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
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
        ///     var @default = AliCloud.Cen.GetTransitRouterRouteTablePropagations.Invoke(new()
        ///     {
        ///         TransitRouterRouteTableId = "rtb-id1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstTransitRouterPeerAttachmentsTransitRouterAttachmentResourceType"] = @default.Apply(getTransitRouterRouteTablePropagationsResult =&gt; getTransitRouterRouteTablePropagationsResult).Apply(@default =&gt; @default.Apply(getTransitRouterRouteTablePropagationsResult =&gt; getTransitRouterRouteTablePropagationsResult.Propagations[0]?.ResourceType)),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTransitRouterRouteTablePropagationsResult> Invoke(GetTransitRouterRouteTablePropagationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitRouterRouteTablePropagationsResult>("alicloud:cen/getTransitRouterRouteTablePropagations:getTransitRouterRouteTablePropagations", args ?? new GetTransitRouterRouteTablePropagationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitRouterRouteTablePropagationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Table Association IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the route table, including `Active`, `Enabling`, `Disabling`, `Deleted`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// ID of the route table of the VPC or VBR.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public string TransitRouterRouteTableId { get; set; } = null!;

        public GetTransitRouterRouteTablePropagationsArgs()
        {
        }
        public static new GetTransitRouterRouteTablePropagationsArgs Empty => new GetTransitRouterRouteTablePropagationsArgs();
    }

    public sealed class GetTransitRouterRouteTablePropagationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Table Association IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the route table, including `Active`, `Enabling`, `Disabling`, `Deleted`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// ID of the route table of the VPC or VBR.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public Input<string> TransitRouterRouteTableId { get; set; } = null!;

        public GetTransitRouterRouteTablePropagationsInvokeArgs()
        {
        }
        public static new GetTransitRouterRouteTablePropagationsInvokeArgs Empty => new GetTransitRouterRouteTablePropagationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitRouterRouteTablePropagationsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CEN Transit Router Route Table Association IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of CEN Transit Router Route Table Propagations. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRouterRouteTablePropagationsPropagationResult> Propagations;
        /// <summary>
        /// The status of the route table.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// ID of the transit router route table.
        /// </summary>
        public readonly string TransitRouterRouteTableId;

        [OutputConstructor]
        private GetTransitRouterRouteTablePropagationsResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetTransitRouterRouteTablePropagationsPropagationResult> propagations,

            string? status,

            string transitRouterRouteTableId)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Propagations = propagations;
            Status = status;
            TransitRouterRouteTableId = transitRouterRouteTableId;
        }
    }
}
