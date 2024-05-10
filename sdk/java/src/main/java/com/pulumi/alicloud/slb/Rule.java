// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.RuleArgs;
import com.pulumi.alicloud.slb.inputs.RuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Lindorm Instance resource.
 * 
 * For information about Load Balancer Forwarding Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/slb/classic-load-balancer/developer-reference/api-slb-2014-05-15-dir-forwarding-rules).
 * 
 * &gt; **NOTE:** Available since v1.6.0.
 * 
 * A forwarding rule is configured in `HTTP`/`HTTPS` listener and it used to listen a list of backend servers which in one specified virtual backend server group.
 * You can add forwarding rules to a listener to forward requests based on the domain names or the URL in the request.
 * 
 * &gt; **NOTE:** One virtual backend server group can be attached in multiple forwarding rules.
 * 
 * &gt; **NOTE:** At least one &#34;Domain&#34; or &#34;Url&#34; must be specified when creating a new rule.
 * 
 * &gt; **NOTE:** Having the same &#39;Domain&#39; and &#39;Url&#39; rule can not be created repeatedly in the one listener.
 * 
 * &gt; **NOTE:** Rule only be created in the `HTTP` or `HTTPS` listener.
 * 
 * &gt; **NOTE:** Only rule&#39;s virtual server group can be modified.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.slb.Listener;
 * import com.pulumi.alicloud.slb.ListenerArgs;
 * import com.pulumi.alicloud.slb.ServerGroup;
 * import com.pulumi.alicloud.slb.ServerGroupArgs;
 * import com.pulumi.alicloud.slb.Rule;
 * import com.pulumi.alicloud.slb.RuleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var slbRuleName = config.get("slbRuleName").orElse("terraform-example");
 *         final var rule = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         var ruleNetwork = new Network("ruleNetwork", NetworkArgs.builder()        
 *             .vpcName(slbRuleName)
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var ruleSwitch = new Switch("ruleSwitch", SwitchArgs.builder()        
 *             .vpcId(ruleNetwork.id())
 *             .cidrBlock("172.16.0.0/16")
 *             .zoneId(rule.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *             .vswitchName(slbRuleName)
 *             .build());
 * 
 *         var ruleApplicationLoadBalancer = new ApplicationLoadBalancer("ruleApplicationLoadBalancer", ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(slbRuleName)
 *             .vswitchId(ruleSwitch.id())
 *             .instanceChargeType("PayByCLCU")
 *             .build());
 * 
 *         var ruleListener = new Listener("ruleListener", ListenerArgs.builder()        
 *             .loadBalancerId(ruleApplicationLoadBalancer.id())
 *             .backendPort(22)
 *             .frontendPort(22)
 *             .protocol("http")
 *             .bandwidth(5)
 *             .healthCheckConnectPort("20")
 *             .build());
 * 
 *         var ruleServerGroup = new ServerGroup("ruleServerGroup", ServerGroupArgs.builder()        
 *             .loadBalancerId(ruleApplicationLoadBalancer.id())
 *             .name(slbRuleName)
 *             .build());
 * 
 *         var ruleRule = new Rule("ruleRule", RuleArgs.builder()        
 *             .loadBalancerId(ruleApplicationLoadBalancer.id())
 *             .frontendPort(ruleListener.frontendPort())
 *             .name(slbRuleName)
 *             .domain("*.aliyun.com")
 *             .url("/image")
 *             .serverGroupId(ruleServerGroup.id())
 *             .cookie("23ffsa")
 *             .cookieTimeout(100)
 *             .healthCheckHttpCode("http_2xx")
 *             .healthCheckInterval(10)
 *             .healthCheckUri("/test")
 *             .healthCheckConnectPort(80)
 *             .healthCheckTimeout(30)
 *             .healthyThreshold(3)
 *             .unhealthyThreshold(5)
 *             .stickySession("on")
 *             .stickySessionType("server")
 *             .listenerSync("off")
 *             .scheduler("rr")
 *             .healthCheckDomain("test")
 *             .healthCheck("on")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Load balancer forwarding rule can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/rule:Rule example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/rule:Rule")
public class Rule extends com.pulumi.resources.CustomResource {
    /**
     * The cookie configured on the server. It is mandatory when `sticky_session` is `on` and `sticky_session_type` is `server`. Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being `1` - `200`. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
     * 
     */
    @Export(name="cookie", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cookie;

    /**
     * @return The cookie configured on the server. It is mandatory when `sticky_session` is `on` and `sticky_session_type` is `server`. Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being `1` - `200`. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.
     * 
     */
    public Output<Optional<String>> cookie() {
        return Codegen.optional(this.cookie);
    }
    /**
     * Cookie timeout. It is mandatory when `sticky_session` is `on` and `sticky_session_type` is `insert`. Otherwise, it will be ignored. Valid values: [1-86400] in seconds.
     * 
     */
    @Export(name="cookieTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cookieTimeout;

    /**
     * @return Cookie timeout. It is mandatory when `sticky_session` is `on` and `sticky_session_type` is `insert`. Otherwise, it will be ignored. Valid values: [1-86400] in seconds.
     * 
     */
    public Output<Optional<Integer>> cookieTimeout() {
        return Codegen.optional(this.cookieTimeout);
    }
    /**
     * Checking DeleteProtection of SLB instance before deleting. If `true`, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default value: `false`.
     * 
     */
    @Export(name="deleteProtectionValidation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteProtectionValidation;

    /**
     * @return Checking DeleteProtection of SLB instance before deleting. If `true`, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> deleteProtectionValidation() {
        return Codegen.optional(this.deleteProtectionValidation);
    }
    /**
     * Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
     * and wildcard characters. The following two domain name formats are supported:
     * - Standard domain name: www.test.com
     * - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domain;

    /**
     * @return Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.),
     * and wildcard characters. The following two domain name formats are supported:
     * - Standard domain name: www.test.com
     * - Wildcard domain name: *.test.com. wildcard (*) must be the first character in the format of (*.)
     * 
     */
    public Output<Optional<String>> domain() {
        return Codegen.optional(this.domain);
    }
    /**
     * The listener frontend port which is used to launch the new forwarding rule. Valid values: [1-65535].
     * 
     */
    @Export(name="frontendPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> frontendPort;

    /**
     * @return The listener frontend port which is used to launch the new forwarding rule. Valid values: [1-65535].
     * 
     */
    public Output<Integer> frontendPort() {
        return this.frontendPort;
    }
    /**
     * Whether to enable health check. Valid values: `on` and `off`. `TCP` and `UDP` listener&#39;s `health_check` is always `on`, so it will be ignore when launching `TCP` or `UDP` listener. **NOTE:** `health_check` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    @Export(name="healthCheck", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheck;

    /**
     * @return Whether to enable health check. Valid values: `on` and `off`. `TCP` and `UDP` listener&#39;s `health_check` is always `on`, so it will be ignore when launching `TCP` or `UDP` listener. **NOTE:** `health_check` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    public Output<Optional<String>> healthCheck() {
        return Codegen.optional(this.healthCheck);
    }
    /**
     * Port used for health check. Valid values: [1-65535]. Default value: `None` means the backend server port is used.
     * 
     */
    @Export(name="healthCheckConnectPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> healthCheckConnectPort;

    /**
     * @return Port used for health check. Valid values: [1-65535]. Default value: `None` means the backend server port is used.
     * 
     */
    public Output<Integer> healthCheckConnectPort() {
        return this.healthCheckConnectPort;
    }
    /**
     * Domain name used for health check. When it used to launch TCP listener, `health_check_type` must be `http`. Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty, Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
     * 
     */
    @Export(name="healthCheckDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckDomain;

    /**
     * @return Domain name used for health check. When it used to launch TCP listener, `health_check_type` must be `http`. Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty, Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.
     * 
     */
    public Output<Optional<String>> healthCheckDomain() {
        return Codegen.optional(this.healthCheckDomain);
    }
    /**
     * Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `health_check` is `on`. Default value: `http_2xx`. Valid values: `http_2xx`, `http_3xx`, `http_4xx` and `http_5xx`.
     * 
     */
    @Export(name="healthCheckHttpCode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckHttpCode;

    /**
     * @return Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when `health_check` is `on`. Default value: `http_2xx`. Valid values: `http_2xx`, `http_3xx`, `http_4xx` and `http_5xx`.
     * 
     */
    public Output<Optional<String>> healthCheckHttpCode() {
        return Codegen.optional(this.healthCheckHttpCode);
    }
    /**
     * Time interval of health checks. It is required when `health_check` is `on`. Valid values: [1-50] in seconds. Default value: `2`.
     * 
     */
    @Export(name="healthCheckInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckInterval;

    /**
     * @return Time interval of health checks. It is required when `health_check` is `on`. Valid values: [1-50] in seconds. Default value: `2`.
     * 
     */
    public Output<Optional<Integer>> healthCheckInterval() {
        return Codegen.optional(this.healthCheckInterval);
    }
    /**
     * Maximum timeout of each health check response. It is required when `health_check` is `on`. Valid values: [1-300] in seconds. Default value: `5`. Note: If `health_check_timeout` &lt; `health_check_interval`, its will be replaced by `health_check_interval`.
     * 
     */
    @Export(name="healthCheckTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckTimeout;

    /**
     * @return Maximum timeout of each health check response. It is required when `health_check` is `on`. Valid values: [1-300] in seconds. Default value: `5`. Note: If `health_check_timeout` &lt; `health_check_interval`, its will be replaced by `health_check_interval`.
     * 
     */
    public Output<Optional<Integer>> healthCheckTimeout() {
        return Codegen.optional(this.healthCheckTimeout);
    }
    /**
     * URI used for health check. When it used to launch TCP listener, `health_check_type` must be `http`. Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&amp;’ are allowed.
     * 
     */
    @Export(name="healthCheckUri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckUri;

    /**
     * @return URI used for health check. When it used to launch TCP listener, `health_check_type` must be `http`. Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&amp;’ are allowed.
     * 
     */
    public Output<Optional<String>> healthCheckUri() {
        return Codegen.optional(this.healthCheckUri);
    }
    /**
     * Threshold determining the result of the health check is success. It is required when `health_check` is `on`. Valid values: [1-10] in seconds. Default value: `3`.
     * 
     */
    @Export(name="healthyThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthyThreshold;

    /**
     * @return Threshold determining the result of the health check is success. It is required when `health_check` is `on`. Valid values: [1-10] in seconds. Default value: `3`.
     * 
     */
    public Output<Optional<Integer>> healthyThreshold() {
        return Codegen.optional(this.healthyThreshold);
    }
    /**
     * Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default value: `on`. Valid values: `on` and `off`.
     * 
     */
    @Export(name="listenerSync", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> listenerSync;

    /**
     * @return Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default value: `on`. Valid values: `on` and `off`.
     * 
     */
    public Output<Optional<String>> listenerSync() {
        return Codegen.optional(this.listenerSync);
    }
    /**
     * The Load Balancer ID which is used to launch the new forwarding rule.
     * 
     */
    @Export(name="loadBalancerId", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerId;

    /**
     * @return The Load Balancer ID which is used to launch the new forwarding rule.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * Name of the forwarding rule. Our plugin provides a default name: &#34;tf-slb-rule&#34;.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the forwarding rule. Our plugin provides a default name: &#34;tf-slb-rule&#34;.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Scheduling algorithm. Valid values: `wrr`, `rr` and `wlc`. Default value: `wrr`. **NOTE:** `scheduler` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    @Export(name="scheduler", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scheduler;

    /**
     * @return Scheduling algorithm. Valid values: `wrr`, `rr` and `wlc`. Default value: `wrr`. **NOTE:** `scheduler` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    public Output<Optional<String>> scheduler() {
        return Codegen.optional(this.scheduler);
    }
    /**
     * ID of a virtual server group that will be forwarded.
     * 
     */
    @Export(name="serverGroupId", refs={String.class}, tree="[0]")
    private Output<String> serverGroupId;

    /**
     * @return ID of a virtual server group that will be forwarded.
     * 
     */
    public Output<String> serverGroupId() {
        return this.serverGroupId;
    }
    /**
     * Whether to enable session persistence. Valid values: `on` and `off`. Default value: `off`. **NOTE:** `sticky_session` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    @Export(name="stickySession", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stickySession;

    /**
     * @return Whether to enable session persistence. Valid values: `on` and `off`. Default value: `off`. **NOTE:** `sticky_session` is required and takes effect only when `listener_sync` is set to `off`.
     * 
     */
    public Output<Optional<String>> stickySession() {
        return Codegen.optional(this.stickySession);
    }
    /**
     * Mode for handling the cookie. If `sticky_session` is `on`, it is mandatory. Otherwise, it will be ignored. Valid values: `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
     * 
     */
    @Export(name="stickySessionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stickySessionType;

    /**
     * @return Mode for handling the cookie. If `sticky_session` is `on`, it is mandatory. Otherwise, it will be ignored. Valid values: `insert` and `server`. `insert` means it is inserted from Server Load Balancer; `server` means the Server Load Balancer learns from the backend server.
     * 
     */
    public Output<Optional<String>> stickySessionType() {
        return Codegen.optional(this.stickySessionType);
    }
    /**
     * Threshold determining the result of the health check is fail. It is required when `health_check` is `on`. Valid values: [1-10] in seconds. Default value: `3`.
     * 
     */
    @Export(name="unhealthyThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> unhealthyThreshold;

    /**
     * @return Threshold determining the result of the health check is fail. It is required when `health_check` is `on`. Valid values: [1-10] in seconds. Default value: `3`.
     * 
     */
    public Output<Optional<Integer>> unhealthyThreshold() {
        return Codegen.optional(this.unhealthyThreshold);
    }
    /**
     * Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9, and characters &#39;-&#39; &#39;/&#39; &#39;?&#39; &#39;%&#39; &#39;#&#39; and &#39;&amp;&#39; are allowed. URLs must be started with the character &#39;/&#39;, but cannot be &#39;/&#39; alone.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9, and characters &#39;-&#39; &#39;/&#39; &#39;?&#39; &#39;%&#39; &#39;#&#39; and &#39;&amp;&#39; are allowed. URLs must be started with the character &#39;/&#39;, but cannot be &#39;/&#39; alone.
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rule(String name) {
        this(name, RuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rule(String name, RuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rule(String name, RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/rule:Rule", name, args == null ? RuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Rule(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/rule:Rule", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Rule get(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rule(name, id, state, options);
    }
}
