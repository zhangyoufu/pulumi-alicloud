// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oss.BucketRefererArgs;
import com.pulumi.alicloud.oss.inputs.BucketRefererState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a OSS Bucket Referer resource. Bucket Referer configuration (Hotlink protection).
 * 
 * For information about OSS Bucket Referer and how to use it, see [What is Bucket Referer](https://www.alibabacloud.com/help/en/oss/user-guide/hotlink-protection).
 * 
 * &gt; **NOTE:** Available since v1.220.0.
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
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketReferer;
 * import com.pulumi.alicloud.oss.BucketRefererArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var name = config.get("name").orElse("terraform-example");
 *         var default_ = new Integer("default", IntegerArgs.builder()
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var createBucket = new Bucket("createBucket", BucketArgs.builder()
 *             .storageClass("Standard")
 *             .bucket(String.format("%s-%s", name,default_.result()))
 *             .build());
 * 
 *         var defaultBucketReferer = new BucketReferer("defaultBucketReferer", BucketRefererArgs.builder()
 *             .allowEmptyReferer("true")
 *             .refererBlacklists("*.forbidden.com")
 *             .bucket(createBucket.bucket())
 *             .truncatePath("false")
 *             .allowTruncateQueryString("true")
 *             .refererLists(            
 *                 "*.aliyun.com",
 *                 "*.example.com")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(createBucket)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OSS Bucket Referer can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:oss/bucketReferer:BucketReferer example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oss/bucketReferer:BucketReferer")
public class BucketReferer extends com.pulumi.resources.CustomResource {
    /**
     * Whether to allow empty Referer request headers.
     * 
     */
    @Export(name="allowEmptyReferer", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowEmptyReferer;

    /**
     * @return Whether to allow empty Referer request headers.
     * 
     */
    public Output<Boolean> allowEmptyReferer() {
        return this.allowEmptyReferer;
    }
    /**
     * Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
     * 
     */
    @Export(name="allowTruncateQueryString", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowTruncateQueryString;

    /**
     * @return Specifies whether to truncate the query string in the URL when the Referer is matched. Valid values: true, false.
     * 
     */
    public Output<Boolean> allowTruncateQueryString() {
        return this.allowTruncateQueryString;
    }
    /**
     * Name of the Bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return Name of the Bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The container that holds the Referer blacklist.
     * 
     */
    @Export(name="refererBlacklists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> refererBlacklists;

    /**
     * @return The container that holds the Referer blacklist.
     * 
     */
    public Output<Optional<List<String>>> refererBlacklists() {
        return Codegen.optional(this.refererBlacklists);
    }
    /**
     * The container that holds the Referer whitelist.
     * 
     */
    @Export(name="refererLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> refererLists;

    /**
     * @return The container that holds the Referer whitelist.
     * 
     */
    public Output<Optional<List<String>>> refererLists() {
        return Codegen.optional(this.refererLists);
    }
    /**
     * Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
     * 
     */
    @Export(name="truncatePath", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> truncatePath;

    /**
     * @return Specifies whether to truncate the path and parts that follow the path in the URL when the Referer is matched. Valid values: true, false. If TruncatePath is set to true, the value of AllowTruncateQueryString must be also true because the query string follows the path component. When the path is truncated, the query string is also truncated.
     * 
     */
    public Output<Optional<Boolean>> truncatePath() {
        return Codegen.optional(this.truncatePath);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketReferer(java.lang.String name) {
        this(name, BucketRefererArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketReferer(java.lang.String name, BucketRefererArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketReferer(java.lang.String name, BucketRefererArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketReferer:BucketReferer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BucketReferer(java.lang.String name, Output<java.lang.String> id, @Nullable BucketRefererState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketReferer:BucketReferer", name, state, makeResourceOptions(options, id), false);
    }

    private static BucketRefererArgs makeArgs(BucketRefererArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BucketRefererArgs.Empty : args;
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
    public static BucketReferer get(java.lang.String name, Output<java.lang.String> id, @Nullable BucketRefererState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketReferer(name, id, state, options);
    }
}
