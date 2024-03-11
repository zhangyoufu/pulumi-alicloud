// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterPrefixListAssociations
    {
        /// <summary>
        /// This data source provides the Cen Transit Router Prefix List Associations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.188.0+.
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
        ///     var @default = AliCloud.Cen.GetTransitRouterPrefixListAssociations.Invoke(new()
        ///     {
        ///         TransitRouterId = "tr-6ehx7q2jze8ch5ji0****",
        ///         TransitRouterTableId = "vtb-6ehgc262hr170qgyc****",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cenTransitRouterPrefixListAssociationId"] = @default.Apply(@default =&gt; @default.Apply(getTransitRouterPrefixListAssociationsResult =&gt; getTransitRouterPrefixListAssociationsResult.Associations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTransitRouterPrefixListAssociationsResult> InvokeAsync(GetTransitRouterPrefixListAssociationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterPrefixListAssociationsResult>("alicloud:cen/getTransitRouterPrefixListAssociations:getTransitRouterPrefixListAssociations", args ?? new GetTransitRouterPrefixListAssociationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cen Transit Router Prefix List Associations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.188.0+.
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
        ///     var @default = AliCloud.Cen.GetTransitRouterPrefixListAssociations.Invoke(new()
        ///     {
        ///         TransitRouterId = "tr-6ehx7q2jze8ch5ji0****",
        ///         TransitRouterTableId = "vtb-6ehgc262hr170qgyc****",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cenTransitRouterPrefixListAssociationId"] = @default.Apply(@default =&gt; @default.Apply(getTransitRouterPrefixListAssociationsResult =&gt; getTransitRouterPrefixListAssociationsResult.Associations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTransitRouterPrefixListAssociationsResult> Invoke(GetTransitRouterPrefixListAssociationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitRouterPrefixListAssociationsResult>("alicloud:cen/getTransitRouterPrefixListAssociations:getTransitRouterPrefixListAssociations", args ?? new GetTransitRouterPrefixListAssociationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitRouterPrefixListAssociationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Cen Transit Router Prefix List Association IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the prefix list belongs.
        /// </summary>
        [Input("ownerUid")]
        public int? OwnerUid { get; set; }

        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The ID of the prefix list.
        /// </summary>
        [Input("prefixListId")]
        public string? PrefixListId { get; set; }

        /// <summary>
        /// The status of the prefix list.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public string TransitRouterId { get; set; } = null!;

        /// <summary>
        /// The ID of the route table of the transit router.
        /// </summary>
        [Input("transitRouterTableId", required: true)]
        public string TransitRouterTableId { get; set; } = null!;

        public GetTransitRouterPrefixListAssociationsArgs()
        {
        }
        public static new GetTransitRouterPrefixListAssociationsArgs Empty => new GetTransitRouterPrefixListAssociationsArgs();
    }

    public sealed class GetTransitRouterPrefixListAssociationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Cen Transit Router Prefix List Association IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ID of the Alibaba Cloud account to which the prefix list belongs.
        /// </summary>
        [Input("ownerUid")]
        public Input<int>? OwnerUid { get; set; }

        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The ID of the prefix list.
        /// </summary>
        [Input("prefixListId")]
        public Input<string>? PrefixListId { get; set; }

        /// <summary>
        /// The status of the prefix list.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        /// <summary>
        /// The ID of the route table of the transit router.
        /// </summary>
        [Input("transitRouterTableId", required: true)]
        public Input<string> TransitRouterTableId { get; set; } = null!;

        public GetTransitRouterPrefixListAssociationsInvokeArgs()
        {
        }
        public static new GetTransitRouterPrefixListAssociationsInvokeArgs Empty => new GetTransitRouterPrefixListAssociationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitRouterPrefixListAssociationsResult
    {
        public readonly ImmutableArray<Outputs.GetTransitRouterPrefixListAssociationsAssociationResult> Associations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly int? OwnerUid;
        public readonly int? PageNumber;
        public readonly int? PageSize;
        public readonly string? PrefixListId;
        public readonly string? Status;
        public readonly string TransitRouterId;
        public readonly string TransitRouterTableId;

        [OutputConstructor]
        private GetTransitRouterPrefixListAssociationsResult(
            ImmutableArray<Outputs.GetTransitRouterPrefixListAssociationsAssociationResult> associations,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            int? ownerUid,

            int? pageNumber,

            int? pageSize,

            string? prefixListId,

            string? status,

            string transitRouterId,

            string transitRouterTableId)
        {
            Associations = associations;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            OwnerUid = ownerUid;
            PageNumber = pageNumber;
            PageSize = pageSize;
            PrefixListId = prefixListId;
            Status = status;
            TransitRouterId = transitRouterId;
            TransitRouterTableId = transitRouterTableId;
        }
    }
}
