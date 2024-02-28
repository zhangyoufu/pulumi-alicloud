// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cdn.DomainConfigArgs;
import com.pulumi.alicloud.cdn.inputs.DomainConfigState;
import com.pulumi.alicloud.cdn.outputs.DomainConfigFunctionArg;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a CDN Accelerated Domain resource.
 * 
 * For information about domain config and how to use it, see [Batch set config](https://www.alibabacloud.com/help/zh/doc-detail/90915.htm)
 * 
 * &gt; **NOTE:** Available since v1.34.0+.
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
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.cdn.DomainNew;
 * import com.pulumi.alicloud.cdn.DomainNewArgs;
 * import com.pulumi.alicloud.cdn.inputs.DomainNewSourceArgs;
 * import com.pulumi.alicloud.cdn.DomainConfig;
 * import com.pulumi.alicloud.cdn.DomainConfigArgs;
 * import com.pulumi.alicloud.cdn.inputs.DomainConfigFunctionArgArgs;
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
 *         var default_ = new RandomInteger(&#34;default&#34;, RandomIntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var domain = new DomainNew(&#34;domain&#34;, DomainNewArgs.builder()        
 *             .domainName(default_.result().applyValue(result -&gt; String.format(&#34;mycdndomain-%s.alicloud-provider.cn&#34;, result)))
 *             .cdnType(&#34;web&#34;)
 *             .scope(&#34;overseas&#34;)
 *             .sources(DomainNewSourceArgs.builder()
 *                 .content(&#34;1.1.1.1&#34;)
 *                 .type(&#34;ipaddr&#34;)
 *                 .priority(&#34;20&#34;)
 *                 .port(80)
 *                 .weight(&#34;15&#34;)
 *                 .build())
 *             .build());
 * 
 *         var config = new DomainConfig(&#34;config&#34;, DomainConfigArgs.builder()        
 *             .domainName(domain.domainName())
 *             .functionName(&#34;ip_allow_list_set&#34;)
 *             .functionArgs(DomainConfigFunctionArgArgs.builder()
 *                 .argName(&#34;ip_list&#34;)
 *                 .argValue(&#34;110.110.110.110&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CDN domain config can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cdn/domainConfig:DomainConfig example &lt;domain_name&gt;:&lt;function_name&gt;:&lt;config_id&gt;
 * ```
 * 
 * ```sh
 * $ pulumi import alicloud:cdn/domainConfig:DomainConfig example &lt;domain_name&gt;:&lt;function_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cdn/domainConfig:DomainConfig")
public class DomainConfig extends com.pulumi.resources.CustomResource {
    /**
     * (Available in 1.132.0+) The ID of the domain config function.
     * 
     */
    @Export(name="configId", refs={String.class}, tree="[0]")
    private Output<String> configId;

    /**
     * @return (Available in 1.132.0+) The ID of the domain config function.
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
     * (Available in 1.132.0+) The Status of the function. Valid values: `success`, `testing`, `failed`, and `configuring`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return (Available in 1.132.0+) The Status of the function. Valid values: `success`, `testing`, `failed`, and `configuring`.
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
        super("alicloud:cdn/domainConfig:DomainConfig", name, args == null ? DomainConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DomainConfig(String name, Output<String> id, @Nullable DomainConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cdn/domainConfig:DomainConfig", name, state, makeResourceOptions(options, id));
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
