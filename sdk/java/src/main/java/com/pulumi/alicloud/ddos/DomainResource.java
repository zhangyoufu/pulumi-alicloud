// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ddos.DomainResourceArgs;
import com.pulumi.alicloud.ddos.inputs.DomainResourceState;
import com.pulumi.alicloud.ddos.outputs.DomainResourceProxyType;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Anti-DDoS Pro Domain Resource resource.
 * 
 * For information about Anti-DDoS Pro Domain Resource and how to use it, see [What is Domain Resource](https://www.alibabacloud.com/help/en/doc-detail/157463.htm).
 * 
 * &gt; **NOTE:** Available in v1.123.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.ddos.DomainResource;
 * import com.pulumi.alicloud.ddos.DomainResourceArgs;
 * import com.pulumi.alicloud.ddos.inputs.DomainResourceProxyTypeArgs;
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
 *         var example = new DomainResource(&#34;example&#34;, DomainResourceArgs.builder()        
 *             .domain(&#34;tftestacc1234.abc&#34;)
 *             .httpsExt(&#34;{\&#34;Http2\&#34;:1,\&#34;Http2https\&#34;:0，\&#34;Https2http\&#34;:0}&#34;)
 *             .instanceIds(&#34;ddoscoo-cn-6ja1rl4j****&#34;)
 *             .proxyTypes(DomainResourceProxyTypeArgs.builder()
 *                 .proxyPorts(443)
 *                 .proxyType(&#34;https&#34;)
 *                 .build())
 *             .realServers(&#34;177.167.32.11&#34;)
 *             .rsType(0)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Anti-DDoS Pro Domain Resource can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ddos/domainResource:DomainResource example &lt;domain&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ddos/domainResource:DomainResource")
public class DomainResource extends com.pulumi.resources.CustomResource {
    /**
     * The domain name of the website that you want to add to the instance.
     * 
     */
    @Export(name="domain", type=String.class, parameters={})
    private Output<String> domain;

    /**
     * @return The domain name of the website that you want to add to the instance.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
     * - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
     * - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
     * - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
     * 
     */
    @Export(name="httpsExt", type=String.class, parameters={})
    private Output<String> httpsExt;

    /**
     * @return The advanced HTTPS settings. This parameter takes effect only when the value of ProxyType includes https. This parameter is a string that contains a JSON struct. The JSON struct includes the following fields:
     * - `Http2https`: specifies whether to turn on Enforce HTTPS Routing. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enforce HTTPS Routing is turned off. The value `1` indicates that Enforce HTTPS Routing is turned on. The default value is `0`. If your website supports both HTTP and HTTPS, this feature suits your needs. If you turn on the switch, all HTTP requests are redirected to HTTPS requests on port 443 by default.
     * - `Https2http`: specifies whether to turn on Enable HTTP. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP is turned off. The value `1` indicates that Enable HTTP is turned on. The default value is `0`. If your website does not support HTTPS, this feature suits your needs. If you turn on the switch, all HTTPS requests are redirected to HTTP requests and forwarded to origin servers. The feature can also redirect WebSockets requests to WebSocket requests. All requests are redirected over port 80.
     * - `Http2`: specifies whether to turn on Enable HTTP/2. This field is optional and must be an integer. Valid values: `0` and `1`. The value `0` indicates that Enable HTTP/2 is turned off. The value `1` indicates that Enable HTTP/2 is turned on. The default value is `0`. After you turn on the switch, the protocol type is HTTP/2.
     * 
     */
    public Output<String> httpsExt() {
        return this.httpsExt;
    }
    @Export(name="instanceIds", type=List.class, parameters={String.class})
    private Output<List<String>> instanceIds;

    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }
    /**
     * Protocol type and port number information.
     * 
     */
    @Export(name="proxyTypes", type=List.class, parameters={DomainResourceProxyType.class})
    private Output<List<DomainResourceProxyType>> proxyTypes;

    /**
     * @return Protocol type and port number information.
     * 
     */
    public Output<List<DomainResourceProxyType>> proxyTypes() {
        return this.proxyTypes;
    }
    /**
     * the IP address. This field is required and must be a string array.
     * 
     */
    @Export(name="realServers", type=List.class, parameters={String.class})
    private Output<List<String>> realServers;

    /**
     * @return the IP address. This field is required and must be a string array.
     * 
     */
    public Output<List<String>> realServers() {
        return this.realServers;
    }
    /**
     * The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
     * 
     */
    @Export(name="rsType", type=Integer.class, parameters={})
    private Output<Integer> rsType;

    /**
     * @return The address type of the origin server. Valid values: `0`: IP address. `1`: domain name. Use the domain name of the origin server if you deploy proxies, such as Web Application Firewall (WAF), between the origin server and the Anti-DDoS Pro or Anti-DDoS Premium instance. If you use the domain name, you must enter the address of the proxy, such as the CNAME of WAF.
     * 
     */
    public Output<Integer> rsType() {
        return this.rsType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DomainResource(String name) {
        this(name, DomainResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DomainResource(String name, DomainResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DomainResource(String name, DomainResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/domainResource:DomainResource", name, args == null ? DomainResourceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DomainResource(String name, Output<String> id, @Nullable DomainResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ddos/domainResource:DomainResource", name, state, makeResourceOptions(options, id));
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
    public static DomainResource get(String name, Output<String> id, @Nullable DomainResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DomainResource(name, id, state, options);
    }
}
