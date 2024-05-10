// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oss.BucketCorsArgs;
import com.pulumi.alicloud.oss.inputs.BucketCorsState;
import com.pulumi.alicloud.oss.outputs.BucketCorsCorsRule;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a OSS Bucket Cors resource. Cross-Origin Resource Sharing (CORS) allows web applications to access resources in other regions.
 * 
 * For information about OSS Bucket Cors and how to use it, see [What is Bucket Cors](https://www.alibabacloud.com/help/en/).
 * 
 * &gt; **NOTE:** Available since v1.223.0.
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
 * import com.pulumi.alicloud.oss.Bucket;
 * import com.pulumi.alicloud.oss.BucketArgs;
 * import com.pulumi.alicloud.oss.BucketCors;
 * import com.pulumi.alicloud.oss.BucketCorsArgs;
 * import com.pulumi.alicloud.oss.inputs.BucketCorsCorsRuleArgs;
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
 *         var createBucket = new Bucket("createBucket", BucketArgs.builder()        
 *             .storageClass("Standard")
 *             .bucket(name)
 *             .build());
 * 
 *         var default_ = new BucketCors("default", BucketCorsArgs.builder()        
 *             .bucket(createBucket.bucket())
 *             .responseVary(true)
 *             .corsRules(BucketCorsCorsRuleArgs.builder()
 *                 .allowedMethods("GET")
 *                 .allowedOrigins("*")
 *                 .allowedHeaders(                
 *                     "x-oss-test",
 *                     "x-oss-abc")
 *                 .exposeHeaders("x-oss-request-id")
 *                 .maxAgeSeconds("1000")
 *                 .build())
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
 * OSS Bucket Cors can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:oss/bucketCors:BucketCors example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oss/bucketCors:BucketCors")
public class BucketCors extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the Bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
     * 
     */
    @Export(name="corsRules", refs={List.class,BucketCorsCorsRule.class}, tree="[0,1]")
    private Output<List<BucketCorsCorsRule>> corsRules;

    /**
     * @return The Cross-Origin Resource Sharing (CORS) configuration of the Bucket. See `cors_rule` below.
     * 
     */
    public Output<List<BucketCorsCorsRule>> corsRules() {
        return this.corsRules;
    }
    /**
     * Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
     * 
     */
    @Export(name="responseVary", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> responseVary;

    /**
     * @return Specifies whether to return the Vary: Origin header. Valid values: true: returns the Vary: Origin header, regardless of whether the request is a cross-origin request or whether the cross-origin request succeeds. false: does not return the Vary: Origin header. This element is valid only when at least one CORS rule is configured.
     * 
     */
    public Output<Boolean> responseVary() {
        return this.responseVary;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketCors(String name) {
        this(name, BucketCorsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketCors(String name, BucketCorsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketCors(String name, BucketCorsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketCors:BucketCors", name, args == null ? BucketCorsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BucketCors(String name, Output<String> id, @Nullable BucketCorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketCors:BucketCors", name, state, makeResourceOptions(options, id));
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
    public static BucketCors get(String name, Output<String> id, @Nullable BucketCorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketCors(name, id, state, options);
    }
}
