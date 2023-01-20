// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cms.Inputs
{

    public sealed class MetricRuleTemplateAlertTemplateEscalationsCriticalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The comparison operator of the threshold for warn-level alerts. Valid values: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanOrEqualToThreshold`, `LessThanThreshold`, `NotEqualToThreshold`, `GreaterThanYesterday`, `LessThanYesterday`, `GreaterThanLastWeek`, `LessThanLastWeek`, `GreaterThanLastPeriod`, `LessThanLastPeriod`.
        /// </summary>
        [Input("comparisonOperator")]
        public Input<string>? ComparisonOperator { get; set; }

        /// <summary>
        /// The statistical aggregation method for warn-level alerts.
        /// </summary>
        [Input("statistics")]
        public Input<string>? Statistics { get; set; }

        /// <summary>
        /// The threshold for warn-level alerts.
        /// </summary>
        [Input("threshold")]
        public Input<string>? Threshold { get; set; }

        /// <summary>
        /// The consecutive number of times for which the metric value is measured before a warn-level alert is triggered.
        /// </summary>
        [Input("times")]
        public Input<string>? Times { get; set; }

        public MetricRuleTemplateAlertTemplateEscalationsCriticalArgs()
        {
        }
        public static new MetricRuleTemplateAlertTemplateEscalationsCriticalArgs Empty => new MetricRuleTemplateAlertTemplateEscalationsCriticalArgs();
    }
}
