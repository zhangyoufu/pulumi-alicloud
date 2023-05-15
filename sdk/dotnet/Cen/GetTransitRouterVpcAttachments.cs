// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterVpcAttachments
    {
        /// <summary>
        /// This data source provides CEN Transit Router VPC Attachments available to the user.[What is Cen Transit Router VPC Attachments](https://help.aliyun.com/document_detail/261222.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Task<GetTransitRouterVpcAttachmentsResult> InvokeAsync(GetTransitRouterVpcAttachmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterVpcAttachmentsResult>("alicloud:cen/getTransitRouterVpcAttachments:getTransitRouterVpcAttachments", args ?? new GetTransitRouterVpcAttachmentsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides CEN Transit Router VPC Attachments available to the user.[What is Cen Transit Router VPC Attachments](https://help.aliyun.com/document_detail/261222.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Output<GetTransitRouterVpcAttachmentsResult> Invoke(GetTransitRouterVpcAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitRouterVpcAttachmentsResult>("alicloud:cen/getTransitRouterVpcAttachments:getTransitRouterVpcAttachments", args ?? new GetTransitRouterVpcAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitRouterVpcAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public string CenId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of resource id. The element value is same as `transit_router_id`.
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
        /// The status of the resource. Valid values `Attached`, `Attaching` and `Detaching`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The transit router ID.
        /// </summary>
        [Input("transitRouterId")]
        public string? TransitRouterId { get; set; }

        public GetTransitRouterVpcAttachmentsArgs()
        {
        }
        public static new GetTransitRouterVpcAttachmentsArgs Empty => new GetTransitRouterVpcAttachmentsArgs();
    }

    public sealed class GetTransitRouterVpcAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of resource id. The element value is same as `transit_router_id`.
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
        /// The status of the resource. Valid values `Attached`, `Attaching` and `Detaching`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The transit router ID.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        public GetTransitRouterVpcAttachmentsInvokeArgs()
        {
        }
        public static new GetTransitRouterVpcAttachmentsInvokeArgs Empty => new GetTransitRouterVpcAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitRouterVpcAttachmentsResult
    {
        /// <summary>
        /// A list of CEN Transit Router VPC Attachments. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRouterVpcAttachmentsAttachmentResult> Attachments;
        public readonly string CenId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the transit router attachment.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// ID of the transit router.
        /// </summary>
        public readonly string? TransitRouterId;

        [OutputConstructor]
        private GetTransitRouterVpcAttachmentsResult(
            ImmutableArray<Outputs.GetTransitRouterVpcAttachmentsAttachmentResult> attachments,

            string cenId,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? status,

            string? transitRouterId)
        {
            Attachments = attachments;
            CenId = cenId;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Status = status;
            TransitRouterId = transitRouterId;
        }
    }
}
