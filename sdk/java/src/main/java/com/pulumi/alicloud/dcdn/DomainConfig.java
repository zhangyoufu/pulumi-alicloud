// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dcdn.DomainConfigArgs;
import com.pulumi.alicloud.dcdn.inputs.DomainConfigState;
import com.pulumi.alicloud.dcdn.outputs.DomainConfigFunctionArg;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a DCDN Accelerated Domain resource.
 * 
 * For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/en/doc-detail/130632.htm).
 * 
 * &gt; **NOTE:** Available since v1.131.0.
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
 * import com.pulumi.alicloud.dcdn.Domain;
 * import com.pulumi.alicloud.dcdn.DomainArgs;
 * import com.pulumi.alicloud.dcdn.inputs.DomainSourceArgs;
 * import com.pulumi.alicloud.dcdn.DomainConfig;
 * import com.pulumi.alicloud.dcdn.DomainConfigArgs;
 * import com.pulumi.alicloud.dcdn.inputs.DomainConfigFunctionArgArgs;
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
 *         final var domainName = config.get("domainName").orElse("alibaba-example.com");
 *         var example = new Domain("example", DomainArgs.builder()
 *             .domainName(domainName)
 *             .scope("overseas")
 *             .status("online")
 *             .sources(DomainSourceArgs.builder()
 *                 .content("1.1.1.1")
 *                 .type("ipaddr")
 *                 .priority(20)
 *                 .port(80)
 *                 .weight(10)
 *                 .build())
 *             .build());
 * 
 *         var ipAllowListSet = new DomainConfig("ipAllowListSet", DomainConfigArgs.builder()
 *             .domainName(example.domainName())
 *             .functionName("ip_allow_list_set")
 *             .functionArgs(DomainConfigFunctionArgArgs.builder()
 *                 .argName("ip_list")
 *                 .argValue("192.168.0.1")
 *                 .build())
 *             .build());
 * 
 *         var refererWhiteListSet = new DomainConfig("refererWhiteListSet", DomainConfigArgs.builder()
 *             .domainName(example.domainName())
 *             .functionName("referer_white_list_set")
 *             .functionArgs(DomainConfigFunctionArgArgs.builder()
 *                 .argName("refer_domain_allow_list")
 *                 .argValue("110.110.110.110")
 *                 .build())
 *             .build());
 * 
 *         var filetypeBasedTtlSet = new DomainConfig("filetypeBasedTtlSet", DomainConfigArgs.builder()
 *             .domainName(example.domainName())
 *             .functionName("filetype_based_ttl_set")
 *             .functionArgs(            
 *                 DomainConfigFunctionArgArgs.builder()
 *                     .argName("ttl")
 *                     .argValue("300")
 *                     .build(),
 *                 DomainConfigFunctionArgArgs.builder()
 *                     .argName("file_type")
 *                     .argValue("jpg")
 *                     .build(),
 *                 DomainConfigFunctionArgArgs.builder()
 *                     .argName("weight")
 *                     .argValue("1")
 *                     .build())
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
 * DCDN domain config can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:dcdn/domainConfig:DomainConfig example &lt;domain_name&gt;:&lt;function_name&gt;:&lt;config_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dcdn/domainConfig:DomainConfig")
public class DomainConfig extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the configuration.
     * 
     */
    @Export(name="configId", refs={String.class}, tree="[0]")
    private Output<String> configId;

    /**
     * @return The ID of the configuration.
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
     * The args of the domain config. See `function_args` below.
     * 
     */
    @Export(name="functionArgs", refs={List.class,DomainConfigFunctionArg.class}, tree="[0,1]")
    private Output<List<DomainConfigFunctionArg>> functionArgs;

    /**
     * @return The args of the domain config. See `function_args` below.
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
     * By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
     * 
     */
    @Export(name="parentId", refs={String.class}, tree="[0]")
    private Output<String> parentId;

    /**
     * @return By configuring the function condition (rule engine) in the domain name configuration function parameters, Rule conditions can be created (Rule conditions can match and filter user requests by identifying various parameters carried in user requests). After each rule condition is created, a corresponding ConfigId will be generated, and the ConfigId can be referenced by other functions as a ParentId parameter, in this way, the rule conditions can be combined with the functional configuration to form a more flexible configuration.
     * 
     */
    public Output<String> parentId() {
        return this.parentId;
    }
    /**
     * The status of the Config.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Config.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DomainConfig(java.lang.String name) {
        this(name, DomainConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DomainConfig(java.lang.String name, DomainConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DomainConfig(java.lang.String name, DomainConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/domainConfig:DomainConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DomainConfig(java.lang.String name, Output<java.lang.String> id, @Nullable DomainConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dcdn/domainConfig:DomainConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static DomainConfigArgs makeArgs(DomainConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DomainConfigArgs.Empty : args;
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
    public static DomainConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable DomainConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DomainConfig(name, id, state, options);
    }
}
