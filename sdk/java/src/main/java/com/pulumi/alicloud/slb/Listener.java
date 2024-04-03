// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.slb.ListenerArgs;
import com.pulumi.alicloud.slb.inputs.ListenerState;
import com.pulumi.alicloud.slb.outputs.ListenerXForwardedFor;
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
 * Provides a Classic Load Balancer (SLB) Load Balancer Listener resource.
 * 
 * For information about Classic Load Balancer (SLB) and how to use it, see [What is Classic Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
 * 
 * For information about listener and how to use it, please see the following:
 * 
 * * [Configure a HTTP Classic Load Balancer (SLB) Listener](https://www.alibabacloud.com/help/doc-detail/27592.htm).
 * * [Configure a HTTPS Classic Load Balancer (SLB) Listener](https://www.alibabacloud.com/help/doc-detail/27593.htm).
 * * [Configure a TCP Classic Load Balancer (SLB) Listener](https://www.alibabacloud.com/help/doc-detail/27594.htm).
 * * [Configure a UDP Classic Load Balancer (SLB) Listener](https://www.alibabacloud.com/help/doc-detail/27595.htm).
 * 
 * &gt; **NOTE:** Available since v1.0.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancer;
 * import com.pulumi.alicloud.slb.ApplicationLoadBalancerArgs;
 * import com.pulumi.alicloud.slb.Acl;
 * import com.pulumi.alicloud.slb.AclArgs;
 * import com.pulumi.alicloud.slb.Listener;
 * import com.pulumi.alicloud.slb.ListenerArgs;
 * import com.pulumi.alicloud.slb.inputs.ListenerXForwardedForArgs;
 * import com.pulumi.alicloud.slb.AclEntryAttachment;
 * import com.pulumi.alicloud.slb.AclEntryAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         var default_ = new RandomInteger(&#34;default&#34;, RandomIntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var listenerApplicationLoadBalancer = new ApplicationLoadBalancer(&#34;listenerApplicationLoadBalancer&#34;, ApplicationLoadBalancerArgs.builder()        
 *             .loadBalancerName(default_.result().applyValue(result -&gt; String.format(&#34;%s-%s&#34;, name,result)))
 *             .internetChargeType(&#34;PayByTraffic&#34;)
 *             .addressType(&#34;internet&#34;)
 *             .instanceChargeType(&#34;PayByCLCU&#34;)
 *             .build());
 * 
 *         var listenerAcl = new Acl(&#34;listenerAcl&#34;, AclArgs.builder()        
 *             .ipVersion(&#34;ipv4&#34;)
 *             .build());
 * 
 *         var listenerListener = new Listener(&#34;listenerListener&#34;, ListenerArgs.builder()        
 *             .loadBalancerId(listenerApplicationLoadBalancer.id())
 *             .backendPort(80)
 *             .frontendPort(80)
 *             .protocol(&#34;http&#34;)
 *             .bandwidth(10)
 *             .stickySession(&#34;on&#34;)
 *             .stickySessionType(&#34;insert&#34;)
 *             .cookieTimeout(86400)
 *             .cookie(&#34;tfslblistenercookie&#34;)
 *             .healthCheck(&#34;on&#34;)
 *             .healthCheckDomain(&#34;ali.com&#34;)
 *             .healthCheckUri(&#34;/cons&#34;)
 *             .healthCheckConnectPort(20)
 *             .healthyThreshold(8)
 *             .unhealthyThreshold(8)
 *             .healthCheckTimeout(8)
 *             .healthCheckInterval(5)
 *             .healthCheckHttpCode(&#34;http_2xx,http_3xx&#34;)
 *             .xForwardedFor(ListenerXForwardedForArgs.builder()
 *                 .retriveSlbIp(true)
 *                 .retriveSlbId(true)
 *                 .build())
 *             .aclStatus(&#34;on&#34;)
 *             .aclType(&#34;white&#34;)
 *             .aclId(listenerAcl.id())
 *             .requestTimeout(80)
 *             .idleTimeout(30)
 *             .build());
 * 
 *         var first = new AclEntryAttachment(&#34;first&#34;, AclEntryAttachmentArgs.builder()        
 *             .aclId(listenerAcl.id())
 *             .entry(&#34;10.10.10.0/24&#34;)
 *             .comment(&#34;first&#34;)
 *             .build());
 * 
 *         var second = new AclEntryAttachment(&#34;second&#34;, AclEntryAttachmentArgs.builder()        
 *             .aclId(listenerAcl.id())
 *             .entry(&#34;168.10.10.0/24&#34;)
 *             .comment(&#34;second&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Classic Load Balancer (SLB) Load Balancer Listener can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:slb/listener:Listener example &lt;load_balancer_id&gt;:&lt;protocol&gt;:&lt;frontend_port&gt;
 * ```
 * 
 * ```sh
 * $ pulumi import alicloud:slb/listener:Listener example &lt;load_balancer_id&gt;:&lt;frontend_port&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:slb/listener:Listener")
public class Listener extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the network ACL that is associated with the listener. **NOTE:** If `acl_status` is set to `on`, `acl_id` is required. Otherwise, it will be ignored.
     * 
     */
    @Export(name="aclId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aclId;

    /**
     * @return The ID of the network ACL that is associated with the listener. **NOTE:** If `acl_status` is set to `on`, `acl_id` is required. Otherwise, it will be ignored.
     * 
     */
    public Output<Optional<String>> aclId() {
        return Codegen.optional(this.aclId);
    }
    /**
     * Specifies whether to enable access control. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    @Export(name="aclStatus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aclStatus;

    /**
     * @return Specifies whether to enable access control. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    public Output<Optional<String>> aclStatus() {
        return Codegen.optional(this.aclStatus);
    }
    /**
     * The type of the network ACL. Valid values: `black`, `white`. **NOTE:** If `acl_status` is set to `on`, `acl_type` is required. Otherwise, it will be ignored.
     * 
     */
    @Export(name="aclType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aclType;

    /**
     * @return The type of the network ACL. Valid values: `black`, `white`. **NOTE:** If `acl_status` is set to `on`, `acl_type` is required. Otherwise, it will be ignored.
     * 
     */
    public Output<Optional<String>> aclType() {
        return Codegen.optional(this.aclType);
    }
    /**
     * The backend port that is used by the CLB instance. Valid values: `1` to `65535`. **NOTE:** If `server_group_id` is not set, `backend_port` is required.
     * 
     */
    @Export(name="backendPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> backendPort;

    /**
     * @return The backend port that is used by the CLB instance. Valid values: `1` to `65535`. **NOTE:** If `server_group_id` is not set, `backend_port` is required.
     * 
     */
    public Output<Optional<Integer>> backendPort() {
        return Codegen.optional(this.backendPort);
    }
    /**
     * The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
     * - `-1`: If you set `bandwidth` to `-1`, the bandwidth of the listener is unlimited.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
     * - `-1`: If you set `bandwidth` to `-1`, the bandwidth of the listener is unlimited.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The ID of the certification authority (CA) certificate.
     * 
     */
    @Export(name="caCertificateId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> caCertificateId;

    /**
     * @return The ID of the certification authority (CA) certificate.
     * 
     */
    public Output<Optional<String>> caCertificateId() {
        return Codegen.optional(this.caCertificateId);
    }
    /**
     * The cookie that is configured on the server. The `cookie` must be `1` to `200` characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($). **NOTE:** If `sticky_session` is set to `on`, and `sticky_session_type` is set to `server`, `cookie` is required. Otherwise, it will be ignored.
     * 
     */
    @Export(name="cookie", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cookie;

    /**
     * @return The cookie that is configured on the server. The `cookie` must be `1` to `200` characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($). **NOTE:** If `sticky_session` is set to `on`, and `sticky_session_type` is set to `server`, `cookie` is required. Otherwise, it will be ignored.
     * 
     */
    public Output<Optional<String>> cookie() {
        return Codegen.optional(this.cookie);
    }
    /**
     * The timeout period of a cookie. Unit: seconds. Valid values: `1` to `86400`. **NOTE:** If `sticky_session` is set to `on`, and `sticky_session_type` is set to `insert`, `cookie_timeout` is required. Otherwise, it will be ignored.
     * 
     */
    @Export(name="cookieTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cookieTimeout;

    /**
     * @return The timeout period of a cookie. Unit: seconds. Valid values: `1` to `86400`. **NOTE:** If `sticky_session` is set to `on`, and `sticky_session_type` is set to `insert`, `cookie_timeout` is required. Otherwise, it will be ignored.
     * 
     */
    public Output<Optional<Integer>> cookieTimeout() {
        return Codegen.optional(this.cookieTimeout);
    }
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default value: `false`.
     * 
     */
    @Export(name="deleteProtectionValidation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteProtectionValidation;

    /**
     * @return Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> deleteProtectionValidation() {
        return Codegen.optional(this.deleteProtectionValidation);
    }
    /**
     * The name of the listener. The name must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The name of the listener. The name must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether to enable HTTP/2. Default value: `on`. Valid values: `on`, `off`.
     * 
     */
    @Export(name="enableHttp2", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enableHttp2;

    /**
     * @return Specifies whether to enable HTTP/2. Default value: `on`. Valid values: `on`, `off`.
     * 
     */
    public Output<Optional<String>> enableHttp2() {
        return Codegen.optional(this.enableHttp2);
    }
    /**
     * The timeout period of a connection. Unit: seconds. Default value: `900`. Valid values: `10` to `900`.
     * 
     */
    @Export(name="establishedTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> establishedTimeout;

    /**
     * @return The timeout period of a connection. Unit: seconds. Default value: `900`. Valid values: `10` to `900`.
     * 
     */
    public Output<Optional<Integer>> establishedTimeout() {
        return Codegen.optional(this.establishedTimeout);
    }
    /**
     * The listening port that is used to redirect HTTP requests to HTTPS.
     * 
     */
    @Export(name="forwardPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> forwardPort;

    /**
     * @return The listening port that is used to redirect HTTP requests to HTTPS.
     * 
     */
    public Output<Optional<Integer>> forwardPort() {
        return Codegen.optional(this.forwardPort);
    }
    /**
     * The frontend port that is used by the CLB instance. Valid values: `1` to `65535`.
     * 
     */
    @Export(name="frontendPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> frontendPort;

    /**
     * @return The frontend port that is used by the CLB instance. Valid values: `1` to `65535`.
     * 
     */
    public Output<Integer> frontendPort() {
        return this.frontendPort;
    }
    /**
     * Specifies whether to enable GZIP compression to compress specific types of files. Default value: `true`. Valid values: `true`, `false`.
     * 
     */
    @Export(name="gzip", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> gzip;

    /**
     * @return Specifies whether to enable GZIP compression to compress specific types of files. Default value: `true`. Valid values: `true`, `false`.
     * 
     */
    public Output<Optional<Boolean>> gzip() {
        return Codegen.optional(this.gzip);
    }
    /**
     * Specifies whether to enable the health check feature. Default value: `on`. Valid values: `on`, `off`. **NOTE:** `TCP` and `UDP` listener&#39;s HealthCheck is always on, so it will be ignored when launching `TCP` or `UDP` listener.
     * 
     */
    @Export(name="healthCheck", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheck;

    /**
     * @return Specifies whether to enable the health check feature. Default value: `on`. Valid values: `on`, `off`. **NOTE:** `TCP` and `UDP` listener&#39;s HealthCheck is always on, so it will be ignored when launching `TCP` or `UDP` listener.
     * 
     */
    public Output<Optional<String>> healthCheck() {
        return Codegen.optional(this.healthCheck);
    }
    /**
     * The backend port that is used for health checks. Valid values: `0` to `65535`. **NOTE:** `health_check_connect_port` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckConnectPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> healthCheckConnectPort;

    /**
     * @return The backend port that is used for health checks. Valid values: `0` to `65535`. **NOTE:** `health_check_connect_port` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Integer> healthCheckConnectPort() {
        return this.healthCheckConnectPort;
    }
    /**
     * The domain name that is used for health checks. **NOTE:** `health_check_domain` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckDomain;

    /**
     * @return The domain name that is used for health checks. **NOTE:** `health_check_domain` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<String>> healthCheckDomain() {
        return Codegen.optional(this.healthCheckDomain);
    }
    /**
     * The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (`,`). Default value: `http_2xx`. Valid values: `http_2xx`, `http_3xx`, `http_4xx` and `http_5xx`. **NOTE:** `health_check_http_code` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckHttpCode", refs={String.class}, tree="[0]")
    private Output<String> healthCheckHttpCode;

    /**
     * @return The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (`,`). Default value: `http_2xx`. Valid values: `http_2xx`, `http_3xx`, `http_4xx` and `http_5xx`. **NOTE:** `health_check_http_code` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<String> healthCheckHttpCode() {
        return this.healthCheckHttpCode;
    }
    /**
     * The interval between two consecutive health checks. Unit: seconds. Default value: `2`. Valid values: `1` to `50`. **NOTE:** `health_check_interval` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckInterval;

    /**
     * @return The interval between two consecutive health checks. Unit: seconds. Default value: `2`. Valid values: `1` to `50`. **NOTE:** `health_check_interval` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<Integer>> healthCheckInterval() {
        return Codegen.optional(this.healthCheckInterval);
    }
    /**
     * The health check method used in HTTP health checks. Valid values: `head`, `get`. **NOTE:** `health_check_method` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckMethod", refs={String.class}, tree="[0]")
    private Output<String> healthCheckMethod;

    /**
     * @return The health check method used in HTTP health checks. Valid values: `head`, `get`. **NOTE:** `health_check_method` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<String> healthCheckMethod() {
        return this.healthCheckMethod;
    }
    /**
     * The timeout period of a health check response. Unit: seconds. Default value: `5`. Valid values: `1` to `300`. **NOTE:** If `health_check_timeout` &lt; `health_check_interval`, `health_check_timeout` will be replaced by `health_check_interval`. `health_check_timeout` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthCheckTimeout;

    /**
     * @return The timeout period of a health check response. Unit: seconds. Default value: `5`. Valid values: `1` to `300`. **NOTE:** If `health_check_timeout` &lt; `health_check_interval`, `health_check_timeout` will be replaced by `health_check_interval`. `health_check_timeout` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<Integer>> healthCheckTimeout() {
        return Codegen.optional(this.healthCheckTimeout);
    }
    /**
     * The type of health checks. Default value: `tcp`. Valid values: `tcp`, `http`.
     * 
     */
    @Export(name="healthCheckType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckType;

    /**
     * @return The type of health checks. Default value: `tcp`. Valid values: `tcp`, `http`.
     * 
     */
    public Output<Optional<String>> healthCheckType() {
        return Codegen.optional(this.healthCheckType);
    }
    /**
     * The URI that is used for health checks. The `health_check_uri` must be `1` to `80` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%!)(MISSING), question marks (?), number signs (#), and ampersands (&amp;). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
     * **NOTE:** `health_check_uri` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthCheckUri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> healthCheckUri;

    /**
     * @return The URI that is used for health checks. The `health_check_uri` must be `1` to `80` characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%!)(MISSING), question marks (?), number signs (#), and ampersands (&amp;). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
     * **NOTE:** `health_check_uri` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<String>> healthCheckUri() {
        return Codegen.optional(this.healthCheckUri);
    }
    /**
     * The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. Default value: `3`. Valid values: `2` to `10`. **NOTE:** `healthy_threshold` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="healthyThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> healthyThreshold;

    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. Default value: `3`. Valid values: `2` to `10`. **NOTE:** `healthy_threshold` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<Integer>> healthyThreshold() {
        return Codegen.optional(this.healthyThreshold);
    }
    /**
     * The timeout period of an idle connection. Unit: seconds. Default value: `15`. Valid values: `1` to `60`.
     * 
     */
    @Export(name="idleTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> idleTimeout;

    /**
     * @return The timeout period of an idle connection. Unit: seconds. Default value: `15`. Valid values: `1` to `60`.
     * 
     */
    public Output<Optional<Integer>> idleTimeout() {
        return Codegen.optional(this.idleTimeout);
    }
    /**
     * @deprecated
     * Field &#39;lb_port&#39; has been removed since 1.211.0.
     * 
     */
    @Deprecated /* Field 'lb_port' has been removed since 1.211.0. */
    @Export(name="lbPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> lbPort;

    public Output<Optional<Integer>> lbPort() {
        return Codegen.optional(this.lbPort);
    }
    /**
     * @deprecated
     * Field &#39;lb_protocol&#39; has been removed since 1.211.0.
     * 
     */
    @Deprecated /* Field 'lb_protocol' has been removed since 1.211.0. */
    @Export(name="lbProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lbProtocol;

    public Output<Optional<String>> lbProtocol() {
        return Codegen.optional(this.lbProtocol);
    }
    /**
     * Specifies whether to enable HTTP-to-HTTPS redirection. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    @Export(name="listenerForward", refs={String.class}, tree="[0]")
    private Output<String> listenerForward;

    /**
     * @return Specifies whether to enable HTTP-to-HTTPS redirection. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    public Output<String> listenerForward() {
        return this.listenerForward;
    }
    /**
     * The Load Balancer ID which is used to launch a new listener.
     * 
     */
    @Export(name="loadBalancerId", refs={String.class}, tree="[0]")
    private Output<String> loadBalancerId;

    /**
     * @return The Load Balancer ID which is used to launch a new listener.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * The ID of the primary/secondary server group. **NOTE:** You cannot set both `server_group_id` and `master_slave_server_group_id`.
     * 
     */
    @Export(name="masterSlaveServerGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> masterSlaveServerGroupId;

    /**
     * @return The ID of the primary/secondary server group. **NOTE:** You cannot set both `server_group_id` and `master_slave_server_group_id`.
     * 
     */
    public Output<Optional<String>> masterSlaveServerGroupId() {
        return Codegen.optional(this.masterSlaveServerGroupId);
    }
    /**
     * The timeout period of session persistence. Unit: seconds. Default value: `0`. Valid values: `0` to `3600`.
     * 
     */
    @Export(name="persistenceTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> persistenceTimeout;

    /**
     * @return The timeout period of session persistence. Unit: seconds. Default value: `0`. Valid values: `0` to `3600`.
     * 
     */
    public Output<Optional<Integer>> persistenceTimeout() {
        return Codegen.optional(this.persistenceTimeout);
    }
    /**
     * The protocol to listen on. Valid values: `http`.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol to listen on. Valid values: `http`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    @Export(name="proxyProtocolV2Enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> proxyProtocolV2Enabled;

    /**
     * @return Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Default value: `false`. Valid values: `true`, `false`.
     * 
     */
    public Output<Boolean> proxyProtocolV2Enabled() {
        return this.proxyProtocolV2Enabled;
    }
    /**
     * The timeout period of a request. Unit: seconds. Default value: `60`. Valid values: `1` to `180`.
     * 
     */
    @Export(name="requestTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> requestTimeout;

    /**
     * @return The timeout period of a request. Unit: seconds. Default value: `60`. Valid values: `1` to `180`.
     * 
     */
    public Output<Optional<Integer>> requestTimeout() {
        return Codegen.optional(this.requestTimeout);
    }
    /**
     * The scheduling algorithm. Default value: `wrr`. Valid values:
     * 
     */
    @Export(name="scheduler", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scheduler;

    /**
     * @return The scheduling algorithm. Default value: `wrr`. Valid values:
     * 
     */
    public Output<Optional<String>> scheduler() {
        return Codegen.optional(this.scheduler);
    }
    /**
     * The ID of the server certificate. **NOTE:** `server_certificate_id` is also required when the value of the `ssl_certificate_id` is Empty.
     * 
     */
    @Export(name="serverCertificateId", refs={String.class}, tree="[0]")
    private Output<String> serverCertificateId;

    /**
     * @return The ID of the server certificate. **NOTE:** `server_certificate_id` is also required when the value of the `ssl_certificate_id` is Empty.
     * 
     */
    public Output<String> serverCertificateId() {
        return this.serverCertificateId;
    }
    /**
     * The ID of the vServer group. It&#39;s the ID of resource `alicloud.slb.ServerGroup`.
     * 
     */
    @Export(name="serverGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverGroupId;

    /**
     * @return The ID of the vServer group. It&#39;s the ID of resource `alicloud.slb.ServerGroup`.
     * 
     */
    public Output<Optional<String>> serverGroupId() {
        return Codegen.optional(this.serverGroupId);
    }
    /**
     * The ID of the server certificate. **NOTE:** Field `ssl_certificate_id` has been deprecated from provider version 1.59.0. New field `server_certificate_id` instead.
     * 
     * @deprecated
     * Field &#39;ssl_certificate_id&#39; has been deprecated from 1.59.0 and using &#39;server_certificate_id&#39; instead.
     * 
     */
    @Deprecated /* Field 'ssl_certificate_id' has been deprecated from 1.59.0 and using 'server_certificate_id' instead. */
    @Export(name="sslCertificateId", refs={String.class}, tree="[0]")
    private Output<String> sslCertificateId;

    /**
     * @return The ID of the server certificate. **NOTE:** Field `ssl_certificate_id` has been deprecated from provider version 1.59.0. New field `server_certificate_id` instead.
     * 
     */
    public Output<String> sslCertificateId() {
        return this.sslCertificateId;
    }
    /**
     * Specifies whether to enable session persistence. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    @Export(name="stickySession", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stickySession;

    /**
     * @return Specifies whether to enable session persistence. Default value: `off`. Valid values: `on`, `off`.
     * 
     */
    public Output<Optional<String>> stickySession() {
        return Codegen.optional(this.stickySession);
    }
    /**
     * The method that is used to handle a cookie. Valid values: `insert`, `server`. **NOTE:** If `sticky_session` is set to `on`, `sticky_session_type` is required. Otherwise, it will be ignored.
     * 
     */
    @Export(name="stickySessionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stickySessionType;

    /**
     * @return The method that is used to handle a cookie. Valid values: `insert`, `server`. **NOTE:** If `sticky_session` is set to `on`, `sticky_session_type` is required. Otherwise, it will be ignored.
     * 
     */
    public Output<Optional<String>> stickySessionType() {
        return Codegen.optional(this.stickySessionType);
    }
    /**
     * The Transport Layer Security (TLS) security policy. Default value: `tls_cipher_policy_1_0`. Valid values: `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`.
     * 
     */
    @Export(name="tlsCipherPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tlsCipherPolicy;

    /**
     * @return The Transport Layer Security (TLS) security policy. Default value: `tls_cipher_policy_1_0`. Valid values: `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`.
     * 
     */
    public Output<Optional<String>> tlsCipherPolicy() {
        return Codegen.optional(this.tlsCipherPolicy);
    }
    /**
     * The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. Default value: `3`. Valid values: `2` to `10`. **NOTE:** `unhealthy_threshold` takes effect only if `health_check` is set to `on`.
     * 
     */
    @Export(name="unhealthyThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> unhealthyThreshold;

    /**
     * @return The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. Default value: `3`. Valid values: `2` to `10`. **NOTE:** `unhealthy_threshold` takes effect only if `health_check` is set to `on`.
     * 
     */
    public Output<Optional<Integer>> unhealthyThreshold() {
        return Codegen.optional(this.unhealthyThreshold);
    }
    /**
     * Whether to set additional HTTP Header field &#34;X-Forwarded-For&#34;. See `x_forwarded_for` below.
     * 
     */
    @Export(name="xForwardedFor", refs={ListenerXForwardedFor.class}, tree="[0]")
    private Output<ListenerXForwardedFor> xForwardedFor;

    /**
     * @return Whether to set additional HTTP Header field &#34;X-Forwarded-For&#34;. See `x_forwarded_for` below.
     * 
     */
    public Output<ListenerXForwardedFor> xForwardedFor() {
        return this.xForwardedFor;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Listener(String name) {
        this(name, ListenerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Listener(String name, ListenerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Listener(String name, ListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/listener:Listener", name, args == null ? ListenerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Listener(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:slb/listener:Listener", name, state, makeResourceOptions(options, id));
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
    public static Listener get(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Listener(name, id, state, options);
    }
}
