// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.VideoSurveillance.Outputs
{

    [OutputType]
    public sealed class GetSystemGroupsGroupResult
    {
        /// <summary>
        /// The App Name of Group.
        /// </summary>
        public readonly string App;
        /// <summary>
        /// The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        /// </summary>
        public readonly string Callback;
        /// <summary>
        /// The creation time of the Group.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether to open Group.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Space of national standard ID. **NOTE:** Available only in the national standard access space.
        /// </summary>
        public readonly string GbId;
        /// <summary>
        /// Space of national standard signaling server address. **NOTE:** Available only in the national standard access space.
        /// </summary>
        public readonly string GbIp;
        /// <summary>
        /// The ID of Group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The name of Group.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// The ID of the Group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The use of the access protocol support `gb28181`,`rtmp`(Real Time Messaging Protocol).
        /// </summary>
        public readonly string InProtocol;
        /// <summary>
        /// The use of space play Protocol multi-valued separate them with commas (,). Valid values: `flv`,`hls`, `rtmp`(Real Time Messaging Protocol).
        /// </summary>
        public readonly string OutProtocol;
        /// <summary>
        /// -The domain name of plan streaming used by the group.
        /// </summary>
        public readonly string PlayDomain;
        /// <summary>
        /// The domain name of push streaming used by the group.
        /// </summary>
        public readonly string PushDomain;
        /// <summary>
        /// The Device statistics of Group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemGroupsGroupStatResult> Stats;

        [OutputConstructor]
        private GetSystemGroupsGroupResult(
            string app,

            string callback,

            string createTime,

            string description,

            bool enabled,

            string gbId,

            string gbIp,

            string groupId,

            string groupName,

            string id,

            string inProtocol,

            string outProtocol,

            string playDomain,

            string pushDomain,

            ImmutableArray<Outputs.GetSystemGroupsGroupStatResult> stats)
        {
            App = app;
            Callback = callback;
            CreateTime = createTime;
            Description = description;
            Enabled = enabled;
            GbId = gbId;
            GbIp = gbIp;
            GroupId = groupId;
            GroupName = groupName;
            Id = id;
            InProtocol = inProtocol;
            OutProtocol = outProtocol;
            PlayDomain = playDomain;
            PushDomain = pushDomain;
            Stats = stats;
        }
    }
}
