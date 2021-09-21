// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms.Outputs
{

    [OutputType]
    public sealed class GetDispatchRulesRuleGroupRuleResult
    {
        public readonly int GroupId;
        /// <summary>
        /// The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
        /// </summary>
        public readonly int GroupInterval;
        /// <summary>
        /// The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
        /// </summary>
        public readonly int GroupWaitTime;
        /// <summary>
        /// The fields that are used to group events. Events with the same field content are assigned to a group. Alerts with the same specified grouping field are sent to the handler in separate notifications.
        /// </summary>
        public readonly ImmutableArray<string> GroupingFields;
        /// <summary>
        /// The silence period of repeated alerts. All alerts are repeatedly sent at specified intervals until the alerts are cleared. The minimum value is 61. Default to 600.
        /// </summary>
        public readonly int RepeatInterval;

        [OutputConstructor]
        private GetDispatchRulesRuleGroupRuleResult(
            int groupId,

            int groupInterval,

            int groupWaitTime,

            ImmutableArray<string> groupingFields,

            int repeatInterval)
        {
            GroupId = groupId;
            GroupInterval = groupInterval;
            GroupWaitTime = groupWaitTime;
            GroupingFields = groupingFields;
            RepeatInterval = repeatInterval;
        }
    }
}
