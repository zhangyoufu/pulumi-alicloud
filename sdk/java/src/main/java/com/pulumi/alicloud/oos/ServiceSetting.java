// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oos.ServiceSettingArgs;
import com.pulumi.alicloud.oos.inputs.ServiceSettingState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a OOS Service Setting resource.
 * 
 * For information about OOS Service Setting and how to use it, see [What is Service Setting](https://www.alibabacloud.com/help/en/doc-detail/268700.html).
 * 
 * &gt; **NOTE:** Available in v1.147.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.oos.ServiceSetting;
 * import com.pulumi.alicloud.oos.ServiceSettingArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-testaccoossetting&#34;);
 *         var defaultBucket = new Bucket(&#34;defaultBucket&#34;, BucketArgs.builder()        
 *             .bucket(name)
 *             .acl(&#34;public-read-write&#34;)
 *             .build());
 * 
 *         var defaultProject = new Project(&#34;defaultProject&#34;);
 * 
 *         var defaultServiceSetting = new ServiceSetting(&#34;defaultServiceSetting&#34;, ServiceSettingArgs.builder()        
 *             .deliveryOssEnabled(true)
 *             .deliveryOssKeyPrefix(&#34;path1/&#34;)
 *             .deliveryOssBucketName(defaultBucket.bucket())
 *             .deliverySlsEnabled(true)
 *             .deliverySlsProjectName(defaultProject.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OOS Service Setting can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:oos/serviceSetting:ServiceSetting example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oos/serviceSetting:ServiceSetting")
public class ServiceSetting extends com.pulumi.resources.CustomResource {
    /**
     * The name of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
     * 
     */
    @Export(name="deliveryOssBucketName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryOssBucketName;

    /**
     * @return The name of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
     * 
     */
    public Output<Optional<String>> deliveryOssBucketName() {
        return Codegen.optional(this.deliveryOssBucketName);
    }
    /**
     * Is the recording function for the OSS delivery template enabled.
     * 
     */
    @Export(name="deliveryOssEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deliveryOssEnabled;

    /**
     * @return Is the recording function for the OSS delivery template enabled.
     * 
     */
    public Output<Optional<Boolean>> deliveryOssEnabled() {
        return Codegen.optional(this.deliveryOssEnabled);
    }
    /**
     * The Directory of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
     * 
     */
    @Export(name="deliveryOssKeyPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliveryOssKeyPrefix;

    /**
     * @return The Directory of the OSS bucket. **NOTE:** When the `delivery_oss_enabled` is `true`, The `delivery_oss_bucket_name` is valid.
     * 
     */
    public Output<Optional<String>> deliveryOssKeyPrefix() {
        return Codegen.optional(this.deliveryOssKeyPrefix);
    }
    /**
     * Is the execution record function to SLS delivery Template turned on.
     * 
     */
    @Export(name="deliverySlsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deliverySlsEnabled;

    /**
     * @return Is the execution record function to SLS delivery Template turned on.
     * 
     */
    public Output<Optional<Boolean>> deliverySlsEnabled() {
        return Codegen.optional(this.deliverySlsEnabled);
    }
    /**
     * The name of SLS  Project. **NOTE:** When the `delivery_sls_enabled` is `true`, The `delivery_sls_project_name` is valid.
     * 
     */
    @Export(name="deliverySlsProjectName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deliverySlsProjectName;

    /**
     * @return The name of SLS  Project. **NOTE:** When the `delivery_sls_enabled` is `true`, The `delivery_sls_project_name` is valid.
     * 
     */
    public Output<Optional<String>> deliverySlsProjectName() {
        return Codegen.optional(this.deliverySlsProjectName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceSetting(String name) {
        this(name, ServiceSettingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceSetting(String name, @Nullable ServiceSettingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceSetting(String name, @Nullable ServiceSettingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/serviceSetting:ServiceSetting", name, args == null ? ServiceSettingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceSetting(String name, Output<String> id, @Nullable ServiceSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/serviceSetting:ServiceSetting", name, state, makeResourceOptions(options, id));
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
    public static ServiceSetting get(String name, Output<String> id, @Nullable ServiceSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceSetting(name, id, state, options);
    }
}
