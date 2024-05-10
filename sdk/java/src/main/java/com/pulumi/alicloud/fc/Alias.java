// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fc.AliasArgs;
import com.pulumi.alicloud.fc.inputs.AliasState;
import com.pulumi.alicloud.fc.outputs.AliasRoutingConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Function Compute service alias. Creates an alias that points to the specified Function Compute service version.
 *  For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/en/fc/developer-reference/api-createalias).
 * 
 * &gt; **NOTE:** Available since v1.104.0.
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
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.fc.Service;
 * import com.pulumi.alicloud.fc.ServiceArgs;
 * import com.pulumi.alicloud.fc.Alias;
 * import com.pulumi.alicloud.fc.AliasArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var defaultService = new Service("defaultService", ServiceArgs.builder()        
 *             .name(String.format("example-value-%s", default_.result()))
 *             .description("example-value")
 *             .publish("true")
 *             .build());
 * 
 *         var example = new Alias("example", AliasArgs.builder()        
 *             .aliasName("example-value")
 *             .description("example-value")
 *             .serviceName(defaultService.name())
 *             .serviceVersion("1")
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
 * Function Compute alias can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:fc/alias:Alias example my_alias_id
 * ```
 * 
 */
@ResourceType(type="alicloud:fc/alias:Alias")
public class Alias extends com.pulumi.resources.CustomResource {
    /**
     * Name for the alias you are creating.
     * 
     */
    @Export(name="aliasName", refs={String.class}, tree="[0]")
    private Output<String> aliasName;

    /**
     * @return Name for the alias you are creating.
     * 
     */
    public Output<String> aliasName() {
        return this.aliasName;
    }
    /**
     * Description of the alias.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the alias.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Function Compute alias&#39; route configuration settings. See `routing_config` below.
     * 
     */
    @Export(name="routingConfig", refs={AliasRoutingConfig.class}, tree="[0]")
    private Output</* @Nullable */ AliasRoutingConfig> routingConfig;

    /**
     * @return The Function Compute alias&#39; route configuration settings. See `routing_config` below.
     * 
     */
    public Output<Optional<AliasRoutingConfig>> routingConfig() {
        return Codegen.optional(this.routingConfig);
    }
    /**
     * The Function Compute service name.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The Function Compute service name.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
     * 
     */
    @Export(name="serviceVersion", refs={String.class}, tree="[0]")
    private Output<String> serviceVersion;

    /**
     * @return The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
     * 
     */
    public Output<String> serviceVersion() {
        return this.serviceVersion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Alias(String name) {
        this(name, AliasArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Alias(String name, AliasArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Alias(String name, AliasArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/alias:Alias", name, args == null ? AliasArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Alias(String name, Output<String> id, @Nullable AliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:fc/alias:Alias", name, state, makeResourceOptions(options, id));
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
    public static Alias get(String name, Output<String> id, @Nullable AliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Alias(name, id, state, options);
    }
}
