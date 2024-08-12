// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oss.BucketAccessMonitorArgs;
import com.pulumi.alicloud.oss.inputs.BucketAccessMonitorState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * OSS Bucket Access Monitor can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:oss/bucketAccessMonitor:BucketAccessMonitor example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oss/bucketAccessMonitor:BucketAccessMonitor")
public class BucketAccessMonitor extends com.pulumi.resources.CustomResource {
    /**
     * The name of the bucket.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Specifies whether to enable access tracking for the bucket. Valid values: Enabled: enables access tracking. Disabled: disables access tracking.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BucketAccessMonitor(java.lang.String name) {
        this(name, BucketAccessMonitorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BucketAccessMonitor(java.lang.String name, BucketAccessMonitorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BucketAccessMonitor(java.lang.String name, BucketAccessMonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketAccessMonitor:BucketAccessMonitor", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private BucketAccessMonitor(java.lang.String name, Output<java.lang.String> id, @Nullable BucketAccessMonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oss/bucketAccessMonitor:BucketAccessMonitor", name, state, makeResourceOptions(options, id), false);
    }

    private static BucketAccessMonitorArgs makeArgs(BucketAccessMonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BucketAccessMonitorArgs.Empty : args;
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
    public static BucketAccessMonitor get(java.lang.String name, Output<java.lang.String> id, @Nullable BucketAccessMonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BucketAccessMonitor(name, id, state, options);
    }
}
