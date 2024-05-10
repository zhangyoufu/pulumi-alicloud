// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.EnvCustomJobArgs;
import com.pulumi.alicloud.arms.inputs.EnvCustomJobState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ARMS Env Custom Job resource. Custom jobs in the arms environment.
 * 
 * For information about ARMS Env Custom Job and how to use it, see [What is Env Custom Job](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-createenvcustomjob).
 * 
 * &gt; **NOTE:** Available since v1.212.0.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.arms.Environment;
 * import com.pulumi.alicloud.arms.EnvironmentArgs;
 * import com.pulumi.alicloud.arms.EnvCustomJob;
 * import com.pulumi.alicloud.arms.EnvCustomJobArgs;
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
 *         var default_ = new Integer("default", IntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         final var name = config.get("name").orElse("terraform-example");
 *         var vpc = new Network("vpc", NetworkArgs.builder()        
 *             .description(name)
 *             .cidrBlock("172.16.0.0/12")
 *             .vpcName(name)
 *             .build());
 * 
 *         var env_ecs = new Environment("env-ecs", EnvironmentArgs.builder()        
 *             .environmentType("ECS")
 *             .environmentName(String.format("terraform-example-%s", default_.result()))
 *             .bindResourceId(vpc.id())
 *             .environmentSubType("ECS")
 *             .build());
 * 
 *         var defaultEnvCustomJob = new EnvCustomJob("defaultEnvCustomJob", EnvCustomJobArgs.builder()        
 *             .status("run")
 *             .environmentId(env_ecs.id())
 *             .envCustomJobName(name)
 *             .configYaml("""
 * scrape_configs:
 * - job_name: job-demo1
 *   honor_timestamps: false
 *   honor_labels: false
 *   scrape_interval: 30s
 *   scheme: http
 *   metrics_path: /metric
 *   static_configs:
 *   - targets:
 *     - 127.0.0.1:9090
 *             """)
 *             .aliyunLang("en")
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
 * ARMS Env Custom Job can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:arms/envCustomJob:EnvCustomJob example &lt;environment_id&gt;:&lt;env_custom_job_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/envCustomJob:EnvCustomJob")
public class EnvCustomJob extends com.pulumi.resources.CustomResource {
    /**
     * The locale. The default is Chinese zh | en.
     * 
     */
    @Export(name="aliyunLang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> aliyunLang;

    /**
     * @return The locale. The default is Chinese zh | en.
     * 
     */
    public Output<Optional<String>> aliyunLang() {
        return Codegen.optional(this.aliyunLang);
    }
    /**
     * Yaml configuration string.
     * 
     */
    @Export(name="configYaml", refs={String.class}, tree="[0]")
    private Output<String> configYaml;

    /**
     * @return Yaml configuration string.
     * 
     */
    public Output<String> configYaml() {
        return this.configYaml;
    }
    /**
     * Custom job name.
     * 
     */
    @Export(name="envCustomJobName", refs={String.class}, tree="[0]")
    private Output<String> envCustomJobName;

    /**
     * @return Custom job name.
     * 
     */
    public Output<String> envCustomJobName() {
        return this.envCustomJobName;
    }
    /**
     * Environment id.
     * 
     */
    @Export(name="environmentId", refs={String.class}, tree="[0]")
    private Output<String> environmentId;

    /**
     * @return Environment id.
     * 
     */
    public Output<String> environmentId() {
        return this.environmentId;
    }
    /**
     * Status: run, stop.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status: run, stop.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnvCustomJob(String name) {
        this(name, EnvCustomJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvCustomJob(String name, EnvCustomJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvCustomJob(String name, EnvCustomJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/envCustomJob:EnvCustomJob", name, args == null ? EnvCustomJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnvCustomJob(String name, Output<String> id, @Nullable EnvCustomJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/envCustomJob:EnvCustomJob", name, state, makeResourceOptions(options, id));
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
    public static EnvCustomJob get(String name, Output<String> id, @Nullable EnvCustomJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvCustomJob(name, id, state, options);
    }
}
