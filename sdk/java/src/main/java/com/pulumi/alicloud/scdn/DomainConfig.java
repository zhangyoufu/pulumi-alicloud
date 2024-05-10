// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.scdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.scdn.DomainConfigArgs;
import com.pulumi.alicloud.scdn.inputs.DomainConfigState;
import com.pulumi.alicloud.scdn.outputs.DomainConfigFunctionArg;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a SCDN Accelerated Domain resource.
 * 
 * For information about domain config and how to use it, see [Batch set config](https://help.aliyun.com/document_detail/92912.html)
 * 
 * &gt; **NOTE:** Available in v1.131.0+.
 * 
 * &gt; **NOTE:** Alibaba Cloud SCDN has stopped new customer purchases from January 26, 2023, and you can choose to buy Alibaba Cloud DCDN products with more comprehensive acceleration and protection capabilities. If you are already a SCDN customer, you can submit a work order at any time to apply for a smooth migration to Alibaba Cloud DCDN products. In the future, we will provide better acceleration and security protection services in Alibaba Cloud DCDN, thank you for your understanding and cooperation.
 * 
 * &gt; **DEPRECATED:**  This resource has been [deprecated](https://www.aliyun.com/product/scdn) from version `1.219.0`.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.scdn.Domain;
 * import com.pulumi.alicloud.scdn.DomainArgs;
 * import com.pulumi.alicloud.scdn.inputs.DomainSourceArgs;
 * import com.pulumi.alicloud.scdn.DomainConfig;
 * import com.pulumi.alicloud.scdn.DomainConfigArgs;
 * import com.pulumi.alicloud.scdn.inputs.DomainConfigFunctionArgArgs;
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
 *         // Create a new Domain config.
 *         var domain = new Domain("domain", DomainArgs.builder()        
 *             .domainName("mydomain.alicloud-provider.cn")
 *             .cdnType("web")
 *             .scope("overseas")
 *             .sources(DomainSourceArgs.builder()
 *                 .content("1.1.1.1")
 *                 .type("ipaddr")
 *                 .priority("20")
 *                 .port(80)
 *                 .build())
 *             .build());
 * 
 *         var config = new DomainConfig("config", DomainConfigArgs.builder()        
 *             .domainName(domain.domainName())
 *             .functionName("ip_allow_list_set")
 *             .functionArgs(DomainConfigFunctionArgArgs.builder()
 *                 .argName("ip_list")
 *                 .argValue("110.110.110.110")
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
 * SCDN domain config can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:scdn/domainConfig:DomainConfig example &lt;domain_name&gt;:&lt;function_name&gt;:&lt;config_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:scdn/domainConfig:DomainConfig")
public class DomainConfig extends com.pulumi.resources.CustomResource {
    /**
     * The SCDN domain config id.
     * 
     */
    @Export(name="configId", refs={String.class}, tree="[0]")
    private Output<String> configId;

    /**
     * @return The SCDN domain config id.
     * 
     */
    public Output<String> configId() {
        return this.configId;
    }
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
    }
    /**
     * The args of the domain config.
     * 
     */
    @Export(name="functionArgs", refs={List.class,DomainConfigFunctionArg.class}, tree="[0,1]")
    private Output<List<DomainConfigFunctionArg>> functionArgs;

    /**
     * @return The args of the domain config.
     * 
     */
    public Output<List<DomainConfigFunctionArg>> functionArgs() {
        return this.functionArgs;
    }
    /**
     * The name of the domain config.
     * 
     */
    @Export(name="functionName", refs={String.class}, tree="[0]")
    private Output<String> functionName;

    /**
     * @return The name of the domain config.
     * 
     */
    public Output<String> functionName() {
        return this.functionName;
    }
    /**
     * The status of this resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of this resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DomainConfig(String name) {
        this(name, DomainConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DomainConfig(String name, DomainConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DomainConfig(String name, DomainConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:scdn/domainConfig:DomainConfig", name, args == null ? DomainConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DomainConfig(String name, Output<String> id, @Nullable DomainConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:scdn/domainConfig:DomainConfig", name, state, makeResourceOptions(options, id));
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
    public static DomainConfig get(String name, Output<String> id, @Nullable DomainConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DomainConfig(name, id, state, options);
    }
}
