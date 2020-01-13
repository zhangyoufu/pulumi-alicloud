// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the listeners related to a server load balancer of the current Alibaba Cloud user.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/slb_listeners.html.markdown.
        /// </summary>
        public static Task<GetListenersResult> GetListeners(GetListenersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetListenersResult>("alicloud:slb/getListeners:getListeners", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetListenersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A regex string to filter results by SLB listener description.
        /// </summary>
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        /// <summary>
        /// Filter listeners by the specified frontend port.
        /// </summary>
        [Input("frontendPort")]
        public int? FrontendPort { get; set; }

        /// <summary>
        /// ID of the SLB with listeners.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public string LoadBalancerId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        public GetListenersArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetListenersResult
    {
        public readonly string? DescriptionRegex;
        /// <summary>
        /// Frontend port used to receive incoming traffic and distribute it to the backend servers.
        /// </summary>
        public readonly int? FrontendPort;
        public readonly string LoadBalancerId;
        public readonly string? OutputFile;
        /// <summary>
        /// Listener protocol. Possible values: `http`, `https`, `tcp` and `udp`.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// A list of SLB listeners. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersSlbListenersResult> SlbListeners;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetListenersResult(
            string? descriptionRegex,
            int? frontendPort,
            string loadBalancerId,
            string? outputFile,
            string? protocol,
            ImmutableArray<Outputs.GetListenersSlbListenersResult> slbListeners,
            string id)
        {
            DescriptionRegex = descriptionRegex;
            FrontendPort = frontendPort;
            LoadBalancerId = loadBalancerId;
            OutputFile = outputFile;
            Protocol = protocol;
            SlbListeners = slbListeners;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetListenersSlbListenersResult
    {
        /// <summary>
        /// Port opened on the backend server to receive requests.
        /// </summary>
        public readonly int BackendPort;
        /// <summary>
        /// Peak bandwidth. If the value is set to -1, the listener is not limited by bandwidth.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// ID of the CA certificate (only required when two-way authentication is used). Only available when the protocol is `https`.
        /// </summary>
        public readonly string CaCertificateId;
        /// <summary>
        /// Cookie configured by the backend server. Only available when the sticky_session_type is `server`.
        /// </summary>
        public readonly string Cookie;
        /// <summary>
        /// Cookie timeout in seconds. Only available when the sticky_session_type is `insert`.
        /// </summary>
        public readonly int CookieTimeout;
        /// <summary>
        /// The description of slb listener.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether to enable https listener support http2 or not. Valid values are `on` and `off`. Default to `on`.
        /// </summary>
        public readonly string EnableHttp2;
        /// <summary>
        /// Connection timeout in seconds for the Layer 4 TCP listener. Only available when the protocol is `tcp`.
        /// </summary>
        public readonly int EstablishedTimeout;
        /// <summary>
        /// Filter listeners by the specified frontend port.
        /// </summary>
        public readonly int FrontendPort;
        /// <summary>
        /// Indicate whether Gzip compression is enabled or not. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string Gzip;
        /// <summary>
        /// Indicate whether health check is enabled of not. Possible values are `on` and `off`.
        /// </summary>
        public readonly string HealthCheck;
        /// <summary>
        /// Port used for health check.
        /// </summary>
        public readonly int HealthCheckConnectPort;
        /// <summary>
        /// Amount of time in seconds to wait for the response for a health check.
        /// </summary>
        public readonly int HealthCheckConnectTimeout;
        /// <summary>
        /// Domain name used for health check. The SLB sends HTTP head requests to the backend server, the domain is useful when the backend server verifies the host field in the requests. Only available when the protocol is `http`, `https` or `tcp` (in this case health_check_type must be `http`).
        /// </summary>
        public readonly string HealthCheckDomain;
        /// <summary>
        /// HTTP status codes indicating that the health check is normal. It can contain several comma-separated values such as "http_2xx,http_3xx". Only available when the protocol is `http`, `https` or `tcp` (in this case health_check_type must be `http`).
        /// </summary>
        public readonly string HealthCheckHttpCode;
        /// <summary>
        /// Time interval between two consecutive health checks.
        /// </summary>
        public readonly int HealthCheckInterval;
        /// <summary>
        /// Amount of time in seconds to wait for the response from a health check. If an ECS instance sends no response within the specified timeout period, the health check fails. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly int HealthCheckTimeout;
        /// <summary>
        /// Health check method. Possible values are `tcp` and `http`. Only available when the protocol is `tcp`.
        /// </summary>
        public readonly string HealthCheckType;
        /// <summary>
        /// URI used for health check. Only available when the protocol is `http`, `https` or `tcp` (in this case health_check_type must be `http`).
        /// </summary>
        public readonly string HealthCheckUri;
        /// <summary>
        /// Number of consecutive successes of health check performed on the same ECS instance (from failure to success).
        /// </summary>
        public readonly int HealthyThreshold;
        /// <summary>
        /// Timeout of http or https listener established connection idle timeout. Valid value range: [1-60] in seconds. Default to 15.
        /// </summary>
        public readonly int IdleTimeout;
        /// <summary>
        /// ID of the active/standby server group.
        /// </summary>
        public readonly string MasterSlaveServerGroupId;
        /// <summary>
        /// Timeout value of the TCP connection in seconds. If the value is 0, the session persistence function is disabled. Only available when the protocol is `tcp`.
        /// </summary>
        public readonly int PersistenceTimeout;
        /// <summary>
        /// Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Timeout of http or https listener request (which does not get response from backend) timeout. Valid value range: [1-180] in seconds. Default to 60.
        /// </summary>
        public readonly int RequestTimeout;
        /// <summary>
        /// Algorithm used to distribute traffic. Possible values: `wrr` (weighted round robin), `wlc` (weighted least connection) and `rr` (round robin).
        /// </summary>
        public readonly string Scheduler;
        /// <summary>
        /// Security status. Only available when the protocol is `https`.
        /// </summary>
        public readonly string SecurityStatus;
        public readonly string ServerCertificateId;
        /// <summary>
        /// ID of the linked VServer group.
        /// </summary>
        public readonly string ServerGroupId;
        /// <summary>
        /// ID of the server certificate. Only available when the protocol is `https`.
        /// </summary>
        public readonly string SslCertificateId;
        /// <summary>
        /// Listener status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Indicate whether session persistence is enabled or not. If enabled, all session requests from the same client are sent to the same backend server. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string StickySession;
        /// <summary>
        /// Method used to handle the cookie. Possible values are `insert` (cookie added to the response) and `server` (cookie set by the backend server). Only available when the protocol is `http` or `https` and sticky_session is `on`.
        /// </summary>
        public readonly string StickySessionType;
        /// <summary>
        /// Https listener TLS cipher policy. Valid values are `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`. Default to `tls_cipher_policy_1_0`.
        /// </summary>
        public readonly string TlsCipherPolicy;
        /// <summary>
        /// Number of consecutive failures of health check performed on the same ECS instance (from success to failure).
        /// </summary>
        public readonly int UnhealthyThreshold;
        /// <summary>
        /// Indicate whether the HTTP header field "X-Forwarded-For" is added or not; it allows the backend server to know about the user's IP address. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string XForwardedFor;
        /// <summary>
        /// Indicate whether the HTTP header field "X-Forwarded-For_SLBID" is added or not; it allows the backend server to know about the SLB ID. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string XForwardedForSlbId;
        /// <summary>
        /// Indicate whether the HTTP header field "X-Forwarded-For_SLBIP" is added or not; it allows the backend server to know about the SLB IP address. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string XForwardedForSlbIp;
        /// <summary>
        /// Indicate whether the HTTP header field "X-Forwarded-For_proto" is added or not; it allows the backend server to know about the user's protocol. Possible values are `on` and `off`. Only available when the protocol is `http` or `https`.
        /// </summary>
        public readonly string XForwardedForSlbProto;

        [OutputConstructor]
        private GetListenersSlbListenersResult(
            int backendPort,
            int bandwidth,
            string caCertificateId,
            string cookie,
            int cookieTimeout,
            string description,
            string enableHttp2,
            int establishedTimeout,
            int frontendPort,
            string gzip,
            string healthCheck,
            int healthCheckConnectPort,
            int healthCheckConnectTimeout,
            string healthCheckDomain,
            string healthCheckHttpCode,
            int healthCheckInterval,
            int healthCheckTimeout,
            string healthCheckType,
            string healthCheckUri,
            int healthyThreshold,
            int idleTimeout,
            string masterSlaveServerGroupId,
            int persistenceTimeout,
            string protocol,
            int requestTimeout,
            string scheduler,
            string securityStatus,
            string serverCertificateId,
            string serverGroupId,
            string sslCertificateId,
            string status,
            string stickySession,
            string stickySessionType,
            string tlsCipherPolicy,
            int unhealthyThreshold,
            string xForwardedFor,
            string xForwardedForSlbId,
            string xForwardedForSlbIp,
            string xForwardedForSlbProto)
        {
            BackendPort = backendPort;
            Bandwidth = bandwidth;
            CaCertificateId = caCertificateId;
            Cookie = cookie;
            CookieTimeout = cookieTimeout;
            Description = description;
            EnableHttp2 = enableHttp2;
            EstablishedTimeout = establishedTimeout;
            FrontendPort = frontendPort;
            Gzip = gzip;
            HealthCheck = healthCheck;
            HealthCheckConnectPort = healthCheckConnectPort;
            HealthCheckConnectTimeout = healthCheckConnectTimeout;
            HealthCheckDomain = healthCheckDomain;
            HealthCheckHttpCode = healthCheckHttpCode;
            HealthCheckInterval = healthCheckInterval;
            HealthCheckTimeout = healthCheckTimeout;
            HealthCheckType = healthCheckType;
            HealthCheckUri = healthCheckUri;
            HealthyThreshold = healthyThreshold;
            IdleTimeout = idleTimeout;
            MasterSlaveServerGroupId = masterSlaveServerGroupId;
            PersistenceTimeout = persistenceTimeout;
            Protocol = protocol;
            RequestTimeout = requestTimeout;
            Scheduler = scheduler;
            SecurityStatus = securityStatus;
            ServerCertificateId = serverCertificateId;
            ServerGroupId = serverGroupId;
            SslCertificateId = sslCertificateId;
            Status = status;
            StickySession = stickySession;
            StickySessionType = stickySessionType;
            TlsCipherPolicy = tlsCipherPolicy;
            UnhealthyThreshold = unhealthyThreshold;
            XForwardedFor = xForwardedFor;
            XForwardedForSlbId = xForwardedForSlbId;
            XForwardedForSlbIp = xForwardedForSlbIp;
            XForwardedForSlbProto = xForwardedForSlbProto;
        }
    }
    }
}
