// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.waf;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.waf.DomainArgs;
import com.pulumi.alicloud.waf.inputs.DomainState;
import com.pulumi.alicloud.waf.outputs.DomainLogHeader;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **DEPRECATED:**  This resource has been deprecated and using alicloud.wafv3.Domain instead.
 * 
 * Provides a WAF Domain resource to create domain in the Web Application Firewall.
 * 
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 * 
 * &gt; **NOTE:** Available in 1.82.0+ .
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
 * import com.pulumi.alicloud.waf.Domain;
 * import com.pulumi.alicloud.waf.DomainArgs;
 * import com.pulumi.alicloud.waf.inputs.DomainLogHeaderArgs;
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
 *         var domain = new Domain("domain", DomainArgs.builder()
 *             .domainName("alicloud-provider.cn")
 *             .instanceId("waf-123455")
 *             .isAccessProduct("On")
 *             .sourceIps("1.1.1.1")
 *             .clusterType("PhysicalCluster")
 *             .http2Ports(443)
 *             .httpPorts(80)
 *             .httpsPorts(443)
 *             .httpToUserIp("Off")
 *             .httpsRedirect("Off")
 *             .loadBalancing("IpHash")
 *             .logHeaders(DomainLogHeaderArgs.builder()
 *                 .key("foo")
 *                 .value("http")
 *                 .build())
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
 * WAF domain can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:waf/domain:Domain domain waf-132435:www.domain.com
 * ```
 * 
 */
@ResourceType(type="alicloud:waf/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterType;

    /**
     * @return The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
     * 
     */
    public Output<Optional<String>> clusterType() {
        return Codegen.optional(this.clusterType);
    }
    /**
     * The CNAME record assigned by the WAF instance to the specified domain.
     * 
     */
    @Export(name="cname", refs={String.class}, tree="[0]")
    private Output<String> cname;

    /**
     * @return The CNAME record assigned by the WAF instance to the specified domain.
     * 
     */
    public Output<String> cname() {
        return this.cname;
    }
    /**
     * The connection timeout for WAF exclusive clusters. Unit: seconds.
     * 
     */
    @Export(name="connectionTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> connectionTime;

    /**
     * @return The connection timeout for WAF exclusive clusters. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> connectionTime() {
        return Codegen.optional(this.connectionTime);
    }
    /**
     * Field `domain` has been deprecated from version 1.94.0. Use `domain_name` instead.
     * 
     * @deprecated
     * Field &#39;domain&#39; has been deprecated from version 1.94.0. Use &#39;domain_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'domain' has been deprecated from version 1.94.0. Use 'domain_name' instead. */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return Field `domain` has been deprecated from version 1.94.0. Use `domain_name` instead.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * The domain that you want to add to WAF. The `domain_name` is required when the value of the `domain`  is Empty.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The domain that you want to add to WAF. The `domain_name` is required when the value of the `domain`  is Empty.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * List of the HTTP 2.0 ports.
     * 
     */
    @Export(name="http2Ports", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> http2Ports;

    /**
     * @return List of the HTTP 2.0 ports.
     * 
     */
    public Output<Optional<List<String>>> http2Ports() {
        return Codegen.optional(this.http2Ports);
    }
    /**
     * List of the HTTP ports.
     * 
     */
    @Export(name="httpPorts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> httpPorts;

    /**
     * @return List of the HTTP ports.
     * 
     */
    public Output<Optional<List<String>>> httpPorts() {
        return Codegen.optional(this.httpPorts);
    }
    /**
     * Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server.
     * By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
     * 
     */
    @Export(name="httpToUserIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpToUserIp;

    /**
     * @return Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server.
     * By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
     * 
     */
    public Output<Optional<String>> httpToUserIp() {
        return Codegen.optional(this.httpToUserIp);
    }
    /**
     * List of the HTTPS ports.
     * 
     */
    @Export(name="httpsPorts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> httpsPorts;

    /**
     * @return List of the HTTPS ports.
     * 
     */
    public Output<Optional<List<String>>> httpsPorts() {
        return Codegen.optional(this.httpsPorts);
    }
    /**
     * Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: &#34;On&#34; and `Off`. Default to `Off`.
     * 
     */
    @Export(name="httpsRedirect", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpsRedirect;

    /**
     * @return Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: &#34;On&#34; and `Off`. Default to `Off`.
     * 
     */
    public Output<Optional<String>> httpsRedirect() {
        return Codegen.optional(this.httpsRedirect);
    }
    /**
     * The ID of the WAF instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the WAF instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
     * 
     */
    @Export(name="isAccessProduct", refs={String.class}, tree="[0]")
    private Output<String> isAccessProduct;

    /**
     * @return Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
     * 
     */
    public Output<String> isAccessProduct() {
        return this.isAccessProduct;
    }
    /**
     * The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
     * 
     */
    @Export(name="loadBalancing", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loadBalancing;

    /**
     * @return The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
     * 
     */
    public Output<Optional<String>> loadBalancing() {
        return Codegen.optional(this.loadBalancing);
    }
    /**
     * The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
     * * key: The key of label
     * * value: The value of label
     * 
     */
    @Export(name="logHeaders", refs={List.class,DomainLogHeader.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DomainLogHeader>> logHeaders;

    /**
     * @return The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
     * * key: The key of label
     * * value: The value of label
     * 
     */
    public Output<Optional<List<DomainLogHeader>>> logHeaders() {
        return Codegen.optional(this.logHeaders);
    }
    /**
     * The read timeout of a WAF exclusive cluster. Unit: seconds.
     * 
     */
    @Export(name="readTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> readTime;

    /**
     * @return The read timeout of a WAF exclusive cluster. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> readTime() {
        return Codegen.optional(this.readTime);
    }
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * List of the IP address or domain of the origin server to which the specified domain points.
     * 
     */
    @Export(name="sourceIps", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> sourceIps;

    /**
     * @return List of the IP address or domain of the origin server to which the specified domain points.
     * 
     */
    public Output<Optional<List<String>>> sourceIps() {
        return Codegen.optional(this.sourceIps);
    }
    /**
     * The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
     * 
     */
    @Export(name="writeTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> writeTime;

    /**
     * @return The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
     * 
     */
    public Output<Optional<Integer>> writeTime() {
        return Codegen.optional(this.writeTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(java.lang.String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(java.lang.String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(java.lang.String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:waf/domain:Domain", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Domain(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:waf/domain:Domain", name, state, makeResourceOptions(options, id), false);
    }

    private static DomainArgs makeArgs(DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DomainArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Domain get(java.lang.String name, Output<java.lang.String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
