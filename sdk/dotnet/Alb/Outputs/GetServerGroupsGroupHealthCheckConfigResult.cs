// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetServerGroupsGroupHealthCheckConfigResult
    {
        /// <summary>
        /// The status code for a successful health check. Multiple status codes can be specified as a list. Valid values: `http_2xx`, `http_3xx`, `http_4xx`, and `http_5xx`. Default value: `http_2xx`. **NOTE:** This parameter exists if the `HealthCheckProtocol` parameter is set to `HTTP`.
        /// </summary>
        public readonly ImmutableArray<string> HealthCheckCodes;
        /// <summary>
        /// The port of the backend server that is used for health checks. Valid values: `0` to `65535`. Default value: `0`. A value of `0` indicates that a backend server port is used for health checks.
        /// </summary>
        public readonly int HealthCheckConnectPort;
        /// <summary>
        /// Indicates whether health checks are enabled. Valid values: `true`, `false`. Default value: `true`.
        /// </summary>
        public readonly bool HealthCheckEnabled;
        /// <summary>
        /// The domain name that is used for health checks.
        /// </summary>
        public readonly string HealthCheckHost;
        /// <summary>
        /// HTTP protocol version. Valid values: `HTTP1.0` and `HTTP1.1`. Default value: `HTTP1.1`. **NOTE:** This parameter exists if the `HealthCheckProtocol` parameter is set to `HTTP`.
        /// </summary>
        public readonly string HealthCheckHttpVersion;
        /// <summary>
        /// The time interval between two consecutive health checks. Unit: seconds. Valid values: `1` to `50`. Default value: `2`.
        /// </summary>
        public readonly int HealthCheckInterval;
        /// <summary>
        /// Health check method. Valid values: `GET` and `HEAD`. Default: `GET`. **NOTE:** This parameter exists if the `HealthCheckProtocol` parameter is set to `HTTP`.
        /// </summary>
        public readonly string HealthCheckMethod;
        /// <summary>
        /// The forwarding rule path of health checks. **NOTE:** This parameter exists if the `HealthCheckProtocol` parameter is set to `HTTP`.
        /// </summary>
        public readonly string HealthCheckPath;
        /// <summary>
        /// Health check protocol. Valid values: `HTTP` and `TCP`.
        /// </summary>
        public readonly string HealthCheckProtocol;
        /// <summary>
        /// The timeout period of a health check response. If a backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the ECS instance is considered unhealthy. Unit: seconds. Valid values: `1` to `300`. Default value: `5`. **NOTE:** If the value of the `HealthCHeckTimeout` parameter is smaller than that of the `HealthCheckInterval` parameter, the value of the `HealthCHeckTimeout` parameter is ignored and the value of the `HealthCheckInterval` parameter is regarded as the timeout period.
        /// </summary>
        public readonly int HealthCheckTimeout;
        /// <summary>
        /// The number of health checks that an unhealthy backend server must pass consecutively before it is declared healthy. In this case, the health check state is changed from fail to success. Valid values: `2` to `10`. Default value: `3`.
        /// </summary>
        public readonly int HealthyThreshold;
        /// <summary>
        /// The number of consecutive health checks that a healthy backend server must consecutively fail before it is declared unhealthy. In this case, the health check state is changed from success to fail. Valid values: `2` to `10`. Default value: `3`.
        /// </summary>
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private GetServerGroupsGroupHealthCheckConfigResult(
            ImmutableArray<string> healthCheckCodes,

            int healthCheckConnectPort,

            bool healthCheckEnabled,

            string healthCheckHost,

            string healthCheckHttpVersion,

            int healthCheckInterval,

            string healthCheckMethod,

            string healthCheckPath,

            string healthCheckProtocol,

            int healthCheckTimeout,

            int healthyThreshold,

            int unhealthyThreshold)
        {
            HealthCheckCodes = healthCheckCodes;
            HealthCheckConnectPort = healthCheckConnectPort;
            HealthCheckEnabled = healthCheckEnabled;
            HealthCheckHost = healthCheckHost;
            HealthCheckHttpVersion = healthCheckHttpVersion;
            HealthCheckInterval = healthCheckInterval;
            HealthCheckMethod = healthCheckMethod;
            HealthCheckPath = healthCheckPath;
            HealthCheckProtocol = healthCheckProtocol;
            HealthCheckTimeout = healthCheckTimeout;
            HealthyThreshold = healthyThreshold;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
