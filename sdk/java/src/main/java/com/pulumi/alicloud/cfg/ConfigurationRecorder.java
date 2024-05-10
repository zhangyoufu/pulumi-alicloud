// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cfg.ConfigurationRecorderArgs;
import com.pulumi.alicloud.cfg.inputs.ConfigurationRecorderState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Alicloud Config Configuration Recorder resource. Cloud Config is a specialized service for evaluating resources. Cloud Config tracks configuration changes of your resources and evaluates configuration compliance. Cloud Config can help you evaluate numerous resources and maintain the continuous compliance of your cloud infrastructure.
 * For information about Alicloud Config Configuration Recorder and how to use it, see [What is Configuration Recorder.](https://www.alibabacloud.com/help/en/cloud-config/latest/startconfigurationrecorder)
 * 
 * &gt; **NOTE:** Available since v1.99.0.
 * 
 * &gt; **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-southeast-1`.
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
 * import com.pulumi.alicloud.cfg.ConfigurationRecorder;
 * import com.pulumi.alicloud.cfg.ConfigurationRecorderArgs;
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
 *         var example = new ConfigurationRecorder("example", ConfigurationRecorderArgs.builder()        
 *             .resourceTypes(            
 *                 "ACS::ECS::Instance",
 *                 "ACS::ECS::Disk")
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
 * Alicloud Config Configuration Recorder can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cfg/configurationRecorder:ConfigurationRecorder example 122378463********
 * ```
 * 
 */
@ResourceType(type="alicloud:cfg/configurationRecorder:ConfigurationRecorder")
public class ConfigurationRecorder extends com.pulumi.resources.CustomResource {
    /**
     * Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
     * 
     */
    @Export(name="enterpriseEdition", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enterpriseEdition;

    /**
     * @return Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
     * 
     */
    public Output<Boolean> enterpriseEdition() {
        return this.enterpriseEdition;
    }
    /**
     * Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
     * 
     */
    @Export(name="organizationEnableStatus", refs={String.class}, tree="[0]")
    private Output<String> organizationEnableStatus;

    /**
     * @return Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
     * 
     */
    public Output<String> organizationEnableStatus() {
        return this.organizationEnableStatus;
    }
    /**
     * The ID of the Enterprise management account.
     * 
     */
    @Export(name="organizationMasterId", refs={Integer.class}, tree="[0]")
    private Output<Integer> organizationMasterId;

    /**
     * @return The ID of the Enterprise management account.
     * 
     */
    public Output<Integer> organizationMasterId() {
        return this.organizationMasterId;
    }
    /**
     * A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
     * * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
     * * If you use an enterprise account, the `resource_types` does not support updating.
     * 
     */
    @Export(name="resourceTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> resourceTypes;

    /**
     * @return A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
     * * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
     * * If you use an enterprise account, the `resource_types` does not support updating.
     * 
     */
    public Output<List<String>> resourceTypes() {
        return this.resourceTypes;
    }
    /**
     * Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConfigurationRecorder(String name) {
        this(name, ConfigurationRecorderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConfigurationRecorder(String name, @Nullable ConfigurationRecorderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConfigurationRecorder(String name, @Nullable ConfigurationRecorderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/configurationRecorder:ConfigurationRecorder", name, args == null ? ConfigurationRecorderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConfigurationRecorder(String name, Output<String> id, @Nullable ConfigurationRecorderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cfg/configurationRecorder:ConfigurationRecorder", name, state, makeResourceOptions(options, id));
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
    public static ConfigurationRecorder get(String name, Output<String> id, @Nullable ConfigurationRecorderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConfigurationRecorder(name, id, state, options);
    }
}
