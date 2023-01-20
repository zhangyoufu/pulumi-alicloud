// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.apigateway.LogConfigArgs;
import com.pulumi.alicloud.apigateway.inputs.LogConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Api Gateway Log Config resource.
 * 
 * For information about Api Gateway Log Config and how to use it, see [What is Log Config](https://help.aliyun.com/document_detail/400392.html).
 * 
 * &gt; **NOTE:** Available in v1.185.0+.
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
 * import com.pulumi.alicloud.apigateway.LogConfig;
 * import com.pulumi.alicloud.apigateway.LogConfigArgs;
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
 *         var default_ = new LogConfig(&#34;default&#34;, LogConfigArgs.builder()        
 *             .logType(&#34;PROVIDER&#34;)
 *             .slsLogStore(&#34;example_value&#34;)
 *             .slsProject(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Api Gateway Log Config can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:apigateway/logConfig:LogConfig example &lt;log_type&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:apigateway/logConfig:LogConfig")
public class LogConfig extends com.pulumi.resources.CustomResource {
    /**
     * The type the of log. Valid values: `PROVIDER`.
     * 
     */
    @Export(name="logType", type=String.class, parameters={})
    private Output<String> logType;

    /**
     * @return The type the of log. Valid values: `PROVIDER`.
     * 
     */
    public Output<String> logType() {
        return this.logType;
    }
    /**
     * The name of the Log Store.
     * 
     */
    @Export(name="slsLogStore", type=String.class, parameters={})
    private Output<String> slsLogStore;

    /**
     * @return The name of the Log Store.
     * 
     */
    public Output<String> slsLogStore() {
        return this.slsLogStore;
    }
    /**
     * The name of the Project.
     * 
     */
    @Export(name="slsProject", type=String.class, parameters={})
    private Output<String> slsProject;

    /**
     * @return The name of the Project.
     * 
     */
    public Output<String> slsProject() {
        return this.slsProject;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogConfig(String name) {
        this(name, LogConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogConfig(String name, LogConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogConfig(String name, LogConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/logConfig:LogConfig", name, args == null ? LogConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LogConfig(String name, Output<String> id, @Nullable LogConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:apigateway/logConfig:LogConfig", name, state, makeResourceOptions(options, id));
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
    public static LogConfig get(String name, Output<String> id, @Nullable LogConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogConfig(name, id, state, options);
    }
}
