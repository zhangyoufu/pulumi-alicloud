// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Slb
{
    /// <summary>
    /// A forwarding rule is configured in `HTTP`/`HTTPS` listener and it used to listen a list of backend servers which in one specified virtual backend server group.
    /// You can add forwarding rules to a listener to forward requests based on the domain names or the URL in the request.
    /// 
    /// &gt; **NOTE:** One virtual backend server group can be attached in multiple forwarding rules.
    /// 
    /// &gt; **NOTE:** At least one "Domain" or "Url" must be specified when creating a new rule.
    /// 
    /// &gt; **NOTE:** Having the same 'Domain' and 'Url' rule can not be created repeatedly in the one listener.
    /// 
    /// &gt; **NOTE:** Rule only be created in the `HTTP` or `HTTPS` listener.
    /// 
    /// &gt; **NOTE:** Only rule's virtual server group can be modified.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var slbRuleName = config.Get("slbRuleName") ?? "terraform-example";
    ///     var ruleZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableDiskCategory = "cloud_efficiency",
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var ruleInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = ruleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 1,
    ///         MemorySize = 2,
    ///     });
    /// 
    ///     var ruleImages = AliCloud.Ecs.GetImages.Invoke(new()
    ///     {
    ///         NameRegex = "^ubuntu_18.*64",
    ///         MostRecent = true,
    ///         Owners = "system",
    ///     });
    /// 
    ///     var ruleNetwork = new AliCloud.Vpc.Network("ruleNetwork", new()
    ///     {
    ///         VpcName = slbRuleName,
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var ruleSwitch = new AliCloud.Vpc.Switch("ruleSwitch", new()
    ///     {
    ///         VpcId = ruleNetwork.Id,
    ///         CidrBlock = "172.16.0.0/16",
    ///         ZoneId = ruleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = slbRuleName,
    ///     });
    /// 
    ///     var ruleSecurityGroup = new AliCloud.Ecs.SecurityGroup("ruleSecurityGroup", new()
    ///     {
    ///         VpcId = ruleNetwork.Id,
    ///     });
    /// 
    ///     var ruleInstance = new AliCloud.Ecs.Instance("ruleInstance", new()
    ///     {
    ///         ImageId = ruleImages.Apply(getImagesResult =&gt; getImagesResult.Images[0]?.Id),
    ///         InstanceType = ruleInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///         SecurityGroups = new[]
    ///         {
    ///             ruleSecurityGroup,
    ///         }.Select(__item =&gt; __item.Id).ToList(),
    ///         InternetChargeType = "PayByTraffic",
    ///         InternetMaxBandwidthOut = 10,
    ///         AvailabilityZone = ruleZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         InstanceChargeType = "PostPaid",
    ///         SystemDiskCategory = "cloud_efficiency",
    ///         VswitchId = ruleSwitch.Id,
    ///         InstanceName = slbRuleName,
    ///     });
    /// 
    ///     var ruleApplicationLoadBalancer = new AliCloud.Slb.ApplicationLoadBalancer("ruleApplicationLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = slbRuleName,
    ///         VswitchId = ruleSwitch.Id,
    ///         InstanceChargeType = "PayByCLCU",
    ///     });
    /// 
    ///     var ruleListener = new AliCloud.Slb.Listener("ruleListener", new()
    ///     {
    ///         LoadBalancerId = ruleApplicationLoadBalancer.Id,
    ///         BackendPort = 22,
    ///         FrontendPort = 22,
    ///         Protocol = "http",
    ///         Bandwidth = 5,
    ///         HealthCheckConnectPort = 20,
    ///     });
    /// 
    ///     var ruleServerGroup = new AliCloud.Slb.ServerGroup("ruleServerGroup", new()
    ///     {
    ///         LoadBalancerId = ruleApplicationLoadBalancer.Id,
    ///     });
    /// 
    ///     var ruleRule = new AliCloud.Slb.Rule("ruleRule", new()
    ///     {
    ///         LoadBalancerId = ruleApplicationLoadBalancer.Id,
    ///         FrontendPort = ruleListener.FrontendPort,
    ///         Domain = "*.aliyun.com",
    ///         Url = "/image",
    ///         ServerGroupId = ruleServerGroup.Id,
    ///         Cookie = "23ffsa",
    ///         CookieTimeout = 100,
    ///         HealthCheckHttpCode = "http_2xx",
    ///         HealthCheckInterval = 10,
    ///         HealthCheckUri = "/test",
    ///         HealthCheckConnectPort = 80,
    ///         HealthCheckTimeout = 30,
    ///         HealthyThreshold = 3,
    ///         UnhealthyThreshold = 5,
    ///         StickySession = "on",
    ///         StickySessionType = "server",
    ///         ListenerSync = "off",
    ///         Scheduler = "rr",
    ///         HealthCheckDomain = "test",
    ///         HealthCheck = "on",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Load balancer forwarding rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:slb/rule:Rule example rule-abc123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:slb/rule:Rule")]
    public partial class Rule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The cookie configured on the server. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
        /// </summary>
        [Output("cookie")]
        public Output<string?> Cookie { get; private set; } = null!;

        /// <summary>
        /// Cookie timeout. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
        /// </summary>
        [Output("cookieTimeout")]
        public Output<int?> CookieTimeout { get; private set; } = null!;

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Output("deleteProtectionValidation")]
        public Output<bool?> DeleteProtectionValidation { get; private set; } = null!;

        /// <summary>
        /// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
        /// and wildcard characters. The following two domain name formats are supported:
        /// - Standard domain name: www.test.com
        /// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
        /// </summary>
        [Output("frontendPort")]
        public Output<int> FrontendPort { get; private set; } = null!;

        /// <summary>
        /// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Output("healthCheck")]
        public Output<string?> HealthCheck { get; private set; } = null!;

        /// <summary>
        /// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
        /// </summary>
        [Output("healthCheckConnectPort")]
        public Output<int> HealthCheckConnectPort { get; private set; } = null!;

        /// <summary>
        /// Domain name used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
        /// </summary>
        [Output("healthCheckDomain")]
        public Output<string?> HealthCheckDomain { get; private set; } = null!;

        /// <summary>
        /// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `health_check` is on. Default to `http_2xx`.  Valid values are: `http_2xx`,  `http_3xx`, `http_4xx` and `http_5xx`.
        /// </summary>
        [Output("healthCheckHttpCode")]
        public Output<string?> HealthCheckHttpCode { get; private set; } = null!;

        /// <summary>
        /// Time interval of health checks. It is required when `health_check` is on. Valid value range: [1-50] in seconds. Default to 2.
        /// </summary>
        [Output("healthCheckInterval")]
        public Output<int?> HealthCheckInterval { get; private set; } = null!;

        /// <summary>
        /// Maximum timeout of each health check response. It is required when `health_check` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `health_check_timeout` &lt; `health_check_interval`, its will be replaced by `health_check_interval`.
        /// </summary>
        [Output("healthCheckTimeout")]
        public Output<int?> HealthCheckTimeout { get; private set; } = null!;

        /// <summary>
        /// URI used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%!’(MISSING), ‘?’, #’ and ‘&amp;’ are allowed.
        /// </summary>
        [Output("healthCheckUri")]
        public Output<string?> HealthCheckUri { get; private set; } = null!;

        /// <summary>
        /// Threshold determining the result of the health check is success. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Output("healthyThreshold")]
        public Output<int?> HealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
        /// </summary>
        [Output("listenerSync")]
        public Output<string?> ListenerSync { get; private set; } = null!;

        /// <summary>
        /// The Load Balancer ID which is used to launch the new forwarding rule.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Output("scheduler")]
        public Output<string?> Scheduler { get; private set; } = null!;

        /// <summary>
        /// ID of a virtual server group that will be forwarded.
        /// </summary>
        [Output("serverGroupId")]
        public Output<string> ServerGroupId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Output("stickySession")]
        public Output<string?> StickySession { get; private set; } = null!;

        /// <summary>
        /// Mode for handling the cookie. If `sticky_session` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
        /// </summary>
        [Output("stickySessionType")]
        public Output<string?> StickySessionType { get; private set; } = null!;

        /// <summary>
        /// Threshold determining the result of the health check is fail. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Output("unhealthyThreshold")]
        public Output<int?> UnhealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
        /// and characters '-' '/' '?' '%!'(MISSING) '#' and '&amp;' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("alicloud:slb/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:slb/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cookie configured on the server. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
        /// </summary>
        [Input("cookie")]
        public Input<string>? Cookie { get; set; }

        /// <summary>
        /// Cookie timeout. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
        /// </summary>
        [Input("cookieTimeout")]
        public Input<int>? CookieTimeout { get; set; }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
        /// and wildcard characters. The following two domain name formats are supported:
        /// - Standard domain name: www.test.com
        /// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
        /// </summary>
        [Input("frontendPort", required: true)]
        public Input<int> FrontendPort { get; set; } = null!;

        /// <summary>
        /// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
        /// </summary>
        [Input("healthCheckConnectPort")]
        public Input<int>? HealthCheckConnectPort { get; set; }

        /// <summary>
        /// Domain name used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
        /// </summary>
        [Input("healthCheckDomain")]
        public Input<string>? HealthCheckDomain { get; set; }

        /// <summary>
        /// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `health_check` is on. Default to `http_2xx`.  Valid values are: `http_2xx`,  `http_3xx`, `http_4xx` and `http_5xx`.
        /// </summary>
        [Input("healthCheckHttpCode")]
        public Input<string>? HealthCheckHttpCode { get; set; }

        /// <summary>
        /// Time interval of health checks. It is required when `health_check` is on. Valid value range: [1-50] in seconds. Default to 2.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// Maximum timeout of each health check response. It is required when `health_check` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `health_check_timeout` &lt; `health_check_interval`, its will be replaced by `health_check_interval`.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<int>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// URI used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%!’(MISSING), ‘?’, #’ and ‘&amp;’ are allowed.
        /// </summary>
        [Input("healthCheckUri")]
        public Input<string>? HealthCheckUri { get; set; }

        /// <summary>
        /// Threshold determining the result of the health check is success. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
        /// </summary>
        [Input("listenerSync")]
        public Input<string>? ListenerSync { get; set; }

        /// <summary>
        /// The Load Balancer ID which is used to launch the new forwarding rule.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// ID of a virtual server group that will be forwarded.
        /// </summary>
        [Input("serverGroupId", required: true)]
        public Input<string> ServerGroupId { get; set; } = null!;

        /// <summary>
        /// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("stickySession")]
        public Input<string>? StickySession { get; set; }

        /// <summary>
        /// Mode for handling the cookie. If `sticky_session` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
        /// </summary>
        [Input("stickySessionType")]
        public Input<string>? StickySessionType { get; set; }

        /// <summary>
        /// Threshold determining the result of the health check is fail. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        /// <summary>
        /// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
        /// and characters '-' '/' '?' '%!'(MISSING) '#' and '&amp;' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }

    public sealed class RuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cookie configured on the server. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
        /// </summary>
        [Input("cookie")]
        public Input<string>? Cookie { get; set; }

        /// <summary>
        /// Cookie timeout. It is mandatory when `sticky_session` is "on" and `sticky_session_type` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.
        /// </summary>
        [Input("cookieTimeout")]
        public Input<int>? CookieTimeout { get; set; }

        /// <summary>
        /// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
        /// </summary>
        [Input("deleteProtectionValidation")]
        public Input<bool>? DeleteProtectionValidation { get; set; }

        /// <summary>
        /// Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
        /// and wildcard characters. The following two domain name formats are supported:
        /// - Standard domain name: www.test.com
        /// - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].
        /// </summary>
        [Input("frontendPort")]
        public Input<int>? FrontendPort { get; set; }

        /// <summary>
        /// Whether to enable health check. Valid values are`on` and `off`. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.
        /// </summary>
        [Input("healthCheckConnectPort")]
        public Input<int>? HealthCheckConnectPort { get; set; }

        /// <summary>
        /// Domain name used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty,  Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
        /// </summary>
        [Input("healthCheckDomain")]
        public Input<string>? HealthCheckDomain { get; set; }

        /// <summary>
        /// Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `health_check` is on. Default to `http_2xx`.  Valid values are: `http_2xx`,  `http_3xx`, `http_4xx` and `http_5xx`.
        /// </summary>
        [Input("healthCheckHttpCode")]
        public Input<string>? HealthCheckHttpCode { get; set; }

        /// <summary>
        /// Time interval of health checks. It is required when `health_check` is on. Valid value range: [1-50] in seconds. Default to 2.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// Maximum timeout of each health check response. It is required when `health_check` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If `health_check_timeout` &lt; `health_check_interval`, its will be replaced by `health_check_interval`.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<int>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// URI used for health check. When it used to launch TCP listener, `health_check_type` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%!’(MISSING), ‘?’, #’ and ‘&amp;’ are allowed.
        /// </summary>
        [Input("healthCheckUri")]
        public Input<string>? HealthCheckUri { get; set; }

        /// <summary>
        /// Threshold determining the result of the health check is success. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.
        /// </summary>
        [Input("listenerSync")]
        public Input<string>? ListenerSync { get; set; }

        /// <summary>
        /// The Load Balancer ID which is used to launch the new forwarding rule.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Scheduling algorithm, Valid values are `wrr`, `rr` and `wlc`.  Default to "wrr". This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// ID of a virtual server group that will be forwarded.
        /// </summary>
        [Input("serverGroupId")]
        public Input<string>? ServerGroupId { get; set; }

        /// <summary>
        /// Whether to enable session persistence, Valid values are `on` and `off`. Default to `off`. This parameter is required  and takes effect only when ListenerSync is set to off.
        /// </summary>
        [Input("stickySession")]
        public Input<string>? StickySession { get; set; }

        /// <summary>
        /// Mode for handling the cookie. If `sticky_session` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
        /// </summary>
        [Input("stickySessionType")]
        public Input<string>? StickySessionType { get; set; }

        /// <summary>
        /// Threshold determining the result of the health check is fail. It is required when `health_check` is on. Valid value range: [1-10] in seconds. Default to 3.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        /// <summary>
        /// Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9,
        /// and characters '-' '/' '?' '%!'(MISSING) '#' and '&amp;' are allowed. URLs must be started with the character '/', but cannot be '/' alone.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RuleState()
        {
        }
        public static new RuleState Empty => new RuleState();
    }
}
