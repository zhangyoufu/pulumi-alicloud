// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Emrv2.Outputs
{

    [OutputType]
    public sealed class ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerTimeConstraint
    {
        public readonly string? EndTime;
        public readonly string? StartTime;

        [OutputConstructor]
        private ClusterNodeGroupAutoScalingPolicyScalingRuleMetricsTriggerTimeConstraint(
            string? endTime,

            string? startTime)
        {
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
