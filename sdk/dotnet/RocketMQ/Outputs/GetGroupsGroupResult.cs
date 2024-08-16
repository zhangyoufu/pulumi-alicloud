// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ.Outputs
{

    [OutputType]
    public sealed class GetGroupsGroupResult
    {
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Specify the protocol applicable to the created Group ID. Valid values: `tcp`, `http`. Default to `tcp`.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.
        /// </summary>
        public readonly bool IndependentNaming;
        /// <summary>
        /// ID of the ONS Instance that owns the groups.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The ID of the group owner, which is the Alibaba Cloud UID.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Remark of the group.
        /// </summary>
        public readonly string Remark;
        /// <summary>
        /// A map of tags assigned to the Ons instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetGroupsGroupResult(
            string groupName,

            string groupType,

            string id,

            bool independentNaming,

            string instanceId,

            string owner,

            string remark,

            ImmutableDictionary<string, string> tags)
        {
            GroupName = groupName;
            GroupType = groupType;
            Id = id;
            IndependentNaming = independentNaming;
            InstanceId = instanceId;
            Owner = owner;
            Remark = remark;
            Tags = tags;
        }
    }
}
