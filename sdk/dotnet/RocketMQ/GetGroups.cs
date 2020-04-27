// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    public static class GetGroups
    {
        /// <summary>
        /// This data source provides a list of ONS Groups in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.53.0+
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("alicloud:rocketmq/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithVersion());
    }


    public sealed class GetGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter results by the group name. 
        /// </summary>
        [Input("groupIdRegex")]
        public string? GroupIdRegex { get; set; }

        /// <summary>
        /// ID of the ONS Instance that owns the groups.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly string? GroupIdRegex;
        /// <summary>
        /// A list of groups. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of group names.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetGroupsResult(
            string? groupIdRegex,

            ImmutableArray<Outputs.GetGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? outputFile)
        {
            GroupIdRegex = groupIdRegex;
            Groups = groups;
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            OutputFile = outputFile;
        }
    }
}
