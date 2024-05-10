// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.actiontrail.TrailDeprecatedArgs;
import com.pulumi.alicloud.actiontrail.inputs.TrailDeprecatedState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **DEPRECATED:**  This resource has been renamed to alicloud.actiontrail.Trail from version 1.95.0.
 * 
 * Provides a new resource to manage [Action Trail](https://www.alibabacloud.com/help/doc-detail/28804.htm).
 * 
 * &gt; **NOTE:** Available in 1.35.0+
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
 * import com.pulumi.alicloud.actiontrail.TrailDeprecated;
 * import com.pulumi.alicloud.actiontrail.TrailDeprecatedArgs;
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
 *         // Create a new action trail.
 *         var foo = new TrailDeprecated("foo", TrailDeprecatedArgs.builder()        
 *             .name("action-trail")
 *             .eventRw("Write-test")
 *             .ossBucketName(bucket.id())
 *             .roleName(attach.roleName())
 *             .ossKeyPrefix("at-product-account-audit-B")
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
 * Action trail can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:actiontrail/trailDeprecated:TrailDeprecated foo abc12345678
 * ```
 * 
 * @deprecated
 * Resource renamed to `Trail`
 * 
 */
@Deprecated /* Resource renamed to `Trail` */
@ResourceType(type="alicloud:actiontrail/trailDeprecated:TrailDeprecated")
public class TrailDeprecated extends com.pulumi.resources.CustomResource {
    /**
     * Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     * 
     */
    @Export(name="eventRw", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventRw;

    /**
     * @return Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.
     * 
     */
    public Output<Optional<String>> eventRw() {
        return Codegen.optional(this.eventRw);
    }
    @Export(name="isOrganizationTrail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isOrganizationTrail;

    public Output<Optional<Boolean>> isOrganizationTrail() {
        return Codegen.optional(this.isOrganizationTrail);
    }
    /**
     * @deprecated
     * Field &#39;mns_topic_arn&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'mns_topic_arn' has been deprecated from version 1.118.0 */
    @Export(name="mnsTopicArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mnsTopicArn;

    public Output<Optional<String>> mnsTopicArn() {
        return Codegen.optional(this.mnsTopicArn);
    }
    /**
     * The name of the trail to be created, which must be unique for an account.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from version 1.95.0. Use &#39;trail_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from version 1.95.0. Use 'trail_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the trail to be created, which must be unique for an account.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    @Export(name="ossBucketName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossBucketName;

    /**
     * @return The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.
     * 
     */
    public Output<Optional<String>> ossBucketName() {
        return Codegen.optional(this.ossBucketName);
    }
    /**
     * The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    @Export(name="ossKeyPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossKeyPrefix;

    /**
     * @return The prefix of the specified OSS bucket name. This parameter can be left empty.
     * 
     */
    public Output<Optional<String>> ossKeyPrefix() {
        return Codegen.optional(this.ossKeyPrefix);
    }
    @Export(name="ossWriteRoleArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ossWriteRoleArn;

    public Output<Optional<String>> ossWriteRoleArn() {
        return Codegen.optional(this.ossWriteRoleArn);
    }
    /**
     * The RAM role in ActionTrail permitted by the user.
     * 
     * @deprecated
     * Field &#39;role_name&#39; has been deprecated from version 1.118.0
     * 
     */
    @Deprecated /* Field 'role_name' has been deprecated from version 1.118.0 */
    @Export(name="roleName", refs={String.class}, tree="[0]")
    private Output<String> roleName;

    /**
     * @return The RAM role in ActionTrail permitted by the user.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * The unique ARN of the Log Service project.
     * 
     */
    @Export(name="slsProjectArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> slsProjectArn;

    /**
     * @return The unique ARN of the Log Service project.
     * 
     */
    public Output<Optional<String>> slsProjectArn() {
        return Codegen.optional(this.slsProjectArn);
    }
    /**
     * The unique ARN of the Log Service role.
     * 
     * &gt; **NOTE:** `sls_project_arn` and `sls_write_role_arn` should be set or not set at the same time when actiontrail delivers logs.
     * 
     */
    @Export(name="slsWriteRoleArn", refs={String.class}, tree="[0]")
    private Output<String> slsWriteRoleArn;

    /**
     * @return The unique ARN of the Log Service role.
     * 
     * &gt; **NOTE:** `sls_project_arn` and `sls_write_role_arn` should be set or not set at the same time when actiontrail delivers logs.
     * 
     */
    public Output<String> slsWriteRoleArn() {
        return this.slsWriteRoleArn;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    @Export(name="trailName", refs={String.class}, tree="[0]")
    private Output<String> trailName;

    public Output<String> trailName() {
        return this.trailName;
    }
    @Export(name="trailRegion", refs={String.class}, tree="[0]")
    private Output<String> trailRegion;

    public Output<String> trailRegion() {
        return this.trailRegion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TrailDeprecated(String name) {
        this(name, TrailDeprecatedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TrailDeprecated(String name, @Nullable TrailDeprecatedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TrailDeprecated(String name, @Nullable TrailDeprecatedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, args == null ? TrailDeprecatedArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TrailDeprecated(String name, Output<String> id, @Nullable TrailDeprecatedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/trailDeprecated:TrailDeprecated", name, state, makeResourceOptions(options, id));
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
    public static TrailDeprecated get(String name, Output<String> id, @Nullable TrailDeprecatedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TrailDeprecated(name, id, state, options);
    }
}
