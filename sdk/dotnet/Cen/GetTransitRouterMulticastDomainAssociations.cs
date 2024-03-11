// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterMulticastDomainAssociations
    {
        /// <summary>
        /// This data source provides the Cen Transit Router Multicast Domain Associations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.195.0+.
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
        ///     var ids = AliCloud.Cen.GetTransitRouterMulticastDomainAssociations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         TransitRouterMulticastDomainId = "your_transit_router_multicast_domain_id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cenTransitRouterMulticastDomainId0"] = ids.Apply(getTransitRouterMulticastDomainAssociationsResult =&gt; getTransitRouterMulticastDomainAssociationsResult.Associations[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTransitRouterMulticastDomainAssociationsResult> InvokeAsync(GetTransitRouterMulticastDomainAssociationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterMulticastDomainAssociationsResult>("alicloud:cen/getTransitRouterMulticastDomainAssociations:getTransitRouterMulticastDomainAssociations", args ?? new GetTransitRouterMulticastDomainAssociationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Cen Transit Router Multicast Domain Associations of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.195.0+.
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
        ///     var ids = AliCloud.Cen.GetTransitRouterMulticastDomainAssociations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "example_id",
        ///         },
        ///         TransitRouterMulticastDomainId = "your_transit_router_multicast_domain_id",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["cenTransitRouterMulticastDomainId0"] = ids.Apply(getTransitRouterMulticastDomainAssociationsResult =&gt; getTransitRouterMulticastDomainAssociationsResult.Associations[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTransitRouterMulticastDomainAssociationsResult> Invoke(GetTransitRouterMulticastDomainAssociationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitRouterMulticastDomainAssociationsResult>("alicloud:cen/getTransitRouterMulticastDomainAssociations:getTransitRouterMulticastDomainAssociations", args ?? new GetTransitRouterMulticastDomainAssociationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitRouterMulticastDomainAssociationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Transit Router Multicast Domain Association IDs.
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
        /// The ID of the resource associated with the multicast domain.
        /// </summary>
        [Input("resourceId")]
        public string? ResourceId { get; set; }

        /// <summary>
        /// The type of resource associated with the multicast domain. Valid Value: `VPC`.
        /// </summary>
        [Input("resourceType")]
        public string? ResourceType { get; set; }

        /// <summary>
        /// The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        [Input("transitRouterAttachmentId")]
        public string? TransitRouterAttachmentId { get; set; }

        /// <summary>
        /// The ID of the multicast domain.
        /// </summary>
        [Input("transitRouterMulticastDomainId", required: true)]
        public string TransitRouterMulticastDomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        [Input("vswitchId")]
        public string? VswitchId { get; set; }

        public GetTransitRouterMulticastDomainAssociationsArgs()
        {
        }
        public static new GetTransitRouterMulticastDomainAssociationsArgs Empty => new GetTransitRouterMulticastDomainAssociationsArgs();
    }

    public sealed class GetTransitRouterMulticastDomainAssociationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Transit Router Multicast Domain Association IDs.
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
        /// The ID of the resource associated with the multicast domain.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The type of resource associated with the multicast domain. Valid Value: `VPC`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        [Input("transitRouterAttachmentId")]
        public Input<string>? TransitRouterAttachmentId { get; set; }

        /// <summary>
        /// The ID of the multicast domain.
        /// </summary>
        [Input("transitRouterMulticastDomainId", required: true)]
        public Input<string> TransitRouterMulticastDomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GetTransitRouterMulticastDomainAssociationsInvokeArgs()
        {
        }
        public static new GetTransitRouterMulticastDomainAssociationsInvokeArgs Empty => new GetTransitRouterMulticastDomainAssociationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitRouterMulticastDomainAssociationsResult
    {
        /// <summary>
        /// A list of Cen Transit Router Multicast Domain Associations. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRouterMulticastDomainAssociationsAssociationResult> Associations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The ID of the resource associated with the multicast domain.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The type of resource associated with the multicast domain.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The status of the associated resource.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        public readonly string? TransitRouterAttachmentId;
        /// <summary>
        /// The ID of the multicast domain.
        /// </summary>
        public readonly string TransitRouterMulticastDomainId;
        /// <summary>
        /// The ID of the vSwitch.
        /// </summary>
        public readonly string? VswitchId;

        [OutputConstructor]
        private GetTransitRouterMulticastDomainAssociationsResult(
            ImmutableArray<Outputs.GetTransitRouterMulticastDomainAssociationsAssociationResult> associations,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? resourceId,

            string? resourceType,

            string? status,

            string? transitRouterAttachmentId,

            string transitRouterMulticastDomainId,

            string? vswitchId)
        {
            Associations = associations;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ResourceId = resourceId;
            ResourceType = resourceType;
            Status = status;
            TransitRouterAttachmentId = transitRouterAttachmentId;
            TransitRouterMulticastDomainId = transitRouterMulticastDomainId;
            VswitchId = vswitchId;
        }
    }
}
