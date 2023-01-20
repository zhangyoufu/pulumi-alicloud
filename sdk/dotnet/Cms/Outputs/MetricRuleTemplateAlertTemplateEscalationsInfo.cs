// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Outputs
{

    [OutputType]
    public sealed class MetricRuleTemplateAlertTemplateEscalationsInfo
    {
        /// <summary>
        /// The comparison operator of the threshold for warn-level alerts. Valid values: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanOrEqualToThreshold`, `LessThanThreshold`, `NotEqualToThreshold`, `GreaterThanYesterday`, `LessThanYesterday`, `GreaterThanLastWeek`, `LessThanLastWeek`, `GreaterThanLastPeriod`, `LessThanLastPeriod`.
        /// </summary>
        public readonly string? ComparisonOperator;
        /// <summary>
        /// The statistical aggregation method for warn-level alerts.
        /// </summary>
        public readonly string? Statistics;
        /// <summary>
        /// The threshold for warn-level alerts.
        /// </summary>
        public readonly string? Threshold;
        /// <summary>
        /// The consecutive number of times for which the metric value is measured before a warn-level alert is triggered.
        /// </summary>
        public readonly string? Times;

        [OutputConstructor]
        private MetricRuleTemplateAlertTemplateEscalationsInfo(
            string? comparisonOperator,

            string? statistics,

            string? threshold,

            string? times)
        {
            ComparisonOperator = comparisonOperator;
            Statistics = statistics;
            Threshold = threshold;
            Times = times;
        }
    }
}
