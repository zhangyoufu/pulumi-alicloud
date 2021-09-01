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
    public sealed class GetListenersListenerResult
    {
        /// <summary>
        /// Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
        /// </summary>
        public readonly bool AccessLogRecordCustomizedHeadersEnabled;
        /// <summary>
        /// Xtrace Configuration Information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerAccessLogTracingConfigResult> AccessLogTracingConfigs;
        /// <summary>
        /// Snooping Binding of the Access Policy Group ID List.
        /// </summary>
        public readonly string AclId;
        /// <summary>
        /// The type of the ACL. Valid values: White: specifies the ACL as a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. Whitelists apply to scenarios where only specific IP addresses are allowed to access an application. Risks may occur if the whitelist is improperly set. After you set a whitelist for an Application Load Balancer (ALB) listener, only requests from IP addresses that are added to the whitelist are distributed by the listener. If the whitelist is enabled without IP addresses specified, the ALB listener does not forward requests. Black: All requests from the IP addresses or CIDR blocks in the ACL are denied. The blacklist is used to prevent specified IP addresses from accessing an application. If the blacklist is enabled but the corresponding ACL does not contain IP addresses, the ALB listener forwards all requests.
        /// </summary>
        public readonly string AclType;
        /// <summary>
        /// Certificate.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerCertificateResult> Certificates;
        /// <summary>
        /// The Default Rule Action List.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerDefaultActionResult> DefaultActions;
        /// <summary>
        /// Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid Values: `True` Or `False`. Default Value: `True`.
        /// </summary>
        public readonly bool GzipEnabled;
        /// <summary>
        /// Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
        /// </summary>
        public readonly bool Http2Enabled;
        /// <summary>
        /// The ID of the Listener.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specify the Connection Idle Timeout Value: 1 to 60.
        /// </summary>
        public readonly int IdleTimeout;
        /// <summary>
        /// Set the IP Address of the Listened Description. Length Is from 2 to 256 Characters. 	* `listener_id` - on Behalf of the Resource Level Id of the Resources Property Fields.
        /// </summary>
        public readonly string ListenerDescription;
        public readonly string ListenerId;
        /// <summary>
        /// The SLB Instance Front-End, and Those of the Ports Used. Value: `1~65535`.
        /// </summary>
        public readonly int ListenerPort;
        /// <summary>
        /// Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
        /// </summary>
        public readonly string ListenerProtocol;
        /// <summary>
        /// The SLB Instance Id.
        /// </summary>
        public readonly string LoadBalancerId;
        /// <summary>
        /// This Request Returned by the Maximum Number of Records.
        /// </summary>
        public readonly string MaxResults;
        /// <summary>
        /// The Current Call Returns to the Position of the Set to Null Represents the Data Has Been Read to the End of.
        /// </summary>
        public readonly string NextToken;
        /// <summary>
        /// Configuration Associated with the QuIC Listening.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerQuicConfigResult> QuicConfigs;
        /// <summary>
        /// The Specified Request Timeout Time. Value: 1~180 Seconds. Default Value: 60. If the Timeout Time Within the Back-End Server Has Not Answered the SLB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
        /// </summary>
        public readonly int RequestTimeout;
        /// <summary>
        /// Security Policy.
        /// </summary>
        public readonly string SecurityPolicyId;
        /// <summary>
        /// The state of the listener. Valid Values: `Running` Or `Stopped`. `Running`: The listener is running. `Stopped`: The listener is stopped.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// xforwardfor Related Attribute Configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenersListenerXforwardedForConfigResult> XforwardedForConfigs;

        [OutputConstructor]
        private GetListenersListenerResult(
            bool accessLogRecordCustomizedHeadersEnabled,

            ImmutableArray<Outputs.GetListenersListenerAccessLogTracingConfigResult> accessLogTracingConfigs,

            string aclId,

            string aclType,

            ImmutableArray<Outputs.GetListenersListenerCertificateResult> certificates,

            ImmutableArray<Outputs.GetListenersListenerDefaultActionResult> defaultActions,

            bool gzipEnabled,

            bool http2Enabled,

            string id,

            int idleTimeout,

            string listenerDescription,

            string listenerId,

            int listenerPort,

            string listenerProtocol,

            string loadBalancerId,

            string maxResults,

            string nextToken,

            ImmutableArray<Outputs.GetListenersListenerQuicConfigResult> quicConfigs,

            int requestTimeout,

            string securityPolicyId,

            string status,

            ImmutableArray<Outputs.GetListenersListenerXforwardedForConfigResult> xforwardedForConfigs)
        {
            AccessLogRecordCustomizedHeadersEnabled = accessLogRecordCustomizedHeadersEnabled;
            AccessLogTracingConfigs = accessLogTracingConfigs;
            AclId = aclId;
            AclType = aclType;
            Certificates = certificates;
            DefaultActions = defaultActions;
            GzipEnabled = gzipEnabled;
            Http2Enabled = http2Enabled;
            Id = id;
            IdleTimeout = idleTimeout;
            ListenerDescription = listenerDescription;
            ListenerId = listenerId;
            ListenerPort = listenerPort;
            ListenerProtocol = listenerProtocol;
            LoadBalancerId = loadBalancerId;
            MaxResults = maxResults;
            NextToken = nextToken;
            QuicConfigs = quicConfigs;
            RequestTimeout = requestTimeout;
            SecurityPolicyId = securityPolicyId;
            Status = status;
            XforwardedForConfigs = xforwardedForConfigs;
        }
    }
}
