// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    /// <summary>
    /// Provides a Application Load Balancer (ALB) Listener resource.
    /// 
    /// For information about Application Load Balancer (ALB) Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/slb/application-load-balancer/developer-reference/api-alb-2020-06-16-createlistener).
    /// 
    /// &gt; **NOTE:** Available since v1.133.0.
    /// 
    /// ## Import
    /// 
    /// Application Load Balancer (ALB) Listener can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:alb/listener:Listener example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alb/listener:Listener")]
    public partial class Listener : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
        /// 
        /// &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
        /// </summary>
        [Output("accessLogRecordCustomizedHeadersEnabled")]
        public Output<bool> AccessLogRecordCustomizedHeadersEnabled { get; private set; } = null!;

        /// <summary>
        /// Xtrace Configuration Information. See `access_log_tracing_config` below for details.
        /// </summary>
        [Output("accessLogTracingConfig")]
        public Output<Outputs.ListenerAccessLogTracingConfig?> AccessLogTracingConfig { get; private set; } = null!;

        /// <summary>
        /// The configurations of the access control lists (ACLs). See `acl_config` below for details. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
        /// </summary>
        [Output("aclConfig")]
        public Output<Outputs.ListenerAclConfig> AclConfig { get; private set; } = null!;

        /// <summary>
        /// The default certificate of the Listener. See `certificates` below for details. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
        /// </summary>
        [Output("certificates")]
        public Output<Outputs.ListenerCertificates?> Certificates { get; private set; } = null!;

        /// <summary>
        /// The Default Rule Action List. See `default_actions` below for details.
        /// </summary>
        [Output("defaultActions")]
        public Output<ImmutableArray<Outputs.ListenerDefaultAction>> DefaultActions { get; private set; } = null!;

        /// <summary>
        /// The dry run.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
        /// </summary>
        [Output("gzipEnabled")]
        public Output<bool> GzipEnabled { get; private set; } = null!;

        /// <summary>
        /// Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Output("http2Enabled")]
        public Output<bool> Http2Enabled { get; private set; } = null!;

        /// <summary>
        /// Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
        /// </summary>
        [Output("listenerDescription")]
        public Output<string?> ListenerDescription { get; private set; } = null!;

        /// <summary>
        /// The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
        /// </summary>
        [Output("listenerPort")]
        public Output<int> ListenerPort { get; private set; } = null!;

        /// <summary>
        /// Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
        /// </summary>
        [Output("listenerProtocol")]
        public Output<string> ListenerProtocol { get; private set; } = null!;

        /// <summary>
        /// The ALB Instance Id.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Configuration Associated with the QuIC Listening. See `quic_config` below for details.
        /// </summary>
        [Output("quicConfig")]
        public Output<Outputs.ListenerQuicConfig> QuicConfig { get; private set; } = null!;

        /// <summary>
        /// The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
        /// </summary>
        [Output("requestTimeout")]
        public Output<int> RequestTimeout { get; private set; } = null!;

        /// <summary>
        /// Security Policy.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The `x_forward_for` Related Attribute Configuration. See `x_forwarded_for_config` below for details. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Output("xForwardedForConfig")]
        public Output<Outputs.ListenerXForwardedForConfig> XForwardedForConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alb/listener:Listener", name, args ?? new ListenerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alb/listener:Listener", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, state, options);
        }
    }

    public sealed class ListenerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
        /// 
        /// &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
        /// </summary>
        [Input("accessLogRecordCustomizedHeadersEnabled")]
        public Input<bool>? AccessLogRecordCustomizedHeadersEnabled { get; set; }

        /// <summary>
        /// Xtrace Configuration Information. See `access_log_tracing_config` below for details.
        /// </summary>
        [Input("accessLogTracingConfig")]
        public Input<Inputs.ListenerAccessLogTracingConfigArgs>? AccessLogTracingConfig { get; set; }

        /// <summary>
        /// The configurations of the access control lists (ACLs). See `acl_config` below for details. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
        /// </summary>
        [Input("aclConfig")]
        public Input<Inputs.ListenerAclConfigArgs>? AclConfig { get; set; }

        /// <summary>
        /// The default certificate of the Listener. See `certificates` below for details. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
        /// </summary>
        [Input("certificates")]
        public Input<Inputs.ListenerCertificatesArgs>? Certificates { get; set; }

        [Input("defaultActions")]
        private InputList<Inputs.ListenerDefaultActionArgs>? _defaultActions;

        /// <summary>
        /// The Default Rule Action List. See `default_actions` below for details.
        /// </summary>
        public InputList<Inputs.ListenerDefaultActionArgs> DefaultActions
        {
            get => _defaultActions ?? (_defaultActions = new InputList<Inputs.ListenerDefaultActionArgs>());
            set => _defaultActions = value;
        }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
        /// </summary>
        [Input("gzipEnabled")]
        public Input<bool>? GzipEnabled { get; set; }

        /// <summary>
        /// Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("http2Enabled")]
        public Input<bool>? Http2Enabled { get; set; }

        /// <summary>
        /// Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
        /// </summary>
        [Input("listenerDescription")]
        public Input<string>? ListenerDescription { get; set; }

        /// <summary>
        /// The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
        /// </summary>
        [Input("listenerPort", required: true)]
        public Input<int> ListenerPort { get; set; } = null!;

        /// <summary>
        /// Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
        /// </summary>
        [Input("listenerProtocol", required: true)]
        public Input<string> ListenerProtocol { get; set; } = null!;

        /// <summary>
        /// The ALB Instance Id.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// Configuration Associated with the QuIC Listening. See `quic_config` below for details.
        /// </summary>
        [Input("quicConfig")]
        public Input<Inputs.ListenerQuicConfigArgs>? QuicConfig { get; set; }

        /// <summary>
        /// The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// Security Policy.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The `x_forward_for` Related Attribute Configuration. See `x_forwarded_for_config` below for details. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("xForwardedForConfig")]
        public Input<Inputs.ListenerXForwardedForConfigArgs>? XForwardedForConfig { get; set; }

        public ListenerArgs()
        {
        }
        public static new ListenerArgs Empty => new ListenerArgs();
    }

    public sealed class ListenerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
        /// 
        /// &gt; **NOTE:** Only Instances outside the Security Group to Access the Log Switch **accesslogenabled** Open, in Order to Set This Parameter to the **True**.
        /// </summary>
        [Input("accessLogRecordCustomizedHeadersEnabled")]
        public Input<bool>? AccessLogRecordCustomizedHeadersEnabled { get; set; }

        /// <summary>
        /// Xtrace Configuration Information. See `access_log_tracing_config` below for details.
        /// </summary>
        [Input("accessLogTracingConfig")]
        public Input<Inputs.ListenerAccessLogTracingConfigGetArgs>? AccessLogTracingConfig { get; set; }

        /// <summary>
        /// The configurations of the access control lists (ACLs). See `acl_config` below for details. **NOTE:** Field `acl_config` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
        /// </summary>
        [Input("aclConfig")]
        public Input<Inputs.ListenerAclConfigGetArgs>? AclConfig { get; set; }

        /// <summary>
        /// The default certificate of the Listener. See `certificates` below for details. **NOTE:** When `listener_protocol` is `HTTPS`, The default certificate must be set one。
        /// </summary>
        [Input("certificates")]
        public Input<Inputs.ListenerCertificatesGetArgs>? Certificates { get; set; }

        [Input("defaultActions")]
        private InputList<Inputs.ListenerDefaultActionGetArgs>? _defaultActions;

        /// <summary>
        /// The Default Rule Action List. See `default_actions` below for details.
        /// </summary>
        public InputList<Inputs.ListenerDefaultActionGetArgs> DefaultActions
        {
            get => _defaultActions ?? (_defaultActions = new InputList<Inputs.ListenerDefaultActionGetArgs>());
            set => _defaultActions = value;
        }

        /// <summary>
        /// The dry run.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
        /// </summary>
        [Input("gzipEnabled")]
        public Input<bool>? GzipEnabled { get; set; }

        /// <summary>
        /// Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("http2Enabled")]
        public Input<bool>? Http2Enabled { get; set; }

        /// <summary>
        /// Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
        /// </summary>
        [Input("listenerDescription")]
        public Input<string>? ListenerDescription { get; set; }

        /// <summary>
        /// The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
        /// </summary>
        [Input("listenerPort")]
        public Input<int>? ListenerPort { get; set; }

        /// <summary>
        /// Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
        /// </summary>
        [Input("listenerProtocol")]
        public Input<string>? ListenerProtocol { get; set; }

        /// <summary>
        /// The ALB Instance Id.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Configuration Associated with the QuIC Listening. See `quic_config` below for details.
        /// </summary>
        [Input("quicConfig")]
        public Input<Inputs.ListenerQuicConfigGetArgs>? QuicConfig { get; set; }

        /// <summary>
        /// The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// Security Policy.
        /// 
        /// &gt; **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The `x_forward_for` Related Attribute Configuration. See `x_forwarded_for_config` below for details. **NOTE:** The attribute is valid when the attribute `listener_protocol` is `HTTPS`.
        /// </summary>
        [Input("xForwardedForConfig")]
        public Input<Inputs.ListenerXForwardedForConfigGetArgs>? XForwardedForConfig { get; set; }

        public ListenerState()
        {
        }
        public static new ListenerState Empty => new ListenerState();
    }
}
