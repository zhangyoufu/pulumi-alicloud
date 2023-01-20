// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb.Outputs
{

    [OutputType]
    public sealed class DbInstancePlanPlanConfigScaleIn
    {
        /// <summary>
        /// The executed time of the Plan.
        /// </summary>
        public readonly string? ExecuteTime;
        /// <summary>
        /// The Cron Time of the plan.
        /// </summary>
        public readonly string? PlanCronTime;
        /// <summary>
        /// The segment Node Num of the Plan.
        /// </summary>
        public readonly string? SegmentNodeNum;

        [OutputConstructor]
        private DbInstancePlanPlanConfigScaleIn(
            string? executeTime,

            string? planCronTime,

            string? segmentNodeNum)
        {
            ExecuteTime = executeTime;
            PlanCronTime = planCronTime;
            SegmentNodeNum = segmentNodeNum;
        }
    }
}
