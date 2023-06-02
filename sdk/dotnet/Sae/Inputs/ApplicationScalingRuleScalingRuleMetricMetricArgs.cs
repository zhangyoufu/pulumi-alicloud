// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class ApplicationScalingRuleScalingRuleMetricMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// According to different `metric_type`, set the target value of the corresponding monitoring index.
        /// </summary>
        [Input("metricTargetAverageUtilization")]
        public Input<int>? MetricTargetAverageUtilization { get; set; }

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
        [Input("metricType")]
        public Input<string>? MetricType { get; set; }

        /// <summary>
        /// SLB ID.
        /// </summary>
        [Input("slbId")]
        public Input<string>? SlbId { get; set; }

        /// <summary>
        /// The log store of the Log Service.
        /// </summary>
        [Input("slbLogStore")]
        public Input<string>? SlbLogStore { get; set; }

        /// <summary>
        /// The project of the Log Service.
        /// </summary>
        [Input("slbProject")]
        public Input<string>? SlbProject { get; set; }

        /// <summary>
        /// SLB listening port.
        /// </summary>
        [Input("vport")]
        public Input<string>? Vport { get; set; }

        public ApplicationScalingRuleScalingRuleMetricMetricArgs()
        {
        }
        public static new ApplicationScalingRuleScalingRuleMetricMetricArgs Empty => new ApplicationScalingRuleScalingRuleMetricMetricArgs();
    }
}
