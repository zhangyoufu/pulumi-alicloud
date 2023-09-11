// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Wafv3.Inputs
{

    public sealed class DomainRedirectArgs : global::Pulumi.ResourceArgs
    {
        [Input("backends")]
        private InputList<string>? _backends;

        /// <summary>
        /// The IP address of the origin server corresponding to the domain name or the back-to-origin domain name of the server.
        /// </summary>
        public InputList<string> Backends
        {
            get => _backends ?? (_backends = new InputList<string>());
            set => _backends = value;
        }

        /// <summary>
        /// Connection timeout. Unit: seconds, value range: 5~120.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// Whether to enable forced HTTP back-to-origin. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value:
        /// - **true**: indicates that forced HTTP back-to-origin is enabled.
        /// - **false**: indicates that forced HTTP back-to-origin is not enabled.
        /// </summary>
        [Input("focusHttpBackend")]
        public Input<bool>? FocusHttpBackend { get; set; }

        /// <summary>
        /// Open long connection, default true.
        /// </summary>
        [Input("keepalive")]
        public Input<bool>? Keepalive { get; set; }

        /// <summary>
        /// Number of long connections,  default: `60`. range :60-1000.
        /// </summary>
        [Input("keepaliveRequests")]
        public Input<int>? KeepaliveRequests { get; set; }

        /// <summary>
        /// Long connection over time, default: `15`. Range: 1-60.
        /// </summary>
        [Input("keepaliveTimeout")]
        public Input<int>? KeepaliveTimeout { get; set; }

        /// <summary>
        /// The load balancing algorithm used when returning to the source. Value:
        /// - **iphash**: indicates the IPHash algorithm.
        /// - **roundRobin**: indicates the polling algorithm.
        /// - **leastTime**: indicates the Least Time algorithm.
        /// - This value can be selected only if the value of **protection_resource** is **gslb** (indicating that the protected resource type uses shared cluster intelligent load balancing).
        /// </summary>
        [Input("loadbalance", required: true)]
        public Input<string> Loadbalance { get; set; } = null!;

        /// <summary>
        /// Read timeout duration. **Unit**: seconds, **Value range**: 5~1800.
        /// </summary>
        [Input("readTimeout")]
        public Input<int>? ReadTimeout { get; set; }

        [Input("requestHeaders")]
        private InputList<Inputs.DomainRedirectRequestHeaderArgs>? _requestHeaders;

        /// <summary>
        /// The traffic tag field and value of the domain name which used to mark the traffic processed by WAF. 
        /// It formats as `[{" k ":"_key_"," v ":"_value_"}]`. Where the `k` represents the specified custom request header field,
        /// and the `v` represents the value set for this field. By specifying the custom request header field and the corresponding value,
        /// when the access traffic of the domain name passes through WAF, WAF automatically adds the specified custom field value
        /// to the request header as the traffic mark, which is convenient for backend service statistics.Explain that if the
        /// custom header field already exists in the request, the system will overwrite the value of the custom field in the
        /// request with the set traffic tag value. See `request_headers` below.
        /// </summary>
        public InputList<Inputs.DomainRedirectRequestHeaderArgs> RequestHeaders
        {
            get => _requestHeaders ?? (_requestHeaders = new InputList<Inputs.DomainRedirectRequestHeaderArgs>());
            set => _requestHeaders = value;
        }

        /// <summary>
        /// Back to Source Retry. default: true, retry 3 times by default.
        /// </summary>
        [Input("retry")]
        public Input<bool>? Retry { get; set; }

        /// <summary>
        /// Whether to enable back-to-source SNI. This parameter is used only if the value of **https_ports** is not empty (indicating that the domain name uses the HTTPS protocol). Value:
        /// - **true**: indicates that the back-to-source SNI is enabled.
        /// - **false** (default) indicates that the back-to-source SNI is not enabled.
        /// </summary>
        [Input("sniEnabled")]
        public Input<bool>? SniEnabled { get; set; }

        /// <summary>
        /// Sets the value of the custom SNI extension field. If this parameter is not set, the value of the **Host** field in the request header is used as the value of the SNI extension field by default.In general, you do not need to customize SNI unless your business has special configuration requirements. You want WAF to use SNI that is inconsistent with the actual request Host in the back-to-origin request (that is, the custom SNI set here).&gt; This parameter is required only when **sni_enalbed** is set to **true** (indicating that back-to-source SNI is enabled).
        /// </summary>
        [Input("sniHost")]
        public Input<string>? SniHost { get; set; }

        /// <summary>
        /// Write timeout duration&gt; **Unit**: seconds, **Value range**: 5~1800.
        /// </summary>
        [Input("writeTimeout")]
        public Input<int>? WriteTimeout { get; set; }

        public DomainRedirectArgs()
        {
        }
        public static new DomainRedirectArgs Empty => new DomainRedirectArgs();
    }
}
