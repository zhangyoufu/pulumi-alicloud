// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sddp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.sddp.ConfigArgs;
import com.pulumi.alicloud.sddp.inputs.ConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Data Security Center Config resource.
 * 
 * For information about Data Security Center Config and how to use it, see [What is Config](https://www.alibabacloud.com/help/en/data-security-center/latest/api-sddp-2019-01-03-createconfig).
 * 
 * &gt; **NOTE:** Available since v1.133.0.
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
 * import com.pulumi.alicloud.sddp.Config;
 * import com.pulumi.alicloud.sddp.ConfigArgs;
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
 *         var default_ = new Config(&#34;default&#34;, ConfigArgs.builder()        
 *             .code(&#34;access_failed_cnt&#34;)
 *             .value(10)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Data Security Center Config can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:sddp/config:Config example &lt;code&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:sddp/config:Config")
public class Config extends com.pulumi.resources.CustomResource {
    /**
     * Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
     * 
     */
    @Export(name="code", type=String.class, parameters={})
    private Output</* @Nullable */ String> code;

    /**
     * @return Abnormal Alarm General Configuration Module by Using the Encoding. Valid values: `access_failed_cnt`, `access_permission_exprie_max_days`, `log_datasize_avg_days`.
     * 
     */
    public Output<Optional<String>> code() {
        return Codegen.optional(this.code);
    }
    /**
     * Abnormal Alarm General Description of the Configuration Item.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return Abnormal Alarm General Description of the Configuration Item.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The language of the request and response. Valid values: `zh`,`en`.
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language of the request and response. Valid values: `zh`,`en`.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
     * 
     */
    @Export(name="value", type=String.class, parameters={})
    private Output</* @Nullable */ String> value;

    /**
     * @return The Specified Exception Alarm Generic by Using the Value. Code Different Values for This Parameter the Specific Meaning of Different:
     * 
     */
    public Output<Optional<String>> value() {
        return Codegen.optional(this.value);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Config(String name) {
        this(name, ConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Config(String name, @Nullable ConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Config(String name, @Nullable ConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sddp/config:Config", name, args == null ? ConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Config(String name, Output<String> id, @Nullable ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:sddp/config:Config", name, state, makeResourceOptions(options, id));
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
    public static Config get(String name, Output<String> id, @Nullable ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Config(name, id, state, options);
    }
}
