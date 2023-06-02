// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Outputs
{

    [OutputType]
    public sealed class ApplicationScalingRuleScalingRuleMetricMetric
    {
        /// <summary>
        /// According to different `metric_type`, set the target value of the corresponding monitoring index.
        /// </summary>
        public readonly int? MetricTargetAverageUtilization;
        /// <summary>
        /// Monitoring indicator trigger condition. Valid values: `CPU`, `MEMORY`, `tcpActiveConn`, `QPS`, `RT`, `SLB_QPS`, `SLB_RT`, `INTRANET_SLB_QPS` and `INTRANET_SLB_RT`. The values are described as follows:
        /// - CPU: CPU usage.
        /// - MEMORY: MEMORY usage.
        /// - tcpActiveConn: The average number of TCP active connections for a single instance in 30 seconds.
        /// - QPS: The average QPS of a single instance within 1 minute of JAVA application.
        /// - RT: The average response time of all service interfaces within 1 minute of JAVA application.
        /// - SLB_QPS: The average public network SLB QPS of a single instance within 15 seconds.
        /// - SLB_RT: The average response time of public network SLB within 15 seconds.
        /// - INTRANET_SLB_QPS: The average private network SLB QPS of a single instance within 15 seconds.
        /// - INTRANET_SLB_RT: The average response time of private network SLB within 15 seconds.
        /// **NOTE:** From version 1.206.0, `metric_type` can be set to `QPS`, `RT`, `INTRANET_SLB_QPS`, `INTRANET_SLB_RT`.
        /// </summary>
        public readonly string? MetricType;
        /// <summary>
        /// SLB ID.
        /// </summary>
        public readonly string? SlbId;
        /// <summary>
        /// The log store of the Log Service.
        /// </summary>
        public readonly string? SlbLogStore;
        /// <summary>
        /// The project of the Log Service.
        /// </summary>
        public readonly string? SlbProject;
        /// <summary>
        /// SLB listening port.
        /// </summary>
        public readonly string? Vport;

        [OutputConstructor]
        private ApplicationScalingRuleScalingRuleMetricMetric(
            int? metricTargetAverageUtilization,

            string? metricType,

            string? slbId,

            string? slbLogStore,

            string? slbProject,

            string? vport)
        {
            MetricTargetAverageUtilization = metricTargetAverageUtilization;
            MetricType = metricType;
            SlbId = slbId;
            SlbLogStore = slbLogStore;
            SlbProject = slbProject;
            Vport = vport;
        }
    }
}
