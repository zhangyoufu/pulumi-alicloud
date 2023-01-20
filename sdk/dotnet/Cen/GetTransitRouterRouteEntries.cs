// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterRouteEntries
    {
        /// <summary>
        /// This data source provides CEN Transit Router Route Entries available to the user.[What is Cen Transit Router Route Entries](https://help.aliyun.com/document_detail/260941.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Task<GetTransitRouterRouteEntriesResult> InvokeAsync(GetTransitRouterRouteEntriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterRouteEntriesResult>("alicloud:cen/getTransitRouterRouteEntries:getTransitRouterRouteEntries", args ?? new GetTransitRouterRouteEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides CEN Transit Router Route Entries available to the user.[What is Cen Transit Router Route Entries](https://help.aliyun.com/document_detail/260941.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Output<GetTransitRouterRouteEntriesResult> Invoke(GetTransitRouterRouteEntriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitRouterRouteEntriesResult>("alicloud:cen/getTransitRouterRouteEntries:getTransitRouterRouteEntries", args ?? new GetTransitRouterRouteEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitRouterRouteEntriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Entry IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        [Input("transitRouterRouteEntryIds")]
        private List<string>? _transitRouterRouteEntryIds;

        /// <summary>
        /// A list of ID of the cen transit router route entry.
        /// </summary>
        public List<string> TransitRouterRouteEntryIds
        {
            get => _transitRouterRouteEntryIds ?? (_transitRouterRouteEntryIds = new List<string>());
            set => _transitRouterRouteEntryIds = value;
        }

        [Input("transitRouterRouteEntryNames")]
        private List<string>? _transitRouterRouteEntryNames;

        /// <summary>
        /// A list of name of the cen transit router route entry.
        /// </summary>
        public List<string> TransitRouterRouteEntryNames
        {
            get => _transitRouterRouteEntryNames ?? (_transitRouterRouteEntryNames = new List<string>());
            set => _transitRouterRouteEntryNames = value;
        }

        /// <summary>
        /// The status of the resource.Valid values `Creating`, `Active` and `Deleting`.
        /// </summary>
        [Input("transitRouterRouteEntryStatus")]
        public string? TransitRouterRouteEntryStatus { get; set; }

        /// <summary>
        /// ID of the CEN Transit Router Route Table.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public string TransitRouterRouteTableId { get; set; } = null!;

        public GetTransitRouterRouteEntriesArgs()
        {
        }
        public static new GetTransitRouterRouteEntriesArgs Empty => new GetTransitRouterRouteEntriesArgs();
    }

    public sealed class GetTransitRouterRouteEntriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Entry IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("transitRouterRouteEntryIds")]
        private InputList<string>? _transitRouterRouteEntryIds;

        /// <summary>
        /// A list of ID of the cen transit router route entry.
        /// </summary>
        public InputList<string> TransitRouterRouteEntryIds
        {
            get => _transitRouterRouteEntryIds ?? (_transitRouterRouteEntryIds = new InputList<string>());
            set => _transitRouterRouteEntryIds = value;
        }

        [Input("transitRouterRouteEntryNames")]
        private InputList<string>? _transitRouterRouteEntryNames;

        /// <summary>
        /// A list of name of the cen transit router route entry.
        /// </summary>
        public InputList<string> TransitRouterRouteEntryNames
        {
            get => _transitRouterRouteEntryNames ?? (_transitRouterRouteEntryNames = new InputList<string>());
            set => _transitRouterRouteEntryNames = value;
        }

        /// <summary>
        /// The status of the resource.Valid values `Creating`, `Active` and `Deleting`.
        /// </summary>
        [Input("transitRouterRouteEntryStatus")]
        public Input<string>? TransitRouterRouteEntryStatus { get; set; }

        /// <summary>
        /// ID of the CEN Transit Router Route Table.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public Input<string> TransitRouterRouteTableId { get; set; } = null!;

        public GetTransitRouterRouteEntriesInvokeArgs()
        {
        }
        public static new GetTransitRouterRouteEntriesInvokeArgs Empty => new GetTransitRouterRouteEntriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitRouterRouteEntriesResult
    {
        /// <summary>
        /// A list of CEN Route Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRouterRouteEntriesEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CEN Transit Router Route Entry IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of CEN Transit Router Route Entry Names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;
        public readonly ImmutableArray<string> TransitRouterRouteEntryIds;
        public readonly ImmutableArray<string> TransitRouterRouteEntryNames;
        /// <summary>
        /// The status of the route entry in CEN.
        /// </summary>
        public readonly string? TransitRouterRouteEntryStatus;
        public readonly string TransitRouterRouteTableId;

        [OutputConstructor]
        private GetTransitRouterRouteEntriesResult(
            ImmutableArray<Outputs.GetTransitRouterRouteEntriesEntryResult> entries,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableArray<string> transitRouterRouteEntryIds,

            ImmutableArray<string> transitRouterRouteEntryNames,

            string? transitRouterRouteEntryStatus,

            string transitRouterRouteTableId)
        {
            Entries = entries;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            TransitRouterRouteEntryIds = transitRouterRouteEntryIds;
            TransitRouterRouteEntryNames = transitRouterRouteEntryNames;
            TransitRouterRouteEntryStatus = transitRouterRouteEntryStatus;
            TransitRouterRouteTableId = transitRouterRouteTableId;
        }
    }
}
