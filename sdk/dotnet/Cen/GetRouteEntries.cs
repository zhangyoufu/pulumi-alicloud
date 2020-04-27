// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    public static class GetRouteEntries
    {
        /// <summary>
        /// This data source provides CEN Route Entries available to the user.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteEntriesResult> InvokeAsync(GetRouteEntriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteEntriesResult>("alicloud:cen/getRouteEntries:getRouteEntries", args ?? new GetRouteEntriesArgs(), options.WithVersion());
    }


    public sealed class GetRouteEntriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry to query.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        /// <summary>
        /// ID of the CEN instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// ID of the route table of the VPC or VBR.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        public GetRouteEntriesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteEntriesResult
    {
        /// <summary>
        /// The destination CIDR block of the conflicted route entry.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// A list of CEN Route Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteEntriesEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the CEN child instance.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// ID of the route table.
        /// </summary>
        public readonly string RouteTableId;

        [OutputConstructor]
        private GetRouteEntriesResult(
            string? cidrBlock,

            ImmutableArray<Outputs.GetRouteEntriesEntryResult> entries,

            string id,

            string instanceId,

            string? outputFile,

            string routeTableId)
        {
            CidrBlock = cidrBlock;
            Entries = entries;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            RouteTableId = routeTableId;
        }
    }
}
